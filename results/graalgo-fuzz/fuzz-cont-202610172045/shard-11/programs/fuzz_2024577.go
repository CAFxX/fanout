package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint; F1 int32 }
type GenS3 struct { F0 bool; F1 string; F2 int }
type GenS4 struct { F0 uint; F1 uint64; F2 uint32 }
type GenS5 struct { F0 uint64; F1 uint8; F2 uint }
type GenS6 struct { F0 uint64; F1 uint16 }
type GenS7 struct { F0 uint16; F1 bool }
type GenS8 struct { F0 uint; F1 int }
type GenS9 struct { F0 bool }
func genFn0() int64 {
    var v5 int64 = -1
    v5 += (v5 | v5)
    for fi6 := 0; fi6 < 18; fi6++ {
        v7 := []int{(^(fi6)), fi6, (^(fi6))}
        v8 := GenS1{F0: 42}
        var v9 string = ""
        _ = fi6
        _ = v7
        _ = v8
        _ = v9
    }
    v10 := GenS1{F0: v5}
    switch "" {
    case "a":
        v11 := GenS1{F0: (^(v5))}
        _ = v11
    }
    _ = v10
    _ = v5
    return v5
}
func genFn1(p0 int64, p1 int32) uint8 {
    fmt.Println(p1, p0)
    fmt.Println(p1, p1)
    for fi12 := 0; fi12 < 18; fi12++ {
        var v13 uint32 = 0
        if (42 <= 0) {
            var v14 int = fi12
            _ = v14
        } else {
            v15 := GenS1{F0: p0}
            v16 := GenS2{F0: 65535, F1: (p1 << ((4294967295) & 63))}
            var v17 int64 = p0
            _ = v15
            _ = v16
            _ = v17
        }
        _ = fi12
        _ = v13
    }
    _ = p0
    _ = p1
    return 241
}

