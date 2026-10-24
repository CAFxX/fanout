package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int; F1 bool }
type GenS3 struct { F0 uint64; F1 int16 }
type GenS4 struct { F0 uint32; F1 bool; F2 uint }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint; F1 uint16 }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 bool; F1 int8; F2 int8 }
type GenS9 struct { F0 uint16; F1 uint16; F2 int32 }

func main() {
    v1 := make(map[string]bool)
    v2 := []string{}
    var v3 uint = 2
    var v4 uint = (v3 & 863)
    var v5 bool = false
    fmt.Println(v3, v4, v5)
    if ((18446744073709551615 < 42) || (-2.25 > float32(0.0))) {
        v5 = v5
        if (4294967295 < 2) {
            v6 := GenS1{F0: 2}
            v4 = v3
            _ = v6
        }
    } else {
        for fi7 := 0; fi7 < 11; fi7++ {
            switch ("\n\t" + "") {
            case "\n\t":
                switch ("a" + "foo bar") {
                case "foo bar":
                    v8 := v1[(("" + "\n\t") + "foo bar")]
                    _ = v8
                case "\n\t":
                    v2 = append(v2, ("hello" + "foo bar"))
                case "a":
                    if len(v2) > 0 {
                        v2[0] = ("\n\t" + (("hello" + "\n\t") + "hello"))
                    }
                }
            default:
                var v9 bool = false
                _ = v9
            }
            fmt.Println(v5, v5, fi7)
            _ = fi7
        }
        if (-38 < 906) {
            var v10 int = -1
            delete(v1, ((("foo bar" + "\n\t") + "foo bar") + ""))
            var v11 float64 = float64(1.5)
            _ = v10
            _ = v11
        } else {
            v3 = (v3 - (v4 / ((v3) | 1)))
            var v12 int32 = 0
            var v13 int8 = 100
            _ = v12
            _ = v13
        }
    }
    if (128 > 128) {
        var v14 string = ("hello" + "hello")
        _ = v14
    }
    for fi15 := 0; fi15 < 8; fi15++ {
        switch -444 {
        case 32767:
            if v5 {
                v16 := GenS2{F0: fi15, F1: v5}
                var v17 float32 = float32(-2.25)
                switch (^(fi15)) {
                case -976:
                    var v18 int8 = 0
                    _ = v18
                case 127:
                    var v19 uint64 = 452
                    _ = v19
                case 1:
                    switch fi15 {
                    case -9223372036854775808:
                        v3 -= v3
                    case 4294967295:
                        _ = len(v2)
                    case 42:
                        if ((1e10 * 3.14159) > float64(v17)) {
                            if ((v5 && ((3.14159 >= float64((v17 - v17))) || (42 == 42))) && ((fi15 / ((127) | 1)) > fi15)) {
                                v5 = v5
                                switch ("\n\t" + "foo bar") {
                                case "":
                                    var v20 uint = 0
                                    _ = v20
                                default:
                                    var v21 uint16 = 135
                                    _ = v21
                                }
                                v22 := GenS3{F0: 329, F1: 42}
                                _ = v22
                            }
                            if true {
                                var v23 int = fi15
                                if false {
                                    var v24 bool = (0 >= -2147483648)
                                    if (-32768 >= 565) {
                                        delete(v1, ("" + ((("hello" + ("a" + (("\n\t" + "a") + "foo bar"))) + ("hello" + "hello")) + ("a" + ("\n\t" + "hello")))))
                                        var v25 uint16 = 2
                                        var v26 uint8 = 1
                                        _ = v25
                                        _ = v26
                                    } else {
                                        if (0 < 1) {
                                            if (-1 != -56) {
                                                var v27 float32 = v17
                                                _ = v27
                                            } else {
                                                var v28 uint = 333
                                                v29 := v1[("foo bar" + ("foo bar" + ""))]
                                                _ = v28
                                                _ = v29
                                            }
                                            switch ("a" + "\n\t") {
                                            case "":
                                                var v30 bool = true
                                                _ = v30
                                            default:
                                                fmt.Println(v4)
                                            }
                                        }
                                        if v24 {
                                            switch ("foo bar" + "") {
                                            case "a":
                                                switch 0 {
                                                case 128:
                                                    v3 -= v4
                                                case -9223372036854775808:
                                                    fmt.Println(fi15)
                                                }
                                            case "":
                                                v23 = (v23 + (fi15 + (^(v23))))
                                            case "foo bar":
                                                fmt.Println(v23)
                                            }
                                            v24 = v5
                                            for fi31 := 0; fi31 < 2; fi31++ {
                                                switch "hello" {
                                                case "\n\t":
                                                    if (1e10 <= 1.5) {
                                                        v24 = (false && true)
                                                    } else {
                                                        if v5 {
                                                            switch 127 {
                                                            case -2147483648:
                                                                v24 = v24
                                                            default:
                                                                fmt.Println(v5, v23, v17)
                                                            }
                                                        }
                                                        v4 = v4
                                                        v32 := GenS2{F0: (fi31 / (((^(v23))) | 1)), F1: true}
                                                        _ = v32
                                                    }
                                                case "hello":
                                                    v33 := GenS4{F0: 2147483647, F1: (((("" + "\n\t") + "hello") + "foo bar") != ("a" + ("hello" + (("a" + ("" + ("\n\t" + "a"))) + "foo bar")))), F2: 4294967295}
                                                    _ = v33
                                                case "foo bar":
                                                    if (770 == (v3 / (((^(v3))) | 1))) {
                                                        v34 := GenS3{F0: 2, F1: 1}
                                                        v5 = (((v4 << (((v3 % (((v3 ^ 700)) | 1))) & 63)) > 42) || (!true))
                                                        _ = v34
                                                    }
                                                default:
                                                    v23 = 32767
                                                }
                                                _ = fi31
                                            }
                                        } else {
                                            if false {
                                                v35 := GenS5{F0: (fi15 / (((fi15 + (-(v23)))) | 1))}
                                                delete(v1, "")
                                                _ = v35
                                            }
                                        }
                                        v3 -= 255
                                    }
                                    v23 -= fi15
                                    _ = v24
                                }
                                _ = v23
                            }
                        }
                    default:
                        _ = len(v2)
                    }
                }
                _ = v16
                _ = v17
            }
        default:
            if v5 {
                _ = len(v1)
                for fi36 := 0; fi36 < 8; fi36++ {
                    var v37 uint32 = 1
                    if v5 {
                        var v38 float64 = 0.0
                        v39 := v1["foo bar"]
                        _ = v38
                        _ = v39
                    }
                    v4 += 42
                    _ = fi36
                    _ = v37
                }
                v40 := GenS1{F0: 270}
                _ = v40
            } else {
                switch "hello" {
                case "hello":
                    v4 = 18446744073709551615
                case "foo bar":
                    if (v5 && v5) {
                        if len(v2) > 0 {
                            v2[0] = ("hello" + "\n\t")
                        }
                        if len(v2) > 0 {
                            v2[0] = ("\n\t" + "\n\t")
                        }
                    } else {
                        v3 = v3
                    }
                case "a":
                    switch "a" {
                    case "foo bar":
                        var v41 uint64 = 42
                        _ = v41
                    case "":
                        var v42 float32 = 1.5
                        _ = v42
                    case "hello":
                        var v43 float64 = 1e10
                        _ = v43
                    default:
                        var v44 float32 = 1.5
                        _ = v44
                    }
                default:
                    for fi45 := 0; fi45 < 2; fi45++ {
                        fmt.Println(fi15)
                        _ = fi45
                    }
                }
                var v46 uint64 = 233
                v5 = (!v5)
                _ = v46
            }
        }
        if (false && v5) {
            if v5 {
                var v47 string = "hello"
                for fi48 := 0; fi48 < 2; fi48++ {
                    switch (fi15 << ((v3) & 63)) {
                    case -129:
                        var v49 int16 = -32768
                        _ = v49
                    case 295:
                        v4 += (v3 << ((v4) & 63))
                    case 32767:
                        v5 = v5
                    }
                    _ = fi48
                }
                v50 := GenS4{F0: 4294967295, F1: false, F2: v4}
                _ = v47
                _ = v50
            }
        } else {
            var v51 int16 = -1
            switch fi15 {
            case 32767:
                fmt.Println(v3)
            case 88:
                _ = len(v2)
            case 544:
                _ = len(v1)
            default:
                for fi52 := 0; fi52 < 16; fi52++ {
                    v4 += (v3 % ((v3) | 1))
                    if v5 {
                        var v53 float64 = 3.14159
                        if v5 {
                            _ = len(v2)
                            if v5 {
                                var v54 uint8 = 1
                                var v55 float32 = 0.0
                                if (true || ((252 <= 655) || v5)) {
                                    fmt.Println(fi15)
                                }
                                _ = v54
                                _ = v55
                            } else {
                                v56 := GenS4{F0: 828, F1: (!(352 < 0)), F2: 246}
                                switch "a" {
                                case "\n\t":
                                    v4 += v4
                                }
                                var v57 uint32 = uint32(v3)
                                v58 := GenS6{F0: 255, F1: 461}
                                switch (("\n\t" + (((("foo bar" + "hello") + ((("hello" + "") + "a") + ("a" + ("a" + "foo bar")))) + "foo bar") + "a")) + "hello") {
                                case "hello":
                                    var v59 uint8 = 175
                                    var v60 bool = (false || v5)
                                    var v61 int16 = (v51 & (v51 - 0))
                                    _ = v59
                                    _ = v60
                                    _ = v61
                                case "a":
                                    v62 := GenS7{F0: 1}
                                    _ = v62
                                case "":
                                    var v63 int64 = 42
                                    _ = v63
                                }
                                _ = v56
                                _ = v57
                                _ = v58
                            }
                        }
                        fmt.Println(v51, v53, v3)
                        _ = v53
                    }
                    var v64 float64 = -2.25
                    _ = fi52
                    _ = v64
                }
            }
            if false {
                var v65 float64 = (0.1 - float64((1e10 * (3.14159 + float32(3.14159)))))
                _ = v65
            }
            _ = v51
        }
        var v66 float64 = -2.25
        _ = fi15
        _ = v66
    }
    if (!v5) {
        for fi67 := 0; fi67 < 14; fi67++ {
            var v68 uint32 = 2
            var v69 float64 = 0.1
            v70 := GenS6{F0: 2, F1: 255}
            _ = fi67
            _ = v68
            _ = v69
            _ = v70
        }
        v4 += v4
    } else {
        fmt.Println(v4, v3, v5)
    }
    delete(v1, "\n\t")
    switch 2147483647 {
    case 0:
        if true {
            var v71 bool = (0 == -9223372036854775808)
            v72 := GenS8{F0: false, F1: 0, F2: 0}
            _ = v71
            _ = v72
        } else {
            v73 := GenS9{F0: 2, F1: 60, F2: 42}
            v74 := v1["\n\t"]
            var v75 string = "hello"
            _ = v73
            _ = v74
            _ = v75
        }
    case 116:
        v76 := GenS8{F0: (194 != 876), F1: 34, F2: 127}
        _ = v76
    }
    v77 := GenS9{F0: 0, F1: 1, F2: 2147483647}
    var v78 uint64 = 2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v78:", v78)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v77
    _ = v78
}
