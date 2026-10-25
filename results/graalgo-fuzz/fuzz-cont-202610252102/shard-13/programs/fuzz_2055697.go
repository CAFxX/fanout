package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32; F2 uint16 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int32; F1 uint8; F2 int64 }
type GenS4 struct { F0 uint32; F1 int; F2 int64 }
type GenS5 struct { F0 int64; F1 int16 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 int32; F1 uint32; F2 uint64 }
type GenS10 struct { F0 int; F1 uint64 }
type GenS11 struct { F0 int; F1 int16; F2 int16 }
type GenS12 struct { F0 int; F1 uint8; F2 int16 }
func genFn0(p0 bool, p1 int16) int32 {
    fmt.Println(p1)
    v5 := GenS1{F0: ("\n\t" + "a"), F1: -2147483648, F2: 32767}
    for fi6 := 0; fi6 < 10; fi6++ {
        var v7 uint32 = 42
        v8 := GenS2{F0: 1}
        _ = fi6
        _ = v7
        _ = v8
    }
    var v9 uint8 = 2
    for fi10 := 0; fi10 < 7; fi10++ {
        switch (("\n\t" + "hello") + "foo bar") {
        case "":
            if (4294967295 <= 0) {
                switch ("a" + "foo bar") {
                case "":
                    fmt.Println(v9, v9)
                case "\n\t":
                    if p0 {
                        var v11 string = "foo bar"
                        v9 += v9
                        v12 := GenS3{F0: 1, F1: (v9 >> ((18446744073709551615) & 63)), F2: -912}
                        _ = v11
                        _ = v12
                    } else {
                        p0 = p0
                    }
                default:
                    for fi13 := 0; fi13 < 9; fi13++ {
                        var v14 int8 = 1
                        v14 = v14
                        var v15 bool = false
                        _ = fi13
                        _ = v14
                        _ = v15
                    }
                }
                var v16 string = "a"
                for fi17 := 0; fi17 < 1; fi17++ {
                    var v18 string = (v16 + v16)
                    _ = fi17
                    _ = v18
                }
                _ = v16
            }
        case "foo bar":
            switch fi10 {
            case -183:
                p1 = -1
            case -1:
                v19 := GenS2{F0: 42}
                _ = v19
            case 127:
                for fi20 := 0; fi20 < 2; fi20++ {
                    v9 -= (^(v9))
                    _ = fi20
                }
            }
        case "\n\t":
            var v21 uint64 = 42
            _ = v21
        }
        v22 := []int{fi10, fi10, fi10}
        if (true || (0 != 1)) {
            p1 = (^(p1))
        } else {
            if true {
                v23 := GenS2{F0: 42}
                if (2 >= 2) {
                    var v24 uint64 = 42
                    var v25 int32 = 0
                    _ = v24
                    _ = v25
                } else {
                    v9 += 0
                    v26 := GenS4{F0: 0, F1: fi10, F2: 42}
                    for fi27 := 0; fi27 < 3; fi27++ {
                        if (false || p0) {
                            v22 = append(v22, -510)
                            var v28 uint8 = v9
                            if false {
                                p1 = p1
                                if (822 < 495) {
                                    fmt.Println(p0, p1, v28)
                                    switch "\n\t" {
                                    case "\n\t":
                                        v29 := GenS5{F0: 1, F1: p1}
                                        _ = v29
                                    case "":
                                        fmt.Println(p1)
                                    default:
                                        p0 = true
                                    }
                                    var v30 bool = true
                                    _ = v30
                                } else {
                                    var v31 int = fi10
                                    var v32 float64 = 1.5
                                    var v33 int16 = (p1 >> ((0) & 63))
                                    if ((!(!(((0.1 * float32(v32)) < float32(1.5)) || (1 < 0)))) || (!((!p0) && (1 != 0)))) {
                                        fmt.Println(v32)
                                    } else {
                                        v34 := GenS1{F0: "", F1: -39, F2: 151}
                                        _ = v34
                                    }
                                    var v35 uint = 42
                                    _ = v31
                                    _ = v32
                                    _ = v33
                                    _ = v35
                                }
                                if (p0 || (255 == 255)) {
                                    _ = len(v22)
                                }
                            } else {
                                var v36 bool = p0
                                var v37 int8 = 100
                                _ = v36
                                _ = v37
                            }
                            if len(v22) >= 2 {
                                v22 = v22[1:len(v22)]
                            }
                            _ = v28
                        }
                        if p0 {
                            v38 := GenS4{F0: 0, F1: int(p1), F2: 1}
                            var v39 uint64 = 681
                            v22 = append(v22, -676)
                            switch ("\n\t" + "foo bar") {
                            case "foo bar":
                                v40 := GenS2{F0: 42}
                                _ = v40
                            case "a":
                                var v41 int64 = 475
                                _ = v41
                            }
                            _ = v38
                            _ = v39
                        }
                        _ = fi27
                    }
                    _ = v26
                }
                p1 = p1
                _ = v23
            }
        }
        _ = fi10
        _ = v22
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v9
    return 945
}

func main() {
    var v1 uint16 = 0
    v2 := []int{-896, 32768, 144}
    v3 := make(map[string]uint32)
    var v4 int64 = 89
    if (3.14159 == float32(((-2.25 / 1.0) / 1.5))) {
        if (!false) {
            if ((v1 % ((v1) | 1)) <= (v1 & v1)) {
                v42 := GenS6{F0: (v1 * v1)}
                v43 := GenS2{F0: v4}
                v4 -= (v4 & (v4 - (v4 - (v4 ^ (v4 ^ (v4 ^ v4))))))
                _ = v42
                _ = v43
            }
            var v44 uint = uint(v1)
            switch 32767 {
            case 127:
                v45 := GenS7{F0: -1}
                _ = v45
            default:
                v3[("a" + "hello")] = 65535
            }
            _ = v44
        }
        v46 := v3[((("" + "hello") + "a") + "a")]
        _ = v46
    } else {
        var v47 uint = 4294967295
        _ = len(v3)
        var v48 int = 0
        _ = v47
        _ = v48
    }
    switch 4294967295 {
    case 666:
        v49 := v3[(("a" + "") + ("foo bar" + "foo bar"))]
        _ = v49
    case -554:
        var v50 float32 = -2.25
        _ = v50
    default:
        var v51 int16 = -1
        _ = v51
    }
    v52 := GenS8{F0: 2}
    if (-2147483648 <= 267) {
        switch 1 {
        case -923:
            v53 := GenS3{F0: 0, F1: 126, F2: v4}
            _ = v53
        case 32768:
            var v54 uint64 = 0
            _ = v54
        default:
            v55 := v3["hello"]
            _ = v55
        }
    }
    if true {
        if (true || true) {
            fmt.Println(v4)
            for fi56 := 0; fi56 < 9; fi56++ {
                if (264 >= 42) {
                    _ = len(v3)
                } else {
                    v4 = v4
                }
                var v57 bool = false
                var v58 int16 = 1
                _ = fi56
                _ = v57
                _ = v58
            }
            v1 += v1
        }
    } else {
        v4 = (-(v4))
        var v59 int64 = (-(v4))
        v60 := GenS9{F0: 10, F1: 42, F2: 18446744073709551615}
        _ = v59
        _ = v60
    }
    for fi61 := 0; fi61 < 2; fi61++ {
        _ = len(v2)
        for fi62 := 0; fi62 < 18; fi62++ {
            switch 32768 {
            case 127:
                v4 += (v4 | (v4 + v4))
            }
            v63 := v3[("hello" + "a")]
            v4 -= 794
            switch ((((("a" + "hello") + ("hello" + "hello")) + (("foo bar" + "") + "foo bar")) + (("foo bar" + "foo bar") + "\n\t")) + "foo bar") {
            case "\n\t":
                v4 = (v4 % (((-(v4))) | 1))
            case "hello":
                var v64 uint8 = 42
                _ = v64
            case "a":
                v63 = (v63 - (v63 | (v63 / (((v63 * (^(v63)))) | 1))))
            default:
                v3[(((("a" + "foo bar") + ("" + "\n\t")) + "\n\t") + "hello")] = (v63 % ((2147483648) | 1))
            }
            _ = fi62
            _ = v63
        }
        _ = len(v2)
        _ = fi61
    }
    v1 -= v1
    v4 -= v4
    v1 += v1
    for fi65 := 0; fi65 < 10; fi65++ {
        if (false || (false && (!(127 >= -1)))) {
            for fi66 := 0; fi66 < 18; fi66++ {
                switch "a" {
                case "a":
                    if len(v2) > 0 {
                        v2[0] = (fi65 & fi65)
                    }
                case "foo bar":
                    v67 := GenS9{F0: -2147483648, F1: 2147483648, F2: 42}
                    _ = v67
                case "":
                    v1 = v1
                default:
                    _ = len(v2)
                }
                if true {
                    v4 -= v4
                    switch ("\n\t" + "hello") {
                    case "\n\t":
                        var v68 uint64 = 0
                        _ = v68
                    default:
                        switch fi65 {
                        case -279:
                            if (!((1.5 - 0.1) == (3.14159 - 1.5))) {
                                if (!false) {
                                    var v69 uint32 = uint32(v1)
                                    v3["a"] = 2
                                    switch "hello" {
                                    case "foo bar":
                                        switch (fi65 % ((745) | 1)) {
                                        case 127:
                                            var v70 int32 = genFn0(true, 42)
                                            _ = v70
                                        case -105:
                                            v69 = 65535
                                        }
                                    default:
                                        v1 = 0
                                    }
                                    _ = v69
                                } else {
                                    v71 := GenS10{F0: 2147483647, F1: 42}
                                    if false {
                                        if (-2147483648 >= 1) {
                                            if (false && (42 <= 255)) {
                                                if false {
                                                    delete(v3, "")
                                                    fmt.Println(fi66, fi65)
                                                    fmt.Println(fi66, fi65, fi66)
                                                    fmt.Println(v1, v1, fi65)
                                                    var v72 uint8 = 255
                                                    _ = v72
                                                } else {
                                                    var v73 int8 = 93
                                                    v1 -= 42
                                                    var v74 int = 576
                                                    var v75 uint32 = 2
                                                    _ = v73
                                                    _ = v74
                                                    _ = v75
                                                }
                                                v76 := GenS8{F0: 1}
                                                switch "a" {
                                                case "a":
                                                    fmt.Println(v1, fi65)
                                                case "":
                                                    if false {
                                                        v4 = (v4 - v4)
                                                    }
                                                }
                                                _ = v76
                                            }
                                            v77 := GenS6{F0: v1}
                                            v78 := GenS2{F0: -9223372036854775808}
                                            switch ("hello" + "a") {
                                            case "foo bar":
                                                fmt.Println(fi66)
                                            }
                                            _ = v77
                                            _ = v78
                                        }
                                        v79 := GenS1{F0: "\n\t", F1: 2147483647, F2: (v1 / ((v1) | 1))}
                                        fmt.Println(fi65, v1, v1)
                                        _ = v79
                                    } else {
                                        v4 = v4
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        v80 := GenS11{F0: fi65, F1: 32767, F2: -711}
                                        var v81 uint = 4294967295
                                        var v82 string = "\n\t"
                                        _ = v80
                                        _ = v81
                                        _ = v82
                                    }
                                    _ = v71
                                }
                                fmt.Println(fi65, v1, fi65)
                                switch (("foo bar" + "") + ("hello" + "a")) {
                                case "foo bar":
                                    var v83 int8 = -1
                                    _ = v83
                                default:
                                    var v84 uint32 = 4294967295
                                    _ = v84
                                }
                            } else {
                                if false {
                                    v85 := GenS11{F0: fi66, F1: -1, F2: 32767}
                                    v4 += (v4 >> ((310) & 63))
                                    _ = v85
                                } else {
                                    v4 = v4
                                }
                                fmt.Println(fi65, fi65, fi65)
                                v86 := GenS12{F0: fi66, F1: 104, F2: 1000}
                                _ = v86
                            }
                        case 0:
                            if len(v2) > 0 {
                                v2[0] = (fi65 << ((42) & 63))
                            }
                        case -934:
                            var v87 uint8 = 127
                            _ = v87
                        }
                        v1 = (v1 | 674)
                        var v88 int32 = 849
                        _ = v88
                    }
                    var v89 int64 = v4
                    _ = v89
                }
                fmt.Println(fi65, v1, v1)
                _ = fi66
            }
            switch (^(fi65)) {
            case -808:
                if false {
                    v90 := v3["hello"]
                    var v91 int = -660
                    _ = v90
                    _ = v91
                }
            case -2147483648:
                v1 -= v1
            case 127:
                fmt.Println(v1, v4, v1)
            }
        }
        v1 = v1
        v4 -= v4
        _ = fi65
    }
    fmt.Println(v4, v4)
    v4 = v4
    if false {
        if true {
            for fi92 := 0; fi92 < 9; fi92++ {
                fmt.Println(v1)
                v1 -= (v1 | v1)
                _ = fi92
            }
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        v1 = (v1 / ((32767) | 1))
    } else {
        var v93 int16 = -398
        _ = v93
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v94 string = "\n\t"
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v94:", v94)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v52
    _ = v94
}
