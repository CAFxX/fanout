package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 bool; F1 uint16 }
type GenS3 struct { F0 int32; F1 int32; F2 int }
type GenS4 struct { F0 bool; F1 uint64 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 int; F1 uint32; F2 bool }
type GenS7 struct { F0 int16; F1 uint8 }
type GenS8 struct { F0 uint64; F1 int8 }
type GenS9 struct { F0 int8 }
type GenS10 struct { F0 uint8 }
type GenS11 struct { F0 int16 }
type GenS12 struct { F0 uint }
func genFn0(p0 uint) uint8 {
    v5 := GenS1{F0: "a"}
    v6 := GenS1{F0: "a"}
    _ = p0
    _ = v5
    _ = v6
    return 1
}

func main() {
    v1 := []int{9223372036854775807, 32768}
    var v2 bool = ((("foo bar" + (("" + "hello") + "hello")) + "a") <= "hello")
    var v3 string = "foo bar"
    var v4 uint8 = 0
    if false {
        var v7 uint = 255
        var v8 float32 = 0.1
        if true {
            var v9 int8 = 100
            _ = v9
        }
        _ = v7
        _ = v8
    }
    fmt.Println(v2)
    for fi10 := 0; fi10 < 1; fi10++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = len(v1)
        for fi11 := 0; fi11 < 12; fi11++ {
            v2 = (!v2)
            var v12 bool = (1 <= -2147483648)
            v13 := GenS1{F0: (("a" + v3) + (v3 + "a"))}
            _ = fi11
            _ = v12
            _ = v13
        }
        _ = fi10
    }
    v14 := GenS1{F0: (v3 + "")}
    var v15 int = 128
    v16 := GenS1{F0: v3}
    for fi17 := 0; fi17 < 20; fi17++ {
        if v2 {
            var v18 uint = 4294967295
            v4 += v4
            if v2 {
                var v19 uint = (v18 & 18446744073709551615)
                _ = v19
            } else {
                v2 = v2
                switch v3 {
                case "foo bar":
                    if (((-54 <= -847) || v2) && false) {
                        v15 = 127
                        switch 266 {
                        case -1:
                            v20 := GenS1{F0: v3}
                            _ = v20
                        case -225:
                            if (!v2) {
                                v2 = v2
                                v3 = (v3 + ("foo bar" + v3))
                            } else {
                                v4 = (^(v4))
                                var v21 int32 = -2147483648
                                v22 := GenS2{F0: (-1 > 542), F1: 32767}
                                _ = v21
                                _ = v22
                            }
                        case 0:
                            if len(v1) > 0 {
                                v1[0] = 4294967295
                            }
                        default:
                            v23 := GenS2{F0: (192 < 42), F1: 65535}
                            _ = v23
                        }
                        v24 := GenS3{F0: -1, F1: 42, F2: fi17}
                        _ = v24
                    }
                default:
                    if (!v2) {
                        if len(v1) > 0 {
                            v1[0] = 9223372036854775807
                        }
                    }
                }
            }
            _ = v18
        } else {
            if (((((1 == 26) && v2) && false) || true) && true) {
                var v25 uint = 1
                v26 := GenS4{F0: v2, F1: 2}
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = v25
                _ = v26
            } else {
                v2 = v2
                if false {
                    var v27 string = v3
                    _ = v27
                } else {
                    var v28 int8 = -128
                    _ = v28
                }
                for fi29 := 0; fi29 < 15; fi29++ {
                    switch "" {
                    case "\n\t":
                        v30 := GenS4{F0: (!(!(v2 || (-1 >= 1)))), F1: 0}
                        _ = v30
                    case "":
                        if (0 <= 60) {
                            v2 = false
                            v31 := GenS1{F0: v3}
                            v4 += (v4 * v4)
                            v3 = "\n\t"
                            _ = v31
                        }
                    default:
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        v15 -= v15
                    }
                    v32 := GenS1{F0: "foo bar"}
                    fmt.Println(v3)
                    _ = fi29
                    _ = v32
                }
            }
            v2 = ((!(0 <= 127)) && v2)
            v15 -= v15
        }
        var v33 int64 = -57
        _ = fi17
        _ = v33
    }
    v4 = 0
    for fi34 := 0; fi34 < 1; fi34++ {
        v3 = v3
        v2 = false
        switch v3 {
        case "\n\t":
            v15 += 42
        default:
            if ((v4 <= 1) || (!v2)) {
                var v35 int32 = 127
                fmt.Println(fi34, v3, v3)
                var v36 float64 = 1.5
                _ = v35
                _ = v36
            } else {
                v37 := GenS5{F0: 2}
                fmt.Println(fi34, v4, v3)
                fmt.Println(fi34)
                _ = v37
            }
        }
        _ = fi34
    }
    v2 = (true && true)
    if false {
        var v38 bool = (1 == 42)
        switch v3 {
        case "foo bar":
            var v39 int64 = -9223372036854775808
            _ = v39
        case "\n\t":
            if (!true) {
                if false {
                    switch (((v3 + (v3 + v3)) + "hello") + "") {
                    case "":
                        for fi40 := 0; fi40 < 3; fi40++ {
                            if (18446744073709551615 > 4294967295) {
                                var v41 uint64 = 4294967295
                                _ = v41
                            } else {
                                if (!((true && (42 <= 100)) || (779 < 18446744073709551615))) {
                                    switch fi40 {
                                    case -9223372036854775808:
                                        v3 = "\n\t"
                                    case 127:
                                        v42 := GenS3{F0: 0, F1: -2147483648, F2: v15}
                                        _ = v42
                                    case -1:
                                        v43 := GenS6{F0: 32767, F1: 0, F2: false}
                                        _ = v43
                                    default:
                                        if false {
                                            v1 = append(v1, fi40)
                                        } else {
                                            switch v3 {
                                            case "":
                                                switch v3 {
                                                case "":
                                                    var v44 int32 = 42
                                                    _ = v44
                                                case "\n\t":
                                                    v3 = "hello"
                                                }
                                            }
                                            v45 := GenS7{F0: 1, F1: 128}
                                            fmt.Println(v15, fi40, v15)
                                            _ = v45
                                        }
                                    }
                                    var v46 float64 = float64(float32(((float64(((0.0 - (float32((float64(-2.25) - 3.14159)) - (0.0 + 1e10))) * 1e10)) - 0.1) * 0.0)))
                                    if v38 {
                                        if v38 {
                                            if ((v4 * v4) <= (v4 & (^(v4)))) {
                                                fmt.Println(v2, fi40)
                                                if ((3.14159 - 1e10) < 3.14159) {
                                                    fmt.Println(v4, v2, v15)
                                                    var v47 uint8 = (^(v4))
                                                    v15 = (fi40 % (((v15 ^ (v15 << ((1) & 63)))) | 1))
                                                    _ = v47
                                                }
                                                v48 := GenS8{F0: 0, F1: -1}
                                                _ = v48
                                            } else {
                                                v15 -= fi40
                                            }
                                            if v2 {
                                                switch v15 {
                                                case 1:
                                                    var v49 float32 = -2.25
                                                    _ = v49
                                                case 128:
                                                    var v50 float64 = v46
                                                    _ = v50
                                                case 575:
                                                    _ = len(v1)
                                                default:
                                                    v4 = 138
                                                }
                                            } else {
                                                v1 = append(v1, 32767)
                                            }
                                            var v51 uint = 589
                                            _ = v51
                                        } else {
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                            v52 := GenS5{F0: 42}
                                            for fi53 := 0; fi53 < 4; fi53++ {
                                                if len(v1) > 0 {
                                                    v1[0] = (^(fi53))
                                                }
                                                if (9223372036854775807 != 9223372036854775807) {
                                                    v4 = 42
                                                    switch ((((v3 + (v3 + v3)) + "") + v3) + (("hello" + "foo bar") + "foo bar")) {
                                                    case "foo bar":
                                                        var v54 uint16 = 82
                                                        _ = v54
                                                    default:
                                                        if ((fi53 - v15) > (fi40 >> ((18446744073709551615) & 63))) {
                                                            switch fi40 {
                                                            case 0:
                                                                v2 = (2 == 42)
                                                                v38 = v38
                                                                if len(v1) >= 2 {
                                                                    v1 = v1[1:len(v1)]
                                                                }
                                                            }
                                                            switch v3 {
                                                            case "hello":
                                                                if len(v1) > 0 {
                                                                    v1[0] = (fi40 + (^(fi40)))
                                                                }
                                                                fmt.Println(v15, v46, fi40)
                                                                if len(v1) > 0 {
                                                                    v1[0] = fi53
                                                                }
                                                            case "\n\t":
                                                                switch int(v4) {
                                                                case -9223372036854775808:
                                                                    var v55 uint = 68
                                                                    _ = v55
                                                                case -212:
                                                                    v56 := GenS9{F0: 0}
                                                                    _ = v56
                                                                case 42:
                                                                    if v38 {
                                                                        v46 += 1.5
                                                                        v57 := GenS6{F0: fi40, F1: 4294967295, F2: v2}
                                                                        _ = v57
                                                                    }
                                                                }
                                                            default:
                                                                if len(v1) >= 2 {
                                                                    v1 = v1[1:len(v1)]
                                                                }
                                                            }
                                                            v1 = append(v1, v15)
                                                        }
                                                        v58 := GenS5{F0: 1}
                                                        _ = v58
                                                    }
                                                    var v59 int32 = 0
                                                    v3 = (v3 + v3)
                                                    _ = v59
                                                } else {
                                                    var v60 bool = v2
                                                    _ = v60
                                                }
                                                switch v15 {
                                                case 9223372036854775807:
                                                    var v61 uint16 = 42
                                                    _ = v61
                                                case 128:
                                                    v15 = (fi53 ^ (fi53 << ((4294967295) & 63)))
                                                }
                                                _ = fi53
                                            }
                                            _ = v52
                                        }
                                        var v62 float64 = -2.25
                                        _ = v62
                                    } else {
                                        switch 2147483647 {
                                        case 2147483647:
                                            for fi63 := 0; fi63 < 6; fi63++ {
                                                fmt.Println(v46)
                                                if v38 {
                                                    var v64 int = 690
                                                    v3 = v3
                                                    _ = v64
                                                } else {
                                                    var v65 uint8 = v4
                                                    var v66 int8 = 33
                                                    if (-222 >= -462) {
                                                        var v67 float32 = 0.1
                                                        if v38 {
                                                            v1 = append(v1, (fi40 / ((fi40) | 1)))
                                                            switch v3 {
                                                            case "":
                                                                fmt.Println(v66)
                                                            case "hello":
                                                                var v68 uint16 = 65535
                                                                _ = v68
                                                            case "\n\t":
                                                                fmt.Println(v2, v15, v46)
                                                            default:
                                                                if false {
                                                                    v69 := GenS10{F0: v4}
                                                                    _ = v69
                                                                } else {
                                                                    fmt.Println(v67, v65, v66)
                                                                    fmt.Println(v38, v15)
                                                                }
                                                            }
                                                        }
                                                        _ = v67
                                                    }
                                                    _ = v65
                                                    _ = v66
                                                }
                                                var v70 float64 = (v46 + (v46 - (v46 + v46)))
                                                var v71 bool = false
                                                _ = fi63
                                                _ = v70
                                                _ = v71
                                            }
                                        case 9223372036854775807:
                                            if ((1e10 / 3.14159) > float32(1.5)) {
                                                var v72 uint64 = 624
                                                _ = v72
                                            } else {
                                                v73 := GenS1{F0: v3}
                                                for fi74 := 0; fi74 < 10; fi74++ {
                                                    if len(v1) > 0 {
                                                        v1[0] = v15
                                                    }
                                                    fmt.Println(v38, fi40, v2)
                                                    _ = fi74
                                                }
                                                switch fi40 {
                                                case -314:
                                                    if true {
                                                        v1 = append(v1, (v15 + 32767))
                                                    }
                                                default:
                                                    switch ((v3 + ("foo bar" + ((v3 + v3) + (v3 + v3)))) + (v3 + v3)) {
                                                    case "":
                                                        if (!true) {
                                                            v75 := GenS9{F0: -1}
                                                            var v76 uint16 = 847
                                                            switch v3 {
                                                            case "":
                                                                v38 = (true && true)
                                                            case "\n\t":
                                                                fmt.Println(v3)
                                                            }
                                                            _ = v75
                                                            _ = v76
                                                        } else {
                                                            v38 = (-747 >= 710)
                                                        }
                                                    case "foo bar":
                                                        if false {
                                                            v2 = false
                                                            var v77 float32 = (0.1 / 1.5)
                                                            var v78 float32 = v77
                                                            _ = v77
                                                            _ = v78
                                                        }
                                                    default:
                                                        v79 := GenS11{F0: 908}
                                                        _ = v79
                                                    }
                                                }
                                                _ = v73
                                            }
                                        }
                                    }
                                    _ = v46
                                } else {
                                    for fi80 := 0; fi80 < 10; fi80++ {
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        fmt.Println(v2, fi80, v15)
                                        var v81 int16 = 37
                                        _ = fi80
                                        _ = v81
                                    }
                                }
                                v38 = v2
                            }
                            fmt.Println(v2, v15)
                            for fi82 := 0; fi82 < 12; fi82++ {
                                v38 = (65535 != 65535)
                                if (v3 >= ("\n\t" + (v3 + "foo bar"))) {
                                    var v83 uint64 = 4294967295
                                    v38 = false
                                    v84 := GenS12{F0: 2}
                                    _ = v83
                                    _ = v84
                                }
                                var v85 uint64 = 4294967295
                                fmt.Println(fi82, v15)
                                _ = fi82
                                _ = v85
                            }
                            _ = fi40
                        }
                    case "a":
                        var v86 int8 = -128
                        _ = v86
                    }
                    var v87 int64 = 777
                    switch "foo bar" {
                    case "hello":
                        switch "foo bar" {
                        case "hello":
                            if len(v1) > 0 {
                                v1[0] = 1
                            }
                        default:
                            var v88 uint16 = 65535
                            _ = v88
                        }
                    }
                    _ = v87
                }
            }
        case "":
            switch -589 {
            case -48:
                v3 = ("hello" + v3)
            }
        default:
            switch v3 {
            case "hello":
                var v89 int = -563
                _ = v89
            }
        }
        _ = v38
    }
    v90 := GenS7{F0: 768, F1: (v4 - v4)}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v90
}
