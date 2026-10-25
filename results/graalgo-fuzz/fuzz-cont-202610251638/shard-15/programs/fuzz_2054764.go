package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string; F2 int8 }

func main() {
    v1 := make(map[string]int)
    v2 := make(map[int]int)
    v3 := make(map[string]int)
    var v4 int8 = 127
    var v5 float32 = 1.5
    v4 += v4
    v5 += v5
    fmt.Println(v5)
    v4 -= 42
    var v6 uint8 = 1
    switch -9223372036854775808 {
    case -987:
        for fi7 := 0; fi7 < 7; fi7++ {
            v8 := GenS1{F0: 42, F1: "hello", F2: (v4 % ((v4) | 1))}
            if true {
                fmt.Println(v5, fi7)
                fmt.Println(fi7, v4, v5)
                v4 = v4
            } else {
                var v9 uint = 18446744073709551615
                v9 = (v9 << ((42) & 63))
                _ = v9
            }
            _ = fi7
            _ = v8
        }
    case -628:
        v1[("foo bar" + ("" + ("a" + "")))] = -63
    case 1:
        _ = len(v2)
    default:
        v10 := GenS1{F0: 354, F1: "foo bar", F2: v4}
        _ = v10
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
