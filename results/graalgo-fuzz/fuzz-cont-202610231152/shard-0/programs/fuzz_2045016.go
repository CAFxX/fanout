package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32 }
type GenS2 struct { F0 uint16; F1 int64; F2 int64 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int16; F1 int16 }
type GenS5 struct { F0 int; F1 uint16; F2 string }
type GenS6 struct { F0 uint64; F1 int32; F2 uint16 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 bool; F1 uint8 }
func genFn0() uint32 {
    v6 := GenS1{F0: 805, F1: 461}
    v7 := GenS2{F0: 0, F1: -747, F2: -9223372036854775808}
    var v8 uint8 = 1
    _ = v6
    _ = v7
    _ = v8
    return 2147483648
}
func genFn1(p0 uint64, p1 int) bool {
    fmt.Println(p0, p0)
    p1 += p1
    var v9 float32 = 3.14159
    if true {
        if ((32767 >= -823) && (false || false)) {
            fmt.Println(v9, p0, p0)
            v10 := []int{p1}
            for fi11 := 0; fi11 < 6; fi11++ {
                v9 -= (3.14159 / v9)
                if true {
                    var v12 int64 = 9223372036854775807
                    var v13 uint8 = 127
                    _ = v12
                    _ = v13
                } else {
                    var v14 uint64 = p0
                    v15 := GenS2{F0: 1, F1: -9223372036854775808, F2: 769}
                    _ = v14
                    _ = v15
                }
                _ = fi11
            }
            _ = v10
        }
        var v16 int32 = -1
        v16 = (-(v16))
        _ = v16
    } else {
        v9 += v9
        switch "a" {
        case "":
            switch ("" + "a") {
            case "":
                if (1 < p0) {
                    v17 := []int{(-(p1))}
                    var v18 uint = 2
                    if len(v17) >= 2 {
                        v17 = v17[1:len(v17)]
                    }
                    _ = v17
                    _ = v18
                } else {
                    var v19 int32 = 42
                    for fi20 := 0; fi20 < 4; fi20++ {
                        v21 := []int{p1, 32768}
                        if (p1 == p1) {
                            v22 := GenS2{F0: 255, F1: 366, F2: -678}
                            if len(v21) > 0 {
                                v21[0] = 2147483647
                            }
                            if len(v21) > 0 {
                                v21[0] = p1
                            }
                            _ = v22
                        } else {
                            var v23 int16 = 1
                            for fi24 := 0; fi24 < 9; fi24++ {
                                v19 += 0
                                _ = fi24
                            }
                            if (163 < 2) {
                                _ = len(v21)
                                if (-2.25 <= 1e10) {
                                    var v25 int64 = 1
                                    for fi26 := 0; fi26 < 11; fi26++ {
                                        var v27 float32 = v9
                                        v28 := GenS2{F0: 447, F1: v25, F2: (v25 - 42)}
                                        _ = fi26
                                        _ = v27
                                        _ = v28
                                    }
                                    _ = v25
                                } else {
                                    var v29 uint = 18446744073709551615
                                    v9 = float32((-2.25 - (0.0 * 1.5)))
                                    switch (-(p1)) {
                                    case 946:
                                        var v30 int8 = 127
                                        _ = v30
                                    default:
                                        v29 += 2
                                    }
                                    _ = v29
                                }
                            }
                            _ = v23
                        }
                        fmt.Println(p1, fi20)
                        _ = fi20
                        _ = v21
                    }
                    fmt.Println(v9, p1)
                    _ = v19
                }
            case "a":
                v31 := GenS3{F0: -1}
                _ = v31
            case "\n\t":
                var v32 uint32 = 194
                _ = v32
            default:
                var v33 float64 = 1e10
                _ = v33
            }
        case "hello":
            v9 -= (((v9 - ((3.14159 - (v9 * v9)) - ((0.0 + 0.1) / (v9 / v9)))) / v9) * (((1.5 + float32(-2.25)) / (3.14159 + (v9 / 0.1))) - v9))
        case "a":
            for fi34 := 0; fi34 < 15; fi34++ {
                v9 -= v9
                _ = fi34
            }
        default:
            switch p1 {
            case -9223372036854775808:
                v35 := GenS4{F0: 0, F1: 42}
                _ = v35
            case -826:
                v9 = -2.25
            }
        }
        v36 := []int{}
        _ = v36
    }
    p0 += (^(p0))
    _ = p0
    _ = p1
    _ = v9
    return (("foo bar" + "\n\t") == "hello")
}

