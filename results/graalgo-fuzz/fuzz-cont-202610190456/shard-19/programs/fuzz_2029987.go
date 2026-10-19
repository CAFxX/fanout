package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16 }
type GenS2 struct { F0 string; F1 uint }
type GenS3 struct { F0 uint8; F1 int8; F2 uint }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int16; F1 bool; F2 uint8 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 uint16; F1 uint32 }
func genFn0() bool {
    var v6 int16 = -1
    switch "" {
    case "a":
        v6 += (v6 >> ((645) & 63))
    case "":
        v6 -= v6
    default:
        for fi7 := 0; fi7 < 11; fi7++ {
            v8 := GenS1{F0: 2147483647, F1: 101}
            var v9 int = 32768
            _ = fi7
            _ = v8
            _ = v9
        }
    }
    switch "" {
    case "\n\t":
        switch -158 {
        case 128:
            v6 = (^(v6))
        default:
            var v10 int16 = v6
            _ = v10
        }
    case "hello":
        switch "foo bar" {
        case "foo bar":
            switch "a" {
            case "hello":
                if (694 <= 18446744073709551615) {
                    v11 := GenS2{F0: ("a" + ""), F1: 18446744073709551615}
                    _ = v11
                } else {
                    var v12 float32 = 1.5
                    _ = v12
                }
            case "\n\t":
                if (42 != 634) {
                    switch ("hello" + "foo bar") {
                    case "":
                        var v13 int64 = 42
                        _ = v13
                    case "foo bar":
                        var v14 int16 = (^(v6))
                        _ = v14
                    case "hello":
                        if (1 == 598) {
                            if true {
                                var v15 float32 = (0.0 * 1e10)
                                _ = v15
                            }
                        }
                    default:
                        if (241 != 42) {
                            if (42 != 18446744073709551615) {
                                var v16 int = -97
                                var v17 uint32 = 65535
                                _ = v16
                                _ = v17
                            } else {
                                v18 := []int{32767, 662}
                                _ = v18
                            }
                            v19 := []int{42, -1, 32767, 128}
                            if len(v19) >= 2 {
                                v19 = v19[1:len(v19)]
                            }
                            _ = v19
                        } else {
                            for fi20 := 0; fi20 < 16; fi20++ {
                                v21 := []int{fi20}
                                var v22 uint8 = 42
                                _ = fi20
                                _ = v21
                                _ = v22
                            }
                        }
                    }
                }
            }
        }
    }
    for fi23 := 0; fi23 < 10; fi23++ {
        switch (("" + ("" + "foo bar")) + "") {
        case "":
            var v24 string = "hello"
            _ = v24
        case "hello":
            fmt.Println(v6)
        case "a":
            v6 = (-(v6))
        }
        _ = fi23
    }
    for fi25 := 0; fi25 < 3; fi25++ {
        if (42 > 42) {
            if false {
                switch fi25 {
                case 570:
                    var v26 uint32 = 2147483648
                    _ = v26
                case -1:
                    fmt.Println(v6)
                default:
                    var v27 int64 = 0
                    _ = v27
                }
                v6 = (^(v6))
                switch (fi25 << ((2) & 63)) {
                case -2147483648:
                    for fi28 := 0; fi28 < 7; fi28++ {
                        if (!(((1e10 / 3.14159) >= 0.0) && (813 >= 928))) {
                            v29 := []int{(fi25 % ((fi28) | 1)), (fi28 >> ((249) & 63)), 2147483647, fi25}
                            _ = v29
                        }
                        v6 += (v6 << ((257) & 63))
                        var v30 float64 = -2.25
                        _ = fi28
                        _ = v30
                    }
                case -1:
                    var v31 float64 = -2.25
                    _ = v31
                default:
                    if (255 == 863) {
                        if (!false) {
                            v6 -= 32767
                        }
                    } else {
                        v32 := []int{fi25, -1, -9223372036854775808, (^(fi25))}
                        _ = len(v32)
                        _ = v32
                    }
                }
            } else {
                v33 := GenS3{F0: 173, F1: -39, F2: 961}
                if true {
                    var v34 uint32 = 2147483647
                    var v35 uint8 = 42
                    v6 -= int16(v35)
                    _ = v34
                    _ = v35
                } else {
                    v36 := []int{fi25, fi25, fi25}
                    _ = v36
                }
                _ = v33
            }
        } else {
            v6 -= v6
        }
        _ = fi25
    }
    _ = v6
    return (false && (255 < 65535))
}
func genFn1() int {
    switch ("foo bar" + (("foo bar" + "\n\t") + ((((((("\n\t" + ("\n\t" + (("foo bar" + "") + (("a" + "\n\t") + (((("" + ("hello" + "foo bar")) + "a") + ("\n\t" + "foo bar")) + ""))))) + "foo bar") + "foo bar") + "") + "") + "hello") + "a"))) {
    case "\n\t":
        var v37 float64 = 0.0
        _ = v37
    case "hello":
        v38 := []int{-170, -1}
        _ = v38
    default:
        v39 := GenS4{F0: 439}
        _ = v39
    }
    var v40 uint64 = 18446744073709551615
    _ = v40
    return -388
}

