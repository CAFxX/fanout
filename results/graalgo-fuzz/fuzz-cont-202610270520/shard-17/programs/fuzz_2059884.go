package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16; F2 int32 }
type GenS2 struct { F0 int; F1 int64; F2 uint }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint8; F1 string }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint32; F1 uint64; F2 uint16 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 int16; F1 uint; F2 int16 }
type GenS9 struct { F0 int8 }
type GenS10 struct { F0 uint16 }
type GenS11 struct { F0 string; F1 uint64 }
type GenS12 struct { F0 uint64 }
func genFn0() uint32 {
    var v7 uint32 = 826
    v8 := GenS1{F0: -2147483648, F1: 255, F2: -425}
    v9 := GenS2{F0: 128, F1: 42, F2: 65535}
    _ = v7
    _ = v8
    _ = v9
    return (v7 | (v7 & 79))
}
func genFn1(p0 uint16) int8 {
    if true {
        var v10 uint32 = 719
        for fi11 := 0; fi11 < 14; fi11++ {
            var v12 bool = false
            _ = fi11
            _ = v12
        }
        for fi13 := 0; fi13 < 7; fi13++ {
            fmt.Println(v10, fi13)
            if ("\n\t" < ("a" + ("\n\t" + ("a" + "\n\t")))) {
                v14 := []int{}
                for fi15 := 0; fi15 < 3; fi15++ {
                    fmt.Println(fi13, v10)
                    switch "hello" {
                    case "a":
                        if false {
                            v16 := GenS2{F0: -2147483648, F1: 761, F2: 65535}
                            var v17 int8 = 0
                            _ = v16
                            _ = v17
                        } else {
                            p0 = 65535
                            if true {
                                v18 := GenS3{F0: fi15}
                                v10 = (v10 * 684)
                                switch -299 {
                                case 4294967295:
                                    var v19 int = (fi15 * fi15)
                                    _ = v19
                                case -926:
                                    v10 -= 758
                                default:
                                    fmt.Println(p0, p0, fi13)
                                }
                                _ = v18
                            }
                            var v20 uint = 4294967295
                            _ = v20
                        }
                    case "hello":
                        var v21 uint8 = 1
                        _ = v21
                    case "foo bar":
                        switch (("foo bar" + ((("\n\t" + "a") + ("\n\t" + "hello")) + ("hello" + "foo bar"))) + "\n\t") {
                        case "a":
                            if true {
                                v22 := GenS4{F0: 2, F1: "a"}
                                v23 := GenS1{F0: -2147483648, F1: (^(p0)), F2: -862}
                                _ = v22
                                _ = v23
                            }
                        case "hello":
                            p0 += p0
                        case "\n\t":
                            v24 := GenS1{F0: -719, F1: p0, F2: int32(v10)}
                            _ = v24
                        default:
                            if true {
                                var v25 bool = ((true && false) || false)
                                switch fi15 {
                                case -9223372036854775808:
                                    var v26 uint64 = 623
                                    _ = v26
                                case 127:
                                    v27 := GenS3{F0: -161}
                                    _ = v27
                                }
                                if len(v14) >= 2 {
                                    v14 = v14[1:len(v14)]
                                }
                                _ = v25
                            } else {
                                if (!false) {
                                    v10 = (v10 ^ v10)
                                    if (!true) {
                                        if ((v10 << ((2) & 63)) > (v10 - (^(v10)))) {
                                            var v28 uint8 = 189
                                            var v29 uint64 = 2
                                            _ = v28
                                            _ = v29
                                        } else {
                                            v10 -= genFn0()
                                            if (-36 > -307) {
                                                v30 := GenS5{F0: 0}
                                                _ = v30
                                            } else {
                                                v31 := GenS6{F0: 23, F1: 1, F2: p0}
                                                v32 := GenS7{F0: 755}
                                                var v33 int32 = 0
                                                _ = v31
                                                _ = v32
                                                _ = v33
                                            }
                                        }
                                        v10 = v10
                                        p0 = (p0 - p0)
                                    }
                                }
                                switch "foo bar" {
                                case "a":
                                    fmt.Println(fi15, p0)
                                }
                                v34 := GenS8{F0: 32767, F1: 555, F2: 1}
                                _ = v34
                            }
                        }
                    }
                    _ = fi15
                }
                _ = v14
            }
            _ = fi13
        }
        _ = v10
    }
    switch 127 {
    case -1:
        for fi35 := 0; fi35 < 6; fi35++ {
            if true {
                v36 := GenS9{F0: 117}
                v37 := GenS10{F0: p0}
                _ = v36
                _ = v37
            }
            var v38 uint8 = 71
            v39 := GenS6{F0: 0, F1: 2, F2: (p0 >> ((42) & 63))}
            _ = fi35
            _ = v38
            _ = v39
        }
    case 127:
        v40 := GenS8{F0: 32767, F1: 42, F2: 32767}
        _ = v40
    default:
        switch "a" {
        case "a":
            var v41 int32 = 1
            _ = v41
        default:
            var v42 uint = 4294967295
            _ = v42
        }
    }
    v43 := GenS8{F0: 261, F1: 65535, F2: 42}
    v44 := []int{128, 4294967295, 94}
    if len(v44) >= 2 {
        v44 = v44[1:len(v44)]
    }
    _ = p0
    _ = v43
    _ = v44
    return 127
}

