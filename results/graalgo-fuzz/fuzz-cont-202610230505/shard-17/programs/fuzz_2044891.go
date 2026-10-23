package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8 }
type GenS2 struct { F0 uint32 }
func genFn0(p0 uint32) bool {
    v5 := []int{-464, -1, -979, 971}
    switch 127 {
    case -2147483648:
        fmt.Println(p0)
    case 0:
        if (128 >= 104) {
            fmt.Println(p0)
        }
    case 128:
        p0 = p0
    default:
        switch ("" + "a") {
        case "\n\t":
            if (65535 > 2) {
                var v6 int64 = -9223372036854775808
                _ = v6
            }
        }
    }
    var v7 float32 = (0.1 * 1e10)
    var v8 int16 = 731
    _ = p0
    _ = v5
    _ = v7
    _ = v8
    return (("hello" != "\n\t") && (255 < 1))
}
func genFn1(p0 uint32, p1 string) uint {
    v9 := GenS1{F0: -1, F1: 68}
    p0 += (^(p0))
    v10 := []int{128, 128, -867, -1}
    fmt.Println(p1, p0)
    _ = len(v10)
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return 0
}

func main() {
    var v1 uint = 893
    var v2 uint8 = 75
    v3 := make(map[string]string)
    var v4 float32 = 0.1
    var v11 uint8 = 1
    switch 193 {
    case 4294967295:
        for fi12 := 0; fi12 < 16; fi12++ {
            if false {
                for fi13 := 0; fi13 < 1; fi13++ {
                    v3["hello"] = ("a" + "")
                    v14 := GenS1{F0: 0, F1: 127}
                    _ = fi13
                    _ = v14
                }
                v15 := GenS2{F0: 2}
                fmt.Println(fi12, v11)
                _ = v15
            }
            switch ("" + "foo bar") {
            case "\n\t":
                var v16 int8 = -128
                _ = v16
            case "":
                v17 := GenS1{F0: 42, F1: 1}
                _ = v17
            case "hello":
                v18 := GenS2{F0: 4294967295}
                _ = v18
            }
            _ = fi12
        }
    default:
        fmt.Println(v1, v1)
    }
    v4 += (float32(float64((v4 + 3.14159))) / v4)
    v19 := []int{}
    var v20 float64 = 1e10
    for fi21 := 0; fi21 < 6; fi21++ {
        v3["hello"] = "hello"
        fmt.Println(v2, v2)
        _ = fi21
    }
    var v22 int16 = 0
    if false {
        v22 += (v22 | v22)
        for fi23 := 0; fi23 < 10; fi23++ {
            var v24 int32 = -268
            var v25 int32 = (v24 & v24)
            var v26 uint32 = 2
            _ = fi23
            _ = v24
            _ = v25
            _ = v26
        }
        v2 = (v2 * (v11 << ((v1) & 63)))
    } else {
        var v27 string = "hello"
        var v28 int32 = 1
        _ = v27
        _ = v28
    }
    switch 42 {
    case -274:
        if len(v19) > 0 {
            v19[0] = 1
        }
    case 32768:
        var v29 uint64 = 1
        _ = v29
    default:
        v1 += genFn1(771, "hello")
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v19
    _ = v2
    _ = v20
    _ = v22
    _ = v3
    _ = v4
}
