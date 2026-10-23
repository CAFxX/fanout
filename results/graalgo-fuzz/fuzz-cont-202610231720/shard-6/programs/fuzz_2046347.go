package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 bool; F1 string; F2 int64 }
type GenS4 struct { F0 int16; F1 uint; F2 int16 }
func genFn0(p0 uint16) uint16 {
    v7 := GenS1{F0: 56, F1: 42}
    var v8 int16 = -32768
    for fi9 := 0; fi9 < 12; fi9++ {
        v10 := []int{fi9, fi9, (fi9 & (^(fi9))), fi9}
        _ = fi9
        _ = v10
    }
    var v11 int64 = -9223372036854775808
    v11 -= (v11 ^ int64(v8))
    _ = p0
    _ = v11
    _ = v7
    _ = v8
    return p0
}
func genFn1(p0 bool, p1 bool) int16 {
    for fi12 := 0; fi12 < 14; fi12++ {
        for fi13 := 0; fi13 < 5; fi13++ {
            if p1 {
                v14 := []int{fi13, fi12, fi13, fi13}
                _ = v14
            } else {
                var v15 uint64 = 18446744073709551615
                var v16 int32 = 1
                v17 := []int{(fi12 * fi13), (^(fi12)), (^(fi13)), 258}
                _ = v15
                _ = v16
                _ = v17
            }
            v18 := []int{}
            _ = fi13
            _ = v18
        }
        v19 := []int{(^(fi12))}
        p1 = p1
        _ = fi12
        _ = v19
    }
    switch 42 {
    case 1:
        p0 = (("" + ("foo bar" + "")) > ("hello" + "hello"))
    default:
        if (!p0) {
            var v20 int16 = 42
            var v21 int16 = (v20 / ((1000) | 1))
            p1 = p1
            _ = v20
            _ = v21
        }
    }
    _ = p0
    _ = p1
    return 1
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[string]int)
    v3 := make(map[int]uint32)
    var v4 uint16 = 32767
    var v5 string = ("hello" + (("a" + "") + (("hello" + "hello") + "")))
    var v6 int32 = 0
    switch "a" {
    case "a":
        v5 = "\n\t"
    case "":
        for fi22 := 0; fi22 < 12; fi22++ {
            var v23 float64 = ((-2.25 / 1.0) * 0.0)
            _ = fi22
            _ = v23
        }
    case "foo bar":
        if (!false) {
            var v24 int16 = -32768
            _ = v24
        } else {
            v6 = 2147483647
            var v25 uint16 = 2
            v26 := GenS2{F0: -1}
            _ = v25
            _ = v26
        }
    default:
        for fi27 := 0; fi27 < 20; fi27++ {
            if false {
                v28 := GenS1{F0: 722, F1: int8(v4)}
                v6 -= 2147483647
                v29 := GenS1{F0: 18446744073709551615, F1: -128}
                _ = v28
                _ = v29
            } else {
                for fi30 := 0; fi30 < 18; fi30++ {
                    switch (fi30 ^ fi30) {
                    case -193:
                        v6 = 561
                    case 1:
                        v6 += (-(v6))
                    case -1:
                        v6 = int32(fi27)
                    default:
                        var v31 float64 = 0.1
                        _ = v31
                    }
                    if false {
                        v6 = -2147483648
                        v32 := GenS3{F0: (v5 != (v5 + v5)), F1: v5, F2: -1}
                        var v33 int16 = -1
                        _ = v32
                        _ = v33
                    }
                    if false {
                        v6 = (v6 ^ (v6 << ((2) & 63)))
                    }
                    _ = fi30
                }
                for fi34 := 0; fi34 < 14; fi34++ {
                    switch fi27 {
                    case 1:
                        fmt.Println(v4)
                        fmt.Println(v6, v6)
                    case 265:
                        var v35 float32 = (1e10 / 1.5)
                        _ = v35
                    default:
                        if (false || true) {
                            v4 -= v4
                        } else {
                            fmt.Println(fi27, fi34, fi34)
                        }
                    }
                    delete(v2, v5)
                    _ = fi34
                }
            }
            _ = fi27
        }
    }
    if (float64((0.0 * 0.0)) == ((float64(1e10) - 1.5) + -2.25)) {
        for fi36 := 0; fi36 < 4; fi36++ {
            var v37 int16 = -934
            _ = fi36
            _ = v37
        }
        switch "" {
        case "foo bar":
            var v38 uint32 = 65535
            _ = v38
        case "\n\t":
            delete(v2, "")
        case "a":
            if true {
                v4 = 857
                v6 = 693
            }
        }
        v6 -= (-(v6))
    }
    _ = len(v2)
    for fi39 := 0; fi39 < 6; fi39++ {
        fmt.Println(fi39, fi39)
        v6 += v6
        var v40 bool = ((v6 / ((v6) | 1)) == -1)
        _ = fi39
        _ = v40
    }
    for fi41 := 0; fi41 < 18; fi41++ {
        if false {
            switch v5 {
            case "\n\t":
                v42 := GenS3{F0: (-1 <= 100), F1: "foo bar", F2: 9223372036854775807}
                _ = v42
            case "hello":
                switch "hello" {
                case "\n\t":
                    for fi43 := 0; fi43 < 15; fi43++ {
                        v5 = "a"
                        switch "foo bar" {
                        case "hello":
                            switch -690 {
                            case 860:
                                if false {
                                    switch (fi43 + fi41) {
                                    case -806:
                                        v4 = v4
                                    case 2147483647:
                                        v44 := GenS3{F0: true, F1: (((v5 + v5) + (v5 + v5)) + "\n\t"), F2: -675}
                                        _ = v44
                                    case 32767:
                                        var v45 int32 = -1
                                        _ = v45
                                    default:
                                        switch -2147483648 {
                                        case -2147483648:
                                            if true {
                                                var v46 uint32 = 2
                                                v5 = "\n\t"
                                                _ = v46
                                            }
                                        default:
                                            var v47 uint32 = 2147483648
                                            _ = v47
                                        }
                                    }
                                    v48 := []int{int(v6), fi43, fi43}
                                    if false {
                                        v6 += 178
                                        if false {
                                            switch v5 {
                                            case "foo bar":
                                                v6 = -1
                                            }
                                            fmt.Println(v5, fi41)
                                        } else {
                                            v5 = (v5 + v5)
                                            v49 := v3[fi41]
                                            _ = v49
                                        }
                                    }
                                    v50 := GenS1{F0: 867, F1: 100}
                                    v51 := GenS2{F0: 42}
                                    if (v6 < v6) {
                                        v6 = v6
                                        v52 := GenS2{F0: v6}
                                        var v53 uint8 = 18
                                        fmt.Println(fi43)
                                        var v54 uint64 = 585
                                        _ = len(v48)
                                        fmt.Println(v53)
                                        _ = v52
                                        _ = v53
                                        _ = v54
                                    } else {
                                        if len(v48) >= 2 {
                                            v48 = v48[1:len(v48)]
                                        }
                                        v4 = v4
                                        var v55 float64 = (0.0 - (float64(0.0) + -2.25))
                                        _ = v55
                                    }
                                    _ = v48
                                    _ = v50
                                    _ = v51
                                } else {
                                    v56 := GenS4{F0: -32768, F1: 2, F2: -32768}
                                    _ = v56
                                }
                            case 9223372036854775807:
                                if (42 != 65535) {
                                    var v57 bool = ((^(fi41)) < fi43)
                                    _ = v57
                                }
                            }
                        default:
                            v58 := GenS4{F0: 0, F1: 4294967295, F2: 1}
                            _ = v58
                        }
                        _ = fi43
                    }
                case "hello":
                    for fi59 := 0; fi59 < 20; fi59++ {
                        v3[fi59] = 71
                        _ = fi59
                    }
                }
            default:
                v5 = "\n\t"
            }
            if false {
                var v60 uint32 = 65535
                for fi61 := 0; fi61 < 17; fi61++ {
                    v62 := v3[(-(fi61))]
                    _ = fi61
                    _ = v62
                }
                _ = v60
            }
        }
        if false {
            fmt.Println(fi41, v5, v4)
        }
        v6 = 0
        _ = fi41
    }
    v63 := []int{32768}
    var v64 uint16 = (v4 | 2)
    var v65 uint64 = 745
    fmt.Println(v5, v5, v5)
    v64 -= v4
    v64 = 0
    v64 -= (^(v4))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v63:", len(v63), v63)
    fmt.Println("v64:", v64)
    fmt.Println("v65:", v65)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v63
    _ = v64
    _ = v65
}
