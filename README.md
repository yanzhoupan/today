# Today
## Intro 
example
video

## Install
curl install.sh | bash

## Usage
| Flags | Usage |
|-------|-------|
| -add | Add points to the list, use <code>&#124;</code> to separate, the points will be automatically numbered. Eg: run `today -add` and then input "first task &#124; second task &#124; ..."|
| -del | Delete points. Eg: today -del=1,2,4 would delete points 1, 2, and 4 |
| -modify | Modify a given point. Eg: today -modify=3, then type the new task |
| -list | List all the history checklists. Can also use with a limit. Eg: today -list=5 will list recent 5 today's checklist|
| -show | Show the content of a given date. Eg: today -show=20210801 |
| -clean | Clean all the points for today. Can also use with a checklist name |

## TODO
Maybe add -addTomo flag?

