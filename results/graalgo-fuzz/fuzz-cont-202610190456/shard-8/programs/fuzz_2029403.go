package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16; F2 string }
type GenS2 struct { F0 uint64; F1 bool }
type GenS3 struct { F0 uint64; F1 uint32 }
func genFn0(p0 int8, p1 uint) int32 {
    p1 = (p1 * p1)
    p0 -= (p0 | (p0 & 0))
    if false {
        p0 = (p0 + -128)
    }
    for fi6 := 0; fi6 < 15; fi6++ {
        fmt.Println(fi6)
        _ = fi6
    }
    _ = p0
    _ = p1
    return -1
}
func genFn1(p0 uint) uint64 {
    if true {
        var v7 uint16 = 2
        switch "\n\t" {
        case "a":
            var v8 uint16 = (v7 ^ v7)
            _ = v8
        case "foo bar":
            v9 := []int{32768, 128}
            _ = v9
        case "\n\t":
            v10 := []int{-2147483648, int(v7)}
            _ = v10
        }
        fmt.Println(p0, v7)
        _ = v7
    } else {
        v11 := []int{0}
        switch "a" {
        case "\n\t":
            fmt.Println(p0)
        }
        p0 += 0
        _ = v11
    }
    switch 32767 {
    case 4294967295:
        p0 += (p0 | p0)
    case 42:
        var v12 int32 = 326
        _ = v12
    case -1:
        p0 += p0
    }
    var v13 int32 = 2147483647
    _ = p0
    _ = v13
    return 987
}

func main() {
    var v1 uint = 2
    var v2 uint8 = 65
    v3 := make(map[uint32]bool)
    var v4 int64 = -1
    v5 := make(map[string]string)
    switch "hello" {
    case "foo bar":
        for fi14 := 0; fi14 < 2; fi14++ {
            v2 -= (v2 / ((uint8(v1)) | 1))
            v4 = 0
            v15 := GenS1{F0: -1, F1: 0, F2: "\n\t"}
            _ = fi14
            _ = v15
        }
    case "\n\t":
        var v16 bool = true
        _ = v16
    case "hello":
        for fi17 := 0; fi17 < 19; fi17++ {
            switch fi17 {
            case 2147483647:
                v2 -= (v2 | v2)
            case 0:
                fmt.Println(v2, fi17)
            case 32767:
                fmt.Println(v1)
            default:
                if (!false) {
                    v1 -= (v1 >> (((v1 | v1)) & 63))
                } else {
                    v1 = v1
                }
            }
            if (v1 == 255) {
                fmt.Println(v4, v2, fi17)
                if (65535 < 0) {
                    for fi18 := 0; fi18 < 4; fi18++ {
                        var v19 string = "a"
                        v20 := v5[(v19 + v19)]
                        v21 := GenS2{F0: 42, F1: (!(!(!(2 >= 649))))}
                        _ = fi18
                        _ = v19
                        _ = v20
                        _ = v21
                    }
                    var v22 int16 = 630
                    _ = v22
                }
                v4 -= (v4 & 1)
            } else {
                if (42 != 42) {
                    v23 := []int{}
                    _ = v23
                } else {
                    var v24 bool = ((v2 - 2) < v2)
                    var v25 float64 = 3.14159
                    for fi26 := 0; fi26 < 7; fi26++ {
                        delete(v3, 42)
                        v27 := GenS3{F0: 540, F1: 2}
                        _ = fi26
                        _ = v27
                    }
                    _ = v24
                    _ = v25
                }
            }
            v2 = (^(v2))
            _ = fi17
        }
    default:
        v1 -= (v1 & v1)
    }
    var v28 int64 = v4
    v29 := GenS3{F0: 282, F1: 2147483647}
    var v30 int16 = 985
    var v31 int = -9223372036854775808
    var v32 int32 = 42
    for fi33 := 0; fi33 < 13; fi33++ {
        var v34 int8 = 127
        for fi35 := 0; fi35 < 3; fi35++ {
            v36 := GenS3{F0: 883, F1: 2147483648}
            _ = fi35
            _ = v36
        }
        _ = fi33
        _ = v34
    }
    var v37 int16 = v30
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v37
    _ = v4
    _ = v5
}
