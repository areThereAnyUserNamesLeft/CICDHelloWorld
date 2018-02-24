# CICDHelloWorld

## Aim 

Create a "simple as possible" starting set up for a CI-CD tool chain

## Tools

- Source control = Git
- Ci = Jenkins / Travis-CI
- Config Management = Ansible
- Deployment = ??? 
- Testiong = ???

## Language 

I'm going to start with Go as Python is my normal weapon of choice but I wnt something compiled to add a small level of
complexity

```
package main

import "fmt"

func main() {
    fmt.Println("Built using:")
    fmt.Println("        Git")
    // Each time I add a new element I'll add a Println...
}

 ```

My starting point was this article:

```https://www.ibm.com/cloud/garage/content/deliver/practice_delivery_pipeline/```

I'll add a bibiliography of other useful resources I came upon on my journey as I go
