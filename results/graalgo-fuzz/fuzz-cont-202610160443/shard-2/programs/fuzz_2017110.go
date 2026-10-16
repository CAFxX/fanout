package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int32; F1 uint32 }
type GenS3 struct { F0 uint64; F1 int64 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 string }
type GenS7 struct { F0 uint8; F1 bool; F2 uint8 }
type GenS8 struct { F0 uint32 }
type GenS9 struct { F0 uint32 }
type GenS10 struct { F0 int; F1 uint; F2 uint16 }
type GenS11 struct { F0 string; F1 uint16; F2 bool }
func genFn0(p0 bool, p1 bool, p2 bool) bool {
    v5 := GenS1{F0: 65535}
    fmt.Println(p2, p0, p0)
    v6 := []int{509, 4294967295, -1, -2147483648}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return (true || p1)
}
func genFn1() bool {
    var v7 int = 0
    switch (v7 >> ((255) & 63)) {
    case 9223372036854775807:
        v8 := []int{}
        _ = v8
    case 32767:
        v9 := GenS2{F0: 0, F1: 2147483647}
        _ = v9
    }
    for fi10 := 0; fi10 < 17; fi10++ {
        switch fi10 {
        case 2147483647:
            for fi11 := 0; fi11 < 6; fi11++ {
                v7 = 4294967295
                fmt.Println(fi10)
                fmt.Println(fi11, fi10)
                v7 = (^(fi10))
                var v12 uint8 = 42
                var v13 string = "hello"
                var v14 int64 = 451
                _ = fi11
                _ = v12
                _ = v13
                _ = v14
            }
        default:
            fmt.Println(v7, v7)
        }
        if false {
            for fi15 := 0; fi15 < 11; fi15++ {
                v16 := GenS2{F0: -1, F1: 74}
                if (float32((float64(((-2.25 + 1.5) - (0.0 * 0.1))) / 3.14159)) >= 0.0) {
                    v17 := GenS3{F0: 800, F1: 9223372036854775807}
                    if false {
                        v18 := []int{fi10, -525, (fi15 & fi10)}
                        var v19 bool = true
                        _ = v18
                        _ = v19
                    }
                    _ = v17
                } else {
                    switch "foo bar" {
                    case "\n\t":
                        fmt.Println(fi15)
                    case "foo bar":
                        v20 := GenS3{F0: 1, F1: -1}
                        v21 := GenS2{F0: -2147483648, F1: 0}
                        _ = v20
                        _ = v21
                    case "hello":
                        v22 := GenS3{F0: 2, F1: 42}
                        _ = v22
                    default:
                        fmt.Println(fi10, fi10)
                    }
                }
                v23 := GenS2{F0: -677, F1: 924}
                _ = fi15
                _ = v16
                _ = v23
            }
            v7 = (fi10 % ((214) | 1))
        } else {
            v24 := GenS1{F0: 38}
            var v25 int8 = -116
            switch ("a" + (("foo bar" + ("" + (("hello" + "") + "foo bar"))) + "hello")) {
            case "":
                for fi26 := 0; fi26 < 18; fi26++ {
                    var v27 float64 = float64((3.14159 / 1e10))
                    _ = fi26
                    _ = v27
                }
            default:
                if false {
                    var v28 uint32 = 2147483648
                    for fi29 := 0; fi29 < 5; fi29++ {
                        v30 := []int{(fi29 + fi29), v7}
                        v31 := GenS4{F0: -14}
                        _ = fi29
                        _ = v30
                        _ = v31
                    }
                    _ = v28
                }
            }
            _ = v24
            _ = v25
        }
        v32 := GenS1{F0: 643}
        _ = fi10
        _ = v32
    }
    switch v7 {
    case 127:
        for fi33 := 0; fi33 < 19; fi33++ {
            v7 -= (fi33 & (fi33 * fi33))
            v7 = v7
            _ = fi33
        }
    case 1:
        var v34 int = v7
        _ = v34
    case 128:
        v7 -= 42
    default:
        switch "hello" {
        case "a":
            var v35 int8 = 1
            _ = v35
        case "foo bar":
            fmt.Println(v7)
        }
    }
    _ = v7
    return (0 > 1)
}

