package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int32; F1 string }
func genFn0(p0 string) uint8 {
    p0 = "hello"
    var v5 int32 = 1
    switch -802 {
    case 127:
        var v6 string = p0
        _ = v6
    case 471:
        var v7 int8 = 11
        _ = v7
    default:
        var v8 uint64 = 4294967295
        _ = v8
    }
    _ = p0
    _ = v5
    return 10
}
func genFn1() int16 {
    v9 := []int{0}
    var v10 int = 1
    for fi11 := 0; fi11 < 5; fi11++ {
        v10 += (v10 - 351)
        v10 = -1
        if false {
            v12 := GenS1{F0: (!(((0.0 * ((((0.1 * float32((0.1 + 1e10))) + (-2.25 - -2.25)) / (3.14159 - (3.14159 - -2.25))) / 1.0)) * 0.0) >= -2.25)), F1: 1}
            for fi13 := 0; fi13 < 12; fi13++ {
                fmt.Println(v10, v10, fi13)
                v10 = 1
                _ = fi13
            }
            _ = v12
        }
        _ = fi11
    }
    fmt.Println(v10)
    _ = v10
    _ = v9
    return 454
}

func main() {
    var v1 int32 = -751
    var v2 float32 = (float32(float64(((1e10 - 0.0) - float32(0.1)))) + ((float32(1.5) - 1.5) - (1e10 + (((-2.25 * (3.14159 * ((0.0 + 1.5) / float32(((-2.25 + 1.5) / 1.5))))) - 0.0) * float32(3.14159)))))
    var v3 int = 128
    var v4 uint16 = 274
    var v14 int8 = 73
    v1 += v1
    v15 := []int{int(v4), 127}
    for fi16 := 0; fi16 < 13; fi16++ {
        for fi17 := 0; fi17 < 7; fi17++ {
            _ = len(v15)
            switch "a" {
            case "a":
                switch (fi16 ^ (fi17 << ((255) & 63))) {
                case 128:
                    v1 = (v1 >> ((0) & 63))
                case 9223372036854775807:
                    v18 := GenS2{F0: (^(fi16))}
                    _ = v18
                case -1:
                    v19 := GenS2{F0: v3}
                    _ = v19
                }
                v20 := GenS3{F0: -2147483648}
                _ = v20
            }
            v15 = append(v15, -1)
            _ = fi17
        }
        _ = fi16
    }
    if (4294967295 > 1) {
        v21 := GenS4{F0: v1, F1: "a"}
        v2 -= (v2 / v2)
        var v22 int32 = v1
        _ = v21
        _ = v22
    } else {
        var v23 int8 = (^(v14))
        _ = v23
    }
    v24 := GenS4{F0: v1, F1: "hello"}
    for fi25 := 0; fi25 < 8; fi25++ {
        v14 = (v14 >> ((0) & 63))
        fmt.Println(v4, v1)
        _ = fi25
    }
    v26 := GenS4{F0: v1, F1: ((("a" + ("foo bar" + "\n\t")) + "hello") + "foo bar")}
    switch (-(v3)) {
    case -843:
        var v27 float64 = 1e10
        _ = v27
    case -1:
        fmt.Println(v4)
    case 2147483647:
        v1 = (v1 % ((1) | 1))
    }
    v3 = v3
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v24
    _ = v26
    _ = v3
    _ = v4
}
