package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64; F2 int64 }
type GenS2 struct { F0 int; F1 uint }
func genFn0(p0 int32) int32 {
    var v6 float64 = 1.5
    switch (("hello" + "a") + "a") {
    case "\n\t":
        if (((((float64(3.14159) + (0.0 + (1.5 / (3.14159 + v6)))) != float64((1.5 + ((0.0 * 1.5) * 1.5)))) || false) && true) && true) {
            v7 := []int{-301, -1, -1}
            v6 = float64(1e10)
            _ = v7
        }
    }
    var v8 uint = 2
    _ = p0
    _ = v6
    _ = v8
    return p0
}

func main() {
    var v1 uint8 = 128
    var v2 uint32 = 2147483648
    v3 := make(map[string]string)
    var v4 uint16 = uint16(v2)
    v5 := make(map[int]uint32)
    if false {
        for fi9 := 0; fi9 < 3; fi9++ {
            for fi10 := 0; fi10 < 20; fi10++ {
                switch ("" + "\n\t") {
                case "\n\t":
                    v2 = v2
                default:
                    switch fi10 {
                    case -377:
                        var v11 float64 = 0.0
                        _ = v11
                    case -9223372036854775808:
                        v4 -= 42
                    }
                }
                v1 = (v1 + 239)
                _ = fi10
            }
            v2 = v2
            v12 := GenS1{F0: 42, F1: 1, F2: -1}
            _ = fi9
            _ = v12
        }
    } else {
        v2 = 65535
    }
    v13 := []int{32767, 995}
    fmt.Println(v1, v4, v4)
    v14 := GenS1{F0: 42, F1: 0, F2: 506}
    if (!(-115 <= 100)) {
        v4 -= (^(v4))
        var v15 uint64 = 0
        _ = v15
    } else {
        var v16 uint64 = 2
        for fi17 := 0; fi17 < 15; fi17++ {
            if (!false) {
                var v18 int64 = -9223372036854775808
                var v19 uint8 = v1
                var v20 int64 = v18
                _ = v18
                _ = v19
                _ = v20
            } else {
                var v21 uint16 = (v4 * (v4 >> ((790) & 63)))
                var v22 int64 = 0
                v23 := GenS2{F0: 0, F1: 1}
                _ = v21
                _ = v22
                _ = v23
            }
            _ = fi17
        }
        _ = v16
    }
    _ = len(v5)
    v2 = v2
    fmt.Println(v4, v2, v4)
    for fi24 := 0; fi24 < 2; fi24++ {
        fmt.Println(fi24, v4)
        for fi25 := 0; fi25 < 18; fi25++ {
            fmt.Println(fi24, fi24, v2)
            v4 = 685
            v2 -= (^(v2))
            _ = fi25
        }
        _ = fi24
    }
    var v26 float64 = -2.25
    if true {
        v27 := GenS2{F0: 128, F1: uint(v1)}
        v13 = append(v13, -790)
        _ = v27
    }
    var v28 bool = (0.1 == float32(((0.0 / 1e10) * (v26 - ((v26 / ((v26 * v26) / 3.14159)) - v26)))))
    switch 42 {
    case -1:
        fmt.Println(v1, v26)
    case 127:
        var v29 bool = (true && v28)
        _ = v29
    default:
        v30 := GenS1{F0: 32767, F1: 19, F2: -394}
        _ = v30
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v26
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
