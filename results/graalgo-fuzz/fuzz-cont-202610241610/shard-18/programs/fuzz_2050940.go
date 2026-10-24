package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int; F1 int16; F2 int }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 bool; F1 uint8 }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 int }
func genFn0(p0 uint32, p1 uint, p2 int32) uint32 {
    p2 = p2
    p0 -= 742
    for fi5 := 0; fi5 < 7; fi5++ {
        v6 := GenS1{F0: (fi5 ^ int(p0))}
        for fi7 := 0; fi7 < 5; fi7++ {
            v8 := []int{fi7}
            p0 -= (^(p0))
            if (false || (true || (!(float32(-2.25) == (((0.0 * (float32(0.0) - 1.5)) * 1.5) + (0.1 - 0.0)))))) {
                switch (fi7 - (fi7 >> ((p1) & 63))) {
                case 32767:
                    v9 := GenS1{F0: -20}
                    _ = v9
                }
                p2 = 669
                if false {
                    var v10 int = fi5
                    fmt.Println(p1, p2, p1)
                    _ = v10
                }
                var v11 int8 = 105
                _ = v11
            }
            if false {
                v12 := GenS1{F0: fi7}
                _ = v12
            } else {
                p1 = p1
                var v13 uint32 = 2147483648
                var v14 int16 = -944
                _ = v13
                _ = v14
            }
            _ = fi7
            _ = v8
        }
        _ = fi5
        _ = v6
    }
    for fi15 := 0; fi15 < 3; fi15++ {
        var v16 int64 = -9223372036854775808
        switch -2147483648 {
        case 369:
            fmt.Println(p1)
        case 175:
            switch -9223372036854775808 {
            case 128:
                v17 := []int{(fi15 >> ((0) & 63)), (^(fi15)), fi15}
                _ = v17
            case 0:
                v18 := []int{(fi15 - 4294967295), (fi15 / (((fi15 & fi15)) | 1))}
                _ = v18
            }
        case -470:
            v19 := GenS2{F0: (v16 * 42)}
            _ = v19
        }
        _ = fi15
        _ = v16
    }
    _ = p0
    _ = p1
    _ = p2
    return 1
}

