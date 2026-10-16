package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string }
type GenS2 struct { F0 int32; F1 int; F2 int32 }
type GenS3 struct { F0 int8 }
func genFn0(p0 int8, p1 string, p2 uint) int64 {
    p0 = 100
    v6 := GenS1{F0: false, F1: "foo bar"}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 9223372036854775807
}

func main() {
    var v1 uint8 = 205
    var v2 float64 = 0.0
    v3 := make(map[int]bool)
    v4 := make(map[int]bool)
    v5 := make(map[uint32]bool)
    v7 := GenS2{F0: -1, F1: -1, F2: 2147483647}
    v2 = (0.0 / v2)
    v8 := GenS2{F0: -187, F1: -2147483648, F2: 0}
    v2 -= 0.0
    for fi9 := 0; fi9 < 11; fi9++ {
        var v10 uint64 = 42
        fmt.Println(v1, fi9, v10)
        _ = fi9
        _ = v10
    }
    if true {
        var v11 int16 = 0
        v2 += v2
        _ = v11
    } else {
        switch "hello" {
        case "\n\t":
            for fi12 := 0; fi12 < 5; fi12++ {
                v13 := GenS3{F0: -1}
                v1 = 0
                _ = fi12
                _ = v13
            }
        case "foo bar":
            if (609 <= 407) {
                for fi14 := 0; fi14 < 8; fi14++ {
                    v1 -= 1
                    _ = fi14
                }
                v2 += 1e10
            } else {
                var v15 float64 = (0.1 / (float64(1e10) / (v2 + float64((3.14159 + (0.1 / float32(v2)))))))
                v1 += 42
                v1 -= v1
                _ = v15
            }
        }
    }
    fmt.Println(v1, v1)
    var v16 bool = (!false)
    v17 := v4[-9223372036854775808]
    v18 := []int{}
    v19 := GenS3{F0: 0}
    var v20 int32 = -449
    delete(v5, 4294967295)
    if len(v18) > 0 {
        v18[0] = 42
    }
    _ = len(v18)
    v21 := GenS3{F0: 1}
    v16 = (v17 && (!v17))
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
