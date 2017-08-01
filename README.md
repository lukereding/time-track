# time-track

I wanted a tool that I can use to keep track of my time that I didn't have to remember to use.

The goal here is to write a CLI in Go that can be run via `cron` so that I don't have to remember to 'punch in' and 'punch out' whenever I switch tasks.

Every `X` minutes while I'm logged in, a terminal window pops up and asks me what I'm working on.

This project is part of my experience learning Go, which is why it's terribly written. Very much a WIP.

### what it looks like

```bash
⚡ time-track
0) surfing the web
1) researching the death of Biggie
2) watching Parks and Rec
3) cat gifs

▽ What are you working on?
3
```

You type in the number corresponding to what you're doing and move on with your life.

### main commands

There are projects. These are things you are working on. You can add and remove projects:

`time-track --add-project "some project"` : add projects you're working on.
`time-track --rm-project "some project"` : remove some project you're done with or want to forget about.

`time-track` : meant to be run as a `cron` job. To run every 15 minutes, use something like

```bash
*/15 * * * * if [ $USER == "YOUR_USERNAME" ]; then open /usr/local/bin/time-track; fi
```

### other notes

`time-track` writes a couple file. `~/.time_track` is the config file that, of now, just stores your projects. You can edit by hand or use `time-track --add-project` or `time-track --rm-project` via the command libe.

`~/.time-track.csv` stores the results. In the future other commands will be possible that will allow the user to generate the top `n` projects she worked on that week, etc.

## to do

- [X] save a csv
- [X] get the date
- [X] present options to a user :: use https://github.com/dixonwille/wmenu
- [X] create config file that holds the activities / projects to load
- [ ] `cron` might not be nec.: check [here](https://gobyexample.com/timers)
- [X] config file should save the activities
- [X] flags to `time-track` to add activities
- [X] flag to `time-track` to remove activities
- [ ] don't store repeated projects
- [ ] write functions (!) to simplify code
- [ ] add command line option for `--report` that summarizes the top job for the week, month, etc.
