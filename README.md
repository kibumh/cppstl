# Useful C++ STL algorithm (only?) impelmented in go
- containers : We have a slice and a map. What else do we need?
- iterators : ?.?
- algorithms : lovely rotate, nthelement, etc.


# PLAN
## MILESTONE 1
* non-modifying sequence operations
  - [ ] AllOf, AnyOf, NoneOf
  - [ ] CountIf, FindIf
modifying sequence operations
* [x] Reverse
  - [x] Rotate
  - [x] StablePartition
  - [ ] NthElement
* others
  - [ ] Slide from Sean Parent
  - [ ] Gather from Sean Parent
## MILESTONE 2
* other algorithms
  - [ ] LowerBound, UpperBound
  - [ ] MinMax, Min, Max
  - [ ] Shuffle
* containers, iterators?

# IDEAS
* Make it generic enough to apply to both a container and a channel.
