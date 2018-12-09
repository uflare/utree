uTree
======
> a very simple tree based data structure based on redis and speaks redis

Available Commands
==================
- `ping`, returns a `PONG`.
- `gen`, generates a new unique id.
- `append <parent>`, appends a new node to the spefified parent and returns the newly added id. 
- `flatten <id>`, returns all nested children for the specified id.
- `tree <id>`, return json string represents a nested tree for the specified id.
- `mv <id> <dst>`, moves the specified id to another parent `dst`.
- `rm <id>`, removes the specified id from the tree.

SDKs
====
> you can connect it with any redis client

Credits
=======
copyright (c) 2018