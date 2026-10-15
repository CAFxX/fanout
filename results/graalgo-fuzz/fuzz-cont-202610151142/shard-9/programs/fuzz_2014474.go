package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string; F2 int8 }
type GenS2 struct { F0 int64; F1 uint16; F2 string }
type GenS3 struct { F0 int8; F1 uint16; F2 int16 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int; F1 uint }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 uint32; F1 int32 }
type GenS8 struct { F0 bool; F1 int }
type GenS9 struct { F0 uint32 }
type GenS10 struct { F0 string; F1 uint; F2 uint64 }
type GenS11 struct { F0 uint64 }
func genFn0() bool {
    switch ("hello" + ("foo bar" + "a")) {
    case "hello":
        var v5 string = "\n\t"
        _ = v5
    default:
        var v6 uint = 233
        _ = v6
    }
    switch 90 {
    case 465:
        v7 := GenS1{F0: 65535, F1: ("\n\t" + ("foo bar" + "a")), F2: 42}
        _ = v7
    case 128:
        var v8 int64 = 1
        _ = v8
    default:
        var v9 bool = false
        _ = v9
    }
    for fi10 := 0; fi10 < 7; fi10++ {
        var v11 bool = true
        if (!v11) {
            v12 := GenS1{F0: 42, F1: "\n\t", F2: 95}
            switch -3 {
            case 346:
                v13 := GenS2{F0: 372, F1: 255, F2: ("a" + "\n\t")}
                _ = v13
            case -1:
                fmt.Println(v11, v11)
            case -792:
                v11 = (!((((3.14159 - 0.1) + 0.0) + 0.0) == 1.5))
            default:
                if v11 {
                    var v14 uint16 = 274
                    fmt.Println(v11, v14)
                    _ = v14
                }
            }
            v11 = true
            _ = v12
        } else {
            v11 = true
            if (v11 || true) {
                v15 := GenS3{F0: 42, F1: 2, F2: 32767}
                for fi16 := 0; fi16 < 2; fi16++ {
                    v11 = false
                    _ = fi16
                }
                v17 := GenS4{F0: 0}
                _ = v15
                _ = v17
            }
            for fi18 := 0; fi18 < 16; fi18++ {
                v11 = (v11 || ("\n\t" != "foo bar"))
                v11 = (845 <= 65535)
                v19 := GenS2{F0: 992, F1: 57, F2: ("\n\t" + "foo bar")}
                _ = fi18
                _ = v19
            }
        }
        switch "foo bar" {
        case "":
            for fi20 := 0; fi20 < 2; fi20++ {
                if (1 == 2) {
                    fmt.Println(fi10)
                    v21 := GenS5{F0: (-(fi10)), F1: 4294967295}
                    switch "\n\t" {
                    case "hello":
                        switch fi10 {
                        case -9223372036854775808:
                            v22 := []int{(^(fi20))}
                            _ = v22
                        case 32768:
                            if (float64(3.14159) <= (float64(((((float32(1.5) / (1.5 + (-2.25 * 1.5))) - -2.25) + float32(1.5)) + (float32(-2.25) / (1.5 + 1.5)))) + 1.5)) {
                                v11 = (!(!(501 != 0)))
                                v23 := []int{-1, (fi10 * -714)}
                                fmt.Println(fi20, fi10, fi10)
                                _ = v23
                            } else {
                                switch fi10 {
                                case 2147483647:
                                    if ((!(v11 || false)) || false) {
                                        if (32767 >= 1) {
                                            v24 := GenS6{F0: ((fi20 * fi20) == 4294967295)}
                                            _ = v24
                                        } else {
                                            fmt.Println(fi10, fi20, fi10)
                                            v11 = v11
                                            if true {
                                                v11 = v11
                                            } else {
                                                fmt.Println(v11, fi10)
                                                var v25 uint16 = 975
                                                if (false || v11) {
                                                    v25 = (^(v25))
                                                    v26 := GenS7{F0: 2, F1: 0}
                                                    switch (fi20 - fi20) {
                                                    case 1:
                                                        var v27 uint64 = 142
                                                        _ = v27
                                                    case -1:
                                                        var v28 int16 = 951
                                                        _ = v28
                                                    case 808:
                                                        if ((!(false && v11)) || v11) {
                                                            v11 = v11
                                                            var v29 uint64 = 18446744073709551615
                                                            switch "\n\t" {
                                                            case "hello":
                                                                var v30 float32 = (0.0 / 1.5)
                                                                _ = v30
                                                            case "a":
                                                                if (-437 < -1) {
                                                                    var v31 int8 = 127
                                                                    _ = v31
                                                                }
                                                            case "foo bar":
                                                                var v32 uint64 = 1
                                                                _ = v32
                                                            default:
                                                                fmt.Println(v29, v29)
                                                            }
                                                            _ = v29
                                                        }
                                                        fmt.Println(fi10, fi20, fi20)
                                                    }
                                                    _ = v26
                                                } else {
                                                    var v33 int16 = 0
                                                    _ = v33
                                                }
                                                _ = v25
                                            }
                                        }
                                    }
                                case 759:
                                    v11 = (162 != 127)
                                default:
                                    var v34 float64 = 3.14159
                                    _ = v34
                                }
                                v11 = false
                            }
                        }
                        fmt.Println(fi10)
                        v35 := GenS6{F0: v11}
                        _ = v35
                    default:
                        var v36 uint16 = 32767
                        _ = v36
                    }
                    _ = v21
                }
                _ = fi20
            }
        case "a":
            var v37 int8 = 1
            _ = v37
        case "\n\t":
            switch 9223372036854775807 {
            case 128:
                var v38 uint64 = 505
                _ = v38
            }
        }
        _ = fi10
        _ = v11
    }
    if (1 <= 15) {
        var v39 uint32 = 2147483647
        _ = v39
    } else {
        switch -1 {
        case -1:
            var v40 int64 = -1000
            _ = v40
        }
    }
    if (128 >= 128) {
        switch 42 {
        case 4294967295:
            v41 := []int{}
            _ = v41
        case -9223372036854775808:
            var v42 string = "foo bar"
            _ = v42
        case -816:
            for fi43 := 0; fi43 < 9; fi43++ {
                v44 := GenS8{F0: false, F1: fi43}
                v45 := []int{(fi43 + (^(fi43))), fi43, (fi43 ^ fi43)}
                fmt.Println(fi43)
                _ = fi43
                _ = v44
                _ = v45
            }
        default:
            v46 := []int{2147483647}
            _ = v46
        }
        switch (((("hello" + (("hello" + "foo bar") + "a")) + "") + "\n\t") + "\n\t") {
        case "\n\t":
            v47 := []int{32767, -766}
            _ = v47
        default:
            if true {
                v48 := []int{}
                _ = v48
            } else {
                switch 42 {
                case -293:
                    switch -518 {
                    case -240:
                        for fi49 := 0; fi49 < 3; fi49++ {
                            switch ("hello" + "") {
                            case "":
                                fmt.Println(fi49)
                            }
                            var v50 uint32 = 393
                            var v51 int64 = -733
                            _ = fi49
                            _ = v50
                            _ = v51
                        }
                    case -88:
                        var v52 int8 = -1
                        _ = v52
                    }
                case -1:
                    v53 := GenS9{F0: 847}
                    _ = v53
                }
            }
        }
    }
    return (!(0 <= 1))
}
func genFn1(p0 int8) uint16 {
    p0 = (^(p0))
    p0 = -128
    var v54 uint16 = 2
    _ = p0
    _ = v54
    return 26
}

