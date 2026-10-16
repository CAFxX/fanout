package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint16; F1 int64 }
type GenS5 struct { F0 uint8; F1 int64 }
func genFn0(p0 int8) int64 {
    var v4 uint64 = 4294967295
    v5 := []int{-918}
    v6 := GenS1{F0: 128, F1: 0}
    for fi7 := 0; fi7 < 16; fi7++ {
        for fi8 := 0; fi8 < 18; fi8++ {
            fmt.Println(v4, fi8)
            _ = len(v5)
            _ = fi8
        }
        fmt.Println(v4)
        var v9 int16 = -1
        _ = fi7
        _ = v9
    }
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    return 42
}
func genFn1() string {
    var v10 float64 = 0.1
    v10 = 3.14159
    for fi11 := 0; fi11 < 2; fi11++ {
        if (!false) {
            var v12 float32 = 1.5
            for fi13 := 0; fi13 < 16; fi13++ {
                var v14 float32 = 0.1
                _ = fi13
                _ = v14
            }
            _ = v12
        } else {
            var v15 uint16 = 255
            _ = v15
        }
        _ = fi11
    }
    _ = v10
    return ("foo bar" + "\n\t")
}

func main() {
    var v1 int32 = 0
    v2 := make(map[string]uint32)
    v3 := make(map[int]int)
    v16 := GenS1{F0: -2147483648, F1: 100}
    for fi17 := 0; fi17 < 8; fi17++ {
        v18 := GenS2{F0: -142}
        fmt.Println(fi17, v1)
        var v19 int32 = (^(v1))
        _ = fi17
        _ = v18
        _ = v19
    }
    var v20 uint8 = 255
    if ((false && ("foo bar" == "")) || false) {
        var v21 int16 = -32768
        for fi22 := 0; fi22 < 11; fi22++ {
            v21 = 1
            _ = fi22
        }
        _ = v21
    } else {
        fmt.Println(v20)
    }
    v20 -= (v20 % ((127) | 1))
    v1 += (v1 * -557)
    delete(v3, -235)
    switch ("hello" + ("" + "foo bar")) {
    case "\n\t":
        var v23 int16 = 1000
        _ = v23
    default:
        v24 := GenS1{F0: 32768, F1: 0}
        _ = v24
    }
    v25 := GenS3{F0: (v20 - v20)}
    var v26 int = -355
    switch "" {
    case "a":
        v27 := []int{}
        _ = v27
    case "hello":
        if (1.5 <= (-2.25 * -2.25)) {
            if (!true) {
                v26 = 32768
                fmt.Println(v26)
                if (false && (1.5 >= 1.5)) {
                    v20 -= 0
                    var v28 float32 = float32(float64(float32(-2.25)))
                    v20 -= v20
                    _ = v28
                } else {
                    var v29 int32 = (v1 % (((v1 / ((-370) | 1))) | 1))
                    _ = v29
                }
            }
        }
    case "":
        fmt.Println(v1, v26, v1)
    default:
        v1 = (v1 / (((v1 ^ (v1 + (^(v1))))) | 1))
    }
    for fi30 := 0; fi30 < 6; fi30++ {
        if false {
            var v31 uint16 = 2
            _ = v31
        } else {
            v32 := GenS4{F0: 2, F1: 1}
            _ = v32
        }
        var v33 float32 = (float32(1e10) + ((-2.25 + (-2.25 - float32((((1.5 - (0.1 + 3.14159)) * float64(float32((0.0 + (-2.25 + 0.1))))) + 3.14159)))) * 1.5))
        v33 = 0.0
        _ = fi30
        _ = v33
    }
    for fi34 := 0; fi34 < 14; fi34++ {
        v35 := []int{(v26 | v26)}
        _ = fi34
        _ = v35
    }
    for fi36 := 0; fi36 < 18; fi36++ {
        v37 := GenS5{F0: v20, F1: -787}
        v26 += (v26 + fi36)
        v38 := GenS1{F0: 709, F1: 1}
        _ = fi36
        _ = v37
        _ = v38
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v16
    _ = v2
    _ = v20
    _ = v25
    _ = v26
    _ = v3
}
