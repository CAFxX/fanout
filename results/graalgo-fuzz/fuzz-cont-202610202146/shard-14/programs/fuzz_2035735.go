package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int; F1 int8 }
type GenS4 struct { F0 uint16; F1 int8; F2 bool }
func genFn0(p0 uint64, p1 int16, p2 int32) int8 {
    p0 = p0
    v4 := []int{-587, 42}
    var v5 int8 = -128
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return v5
}
func genFn1(p0 int64, p1 int8, p2 int32) int8 {
    p1 += genFn0(42, 1000, (p2 | (^(p2))))
    if false {
        var v6 bool = true
        _ = v6
    }
    v7 := GenS1{F0: (-96 <= p1), F1: 2147483647}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return p1
}

func main() {
    var v1 float32 = 1.5
    v2 := make(map[string]bool)
    var v3 uint32 = 65535
    for fi8 := 0; fi8 < 8; fi8++ {
        switch fi8 {
        case 32768:
            var v9 uint16 = 65535
            _ = v9
        case -1:
            for fi10 := 0; fi10 < 4; fi10++ {
                v1 = ((0.1 - ((v1 / 1.5) / v1)) + ((3.14159 + ((0.0 * ((1.5 * (v1 * v1)) * (0.0 * v1))) / 1e10)) + v1))
                v2[""] = (!false)
                v3 += v3
                _ = fi10
            }
        default:
            var v11 float32 = 0.0
            _ = v11
        }
        _ = fi8
    }
    var v12 uint64 = 410
    v1 += v1
    fmt.Println(v3, v1)
    var v13 uint8 = 158
    fmt.Println(v12, v1, v13)
    if false {
        v12 += v12
        switch "hello" {
        case "":
            switch "foo bar" {
            case "hello":
                for fi14 := 0; fi14 < 4; fi14++ {
                    for fi15 := 0; fi15 < 20; fi15++ {
                        var v16 string = "foo bar"
                        v17 := GenS2{F0: -1}
                        v1 = (v1 * v1)
                        _ = fi15
                        _ = v16
                        _ = v17
                    }
                    v3 -= v3
                    fmt.Println(v12, v3, v1)
                    _ = fi14
                }
            }
        case "\n\t":
            v1 = 0.1
        default:
            v18 := []int{}
            _ = v18
        }
        v19 := GenS1{F0: false, F1: 704}
        _ = v19
    }
    var v20 uint = 255
    var v21 int = -1
    v22 := GenS3{F0: (v21 - v21), F1: -128}
    v20 = v20
    if ((^(v3)) >= v3) {
        var v23 float32 = v1
        _ = v23
    }
    var v24 uint64 = v12
    fmt.Println(v13)
    v25 := GenS4{F0: 2, F1: -1, F2: ((false || false) || (!true))}
    v24 -= (^(v24))
    for fi26 := 0; fi26 < 7; fi26++ {
        fmt.Println(v24, v12)
        _ = fi26
    }
    for fi27 := 0; fi27 < 7; fi27++ {
        fmt.Println(v24)
        _ = fi27
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v24
    _ = v25
    _ = v3
}
