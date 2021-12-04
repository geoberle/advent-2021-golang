# Day 1
I learned a lot about golang in this challenge.

Coming from Python with context managers, the defer call to close the file seemed odd at first, but hey each language has its own ways.

I played a bit with different ways of solving part 2. I wanted to implement it in a way that only the necessary data is kept in memory. So i moved from the "reading everythin into memory" approach i used in part 1 to a "processing line by line" approach. For that i introduced a ring buffer to keep the last three read values in memory, just enough to compare the "oldest" value with the currently read one from the file. See `ProcessSectionSonarLineByLine`.

I was unhappy with the way that file handling and type conversion was making the actual calculation harder to read, so i was looking for something like Pythons generators to decouple this and found golang channels. See `ProcessSectionSonarFromChannel` and `ReadInputIntoChannel`. This way lines are still read and processed line by line but the actual processing is only taking care of ... well, the processing.

I don't know if channels are really the way so solve something like this or if i'm just pythoning up my go code for no reason at all.
