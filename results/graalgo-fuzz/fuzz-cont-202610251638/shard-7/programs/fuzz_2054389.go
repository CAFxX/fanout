package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int8; F1 int16 }
type GenS3 struct { F0 string; F1 string; F2 int8 }
type GenS4 struct { F0 uint64; F1 string; F2 uint64 }
type GenS5 struct { F0 string; F1 uint16; F2 int32 }
type GenS6 struct { F0 uint16; F1 int64; F2 int }
type GenS7 struct { F0 string; F1 uint16; F2 uint8 }
func genFn0(p0 uint, p1 uint16, p2 uint8) int64 {
    p1 -= p1
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = p2
    return -9223372036854775808
}
func genFn1(p0 int32) int {
    for fi4 := 0; fi4 < 12; fi4++ {
        if (!false) {
            var v5 float32 = -2.25
            p0 = (p0 | p0)
            _ = v5
        }
        v6 := GenS1{F0: (fi4 << ((18446744073709551615) & 63))}
        _ = fi4
        _ = v6
    }
    fmt.Println(p0)
    if (2 >= 725) {
        var v7 int64 = 9223372036854775807
        var v8 float32 = float32(((1.5 - 1.5) - float64(-2.25)))
        switch "a" {
        case "foo bar":
            v9 := GenS2{F0: -110, F1: 1000}
            _ = v9
        case "":
            var v10 int16 = 42
            _ = v10
        case "a":
            for fi11 := 0; fi11 < 10; fi11++ {
                for fi12 := 0; fi12 < 1; fi12++ {
                    switch "a" {
                    case "\n\t":
                        p0 -= p0
                    case "":
                        switch (fi11 - (-(fi11))) {
                        case 0:
                            var v13 uint16 = 0
                            _ = v13
                        case -62:
                            switch (fi11 + fi11) {
                            case -716:
                                v7 += genFn0(4294967295, 301, 81)
                            case 4294967295:
                                v8 += v8
                            case -353:
                                v14 := []int{fi11, fi11}
                                _ = v14
                            default:
                                fmt.Println(fi11)
                            }
                        default:
                            v8 = (float32(1.5) + (1e10 + v8))
                        }
                    default:
                        v15 := GenS2{F0: 1, F1: 0}
                        _ = v15
                    }
                    if (255 >= 65535) {
                        var v16 uint64 = 18446744073709551615
                        fmt.Println(v16, fi12)
                        var v17 uint = 255
                        v18 := []int{(fi12 / (((fi11 + -845)) | 1)), 1, (fi11 | (^(fi11)))}
                        _ = v16
                        _ = v17
                        _ = v18
                    } else {
                        v7 = (v7 + -255)
                    }
                    var v19 int64 = v7
                    _ = fi12
                    _ = v19
                }
                fmt.Println(v8, v7)
                _ = fi11
            }
        default:
            v20 := GenS2{F0: 127, F1: 32767}
            _ = v20
        }
        _ = v7
        _ = v8
    } else {
        v21 := []int{-1, 2147483647}
        if (false && (65535 >= 969)) {
            var v22 int8 = -49
            _ = v22
        }
        var v23 int16 = -32768
        _ = v21
        _ = v23
    }
    _ = p0
    return 9223372036854775807
}

func main() {
    var v1 uint64 = 1
    var v2 float64 = (-2.25 / 1.0)
    v3 := []string{}
    switch "foo bar" {
    case "hello":
        var v24 int32 = 2147483647
        _ = v24
    case "":
        _ = len(v3)
    case "\n\t":
        switch ("hello" + ("hello" + "")) {
        case "foo bar":
            v25 := GenS2{F0: -128, F1: 590}
            _ = v25
        default:
            v26 := GenS3{F0: (("" + "foo bar") + "hello"), F1: "a", F2: 42}
            _ = v26
        }
    }
    switch "" {
    case "a":
        var v27 uint16 = 2
        _ = v27
    case "\n\t":
        for fi28 := 0; fi28 < 15; fi28++ {
            var v29 uint64 = v1
            v30 := GenS3{F0: "\n\t", F1: "a", F2: 71}
            _ = fi28
            _ = v29
            _ = v30
        }
    default:
        v1 += (v1 & (v1 | 4294967295))
    }
    var v31 int16 = 1000
    var v32 uint16 = 613
    if false {
        var v33 bool = false
        fmt.Println(v32, v2, v1)
        _ = v33
    } else {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    v1 += (v1 ^ v1)
    for fi34 := 0; fi34 < 19; fi34++ {
        for fi35 := 0; fi35 < 12; fi35++ {
            _ = len(v3)
            _ = fi35
        }
        fmt.Println(v32)
        v1 = v1
        _ = fi34
    }
    switch "foo bar" {
    case "":
        fmt.Println(v2, v32, v32)
    case "foo bar":
        if (false && true) {
            var v36 bool = (v31 == v31)
            _ = v36
        }
    default:
        if (v2 != (float64(1e10) / (v2 / v2))) {
            if (42 != 0) {
                fmt.Println(v31, v1)
            }
            var v37 float64 = 1e10
            _ = v37
        } else {
            v2 = 0.1
            v31 -= -32768
            if len(v3) > 0 {
                v3[0] = "foo bar"
            }
        }
    }
    v38 := GenS4{F0: v1, F1: "\n\t", F2: 2}
    switch genFn1(224) {
    case 9223372036854775807:
        fmt.Println(v32)
    case -1:
        var v39 int32 = 1
        _ = v39
    case -2147483648:
        var v40 float32 = float32(-2.25)
        _ = v40
    default:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    if (4294967295 <= 1) {
        v41 := GenS5{F0: "a", F1: (v32 / (((v32 - 416)) | 1)), F2: 0}
        _ = v41
    } else {
        if len(v3) > 0 {
            v3[0] = ("" + ("hello" + "hello"))
        }
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v42 int16 = v31
    var v43 uint64 = v1
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi44 := 0; fi44 < 17; fi44++ {
        v45 := GenS6{F0: uint16(v1), F1: 1, F2: -219}
        if (!(false || true)) {
            v31 = (v31 % ((v42) | 1))
        } else {
            if true {
                var v46 uint8 = 0
                for fi47 := 0; fi47 < 16; fi47++ {
                    var v48 int32 = int32(v32)
                    v48 += v48
                    switch "\n\t" {
                    case "hello":
                        v49 := GenS5{F0: "foo bar", F1: v32, F2: v48}
                        _ = v49
                    case "foo bar":
                        var v50 string = "hello"
                        _ = v50
                    }
                    if false {
                        v51 := GenS7{F0: "a", F1: v32, F2: v46}
                        _ = v51
                    }
                    _ = fi47
                    _ = v48
                }
                _ = v46
            }
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            v3 = append(v3, "")
        }
        if true {
            v2 = v2
        }
        _ = fi44
        _ = v45
    }
    v32 = (v32 % (((v32 & v32)) | 1))
    var v52 float64 = 3.14159
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v38
    _ = v42
    _ = v43
    _ = v52
}
