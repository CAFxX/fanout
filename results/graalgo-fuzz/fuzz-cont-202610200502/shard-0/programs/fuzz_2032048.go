package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string; F2 int8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 string; F1 int32 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint64 }

func main() {
    v1 := make(map[int]bool)
    var v2 string = "foo bar"
    v3 := make(map[uint32]uint32)
    var v4 int16 = 1
    v5 := GenS1{F0: 42, F1: v2, F2: 127}
    v6 := []int{0, 1}
    for fi7 := 0; fi7 < 16; fi7++ {
        fmt.Println(v4, fi7, v4)
        fmt.Println(fi7, v2)
        _ = fi7
    }
    v4 -= (v4 >> ((786) & 63))
    var v8 int32 = 0
    for fi9 := 0; fi9 < 7; fi9++ {
        var v10 uint8 = 42
        v8 -= (^(v8))
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        _ = fi9
        _ = v10
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    switch 0 {
    case 9223372036854775807:
        v11 := GenS2{F0: 0}
        _ = v11
    case 42:
        v6 = append(v6, 32768)
    case -234:
        v2 = v2
    }
    var v12 int = -1
    var v13 uint = 18446744073709551615
    if false {
        var v14 int8 = -94
        var v15 bool = true
        v16 := GenS3{F0: v2, F1: v8}
        _ = v14
        _ = v15
        _ = v16
    } else {
        switch v2 {
        case "a":
            if len(v6) > 0 {
                v6[0] = 1
            }
        case "hello":
            if (-1 != -2147483648) {
                if (!false) {
                    for fi17 := 0; fi17 < 19; fi17++ {
                        if (((1.5 / ((1.5 * float32(float64(1e10))) / 1.5)) - float32((3.14159 - (1e10 + ((1e10 + float64(0.0)) / 1e10))))) > 3.14159) {
                            for fi18 := 0; fi18 < 15; fi18++ {
                                fmt.Println(v13, v2)
                                _ = fi18
                            }
                            v13 -= 1
                        }
                        _ = fi17
                    }
                }
            }
        }
    }
    v19 := GenS4{F0: 0}
    switch v2 {
    case "a":
        v4 = v4
    default:
        if (int(v4) == 2147483647) {
            v20 := GenS5{F0: uint64(v13)}
            _ = v20
        } else {
            v3[2147483647] = 146
            if (false && true) {
                var v21 int32 = v8
                var v22 int = (v12 / ((v12) | 1))
                if true {
                    var v23 int32 = v8
                    _ = v23
                } else {
                    fmt.Println(v8, v22)
                }
                _ = v21
                _ = v22
            }
            var v24 int16 = 1000
            _ = v24
        }
    }
    var v25 int16 = v4
    fmt.Println(v13, v2, v25)
    var v26 int64 = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v13
    _ = v19
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
