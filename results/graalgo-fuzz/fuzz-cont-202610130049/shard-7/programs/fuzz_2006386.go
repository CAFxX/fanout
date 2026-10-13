package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint32; F1 string; F2 int16 }
type GenS4 struct { F0 uint16; F1 int; F2 int }
func genFn0(p0 uint, p1 int64, p2 uint16) string {
    p2 -= p2
    for fi6 := 0; fi6 < 8; fi6++ {
        v7 := GenS1{F0: p0, F1: uint64(p0)}
        _ = fi6
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    return ""
}
func genFn1(p0 string, p1 int64) int32 {
    p0 = p0
    var v8 uint64 = 954
    for fi9 := 0; fi9 < 4; fi9++ {
        var v10 int32 = -2147483648
        if ((4294967295 != 699) && false) {
            v11 := GenS2{F0: 1}
            _ = v11
        }
        v12 := GenS1{F0: 2, F1: v8}
        _ = fi9
        _ = v10
        _ = v12
    }
    switch (p0 + p0) {
    case "a":
        var v13 uint32 = 65535
        _ = v13
    }
    fmt.Println(p0, p0)
    _ = p0
    _ = p1
    _ = v8
    return -136
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[int]uint32)
    v3 := []string{"a", ("\n\t" + (("hello" + "") + ""))}
    v4 := make(map[int]string)
    v5 := make(map[string]uint32)
    v14 := GenS1{F0: 42, F1: 42}
    switch 493 {
    case -9223372036854775808:
        if len(v3) > 0 {
            v3[0] = "\n\t"
        }
    default:
        var v15 int16 = 891
        _ = v15
    }
    var v16 string = "hello"
    switch 9223372036854775807 {
    case 2147483647:
        switch v16 {
        case "":
            var v17 int64 = 42
            _ = v17
        case "a":
            v18 := v2[-1]
            _ = v18
        case "hello":
            var v19 int64 = -9223372036854775808
            _ = v19
        }
    case 128:
        var v20 uint = 702
        _ = v20
    default:
        for fi21 := 0; fi21 < 1; fi21++ {
            v22 := GenS3{F0: 77, F1: v16, F2: -390}
            _ = len(v1)
            for fi23 := 0; fi23 < 1; fi23++ {
                fmt.Println(fi21)
                if ((fi21 - (fi23 * fi21)) != (fi21 | 127)) {
                    switch v16 {
                    case "\n\t":
                        if false {
                            var v24 uint64 = 64
                            var v25 uint = 42
                            var v26 string = ""
                            _ = v24
                            _ = v25
                            _ = v26
                        }
                    case "foo bar":
                        fmt.Println(fi21, fi23, fi21)
                    case "hello":
                        var v27 int16 = 1
                        _ = v27
                    }
                    if (true || true) {
                        var v28 uint64 = 42
                        if false {
                            switch (fi23 >> ((255) & 63)) {
                            case 842:
                                v29 := GenS4{F0: 468, F1: -519, F2: fi23}
                                _ = v29
                            case -1:
                                fmt.Println(fi23)
                            default:
                                if true {
                                    delete(v4, fi21)
                                    switch 781 {
                                    case 42:
                                        fmt.Println(v16, v16, v16)
                                    }
                                    v28 += v28
                                }
                            }
                        }
                        _ = v28
                    }
                }
                if true {
                    v16 = ""
                }
                _ = fi23
            }
            _ = fi21
            _ = v22
        }
    }
    v16 = ("foo bar" + (v16 + "\n\t"))
    var v30 string = v16
    var v31 int8 = -128
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v14
    _ = v16
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
}
