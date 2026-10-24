package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64; F2 uint8 }
type GenS2 struct { F0 uint8; F1 uint32 }
type GenS3 struct { F0 uint8; F1 uint32; F2 uint }
type GenS4 struct { F0 uint8; F1 int; F2 int32 }
func genFn0(p0 uint) uint32 {
    fmt.Println(p0)
    if true {
        for fi7 := 0; fi7 < 12; fi7++ {
            p0 += 55
            _ = fi7
        }
        p0 -= p0
    }
    v8 := []int{-110}
    p0 = (p0 - 0)
    p0 = (p0 | 454)
    _ = p0
    _ = v8
    return 375
}
func genFn1(p0 uint64, p1 int64, p2 int) uint {
    v9 := []int{(-(p2)), p2, (p2 << ((0) & 63)), p2}
    for fi10 := 0; fi10 < 16; fi10++ {
        p2 = p2
        var v11 uint8 = 0
        _ = fi10
        _ = v11
    }
    _ = len(v9)
    if false {
        p1 = (p1 >> ((18446744073709551615) & 63))
        fmt.Println(p0, p2, p1)
        if false {
            p0 += (p0 + 781)
        }
    } else {
        if ((p0 + 42) != (p0 - p0)) {
            if false {
                if true {
                    var v12 int64 = -1
                    if true {
                        var v13 string = "foo bar"
                        if (false && false) {
                            var v14 uint64 = p0
                            v13 = ("hello" + "\n\t")
                            v15 := GenS1{F0: 18446744073709551615, F1: -9223372036854775808, F2: 2}
                            _ = v14
                            _ = v15
                        }
                        _ = v13
                    } else {
                        var v16 uint64 = p0
                        _ = v16
                    }
                    _ = v12
                }
            }
            p1 = (^(p1))
            v17 := GenS2{F0: 1, F1: 65535}
            _ = v17
        }
    }
    switch p2 {
    case -910:
        fmt.Println(p1)
    case 2147483647:
        switch p2 {
        case 127:
            v18 := GenS3{F0: 37, F1: 1, F2: 712}
            _ = v18
        default:
            p2 = p2
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 0
}

func main() {
    v1 := []uint32{65535, 2147483647, 902, 42}
    var v2 uint = 4294967295
    v3 := make(map[int]int)
    v4 := make(map[int]uint32)
    v5 := make(map[uint32]int)
    var v6 int = 490
    for fi19 := 0; fi19 < 13; fi19++ {
        v20 := GenS2{F0: 42, F1: 65535}
        _ = fi19
        _ = v20
    }
    fmt.Println(v6, v2)
    if ((true && false) && true) {
        for fi21 := 0; fi21 < 20; fi21++ {
            var v22 uint32 = 4294967295
            _ = fi21
            _ = v22
        }
    } else {
        var v23 int = (v6 - (v6 / ((-1) | 1)))
        var v24 uint = v2
        _ = v23
        _ = v24
    }
    var v25 float32 = float32(0.0)
    v26 := GenS4{F0: 127, F1: v6, F2: 2147483647}
    switch "\n\t" {
    case "\n\t":
        fmt.Println(v6, v25, v6)
    case "foo bar":
        _ = len(v4)
    case "":
        fmt.Println(v6)
    }
    v25 += (float32(((((0.1 - (0.0 / 1.5)) + 1.5) - float64(((v25 / 0.1) + 0.1))) * 3.14159)) * v25)
    fmt.Println(v6)
    if len(v1) > 0 {
        v1[0] = genFn0(v2)
    }
    v27 := GenS3{F0: 36, F1: 42, F2: (^(v2))}
    var v28 bool = false
    v6 = v6
    v2 = (v2 % ((v2) | 1))
    var v29 string = ("hello" + ("hello" + ""))
    fmt.Println(v6)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
