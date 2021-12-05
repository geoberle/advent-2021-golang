# Day 4
There are several solutions that came to my mind. I though about building
up a big data structure to keep track what number were observed are
located on what position and if they are marked or not.
This would have been an easy solution but, as with the previous challenges,
i wanted to keep only the necessary things in memory: the numbers in flat list.

In the end this, resulted in much more code. But it was still a good golang kata.

I could have reduced memory consumtion event further by keeping the entire data
of the file as string in memory without converting them to numbers. The offset
calculations would have been fun :)

Besides the challenge itself, i played with a couple of new
golang things, namely structs and goroutines. While not really required,
goroutines were easily applied to calculate the score of each board
concurrently.
