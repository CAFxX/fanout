package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 bool; F1 uint8 }
type GenS6 struct { F0 uint; F1 uint32; F2 uint32 }
type GenS7 struct { F0 int8; F1 int32; F2 uint32 }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[int]int)
    var v3 uint64 = 18446744073709551615
    var v4 int = 4294967295
    switch v4 {
    case 2147483647:
        v5 := GenS1{F0: -32768}
        _ = v5
    case 1:
        for fi6 := 0; fi6 < 8; fi6++ {
            if true {
                fmt.Println(v3, v4, v3)
            } else {
                for fi7 := 0; fi7 < 5; fi7++ {
                    v4 -= (v4 >> ((757) & 63))
                    _ = fi7
                }
            }
            if true {
                var v8 uint16 = 399
                _ = v8
            } else {
                var v9 int32 = -2147483648
                if true {
                    if (!(((v3 & 183) >= v3) && false)) {
                        v1[2] = 2147483647
                        var v10 int32 = (v9 | v9)
                        v9 -= (v9 - v10)
                        _ = v10
                    }
                    var v11 uint32 = 842
                    if false {
                        _ = len(v2)
                        fmt.Println(v9, v11)
                    }
                    _ = v11
                } else {
                    v12 := GenS1{F0: 32767}
                    v13 := GenS1{F0: 1000}
                    _ = v12
                    _ = v13
                }
                _ = v9
            }
            var v14 int64 = 199
            _ = fi6
            _ = v14
        }
    }
    if (!(2147483648 > 1)) {
        var v15 int8 = 100
        if (-1 == 42) {
            if false {
                fmt.Println(v4)
            } else {
                switch -9223372036854775808 {
                case -39:
                    if false {
                        v3 += (v3 & v3)
                    }
                default:
                    switch "foo bar" {
                    case "\n\t":
                        if (v15 != (v15 << ((18446744073709551615) & 63))) {
                            for fi16 := 0; fi16 < 6; fi16++ {
                                fmt.Println(v3, v4, v4)
                                var v17 float32 = (1.5 * (3.14159 * (-2.25 * 0.1)))
                                fmt.Println(v17, v3, v3)
                                _ = fi16
                                _ = v17
                            }
                            for fi18 := 0; fi18 < 18; fi18++ {
                                v19 := GenS1{F0: -705}
                                v3 = (v3 >> ((42) & 63))
                                _ = fi18
                                _ = v19
                            }
                            var v20 bool = (9223372036854775807 > (v4 / (((v4 << ((528) & 63))) | 1)))
                            _ = v20
                        }
                    }
                }
                var v21 uint8 = 1
                _ = v21
            }
            var v22 int16 = -508
            for fi23 := 0; fi23 < 20; fi23++ {
                var v24 float64 = 1.5
                if (65535 < 255) {
                    var v25 float32 = 1e10
                    fmt.Println(v15, fi23)
                    _ = v25
                } else {
                    for fi26 := 0; fi26 < 8; fi26++ {
                        var v27 uint32 = 65535
                        v27 -= v27
                        var v28 string = "hello"
                        _ = fi26
                        _ = v27
                        _ = v28
                    }
                    var v29 uint32 = 2147483647
                    _ = v29
                }
                switch (("a" + "foo bar") + "\n\t") {
                case "":
                    switch "\n\t" {
                    case "a":
                        var v30 uint = 18446744073709551615
                        _ = v30
                    }
                case "a":
                    v31 := GenS1{F0: v22}
                    _ = v31
                }
                _ = fi23
                _ = v24
            }
            _ = v22
        }
        v15 -= 42
        _ = v15
    } else {
        var v32 bool = ((0.0 < float64(3.14159)) && false)
        _ = len(v2)
        _ = v32
    }
    if (true || false) {
        for fi33 := 0; fi33 < 18; fi33++ {
            for fi34 := 0; fi34 < 10; fi34++ {
                var v35 int16 = 32767
                _ = fi34
                _ = v35
            }
            for fi36 := 0; fi36 < 15; fi36++ {
                var v37 bool = true
                var v38 int16 = 1000
                v39 := GenS1{F0: v38}
                v40 := GenS1{F0: v38}
                _ = fi36
                _ = v37
                _ = v38
                _ = v39
                _ = v40
            }
            if false {
                v4 = fi33
                if ((0.1 - -2.25) < -2.25) {
                    v4 += int(v3)
                }
            } else {
                fmt.Println(v3, v3, fi33)
                for fi41 := 0; fi41 < 12; fi41++ {
                    if true {
                        fmt.Println(fi33, v4, v3)
                        v4 += (^(fi41))
                        v42 := []int{fi41, fi41}
                        _ = v42
                    }
                    _ = fi41
                }
            }
            _ = fi33
        }
        var v43 float32 = 1e10
        _ = v43
    } else {
        v4 += (v4 | v4)
        var v44 uint = 18446744073709551615
        _ = v44
    }
    for fi45 := 0; fi45 < 1; fi45++ {
        v4 = 718
        fmt.Println(fi45, v4)
        _ = fi45
    }
    v4 += v4
    switch 1 {
    case 1:
        if true {
            fmt.Println(v3, v3)
            if (284 == 65535) {
                if true {
                    var v46 int64 = 312
                    v47 := GenS1{F0: 919}
                    _ = v46
                    _ = v47
                }
                var v48 int16 = 409
                v3 += v3
                _ = v48
            } else {
                var v49 int = v4
                if true {
                    fmt.Println(v49, v49, v3)
                } else {
                    if (!true) {
                        var v50 uint8 = 204
                        _ = v50
                    }
                    v51 := GenS1{F0: -202}
                    v52 := GenS2{F0: 888}
                    _ = v51
                    _ = v52
                }
                if (558 > 1) {
                    v49 -= (-(v49))
                    v53 := GenS3{F0: (v3 - 4294967295)}
                    for fi54 := 0; fi54 < 17; fi54++ {
                        v3 = (v3 * v3)
                        var v55 uint16 = 2
                        fmt.Println(v4, v49, v4)
                        _ = fi54
                        _ = v55
                    }
                    _ = v53
                }
                _ = v49
            }
            v4 -= v4
        } else {
            v3 = (v3 - v3)
            v3 = (v3 | v3)
            if (2 >= 31) {
                for fi56 := 0; fi56 < 6; fi56++ {
                    var v57 int16 = 1000
                    _ = fi56
                    _ = v57
                }
                var v58 uint8 = 25
                _ = v58
            } else {
                switch "\n\t" {
                case "a":
                    v59 := GenS2{F0: v4}
                    _ = v59
                default:
                    var v60 uint8 = 255
                    _ = v60
                }
                var v61 int64 = -1
                var v62 int32 = 42
                _ = v61
                _ = v62
            }
        }
    case 142:
        fmt.Println(v3)
    }
    if (true || (42 < 4294967295)) {
        v63 := GenS4{F0: true}
        _ = v63
    } else {
        v3 += (v3 >> ((0) & 63))
        if false {
            var v64 int16 = 42
            var v65 int64 = -9223372036854775808
            v64 += v64
            _ = v64
            _ = v65
        } else {
            for fi66 := 0; fi66 < 19; fi66++ {
                if (false && (false || (442 == 4294967295))) {
                    v67 := GenS4{F0: false}
                    _ = v67
                } else {
                    v2[(fi66 >> ((2) & 63))] = 218
                }
                _ = fi66
            }
            v4 -= (v4 / ((v4) | 1))
            var v68 float64 = 0.0
            _ = v68
        }
    }
    v3 -= v3
    if (578 > -1) {
        fmt.Println(v4)
    } else {
        if true {
            v3 = 1
            _ = len(v2)
        } else {
            if (v3 > v3) {
                if false {
                    switch "a" {
                    case "foo bar":
                        switch "\n\t" {
                        case "foo bar":
                            var v69 uint16 = 42
                            _ = v69
                        case "\n\t":
                            for fi70 := 0; fi70 < 7; fi70++ {
                                v71 := []int{9223372036854775807, (-(v4))}
                                v4 = (v4 - fi70)
                                _ = fi70
                                _ = v71
                            }
                        }
                    case "\n\t":
                        v3 += 0
                    case "":
                        v3 += 1
                    default:
                        for fi72 := 0; fi72 < 7; fi72++ {
                            v73 := GenS5{F0: (9223372036854775807 == -391), F1: 255}
                            var v74 uint = 886
                            _ = fi72
                            _ = v73
                            _ = v74
                        }
                    }
                    fmt.Println(v4, v3)
                    if (42 != 618) {
                        switch (-(v4)) {
                        case 1:
                            v4 = (v4 ^ 0)
                        case -422:
                            v75 := []int{v4, -224, v4}
                            _ = v75
                        case 4294967295:
                            fmt.Println(v3)
                        }
                        var v76 int64 = -1
                        for fi77 := 0; fi77 < 4; fi77++ {
                            var v78 int8 = 127
                            v79 := GenS5{F0: (true && true), F1: 165}
                            _ = fi77
                            _ = v78
                            _ = v79
                        }
                        _ = v76
                    } else {
                        v80 := GenS1{F0: 275}
                        _ = v80
                    }
                }
                fmt.Println(v4)
                var v81 uint8 = 1
                _ = v81
            } else {
                if true {
                    v82 := GenS2{F0: (v4 ^ (v4 % ((v4) | 1)))}
                    v3 -= (v3 * (^(v3)))
                    _ = v82
                } else {
                    v4 -= (v4 % ((v4) | 1))
                }
            }
        }
        v83 := GenS2{F0: (v4 % (((v4 >> ((4294967295) & 63))) | 1))}
        fmt.Println(v3, v4)
        _ = v83
    }
    v84 := []int{(-(v4)), 603}
    for fi85 := 0; fi85 < 17; fi85++ {
        for fi86 := 0; fi86 < 2; fi86++ {
            v1[42] = 4294967295
            switch "foo bar" {
            case "hello":
                v3 -= v3
            case "a":
                if true {
                    var v87 bool = (!true)
                    _ = v87
                } else {
                    v88 := GenS2{F0: v4}
                    v89 := GenS4{F0: (!false)}
                    _ = v88
                    _ = v89
                }
                v3 = (^(v3))
            case "":
                var v90 uint64 = (^(v3))
                _ = v90
            }
            switch "\n\t" {
            case "":
                v4 = v4
                fmt.Println(fi85)
            default:
                var v91 uint32 = 1
                _ = v91
            }
            _ = fi86
        }
        _ = len(v1)
        _ = fi85
    }
    fmt.Println(v4)
    v92 := GenS6{F0: 255, F1: 1, F2: 2147483648}
    if len(v84) > 0 {
        v84[0] = -2147483648
    }
    var v93 int16 = -250
    v94 := GenS7{F0: 1, F1: -2147483648, F2: 2147483647}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v84:", len(v84), v84)
    fmt.Println("v93:", v93)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v84
    _ = v92
    _ = v93
    _ = v94
}
