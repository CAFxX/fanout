package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64 }
type GenS2 struct { F0 string; F1 uint32 }
func genFn0(p0 uint16, p1 int16) int8 {
    var v4 bool = (!(!true))
    p0 = p0
    p1 = p1
    _ = p0
    _ = p1
    _ = v4
    return -30
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[uint32]uint32)
    var v3 uint64 = 42
    v3 -= (v3 - (v3 + 471))
    v5 := GenS1{F0: v3, F1: 892}
    v3 = (v3 % ((4294967295) | 1))
    switch ("a" + "") {
    case "hello":
        if (1.5 > (1.5 * 0.1)) {
            for fi6 := 0; fi6 < 13; fi6++ {
                for fi7 := 0; fi7 < 18; fi7++ {
                    v8 := GenS1{F0: v3, F1: 42}
                    var v9 int64 = -1
                    v10 := []int{32768, fi7, (fi7 << ((794) & 63)), (fi6 + (fi7 | 9223372036854775807))}
                    _ = fi7
                    _ = v10
                    _ = v8
                    _ = v9
                }
                fmt.Println(v3, fi6)
                _ = fi6
            }
        }
    default:
        var v11 int64 = -693
        _ = v11
    }
    var v12 int64 = 339
    fmt.Println(v3, v3)
    v13 := GenS2{F0: "foo bar", F1: 0}
    var v14 int8 = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v5
}
