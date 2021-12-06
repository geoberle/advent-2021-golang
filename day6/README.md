# Day 6
This was one of the most enjoable challanges from the story behind it :)

I stored the lanternfish in an array of 9 fields, using their days left
until reproduction as index and the number of fish in that particular phase of their
live as value. Based on this data structure, the simulation of their population
increase was moving the fish each day one field to the left, while having special
treatment for the fish on their special day. See `SimulatePopulation`.

As a first naive implementation, that was fine, but recreating the array in
each iteration is not that efficient. My second implementation does not recreate
the array but instead moves the zero index (day 0) virtually one to the right in
each iteration. That implies that all index operations must be based on that
index offset mod 9 to wrap at the end of the array. Also nice: that way, the new
born children don't need to be copied to offset+8 - they naturally land
there because their parents were in offset+0 before moving the offset.
See `SimulatePopulationOffset`
