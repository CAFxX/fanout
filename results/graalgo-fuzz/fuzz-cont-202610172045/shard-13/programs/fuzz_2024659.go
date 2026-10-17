package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint64; F1 bool }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 string; F1 uint }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 bool; F1 int8 }
func genFn0(p0 string, p1 string) uint16 {
    if true {
        var v7 int8 = 1
        v8 := GenS1{F0: (v7 / ((127) | 1))}
        _ = v7
        _ = v8
    }
    fmt.Println(p0, p0)
    fmt.Println(p1, p1)
    p0 = p1
    var v9 int16 = 322
    _ = p0
    _ = p1
    _ = v9
    return 273
}
func genFn1(p0 int, p1 int32) uint32 {
    p0 = 127
    v10 := []int{(p0 >> ((1) & 63)), 386, -456, p0}
    var v11 uint = 65535
    if (int64(p1) > 9223372036854775807) {
        if false {
            switch (p0 << ((v11) & 63)) {
            case -580:
                v11 -= v11
            case 9223372036854775807:
                _ = len(v10)
            default:
                for fi12 := 0; fi12 < 14; fi12++ {
                    if true {
                        p0 -= 42
                        var v13 bool = false
                        _ = v13
                    }
                    for fi14 := 0; fi14 < 12; fi14++ {
                        p1 = p1
                        _ = fi14
                    }
                    _ = fi12
                }
            }
            p0 -= 32768
            var v15 string = ("hello" + "")
            _ = v15
        } else {
            v16 := GenS2{F0: 1}
            if false {
                v10 = append(v10, p0)
                var v17 string = "a"
                for fi18 := 0; fi18 < 13; fi18++ {
                    v19 := GenS2{F0: 801}
                    _ = fi18
                    _ = v19
                }
                _ = v17
            }
            var v20 float32 = float32(0.0)
            _ = v16
            _ = v20
        }
        v10 = append(v10, p0)
        for fi21 := 0; fi21 < 5; fi21++ {
            if true {
                if len(v10) > 0 {
                    v10[0] = 32767
                }
                var v22 uint32 = 2147483648
                if len(v10) >= 2 {
                    v10 = v10[1:len(v10)]
                }
                _ = v22
            } else {
                var v23 float64 = ((float64(((((float32(3.14159) + 1e10) - -2.25) / 1e10) / 1.0)) + float64(1.5)) * 1e10)
                if ((float64(3.14159) - v23) >= 1.5) {
                    p0 = 9223372036854775807
                    _ = len(v10)
                    p0 = (fi21 & p0)
                } else {
                    for fi24 := 0; fi24 < 12; fi24++ {
                        if len(v10) > 0 {
                            v10[0] = fi21
                        }
                        _ = fi24
                    }
                    var v25 uint16 = 42
                    _ = v25
                }
                _ = v23
            }
            _ = fi21
        }
    }
    fmt.Println(p0, v11)
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    return 42
}

