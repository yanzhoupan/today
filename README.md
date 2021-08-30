# Today
## Intro 
example
video

## Install
curl install.sh | bash

## Usage
| Flags | Usage |
|-------|-------|
| -add | Add points to the list, use <code>&#124;</code> to separate, the points will be automatically numbered. Eg: run `today -add`, hit enter and then input "first task &#124; second task &#124; ..."|
| -check | Mark on tasks as DONE. Eg: today -check=1,2,5 |
| -remove | Delete points. Eg: today -del=1,2,4 would delete points 1, 2, and 4 |
| -modify | Modify a given point. Eg: today -mod=3, then type the new task |
| -history | List a limit number of the history todays. Eg: today -history=5 will list recent 5 today's checklist. Use -history=-1 to list all. |
| -show | Show the content of a given date. Eg: today -show=2021-08-01 |
| -clear | Clear all the todos for today. Can also use with a checklist name |

## TODO
Support export to txt file and download with time range
Support word cloud analysis of the checklists
Maybe add -addTomorrow flag?

