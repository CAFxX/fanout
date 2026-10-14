package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint8 }
func genFn0() uint64 {
    var v5 int64 = -137
    for fi6 := 0; fi6 < 4; fi6++ {
        v7 := GenS1{F0: 1000, F1: "foo bar"}
        var v8 uint32 = 2
        _ = fi6
        _ = v7
        _ = v8
    }
    var v9 bool = false
    v5 -= 9223372036854775807
    var v10 string = ("\n\t" + "")
    _ = v10
    _ = v5
    _ = v9
    return 2
}
func genFn1(p0 uint32) int32 {
    var v11 int = -317
    var v12 int = (v11 << ((18446744073709551615) & 63))
    var v13 int = v12
    _ = p0
    _ = v11
    _ = v12
    _ = v13
    return 1
}

func main() {
    v1 := make(map[uint32]string)
    var v2 uint = 0
    v3 := []string{"\n\t"}
    v4 := make(map[int]int)
    v2 = 453
    v2 -= (v2 % ((522) | 1))
    for fi14 := 0; fi14 < 14; fi14++ {
        if (1e10 >= -2.25) {
            switch "hello" {
            case "hello":
                switch "hello" {
                case "":
                    if true {
                        var v15 uint = v2
                        _ = v15
                    }
                case "hello":
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                }
            case "":
                fmt.Println(v2, v2)
            default:
                v2 -= v2
            }
            var v16 uint = (v2 ^ v2)
            for fi17 := 0; fi17 < 20; fi17++ {
                v16 = (v2 << ((v2) & 63))
                _ = fi17
            }
            _ = v16
        }
        _ = fi14
    }
    fmt.Println(v2)
    var v18 int8 = 100
    var v19 int8 = 103
    if (2147483647 >= 9223372036854775807) {
        if (1 > -1) {
            v19 = (-(v19))
            for fi20 := 0; fi20 < 5; fi20++ {
                v21 := GenS2{F0: -9223372036854775808}
                if false {
                    var v22 uint32 = 286
                    var v23 uint8 = 42
                    _ = v22
                    _ = v23
                } else {
                    v2 = v2
                }
                for fi24 := 0; fi24 < 16; fi24++ {
                    if (false && (false && ((531 < 2147483647) && true))) {
                        v25 := GenS1{F0: -423, F1: "hello"}
                        v3 = append(v3, "\n\t")
                        switch fi24 {
                        case -531:
                            _ = len(v4)
                        case 1:
                            v19 -= (v19 % ((-1) | 1))
                        case 124:
                            fmt.Println(fi20, v19, fi20)
                        }
                        switch 42 {
                        case 32768:
                            switch "a" {
                            case "a":
                                v2 += (^(v2))
                            }
                        case 128:
                            v19 += (-(v19))
                        }
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        _ = v25
                    } else {
                        v26 := GenS2{F0: 9223372036854775807}
                        _ = v26
                    }
                    var v27 int32 = 0
                    v28 := v4[(fi20 & fi24)]
                    _ = fi24
                    _ = v27
                    _ = v28
                }
                _ = fi20
                _ = v21
            }
        } else {
            for fi29 := 0; fi29 < 12; fi29++ {
                for fi30 := 0; fi30 < 17; fi30++ {
                    v18 -= 90
                    if false {
                        if (true || false) {
                            v2 = 255
                            if false {
                                v4[-262] = fi30
                            }
                            v18 -= 100
                        }
                        if true {
                            fmt.Println(v2, fi29)
                            if false {
                                var v31 uint32 = 2147483647
                                _ = v31
                            }
                            _ = len(v3)
                        }
                    }
                    v2 = (^(v2))
                    v18 = (v18 % ((1) | 1))
                    v18 = (v18 - 123)
                    _ = fi30
                }
                v32 := v1[165]
                _ = fi29
                _ = v32
            }
            if (false || ((true && false) || false)) {
                v2 -= (v2 + v2)
            }
        }
        for fi33 := 0; fi33 < 1; fi33++ {
            fmt.Println(fi33, v18, v18)
            fmt.Println(fi33)
            fmt.Println(v2)
            _ = fi33
        }
    } else {
        var v34 string = "hello"
        var v35 int16 = 0
        switch 1 {
        case -2147483648:
            v36 := GenS1{F0: v35, F1: v34}
            _ = v36
        case 128:
            var v37 uint64 = 4294967295
            _ = v37
        case 32767:
            v34 = v34
        default:
            v2 -= 386
        }
        _ = v34
        _ = v35
    }
    var v38 int32 = 891
    v38 = -987
    fmt.Println(v18, v18)
    v3 = append(v3, ("hello" + "hello"))
    for fi39 := 0; fi39 < 6; fi39++ {
        v1[42] = "hello"
        _ = fi39
    }
    for fi40 := 0; fi40 < 18; fi40++ {
        v41 := GenS3{F0: 128}
        v42 := v4[(fi40 & 1)]
        _ = fi40
        _ = v41
        _ = v42
    }
    v18 = 42
    var v43 uint16 = 533
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v43
}
