package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint64; F1 bool }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 string }
func genFn0(p0 int8, p1 int32, p2 bool) uint8 {
    v6 := GenS1{F0: false}
    fmt.Println(p1, p1, p0)
    switch -774 {
    case 127:
        if ((^(p0)) == -1) {
            if (!p2) {
                var v7 uint8 = 128
                _ = v7
            } else {
                if p2 {
                    fmt.Println(p2, p1)
                    for fi8 := 0; fi8 < 14; fi8++ {
                        fmt.Println(p1)
                        v9 := GenS2{F0: 229, F1: false}
                        _ = fi8
                        _ = v9
                    }
                }
            }
            for fi10 := 0; fi10 < 11; fi10++ {
                p1 -= p1
                p2 = (2 != 42)
                v11 := GenS2{F0: 18446744073709551615, F1: true}
                _ = fi10
                _ = v11
            }
        } else {
            p2 = (p2 || true)
        }
    default:
        v12 := []int{128, 32767}
        _ = v12
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 42
}
func genFn1(p0 uint, p1 uint32) int16 {
    var v13 bool = (32767 == 1000)
    v14 := []int{}
    v15 := GenS3{F0: -1}
    if v13 {
        var v16 int64 = 230
        if len(v14) > 0 {
            v14[0] = -485
        }
        if len(v14) > 0 {
            v14[0] = 127
        }
        _ = v16
    }
    fmt.Println(p0, p0)
    _ = p0
    _ = p1
    _ = v13
    _ = v14
    _ = v15
    return 1
}

func main() {
    var v1 uint32 = 2147483647
    var v2 float32 = -2.25
    var v3 uint16 = 378
    var v4 float32 = float32((1e10 - 0.1))
    v5 := make(map[uint32]bool)
    if true {
        if true {
            fmt.Println(v4)
        }
    }
    for fi17 := 0; fi17 < 1; fi17++ {
        fmt.Println(v2)
        _ = fi17
    }
    var v18 uint16 = v3
    if ((3.14159 / 1e10) < 1.5) {
        fmt.Println(v4)
        v19 := []int{}
        _ = v19
    } else {
        delete(v5, (v1 + (v1 >> ((65535) & 63))))
        v20 := GenS4{F0: "a"}
        _ = v20
    }
    for fi21 := 0; fi21 < 8; fi21++ {
        var v22 int16 = genFn1(4294967295, v1)
        _ = fi21
        _ = v22
    }
    v3 = v18
    v23 := GenS3{F0: genFn1(166, uint32(v18))}
    var v24 uint = 0
    if true {
        v25 := GenS1{F0: (true || false)}
        v26 := GenS2{F0: 193, F1: true}
        _ = v25
        _ = v26
    } else {
        switch "a" {
        case "hello":
            for fi27 := 0; fi27 < 7; fi27++ {
                v28 := GenS1{F0: (((-2.25 - 0.1) - float64(v2)) < 1e10)}
                v4 = (((v4 - 1.5) + v4) - -2.25)
                _ = fi27
                _ = v28
            }
        case "\n\t":
            v29 := GenS3{F0: int16(v3)}
            _ = v29
        default:
            v30 := GenS4{F0: "a"}
            _ = v30
        }
    }
    v18 -= v3
    for fi31 := 0; fi31 < 12; fi31++ {
        var v32 int = (fi31 * fi31)
        delete(v5, v1)
        v33 := []int{}
        _ = fi31
        _ = v32
        _ = v33
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v18
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