func main() {
    var v1 float32 = 1e10
    var v2 float64 = 3.14159
    var v3 float32 = v1
    v4 := []int{0, -9223372036854775808, 127, 724}
    if (2147483648 > 2147483648) {
        for fi55 := 0; fi55 < 10; fi55++ {
            v4 = append(v4, fi55)
            _ = fi55
        }
    }
    v3 = (v1 - float32(v2))
    switch 555 {
    case -9223372036854775808:
        v56 := GenS8{F0: ((-128 >= -53) && false), F1: 32768}
        _ = v56
    case 4294967295:
        v4 = append(v4, 1)
    default:
        switch 9223372036854775807 {
        case 127:
            _ = len(v4)
        case 766:
            if (false && (!(0 == 485))) {
                if len(v4) > 0 {
                    v4[0] = 32768
                }
                fmt.Println(v1)
            } else {
                v57 := GenS8{F0: true, F1: 42}
                _ = v57
            }
        case 0:
            switch -195 {
            case 32768:
                if len(v4) > 0 {
                    v4[0] = 2147483647
                }
            }
        }
    }
    switch "\n\t" {
    case "hello":
        var v58 string = "\n\t"
        _ = v58
    case "foo bar":
        if (!(!(9223372036854775807 <= -9223372036854775808))) {
            if (false || true) {
                v2 += (v2 * float64(v3))
                switch 581 {
                case 0:
                    if (42 != 235) {
                        var v59 float32 = (0.1 / 1e10)
                        if true {
                            var v60 float32 = 1.5
                            _ = v60
                        }
                        _ = v59
                    }
                case 689:
                    v4 = append(v4, 127)
                }
                v3 = v1
            } else {
                for fi61 := 0; fi61 < 1; fi61++ {
                    switch ((("hello" + "foo bar") + "foo bar") + ("a" + "")) {
                    case "hello":
                        v1 = v3
                    default:
                        fmt.Println(v1)
                    }
                    _ = fi61
                }
                var v62 uint32 = 2
                _ = v62
            }
            v63 := GenS5{F0: -346, F1: 2}
            fmt.Println(v1)
            _ = v63
        } else {
            if ("foo bar" > "foo bar") {
                v4 = append(v4, 771)
                if true {
                    var v64 int16 = -1
                    var v65 float64 = float64(float32((float64(1e10) * float64(v1))))
                    v1 = 1e10
                    _ = v64
                    _ = v65
                } else {
                    v66 := GenS10{F0: "a", F1: 859, F2: 2}
                    var v67 float32 = -2.25
                    v3 += v1
                    _ = v66
                    _ = v67
                }
            }
            v68 := GenS1{F0: 1, F1: ("\n\t" + ("foo bar" + ("foo bar" + "\n\t"))), F2: 58}
            for fi69 := 0; fi69 < 14; fi69++ {
                v1 = 0.0
                var v70 int32 = 885
                _ = fi69
                _ = v70
            }
            _ = v68
        }
    default:
        var v71 string = "hello"
        _ = v71
    }
    switch 530 {
    case -780:
        v1 = 0.0
    }
    if (false || (2 > 0)) {
        v1 = v1
    } else {
        v2 = v2
    }
    var v72 int = -144
    v73 := GenS2{F0: 112, F1: 188, F2: "foo bar"}
    v3 -= v3
    v4 = append(v4, (v72 + 597))
    var v74 int = -2147483648
    v72 += (v74 | -9223372036854775808)
    v75 := GenS2{F0: 226, F1: 255, F2: ("a" + ("hello" + ("foo bar" + "hello")))}
    v76 := GenS11{F0: 244}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v72:", v72)
    fmt.Println("v74:", v74)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v72
    _ = v73
    _ = v74
    _ = v75
    _ = v76
}
