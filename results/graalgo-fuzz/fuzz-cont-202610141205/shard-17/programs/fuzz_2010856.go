package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 string; F1 int }
type GenS3 struct { F0 int64; F1 string; F2 uint16 }
type GenS4 struct { F0 int; F1 int16 }
func genFn0(p0 uint8) uint16 {
    var v6 uint32 = 2
    var v7 uint = 255
    var v8 string = "a"
    v9 := []int{-9223372036854775808, 127}
    p0 += (p0 ^ 177)
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 32767
}

func main() {
    var v1 uint32 = 323
    var v2 int = 42
    var v3 int64 = 74
    var v4 int = v2
    var v5 uint64 = 2
    v10 := GenS1{F0: ((v5 & 667) > (v5 << ((982) & 63)))}
    for fi11 := 0; fi11 < 19; fi11++ {
        v12 := []int{32768, (^(v4)), v4}
        var v13 uint8 = 143
        var v14 int64 = (-(v3))
        _ = fi11
        _ = v12
        _ = v13
        _ = v14
    }
    v15 := GenS2{F0: "\n\t", F1: (v4 * (v4 ^ int(v3)))}
    v16 := GenS3{F0: v3, F1: "\n\t", F2: 2}
    v5 = 0
    var v17 int8 = 1
    v18 := GenS4{F0: (^(v2)), F1: 704}
    v19 := GenS4{F0: v2, F1: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
