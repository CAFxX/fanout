package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32; F2 string }
type GenS2 struct { F0 bool; F1 int64 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int32; F1 bool }
type GenS5 struct { F0 uint64; F1 int32 }
type GenS6 struct { F0 uint32; F1 uint64; F2 uint16 }
type GenS7 struct { F0 int16; F1 uint8; F2 string }
func genFn0(p0 uint32) uint64 {
    for fi5 := 0; fi5 < 15; fi5++ {
        p0 -= (p0 ^ 0)
        _ = fi5
    }
    v6 := []int{4294967295}
    for fi7 := 0; fi7 < 16; fi7++ {
        if false {
            v8 := GenS1{F0: 128, F1: 42, F2: "a"}
            _ = v8
        } else {
            var v9 int8 = -128
            v9 = (-(v9))
            _ = v9
        }
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        for fi10 := 0; fi10 < 4; fi10++ {
            if (true || true) {
                var v11 uint64 = 42
                v12 := GenS2{F0: ((v11 >> ((289) & 63)) > v11), F1: 9223372036854775807}
                _ = v11
                _ = v12
            }
            p0 += (p0 * 1)
            switch ("hello" + "") {
            case "hello":
                var v13 bool = false
                _ = v13
            case "foo bar":
                switch -9223372036854775808 {
                case 32768:
                    var v14 uint = 2
                    _ = v14
                case -1:
                    _ = len(v6)
                case 42:
                    v15 := GenS3{F0: 255}
                    _ = v15
                default:
                    fmt.Println(p0, fi10)
                }
            }
            _ = fi10
        }
        _ = fi7
    }
    p0 = (p0 % (((p0 % ((2) | 1))) | 1))
    var v16 uint32 = p0
    _ = p0
    _ = v16
    _ = v6
    return 4294967295
}
func genFn1(p0 uint64, p1 int64, p2 int8) uint32 {
    p1 = (p1 % (((^(p1))) | 1))
    v17 := GenS4{F0: 875, F1: false}
    if true {
        fmt.Println(p1, p0, p1)
        if false {
            var v18 int16 = 1000
            if (0.0 < 1.5) {
                var v19 int16 = v18
                if false {
                    for fi20 := 0; fi20 < 19; fi20++ {
                        p0 = (p0 - (p0 % ((112) | 1)))
                        _ = fi20
                    }
                    for fi21 := 0; fi21 < 9; fi21++ {
                        var v22 int32 = -977
                        v23 := GenS5{F0: 0, F1: (v22 % ((v22) | 1))}
                        _ = fi21
                        _ = v22
                        _ = v23
                    }
                    fmt.Println(p1, p0, p0)
                } else {
                    v24 := GenS3{F0: 185}
                    _ = v24
                }
                if ((127 == 0) || true) {
                    p2 = 0
                    v19 -= (v19 * 1)
                } else {
                    switch "\n\t" {
                    case "foo bar":
                        v25 := GenS2{F0: true, F1: (p1 << ((4294967295) & 63))}
                        _ = v25
                    case "\n\t":
                        v26 := []int{-133, -929, 408}
                        _ = v26
                    default:
                        p1 = p1
                    }
                    v27 := []int{}
                    var v28 int = 1
                    _ = v27
                    _ = v28
                }
                _ = v19
            }
            p0 += p0
            _ = v18
        }
    }
    v29 := []int{-1, -642}
    p2 = (p2 & (p2 * (p2 | p2)))
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    _ = v29
    return 2147483648
}

