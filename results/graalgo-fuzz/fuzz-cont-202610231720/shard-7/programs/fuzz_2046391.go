package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 string; F1 uint8 }
type GenS3 struct { F0 bool; F1 uint8; F2 uint64 }
type GenS4 struct { F0 int16; F1 uint64; F2 int64 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int32; F1 uint32; F2 int8 }
type GenS7 struct { F0 int8; F1 int64; F2 int }

func main() {
    var v1 string = "a"
    v2 := make(map[uint32]string)
    var v3 int8 = -128
    var v4 uint64 = 0
    v5 := GenS1{F0: 32767}
    if (18446744073709551615 >= 255) {
        var v6 uint16 = 42
        v7 := GenS1{F0: 32767}
        _ = v6
        _ = v7
    } else {
        var v8 int8 = (^(v3))
        v3 += (v8 & -128)
        _ = v8
    }
    if (!((-66 <= -779) || false)) {
        v4 = 875
        if false {
            switch "foo bar" {
            case "foo bar":
                v4 -= 42
            case "\n\t":
                for fi9 := 0; fi9 < 7; fi9++ {
                    var v10 string = (v1 + (v1 + "foo bar"))
                    _ = fi9
                    _ = v10
                }
            case "hello":
                v2[uint32(v4)] = ((((((v1 + v1) + "") + v1) + (v1 + "")) + v1) + ("hello" + v1))
            }
            for fi11 := 0; fi11 < 20; fi11++ {
                v12 := []int{(^(fi11)), (^(fi11)), (fi11 - (fi11 | 32768)), -9223372036854775808}
                _ = fi11
                _ = v12
            }
        }
    }
    v13 := GenS1{F0: -357}
    var v14 string = v1
    var v15 float32 = 0.0
    for fi16 := 0; fi16 < 2; fi16++ {
        if (2147483648 < 2) {
            var v17 int32 = 173
            v17 = (v17 >> ((96) & 63))
            _ = v17
        } else {
            if true {
                v18 := GenS1{F0: -197}
                var v19 bool = true
                v20 := []int{(fi16 - fi16), fi16}
                _ = v18
                _ = v19
                _ = v20
            }
            if false {
                var v21 bool = (true || (v4 <= v4))
                v22 := []int{}
                for fi23 := 0; fi23 < 2; fi23++ {
                    var v24 uint16 = 211
                    var v25 int = fi16
                    _ = len(v22)
                    v26 := GenS1{F0: 1}
                    switch 127 {
                    case -2147483648:
                        v27 := v2[65535]
                        _ = v27
                    }
                    if (((3.14159 > -2.25) || (1 >= 2147483647)) && false) {
                        v4 += uint64(v24)
                        if ("foo bar" > "") {
                            v28 := GenS2{F0: (("a" + v1) + "foo bar"), F1: 226}
                            var v29 int = (v25 - (-(v25)))
                            _ = v28
                            _ = v29
                        } else {
                            v30 := GenS1{F0: 1}
                            if v21 {
                                if true {
                                    v24 = 0
                                    switch v14 {
                                    case "a":
                                        v31 := v2[2147483648]
                                        v32 := GenS1{F0: -246}
                                        _ = v31
                                        _ = v32
                                    }
                                } else {
                                    var v33 uint16 = v24
                                    _ = v33
                                }
                                if v21 {
                                    if len(v22) > 0 {
                                        v22[0] = 42
                                    }
                                } else {
                                    switch v14 {
                                    case "hello":
                                        if v21 {
                                            if v21 {
                                                var v34 int = v25
                                                fmt.Println(v3, v15, v34)
                                                v1 = "\n\t"
                                                _ = v34
                                            }
                                            var v35 int16 = 32767
                                            _ = v35
                                        } else {
                                            switch (v25 >> ((0) & 63)) {
                                            case 9223372036854775807:
                                                switch v25 {
                                                case -1:
                                                    if (v21 || ((v4 / (((v4 & v4)) | 1)) != v4)) {
                                                        if (407 != uint32(v4)) {
                                                            v21 = v21
                                                        } else {
                                                            fmt.Println(v15, v24, v3)
                                                            v36 := GenS2{F0: "", F1: 86}
                                                            if v21 {
                                                                var v37 uint16 = (v24 * 2)
                                                                if v21 {
                                                                    v1 = v14
                                                                    fmt.Println(v3, v24)
                                                                }
                                                                _ = len(v22)
                                                                _ = v37
                                                            }
                                                            _ = v36
                                                        }
                                                        v38 := GenS2{F0: "hello", F1: 128}
                                                        _ = v38
                                                    } else {
                                                        var v39 int64 = -9223372036854775808
                                                        _ = v39
                                                    }
                                                }
                                            }
                                            switch fi23 {
                                            case -2147483648:
                                                v40 := GenS3{F0: (((-2.25 - (v15 / (v15 / (v15 - 3.14159)))) * (v15 - v15)) <= ((3.14159 * (v15 - float32(3.14159))) + 0.0)), F1: 0, F2: v4}
                                                _ = v40
                                            case -685:
                                                v14 = "hello"
                                            default:
                                                v2[2] = (v1 + v14)
                                            }
                                            if true {
                                                fmt.Println(v15)
                                            } else {
                                                delete(v2, 650)
                                                v41 := GenS4{F0: 42, F1: (^(v4)), F2: 9223372036854775807}
                                                if (!(!v21)) {
                                                    v42 := GenS3{F0: ((v24 << ((2) & 63)) == v24), F1: 148, F2: v4}
                                                    v15 -= (1.5 / v15)
                                                    _ = v42
                                                }
                                                _ = v41
                                            }
                                            var v43 string = v14
                                            v44 := GenS5{F0: (fi23 * fi16)}
                                            if false {
                                                fmt.Println(v1, v43)
                                                v4 = (v4 ^ v4)
                                                var v45 float64 = (float64((float32(0.0) - (0.0 / v15))) * 3.14159)
                                                _ = v45
                                            } else {
                                                switch ("" + v1) {
                                                case "\n\t":
                                                    var v46 uint = 255
                                                    _ = v46
                                                }
                                            }
                                            _ = len(v2)
                                            fmt.Println(fi16, v24, v3)
                                            switch v14 {
                                            case "a":
                                                var v47 int = fi16
                                                _ = v47
                                            }
                                            var v48 string = (v43 + (v1 + v43))
                                            var v49 int64 = -9223372036854775808
                                            _ = v43
                                            _ = v44
                                            _ = v48
                                            _ = v49
                                        }
                                    case "foo bar":
                                        fmt.Println(v25, fi23, v25)
                                    default:
                                        v50 := GenS1{F0: 106}
                                        _ = v50
                                    }
                                    var v51 int16 = -1
                                    delete(v2, 2147483648)
                                    _ = v51
                                }
                            }
                            if ((v15 - v15) > 1.5) {
                                v52 := GenS3{F0: (!(904 >= 32767)), F1: 255, F2: v4}
                                _ = v52
                            }
                            v1 = v14
                            v15 += v15
                            v2[0] = "a"
                            if true {
                                if (!((v25 * v25) != (fi16 + fi23))) {
                                    v53 := GenS5{F0: fi23}
                                    v1 = v14
                                    _ = v53
                                }
                                v3 -= -76
                            } else {
                                var v54 uint = 4294967295
                                v2[367] = v14
                                _ = v54
                            }
                            _ = v30
                        }
                        v1 = ((v1 + v14) + v1)
                    } else {
                        v3 = (v3 * (v3 % (((^(v3))) | 1)))
                        if false {
                            var v55 uint64 = (v4 % ((v4) | 1))
                            v55 -= v55
                            _ = len(v2)
                            v1 = ("hello" + v14)
                            _ = v55
                        } else {
                            v21 = (!(1 >= -1))
                        }
                        fmt.Println(v21, v15)
                    }
                    fmt.Println(fi16)
                    var v56 int = 1
                    fmt.Println(v15, v1, v56)
                    _ = fi23
                    _ = v24
                    _ = v25
                    _ = v26
                    _ = v56
                }
                _ = v21
                _ = v22
            } else {
                var v57 uint16 = 32767
                v58 := GenS6{F0: -779, F1: 164, F2: (v3 << ((2) & 63))}
                _ = v57
                _ = v58
            }
            for fi59 := 0; fi59 < 7; fi59++ {
                v60 := GenS7{F0: (v3 << ((4294967295) & 63)), F1: 491, F2: (fi16 ^ fi59)}
                _ = fi59
                _ = v60
            }
        }
        _ = fi16
    }
    v15 = v15
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
