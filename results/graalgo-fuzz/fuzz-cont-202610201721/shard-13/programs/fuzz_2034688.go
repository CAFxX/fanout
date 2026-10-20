package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int64 }
func genFn0() int16 {
    v6 := []int{}
    if false {
        var v7 string = "a"
        _ = v7
    }
    for fi8 := 0; fi8 < 17; fi8++ {
        v9 := GenS1{F0: 99}
        var v10 int32 = 0
        _ = fi8
        _ = v10
        _ = v9
    }
    var v11 int32 = -1
    _ = v11
    _ = v6
    return -187
}

func main() {
    var v1 float32 = ((-2.25 / 1.5) / (float32(0.1) + 0.0))
    v2 := make(map[int]bool)
    var v3 bool = (!false)
    v4 := []int{128}
    var v5 uint16 = 193
    v5 = 32767
    v12 := GenS1{F0: 100}
    for fi13 := 0; fi13 < 14; fi13++ {
        if true {
            var v14 int8 = 0
            v5 += v5
            _ = v14
        } else {
            var v15 string = ""
            var v16 uint8 = 239
            _ = v15
            _ = v16
        }
        switch fi13 {
        case -474:
            for fi17 := 0; fi17 < 2; fi17++ {
                v4 = append(v4, fi17)
                v5 = v5
                _ = fi17
            }
        }
        _ = fi13
    }
    v18 := v2[-379]
    switch 834 {
    case 127:
        v5 = v5
    case -2147483648:
        v1 = 0.0
    default:
        for fi19 := 0; fi19 < 15; fi19++ {
            var v20 int = -9223372036854775808
            var v21 int8 = 100
            v21 += 48
            _ = fi19
            _ = v20
            _ = v21
        }
    }
    v22 := GenS1{F0: -128}
    var v23 float64 = 3.14159
    v24 := GenS2{F0: int64(v5)}
    if v3 {
        for fi25 := 0; fi25 < 7; fi25++ {
            for fi26 := 0; fi26 < 14; fi26++ {
                v27 := GenS2{F0: -1}
                _ = fi26
                _ = v27
            }
            _ = fi25
        }
        v28 := GenS1{F0: 100}
        _ = v28
    }
    v5 = (v5 * v5)
    var v29 int16 = -165
    v5 = v5
    var v30 int64 = -9223372036854775808
    if (v18 || (65535 < 139)) {
        fmt.Println(v1, v30, v23)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v18
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
