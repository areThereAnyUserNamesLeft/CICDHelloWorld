# CICDHelloWorld
[![Build Status](https://travis-ci.org/areThereAnyUserNamesLeft/CICDHelloWorld.svg?branch=master)](https://travis-ci.org/areThereAnyUserNamesLeft/CICDHelloWorld)
## Aim 

Create a "simple as possible" starting set up for a CI-CD tool chain

## Tools

- Source control = Git
- Ci = Jenkins / Travis-CI
- Config Management = Ansible
- Deployment = ??? 
- Testing = ???

#### Language 

I'm going to start with Go as Python is my normal weapon of choice but I want something compiled to add a small level of
low complexity like:

```go
package main

import "fmt"

func main() {
    fmt.Println("Built using:")
    fmt.Println("        Git")
    // Each time I add a new element I'll add a Println...
}

 ```

My starting point was [this article from IBM](https://www.ibm.com/cloud/garage/content/deliver/practice_delivery_pipeline/)

Here is a high level list of the steps taken:

```// I'll add links of other useful resources I came upon on my journey as I go along and what I used them for```
#### Foundations
1. Made a repo on Git...
2. Added a go file 
3. Added a test file

#### Travis-CI

4. Added Travis using instructions on these two
   pages 

[Travis-ci - getting started](https://docs.travis-ci.com/user/getting-started/) 

[Travis-ci - GO](https://docs.travis-ci.com/user/languages/go)

**Now I have Travis passing tests so  we can move on to Jenkins**

#### Jenkins

1. ~~install Jenkins (on arch) ```"sudo packman -Syu jenkins"``` then start the service ```"sudo systemctrl
   jenkins.service"```~~
6. ~~Logged on the ```localhost:8090``` and followed instructions and installed plugins - folders, timestamper, pipeline,
   git, LDAP, GitHub, OWASP Markup Formatter, Workspace Cleanup, Git branch Source SSH Slaves, Email Extention, Ant,
...~~
7. Tried again using the [Jenkins official docker image](https://github.com/jenkinsci/docker) rather than installing it locally where there might be some java
   dependancy issues occuring due to my OS. I ended up running into similar issues  
8. Keep retrying the recommended plugins install until everything is deployed...
9. Once every thing is present and correct install the [go plugin](https://wiki.jenkins.io/display/JENKINS/Go+Plugin)
   and follow the instructions to set up the tool

**WARNING - I am not securing my server but in reality I would if I were working on a real project rather than just
getting some low level experience - perhaps I 'll have that as a todo once I have the whole pipeline set up**

10. Make a "New Item" - Give it a name and you should be able to add some commands into the execute shell section, in my
    case 

```go test -v && go build CICDHelloWorld.go``` 

This does echo the Travis test but it will do for now...

**Now we have jenkins working I can look at Anisible - then the fun can begin**

#### Anisible

1. 
