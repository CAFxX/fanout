package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32 }
type GenS2 struct { F0 uint16 }
func genFn0(p0 int8, p1 string) int64 {
    for fi5 := 0; fi5 < 19; fi5++ {
        if true {
            var v6 uint = 2
            if (!false) {
                v7 := []int{}
                _ = v7
            } else {
                for fi8 := 0; fi8 < 17; fi8++ {
                    p0 = (p0 % (((p0 * 42)) | 1))
                    _ = fi8
                }
                if (!false) {
                    fmt.Println(p1)
                } else {
                    var v9 float64 = 3.14159
                    fmt.Println(p1, fi5, v6)
                    var v10 int64 = 827
                    _ = v10
                    _ = v9
                }
            }
            _ = v6
        }
        v11 := GenS1{F0: 127, F1: 2147483648}
        _ = fi5
        _ = v11
    }
    p1 = p1
    _ = p0
    _ = p1
    return -1
}

func main() {
    v1 := []int{-991, 128, 127}
    var v2 int32 = 324
    v3 := make(map[int]int)
    v4 := []int{-2147483648}
    if true {
        v2 = 0
        switch ("\n\t" + "foo bar") {
        case "foo bar":
            for fi12 := 0; fi12 < 9; fi12++ {
                var v13 uint8 = 128
                _ = fi12
                _ = v13
            }
        case "\n\t":
            var v14 string = ("a" + ("" + ("" + "foo bar")))
            _ = v14
        default:
            v15 := GenS1{F0: 2, F1: 2147483647}
            _ = v15
        }
    } else {
        var v16 uint64 = 42
        _ = len(v4)
        _ = v16
    }
    if (940 >= 958) {
        v2 -= -1
        var v17 float64 = 1.5
        _ = v17
    }
    v2 -= (v2 << ((65535) & 63))
    v18 := GenS1{F0: 247, F1: 1}
    fmt.Println(v2)
    v19 := GenS2{F0: 461}
    switch 0 {
    case 306:
        var v20 uint32 = 887
        _ = v20
    default:
        v2 = v2
    }
    if false {
        fmt.Println(v2)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