func main() {
    var v1 float32 = 0.0
    var v2 uint32 = 42
    var v3 uint64 = 0
    v4 := make(map[string]uint32)
    v5 := make(map[uint32]bool)
    var v6 int32 = 42
    if (((1.5 - (float32(0.1) + (float32(0.0) / (v1 + v1)))) - (v1 + (float32((1e10 + (3.14159 * 1.5))) / v1))) > (v1 - v1)) {
        v3 += v3
        v1 -= v1
    } else {
        fmt.Println(v1, v2)
    }
    for fi45 := 0; fi45 < 7; fi45++ {
        v3 -= 4294967295
        var v46 int = 128
        v47 := GenS5{F0: 236}
        _ = fi45
        _ = v46
        _ = v47
    }
    v48 := GenS7{F0: 393}
    for fi49 := 0; fi49 < 9; fi49++ {
        v50 := GenS3{F0: 4294967295}
        switch "\n\t" {
        case "foo bar":
            var v51 float64 = 3.14159
            _ = v51
        case "":
            var v52 int16 = 0
            _ = v52
        case "\n\t":
            var v53 int16 = -32768
            _ = v53
        default:
            var v54 int64 = -545
            _ = v54
        }
        _ = fi49
        _ = v50
    }
    v55 := GenS8{F0: int16(v3), F1: 4294967295, F2: 1000}
    v2 -= uint32(v3)
    v56 := GenS9{F0: 42}
    fmt.Println(v2, v6)
    _ = len(v4)
    fmt.Println(v1, v2)
    v6 -= (^(v6))
    fmt.Println(v3)
    switch 248 {
    case 0:
        v57 := GenS5{F0: 0}
        _ = v57
    default:
        for fi58 := 0; fi58 < 9; fi58++ {
            if true {
                fmt.Println(v3, v6, fi58)
                switch (^(fi58)) {
                case 42:
                    v59 := []int{}
                    _ = v59
                default:
                    v5[genFn0()] = true
                }
                switch ("\n\t" + "\n\t") {
                case "\n\t":
                    if true {
                        if true {
                            v1 = ((((v1 / v1) + 0.0) * v1) - 1e10)
                            var v60 uint32 = v2
                            var v61 int16 = -32768
                            _ = v60
                            _ = v61
                        } else {
                            v62 := []int{fi58}
                            _ = v62
                        }
                    }
                default:
                    for fi63 := 0; fi63 < 6; fi63++ {
                        v6 -= v6
                        if false {
                            switch "a" {
                            case "hello":
                                if (false && (((-89 < -367) && true) && (!((v2 << ((1) & 63)) > v2)))) {
                                    var v64 float64 = 3.14159
                                    var v65 int16 = 0
                                    _ = v64
                                    _ = v65
                                } else {
                                    delete(v4, ("hello" + "a"))
                                    var v66 int8 = 0
                                    v67 := GenS11{F0: ("" + "a"), F1: v3}
                                    _ = v66
                                    _ = v67
                                }
                            }
                            var v68 string = "hello"
                            _ = v68
                        } else {
                            switch ("\n\t" + (("hello" + "hello") + ("a" + "hello"))) {
                            case "foo bar":
                                var v69 int16 = -1
                                _ = v69
                            case "hello":
                                v2 = v2
                            case "\n\t":
                                var v70 uint8 = 214
                                _ = v70
                            }
                            switch ("\n\t" + "\n\t") {
                            case "\n\t":
                                v1 += v1
                            default:
                                fmt.Println(fi58, fi58)
                            }
                            var v71 int32 = (v6 - (v6 | v6))
                            _ = v71
                        }
                        _ = fi63
                    }
                }
            }
            switch "" {
            case "a":
                v2 = genFn0()
            case "":
                for fi72 := 0; fi72 < 9; fi72++ {
                    v4["foo bar"] = (v2 - (v2 / ((v2) | 1)))
                    fmt.Println(v1)
                    v73 := GenS12{F0: (v3 & (v3 % ((1) | 1)))}
                    _ = fi72
                    _ = v73
                }
            }
            _ = fi58
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v5
    _ = v55
    _ = v56
    _ = v6
}
