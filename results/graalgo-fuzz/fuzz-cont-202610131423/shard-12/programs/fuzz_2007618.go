package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32; F2 int32 }
type GenS2 struct { F0 int64; F1 string }
type GenS3 struct { F0 int; F1 int16 }
type GenS4 struct { F0 uint; F1 uint8; F2 uint64 }
type GenS5 struct { F0 int16; F1 int64 }
type GenS6 struct { F0 int; F1 bool; F2 uint }
func genFn0(p0 uint64) int16 {
    p0 = (^(p0))
    if (!false) {
        p0 -= p0
        switch -2147483648 {
        case -9223372036854775808:
            v6 := GenS1{F0: 0, F1: 517, F2: -1}
            _ = v6
        case -2147483648:
            if (false && false) {
                p0 = p0
                v7 := GenS2{F0: 42, F1: ("" + "")}
                _ = v7
            }
        default:
            for fi8 := 0; fi8 < 6; fi8++ {
                v9 := GenS1{F0: 1, F1: 0, F2: 1}
                for fi10 := 0; fi10 < 2; fi10++ {
                    if true {
                        v11 := GenS1{F0: -2147483648, F1: 605, F2: -1}
                        v12 := GenS1{F0: 0, F1: -1, F2: 0}
                        switch fi10 {
                        case 1:
                            v13 := []int{(-(fi10))}
                            _ = v13
                        case -402:
                            p0 += 4294967295
                        default:
                            v14 := GenS1{F0: 2147483647, F1: 0, F2: 2147483647}
                            _ = v14
                        }
                        _ = v11
                        _ = v12
                    }
                    switch "\n\t" {
                    case "foo bar":
                        var v15 float64 = 1.5
                        _ = v15
                    }
                    v16 := GenS3{F0: -1, F1: -503}
                    _ = fi10
                    _ = v16
                }
                v17 := GenS4{F0: 654, F1: 174, F2: p0}
                _ = fi8
                _ = v17
                _ = v9
            }
        }
    } else {
        if ((((p0 + p0) >= 42) || true) || ("\n\t" < "foo bar")) {
            fmt.Println(p0)
        } else {
            p0 += (^(p0))
            p0 = p0
        }
        for fi18 := 0; fi18 < 19; fi18++ {
            if (false || (!false)) {
                p0 = p0
                var v19 uint16 = 1
                p0 += p0
                _ = v19
            }
            v20 := GenS4{F0: 2, F1: 27, F2: p0}
            _ = fi18
            _ = v20
        }
    }
    _ = p0
    return -969
}

func main() {
    v1 := make(map[string]uint32)
    v2 := []int{-44}
    v3 := make(map[uint32]string)
    var v4 bool = false
    var v5 uint = 438
    var v21 uint = v5
    if v4 {
        fmt.Println(v5)
        for fi22 := 0; fi22 < 19; fi22++ {
            for fi23 := 0; fi23 < 8; fi23++ {
                v24 := GenS5{F0: -1, F1: 42}
                v25 := GenS3{F0: fi23, F1: 32767}
                if (10 == 0) {
                    if (94 < 4294967295) {
                        switch "a" {
                        case "a":
                            var v26 int32 = 0
                            _ = v26
                        case "hello":
                            v5 = v21
                        default:
                            if (323 <= 813) {
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                var v27 int = (fi23 ^ 32767)
                                _ = v27
                            } else {
                                var v28 uint8 = 42
                                _ = v28
                            }
                        }
                        switch (fi23 ^ fi23) {
                        case 555:
                            switch fi23 {
                            case -735:
                                v5 = (v21 % ((v5) | 1))
                            case -479:
                                var v29 uint32 = 792
                                _ = v29
                            case 32767:
                                fmt.Println(v4, v21, fi23)
                            }
                        case -262:
                            v30 := GenS6{F0: -2147483648, F1: false, F2: v21}
                            var v31 string = "hello"
                            _ = v30
                            _ = v31
                        default:
                            v21 += (v5 / ((v5) | 1))
                        }
                        fmt.Println(v5, v21)
                        v5 = (v5 & (v5 ^ 818))
                        v32 := GenS2{F0: 0, F1: "foo bar"}
                        v33 := GenS2{F0: 9223372036854775807, F1: "hello"}
                        if v4 {
                            v34 := GenS5{F0: 664, F1: 42}
                            var v35 int32 = 1
                            _ = v34
                            _ = v35
                        } else {
                            if v4 {
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                var v36 int16 = -1
                                _ = v36
                            } else {
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                var v37 uint8 = 255
                                v5 += v5
                                _ = v37
                            }
                        }
                        _ = v32
                        _ = v33
                    } else {
                        var v38 int16 = 0
                        fmt.Println(fi23, v21)
                        v38 += int16(fi23)
                        _ = v38
                    }
                    var v39 float32 = 3.14159
                    _ = v39
                } else {
                    _ = len(v2)
                }
                var v40 float32 = (3.14159 - 0.0)
                _ = fi23
                _ = v24
                _ = v25
                _ = v40
            }
            if len(v2) > 0 {
                v2[0] = (fi22 % (((fi22 << ((v21) & 63))) | 1))
            }
            v41 := GenS5{F0: -1, F1: 397}
            _ = fi22
            _ = v41
        }
    } else {
        v4 = ((1 != 42) && (true || ((!(740 == 9223372036854775807)) || (2 < 18446744073709551615))))
    }
    v42 := GenS6{F0: 32767, F1: (((18446744073709551615 > 233) || v4) || (3.14159 == 1e10)), F2: (v5 % ((v5) | 1))}
    switch 32767 {
    case 128:
        v43 := GenS1{F0: -2147483648, F1: 42, F2: 2147483647}
        _ = v43
    default:
        v21 += (v5 + (v21 >> ((18446744073709551615) & 63)))
    }
    fmt.Println(v21, v21, v4)
    if true {
        v5 = (v5 / (((v21 | 255)) | 1))
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    v5 += 42
    var v44 uint8 = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v42
    _ = v44
    _ = v5
}
