package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint32; F1 int32; F2 int }
type GenS3 struct { F0 int; F1 int16 }
type GenS4 struct { F0 uint8; F1 uint; F2 bool }
type GenS5 struct { F0 int16; F1 string }
type GenS6 struct { F0 int16; F1 int8 }
type GenS7 struct { F0 int8; F1 uint32; F2 uint8 }
type GenS8 struct { F0 uint; F1 int32 }
type GenS9 struct { F0 int16; F1 uint32; F2 int16 }

func main() {
    v1 := make(map[int]string)
    var v2 uint32 = 0
    v3 := make(map[string]int)
    var v4 bool = true
    for fi5 := 0; fi5 < 17; fi5++ {
        if v4 {
            var v6 int16 = 42
            _ = v6
        }
        if false {
            if true {
                var v7 int = fi5
                var v8 float32 = (3.14159 / (0.1 - (0.0 / 1.5)))
                _ = v7
                _ = v8
            } else {
                if (32767 > 65535) {
                    for fi9 := 0; fi9 < 12; fi9++ {
                        v4 = v4
                        _ = fi9
                    }
                    switch "hello" {
                    case "a":
                        v4 = (958 >= 2)
                    case "":
                        v3[("foo bar" + "\n\t")] = fi5
                    default:
                        v3[(("hello" + "hello") + "\n\t")] = (^(fi5))
                    }
                    v10 := GenS1{F0: fi5}
                    _ = v10
                } else {
                    var v11 int8 = 47
                    _ = v11
                }
            }
        } else {
            for fi12 := 0; fi12 < 6; fi12++ {
                v13 := GenS1{F0: (fi12 % ((int(v2)) | 1))}
                v2 += (v2 - (v2 & v2))
                var v14 uint16 = 2
                _ = fi12
                _ = v13
                _ = v14
            }
            v2 = v2
        }
        _ = fi5
    }
    for fi15 := 0; fi15 < 1; fi15++ {
        if (((!v4) || (3.14159 <= 0.1)) || (4294967295 == 427)) {
            var v16 int32 = 2147483647
            for fi17 := 0; fi17 < 9; fi17++ {
                v18 := v3["foo bar"]
                switch ("a" + "a") {
                case "\n\t":
                    v19 := GenS2{F0: (v2 * v2), F1: (^(v16)), F2: v18}
                    _ = v19
                case "foo bar":
                    fmt.Println(v18, v16)
                }
                _ = fi17
                _ = v18
            }
            v4 = (("\n\t" + (("foo bar" + ((("foo bar" + "hello") + "\n\t") + (("" + "foo bar") + ("a" + "a")))) + "foo bar")) < "")
            _ = v16
        } else {
            var v20 uint = 65535
            v21 := []int{-222}
            for fi22 := 0; fi22 < 19; fi22++ {
                var v23 float32 = 1.5
                v21 = append(v21, (-(fi15)))
                _ = fi22
                _ = v23
            }
            _ = v20
            _ = v21
        }
        for fi24 := 0; fi24 < 5; fi24++ {
            if (("hello" != ((("hello" + ("" + (("\n\t" + ("foo bar" + ("foo bar" + "a"))) + "\n\t"))) + (((("" + "foo bar") + "hello") + "\n\t") + "a")) + "hello")) || false) {
                v2 = v2
                v2 = 0
                var v25 int64 = -1
                _ = v25
            }
            _ = fi24
        }
        switch fi15 {
        case -310:
            if v4 {
                _ = len(v3)
            }
        case 4294967295:
            if (v4 && (!("hello" == "a"))) {
                v26 := []int{fi15, -9223372036854775808}
                for fi27 := 0; fi27 < 11; fi27++ {
                    var v28 uint = 255
                    if len(v26) > 0 {
                        v26[0] = (fi15 * (fi15 | (fi15 & (fi15 % (((^(fi27))) | 1)))))
                    }
                    if true {
                        if true {
                            if v4 {
                                var v29 bool = (48 != 119)
                                fmt.Println(v28, v4, v28)
                                if v29 {
                                    v4 = v4
                                    v30 := GenS2{F0: 65535, F1: -1, F2: fi15}
                                    var v31 uint16 = 163
                                    _ = v30
                                    _ = v31
                                }
                                var v32 string = ("\n\t" + "hello")
                                if (v32 > "") {
                                    var v33 uint32 = (v2 >> (((v28 * v28)) & 63))
                                    _ = v33
                                }
                                _ = v29
                                _ = v32
                            } else {
                                if len(v26) >= 2 {
                                    v26 = v26[1:len(v26)]
                                }
                                if ((fi27 * 127) < 9223372036854775807) {
                                    switch ("\n\t" + ((("hello" + "") + "hello") + ("hello" + "\n\t"))) {
                                    case "":
                                        fmt.Println(fi27, v2)
                                    case "\n\t":
                                        if v4 {
                                            v4 = false
                                        }
                                    default:
                                        v4 = v4
                                    }
                                    switch (fi27 << (((v28 / (((^(v28))) | 1))) & 63)) {
                                    case 252:
                                        var v34 uint64 = 1
                                        _ = v34
                                    default:
                                        var v35 int = (fi15 * fi15)
                                        _ = v35
                                    }
                                }
                            }
                            if v4 {
                                if len(v26) > 0 {
                                    v26[0] = fi15
                                }
                                switch fi15 {
                                case 42:
                                    switch "" {
                                    case "\n\t":
                                        v2 = 812
                                    case "hello":
                                        fmt.Println(v4, v28, v28)
                                        fmt.Println(v28)
                                        v36 := v3["\n\t"]
                                        _ = v36
                                    default:
                                        var v37 int32 = -2147483648
                                        _ = v37
                                    }
                                default:
                                    var v38 bool = false
                                    _ = v38
                                }
                                if v4 {
                                    switch (fi15 - fi15) {
                                    case -9223372036854775808:
                                        if (1 <= 42) {
                                            if v4 {
                                                var v39 uint64 = 18446744073709551615
                                                var v40 float64 = (0.1 - -2.25)
                                                v41 := GenS3{F0: fi27, F1: -1}
                                                _ = v39
                                                _ = v40
                                                _ = v41
                                            }
                                            var v42 float64 = 1e10
                                            _ = v42
                                        }
                                    case 734:
                                        if len(v26) > 0 {
                                            v26[0] = 4294967295
                                        }
                                    default:
                                        delete(v3, "foo bar")
                                    }
                                    if (float32(((1e10 + 0.0) / ((0.1 - (float64(-2.25) - -2.25)) / 1.0))) >= 1e10) {
                                        var v43 int64 = -478
                                        v4 = (3.14159 >= 0.1)
                                        if v4 {
                                            v44 := GenS2{F0: v2, F1: 511, F2: 128}
                                            _ = v44
                                        } else {
                                            var v45 int = fi27
                                            switch "a" {
                                            case "\n\t":
                                                v46 := GenS1{F0: (fi15 % ((v45) | 1))}
                                                _ = v46
                                            case "a":
                                                v43 = -9223372036854775808
                                            }
                                            _ = v45
                                        }
                                        _ = v43
                                    }
                                } else {
                                    var v47 string = "\n\t"
                                    _ = v47
                                }
                                v48 := GenS2{F0: v2, F1: 1, F2: (fi15 << ((v28) & 63))}
                                _ = v48
                            } else {
                                var v49 uint16 = uint16(v2)
                                var v50 uint16 = v49
                                v4 = ((!true) || v4)
                                _ = v49
                                _ = v50
                            }
                        } else {
                            if ("" <= "hello") {
                                v4 = true
                            }
                        }
                        v51 := v3[("\n\t" + "hello")]
                        if len(v26) >= 2 {
                            v26 = v26[1:len(v26)]
                        }
                        _ = v51
                    } else {
                        var v52 uint = v28
                        var v53 float32 = -2.25
                        if false {
                            switch "hello" {
                            case "foo bar":
                                var v54 uint8 = 255
                                _ = v54
                            case "":
                                v53 = (v53 - float32((3.14159 - float64(1e10))))
                            default:
                                v26 = append(v26, int(v2))
                            }
                            if len(v26) >= 2 {
                                v26 = v26[1:len(v26)]
                            }
                        }
                        _ = v52
                        _ = v53
                    }
                    switch ("hello" + ("hello" + ("hello" + ("\n\t" + "hello")))) {
                    case "":
                        var v55 int16 = -426
                        _ = v55
                    case "hello":
                        v56 := GenS4{F0: uint8(v28), F1: (v28 ^ v28), F2: ((0 >= 1) && (false && (((((0.1 + -2.25) * 1e10) * 0.0) == (1.5 * 1.5)) || (0.0 == -2.25))))}
                        _ = v56
                    default:
                        switch (fi27 % ((fi27) | 1)) {
                        case 469:
                            if ((1 == 32767) || (v4 || true)) {
                                var v57 int16 = 0
                                fmt.Println(fi27)
                                _ = v57
                            }
                            var v58 uint8 = 128
                            _ = v58
                        default:
                            v59 := GenS3{F0: (fi27 % (((^(fi27))) | 1)), F1: 472}
                            _ = v59
                        }
                    }
                    var v60 int = fi15
                    v61 := GenS5{F0: 42, F1: (((("foo bar" + "") + "") + "hello") + (("a" + ("a" + "foo bar")) + "foo bar"))}
                    var v62 uint32 = v2
                    _ = fi27
                    _ = v28
                    _ = v60
                    _ = v61
                    _ = v62
                }
                _ = v26
            } else {
                v1[(fi15 << ((42) & 63))] = ""
                delete(v1, fi15)
                fmt.Println(v4, fi15)
            }
        case 1:
            v63 := GenS6{F0: 25, F1: 43}
            _ = v63
        }
        _ = fi15
    }
    fmt.Println(v2, v4)
    var v64 float64 = 1.5
    var v65 float32 = ((0.0 * 0.0) / 1.5)
    v66 := GenS6{F0: -1, F1: 0}
    switch 9223372036854775807 {
    case 42:
        v67 := GenS7{F0: 83, F1: (v2 >> ((2) & 63)), F2: 42}
        _ = v67
    case 2147483647:
        if v4 {
            var v68 float32 = (0.1 + 0.1)
            switch ("\n\t" + (((("hello" + "\n\t") + ((("hello" + ("hello" + "foo bar")) + "hello") + ("" + ("" + "\n\t")))) + ((("a" + "hello") + "") + "hello")) + "hello")) {
            case "\n\t":
                v69 := GenS8{F0: 2, F1: -1}
                _ = v69
            case "":
                v65 = (v68 / 3.14159)
            case "hello":
                fmt.Println(v2)
            default:
                var v70 uint16 = 42
                _ = v70
            }
            switch 615 {
            case 306:
                switch ("a" + "\n\t") {
                case "a":
                    _ = len(v1)
                case "\n\t":
                    v71 := []int{128}
                    _ = v71
                case "hello":
                    v72 := []int{32767, 42, 4294967295, -2147483648}
                    _ = v72
                default:
                    v68 = v65
                }
            case -9223372036854775808:
                for fi73 := 0; fi73 < 11; fi73++ {
                    var v74 int64 = 9223372036854775807
                    var v75 uint64 = 0
                    _ = fi73
                    _ = v74
                    _ = v75
                }
            default:
                if v4 {
                    var v76 uint16 = 32767
                    v64 = (v64 / v64)
                    if v4 {
                        v2 = v2
                        for fi77 := 0; fi77 < 19; fi77++ {
                            v78 := []int{fi77}
                            _ = fi77
                            _ = v78
                        }
                        v79 := GenS9{F0: 727, F1: v2, F2: 1}
                        _ = v79
                    }
                    _ = v76
                } else {
                    for fi80 := 0; fi80 < 17; fi80++ {
                        var v81 string = ("\n\t" + "")
                        v65 = v65
                        _ = fi80
                        _ = v81
                    }
                }
            }
            _ = v68
        }
    }
    for fi82 := 0; fi82 < 19; fi82++ {
        v83 := GenS5{F0: 32767, F1: (("" + "a") + "")}
        for fi84 := 0; fi84 < 5; fi84++ {
            switch fi82 {
            case 127:
                if (1 > 4294967295) {
                    switch fi84 {
                    case 32767:
                        if (!true) {
                            var v85 float32 = v65
                            var v86 int16 = 1000
                            _ = v85
                            _ = v86
                        }
                        v2 = (^(v2))
                        var v87 uint = 965
                        _ = v87
                    case 128:
                        if ((!((v2 << ((18446744073709551615) & 63)) < 2)) || v4) {
                            var v88 uint = 42
                            switch -48 {
                            case 534:
                                v89 := []int{fi84}
                                _ = v89
                            default:
                                var v90 uint = v88
                                _ = v90
                            }
                            var v91 int = (fi84 * 1)
                            _ = v88
                            _ = v91
                        } else {
                            v92 := []int{(-(fi82))}
                            _ = len(v1)
                            _ = v92
                        }
                    case 32768:
                        switch "" {
                        case "hello":
                            v2 -= v2
                        case "a":
                            v93 := GenS9{F0: -1, F1: v2, F2: -428}
                            _ = v93
                        case "foo bar":
                            v94 := GenS7{F0: 127, F1: (^(v2)), F2: 42}
                            _ = v94
                        default:
                            v95 := GenS1{F0: -976}
                            _ = v95
                        }
                        v96 := []int{42, (fi82 | (fi84 >> ((18446744073709551615) & 63)))}
                        v2 = (v2 >> ((2) & 63))
                        _ = v96
                    }
                    v65 -= -2.25
                } else {
                    v2 = 0
                    fmt.Println(v65)
                }
            default:
                v97 := GenS3{F0: (-(fi84)), F1: -1}
                _ = v97
            }
            var v98 bool = v4
            _ = fi84
            _ = v98
        }
        v2 = (v2 * 383)
        _ = fi82
        _ = v83
    }
    for fi99 := 0; fi99 < 9; fi99++ {
        fmt.Println(v65, v65, v64)
        _ = fi99
    }
    fmt.Println(v2, v64)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v64:", v64)
    fmt.Println("v65:", v65)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v64
    _ = v65
    _ = v66
}
