Concurrent reads and writes
===========================
The purpose of this code is to create a data structure that permits
concurrent reads and writes and learn a bit about go's various
concurrency primitives along the way.

# Results
BenchmarkUnsafe-4            500           4239636 ns/op
BenchmarkLocked-4             50          35497691 ns/op
BenchmarkLockSwap-4          100          25712943 ns/op
BenchmarkSwap-4              500           4282290 ns/op
