# Today
## Intro 
`today` - a commandline todo list. 

introduction video here

## Install (linux and macOS)
curl -fsSL https://raw.githubusercontent.com/yanzhoupan/today/main/install.sh | bash

## Usage
| Flags | Usage |
|---------|---------|
| -add | Add points to the list, use <code>&#124;</code> to separate, the points will be automatically numbered. Eg: run `today -add`, hit enter and then input "first task &#124; second task &#124; ..."|
| -check | Mark on tasks as DONE. Eg: `today -check=1,2,5` |
| -rm | Delete points from today's todo list. Eg: `today -rm=1,2,4` would delete points 1, 2, and 4 |
| -mod | Modify a given point. Eg: `today -mod=3`, then type the new task |
| -ls | List a limit number of the history todays. Eg: `today -list=5` will list recent 5 today's checklist. Use -history=-1 to list all. |
| -show | Show the content of a given date. Eg: `today -show=2021-08-01` |
| -clear | Clear all the todos for today. |

## TODO
Support export to txt file and download with time range
Support word cloud analysis of the checklists
Maybe add -addTomorrow flag?

