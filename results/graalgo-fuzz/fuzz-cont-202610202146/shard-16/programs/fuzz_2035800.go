package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint8 }
func genFn0() int8 {
    var v5 int32 = -2147483648
    var v6 uint32 = 395
    v7 := GenS1{F0: 612, F1: 280}
    v8 := GenS2{F0: 2}
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return 127
}

func main() {
    var v1 float32 = (1e10 / 1.0)
    var v2 string = ""
    v3 := make(map[int]int)
    var v4 uint16 = 65535
    v1 = (((v1 / v1) * (v1 * float32((0.1 / 3.14159)))) - ((v1 - v1) * 0.1))
    for fi9 := 0; fi9 < 12; fi9++ {
        var v10 int32 = 251
        _ = fi9
        _ = v10
    }
    v4 += (^(v4))
    v1 += v1
    fmt.Println(v4, v4)
    v4 = 32767
    v11 := []int{-1}
    var v12 int16 = 42
    switch 0 {
    case 731:
        v13 := GenS3{F0: 255}
        _ = v13
    case 2147483647:
        if (true && ((((v1 + v1) * (v1 / v1)) * 1.5) > (-2.25 - (v1 / v1)))) {
            fmt.Println(v4, v4, v12)
        } else {
            for fi14 := 0; fi14 < 15; fi14++ {
                var v15 int8 = 1
                var v16 uint16 = (v4 & v4)
                _ = fi14
                _ = v15
                _ = v16
            }
        }
    case 146:
        v2 = v2
    }
    if false {
        var v17 int16 = 0
        _ = v17
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
}
