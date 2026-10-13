package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint8; F1 int }
type GenS3 struct { F0 uint16; F1 uint32; F2 uint32 }
type GenS4 struct { F0 int32; F1 int }
type GenS5 struct { F0 int32; F1 int64; F2 uint16 }
func genFn0(p0 uint32, p1 int) int32 {
    v5 := GenS1{F0: 219}
    v6 := []int{42}
    switch p1 {
    case -568:
        var v7 string = ""
        _ = v7
    case 341:
        v8 := GenS2{F0: 1, F1: (p1 >> ((286) & 63))}
        _ = v8
    case 2147483647:
        p0 = 875
    default:
        var v9 uint = 2
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 0
}

func main() {
    var v1 uint64 = 4294967295
    var v2 uint8 = 255
    var v3 int8 = -1
    v4 := make(map[string]uint32)
    for fi10 := 0; fi10 < 18; fi10++ {
        v11 := GenS1{F0: 2}
        _ = fi10
        _ = v11
    }
    var v12 int8 = 106
    switch "\n\t" {
    case "\n\t":
        v13 := GenS2{F0: (v2 - 127), F1: 127}
        _ = v13
    default:
        switch ("a" + "") {
        case "a":
            fmt.Println(v1, v1)
        default:
            fmt.Println(v12, v2, v2)
        }
    }
    delete(v4, ((("" + (("" + "a") + ("\n\t" + "foo bar"))) + "a") + "foo bar"))
    v14 := GenS3{F0: uint16(v1), F1: uint32(v1), F2: 65535}
    v2 = v2
    for fi15 := 0; fi15 < 19; fi15++ {
        var v16 int64 = 1
        fmt.Println(v12, fi15)
        _ = fi15
        _ = v16
    }
    switch -1 {
    case 103:
        for fi17 := 0; fi17 < 4; fi17++ {
            if false {
                v1 = v1
            }
            v12 = (v12 | (-(v3)))
            _ = fi17
        }
    }
    v18 := []int{-1, 9223372036854775807}
    var v19 uint8 = v2
    v19 -= (v2 % ((2) | 1))
    v20 := GenS4{F0: 1, F1: -2147483648}
    v1 += (^(v1))
    fmt.Println(v3, v3)
    for fi21 := 0; fi21 < 4; fi21++ {
        v3 = v3
        v22 := GenS5{F0: -684, F1: -830, F2: 1}
        if (-9223372036854775808 <= -1) {
            var v23 float64 = 1.5
            _ = v23
        }
        _ = fi21
        _ = v22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v12
    _ = v14
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
