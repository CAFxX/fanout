package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string; F2 uint64 }
type GenS2 struct { F0 uint; F1 string }
type GenS3 struct { F0 uint16; F1 uint; F2 uint }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint; F1 int64; F2 int64 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 int8; F1 int16; F2 int64 }
type GenS8 struct { F0 uint; F1 int; F2 uint }
type GenS9 struct { F0 bool; F1 int32 }
func genFn0(p0 int8) int16 {
    p0 -= p0
    p0 = p0
    _ = p0
    return -787
}
func genFn1(p0 int) int32 {
    var v6 uint32 = 2147483647
    v7 := []int{(p0 & p0), p0, p0}
    for fi8 := 0; fi8 < 13; fi8++ {
        var v9 uint64 = 486
        _ = fi8
        _ = v9
    }
    v10 := GenS1{F0: 32767, F1: "foo bar", F2: 42}
    _ = p0
    _ = v10
    _ = v6
    _ = v7
    return 2147483647
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[string]uint32)
    var v3 float32 = 0.0
    v4 := []uint32{73, 2147483647}
    var v5 int = 1
    if true {
        var v11 bool = false
        v12 := GenS1{F0: 0, F1: "a", F2: 4294967295}
        v13 := GenS2{F0: 42, F1: "hello"}
        _ = v11
        _ = v12
        _ = v13
    } else {
        if (!false) {
            var v14 string = ("\n\t" + "a")
            v2[v14] = 65535
            if false {
                if (false || true) {
                    for fi15 := 0; fi15 < 8; fi15++ {
                        if true {
                            for fi16 := 0; fi16 < 14; fi16++ {
                                v17 := GenS3{F0: 348, F1: 2, F2: 4294967295}
                                _ = fi16
                                _ = v17
                            }
                        } else {
                            switch v14 {
                            case "a":
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            case "foo bar":
                                var v18 int = -746
                                _ = v18
                            case "":
                                v3 -= v3
                            }
                        }
                        fmt.Println(fi15, v5, v3)
                        _ = fi15
                    }
                    if false {
                        var v19 uint32 = 65535
                        _ = v19
                    }
                    switch 0 {
                    case 32768:
                        fmt.Println(v3, v14, v5)
                    default:
                        v5 -= v5
                    }
                }
                if false {
                    switch v14 {
                    case "a":
                        if len(v4) > 0 {
                            v4[0] = 2
                        }
                    case "foo bar":
                        v20 := GenS4{F0: -1}
                        _ = v20
                    case "hello":
                        v5 -= (v5 - (v5 * (v5 / (((v5 | (v5 >> ((4294967295) & 63)))) | 1))))
                    }
                    fmt.Println(v3, v3, v14)
                    switch "a" {
                    case "a":
                        switch "" {
                        case "\n\t":
                            var v21 float32 = (((0.0 * float32(1.5)) / v3) * v3)
                            _ = v21
                        case "":
                            for fi22 := 0; fi22 < 6; fi22++ {
                                for fi23 := 0; fi23 < 13; fi23++ {
                                    v14 = ""
                                    if true {
                                        v2[(v14 + v14)] = 0
                                    }
                                    _ = fi23
                                }
                                _ = fi22
                            }
                        default:
                            switch (v14 + "hello") {
                            case "foo bar":
                                var v24 uint16 = 439
                                _ = v24
                            default:
                                for fi25 := 0; fi25 < 14; fi25++ {
                                    var v26 float64 = float64(v3)
                                    for fi27 := 0; fi27 < 4; fi27++ {
                                        v28 := GenS5{F0: 42, F1: 555, F2: 42}
                                        switch (v5 * fi27) {
                                        case -97:
                                            v14 = ("" + ((v14 + v14) + v14))
                                            switch (("foo bar" + v14) + (v14 + "\n\t")) {
                                            case "hello":
                                                var v29 uint8 = 42
                                                _ = v29
                                            default:
                                                fmt.Println(v14)
                                            }
                                            v5 = fi25
                                            _ = len(v4)
                                        }
                                        _ = fi27
                                        _ = v28
                                    }
                                    for fi30 := 0; fi30 < 2; fi30++ {
                                        v31 := GenS6{F0: v14}
                                        var v32 int32 = 2147483647
                                        v33 := GenS3{F0: 32767, F1: 18446744073709551615, F2: 0}
                                        _ = fi30
                                        _ = v31
                                        _ = v32
                                        _ = v33
                                    }
                                    _ = fi25
                                    _ = v26
                                }
                            }
                        }
                    default:
                        delete(v2, v14)
                    }
                } else {
                    v3 = 1e10
                    if false {
                        if ((492 >= 124) && false) {
                            if (18446744073709551615 < 2) {
                                v34 := GenS6{F0: ("" + ("a" + v14))}
                                _ = v34
                            }
                            if false {
                                v35 := GenS1{F0: 42, F1: "a", F2: 802}
                                v14 = ("foo bar" + v14)
                                if (0.1 > (3.14159 / (1e10 - 1.5))) {
                                    v36 := GenS7{F0: 61, F1: -476, F2: -270}
                                    fmt.Println(v5)
                                    _ = v36
                                }
                                _ = v35
                            } else {
                                _ = len(v2)
                            }
                        }
                    } else {
                        switch v5 {
                        case -1:
                            v37 := GenS8{F0: 0, F1: v5, F2: 789}
                            _ = v37
                        }
                    }
                    for fi38 := 0; fi38 < 10; fi38++ {
                        var v39 uint64 = 2
                        var v40 int8 = 1
                        _ = fi38
                        _ = v39
                        _ = v40
                    }
                }
            } else {
                var v41 uint64 = 42
                var v42 uint64 = 18446744073709551615
                _ = v41
                _ = v42
            }
            _ = v14
        }
    }
    v3 += (v3 + 0.0)
    v3 = v3
    var v43 uint16 = 693
    if (2 <= 42) {
        fmt.Println(v5, v3)
    } else {
        fmt.Println(v3)
    }
    switch "hello" {
    case "\n\t":
        var v44 int32 = 2147483647
        _ = v44
    case "foo bar":
        fmt.Println(v3, v43, v3)
    default:
        var v45 uint = 476
        _ = v45
    }
    var v46 uint8 = 42
    if (1.5 < ((v3 - 1e10) * 0.1)) {
        var v47 uint8 = 42
        v48 := GenS4{F0: -128}
        v49 := GenS9{F0: (true || true), F1: -1}
        _ = v47
        _ = v48
        _ = v49
    }
    var v50 int = (v5 / ((913) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v43:", v43)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v46
    _ = v5
    _ = v50
}
