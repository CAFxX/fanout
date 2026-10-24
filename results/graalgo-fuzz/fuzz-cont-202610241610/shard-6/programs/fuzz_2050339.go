package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int }

func main() {
    v1 := make(map[int]bool)
    var v2 int = 485
    var v3 uint = 198
    v4 := []int{(v2 % ((-1) | 1))}
    var v5 uint = v3
    var v6 uint = (v3 - (^(v3)))
    switch (v2 ^ v2) {
    case 4294967295:
        if true {
            fmt.Println(v6)
            var v7 int8 = 0
            for fi8 := 0; fi8 < 3; fi8++ {
                v7 += 1
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                v9 := GenS1{F0: 9223372036854775807}
                _ = fi8
                _ = v9
            }
            _ = v7
        } else {
            v2 -= (v2 ^ (v2 << ((v6) & 63)))
            v2 = -860
            v6 = v3
        }
    case 9223372036854775807:
        v10 := GenS1{F0: 39}
        _ = v10
    case 127:
        v3 += (v6 >> ((v6) & 63))
    default:
        fmt.Println(v6, v6, v3)
    }
    if false {
        if (float64(1e10) != (float64(1e10) + 0.0)) {
            v5 -= 65535
            var v11 uint16 = 255
            _ = v11
        }
        for fi12 := 0; fi12 < 17; fi12++ {
            v2 = fi12
            _ = fi12
        }
        v3 -= (v6 * (v3 / ((2) | 1)))
    } else {
        var v13 int64 = -9223372036854775808
        v14 := GenS2{F0: 1}
        fmt.Println(v2, v13, v13)
        _ = v13
        _ = v14
    }
    v2 -= v2
    for fi15 := 0; fi15 < 12; fi15++ {
        delete(v1, (fi15 >> ((v3) & 63)))
        _ = fi15
    }
    v16 := GenS2{F0: 42}
    switch "a" {
    case "a":
        var v17 int64 = 9223372036854775807
        _ = v17
    case "hello":
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    fmt.Println(v5, v6, v3)
    var v18 uint16 = 916
    fmt.Println(v3, v6, v6)
    var v19 int = (v2 | v2)
    v18 -= 714
    v4 = append(v4, v19)
    for fi20 := 0; fi20 < 9; fi20++ {
        fmt.Println(v18)
        v21 := GenS3{F0: 463}
        _ = fi20
        _ = v21
    }
    if (-1 < -553) {
        v22 := v1[v2]
        _ = v22
    }
    var v23 uint = (v3 - (v3 ^ v5))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v18
    _ = v19
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
