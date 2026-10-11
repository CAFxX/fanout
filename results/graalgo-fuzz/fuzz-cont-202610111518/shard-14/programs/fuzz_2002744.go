package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8; F2 uint8 }
type GenS2 struct { F0 uint8; F1 int }
type GenS3 struct { F0 uint64; F1 uint; F2 int16 }
type GenS4 struct { F0 uint64; F1 uint }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 uint16; F1 int64 }
type GenS7 struct { F0 string }
func genFn0(p0 uint32) uint16 {
    p0 -= p0
    for fi5 := 0; fi5 < 17; fi5++ {
        var v6 uint32 = (p0 & (p0 & p0))
        v7 := []int{int(v6), (fi5 * 128), 9223372036854775807, 32768}
        for fi8 := 0; fi8 < 1; fi8++ {
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
            v6 = (v6 ^ 4294967295)
            _ = fi8
        }
        _ = fi5
        _ = v6
        _ = v7
    }
    switch ("hello" + "hello") {
    case "a":
        for fi9 := 0; fi9 < 11; fi9++ {
            for fi10 := 0; fi10 < 15; fi10++ {
                p0 = 2147483647
                _ = fi10
            }
            v11 := []int{-104, 9223372036854775807, fi9}
            p0 += 406
            _ = fi9
            _ = v11
        }
    case "\n\t":
        if (!(true || (!true))) {
            var v12 uint = 4294967295
            for fi13 := 0; fi13 < 6; fi13++ {
                v12 += v12
                for fi14 := 0; fi14 < 5; fi14++ {
                    v15 := GenS1{F0: (fi13 - 1), F1: 108, F2: 2}
                    _ = fi14
                    _ = v15
                }
                v16 := GenS1{F0: fi13, F1: 42, F2: 0}
                _ = fi13
                _ = v16
            }
            var v17 int32 = 874
            _ = v12
            _ = v17
        }
    case "foo bar":
        for fi18 := 0; fi18 < 18; fi18++ {
            p0 += (p0 - p0)
            var v19 string = ("foo bar" + "a")
            _ = fi18
            _ = v19
        }
    default:
        v20 := GenS2{F0: 154, F1: -1}
        _ = v20
    }
    p0 += (^(p0))
    if true {
        if false {
            p0 = 36
            switch 32768 {
            case 32767:
                if (750 == 1) {
                    if false {
                        p0 = (p0 + p0)
                    }
                } else {
                    var v21 float32 = 0.1
                    v21 -= v21
                    v22 := GenS2{F0: 255, F1: 201}
                    _ = v21
                    _ = v22
                }
            case -45:
                p0 = p0
            case -3:
                p0 += p0
            }
            for fi23 := 0; fi23 < 14; fi23++ {
                var v24 int32 = -1
                switch (fi23 - fi23) {
                case -9223372036854775808:
                    var v25 float32 = -2.25
                    _ = v25
                case -2147483648:
                    fmt.Println(p0, fi23, p0)
                default:
                    if (false && false) {
                        if (false || true) {
                            if false {
                                var v26 uint8 = 255
                                _ = v26
                            }
                            v27 := GenS3{F0: 4294967295, F1: 18446744073709551615, F2: -758}
                            var v28 uint32 = p0
                            _ = v27
                            _ = v28
                        }
                        v24 = v24
                        v29 := []int{fi23}
                        _ = v29
                    } else {
                        fmt.Println(p0, p0)
                        fmt.Println(fi23, v24)
                        for fi30 := 0; fi30 < 16; fi30++ {
                            if (0 < 42) {
                                v31 := []int{-2147483648, fi23, fi23, fi23}
                                v32 := GenS1{F0: 1, F1: -1, F2: 127}
                                var v33 uint64 = 18446744073709551615
                                _ = v31
                                _ = v32
                                _ = v33
                            }
                            var v34 string = (("\n\t" + "") + "a")
                            v35 := []int{fi30, fi30}
                            if (1 < 255) {
                                fmt.Println(fi23, fi30)
                                _ = len(v35)
                                if (((0 > 855) || false) && (true && false)) {
                                    v34 = "foo bar"
                                    v24 = v24
                                    if len(v35) > 0 {
                                        v35[0] = 1
                                    }
                                    switch "foo bar" {
                                    case "foo bar":
                                        fmt.Println(v34)
                                        var v36 uint8 = 1
                                        v37 := GenS1{F0: (fi23 ^ 894), F1: 127, F2: v36}
                                        v38 := GenS4{F0: uint64(v36), F1: 255}
                                        var v39 int64 = -9223372036854775808
                                        _ = v36
                                        _ = v37
                                        _ = v38
                                        _ = v39
                                    case "a":
                                        fmt.Println(fi30, fi30, v34)
                                    }
                                } else {
                                    p0 = p0
                                }
                                fmt.Println(fi23, fi30, fi30)
                            }
                            switch int(p0) {
                            case 12:
                                if ((1 != 605) || (-15 != -2)) {
                                    if ((!((1e10 + 3.14159) >= 3.14159)) || (!true)) {
                                        v35 = append(v35, (fi30 | fi30))
                                        switch (v34 + v34) {
                                        case "foo bar":
                                            if (42 == 1) {
                                                p0 -= 2147483647
                                            }
                                        case "hello":
                                            switch 7 {
                                            case 96:
                                                if true {
                                                    if true {
                                                        p0 += (p0 - (p0 - 714))
                                                        if true {
                                                            _ = len(v35)
                                                            v40 := GenS2{F0: 80, F1: -2147483648}
                                                            _ = v40
                                                        }
                                                        var v41 uint64 = 331
                                                        _ = v41
                                                    } else {
                                                        var v42 uint = 4294967295
                                                        var v43 int = 9223372036854775807
                                                        if len(v35) > 0 {
                                                            v35[0] = fi23
                                                        }
                                                        _ = v42
                                                        _ = v43
                                                    }
                                                    v44 := GenS2{F0: 1, F1: fi23}
                                                    _ = v44
                                                }
                                            case 127:
                                                var v45 int64 = -1
                                                _ = v45
                                            case 4294967295:
                                                v34 = v34
                                            default:
                                                p0 += (p0 + (^(p0)))
                                            }
                                        case "\n\t":
                                            var v46 uint32 = p0
                                            _ = v46
                                        default:
                                            var v47 uint64 = 42
                                            _ = v47
                                        }
                                        v48 := GenS1{F0: 235, F1: 100, F2: 139}
                                        switch "" {
                                        case "hello":
                                            p0 -= p0
                                        }
                                        switch (fi23 ^ (fi23 % ((fi30) | 1))) {
                                        case -9223372036854775808:
                                            var v49 uint16 = 0
                                            var v50 uint64 = 4294967295
                                            p0 += (p0 - 1)
                                            _ = v49
                                            _ = v50
                                        }
                                        _ = v48
                                    }
                                    _ = len(v35)
                                    var v51 int32 = (v24 / (((v24 & (v24 << ((4294967295) & 63)))) | 1))
                                    fmt.Println(v34, p0)
                                    if (-2.25 >= -2.25) {
                                        if true {
                                            if true {
                                                if len(v35) > 0 {
                                                    v35[0] = 32768
                                                }
                                                var v52 int32 = v24
                                                v53 := GenS5{F0: -128}
                                                _ = v52
                                                _ = v53
                                            }
                                        } else {
                                            switch v34 {
                                            case "":
                                                if (446 == 2) {
                                                    var v54 float32 = ((-2.25 - 0.0) + ((3.14159 / 1.5) * (0.0 + 0.0)))
                                                    if (0 <= -1) {
                                                        var v55 string = v34
                                                        _ = v55
                                                    }
                                                    _ = v54
                                                } else {
                                                    fmt.Println(v34)
                                                    switch ("" + "hello") {
                                                    case "hello":
                                                        fmt.Println(fi23)
                                                    default:
                                                        v56 := GenS4{F0: 4294967295, F1: 104}
                                                        _ = v56
                                                    }
                                                    var v57 int = -308
                                                    _ = v57
                                                }
                                            }
                                        }
                                        fmt.Println(v24, fi30, p0)
                                        var v58 float32 = 0.0
                                        if len(v35) >= 2 {
                                            v35 = v35[1:len(v35)]
                                        }
                                        if false {
                                            var v59 int = 9223372036854775807
                                            if true {
                                                var v60 int = fi23
                                                var v61 uint = 954
                                                _ = v60
                                                _ = v61
                                            }
                                            v62 := GenS6{F0: 497, F1: -1}
                                            _ = v59
                                            _ = v62
                                        } else {
                                            if (("a" + ("a" + v34)) < v34) {
                                                v63 := GenS1{F0: (fi23 / ((9223372036854775807) | 1)), F1: 0, F2: 42}
                                                if false {
                                                    var v64 int64 = -879
                                                    _ = len(v35)
                                                    _ = v64
                                                }
                                                _ = v63
                                            }
                                            v34 = ("a" + v34)
                                            v65 := GenS4{F0: 0, F1: 458}
                                            _ = v65
                                        }
                                        _ = v58
                                    } else {
                                        var v66 int64 = -1
                                        var v67 bool = true
                                        _ = v66
                                        _ = v67
                                    }
                                    v35 = append(v35, fi23)
                                    _ = v51
                                } else {
                                    if (!true) {
                                        v24 = v24
                                    }
                                    var v68 uint8 = 8
                                    _ = v68
                                }
                            }
                            var v69 uint = 997
                            _ = fi30
                            _ = v34
                            _ = v35
                            _ = v69
                        }
                    }
                }
                v24 += v24
                _ = fi23
                _ = v24
            }
        } else {
            if (-1 > -32768) {
                p0 = p0
            }
            fmt.Println(p0)
        }
        if ((false || ((p0 | p0) != p0)) && ((2 != uint8(p0)) || true)) {
            var v70 string = "foo bar"
            _ = v70
        } else {
            v71 := GenS6{F0: 65535, F1: -9223372036854775808}
            for fi72 := 0; fi72 < 12; fi72++ {
                p0 = p0
                _ = fi72
            }
            _ = v71
        }
        fmt.Println(p0)
    }
    _ = p0
    return 721
}

