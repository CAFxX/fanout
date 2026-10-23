package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string; F2 uint16 }
type GenS2 struct { F0 uint8; F1 uint32; F2 int16 }
type GenS3 struct { F0 uint16; F1 int16 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 string; F1 uint }
func genFn0(p0 int16, p1 string) bool {
    var v5 int32 = 2147483647
    if (!true) {
        var v6 uint16 = 78
        if ((true || true) || true) {
            for fi7 := 0; fi7 < 4; fi7++ {
                if ((2 > 747) && true) {
                    v5 += (-(v5))
                    var v8 int64 = 9223372036854775807
                    _ = v8
                }
                _ = fi7
            }
            for fi9 := 0; fi9 < 7; fi9++ {
                v10 := []int{fi9, int(v5), fi9}
                if len(v10) > 0 {
                    v10[0] = fi9
                }
                _ = fi9
                _ = v10
            }
        } else {
            if false {
                switch p1 {
                case "foo bar":
                    v6 = (v6 % ((v6) | 1))
                case "hello":
                    v6 += 429
                }
                v11 := []int{-953, 32768, 45, -612}
                _ = v11
            }
        }
        _ = v6
    }
    fmt.Println(p1)
    v12 := GenS1{F0: (p0 ^ p0), F1: "hello", F2: 464}
    v13 := []int{0, -1, 42}
    _ = p0
    _ = p1
    _ = v12
    _ = v13
    _ = v5
    return (!false)
}

func main() {
    var v1 uint32 = 701
    var v2 int8 = 0
    var v3 string = "a"
    var v4 int32 = 0
    v1 += 2147483647
    var v14 float32 = 1.5
    fmt.Println(v3, v4, v14)
    switch v3 {
    case "":
        if false {
            v15 := GenS2{F0: 127, F1: 927, F2: 32767}
            for fi16 := 0; fi16 < 11; fi16++ {
                v17 := []int{127, fi16, 0}
                var v18 int16 = -1
                var v19 string = ("a" + (v3 + v3))
                _ = fi16
                _ = v17
                _ = v18
                _ = v19
            }
            _ = v15
        }
    }
    for fi20 := 0; fi20 < 16; fi20++ {
        for fi21 := 0; fi21 < 3; fi21++ {
            v22 := []int{-9223372036854775808, fi20, -9223372036854775808, fi21}
            fmt.Println(v3, v2, fi20)
            _ = fi21
            _ = v22
        }
        _ = fi20
    }
    if false {
        for fi23 := 0; fi23 < 12; fi23++ {
            v24 := []int{}
            var v25 float64 = ((((0.1 - 3.14159) * 1e10) + 0.0) * -2.25)
            for fi26 := 0; fi26 < 19; fi26++ {
                var v27 uint64 = 4294967295
                switch (fi23 | -877) {
                case 32767:
                    v25 -= ((-2.25 + 3.14159) * float64(v14))
                case 2:
                    if (42 < 42) {
                        v2 = v2
                        var v28 int64 = -823
                        _ = v28
                    } else {
                        if len(v24) > 0 {
                            v24[0] = (fi26 / ((fi23) | 1))
                        }
                    }
                }
                var v29 uint32 = (v1 / ((v1) | 1))
                _ = fi26
                _ = v27
                _ = v29
            }
            _ = fi23
            _ = v24
            _ = v25
        }
        fmt.Println(v1)
    }
    var v30 int32 = v4
    var v31 int16 = 1
    if (0 > 32767) {
        switch v3 {
        case "\n\t":
            for fi32 := 0; fi32 < 9; fi32++ {
                var v33 int8 = -128
                _ = fi32
                _ = v33
            }
        default:
            for fi34 := 0; fi34 < 14; fi34++ {
                for fi35 := 0; fi35 < 6; fi35++ {
                    switch v3 {
                    case "\n\t":
                        v36 := []int{32768, fi35}
                        _ = v36
                    case "hello":
                        v37 := []int{128, int(v30)}
                        v14 -= (v14 * float32((0.1 * float64(v14))))
                        if (v4 > (-(v30))) {
                            var v38 string = "a"
                            _ = v38
                        } else {
                            var v39 int16 = 32767
                            _ = v39
                        }
                        v40 := GenS3{F0: 2, F1: (v31 ^ (-(v31)))}
                        _ = v37
                        _ = v40
                    case "":
                        v41 := []int{127, fi34}
                        _ = v41
                    default:
                        fmt.Println(fi34, fi34)
                    }
                    v42 := GenS4{F0: 0}
                    if (v14 < (v14 - 0.0)) {
                        var v43 uint8 = 255
                        if (true && ((v1 | uint32(v43)) == 4294967295)) {
                            v3 = "hello"
                        } else {
                            var v44 bool = (false && true)
                            _ = v44
                        }
                        _ = v43
                    } else {
                        var v45 int32 = v30
                        var v46 string = ("\n\t" + v3)
                        _ = v45
                        _ = v46
                    }
                    _ = fi35
                    _ = v42
                }
                v47 := []int{(fi34 - fi34), fi34, fi34}
                _ = fi34
                _ = v47
            }
        }
    } else {
        v48 := GenS4{F0: v30}
        if true {
            if true {
                switch v3 {
                case "hello":
                    v30 -= 193
                case "foo bar":
                    if (!(0 > (v30 - (v4 << ((uint(v1)) & 63))))) {
                        if false {
                            v31 = v31
                            if true {
                                fmt.Println(v14, v1)
                                if false {
                                    fmt.Println(v4, v4)
                                } else {
                                    v49 := []int{42, 4294967295, 1, int(v1)}
                                    var v50 int32 = 718
                                    _ = v49
                                    _ = v50
                                }
                                v51 := []int{42, 32767}
                                _ = v51
                            }
                            var v52 float64 = (1.5 * (float64(v14) * (1.5 * 0.0)))
                            _ = v52
                        } else {
                            v2 += 127
                            v53 := []int{-799, 42}
                            v54 := GenS2{F0: 128, F1: (v1 % ((v1) | 1)), F2: v31}
                            _ = v53
                            _ = v54
                        }
                        v55 := []int{}
                        v56 := GenS5{F0: "", F1: 233}
                        _ = v55
                        _ = v56
                    }
                case "":
                    v57 := GenS2{F0: 195, F1: 1, F2: (v31 + v31)}
                    _ = v57
                default:
                    switch -1 {
                    case 0:
                        v3 = ("\n\t" + v3)
                    }
                }
                var v58 uint = 485
                fmt.Println(v58, v31, v3)
                _ = v58
            }
        }
        _ = v48
    }
    var v59 bool = false
    v60 := GenS2{F0: 1, F1: v1, F2: (-(v31))}
    var v61 uint32 = v1
    v3 = ""
    var v62 int8 = v2
    v63 := GenS2{F0: 131, F1: (v1 - 56), F2: 960}
    v61 = 42
    if (72 < 2) {
        v30 = (-(v4))
        var v64 int16 = v31
        v65 := []int{9223372036854775807, 42, 42, -9223372036854775808}
        _ = v64
        _ = v65
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v59:", v59)
    fmt.Println("v61:", v61)
    fmt.Println("v62:", v62)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v59
    _ = v60
    _ = v61
    _ = v62
    _ = v63
}
