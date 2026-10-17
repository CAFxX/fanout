package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint32; F1 string }
type GenS4 struct { F0 bool; F1 int64; F2 uint32 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 uint16; F1 uint32; F2 int }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 string }
type GenS10 struct { F0 int32 }
func genFn0(p0 uint, p1 int8) uint32 {
    var v5 string = "\n\t"
    fmt.Println(p1, v5, p0)
    for fi6 := 0; fi6 < 7; fi6++ {
        for fi7 := 0; fi7 < 11; fi7++ {
            v8 := GenS1{F0: (fi7 + (fi6 % ((-9223372036854775808) | 1))), F1: 812}
            var v9 float32 = (-2.25 * ((0.0 - ((float32((((3.14159 + 0.1) * (-2.25 - (1.5 - 0.0))) * 1.5)) / 3.14159) + 1.5)) - 1e10))
            v10 := GenS1{F0: fi7, F1: 2147483647}
            _ = fi7
            _ = v10
            _ = v8
            _ = v9
        }
        v11 := GenS2{F0: 206}
        _ = fi6
        _ = v11
    }
    if true {
        if ((764 == 559) && false) {
            switch 42 {
            case -9223372036854775808:
                for fi12 := 0; fi12 < 16; fi12++ {
                    if false {
                        v5 = "foo bar"
                        switch (fi12 << (((p0 | 2)) & 63)) {
                        case 550:
                            fmt.Println(fi12, p0)
                        case 32767:
                            fmt.Println(v5, p0, fi12)
                        default:
                            v13 := GenS1{F0: (fi12 ^ fi12), F1: 42}
                            _ = v13
                        }
                        for fi14 := 0; fi14 < 7; fi14++ {
                            if true {
                                p0 = (p0 / (((p0 ^ p0)) | 1))
                                var v15 uint8 = 207
                                _ = v15
                            } else {
                                switch fi12 {
                                case -2147483648:
                                    v5 = "foo bar"
                                case 0:
                                    v16 := GenS1{F0: fi14, F1: 42}
                                    _ = v16
                                }
                                var v17 uint32 = 4294967295
                                var v18 uint32 = v17
                                _ = v17
                                _ = v18
                            }
                            _ = fi14
                        }
                    } else {
                        for fi19 := 0; fi19 < 7; fi19++ {
                            p1 = p1
                            v20 := GenS3{F0: 1, F1: v5}
                            _ = fi19
                            _ = v20
                        }
                        if ((3.14159 + 0.0) != 1e10) {
                            for fi21 := 0; fi21 < 5; fi21++ {
                                var v22 int = 9223372036854775807
                                v23 := GenS4{F0: false, F1: 42, F2: 1}
                                _ = fi21
                                _ = v22
                                _ = v23
                            }
                        } else {
                            var v24 int64 = 0
                            _ = v24
                        }
                        for fi25 := 0; fi25 < 5; fi25++ {
                            v26 := []int{-9223372036854775808, -337}
                            _ = fi25
                            _ = v26
                        }
                    }
                    p1 = (p1 ^ -53)
                    _ = fi12
                }
            case 128:
                v27 := GenS1{F0: 270, F1: 42}
                _ = v27
            default:
                v28 := []int{0}
                _ = v28
            }
            switch 1 {
            case 32767:
                v29 := GenS5{F0: 42}
                _ = v29
            case 9223372036854775807:
                if false {
                    v30 := []int{128, 486, 4294967295, 1}
                    _ = v30
                } else {
                    if ((-364 <= 0) || (547 < 18446744073709551615)) {
                        v31 := GenS4{F0: true, F1: 0, F2: 2147483648}
                        fmt.Println(v5)
                        var v32 uint = (p0 + p0)
                        _ = v31
                        _ = v32
                    } else {
                        if (!(!true)) {
                            var v33 uint64 = 434
                            _ = v33
                        }
                        var v34 bool = false
                        _ = v34
                    }
                }
            case -169:
                v35 := []int{274, -291, 9223372036854775807}
                _ = v35
            default:
                v36 := GenS5{F0: 42}
                _ = v36
            }
            for fi37 := 0; fi37 < 13; fi37++ {
                var v38 string = v5
                var v39 int32 = 2147483647
                _ = fi37
                _ = v38
                _ = v39
            }
        }
        for fi40 := 0; fi40 < 6; fi40++ {
            v5 = v5
            _ = fi40
        }
    }
    if false {
        if (4294967295 >= 1) {
            var v41 int = 0
            var v42 int8 = (^(p1))
            _ = v41
            _ = v42
        } else {
            fmt.Println(p1, v5)
        }
        var v43 bool = false
        _ = v43
    } else {
        fmt.Println(v5)
        p1 -= -1
    }
    _ = p0
    _ = p1
    _ = v5
    return 65535
}
func genFn1() string {
    var v44 bool = true
    v44 = v44
    _ = v44
    return (((("foo bar" + "") + "foo bar") + "") + ("" + (("foo bar" + "hello") + (("foo bar" + "a") + "\n\t"))))
}