func main() {
    v1 := make(map[uint32]int)
    var v2 string = ("\n\t" + ("foo bar" + "foo bar"))
    var v3 int = 13
    var v4 int32 = 42
    var v36 string = "hello"
    for fi37 := 0; fi37 < 11; fi37++ {
        fmt.Println(v3, v3, fi37)
        _ = fi37
    }
    fmt.Println(v4, v4)
    var v38 uint64 = 0
    if (("hello" + "\n\t") == "foo bar") {
        if true {
            var v39 int32 = -667
            _ = v39
        } else {
            v40 := GenS5{F0: ""}
            _ = v40
        }
        fmt.Println(v2, v4)
        if false {
            var v41 uint = 184
            if false {
                var v42 string = "foo bar"
                if (false || (true || true)) {
                    v43 := []int{v3, (v3 & (-(v3))), (-(v3)), v3}
                    v41 -= v41
                    if len(v43) >= 2 {
                        v43 = v43[1:len(v43)]
                    }
                    _ = v43
                }
                _ = v42
            }
            _ = v41
        } else {
            delete(v1, 42)
        }
    } else {
        if (2147483647 > 1) {
            switch v3 {
            case 32768:
                delete(v1, 490)
            }
            v1[621] = (v3 % ((42) | 1))
        } else {
            fmt.Println(v38, v4, v2)
            v2 = v36
        }
        if ((42 <= 194) || (!(848 == 0))) {
            if (!true) {
                for fi44 := 0; fi44 < 14; fi44++ {
                    for fi45 := 0; fi45 < 1; fi45++ {
                        if true {
                            delete(v1, 1)
                            var v46 uint8 = 255
                            _ = v46
                        }
                        fmt.Println(v38)
                        var v47 int = fi44
                        _ = fi45
                        _ = v47
                    }
                    _ = fi44
                }
            }
            if false {
                v48 := []int{(^(v3)), v3}
                v49 := GenS4{F0: 139}
                var v50 bool = (0 != -1)
                _ = v48
                _ = v49
                _ = v50
            }
        } else {
            if (!false) {
                v2 = v2
            }
            for fi51 := 0; fi51 < 17; fi51++ {
                if false {
                    switch v3 {
                    case 128:
                        v52 := GenS4{F0: -459}
                        _ = v52
                    case -581:
                        fmt.Println(fi51, v36)
                    }
                }
                _ = fi51
            }
        }
        if false {
            if (65535 < 2147483647) {
                v53 := []int{}
                v2 = v36
                _ = v53
            } else {
                if (!((-2.25 + 1.5) >= (-2.25 + 1.5))) {
                    v54 := GenS6{F0: v2}
                    v4 -= v4
                    _ = v54
                }
                var v55 float64 = float64(float32(3.14159))
                v55 = float64(3.14159)
                _ = v55
            }
        } else {
            var v56 int16 = -875
            switch v2 {
            case "":
                var v57 float32 = ((((0.0 / 1.0) / 1.0) + -2.25) / ((0.1 / -2.25) - 0.1))
                _ = v57
            }
            if (!false) {
                var v58 bool = (0 <= 255)
                for fi59 := 0; fi59 < 14; fi59++ {
                    var v60 int64 = int64(v4)
                    v61 := GenS2{F0: v4, F1: 65535}
                    v2 = "\n\t"
                    _ = fi59
                    _ = v60
                    _ = v61
                }
                fmt.Println(v38, v3)
                _ = v58
            } else {
                v1[0] = 32767
                _ = len(v1)
            }
            _ = v56
        }
    }
    v3 -= 1
    v62 := []int{}
    if len(v62) >= 2 {
        v62 = v62[1:len(v62)]
    }
    if false {
        v63 := GenS7{F0: 165, F1: (1e10 != (1.5 - 1e10)), F2: 255}
        v4 = 42
        _ = v63
    } else {
        v62 = append(v62, v3)
        var v64 int16 = -32768
        v3 += (v3 & 434)
        _ = v64
    }
    v65 := GenS8{F0: 65535}
    if (true || false) {
        if (255 <= 965) {
            if false {
                delete(v1, 2147483647)
            }
            v3 = 128
            v38 += (v38 ^ (v38 * v38))
        }
        v4 = (v4 ^ 368)
        _ = len(v1)
    }
    if true {
        switch (v3 % ((42) | 1)) {
        case -2147483648:
            var v66 string = v2
            _ = v66
        case 9223372036854775807:
            v36 = (v36 + ("" + v36))
        case 127:
            if ((!(-1 <= 100)) && (v4 >= v4)) {
                var v67 uint32 = 0
                for fi68 := 0; fi68 < 10; fi68++ {
                    var v69 uint = 393
                    _ = fi68
                    _ = v69
                }
                _ = v67
            }
        default:
            for fi70 := 0; fi70 < 8; fi70++ {
                if (!(false || ((0 <= 18446744073709551615) && false))) {
                    for fi71 := 0; fi71 < 2; fi71++ {
                        v72 := GenS9{F0: 0}
                        _ = fi71
                        _ = v72
                    }
                } else {
                    v73 := v1[1]
                    _ = v73
                }
                var v74 float32 = (1e10 + 0.0)
                _ = fi70
                _ = v74
            }
        }
        switch (v36 + v36) {
        case "":
            var v75 int8 = -1
            _ = v75
        case "a":
            if len(v62) >= 2 {
                v62 = v62[1:len(v62)]
            }
        default:
            for fi76 := 0; fi76 < 18; fi76++ {
                v77 := GenS10{F0: v3, F1: 0, F2: 351}
                _ = fi76
                _ = v77
            }
        }
        for fi78 := 0; fi78 < 9; fi78++ {
            v2 = "\n\t"
            v79 := GenS7{F0: 0, F1: (v38 <= 208), F2: 118}
            _ = fi78
            _ = v79
        }
    }
    for fi80 := 0; fi80 < 10; fi80++ {
        if (0 <= -1) {
            if false {
                if false {
                    v4 -= int32(v3)
                    var v81 bool = (-20 < 100)
                    if ((v4 <= 2147483647) || (((v36 > (v36 + "\n\t")) || (!v81)) || true)) {
                        if (!v81) {
                            for fi82 := 0; fi82 < 16; fi82++ {
                                var v83 int = (fi82 + -27)
                                if len(v62) > 0 {
                                    v62[0] = -2147483648
                                }
                                v81 = (false || (13 >= 65535))
                                _ = fi82
                                _ = v83
                            }
                            var v84 float32 = 1.5
                            v85 := GenS9{F0: 1}
                            _ = v84
                            _ = v85
                        } else {
                            for fi86 := 0; fi86 < 16; fi86++ {
                                v2 = ("" + "a")
                                var v87 int = 127
                                v38 -= v38
                                _ = fi86
                                _ = v87
                            }
                            for fi88 := 0; fi88 < 20; fi88++ {
                                if (18446744073709551615 <= 151) {
                                    var v89 uint64 = v38
                                    v89 = v89
                                    var v90 int32 = v4
                                    _ = v89
                                    _ = v90
                                }
                                _ = fi88
                            }
                        }
                    }
                    _ = v81
                }
                var v91 uint8 = 2
                if (!((v38 / ((40) | 1)) <= v38)) {
                    v36 = ("" + v2)
                }
                _ = v91
            } else {
                if true {
                    v92 := GenS11{F0: v2, F1: 42, F2: true}
                    v38 = v38
                    _ = v92
                }
            }
            var v93 string = "foo bar"
            for fi94 := 0; fi94 < 11; fi94++ {
                v93 = "a"
                _ = fi94
            }
            _ = v93
        }
        v4 = (v4 - 1)
        var v95 uint = 471
        _ = fi80
        _ = v95
    }
    _ = len(v1)
    v1[0] = v3
    var v96 uint8 = 115
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v62:", len(v62), v62)
    fmt.Println("v96:", v96)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v38
    _ = v4
    _ = v62
    _ = v65
    _ = v96
}
