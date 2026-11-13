package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int; F2 int }
type GenS2 struct { F0 int; F1 int32; F2 uint }
func genFn0() int {
    v7 := []int{4294967295, -2147483648, 32768}
    var v8 float64 = float64(1e10)
    v8 = v8
    _ = v7
    _ = v8
    return 4294967295
}

func main() {
    var v1 uint = 382
    v2 := make(map[string]uint32)
    v3 := make(map[int]bool)
    var v4 float64 = float64((0.1 * float32(float64(1e10))))
    var v5 float64 = v4
    v6 := make(map[int]int)
    fmt.Println(v5, v5)
    v4 = (0.0 - ((((v5 + (0.1 * v5)) * -2.25) * (v4 / (v5 - v4))) * v5))
    if (false || false) {
        v9 := GenS1{F0: 4294967295, F1: genFn0(), F2: -830}
        _ = v9
    } else {
        var v10 uint16 = 2
        _ = v10
    }
    v11 := v3[4294967295]
    if (false && v11) {
        fmt.Println(v5, v11)
    } else {
        v12 := GenS2{F0: 759, F1: 2147483647, F2: v1}
        var v13 uint8 = 127
        v14 := []int{748, -64, 2147483647, 9223372036854775807}
        _ = v12
        _ = v13
        _ = v14
    }
    var v15 float64 = float64(float32(v5))
    for fi16 := 0; fi16 < 7; fi16++ {
        fmt.Println(v11, v5, v4)
        _ = fi16
    }
    var v17 uint32 = 916
    var v18 uint = 65535
    delete(v2, "hello")
    v18 -= (v1 * v18)
    if false {
        fmt.Println(v11, v4, v11)
        var v19 bool = v11
        if (v17 < (v17 * 0)) {
            v20 := []int{-967, -350, int(v17)}
            v15 = v15
            switch ("a" + "a") {
            case "foo bar":
                v18 -= (v1 + (v1 - (^(v1))))
            }
            _ = v20
        }
        _ = v19
    }
    v21 := GenS1{F0: (^(v1)), F1: 42, F2: 1}
    for fi22 := 0; fi22 < 6; fi22++ {
        var v23 float32 = float32(float64((0.1 / 1e10)))
        var v24 float64 = (1.5 * ((1e10 + ((v15 / v5) * v4)) * -2.25))
        _ = fi22
        _ = v23
        _ = v24
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v15
    _ = v17
    _ = v18
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
