package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64; F2 uint }
type GenS2 struct { F0 uint8; F1 int64; F2 uint16 }
type GenS3 struct { F0 uint64; F1 int16 }
type GenS4 struct { F0 string; F1 int64 }
func genFn0(p0 string, p1 int32, p2 uint) uint64 {
    v6 := []int{2147483647, -1}
    fmt.Println(p1)
    switch 127 {
    case 522:
        v7 := GenS1{F0: 1, F1: 0, F2: 0}
        _ = v7
    case 339:
        var v8 uint32 = 2147483648
        _ = v8
    default:
        if (false || (false && true)) {
            var v9 int64 = 924
            var v10 float64 = 3.14159
            _ = v10
            _ = v9
        }
    }
    var v11 uint = p2
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v6
    return uint64(p2)
}
func genFn1() uint {
    switch "" {
    case "\n\t":
        v12 := []int{-1, 32768, 2147483647}
        _ = v12
    default:
        if (true || ((!false) || true)) {
            for fi13 := 0; fi13 < 14; fi13++ {
                fmt.Println(fi13)
                for fi14 := 0; fi14 < 12; fi14++ {
                    var v15 uint64 = 4294967295
                    _ = fi14
                    _ = v15
                }
                _ = fi13
            }
            switch -2147483648 {
            case 669:
                if (!false) {
                    var v16 string = "foo bar"
                    v17 := []int{-77, 0, 277}
                    _ = v16
                    _ = v17
                }
            case 0:
                var v18 int = 32768
                _ = v18
            case -1:
                var v19 uint16 = 245
                _ = v19
            default:
                v20 := GenS1{F0: 255, F1: 1, F2: 255}
                _ = v20
            }
        } else {
            var v21 int8 = 94
            _ = v21
        }
    }
    for fi22 := 0; fi22 < 13; fi22++ {
        v23 := GenS2{F0: 182, F1: 400, F2: 0}
        v24 := GenS3{F0: 644, F1: 1000}
        var v25 uint = 18446744073709551615
        _ = fi22
        _ = v23
        _ = v24
        _ = v25
    }
    switch -9223372036854775808 {
    case 329:
        switch ("" + "hello") {
        case "a":
            var v26 float32 = 3.14159
            _ = v26
        case "foo bar":
            var v27 uint32 = 2147483648
            _ = v27
        case "hello":
            for fi28 := 0; fi28 < 4; fi28++ {
                var v29 int8 = -128
                switch -1 {
                case 9223372036854775807:
                    for fi30 := 0; fi30 < 3; fi30++ {
                        var v31 uint16 = 0
                        if true {
                            var v32 uint8 = 42
                            _ = v32
                        }
                        _ = fi30
                        _ = v31
                    }
                }
                _ = fi28
                _ = v29
            }
        default:
            if true {
                for fi33 := 0; fi33 < 11; fi33++ {
                    if false {
                        var v34 uint8 = 0
                        v35 := GenS3{F0: 1, F1: 1000}
                        _ = v34
                        _ = v35
                    }
                    var v36 string = ("hello" + (("foo bar" + ((("hello" + "hello") + "hello") + ("" + (("hello" + "") + "foo bar")))) + "\n\t"))
                    var v37 float64 = 0.0
                    _ = fi33
                    _ = v36
                    _ = v37
                }
                var v38 uint = 0
                _ = v38
            }
        }
    case 32768:
        var v39 bool = (genFn0("\n\t", -28, 569) != 18446744073709551615)
        _ = v39
    case -1:
        switch ((("" + "") + ((("foo bar" + ((("hello" + (("hello" + "foo bar") + "a")) + "a") + ("hello" + "foo bar"))) + ("hello" + "a")) + "hello")) + "") {
        case "":
            switch ("\n\t" + "") {
            case "foo bar":
                var v40 int64 = -187
                _ = v40
            case "\n\t":
                var v41 uint = 0
                _ = v41
            default:
                v42 := []int{}
                _ = v42
            }
        case "foo bar":
            var v43 int = 32767
            _ = v43
        default:
            var v44 uint8 = 2
            _ = v44
        }
    }
    v45 := []int{128, -9223372036854775808}
    if false {
        switch 756 {
        case 127:
            var v46 uint32 = 42
            _ = v46
        case 2147483647:
            switch 211 {
            case 0:
                v47 := GenS4{F0: "hello", F1: -9223372036854775808}
                _ = v47
            case 42:
                var v48 int32 = 975
                _ = v48
            case 2147483647:
                var v49 int32 = 0
                _ = v49
            default:
                v50 := GenS1{F0: 0, F1: 2, F2: 18446744073709551615}
                _ = v50
            }
        default:
            var v51 bool = (false || false)
            _ = v51
        }
        for fi52 := 0; fi52 < 12; fi52++ {
            var v53 int = -9223372036854775808
            _ = fi52
            _ = v53
        }
        var v54 float32 = 0.0
        _ = v54
    } else {
        v55 := GenS2{F0: 127, F1: -1, F2: 2}
        if true {
            var v56 float64 = 0.1
            v56 = (-2.25 / 1.0)
            _ = v56
        } else {
            if false {
                if (-2147483648 > 0) {
                    for fi57 := 0; fi57 < 12; fi57++ {
                        if (((1e10 - 1.5) + ((0.1 * 1e10) - ((((0.0 / 1.0) - -2.25) * (1e10 * 3.14159)) - 0.1))) >= (float64(0.0) * ((0.1 / 1.5) + ((((1e10 / 3.14159) - 3.14159) - float64(float32(float64(0.0)))) + (3.14159 + 1e10))))) {
                            var v58 int8 = 0
                            _ = v58
                        }
                        _ = fi57
                    }
                    var v59 bool = (!(-2.25 != float32(float64(((((0.0 + -2.25) + 1.5) + 0.0) * float32(1e10))))))
                    _ = v59
                }
                var v60 int = 0
                v61 := GenS3{F0: 4294967295, F1: -180}
                _ = v60
                _ = v61
            } else {
                v62 := GenS3{F0: 611, F1: -1}
                var v63 int16 = 32767
                _ = v62
                _ = v63
            }
        }
        if len(v45) >= 2 {
            v45 = v45[1:len(v45)]
        }
        _ = v55
    }
    _ = v45
    return 65535
}

