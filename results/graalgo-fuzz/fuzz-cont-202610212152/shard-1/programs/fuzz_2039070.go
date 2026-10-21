package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int }
type GenS2 struct { F0 uint64; F1 int32; F2 uint64 }
type GenS3 struct { F0 int8; F1 int8; F2 uint64 }
type GenS4 struct { F0 uint16; F1 int16; F2 bool }
type GenS5 struct { F0 string; F1 uint16 }
type GenS6 struct { F0 uint16 }
func genFn0(p0 int32) uint16 {
    if true {
        for fi7 := 0; fi7 < 20; fi7++ {
            for fi8 := 0; fi8 < 4; fi8++ {
                v9 := []int{fi8}
                var v10 uint = 0
                _ = fi8
                _ = v10
                _ = v9
            }
            if true {
                for fi11 := 0; fi11 < 9; fi11++ {
                    v12 := []int{fi7, (^(fi7))}
                    var v13 uint32 = 4294967295
                    _ = fi11
                    _ = v12
                    _ = v13
                }
                for fi14 := 0; fi14 < 3; fi14++ {
                    v15 := GenS1{F0: 4294967295, F1: (^(fi7))}
                    var v16 uint32 = 1
                    _ = fi14
                    _ = v15
                    _ = v16
                }
                var v17 uint = 341
                _ = v17
            } else {
                fmt.Println(p0, p0)
            }
            switch ("a" + "hello") {
            case "foo bar":
                if false {
                    fmt.Println(fi7)
                } else {
                    fmt.Println(p0, fi7)
                    var v18 uint8 = 159
                    v19 := GenS2{F0: 1, F1: p0, F2: 42}
                    _ = v18
                    _ = v19
                }
            }
            _ = fi7
        }
    }
    v20 := []int{1, 4294967295, -586, -395}
    _ = p0
    _ = v20
    return 820
}