func main() {
    var v1 int16 = -15
    var v2 uint16 = 406
    v3 := []int{0, -1, 32767}
    v4 := make(map[uint32]string)
    switch -447 {
    case -2147483648:
        v2 = v2
    case 509:
        for fi18 := 0; fi18 < 2; fi18++ {
            var v19 string = ""
            v1 = 42
            for fi20 := 0; fi20 < 12; fi20++ {
                if (fi20 > fi18) {
                    if true {
                        var v21 uint = 215
                        var v22 float64 = float64(-2.25)
                        _ = v21
                        _ = v22
                    } else {
                        v23 := GenS2{F0: 18446744073709551615, F1: -1}
                        var v24 int64 = 957
                        fmt.Println(v2, fi18, v24)
                        _ = v23
                        _ = v24
                    }
                    if true {
                        var v25 int32 = -878
                        _ = v25
                    }
                    v26 := GenS3{F0: true, F1: "\n\t", F2: fi18}
                    switch v19 {
                    case "foo bar":
                        v1 -= v1
                    case "a":
                        switch fi18 {
                        case 847:
                            v27 := GenS4{F0: 2, F1: uint64(v2), F2: 0}
                            _ = v27
                        case -514:
                            v28 := GenS5{F0: 738, F1: genFn1(9223372036854775807, 42), F2: 0}
                            _ = v28
                        case 992:
                            if true {
                                var v29 uint16 = v2
                                if len(v3) > 0 {
                                    v3[0] = fi20
                                }
                                if ((2 >= 255) && true) {
                                    _ = len(v3)
                                    fmt.Println(v19, v29)
                                }
                                _ = v29
                            } else {
                                switch (v19 + ("hello" + v19)) {
                                case "foo bar":
                                    v2 += 1
                                case "\n\t":
                                    var v30 uint64 = 1
                                    _ = v30
                                default:
                                    v2 = 794
                                }
                            }
                        }
                    default:
                        v1 -= (v1 << ((986) & 63))
                    }
                    _ = v26
                } else {
                    switch 0 {
                    case -316:
                        fmt.Println(v19, v19)
                    case 523:
                        var v31 bool = ((false && true) || true)
                        _ = v31
                    }
                    v19 = "a"
                }
                if (!(!false)) {
                    v2 += 42
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    v4[0] = (v19 + v19)
                    var v32 uint64 = 700
                    var v33 uint = 255
                    _ = v32
                    _ = v33
                } else {
                    switch (v19 + v19) {
                    case "\n\t":
                        v19 = ("" + "hello")
                    case "foo bar":
                        if (genFn1(int64(v1), 2147483647) < 127) {
                            var v34 int32 = -715
                            switch "foo bar" {
                            case "\n\t":
                                v35 := GenS6{F0: 42, F1: 0}
                                _ = v35
                            }
                            switch 32767 {
                            case -209:
                                var v36 int = fi18
                                _ = v36
                            case -709:
                                v37 := GenS2{F0: 1, F1: (v34 & v34)}
                                _ = v37
                            case 9223372036854775807:
                                var v38 int = (fi20 * fi18)
                                _ = v38
                            }
                            v34 -= (v34 / ((v34) | 1))
                            _ = v34
                        } else {
                            fmt.Println(fi18)
                        }
                    default:
                        var v39 uint8 = 214
                        _ = v39
                    }
                    var v40 int64 = 9223372036854775807
                    _ = v40
                }
                if true {
                    v41 := GenS7{F0: (v2 / ((v2) | 1)), F1: false}
                    v19 = ("" + v19)
                    v1 -= v1
                    v2 = v2
                    v42 := GenS4{F0: 42, F1: 42, F2: 4294967295}
                    var v43 bool = false
                    _ = v41
                    _ = v42
                    _ = v43
                }
                _ = fi20
            }
            _ = fi18
            _ = v19
        }
    case -9223372036854775808:
        switch "\n\t" {
        case "a":
            if false {
                fmt.Println(v1)
                v1 += 42
            }
        }
    }
    var v44 uint8 = 142
    var v45 uint32 = 65535
    switch ("foo bar" + "\n\t") {
    case "":
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    case "hello":
        delete(v4, (v45 - v45))
    case "a":
        if false {
            var v46 int8 = -1
            v1 = (v1 >> ((1) & 63))
            _ = v46
        } else {
            if false {
                for fi47 := 0; fi47 < 7; fi47++ {
                    var v48 string = ("" + "\n\t")
                    _ = len(v4)
                    if (!((v48 == (v48 + "hello")) || true)) {
                        var v49 int = -976
                        _ = v49
                    } else {
                        if false {
                            var v50 uint = 18446744073709551615
                            _ = v50
                        }
                    }
                    _ = fi47
                    _ = v48
                }
            }
            var v51 int64 = -9223372036854775808
            for fi52 := 0; fi52 < 9; fi52++ {
                if false {
                    fmt.Println(v51)
                }
                v1 += v1
                _ = fi52
            }
            _ = v51
        }
    }
    var v53 int = 128
    _ = len(v3)
    v1 = 42
    if false {
        fmt.Println(v44, v53, v1)
    }
    var v54 int8 = 0
    fmt.Println(v44, v44, v2)
    v55 := v4[(v45 * 600)]
    if false {
        delete(v4, v45)
    }
    fmt.Println(v45)
    for fi56 := 0; fi56 < 11; fi56++ {
        if len(v3) > 0 {
            v3[0] = 4294967295
        }
        if (!false) {
            _ = len(v3)
            if true {
                v57 := GenS8{F0: 4294967295, F1: (-(v53))}
                v58 := GenS7{F0: (v2 - v2), F1: (529 != 598)}
                _ = v57
                _ = v58
            } else {
                v59 := GenS9{F0: false}
                if (fi56 != fi56) {
                    v60 := GenS4{F0: 0, F1: 4294967295, F2: v45}
                    for fi61 := 0; fi61 < 19; fi61++ {
                        v54 = v54
                        if (!true) {
                            var v62 uint64 = 1
                            _ = v62
                        }
                        var v63 int16 = 1
                        _ = fi61
                        _ = v63
                    }
                    _ = v60
                }
                _ = v59
            }
            if false {
                v64 := GenS4{F0: 18446744073709551615, F1: 672, F2: (v45 * (v45 * 2147483648))}
                for fi65 := 0; fi65 < 18; fi65++ {
                    if false {
                        v44 = v44
                        var v66 float32 = (0.1 / ((((1.5 - 1.5) * 1e10) + 3.14159) - 0.0))
                        _ = v66
                    }
                    _ = fi65
                }
                v45 = (v45 + v45)
                _ = v64
            } else {
                v1 += int16(v54)
                v45 += 4294967295
            }
        } else {
            for fi67 := 0; fi67 < 16; fi67++ {
                var v68 int = 0
                fmt.Println(v54, fi56)
                _ = fi67
                _ = v68
            }
            v54 = v54
        }
        _ = fi56
    }
    fmt.Println(v53, v1)
    switch v53 {
    case 4294967295:
        fmt.Println(v2, v55, v2)
    default:
        for fi69 := 0; fi69 < 8; fi69++ {
            fmt.Println(fi69, v1, v44)
            v1 += v1
            var v70 int16 = 42
            _ = fi69
            _ = v70
        }
    }
    v44 -= v44
    if len(v3) > 0 {
        v3[0] = (^(v53))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v53:", v53)
    fmt.Println("v54:", v54)
    fmt.Println("v55:", v55)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v53
    _ = v54
    _ = v55
}
