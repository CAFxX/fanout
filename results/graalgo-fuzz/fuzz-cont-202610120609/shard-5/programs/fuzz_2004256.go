package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint }
type GenS2 struct { F0 bool; F1 bool; F2 bool }
type GenS3 struct { F0 uint8; F1 int8; F2 int64 }
type GenS4 struct { F0 string; F1 uint32; F2 uint32 }
type GenS5 struct { F0 uint; F1 int }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 string; F1 int16 }
func genFn0(p0 uint16) uint16 {
    for fi7 := 0; fi7 < 16; fi7++ {
        if (!true) {
            v8 := []int{(fi7 >> ((255) & 63)), 52, (fi7 + (fi7 - fi7)), fi7}
            _ = v8
        } else {
            p0 += 1
        }
        _ = fi7
    }
    p0 -= p0
    var v9 uint = 65535
    v10 := []int{-1, 1}
    _ = p0
    _ = v10
    _ = v9
    return p0
}
func genFn1(p0 uint64, p1 uint8, p2 int64) int64 {
    var v11 uint = 255
    if (!(2147483647 <= 314)) {
        var v12 int8 = 0
        for fi13 := 0; fi13 < 16; fi13++ {
            for fi14 := 0; fi14 < 16; fi14++ {
                switch "a" {
                case "\n\t":
                    v15 := GenS1{F0: false, F1: 546}
                    _ = v15
                }
                if false {
                    var v16 int8 = (v12 | v12)
                    if true {
                        p1 = p1
                        v17 := GenS2{F0: ((v11 - (v11 | (v11 >> (((v11 * (v11 ^ (v11 ^ v11)))) & 63)))) == 2), F1: ((0.1 != (0.1 * (float32(3.14159) + -2.25))) && true), F2: false}
                        fmt.Println(fi14, v16, v11)
                        _ = v17
                    } else {
                        v11 = v11
                        v18 := []int{(fi13 & fi14), fi14}
                        p2 = p2
                        _ = v18
                    }
                    v19 := GenS3{F0: (p1 + (p1 % ((24) | 1))), F1: v16, F2: 1}
                    _ = v16
                    _ = v19
                }
                _ = fi14
            }
            _ = fi13
        }
        _ = v12
    } else {
        var v20 int = -1
        var v21 bool = (true && true)
        _ = v20
        _ = v21
    }
    p0 += p0
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return p2
}

func main() {
    var v1 int = 42
    var v2 int16 = -401
    v3 := []int{v1, v1, (-(v1))}
    var v4 uint16 = 453
    v5 := make(map[uint32]bool)
    v6 := []string{}
    if len(v6) > 0 {
        v6[0] = ""
    }
    v4 -= v4
    var v22 float64 = 3.14159
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    var v23 float64 = v22
    for fi24 := 0; fi24 < 2; fi24++ {
        if (1 == 0) {
            v4 = (^(v4))
            switch (fi24 << ((65535) & 63)) {
            case 30:
                v25 := GenS4{F0: "\n\t", F1: 42, F2: 1}
                _ = v25
            default:
                for fi26 := 0; fi26 < 10; fi26++ {
                    v23 = (v23 - v23)
                    v2 = v2
                    v27 := GenS5{F0: 1, F1: -1}
                    _ = fi26
                    _ = v27
                }
            }
            if true {
                v1 += (v1 | 127)
                fmt.Println(v22)
            } else {
                if true {
                    v23 = (float64((0.1 * 1.5)) * 1.5)
                } else {
                    for fi28 := 0; fi28 < 17; fi28++ {
                        fmt.Println(v23)
                        v23 = v22
                        _ = fi28
                    }
                    v2 -= 32767
                }
                v6 = append(v6, "")
            }
        } else {
            v3 = append(v3, fi24)
            if (((1.5 * 1.5) * (3.14159 - 0.1)) <= 3.14159) {
                v29 := GenS6{F0: v2}
                _ = v29
            }
            var v30 int8 = 100
            _ = v30
        }
        v23 = v22
        v31 := GenS3{F0: 131, F1: 127, F2: 9223372036854775807}
        _ = fi24
        _ = v31
    }
    switch ("hello" + "foo bar") {
    case "foo bar":
        for fi32 := 0; fi32 < 16; fi32++ {
            var v33 float32 = (0.0 / 1e10)
            v4 += (v4 / (((v4 / (((v4 << ((983) & 63))) | 1))) | 1))
            _ = fi32
            _ = v33
        }
    case "hello":
        if false {
            for fi34 := 0; fi34 < 3; fi34++ {
                v35 := GenS7{F0: 9223372036854775807}
                _ = fi34
                _ = v35
            }
            for fi36 := 0; fi36 < 20; fi36++ {
                v2 -= 0
                for fi37 := 0; fi37 < 5; fi37++ {
                    v38 := GenS8{F0: "\n\t", F1: v2}
                    var v39 float32 = 1e10
                    if (65535 < 65535) {
                        var v40 int64 = -1
                        _ = v40
                    } else {
                        var v41 int64 = 1
                        _ = v41
                    }
                    _ = len(v6)
                    fmt.Println(v39)
                    _ = fi37
                    _ = v38
                    _ = v39
                }
                v2 = v2
                _ = fi36
            }
        } else {
            v22 = (v22 / 1.0)
            v4 -= 696
            fmt.Println(v4)
        }
    default:
        if true {
            if len(v3) > 0 {
                v3[0] = -1
            }
            var v42 int8 = -94
            _ = v42
        }
    }
    var v43 float32 = 1.5
    var v44 float32 = v43
    var v45 int64 = -1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v45
    _ = v5
    _ = v6
}
