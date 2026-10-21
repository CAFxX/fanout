package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 string; F1 int8; F2 int8 }
type GenS3 struct { F0 int; F1 int8; F2 uint }
func genFn0(p0 int32, p1 uint32, p2 int8) uint64 {
    p1 = (p1 * p1)
    var v4 int32 = p0
    if true {
        var v5 int8 = p2
        v6 := GenS1{F0: (p2 | (^(p2)))}
        if true {
            var v7 uint16 = 12
            _ = v7
        } else {
            p1 = (p1 * 4294967295)
            v8 := GenS2{F0: "hello", F1: (v5 % (((p2 >> ((335) & 63))) | 1)), F2: v5}
            _ = v8
        }
        _ = v5
        _ = v6
    } else {
        p1 -= p1
        fmt.Println(v4)
        if (9223372036854775807 <= int64(p2)) {
            if false {
                v9 := []int{32768, 9223372036854775807}
                _ = v9
            }
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 0
}

func main() {
    v1 := make(map[uint32]string)
    v2 := []string{"hello", "a", "foo bar"}
    var v3 int = -2147483648
    v3 = (v3 - (^(v3)))
    var v10 int16 = 1
    v11 := GenS3{F0: int(v10), F1: 127, F2: 560}
    fmt.Println(v3, v10)
    var v12 int = 4294967295
    fmt.Println(v10)
    var v13 int8 = 0
    for fi14 := 0; fi14 < 2; fi14++ {
        for fi15 := 0; fi15 < 19; fi15++ {
            v2 = append(v2, ("\n\t" + "\n\t"))
            var v16 uint16 = 0
            var v17 int32 = -1
            _ = fi15
            _ = v16
            _ = v17
        }
        for fi18 := 0; fi18 < 5; fi18++ {
            v13 = (v13 | v13)
            fmt.Println(v10, v12, v13)
            var v19 float32 = 0.0
            _ = fi18
            _ = v19
        }
        var v20 int8 = (v13 ^ (v13 * v13))
        _ = fi14
        _ = v20
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
}
