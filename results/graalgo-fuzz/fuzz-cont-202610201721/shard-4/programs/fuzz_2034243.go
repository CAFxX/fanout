package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint64; F1 bool; F2 int16 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 bool; F1 int }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 int64; F1 uint; F2 int16 }
type GenS8 struct { F0 string; F1 int16 }
func genFn0(p0 uint16, p1 uint8, p2 int) string {
    p2 = (p2 | 42)
    p0 += (p0 | 69)
    _ = p0
    _ = p1
    _ = p2
    return ((("a" + ("a" + "")) + "\n\t") + "foo bar")
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[string]string)
    v3 := make(map[int]string)
    switch ("a" + "foo bar") {
    case "\n\t":
        if (0 < 42) {
            switch 2147483647 {
            case -1:
                var v4 uint32 = 2
                _ = v4
            case 9223372036854775807:
                var v5 uint8 = 255
                _ = v5
            }
            var v6 uint16 = 0
            _ = v6
        }
    }
    var v7 int32 = 0
    for fi8 := 0; fi8 < 15; fi8++ {
        if (false && (!true)) {
            var v9 int64 = 9223372036854775807
            v10 := []int{}
            _ = v10
            _ = v9
        }
        _ = fi8
    }
    fmt.Println(v7)
    v7 = (^(v7))
    v7 = v7
    for fi11 := 0; fi11 < 19; fi11++ {
        v12 := GenS1{F0: v7}
        for fi13 := 0; fi13 < 20; fi13++ {
            v7 -= (v7 / (((v7 << ((0) & 63))) | 1))
            var v14 int64 = -1
            if ((("a" + "hello") >= "a") || false) {
                var v15 bool = (!true)
                _ = v15
            } else {
                v16 := GenS2{F0: 820, F1: true, F2: -993}
                v17 := []int{(fi11 & fi11), -123}
                _ = len(v2)
                var v18 bool = (false || (76 <= 4294967295))
                if (v18 && (false && false)) {
                    v17 = append(v17, (fi11 & fi13))
                    v19 := v1[fi11]
                    _ = v19
                }
                _ = v16
                _ = v17
                _ = v18
            }
            v20 := GenS3{F0: v7}
            if false {
                v7 -= 415
            } else {
                v21 := GenS4{F0: false, F1: fi13}
                var v22 int8 = -68
                _ = v21
                _ = v22
            }
            var v23 string = ("" + "")
            v7 = (v7 & (-(v7)))
            _ = fi13
            _ = v14
            _ = v20
            _ = v23
        }
        _ = fi11
        _ = v12
    }
    v7 = (^(v7))
    if (("foo bar" + "foo bar") <= "\n\t") {
        v24 := v1[939]
        v7 = -387
        v7 = (^(v7))
        _ = v24
    } else {
        var v25 int64 = -9223372036854775808
        v26 := GenS2{F0: 2, F1: ((378 > 42) || true), F2: 0}
        switch (("a" + "\n\t") + "a") {
        case "hello":
            for fi27 := 0; fi27 < 17; fi27++ {
                if false {
                    var v28 int = 2147483647
                    _ = v28
                } else {
                    for fi29 := 0; fi29 < 19; fi29++ {
                        v25 = 42
                        var v30 float64 = -2.25
                        if (false || (true && true)) {
                            var v31 string = (("" + "\n\t") + "\n\t")
                            _ = v31
                        } else {
                            var v32 string = "foo bar"
                            _ = v32
                        }
                        var v33 int16 = -441
                        if ((!false) && false) {
                            v30 -= (v30 - v30)
                            var v34 uint16 = 1
                            v25 = 9223372036854775807
                            var v35 bool = false
                            var v36 string = ("hello" + "")
                            v37 := GenS1{F0: (v7 | (v7 >> ((65535) & 63)))}
                            v36 = "foo bar"
                            fmt.Println(v35)
                            if true {
                                v25 += 9223372036854775807
                                v38 := GenS2{F0: 889, F1: true, F2: 42}
                                _ = len(v1)
                                _ = v38
                            } else {
                                fmt.Println(fi29, fi29)
                            }
                            switch v36 {
                            case "":
                                switch "hello" {
                                case "foo bar":
                                    var v39 int = -9223372036854775808
                                    _ = v39
                                }
                            }
                            _ = v34
                            _ = v35
                            _ = v36
                            _ = v37
                        }
                        var v40 int = fi29
                        v41 := GenS5{F0: 32767}
                        _ = fi29
                        _ = v30
                        _ = v33
                        _ = v40
                        _ = v41
                    }
                    switch "" {
                    case "foo bar":
                        switch int(v25) {
                        case -340:
                            v25 = 864
                        default:
                            switch (("foo bar" + "foo bar") + "hello") {
                            case "\n\t":
                                for fi42 := 0; fi42 < 4; fi42++ {
                                    v25 = -1
                                    _ = fi42
                                }
                            case "foo bar":
                                var v43 float64 = 0.1
                                _ = v43
                            case "hello":
                                for fi44 := 0; fi44 < 15; fi44++ {
                                    var v45 uint16 = 33
                                    if false {
                                        v46 := GenS4{F0: false, F1: 42}
                                        switch fi44 {
                                        case -2147483648:
                                            v45 -= 65535
                                        case 128:
                                            v7 += (-(v7))
                                        case 127:
                                            v45 = v45
                                        default:
                                            var v47 uint32 = uint32(v45)
                                            _ = v47
                                        }
                                        var v48 float32 = 0.1
                                        _ = v46
                                        _ = v48
                                    } else {
                                        switch fi27 {
                                        case 9223372036854775807:
                                            v45 = (v45 + v45)
                                        case 127:
                                            v7 = v7
                                        default:
                                            fmt.Println(fi44)
                                        }
                                    }
                                    _ = fi44
                                    _ = v45
                                }
                            default:
                                if false {
                                    if (42 > 448) {
                                        v49 := GenS6{F0: (("" + ("a" + ("hello" + "\n\t"))) + "foo bar")}
                                        _ = v49
                                    } else {
                                        v7 = v7
                                        v7 -= -1
                                        v50 := GenS1{F0: -1}
                                        _ = v50
                                    }
                                    var v51 uint8 = 42
                                    switch "a" {
                                    case "":
                                        var v52 uint8 = 128
                                        _ = v52
                                    default:
                                        var v53 uint = 601
                                        _ = v53
                                    }
                                    _ = v51
                                } else {
                                    v25 = (v25 * (v25 << ((42) & 63)))
                                    for fi54 := 0; fi54 < 2; fi54++ {
                                        if (255 >= 0) {
                                            v7 += (v7 & -919)
                                            v7 += 42
                                            switch ("a" + "hello") {
                                            case "a":
                                                v7 = v7
                                            }
                                        } else {
                                            v55 := GenS7{F0: 0, F1: 42, F2: 32767}
                                            var v56 uint64 = 4294967295
                                            if ("" == "foo bar") {
                                                v57 := GenS2{F0: (v56 / ((0) | 1)), F1: (!(uint8(v56) < 128)), F2: 32767}
                                                _ = v57
                                            } else {
                                                switch fi27 {
                                                case 128:
                                                    switch "\n\t" {
                                                    case "\n\t":
                                                        switch (("" + "foo bar") + "") {
                                                        case "\n\t":
                                                            var v58 uint64 = (v56 ^ 4294967295)
                                                            _ = v58
                                                        case "":
                                                            if false {
                                                                v7 -= (v7 * 1)
                                                                if true {
                                                                    switch (^(fi27)) {
                                                                    case 59:
                                                                        v25 = v25
                                                                    case 1:
                                                                        fmt.Println(v25, v7)
                                                                    case -9223372036854775808:
                                                                        v7 = 42
                                                                    }
                                                                    if (-1 <= 0) {
                                                                        v25 -= 1
                                                                    }
                                                                } else {
                                                                    if false {
                                                                        var v59 int32 = v7
                                                                        var v60 string = ""
                                                                        var v61 float32 = (1e10 + -2.25)
                                                                        _ = v59
                                                                        _ = v60
                                                                        _ = v61
                                                                    }
                                                                    if true {
                                                                        var v62 uint32 = 4294967295
                                                                        var v63 bool = (!("foo bar" != "\n\t"))
                                                                        v64 := v2[((("a" + "hello") + "a") + ("foo bar" + "hello"))]
                                                                        _ = v62
                                                                        _ = v63
                                                                        _ = v64
                                                                    }
                                                                    var v65 string = ("" + "hello")
                                                                    _ = v65
                                                                }
                                                                var v66 uint8 = 26
                                                                _ = v66
                                                            } else {
                                                                v67 := GenS4{F0: (!false), F1: fi54}
                                                                _ = v67
                                                            }
                                                        }
                                                    case "foo bar":
                                                        var v68 string = ((("foo bar" + "foo bar") + "a") + "")
                                                        _ = v68
                                                    default:
                                                        v7 -= (v7 & v7)
                                                    }
                                                default:
                                                    switch -557 {
                                                    case 32767:
                                                        fmt.Println(fi27, v56)
                                                        v69 := v2[""]
                                                        var v70 uint16 = 236
                                                        switch fi54 {
                                                        case 127:
                                                            v71 := GenS8{F0: "hello", F1: 32767}
                                                            _ = v71
                                                        case 0:
                                                            v56 = (v56 - (^(v56)))
                                                        default:
                                                            v3[-1] = "foo bar"
                                                        }
                                                        var v72 uint16 = (v70 & (v70 + v70))
                                                        v70 -= 255
                                                        _ = v69
                                                        _ = v70
                                                        _ = v72
                                                    case -1:
                                                        v2["foo bar"] = "a"
                                                    case 2147483647:
                                                        fmt.Println(v7)
                                                    }
                                                }
                                            }
                                            _ = v55
                                            _ = v56
                                        }
                                        _ = fi54
                                    }
                                    switch ("foo bar" + "") {
                                    case "hello":
                                        if (!false) {
                                            v7 = -788
                                            for fi73 := 0; fi73 < 10; fi73++ {
                                                v25 += (v25 | v25)
                                                v7 = v7
                                                _ = fi73
                                            }
                                        }
                                    }
                                }
                            }
                        }
                    default:
                        v74 := GenS6{F0: "\n\t"}
                        _ = v74
                    }
                }
                for fi75 := 0; fi75 < 17; fi75++ {
                    var v76 int = fi75
                    if true {
                        v25 -= v25
                    }
                    _ = fi75
                    _ = v76
                }
                v25 -= v25
                _ = fi27
            }
        }
        _ = v25
        _ = v26
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v7
}
