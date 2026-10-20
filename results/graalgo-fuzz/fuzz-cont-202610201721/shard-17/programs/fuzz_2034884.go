package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint64; F1 uint16 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 float64 = 0.0
    v3 := make(map[string]int)
    v4 := []uint32{4294967295, 42, 487, 1}
    var v5 uint32 = 129
    v5 += v5
    var v6 uint16 = 263
    for fi7 := 0; fi7 < 15; fi7++ {
        v4 = append(v4, (v5 | 2147483648))
        for fi8 := 0; fi8 < 16; fi8++ {
            var v9 float64 = v2
            var v10 bool = (false && false)
            _ = fi8
            _ = v10
            _ = v9
        }
        delete(v3, ("a" + ("\n\t" + "hello")))
        _ = fi7
    }
    var v11 float64 = float64(((1e10 / float32(v2)) - ((0.1 * 0.1) - 0.0)))
    for fi12 := 0; fi12 < 17; fi12++ {
        fmt.Println(v6)
        if true {
            for fi13 := 0; fi13 < 3; fi13++ {
                v5 = 1
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                fmt.Println(v5)
                _ = fi13
            }
            if false {
                v11 = v2
                for fi14 := 0; fi14 < 8; fi14++ {
                    if len(v4) > 0 {
                        v4[0] = (v5 >> ((4294967295) & 63))
                    }
                    _ = fi14
                }
            }
        }
        _ = fi12
    }
    v15 := GenS1{F0: 967, F1: 255}
    v6 += 65535
    if len(v4) > 0 {
        v4[0] = 0
    }
    switch "\n\t" {
    case "\n\t":
        v16 := GenS2{F0: 219}
        _ = v16
    case "":
        if ((20 != 255) || (!(float64((-2.25 / float32(v11))) < float64(0.0)))) {
            for fi17 := 0; fi17 < 10; fi17++ {
                v18 := GenS3{F0: 2, F1: v6}
                _ = fi17
                _ = v18
            }
            fmt.Println(v5, v6, v11)
            fmt.Println(v11, v11)
        } else {
            var v19 int8 = -1
            for fi20 := 0; fi20 < 13; fi20++ {
                var v21 int = 2147483647
                _ = fi20
                _ = v21
            }
            _ = v19
        }
    default:
        v22 := GenS1{F0: 0, F1: 0}
        _ = v22
    }
    v2 = (-2.25 / v11)
    var v23 int = 127
    var v24 float32 = (0.0 - (1e10 / 1e10))
    var v25 float64 = v2
    fmt.Println(v11, v11)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v15
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
