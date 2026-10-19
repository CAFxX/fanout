package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64; F2 bool }
type GenS2 struct { F0 uint; F1 int }
type GenS3 struct { F0 uint16; F1 uint32; F2 uint8 }
type GenS4 struct { F0 string; F1 int32 }

func main() {
    v1 := make(map[string]int)
    v2 := []uint32{42, 2147483648, 0, 42}
    v3 := []int{-788, 617, -905, -9223372036854775808}
    v4 := make(map[uint32]string)
    var v5 int32 = 0
    var v6 uint64 = 4294967295
    v2 = append(v2, 2147483648)
    v5 = v5
    for fi7 := 0; fi7 < 20; fi7++ {
        v5 = (v5 >> ((817) & 63))
        if (v5 < -2147483648) {
            for fi8 := 0; fi8 < 17; fi8++ {
                fmt.Println(fi7)
                v5 = 42
                var v9 int8 = 14
                _ = fi8
                _ = v9
            }
        } else {
            for fi10 := 0; fi10 < 4; fi10++ {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = fi10
            }
        }
        var v11 uint16 = 42
        _ = fi7
        _ = v11
    }
    var v12 uint32 = 2
    fmt.Println(v5, v6, v12)
    if (((0 >= 411) && (v5 == 2147483647)) || (575 < 9223372036854775807)) {
        var v13 uint32 = (v12 / (((v12 ^ 0)) | 1))
        var v14 uint = 255
        if ((false || true) && (false || false)) {
            if (-82 != -34) {
                var v15 int = -9223372036854775808
                if ((42 < 0) || false) {
                    switch (-(v15)) {
                    case 42:
                        v16 := GenS1{F0: v14, F1: -233, F2: true}
                        _ = v16
                    }
                    switch "\n\t" {
                    case "":
                        for fi17 := 0; fi17 < 10; fi17++ {
                            var v18 bool = ((false && (v5 <= (v5 >> (((v14 << (((^(v14))) & 63))) & 63)))) || ((v15 < 678) || false))
                            _ = fi17
                            _ = v18
                        }
                    case "hello":
                        v19 := GenS2{F0: v14, F1: v15}
                        _ = v19
                    case "\n\t":
                        for fi20 := 0; fi20 < 4; fi20++ {
                            fmt.Println(v13)
                            v12 -= (v12 ^ v13)
                            _ = fi20
                        }
                    }
                    fmt.Println(v14)
                }
                _ = v15
            }
        }
        _ = v13
        _ = v14
    } else {
        var v21 uint32 = 2147483648
        _ = v21
    }
    _ = len(v1)
    for fi22 := 0; fi22 < 19; fi22++ {
        v12 = (v12 % ((v12) | 1))
        _ = fi22
    }
    v12 = (v12 ^ (^(v12)))
    var v23 float64 = (-2.25 - 1e10)
    switch 32767 {
    case 102:
        var v24 uint16 = 255
        _ = v24
    }
    var v25 bool = true
    var v26 int16 = 928
    v23 += v23
    if (992 <= 42) {
        var v27 string = (("" + "") + "foo bar")
        switch v27 {
        case "a":
            if true {
                v28 := GenS3{F0: 65535, F1: 2147483647, F2: 29}
                v29 := v4[v12]
                var v30 int32 = (^(v5))
                _ = v28
                _ = v29
                _ = v30
            }
        case "\n\t":
            v31 := GenS2{F0: 0, F1: -2147483648}
            _ = v31
        case "hello":
            for fi32 := 0; fi32 < 5; fi32++ {
                if v25 {
                    if v25 {
                        switch v27 {
                        case "\n\t":
                            switch "hello" {
                            case "\n\t":
                                v2 = append(v2, (v12 % ((v12) | 1)))
                            case "":
                                v5 -= v5
                            case "foo bar":
                                if v25 {
                                    v27 = v27
                                } else {
                                    var v33 int16 = 42
                                    if len(v3) > 0 {
                                        v3[0] = int(v5)
                                    }
                                    fmt.Println(v25, v5)
                                    _ = v33
                                }
                            }
                        }
                    }
                    if v25 {
                        for fi34 := 0; fi34 < 1; fi34++ {
                            _ = len(v2)
                            var v35 int32 = v5
                            fmt.Println(v27, v26)
                            fmt.Println(fi34, v5)
                            v36 := v4[v12]
                            _ = fi34
                            _ = v35
                            _ = v36
                        }
                        v6 = (v6 >> ((917) & 63))
                    }
                } else {
                    v23 = v23
                    var v37 uint = 1
                    _ = v37
                }
                if (v12 > (v12 >> ((4294967295) & 63))) {
                    switch v27 {
                    case "\n\t":
                        v3 = append(v3, (fi32 / ((-119) | 1)))
                    case "foo bar":
                        var v38 string = (("" + v27) + (v27 + "\n\t"))
                        _ = v38
                    default:
                        if v25 {
                            v39 := v1[""]
                            fmt.Println(v27)
                            _ = v39
                        } else {
                            var v40 string = v27
                            _ = v40
                        }
                    }
                } else {
                    switch v27 {
                    case "hello":
                        fmt.Println(v27, fi32, v6)
                    case "a":
                        var v41 float32 = 0.1
                        _ = v41
                    case "":
                        if true {
                            switch "hello" {
                            case "\n\t":
                                v12 = (v12 + (v12 * v12))
                            case "a":
                                for fi42 := 0; fi42 < 17; fi42++ {
                                    v6 = v6
                                    _ = fi42
                                }
                            case "foo bar":
                                var v43 uint = 2
                                _ = v43
                            }
                            v44 := GenS3{F0: 2, F1: v12, F2: 95}
                            for fi45 := 0; fi45 < 1; fi45++ {
                                var v46 string = v27
                                fmt.Println(v12, fi45)
                                switch (fi45 + 1) {
                                case -1:
                                    fmt.Println(v12, v46)
                                }
                                _ = fi45
                                _ = v46
                            }
                            _ = v44
                        }
                    default:
                        var v47 uint32 = (v12 | (^(v12)))
                        _ = v47
                    }
                    v27 = v27
                    v26 -= 32767
                }
                var v48 int64 = -1
                _ = fi32
                _ = v48
            }
        }
        _ = v27
    } else {
        _ = len(v4)
        var v49 float32 = -2.25
        v6 -= v6
        _ = v49
    }
    v50 := GenS4{F0: "a", F1: v5}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v23
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v6
}
