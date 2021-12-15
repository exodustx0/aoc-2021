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

On the flipside, Go's opinionatedness has really shown it's — in my opinion — ugly side today, specifically in how it "beautifies" math expressions. Take [this line](https://github.com/exodustx0/aoc-2021/blob/8194544ed3/day-4/main.go#L154) that I ended up with (it was cleaned up a bit on day 6):

```go
(*boards)[lastBoard].unmarkedSum()*int((*calls)[lastCall])
```

Dunno about you, but that reads like yikes to me. I'm a big proponent for space-separating tokens in expressions like these. Even though the difference would be only two spaces, look:

```go
(*boards)[lastBoard].unmarkedSum() * int((*calls)[lastCall])
```

Especially since `*` is used for both multiplying and dereferencing, this is _much_ more legible to me. Note that Go will format to this only when this expression is _one_ of the arguments to a function call; for some reason, however, it'll actually format to the _space-separated_ version if it's the only argument in a function call. I guess that Go wants spaces to separate arguments _or_ expression tokens, but never both at the same time. One space-separation to rule them all.

(Also, in case it comes up at some point: I'm not averse to going back and reworking previous days' code. I'm a perfectionist. Posterity can still have the commit history!)

### Day 5

When I first went most of the way through [A Tour of Go](https://go.dev/tour), I anticipated annoyance-to-come with how `if` conditions require brackets and there are no ternary conditions. The `(*Grid).drawLines` method that I wrote today made that come true. I tried to work around it by using a `switch` block and a label break, but any (significant) improvements in cleanliness were dashed by Go's formatter. I'm not saying it can't be done better; after all, I've only recently taken the plunge into Go. But I'm very much worried that all I'll end up learning about this is "deal with it". I suppose time will tell.

### Day 6

Nice breather, this :) very enjoyable challenge today, and it resulted in quite nicely legible code in my opinion. Go shone brightly for me today!

### Day 8

OK, today's puzzle got involved. I'm liking this. I'm not fully confident that what I wrote is optimal, especially concerning the handling of pointers (I'm still somewhat new to working with them directly), but I got through it and had fun doing it.

With regards to the pointer stuff, I should probably look into ways to get memory metrics, so I can try different stuff, see how much memory they take. So I can tell if I'm doing the right thing, in terms of learning the right habits regarding using pointers.

### Day 11

Welp. I sure used `goto`. Please forgive me.

### Day 12

I feel like I've got the hang of Go now. At that, I've also come to terms with how Go can, at times, get to look verbose (as I described on [day 5](#day-5)); after all, Go is more similar to C than to, say, TypeScript (yep, that's my frame of reference). I can see myself using this for real-world stuff.

### Day 15

Right... Well, I got what I asked for: a spike in difficulty! This day delivered. Pathfinding algorithms, y'all. So, I've looked into pathfinding algorithms before in passing, for their use in solving mazes which I find fascinating. I remembered a few minutes into the challenge that I'd once seen [a video that demonstrates the A* algorithm](https://youtu.be/icZj67PTFhc) (followed this guy for his NES stuff, stuck around for videos like these), so I rewatched it, reacquainted myself with the algorithm, and mostly copied what he did for my implementation.

The heuristics part of the algorithm _properly_ baffles me. Simply returning `1` resulted in the fastest (~8s on my machine) correct solve. I might at some point return to this and try to figure out how this _should_ work, but for now, this works, it netted me the points, and I want to clear my head of this!
