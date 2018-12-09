package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/alash3al/go-color"

	"github.com/rs/xid"
	"github.com/tidwall/redcon"
)

func main() {
	fmt.Println(color.CyanString(banner))
	fmt.Printf("⇨ uTree started on port %s \n", color.GreenString(*flagListenAddr))
	redcon.ListenAndServe(*flagListenAddr, handleProcessCommand, handleAcceptConn, nil)
}

func handleProcessCommand(conn redcon.Conn, cmd redcon.Command) {
	action := strings.ToLower(string(cmd.Args[0]))
	args := []string{}
	for _, v := range cmd.Args[1:] {
		args = append(args, string(v))
	}
	switch action {
	default:
		conn.WriteError("unknown command: " + action)

	// ping
	// returns PONG
	case "ping":
		conn.WriteString("PONG")

	// gen
	// generates a new unique id
	case "gen":
		id := xid.New().String()
		conn.WriteString(id)

	// append
	// adds a new node to the tree
	case "append":
		if len(args) < 1 {
			conn.WriteError("you must specify the parent")
			break
		}
		parent, child := args[0], xid.New().String()
		redisConn.SAdd("utree:tree:"+parent+":children", child).Val()
		redisConn.Set("utree:tree:"+child+":parent", parent, -1).Val()
		conn.WriteString(child)

	// flatten
	// returns the elements of the node [tree]
	// as an array of strings
	case "flatten":
		if len(args) < 1 {
			conn.WriteError("you must specify the parent")
			break
		}
		var nest func(string) []string
		parent := args[0]
		loaded := map[string]bool{}
		nest = func(parent string) []string {
			ret := []string{}
			children := redisConn.SMembers("utree:tree:" + parent + ":children").Val()
			for _, id := range children {
				if loaded[id] {
					continue
				}
				loaded[id] = true
				ret = append(ret, id)
				ret = append(ret, nest(id)...)
			}
			return ret
		}

		all := nest(parent)
		conn.WriteArray(len(all))
		for _, v := range all {
			conn.WriteBulkString(v)
		}

	// tree
	// returns the node elements
	// as a tree as json encoded string
	case "tree":
		if len(args) < 1 {
			conn.WriteError("you must specify the parent")
			break
		}
		var nest func(string) []Node
		parent := args[0]
		loaded := map[string]*Node{}
		nest = func(parent string) []Node {
			node := new(Node)
			node.Node = parent
			node.Tree = []Node{}
			children := redisConn.SMembers("utree:tree:" + parent + ":children").Val()
			for _, id := range children {
				if loaded[id] != nil {
					continue
				}
				loaded[id] = new(Node)
				loaded[id].Node = id
				loaded[id].Tree = nest(id)
				node.Tree = append(node.Tree, *(loaded[id]))
			}
			return node.Tree
		}

		d, _ := json.Marshal(nest(parent))
		conn.WriteString(string(d))

	// move a node to another parent
	// 1)- fetch the old parent
	// 2)- remove the node from the old parent children
	// 3)- add the node to the dst's children
	// 4)- set the parent of the node as dst
	case "mv":
		if len(args) < 2 {
			conn.WriteError("you must specify the node and the dst")
			break
		}
		n, dst := args[0], args[1]
		oldParentID := redisConn.Get("utree:tree:" + n + ":parent").Val()
		redisConn.SRem("utree:tree:"+oldParentID+":children", n).Val()
		redisConn.SAdd("utree:tree:"+dst+":children", n).Val()
		redisConn.Set("utree:tree:"+n+":parent", dst, -1).Val()
		conn.WriteInt(1)
		break

	// remove a node
	// 1)- fetch the old parent
	// 2)- remove the node from the old parent children
	// 3)- clean the parent from the node
	case "rm":
		if len(args) < 1 {
			conn.WriteError("you must specify the deletable node")
			break
		}
		n := args[0]
		oldParentID := redisConn.Get("utree:tree:" + n + ":parent").Val()
		redisConn.SRem("utree:tree:"+oldParentID+":children", n).Val()
		redisConn.Del("utree:tree:" + n + ":parent").Val()
	}
}

func handleAcceptConn(conn redcon.Conn) bool {
	return true
}