func main() {
    var v1 bool = ((true || false) || (true || true))
    v2 := []int{42, 42}
    var v3 string = "hello"
    var v4 uint32 = 717
    var v20 int32 = 2147483647
    for fi21 := 0; fi21 < 14; fi21++ {
        v20 = 42
        _ = fi21
    }
    for fi22 := 0; fi22 < 17; fi22++ {
        for fi23 := 0; fi23 < 20; fi23++ {
            if (true && (v1 || v1)) {
                v20 = (v20 >> ((255) & 63))
                var v24 int32 = v20
                var v25 float32 = (-2.25 / -2.25)
                _ = v24
                _ = v25
            } else {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                v4 += v4
                v3 = (v3 + v3)
            }
            if (255 > 255) {
                var v26 int16 = -32768
                if false {
                    if (1 != -431) {
                        if (!v1) {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            fmt.Println(v20, fi22, fi22)
                        }
                        if true {
                            v27 := GenS1{F0: -1}
                            switch "hello" {
                            case "hello":
                                if v1 {
                                    v28 := GenS3{F0: 128}
                                    v20 = (v20 << ((uint(v4)) & 63))
                                    if (false && (("\n\t" + (v3 + (((v3 + v3) + ((("\n\t" + v3) + v3) + (v3 + ""))) + (v3 + v3)))) < (v3 + (v3 + "foo bar")))) {
                                        if (514 <= 42) {
                                            v26 += v26
                                            var v29 uint32 = 2147483648
                                            var v30 int16 = -192
                                            v30 += v30
                                            if v1 {
                                                var v31 string = "\n\t"
                                                var v32 float64 = -2.25
                                                _ = v31
                                                _ = v32
                                            }
                                            _ = v29
                                            _ = v30
                                        }
                                    }
                                    _ = v28
                                } else {
                                    var v33 string = "\n\t"
                                    if (!(v33 == "a")) {
                                        if (42 > 2) {
                                            fmt.Println(v3, v33)
                                            if len(v2) > 0 {
                                                v2[0] = 1
                                            }
                                        } else {
                                            v34 := GenS2{F0: 654}
                                            if len(v2) > 0 {
                                                v2[0] = (fi23 ^ (fi23 % (((fi23 >> ((414) & 63))) | 1)))
                                            }
                                            _ = v34
                                        }
                                        if len(v2) > 0 {
                                            v2[0] = (-(fi22))
                                        }
                                        var v35 float64 = 0.1
                                        _ = v35
                                    }
                                    v2 = append(v2, -1)
                                    _ = v33
                                }
                            case "":
                                var v36 uint = 42
                                _ = v36
                            case "foo bar":
                                fmt.Println(fi23, v26)
                            default:
                                if true {
                                    v4 = (v4 | (^(v4)))
                                    v26 = 1000
                                }
                            }
                            _ = v27
                        }
                    } else {
                        fmt.Println(v1, fi23, v3)
                    }
                    v2 = append(v2, 42)
                    v37 := GenS2{F0: 0}
                    v4 -= genFn0(v4, 65535, v20)
                    _ = v37
                }
                v20 -= v20
                _ = v26
            }
            _ = fi23
        }
        var v38 uint8 = 35
        _ = fi22
        _ = v38
    }
    switch 32768 {
    case -546:
        var v39 uint64 = 808
        _ = v39
    case 1:
        v1 = false
    case 2147483647:
        if v1 {
            fmt.Println(v20, v3, v1)
        } else {
            switch v3 {
            case "\n\t":
                if (v1 || v1) {
                    switch "\n\t" {
                    case "a":
                        v3 = v3
                    default:
                        if v1 {
                            for fi40 := 0; fi40 < 16; fi40++ {
                                var v41 int32 = -1
                                for fi42 := 0; fi42 < 13; fi42++ {
                                    var v43 uint = 0
                                    if v1 {
                                        if v1 {
                                            _ = len(v2)
                                            if ((42 <= 580) && v1) {
                                                var v44 bool = v1
                                                switch 699 {
                                                case -266:
                                                    var v45 string = ("" + "")
                                                    var v46 uint32 = v4
                                                    v47 := GenS3{F0: fi42}
                                                    fmt.Println(v44)
                                                    _ = v45
                                                    _ = v46
                                                    _ = v47
                                                case 0:
                                                    switch (v3 + (v3 + (v3 + (v3 + ("a" + ((((("" + (v3 + (v3 + v3))) + v3) + ("\n\t" + "hello")) + v3) + v3)))))) {
                                                    case "foo bar":
                                                        v48 := GenS4{F0: fi42}
                                                        _ = v48
                                                    case "":
                                                        var v49 bool = false
                                                        if len(v2) > 0 {
                                                            v2[0] = fi42
                                                        }
                                                        _ = v49
                                                    }
                                                }
                                                _ = v44
                                            }
                                        } else {
                                            fmt.Println(v4, v43, fi42)
                                            if (fi40 <= (^(fi40))) {
                                                fmt.Println(v3, v1, v41)
                                                fmt.Println(v41)
                                            }
                                            fmt.Println(v41, v43)
                                        }
                                        v1 = v1
                                    }
                                    var v50 int32 = (v41 * (v20 / (((v20 + 0)) | 1)))
                                    _ = fi42
                                    _ = v43
                                    _ = v50
                                }
                                var v51 int16 = -1
                                _ = fi40
                                _ = v41
                                _ = v51
                            }
                        }
                    }
                    v52 := GenS5{F0: 4294967295, F1: int16(v20), F2: 490}
                    _ = v52
                } else {
                    v53 := GenS3{F0: 32768}
                    var v54 float64 = (3.14159 + 0.1)
                    if v1 {
                        if (114 != 1) {
                            v54 = v54
                            v54 = -2.25
                        } else {
                            var v55 uint = 2
                            var v56 int64 = int64(v20)
                            for fi57 := 0; fi57 < 3; fi57++ {
                                var v58 float32 = (float32((1.5 - v54)) * float32(1e10))
                                _ = fi57
                                _ = v58
                            }
                            _ = v55
                            _ = v56
                        }
                    } else {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    }
                    _ = v53
                    _ = v54
                }
            case "":
                v1 = v1
            default:
                v1 = (!(!(1 > 42)))
            }
            _ = len(v2)
        }
    }
    if v1 {
        _ = len(v2)
    } else {
        if len(v2) > 0 {
            v2[0] = 128
        }
        if true {
            v3 = v3
            var v59 int32 = -846
            _ = v59
        } else {
            for fi60 := 0; fi60 < 12; fi60++ {
                for fi61 := 0; fi61 < 19; fi61++ {
                    v62 := GenS6{F0: v4}
                    var v63 bool = v1
                    _ = fi61
                    _ = v62
                    _ = v63
                }
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = fi60
            }
            var v64 uint64 = 18446744073709551615
            _ = v64
        }
        v65 := GenS7{F0: (926 != 590), F1: 1}
        _ = v65
    }
    for fi66 := 0; fi66 < 12; fi66++ {
        if false {
            if (((fi66 >> ((4294967295) & 63)) > 42) && (!(1 >= 127))) {
                for fi67 := 0; fi67 < 6; fi67++ {
                    v68 := GenS8{F0: v1}
                    v4 -= 597
                    switch fi67 {
                    case 0:
                        v69 := GenS9{F0: fi66}
                        _ = v69
                    }
                    _ = fi67
                    _ = v68
                }
                if len(v2) > 0 {
                    v2[0] = fi66
                }
            }
        } else {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            fmt.Println(v3, v4, v1)
            v1 = ((v4 - (v4 / (((v4 & genFn0(2147483647, uint(v4), v20))) | 1))) > v4)
        }
        v70 := GenS8{F0: v1}
        v4 += (v4 ^ (v4 | (v4 - (^(v4)))))
        _ = fi66
        _ = v70
    }
    fmt.Println(v20, v3)
    v4 = (v4 << ((720) & 63))
    var v71 int16 = -293
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v1 = true
    var v72 int8 = 42
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v71:", v71)
    fmt.Println("v72:", v72)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v71
    _ = v72
}
