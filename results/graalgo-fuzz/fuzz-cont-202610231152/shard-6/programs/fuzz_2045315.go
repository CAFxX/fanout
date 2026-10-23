package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32; F2 string }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int16; F1 int }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int64; F1 int64; F2 int16 }

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[int]string)
    var v3 uint8 = 255
    var v4 uint32 = 1
    if true {
        switch 32767 {
        case 2147483647:
            var v5 uint64 = 4294967295
            _ = v5
        case -2147483648:
            for fi6 := 0; fi6 < 15; fi6++ {
                if true {
                    if false {
                        fmt.Println(fi6)
                    } else {
                        fmt.Println(v3)
                        v4 += (v4 >> ((338) & 63))
                    }
                    for fi7 := 0; fi7 < 2; fi7++ {
                        v4 = v4
                        _ = fi7
                    }
                    v3 = (v3 - (v3 & v3))
                }
                var v8 float32 = 1.5
                var v9 int = (-(fi6))
                _ = fi6
                _ = v8
                _ = v9
            }
        }
    } else {
        var v10 int8 = -128
        if (1.5 > (-2.25 - 1e10)) {
            v4 = v4
            var v11 uint32 = v4
            _ = v11
        }
        if (!("" >= "a")) {
            v12 := GenS1{F0: (v4 % ((86) | 1)), F1: (v4 & (v4 + (v4 << ((65535) & 63)))), F2: "hello"}
            for fi13 := 0; fi13 < 1; fi13++ {
                switch ("a" + "a") {
                case "a":
                    v4 -= (v4 * (v4 ^ (v4 + (v4 % ((1) | 1)))))
                case "foo bar":
                    if false {
                        _ = len(v2)
                        if true {
                            v4 = 822
                        } else {
                            var v14 uint16 = uint16(v3)
                            v15 := GenS1{F0: v4, F1: v4, F2: "foo bar"}
                            var v16 int = fi13
                            _ = v14
                            _ = v15
                            _ = v16
                        }
                    }
                case "":
                    fmt.Println(v4, v10, v10)
                }
                _ = fi13
            }
            _ = v12
        } else {
            v17 := GenS2{F0: "foo bar"}
            var v18 int64 = 42
            _ = v17
            _ = v18
        }
        _ = v10
    }
    v19 := []int{32767, 128, 4294967295}
    v3 = v3
    v20 := GenS1{F0: 246, F1: v4, F2: ("foo bar" + ("\n\t" + ("a" + "hello")))}
    v21 := GenS1{F0: v4, F1: v4, F2: "foo bar"}
    var v22 int8 = -1
    var v23 float32 = 0.1
    v3 = v3
    if (42 == -1) {
        if (255 != 873) {
            fmt.Println(v23, v3)
            if (-599 <= 0) {
                fmt.Println(v3, v23)
                fmt.Println(v22)
                if (v22 > v22) {
                    delete(v1, v4)
                } else {
                    switch ("\n\t" + ("hello" + ("" + ""))) {
                    case "foo bar":
                        fmt.Println(v22)
                    case "":
                        v24 := v1[v4]
                        _ = v24
                    default:
                        if (-1 != 871) {
                            for fi25 := 0; fi25 < 7; fi25++ {
                                var v26 int = (fi25 ^ fi25)
                                _ = fi25
                                _ = v26
                            }
                        } else {
                            var v27 int16 = -32768
                            _ = v27
                        }
                    }
                    v28 := GenS1{F0: (v4 - 113), F1: v4, F2: (("hello" + ("a" + ("\n\t" + "foo bar"))) + "\n\t")}
                    _ = v28
                }
            }
        } else {
            var v29 float32 = float32(3.14159)
            _ = v29
        }
        switch "\n\t" {
        case "":
            for fi30 := 0; fi30 < 6; fi30++ {
                v31 := v2[fi30]
                _ = fi30
                _ = v31
            }
        default:
            switch ("a" + "") {
            case "\n\t":
                if (!(1 == -9223372036854775808)) {
                    for fi32 := 0; fi32 < 18; fi32++ {
                        var v33 uint32 = 4294967295
                        v22 -= v22
                        var v34 int16 = 682
                        _ = fi32
                        _ = v33
                        _ = v34
                    }
                } else {
                    for fi35 := 0; fi35 < 20; fi35++ {
                        for fi36 := 0; fi36 < 19; fi36++ {
                            switch (fi35 % ((-646) | 1)) {
                            case 128:
                                switch 128 {
                                case 1:
                                    v4 -= (^(v4))
                                    v37 := GenS1{F0: v4, F1: (v4 - (v4 - v4)), F2: ((("" + "hello") + "a") + "a")}
                                    _ = v37
                                case 32767:
                                    var v38 int = fi36
                                    _ = v38
                                case 32768:
                                    _ = len(v19)
                                default:
                                    fmt.Println(v4, v4)
                                }
                            case -1:
                                v3 = (v3 - v3)
                            case -2147483648:
                                switch -640 {
                                case -2147483648:
                                    v4 -= 2
                                default:
                                    _ = len(v2)
                                }
                            default:
                                var v39 int64 = 64
                                _ = v39
                            }
                            fmt.Println(v3)
                            v3 = (^(v3))
                            _ = fi36
                        }
                        switch "foo bar" {
                        case "hello":
                            switch ("\n\t" + "foo bar") {
                            case "a":
                                v23 -= v23
                            case "foo bar":
                                v22 -= (^(v22))
                            default:
                                fmt.Println(v3)
                            }
                        default:
                            for fi40 := 0; fi40 < 14; fi40++ {
                                if false {
                                    switch fi35 {
                                    case 256:
                                        var v41 int64 = int64(v3)
                                        _ = v41
                                    case -2147483648:
                                        delete(v1, v4)
                                    case 809:
                                        var v42 uint8 = v3
                                        _ = v42
                                    }
                                    if (0 < 34) {
                                        var v43 int32 = 0
                                        _ = v43
                                    } else {
                                        switch (("a" + "hello") + "\n\t") {
                                        case "":
                                            v44 := GenS1{F0: 2147483648, F1: v4, F2: "a"}
                                            _ = v44
                                        case "foo bar":
                                            if (true || (!false)) {
                                                if (false && true) {
                                                    v23 -= v23
                                                    switch "\n\t" {
                                                    case "\n\t":
                                                        var v45 string = ""
                                                        _ = v45
                                                    case "foo bar":
                                                        switch ("\n\t" + "a") {
                                                        case "foo bar":
                                                            v46 := GenS2{F0: "hello"}
                                                            _ = v46
                                                        case "\n\t":
                                                            v22 += v22
                                                        }
                                                    default:
                                                        v3 = v3
                                                        if len(v19) >= 2 {
                                                            v19 = v19[1:len(v19)]
                                                        }
                                                        v47 := GenS1{F0: v4, F1: v4, F2: (("\n\t" + "foo bar") + "foo bar")}
                                                        _ = v47
                                                    }
                                                    fmt.Println(v22, v22)
                                                } else {
                                                    v3 = v3
                                                    fmt.Println(fi40, v22)
                                                    if false {
                                                        var v48 float32 = (v23 * v23)
                                                        fmt.Println(v22)
                                                        _ = v48
                                                    }
                                                }
                                                if ("" == "\n\t") {
                                                    var v49 uint32 = 2147483648
                                                    _ = len(v2)
                                                    var v50 uint64 = 2
                                                    _ = v49
                                                    _ = v50
                                                }
                                                if (((!("foo bar" > ("hello" + "hello"))) || true) && (-1 != 9223372036854775807)) {
                                                    if (65535 < 42) {
                                                        v22 -= v22
                                                        fmt.Println(v23, v3, v23)
                                                        var v51 int8 = 100
                                                        _ = v51
                                                    }
                                                    if (!((!((!((-2.25 * 0.1) != (float64(v23) / ((1.5 + 1e10) - float64((((float32(3.14159) * 0.0) + (0.1 + v23)) * 0.1)))))) || (1 <= 392))) || (0 >= -218))) {
                                                        var v52 int16 = 42
                                                        delete(v1, v4)
                                                        v19 = append(v19, -9223372036854775808)
                                                        if false {
                                                            var v53 int8 = 103
                                                            _ = v53
                                                        }
                                                        _ = v52
                                                    }
                                                } else {
                                                    if true {
                                                        v22 = v22
                                                        var v54 uint = 0
                                                        switch fi40 {
                                                        case 127:
                                                            fmt.Println(v22, v23, fi40)
                                                        }
                                                        _ = v54
                                                    }
                                                }
                                            }
                                        }
                                        if ((v4 >> ((1) & 63)) < 901) {
                                            v22 = int8(v4)
                                        } else {
                                            v55 := GenS3{F0: 42, F1: 840}
                                            _ = len(v2)
                                            var v56 int64 = -9223372036854775808
                                            if (!false) {
                                                v57 := GenS1{F0: 0, F1: v4, F2: ("\n\t" + "foo bar")}
                                                v23 = v23
                                                _ = len(v19)
                                                if false {
                                                    var v58 uint8 = v3
                                                    v4 = 1
                                                    v59 := v1[v4]
                                                    v4 += v4
                                                    v19 = append(v19, 4294967295)
                                                    v22 += v22
                                                    var v60 uint16 = 432
                                                    v2[(^(fi35))] = v59
                                                    _ = v58
                                                    _ = v59
                                                    _ = v60
                                                }
                                                _ = v57
                                            }
                                            _ = v55
                                            _ = v56
                                        }
                                        v3 += (v3 & (^(v3)))
                                        v4 += 1
                                    }
                                    var v61 int = fi40
                                    v3 -= v3
                                    v3 += v3
                                    if (v23 == (1.5 * -2.25)) {
                                        var v62 uint32 = (v4 + v4)
                                        if true {
                                            v63 := GenS3{F0: 32767, F1: fi35}
                                            switch (fi35 / ((1) | 1)) {
                                            case 128:
                                                v64 := v2[int(v3)]
                                                _ = v64
                                            case 668:
                                                switch (fi40 ^ (fi35 << ((uint(v3)) & 63))) {
                                                case 9223372036854775807:
                                                    v22 = (v22 ^ 107)
                                                case 32767:
                                                    if ("\n\t" > (("hello" + ("hello" + "hello")) + "a")) {
                                                        var v65 int32 = 2147483647
                                                        fmt.Println(fi40)
                                                        _ = v65
                                                    } else {
                                                        if len(v19) > 0 {
                                                            v19[0] = int(v22)
                                                        }
                                                    }
                                                case -2147483648:
                                                    v22 = (v22 >> ((195) & 63))
                                                }
                                            case 796:
                                                switch "\n\t" {
                                                case "hello":
                                                    var v66 uint16 = 65535
                                                    _ = v66
                                                default:
                                                    fmt.Println(v22)
                                                    v23 -= 1e10
                                                    var v67 uint32 = (v4 % (((v62 - v4)) | 1))
                                                    _ = v67
                                                }
                                            }
                                            _ = v63
                                        } else {
                                            v23 += v23
                                            v68 := GenS4{F0: (false && false)}
                                            _ = v68
                                        }
                                        _ = v62
                                    } else {
                                        switch fi40 {
                                        case 128:
                                            if true {
                                                fmt.Println(v61, v22)
                                            }
                                        }
                                        var v69 string = ("hello" + "")
                                        _ = v69
                                    }
                                    _ = v61
                                } else {
                                    var v70 uint64 = 145
                                    _ = v70
                                }
                                switch (fi40 % ((1) | 1)) {
                                case -410:
                                    v2[(^(fi35))] = ""
                                default:
                                    if (1 <= (v3 & 61)) {
                                        v23 += v23
                                        var v71 uint16 = 65535
                                        _ = v71
                                    } else {
                                        v23 += float32((1.5 - 0.0))
                                        var v72 bool = (42 >= 709)
                                        delete(v1, 2147483647)
                                        _ = v72
                                    }
                                }
                                _ = fi40
                            }
                        }
                        _ = fi35
                    }
                    if ((-21 >= 288) && true) {
                        var v73 int8 = -61
                        v3 = 246
                        v74 := GenS5{F0: 288, F1: 42, F2: -483}
                        _ = v73
                        _ = v74
                    }
                }
            case "":
                v4 += 285
            }
        }
        v23 = 0.0
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
}