func main() {
    v1 := []uint32{65535}
    var v2 uint64 = 42
    v3 := []string{"foo bar", ("\n\t" + "foo bar"), "a"}
    v4 := make(map[int]bool)
    var v5 float64 = 1e10
    for fi37 := 0; fi37 < 15; fi37++ {
        for fi38 := 0; fi38 < 15; fi38++ {
            if (false && true) {
                v5 -= float64(-2.25)
                var v39 uint = 42
                _ = v39
            }
            _ = fi38
        }
        v40 := GenS2{F0: 65535, F1: -320, F2: 0}
        _ = fi37
        _ = v40
    }
    var v41 bool = (true || (!false))
    v42 := GenS3{F0: -2147483648}
    v41 = v41
    for fi43 := 0; fi43 < 2; fi43++ {
        if false {
            var v44 uint64 = 42
            _ = v44
        } else {
            fmt.Println(v41)
            if v41 {
                var v45 uint16 = 1
                fmt.Println(v41, v5)
                _ = v45
            }
            var v46 float32 = (1.5 * float32(v5))
            _ = v46
        }
        var v47 int64 = 214
        switch ("" + "hello") {
        case "hello":
            if (!false) {
                v41 = (!(!v41))
                if ("\n\t" == "a") {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                }
                if (!v41) {
                    var v48 int64 = 399
                    _ = v48
                } else {
                    if (true && false) {
                        v2 -= v2
                    } else {
                        var v49 float32 = 1.5
                        v50 := GenS2{F0: 1, F1: v47, F2: (v47 & v47)}
                        var v51 uint16 = 42
                        _ = v49
                        _ = v50
                        _ = v51
                    }
                }
            }
        case "\n\t":
            for fi52 := 0; fi52 < 10; fi52++ {
                v53 := GenS2{F0: 619, F1: v47, F2: (v47 * -9223372036854775808)}
                _ = len(v1)
                if v41 {
                    v47 -= (-(v47))
                    v47 = -379
                    fmt.Println(v5)
                } else {
                    delete(v4, (fi43 >> ((uint(v2)) & 63)))
                    var v54 uint8 = 1
                    v5 += (float64((3.14159 * (((1e10 * float32(1.5)) / (1.5 + 0.1)) - 0.1))) * (1e10 * v5))
                    _ = v54
                }
                v55 := GenS5{F0: (^(fi43)), F1: 255, F2: ("\n\t" + ("foo bar" + (("a" + "") + "a")))}
                v5 = float64(0.1)
                _ = fi52
                _ = v53
                _ = v55
            }
        case "a":
            v56 := GenS6{F0: (v2 & v2), F1: 439, F2: 65535}
            _ = v56
        }
        _ = fi43
        _ = v47
    }
    v41 = (v5 < float64(0.1))
    switch ("a" + "foo bar") {
    case "foo bar":
        if len(v1) > 0 {
            v1[0] = 42
        }
    case "\n\t":
        v2 = v2
    case "":
        v57 := GenS5{F0: 9223372036854775807, F1: 465, F2: ((("" + "hello") + "a") + "\n\t")}
        _ = v57
    default:
        v58 := GenS7{F0: "foo bar"}
        _ = v58
    }
    switch 2147483647 {
    case 127:
        v59 := GenS6{F0: 2, F1: -737, F2: 750}
        _ = v59
    case 32768:
        switch "" {
        case "foo bar":
            v60 := GenS2{F0: 48, F1: -648, F2: -859}
            _ = v60
        case "":
            v5 = v5
        case "\n\t":
            for fi61 := 0; fi61 < 14; fi61++ {
                v62 := v4[32767]
                _ = fi61
                _ = v62
            }
        default:
            v63 := GenS5{F0: 2147483647, F1: 0, F2: "\n\t"}
            _ = v63
        }
    case 0:
        fmt.Println(v2, v2)
    default:
        v64 := GenS8{F0: v41, F1: 42}
        _ = v64
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v5
}
