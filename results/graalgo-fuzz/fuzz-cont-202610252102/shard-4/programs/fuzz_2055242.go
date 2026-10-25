package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint64; F1 uint32 }
type GenS3 struct { F0 bool; F1 uint32; F2 int32 }
type GenS4 struct { F0 bool; F1 int16; F2 int16 }

func main() {
    v1 := []int{-808, 33, -2147483648, 32767}
    var v2 float32 = 1e10
    var v3 int32 = 1
    var v4 int8 = 0
    var v5 string = ("a" + "foo bar")
    v6 := make(map[string]int)
    v7 := GenS1{F0: v4}
    var v8 uint16 = 32767
    v6[("a" + v5)] = 0
    switch 127 {
    case 9223372036854775807:
        switch (("a" + v5) + "foo bar") {
        case "a":
            switch 32768 {
            case 349:
                fmt.Println(v4, v4, v3)
            default:
                var v9 uint = 0
                _ = v9
            }
        }
    case -1:
        v10 := GenS2{F0: 0, F1: 4294967295}
        _ = v10
    }
    v5 = (v5 + v5)
    v8 -= 65535
    for fi11 := 0; fi11 < 16; fi11++ {
        var v12 uint32 = 1
        _ = fi11
        _ = v12
    }
    _ = len(v1)
    v4 += v4
    fmt.Println(v3, v8)
    if false {
        v2 += 3.14159
        for fi13 := 0; fi13 < 8; fi13++ {
            v14 := GenS3{F0: (!(1.5 <= (1.5 / (v2 * 0.1)))), F1: 313, F2: v3}
            v15 := GenS2{F0: 2, F1: 2}
            fmt.Println(v5, v2, v2)
            _ = fi13
            _ = v14
            _ = v15
        }
    } else {
        var v16 float64 = -2.25
        var v17 float64 = 1e10
        _ = v16
        _ = v17
    }
    fmt.Println(v8, v5, v2)
    for fi18 := 0; fi18 < 5; fi18++ {
        var v19 uint32 = 893
        v1 = append(v1, (^(fi18)))
        var v20 float64 = 1e10
        _ = fi18
        _ = v19
        _ = v20
    }
    for fi21 := 0; fi21 < 6; fi21++ {
        for fi22 := 0; fi22 < 15; fi22++ {
            var v23 int32 = v3
            v24 := GenS4{F0: false, F1: 572, F2: -32768}
            if false {
                v5 = "hello"
                if (!true) {
                    v25 := v6[(((((v5 + v5) + v5) + "hello") + ("\n\t" + (v5 + v5))) + "\n\t")]
                    var v26 uint64 = 42
                    v5 = ((("" + v5) + (("hello" + v5) + (v5 + "a"))) + (v5 + (v5 + v5)))
                    if true {
                        fmt.Println(v3, fi22, v2)
                        if false {
                            if (100 > -1) {
                                if false {
                                    v6[((("foo bar" + v5) + ("foo bar" + ("foo bar" + "a"))) + ("\n\t" + v5))] = v25
                                    v2 = v2
                                    var v27 int = int(v23)
                                    var v28 uint32 = 42
                                    var v29 uint = 4294967295
                                    _ = v27
                                    _ = v28
                                    _ = v29
                                }
                                fmt.Println(v8, v4, v26)
                                switch ((v5 + v5) + "") {
                                case "\n\t":
                                    switch (v5 + v5) {
                                    case "\n\t":
                                        var v30 int32 = -492
                                        _ = v30
                                    case "foo bar":
                                        v26 = (v26 % ((v26) | 1))
                                    }
                                case "hello":
                                    fmt.Println(fi22)
                                case "":
                                    if true {
                                        v4 -= v4
                                        switch v5 {
                                        case "a":
                                            v26 = (v26 | uint64(v8))
                                            _ = len(v1)
                                        case "hello":
                                            if (!false) {
                                                v31 := GenS1{F0: 126}
                                                var v32 int16 = 32767
                                                _ = v31
                                                _ = v32
                                            }
                                        case "":
                                            var v33 int16 = 98
                                            _ = v33
                                        default:
                                            _ = len(v6)
                                        }
                                        v34 := GenS1{F0: v4}
                                        _ = v34
                                    }
                                default:
                                    fmt.Println(v26, fi21)
                                }
                            } else {
                                v5 = ""
                                if (0 == 599) {
                                    switch 32767 {
                                    case 1:
                                        var v35 int8 = v4
                                        _ = v35
                                    case -883:
                                        var v36 float32 = v2
                                        _ = v36
                                    case 0:
                                        if ((v26 / ((0) | 1)) != (v26 * 18446744073709551615)) {
                                            v8 += (^(v8))
                                            if true {
                                                v25 -= (fi22 ^ (-(fi21)))
                                                if ((v26 > v26) && (float64((v2 + 1e10)) < 0.1)) {
                                                    var v37 string = "\n\t"
                                                    _ = v37
                                                }
                                                v3 -= (v23 & -1)
                                                v38 := v6[v5]
                                                v3 -= v3
                                                var v39 uint64 = (v26 >> ((65535) & 63))
                                                var v40 uint8 = 42
                                                var v41 uint = 255
                                                _ = v38
                                                _ = v39
                                                _ = v40
                                                _ = v41
                                            }
                                            if len(v1) > 0 {
                                                v1[0] = v25
                                            }
                                        }
                                    }
                                    var v42 int64 = 0
                                    _ = v42
                                } else {
                                    v26 = v26
                                    fmt.Println(v5, v5)
                                    var v43 string = (v5 + v5)
                                    _ = v43
                                }
                                v4 -= (v4 & v4)
                                v1 = append(v1, fi22)
                            }
                            v44 := GenS4{F0: false, F1: 1, F2: 340}
                            if false {
                                var v45 string = (v5 + v5)
                                _ = v45
                            }
                            _ = v44
                        }
                    } else {
                        v5 = (v5 + (((v5 + (v5 + ((v5 + (v5 + "foo bar")) + v5))) + "foo bar") + "a"))
                        var v46 int32 = (^(v23))
                        if (999 == v8) {
                            if (1 < 536) {
                                switch "\n\t" {
                                case "":
                                    v6[v5] = (fi21 >> ((255) & 63))
                                default:
                                    if false {
                                        _ = len(v1)
                                        if len(v1) > 0 {
                                            v1[0] = (^(fi21))
                                        }
                                        fmt.Println(v23)
                                    }
                                }
                            } else {
                                v47 := GenS4{F0: (("foo bar" < v5) || true), F1: 32767, F2: -541}
                                _ = v47
                            }
                            var v48 int8 = (v4 | (v4 ^ 1))
                            fmt.Println(v46, v8, fi22)
                            _ = v48
                        }
                        if (v46 >= v46) {
                            switch v5 {
                            case "":
                                var v49 int8 = (v4 ^ (v4 >> ((4294967295) & 63)))
                                _ = v49
                            case "hello":
                                var v50 uint64 = 0
                                _ = v50
                            default:
                                if false {
                                    var v51 int8 = -128
                                    var v52 uint = 318
                                    v4 = 1
                                    _ = v51
                                    _ = v52
                                }
                            }
                        }
                        _ = v46
                    }
                    var v53 int64 = -9223372036854775808
                    fmt.Println(v2, v2)
                    if true {
                        var v54 uint = 1
                        _ = v54
                    }
                    delete(v6, ("" + (v5 + v5)))
                    fmt.Println(v26, v23)
                    _ = v25
                    _ = v26
                    _ = v53
                } else {
                    var v55 int = 128
                    _ = v55
                }
                if len(v1) > 0 {
                    v1[0] = (fi21 - fi22)
                }
            } else {
                if (!true) {
                    if true {
                        v2 = (-2.25 * 1e10)
                    }
                }
                fmt.Println(v8, v3, fi21)
                v56 := v6[v5]
                var v57 uint64 = 915
                var v58 int32 = v3
                v58 -= (v23 * v58)
                _ = v56
                _ = v57
                _ = v58
            }
            _ = fi22
            _ = v23
            _ = v24
        }
        if (v5 < "foo bar") {
            if (v3 > v3) {
                v59 := GenS4{F0: (0.0 != 0.1), F1: -532, F2: 42}
                _ = v59
            }
            var v60 uint64 = 0
            for fi61 := 0; fi61 < 18; fi61++ {
                if (2 < 346) {
                    if len(v1) > 0 {
                        v1[0] = fi61
                    }
                    var v62 uint8 = 42
                    _ = v62
                }
                if true {
                    if false {
                        v5 = v5
                    } else {
                        var v63 uint64 = v60
                        _ = len(v1)
                        _ = v63
                    }
                    if (!(v8 >= v8)) {
                        var v64 string = v5
                        if false {
                            v65 := GenS1{F0: (v4 ^ v4)}
                            switch (((v64 + ("" + v64)) + v5) + "a") {
                            case "a":
                                if (true && false) {
                                    var v66 float32 = (((v2 + v2) - (v2 - 1.5)) + v2)
                                    _ = v66
                                }
                            }
                            v67 := GenS3{F0: (false && (!true)), F1: 1, F2: (v3 / (((v3 >> ((954) & 63))) | 1))}
                            _ = v65
                            _ = v67
                        } else {
                            switch (v64 + v64) {
                            case "foo bar":
                                if (!false) {
                                    if (-1 <= 140) {
                                        fmt.Println(v64)
                                        v5 = v64
                                    }
                                    var v68 int32 = -1
                                    v4 = 42
                                    _ = v68
                                } else {
                                    v69 := GenS3{F0: ((true || (!true)) && false), F1: 2, F2: (v3 % ((v3) | 1))}
                                    if (false && true) {
                                        if false {
                                            var v70 bool = false
                                            var v71 float32 = float32(0.1)
                                            _ = v70
                                            _ = v71
                                        } else {
                                            fmt.Println(v64, v3, v2)
                                            v60 = (v60 + (v60 + (v60 & (v60 % ((v60) | 1)))))
                                        }
                                        delete(v6, ("\n\t" + "hello"))
                                        v64 = v64
                                    }
                                    _ = v69
                                }
                            case "":
                                if (!(uint8(v60) > 199)) {
                                    fmt.Println(v8)
                                    var v72 uint64 = (v60 * v60)
                                    switch v5 {
                                    case "":
                                        if (v8 <= v8) {
                                            var v73 uint = 0
                                            fmt.Println(v5, fi21, v8)
                                            fmt.Println(v72, v8)
                                            _ = v73
                                        }
                                    }
                                    _ = v72
                                } else {
                                    v64 = v64
                                    var v74 int32 = -657
                                    _ = v74
                                }
                            case "\n\t":
                                var v75 uint16 = (v8 % ((2) | 1))
                                var v76 int32 = v3
                                _ = v75
                                _ = v76
                            default:
                                var v77 float32 = v2
                                var v78 uint64 = (^(v60))
                                _ = v77
                                _ = v78
                            }
                            v3 += v3
                        }
                        _ = v64
                    } else {
                        v5 = v5
                    }
                    v79 := GenS4{F0: false, F1: 32767, F2: 32767}
                    v3 -= (v3 << ((4294967295) & 63))
                    _ = v79
                } else {
                    var v80 uint32 = 508
                    v8 = (^(v8))
                    switch (-(fi21)) {
                    case -858:
                        switch fi61 {
                        case 409:
                            switch int(v8) {
                            case -225:
                                var v81 int8 = (-(v4))
                                _ = v81
                            case -793:
                                fmt.Println(v60, v5, v5)
                            default:
                                var v82 bool = (true || true)
                                _ = v82
                            }
                        case 1:
                            var v83 uint16 = 629
                            _ = v83
                        }
                    }
                    delete(v6, v5)
                    _ = v80
                }
                _ = fi61
            }
            _ = v60
        } else {
            var v84 int8 = v4
            switch (fi21 & -2147483648) {
            case 4294967295:
                v5 = v5
            case 32767:
                var v85 uint64 = 4294967295
                _ = v85
            case 465:
                for fi86 := 0; fi86 < 20; fi86++ {
                    v3 = v3
                    v3 -= 0
                    v4 -= v4
                    _ = fi86
                }
            default:
                v87 := v6[(v5 + v5)]
                _ = v87
            }
            var v88 int64 = 786
            _ = v84
            _ = v88
        }
        _ = fi21
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
