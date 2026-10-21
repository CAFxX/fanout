package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32 }
type GenS2 struct { F0 int32; F1 uint16 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int8; F1 uint8 }
type GenS5 struct { F0 uint16; F1 int64 }
type GenS6 struct { F0 uint32 }
func genFn0(p0 int, p1 uint8, p2 uint8) uint {
    var v7 bool = true
    for fi8 := 0; fi8 < 7; fi8++ {
        switch ("foo bar" + "a") {
        case "a":
            fmt.Println(p0)
        case "hello":
            var v9 int16 = 32767
            _ = v9
        case "foo bar":
            v10 := []int{}
            _ = v10
        default:
            fmt.Println(v7)
        }
        v11 := []int{(p0 - (p0 - (p0 + p0))), -850}
        _ = fi8
        _ = v11
    }
    var v12 int16 = 391
    var v13 uint64 = 135
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v13
    _ = v7
    return 527
}

func main() {
    var v1 int = 127
    var v2 float64 = (0.0 - 1e10)
    v3 := []int{(v1 / ((v1) | 1)), (v1 ^ (v1 ^ v1)), v1}
    var v4 int8 = -128
    var v5 float32 = -2.25
    v6 := []string{}
    v14 := GenS1{F0: -2147483648, F1: 384}
    switch v1 {
    case 491:
        v1 = v1
    default:
        if ((0 >= 0) || false) {
            v1 = 128
        }
    }
    v3 = append(v3, 32768)
    var v15 uint = 4294967295
    var v16 int8 = v4
    v17 := GenS2{F0: 42, F1: 65535}
    if (!true) {
        v15 = (^(v15))
        for fi18 := 0; fi18 < 11; fi18++ {
            v15 = (^(v15))
            for fi19 := 0; fi19 < 17; fi19++ {
                var v20 float64 = v2
                if (!true) {
                    var v21 uint8 = 130
                    if true {
                        v20 = (0.1 / v2)
                    } else {
                        var v22 int8 = (^(v16))
                        _ = v22
                    }
                    var v23 uint16 = 712
                    var v24 int = fi19
                    v25 := GenS3{F0: -32768}
                    v21 -= 0
                    _ = v21
                    _ = v23
                    _ = v24
                    _ = v25
                } else {
                    fmt.Println(v4, v5)
                    v20 = v2
                    v26 := GenS4{F0: 1, F1: 42}
                    if ((0 >= -1) && (true || (1e10 == v5))) {
                        if false {
                            if len(v3) > 0 {
                                v3[0] = v1
                            }
                            fmt.Println(v4, v16)
                        }
                        var v27 float64 = 1.5
                        _ = v27
                    }
                    _ = len(v3)
                    _ = v26
                }
                v4 += v4
                _ = fi19
                _ = v20
            }
            if false {
                v28 := GenS5{F0: 42, F1: 0}
                _ = v28
            }
            _ = fi18
        }
    }
    var v29 int16 = 636
    for fi30 := 0; fi30 < 18; fi30++ {
        if true {
            fmt.Println(v2, v1, fi30)
        }
        v31 := GenS3{F0: v29}
        for fi32 := 0; fi32 < 13; fi32++ {
            fmt.Println(v5, fi32)
            if ((false && (0 > -2147483648)) && true) {
                switch fi30 {
                case -2147483648:
                    var v33 uint64 = 18446744073709551615
                    _ = v33
                default:
                    fmt.Println(v4, v16, fi32)
                }
                v34 := GenS6{F0: 65535}
                switch ("a" + "") {
                case "hello":
                    v1 = (-(fi30))
                case "\n\t":
                    var v35 bool = (2 < 2)
                    v1 -= fi32
                    switch "hello" {
                    case "a":
                        fmt.Println(v29, v2)
                    }
                    _ = v35
                }
                _ = v34
            } else {
                if len(v3) > 0 {
                    v3[0] = -2147483648
                }
                v6 = append(v6, "foo bar")
                switch (-(v1)) {
                case -9223372036854775808:
                    v3 = append(v3, (fi30 >> ((18446744073709551615) & 63)))
                case 702:
                    switch (v1 | fi32) {
                    case 32768:
                        fmt.Println(v4, v5)
                    case -757:
                        if ((v15 * (v15 | v15)) <= (v15 * genFn0(-213, 128, 255))) {
                            if false {
                                var v36 uint = 2
                                if true {
                                    v6 = append(v6, "foo bar")
                                    var v37 bool = (314 < 552)
                                    var v38 uint64 = 998
                                    _ = v37
                                    _ = v38
                                }
                                v1 = v1
                                if (!(2 > 2147483647)) {
                                    fmt.Println(v1, v5, v5)
                                }
                                var v39 int64 = 1
                                _ = v36
                                _ = v39
                            } else {
                                v40 := GenS3{F0: v29}
                                fmt.Println(v16, v29, fi30)
                                v15 = v15
                                _ = v40
                            }
                            v3 = append(v3, (v1 ^ (fi32 % (((v1 + (v1 * fi30))) | 1))))
                        } else {
                            v16 += 57
                            v41 := GenS4{F0: (v16 ^ 100), F1: 42}
                            var v42 int8 = 100
                            _ = v41
                            _ = v42
                        }
                    case 4294967295:
                        v43 := GenS6{F0: 2147483647}
                        _ = v43
                    }
                case 210:
                    v44 := GenS3{F0: (v29 * 32767)}
                    _ = v44
                }
            }
            v45 := GenS5{F0: 621, F1: -9223372036854775808}
            _ = fi32
            _ = v45
        }
        _ = fi30
        _ = v31
    }
    v15 -= v15
    fmt.Println(v5)
    switch "\n\t" {
    case "\n\t":
        var v46 int64 = -9223372036854775808
        _ = v46
    case "":
        v47 := GenS3{F0: v29}
        _ = v47
    case "foo bar":
        var v48 float64 = 0.0
        _ = v48
    default:
        switch -1 {
        case 465:
            var v49 int8 = (^(v4))
            _ = v49
        case 9223372036854775807:
            var v50 bool = (int64(v29) < -276)
            _ = v50
        }
    }
    for fi51 := 0; fi51 < 14; fi51++ {
        _ = len(v6)
        if false {
            if true {
                if false {
                    if false {
                        if (((v2 - v2) * (v2 / 3.14159)) < ((v2 * v2) / ((v2 / v2) - 0.1))) {
                            fmt.Println(v29, v4, fi51)
                            v52 := GenS1{F0: -751, F1: 1}
                            _ = v52
                        }
                    } else {
                        fmt.Println(v5, v15, v4)
                        if (4294967295 <= 1) {
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                            var v53 int16 = 0
                            if len(v3) > 0 {
                                v3[0] = v1
                            }
                            _ = v53
                        } else {
                            for fi54 := 0; fi54 < 8; fi54++ {
                                var v55 uint32 = 2
                                var v56 uint32 = 4294967295
                                _ = fi54
                                _ = v55
                                _ = v56
                            }
                            switch ("\n\t" + "a") {
                            case "":
                                if (true && false) {
                                    v57 := GenS1{F0: 309, F1: 42}
                                    v15 += 18446744073709551615
                                    _ = v57
                                } else {
                                    for fi58 := 0; fi58 < 5; fi58++ {
                                        if len(v6) >= 2 {
                                            v6 = v6[1:len(v6)]
                                        }
                                        if false {
                                            var v59 bool = (((float32(v2) / 0.1) / v5) != v5)
                                            _ = v59
                                        }
                                        var v60 string = "foo bar"
                                        if (((fi51 | (fi58 - (v1 / (((v1 >> ((genFn0(fi51, 2, 127)) & 63))) | 1)))) < fi58) || (!false)) {
                                            var v61 int32 = -2147483648
                                            v4 = 127
                                            _ = v61
                                        }
                                        _ = fi58
                                        _ = v60
                                    }
                                    if len(v6) > 0 {
                                        v6[0] = "foo bar"
                                    }
                                }
                            }
                        }
                    }
                    if len(v3) > 0 {
                        v3[0] = (fi51 + 4294967295)
                    }
                    var v62 uint = v15
                    _ = v62
                } else {
                    var v63 string = "\n\t"
                    v4 -= v16
                    _ = v63
                }
                for fi64 := 0; fi64 < 2; fi64++ {
                    v15 -= (^(v15))
                    _ = fi64
                }
                v15 = (v15 >> ((255) & 63))
            }
            fmt.Println(v4, v16)
        }
        _ = fi51
    }
    var v65 uint32 = 65535
    for fi66 := 0; fi66 < 20; fi66++ {
        v67 := GenS1{F0: -2147483648, F1: 1}
        v5 -= v5
        _ = fi66
        _ = v67
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v65:", v65)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v65
}
