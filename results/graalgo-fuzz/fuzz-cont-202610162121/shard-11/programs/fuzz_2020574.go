package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8 }
type GenS2 struct { F0 int64; F1 int64; F2 int32 }
type GenS3 struct { F0 uint8; F1 uint64 }
type GenS4 struct { F0 bool; F1 int64 }
func genFn0() int16 {
    v5 := GenS1{F0: -128, F1: 42}
    for fi6 := 0; fi6 < 11; fi6++ {
        for fi7 := 0; fi7 < 9; fi7++ {
            v8 := []int{(fi7 & 42)}
            _ = fi7
            _ = v8
        }
        _ = fi6
    }
    _ = v5
    return -1
}
func genFn1(p0 uint64) uint64 {
    for fi9 := 0; fi9 < 3; fi9++ {
        if false {
            p0 = (p0 / ((714) | 1))
            for fi10 := 0; fi10 < 13; fi10++ {
                fmt.Println(fi9, p0, fi9)
                v11 := []int{(-(fi10)), (-(fi9))}
                _ = fi10
                _ = v11
            }
            p0 = p0
        } else {
            switch "hello" {
            case "":
                if (!false) {
                    for fi12 := 0; fi12 < 7; fi12++ {
                        fmt.Println(p0, fi12, p0)
                        _ = fi12
                    }
                    var v13 int = fi9
                    v13 -= (v13 & v13)
                    _ = v13
                }
            default:
                for fi14 := 0; fi14 < 1; fi14++ {
                    switch (("foo bar" + ("a" + "foo bar")) + (("foo bar" + "\n\t") + "a")) {
                    case "hello":
                        if true {
                            var v15 int32 = 42
                            var v16 bool = false
                            switch "a" {
                            case "hello":
                                switch fi9 {
                                case -1:
                                    v17 := GenS1{F0: -128, F1: 32}
                                    _ = v17
                                case 2147483647:
                                    if false {
                                        fmt.Println(v15, v15)
                                        v18 := GenS1{F0: 100, F1: 1}
                                        switch "" {
                                        case "foo bar":
                                            v16 = v16
                                        case "a":
                                            fmt.Println(fi9, p0)
                                        default:
                                            if v16 {
                                                var v19 uint8 = 128
                                                _ = v19
                                            } else {
                                                fmt.Println(v16, v15, v15)
                                                var v20 int32 = 0
                                                if v16 {
                                                    var v21 uint64 = (p0 & (p0 | (^(p0))))
                                                    v16 = false
                                                    var v22 string = (((("a" + ("" + ("hello" + ("" + (((("\n\t" + "hello") + "\n\t") + (("" + ("foo bar" + "foo bar")) + "foo bar")) + (("hello" + (("hello" + "foo bar") + ((("a" + ("a" + ("a" + ((("foo bar" + "\n\t") + (("hello" + ("\n\t" + "a")) + "")) + ("\n\t" + "\n\t"))))) + "") + "a"))) + "a")))))) + "a") + "hello") + "\n\t")
                                                    _ = v21
                                                    _ = v22
                                                }
                                                fmt.Println(fi14)
                                                _ = v20
                                            }
                                        }
                                        _ = v18
                                    }
                                case 32768:
                                    v23 := []int{-9223372036854775808, fi14}
                                    switch ("" + "foo bar") {
                                    case "":
                                        switch -102 {
                                        case -865:
                                            v15 -= v15
                                        case -9223372036854775808:
                                            var v24 uint32 = 4294967295
                                            _ = v24
                                        default:
                                            var v25 uint = 4294967295
                                            _ = v25
                                        }
                                    }
                                    var v26 bool = v16
                                    _ = v23
                                    _ = v26
                                default:
                                    var v27 uint8 = 255
                                    _ = v27
                                }
                            case "a":
                                fmt.Println(p0)
                                var v28 uint8 = 0
                                _ = v28
                            }
                            _ = v15
                            _ = v16
                        }
                    case "a":
                        var v29 int64 = 1
                        _ = v29
                    }
                    _ = fi14
                }
            }
            for fi30 := 0; fi30 < 10; fi30++ {
                var v31 float64 = 1e10
                var v32 int8 = -17
                if (!false) {
                    v33 := GenS2{F0: 0, F1: 77, F2: 42}
                    var v34 uint = 65535
                    _ = v33
                    _ = v34
                } else {
                    if (!false) {
                        var v35 bool = (((!(("" >= ((("" + ("foo bar" + "foo bar")) + ("a" + "foo bar")) + "\n\t")) || true)) || (false && (0.1 > 1e10))) || ((-32768 >= 32767) || (!(!(!true)))))
                        _ = v35
                    } else {
                        v32 = 100
                        var v36 bool = (!(1 <= 149))
                        if (!v36) {
                            fmt.Println(p0)
                            var v37 int16 = -1
                            _ = v37
                        }
                        _ = v36
                    }
                    var v38 int64 = -1
                    var v39 bool = true
                    _ = v38
                    _ = v39
                }
                _ = fi30
                _ = v31
                _ = v32
            }
            v40 := []int{(fi9 & (fi9 & (fi9 | (fi9 % (((fi9 - fi9)) | 1))))), (fi9 + 32767), fi9}
            _ = v40
        }
        _ = fi9
    }
    var v41 uint32 = 2147483648
    v41 -= (v41 * v41)
    _ = p0
    _ = v41
    return p0
}

