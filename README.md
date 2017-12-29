[![Build Status](https://travis-ci.org/kibumh/cppstl.png)](https://travis-ci.org/kibumh/cppstl)

# Useful C++ STL algorithms impelmented in go

# Quick Start
```
go get -u github.com/kibumh/cppstl/algorithm
```

# Go Doc
https://godoc.org/github.com/kibumh/cppstl/algorithm

# PLAN
## MILESTONE 1
* non-modifying sequence operations
  - [x] AllOf, AnyOf, NoneOf
  - [ ] CountIf, FindIf
* modifying sequence operations
  - [x] Reverse
  - [x] Rotate
  - [x] StablePartition
  - [ ] NthElement
* others
  - [x] make them receive a slice. (like sort.Slice)
  - [ ] Slide, Gather (See C++ Seasoning talk presented by Sean Parent)
## MILESTONE 2
* other algorithms
  - [ ] LowerBound, UpperBound
  - [ ] MinMax, Min, Max
  - [ ] Shuffle
* containers, iterators?

# IDEAS
* Make it generic enough to apply to both a container and a channel.