func main() {
    var v1 int = 42
    var v2 int = v1
    v3 := make(map[string]uint32)
    var v4 int = (^(v1))
    var v5 uint64 = 0
    var v6 int64 = 1
    if false {
        v21 := GenS2{F0: v5, F1: -396, F2: (v5 - v5)}
        _ = v21
    } else {
        if false {
            v6 -= 748
        } else {
            if (623 < 0) {
                var v22 int8 = -14
                switch (v2 | (^(v1))) {
                case 0:
                    fmt.Println(v6, v6, v2)
                case 42:
                    var v23 uint32 = 65535
                    _ = v23
                case 1:
                    v22 -= (v22 & (-(v22)))
                default:
                    v4 += 32768
                }
                _ = v22
            }
            switch 390 {
            case 404:
                v24 := []int{(v1 & 32767), (v2 ^ (-(v2)))}
                _ = v24
            case 9223372036854775807:
                switch "a" {
                case "hello":
                    switch ("a" + "hello") {
                    case "hello":
                        v25 := []int{}
                        _ = v25
                    case "\n\t":
                        fmt.Println(v4, v6)
                    case "a":
                        switch ("" + (("\n\t" + ("\n\t" + ("" + "foo bar"))) + "foo bar")) {
                        case "\n\t":
                            fmt.Println(v4, v5)
                        case "":
                            v1 += v4
                        case "a":
                            for fi26 := 0; fi26 < 10; fi26++ {
                                if (1 < 2) {
                                    if ((1 > -128) || (1000 > 32767)) {
                                        if true {
                                            switch ((("foo bar" + "foo bar") + "foo bar") + "") {
                                            case "hello":
                                                v6 -= v6
                                            case "a":
                                                v2 += v2
                                            case "foo bar":
                                                var v27 int = (fi26 % ((26) | 1))
                                                _ = v27
                                            }
                                        } else {
                                            if false {
                                                v28 := GenS3{F0: 0, F1: -94, F2: 463}
                                                var v29 int32 = 42
                                                for fi30 := 0; fi30 < 15; fi30++ {
                                                    var v31 int = v4
                                                    _ = fi30
                                                    _ = v31
                                                }
                                                _ = v28
                                                _ = v29
                                            }
                                        }
                                    }
                                    v32 := []int{fi26}
                                    for fi33 := 0; fi33 < 18; fi33++ {
                                        _ = len(v32)
                                        if ("foo bar" != "\n\t") {
                                            v5 -= (^(v5))
                                            v2 = (fi33 | fi33)
                                            v4 += (fi26 / (((v1 & fi26)) | 1))
                                        }
                                        _ = fi33
                                    }
                                    _ = v32
                                } else {
                                    for fi34 := 0; fi34 < 9; fi34++ {
                                        delete(v3, "hello")
                                        var v35 uint64 = (v5 * 18446744073709551615)
                                        _ = fi34
                                        _ = v35
                                    }
                                    v36 := GenS3{F0: 100, F1: 1, F2: v5}
                                    v1 += v4
                                    _ = v36
                                }
                                v6 -= (v6 / ((v6) | 1))
                                for fi37 := 0; fi37 < 6; fi37++ {
                                    var v38 uint32 = 65535
                                    v6 = (v6 * (v6 << ((uint(v5)) & 63)))
                                    _ = fi37
                                    _ = v38
                                }
                                _ = fi26
                            }
                        default:
                            if true {
                                v1 = v1
                                switch v1 {
                                case 9223372036854775807:
                                    v4 = -9223372036854775808
                                }
                                if false {
                                    v2 += v1
                                    for fi39 := 0; fi39 < 6; fi39++ {
                                        _ = len(v3)
                                        fmt.Println(v4, v5)
                                        _ = fi39
                                    }
                                    if (true && (174 < 128)) {
                                        v1 -= (v1 % ((42) | 1))
                                        _ = len(v3)
                                        v2 += v1
                                    }
                                }
                            } else {
                                switch ("a" + ("hello" + (("foo bar" + ((("a" + "\n\t") + ("" + "foo bar")) + "\n\t")) + "hello"))) {
                                case "\n\t":
                                    v40 := []int{v1}
                                    _ = v40
                                case "a":
                                    var v41 int64 = v6
                                    _ = v41
                                }
                            }
                        }
                    }
                case "\n\t":
                    for fi42 := 0; fi42 < 17; fi42++ {
                        var v43 uint8 = 0
                        _ = fi42
                        _ = v43
                    }
                case "foo bar":
                    if false {
                        var v44 int8 = -9
                        v2 = 32768
                        _ = v44
                    }
                default:
                    switch v4 {
                    case 128:
                        if true {
                            fmt.Println(v4, v1)
                            var v45 int16 = -1
                            _ = v45
                        }
                    default:
                        fmt.Println(v1)
                    }
                }
            }
            for fi46 := 0; fi46 < 1; fi46++ {
                if (true && false) {
                    var v47 uint16 = genFn0(2147483647)
                    for fi48 := 0; fi48 < 15; fi48++ {
                        var v49 uint8 = 138
                        v50 := GenS2{F0: 2, F1: -938, F2: v5}
                        fmt.Println(fi46)
                        _ = fi48
                        _ = v49
                        _ = v50
                    }
                    if (false || false) {
                        v51 := v3[("hello" + "")]
                        var v52 int64 = -1
                        for fi53 := 0; fi53 < 6; fi53++ {
                            v54 := GenS1{F0: (v51 * v51), F1: (v4 * fi46)}
                            fmt.Println(v6)
                            fmt.Println(v4, v4)
                            _ = fi53
                            _ = v54
                        }
                        _ = v51
                        _ = v52
                    }
                    _ = v47
                } else {
                    v55 := GenS4{F0: 195, F1: -472, F2: (957 <= 695)}
                    v2 += fi46
                    v6 -= v6
                    _ = v55
                }
                if ("foo bar" >= ("foo bar" + "a")) {
                    v56 := []int{(-(fi46)), (v1 * (fi46 ^ (^(v1))))}
                    fmt.Println(v1)
                    var v57 float32 = 0.0
                    _ = v56
                    _ = v57
                }
                var v58 uint = 42
                _ = fi46
                _ = v58
            }
        }
        if true {
            switch "foo bar" {
            case "foo bar":
                v59 := GenS4{F0: 248, F1: -32768, F2: false}
                _ = v59
            case "":
                var v60 int16 = 934
                _ = v60
            }
            v61 := []int{(v4 >> ((18446744073709551615) & 63)), (v2 % (((v4 | v4)) | 1))}
            fmt.Println(v1)
            _ = v61
        } else {
            v62 := []int{v1, int(v6), (v4 * -1), v1}
            _ = v62
        }
    }
    var v63 string = "\n\t"
    v2 = -48
    v64 := []int{v2, 42, int(v6), v1}
    fmt.Println(v5, v63, v6)
    for fi65 := 0; fi65 < 10; fi65++ {
        for fi66 := 0; fi66 < 17; fi66++ {
            fmt.Println(v6, v2, fi66)
            var v67 uint8 = 128
            fmt.Println(v5, v4, v5)
            _ = fi66
            _ = v67
        }
        _ = fi65
    }
    switch 42 {
    case 42:
        if false {
            var v68 bool = false
            switch "" {
            case "hello":
                for fi69 := 0; fi69 < 6; fi69++ {
                    v70 := GenS5{F0: v63, F1: 913}
                    _ = fi69
                    _ = v70
                }
            case "\n\t":
                var v71 uint64 = (v5 | v5)
                _ = v71
            case "foo bar":
                v2 -= (v2 / ((v2) | 1))
            }
            _ = v68
        }
    }
    v72 := GenS6{F0: genFn0(-2147483648)}
    var v73 uint = 0
    v5 = (v5 | 4294967295)
    delete(v3, ("foo bar" + ((("foo bar" + "hello") + ("\n\t" + v63)) + v63)))
    fmt.Println(v5)
    var v74 int = v4
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v63:", v63)
    fmt.Println("slicelen v64:", len(v64), v64)
    fmt.Println("v73:", v73)
    fmt.Println("v74:", v74)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v63
    _ = v64
    _ = v72
    _ = v73
    _ = v74
}
