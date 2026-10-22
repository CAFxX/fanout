package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64; F2 int8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 int; F1 int; F2 string }
func genFn0(p0 uint8, p1 int64) int16 {
    if true {
        p1 = int64(p0)
    }
    if (-1 <= -1) {
        for fi4 := 0; fi4 < 17; fi4++ {
            var v5 int8 = 127
            v6 := GenS1{F0: 127, F1: int64(p0), F2: (v5 | (v5 * v5))}
            v7 := GenS1{F0: (-(v5)), F1: (-(p1)), F2: (v5 ^ v5)}
            _ = fi4
            _ = v5
            _ = v6
            _ = v7
        }
    } else {
        var v8 int8 = -1
        switch -204 {
        case -1:
            var v9 uint16 = 2
            _ = v9
        default:
            p1 = p1
        }
        p0 -= (p0 - p0)
        _ = v8
    }
    for fi10 := 0; fi10 < 1; fi10++ {
        var v11 uint8 = 1
        _ = fi10
        _ = v11
    }
    p1 += 9223372036854775807
    _ = p0
    _ = p1
    return -1
}

func main() {
    v1 := []int{-365, 128, 2147483647, 2147483647}
    v2 := make(map[int]uint32)
    v3 := []uint32{}
    v3 = append(v3, 2147483647)
    if (false && ((false || (!false)) || true)) {
        var v12 int = -9223372036854775808
        v12 = v12
        _ = v12
    } else {
        var v13 int8 = 0
        var v14 uint = 1
        v15 := GenS2{F0: uint64(v14)}
        _ = v13
        _ = v14
        _ = v15
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    delete(v2, 370)
    v16 := GenS3{F0: 878}
    switch 127 {
    case 761:
        for fi17 := 0; fi17 < 8; fi17++ {
            switch "" {
            case "a":
                var v18 uint = 42
                _ = v18
            }
            _ = fi17
        }
    case -1:
        var v19 int32 = -1
        _ = v19
    default:
        var v20 int32 = 0
        _ = v20
    }
    if true {
        v21 := GenS2{F0: 614}
        _ = v21
    }
    if true {
        if true {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
        var v22 bool = (false && (!false))
        if (255 != 79) {
            var v23 string = "foo bar"
            v24 := GenS4{F0: 513}
            _ = v23
            _ = v24
        } else {
            switch "foo bar" {
            case "hello":
                var v25 uint8 = 127
                _ = v25
            case "foo bar":
                v22 = v22
            case "\n\t":
                for fi26 := 0; fi26 < 6; fi26++ {
                    fmt.Println(v22, v22)
                    _ = fi26
                }
            default:
                v22 = v22
            }
            var v27 uint8 = 64
            if (2 <= v27) {
                if true {
                    fmt.Println(v22, v27)
                    switch -907 {
                    case 643:
                        var v28 uint8 = v27
                        _ = v28
                    case 42:
                        fmt.Println(v22, v27)
                    }
                    _ = len(v2)
                } else {
                    var v29 int16 = genFn0(v27, -1)
                    _ = v29
                }
                var v30 float64 = float64(-2.25)
                switch "a" {
                case "":
                    v31 := GenS2{F0: 1}
                    _ = v31
                }
                _ = v30
            }
            _ = v27
        }
        _ = v22
    } else {
        var v32 uint = 224
        _ = v32
    }
    var v33 int64 = 0
    v34 := GenS1{F0: 1, F1: (v33 % ((v33) | 1)), F2: -47}
    if false {
        for fi35 := 0; fi35 < 17; fi35++ {
            if (!(true && ((1.5 > float32(1.5)) || false))) {
                if (2147483647 >= 2147483647) {
                    v33 -= v33
                    var v36 string = "a"
                    _ = v36
                }
            }
            switch ((("hello" + ("a" + "hello")) + "") + "foo bar") {
            case "foo bar":
                for fi37 := 0; fi37 < 19; fi37++ {
                    _ = len(v2)
                    _ = fi37
                }
            case "":
                v2[(fi35 + (-(fi35)))] = 332
            case "hello":
                v38 := GenS5{F0: true}
                _ = v38
            default:
                v33 = (v33 ^ v33)
            }
            _ = fi35
        }
        var v39 uint8 = 0
        _ = v39
    }
    v40 := GenS4{F0: 65535}
    v33 = -1
    switch "hello" {
    case "\n\t":
        if false {
            v33 = (v33 * (v33 - v33))
        } else {
            var v41 float64 = -2.25
            if true {
                v42 := GenS6{F0: 2147483647, F1: 32767, F2: "foo bar"}
                if (!true) {
                    var v43 string = "hello"
                    if (false && ((v41 - 3.14159) > (v41 * v41))) {
                        var v44 int32 = -617
                        v41 += v41
                        var v45 string = "a"
                        _ = v44
                        _ = v45
                    } else {
                        v1 = append(v1, 32768)
                        var v46 int = 858
                        _ = v46
                    }
                    _ = v43
                } else {
                    if (!(-38 <= 42)) {
                        var v47 uint = 255
                        _ = v47
                    }
                    fmt.Println(v33, v41)
                    v48 := v2[-595]
                    _ = v48
                }
                _ = v42
            }
            _ = v41
        }
    case "":
        var v49 uint8 = 2
        _ = v49
    }
    var v50 uint32 = 969
    var v51 int = -9223372036854775808
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v40
    _ = v50
    _ = v51
}
