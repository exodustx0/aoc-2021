# Advent of Code 2021

These are my solutions to this year's [Advent of Code](https://adventofcode.com/2021). I used it as an excuse to practice with a language which I'm not yet very familiar with but want to use more: Go.

## Daily notes

### Day 3

Decided to create this repository, private as of writing, with the intent of making it public after the event is over.

I think the setup I've got here is best, Go module-wise. I had some trouble with figuring out how to organise this project, but I can't seem to find a way to run individual days' scripts without `cd`-ing into their respective directories and `go run`ning. I'm pretty sure I've overlooked something, because this rather makes me a sad panda. Maybe it's part of Go's opinionatedness though. I'm sure I'll find out at some point, maybe during the course of this event. We'll see.

It took me an annoyingly long time to debug for part two. It was a few hours in when I _finally_ noticed that I forgot a `break` statement. Figures!

My clumsiness of the day aside (it took me too long to pull in the example input and test against _that_), I've got mixed feelings on how the debugging process felt, in terms of how doable it was with the tools that Go offers. No conclusions as of yet, but I'll be paying special attention from here on out.
