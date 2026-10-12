package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint }
type GenS2 struct { F0 string; F1 bool; F2 int16 }
type GenS3 struct { F0 int32 }
func genFn0() int32 {
    v7 := []int{-895, -1, -718, -9223372036854775808}
    if (!(0.1 <= 0.0)) {
        var v8 string = ""
        switch "" {
        case "a":
            var v9 int8 = 127
            _ = v9
        default:
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
        }
        _ = v8
    } else {
        if true {
            for fi10 := 0; fi10 < 19; fi10++ {
                v7 = append(v7, -154)
                if len(v7) >= 2 {
                    v7 = v7[1:len(v7)]
                }
                _ = fi10
            }
        } else {
            if len(v7) > 0 {
                v7[0] = 32767
            }
            var v11 int32 = -759
            _ = v11
        }
    }
    v7 = append(v7, 111)
    switch 911 {
    case 127:
        var v12 int16 = -32768
        _ = v12
    default:
        var v13 float32 = -2.25
        _ = v13
    }
    _ = v7
    return -395
}
func genFn1(p0 bool, p1 bool, p2 int16) uint32 {
    for fi14 := 0; fi14 < 11; fi14++ {
        fmt.Println(p0, fi14)
        if p0 {
            var v15 uint16 = 710
            _ = v15
        } else {
            for fi16 := 0; fi16 < 13; fi16++ {
                p1 = p1
                _ = fi16
            }
        }
        _ = fi14
    }
    for fi17 := 0; fi17 < 6; fi17++ {
        var v18 float64 = float64(-2.25)
        _ = fi17
        _ = v18
    }
    switch ("" + "foo bar") {
    case "\n\t":
        var v19 uint = 4294967295
        _ = v19
    case "":
        v20 := GenS1{F0: 1, F1: 65535}
        _ = v20
    default:
        p0 = p1
    }
    _ = p0
    _ = p1
    _ = p2
    return 2
}

func main() {
    var v1 int = 517
    var v2 uint16 = 32767
    var v3 float64 = 3.14159
    var v4 uint8 = 1
    v5 := []uint32{4294967295, 42, 53}
    v6 := make(map[uint32]string)
    switch "" {
    case "\n\t":
        var v21 float64 = v3
        _ = v21
    case "":
        var v22 uint64 = 185
        _ = v22
    }
    fmt.Println(v2, v2)
    for fi23 := 0; fi23 < 20; fi23++ {
        if len(v5) > 0 {
            v5[0] = 4294967295
        }
        v5 = append(v5, 2147483647)
        _ = fi23
    }
    var v24 float64 = v3
    for fi25 := 0; fi25 < 8; fi25++ {
        var v26 int16 = 1000
        fmt.Println(v4, v3, v26)
        fmt.Println(v4, v2)
        _ = fi25
        _ = v26
    }
    var v27 uint16 = (^(v2))
    v28 := GenS2{F0: "foo bar", F1: true, F2: 1000}
    var v29 float32 = (1.5 * 1e10)
    if true {
        var v30 int16 = 42
        v31 := GenS3{F0: -1}
        _ = v30
        _ = v31
    }
    var v32 uint = 42
    v29 = (v29 * v29)
    for fi33 := 0; fi33 < 20; fi33++ {
        if false {
            var v34 float64 = float64(3.14159)
            v24 += 0.0
            for fi35 := 0; fi35 < 15; fi35++ {
                if len(v5) > 0 {
                    v5[0] = uint32(v27)
                }
                if false {
                    v1 = (fi35 | (fi35 | (fi35 >> ((18446744073709551615) & 63))))
                    v36 := GenS1{F0: uint8(v32), F1: 0}
                    if (fi35 >= (fi35 ^ v1)) {
                        var v37 uint = 1
                        v32 = (v32 - uint(v4))
                        _ = v37
                    }
                    _ = v36
                } else {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                }
                _ = fi35
            }
            _ = v34
        }
        _ = fi33
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v24
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
