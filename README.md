# Toy Google File System

This is a toy implementation of the Google File System written in Golang, based on [this paper](https://static.googleusercontent.com/media/research.google.com/en//archive/gfs-sosp2003.pdf).

The intent is to show a possible implementation of a "create" API call. The paper details architecture flow for reads and writes, but doesn't go deeply into how file creation is done. From the information given in the paper, this is my crack at how file creation may have been done.

This project is incomplete. The work so far just simulates interactions between a client and the coordinator node at file creation.

# Future work:
## Chunk namespace manager
- A chunk namespace manager is not implemented. The paper states that file chunks are given unique chunk handles (ID numbers) at creation, and the way that's done here is just a global int variable that's incremented by 1 whenever the number is assigned to a chunk.

## Chunkservers
- Actual file creation is not implemented. The intent was to run a bunch of docker containers, each running a copy of chunkserver code. File creation should be done at the chunkserver level.