func main() {
    var v1 float64 = (1e10 / 1.5)
    v2 := make(map[uint32]string)
    var v3 int = 0
    v4 := make(map[uint32]uint32)
    var v5 int = v3
    var v6 uint16 = 2
    _ = len(v4)
    if (false && (2147483648 > 2)) {
        if ((v1 + v1) <= (-2.25 - v1)) {
            var v26 int32 = -200
            var v27 uint = 658
            _ = v26
            _ = v27
        }
        v28 := []int{v5}
        _ = v28
    } else {
        if (false && true) {
            if (2147483647 < 2147483648) {
                var v29 bool = true
                var v30 bool = (32767 < 0)
                switch "" {
                case "a":
                    for fi31 := 0; fi31 < 5; fi31++ {
                        if true {
                            if (!v30) {
                                for fi32 := 0; fi32 < 16; fi32++ {
                                    v6 = (v6 >> ((42) & 63))
                                    v33 := GenS3{F0: 0, F1: true}
                                    _ = fi32
                                    _ = v33
                                }
                                var v34 string = ("hello" + "hello")
                                fmt.Println(v29, v6, v6)
                                _ = v34
                            }
                            var v35 uint32 = genFn1(v5, 2147483647)
                            _ = v35
                        }
                        _ = fi31
                    }
                case "":
                    if (2147483647 != -487) {
                        fmt.Println(v6, v3)
                        var v36 float32 = 1e10
                        _ = v36
                    }
                case "foo bar":
                    v30 = ((!(((1e10 * ((((1.5 * 1.5) * (0.0 + 1e10)) * ((((3.14159 + float32(((float64(0.0) - (((v1 + (v1 / v1)) - 1.5) / v1)) / float64(1.5)))) / 0.1) * 1e10) + (3.14159 / (float32(v1) / ((float32(1.5) + 0.1) * -2.25))))) / 1.0)) - (-2.25 - 1e10)) >= 0.1)) && (((v6 - (v6 - genFn0(("foo bar" + ("\n\t" + (("" + ("" + "")) + "foo bar"))), ("foo bar" + "foo bar")))) > 255) && (!((-852 != 42) && true))))
                default:
                    v37 := GenS4{F0: ((v3 / (((v3 >> ((0) & 63))) | 1)) > (v3 + v5))}
                    _ = v37
                }
                _ = v29
                _ = v30
            }
            v1 = (v1 * v1)
        }
        var v38 uint64 = 2
        _ = v38
    }
    var v39 uint = 0
    v3 -= v3
    delete(v4, genFn1((v3 << (((v39 + v39)) & 63)), 0))
    for fi40 := 0; fi40 < 6; fi40++ {
        v4[174] = 2147483648
        v3 = (fi40 >> ((v39) & 63))
        for fi41 := 0; fi41 < 12; fi41++ {
            if false {
                v6 += v6
                v2[2] = ("a" + "a")
                v1 = 0.0
            } else {
                fmt.Println(v1)
                var v42 int8 = -36
                switch ("\n\t" + "\n\t") {
                case "hello":
                    v6 += (v6 + (v6 * v6))
                case "":
                    v5 -= (v5 | (v3 >> (((^(v39))) & 63)))
                }
                if false {
                    v39 = (v39 >> ((v39) & 63))
                    fmt.Println(v3)
                    v43 := GenS3{F0: 2, F1: (-251 >= 0)}
                    v3 = (v3 & (v5 * 458))
                    _ = v43
                }
                var v44 uint = (v39 - (v39 + v39))
                _ = v42
                _ = v44
            }
            if true {
                v45 := []int{fi41}
                var v46 uint32 = 2147483647
                _ = v45
                _ = v46
            } else {
                if false {
                    if false {
                        v39 = 18446744073709551615
                        if false {
                            if false {
                                fmt.Println(v39, fi41, v5)
                            } else {
                                if false {
                                    v39 -= v39
                                } else {
                                    v47 := []int{}
                                    _ = v47
                                }
                                if false {
                                    switch (-(fi40)) {
                                    case -150:
                                        var v48 int16 = -948
                                        _ = v48
                                    case 32767:
                                        v6 = (v6 + 32767)
                                    case 0:
                                        v39 = 876
                                    }
                                }
                            }
                            var v49 float64 = (1e10 - v1)
                            _ = v49
                        }
                    }
                } else {
                    if false {
                        switch "\n\t" {
                        case "":
                            v5 = (-(v3))
                        case "a":
                            v39 += v39
                        default:
                            var v50 string = "hello"
                            _ = v50
                        }
                        v39 = 4294967295
                    } else {
                        switch (fi40 / ((-225) | 1)) {
                        case 127:
                            fmt.Println(v5, fi40)
                        case 2147483647:
                            if ("\n\t" != "") {
                                var v51 uint16 = v6
                                v52 := GenS5{F0: ("hello" + ("" + "hello")), F1: uint(v51)}
                                _ = v51
                                _ = v52
                            } else {
                                delete(v4, 2147483647)
                                fmt.Println(v5, v5, v1)
                            }
                        default:
                            v1 += 0.1
                        }
                        switch v5 {
                        case 9223372036854775807:
                            if true {
                                fmt.Println(v39)
                                fmt.Println(v39, v6)
                                if true {
                                    var v53 int32 = -1
                                    _ = v53
                                }
                            }
                        }
                        fmt.Println(v39)
                        v54 := GenS2{F0: 263}
                        _ = v54
                    }
                    v55 := GenS5{F0: "\n\t", F1: v39}
                    _ = v55
                }
            }
            if false {
                var v56 int64 = -1
                _ = v56
            }
            var v57 int32 = -2147483648
            _ = fi41
            _ = v57
        }
        _ = fi40
    }
    if true {
        var v58 int8 = 100
        _ = v58
    } else {
        v1 = v1
        v59 := GenS6{F0: -2147483648}
        if ((!false) && (!((v3 << (((v39 >> (((^(v39))) & 63))) & 63)) > -9223372036854775808))) {
            for fi60 := 0; fi60 < 4; fi60++ {
                if false {
                    _ = len(v4)
                }
                v61 := []int{(^(fi60))}
                _ = fi60
                _ = v61
            }
            if true {
                v62 := []int{v3, (v3 * 1), (v5 + (v5 ^ 32768))}
                for fi63 := 0; fi63 < 20; fi63++ {
                    if ((((("\n\t" + ("" + "")) + ("foo bar" + "\n\t")) + "a") + ("" + "\n\t")) <= ("\n\t" + ("" + "hello"))) {
                        for fi64 := 0; fi64 < 6; fi64++ {
                            v3 = v5
                            var v65 uint8 = uint8(v39)
                            if (v65 <= v65) {
                                v6 -= v6
                                switch (("a" + "") + "foo bar") {
                                case "":
                                    v6 = uint16(v65)
                                case "hello":
                                    if true {
                                        if len(v62) > 0 {
                                            v62[0] = v3
                                        }
                                        var v66 uint16 = genFn0("", "")
                                        fmt.Println(fi63, v6)
                                        _ = v66
                                    }
                                case "\n\t":
                                    switch (fi64 ^ v5) {
                                    case 32767:
                                        var v67 int32 = 2147483647
                                        _ = v67
                                    default:
                                        fmt.Println(v3, v5, fi64)
                                    }
                                }
                            }
                            _ = fi64
                            _ = v65
                        }
                        for fi68 := 0; fi68 < 19; fi68++ {
                            v1 -= v1
                            _ = len(v62)
                            _ = fi68
                        }
                        var v69 int32 = 42
                        _ = v69
                    } else {
                        var v70 int = (v3 * fi63)
                        _ = v70
                    }
                    v3 -= fi63
                    _ = fi63
                }
                _ = v62
            }
            for fi71 := 0; fi71 < 9; fi71++ {
                var v72 uint32 = 65535
                if (1 > genFn1(-9223372036854775808, 42)) {
                    switch (-(v5)) {
                    case 128:
                        v1 = (v1 - 1e10)
                    case -809:
                        var v73 int16 = 42
                        _ = v73
                    case 42:
                        var v74 int8 = 79
                        _ = v74
                    default:
                        v39 = (v39 >> ((v39) & 63))
                    }
                    if (false || true) {
                        fmt.Println(v72, v5)
                        switch (v5 - (v5 & (^(fi71)))) {
                        case 241:
                            if (genFn0("\n\t", ("a" + "\n\t")) <= (v6 / ((42) | 1))) {
                                var v75 uint32 = 0
                                for fi76 := 0; fi76 < 3; fi76++ {
                                    delete(v4, 252)
                                    switch (-(v3)) {
                                    case 4294967295:
                                        fmt.Println(v6)
                                    }
                                    _ = fi76
                                }
                                _ = v75
                            } else {
                                v39 = (v39 / ((uint(v6)) | 1))
                                fmt.Println(v6, v3)
                            }
                        }
                        v77 := GenS7{F0: false, F1: -128}
                        _ = v77
                    }
                } else {
                    v3 = v3
                }
                _ = fi71
                _ = v72
            }
        }
        _ = v59
    }
    var v78 int16 = 1000
    v79 := GenS4{F0: (true || (((false && (((213 != 747) && false) && false)) && (!false)) || true))}
    v78 = 176
    var v80 bool = false
    v81 := GenS2{F0: (v78 * (v78 % ((v78) | 1)))}
    v5 = (-(v5))
    var v82 int16 = v78
    v80 = (2 >= 0)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v78:", v78)
    fmt.Println("v80:", v80)
    fmt.Println("v82:", v82)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v5
    _ = v6
    _ = v78
    _ = v79
    _ = v80
    _ = v81
    _ = v82
}
