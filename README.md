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

I'm going to start with Go as Python is my normal weapon of choice but I want something compiled to add a small level of
low complexity like:

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

Here is a high level list of the steps taken:

```// I'll add links of other useful resources I came upon on my journey as I go along and what I used them for```

1. Made a repo on Git...
2. Added a go file 
3. Added a test file
4. Added Travis using these two pages
    ``` https://docs.travis-ci.com/user/getting-started/ ```
    ``` https://docs.travis-ci.com/user/languages/go/ ```
5. Install Jenkins (on arch) ```"sudo packman -Syu jenkins"``` then start the service ```"sudo systemctrl
   jenkins.service"```
6. Logged on the ```localhost:8090``` and followed instructions and installed plugins - folders, timestamper, pipeline,
   git, LDAP, GitHub, OWASP Markup Formatter, Workspace Cleanup, Git branch Source SSH Slaves, Email Extention, Ant,
...




-------------------------
To do: Add links in Markdown
