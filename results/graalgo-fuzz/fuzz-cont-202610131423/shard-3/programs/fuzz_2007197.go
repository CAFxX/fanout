package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int8 }
func genFn0(p0 int64, p1 bool, p2 bool) int8 {
    var v7 uint32 = 156
    p1 = false
    var v8 uint = 65535
    for fi9 := 0; fi9 < 11; fi9++ {
        p0 = int64(fi9)
        fmt.Println(v8, v7, v7)
        p0 += p0
        _ = fi9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return -128
}

func main() {
    v1 := make(map[string]string)
    var v2 int16 = -780
    var v3 uint32 = 922
    v4 := make(map[uint32]string)
    v5 := make(map[uint32]uint32)
    var v6 uint = 1
    switch "" {
    case "hello":
        fmt.Println(v6, v6, v2)
    case "foo bar":
        if (3.14159 < -2.25) {
            v3 += v3
            var v10 uint8 = 127
            _ = v10
        }
    case "a":
        v3 = v3
    default:
        v3 = v3
    }
    v2 -= v2
    var v11 int16 = (v2 & (v2 % ((v2) | 1)))
    if true {
        v11 = (v2 ^ 32767)
        fmt.Println(v6)
    }
    for fi12 := 0; fi12 < 11; fi12++ {
        v11 += 32767
        v13 := GenS1{F0: -592, F1: 100}
        for fi14 := 0; fi14 < 16; fi14++ {
            switch "" {
            case "hello":
                v2 = v2
            case "\n\t":
                v11 += 825
            }
            var v15 uint64 = 0
            _ = fi14
            _ = v15
        }
        _ = fi12
        _ = v13
    }
    for fi16 := 0; fi16 < 15; fi16++ {
        v3 += v3
        if true {
            v2 = v2
            fmt.Println(v6, fi16)
            v11 = (v11 ^ 45)
        } else {
            v3 = (v3 / ((0) | 1))
            v17 := GenS1{F0: 0, F1: -52}
            v18 := GenS2{F0: (v6 << (((v6 * 90)) & 63))}
            _ = v17
            _ = v18
        }
        v3 = v3
        _ = fi16
    }
    switch 1 {
    case 128:
        switch "" {
        case "hello":
            if ((42 >= 849) || false) {
                fmt.Println(v2)
                v6 = 255
                v2 = 42
            } else {
                for fi19 := 0; fi19 < 6; fi19++ {
                    if (!true) {
                        var v20 string = ""
                        _ = v20
                    } else {
                        v2 += v11
                        if (true && (2 < 1)) {
                            var v21 uint32 = v3
                            _ = v21
                        }
                        fmt.Println(v2, fi19, v6)
                    }
                    var v22 uint64 = 1
                    for fi23 := 0; fi23 < 19; fi23++ {
                        var v24 int16 = 1
                        v25 := []int{(fi23 % (((fi23 % ((4294967295) | 1))) | 1)), -2147483648, (fi19 % ((fi19) | 1)), (fi19 ^ (fi23 % ((fi19) | 1)))}
                        _ = fi23
                        _ = v24
                        _ = v25
                    }
                    _ = fi19
                    _ = v22
                }
                fmt.Println(v11)
            }
        default:
            if false {
                v5[(v3 & v3)] = (v3 | uint32(v6))
                v26 := GenS3{F0: genFn0(0, (v3 != v3), true)}
                for fi27 := 0; fi27 < 16; fi27++ {
                    v2 += (v2 ^ v2)
                    _ = fi27
                }
                _ = v26
            }
        }
    case 469:
        var v28 uint16 = 255
        _ = v28
    }
    fmt.Println(v6, v3, v2)
    fmt.Println(v2, v11, v6)
    for fi29 := 0; fi29 < 20; fi29++ {
        v11 -= (v2 + v2)
        var v30 float64 = -2.25
        fmt.Println(fi29, v3)
        _ = fi29
        _ = v30
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
