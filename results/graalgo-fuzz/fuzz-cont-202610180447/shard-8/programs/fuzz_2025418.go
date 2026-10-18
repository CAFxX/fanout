package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16; F2 uint }
type GenS2 struct { F0 uint16 }
func genFn0(p0 int64) int8 {
    fmt.Println(p0)
    for fi7 := 0; fi7 < 17; fi7++ {
        var v8 int = -2147483648
        var v9 uint = 1
        fmt.Println(v9)
        _ = fi7
        _ = v8
        _ = v9
    }
    v10 := GenS1{F0: true, F1: 32767, F2: 56}
    switch -1 {
    case 32768:
        for fi11 := 0; fi11 < 2; fi11++ {
            var v12 float32 = 0.0
            _ = fi11
            _ = v12
        }
    }
    _ = p0
    _ = v10
    return 100
}
func genFn1(p0 uint8, p1 int32, p2 uint16) int {
    var v13 int16 = int16(p0)
    for fi14 := 0; fi14 < 17; fi14++ {
        fmt.Println(fi14)
        _ = fi14
    }
    var v15 uint8 = p0
    var v16 uint = 42
    fmt.Println(v13, p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v15
    _ = v16
    return 0
}

func main() {
    var v1 uint16 = 255
    v2 := make(map[int]uint32)
    v3 := make(map[string]string)
    var v4 bool = (!true)
    var v5 int64 = 9223372036854775807
    var v6 string = ""
    for fi17 := 0; fi17 < 17; fi17++ {
        v18 := GenS1{F0: v4, F1: (^(v1)), F2: 255}
        var v19 uint16 = v1
        var v20 uint8 = 127
        _ = fi17
        _ = v18
        _ = v19
        _ = v20
    }
    v4 = (1 > 42)
    v21 := []int{}
    v4 = true
    for fi22 := 0; fi22 < 1; fi22++ {
        delete(v3, v6)
        var v23 uint8 = 255
        _ = fi22
        _ = v23
    }
    var v24 string = v6
    var v25 int = 9223372036854775807
    if v4 {
        switch genFn1(1, -1, v1) {
        case 128:
            if (v25 <= 42) {
                var v26 bool = ("hello" < "\n\t")
                _ = v26
            }
        case 32767:
            v27 := GenS2{F0: (v1 + (v1 / (((v1 * (v1 << ((65535) & 63)))) | 1)))}
            _ = v27
        default:
            v1 += (v1 + (v1 - v1))
        }
    }
    v28 := GenS2{F0: (v1 - (^(v1)))}
    var v29 int64 = (v5 / (((v5 >> ((262) & 63))) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v24
    _ = v25
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
