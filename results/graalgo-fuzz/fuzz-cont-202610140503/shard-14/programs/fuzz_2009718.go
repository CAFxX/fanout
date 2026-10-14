package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8 }
type GenS2 struct { F0 uint8; F1 int32; F2 int16 }
type GenS3 struct { F0 int32; F1 uint64; F2 string }
type GenS4 struct { F0 bool; F1 string }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int16; F1 bool; F2 int16 }
type GenS7 struct { F0 string; F1 uint16 }
func genFn0() string {
    for fi5 := 0; fi5 < 10; fi5++ {
        v6 := GenS1{F0: 608, F1: 0}
        for fi7 := 0; fi7 < 6; fi7++ {
            if false {
                v8 := []int{0, (fi7 << ((2) & 63))}
                var v9 int64 = 1
                _ = v8
                _ = v9
            }
            _ = fi7
        }
        fmt.Println(fi5)
        _ = fi5
        _ = v6
    }
    var v10 int64 = 0
    v10 = v10
    _ = v10
    return "foo bar"
}

func main() {
    var v1 uint = 1
    v2 := []int{-168, 0, 32767, 766}
    v3 := make(map[string]bool)
    v4 := []int{}
    v11 := v3["a"]
    var v12 int8 = 100
    v13 := v3[("" + "hello")]
    v11 = v11
    v14 := GenS2{F0: 42, F1: 204, F2: 1000}
    for fi15 := 0; fi15 < 11; fi15++ {
        for fi16 := 0; fi16 < 10; fi16++ {
            if v13 {
                var v17 string = ""
                switch fi16 {
                case -616:
                    var v18 float64 = ((-2.25 * ((0.1 - 3.14159) * float64(float32(0.0)))) / 1.5)
                    _ = v18
                case -438:
                    v17 = v17
                }
                v19 := GenS2{F0: 225, F1: 42, F2: 42}
                _ = v17
                _ = v19
            } else {
                var v20 int64 = int64(v1)
                if true {
                    var v21 bool = (true || v11)
                    _ = v21
                }
                var v22 uint8 = 100
                _ = v20
                _ = v22
            }
            if v11 {
                v23 := v3["\n\t"]
                _ = v23
            }
            var v24 float64 = 1e10
            var v25 bool = (false && (-1 <= 2147483647))
            _ = len(v2)
            _ = fi16
            _ = v24
            _ = v25
        }
        fmt.Println(v13, v11, v13)
        _ = fi15
    }
    if true {
        fmt.Println(v12)
        for fi26 := 0; fi26 < 13; fi26++ {
            for fi27 := 0; fi27 < 6; fi27++ {
                var v28 bool = v13
                switch (fi27 % ((int(v12)) | 1)) {
                case -905:
                    fmt.Println(fi27, v1)
                case -9223372036854775808:
                    if v13 {
                        fmt.Println(v28, v13)
                        v12 = v12
                    }
                case 4294967295:
                    v28 = ("a" < "")
                default:
                    var v29 int32 = int32(v12)
                    _ = v29
                }
                if true {
                    fmt.Println(v11, v12, v1)
                    v30 := v3[("foo bar" + "a")]
                    _ = v30
                }
                _ = fi27
                _ = v28
            }
            switch "hello" {
            case "\n\t":
                v31 := GenS3{F0: 912, F1: 1, F2: "foo bar"}
                _ = v31
            }
            var v32 int32 = -1
            _ = fi26
            _ = v32
        }
        for fi33 := 0; fi33 < 7; fi33++ {
            var v34 float64 = 0.0
            v12 = (v12 & v12)
            _ = fi33
            _ = v34
        }
    } else {
        for fi35 := 0; fi35 < 3; fi35++ {
            switch "foo bar" {
            case "\n\t":
                switch fi35 {
                case 36:
                    var v36 uint = 1
                    _ = v36
                case -953:
                    if (fi35 > fi35) {
                        switch -393 {
                        case -2147483648:
                            var v37 int = (fi35 / ((fi35) | 1))
                            _ = v37
                        case -569:
                            v38 := GenS4{F0: (!(!v13)), F1: ("hello" + ("" + (("" + "\n\t") + (("" + "\n\t") + "\n\t"))))}
                            _ = v38
                        }
                        for fi39 := 0; fi39 < 12; fi39++ {
                            fmt.Println(fi39, v11, fi39)
                            switch "foo bar" {
                            case "foo bar":
                                switch "hello" {
                                case "hello":
                                    v40 := GenS3{F0: 42, F1: 2, F2: "foo bar"}
                                    _ = v40
                                }
                            case "a":
                                _ = len(v4)
                            case "":
                                var v41 float32 = 1e10
                                _ = v41
                            }
                            switch fi35 {
                            case -1:
                                v13 = (("" + ("a" + "foo bar")) >= "a")
                                var v42 int64 = -9223372036854775808
                                switch fi39 {
                                case 32768:
                                    if ((v12 + v12) < (v12 - (v12 + v12))) {
                                        v1 += (v1 & (v1 >> ((1) & 63)))
                                        v43 := GenS5{F0: int16(v12)}
                                        _ = v43
                                    } else {
                                        v3[""] = (!v11)
                                        fmt.Println(fi39, v13)
                                        if ((((("foo bar" + ("foo bar" + "a")) + "foo bar") + ("" + "hello")) + ("a" + "a")) != ("a" + ("" + (("\n\t" + "hello") + "foo bar")))) {
                                            var v44 int32 = -20
                                            _ = v44
                                        }
                                    }
                                }
                                _ = v42
                            case 128:
                                if ((("a" + "a") + "") < ("foo bar" + (("" + ((("a" + ("\n\t" + (("" + ("\n\t" + "hello")) + ("a" + "a")))) + "") + "\n\t")) + "\n\t"))) {
                                    v13 = (v13 || (202 < 676))
                                    var v45 float64 = (float64(0.0) + 1e10)
                                    _ = v45
                                }
                            case 127:
                                v12 = 127
                            default:
                                v11 = v11
                            }
                            _ = fi39
                        }
                        if (v12 < (v12 << ((1) & 63))) {
                            v11 = (1.5 > 3.14159)
                            var v46 float32 = (0.1 / 3.14159)
                            v12 = v12
                            _ = v46
                        } else {
                            var v47 uint32 = uint32(v1)
                            _ = v47
                        }
                    } else {
                        if (!(128 == 42)) {
                            fmt.Println(fi35, v13, v12)
                            var v48 uint64 = 2
                            _ = v48
                        }
                        var v49 uint8 = 2
                        if ((v49 & v49) >= 42) {
                            for fi50 := 0; fi50 < 1; fi50++ {
                                v51 := GenS6{F0: -32768, F1: (v13 || false), F2: -64}
                                _ = len(v4)
                                if v13 {
                                    var v52 uint8 = (v49 ^ 42)
                                    var v53 uint8 = (^(v52))
                                    _ = v52
                                    _ = v53
                                }
                                v12 = (v12 % (((v12 % (((v12 & v12)) | 1))) | 1))
                                v54 := GenS5{F0: 1}
                                fmt.Println(fi50)
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                _ = fi50
                                _ = v51
                                _ = v54
                            }
                        } else {
                            for fi55 := 0; fi55 < 12; fi55++ {
                                v2 = append(v2, -9223372036854775808)
                                _ = fi55
                            }
                            var v56 int = (fi35 * fi35)
                            _ = v56
                        }
                        _ = v49
                    }
                case 9223372036854775807:
                    if v11 {
                        var v57 float64 = (0.1 / ((((0.1 + 0.1) * 1e10) / float64((3.14159 / (((1e10 - ((0.1 - (0.1 / 1.5)) / 1.0)) + 1e10) * -2.25)))) * float64(3.14159)))
                        for fi58 := 0; fi58 < 5; fi58++ {
                            if v11 {
                                var v59 int32 = 0
                                _ = v59
                            }
                            _ = fi58
                        }
                        _ = v57
                    }
                }
            }
            _ = fi35
        }
        if (false || (902 < 9223372036854775807)) {
            switch "" {
            case "foo bar":
                if (991 == 42) {
                    fmt.Println(v1, v12, v13)
                    v12 = (v12 >> (((v1 / ((v1) | 1))) & 63))
                } else {
                    fmt.Println(v12)
                    var v60 float32 = 3.14159
                    _ = v60
                }
            }
            switch -687 {
            case 42:
                if v11 {
                    for fi61 := 0; fi61 < 12; fi61++ {
                        var v62 uint64 = 42
                        var v63 float64 = 1.5
                        _ = fi61
                        _ = v62
                        _ = v63
                    }
                    v64 := GenS7{F0: ("hello" + "foo bar"), F1: 281}
                    _ = v64
                }
            case 128:
                for fi65 := 0; fi65 < 3; fi65++ {
                    v12 -= (v12 << (((v1 << ((207) & 63))) & 63))
                    if v11 {
                        v1 += (^(v1))
                        switch 42 {
                        case 271:
                            var v66 string = ("a" + "\n\t")
                            _ = v66
                        case -86:
                            if (!(127 <= 127)) {
                                fmt.Println(fi65, fi65, v11)
                                fmt.Println(v11, v1, v1)
                            }
                        default:
                            var v67 int16 = -32768
                            _ = v67
                        }
                    }
                    _ = fi65
                }
            case 223:
                var v68 uint64 = 18446744073709551615
                _ = v68
            }
        } else {
            if false {
                for fi69 := 0; fi69 < 10; fi69++ {
                    for fi70 := 0; fi70 < 15; fi70++ {
                        v71 := GenS1{F0: 2147483648, F1: 127}
                        _ = fi70
                        _ = v71
                    }
                    for fi72 := 0; fi72 < 16; fi72++ {
                        switch "a" {
                        case "hello":
                            var v73 int64 = -1
                            switch fi72 {
                            case -9223372036854775808:
                                switch ("hello" + "foo bar") {
                                case "a":
                                    var v74 string = ""
                                    _ = v74
                                }
                            case 499:
                                v75 := GenS4{F0: ((4294967295 != 699) && (22 > 2)), F1: "a"}
                                _ = v75
                            case 116:
                                var v76 bool = (!(false && (v1 < (v1 << ((v1) & 63)))))
                                _ = v76
                            default:
                                var v77 uint16 = 65535
                                _ = v77
                            }
                            _ = v73
                        case "\n\t":
                            v13 = (88 != 4294967295)
                        default:
                            if (4294967295 < 888) {
                                var v78 float32 = (-2.25 + 1e10)
                                fmt.Println(v1, v78)
                                var v79 uint32 = 4294967295
                                var v80 uint16 = 1
                                _ = v78
                                _ = v79
                                _ = v80
                            } else {
                                switch fi72 {
                                case 9223372036854775807:
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                case -9223372036854775808:
                                    var v81 uint64 = 42
                                    _ = v81
                                case 0:
                                    if (v12 > -1) {
                                        if len(v4) > 0 {
                                            v4[0] = -2147483648
                                        }
                                        fmt.Println(v13)
                                    }
                                }
                                if (v13 && (!v11)) {
                                    v82 := GenS6{F0: 1000, F1: (42 >= int16(v1)), F2: 32767}
                                    fmt.Println(fi72, fi72)
                                    v1 += v1
                                    _ = v82
                                }
                            }
                        }
                        _ = fi72
                    }
                    _ = fi69
                }
                for fi83 := 0; fi83 < 15; fi83++ {
                    fmt.Println(v13, v11, v1)
                    fmt.Println(v13)
                    for fi84 := 0; fi84 < 17; fi84++ {
                        var v85 int = (fi84 << (((v1 << ((v1) & 63))) & 63))
                        fmt.Println(v11, v12)
                        _ = fi84
                        _ = v85
                    }
                    _ = fi83
                }
                delete(v3, (("foo bar" + "a") + "a"))
            } else {
                fmt.Println(v12, v1)
                fmt.Println(v13)
            }
        }
    }
    if (v12 <= (v12 | (^(v12)))) {
        fmt.Println(v1, v13, v12)
        var v86 uint32 = 0
        _ = v86
    } else {
        var v87 float64 = 0.1
        if ((-1 > -955) && (-166 == -32768)) {
            switch 1 {
            case -631:
                v87 = v87
            case 42:
                if (v11 || (1 >= 1)) {
                    var v88 float32 = 3.14159
                    _ = v88
                }
            case -1:
                if false {
                    switch "" {
                    case "hello":
                        for fi89 := 0; fi89 < 19; fi89++ {
                            v4 = append(v4, fi89)
                            for fi90 := 0; fi90 < 1; fi90++ {
                                v1 = v1
                                _ = fi90
                            }
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            _ = fi89
                        }
                    case "a":
                        for fi91 := 0; fi91 < 18; fi91++ {
                            var v92 int32 = 2147483647
                            v11 = false
                            _ = fi91
                            _ = v92
                        }
                    default:
                        fmt.Println(v1, v12)
                    }
                } else {
                    var v93 int = 716
                    var v94 uint16 = 1
                    _ = v93
                    _ = v94
                }
            }
        } else {
            fmt.Println(v11, v12)
            fmt.Println(v11)
        }
        var v95 float32 = 3.14159
        _ = v87
        _ = v95
    }
    v1 -= v1
    v3["a"] = true
    fmt.Println(v1)
    var v96 uint16 = 981
    fmt.Println(v12, v13)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v96:", v96)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v96
}
