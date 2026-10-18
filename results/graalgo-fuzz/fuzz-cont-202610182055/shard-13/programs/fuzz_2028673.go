package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64; F2 uint32 }
type GenS2 struct { F0 int32; F1 int; F2 int }
type GenS3 struct { F0 int; F1 uint }
func genFn0(p0 uint32, p1 int16) int16 {
    var v4 string = "hello"
    for fi5 := 0; fi5 < 13; fi5++ {
        switch "foo bar" {
        case "foo bar":
            if (0 < 65535) {
                fmt.Println(p1, fi5)
            } else {
                v6 := []int{}
                _ = v6
            }
        case "a":
            p0 += 1
        case "":
            fmt.Println(p1, p1, p0)
        }
        for fi7 := 0; fi7 < 18; fi7++ {
            fmt.Println(v4)
            _ = fi7
        }
        _ = fi5
    }
    _ = p0
    _ = p1
    _ = v4
    return p1
}
func genFn1(p0 int16, p1 int8, p2 int32) bool {
    var v8 float64 = (0.1 / float64(1.5))
    fmt.Println(p2, p2)
    var v9 uint = 720
    v8 = float64((-2.25 + (1.5 + -2.25)))
    v10 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v8
    _ = v9
    return false
}

func main() {
    var v1 int8 = -1
    var v2 uint8 = 253
    var v3 uint32 = 42
    v11 := []int{-2147483648, 4294967295}
    v11 = append(v11, 0)
    v2 += (v2 ^ (v2 | 127))
    if true {
        if (genFn0(v3, 1) <= 1000) {
            var v12 int16 = 1000
            v13 := GenS1{F0: 797, F1: 170, F2: (v3 & v3)}
            v11 = append(v11, -9223372036854775808)
            _ = v12
            _ = v13
        }
        var v14 int = 32768
        _ = v14
    }
    if len(v11) > 0 {
        v11[0] = 594
    }
    fmt.Println(v3, v3, v2)
    var v15 float64 = 0.0
    v16 := GenS2{F0: -914, F1: -2147483648, F2: 127}
    v17 := GenS1{F0: 32767, F1: 9223372036854775807, F2: v3}
    if len(v11) >= 2 {
        v11 = v11[1:len(v11)]
    }
    switch -252 {
    case 206:
        if (true || (33 <= -9223372036854775808)) {
            var v18 float64 = 0.1
            fmt.Println(v15)
            _ = v18
        } else {
            fmt.Println(v3, v2, v2)
            if len(v11) > 0 {
                v11[0] = 0
            }
        }
    case 842:
        var v19 float64 = 0.1
        _ = v19
    case -206:
        var v20 int16 = 32767
        _ = v20
    default:
        var v21 float32 = float32(float64((float32(float64(0.1)) / ((3.14159 + -2.25) / 3.14159))))
        _ = v21
    }
    for fi22 := 0; fi22 < 9; fi22++ {
        if len(v11) >= 2 {
            v11 = v11[1:len(v11)]
        }
        _ = fi22
    }
    v23 := GenS3{F0: -377, F1: 810}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v23
    _ = v3
}
