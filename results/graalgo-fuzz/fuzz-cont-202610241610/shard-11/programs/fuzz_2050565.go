package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string; F2 int64 }
type GenS2 struct { F0 bool; F1 int }
type GenS3 struct { F0 bool; F1 uint8 }

func main() {
    var v1 float32 = float32(((((3.14159 + 3.14159) + float64(1.5)) * (((0.0 - (((1.5 - 3.14159) * ((3.14159 + ((3.14159 / float64(0.1)) * ((0.0 + float64((3.14159 + -2.25))) / 1.0))) - (1e10 / 1.0))) + 0.0)) * 0.0) - (float64(1.5) / 1.0))) * (0.1 / 1.5)))
    var v2 int16 = 42
    var v3 float32 = 0.0
    if false {
        v1 = -2.25
        v4 := []int{32768, 2147483647, int(v2), 9223372036854775807}
        var v5 int32 = -1
        _ = v4
        _ = v5
    }
    var v6 int32 = -2147483648
    v7 := GenS1{F0: 0, F1: "a", F2: 0}
    for fi8 := 0; fi8 < 18; fi8++ {
        v3 += (float32(-2.25) / 1.0)
        switch fi8 {
        case -1:
            v9 := GenS1{F0: fi8, F1: "", F2: 1}
            _ = v9
        case 127:
            if false {
                fmt.Println(v6, v3)
                var v10 int = fi8
                _ = v10
            } else {
                v1 += v3
            }
        case 4294967295:
            var v11 int16 = (v2 ^ v2)
            _ = v11
        default:
            v12 := []int{}
            _ = v12
        }
        _ = fi8
    }
    v13 := []int{4294967295, 291, 9223372036854775807}
    if false {
        var v14 float64 = 1.5
        fmt.Println(v6, v3, v6)
        _ = v14
    } else {
        var v15 int64 = 640
        v15 = v15
        var v16 uint32 = 65535
        _ = v15
        _ = v16
    }
    v17 := GenS1{F0: 0, F1: "hello", F2: -1}
    if false {
        v6 -= v6
    }
    var v18 uint16 = 255
    for fi19 := 0; fi19 < 11; fi19++ {
        var v20 uint64 = 0
        fmt.Println(v6, v3, v2)
        if (-1 <= 127) {
            var v21 int32 = (v6 / ((v6) | 1))
            _ = len(v13)
            _ = v21
        } else {
            var v22 bool = (!(255 <= 2))
            _ = v22
        }
        _ = fi19
        _ = v20
    }
    var v23 int64 = 329
    v24 := GenS1{F0: 171, F1: "foo bar", F2: -1}
    v25 := GenS2{F0: true, F1: 783}
    for fi26 := 0; fi26 < 7; fi26++ {
        if len(v13) >= 2 {
            v13 = v13[1:len(v13)]
        }
        for fi27 := 0; fi27 < 15; fi27++ {
            if len(v13) > 0 {
                v13[0] = fi27
            }
            v2 -= v2
            _ = fi27
        }
        if ((("a" + ("\n\t" + "")) != "\n\t") || (false && true)) {
            v28 := GenS3{F0: (0 < v6), F1: 183}
            _ = v28
        }
        _ = fi26
    }
    var v29 int8 = -128
    fmt.Println(v23)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v17
    _ = v18
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v29
    _ = v3
    _ = v6
    _ = v7
}
