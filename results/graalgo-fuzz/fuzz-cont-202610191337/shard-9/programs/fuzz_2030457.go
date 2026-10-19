package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int32; F1 int32; F2 int32 }
type GenS3 struct { F0 uint16; F1 bool }
type GenS4 struct { F0 uint; F1 int16; F2 int8 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 string; F1 string }
type GenS7 struct { F0 bool; F1 uint16; F2 int8 }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 int64; F1 int16; F2 int16 }
func genFn0(p0 int16, p1 int16) int {
    for fi5 := 0; fi5 < 10; fi5++ {
        if true {
            v6 := GenS1{F0: -63}
            _ = v6
        }
        p0 = (p0 + p0)
        _ = fi5
    }
    for fi7 := 0; fi7 < 7; fi7++ {
        v8 := GenS1{F0: -1}
        fmt.Println(fi7, p0)
        _ = fi7
        _ = v8
    }
    v9 := []int{804, 128, -9223372036854775808, 127}
    for fi10 := 0; fi10 < 18; fi10++ {
        var v11 uint16 = 189
        v12 := GenS1{F0: 1}
        _ = fi10
        _ = v11
        _ = v12
    }
    p1 = (-(p1))
    _ = p0
    _ = p1
    _ = v9
    return -9223372036854775808
}
func genFn1(p0 string, p1 bool, p2 bool) uint8 {
    p0 = ("\n\t" + "\n\t")
    var v13 int64 = 42
    fmt.Println(p2, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return 1
}

func main() {
    v1 := []int{}
    var v2 uint16 = 42
    var v3 uint8 = 85
    var v4 int = 1
    fmt.Println(v4, v2)
    v3 += 127
    var v14 float64 = 3.14159
    if ("foo bar" < "foo bar") {
        switch "a" {
        case "a":
            if (-128 <= -60) {
                switch 913 {
                case 0:
                    switch int(v3) {
                    case 859:
                        v2 = v2
                    case 32767:
                        if false {
                            var v15 uint8 = (v3 + v3)
                            _ = v15
                        } else {
                            var v16 uint16 = (^(v2))
                            for fi17 := 0; fi17 < 9; fi17++ {
                                var v18 uint8 = (v3 >> ((382) & 63))
                                v16 -= v16
                                if true {
                                    var v19 uint64 = 0
                                    if true {
                                        for fi20 := 0; fi20 < 17; fi20++ {
                                            if (3.14159 < 1e10) {
                                                v1 = append(v1, (fi20 / ((fi20) | 1)))
                                                if len(v1) > 0 {
                                                    v1[0] = 4294967295
                                                }
                                                if (!(true || (!("foo bar" <= "hello")))) {
                                                    v21 := GenS1{F0: 127}
                                                    _ = v21
                                                }
                                            }
                                            if ((fi20 ^ (fi20 / ((42) | 1))) < genFn0(-1, 520)) {
                                                var v22 uint16 = v16
                                                var v23 int16 = -836
                                                var v24 float32 = float32(((v14 * (v14 / ((((((float64(0.1) - v14) / v14) + v14) * float64(0.0)) + v14) * float64((0.1 / float32(1.5)))))) / 3.14159))
                                                _ = v22
                                                _ = v23
                                                _ = v24
                                            }
                                            v25 := GenS2{F0: 1, F1: 712, F2: 1}
                                            v16 -= 65535
                                            _ = fi20
                                            _ = v25
                                        }
                                        var v26 int = -886
                                        _ = v26
                                    } else {
                                        if len(v1) > 0 {
                                            v1[0] = v4
                                        }
                                    }
                                    v16 += uint16(v19)
                                    _ = v19
                                }
                                _ = fi17
                                _ = v18
                            }
                            switch (v4 & (v4 + int(v3))) {
                            case 9223372036854775807:
                                for fi27 := 0; fi27 < 4; fi27++ {
                                    var v28 uint8 = (v3 - (v3 ^ v3))
                                    _ = fi27
                                    _ = v28
                                }
                            }
                            _ = v16
                        }
                    }
                }
                for fi29 := 0; fi29 < 11; fi29++ {
                    switch "hello" {
                    case "\n\t":
                        v4 += (v4 * 42)
                    case "a":
                        v4 = fi29
                    case "":
                        var v30 int64 = 9223372036854775807
                        _ = v30
                    default:
                        if false {
                            if (v14 == float64((((0.1 / (3.14159 + 3.14159)) + (3.14159 - (3.14159 * (3.14159 + 0.0)))) * (((3.14159 + 3.14159) - 1e10) / 0.1)))) {
                                v31 := GenS3{F0: v2, F1: false}
                                for fi32 := 0; fi32 < 13; fi32++ {
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    _ = len(v1)
                                    v3 = 191
                                    _ = fi32
                                }
                                v33 := GenS4{F0: 0, F1: -345, F2: 117}
                                _ = v31
                                _ = v33
                            }
                            switch fi29 {
                            case 32768:
                                v34 := GenS5{F0: 0}
                                _ = v34
                            case -1:
                                var v35 int64 = -473
                                _ = v35
                            case 127:
                                v36 := GenS6{F0: ("\n\t" + ("foo bar" + "foo bar")), F1: "hello"}
                                _ = v36
                            }
                        } else {
                            v14 -= ((float64(-2.25) - float64(3.14159)) * float64(0.0))
                        }
                    }
                    v37 := GenS7{F0: true, F1: 65535, F2: 49}
                    _ = fi29
                    _ = v37
                }
                switch v4 {
                case 367:
                    for fi38 := 0; fi38 < 19; fi38++ {
                        var v39 float64 = v14
                        _ = fi38
                        _ = v39
                    }
                case 42:
                    if (!(3.14159 == ((1.5 + (float64((-2.25 - 1e10)) + v14)) * v14))) {
                        v2 -= (v2 % (((v2 ^ (^(v2)))) | 1))
                        var v40 uint16 = (^(v2))
                        _ = v40
                    }
                }
            } else {
                for fi41 := 0; fi41 < 12; fi41++ {
                    if true {
                        v1 = append(v1, fi41)
                        v42 := GenS1{F0: 42}
                        if (((float64(0.1) * (v14 / (((v14 / float64(1.5)) * v14) * (v14 * float64((float32((-2.25 / 0.1)) + -2.25)))))) + v14) <= (3.14159 * (((((v14 + 0.1) + (v14 + v14)) + v14) * (((v14 - (v14 - (1.5 * (v14 * v14)))) - float64((3.14159 - (float32(3.14159) * 0.1)))) / v14)) - float64((1.5 / 1.5))))) {
                            var v43 float64 = (v14 - v14)
                            v1 = append(v1, fi41)
                            var v44 float64 = v14
                            _ = v43
                            _ = v44
                        } else {
                            if false {
                                v45 := GenS8{F0: 4294967295}
                                v46 := GenS6{F0: (("hello" + ("foo bar" + "hello")) + "foo bar"), F1: ("hello" + "hello")}
                                fmt.Println(v2, v14, fi41)
                                _ = v45
                                _ = v46
                            } else {
                                var v47 int64 = 395
                                v14 = (float64(-2.25) * 0.1)
                                if ("" != "\n\t") {
                                    v48 := GenS8{F0: 42}
                                    _ = v48
                                }
                                _ = v47
                            }
                        }
                        _ = v42
                    } else {
                        for fi49 := 0; fi49 < 14; fi49++ {
                            fmt.Println(v3, v14)
                            if false {
                                if (255 > 18446744073709551615) {
                                    var v50 uint = 65535
                                    _ = v50
                                } else {
                                    var v51 string = "\n\t"
                                    var v52 int16 = 32767
                                    fmt.Println(v3)
                                    var v53 uint = 541
                                    v2 -= v2
                                    _ = v51
                                    _ = v52
                                    _ = v53
                                }
                                var v54 int16 = -32768
                                if true {
                                    _ = len(v1)
                                } else {
                                    v2 += (^(v2))
                                    if (4294967295 != 1) {
                                        if ((!(1.5 == 0.1)) && (true || ((uint32(v3) >= 42) || false))) {
                                            v55 := GenS7{F0: true, F1: v2, F2: 1}
                                            var v56 uint = 255
                                            _ = v55
                                            _ = v56
                                        }
                                    } else {
                                        if (true || true) {
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                            var v57 float64 = (v14 + (v14 + v14))
                                            _ = v57
                                        } else {
                                            v58 := GenS2{F0: 2147483647, F1: -606, F2: -1}
                                            v14 += v14
                                            _ = v58
                                        }
                                    }
                                    var v59 bool = (v2 >= (v2 << ((0) & 63)))
                                    var v60 uint = 2
                                    var v61 float32 = (0.1 + -2.25)
                                    _ = v59
                                    _ = v60
                                    _ = v61
                                }
                                _ = v54
                            } else {
                                v2 = v2
                            }
                            v62 := GenS3{F0: 1, F1: (1 < 2147483647)}
                            _ = fi49
                            _ = v62
                        }
                    }
                    _ = fi41
                }
                v1 = append(v1, (v4 + v4))
            }
        }
        var v63 int16 = 798
        var v64 int8 = -1
        _ = v63
        _ = v64
    } else {
        v65 := GenS9{F0: 1, F1: -941, F2: -223}
        _ = v65
    }
    var v66 int16 = 771
    var v67 uint = 473
    v68 := GenS8{F0: (v67 / (((v67 * (^(v67)))) | 1))}
    v4 -= v4
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v66:", v66)
    fmt.Println("v67:", v67)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v66
    _ = v67
    _ = v68
}
