package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }

func main() {
    var v1 int32 = 1
    v2 := []string{("hello" + ""), ((("hello" + (("" + "\n\t") + ("\n\t" + "a"))) + "hello") + ""), (("foo bar" + "a") + (("hello" + "") + ("hello" + "a")))}
    v3 := make(map[uint32]int)
    var v4 uint32 = 4294967295
    var v5 uint64 = 42
    if (!(1 <= -1)) {
        v4 += v4
        for fi6 := 0; fi6 < 19; fi6++ {
            if (0 < 1) {
                _ = len(v3)
            } else {
                v3[v4] = (^(fi6))
            }
            _ = fi6
        }
        fmt.Println(v4)
    } else {
        v7 := GenS1{F0: v4}
        fmt.Println(v1)
        var v8 int64 = -9223372036854775808
        _ = v7
        _ = v8
    }
    v9 := GenS1{F0: (v4 + (v4 ^ v4))}
    v5 -= v5
    v10 := GenS1{F0: 793}
    for fi11 := 0; fi11 < 13; fi11++ {
        for fi12 := 0; fi12 < 14; fi12++ {
            v2 = append(v2, "\n\t")
            v1 = 1
            _ = fi12
        }
        _ = fi11
    }
    var v13 int64 = -9223372036854775808
    v1 += v1
    fmt.Println(v4, v1)
    v13 -= (v13 / ((0) | 1))
    for fi14 := 0; fi14 < 3; fi14++ {
        delete(v3, (v4 / (((v4 - (v4 + uint32(v5)))) | 1)))
        var v15 int = fi14
        _ = fi14
        _ = v15
    }
    for fi16 := 0; fi16 < 15; fi16++ {
        v1 += v1
        _ = fi16
    }
    v17 := GenS1{F0: v4}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v13
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
