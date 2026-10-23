package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64; F2 int16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int64; F1 bool; F2 uint8 }
func genFn0() int {
    v4 := GenS1{F0: 1, F1: 9223372036854775807, F2: 0}
    var v5 float64 = 1e10
    v6 := []int{806, 4294967295, -399, 32767}
    if ((-2.25 * 1e10) <= v5) {
        v7 := GenS1{F0: 255, F1: -909, F2: -155}
        _ = v7
    }
    if (4 == 1000) {
        v5 -= 1e10
        var v8 uint32 = 2147483647
        _ = v8
    }
    _ = v4
    _ = v5
    _ = v6
    return -480
}

func main() {
    var v1 uint = 42
    var v2 uint64 = 42
    var v3 bool = (!true)
    v9 := []int{}
    for fi10 := 0; fi10 < 9; fi10++ {
        fmt.Println(v1, v3, v1)
        v11 := GenS2{F0: 32767}
        _ = fi10
        _ = v11
    }
    v12 := GenS3{F0: 0, F1: v3, F2: 88}
    _ = len(v9)
    var v13 string = ""
    fmt.Println(v3, v2, v3)
    var v14 int8 = -53
    switch (("a" + v13) + v13) {
    case "a":
        if ((v3 && (128 != 42)) || true) {
            if ((v3 && (false && v3)) && v3) {
                v14 += 80
            }
            v13 = (v13 + (v13 + v13))
            v3 = v3
        }
    default:
        v14 += v14
    }
    _ = len(v9)
    v15 := GenS2{F0: 2}
    v1 += v1
    var v16 bool = v3
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v9
}
