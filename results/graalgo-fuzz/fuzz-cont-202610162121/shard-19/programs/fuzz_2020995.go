package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint; F2 int }
type GenS2 struct { F0 int16; F1 uint16; F2 uint }
type GenS3 struct { F0 uint16; F1 int32; F2 bool }
func genFn0() int16 {
    v7 := []int{-143, 0}
    var v8 uint32 = 42
    fmt.Println(v8)
    switch "hello" {
    case "\n\t":
        v8 += (v8 ^ 286)
    default:
        switch "foo bar" {
        case "hello":
            for fi9 := 0; fi9 < 19; fi9++ {
                fmt.Println(fi9)
                for fi10 := 0; fi10 < 18; fi10++ {
                    if len(v7) > 0 {
                        v7[0] = fi10
                    }
                    if false {
                        var v11 float64 = float64(((float32(1.5) - -2.25) * 0.0))
                        switch (fi9 & fi10) {
                        case 4294967295:
                            var v12 int32 = int32(fi10)
                            _ = v12
                        default:
                            var v13 float64 = v11
                            _ = v13
                        }
                        fmt.Println(fi10, v8, v8)
                        _ = v11
                    } else {
                        v14 := GenS1{F0: -481, F1: 255, F2: fi10}
                        v15 := GenS1{F0: 1, F1: 1, F2: (fi9 / ((1) | 1))}
                        _ = v14
                        _ = v15
                    }
                    _ = fi10
                }
                _ = len(v7)
                _ = fi9
            }
        }
    }
    _ = v7
    _ = v8
    return 1000
}
func genFn1(p0 string) int64 {
    switch p0 {
    case "hello":
        p0 = ("foo bar" + (((((p0 + p0) + p0) + "\n\t") + "hello") + p0))
    default:
        var v16 uint = 42
        _ = v16
    }
    if false {
        v17 := []int{32768}
        var v18 string = (p0 + p0)
        fmt.Println(v18, p0)
        _ = v17
        _ = v18
    }
    for fi19 := 0; fi19 < 17; fi19++ {
        if true {
            if false {
                for fi20 := 0; fi20 < 5; fi20++ {
                    v21 := GenS2{F0: -184, F1: 0, F2: 1}
                    p0 = ((p0 + p0) + (p0 + ("\n\t" + p0)))
                    _ = fi20
                    _ = v21
                }
            } else {
                p0 = "foo bar"
                p0 = p0
                p0 = (p0 + p0)
            }
            p0 = "hello"
            var v22 uint64 = 2
            _ = v22
        } else {
            var v23 string = p0
            _ = v23
        }
        _ = fi19
    }
    var v24 uint32 = 4294967295
    for fi25 := 0; fi25 < 19; fi25++ {
        p0 = p0
        _ = fi25
    }
    _ = p0
    _ = v24
    return 42
}

func main() {
    v1 := []string{}
    var v2 uint8 = 54
    v3 := []uint32{65535, 65535, 0, 4294967295}
    var v4 uint16 = 184
    var v5 float64 = (1.5 + 0.1)
    v6 := make(map[uint32]string)
    var v26 int32 = 2147483647
    v27 := GenS3{F0: (v4 ^ (v4 | v4)), F1: -1, F2: (("\n\t" < "hello") || false)}
    var v28 bool = true
    var v29 int32 = v26
    switch "foo bar" {
    case "a":
        switch 297 {
        case 1:
            var v30 uint8 = 1
            _ = v30
        default:
            fmt.Println(v4)
        }
    default:
        var v31 int64 = -9223372036854775808
        _ = v31
    }
    var v32 uint64 = 18446744073709551615
    v33 := GenS2{F0: 1000, F1: 42, F2: 1}
    switch ("a" + (("hello" + "a") + (("" + "hello") + "\n\t"))) {
    case "hello":
        var v34 uint = 255
        _ = v34
    case "foo bar":
        var v35 string = "\n\t"
        _ = v35
    default:
        for fi36 := 0; fi36 < 12; fi36++ {
            for fi37 := 0; fi37 < 7; fi37++ {
                v2 += 194
                _ = fi37
            }
            _ = fi36
        }
    }
    for fi38 := 0; fi38 < 7; fi38++ {
        _ = len(v6)
        _ = fi38
    }
    v39 := GenS2{F0: 931, F1: (v4 ^ v4), F2: 1}
    var v40 uint16 = v4
    var v41 uint32 = uint32(v4)
    var v42 uint32 = v41
    v2 -= (v2 << ((1) & 63))
    fmt.Println(v4)
    if (v28 && v28) {
        if len(v1) > 0 {
            v1[0] = ("hello" + "")
        }
    } else {
        var v43 uint64 = (v32 + v32)
        _ = v43
    }
    var v44 uint64 = v32
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v32
    _ = v33
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v44
    _ = v5
    _ = v6
}
