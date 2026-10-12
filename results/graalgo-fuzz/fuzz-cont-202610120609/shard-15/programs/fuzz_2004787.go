package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8; F2 int64 }
type GenS2 struct { F0 uint; F1 uint32 }
type GenS3 struct { F0 uint32; F1 uint }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int; F1 bool; F2 uint64 }
type GenS6 struct { F0 int64; F1 int16 }
type GenS7 struct { F0 int16; F1 uint32 }
func genFn0(p0 int64) int16 {
    p0 = (-(p0))
    p0 = (p0 & -9223372036854775808)
    v7 := []int{-600, -43, -941}
    _ = p0
    _ = v7
    return 1
}
func genFn1(p0 uint8, p1 int64) uint16 {
    var v8 bool = false
    v9 := GenS1{F0: 127, F1: p0, F2: (p1 - (p1 * (p1 >> ((18446744073709551615) & 63))))}
    v10 := GenS2{F0: 679, F1: 42}
    v11 := []int{127, 128, 256}
    v8 = ((p0 / (((p0 + p0)) | 1)) >= 127)
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v8
    _ = v9
    return 42
}

func main() {
    v1 := make(map[int]int)
    var v2 int8 = 104
    var v3 float32 = 1.5
    var v4 int = 128
    var v5 int8 = v2
    v6 := make(map[int]uint32)
    if true {
        v3 = v3
        switch 1 {
        case 9223372036854775807:
            switch v4 {
            case -580:
                if (!true) {
                    delete(v1, (v4 + (v4 / ((v4) | 1))))
                    for fi12 := 0; fi12 < 20; fi12++ {
                        fmt.Println(v5, v5, v2)
                        fmt.Println(v5)
                        for fi13 := 0; fi13 < 10; fi13++ {
                            var v14 int64 = -9223372036854775808
                            if (false || true) {
                                if false {
                                    v15 := GenS2{F0: 412, F1: 2147483648}
                                    v5 -= v2
                                    if true {
                                        v6[-2147483648] = 2
                                        v16 := GenS1{F0: (v2 << ((4294967295) & 63)), F1: 229, F2: (^(v14))}
                                        fmt.Println(v5, fi13)
                                        var v17 float64 = ((((1e10 * -2.25) * 3.14159) * 0.1) - -2.25)
                                        _ = v16
                                        _ = v17
                                    } else {
                                        if (v14 > 42) {
                                            fmt.Println(fi13, fi12, v3)
                                        }
                                    }
                                    _ = v15
                                } else {
                                    if (4294967295 != 0) {
                                        delete(v1, (fi13 * 4294967295))
                                        v18 := GenS3{F0: 2147483648, F1: 6}
                                        _ = v18
                                    } else {
                                        var v19 bool = true
                                        v20 := v1[fi12]
                                        _ = v19
                                        _ = v20
                                    }
                                }
                            }
                            v21 := []int{}
                            if true {
                                v22 := GenS4{F0: 160}
                                var v23 uint64 = 1
                                _ = v22
                                _ = v23
                            }
                            _ = fi13
                            _ = v14
                            _ = v21
                        }
                        _ = fi12
                    }
                }
            case 128:
                v24 := []int{(v4 | v4), v4, 127, v4}
                _ = v24
            }
        case -688:
            v2 -= v2
        case 127:
            var v25 string = "a"
            _ = v25
        default:
            var v26 uint8 = 128
            _ = v26
        }
        if true {
            v27 := GenS5{F0: (v4 + 9223372036854775807), F1: true, F2: 42}
            for fi28 := 0; fi28 < 5; fi28++ {
                var v29 int32 = -929
                v3 += (0.1 * v3)
                v29 = v29
                _ = fi28
                _ = v29
            }
            var v30 int16 = -101
            _ = v27
            _ = v30
        }
    } else {
        v31 := GenS6{F0: 84, F1: 42}
        _ = v31
    }
    switch "\n\t" {
    case "\n\t":
        v2 = (v2 / ((v5) | 1))
    case "hello":
        var v32 uint16 = 121
        _ = v32
    case "foo bar":
        if (true || false) {
            v5 = v2
            var v33 int = (v4 ^ v4)
            _ = v33
        } else {
            v34 := GenS5{F0: (v4 >> ((42) & 63)), F1: (true && (v3 <= 1.5)), F2: 42}
            if true {
                v35 := GenS7{F0: -185, F1: 4294967295}
                _ = v35
            }
            var v36 uint32 = 1
            _ = v34
            _ = v36
        }
    }
    fmt.Println(v5, v5, v3)
    for fi37 := 0; fi37 < 13; fi37++ {
        switch "foo bar" {
        case "":
            v4 -= fi37
        }
        var v38 uint16 = 255
        _ = fi37
        _ = v38
    }
    var v39 int32 = 2147483647
    for fi40 := 0; fi40 < 3; fi40++ {
        switch (fi40 & 4294967295) {
        case -391:
            var v41 int32 = v39
            _ = v41
        default:
            fmt.Println(v39, v4, v3)
        }
        _ = fi40
    }
    fmt.Println(v4)
    for fi42 := 0; fi42 < 1; fi42++ {
        v43 := []int{(fi42 / (((^(fi42))) | 1)), fi42}
        v5 += v5
        if (!(32767 == 42)) {
            v3 = v3
        }
        _ = fi42
        _ = v43
    }
    var v44 int = 1
    var v45 int8 = v5
    var v46 int8 = 0
    v47 := []int{v44, v44, v4}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("slicelen v47:", len(v47), v47)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v44
    _ = v45
    _ = v46
    _ = v47
    _ = v5
    _ = v6
}
