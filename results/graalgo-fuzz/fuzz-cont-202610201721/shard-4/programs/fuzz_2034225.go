package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }

func main() {
    var v1 int8 = 127
    var v2 float64 = 0.1
    var v3 float64 = (v2 * 1e10)
    v4 := make(map[uint32]uint32)
    v5 := make(map[int]string)
    var v6 uint64 = 443
    for fi7 := 0; fi7 < 7; fi7++ {
        var v8 bool = false
        _ = fi7
        _ = v8
    }
    switch "" {
    case "\n\t":
        v6 -= (v6 << ((756) & 63))
    case "":
        v3 = v2
    case "a":
        if (1 != 0) {
            var v9 int16 = 1
            _ = v9
        } else {
            _ = len(v4)
            var v10 float64 = v2
            _ = v10
        }
    default:
        v3 += v2
    }
    if false {
        v2 = v3
    }
    v2 += (0.0 / v3)
    v11 := []int{}
    var v12 int64 = -1
    var v13 string = "a"
    for fi14 := 0; fi14 < 14; fi14++ {
        switch 32768 {
        case -510:
            v12 = v12
        default:
            var v15 int8 = -1
            _ = v15
        }
        v16 := GenS1{F0: false}
        _ = fi14
        _ = v16
    }
    v12 -= (v12 % (((v12 << ((18446744073709551615) & 63))) | 1))
    var v17 int32 = -1
    v17 = v17
    v13 = ""
    fmt.Println(v17, v12, v6)
    switch v13 {
    case "a":
        if false {
            v6 = 1
        }
    default:
        if false {
            if false {
                var v18 int64 = 1
                fmt.Println(v1)
                v6 -= (v6 & 65)
                _ = v18
            }
            if len(v11) > 0 {
                v11[0] = 2147483647
            }
            _ = len(v11)
        } else {
            if len(v11) >= 2 {
                v11 = v11[1:len(v11)]
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
