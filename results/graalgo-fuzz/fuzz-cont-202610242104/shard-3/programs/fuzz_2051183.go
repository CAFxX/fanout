package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8; F2 int16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint8; F1 uint }
type GenS4 struct { F0 bool; F1 int8 }
func genFn0() uint {
    switch "foo bar" {
    case "hello":
        v4 := GenS1{F0: 2147483647, F1: 6, F2: 42}
        _ = v4
    case "foo bar":
        var v5 uint8 = 176
        _ = v5
    }
    var v6 int32 = 0
    switch 9223372036854775807 {
    case -9223372036854775808:
        switch 942 {
        case 32767:
            v6 -= (v6 ^ v6)
        case -9223372036854775808:
            switch "" {
            case "foo bar":
                fmt.Println(v6)
            }
        default:
            v7 := []int{-150, -1}
            _ = v7
        }
    case 1:
        var v8 bool = false
        _ = v8
    case 42:
        fmt.Println(v6)
    }
    _ = v6
    return 255
}
func genFn1(p0 uint16, p1 uint64, p2 uint16) int {
    for fi9 := 0; fi9 < 13; fi9++ {
        if (3.14159 != 1.5) {
            v10 := []int{fi9, 0}
            fmt.Println(fi9, fi9)
            _ = v10
        }
        v11 := []int{}
        for fi12 := 0; fi12 < 20; fi12++ {
            var v13 string = ("\n\t" + "a")
            if false {
                v14 := GenS2{F0: fi9}
                v15 := GenS3{F0: 128, F1: 1}
                _ = v14
                _ = v15
            }
            v11 = append(v11, fi9)
            _ = fi12
            _ = v13
        }
        _ = fi9
        _ = v11
    }
    for fi16 := 0; fi16 < 2; fi16++ {
        p2 += p2
        p1 = p1
        _ = fi16
    }
    var v17 float64 = 0.1
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    return 32768
}

func main() {
    v1 := []int{-87, 32767}
    var v2 uint16 = 2
    var v3 int64 = -199
    var v18 uint8 = 109
    var v19 uint32 = 1
    fmt.Println(v18, v18, v2)
    v20 := GenS3{F0: (^(v18)), F1: 18446744073709551615}
    for fi21 := 0; fi21 < 14; fi21++ {
        var v22 int16 = 659
        var v23 int32 = 819
        if len(v1) > 0 {
            v1[0] = (fi21 - genFn1(2, 0, (^(v2))))
        }
        _ = fi21
        _ = v22
        _ = v23
    }
    fmt.Println(v3, v3)
    _ = len(v1)
    v18 = (v18 & (^(v18)))
    var v24 uint16 = (v2 - v2)
    for fi25 := 0; fi25 < 19; fi25++ {
        switch -2147483648 {
        case 800:
            v18 -= 1
        case 994:
            fmt.Println(fi25, v24)
        case 0:
            fmt.Println(v19, v18, v18)
        default:
            v26 := GenS2{F0: fi25}
            _ = v26
        }
        var v27 uint32 = 2147483648
        if ((v18 / ((v18) | 1)) >= 255) {
            var v28 float32 = -2.25
            _ = v28
        }
        _ = fi25
        _ = v27
    }
    for fi29 := 0; fi29 < 6; fi29++ {
        var v30 uint = 0
        _ = fi29
        _ = v30
    }
    v31 := GenS4{F0: ((v3 & (v3 * v3)) > -1), F1: 127}
    switch 91 {
    case 9223372036854775807:
        fmt.Println(v18, v19)
    default:
        var v32 int = 114
        _ = v32
    }
    fmt.Println(v2, v18, v19)
    v19 = 0
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v24
    _ = v3
    _ = v31
}
