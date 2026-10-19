package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int64; F1 uint32 }
func genFn0() int32 {
    var v5 uint8 = 143
    v5 -= (^(v5))
    for fi6 := 0; fi6 < 7; fi6++ {
        fmt.Println(fi6, v5)
        v5 += 2
        var v7 int8 = -1
        _ = fi6
        _ = v7
    }
    v5 = v5
    v8 := []int{-2147483648, 2147483647, 9223372036854775807, 1}
    _ = v5
    _ = v8
    return 0
}

func main() {
    var v1 int = 128
    var v2 int64 = -1
    v3 := make(map[int]bool)
    v4 := make(map[int]uint32)
    v1 -= 1
    switch "foo bar" {
    case "\n\t":
        v1 -= -9223372036854775808
    case "hello":
        for fi9 := 0; fi9 < 10; fi9++ {
            switch v1 {
            case 9223372036854775807:
                if false {
                    var v10 bool = false
                    v10 = (v10 || v10)
                    _ = v10
                } else {
                    v1 += (v1 & (fi9 + (v1 / ((-9223372036854775808) | 1))))
                    var v11 int64 = v2
                    _ = v11
                }
            case 4294967295:
                v2 += (v2 - v2)
            case 128:
                fmt.Println(fi9, fi9, fi9)
            }
            v1 = (v1 << ((4294967295) & 63))
            _ = fi9
        }
    case "":
        v12 := v3[(v1 ^ 301)]
        _ = v12
    default:
        for fi13 := 0; fi13 < 1; fi13++ {
            switch "" {
            case "a":
                var v14 int = -9223372036854775808
                _ = v14
            case "":
                v2 = v2
            case "\n\t":
                var v15 uint16 = 32767
                _ = v15
            }
            var v16 int = -562
            v17 := []int{}
            _ = fi13
            _ = v16
            _ = v17
        }
    }
    fmt.Println(v2)
    v18 := []int{(v1 + (-(v1)))}
    if len(v18) >= 2 {
        v18 = v18[1:len(v18)]
    }
    v1 -= v1
    for fi19 := 0; fi19 < 11; fi19++ {
        v4[42] = 43
        v20 := GenS1{F0: 128}
        var v21 string = ""
        _ = fi19
        _ = v20
        _ = v21
    }
    if (true || (!(!(true && false)))) {
        v22 := GenS2{F0: (v2 << ((255) & 63)), F1: 0}
        if false {
            var v23 uint16 = 32767
            v1 -= 252
            _ = v23
        } else {
            v1 = (v1 >> ((176) & 63))
            if true {
                fmt.Println(v1)
                if (0 <= 1) {
                    var v24 float64 = 1e10
                    _ = v24
                } else {
                    var v25 bool = true
                    _ = v25
                }
            }
        }
        _ = v22
    }
    var v26 bool = true
    v26 = ((65535 != 2) || (420 == 42))
    for fi27 := 0; fi27 < 4; fi27++ {
        var v28 bool = (2 == 244)
        if (v2 != v2) {
            var v29 float32 = (float32((((float64(1.5) - ((-2.25 - ((-2.25 + 1.5) - 0.0)) * 1.5)) * (3.14159 - 1.5)) + ((float64((3.14159 + (((1.5 / (-2.25 - 1.5)) / -2.25) + float32(float64(1.5))))) / ((float64((1.5 - 0.0)) * float64((1.5 + float32(-2.25)))) * 3.14159)) + 3.14159))) - float32((1.5 - (1e10 + 0.1))))
            v29 -= v29
            _ = v29
        }
        _ = fi27
        _ = v28
    }
    fmt.Println(v2, v2, v1)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v18
    _ = v2
    _ = v26
    _ = v3
    _ = v4
}
