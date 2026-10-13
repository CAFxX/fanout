package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16; F2 int8 }
type GenS2 struct { F0 uint8; F1 int32; F2 bool }
type GenS3 struct { F0 uint32; F1 uint16 }
type GenS4 struct { F0 int32; F1 int8 }
type GenS5 struct { F0 uint8; F1 bool }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 string; F1 uint8; F2 int16 }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 uint16; F1 uint32; F2 uint }
type GenS10 struct { F0 int32; F1 uint64; F2 uint64 }
func genFn0(p0 uint16) uint {
    for fi5 := 0; fi5 < 11; fi5++ {
        fmt.Println(fi5, p0)
        var v6 int = fi5
        _ = fi5
        _ = v6
    }
    if (128 <= 0) {
        v7 := GenS1{F0: "", F1: 606, F2: -1}
        v8 := []int{243, 4294967295, -1, 32768}
        _ = v7
        _ = v8
    } else {
        var v9 int16 = -138
        p0 = p0
        p0 -= 1
        _ = v9
    }
    for fi10 := 0; fi10 < 9; fi10++ {
        v11 := []int{}
        _ = fi10
        _ = v11
    }
    v12 := []int{4294967295, -1, -369}
    for fi13 := 0; fi13 < 2; fi13++ {
        var v14 uint = 42
        v15 := GenS1{F0: (("a" + "") + ("a" + "\n\t")), F1: -83, F2: 0}
        for fi16 := 0; fi16 < 1; fi16++ {
            var v17 string = "a"
            fmt.Println(v17, fi13)
            fmt.Println(fi13, fi16, v17)
            _ = fi16
            _ = v17
        }
        _ = fi13
        _ = v14
        _ = v15
    }
    _ = p0
    _ = v12
    return 86
}
func genFn1(p0 uint8, p1 bool) uint16 {
    p1 = (472 != 18446744073709551615)
    for fi18 := 0; fi18 < 18; fi18++ {
        if (true && (500 == -1)) {
            for fi19 := 0; fi19 < 3; fi19++ {
                p0 = 128
                _ = fi19
            }
            for fi20 := 0; fi20 < 15; fi20++ {
                fmt.Println(p0, p1)
                _ = fi20
            }
            p1 = (int8(fi18) <= -96)
        }
        switch ("foo bar" + "hello") {
        case "foo bar":
            var v21 string = "\n\t"
            _ = v21
        case "a":
            p1 = (!(100 >= 100))
        case "hello":
            var v22 int32 = -1
            _ = v22
        }
        _ = fi18
    }
    _ = p0
    _ = p1
    return 42
}

