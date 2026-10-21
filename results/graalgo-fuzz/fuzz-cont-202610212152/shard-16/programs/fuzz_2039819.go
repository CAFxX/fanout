package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16; F2 uint64 }
type GenS2 struct { F0 uint32; F1 string }
type GenS3 struct { F0 bool; F1 uint16 }
type GenS4 struct { F0 int16; F1 int16; F2 int16 }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[uint32]bool)
    var v3 int32 = 2147483647
    var v4 float64 = 0.1
    v5 := []uint32{2147483647, 2147483647, 1, 2147483647}
    var v6 int32 = v3
    if (4294967295 > 711) {
        var v7 int64 = 9223372036854775807
        _ = v7
    } else {
        fmt.Println(v4, v6, v3)
        v8 := v2[0]
        v3 = v3
        _ = v8
    }
    if (!((1e10 != (-2.25 + 0.0)) || ((false || ((float32(v4) - float32(((((v4 - 0.1) + v4) / 3.14159) * ((((v4 * v4) - 0.0) * 0.0) - (1.5 * (v4 / ((1.5 / -2.25) / (v4 + float64(float32((v4 * (3.14159 - float64(1e10))))))))))))) >= 0.0)) && true))) {
        switch 9223372036854775807 {
        case 127:
            for fi9 := 0; fi9 < 5; fi9++ {
                for fi10 := 0; fi10 < 11; fi10++ {
                    var v11 float64 = v4
                    var v12 bool = false
                    v13 := GenS1{F0: 32767, F1: 32767, F2: 421}
                    _ = fi10
                    _ = v11
                    _ = v12
                    _ = v13
                }
                _ = fi9
            }
        case -772:
            v14 := GenS2{F0: 114, F1: ""}
            _ = v14
        }
    } else {
        fmt.Println(v6, v3)
        if (0 <= 1) {
            if len(v5) > 0 {
                v5[0] = 2147483648
            }
            v15 := GenS2{F0: 2147483647, F1: ""}
            _ = v15
        } else {
            var v16 uint64 = 18446744073709551615
            var v17 uint8 = 0
            _ = v16
            _ = v17
        }
        var v18 int8 = 100
        _ = v18
    }
    v2[893] = (v6 == -2147483648)
    switch "hello" {
    case "":
        var v19 int = 4294967295
        _ = v19
    default:
        fmt.Println(v4, v6, v4)
    }
    v20 := GenS2{F0: 2147483648, F1: ("hello" + "a")}
    v4 -= v4
    v6 += v6
    if false {
        if false {
            if (v4 < 1.5) {
                v5 = append(v5, 892)
                fmt.Println(v4, v3)
                switch -2147483648 {
                case 42:
                    if true {
                        v3 = (-(v6))
                    } else {
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        var v21 uint8 = 127
                        var v22 uint32 = 4294967295
                        _ = v21
                        _ = v22
                    }
                case -983:
                    fmt.Println(v3, v3, v3)
                case 211:
                    fmt.Println(v4, v4)
                }
            } else {
                v3 = (v6 & v6)
            }
            var v23 bool = true
            var v24 uint64 = 42
            _ = v23
            _ = v24
        } else {
            var v25 string = ("hello" + "a")
            var v26 uint16 = 32767
            _ = v25
            _ = v26
        }
    }
    if (196 <= 139) {
        v27 := GenS3{F0: (!true), F1: 983}
        fmt.Println(v6)
        _ = v27
    }
    v28 := GenS4{F0: 42, F1: 0, F2: 1}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