func main() {
    var v1 uint8 = 1
    var v2 int = -148
    v3 := make(map[int]uint32)
    var v4 int64 = 0
    v30 := GenS5{F0: 719, F1: 42}
    fmt.Println(v4, v2, v1)
    v4 -= 559
    if false {
        var v31 uint32 = 4294967295
        for fi32 := 0; fi32 < 10; fi32++ {
            switch "hello" {
            case "foo bar":
                var v33 uint64 = 997
                _ = v33
            case "\n\t":
                var v34 uint16 = 412
                _ = v34
            case "a":
                v2 = v2
            default:
                v35 := GenS2{F0: false, F1: (^(v4))}
                _ = v35
            }
            _ = fi32
        }
        switch (v2 - (v2 | v2)) {
        case 4294967295:
            if false {
                var v36 int = 9223372036854775807
                _ = v36
            }
        case 128:
            v2 = 32767
        case 1:
            if (2 > 18446744073709551615) {
                if (!(!true)) {
                    fmt.Println(v2, v2, v2)
                    v31 -= (^(v31))
                    v1 = (v1 / (((v1 + (v1 | 127))) | 1))
                } else {
                    v2 = v2
                    delete(v3, v2)
                    v4 = v4
                }
            }
        }
        _ = v31
    }
    switch "a" {
    case "foo bar":
        v4 += 925
    default:
        v37 := GenS3{F0: v1}
        _ = v37
    }
    switch "\n\t" {
    case "":
        switch (((("a" + ("hello" + "foo bar")) + "\n\t") + ((("foo bar" + "hello") + "\n\t") + (("" + "") + "a"))) + ("foo bar" + "\n\t")) {
        case "a":
            for fi38 := 0; fi38 < 16; fi38++ {
                v1 -= v1
                for fi39 := 0; fi39 < 17; fi39++ {
                    v40 := GenS1{F0: (v1 & (^(v1))), F1: 1, F2: "foo bar"}
                    fmt.Println(fi38, v4)
                    v1 -= v1
                    if false {
                        v41 := []int{}
                        fmt.Println(v2, v2)
                        if false {
                            v2 += fi39
                            v42 := GenS1{F0: (v1 * 127), F1: 0, F2: ""}
                            switch (^(fi39)) {
                            case 32767:
                                switch "foo bar" {
                                case "hello":
                                    v43 := GenS1{F0: v1, F1: int32(v1), F2: ("\n\t" + ("a" + "foo bar"))}
                                    _ = v43
                                }
                            case -2147483648:
                                switch "a" {
                                case "":
                                    fmt.Println(fi38)
                                case "\n\t":
                                    var v44 uint64 = 42
                                    _ = v44
                                case "a":
                                    if false {
                                        var v45 float32 = (1.5 - (-2.25 * 1.5))
                                        _ = v45
                                    } else {
                                        v2 += (fi38 / (((v2 & 128)) | 1))
                                        v1 -= (v1 ^ (v1 - (v1 * (v1 * 208))))
                                    }
                                default:
                                    var v46 int = -181
                                    _ = v46
                                }
                            case 42:
                                if false {
                                    v4 = v4
                                    if ("" != "\n\t") {
                                        var v47 int16 = 353
                                        var v48 int32 = -411
                                        _ = v47
                                        _ = v48
                                    } else {
                                        v1 -= (v1 << ((42) & 63))
                                    }
                                }
                            default:
                                v49 := GenS3{F0: v1}
                                _ = v49
                            }
                            if false {
                                v50 := GenS6{F0: 65535, F1: 42, F2: 42}
                                var v51 int32 = 404
                                v52 := v3[v2]
                                var v53 float32 = 1.5
                                v41 = append(v41, fi39)
                                _ = len(v41)
                                _ = v50
                                _ = v51
                                _ = v52
                                _ = v53
                            }
                            fmt.Println(fi39, fi38)
                            _ = v42
                        } else {
                            var v54 uint64 = 555
                            if false {
                                fmt.Println(v54)
                            } else {
                                v55 := v3[(v2 | 2147483647)]
                                _ = v55
                            }
                            _ = v54
                        }
                        _ = v41
                    }
                    _ = fi39
                    _ = v40
                }
                for fi56 := 0; fi56 < 9; fi56++ {
                    switch "" {
                    case "a":
                        var v57 uint64 = 0
                        _ = v57
                    case "\n\t":
                        var v58 int64 = (v4 % (((v4 >> ((839) & 63))) | 1))
                        _ = v58
                    }
                    if (fi56 == 214) {
                        v4 = (v4 << ((65535) & 63))
                    } else {
                        v59 := []int{32767}
                        v2 = 4294967295
                        _ = v59
                    }
                    _ = fi56
                }
                _ = fi38
            }
        }
    case "\n\t":
        for fi60 := 0; fi60 < 20; fi60++ {
            v2 = (fi60 >> ((65535) & 63))
            switch ("hello" + "") {
            case "hello":
                v4 = 42
            case "foo bar":
                switch "hello" {
                case "\n\t":
                    switch v2 {
                    case 0:
                        v61 := v3[(v2 % ((911) | 1))]
                        _ = v61
                    case 370:
                        for fi62 := 0; fi62 < 2; fi62++ {
                            v63 := []int{fi60, fi62, 32768}
                            v64 := GenS1{F0: (^(v1)), F1: -2147483648, F2: "\n\t"}
                            v2 -= fi62
                            var v65 uint = 255
                            switch ("\n\t" + ("a" + "foo bar")) {
                            case "":
                                v2 = (v2 % ((fi62) | 1))
                            }
                            v4 = 1
                            _ = fi62
                            _ = v63
                            _ = v64
                            _ = v65
                        }
                    }
                case "hello":
                    fmt.Println(v4, v2, v2)
                case "foo bar":
                    switch (("a" + "\n\t") + ("\n\t" + "foo bar")) {
                    case "\n\t":
                        v66 := GenS3{F0: v1}
                        _ = v66
                    }
                default:
                    v67 := GenS7{F0: 627, F1: v1, F2: ("foo bar" + "hello")}
                    _ = v67
                }
            default:
                v4 -= -9223372036854775808
            }
            var v68 uint64 = 920
            _ = fi60
            _ = v68
        }
    default:
        if (!(v1 == v1)) {
            var v69 float32 = ((1e10 / ((((((0.0 / ((1e10 * 3.14159) / 1e10)) + (-2.25 * -2.25)) / 3.14159) - 0.0) / 0.1) / (3.14159 + 0.0))) - (3.14159 - float32((3.14159 / 0.1))))
            _ = v69
        } else {
            v70 := GenS4{F0: 1, F1: (!false)}
            _ = v70
        }
    }
    v2 -= (v2 / ((-691) | 1))
    var v71 int64 = -9223372036854775808
    v72 := v3[(v2 >> ((731) & 63))]
    if true {
        var v73 float32 = 3.14159
        v74 := GenS2{F0: false, F1: (v4 | 0)}
        switch "foo bar" {
        case "\n\t":
            v73 = (1.5 - float32(-2.25))
        case "a":
            v75 := GenS5{F0: 42, F1: 42}
            _ = v75
        }
        _ = v73
        _ = v74
    }
    v76 := v3[(v2 * 507)]
    var v77 uint32 = 17
    var v78 int64 = v4
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v71:", v71)
    fmt.Println("v72:", v72)
    fmt.Println("v76:", v76)
    fmt.Println("v77:", v77)
    fmt.Println("v78:", v78)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v71
    _ = v72
    _ = v76
    _ = v77
    _ = v78
}
