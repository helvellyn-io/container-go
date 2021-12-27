# container-go
[![Build Status]Passing]
Container actions using Golang Docker SDK 

- For use with Temporal build task queues. 
  For guidance on how to run this as a Temporal task please refer to: 
  * https://github.com/helvellyn-io/temporal-template


```
* cd build 
* make articats //gets build artifacts from Git repository.
* make build    //build container.
* make push     //pushes containre to registry.
* make test     //runs ad-hoc testing.
* make clean    //deletes artifacts.
```


