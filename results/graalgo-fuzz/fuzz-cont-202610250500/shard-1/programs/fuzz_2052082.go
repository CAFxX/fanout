package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int; F2 int16 }
type GenS2 struct { F0 uint32; F1 bool; F2 int32 }
type GenS3 struct { F0 uint; F1 int32; F2 uint64 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint16; F1 bool }
type GenS6 struct { F0 int32; F1 uint64; F2 int8 }
type GenS7 struct { F0 uint16; F1 int64; F2 int32 }
func genFn0(p0 int, p1 uint32, p2 uint32) uint16 {
    var v4 int8 = -128
    var v5 bool = false
    v6 := GenS1{F0: v5, F1: p0, F2: 1000}
    v7 := GenS1{F0: v5, F1: (-(p0)), F2: 1}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return 2
}
func genFn1(p0 int32, p1 uint16) int64 {
    v8 := GenS1{F0: (33 <= 65535), F1: -973, F2: 1000}
    switch ("hello" + "") {
    case "hello":
        v9 := []int{}
        _ = v9
    case "foo bar":
        v10 := GenS1{F0: false, F1: 32767, F2: 457}
        _ = v10
    case "\n\t":
        v11 := []int{-2147483648, 32767}
        _ = v11
    }
    switch -1 {
    case 9223372036854775807:
        v12 := []int{-642, -9223372036854775808}
        _ = v12
    case 17:
        p1 -= genFn0(263, 780, 0)
    case 32768:
        var v13 float64 = 3.14159
        _ = v13
    default:
        switch "foo bar" {
        case "hello":
            var v14 uint = 18446744073709551615
            _ = v14
        case "\n\t":
            if false {
                p0 += p0
            }
        case "foo bar":
            var v15 uint32 = 2147483647
            _ = v15
        default:
            v16 := GenS2{F0: 65535, F1: true, F2: p0}
            _ = v16
        }
    }
    p0 -= (p0 & p0)
    v17 := []int{0, 127}
    _ = p0
    _ = p1
    _ = v17
    _ = v8
    return 180
}

func main() {
    var v1 float32 = ((3.14159 + -2.25) / 0.1)
    var v2 uint = 853
    var v3 int8 = 121
    if true {
        var v18 uint = (^(v2))
        fmt.Println(v1, v18)
        _ = v18
    } else {
        for fi19 := 0; fi19 < 19; fi19++ {
            switch (fi19 * fi19) {
            case -2147483648:
                v2 += v2
            case 0:
                v20 := GenS3{F0: 65535, F1: -66, F2: 4294967295}
                _ = v20
            case 2147483647:
                var v21 bool = (!(!(2147483647 == 548)))
                _ = v21
            }
            switch 4294967295 {
            case 9223372036854775807:
                v22 := GenS2{F0: 191, F1: true, F2: -1}
                _ = v22
            case 32768:
                v23 := []int{2147483647, fi19, fi19}
                _ = v23
            case 128:
                fmt.Println(v2)
            default:
                var v24 uint32 = 0
                _ = v24
            }
            _ = fi19
        }
        v2 += v2
        v25 := []int{-165}
        _ = v25
    }
    v3 = v3
    if true {
        fmt.Println(v3, v3, v3)
        v1 += v1
        v1 = float32(1.5)
    } else {
        for fi26 := 0; fi26 < 14; fi26++ {
            switch (((("\n\t" + "a") + "foo bar") + ((("a" + "") + "") + "hello")) + "a") {
            case "foo bar":
                fmt.Println(v1)
            }
            fmt.Println(v3, v2)
            _ = fi26
        }
    }
    fmt.Println(v2)
    for fi27 := 0; fi27 < 20; fi27++ {
        if ((fi27 | (fi27 | fi27)) != -2147483648) {
            v3 = (v3 ^ v3)
            if (v1 < v1) {
                var v28 bool = false
                fmt.Println(v3, v1)
                var v29 uint16 = 1
                _ = v28
                _ = v29
            } else {
                v30 := []int{fi27}
                _ = v30
            }
            v2 += v2
        }
        if false {
            v31 := GenS1{F0: ((^(v3)) >= (v3 & 0)), F1: fi27, F2: 124}
            _ = v31
        }
        _ = fi27
    }
    v32 := GenS4{F0: 128}
    v33 := []int{2147483647, 4294967295, -2147483648}
    if false {
        for fi34 := 0; fi34 < 7; fi34++ {
            v35 := GenS4{F0: (fi34 ^ 4294967295)}
            v3 = (v3 - v3)
            _ = fi34
            _ = v35
        }
        if (!false) {
            v36 := GenS1{F0: true, F1: 128, F2: -675}
            v33 = append(v33, -1)
            v37 := GenS5{F0: 42, F1: ((v1 + (float32((1e10 - 3.14159)) - 0.1)) < v1)}
            _ = v36
            _ = v37
        } else {
            if (true && false) {
                v3 -= v3
            } else {
                var v38 uint8 = 0
                switch "\n\t" {
                case "\n\t":
                    for fi39 := 0; fi39 < 7; fi39++ {
                        v40 := GenS6{F0: 1, F1: 4294967295, F2: (-(v3))}
                        var v41 int32 = -118
                        _ = fi39
                        _ = v40
                        _ = v41
                    }
                case "a":
                    if ((v2 + 360) != v2) {
                        v38 = (v38 - (v38 & v38))
                        fmt.Println(v3)
                    } else {
                        for fi42 := 0; fi42 < 12; fi42++ {
                            var v43 bool = (4294967295 >= 770)
                            fmt.Println(v43, v43)
                            _ = fi42
                            _ = v43
                        }
                        v44 := GenS2{F0: 1, F1: false, F2: -427}
                        for fi45 := 0; fi45 < 11; fi45++ {
                            if ((v2 << ((v2) & 63)) >= (v2 * (v2 + (v2 * v2)))) {
                                v46 := GenS5{F0: 65535, F1: ((true || (-9 <= 42)) && (18446744073709551615 < 42))}
                                v2 = 2
                                if (65535 != 511) {
                                    fmt.Println(fi45, v2, fi45)
                                    v2 = 65535
                                    if true {
                                        var v47 string = ("\n\t" + "a")
                                        _ = v47
                                    }
                                }
                                _ = v46
                            }
                            v3 -= v3
                            for fi48 := 0; fi48 < 17; fi48++ {
                                v49 := GenS6{F0: 973, F1: 2, F2: (^(v3))}
                                if (false || false) {
                                    fmt.Println(fi48, fi45)
                                }
                                _ = fi48
                                _ = v49
                            }
                            _ = fi45
                        }
                        _ = v44
                    }
                case "foo bar":
                    v2 = 65535
                }
                _ = v38
            }
            var v50 int32 = -1
            v51 := GenS3{F0: (v2 % (((v2 - 114)) | 1)), F1: (^(v50)), F2: 0}
            _ = v50
            _ = v51
        }
    }
    v52 := GenS7{F0: 783, F1: genFn1(-1, 0), F2: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v33:", len(v33), v33)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v52
}
