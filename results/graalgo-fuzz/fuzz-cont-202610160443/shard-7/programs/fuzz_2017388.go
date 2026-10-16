package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int; F1 uint32 }
type GenS3 struct { F0 uint64; F1 int }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int32; F1 bool; F2 uint16 }
type GenS6 struct { F0 uint; F1 uint; F2 int8 }
func genFn0(p0 uint64, p1 int16, p2 string) uint32 {
    v7 := []int{807, 9223372036854775807, -2147483648}
    var v8 string = p2
    if (false && true) {
        p0 = p0
        fmt.Println(p1, p0, v8)
    }
    var v9 int = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return 2147483647
}
func genFn1(p0 string, p1 int64) int {
    v10 := []int{4294967295, 9223372036854775807, 128, 32768}
    var v11 float32 = 1e10
    p1 -= p1
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    return -9223372036854775808
}

func main() {
    v1 := make(map[string]int)
    v2 := []int{-9223372036854775808, 32768, 9223372036854775807, -2147483648}
    v3 := make(map[int]bool)
    v4 := make(map[uint32]bool)
    var v5 string = ("foo bar" + ("hello" + ("\n\t" + (((("" + "foo bar") + "a") + "hello") + "a"))))
    var v6 uint32 = 42
    v12 := GenS1{F0: -128}
    for fi13 := 0; fi13 < 19; fi13++ {
        _ = len(v1)
        v14 := GenS2{F0: fi13, F1: v6}
        _ = fi13
        _ = v14
    }
    for fi15 := 0; fi15 < 13; fi15++ {
        for fi16 := 0; fi16 < 16; fi16++ {
            v17 := GenS1{F0: 1}
            var v18 float32 = 3.14159
            if (!true) {
                switch -2147483648 {
                case 32768:
                    var v19 string = "hello"
                    _ = v19
                case 9223372036854775807:
                    v5 = ((((v5 + "") + "") + v5) + v5)
                case 128:
                    v6 = v6
                default:
                    v18 += (((1e10 - 0.0) / (v18 / v18)) + v18)
                }
            } else {
                if (0 != -2147483648) {
                    var v20 bool = true
                    v21 := GenS3{F0: 42, F1: 32767}
                    if v20 {
                        var v22 uint8 = 0
                        var v23 uint32 = (v6 & 348)
                        fmt.Println(v20, v20, v18)
                        _ = v22
                        _ = v23
                    } else {
                        var v24 int16 = 32767
                        var v25 int32 = 2147483647
                        if ((0 <= 2) || (!((float32(1.5) > ((v18 / 1.0) / v18)) && v20))) {
                            v5 = v5
                            var v26 int = 4294967295
                            v27 := GenS4{F0: 42}
                            _ = v26
                            _ = v27
                        }
                        _ = v24
                        _ = v25
                    }
                    v20 = (1 >= 255)
                    switch ((v5 + "\n\t") + (v5 + v5)) {
                    case "":
                        v6 = genFn0(536, 32767, ("foo bar" + v5))
                    case "foo bar":
                        if v20 {
                            var v28 float32 = v18
                            v29 := v4[(v6 ^ v6)]
                            _ = v28
                            _ = v29
                        }
                    case "\n\t":
                        var v30 int16 = -247
                        _ = v30
                    }
                    _ = v20
                    _ = v21
                } else {
                    v5 = (((v5 + v5) + v5) + ("hello" + v5))
                    v31 := v1[("" + v5)]
                    var v32 float64 = (float64(v18) * 0.0)
                    _ = v31
                    _ = v32
                }
                var v33 uint = 1
                _ = v33
            }
            _ = fi16
            _ = v17
            _ = v18
        }
        if false {
            if false {
                v6 += (v6 - v6)
                v1[("a" + v5)] = fi15
            }
            var v34 uint16 = 42
            _ = v34
        }
        v35 := GenS3{F0: 399, F1: (fi15 | genFn1(((v5 + v5) + v5), 616))}
        _ = fi15
        _ = v35
    }
    if false {
        fmt.Println(v5, v6)
        var v36 uint8 = 128
        _ = v36
    }
    v37 := v1[v5]
    switch v5 {
    case "hello":
        if ((9223372036854775807 == 1) || false) {
            v37 = 32768
        }
    default:
        fmt.Println(v6)
    }
    switch (("foo bar" + (v5 + "\n\t")) + "") {
    case "foo bar":
        switch "a" {
        case "\n\t":
            v5 = v5
        case "hello":
            if (float64(3.14159) != -2.25) {
                var v38 bool = true
                v5 = "\n\t"
                for fi39 := 0; fi39 < 12; fi39++ {
                    var v40 uint64 = 18446744073709551615
                    v41 := GenS2{F0: fi39, F1: 2}
                    _ = fi39
                    _ = v40
                    _ = v41
                }
                _ = v38
            }
        }
    case "a":
        fmt.Println(v6)
    case "\n\t":
        switch (v37 - (v37 - genFn1(v5, 711))) {
        case 4294967295:
            for fi42 := 0; fi42 < 1; fi42++ {
                var v43 int64 = 1
                _ = fi42
                _ = v43
            }
        default:
            for fi44 := 0; fi44 < 19; fi44++ {
                v45 := GenS5{F0: 42, F1: true, F2: 132}
                _ = fi44
                _ = v45
            }
        }
    }
    v46 := GenS6{F0: 65535, F1: 1, F2: int8(v6)}
    for fi47 := 0; fi47 < 1; fi47++ {
        var v48 int8 = 42
        v1[v5] = genFn1(((v5 + v5) + v5), 0)
        _ = fi47
        _ = v48
    }
    for fi49 := 0; fi49 < 9; fi49++ {
        v2 = append(v2, (fi49 | v37))
        fmt.Println(v37)
        _ = fi49
    }
    var v50 int = v37
    v51 := GenS4{F0: uint64(v6)}
    var v52 float64 = 0.1
    switch (v5 + ("foo bar" + (v5 + v5))) {
    case "a":
        for fi53 := 0; fi53 < 20; fi53++ {
            v54 := v3[0]
            if v54 {
                fmt.Println(fi53, v52, v54)
                switch v37 {
                case -2147483648:
                    switch 32767 {
                    case -9223372036854775808:
                        v5 = v5
                    case 1:
                        var v55 int32 = 534
                        _ = v55
                    default:
                        if (32 < 1) {
                            var v56 uint32 = (v6 / (((v6 & (^(v6)))) | 1))
                            _ = v56
                        }
                    }
                case -690:
                    for fi57 := 0; fi57 < 8; fi57++ {
                        v2 = append(v2, -1)
                        _ = fi57
                    }
                default:
                    v54 = v54
                }
                v52 += v52
            } else {
                v37 -= v50
            }
            _ = fi53
            _ = v54
        }
    case "":
        v6 -= genFn0(93, 0, v5)
    case "hello":
        switch v37 {
        case 42:
            v58 := v1[v5]
            _ = v58
        case 32767:
            _ = len(v2)
        case 60:
            v52 -= float64(0.0)
        default:
            fmt.Println(v37, v5, v52)
        }
    default:
        var v59 float32 = ((0.0 - (0.0 * 0.1)) + 1.5)
        _ = v59
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v52:", v52)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v46
    _ = v5
    _ = v50
    _ = v51
    _ = v52
    _ = v6
}
