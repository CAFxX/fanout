package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int; F2 string }
type GenS2 struct { F0 uint64 }
func genFn0(p0 uint, p1 uint64, p2 int8) bool {
    p0 = (^(p0))
    switch 32768 {
    case -653:
        fmt.Println(p0, p1)
    case 266:
        v4 := []int{514, -676, 2147483647, -593}
        _ = v4
    default:
        v5 := []int{-1}
        _ = v5
    }
    if false {
        fmt.Println(p1, p2)
        for fi6 := 0; fi6 < 14; fi6++ {
            var v7 string = "a"
            _ = fi6
            _ = v7
        }
        var v8 float32 = 0.1
        _ = v8
    } else {
        var v9 float64 = 0.0
        var v10 uint8 = 143
        _ = v10
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    return false
}

func main() {
    var v1 float32 = 0.0
    var v2 uint8 = 255
    v3 := make(map[int]bool)
    switch 864 {
    case 455:
        v11 := []int{-1, 828, 764, -638}
        _ = v11
    }
    var v12 uint32 = 2
    for fi13 := 0; fi13 < 18; fi13++ {
        switch "a" {
        case "hello":
            if false {
                v3[fi13] = (32767 > 0)
            }
        }
        for fi14 := 0; fi14 < 9; fi14++ {
            if false {
                var v15 int8 = 127
                var v16 int16 = 1000
                v17 := v3[fi13]
                _ = v15
                _ = v16
                _ = v17
            } else {
                v18 := v3[fi14]
                v12 += v12
                fmt.Println(v12)
                v19 := GenS1{F0: v18, F1: (fi13 >> ((65535) & 63)), F2: "a"}
                if (99 == 162) {
                    v1 -= v1
                    fmt.Println(v18, v18)
                    if (!(!((!((v1 / (v1 / ((v1 + v1) * v1))) >= (v1 - v1))) || (!true)))) {
                        switch "hello" {
                        case "":
                            switch "\n\t" {
                            case "hello":
                                v12 = v12
                            case "\n\t":
                                v20 := GenS1{F0: (fi13 == (fi13 / (((fi14 & fi13)) | 1))), F1: fi13, F2: "foo bar"}
                                _ = v20
                            }
                            fmt.Println(v12, v1, v2)
                            var v21 uint8 = (v2 | (v2 & v2))
                            var v22 uint16 = 2
                            v18 = ("hello" < "hello")
                            _ = v21
                            _ = v22
                        case "hello":
                            v12 += 4294967295
                        case "foo bar":
                            v23 := GenS1{F0: (v12 == (v12 & (v12 * v12))), F1: (-(fi14)), F2: "\n\t"}
                            _ = v23
                        default:
                            var v24 int64 = -9223372036854775808
                            _ = v24
                        }
                        switch "foo bar" {
                        case "hello":
                            v12 += (^(v12))
                        case "\n\t":
                            _ = len(v3)
                        default:
                            var v25 uint8 = (v2 - v2)
                            _ = v25
                        }
                    }
                }
                _ = v18
                _ = v19
            }
            v26 := GenS1{F0: (!false), F1: -9223372036854775808, F2: ("foo bar" + ("" + ""))}
            v2 -= v2
            v12 = v12
            _ = fi14
            _ = v26
        }
        _ = fi13
    }
    v3[-1] = (!false)
    v27 := []int{128}
    switch "foo bar" {
    case "hello":
        switch 641 {
        case 1:
            v2 = (v2 | 127)
        default:
            if len(v27) >= 2 {
                v27 = v27[1:len(v27)]
            }
        }
    case "a":
        if len(v27) > 0 {
            v27[0] = 882
        }
    case "":
        var v28 uint = 63
        _ = v28
    default:
        var v29 uint16 = 255
        _ = v29
    }
    v12 = 0
    var v30 uint16 = 255
    for fi31 := 0; fi31 < 5; fi31++ {
        var v32 int64 = 1
        switch fi31 {
        case 2147483647:
            v33 := GenS1{F0: ((1e10 / 3.14159) < 0.1), F1: fi31, F2: "foo bar"}
            _ = v33
        default:
            switch "" {
            case "":
                v32 += v32
            case "hello":
                switch (fi31 - (fi31 >> ((4294967295) & 63))) {
                case -2147483648:
                    v32 -= v32
                case 436:
                    v34 := GenS1{F0: ((-128 >= 1) && (323 != -2147483648)), F1: fi31, F2: "a"}
                    _ = v34
                case 4294967295:
                    if len(v27) >= 2 {
                        v27 = v27[1:len(v27)]
                    }
                }
            case "foo bar":
                var v35 uint = 609
                _ = v35
            }
        }
        _ = fi31
        _ = v32
    }
    var v36 int8 = -95
    v37 := GenS2{F0: 0}
    v36 = -35
    var v38 uint8 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v36:", v36)
    fmt.Println("v38:", v38)
    _ = v1
    _ = v12
    _ = v2
    _ = v27
    _ = v3
    _ = v30
    _ = v36
    _ = v37
    _ = v38
}
