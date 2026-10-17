package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64; F2 uint16 }
type GenS2 struct { F0 int8; F1 int16; F2 int16 }
type GenS3 struct { F0 string; F1 uint64; F2 int32 }
func genFn0(p0 uint8) uint {
    switch "\n\t" {
    case "a":
        fmt.Println(p0)
    case "hello":
        switch 42 {
        case 1:
            if (!((32767 < 1000) && (((!(9223372036854775807 >= 9223372036854775807)) || false) || (4294967295 == 41)))) {
                var v5 uint16 = 65535
                p0 = p0
                v6 := []int{2147483647, 176}
                _ = v5
                _ = v6
            }
        case 60:
            if (!false) {
                for fi7 := 0; fi7 < 15; fi7++ {
                    p0 = p0
                    for fi8 := 0; fi8 < 4; fi8++ {
                        v9 := []int{fi8, (fi7 / (((fi8 >> ((4294967295) & 63))) | 1)), fi7}
                        p0 = 127
                        switch "\n\t" {
                        case "\n\t":
                            if false {
                                v10 := GenS1{F0: 0, F1: -1, F2: 2}
                                _ = v10
                            } else {
                                if len(v9) > 0 {
                                    v9[0] = (-(fi8))
                                }
                                var v11 int64 = 0
                                _ = v11
                            }
                        case "hello":
                            p0 -= 0
                        default:
                            fmt.Println(fi7, fi8, fi7)
                        }
                        _ = fi8
                        _ = v9
                    }
                    _ = fi7
                }
            }
        case 127:
            for fi12 := 0; fi12 < 6; fi12++ {
                for fi13 := 0; fi13 < 13; fi13++ {
                    if (1 <= -1) {
                        p0 = p0
                        var v14 float64 = float64((0.1 + 0.1))
                        _ = v14
                    } else {
                        v15 := []int{-1, (fi12 & 803), (fi13 >> ((509) & 63)), 42}
                        var v16 int32 = 2147483647
                        _ = len(v15)
                        p0 = p0
                        _ = v15
                        _ = v16
                    }
                    v17 := GenS1{F0: 539, F1: -104, F2: 65535}
                    _ = fi13
                    _ = v17
                }
                _ = fi12
            }
        }
    }
    p0 = 42
    for fi18 := 0; fi18 < 16; fi18++ {
        var v19 uint64 = 18446744073709551615
        v20 := GenS1{F0: 329, F1: 1, F2: uint16(v19)}
        _ = fi18
        _ = v19
        _ = v20
    }
    p0 -= p0
    fmt.Println(p0)
    _ = p0
    return 42
}

func main() {
    var v1 int32 = -210
    var v2 uint64 = 18446744073709551615
    var v3 uint32 = 2147483648
    var v4 float32 = 0.1
    if true {
        switch 0 {
        case 42:
            for fi21 := 0; fi21 < 16; fi21++ {
                var v22 uint16 = 65535
                var v23 int32 = (v1 - 42)
                _ = fi21
                _ = v22
                _ = v23
            }
        default:
            v2 += (v2 ^ (v2 * (v2 * v2)))
        }
    }
    if (("foo bar" + "") <= "") {
        v2 = 4
        for fi24 := 0; fi24 < 6; fi24++ {
            v3 = 2
            v25 := []int{}
            fmt.Println(v4)
            _ = fi24
            _ = v25
        }
        var v26 uint = 401
        _ = v26
    } else {
        var v27 uint64 = v2
        var v28 bool = false
        for fi29 := 0; fi29 < 11; fi29++ {
            fmt.Println(fi29, v27)
            _ = fi29
        }
        _ = v27
        _ = v28
    }
    if (false && false) {
        if (!(!false)) {
            v30 := []int{42, 0, 127}
            for fi31 := 0; fi31 < 3; fi31++ {
                if len(v30) > 0 {
                    v30[0] = fi31
                }
                if len(v30) > 0 {
                    v30[0] = (fi31 & fi31)
                }
                _ = fi31
            }
            for fi32 := 0; fi32 < 20; fi32++ {
                v33 := GenS2{F0: 0, F1: 611, F2: 32767}
                _ = fi32
                _ = v33
            }
            _ = v30
        } else {
            if true {
                switch "foo bar" {
                case "a":
                    v34 := GenS1{F0: 65535, F1: 1, F2: 303}
                    _ = v34
                case "hello":
                    v4 += (0.1 + v4)
                case "":
                    v35 := []int{908, 127}
                    _ = v35
                default:
                    var v36 uint64 = v2
                    _ = v36
                }
            } else {
                if false {
                    var v37 uint64 = 2
                    if false {
                        var v38 int32 = 210
                        _ = v38
                    }
                    var v39 int32 = -1
                    _ = v37
                    _ = v39
                }
                var v40 int16 = -32768
                fmt.Println(v40, v2, v4)
                _ = v40
            }
            v4 -= v4
            if (((1 <= 0) || true) || (1.5 < 1e10)) {
                var v41 uint32 = v3
                v4 = v4
                var v42 string = ""
                _ = v41
                _ = v42
            }
        }
    } else {
        v1 += (v1 & 0)
    }
    if (!((genFn0(9) == uint(v3)) || true)) {
        var v43 uint32 = (v3 * (v3 | v3))
        if (((^(v43)) != 2) && true) {
            var v44 int64 = 0
            for fi45 := 0; fi45 < 5; fi45++ {
                fmt.Println(v44, v1, v4)
                v46 := GenS3{F0: "foo bar", F1: (^(v2)), F2: v1}
                var v47 int32 = v1
                _ = fi45
                _ = v46
                _ = v47
            }
            _ = v44
        } else {
            switch "a" {
            case "hello":
                v48 := GenS1{F0: 65535, F1: 0, F2: 0}
                _ = v48
            case "\n\t":
                var v49 uint = 65535
                _ = v49
            default:
                v3 -= 1
            }
        }
        v3 = v43
        _ = v43
    }
    v50 := []int{9223372036854775807, 32768, 0}
    if (!false) {
        if len(v50) > 0 {
            v50[0] = 9223372036854775807
        }
    }
    v3 += v3
    var v51 uint8 = 128
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v50:", len(v50), v50)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v50
    _ = v51
}