func main() {
    var v1 uint32 = 408
    v2 := make(map[string]string)
    var v3 float64 = 0.1
    var v4 uint = 590
    if false {
        var v23 int16 = 1000
        switch "\n\t" {
        case "\n\t":
            v24 := GenS1{F0: ("\n\t" + "\n\t"), F1: (v23 >> ((18446744073709551615) & 63)), F2: 0}
            _ = v24
        case "a":
            for fi25 := 0; fi25 < 10; fi25++ {
                for fi26 := 0; fi26 < 8; fi26++ {
                    var v27 uint = genFn0(32767)
                    _ = fi26
                    _ = v27
                }
                if (true || false) {
                    var v28 uint64 = 2
                    v29 := v2[("\n\t" + ((("" + "") + "hello") + (((("a" + "") + "a") + "foo bar") + "hello")))]
                    for fi30 := 0; fi30 < 18; fi30++ {
                        v31 := GenS1{F0: "", F1: v23, F2: 31}
                        _ = fi30
                        _ = v31
                    }
                    _ = v28
                    _ = v29
                } else {
                    v3 -= v3
                    fmt.Println(v1)
                    switch "" {
                    case "foo bar":
                        fmt.Println(v1)
                    }
                }
                if (true || true) {
                    switch (fi25 >> (((v4 >> ((v4) & 63))) & 63)) {
                    case 4294967295:
                        v23 -= -32768
                    case -878:
                        var v32 int8 = 61
                        _ = v32
                    }
                    switch "foo bar" {
                    case "foo bar":
                        v2[""] = ("" + "hello")
                    case "a":
                        for fi33 := 0; fi33 < 4; fi33++ {
                            fmt.Println(v1, fi33)
                            var v34 float64 = v3
                            if (!false) {
                                if true {
                                    var v35 uint32 = 1
                                    v36 := []int{4294967295, -33, -615}
                                    _ = v35
                                    _ = v36
                                } else {
                                    if true {
                                        switch fi25 {
                                        case 127:
                                            fmt.Println(fi33, v4)
                                        case 4294967295:
                                            v37 := GenS2{F0: 128, F1: 2147483647, F2: (216 == 128)}
                                            _ = v37
                                        }
                                    } else {
                                        var v38 float64 = ((v3 * (v3 / v34)) + 3.14159)
                                        if ((true || false) && (v23 < int16(fi25))) {
                                            switch "" {
                                            case "foo bar":
                                                fmt.Println(v38, fi33)
                                            case "hello":
                                                fmt.Println(v3, v38)
                                            }
                                            fmt.Println(v3, fi25)
                                            v39 := []int{(fi25 / ((475) | 1))}
                                            _ = v39
                                        }
                                        _ = v38
                                    }
                                    var v40 int16 = v23
                                    switch (fi25 / ((fi33) | 1)) {
                                    case -2147483648:
                                        v34 = (1.5 / (v3 * v3))
                                    case 32768:
                                        var v41 uint64 = 490
                                        _ = v41
                                    default:
                                        v34 -= float64(float32(1e10))
                                    }
                                    v42 := []int{(fi25 + fi25)}
                                    if (v4 >= (v4 - genFn0(22))) {
                                        v2["hello"] = "a"
                                        fmt.Println(v34)
                                        fmt.Println(v34, v4, v4)
                                    } else {
                                        var v43 float32 = (0.0 / ((-2.25 / 1.0) / -2.25))
                                        v44 := GenS3{F0: v1, F1: 32767}
                                        v1 -= 65535
                                        _ = v43
                                        _ = v44
                                    }
                                    if len(v42) > 0 {
                                        v42[0] = (-(fi25))
                                    }
                                    var v45 uint32 = 0
                                    _ = v40
                                    _ = v42
                                    _ = v45
                                }
                                switch ("hello" + ("foo bar" + "a")) {
                                case "foo bar":
                                    fmt.Println(v1, v23)
                                }
                                var v46 int32 = 0
                                _ = v46
                            } else {
                                v34 = v3
                                v47 := GenS3{F0: v1, F1: 65535}
                                v48 := GenS1{F0: "\n\t", F1: v23, F2: 1}
                                v3 += v34
                                _ = v47
                                _ = v48
                            }
                            switch "hello" {
                            case "a":
                                v3 += v3
                            case "\n\t":
                                if false {
                                    v49 := GenS4{F0: 0, F1: 107}
                                    _ = v49
                                } else {
                                    fmt.Println(fi25, v3, fi33)
                                    v3 = ((((v34 + v3) + (0.0 * v34)) * v3) / v34)
                                }
                            }
                            switch "" {
                            case "\n\t":
                                v50 := GenS5{F0: 173, F1: false}
                                _ = v50
                            default:
                                v34 = -2.25
                            }
                            _ = fi33
                            _ = v34
                        }
                    case "hello":
                        v3 += (0.1 + 1e10)
                    }
                }
                _ = fi25
            }
        case "foo bar":
            v4 = 47
        }
        _ = v23
    } else {
        for fi51 := 0; fi51 < 2; fi51++ {
            for fi52 := 0; fi52 < 9; fi52++ {
                var v53 int = int(v4)
                _ = fi52
                _ = v53
            }
            switch (fi51 << ((v4) & 63)) {
            case -383:
                v54 := GenS6{F0: (v4 & genFn0(2))}
                _ = v54
            case 32767:
                v55 := GenS2{F0: 42, F1: 150, F2: (false || (-79 > -1))}
                _ = v55
            case 0:
                v3 = -2.25
            default:
                for fi56 := 0; fi56 < 20; fi56++ {
                    v57 := GenS3{F0: (v1 - v1), F1: 32767}
                    if (949 < -9223372036854775808) {
                        if true {
                            var v58 float32 = 0.0
                            _ = v58
                        }
                    } else {
                        v59 := []int{(-(fi51)), fi51, 4294967295, (fi51 % (((fi56 * fi56)) | 1))}
                        var v60 int8 = -69
                        _ = v59
                        _ = v60
                    }
                    _ = fi56
                    _ = v57
                }
            }
            var v61 float32 = 0.1
            _ = fi51
            _ = v61
        }
    }
    if false {
        v62 := GenS7{F0: ("\n\t" + "\n\t"), F1: 255, F2: 32767}
        _ = v62
    } else {
        for fi63 := 0; fi63 < 16; fi63++ {
            if (!(!((747 <= 1) && false))) {
                var v64 int8 = 0
                _ = v64
            } else {
                if false {
                    if (!(-986 <= -627)) {
                        switch ("" + "") {
                        case "":
                            v3 += ((1.5 - float64(-2.25)) + (-2.25 * -2.25))
                        case "foo bar":
                            var v65 int16 = -1
                            _ = v65
                        }
                        switch "" {
                        case "":
                            v66 := GenS8{F0: 0}
                            _ = v66
                        default:
                            v67 := GenS8{F0: 18446744073709551615}
                            _ = v67
                        }
                    }
                    var v68 uint16 = 255
                    _ = v68
                } else {
                    var v69 int32 = -1
                    var v70 int = (fi63 ^ fi63)
                    _ = v69
                    _ = v70
                }
                v71 := GenS2{F0: 127, F1: 42, F2: (-735 == -32768)}
                v4 = genFn0(genFn1(42, ("a" != ("a" + "\n\t"))))
                _ = v71
            }
            delete(v2, "")
            _ = fi63
        }
        for fi72 := 0; fi72 < 1; fi72++ {
            switch (-(fi72)) {
            case -998:
                for fi73 := 0; fi73 < 14; fi73++ {
                    var v74 uint16 = genFn1(255, true)
                    v75 := []int{(fi73 % ((fi72) | 1)), (fi72 + (fi72 << (((v4 | uint(v1))) & 63))), fi73}
                    if ((-14 >= -128) || (((1.5 - -2.25) * 1.5) == -2.25)) {
                        v74 = v74
                    }
                    _ = fi73
                    _ = v74
                    _ = v75
                }
            default:
                v4 = v4
            }
            if ((47 != 2) || (v1 >= (v1 & (v1 * v1)))) {
                v76 := []int{(fi72 ^ (fi72 / ((fi72) | 1))), fi72, fi72, (fi72 & -587)}
                _ = v76
            } else {
                v1 += v1
            }
            _ = fi72
        }
    }
    if false {
        _ = len(v2)
    } else {
        v4 += (v4 + uint(v1))
        fmt.Println(v3, v1, v4)
    }
    v77 := GenS9{F0: 859, F1: (v1 & v1), F2: (^(v4))}
    v1 += 1
    v78 := GenS9{F0: 32767, F1: v1, F2: (v4 * (v4 - 255))}
    var v79 int32 = 42
    for fi80 := 0; fi80 < 15; fi80++ {
        v2["hello"] = "foo bar"
        v81 := GenS10{F0: (v79 >> ((65535) & 63)), F1: 130, F2: 318}
        fmt.Println(fi80)
        _ = fi80
        _ = v81
    }
    v3 += (v3 * v3)
    var v82 uint = (v4 << ((255) & 63))
    var v83 uint = (v82 & 2)
    var v84 int8 = 42
    var v85 string = "hello"
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v79:", v79)
    fmt.Println("v82:", v82)
    fmt.Println("v83:", v83)
    fmt.Println("v84:", v84)
    fmt.Println("v85:", v85)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v77
    _ = v78
    _ = v79
    _ = v82
    _ = v83
    _ = v84
    _ = v85
}
