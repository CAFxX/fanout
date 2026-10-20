package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 string; F1 int32; F2 uint64 }
func genFn0() int8 {
    if false {
        v4 := []int{127}
        var v5 int64 = 0
        _ = v4
        _ = v5
    }
    var v6 int64 = 0
    if (true && (!true)) {
        var v7 uint16 = 139
        for fi8 := 0; fi8 < 8; fi8++ {
            v7 = 146
            _ = fi8
        }
        _ = v7
    }
    v6 += (v6 ^ 9223372036854775807)
    _ = v6
    return -127
}
func genFn1() int64 {
    v9 := GenS1{F0: 0}
    v10 := []int{128, -877, 9223372036854775807, 32768}
    var v11 uint64 = 2
    _ = v10
    _ = v11
    _ = v9
    return 532
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 string = "\n\t"
    var v3 int16 = 1000
    fmt.Println(v3, v2)
    var v12 int32 = -377
    v3 += v3
    if (255 != 1) {
        v13 := GenS2{F0: -4}
        if false {
            for fi14 := 0; fi14 < 16; fi14++ {
                v15 := GenS1{F0: -95}
                var v16 uint64 = 2
                v16 = v16
                _ = fi14
                _ = v15
                _ = v16
            }
        }
        _ = v13
    }
    var v17 uint32 = 42
    var v18 uint64 = 18446744073709551615
    if ((!true) && true) {
        switch 1 {
        case -884:
            v2 = v2
        default:
            v19 := GenS3{F0: (v2 + "foo bar"), F1: v12, F2: 4294967295}
            _ = v19
        }
    } else {
        v18 = (v18 * v18)
        v18 += 119
        v3 += v3
    }
    var v20 int8 = -1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v3
}
