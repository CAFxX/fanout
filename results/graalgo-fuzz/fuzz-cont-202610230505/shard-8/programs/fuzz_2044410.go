package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64 }
type GenS2 struct { F0 uint32; F1 uint32 }
type GenS3 struct { F0 int; F1 int64 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int16; F1 uint16 }
type GenS7 struct { F0 int8; F1 uint16; F2 bool }
type GenS8 struct { F0 int16; F1 uint8; F2 uint64 }
func genFn0(p0 uint16, p1 uint32) bool {
    var v4 int32 = -1
    v5 := []int{-1}
    var v6 uint = 246
    v7 := GenS1{F0: 1, F1: 9223372036854775807}
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return false
}
func genFn1(p0 uint64, p1 int16, p2 uint) bool {
    if true {
        p0 += 42
        for fi8 := 0; fi8 < 15; fi8++ {
            v9 := []int{(fi8 / ((fi8) | 1)), (fi8 + fi8), fi8, fi8}
            _ = fi8
            _ = v9
        }
    }
    p0 = p0
    v10 := []int{4294967295, 744, -809, 2147483647}
    fmt.Println(p1, p2)
    var v11 int32 = 2147483647
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    return false
}

func main() {
    var v1 int32 = 789
    v2 := make(map[uint32]int)
    v3 := make(map[string]string)
    var v12 float32 = float32(((-2.25 * ((0.1 - (1.5 + (3.14159 / 1.0))) * (1e10 / 1.0))) / float64(float32(1e10))))
    if true {
        v13 := GenS2{F0: 65535, F1: 2}
        v14 := GenS3{F0: 4294967295, F1: 9223372036854775807}
        v15 := GenS4{F0: 162}
        _ = v13
        _ = v14
        _ = v15
    } else {
        v3["hello"] = ("" + "foo bar")
        var v16 uint32 = 2147483647
        v17 := GenS2{F0: (v16 >> ((42) & 63)), F1: (v16 * v16)}
        _ = v16
        _ = v17
    }
    fmt.Println(v1)
    if true {
        for fi18 := 0; fi18 < 14; fi18++ {
            if (!true) {
                v19 := []int{(fi18 | (-(fi18))), (fi18 << ((908) & 63)), (-(fi18))}
                _ = v19
            } else {
                var v20 float64 = 0.0
                if (565 <= 42) {
                    switch fi18 {
                    case 32768:
                        for fi21 := 0; fi21 < 12; fi21++ {
                            if (!("hello" != "hello")) {
                                var v22 float32 = v12
                                if true {
                                    if true {
                                        v3["foo bar"] = (("foo bar" + ("foo bar" + "hello")) + "a")
                                        switch -79 {
                                        case 2147483647:
                                            var v23 int64 = int64(v1)
                                            _ = v23
                                        default:
                                            switch fi18 {
                                            case -1:
                                                v20 = float64(1e10)
                                            case -9223372036854775808:
                                                var v24 bool = (!(false && false))
                                                _ = v24
                                            default:
                                                v22 = (1.5 / (v12 + v12))
                                            }
                                        }
                                    } else {
                                        switch fi18 {
                                        case 2147483647:
                                            v1 = (-(v1))
                                        }
                                    }
                                    switch "foo bar" {
                                    case "":
                                        var v25 uint64 = 2
                                        var v26 float32 = v22
                                        _ = v25
                                        _ = v26
                                    case "hello":
                                        var v27 int8 = 0
                                        _ = v27
                                    default:
                                        var v28 uint64 = 0
                                        _ = v28
                                    }
                                    var v29 uint = 255
                                    _ = v29
                                } else {
                                    var v30 bool = true
                                    var v31 string = "\n\t"
                                    v20 = -2.25
                                    _ = v30
                                    _ = v31
                                }
                                v22 += float32((0.1 + 3.14159))
                                _ = v22
                            }
                            if ("foo bar" != "a") {
                                v32 := []int{(fi21 | fi21), (fi18 | (fi18 << ((65535) & 63))), fi21, 360}
                                _ = v32
                            }
                            v12 = (v12 / ((v12 - float32(v20)) * v12))
                            _ = fi21
                        }
                    case 128:
                        var v33 int = fi18
                        _ = v33
                    }
                    v34 := GenS5{F0: 2}
                    for fi35 := 0; fi35 < 3; fi35++ {
                        var v36 int8 = 100
                        var v37 uint16 = 65535
                        v37 = (v37 / ((v37) | 1))
                        _ = fi35
                        _ = v36
                        _ = v37
                    }
                    _ = v34
                } else {
                    switch 4294967295 {
                    case 1:
                        if ("a" >= "hello") {
                            v38 := GenS6{F0: 42, F1: 1}
                            _ = v38
                        }
                    case -466:
                        v20 -= float64((v12 / v12))
                    case 9223372036854775807:
                        var v39 int = 128
                        _ = v39
                    }
                    var v40 int64 = -9223372036854775808
                    _ = v40
                }
                _ = v20
            }
            _ = fi18
        }
    }
    var v41 uint16 = 950
    if true {
        switch "a" {
        case "\n\t":
            if (true || ((1.5 / 0.1) > 3.14159)) {
                v42 := GenS5{F0: 57}
                v43 := []int{32768, 127, -1}
                v44 := GenS4{F0: 248}
                _ = v42
                _ = v43
                _ = v44
            }
        case "a":
            fmt.Println(v41)
        }
        switch "foo bar" {
        case "":
            switch 9223372036854775807 {
            case -209:
                v45 := []int{}
                _ = v45
            case 128:
                fmt.Println(v1, v1, v12)
            case 866:
                v46 := GenS7{F0: 1, F1: (v41 / ((v41) | 1)), F2: true}
                _ = v46
            }
        case "\n\t":
            v47 := GenS4{F0: 127}
            _ = v47
        }
        v48 := GenS6{F0: 1000, F1: (^(v41))}
        _ = v48
    } else {
        var v49 int32 = (-(v1))
        _ = v49
    }
    var v50 int16 = -32768
    var v51 uint = 687
    v52 := GenS8{F0: int16(v1), F1: 2, F2: 1}
    v1 = v1
    fmt.Println(v50, v41)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v41:", v41)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v41
    _ = v50
    _ = v51
    _ = v52
}
