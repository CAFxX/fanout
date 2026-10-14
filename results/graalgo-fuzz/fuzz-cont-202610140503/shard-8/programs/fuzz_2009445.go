package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint8; F1 int64; F2 string }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint64; F1 uint16 }
type GenS6 struct { F0 uint; F1 uint64; F2 int }
type GenS7 struct { F0 int8 }
func genFn0(p0 uint64) int32 {
    v6 := []int{1, 42}
    _ = len(v6)
    var v7 int8 = 84
    _ = p0
    _ = v6
    _ = v7
    return 41
}
func genFn1(p0 uint32, p1 int16) int64 {
    var v8 uint8 = uint8(p0)
    v9 := GenS1{F0: 18446744073709551615}
    p0 = (p0 << ((2) & 63))
    v10 := GenS1{F0: uint(v8)}
    _ = p0
    _ = p1
    _ = v10
    _ = v8
    _ = v9
    return 1
}

func main() {
    v1 := []uint32{}
    var v2 int = -1
    v3 := make(map[string]int)
    var v4 float64 = float64(0.0)
    v5 := []string{"a", "\n\t", "a"}
    v4 -= v4
    v2 = 0
    v4 += v4
    for fi11 := 0; fi11 < 16; fi11++ {
        for fi12 := 0; fi12 < 9; fi12++ {
            var v13 int32 = 511
            if (v4 <= (1.5 * 0.0)) {
                switch "foo bar" {
                case "foo bar":
                    var v14 uint64 = 465
                    _ = v14
                }
            } else {
                v15 := GenS2{F0: true}
                _ = v15
            }
            _ = fi12
            _ = v13
        }
        var v16 int32 = -1
        fmt.Println(v4, v2)
        _ = fi11
        _ = v16
    }
    v17 := GenS1{F0: 501}
    v4 = 0.1
    switch "foo bar" {
    case "a":
        if false {
            switch ("hello" + "\n\t") {
            case "\n\t":
                v4 = 3.14159
            case "a":
                var v18 float32 = (float32((v4 / v4)) - 0.0)
                _ = v18
            case "hello":
                var v19 int = v2
                _ = v19
            default:
                v4 += float64(3.14159)
            }
        }
    case "":
        switch (v2 | (v2 >> ((1) & 63))) {
        case 42:
            v20 := GenS2{F0: true}
            _ = v20
        case 329:
            if false {
                if ("hello" == "foo bar") {
                    fmt.Println(v4, v2)
                    if false {
                        var v21 uint16 = 0
                        _ = v21
                    } else {
                        var v22 string = ""
                        v3["hello"] = v2
                        v23 := GenS1{F0: 255}
                        _ = v22
                        _ = v23
                    }
                }
                var v24 float64 = v4
                fmt.Println(v2, v2)
                _ = v24
            } else {
                v25 := GenS1{F0: 37}
                _ = v25
            }
        case 1:
            v5 = append(v5, "foo bar")
        }
    case "\n\t":
        switch "\n\t" {
        case "\n\t":
            if true {
                if len(v5) > 0 {
                    v5[0] = "hello"
                }
                var v26 uint = 4294967295
                _ = v26
            }
        case "hello":
            var v27 int8 = -1
            _ = v27
        default:
            fmt.Println(v4, v2)
        }
    default:
        if (1 == 42) {
            for fi28 := 0; fi28 < 2; fi28++ {
                v29 := GenS2{F0: ((42 != 0) || true)}
                if (2 <= 18446744073709551615) {
                    switch fi28 {
                    case -1:
                        for fi30 := 0; fi30 < 6; fi30++ {
                            v31 := v3[("foo bar" + "a")]
                            switch 127 {
                            case 0:
                                v32 := GenS3{F0: 0, F1: 42, F2: "\n\t"}
                                var v33 int8 = -128
                                if false {
                                    v34 := GenS1{F0: 65535}
                                    var v35 uint32 = 42
                                    _ = v34
                                    _ = v35
                                }
                                _ = v32
                                _ = v33
                            case 412:
                                v36 := GenS2{F0: (!true)}
                                _ = v36
                            default:
                                fmt.Println(fi30, v4)
                            }
                            v31 = fi30
                            _ = fi30
                            _ = v31
                        }
                    }
                    v4 = (float64((3.14159 + 1.5)) * (((1.5 / ((1e10 + 3.14159) + 1e10)) + 0.0) / v4))
                    v3[((("\n\t" + ("\n\t" + ("hello" + "hello"))) + "hello") + ("a" + "a"))] = (^(fi28))
                } else {
                    if false {
                        fmt.Println(fi28)
                        if len(v1) > 0 {
                            v1[0] = 4294967295
                        }
                    } else {
                        for fi37 := 0; fi37 < 5; fi37++ {
                            var v38 int8 = -128
                            var v39 float64 = (0.0 * 3.14159)
                            if (65535 <= 680) {
                                v40 := GenS4{F0: "foo bar"}
                                _ = v40
                            } else {
                                v4 += float64(1.5)
                                var v41 int64 = 0
                                _ = v41
                            }
                            if (((v39 / v39) + 0.1) >= v39) {
                                v39 = 1.5
                                v42 := GenS2{F0: true}
                                switch fi28 {
                                case 9223372036854775807:
                                    switch "hello" {
                                    case "foo bar":
                                        v4 = (v4 / v4)
                                    case "":
                                        v43 := GenS5{F0: 2, F1: 384}
                                        _ = v43
                                    case "a":
                                        var v44 float32 = (0.0 - (float32(3.14159) / (0.0 / float32(((-2.25 * v4) - 3.14159)))))
                                        _ = v44
                                    }
                                case 128:
                                    v3[("foo bar" + "foo bar")] = 9223372036854775807
                                case 2147483647:
                                    if true {
                                        v1 = append(v1, 410)
                                        v3["foo bar"] = fi37
                                        if (false && (!true)) {
                                            fmt.Println(v2, fi28)
                                        }
                                        if (true || ((-219 == 644) && (v2 <= fi28))) {
                                            var v45 int64 = -683
                                            fmt.Println(v2, fi37, v2)
                                            _ = len(v3)
                                            _ = v45
                                        }
                                    }
                                }
                                var v46 float32 = 3.14159
                                var v47 uint64 = 4294967295
                                var v48 int = (fi28 - (v2 << ((374) & 63)))
                                v46 = (v46 * (v46 - v46))
                                _ = v42
                                _ = v46
                                _ = v47
                                _ = v48
                            } else {
                                if ((!false) || (-1 <= 2147483647)) {
                                    v49 := v3["hello"]
                                    _ = v49
                                }
                                v50 := GenS3{F0: 127, F1: 9223372036854775807, F2: "hello"}
                                var v51 bool = false
                                _ = v50
                                _ = v51
                            }
                            _ = fi37
                            _ = v38
                            _ = v39
                        }
                    }
                }
                _ = fi28
                _ = v29
            }
        } else {
            v3["hello"] = v2
            v52 := GenS2{F0: (127 <= 212)}
            _ = v52
        }
    }
    for fi53 := 0; fi53 < 12; fi53++ {
        fmt.Println(v4)
        var v54 int32 = -213
        for fi55 := 0; fi55 < 18; fi55++ {
            v2 = 128
            if (!true) {
                v56 := GenS6{F0: 42, F1: 0, F2: (fi53 >> ((830) & 63))}
                var v57 uint = 278
                _ = v56
                _ = v57
            } else {
                fmt.Println(fi55)
                delete(v3, "\n\t")
            }
            if (("hello" + "") == "") {
                _ = len(v1)
                v58 := GenS4{F0: "foo bar"}
                _ = v58
            }
            _ = fi55
        }
        _ = fi53
        _ = v54
    }
    v59 := GenS5{F0: 118, F1: 661}
    v60 := GenS1{F0: 0}
    var v61 float64 = v4
    var v62 int32 = 307
    switch (-(v2)) {
    case 1:
        v5 = append(v5, "")
    case 32768:
        v62 = -678
    }
    v4 -= v4
    var v63 int32 = v62
    v5 = append(v5, ("foo bar" + "a"))
    switch ("foo bar" + ("a" + "foo bar")) {
    case "hello":
        var v64 int8 = 42
        _ = v64
    case "\n\t":
        var v65 int64 = 157
        _ = v65
    case "foo bar":
        if (false || (255 != 513)) {
            for fi66 := 0; fi66 < 12; fi66++ {
                delete(v3, ("foo bar" + "a"))
                _ = fi66
            }
            fmt.Println(v62)
            var v67 uint16 = 1
            _ = v67
        }
    }
    if false {
        switch (^(v2)) {
        case -864:
            for fi68 := 0; fi68 < 12; fi68++ {
                var v69 float32 = (1e10 + (float32(v4) / 0.1))
                var v70 uint16 = 125
                if false {
                    var v71 int32 = (v62 % ((int32(fi68)) | 1))
                    var v72 float64 = float64(3.14159)
                    _ = v71
                    _ = v72
                } else {
                    switch "" {
                    case "":
                        delete(v3, "\n\t")
                    case "\n\t":
                        v2 = v2
                    }
                    switch (("foo bar" + "") + "hello") {
                    case "a":
                        v73 := GenS4{F0: ("\n\t" + "\n\t")}
                        _ = v73
                    case "hello":
                        if true {
                            fmt.Println(v62, v62)
                            var v74 int = (-(fi68))
                            _ = v74
                        } else {
                            v75 := GenS6{F0: 4294967295, F1: 0, F2: (v2 | (fi68 + fi68))}
                            _ = v75
                        }
                    case "foo bar":
                        if true {
                            for fi76 := 0; fi76 < 11; fi76++ {
                                v77 := GenS1{F0: 1}
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                var v78 int64 = 1
                                _ = fi76
                                _ = v77
                                _ = v78
                            }
                            for fi79 := 0; fi79 < 17; fi79++ {
                                v4 += 0.1
                                v80 := GenS7{F0: 127}
                                _ = fi79
                                _ = v80
                            }
                            switch "a" {
                            case "\n\t":
                                switch ("" + "a") {
                                case "hello":
                                    var v81 int8 = -85
                                    _ = v81
                                default:
                                    for fi82 := 0; fi82 < 5; fi82++ {
                                        var v83 uint = 4294967295
                                        _ = fi82
                                        _ = v83
                                    }
                                }
                            default:
                                for fi84 := 0; fi84 < 1; fi84++ {
                                    var v85 int64 = 9223372036854775807
                                    switch (v2 ^ fi68) {
                                    case 592:
                                        v62 = v63
                                    case 32768:
                                        v62 -= (v62 ^ (v62 % (((v62 + (-(v62)))) | 1)))
                                        v62 = v63
                                    }
                                    _ = fi84
                                    _ = v85
                                }
                            }
                        }
                    }
                    var v86 float32 = ((1e10 - (1e10 * v69)) * ((((v69 - 1e10) / 1.5) + (v69 / v69)) + 3.14159))
                    _ = v86
                }
                _ = fi68
                _ = v69
                _ = v70
            }
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v61:", v61)
    fmt.Println("v62:", v62)
    fmt.Println("v63:", v63)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v59
    _ = v60
    _ = v61
    _ = v62
    _ = v63
}
