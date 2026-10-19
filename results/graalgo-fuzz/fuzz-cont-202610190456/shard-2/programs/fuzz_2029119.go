package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }

func main() {
    var v1 uint64 = 2
    var v2 uint32 = uint32(v1)
    v3 := make(map[string]int)
    var v4 int16 = -32768
    var v5 int = 939
    v6 := make(map[string]int)
    v7 := GenS1{F0: -58}
    v1 -= (v1 << ((0) & 63))
    v8 := []int{v5}
    v4 = v4
    var v9 int = (v5 >> ((4294967295) & 63))
    for fi10 := 0; fi10 < 6; fi10++ {
        if ((((1e10 + -2.25) + (0.1 / (-2.25 * (1e10 - 0.0)))) / 1.0) != 1e10) {
            if false {
                v6["\n\t"] = v5
                v2 += v2
                var v11 uint64 = v1
                _ = v11
            } else {
                var v12 int = (fi10 * v5)
                _ = v12
            }
        }
        var v13 bool = (-128 < 1)
        switch ("\n\t" + "a") {
        case "":
            v5 -= (v9 << ((65535) & 63))
        }
        _ = fi10
        _ = v13
    }
    var v14 uint = 2
    for fi15 := 0; fi15 < 9; fi15++ {
        if len(v8) >= 2 {
            v8 = v8[1:len(v8)]
        }
        var v16 int16 = (v4 << ((v14) & 63))
        if ((v14 ^ 42) >= (v14 - v14)) {
            var v17 int8 = 127
            _ = v17
        } else {
            var v18 int64 = -1
            _ = v18
        }
        _ = fi15
        _ = v16
    }
    fmt.Println(v1, v2, v14)
    if (!false) {
        v19 := v6[("a" + "a")]
        _ = v19
    } else {
        v20 := GenS1{F0: 0}
        _ = v20
    }
    if (1 <= 9223372036854775807) {
        fmt.Println(v5, v1)
        fmt.Println(v1, v2)
    }
    v4 = v4
    v9 = (v5 % ((v5) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
