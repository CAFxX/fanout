package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32; F2 uint8 }
type GenS2 struct { F0 int8; F1 int16 }
type GenS3 struct { F0 int32; F1 int8 }
type GenS4 struct { F0 int8; F1 uint }
func genFn0() int8 {
    v4 := GenS1{F0: 32767, F1: -718, F2: 162}
    var v5 uint16 = 947
    v5 -= 32767
    _ = v4
    _ = v5
    return 0
}
func genFn1(p0 uint32) string {
    v6 := []int{}
    var v7 uint64 = 18446744073709551615
    _ = p0
    _ = v6
    _ = v7
    return ""
}

func main() {
    v1 := make(map[int]uint32)
    var v2 uint16 = 0
    var v3 float64 = float64(3.14159)
    v2 = (v2 - 42)
    v8 := GenS1{F0: 42, F1: -418, F2: 42}
    for fi9 := 0; fi9 < 5; fi9++ {
        var v10 uint16 = (v2 >> ((42) & 63))
        v2 = (^(v10))
        _ = fi9
        _ = v10
    }
    if (!false) {
        delete(v1, 1)
        v11 := GenS1{F0: -32768, F1: -2147483648, F2: 42}
        _ = v11
    }
    switch "\n\t" {
    case "foo bar":
        v3 -= v3
    case "\n\t":
        if (42 > 691) {
            switch "hello" {
            case "foo bar":
                v3 = (v3 * ((v3 + 3.14159) * (((v3 * v3) * 1e10) * float64(3.14159))))
            default:
                v12 := GenS2{F0: 0, F1: 1000}
                _ = v12
            }
        } else {
            for fi13 := 0; fi13 < 4; fi13++ {
                if ((4294967295 == 18446744073709551615) && true) {
                    var v14 int8 = 127
                    fmt.Println(fi13)
                    for fi15 := 0; fi15 < 17; fi15++ {
                        v3 = (1e10 + 0.1)
                        fmt.Println(v3, fi15, fi15)
                        var v16 float32 = 0.0
                        if true {
                            var v17 uint16 = 65535
                            v14 = (v14 & 127)
                            var v18 uint64 = 42
                            var v19 int16 = 1000
                            _ = v17
                            _ = v18
                            _ = v19
                        } else {
                            v3 += ((v3 / (1.5 / v3)) / 1.0)
                            if true {
                                delete(v1, fi15)
                            }
                        }
                        v14 += (v14 * (v14 / ((v14) | 1)))
                        _ = fi15
                        _ = v16
                    }
                    _ = v14
                }
                _ = fi13
            }
            if true {
                v1[127] = 570
            } else {
                for fi20 := 0; fi20 < 2; fi20++ {
                    v3 += (3.14159 + (float64(0.0) / v3))
                    v21 := GenS2{F0: -128, F1: 774}
                    _ = fi20
                    _ = v21
                }
                v3 = v3
                fmt.Println(v3, v2)
            }
        }
    case "hello":
        var v22 int = 128
        _ = v22
    }
    v3 += (3.14159 / 1e10)
    var v23 float64 = float64(0.1)
    if (!(!((v2 ^ 0) == (v2 % ((v2) | 1))))) {
        var v24 uint64 = 1
        v25 := GenS3{F0: 0, F1: genFn0()}
        _ = v24
        _ = v25
    } else {
        var v26 int32 = 42
        _ = v26
    }
    if (3.14159 != 0.1) {
        v27 := GenS4{F0: 1, F1: 255}
        var v28 uint8 = 127
        _ = v27
        _ = v28
    } else {
        var v29 int8 = -41
        v29 += (v29 - genFn0())
        _ = v29
    }
    switch (("foo bar" + ("hello" + "")) + "\n\t") {
    case "foo bar":
        if true {
            for fi30 := 0; fi30 < 17; fi30++ {
                var v31 string = ("" + ("" + ""))
                _ = fi30
                _ = v31
            }
            var v32 bool = false
            _ = v32
        } else {
            v2 -= (^(v2))
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v8
}
