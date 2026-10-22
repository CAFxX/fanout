package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int64; F1 uint8; F2 int32 }
type GenS4 struct { F0 int; F1 int8; F2 bool }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int }
func genFn0(p0 string) uint8 {
    v5 := GenS1{F0: 119}
    var v6 uint16 = 32767
    for fi7 := 0; fi7 < 2; fi7++ {
        for fi8 := 0; fi8 < 14; fi8++ {
            v6 += 32767
            _ = fi8
        }
        p0 = (p0 + p0)
        _ = fi7
    }
    if true {
        switch (("a" + ((p0 + p0) + p0)) + (p0 + "hello")) {
        case "\n\t":
            for fi9 := 0; fi9 < 20; fi9++ {
                v10 := []int{}
                _ = fi9
                _ = v10
            }
        }
        p0 = ""
    } else {
        v11 := GenS2{F0: 42}
        var v12 bool = (-115 >= 1000)
        _ = v11
        _ = v12
    }
    v13 := GenS3{F0: -9223372036854775808, F1: 1, F2: 1}
    _ = p0
    _ = v13
    _ = v5
    _ = v6
    return 40
}

func main() {
    v1 := make(map[string]int)
    var v2 uint16 = 150
    v3 := make(map[string]string)
    v4 := []int{32767, 394, 9223372036854775807, 2147483647}
    v14 := v1["\n\t"]
    for fi15 := 0; fi15 < 11; fi15++ {
        if (1e10 == float32(((3.14159 / 1.0) - (0.1 - 1.5)))) {
            fmt.Println(v2)
            var v16 int16 = 727
            _ = v16
        } else {
            for fi17 := 0; fi17 < 9; fi17++ {
                var v18 int8 = 42
                v14 += (-(fi15))
                switch v14 {
                case 128:
                    switch (v14 * v14) {
                    case 550:
                        fmt.Println(v14)
                    case -420:
                        if (((((3.14159 * 0.0) + 3.14159) * ((1.5 + (float32(-2.25) - (0.0 / 1.5))) * 1e10)) != -2.25) || false) {
                            var v19 int32 = -323
                            _ = v19
                        } else {
                            switch fi17 {
                            case 482:
                                v20 := GenS4{F0: 128, F1: (v18 + v18), F2: (131 > 255)}
                                _ = v20
                            }
                            delete(v3, "\n\t")
                            v3["foo bar"] = ""
                        }
                    default:
                        v21 := GenS1{F0: 1}
                        _ = v21
                    }
                    v14 += fi17
                    switch "a" {
                    case "a":
                        switch "\n\t" {
                        case "a":
                            var v22 int16 = -984
                            _ = v22
                        case "foo bar":
                            var v23 uint32 = 0
                            _ = v23
                        case "hello":
                            v18 += v18
                        }
                    }
                    var v24 uint8 = 128
                    fmt.Println(fi17)
                    var v25 int8 = v18
                    _ = v24
                    _ = v25
                default:
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                }
                _ = fi17
                _ = v18
            }
            if (1e10 == 0.0) {
                switch "" {
                case "\n\t":
                    if (!false) {
                        fmt.Println(v14, fi15)
                    } else {
                        v2 += 1
                        v14 = (v14 & v14)
                        v14 += (fi15 * (^(v14)))
                    }
                default:
                    _ = len(v3)
                }
                if true {
                    v2 += (^(v2))
                    for fi26 := 0; fi26 < 17; fi26++ {
                        switch "a" {
                        case "a":
                            switch fi26 {
                            case 9223372036854775807:
                                if true {
                                    var v27 uint = 792
                                    _ = v27
                                } else {
                                    fmt.Println(v2, v2)
                                    v2 = (v2 / ((42) | 1))
                                    v14 -= 500
                                }
                            case 32767:
                                if (!true) {
                                    v14 -= fi26
                                    v28 := GenS3{F0: 42, F1: genFn0("\n\t"), F2: 2147483647}
                                    _ = v28
                                }
                            case 547:
                                switch "a" {
                                case "":
                                    var v29 int8 = int8(fi15)
                                    _ = v29
                                case "a":
                                    var v30 string = ""
                                    _ = v30
                                case "foo bar":
                                    if len(v4) > 0 {
                                        v4[0] = -2147483648
                                    }
                                default:
                                    var v31 int64 = -1
                                    _ = v31
                                }
                            }
                        case "\n\t":
                            var v32 int64 = 460
                            _ = v32
                        case "hello":
                            var v33 uint64 = 42
                            switch (v14 >> ((982) & 63)) {
                            case 9223372036854775807:
                                fmt.Println(v33, v33, v33)
                            }
                            switch 4294967295 {
                            case 929:
                                v14 = (^(v14))
                            case -366:
                                v14 += v14
                                var v34 string = "hello"
                                _ = v34
                            }
                            _ = v33
                        default:
                            var v35 uint64 = 4294967295
                            _ = v35
                        }
                        v14 = 2147483647
                        fmt.Println(v2)
                        _ = fi26
                    }
                }
                var v36 uint = 0
                _ = v36
            }
            if false {
                switch v14 {
                case 874:
                    var v37 uint16 = (v2 << ((18446744073709551615) & 63))
                    _ = v37
                }
                v2 = v2
            } else {
                v14 -= 516
                for fi38 := 0; fi38 < 16; fi38++ {
                    var v39 int32 = -2147483648
                    var v40 uint32 = 2
                    switch "a" {
                    case "a":
                        v41 := GenS3{F0: 9223372036854775807, F1: 128, F2: 0}
                        _ = v41
                    case "":
                        switch (("" + "") + (("foo bar" + "foo bar") + ("a" + (("\n\t" + ("" + "")) + ("\n\t" + "a"))))) {
                        case "foo bar":
                            v42 := GenS3{F0: 1, F1: 162, F2: v39}
                            _ = v42
                        case "hello":
                            v43 := GenS3{F0: 0, F1: genFn0("hello"), F2: (v39 ^ v39)}
                            _ = v43
                        case "a":
                            var v44 uint16 = v2
                            var v45 bool = false
                            v2 += 255
                            _ = v44
                            _ = v45
                        default:
                            switch "hello" {
                            case "a":
                                var v46 float32 = float32(float64(0.1))
                                _ = v46
                            case "":
                                v4 = append(v4, (v14 * 1))
                            }
                        }
                    }
                    _ = fi38
                    _ = v39
                    _ = v40
                }
                var v47 uint = 2
                _ = v47
            }
        }
        for fi48 := 0; fi48 < 9; fi48++ {
            var v49 int64 = -1
            _ = fi48
            _ = v49
        }
        _ = fi15
    }
    for fi50 := 0; fi50 < 15; fi50++ {
        if true {
            var v51 int64 = -9223372036854775808
            _ = v51
        } else {
            var v52 float64 = 1e10
            _ = v52
        }
        _ = fi50
    }
    var v53 int8 = 100
    v54 := GenS5{F0: -2147483648}
    var v55 int8 = 1
    v56 := GenS5{F0: 1}
    if false {
        var v57 bool = true
        var v58 uint16 = 768
        _ = v57
        _ = v58
    } else {
        v53 += -1
        v59 := GenS6{F0: -1}
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = v59
    }
    v14 = (v14 | (-(v14)))
    var v60 uint16 = (v2 % ((v2) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v53:", v53)
    fmt.Println("v55:", v55)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v53
    _ = v54
    _ = v55
    _ = v56
    _ = v60
}
