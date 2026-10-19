package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint }
type GenS2 struct { F0 int32; F1 int16; F2 int8 }
type GenS3 struct { F0 uint32; F1 uint32; F2 bool }
type GenS4 struct { F0 uint64; F1 int; F2 uint64 }
func genFn0(p0 string, p1 uint8, p2 int32) int32 {
    if false {
        if ((1e10 - -2.25) <= (1e10 / -2.25)) {
            for fi7 := 0; fi7 < 12; fi7++ {
                p0 = ("a" + "\n\t")
                fmt.Println(p1, p1, p2)
                _ = fi7
            }
            switch p0 {
            case "\n\t":
                p0 = p0
            }
            p2 += 42
        }
        fmt.Println(p0, p2, p1)
    }
    if false {
        var v8 int8 = -1
        if false {
            fmt.Println(p1, v8, p1)
        }
        _ = v8
    } else {
        var v9 int16 = 77
        v10 := GenS1{F0: (0.0 != 1.5), F1: 475}
        _ = v10
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    return (p2 ^ -1)
}

func main() {
    v1 := []uint32{}
    var v2 uint16 = 656
    v3 := make(map[int]string)
    var v4 float32 = -2.25
    var v5 int16 = 1000
    v6 := make(map[uint32]int)
    if (v5 == v5) {
        var v11 uint64 = 745
        if (true || (65535 < 107)) {
            var v12 bool = (1 <= 0)
            _ = v12
        } else {
            if (!true) {
                var v13 int64 = 1
                v11 = v11
                _ = v13
            } else {
                for fi14 := 0; fi14 < 1; fi14++ {
                    v15 := GenS1{F0: (!(!true)), F1: 0}
                    v11 -= 42
                    _ = fi14
                    _ = v15
                }
                if true {
                    var v16 int16 = -15
                    var v17 uint16 = (v2 * 0)
                    _ = v16
                    _ = v17
                }
                v2 += 65535
            }
            var v18 uint32 = 0
            _ = v18
        }
        if (!(42 != -1)) {
            _ = len(v3)
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
        _ = v11
    }
    v19 := GenS2{F0: 42, F1: (^(v5)), F2: 0}
    var v20 uint = 570
    fmt.Println(v5)
    for fi21 := 0; fi21 < 11; fi21++ {
        v22 := GenS3{F0: 42, F1: 2147483648, F2: (-126 > 5)}
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        v20 = 42
        _ = fi21
        _ = v22
    }
    if true {
        v5 = (v5 / (((v5 / (((v5 % ((v5) | 1))) | 1))) | 1))
        fmt.Println(v5, v2, v4)
    }
    if (!(!((v20 + (v20 * v20)) == v20))) {
        v2 = v2
        v2 = (v2 ^ 0)
    } else {
        fmt.Println(v20, v4, v2)
    }
    switch "hello" {
    case "a":
        if true {
            if (!(float64(float32(((1e10 + 1.5) / (1e10 / 1e10)))) < float64((((1.5 - 3.14159) - (v4 * v4)) / v4)))) {
                switch 128 {
                case 932:
                    v2 = (v2 + 532)
                default:
                    fmt.Println(v20)
                }
            }
        } else {
            v23 := GenS4{F0: 57, F1: 2147483647, F2: 874}
            _ = v23
        }
    case "hello":
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    default:
        v20 -= (v20 % ((854) | 1))
    }
    v20 -= v20
    if true {
        fmt.Println(v4, v20)
        v24 := v6[407]
        _ = v24
    }
    v5 = v5
    if (!(!true)) {
        v4 = v4
        for fi25 := 0; fi25 < 6; fi25++ {
            v26 := GenS4{F0: 274, F1: (fi25 - (^(fi25))), F2: 212}
            v27 := GenS3{F0: 824, F1: 2147483648, F2: false}
            _ = fi25
            _ = v26
            _ = v27
        }
    } else {
        v2 = v2
        v28 := GenS4{F0: 116, F1: -2147483648, F2: 858}
        _ = v28
    }
    fmt.Println(v5)
    v29 := GenS2{F0: -2147483648, F1: 0, F2: 100}
    v4 = v4
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
