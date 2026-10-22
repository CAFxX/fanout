package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32; F2 int16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int8; F1 uint16 }
type GenS5 struct { F0 string; F1 uint8 }
type GenS6 struct { F0 uint; F1 bool; F2 uint }
type GenS7 struct { F0 int; F1 int16; F2 int }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 uint8; F1 int16 }
type GenS10 struct { F0 uint16; F1 int64; F2 uint8 }
type GenS11 struct { F0 uint16; F1 bool }
type GenS12 struct { F0 int32; F1 bool }

func main() {
    var v1 uint8 = 2
    var v2 int16 = 581
    var v3 bool = false
    v4 := make(map[string]int)
    v5 := GenS1{F0: -434, F1: -729, F2: -303}
    var v6 int8 = 127
    if (-646 < 32767) {
        switch "foo bar" {
        case "\n\t":
            for fi7 := 0; fi7 < 19; fi7++ {
                v1 = v1
                _ = fi7
            }
        case "":
            if (false && v3) {
                v8 := []int{9223372036854775807, -235, 0}
                _ = v8
            } else {
                for fi9 := 0; fi9 < 19; fi9++ {
                    v6 = (v6 % ((v6) | 1))
                    _ = fi9
                }
                if v3 {
                    v3 = ((!true) || v3)
                    if (v6 != (v6 * 42)) {
                        var v10 int = 1
                        var v11 float64 = -2.25
                        _ = v10
                        _ = v11
                    } else {
                        fmt.Println(v6, v6, v1)
                    }
                    v12 := GenS2{F0: -928}
                    _ = v12
                } else {
                    var v13 int8 = 100
                    v3 = v3
                    _ = v13
                }
                for fi14 := 0; fi14 < 7; fi14++ {
                    var v15 bool = (-2147483648 == 0)
                    if true {
                        v16 := GenS1{F0: 42, F1: 1, F2: -1}
                        _ = v16
                    } else {
                        for fi17 := 0; fi17 < 3; fi17++ {
                            switch (fi14 + 1) {
                            case -573:
                                v1 -= v1
                                v18 := v4[("\n\t" + "")]
                                v19 := GenS3{F0: (v2 ^ v2)}
                                fmt.Println(v15, v6)
                                _ = v18
                                _ = v19
                            default:
                                var v20 uint64 = 101
                                _ = v20
                            }
                            _ = fi17
                        }
                        if v15 {
                            if v15 {
                                v21 := GenS4{F0: (^(v6)), F1: 2}
                                fmt.Println(v2)
                                v22 := GenS5{F0: ((("foo bar" + ("\n\t" + ((("\n\t" + "foo bar") + "a") + "hello"))) + ("foo bar" + "foo bar")) + ("foo bar" + "\n\t")), F1: v1}
                                _ = v21
                                _ = v22
                            } else {
                                fmt.Println(v6)
                                v2 = (v2 & v2)
                            }
                            var v23 int64 = 1
                            delete(v4, ("a" + (("hello" + (("foo bar" + "a") + (("" + "") + ("foo bar" + "\n\t")))) + ("foo bar" + "foo bar"))))
                            _ = v23
                        }
                    }
                    _ = fi14
                    _ = v15
                }
            }
        case "hello":
            var v24 bool = false
            _ = v24
        default:
            var v25 int64 = 655
            _ = v25
        }
        if true {
            switch 32768 {
            case 127:
                v26 := GenS5{F0: "foo bar", F1: v1}
                _ = v26
            }
            v1 += 128
            var v27 int = 4294967295
            _ = v27
        } else {
            if (9223372036854775807 == -1) {
                var v28 int8 = v6
                if (false && (int16(v6) <= v2)) {
                    for fi29 := 0; fi29 < 5; fi29++ {
                        for fi30 := 0; fi30 < 6; fi30++ {
                            v31 := []int{(fi29 * (fi30 >> ((2) & 63))), 32767, -288, fi29}
                            v1 += (^(v1))
                            switch "foo bar" {
                            case "\n\t":
                                var v32 float64 = 0.0
                                _ = v32
                            case "foo bar":
                                v1 = 255
                            case "":
                                var v33 int64 = -688
                                _ = v33
                            }
                            _ = fi30
                            _ = v31
                        }
                        _ = fi29
                    }
                    var v34 uint8 = (v1 & 128)
                    var v35 int8 = 1
                    _ = v34
                    _ = v35
                }
                var v36 int16 = (v2 % ((42) | 1))
                _ = v28
                _ = v36
            }
            var v37 int32 = -538
            v6 = (v6 / ((-1) | 1))
            _ = v37
        }
        delete(v4, (("hello" + "\n\t") + "\n\t"))
    }
    v3 = v3
    v38 := []int{-2147483648, -9223372036854775808}
    var v39 int32 = 524
    v40 := GenS2{F0: 576}
    for fi41 := 0; fi41 < 16; fi41++ {
        if v3 {
            fmt.Println(v1)
            delete(v4, "a")
        } else {
            var v42 bool = true
            delete(v4, ("a" + ("hello" + ("foo bar" + "hello"))))
            _ = len(v4)
            _ = v42
        }
        var v43 uint16 = uint16(v1)
        _ = fi41
        _ = v43
    }
    v44 := v4[("hello" + (("" + ("" + ("a" + (((((("hello" + "hello") + (("a" + "hello") + "foo bar")) + ("\n\t" + ("" + "\n\t"))) + "foo bar") + "a") + ("" + "\n\t"))))) + ("" + "hello")))]
    for fi45 := 0; fi45 < 17; fi45++ {
        for fi46 := 0; fi46 < 17; fi46++ {
            v3 = (41 > 0)
            if v3 {
                v47 := GenS2{F0: 42}
                fmt.Println(v2)
                switch "a" {
                case "\n\t":
                    fmt.Println(v6, v39)
                case "":
                    var v48 float32 = (1.5 * (((((((0.1 + 0.1) / float32(float64(((float32(1e10) / float32(3.14159)) / 1.5)))) * 0.1) + 0.1) + (1.5 + (float32((((0.1 - -2.25) + (((((1.5 - (0.0 * float64(0.0))) / float64((float32((0.1 - 1e10)) + 3.14159))) - 3.14159) * float64(((-2.25 / 3.14159) - 0.1))) - (float64(float32((1.5 - 3.14159))) / -2.25))) * -2.25)) - ((0.0 + 1.5) + 1.5)))) + (float32(1.5) - 1.5)) + 1e10))
                    _ = v48
                case "foo bar":
                    if len(v38) >= 2 {
                        v38 = v38[1:len(v38)]
                    }
                default:
                    v39 = v39
                }
                _ = v47
            } else {
                v3 = v3
                var v49 uint16 = 65535
                _ = v49
            }
            _ = fi46
        }
        _ = fi45
    }
    for fi50 := 0; fi50 < 11; fi50++ {
        v51 := GenS6{F0: 255, F1: (50 < (v1 ^ v1)), F2: 2}
        _ = fi50
        _ = v51
    }
    switch (v44 & v44) {
    case 42:
        v52 := GenS3{F0: (-(v2))}
        _ = v52
    case 128:
        v2 += (v2 >> ((4294967295) & 63))
    default:
        for fi53 := 0; fi53 < 4; fi53++ {
            fmt.Println(v1, v3)
            v54 := GenS7{F0: (v44 ^ fi53), F1: (v2 * v2), F2: 0}
            _ = fi53
            _ = v54
        }
    }
    switch (v44 ^ v44) {
    case 1:
        var v55 string = (("\n\t" + ("hello" + "a")) + (("" + (("foo bar" + ("\n\t" + "foo bar")) + ((("foo bar" + "hello") + "foo bar") + ("a" + "hello")))) + "hello"))
        _ = v55
    case 152:
        v3 = v3
    case 42:
        v56 := v4["\n\t"]
        _ = v56
    default:
        var v57 uint32 = 2147483648
        _ = v57
    }
    var v58 uint = 18446744073709551615
    if v3 {
        fmt.Println(v39, v6)
        if ("foo bar" != ("a" + "")) {
            if v3 {
                if v3 {
                    fmt.Println(v3)
                    v59 := GenS3{F0: v2}
                    _ = v59
                } else {
                    for fi60 := 0; fi60 < 19; fi60++ {
                        if true {
                            v44 = (fi60 % ((fi60) | 1))
                        } else {
                            if v3 {
                                v61 := GenS8{F0: 9223372036854775807}
                                fmt.Println(fi60, fi60)
                                _ = v61
                            } else {
                                var v62 int32 = v39
                                v63 := v4[(("foo bar" + ("foo bar" + (("a" + "a") + ("hello" + "a")))) + (("hello" + "a") + "foo bar"))]
                                for fi64 := 0; fi64 < 3; fi64++ {
                                    v38 = append(v38, v44)
                                    if v3 {
                                        switch (-(fi60)) {
                                        case 819:
                                            switch -9223372036854775808 {
                                            case 2147483647:
                                                v65 := GenS3{F0: 687}
                                                _ = v65
                                            default:
                                                var v66 int64 = 940
                                                _ = v66
                                            }
                                        case 149:
                                            switch "hello" {
                                            case "\n\t":
                                                fmt.Println(v3, v44)
                                            default:
                                                if ("\n\t" <= "\n\t") {
                                                    var v67 uint64 = 4294967295
                                                    _ = v67
                                                }
                                            }
                                        }
                                        v44 -= (v44 | fi64)
                                        switch ("a" + (("hello" + "a") + "foo bar")) {
                                        case "foo bar":
                                            _ = len(v38)
                                        }
                                    } else {
                                        v1 = 97
                                        fmt.Println(v63, v3, v44)
                                        if v3 {
                                            v68 := GenS9{F0: 1, F1: int16(v63)}
                                            _ = v68
                                        } else {
                                            if v3 {
                                                v2 = v2
                                                if v3 {
                                                    v2 = (v2 >> ((v58) & 63))
                                                }
                                            }
                                            if (false && ("foo bar" != ("foo bar" + "a"))) {
                                                fmt.Println(v58)
                                            }
                                            v44 = (^(v44))
                                        }
                                    }
                                    var v69 uint = (v58 + (v58 << ((v58) & 63)))
                                    switch "a" {
                                    case "a":
                                        if ((!v3) && v3) {
                                            if v3 {
                                                v2 = 0
                                                var v70 float32 = ((0.0 + 1.5) + ((3.14159 * 1e10) * 1.5))
                                                _ = v70
                                            } else {
                                                if v3 {
                                                    switch (fi60 ^ (fi60 & 178)) {
                                                    case 32767:
                                                        _ = len(v38)
                                                    case 773:
                                                        fmt.Println(v6, v39, v62)
                                                    case 796:
                                                        if (65535 <= 65535) {
                                                            var v71 uint = v58
                                                            _ = v71
                                                        }
                                                    default:
                                                        v4[("hello" + "hello")] = int(v69)
                                                    }
                                                    var v72 int16 = -1
                                                    if len(v38) >= 2 {
                                                        v38 = v38[1:len(v38)]
                                                    }
                                                    _ = v72
                                                }
                                                v6 -= (-(v6))
                                            }
                                            v69 = (v58 << ((v58) & 63))
                                            switch (fi60 % ((int(v58)) | 1)) {
                                            case 32768:
                                                var v73 int = (v44 << ((v69) & 63))
                                                _ = v73
                                            }
                                        } else {
                                            v58 -= (^(v58))
                                            v74 := GenS7{F0: fi64, F1: -1, F2: 55}
                                            _ = v74
                                        }
                                        var v75 string = ("hello" + "")
                                        _ = len(v38)
                                        _ = v75
                                    default:
                                        if v3 {
                                            v44 -= fi60
                                        }
                                        var v76 uint = 18446744073709551615
                                        _ = v76
                                    }
                                    _ = fi64
                                    _ = v69
                                }
                                _ = v62
                                _ = v63
                            }
                        }
                        var v77 uint32 = 4294967295
                        _ = fi60
                        _ = v77
                    }
                    for fi78 := 0; fi78 < 4; fi78++ {
                        fmt.Println(v44, fi78)
                        _ = fi78
                    }
                    switch (v44 % ((32768) | 1)) {
                    case 128:
                        fmt.Println(v1)
                    case 42:
                        v79 := GenS10{F0: 906, F1: 9223372036854775807, F2: (v1 ^ 1)}
                        _ = v79
                    default:
                        v80 := v4["a"]
                        _ = v80
                    }
                }
                if (!(v39 == (v39 * v39))) {
                    if true {
                        v81 := GenS11{F0: 42, F1: v3}
                        if (-412 > v44) {
                            switch "\n\t" {
                            case "foo bar":
                                switch (v44 % (((v44 | v44)) | 1)) {
                                case 0:
                                    if v3 {
                                        v1 -= v1
                                        if (!((v1 % ((v1) | 1)) != v1)) {
                                            var v82 uint32 = 990
                                            _ = v82
                                        } else {
                                            if v3 {
                                                for fi83 := 0; fi83 < 10; fi83++ {
                                                    for fi84 := 0; fi84 < 8; fi84++ {
                                                        v85 := GenS9{F0: (v1 << ((v58) & 63)), F1: (v2 ^ -1)}
                                                        v86 := GenS12{F0: v39, F1: ((v6 >> (((v58 * (v58 >> ((v58) & 63)))) & 63)) > (-(v6)))}
                                                        v87 := GenS2{F0: 42}
                                                        _ = fi84
                                                        _ = v85
                                                        _ = v86
                                                        _ = v87
                                                    }
                                                    _ = fi83
                                                }
                                            }
                                        }
                                        var v88 uint32 = 2147483648
                                        _ = v88
                                    } else {
                                        v39 = (v39 / ((v39) | 1))
                                        var v89 int32 = (^(v39))
                                        if len(v38) > 0 {
                                            v38[0] = (v44 / (((v44 - v44)) | 1))
                                        }
                                        _ = v89
                                    }
                                }
                            case "a":
                                var v90 int = v44
                                _ = v90
                            default:
                                var v91 uint = v58
                                _ = v91
                            }
                            var v92 float32 = (0.1 / 3.14159)
                            _ = v92
                        }
                        v2 = (v2 % ((int16(v39)) | 1))
                        _ = v81
                    }
                    var v93 int = -1
                    _ = v93
                } else {
                    var v94 uint8 = v1
                    _ = v94
                }
            }
        }
        if v3 {
            for fi95 := 0; fi95 < 4; fi95++ {
                v2 += (^(v2))
                var v96 int8 = int8(v2)
                _ = fi95
                _ = v96
            }
        }
    } else {
        v1 = (v1 >> (((^(v58))) & 63))
        fmt.Println(v58, v58)
    }
    var v97 bool = v3
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v38:", len(v38), v38)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v44:", v44)
    fmt.Println("v58:", v58)
    fmt.Println("v6:", v6)
    fmt.Println("v97:", v97)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v44
    _ = v5
    _ = v58
    _ = v6
    _ = v97
}
