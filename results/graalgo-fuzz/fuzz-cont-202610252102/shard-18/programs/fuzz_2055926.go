package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8; F2 string }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 string; F1 uint8; F2 uint }
type GenS4 struct { F0 int }
func genFn0(p0 uint32) int8 {
    if false {
        for fi6 := 0; fi6 < 14; fi6++ {
            p0 -= (p0 ^ p0)
            _ = fi6
        }
        switch "\n\t" {
        case "foo bar":
            fmt.Println(p0)
        default:
            if true {
                v7 := []int{0, 0, -9223372036854775808, 1}
                _ = v7
            }
        }
    } else {
        p0 = 42
        var v8 int = -58
        switch "a" {
        case "\n\t":
            var v9 uint32 = (p0 ^ p0)
            _ = v9
        case "":
            v8 = v8
        default:
            for fi10 := 0; fi10 < 16; fi10++ {
                var v11 int = (fi10 ^ fi10)
                v12 := GenS1{F0: p0, F1: 42, F2: (((("a" + "") + "foo bar") + ("" + "foo bar")) + ("\n\t" + (("a" + "a") + "hello")))}
                if true {
                    for fi13 := 0; fi13 < 12; fi13++ {
                        if (!false) {
                            var v14 int16 = 975
                            p0 = p0
                            _ = v14
                        }
                        p0 -= p0
                        p0 += (p0 % ((p0) | 1))
                        _ = fi13
                    }
                    v15 := GenS2{F0: false}
                    fmt.Println(v8, p0)
                    _ = v15
                }
                _ = fi10
                _ = v11
                _ = v12
            }
        }
        _ = v8
    }
    for fi16 := 0; fi16 < 17; fi16++ {
        var v17 float64 = 1e10
        _ = fi16
        _ = v17
    }
    if false {
        var v18 int32 = -580
        fmt.Println(p0, p0)
        _ = v18
    }
    _ = p0
    return 118
}
func genFn1(p0 int16) uint32 {
    var v19 uint = 1
    if true {
        p0 = (p0 << (((v19 + v19)) & 63))
        if true {
            var v20 float32 = 3.14159
            _ = v20
        } else {
            v21 := GenS3{F0: (("\n\t" + "\n\t") + "foo bar"), F1: 0, F2: (v19 ^ v19)}
            if (!(42 == genFn0(2147483647))) {
                v22 := GenS2{F0: (863 != 42)}
                if true {
                    var v23 float32 = 3.14159
                    _ = v23
                }
                v24 := []int{1}
                _ = v22
                _ = v24
            }
            _ = v21
        }
    }
    if false {
        var v25 uint = (v19 | (v19 ^ (v19 / (((v19 - (^(v19)))) | 1))))
        _ = v25
    } else {
        for fi26 := 0; fi26 < 2; fi26++ {
            v19 = v19
            v27 := []int{fi26}
            _ = fi26
            _ = v27
        }
        v19 = 18446744073709551615
        switch -2147483648 {
        case 282:
            var v28 int8 = genFn0(594)
            _ = v28
        case -219:
            v29 := GenS2{F0: (42 >= 0)}
            _ = v29
        default:
            if false {
                for fi30 := 0; fi30 < 6; fi30++ {
                    v31 := GenS3{F0: "a", F1: 2, F2: v19}
                    v32 := GenS2{F0: (uint8(v19) < 2)}
                    _ = fi30
                    _ = v31
                    _ = v32
                }
            }
        }
    }
    v19 -= (v19 / ((2) | 1))
    _ = p0
    _ = v19
    return 718
}

func main() {
    var v1 int8 = 100
    var v2 uint8 = 127
    var v3 int8 = 1
    var v4 int64 = -983
    var v5 float32 = 1e10
    v33 := GenS1{F0: genFn1(int16(v4)), F1: v2, F2: ((((("hello" + "hello") + ((("\n\t" + "foo bar") + "foo bar") + (("foo bar" + "foo bar") + ""))) + "foo bar") + ("a" + "")) + ("hello" + "hello"))}
    switch "foo bar" {
    case "hello":
        var v34 float64 = (1e10 / 1.5)
        _ = v34
    case "\n\t":
        v35 := GenS1{F0: 448, F1: v2, F2: "\n\t"}
        _ = v35
    case "":
        fmt.Println(v4, v1)
    }
    if false {
        v36 := []int{764, 2147483647}
        v4 -= (v4 * (v4 | (^(v4))))
        v37 := GenS4{F0: -467}
        _ = v36
        _ = v37
    } else {
        v2 = v2
    }
    var v38 bool = false
    v39 := []int{}
    for fi40 := 0; fi40 < 15; fi40++ {
        _ = len(v39)
        _ = fi40
    }
    v2 = (v2 ^ 0)
    v5 = v5
    v38 = v38
    for fi41 := 0; fi41 < 2; fi41++ {
        if len(v39) > 0 {
            v39[0] = (^(fi41))
        }
        _ = fi41
    }
    var v42 uint64 = 130
    for fi43 := 0; fi43 < 3; fi43++ {
        var v44 float32 = (1e10 - v5)
        var v45 float64 = 0.1
        _ = fi43
        _ = v44
        _ = v45
    }
    v2 = (v2 / (((v2 % (((v2 << ((255) & 63))) | 1))) | 1))
    fmt.Println(v3, v1, v2)
    fmt.Println(v1, v2, v38)
    fmt.Println(v3, v38, v38)
    v5 = (float32(-2.25) * v5)
    v42 += 274
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v39:", len(v39), v39)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v38
    _ = v39
    _ = v4
    _ = v42
    _ = v5
}
