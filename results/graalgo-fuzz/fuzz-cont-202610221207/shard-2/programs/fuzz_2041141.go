package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string }
type GenS2 struct { F0 uint8; F1 uint64; F2 int64 }
type GenS3 struct { F0 string; F1 int32; F2 bool }
type GenS4 struct { F0 uint16 }
func genFn0() int8 {
    v7 := GenS1{F0: 1, F1: ((("\n\t" + "hello") + "a") + "\n\t")}
    if false {
        var v8 int16 = 0
        _ = v8
    } else {
        var v9 float64 = (1e10 - 0.1)
        var v10 uint8 = 205
        fmt.Println(v9, v9)
        _ = v10
        _ = v9
    }
    v11 := GenS2{F0: 2, F1: 315, F2: 0}
    v12 := GenS3{F0: ("" + "hello"), F1: -1, F2: (0 >= 255)}
    _ = v11
    _ = v12
    _ = v7
    return -128
}
func genFn1(p0 int8, p1 uint16, p2 int16) int32 {
    switch "hello" {
    case "":
        if (0 <= 0) {
            if (2 > 42) {
                v13 := []int{4294967295, 32768}
                _ = v13
            }
            for fi14 := 0; fi14 < 16; fi14++ {
                switch fi14 {
                case 135:
                    v15 := GenS4{F0: p1}
                    _ = v15
                case 980:
                    var v16 string = "hello"
                    _ = v16
                default:
                    var v17 float64 = ((-2.25 * 0.0) - 1.5)
                    _ = v17
                }
                if (!(((3.14159 - (float64(float32(0.1)) * 1e10)) <= 1.5) && false)) {
                    v18 := GenS2{F0: 255, F1: 1, F2: 1}
                    _ = v18
                }
                var v19 uint8 = 58
                _ = fi14
                _ = v19
            }
        } else {
            for fi20 := 0; fi20 < 9; fi20++ {
                var v21 uint8 = 127
                _ = fi20
                _ = v21
            }
        }
    }
    p2 = (p2 * 496)
    if (!(p0 > genFn0())) {
        v22 := []int{-9223372036854775808, 42, 35, int(p2)}
        p2 = p2
        _ = v22
    } else {
        if false {
            p2 = (p2 << ((26) & 63))
            var v23 int = 4294967295
            switch v23 {
            case 42:
                switch ("\n\t" + ("\n\t" + (("\n\t" + "") + ("hello" + ("\n\t" + "foo bar"))))) {
                case "":
                    var v24 int32 = -2147483648
                    _ = v24
                case "a":
                    for fi25 := 0; fi25 < 16; fi25++ {
                        for fi26 := 0; fi26 < 6; fi26++ {
                            v23 -= v23
                            _ = fi26
                        }
                        _ = fi25
                    }
                case "\n\t":
                    fmt.Println(p0)
                default:
                    switch -602 {
                    case -306:
                        v27 := GenS2{F0: 255, F1: 0, F2: -830}
                        _ = v27
                    case 32767:
                        p0 -= 73
                    }
                }
            default:
                if true {
                    var v28 uint = uint(p1)
                    _ = v28
                }
            }
            _ = v23
        }
        if false {
            fmt.Println(p1, p2)
            fmt.Println(p0, p0)
        }
    }
    if true {
        switch -707 {
        case 42:
            for fi29 := 0; fi29 < 12; fi29++ {
                p2 -= p2
                p2 += -862
                _ = fi29
            }
        }
        v30 := GenS3{F0: "\n\t", F1: 42, F2: ((true || false) && false)}
        p0 = (^(p0))
        _ = v30
    }
    _ = p0
    _ = p1
    _ = p2
    return -953
}

func main() {
    v1 := []uint32{526}
    var v2 int64 = -115
    var v3 uint16 = 1
    v4 := make(map[int]uint32)
    var v5 int8 = 32
    v6 := make(map[int]string)
    var v31 int64 = (v2 * v2)
    var v32 uint16 = (v3 >> ((4294967295) & 63))
    var v33 bool = false
    v33 = ((!v33) && false)
    v2 = 0
    var v34 float32 = -2.25
    v33 = (!(float64(v34) <= float64(0.0)))
    v35 := GenS1{F0: 2, F1: "\n\t"}
    fmt.Println(v3, v31)
    for fi36 := 0; fi36 < 18; fi36++ {
        var v37 uint32 = 65535
        for fi38 := 0; fi38 < 3; fi38++ {
            fmt.Println(v5, v5)
            var v39 string = "foo bar"
            _ = fi38
            _ = v39
        }
        v2 = (v31 + (v2 << ((440) & 63)))
        _ = fi36
        _ = v37
    }
    v40 := v4[0]
    if len(v1) > 0 {
        v1[0] = (v40 ^ (v40 % ((v40) | 1)))
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