func main() {
    var v1 float64 = 1e10
    var v2 float32 = (((float32(1.5) - 0.0) * 0.1) - 0.1)
    v3 := []int{127, 0, -2147483648, -2147483648}
    var v4 bool = false
    v4 = (true || v4)
    _ = len(v3)
    v1 += 0.1
    if len(v3) > 0 {
        v3[0] = 0
    }
    var v42 int64 = 761
    v3 = append(v3, 354)
    switch 32768 {
    case 2147483647:
        var v43 int16 = 1
        _ = v43
    case -1:
        if len(v3) > 0 {
            v3[0] = -63
        }
    }
    if true {
        var v44 float32 = v2
        for fi45 := 0; fi45 < 9; fi45++ {
            switch (fi45 & -1) {
            case -1:
                if (((((v1 / (v1 / 1.5)) + 0.0) - v1) != v1) || v4) {
                    switch (("foo bar" + ("a" + "")) + "foo bar") {
                    case "foo bar":
                        fmt.Println(fi45)
                    case "":
                        var v46 uint64 = 0
                        _ = v46
                    case "hello":
                        var v47 int64 = v42
                        _ = v47
                    default:
                        var v48 uint8 = 43
                        _ = v48
                    }
                    for fi49 := 0; fi49 < 14; fi49++ {
                        v42 -= v42
                        var v50 int16 = 1000
                        var v51 float64 = (float64(v2) - ((v1 - v1) + v1))
                        _ = fi49
                        _ = v50
                        _ = v51
                    }
                } else {
                    var v52 float32 = v44
                    _ = v52
                }
            case 32767:
                fmt.Println(v1, v2)
            }
            var v53 uint32 = 42
            v44 += v44
            _ = fi45
            _ = v53
        }
        _ = v44
    }
    v54 := GenS3{F0: 131, F1: 2}
    switch 2147483647 {
    case -644:
        var v55 uint = 509
        _ = v55
    case 9223372036854775807:
        var v56 bool = (2147483648 != 2147483648)
        _ = v56
    }
    switch "" {
    case "\n\t":
        if (v4 && v4) {
            switch "a" {
            case "":
                v1 += -2.25
            default:
                v4 = (v1 == v1)
            }
            var v57 uint16 = 0
            v58 := GenS1{F0: 127, F1: 4}
            _ = v57
            _ = v58
        } else {
            var v59 bool = false
            fmt.Println(v42, v2)
            _ = v59
        }
    case "a":
        for fi60 := 0; fi60 < 2; fi60++ {
            var v61 uint32 = 65535
            v62 := GenS2{F0: (v42 * (^(v42))), F1: v42, F2: 0}
            if len(v3) > 0 {
                v3[0] = (fi60 * 4294967295)
            }
            _ = fi60
            _ = v61
            _ = v62
        }
    default:
        for fi63 := 0; fi63 < 3; fi63++ {
            var v64 string = ""
            _ = fi63
            _ = v64
        }
    }
    var v65 string = ((("hello" + ("hello" + "hello")) + (("\n\t" + "foo bar") + "foo bar")) + (((("\n\t" + "a") + ("hello" + "hello")) + "\n\t") + "foo bar"))
    fmt.Println(v65, v2)
    for fi66 := 0; fi66 < 8; fi66++ {
        var v67 uint16 = 1
        switch (fi66 | 42) {
        case 127:
            for fi68 := 0; fi68 < 12; fi68++ {
                switch 0 {
                case 4294967295:
                    v69 := GenS4{F0: v4, F1: (v42 ^ v42)}
                    _ = v69
                case 128:
                    var v70 float32 = v2
                    _ = v70
                default:
                    var v71 float32 = (v2 - v2)
                    _ = v71
                }
                var v72 int8 = 100
                _ = fi68
                _ = v72
            }
        case 640:
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        case -2147483648:
            if (!((v4 || v4) && v4)) {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                v73 := GenS3{F0: uint8(v67), F1: 537}
                _ = v73
            } else {
                if (!((fi66 >= 1) && (!v4))) {
                    switch ("\n\t" + ("\n\t" + v65)) {
                    case "a":
                        v67 = 664
                    case "":
                        fmt.Println(v2)
                    }
                    var v74 uint16 = v67
                    _ = v74
                }
            }
        default:
            v75 := GenS3{F0: 2, F1: 42}
            _ = v75
        }
        v76 := GenS2{F0: (v42 * v42), F1: -438, F2: -2147483648}
        _ = fi66
        _ = v67
        _ = v76
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v65:", v65)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v54
    _ = v65
}
