package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32 }
type GenS2 struct { F0 uint; F1 int16; F2 bool }
func genFn0(p0 uint, p1 uint32) uint {
    switch "hello" {
    case "a":
        v6 := GenS1{F0: 9223372036854775807, F1: 1}
        _ = v6
    case "":
        switch 524 {
        case -572:
            if true {
                if (!true) {
                    switch "\n\t" {
                    case "":
                        var v7 int32 = 42
                        _ = v7
                    case "a":
                        var v8 float64 = -2.25
                        _ = v8
                    }
                    fmt.Println(p1)
                } else {
                    if true {
                        if false {
                            switch 4294967295 {
                            case 598:
                                v9 := []int{127, int(p0)}
                                _ = v9
                            case 4294967295:
                                p0 = (^(p0))
                            case -692:
                                p1 = 65535
                            }
                        } else {
                            var v10 int = 1
                            v11 := GenS1{F0: v10, F1: 1}
                            _ = v10
                            _ = v11
                        }
                    }
                    p0 = p0
                }
            } else {
                p0 -= 861
            }
        case 0:
            var v12 float32 = (3.14159 / -2.25)
            _ = v12
        case -183:
            p1 = 274
        }
    default:
        fmt.Println(p1)
    }
    var v13 int = int(p1)
    if true {
        p1 += 2
        p1 = 395
    }
    _ = p0
    _ = p1
    _ = v13
    return p0
}
func genFn1(p0 uint32, p1 uint32, p2 int) uint16 {
    var v14 int = (^(p2))
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    return 0
}

func main() {
    var v1 uint32 = 805
    var v2 uint64 = 0
    v3 := make(map[int]int)
    var v4 float32 = (-2.25 + 3.14159)
    var v5 uint8 = 114
    if false {
        v15 := []int{128}
        _ = v15
    } else {
        var v16 string = "a"
        _ = v16
    }
    v1 -= (v1 << ((18446744073709551615) & 63))
    v2 = v2
    for fi17 := 0; fi17 < 1; fi17++ {
        fmt.Println(v5, v2, v1)
        if (65535 <= 255) {
            v4 = (v4 - v4)
        } else {
            v18 := v3[(^(fi17))]
            _ = v18
        }
        var v19 uint64 = 18446744073709551615
        _ = fi17
        _ = v19
    }
    for fi20 := 0; fi20 < 6; fi20++ {
        var v21 uint = 782
        _ = fi20
        _ = v21
    }
    v22 := []int{2147483647}
    fmt.Println(v2, v5, v5)
    var v23 int16 = -32768
    v23 += v23
    if (int64(v23) == 9223372036854775807) {
        var v24 float32 = ((v4 / (0.0 + float32(float64(((3.14159 + float32(3.14159)) - v4))))) / v4)
        v5 = v5
        v1 += 65535
        _ = v24
    } else {
        v23 -= (v23 / ((int16(v2)) | 1))
        switch -9223372036854775808 {
        case 2147483647:
            v22 = append(v22, -2147483648)
        case -9223372036854775808:
            var v25 uint8 = (v5 * 5)
            _ = v25
        case 42:
            v26 := GenS2{F0: 0, F1: 32767, F2: true}
            _ = v26
        default:
            if (!false) {
                for fi27 := 0; fi27 < 16; fi27++ {
                    for fi28 := 0; fi28 < 8; fi28++ {
                        fmt.Println(fi28, v5)
                        v29 := GenS2{F0: uint(v1), F1: v23, F2: (339 > 42)}
                        _ = fi28
                        _ = v29
                    }
                    _ = fi27
                }
                var v30 uint8 = (^(v5))
                v31 := GenS1{F0: 0, F1: 2147483647}
                _ = v30
                _ = v31
            } else {
                v32 := GenS1{F0: -1, F1: 2147483647}
                switch 1 {
                case 7:
                    if len(v22) >= 2 {
                        v22 = v22[1:len(v22)]
                    }
                case 399:
                    v4 -= v4
                }
                _ = v32
            }
        }
        var v33 uint64 = uint64(v1)
        _ = v33
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