func main() {
    v1 := []int{226, 983, 2147483647, -2147483648}
    v2 := []string{"foo bar", ("foo bar" + ("foo bar" + "a")), ("hello" + "a"), ""}
    var v3 int8 = 127
    var v4 uint32 = 0
    v2 = append(v2, ("hello" + ((("a" + "") + "a") + "foo bar")))
    if true {
        fmt.Println(v4, v3)
    }
    for fi73 := 0; fi73 < 4; fi73++ {
        switch "\n\t" {
        case "hello":
            var v74 string = "foo bar"
            _ = v74
        default:
            v3 = v3
        }
        v3 -= (v3 | (v3 | (^(v3))))
        if (1000 <= 1000) {
            if (!(true && ((v3 * -79) == 0))) {
                var v75 float32 = 0.0
                _ = v75
            } else {
                v76 := GenS6{F0: 0, F1: 0}
                fmt.Println(v4, v3)
                _ = v76
            }
            if (1e10 <= 0.0) {
                for fi77 := 0; fi77 < 7; fi77++ {
                    v4 -= (v4 / (((v4 * v4)) | 1))
                    fmt.Println(v4)
                    _ = fi77
                }
                var v78 int32 = 2147483647
                _ = v78
            }
            v79 := GenS7{F0: "a"}
            _ = v79
        }
        _ = fi73
    }
    switch ("foo bar" + "") {
    case "foo bar":
        _ = len(v1)
    case "a":
        v4 -= (v4 >> ((18446744073709551615) & 63))
    }
    v80 := GenS4{F0: 42, F1: 509}
    fmt.Println(v4, v3)
    v4 += v4
    v81 := GenS7{F0: ""}
    v4 += v4
    var v82 int64 = 0
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v82:", v82)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v80
    _ = v81
    _ = v82
}
