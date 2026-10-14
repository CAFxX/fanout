package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16; F2 bool }
type GenS2 struct { F0 bool }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []uint32{42, 949, 2147483647, 2}
    var v3 string = "a"
    v4 := []int{4294967295, 986}
    var v5 int16 = -32768
    v5 += (v5 | v5)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v6 int64 = -9223372036854775808
    fmt.Println(v6)
    _ = len(v2)
    var v7 float32 = 0.0
    for fi8 := 0; fi8 < 12; fi8++ {
        v7 = v7
        v7 += v7
        _ = fi8
    }
    if (false && (true || (true || (42 == 0)))) {
        var v9 int8 = 100
        _ = v9
    }
    v6 += (v6 * 215)
    var v10 bool = (-2147483648 == -1)
    v11 := GenS1{F0: (false && (false || (!(!v10)))), F1: 343, F2: v10}
    switch v3 {
    case "\n\t":
        v6 -= v6
    case "a":
        if ((v7 - (1e10 * v7)) != v7) {
            v10 = v10
        }
    case "":
        for fi12 := 0; fi12 < 20; fi12++ {
            fmt.Println(v7)
            for fi13 := 0; fi13 < 14; fi13++ {
                v14 := GenS1{F0: true, F1: 658, F2: v10}
                fmt.Println(v5, v10)
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = fi13
                _ = v14
            }
            if (!v10) {
                for fi15 := 0; fi15 < 5; fi15++ {
                    switch fi12 {
                    case -746:
                        _ = len(v1)
                    case 1:
                        v16 := GenS1{F0: v10, F1: 790, F2: (!(32767 == 395))}
                        _ = v16
                    default:
                        if false {
                            var v17 float32 = 1e10
                            switch fi12 {
                            case 32767:
                                if (!(!(!(v10 && true)))) {
                                    if v10 {
                                        _ = len(v4)
                                        fmt.Println(v3, v7, v10)
                                    }
                                }
                            }
                            _ = v17
                        }
                    }
                    v18 := GenS2{F0: (0.1 != (0.0 - (v7 + -2.25)))}
                    var v19 int8 = -85
                    _ = fi15
                    _ = v18
                    _ = v19
                }
            }
            _ = fi12
        }
    default:
        for fi20 := 0; fi20 < 8; fi20++ {
            fmt.Println(v5)
            delete(v1, 65535)
            v5 -= -1
            _ = fi20
        }
    }
    _ = len(v1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