func main() {
    var v1 int8 = 62
    v2 := []int{}
    var v3 int32 = -2147483648
    v4 := []string{"hello"}
    var v5 int64 = 9223372036854775807
    v3 += v3
    var v64 bool = true
    for fi65 := 0; fi65 < 5; fi65++ {
        switch "foo bar" {
        case "hello":
            for fi66 := 0; fi66 < 7; fi66++ {
                switch ("hello" + ("\n\t" + "foo bar")) {
                case "foo bar":
                    v1 = -128
                case "hello":
                    v4 = append(v4, (("a" + "hello") + ""))
                case "":
                    v5 += v5
                }
                v5 += 9223372036854775807
                _ = fi66
            }
        case "foo bar":
            switch 42 {
            case 562:
                for fi67 := 0; fi67 < 10; fi67++ {
                    v68 := GenS2{F0: 85, F1: (^(v5)), F2: 65535}
                    v1 -= (v1 % ((v1) | 1))
                    var v69 uint64 = 640
                    switch 128 {
                    case 373:
                        switch -1 {
                        case 9223372036854775807:
                            var v70 int16 = 438
                            _ = v70
                        case 32767:
                            v5 -= v5
                        case -422:
                            var v71 int16 = 1
                            _ = v71
                        default:
                            v5 -= (v5 % ((-1) | 1))
                        }
                    default:
                        v64 = (!v64)
                    }
                    v72 := GenS2{F0: 138, F1: v5, F2: 0}
                    _ = fi67
                    _ = v68
                    _ = v69
                    _ = v72
                }
            case 2147483647:
                for fi73 := 0; fi73 < 20; fi73++ {
                    var v74 int8 = (v1 << ((642) & 63))
                    v5 -= (v5 >> ((607) & 63))
                    _ = fi73
                    _ = v74
                }
            case 32767:
                if true {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                } else {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    for fi75 := 0; fi75 < 14; fi75++ {
                        v76 := GenS4{F0: "foo bar", F1: 0}
                        fmt.Println(v5, fi75)
                        _ = fi75
                        _ = v76
                    }
                    fmt.Println(v5)
                }
            }
        case "":
            if ((-1 >= v5) || v64) {
                for fi77 := 0; fi77 < 1; fi77++ {
                    var v78 int8 = v1
                    _ = fi77
                    _ = v78
                }
                var v79 string = "foo bar"
                _ = v79
            }
        }
        v80 := GenS2{F0: 157, F1: (v5 ^ 414), F2: 255}
        _ = fi65
        _ = v80
    }
    switch 42 {
    case 4294967295:
        v81 := GenS4{F0: (("" + "") + ("foo bar" + "foo bar")), F1: (-(v5))}
        _ = v81
    case 128:
        _ = len(v4)
    }
    for fi82 := 0; fi82 < 3; fi82++ {
        switch fi82 {
        case 32768:
            v5 += 0
        }
        _ = fi82
    }
    switch 128 {
    case -621:
        switch 42 {
        case 32767:
            var v83 uint64 = 256
            _ = v83
        case 127:
            switch 9223372036854775807 {
            case 0:
                var v84 uint16 = 255
                _ = v84
            }
        case 2147483647:
            var v85 uint = 18446744073709551615
            _ = v85
        }
    case 18:
        var v86 uint32 = 804
        _ = v86
    case 709:
        var v87 uint32 = 1
        _ = v87
    }
    for fi88 := 0; fi88 < 1; fi88++ {
        fmt.Println(v5, v1, v1)
        _ = len(v4)
        _ = fi88
    }
    v89 := GenS3{F0: 18446744073709551615, F1: -1}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v64:", v64)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v64
    _ = v89
}
