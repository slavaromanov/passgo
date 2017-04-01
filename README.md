# passgo
#### Password generator with cli interface.
### common info
  - **collision**   - 0.016129 average frequency on 1m chars (ascii letters + digits)
  - **performance** - 0.01233s 100 chars, 0.33s 1000c, 3.62s 10000c
  - **ascetic**     - independence from third-party libarys and sys requirements 

# gnu-time
| User time (seconds)                         | 287.54  |
|:-------------------------------------------:|--------:|
| System time (seconds)                       | 19.91   |
| Percent of CPU this job got                 | 96%     |
| Elapsed (wall clock) time (h:mm:ss or m:ss) | 5:19.17 |
| Maximum resident set size (kbytes)          | 30736   |
| Minor (reclaiming a frame) page faults      | 7697    |   
| Voluntary context switches                  | 1664574 |
| Involuntary context switches                | 680082  |

### used
- [arc4random](https://www.freebsd.org/cgi/man.cgi?query=arc4random) - from BSD libc
- Binary Power - **D. Knuth: ArtComputerProgramming**, _1969_
