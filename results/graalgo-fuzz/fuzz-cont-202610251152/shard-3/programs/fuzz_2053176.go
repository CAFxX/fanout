package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32 }
type GenS2 struct { F0 int32 }
func genFn0(p0 bool) int16 {
    var v4 string = "hello"
    switch v4 {
    case "foo bar":
        switch 127 {
        case 1:
            v5 := []int{4294967295, -2147483648, -764}
            _ = v5
        case 128:
            for fi6 := 0; fi6 < 6; fi6++ {
                p0 = true
                p0 = p0
                v7 := []int{(fi6 & (fi6 % ((fi6) | 1))), 128}
                _ = fi6
                _ = v7
            }
        default:
            if (2 >= 65535) {
                var v8 float32 = 1e10
                if p0 {
                    if (4294967295 >= 65535) {
                        v9 := []int{}
                        v8 -= v8
                        fmt.Println(v4)
                        _ = v9
                    }
                    v10 := GenS1{F0: 65535, F1: 1}
                    v4 = ("foo bar" + v4)
                    _ = v10
                }
                _ = v8
            }
        }
    case "hello":
        v11 := GenS1{F0: 0, F1: 2}
        _ = v11
    case "":
        var v12 int8 = 107
        _ = v12
    }
    _ = p0
    _ = v4
    return 1
}
func genFn1(p0 int64, p1 uint8, p2 uint8) uint16 {
    fmt.Println(p0, p2, p0)
    var v13 uint16 = 2
    v13 = v13
    var v14 uint16 = 0
    var v15 float64 = 0.1
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v14
    _ = v15
    return v14
}

func main() {
    var v1 uint = 460
    var v2 float32 = 3.14159
    v3 := make(map[int]uint32)
    if false {
        v2 = 1.5
        switch "" {
        case "a":
            fmt.Println(v2, v1)
        case "":
            v16 := v3[-66]
            _ = v16
        case "foo bar":
            v1 = v1
        }
        if false {
            var v17 int8 = -128
            v3[0] = 2147483647
            _ = v17
        }
    }
    v18 := GenS1{F0: 1, F1: 981}
    for fi19 := 0; fi19 < 10; fi19++ {
        var v20 string = "\n\t"
        _ = fi19
        _ = v20
    }
    v1 = 381
    if false {
        if (0.0 <= v2) {
            var v21 float32 = (-2.25 / 0.1)
            var v22 float32 = v2
            _ = v21
            _ = v22
        } else {
            fmt.Println(v1, v2)
        }
    }
    v23 := GenS2{F0: -2147483648}
    var v24 int = 9223372036854775807
    for fi25 := 0; fi25 < 2; fi25++ {
        switch "" {
        case "":
            v2 = (1e10 * v2)
        case "foo bar":
            var v26 int16 = genFn0((!true))
            _ = v26
        }
        _ = fi25
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v18
    _ = v2
    _ = v23
    _ = v24
    _ = v3
}
