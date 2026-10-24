package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32; F2 int8 }
type GenS2 struct { F0 uint; F1 uint64; F2 int16 }
type GenS3 struct { F0 bool; F1 string; F2 int8 }
type GenS4 struct { F0 uint64; F1 int32 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 bool; F1 uint16 }
func genFn0() uint {
    var v7 uint64 = 0
    v8 := []int{187, 1, 427}
    for fi9 := 0; fi9 < 16; fi9++ {
        v8 = append(v8, (fi9 ^ 9223372036854775807))
        _ = len(v8)
        if len(v8) > 0 {
            v8[0] = (^(fi9))
        }
        _ = fi9
    }
    _ = v7
    _ = v8
    return 158
}
func genFn1(p0 uint, p1 uint32) uint8 {
    v10 := GenS1{F0: (!(255 != 32767)), F1: (p1 / (((p1 << ((p0) & 63))) | 1)), F2: 0}
    for fi11 := 0; fi11 < 10; fi11++ {
        p0 += (p0 % (((^(p0))) | 1))
        _ = fi11
    }
    var v12 bool = true
    _ = p0
    _ = p1
    _ = v10
    _ = v12
    return 255
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[int]string)
    var v3 int32 = 2147483647
    v4 := make(map[uint32]string)
    var v5 int8 = 0
    var v6 uint = 0
    if (0 == 128) {
        var v13 float64 = (3.14159 * 3.14159)
        _ = v13
    } else {
        v6 = 0
        var v14 uint16 = 65535
        _ = v14
    }
    v15 := GenS1{F0: true, F1: 4294967295, F2: v5}
    v3 += v3
    var v16 int8 = int8(v6)
    v17 := GenS1{F0: ((("\n\t" + (("foo bar" + "hello") + "\n\t")) + "a") > "a"), F1: 121, F2: 127}
    if false {
        _ = len(v4)
        if true {
            v18 := []int{}
            _ = v18
        }
    } else {
        v19 := []int{4294967295}
        _ = v19
    }
    v16 += (v5 | (v16 << ((42) & 63)))
    if true {
        if ((v6 % ((65535) | 1)) < 2) {
            if (true || false) {
                v20 := GenS2{F0: 18446744073709551615, F1: 0, F2: 42}
                v21 := GenS2{F0: v6, F1: 483, F2: 190}
                fmt.Println(v6, v6, v3)
                _ = v20
                _ = v21
            } else {
                v5 -= (v5 << (((v6 + v6)) & 63))
            }
            var v22 bool = true
            v23 := GenS3{F0: (v22 && (v22 && ((((1e10 * -2.25) + (0.1 * (1.5 - 1e10))) / 0.1) == 0.1))), F1: ("hello" + ("foo bar" + "")), F2: (v16 * (v16 | 94))}
            _ = v22
            _ = v23
        } else {
            if (-1 > -1) {
                fmt.Println(v6, v16)
                var v24 float64 = (float64(-2.25) / 3.14159)
                for fi25 := 0; fi25 < 2; fi25++ {
                    var v26 int = fi25
                    _ = fi25
                    _ = v26
                }
                _ = v24
            }
        }
        if true {
            if ((v16 * v5) <= v16) {
                for fi27 := 0; fi27 < 11; fi27++ {
                    if false {
                        v6 -= (v6 << ((v6) & 63))
                    } else {
                        v3 = v3
                        v28 := GenS3{F0: true, F1: (("\n\t" + "a") + ("" + ("\n\t" + "\n\t"))), F2: (v16 - -1)}
                        _ = v28
                    }
                    if ((!true) || ("" < "a")) {
                        v3 = -2147483648
                        if (2147483647 < 498) {
                            var v29 uint8 = 255
                            switch fi27 {
                            case -811:
                                var v30 int64 = 0
                                _ = v30
                            }
                            for fi31 := 0; fi31 < 2; fi31++ {
                                fmt.Println(v3)
                                _ = fi31
                            }
                            _ = v29
                        }
                        switch (fi27 >> ((18446744073709551615) & 63)) {
                        case 1:
                            v5 += 1
                        case 72:
                            v32 := GenS4{F0: 4294967295, F1: (v3 % (((v3 % (((v3 * v3)) | 1))) | 1))}
                            _ = v32
                        case 844:
                            var v33 uint16 = 42
                            _ = v33
                        default:
                            switch fi27 {
                            case 911:
                                var v34 uint = v6
                                _ = v34
                            case -2147483648:
                                fmt.Println(v3, v6, v16)
                            case 127:
                                v2[(fi27 / ((fi27) | 1))] = "hello"
                            default:
                                var v35 int16 = 42
                                _ = v35
                            }
                        }
                    }
                    switch fi27 {
                    case -396:
                        if true {
                            v36 := GenS3{F0: false, F1: (("" + "a") + ("" + "foo bar")), F2: (v5 | (v5 * (v16 | v16)))}
                            if false {
                                switch (fi27 << ((v6) & 63)) {
                                case 1:
                                    var v37 uint64 = 42
                                    _ = v37
                                case 32767:
                                    var v38 uint = v6
                                    _ = v38
                                case 42:
                                    var v39 bool = (624 < 32767)
                                    _ = v39
                                default:
                                    v3 = v3
                                }
                                var v40 string = "hello"
                                switch ((v40 + v40) + v40) {
                                case "a":
                                    for fi41 := 0; fi41 < 16; fi41++ {
                                        var v42 uint64 = 944
                                        _ = fi41
                                        _ = v42
                                    }
                                }
                                _ = v40
                            }
                            _ = v36
                        }
                    case 1:
                        switch "hello" {
                        case "hello":
                            var v43 uint8 = 1
                            _ = v43
                        case "\n\t":
                            for fi44 := 0; fi44 < 14; fi44++ {
                                v45 := GenS2{F0: genFn0(), F1: 42, F2: 511}
                                switch fi27 {
                                case -645:
                                    switch "a" {
                                    case "a":
                                        v3 = v3
                                    case "":
                                        v46 := []int{-538, 42}
                                        _ = v46
                                    case "\n\t":
                                        v2[-1] = ""
                                    default:
                                        v47 := GenS5{F0: 1}
                                        _ = v47
                                    }
                                default:
                                    var v48 float32 = 1e10
                                    _ = v48
                                }
                                _ = fi44
                                _ = v45
                            }
                        }
                    case -1:
                        switch fi27 {
                        case 1:
                            for fi49 := 0; fi49 < 4; fi49++ {
                                if ("" >= ((("\n\t" + ("a" + (("hello" + "hello") + (("a" + "hello") + ("" + ("\n\t" + "foo bar")))))) + ("" + "foo bar")) + "")) {
                                    v3 -= -431
                                }
                                v50 := GenS2{F0: v6, F1: 461, F2: -1}
                                if (!false) {
                                    var v51 uint32 = 4294967295
                                    _ = v51
                                }
                                _ = fi49
                                _ = v50
                            }
                        case 9223372036854775807:
                            var v52 int16 = 42
                            _ = v52
                        default:
                            v53 := GenS2{F0: v6, F1: 2, F2: 1}
                            _ = v53
                        }
                    }
                    _ = fi27
                }
            }
        }
    }
    if ((!(-9223372036854775808 != -9223372036854775808)) || (!((^(v3)) == 634))) {
        v54 := GenS2{F0: (v6 | (^(v6))), F1: 1, F2: 963}
        v5 -= v16
        v5 = (^(v16))
        _ = v54
    } else {
        for fi55 := 0; fi55 < 13; fi55++ {
            if true {
                if false {
                    var v56 float32 = 0.1
                    var v57 uint32 = 2147483647
                    var v58 uint32 = v57
                    _ = v56
                    _ = v57
                    _ = v58
                } else {
                    var v59 uint16 = 32767
                    var v60 uint32 = 947
                    v2[fi55] = "foo bar"
                    _ = v59
                    _ = v60
                }
                var v61 uint = 255
                switch -1 {
                case 105:
                    v6 += v61
                }
                _ = v61
            } else {
                v5 -= (^(v16))
                var v62 uint64 = 0
                _ = v62
            }
            if true {
                var v63 int8 = (^(v16))
                if true {
                    v64 := GenS5{F0: 82}
                    if ((-(v3)) != (^(v3))) {
                        var v65 int32 = 0
                        _ = v65
                    }
                    _ = v64
                } else {
                    v66 := []int{(fi55 & (fi55 % ((-1) | 1))), fi55}
                    v5 -= int8(v3)
                    if ((v16 | v63) <= 42) {
                        fmt.Println(v16, v6)
                        fmt.Println(fi55, v5)
                    } else {
                        v63 += (v5 >> ((v6) & 63))
                        if (!(-2.25 > (-2.25 + 3.14159))) {
                            if len(v66) > 0 {
                                v66[0] = (fi55 + fi55)
                            }
                            var v67 float32 = float32(((((-2.25 / ((float64(1.5) - (-2.25 + -2.25)) * -2.25)) / 1.0) + (0.1 * 0.1)) / 3.14159))
                            fmt.Println(fi55, v5)
                            _ = v67
                        }
                    }
                    _ = v66
                }
                for fi68 := 0; fi68 < 6; fi68++ {
                    fmt.Println(v63)
                    if true {
                        v69 := []int{}
                        v70 := v1[(("hello" + ("foo bar" + (("a" + "\n\t") + "\n\t"))) + "")]
                        _ = v69
                        _ = v70
                    }
                    v4[2147483647] = ""
                    _ = fi68
                }
                _ = v63
            }
            fmt.Println(v3, v16, fi55)
            _ = fi55
        }
        fmt.Println(v6, v16, v6)
        var v71 uint = (v6 << (((v6 << ((v6) & 63))) & 63))
        _ = v71
    }
    v5 = v5
    v72 := []int{-9223372036854775808, 32768, 42}
    var v73 int32 = v3
    v74 := GenS6{F0: ((genFn0() <= (v6 * 4294967295)) || false), F1: 1}
    v75 := GenS2{F0: (v6 >> ((genFn0()) & 63)), F1: 2, F2: 1}
    v16 -= (v16 >> (((v6 - 1)) & 63))
    switch -129 {
    case -843:
        if ((!(float32(1e10) > 0.1)) && (-835 != 0)) {
            v5 += v16
            v5 = (v16 - -38)
        }
    case 4294967295:
        v76 := GenS6{F0: (v3 == (v73 - 113)), F1: 803}
        _ = v76
    case 127:
        for fi77 := 0; fi77 < 10; fi77++ {
            v72 = append(v72, fi77)
            var v78 int32 = (-(v73))
            v79 := v4[2147483648]
            _ = fi77
            _ = v78
            _ = v79
        }
    default:
        var v80 int32 = -2147483648
        _ = v80
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v72:", len(v72), v72)
    fmt.Println("v73:", v73)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v72
    _ = v73
    _ = v74
    _ = v75
}
