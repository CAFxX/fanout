package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int32; F1 bool }
func genFn0() uint8 {
    var v7 float64 = 1e10
    if (false && ((v7 / ((v7 / v7) - (v7 + v7))) >= 3.14159)) {
        v8 := []int{-93, 687, 32768, 806}
        v7 = (v7 - v7)
        _ = v8
    } else {
        if false {
            v9 := []int{1, 42, -751, 0}
            _ = v9
        } else {
            v10 := []int{}
            _ = v10
        }
        for fi11 := 0; fi11 < 11; fi11++ {
            if ((1e10 > (-2.25 + 3.14159)) && ((-2.25 * (0.1 - -2.25)) >= float64(float32(v7)))) {
                var v12 uint32 = 2
                _ = v12
            }
            fmt.Println(v7, v7)
            switch 2147483647 {
            case 187:
                var v13 float32 = 1e10
                _ = v13
            case 32767:
                var v14 uint16 = 65535
                _ = v14
            case 128:
                var v15 int16 = 32767
                _ = v15
            default:
                var v16 int8 = -128
                _ = v16
            }
            _ = fi11
        }
        v17 := []int{-570, 601, -1}
        _ = v17
    }
    _ = v7
    return 2
}

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 string = "\n\t"
    var v3 string = (((v2 + v2) + (v2 + v2)) + v2)
    v4 := make(map[string]uint32)
    var v5 int16 = 42
    var v6 int = 128
    var v18 float64 = (((3.14159 / 1.5) / 1.5) / ((0.1 - (0.1 + (-2.25 / 1e10))) + (3.14159 / float64(1.5))))
    fmt.Println(v2, v2)
    v19 := GenS1{F0: 42}
    var v20 uint64 = 18446744073709551615
    v21 := GenS2{F0: 1, F1: true}
    for fi22 := 0; fi22 < 17; fi22++ {
        for fi23 := 0; fi23 < 1; fi23++ {
            v18 += float64(float32(((float64((0.1 * ((1e10 * (((0.0 * (((float32(float64(((3.14159 - ((0.0 / 1e10) + 0.1)) / 1e10))) / -2.25) + (1e10 + 0.0)) - 1e10)) / 1.0) / 3.14159)) - 1e10))) - v18) + float64(1e10))))
            _ = fi23
        }
        fmt.Println(v18, v6, v6)
        var v24 uint16 = 819
        _ = fi22
        _ = v24
    }
    var v25 int32 = -622
    fmt.Println(v5, v2, v20)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
