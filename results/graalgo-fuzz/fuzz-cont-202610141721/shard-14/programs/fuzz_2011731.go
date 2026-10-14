package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64 }
type GenS2 struct { F0 uint32; F1 uint64; F2 string }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int; F1 uint }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 uint32; F1 uint32; F2 uint32 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 uint64; F1 int32; F2 bool }
type GenS9 struct { F0 int8; F1 int32; F2 uint16 }
type GenS10 struct { F0 int8; F1 int16; F2 uint16 }

func main() {
    var v1 uint8 = 2
    var v2 float32 = 0.1
    var v3 bool = ("\n\t" > ("a" + ("a" + (("\n\t" + ("foo bar" + "")) + (("" + "\n\t") + "")))))
    if (1 < -1) {
        v4 := GenS1{F0: 0, F1: 0}
        var v5 int64 = 722
        for fi6 := 0; fi6 < 8; fi6++ {
            var v7 int = fi6
            _ = fi6
            _ = v7
        }
        _ = v4
        _ = v5
    } else {
        v8 := GenS2{F0: 217, F1: 357, F2: ("a" + "\n\t")}
        if (!(!v3)) {
            var v9 float32 = float32(0.1)
            if false {
                if (v3 && (float32(((-2.25 + 3.14159) * -2.25)) == v9)) {
                    v10 := []int{431, 32767, 752}
                    v3 = (((-2.25 - v2) <= v2) || v3)
                    _ = v10
                }
            }
            _ = v9
        } else {
            for fi11 := 0; fi11 < 17; fi11++ {
                var v12 uint = 1
                if (1 != -1) {
                    var v13 string = (("foo bar" + "foo bar") + "foo bar")
                    v1 = 234
                    fmt.Println(fi11, v13)
                    _ = v13
                }
                _ = fi11
                _ = v12
            }
            if false {
                for fi14 := 0; fi14 < 17; fi14++ {
                    v15 := GenS3{F0: 2}
                    _ = fi14
                    _ = v15
                }
                var v16 int = -2147483648
                for fi17 := 0; fi17 < 6; fi17++ {
                    var v18 int16 = 1
                    v19 := []int{32768}
                    v18 += (^(v18))
                    _ = fi17
                    _ = v18
                    _ = v19
                }
                _ = v16
            } else {
                if v3 {
                    var v20 string = ("foo bar" + ("\n\t" + ((("a" + "hello") + ("\n\t" + ("hello" + ""))) + "\n\t")))
                    for fi21 := 0; fi21 < 20; fi21++ {
                        v1 -= (v1 + (v1 % ((91) | 1)))
                        v2 = v2
                        _ = fi21
                    }
                    switch (v20 + "a") {
                    case "\n\t":
                        v22 := GenS4{F0: -1, F1: 255}
                        _ = v22
                    case "hello":
                        fmt.Println(v3)
                    case "a":
                        v1 -= (v1 - 16)
                    default:
                        if ((-9223372036854775808 == -270) || v3) {
                            v23 := GenS3{F0: 2}
                            v1 = v1
                            var v24 string = "\n\t"
                            _ = v23
                            _ = v24
                        } else {
                            v25 := GenS1{F0: 2147483648, F1: 42}
                            var v26 string = v20
                            switch (("a" + ("a" + "hello")) + "foo bar") {
                            case "hello":
                                v2 -= ((v2 - ((v2 / float32((3.14159 / (((3.14159 + 1.5) - 0.1) + 1.5)))) + v2)) + 1.5)
                            default:
                                v27 := []int{0, 1, 32767, 0}
                                _ = v27
                            }
                            _ = v25
                            _ = v26
                        }
                    }
                    _ = v20
                }
                v28 := []int{-937}
                _ = v28
            }
            v2 = (v2 * v2)
        }
        for fi29 := 0; fi29 < 5; fi29++ {
            v1 -= v1
            _ = fi29
        }
        _ = v8
    }
    if v3 {
        v3 = true
        v30 := GenS5{F0: 635}
        _ = v30
    } else {
        v31 := GenS3{F0: 4294967295}
        _ = v31
    }
    for fi32 := 0; fi32 < 3; fi32++ {
        for fi33 := 0; fi33 < 12; fi33++ {
            var v34 int32 = -1
            _ = fi33
            _ = v34
        }
        switch (fi32 & (fi32 * (fi32 * fi32))) {
        case -940:
            v35 := GenS2{F0: 65535, F1: 42, F2: "hello"}
            _ = v35
        case 2147483647:
            v1 = 73
        }
        _ = fi32
    }
    v1 = v1
    v36 := []int{9223372036854775807}
    fmt.Println(v2)
    v37 := GenS1{F0: 0, F1: 271}
    v3 = (2 < 24)
    if (false && (1 != 655)) {
        v1 = (v1 - v1)
        var v38 int64 = 525
        _ = v38
    } else {
        v39 := GenS6{F0: 503, F1: 2, F2: 51}
        v2 = v2
        if true {
            for fi40 := 0; fi40 < 12; fi40++ {
                if len(v36) >= 2 {
                    v36 = v36[1:len(v36)]
                }
                v41 := GenS7{F0: v3}
                _ = fi40
                _ = v41
            }
        }
        _ = v39
    }
    if v3 {
        fmt.Println(v2, v1)
        if (false && (!(!((v3 && (v1 < (v1 ^ v1))) || (-946 != -287))))) {
            fmt.Println(v1, v3, v3)
            v1 = (^(v1))
            v42 := GenS3{F0: 0}
            _ = v42
        }
    }
    if true {
        if v3 {
            if len(v36) >= 2 {
                v36 = v36[1:len(v36)]
            }
        } else {
            for fi43 := 0; fi43 < 16; fi43++ {
                v3 = v3
                _ = fi43
            }
            v2 += (v2 - v2)
            for fi44 := 0; fi44 < 6; fi44++ {
                fmt.Println(v3)
                if true {
                    if (-1 == -863) {
                        if len(v36) >= 2 {
                            v36 = v36[1:len(v36)]
                        }
                    } else {
                        v2 = (v2 + float32((0.0 / 1.0)))
                    }
                    v2 -= v2
                    v2 = ((((1e10 * v2) * 3.14159) * v2) + (v2 + -2.25))
                }
                v3 = (!false)
                _ = fi44
            }
        }
    } else {
        switch ("a" + "") {
        case "hello":
            var v45 int64 = 42
            _ = v45
        case "\n\t":
            switch int(v1) {
            case 944:
                for fi46 := 0; fi46 < 7; fi46++ {
                    for fi47 := 0; fi47 < 1; fi47++ {
                        if (v1 > v1) {
                            if v3 {
                                var v48 float64 = (3.14159 + 3.14159)
                                fmt.Println(v48)
                                _ = v48
                            }
                            switch fi47 {
                            case 320:
                                v1 = v1
                            case -455:
                                fmt.Println(v2, v3, v2)
                            case 32767:
                                v49 := GenS8{F0: 2, F1: -2147483648, F2: (798 == 65535)}
                                _ = v49
                            }
                            var v50 int = (^(fi46))
                            var v51 uint8 = (v1 - v1)
                            _ = v50
                            _ = v51
                        }
                        v3 = (!((v3 && v3) && v3))
                        var v52 int64 = -1
                        _ = fi47
                        _ = v52
                    }
                    v3 = v3
                    if len(v36) >= 2 {
                        v36 = v36[1:len(v36)]
                    }
                    _ = fi46
                }
            case 430:
                if false {
                    switch 127 {
                    case 128:
                        var v53 int8 = -1
                        _ = v53
                    case 9223372036854775807:
                        var v54 uint32 = 1
                        _ = v54
                    default:
                        if len(v36) > 0 {
                            v36[0] = 642
                        }
                    }
                }
            case 808:
                if v3 {
                    switch "a" {
                    case "foo bar":
                        var v55 int16 = -32768
                        _ = v55
                    case "":
                        var v56 int8 = 127
                        _ = v56
                    default:
                        v2 = v2
                    }
                    v57 := GenS9{F0: -28, F1: 847, F2: 32767}
                    fmt.Println(v1, v1, v1)
                    _ = v57
                } else {
                    v58 := GenS3{F0: uint64(v1)}
                    if v3 {
                        var v59 int8 = 127
                        v60 := GenS4{F0: 712, F1: 42}
                        _ = v59
                        _ = v60
                    } else {
                        switch "a" {
                        case "\n\t":
                            v3 = true
                        case "":
                            fmt.Println(v1)
                        case "hello":
                            v61 := GenS10{F0: 42, F1: 32767, F2: 32767}
                            _ = v61
                        }
                    }
                    var v62 uint8 = 63
                    _ = v58
                    _ = v62
                }
            default:
                fmt.Println(v2, v2)
            }
        case "foo bar":
            var v63 uint = 391
            _ = v63
        default:
            var v64 float64 = 0.0
            _ = v64
        }
        v3 = (!(42 != 18446744073709551615))
    }
    v36 = append(v36, 4294967295)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v36:", len(v36), v36)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
}
