package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint8; F1 uint32; F2 int }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 int8; F1 uint32; F2 int64 }
func genFn0() int8 {
    for fi7 := 0; fi7 < 13; fi7++ {
        for fi8 := 0; fi8 < 11; fi8++ {
            var v9 uint = 255
            _ = fi8
            _ = v9
        }
        fmt.Println(fi7)
        if (461 <= 609) {
            var v10 float64 = 1.5
            for fi11 := 0; fi11 < 12; fi11++ {
                if true {
                    var v12 string = "\n\t"
                    var v13 int32 = -436
                    _ = v12
                    _ = v13
                } else {
                    v10 = v10
                    if true {
                        var v14 float32 = (float32(v10) - 3.14159)
                        fmt.Println(v10, v10)
                        switch "foo bar" {
                        case "":
                            var v15 int64 = 9223372036854775807
                            _ = v15
                        case "a":
                            var v16 uint64 = 4294967295
                            _ = v16
                        default:
                            fmt.Println(fi11, fi11)
                        }
                        _ = v14
                    } else {
                        if false {
                            v17 := GenS1{F0: 0, F1: 42}
                            switch (-(fi11)) {
                            case 32768:
                                if true {
                                    var v18 float32 = 1.5
                                    _ = v18
                                }
                            default:
                                fmt.Println(fi7, fi7)
                            }
                            _ = v17
                        } else {
                            var v19 int = fi7
                            var v20 uint16 = 37
                            _ = v19
                            _ = v20
                        }
                    }
                    if ((-52 != 127) && (1.5 != v10)) {
                        v21 := GenS1{F0: 0, F1: 1}
                        var v22 bool = (!false)
                        _ = v21
                        _ = v22
                    } else {
                        v23 := GenS1{F0: -808, F1: 2147483647}
                        _ = v23
                    }
                }
                v10 = ((v10 - (v10 - v10)) + (1.5 * 0.0))
                fmt.Println(fi11, fi11, fi11)
                _ = fi11
            }
            _ = v10
        }
        _ = fi7
    }
    for fi24 := 0; fi24 < 18; fi24++ {
        switch -2147483648 {
        case 128:
            switch "\n\t" {
            case "a":
                var v25 int8 = 127
                _ = v25
            default:
                var v26 string = ""
                _ = v26
            }
        case 9223372036854775807:
            var v27 float64 = 0.0
            _ = v27
        case 55:
            v28 := []int{991, (-(fi24)), (^(fi24)), (fi24 | (fi24 / (((-(fi24))) | 1)))}
            _ = v28
        }
        v29 := []int{fi24, 456, (fi24 + (fi24 << ((4294967295) & 63))), (fi24 ^ 32768)}
        var v30 float64 = -2.25
        _ = fi24
        _ = v29
        _ = v30
    }
    var v31 int32 = 42
    v32 := GenS1{F0: 32767, F1: v31}
    _ = v31
    _ = v32
    return -128
}

func main() {
    var v1 uint = 989
    v2 := []string{"hello", "foo bar", "a", ("foo bar" + "hello")}
    var v3 bool = false
    v4 := []uint32{1, 458, 2, 4294967295}
    var v5 int16 = 32767
    var v6 bool = (9223372036854775807 <= 694)
    for fi33 := 0; fi33 < 18; fi33++ {
        if (1 > 128) {
            if false {
                fmt.Println(fi33, fi33, fi33)
            } else {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                v4 = append(v4, 65535)
                v3 = v3
            }
        } else {
            for fi34 := 0; fi34 < 12; fi34++ {
                if (3.14159 > (0.0 - (((1e10 - (-2.25 + 0.0)) * 1e10) + (float32(float64(-2.25)) - 1.5)))) {
                    if len(v4) > 0 {
                        v4[0] = 42
                    }
                } else {
                    fmt.Println(v3)
                    var v35 uint = v1
                    _ = v35
                }
                v1 = 42
                var v36 uint32 = 55
                _ = fi34
                _ = v36
            }
        }
        if (18446744073709551615 >= 2) {
            if v3 {
                if (!(2147483647 < 511)) {
                    var v37 int8 = genFn0()
                    v38 := GenS2{F0: (v1 << ((v1) & 63))}
                    _ = v37
                    _ = v38
                }
                v39 := GenS3{F0: 655}
                v6 = (-9223372036854775808 > -359)
                _ = v39
            } else {
                for fi40 := 0; fi40 < 6; fi40++ {
                    if true {
                        v1 -= (^(v1))
                    } else {
                        v41 := GenS2{F0: v1}
                        var v42 uint8 = 1
                        v3 = (!v3)
                        v43 := GenS4{F0: v42, F1: 1, F2: fi33}
                        _ = v41
                        _ = v42
                        _ = v43
                    }
                    v2 = append(v2, "\n\t")
                    _ = fi40
                }
                v4 = append(v4, 4294967295)
            }
        } else {
            fmt.Println(v1)
            switch "foo bar" {
            case "":
                v3 = (!true)
            default:
                var v44 int8 = 42
                _ = v44
            }
        }
        _ = fi33
    }
    v6 = (!v6)
    if (0 < 78) {
        v45 := GenS4{F0: 127, F1: 65535, F2: 337}
        v1 -= (v1 - 255)
        _ = v45
    }
    if ((0 != 128) || v3) {
        for fi46 := 0; fi46 < 8; fi46++ {
            v5 = (v5 % (((v5 | v5)) | 1))
            for fi47 := 0; fi47 < 2; fi47++ {
                switch (fi47 ^ (fi47 ^ (fi47 + fi47))) {
                case 127:
                    var v48 uint8 = 255
                    _ = v48
                default:
                    var v49 bool = v6
                    _ = v49
                }
                var v50 int32 = 8
                v51 := GenS5{F0: (v1 | 851)}
                _ = fi47
                _ = v50
                _ = v51
            }
            v52 := GenS6{F0: 127, F1: 831, F2: 9223372036854775807}
            _ = fi46
            _ = v52
        }
    } else {
        v53 := GenS1{F0: (v5 ^ (v5 >> (((v1 | 42)) & 63))), F1: 183}
        var v54 uint8 = 81
        v5 -= (v5 << ((v1) & 63))
        _ = v53
        _ = v54
    }
    var v55 float32 = 1e10
    var v56 int16 = (^(v5))
    v5 = v56
    if true {
        if (255 != (^(v1))) {
            switch 0 {
            case 2147483647:
                v4 = append(v4, 42)
            default:
                for fi57 := 0; fi57 < 3; fi57++ {
                    v5 -= (v56 * v56)
                    if ("a" >= ("" + "a")) {
                        var v58 uint32 = 42
                        v3 = v3
                        for fi59 := 0; fi59 < 5; fi59++ {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            _ = fi59
                        }
                        _ = v58
                    }
                    _ = fi57
                }
            }
            v56 += (v5 + -1)
            for fi60 := 0; fi60 < 14; fi60++ {
                var v61 uint32 = 0
                _ = fi60
                _ = v61
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v55:", v55)
    fmt.Println("v56:", v56)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v55
    _ = v56
    _ = v6
}
