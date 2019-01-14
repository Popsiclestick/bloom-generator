# Bloom Generator

Small go project for writing out bloom filters to disk, based off [willf's library](https://github.com/willf/bloom).

This code reads in a newline delimited file and returns a binary bloom filter.

[Calculator](https://hur.st/bloomfilter/) for finding _m_ & _k_

```
:; ./bloom-generator --help
Usage of ./bloom-generator:
  -b uint
    	Bits (m) (default 1)
  -f string
    	New line delimited file (default "bloom-input")
  -h uint
    	Hashing functions (k) (default 1)
```
