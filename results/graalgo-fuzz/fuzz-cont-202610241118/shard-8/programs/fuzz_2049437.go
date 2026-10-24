package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string; F2 bool }
type GenS2 struct { F0 string; F1 int16 }
type GenS3 struct { F0 uint }
func genFn0(p0 int8) uint8 {
    p0 = p0
    p0 += (p0 % (((-(p0))) | 1))
    var v5 int32 = 2147483647
    _ = p0
    _ = v5
    return 1
}

func main() {
    var v1 uint8 = 255
    var v2 uint16 = 65535
    var v3 float32 = (3.14159 / (1e10 - -2.25))
    var v4 int = 128
    if true {
        var v6 int32 = 1
        _ = v6
    } else {
        if true {
            v2 += v2
            fmt.Println(v4, v1, v1)
        }
    }
    v7 := []int{(v4 % (((v4 * -747)) | 1)), 1, v4, 42}
    for fi8 := 0; fi8 < 16; fi8++ {
        if len(v7) > 0 {
            v7[0] = -2147483648
        }
        var v9 int64 = -1
        _ = fi8
        _ = v9
    }
    var v10 int = (v4 - (v4 >> ((221) & 63)))
    fmt.Println(v1, v4)
    v2 -= 0
    for fi11 := 0; fi11 < 10; fi11++ {
        v3 = -2.25
        _ = fi11
    }
    for fi12 := 0; fi12 < 19; fi12++ {
        switch v4 {
        case 4294967295:
            var v13 float64 = float64(0.0)
            _ = v13
        case 2147483647:
            if ((!false) || false) {
                for fi14 := 0; fi14 < 5; fi14++ {
                    if (18446744073709551615 != 42) {
                        v15 := GenS1{F0: 201, F1: "\n\t", F2: true}
                        _ = v15
                    }
                    v2 -= (v2 & (v2 - v2))
                    _ = fi14
                }
                if true {
                    v16 := GenS2{F0: "a", F1: -1}
                    _ = v16
                }
            }
        }
        _ = fi12
    }
    v3 -= v3
    if false {
        var v17 uint = 754
        if (!(uint8(v17) <= v1)) {
            v10 = 128
        } else {
            var v18 uint8 = v1
            var v19 int64 = 1
            _ = v18
            _ = v19
        }
        _ = v17
    }
    if (v2 > v2) {
        v20 := GenS3{F0: 255}
        _ = v20
    }
    var v21 float64 = 0.0
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v10
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v7
}
