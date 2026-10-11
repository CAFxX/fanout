package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint64; F1 int; F2 uint }
type GenS3 struct { F0 int64; F1 int64; F2 int16 }
func genFn0(p0 uint32, p1 int32, p2 int64) uint16 {
    var v7 int8 = -1
    var v8 bool = (!false)
    if v8 {
        if v8 {
            if v8 {
                switch "\n\t" {
                case "foo bar":
                    for fi9 := 0; fi9 < 12; fi9++ {
                        var v10 float32 = 0.0
                        var v11 uint32 = p0
                        _ = fi9
                        _ = v10
                        _ = v11
                    }
                case "hello":
                    var v12 uint = 2
                    _ = v12
                case "\n\t":
                    v13 := []int{-9223372036854775808, 2147483647}
                    _ = v13
                }
                for fi14 := 0; fi14 < 15; fi14++ {
                    switch "hello" {
                    case "a":
                        p2 += (^(p2))
                    case "hello":
                        for fi15 := 0; fi15 < 19; fi15++ {
                            if ((345 != 42) && v8) {
                                var v16 uint64 = 2
                                fmt.Println(fi15, p0)
                                v17 := []int{int(p1), fi14, (fi15 / (((fi15 % ((fi15) | 1))) | 1)), int(p2)}
                                if len(v17) >= 2 {
                                    v17 = v17[1:len(v17)]
                                }
                                _ = v16
                                _ = v17
                            } else {
                                v18 := []int{(fi14 / (((fi14 | fi14)) | 1)), 584, (fi15 >> ((2) & 63)), fi15}
                                if true {
                                    fmt.Println(fi14)
                                    switch "\n\t" {
                                    case "hello":
                                        fmt.Println(fi15)
                                        v19 := GenS1{F0: 32767}
                                        _ = v19
                                    }
                                    if len(v18) > 0 {
                                        v18[0] = (-(fi15))
                                    }
                                    var v20 uint8 = 2
                                    _ = v20
                                }
                                switch (("foo bar" + ((("" + "foo bar") + "foo bar") + "\n\t")) + "") {
                                case "foo bar":
                                    fmt.Println(p0, v7)
                                    var v21 float64 = ((((-2.25 + 0.1) * (1e10 + 1.5)) - 0.0) / 1e10)
                                    _ = v21
                                case "":
                                    var v22 uint32 = p0
                                    switch "hello" {
                                    case "":
                                        if (-2.25 != 0.1) {
                                            if (fi15 < (^(fi15))) {
                                                if len(v18) > 0 {
                                                    v18[0] = 2147483647
                                                }
                                                var v23 uint = 42
                                                fmt.Println(p2, v23, fi15)
                                                _ = v23
                                            }
                                            var v24 float64 = 1.5
                                            if v8 {
                                                v24 = float64((0.0 + 1.5))
                                            }
                                            _ = v24
                                        } else {
                                            p1 -= int32(fi15)
                                        }
                                    case "a":
                                        v25 := GenS1{F0: 1}
                                        _ = v25
                                    }
                                    switch "a" {
                                    case "a":
                                        var v26 uint32 = p0
                                        _ = v26
                                    case "foo bar":
                                        p0 -= (^(p0))
                                    case "\n\t":
                                        if false {
                                            v27 := GenS1{F0: 1000}
                                            _ = v27
                                        }
                                    default:
                                        p2 = 1
                                    }
                                    _ = v22
                                case "hello":
                                    v28 := GenS2{F0: 18446744073709551615, F1: fi15, F2: 903}
                                    _ = v28
                                }
                                v8 = (v8 && true)
                                p2 = 0
                                _ = v18
                            }
                            _ = fi15
                        }
                    default:
                        if false {
                            v29 := []int{}
                            for fi30 := 0; fi30 < 11; fi30++ {
                                p2 -= p2
                                var v31 int8 = 42
                                _ = fi30
                                _ = v31
                            }
                            v32 := GenS1{F0: 1000}
                            _ = v29
                            _ = v32
                        }
                    }
                    for fi33 := 0; fi33 < 9; fi33++ {
                        fmt.Println(fi33)
                        v34 := GenS1{F0: -1}
                        _ = fi33
                        _ = v34
                    }
                    for fi35 := 0; fi35 < 19; fi35++ {
                        switch fi14 {
                        case 920:
                            v8 = false
                        case -939:
                            v36 := []int{(fi14 - 199), (fi14 * fi14)}
                            _ = v36
                        }
                        var v37 int32 = (^(p1))
                        _ = fi35
                        _ = v37
                    }
                    _ = fi14
                }
                v38 := GenS2{F0: 547, F1: 9223372036854775807, F2: 65535}
                _ = v38
            }
            fmt.Println(v8, p1)
        }
    }
    v39 := []int{9223372036854775807, 4294967295}
    var v40 uint32 = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v39
    _ = v40
    _ = v7
    _ = v8
    return 65535
}
func genFn1(p0 uint) string {
    var v41 uint64 = 251
    var v42 float32 = 0.0
    var v43 int8 = -17
    var v44 uint32 = 2147483648
    v45 := GenS1{F0: 0}
    _ = p0
    _ = v41
    _ = v42
    _ = v43
    _ = v44
    _ = v45
    return ((("hello" + (("a" + (("hello" + "\n\t") + ("a" + "a"))) + "foo bar")) + "foo bar") + "hello")
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[string]bool)
    var v3 int = 2147483647
    v4 := []string{}
    var v5 int8 = 1
    v6 := []int{(v3 ^ (v3 / ((v3) | 1))), 4294967295, 2147483647, 786}
    var v46 int16 = 0
    if false {
        v46 = -32768
    }
    if true {
        v4 = append(v4, "")
    }
    v46 = 943
    v47 := GenS2{F0: 880, F1: v3, F2: 65535}
    v48 := GenS1{F0: (v46 % (((v46 % ((42) | 1))) | 1))}
    switch (-(v3)) {
    case 127:
        var v49 int = -1
        _ = v49
    case 9223372036854775807:
        v50 := GenS3{F0: int64(v46), F1: 537, F2: (v46 ^ (v46 - 0))}
        _ = v50
    }
    if (0 != 1) {
        if false {
            v51 := GenS2{F0: 2, F1: (v3 ^ (v3 & 32768)), F2: 255}
            v3 -= (v3 - (v3 / ((9223372036854775807) | 1)))
            _ = v51
        }
        for fi52 := 0; fi52 < 19; fi52++ {
            var v53 bool = ((!(false && (!true))) || (127 == 42))
            _ = fi52
            _ = v53
        }
    }
    var v54 uint64 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("v54:", v54)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v47
    _ = v48
    _ = v5
    _ = v54
    _ = v6
}
