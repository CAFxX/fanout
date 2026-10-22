package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint8; F1 int8 }
type GenS4 struct { F0 int; F1 uint; F2 int32 }
type GenS5 struct { F0 uint8; F1 uint }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 uint16 }
func genFn0(p0 bool) bool {
    if p0 {
        fmt.Println(p0)
        v4 := GenS1{F0: 2}
        v5 := GenS1{F0: 2}
        _ = v4
        _ = v5
    } else {
        p0 = (327 < 4294967295)
        if (("foo bar" + "hello") <= "") {
            var v6 uint16 = 1
            p0 = (1 != -53)
            for fi7 := 0; fi7 < 5; fi7++ {
                var v8 int = -849
                for fi9 := 0; fi9 < 5; fi9++ {
                    if p0 {
                        if p0 {
                            v10 := GenS1{F0: 0}
                            v6 = (v6 >> ((987) & 63))
                            _ = v10
                        } else {
                            v11 := []int{127}
                            fmt.Println(v6)
                            _ = v11
                        }
                        fmt.Println(fi7, p0)
                        var v12 float64 = 3.14159
                        _ = v12
                    } else {
                        fmt.Println(v8, fi7)
                        v13 := GenS1{F0: 2}
                        var v14 int8 = 1
                        _ = v13
                        _ = v14
                    }
                    var v15 string = "\n\t"
                    v8 += (^(fi7))
                    _ = fi9
                    _ = v15
                }
                _ = fi7
                _ = v8
            }
            _ = v6
        }
        var v16 int = 318
        _ = v16
    }
    v17 := GenS1{F0: 65535}
    for fi18 := 0; fi18 < 20; fi18++ {
        fmt.Println(fi18)
        _ = fi18
    }
    for fi19 := 0; fi19 < 13; fi19++ {
        var v20 int32 = -553
        if false {
            var v21 uint64 = 845
            switch 572 {
            case -773:
                var v22 int64 = 42
                _ = v22
            case 9223372036854775807:
                if false {
                    for fi23 := 0; fi23 < 13; fi23++ {
                        v24 := GenS2{F0: 32767}
                        v25 := GenS2{F0: 8}
                        _ = fi23
                        _ = v24
                        _ = v25
                    }
                } else {
                    v21 = (v21 - v21)
                }
            case 270:
                for fi26 := 0; fi26 < 18; fi26++ {
                    switch (fi26 & fi19) {
                    case 128:
                        fmt.Println(fi26, p0, v20)
                    }
                    var v27 uint8 = 17
                    _ = fi26
                    _ = v27
                }
            default:
                if true {
                    var v28 bool = true
                    _ = v28
                } else {
                    fmt.Println(p0, v20)
                }
            }
            _ = v21
        } else {
            var v29 int16 = 0
            _ = v29
        }
        _ = fi19
        _ = v20
    }
    _ = p0
    _ = v17
    return ((0.0 - (0.0 + float64((3.14159 / 0.1)))) == (float64(((0.0 + -2.25) + (3.14159 * 1.5))) * float64((0.0 - 3.14159))))
}
func genFn1() uint64 {
    v30 := []int{}
    var v31 int64 = 1
    if true {
        var v32 int = -2147483648
        if true {
            fmt.Println(v32, v32)
        } else {
            v33 := GenS3{F0: 2, F1: 127}
            _ = v33
        }
        _ = v32
    }
    fmt.Println(v31)
    v34 := GenS4{F0: 137, F1: 255, F2: 1}
    _ = v30
    _ = v31
    _ = v34
    return 18446744073709551615
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[int]uint32)
    v3 := make(map[int]int)
    var v35 int16 = 749
    v35 = v35
    fmt.Println(v35)
    v36 := []int{32768, 85}
    for fi37 := 0; fi37 < 3; fi37++ {
        v35 += (v35 - v35)
        _ = fi37
    }
    fmt.Println(v35)
    v38 := GenS5{F0: 128, F1: 1}
    fmt.Println(v35)
    v35 = (v35 / (((v35 % ((v35) | 1))) | 1))
    v39 := v2[-1]
    for fi40 := 0; fi40 < 19; fi40++ {
        var v41 int64 = -775
        v42 := GenS5{F0: 42, F1: 255}
        _ = fi40
        _ = v41
        _ = v42
    }
    v43 := v1["a"]
    switch ((v43 + v43) + (v43 + v43)) {
    case "hello":
        for fi44 := 0; fi44 < 8; fi44++ {
            for fi45 := 0; fi45 < 16; fi45++ {
                if true {
                    _ = len(v36)
                } else {
                    if true {
                        var v46 int8 = -1
                        switch (fi44 ^ (fi45 - fi45)) {
                        case 0:
                            if len(v36) > 0 {
                                v36[0] = fi45
                            }
                        case -9223372036854775808:
                            var v47 uint64 = 42
                            _ = v47
                        case -2147483648:
                            v35 += v35
                        default:
                            var v48 uint8 = 128
                            _ = v48
                        }
                        _ = v46
                    }
                    v49 := v1[(v43 + v43)]
                    _ = v49
                }
                var v50 uint32 = v39
                _ = fi45
                _ = v50
            }
            for fi51 := 0; fi51 < 13; fi51++ {
                if (!false) {
                    var v52 float64 = 1e10
                    if len(v36) >= 2 {
                        v36 = v36[1:len(v36)]
                    }
                    fmt.Println(v52, fi44, v35)
                    _ = v52
                } else {
                    v53 := GenS4{F0: -955, F1: 586, F2: 1}
                    if (0.1 >= (1e10 + -2.25)) {
                        v54 := GenS4{F0: 42, F1: 1, F2: 42}
                        _ = v54
                    } else {
                        if len(v36) >= 2 {
                            v36 = v36[1:len(v36)]
                        }
                        if true {
                            switch 791 {
                            case -9223372036854775808:
                                v55 := GenS1{F0: 551}
                                _ = v55
                            }
                        } else {
                            _ = len(v1)
                        }
                        v56 := GenS5{F0: 90, F1: 2}
                        _ = v56
                    }
                    var v57 bool = false
                    _ = v53
                    _ = v57
                }
                if ((1e10 == float32(1.5)) || true) {
                    if (!true) {
                        if ((fi44 / ((875) | 1)) < (fi44 - fi44)) {
                            fmt.Println(v35)
                            v58 := GenS6{F0: 218}
                            _ = len(v36)
                            _ = v58
                        }
                        fmt.Println(fi51)
                    } else {
                        if true {
                            if false {
                                var v59 int32 = 42
                                _ = v59
                            }
                        } else {
                            fmt.Println(v39)
                            if true {
                                var v60 int = fi51
                                _ = v60
                            } else {
                                v35 = v35
                                fmt.Println(fi51, fi44)
                                switch (^(fi44)) {
                                case 32768:
                                    fmt.Println(fi51, v43, v39)
                                case -861:
                                    switch 553 {
                                    case 80:
                                        v43 = (((((v43 + ("hello" + v43)) + "\n\t") + (v43 + v43)) + "foo bar") + v43)
                                    default:
                                        var v61 uint = 2
                                        _ = v61
                                    }
                                default:
                                    var v62 int64 = 42
                                    _ = v62
                                }
                            }
                        }
                    }
                } else {
                    if len(v36) >= 2 {
                        v36 = v36[1:len(v36)]
                    }
                    v35 += 1
                    var v63 float32 = (3.14159 * 1.5)
                    _ = v63
                }
                fmt.Println(fi51)
                _ = fi51
            }
            _ = fi44
        }
    default:
        if (true || false) {
            if ((0.1 * (float64(-2.25) + 1e10)) == -2.25) {
                switch ("\n\t" + "\n\t") {
                case "\n\t":
                    switch v43 {
                    case "\n\t":
                        if true {
                            var v64 uint32 = (^(v39))
                            _ = v64
                        } else {
                            if len(v36) >= 2 {
                                v36 = v36[1:len(v36)]
                            }
                        }
                    }
                case "foo bar":
                    switch "foo bar" {
                    case "\n\t":
                        switch 127 {
                        case 127:
                            var v65 int64 = 0
                            _ = v65
                        case -626:
                            var v66 float64 = (float64(3.14159) + (0.0 * float64(3.14159)))
                            _ = v66
                        case 42:
                            if len(v36) > 0 {
                                v36[0] = -700
                            }
                        }
                    default:
                        for fi67 := 0; fi67 < 20; fi67++ {
                            switch "foo bar" {
                            case "\n\t":
                                v35 = (v35 / ((v35) | 1))
                            case "foo bar":
                                switch (v43 + ("hello" + (v43 + v43))) {
                                case "\n\t":
                                    for fi68 := 0; fi68 < 4; fi68++ {
                                        if false {
                                            if false {
                                                if (!(42 < 9223372036854775807)) {
                                                    v43 = (v43 + "a")
                                                    var v69 int64 = -1
                                                    if true {
                                                        v39 += (v39 << ((1) & 63))
                                                        v39 = v39
                                                        var v70 string = v43
                                                        switch -587 {
                                                        case 112:
                                                            if len(v36) >= 2 {
                                                                v36 = v36[1:len(v36)]
                                                            }
                                                        case -1:
                                                            var v71 uint16 = 0
                                                            _ = v71
                                                        case 9223372036854775807:
                                                            var v72 uint16 = 42
                                                            _ = v72
                                                        }
                                                        v35 = v35
                                                        _ = v70
                                                    } else {
                                                        fmt.Println(fi67)
                                                        v69 = (v69 % ((466) | 1))
                                                    }
                                                    _ = v69
                                                } else {
                                                    v36 = append(v36, 4294967295)
                                                }
                                                v43 = ((v43 + "a") + ((v43 + "") + "foo bar"))
                                                fmt.Println(fi67)
                                            } else {
                                                v73 := GenS5{F0: 127, F1: uint(v39)}
                                                var v74 uint = 613
                                                _ = v73
                                                _ = v74
                                            }
                                            fmt.Println(fi67, fi68)
                                            _ = len(v36)
                                        }
                                        v75 := GenS7{F0: 42}
                                        switch (fi67 + fi68) {
                                        case 42:
                                            v76 := GenS4{F0: (fi67 * 1), F1: 65535, F2: -2147483648}
                                            _ = v76
                                        case -781:
                                            v39 -= v39
                                        case 32768:
                                            v39 += v39
                                        default:
                                            if true {
                                                var v77 int = (fi68 * fi68)
                                                if true {
                                                    if (4294967295 < v77) {
                                                        if (356 == 906) {
                                                            v2[fi68] = 65535
                                                            var v78 float64 = (-2.25 / 3.14159)
                                                            _ = v78
                                                        } else {
                                                            v35 += (v35 - (v35 * v35))
                                                        }
                                                    }
                                                }
                                                _ = v77
                                            } else {
                                                var v79 uint16 = 642
                                                var v80 int64 = 553
                                                fmt.Println(v39)
                                                _ = v79
                                                _ = v80
                                            }
                                        }
                                        _ = fi68
                                        _ = v75
                                    }
                                case "foo bar":
                                    v35 -= 893
                                case "a":
                                    v35 -= -702
                                }
                            case "a":
                                v35 -= v35
                            default:
                                v35 = -679
                            }
                            v43 = v43
                            var v81 int32 = 42
                            _ = fi67
                            _ = v81
                        }
                    }
                }
            } else {
                for fi82 := 0; fi82 < 14; fi82++ {
                    switch v43 {
                    case "":
                        fmt.Println(v35, v35)
                    default:
                        if false {
                            if len(v36) >= 2 {
                                v36 = v36[1:len(v36)]
                            }
                            fmt.Println(v35)
                        }
                    }
                    var v83 int32 = 42
                    v39 = 1
                    _ = fi82
                    _ = v83
                }
                if true {
                    var v84 float32 = ((1.5 * 0.1) + 0.0)
                    _ = v84
                }
            }
            v43 = v43
            var v85 float64 = -2.25
            _ = v85
        } else {
            v39 = v39
            v43 = "hello"
        }
    }
    v86 := GenS8{F0: 65535}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v36:", len(v36), v36)
    fmt.Println("v39:", v39)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v38
    _ = v39
    _ = v43
    _ = v86
}
