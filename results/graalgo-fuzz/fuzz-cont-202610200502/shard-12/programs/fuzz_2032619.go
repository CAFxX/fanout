package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 string; F1 uint; F2 uint }

func main() {
    v1 := []int{32768, -451}
    v2 := []uint32{2, 2, 2, 42}
    var v3 float64 = 0.1
    var v4 int64 = -9223372036854775808
    var v5 uint = 4294967295
    var v6 uint32 = 2
    v1 = append(v1, -153)
    v7 := GenS1{F0: 793}
    fmt.Println(v3)
    var v8 uint64 = 42
    v9 := GenS2{F0: ("foo bar" + ""), F1: v5, F2: (v5 - v5)}
    v4 -= v4
    var v10 uint8 = 1
    var v11 bool = (!(true || true))
    v4 -= (v4 + (v4 | -9223372036854775808))
    switch 407 {
    case 128:
        v4 = (v4 % (((v4 % (((v4 ^ v4)) | 1))) | 1))
    }
    for fi12 := 0; fi12 < 10; fi12++ {
        if true {
            v6 = 719
            var v13 int64 = v4
            var v14 float32 = (3.14159 - ((1.5 + 1e10) * float32(v3)))
            _ = v13
            _ = v14
        } else {
            var v15 int = fi12
            _ = v15
        }
        if false {
            v8 += v8
            v10 = (v10 / (((v10 - 127)) | 1))
        } else {
            var v16 uint = (v5 | v5)
            _ = v16
        }
        var v17 float64 = 1e10
        _ = fi12
        _ = v17
    }
    switch 128 {
    case -922:
        var v18 int32 = -1
        _ = v18
    case 582:
        switch 32767 {
        case -2147483648:
            _ = len(v1)
        }
    case 636:
        var v19 float32 = (1e10 * 1e10)
        _ = v19
    default:
        if (false && v11) {
            fmt.Println(v3, v3, v4)
        } else {
            _ = len(v2)
            _ = len(v2)
        }
    }
    if false {
        if false {
            v10 -= v10
            if true {
                var v20 float64 = ((v3 + 1e10) - v3)
                _ = v20
            }
            fmt.Println(v5, v5, v11)
        } else {
            fmt.Println(v8)
            if v11 {
                v3 += v3
                var v21 int16 = 32767
                for fi22 := 0; fi22 < 7; fi22++ {
                    if v11 {
                        switch "" {
                        case "foo bar":
                            fmt.Println(v10, v3)
                        default:
                            v21 += 514
                        }
                        v8 = 2
                    } else {
                        v10 += (v10 >> (((v5 + v5)) & 63))
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    }
                    _ = fi22
                }
                _ = v21
            }
        }
    } else {
        var v23 int8 = -128
        var v24 float64 = 1.5
        _ = v23
        _ = v24
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