func main() {
    var v1 int16 = 0
    v2 := make(map[int]int)
    var v3 float32 = 1.5
    v4 := make(map[int]uint32)
    switch -1 {
    case 113:
        var v45 uint = 112
        _ = v45
    case 9223372036854775807:
        v46 := GenS2{F0: 476}
        _ = v46
    case 2147483647:
        v47 := GenS6{F0: 65535}
        _ = v47
    default:
        v48 := GenS3{F0: 887, F1: "\n\t"}
        _ = v48
    }
    var v49 uint32 = genFn0(941, -1)
    for fi50 := 0; fi50 < 1; fi50++ {
        fmt.Println(fi50, v3)
        for fi51 := 0; fi51 < 5; fi51++ {
            if ("foo bar" < (("foo bar" + "hello") + "a")) {
                v52 := []int{(fi50 * fi50), (fi51 * -2147483648), fi51, 9223372036854775807}
                v3 += ((3.14159 + (float32((float64((v3 * v3)) * (1.5 / (0.1 - (1e10 + float64(v3)))))) - float32(float64(1e10)))) * ((1.5 / v3) * 1e10))
                switch "hello" {
                case "foo bar":
                    v1 -= v1
                case "\n\t":
                    if (!(0.1 != -2.25)) {
                        var v53 int32 = -1
                        var v54 int = -927
                        var v55 uint32 = 2
                        _ = v53
                        _ = v54
                        _ = v55
                    } else {
                        var v56 int = (fi51 ^ fi51)
                        v57 := GenS7{F0: 487, F1: 2, F2: -863}
                        switch (fi51 / ((-2147483648) | 1)) {
                        case 127:
                            v58 := GenS8{F0: -692}
                            _ = v58
                        case 4294967295:
                            if len(v52) >= 2 {
                                v52 = v52[1:len(v52)]
                            }
                        case 193:
                            var v59 bool = false
                            _ = v59
                        }
                        _ = v56
                        _ = v57
                    }
                default:
                    switch fi50 {
                    case -2147483648:
                        _ = len(v52)
                    }
                }
                _ = v52
            }
            if false {
                switch 1 {
                case 1:
                    v3 -= 0.0
                case 42:
                    v1 = (^(v1))
                case 2147483647:
                    var v60 string = ((("a" + "foo bar") + ("hello" + ((((((("foo bar" + "a") + ("hello" + ("a" + ("" + "")))) + "hello") + "a") + (("" + "") + "hello")) + (("foo bar" + "\n\t") + "hello")) + ("foo bar" + ((("hello" + "\n\t") + "\n\t") + "\n\t"))))) + "a")
                    _ = v60
                }
                if false {
                    if true {
                        v49 -= (v49 * v49)
                        v61 := GenS9{F0: (("foo bar" + "foo bar") + "foo bar")}
                        _ = v61
                    }
                    v62 := v2[(fi51 + fi50)]
                    _ = v62
                }
            }
            var v63 string = "foo bar"
            _ = fi51
            _ = v63
        }
        switch (fi50 ^ (fi50 + (fi50 * fi50))) {
        case -978:
            v3 += v3
        case 127:
            v64 := GenS6{F0: 42}
            _ = v64
        }
        _ = fi50
    }
    var v65 bool = false
    fmt.Println(v1, v65)
    if v65 {
        if ((-2.25 - v3) > float32((((float64(v3) - (1e10 * -2.25)) * (1.5 * (0.0 + ((3.14159 * float64(1e10)) - ((3.14159 / 1.0) - -2.25))))) / 1.0))) {
            switch -696 {
            case 9223372036854775807:
                v49 = v49
            case 32768:
                v66 := GenS7{F0: 184, F1: (^(v49)), F2: -2147483648}
                _ = v66
            }
        }
        var v67 int32 = 596
        fmt.Println(v67)
        _ = v67
    }
    fmt.Println(v49, v1, v3)
    var v68 uint = 2
    var v69 int16 = (v1 >> ((v68) & 63))
    var v70 string = "foo bar"
    var v71 uint16 = 1
    for fi72 := 0; fi72 < 11; fi72++ {
        v73 := []int{fi72, (fi72 + (fi72 % ((fi72) | 1))), fi72, -585}
        var v74 float32 = 1.5
        for fi75 := 0; fi75 < 19; fi75++ {
            v76 := GenS10{F0: -22}
            switch v70 {
            case "foo bar":
                var v77 int16 = (v69 + v69)
                _ = v77
            default:
                v2[fi72] = fi72
            }
            if len(v73) > 0 {
                v73[0] = (fi75 + int(v69))
            }
            switch v70 {
            case "hello":
                if true {
                    v68 = v68
                } else {
                    var v78 int8 = 127
                    _ = v78
                }
            default:
                var v79 string = v70
                _ = v79
            }
            _ = fi75
            _ = v76
        }
        _ = fi72
        _ = v73
        _ = v74
    }
    v3 += (0.1 * -2.25)
    var v80 int64 = 481
    for fi81 := 0; fi81 < 7; fi81++ {
        for fi82 := 0; fi82 < 8; fi82++ {
            v83 := GenS6{F0: 862}
            fmt.Println(v80, fi81)
            var v84 uint = v68
            _ = fi82
            _ = v83
            _ = v84
        }
        _ = fi81
    }
    if ((4294967295 > -2147483648) && (((^(v68)) == v68) || (!(v65 || (1 == 1))))) {
        for fi85 := 0; fi85 < 3; fi85++ {
            v2[32767] = fi85
            var v86 int64 = (-(v80))
            v86 -= 42
            _ = fi85
            _ = v86
        }
    }
    switch 9223372036854775807 {
    case 2147483647:
        var v87 uint64 = 1
        _ = v87
    default:
        if false {
            switch 0 {
            case 32768:
                var v88 uint32 = 2
                _ = v88
            case 1:
                var v89 int8 = 57
                _ = v89
            case 128:
                v90 := []int{}
                _ = v90
            default:
                if v65 {
                    var v91 float32 = (-2.25 / v3)
                    _ = v91
                }
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v49:", v49)
    fmt.Println("v65:", v65)
    fmt.Println("v68:", v68)
    fmt.Println("v69:", v69)
    fmt.Println("v70:", v70)
    fmt.Println("v71:", v71)
    fmt.Println("v80:", v80)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v65
    _ = v68
    _ = v69
    _ = v70
    _ = v71
    _ = v80
}
