package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64 }
type GenS2 struct { F0 uint64; F1 uint64 }
type GenS3 struct { F0 int8; F1 bool; F2 int64 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint64; F1 int16 }
type GenS6 struct { F0 uint8 }
func genFn0() uint8 {
    var v7 int64 = 9223372036854775807
    if false {
        switch "hello" {
        case "a":
            var v8 float32 = 0.1
            _ = v8
        case "\n\t":
            v7 = 42
        }
    }
    v9 := GenS1{F0: 127, F1: 9223372036854775807}
    switch 305 {
    case 77:
        v10 := GenS2{F0: 1, F1: 824}
        _ = v10
    case -9223372036854775808:
        var v11 int16 = 32767
        _ = v11
    case -601:
        if ((v7 ^ 385) > v7) {
            for fi12 := 0; fi12 < 20; fi12++ {
                if (0 >= -32768) {
                    v13 := GenS2{F0: 42, F1: 18446744073709551615}
                    v14 := GenS2{F0: 0, F1: 600}
                    var v15 int64 = (v7 / (((v7 >> ((18446744073709551615) & 63))) | 1))
                    _ = v13
                    _ = v14
                    _ = v15
                }
                v7 = v7
                _ = fi12
            }
        } else {
            fmt.Println(v7)
            v16 := []int{-9223372036854775808, 9223372036854775807, 32768}
            switch "" {
            case "":
                v7 -= v7
            }
            _ = v16
        }
    default:
        v7 += v7
    }
    v7 = (v7 * (v7 + v7))
    _ = v7
    _ = v9
    return 127
}
func genFn1(p0 int32, p1 int8) int8 {
    var v17 uint32 = 2147483647
    switch "a" {
    case "":
        v18 := []int{-732}
        _ = v18
    case "\n\t":
        switch -990 {
        case -9223372036854775808:
            v17 -= (v17 << ((883) & 63))
        }
    }
    v19 := GenS3{F0: p1, F1: (1e10 <= 0.0), F2: int64(v17)}
    v17 = 838
    v20 := []int{2147483647, int(v17), -363}
    _ = p0
    _ = p1
    _ = v17
    _ = v19
    _ = v20
    return -128
}

func main() {
    var v1 int64 = -937
    var v2 int = 32768
    v3 := make(map[string]string)
    v4 := []uint32{2147483647, 0}
    var v5 int8 = -128
    v6 := []string{(("a" + "\n\t") + "foo bar"), "foo bar"}
    var v21 int64 = 233
    v1 = v1
    var v22 uint16 = 0
    if true {
        v3["a"] = "a"
        switch 4294967295 {
        case 127:
            switch "foo bar" {
            case "hello":
                v1 -= (v21 | v1)
            case "foo bar":
                _ = len(v3)
            case "\n\t":
                for fi23 := 0; fi23 < 16; fi23++ {
                    for fi24 := 0; fi24 < 19; fi24++ {
                        var v25 bool = (668 > 136)
                        v26 := GenS4{F0: v22}
                        _ = fi24
                        _ = v25
                        _ = v26
                    }
                    switch (fi23 / ((fi23) | 1)) {
                    case 494:
                        var v27 uint32 = 4294967295
                        _ = v27
                    case 128:
                        v28 := GenS1{F0: (v5 % (((v5 / ((v5) | 1))) | 1)), F1: (v21 | v1)}
                        _ = v28
                    case 1:
                        var v29 float32 = 1.5
                        _ = v29
                    default:
                        if false {
                            v30 := v3["hello"]
                            if ((-(v2)) >= (v2 >> ((0) & 63))) {
                                var v31 int16 = 0
                                if true {
                                    v32 := GenS4{F0: 2}
                                    var v33 int64 = (v1 ^ 42)
                                    _ = v32
                                    _ = v33
                                } else {
                                    for fi34 := 0; fi34 < 10; fi34++ {
                                        v35 := GenS4{F0: v22}
                                        switch "hello" {
                                        case "a":
                                            _ = len(v4)
                                        }
                                        _ = fi34
                                        _ = v35
                                    }
                                }
                                _ = v31
                            } else {
                                v36 := GenS5{F0: 122, F1: -591}
                                if (!true) {
                                    fmt.Println(fi23, v22)
                                    var v37 int64 = (-(v1))
                                    _ = v37
                                }
                                _ = v36
                            }
                            _ = v30
                        }
                    }
                    fmt.Println(v2, v1)
                    _ = fi23
                }
            }
        }
    }
    var v38 float64 = 1e10
    var v39 uint32 = 2
    _ = len(v4)
    var v40 uint16 = v22
    v1 -= (^(v1))
    var v41 int8 = v5
    if false {
        fmt.Println(v40, v21)
        _ = len(v4)
        v40 += v22
    } else {
        if true {
            if (true || false) {
                for fi42 := 0; fi42 < 7; fi42++ {
                    v43 := GenS1{F0: v41, F1: v1}
                    for fi44 := 0; fi44 < 12; fi44++ {
                        v45 := GenS3{F0: (^(v5)), F1: false, F2: v21}
                        if true {
                            v39 = v39
                            v5 += (-(v41))
                        }
                        v40 = v22
                        _ = fi44
                        _ = v45
                    }
                    var v46 int64 = v1
                    _ = fi42
                    _ = v43
                    _ = v46
                }
            }
            v47 := GenS6{F0: 127}
            if (1 > (v2 - (v2 ^ v2))) {
                for fi48 := 0; fi48 < 6; fi48++ {
                    var v49 int64 = v21
                    _ = fi48
                    _ = v49
                }
            } else {
                fmt.Println(v39)
            }
            _ = v47
        }
        if ("\n\t" == (("\n\t" + "") + "\n\t")) {
            for fi50 := 0; fi50 < 1; fi50++ {
                fmt.Println(v38)
                v51 := GenS5{F0: 33, F1: 1000}
                _ = fi50
                _ = v51
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v6
}
