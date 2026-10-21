package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8; F2 string }
type GenS2 struct { F0 uint8 }
func genFn0(p0 int8, p1 int, p2 string) uint16 {
    v7 := []int{}
    p1 += p1
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 42
}

func main() {
    var v1 float64 = 3.14159
    v2 := make(map[int]uint32)
    v3 := []int{136, 752, 535}
    var v4 int32 = 306
    v5 := make(map[string]bool)
    var v6 int8 = 127
    if (42 >= 18446744073709551615) {
        var v8 int8 = (v6 >> ((406) & 63))
        fmt.Println(v8, v6, v8)
        var v9 uint16 = 808
        _ = v8
        _ = v9
    } else {
        if (42 <= 24) {
            v1 = (float64(1.5) / (v1 + float64(0.1)))
        }
        for fi10 := 0; fi10 < 11; fi10++ {
            v1 = v1
            _ = fi10
        }
        if (-2.25 >= (3.14159 + 1.5)) {
            var v11 string = "\n\t"
            _ = v11
        }
    }
    for fi12 := 0; fi12 < 7; fi12++ {
        var v13 bool = (0 != 255)
        fmt.Println(fi12, v6)
        _ = fi12
        _ = v13
    }
    if (!false) {
        v14 := GenS1{F0: 42, F1: -47, F2: "foo bar"}
        for fi15 := 0; fi15 < 3; fi15++ {
            v6 = v6
            _ = fi15
        }
        _ = v14
    }
    if (!((!false) && true)) {
        if false {
            v1 += v1
        }
        var v16 int16 = -727
        v3 = append(v3, 2147483647)
        _ = v16
    } else {
        v4 += v4
    }
    var v17 bool = (!(4294967295 <= 316))
    if (!v17) {
        v18 := GenS2{F0: 0}
        v4 = (v4 * (v4 - (^(v4))))
        _ = v18
    }
    var v19 uint64 = 4294967295
    fmt.Println(v19, v19)
    delete(v2, 1)
    fmt.Println(v19)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
