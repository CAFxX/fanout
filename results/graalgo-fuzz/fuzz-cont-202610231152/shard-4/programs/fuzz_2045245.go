package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool }
type GenS2 struct { F0 bool; F1 bool; F2 string }
func genFn0(p0 uint, p1 uint16) int {
    var v7 uint64 = 0
    for fi8 := 0; fi8 < 4; fi8++ {
        for fi9 := 0; fi9 < 13; fi9++ {
            fmt.Println(p0)
            p0 += p0
            switch fi8 {
            case 32767:
                v7 += 795
            case -1:
                var v10 uint16 = (p1 / ((255) | 1))
                _ = v10
            default:
                fmt.Println(p0)
            }
            _ = fi9
        }
        v7 = 18446744073709551615
        _ = fi8
    }
    p1 -= p1
    var v11 float32 = 1.5
    _ = p0
    _ = p1
    _ = v11
    _ = v7
    return 42
}
func genFn1(p0 int16) int8 {
    if false {
        v12 := GenS1{F0: -420, F1: (255 < 2)}
        _ = v12
    }
    p0 += 42
    switch 433 {
    case 4294967295:
        switch "a" {
        case "a":
            var v13 uint = 4294967295
            _ = v13
        case "":
            switch "foo bar" {
            case "\n\t":
                v14 := []int{0, 2147483647}
                _ = v14
            }
        case "\n\t":
            for fi15 := 0; fi15 < 15; fi15++ {
                for fi16 := 0; fi16 < 8; fi16++ {
                    var v17 uint = 0
                    if true {
                        var v18 int16 = (p0 + -117)
                        v19 := GenS1{F0: 0, F1: true}
                        _ = v18
                        _ = v19
                    }
                    _ = fi16
                    _ = v17
                }
                fmt.Println(p0)
                if (331 < 218) {
                    p0 = p0
                    v20 := []int{(fi15 / ((fi15) | 1))}
                    var v21 float64 = 1.5
                    _ = v20
                    _ = v21
                } else {
                    p0 -= p0
                }
                _ = fi15
            }
        default:
            v22 := GenS1{F0: 42, F1: ((false || (32767 <= 176)) && (0 >= p0))}
            _ = v22
        }
    default:
        v23 := GenS1{F0: -741, F1: true}
        _ = v23
    }
    _ = p0
    return -128
}

func main() {
    v1 := []uint32{2147483648}
    v2 := make(map[int]string)
    var v3 int64 = 0
    var v4 bool = (!true)
    var v5 uint64 = 992
    v6 := make(map[uint32]uint32)
    var v24 uint8 = 166
    var v25 uint32 = 2147483647
    v25 = (v25 - (v25 | 42))
    for fi26 := 0; fi26 < 19; fi26++ {
        v3 += 0
        v25 = uint32(v5)
        _ = fi26
    }
    v5 -= 4294967295
    v27 := GenS2{F0: v4, F1: false, F2: ("foo bar" + (("foo bar" + "hello") + ("foo bar" + "foo bar")))}
    fmt.Println(v4)
    var v28 int8 = 127
    for fi29 := 0; fi29 < 10; fi29++ {
        switch 621 {
        case 128:
            v25 = (v25 * 65535)
        case -2147483648:
            switch fi29 {
            case -201:
                var v30 string = "foo bar"
                _ = v30
            case 0:
                var v31 int8 = genFn1(-1)
                _ = v31
            default:
                var v32 uint16 = 65535
                _ = v32
            }
        }
        v33 := GenS2{F0: (!(32767 <= 0)), F1: v4, F2: ("hello" + "hello")}
        if false {
            v3 = v3
            v34 := GenS2{F0: (!true), F1: v4, F2: "a"}
            for fi35 := 0; fi35 < 17; fi35++ {
                var v36 int32 = 48
                var v37 float64 = 3.14159
                _ = fi35
                _ = v36
                _ = v37
            }
            _ = v34
        }
        _ = fi29
        _ = v33
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
