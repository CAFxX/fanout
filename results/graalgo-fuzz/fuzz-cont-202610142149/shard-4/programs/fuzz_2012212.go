package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8; F2 int8 }
type GenS2 struct { F0 uint16; F1 int; F2 int32 }
type GenS3 struct { F0 bool; F1 int64; F2 string }
type GenS4 struct { F0 bool; F1 int }
type GenS5 struct { F0 uint64; F1 uint64 }
type GenS6 struct { F0 uint16 }
func genFn0(p0 uint16, p1 uint32, p2 bool) uint16 {
    for fi4 := 0; fi4 < 7; fi4++ {
        v5 := GenS1{F0: 0, F1: 255, F2: 100}
        v6 := GenS2{F0: p0, F1: (^(fi4)), F2: 0}
        _ = fi4
        _ = v5
        _ = v6
    }
    v7 := GenS1{F0: 2, F1: 77, F2: 42}
    if ((2 != 1) && (p2 || (1 >= p0))) {
        for fi8 := 0; fi8 < 14; fi8++ {
            switch "\n\t" {
            case "\n\t":
                v9 := GenS2{F0: (p0 >> ((18446744073709551615) & 63)), F1: (-(fi8)), F2: -2147483648}
                _ = v9
            case "hello":
                v10 := []int{}
                _ = v10
            default:
                v11 := GenS3{F0: p2, F1: -9223372036854775808, F2: ""}
                _ = v11
            }
            _ = fi8
        }
        if (-316 < 1) {
            v12 := []int{}
            _ = v12
        }
    } else {
        p2 = (true || (1e10 <= -2.25))
        var v13 uint16 = 1
        _ = v13
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 2
}

func main() {
    var v1 int8 = -128
    var v2 int64 = 9223372036854775807
    v3 := []string{("\n\t" + (((("\n\t" + "\n\t") + "\n\t") + "hello") + "a"))}
    v14 := GenS1{F0: 167, F1: 42, F2: v1}
    if true {
        if (1 <= 2) {
            switch "foo bar" {
            case "\n\t":
                v15 := GenS3{F0: true, F1: 1, F2: ("foo bar" + "")}
                _ = v15
            case "a":
                v16 := GenS3{F0: (false || (-1 > -1)), F1: int64(v1), F2: "\n\t"}
                _ = v16
            case "hello":
                if len(v3) > 0 {
                    v3[0] = (("" + ("foo bar" + ("a" + "a"))) + "")
                }
            }
            if false {
                var v17 uint32 = 2147483647
                fmt.Println(v17)
                v18 := GenS4{F0: true, F1: -9223372036854775808}
                _ = v17
                _ = v18
            } else {
                switch "" {
                case "\n\t":
                    fmt.Println(v2)
                case "":
                    v19 := GenS5{F0: 0, F1: 2}
                    _ = v19
                case "a":
                    if ("foo bar" > "\n\t") {
                        v2 = (^(v2))
                        switch "" {
                        case "":
                            v20 := GenS5{F0: 2, F1: 371}
                            _ = v20
                        }
                        if (!true) {
                            v1 -= -128
                            v1 -= (v1 - (v1 - 0))
                        } else {
                            fmt.Println(v2)
                        }
                    }
                default:
                    fmt.Println(v1, v1)
                }
                v3 = append(v3, "")
                var v21 bool = (!false)
                _ = v21
            }
            if (!true) {
                if false {
                    var v22 uint = 42
                    if (3.14159 <= -2.25) {
                        if true {
                            var v23 bool = true
                            v1 += (v1 >> (((^(v22))) & 63))
                            _ = v23
                        } else {
                            v1 = 100
                            var v24 int = 127
                            _ = v24
                        }
                        switch ("a" + "foo bar") {
                        case "":
                            fmt.Println(v22, v22)
                        default:
                            v1 = 24
                        }
                    } else {
                        v2 += v2
                    }
                    _ = v22
                } else {
                    var v25 uint8 = 2
                    var v26 uint8 = (v25 / ((v25) | 1))
                    _ = v25
                    _ = v26
                }
            } else {
                var v27 uint16 = 2
                _ = v27
            }
        } else {
            v2 = (v2 * (v2 << ((4294967295) & 63)))
            v2 += v2
            if false {
                var v28 bool = (true || false)
                var v29 int16 = 615
                v2 -= 606
                _ = v28
                _ = v29
            } else {
                var v30 uint64 = 0
                _ = v30
            }
        }
    } else {
        _ = len(v3)
        var v31 float64 = 1.5
        var v32 int32 = 243
        _ = v31
        _ = v32
    }
    var v33 int8 = v1
    if (1000 > 32767) {
        switch ("" + "a") {
        case "":
            v2 -= (v2 / (((-(v2))) | 1))
        case "\n\t":
            v2 = (v2 & v2)
        case "a":
            for fi34 := 0; fi34 < 13; fi34++ {
                v2 += (v2 * (-(v2)))
                _ = fi34
            }
        default:
            v1 = v1
        }
        var v35 uint16 = 32767
        v33 = (v33 >> ((18446744073709551615) & 63))
        _ = v35
    }
    if len(v3) > 0 {
        v3[0] = "\n\t"
    }
    v2 = (^(v2))
    fmt.Println(v2, v33, v33)
    var v36 bool = false
    v37 := GenS5{F0: 2, F1: 545}
    var v38 int64 = (v2 - (^(v2)))
    switch "a" {
    case "":
        var v39 int32 = 391
        _ = v39
    default:
        for fi40 := 0; fi40 < 14; fi40++ {
            for fi41 := 0; fi41 < 10; fi41++ {
                if (1 != 681) {
                    v42 := GenS4{F0: v36, F1: (fi40 << ((4294967295) & 63))}
                    v43 := GenS5{F0: 1, F1: 831}
                    _ = v42
                    _ = v43
                }
                _ = fi41
            }
            v36 = v36
            for fi44 := 0; fi44 < 9; fi44++ {
                v1 += (-(v1))
                var v45 uint32 = 1
                _ = len(v3)
                v36 = ((!(((!v36) && (!(genFn0(65535, (^(v45)), false) <= 65535))) || v36)) && v36)
                var v46 int32 = -276
                _ = fi44
                _ = v45
                _ = v46
            }
            _ = fi40
        }
    }
    v47 := GenS6{F0: 1}
    if v36 {
        v38 = v2
        fmt.Println(v33, v38, v33)
    } else {
        v3 = append(v3, "\n\t")
        var v48 int8 = v33
        v36 = v36
        _ = v48
    }
    v1 = (v33 | v1)
    for fi49 := 0; fi49 < 19; fi49++ {
        v50 := GenS5{F0: 42, F1: 481}
        v2 = 773
        _ = fi49
        _ = v50
    }
    v38 += v38
    var v51 float64 = 0.0
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("v36:", v36)
    fmt.Println("v38:", v38)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v33
    _ = v36
    _ = v37
    _ = v38
    _ = v47
    _ = v51
}
