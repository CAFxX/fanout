package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32 }
type GenS2 struct { F0 uint32; F1 uint8 }
type GenS3 struct { F0 string; F1 uint8 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 string; F1 uint16 }
func genFn0(p0 uint64, p1 int32, p2 int) string {
    var v5 uint16 = 0
    v6 := GenS1{F0: 2, F1: p1}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return ("foo bar" + "a")
}
func genFn1() uint16 {
    for fi7 := 0; fi7 < 1; fi7++ {
        var v8 int = fi7
        for fi9 := 0; fi9 < 5; fi9++ {
            v8 = (fi9 | (^(v8)))
            v8 += fi9
            _ = fi9
        }
        var v10 bool = ((("\n\t" + ("foo bar" + "")) + "hello") == (("" + "foo bar") + "hello"))
        _ = fi7
        _ = v10
        _ = v8
    }
    if true {
        v11 := []int{}
        var v12 uint16 = 42
        _ = v11
        _ = v12
    } else {
        var v13 bool = true
        _ = v13
    }
    var v14 int8 = 0
    _ = v14
    return 65535
}

func main() {
    var v1 int8 = 127
    v2 := []uint32{4294967295}
    var v3 uint64 = 0
    v4 := make(map[uint32]string)
    v15 := GenS1{F0: 286, F1: 988}
    switch -1 {
    case 738:
        v3 = 276
    default:
        v2 = append(v2, 202)
    }
    if true {
        fmt.Println(v3, v3)
        if (319 != 1) {
            v3 += v3
        }
        v1 = -128
    }
    var v16 int8 = (-(v1))
    var v17 int64 = 363
    if (false || (32767 != 1)) {
        for fi18 := 0; fi18 < 19; fi18++ {
            fmt.Println(v1, v16)
            var v19 int16 = 1
            _ = fi18
            _ = v19
        }
    }
    v16 += (-(v1))
    fmt.Println(v3, v3)
    var v20 float64 = -2.25
    delete(v4, 4294967295)
    var v21 uint = 255
    var v22 bool = (0 != -2147483648)
    switch -95 {
    case -642:
        var v23 uint8 = 193
        _ = v23
    }
    switch "hello" {
    case "":
        v24 := GenS1{F0: 2, F1: 2147483647}
        _ = v24
    case "hello":
        if v22 {
            var v25 int8 = v16
            _ = v25
        } else {
            v26 := GenS1{F0: 412, F1: -1}
            _ = v26
        }
    case "foo bar":
        var v27 uint16 = 32767
        _ = v27
    default:
        v17 += (v17 % ((-9223372036854775808) | 1))
    }
    fmt.Println(v20)
    if v22 {
        for fi28 := 0; fi28 < 17; fi28++ {
            switch "\n\t" {
            case "":
                if (-496 <= 1) {
                    var v29 int32 = 0
                    for fi30 := 0; fi30 < 2; fi30++ {
                        v16 = v16
                        if (true || v22) {
                            v29 -= 345
                            switch (fi30 >> ((v21) & 63)) {
                            case 32768:
                                var v31 int32 = 449
                                _ = v31
                            case 0:
                                fmt.Println(v17, v17)
                            default:
                                var v32 int8 = (v1 | -36)
                                v22 = v22
                                var v33 uint16 = 65535
                                _ = v32
                                _ = v33
                            }
                        } else {
                            if v22 {
                                if true {
                                    if (0.1 != v20) {
                                        v2 = append(v2, 2147483647)
                                    } else {
                                        v22 = ((!((fi28 + (fi30 ^ fi30)) < fi30)) || (v22 && v22))
                                        var v34 int = fi28
                                        v21 = 2
                                        _ = v34
                                    }
                                    v4[370] = ("foo bar" + (("foo bar" + ("hello" + ("foo bar" + ("" + "foo bar")))) + (("hello" + "\n\t") + ("hello" + ("" + ("foo bar" + "\n\t"))))))
                                }
                            }
                            fmt.Println(v17)
                        }
                        v35 := GenS1{F0: 42, F1: v29}
                        if v22 {
                            v36 := GenS1{F0: 0, F1: 42}
                            var v37 bool = (((^(v21)) == (^(v21))) && true)
                            _ = v36
                            _ = v37
                        }
                        _ = fi30
                        _ = v35
                    }
                    _ = v29
                } else {
                    if (172 != 918) {
                        var v38 uint8 = 1
                        v39 := GenS1{F0: 2147483648, F1: 758}
                        if (-291 > int16(v3)) {
                            for fi40 := 0; fi40 < 10; fi40++ {
                                v20 = (3.14159 / ((3.14159 / ((v20 / (float64(1.5) * -2.25)) - (float64(0.1) + v20))) / 1.0))
                                _ = fi40
                            }
                        }
                        _ = v38
                        _ = v39
                    } else {
                        if (v22 || false) {
                            switch "foo bar" {
                            case "hello":
                                if v22 {
                                    v41 := GenS1{F0: 42, F1: 42}
                                    if ((false && (655 < 450)) && true) {
                                        v42 := GenS1{F0: 2147483647, F1: 42}
                                        _ = v42
                                    } else {
                                        var v43 uint8 = 128
                                        if ((((v3 / (((v3 & v3)) | 1)) < v3) || (2147483647 > -1)) && (true || ((fi28 ^ (fi28 << ((4294967295) & 63))) < fi28))) {
                                            if (-2147483648 != -97) {
                                                fmt.Println(v17, v21, v21)
                                                v44 := GenS2{F0: 362, F1: 127}
                                                _ = v44
                                            }
                                        } else {
                                            fmt.Println(v17, v21, v17)
                                            v43 += 42
                                        }
                                        _ = v43
                                    }
                                    v45 := GenS2{F0: 2, F1: 127}
                                    _ = v41
                                    _ = v45
                                }
                            case "\n\t":
                                fmt.Println(v20)
                            case "":
                                if (!(402 != 392)) {
                                    v1 -= v16
                                    v2 = append(v2, 849)
                                }
                            default:
                                var v46 bool = (!(((3.14159 + 3.14159) != v20) || v22))
                                _ = v46
                            }
                            if (false && ((1 >= 0) || (255 == 2))) {
                                if false {
                                    switch 9223372036854775807 {
                                    case 128:
                                        v2 = append(v2, 65535)
                                    case 2147483647:
                                        for fi47 := 0; fi47 < 6; fi47++ {
                                            v4[42] = ("a" + "\n\t")
                                            var v48 int64 = v17
                                            if true {
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                            } else {
                                                v4[907] = "hello"
                                            }
                                            _ = fi47
                                            _ = v48
                                        }
                                    case -9223372036854775808:
                                        v22 = (true && v22)
                                    }
                                    v49 := GenS2{F0: 4294967295, F1: 1}
                                    _ = v49
                                }
                                v50 := GenS1{F0: 0, F1: 0}
                                _ = v50
                            } else {
                                if (!(int64(v21) <= (v17 >> ((v21) & 63)))) {
                                    for fi51 := 0; fi51 < 20; fi51++ {
                                        var v52 float32 = 0.0
                                        _ = fi51
                                        _ = v52
                                    }
                                }
                                for fi53 := 0; fi53 < 18; fi53++ {
                                    fmt.Println(fi53, v3)
                                    v54 := GenS3{F0: ((("hello" + "a") + "foo bar") + "a"), F1: 29}
                                    v55 := v4[42]
                                    _ = fi53
                                    _ = v54
                                    _ = v55
                                }
                                v3 += v3
                            }
                        } else {
                            if v22 {
                                v21 = 610
                                fmt.Println(v20, v21, v16)
                            } else {
                                v3 -= (v3 ^ v3)
                                for fi56 := 0; fi56 < 11; fi56++ {
                                    switch fi28 {
                                    case -2147483648:
                                        v22 = v22
                                    default:
                                        if ((!(float32((0.1 + v20)) <= 1e10)) || v22) {
                                            var v57 int16 = 0
                                            delete(v4, 0)
                                            _ = v57
                                        }
                                    }
                                    if (true || true) {
                                        v22 = (1.5 < -2.25)
                                        var v58 int8 = v16
                                        var v59 float64 = v20
                                        _ = v58
                                        _ = v59
                                    } else {
                                        var v60 int64 = -599
                                        _ = v60
                                    }
                                    v61 := GenS3{F0: "hello", F1: 154}
                                    _ = fi56
                                    _ = v61
                                }
                                var v62 float32 = (3.14159 * (0.0 / ((3.14159 * 1e10) + 1e10)))
                                _ = v62
                            }
                        }
                    }
                }
            case "a":
                var v63 float32 = -2.25
                _ = v63
            case "foo bar":
                for fi64 := 0; fi64 < 6; fi64++ {
                    v22 = (v22 || ("" == "\n\t"))
                    switch 9223372036854775807 {
                    case 432:
                        if len(v2) > 0 {
                            v2[0] = 964
                        }
                    case 32767:
                        v1 += -1
                    default:
                        if v22 {
                            var v65 uint8 = 255
                            if (1 >= 0) {
                                if (2 >= 4294967295) {
                                    switch ((("a" + "foo bar") + "") + ("foo bar" + "foo bar")) {
                                    case "\n\t":
                                        var v66 int = (fi64 << ((v21) & 63))
                                        _ = v66
                                    case "a":
                                        v4[2147483647] = "hello"
                                    default:
                                        v67 := v4[65535]
                                        _ = v67
                                    }
                                } else {
                                    var v68 uint8 = 127
                                    _ = v68
                                }
                                fmt.Println(v16, v3, v16)
                            }
                            _ = v65
                        } else {
                            v17 -= (v17 % ((v17) | 1))
                            var v69 uint64 = (^(v3))
                            if (2147483647 > 1) {
                                v70 := GenS2{F0: 65535, F1: 42}
                                _ = v70
                            } else {
                                v20 -= (3.14159 - 1.5)
                                if (85 >= 42) {
                                    switch (("hello" + "a") + ("" + ("a" + "\n\t"))) {
                                    case "hello":
                                        var v71 int64 = (v17 - (-(v17)))
                                        _ = v71
                                    case "":
                                        v4[uint32(v21)] = ("" + "\n\t")
                                        switch fi64 {
                                        case -1:
                                            var v72 int = fi28
                                            _ = v72
                                        case -9223372036854775808:
                                            switch ("" + ("" + ((((((("hello" + (((((((("" + "hello") + "") + "hello") + "a") + ("foo bar" + "\n\t")) + "foo bar") + "a") + (("hello" + ("a" + "\n\t")) + "hello"))) + "hello") + "\n\t") + "\n\t") + "a") + "\n\t") + ("hello" + "\n\t")))) {
                                            case "a":
                                                v73 := GenS4{F0: 303}
                                                _ = v73
                                            case "foo bar":
                                                delete(v4, 2147483647)
                                            case "":
                                                v74 := GenS3{F0: (("foo bar" + "foo bar") + "a"), F1: 1}
                                                _ = v74
                                            default:
                                                fmt.Println(v20, fi64)
                                            }
                                        default:
                                            var v75 uint64 = v3
                                            _ = v75
                                        }
                                    default:
                                        if false {
                                            switch fi64 {
                                            case 4294967295:
                                                v76 := GenS5{F0: ("" + "a"), F1: 2}
                                                var v77 uint8 = 255
                                                _ = v76
                                                _ = v77
                                            case 1:
                                                var v78 uint8 = 0
                                                _ = v78
                                            }
                                        }
                                    }
                                }
                                if v22 {
                                    var v79 int32 = -110
                                    v79 = (v79 - v79)
                                    _ = v79
                                } else {
                                    v21 -= v21
                                    v20 += ((float64(0.0) / -2.25) + v20)
                                    switch fi28 {
                                    case 127:
                                        switch "foo bar" {
                                        case "":
                                            var v80 string = ("a" + "foo bar")
                                            _ = v80
                                        case "hello":
                                            v22 = (!(v22 && (382 < 1)))
                                        }
                                    default:
                                        var v81 int8 = v16
                                        _ = v81
                                    }
                                    v82 := GenS5{F0: ((("" + ("hello" + (((("\n\t" + "\n\t") + ("" + ("" + ""))) + "") + ((("hello" + ("" + ("" + "\n\t"))) + "\n\t") + ((("a" + "foo bar") + "foo bar") + "\n\t"))))) + "foo bar") + ("" + "hello")), F1: 32767}
                                    fmt.Println(v22, v17)
                                    _ = v82
                                }
                            }
                            var v83 uint16 = 2
                            _ = v69
                            _ = v83
                        }
                    }
                    fmt.Println(v20, v16, v17)
                    _ = fi64
                }
            default:
                var v84 uint16 = genFn1()
                _ = v84
            }
            _ = fi28
        }
    }
    v4[540] = "a"
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
}
