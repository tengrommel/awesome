# GFS

storage


What would we like for consistency?

  Ideal model: same behavior as a single server
  
  server uses disk storage
  
  server executes client operations one at a time (even if concurrent)
  
  reads reflect previous writes
    even if server crashes and restarts
  
  thus:
    suppose C1 and C2 write concurrently, and after the writes have
      completed, C3 and C4 read. what can they see?
      
    C1: Wx1
    C2: Wx2
    C3:     Rx?
    C4:         Rx?
    answer: either 1 or 2, but both have to see the same value.
  
  This is a "strong" consistency model.
  
  But a single server has poor fault-tolerance.

Replication for fault-tolerance makes strong consistency tricky.
  
  a simple but broken replication scheme:
    
    - two replica servers, S1 and S2
    - clients send writes to both, in parallel
    - clients send reads to either
  
  in our example, C1's and C2's write messages could arrive in different orders at the two replicas
    
    - if C3 reads S1, it might see x=1
    - if C4 reads S2, it might see x=2
  
  or what if S1 receives a write, but 
    the client crashes before sending the write to S2?
  
  that's not strong consistency!
  
  better consistency usually requires communication to
    ensure the replicas stay in sync -- can be slow!

  lots of tradeoffs possible between performance and consistency
    we'll see one today
    
# GFS

Context:
  > Many Google services needed a big fast unified storage system
    Mapreduce, crawler/indexer, log storage/analysis, Youtube (?)
    
  Global (over a single data center): any client can read any file
    Allows sharing of data among applications
  
  Automatic "sharding" of each file over many servers/disks
    
    - For parallel performance
    - To increase space available
  
  Automatic recovery from failures
  
  Just one data center per deployment
  
  Just Google applications/users
  
  Aimed at sequential access to huge files; read or append
    I.e. not a low-latency DB for small items


What was new about this in 2003? How did they get an SOSP paper accepted?
  
  - Not the basic ideas of distribution, sharding, fault-tolerance.
  - Huge scale.
  - Used in industry, real-world experience.
  - Successful use of weak consistency.
  - Successful use of single master.

Overall structure

  - clients (library, RPC -- but not visible as a UNIX FS)
  - each file split into independent 64 MB chunks
  - chunk servers, each chunk replicated on 3
  - every file's chunks are spread over the chunk servers
    for parallel read/write (e.g. MapReduce), and to allow huge files
  - single master (!), and master replicas
  - division of work: master deals w/ naming, chunkservers w/ data

Master state
  in RAM (for speed, must be smallish):
    file name -> array of chunk handles (nv)
    chunk handle -> version # (nv)
                    list of chunkservers (v)
                    primary (v)
                    lease time (v)
  on disk:
    log
    checkpoint

Why a log? and checkpoint?

Why big chunks?

What are the steps when client C wants to read a file?

  - 1. C sends filename and offset to master M (if not cached)
  - 2. M finds chunk handle for that offset
  - 3. M replies with list of chunkservers
     only those with latest version
  - 4. C caches handle + chunkserver list
  - 5. C sends request to nearest chunkserver
     chunk handle, offset
  - 6. chunk server reads from chunk file on disk, returns

How does the master know what chunkservers have a given chunk?