func main() {
    v1 := []int{736}
    var v2 int = 0
    var v3 uint16 = 32767
    v4 := make(map[int]int)
    v5 := []uint32{}
    v2 += v2
    var v41 int8 = -1
    var v42 uint16 = 2
    var v43 int32 = 42
    if len(v1) > 0 {
        v1[0] = (v2 + v2)
    }
    delete(v4, genFn1())
    for fi44 := 0; fi44 < 2; fi44++ {
        if true {
            switch v2 {
            case 2147483647:
                var v45 string = "a"
                _ = v45
            case 0:
                for fi46 := 0; fi46 < 17; fi46++ {
                    var v47 uint = 42
                    _ = fi46
                    _ = v47
                }
            case 1:
                for fi48 := 0; fi48 < 14; fi48++ {
                    v41 = 20
                    delete(v4, v2)
                    var v49 uint8 = 127
                    _ = fi48
                    _ = v49
                }
            default:
                var v50 int32 = -338
                _ = v50
            }
            v51 := GenS4{F0: genFn1()}
            var v52 int8 = v41
            _ = v51
            _ = v52
        }
        var v53 string = "foo bar"
        v41 = v41
        _ = fi44
        _ = v53
    }
    _ = len(v4)
    v54 := GenS1{F0: 1, F1: 0}
    var v55 int16 = -32768
    if ((929 < 2) || false) {
        if true {
            v41 += v41
            var v56 uint64 = 0
            switch ("hello" + ("\n\t" + "foo bar")) {
            case "a":
                if (-2.25 == ((((1.5 / (3.14159 + 3.14159)) - (((0.0 / ((0.1 / 1.0) / 0.1)) / 0.1) * 0.0)) / 1.5) / 1e10)) {
                    _ = len(v4)
                }
            default:
                v55 += v55
            }
            _ = v56
        }
        for fi57 := 0; fi57 < 10; fi57++ {
            for fi58 := 0; fi58 < 16; fi58++ {
                v41 += -1
                var v59 uint32 = 65535
                var v60 uint = 0
                _ = fi58
                _ = v59
                _ = v60
            }
            _ = len(v1)
            _ = fi57
        }
        fmt.Println(v55, v43, v2)
    }
    for fi61 := 0; fi61 < 12; fi61++ {
        var v62 uint = 18446744073709551615
        v2 += int(v43)
        _ = fi61
        _ = v62
    }
    if (!((3.14159 >= 1e10) && false)) {
        for fi63 := 0; fi63 < 8; fi63++ {
            var v64 float32 = float32(0.0)
            v64 -= 1.5
            if false {
                if true {
                    for fi65 := 0; fi65 < 5; fi65++ {
                        switch ("\n\t" + "a") {
                        case "":
                            v2 = v2
                        case "a":
                            v64 = (0.0 - ((v64 - (-2.25 - v64)) + v64))
                        case "\n\t":
                            v66 := GenS5{F0: (v55 >> ((0) & 63)), F1: (1 == 4294967295), F2: 1}
                            _ = v66
                        default:
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                        }
                        _ = fi65
                    }
                    for fi67 := 0; fi67 < 3; fi67++ {
                        v2 += (-(fi63))
                        v68 := GenS2{F0: ("\n\t" + ("hello" + "")), F1: 65535}
                        var v69 int16 = v55
                        _ = fi67
                        _ = v68
                        _ = v69
                    }
                    v70 := GenS4{F0: int(v41)}
                    _ = v70
                }
                if len(v1) > 0 {
                    v1[0] = (v2 + (v2 ^ (v2 << ((420) & 63))))
                }
                v3 -= (v3 + (v42 << ((4294967295) & 63)))
            } else {
                for fi71 := 0; fi71 < 3; fi71++ {
                    v4[4294967295] = fi63
                    if false {
                        v55 -= (^(v55))
                        var v72 uint16 = v42
                        fmt.Println(v43, v3, v72)
                        _ = v72
                    }
                    _ = fi71
                }
            }
            _ = fi63
            _ = v64
        }
        var v73 float64 = (float64(float32(1.5)) / 1.0)
        _ = v73
    } else {
        v74 := GenS6{F0: 4294967295}
        v75 := GenS7{F0: (v42 % (((v3 | v3)) | 1)), F1: 0}
        var v76 bool = true
        _ = v74
        _ = v75
        _ = v76
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v55:", v55)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v5
    _ = v54
    _ = v55
}
