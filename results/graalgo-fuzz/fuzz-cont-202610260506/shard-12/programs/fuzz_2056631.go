package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 bool; F1 uint; F2 uint32 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int32; F1 uint }
func genFn0(p0 string, p1 uint) int32 {
    for fi5 := 0; fi5 < 1; fi5++ {
        v6 := GenS1{F0: 2147483647}
        _ = fi5
        _ = v6
    }
    v7 := []int{0, 42, 127}
    for fi8 := 0; fi8 < 18; fi8++ {
        var v9 uint8 = 207
        var v10 uint32 = 0
        if (!false) {
            fmt.Println(p1, fi8, p0)
            p1 += (^(p1))
        }
        _ = fi8
        _ = v10
        _ = v9
    }
    if len(v7) > 0 {
        v7[0] = -171
    }
    if len(v7) > 0 {
        v7[0] = 2147483647
    }
    _ = p0
    _ = p1
    _ = v7
    return 2147483647
}
func genFn1(p0 bool, p1 int16, p2 uint32) int16 {
    var v11 uint = 65535
    fmt.Println(p0, p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return (p1 << ((v11) & 63))
}

func main() {
    var v1 uint64 = 941
    var v2 float32 = 1.5
    v3 := []int{744, 1}
    var v4 int = 42
    if true {
        v2 = float32(float64((((3.14159 - (1e10 / ((0.0 / float32(((1e10 * 3.14159) * -2.25))) - v2))) + float32(-2.25)) + ((v2 - float32(float64(((v2 + float32((float64(v2) + (((((float64(float32(3.14159)) + (0.0 / ((1e10 + (0.1 + (((1e10 * 0.0) / (0.1 + ((3.14159 + 0.0) / 1e10))) - (-2.25 / (((float64(float32((-2.25 - 3.14159))) - (((3.14159 + (3.14159 + 0.1)) - -2.25) + -2.25)) - 1e10) / 1.0))))) + 1e10))) + (1.5 - 1e10)) / float64(v2)) * 3.14159) * ((-2.25 / 1.5) + 3.14159))))) * 1e10)))) - (v2 * 3.14159)))))
        fmt.Println(v1, v4)
        fmt.Println(v4, v4)
    }
    var v12 uint16 = 65535
    v13 := GenS1{F0: 2147483647}
    v1 = v1
    switch 32767 {
    case -1:
        var v14 int16 = -1
        _ = v14
    case 1:
        for fi15 := 0; fi15 < 19; fi15++ {
            if false {
                v3 = append(v3, fi15)
                var v16 float32 = float32(1e10)
                _ = v16
            } else {
                var v17 float32 = v2
                _ = v17
            }
            _ = fi15
        }
    case -638:
        switch "hello" {
        case "\n\t":
            v12 += (^(v12))
        case "":
            fmt.Println(v12, v4, v4)
        default:
            fmt.Println(v12, v4)
        }
    }
    if (false && ("" <= "foo bar")) {
        var v18 int = v4
        _ = v18
    } else {
        v19 := GenS2{F0: (!(42 != 42)), F1: 65535, F2: 2147483648}
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = v19
    }
    for fi20 := 0; fi20 < 13; fi20++ {
        fmt.Println(fi20, v4)
        switch "foo bar" {
        case "foo bar":
            var v21 uint8 = 1
            _ = v21
        case "hello":
            switch ("\n\t" + "foo bar") {
            case "hello":
                _ = len(v3)
            default:
                var v22 uint8 = 42
                _ = v22
            }
        case "":
            if len(v3) > 0 {
                v3[0] = (fi20 << ((78) & 63))
            }
        default:
            v23 := GenS3{F0: int64(v1)}
            _ = v23
        }
        _ = fi20
    }
    if false {
        v24 := GenS2{F0: false, F1: 830, F2: 938}
        v25 := GenS4{F0: 1, F1: 65535}
        var v26 int32 = genFn0("foo bar", 382)
        _ = v24
        _ = v25
        _ = v26
    }
    v12 = (^(v12))
    v4 += v4
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
}
