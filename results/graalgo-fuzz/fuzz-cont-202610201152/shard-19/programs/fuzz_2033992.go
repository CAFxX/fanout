package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 bool; F1 uint16 }
func genFn0(p0 int64, p1 uint8) uint16 {
    if (2147483647 == 42) {
        if false {
            var v4 float32 = 1.5
            _ = v4
        }
        p0 = 0
    } else {
        var v5 int16 = 1
        v6 := GenS1{F0: 42}
        _ = v5
        _ = v6
    }
    var v7 int32 = 1
    _ = p0
    _ = p1
    _ = v7
    return 2
}
func genFn1(p0 uint32, p1 uint64, p2 uint8) uint16 {
    v8 := []int{}
    p1 = (p1 / (((^(p1))) | 1))
    var v9 float32 = 3.14159
    v10 := GenS1{F0: 59}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v8
    _ = v9
    return genFn0(9223372036854775807, p2)
}

func main() {
    var v1 int = -773
    var v2 int8 = 1
    v3 := make(map[uint32]uint32)
    v11 := []int{(v1 ^ v1)}
    fmt.Println(v1, v1)
    v2 = 42
    v12 := GenS1{F0: v2}
    fmt.Println(v2)
    switch "\n\t" {
    case "":
        v13 := GenS2{F0: true, F1: 32767}
        _ = v13
    default:
        if false {
            var v14 int64 = 9223372036854775807
            var v15 uint8 = 0
            _ = v14
            _ = v15
        }
    }
    switch 128 {
    case -18:
        fmt.Println(v2, v2)
    case -9223372036854775808:
        if (!((v2 >= v2) && (!(v2 > (^(v2)))))) {
            if len(v11) >= 2 {
                v11 = v11[1:len(v11)]
            }
        }
    case -259:
        var v16 int8 = (v2 | int8(v1))
        _ = v16
    }
    v1 -= v1
    fmt.Println(v1, v2)
    if (!(!(!true))) {
        delete(v3, 42)
    } else {
        delete(v3, 2147483648)
        if false {
            if false {
                for fi17 := 0; fi17 < 1; fi17++ {
                    switch v1 {
                    case 1:
                        fmt.Println(v1, v1, v1)
                    case 32767:
                        switch "\n\t" {
                        case "\n\t":
                            var v18 uint16 = 2
                            _ = v18
                        default:
                            fmt.Println(fi17, fi17)
                        }
                    }
                    _ = len(v11)
                    _ = fi17
                }
            }
        } else {
            fmt.Println(v1, v2)
            switch v1 {
            case -9223372036854775808:
                var v19 string = "hello"
                _ = v19
            case 0:
                if true {
                    switch "" {
                    case "hello":
                        _ = len(v11)
                    case "\n\t":
                        for fi20 := 0; fi20 < 16; fi20++ {
                            if len(v11) > 0 {
                                v11[0] = fi20
                            }
                            _ = fi20
                        }
                    default:
                        v2 -= (v2 % (((v2 << ((882) & 63))) | 1))
                    }
                }
            default:
                switch "a" {
                case "hello":
                    fmt.Println(v1, v1)
                case "\n\t":
                    v11 = append(v11, 32767)
                case "foo bar":
                    fmt.Println(v1)
                }
            }
            if (float64(1.5) != (0.1 - float64(float32(0.0)))) {
                v1 = v1
                for fi21 := 0; fi21 < 2; fi21++ {
                    _ = len(v3)
                    _ = len(v3)
                    _ = len(v11)
                    _ = fi21
                }
            }
        }
    }
    v22 := v3[1]
    for fi23 := 0; fi23 < 6; fi23++ {
        fmt.Println(v2, v22)
        for fi24 := 0; fi24 < 16; fi24++ {
            v2 -= (v2 << ((4294967295) & 63))
            _ = len(v3)
            _ = fi24
        }
        _ = fi23
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v22
    _ = v3
}
