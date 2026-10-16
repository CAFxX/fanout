package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16 }
type GenS2 struct { F0 uint64; F1 bool }
type GenS3 struct { F0 uint32; F1 string; F2 uint32 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 string; F1 string; F2 uint64 }
type GenS9 struct { F0 int64 }
type GenS10 struct { F0 int32 }
type GenS11 struct { F0 int32; F1 uint16 }
func genFn0(p0 int32, p1 int, p2 uint) uint64 {
    v6 := []int{p1, p1}
    var v7 int = p1
    fmt.Println(p2, v7, p0)
    v6 = append(v6, (-(v7)))
    if true {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        fmt.Println(p2, v7, v7)
        for fi8 := 0; fi8 < 17; fi8++ {
            for fi9 := 0; fi9 < 18; fi9++ {
                var v10 float32 = (1e10 / 1.0)
                _ = fi9
                _ = v10
            }
            if false {
                if (206 < 65535) {
                    if len(v6) > 0 {
                        v6[0] = (v7 | p1)
                    }
                    v11 := GenS1{F0: 2, F1: 42}
                    _ = v11
                } else {
                    if true {
                        _ = len(v6)
                        if true {
                            var v12 bool = (127 <= 127)
                            var v13 uint16 = 255
                            for fi14 := 0; fi14 < 20; fi14++ {
                                fmt.Println(v12, fi14)
                                var v15 float32 = 0.1
                                v16 := GenS1{F0: 636, F1: (^(v13))}
                                _ = fi14
                                _ = v15
                                _ = v16
                            }
                            _ = v12
                            _ = v13
                        } else {
                            var v17 float32 = 0.1
                            _ = v17
                        }
                        switch int(p2) {
                        case -167:
                            var v18 uint16 = 0
                            _ = v18
                        case -231:
                            var v19 float32 = (1e10 * 3.14159)
                            _ = v19
                        case 4294967295:
                            v20 := GenS1{F0: 404, F1: 42}
                            _ = v20
                        default:
                            for fi21 := 0; fi21 < 3; fi21++ {
                                p1 -= v7
                                v6 = append(v6, v7)
                                if false {
                                    p2 -= p2
                                } else {
                                    var v22 int = (-(p1))
                                    _ = v22
                                }
                                var v23 int16 = 354
                                if ((1 <= 18446744073709551615) && ((("hello" + ("foo bar" + ("" + "foo bar"))) + "foo bar") == ((("\n\t" + ("foo bar" + "foo bar")) + "\n\t") + "a"))) {
                                    fmt.Println(fi21)
                                    var v24 uint64 = 4294967295
                                    var v25 int8 = 125
                                    _ = v24
                                    _ = v25
                                }
                                switch "hello" {
                                case "foo bar":
                                    if true {
                                        var v26 uint8 = 146
                                        _ = v26
                                    }
                                case "a":
                                    if ("a" < "hello") {
                                        v27 := GenS2{F0: 98, F1: true}
                                        if true {
                                            if false {
                                                v28 := GenS3{F0: 2, F1: ("hello" + "\n\t"), F2: 715}
                                                var v29 uint64 = 151
                                                _ = v28
                                                _ = v29
                                            } else {
                                                v7 = v7
                                                if false {
                                                    var v30 uint8 = 2
                                                    p1 += 2147483647
                                                    p0 += (p0 | (p0 & (p0 ^ (p0 + 8))))
                                                    _ = v30
                                                }
                                                var v31 uint16 = 42
                                                _ = v31
                                            }
                                            p1 = v7
                                        } else {
                                            var v32 uint16 = 255
                                            v6 = append(v6, (fi21 - 128))
                                            _ = v32
                                        }
                                        fmt.Println(p0, p2, v23)
                                        switch "hello" {
                                        case "a":
                                            v33 := GenS4{F0: int8(v23)}
                                            _ = v33
                                        case "":
                                            p0 = p0
                                        case "hello":
                                            var v34 uint32 = 4294967295
                                            _ = v34
                                        }
                                        _ = v27
                                    } else {
                                        v7 -= (v7 << ((541) & 63))
                                        switch "foo bar" {
                                        case "\n\t":
                                            var v35 int32 = p0
                                            var v36 uint64 = 18446744073709551615
                                            fmt.Println(p2)
                                            _ = v35
                                            _ = v36
                                        case "":
                                            fmt.Println(fi21, p0)
                                        case "foo bar":
                                            p0 += (p0 * 0)
                                        default:
                                            switch "foo bar" {
                                            case "hello":
                                                _ = len(v6)
                                            case "\n\t":
                                                v37 := GenS2{F0: 924, F1: false}
                                                _ = v37
                                            case "":
                                                v6 = append(v6, p1)
                                            }
                                        }
                                    }
                                }
                                v38 := GenS5{F0: 127}
                                _ = fi21
                                _ = v23
                                _ = v38
                            }
                        }
                    } else {
                        v39 := GenS6{F0: ("a" + ("foo bar" + "\n\t"))}
                        _ = v39
                    }
                }
                p0 -= p0
            }
            if len(v6) > 0 {
                v6[0] = 1
            }
            _ = fi8
        }
    } else {
        fmt.Println(p1, p0, v7)
        p0 -= p0
        var v40 int = p1
        _ = v40
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 0
}

func main() {
    v1 := make(map[int]int)
    v2 := []string{"", "", ("\n\t" + ("hello" + "a"))}
    v3 := make(map[string]uint32)
    v4 := []uint32{1, 2147483648}
    var v5 float64 = ((-2.25 - (0.1 * (0.1 - 0.1))) * float64(1e10))
    for fi41 := 0; fi41 < 14; fi41++ {
        fmt.Println(fi41, v5)
        switch (fi41 - 136) {
        case 9223372036854775807:
            v5 = float64((-2.25 + (1e10 - ((0.1 * (3.14159 - (0.0 * (-2.25 - ((0.0 + ((((1e10 + float32(1e10)) + (-2.25 / 3.14159)) + -2.25) * (-2.25 - (((3.14159 + 0.0) + float32((float64((-2.25 - ((1.5 / 3.14159) * 1.5))) * ((v5 - -2.25) * v5)))) + 1e10)))) - ((0.0 * ((1.5 / (float32(0.1) - (((float32(v5) + 0.0) - float32((-2.25 * -2.25))) - -2.25))) * ((0.0 - float32(((1.5 - v5) / (v5 / (float64((-2.25 - float32(v5))) - (v5 / 1.5)))))) - 1e10))) / -2.25)))))) - 1.5))))
        case 128:
            if (-1 < 1) {
                v42 := GenS7{F0: 186}
                var v43 float32 = 0.0
                _ = v42
                _ = v43
            } else {
                var v44 uint16 = 255
                switch -2147483648 {
                case 809:
                    v44 -= 180
                default:
                    v5 += (v5 - v5)
                }
                var v45 string = "\n\t"
                _ = v44
                _ = v45
            }
        case 127:
            for fi46 := 0; fi46 < 2; fi46++ {
                if false {
                    v5 -= float64(0.0)
                    switch (fi46 % (((fi46 >> ((18446744073709551615) & 63))) | 1)) {
                    case 32768:
                        var v47 int8 = 0
                        _ = v47
                    case 2147483647:
                        var v48 string = "a"
                        _ = v48
                    case 32767:
                        var v49 string = "hello"
                        _ = v49
                    }
                    if false {
                        if true {
                            fmt.Println(v5)
                            switch (fi46 ^ (fi41 % ((-739) | 1))) {
                            case -1:
                                switch fi41 {
                                case -840:
                                    v50 := GenS3{F0: 2147483647, F1: "\n\t", F2: 1}
                                    _ = v50
                                case 42:
                                    v51 := GenS8{F0: "", F1: "", F2: 0}
                                    _ = v51
                                default:
                                    switch ("a" + "hello") {
                                    case "":
                                        switch ("hello" + "") {
                                        case "\n\t":
                                            switch (fi41 << ((727) & 63)) {
                                            case -672:
                                                var v52 int64 = 473
                                                v5 += (0.1 / (1.5 * v5))
                                                _ = v52
                                            case 32767:
                                                switch fi41 {
                                                case 32767:
                                                    v53 := GenS9{F0: 413}
                                                    _ = v53
                                                case 127:
                                                    v5 += v5
                                                case 717:
                                                    v4 = append(v4, 1)
                                                }
                                            default:
                                                fmt.Println(v5)
                                            }
                                        case "hello":
                                            var v54 uint32 = 4294967295
                                            fmt.Println(fi41, v54)
                                            var v55 float64 = 0.1
                                            _ = v54
                                            _ = v55
                                        case "foo bar":
                                            if true {
                                                var v56 int = 4294967295
                                                v57 := GenS9{F0: -1}
                                                var v58 uint = 255
                                                _ = v56
                                                _ = v57
                                                _ = v58
                                            }
                                        default:
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                            var v59 string = "hello"
                                            _ = v59
                                        }
                                    case "a":
                                        var v60 int64 = 9223372036854775807
                                        var v61 int = (fi46 & 240)
                                        var v62 uint = 4294967295
                                        _ = v60
                                        _ = v61
                                        _ = v62
                                    }
                                }
                            case 32767:
                                var v63 int16 = 1000
                                _ = v63
                            default:
                                v5 = ((((v5 * v5) - ((v5 * (3.14159 - v5)) * 1e10)) + (v5 - 0.0)) * v5)
                            }
                        } else {
                            switch ("\n\t" + ("a" + "hello")) {
                            case "hello":
                                if (((!(!false)) || (975 > -971)) && (1 < 13)) {
                                    var v64 string = "hello"
                                    var v65 uint64 = 321
                                    var v66 bool = true
                                    _ = v64
                                    _ = v65
                                    _ = v66
                                }
                            case "a":
                                var v67 int32 = 917
                                _ = v67
                            case "":
                                var v68 int = (fi46 / (((fi41 << ((0) & 63))) | 1))
                                _ = v68
                            }
                            v5 = 1.5
                            fmt.Println(v5, v5)
                        }
                    } else {
                        var v69 uint16 = 32767
                        _ = v69
                    }
                    v5 = 0.1
                } else {
                    v5 -= float64(1e10)
                }
                var v70 uint32 = 2147483647
                var v71 uint = 747
                var v72 float64 = v5
                _ = fi46
                _ = v70
                _ = v71
                _ = v72
            }
        default:
            var v73 string = "\n\t"
            _ = v73
        }
        if true {
            var v74 int8 = -102
            v5 = 0.0
            var v75 float32 = 0.1
            _ = v74
            _ = v75
        } else {
            v5 -= v5
            var v76 uint64 = genFn0(-1, 4294967295, 0)
            _ = v76
        }
        _ = fi41
    }
    if true {
        v5 = float64(-2.25)
        v2 = append(v2, ("foo bar" + "\n\t"))
        if false {
            var v77 uint64 = 42
            v78 := GenS1{F0: 611, F1: 345}
            _ = v77
            _ = v78
        }
    } else {
        var v79 string = ("" + "hello")
        var v80 float32 = 1e10
        _ = v79
        _ = v80
    }
    fmt.Println(v5)
    var v81 int64 = 1
    for fi82 := 0; fi82 < 15; fi82++ {
        var v83 string = "a"
        switch "\n\t" {
        case "":
            _ = len(v2)
        case "foo bar":
            for fi84 := 0; fi84 < 6; fi84++ {
                var v85 int = fi84
                v85 -= (fi82 >> ((18446744073709551615) & 63))
                if true {
                    if (true || (fi84 >= fi84)) {
                        v5 = v5
                        if true {
                            switch (fi84 / (((-(fi84))) | 1)) {
                            case 32767:
                                v83 = v83
                                var v86 uint32 = 747
                                _ = v86
                            case 0:
                                _ = len(v1)
                            }
                        }
                        switch -2147483648 {
                        case -9223372036854775808:
                            _ = len(v3)
                        case 588:
                            v87 := GenS10{F0: -159}
                            _ = v87
                        default:
                            _ = len(v4)
                        }
                        fmt.Println(v85, v85)
                    }
                    var v88 int16 = 290
                    if false {
                        fmt.Println(v85, v5)
                        v88 = v88
                    }
                    _ = v88
                }
                _ = fi84
                _ = v85
            }
        case "hello":
            v89 := GenS11{F0: 0, F1: 255}
            _ = v89
        }
        fmt.Println(v83, v81)
        _ = fi82
        _ = v83
    }
    v81 = 395
    var v90 int16 = -32768
    v91 := GenS8{F0: "\n\t", F1: "", F2: 4294967295}
    v90 += (-(v90))
    v81 = v81
    switch -324 {
    case -1:
        if true {
            _ = len(v4)
            var v92 int16 = (v90 & 1)
            for fi93 := 0; fi93 < 18; fi93++ {
                var v94 int64 = v81
                _ = fi93
                _ = v94
            }
            _ = v92
        }
    case 32768:
        _ = len(v4)
    }
    _ = len(v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v81:", v81)
    fmt.Println("v90:", v90)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v81
    _ = v90
    _ = v91
}
