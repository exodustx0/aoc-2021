# Advent of Code 2021

These are my solutions to this year's [Advent of Code](https://adventofcode.com/2021). I used it as an excuse to practice with a language which I'm not yet very familiar with but want to use more: Go.

## Daily notes

### Day 3

Decided to create this repository, private as of writing, with the intent of making it public after the event is over.

I think the setup I've got here is best, Go module-wise. I had some trouble with figuring out how to organise this project, but I can't seem to find a way to run individual days' scripts without `cd`-ing into their respective directories and `go run`ning. I'm pretty sure I've overlooked something, because this rather makes me a sad panda. Maybe it's part of Go's opinionatedness though. I'm sure I'll find out at some point, maybe during the course of this event. We'll see.

It took me an annoyingly long time to debug for part two. It was a few hours in when I _finally_ noticed that I forgot a `break` statement. Figures!

My clumsiness of the day aside (it took me too long to pull in the example input and test against _that_), I've got mixed feelings on how the debugging process felt, in terms of how doable it was with the tools that Go offers. No conclusions as of yet, but I'll be paying special attention from here on out.

### Day 4

This repo is now public. What the heck, might as well. Not like I'm a contender for the global leaderboard anyway :P

I loved today's challenge. I'm not trying to be done as fast as possible; rather, I try to learn where I feel that I can do better, to see where I can break free from thinking about problems in the usual chronological order (e.g. rather than mark every bingo board after every call, do the entire call sequence per board), etc. I really enjoy working with Go struct methods, and though I don't have much previous experience with pointers (only through studying certain big C++ projects), I really find that I love working with them and can do so intuitively.

On the flipside, Go's opinionatedness has really shown it's — in my opinion — ugly side today, specifically in how it "beautifies" math expressions. Take this line that I ended up with:

```go
(*boards)[lastBoard].unmarkedSum()*int((*calls)[lastCall])
```

Dunno about you, but that reads like yikes to me. I'm a big proponent for space-separating tokens in expressions like these. Even though the difference would be only two spaces, look:

```go
(*boards)[lastBoard].unmarkedSum() * int((*calls)[lastCall])
```

Especially since `*` is used for both multiplying and dereferencing, this is _much_ more legible to me. Note that Go will format to this only when this expression is _one_ of the arguments to a function call; for some reason, however, it'll actually format to the _space-separated_ version if it's the only argument in a function call. I guess that Go wants spaces to separate both arguments and expression tokens. One space-separation to rule them all.

(Also, in case it comes up at some point: I don't mind going back and reworking previous days' code. I'm a perfectionist. Posterity can still have the commit history!)
