package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32; F2 uint }

func main() {
    var v1 float32 = -2.25
    v2 := make(map[string]string)
    var v3 float64 = 3.14159
    var v4 uint8 = 42
    v1 = v1
    v3 += (float64(v1) + v3)
    switch 128 {
    case 18:
        if false {
            delete(v2, "\n\t")
            v5 := GenS1{F0: (1 > 797), F1: uint32(v4), F2: 0}
            for fi6 := 0; fi6 < 9; fi6++ {
                var v7 int64 = 651
                for fi8 := 0; fi8 < 6; fi8++ {
                    var v9 string = "hello"
                    v7 = 310
                    var v10 uint32 = 489
                    _ = fi8
                    _ = v10
                    _ = v9
                }
                v4 = (v4 | (v4 | (^(v4))))
                _ = fi6
                _ = v7
            }
            _ = v5
        }
    default:
        for fi11 := 0; fi11 < 4; fi11++ {
            v3 = float64(v1)
            var v12 uint64 = 18446744073709551615
            v4 += 128
            _ = fi11
            _ = v12
        }
    }
    switch 32768 {
    case 128:
        var v13 uint32 = 0
        _ = v13
    case -435:
        fmt.Println(v3)
    }
    v14 := GenS1{F0: false, F1: 4294967295, F2: 42}
    v15 := []int{1, -864}
    var v16 float32 = (0.0 * (v1 * 3.14159))
    if false {
        var v17 int16 = -32768
        v15 = append(v15, 588)
        _ = v17
    }
    fmt.Println(v16)
    var v18 int8 = 100
    switch -258 {
    case -428:
        delete(v2, "a")
    case 32767:
        v18 = (v18 - (v18 * v18))
    default:
        var v19 uint = 255
        _ = v19
    }
    v15 = append(v15, -33)
    v16 -= float32(1.5)
    var v20 int = 1
    switch "foo bar" {
    case "":
        for fi21 := 0; fi21 < 19; fi21++ {
            v18 += v18
            _ = fi21
        }
    }
    fmt.Println(v1, v4, v16)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v18
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
