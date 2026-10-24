package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64; F2 bool }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint16; F1 string }
func genFn0(p0 int64, p1 uint64, p2 uint) uint64 {
    var v6 int32 = -2147483648
    if (0.0 >= -2.25) {
        var v7 int = -9223372036854775808
        _ = v7
    } else {
        fmt.Println(p2)
        for fi8 := 0; fi8 < 12; fi8++ {
            if (137 != 127) {
                v9 := GenS1{F0: (p1 + p1), F1: p1, F2: (42 > v6)}
                v10 := GenS2{F0: (^(p1))}
                _ = v10
                _ = v9
            }
            var v11 int64 = (-(p0))
            _ = fi8
            _ = v11
        }
        p1 -= 0
    }
    var v12 uint16 = 32767
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v6
    return uint64(v12)
}

func main() {
    var v1 uint32 = 2
    var v2 float32 = (-2.25 + 0.1)
    v3 := make(map[uint32]int)
    v4 := make(map[uint32]bool)
    var v5 int64 = -1
    v2 += (1e10 / v2)
    v1 = v1
    var v13 int = 1
    v14 := []int{907}
    if false {
        _ = len(v3)
        for fi15 := 0; fi15 < 11; fi15++ {
            var v16 uint8 = 70
            switch "a" {
            case "\n\t":
                for fi17 := 0; fi17 < 14; fi17++ {
                    var v18 int64 = v5
                    v16 = 106
                    switch "foo bar" {
                    case "hello":
                        var v19 int8 = 102
                        _ = v19
                    }
                    _ = fi17
                    _ = v18
                }
            case "foo bar":
                switch "hello" {
                case "":
                    if len(v14) >= 2 {
                        v14 = v14[1:len(v14)]
                    }
                case "foo bar":
                    v20 := GenS2{F0: 759}
                    _ = v20
                default:
                    for fi21 := 0; fi21 < 7; fi21++ {
                        if (42 < v1) {
                            fmt.Println(fi15)
                        }
                        _ = fi21
                    }
                }
            case "hello":
                var v22 uint = 2
                _ = v22
            default:
                v1 += (v1 * 2)
            }
            _ = fi15
            _ = v16
        }
    } else {
        v2 -= -2.25
    }
    v13 += (v13 - (^(v13)))
    v1 = v1
    if len(v14) > 0 {
        v14[0] = 32767
    }
    v23 := GenS3{F0: 255, F1: (("" + ("\n\t" + "a")) + ("foo bar" + ""))}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
