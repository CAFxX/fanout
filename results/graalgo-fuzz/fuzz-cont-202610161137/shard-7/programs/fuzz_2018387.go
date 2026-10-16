package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int; F2 int32 }
type GenS2 struct { F0 int16 }

func main() {
    var v1 uint32 = 417
    v2 := make(map[int]uint32)
    v3 := make(map[string]string)
    var v4 uint8 = 128
    var v5 uint16 = 2
    for fi6 := 0; fi6 < 9; fi6++ {
        var v7 uint32 = v1
        var v8 int32 = -1
        _ = fi6
        _ = v7
        _ = v8
    }
    v9 := GenS1{F0: -2147483648, F1: 0, F2: -38}
    var v10 int16 = -785
    v10 = v10
    switch (("foo bar" + "hello") + "hello") {
    case "a":
        var v11 float64 = 0.1
        _ = v11
    case "\n\t":
        var v12 float32 = 0.1
        _ = v12
    default:
        var v13 int = -716
        _ = v13
    }
    v14 := []int{968, 910, 0}
    var v15 int8 = 66
    fmt.Println(v5, v5)
    for fi16 := 0; fi16 < 11; fi16++ {
        fmt.Println(v10, v1)
        _ = fi16
    }
    v17 := v2[int(v4)]
    for fi18 := 0; fi18 < 7; fi18++ {
        var v19 float32 = (0.1 + 0.1)
        _ = fi18
        _ = v19
    }
    fmt.Println(v4, v17, v15)
    if (127 < 32767) {
        for fi20 := 0; fi20 < 13; fi20++ {
            for fi21 := 0; fi21 < 4; fi21++ {
                _ = len(v3)
                v17 = v17
                v15 += v15
                fmt.Println(v10)
                _ = fi21
            }
            _ = fi20
        }
    }
    v22 := GenS2{F0: v10}
    for fi23 := 0; fi23 < 2; fi23++ {
        v14 = append(v14, (fi23 << ((854) & 63)))
        v5 += (v5 / ((1) | 1))
        _ = fi23
    }
    var v24 int16 = 32767
    v25 := GenS2{F0: -274}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v17
    _ = v2
    _ = v22
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
