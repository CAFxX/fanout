package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string; F2 int64 }
type GenS2 struct { F0 int; F1 int64; F2 uint32 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 string; F1 int32 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 uint32 }
func genFn0(p0 uint64) int8 {
    if false {
        v6 := GenS1{F0: 1, F1: "foo bar", F2: -24}
        switch ((("a" + "foo bar") + ("foo bar" + ("\n\t" + ""))) + "\n\t") {
        case "a":
            v7 := []int{}
            _ = v7
        case "foo bar":
            var v8 uint = 4294967295
            _ = v8
        case "\n\t":
            v9 := []int{32767, -410}
            _ = v9
        default:
            if ("" < ("" + (("hello" + (("" + "\n\t") + "foo bar")) + ("a" + "\n\t")))) {
                v10 := []int{127, 128, 533, 0}
                var v11 uint = 4294967295
                _ = v10
                _ = v11
            }
        }
        _ = v6
    }
    switch 32767 {
    case 128:
        var v12 uint = 1
        _ = v12
    case 9223372036854775807:
        p0 = (p0 - p0)
    }
    if true {
        v13 := []int{9223372036854775807, 42}
        var v14 bool = true
        _ = v13
        _ = v14
    } else {
        p0 = p0
    }
    switch ("" + "foo bar") {
    case "":
        v15 := GenS1{F0: 100, F1: "\n\t", F2: 0}
        _ = v15
    case "hello":
        v16 := GenS2{F0: 127, F1: 632, F2: 2}
        _ = v16
    case "foo bar":
        if (!false) {
            p0 += 4294967295
            p0 = p0
        } else {
            var v17 uint16 = 993
            for fi18 := 0; fi18 < 20; fi18++ {
                p0 = 0
                if (true && true) {
                    var v19 bool = true
                    var v20 uint32 = 0
                    _ = v19
                    _ = v20
                } else {
                    var v21 int32 = 642
                    var v22 uint32 = 587
                    _ = v21
                    _ = v22
                }
                if false {
                    if true {
                        p0 += p0
                        var v23 uint16 = 32767
                        var v24 int16 = 1
                        _ = v23
                        _ = v24
                    } else {
                        v25 := GenS3{F0: 255}
                        v26 := GenS4{F0: 42}
                        _ = v25
                        _ = v26
                    }
                } else {
                    v27 := GenS5{F0: v17}
                    var v28 uint64 = (p0 << ((255) & 63))
                    if false {
                        p0 -= (v28 + (p0 / ((42) | 1)))
                        v17 = v17
                        v17 += v17
                    } else {
                        v17 += (v17 + 345)
                        var v29 uint = 255
                        v30 := GenS6{F0: (("a" + "") + "\n\t"), F1: 0}
                        _ = v29
                        _ = v30
                    }
                    _ = v27
                    _ = v28
                }
                _ = fi18
            }
            _ = v17
        }
    }
    switch "" {
    case "foo bar":
        v31 := GenS6{F0: (("\n\t" + ("\n\t" + "hello")) + "foo bar"), F1: -1}
        _ = v31
    case "hello":
        v32 := []int{9223372036854775807, -117, 880}
        _ = v32
    default:
        p0 = 1
    }
    _ = p0
    return -128
}
func genFn1(p0 string) int64 {
    if false {
        for fi33 := 0; fi33 < 13; fi33++ {
            p0 = p0
            _ = fi33
        }
        switch 4 {
        case -910:
            for fi34 := 0; fi34 < 6; fi34++ {
                var v35 int = fi34
                for fi36 := 0; fi36 < 1; fi36++ {
                    var v37 int8 = 11
                    v38 := GenS6{F0: "", F1: -929}
                    v37 = 127
                    v35 = 42
                    _ = fi36
                    _ = v37
                    _ = v38
                }
                if (-128 < -35) {
                    fmt.Println(v35)
                    v39 := GenS3{F0: 169}
                    _ = v39
                } else {
                    v40 := []int{v35, v35, fi34, 1}
                    var v41 int64 = 647
                    p0 = (p0 + ((p0 + "\n\t") + (p0 + ((p0 + (((p0 + p0) + ("" + (p0 + (p0 + p0)))) + ((p0 + p0) + p0))) + "hello"))))
                    _ = v40
                    _ = v41
                }
                _ = fi34
                _ = v35
            }
        case 0:
            v42 := []int{470, 597}
            _ = v42
        case 4294967295:
            var v43 uint32 = 0
            _ = v43
        }
    } else {
        fmt.Println(p0)
        v44 := GenS7{F0: false}
        if (true && ((0.1 * 0.1) >= 0.1)) {
            var v45 int32 = 0
            if false {
                p0 = p0
                v45 = (v45 + (v45 - -1))
            } else {
                v45 -= 0
                fmt.Println(p0)
                fmt.Println(p0, p0)
            }
            var v46 uint64 = 0
            _ = v45
            _ = v46
        } else {
            p0 = "a"
            var v47 float64 = float64(0.1)
            p0 = (p0 + ("a" + p0))
            _ = v47
        }
        _ = v44
    }
    for fi48 := 0; fi48 < 6; fi48++ {
        v49 := []int{(fi48 >> ((499) & 63))}
        p0 = (p0 + ("foo bar" + (p0 + p0)))
        _ = fi48
        _ = v49
    }
    var v50 int32 = 2147483647
    v50 = v50
    _ = p0
    _ = v50
    return -9223372036854775808
}

func main() {
    v1 := []string{}
    var v2 bool = false
    v3 := make(map[uint32]bool)
    var v4 float64 = 3.14159
    var v5 uint16 = 255
    v1 = append(v1, "foo bar")
    var v51 int16 = 32767
    fmt.Println(v5, v4, v4)
    var v52 uint = 65535
    v53 := GenS6{F0: ("" + "a"), F1: -659}
    v3[42] = (v2 && true)
    _ = len(v3)
    v3[1] = v2
    var v54 string = ("\n\t" + "foo bar")
    v1 = append(v1, v54)
    _ = len(v3)
    switch v54 {
    case "\n\t":
        v55 := GenS1{F0: genFn0(275), F1: v54, F2: 848}
        _ = v55
    case "foo bar":
        v4 = (((3.14159 - ((0.1 - float64(((1e10 * (0.1 + -2.25)) - 1.5))) - (float64(((1.5 - 1e10) - float32(0.1))) - v4))) - 0.0) + (3.14159 / 0.1))
    case "a":
        v54 = v54
    default:
        switch v54 {
        case "a":
            if ((!(v4 <= (v4 / v4))) || v2) {
                var v56 uint64 = 429
                _ = v56
            } else {
                var v57 uint8 = 2
                if v2 {
                    var v58 int = 127
                    _ = v58
                } else {
                    switch 32768 {
                    case -48:
                        v59 := GenS8{F0: 410}
                        _ = v59
                    case 4294967295:
                        var v60 int32 = -740
                        _ = v60
                    }
                }
                _ = v57
            }
        case "hello":
            for fi61 := 0; fi61 < 6; fi61++ {
                if ((true && (!true)) && v2) {
                    v62 := GenS6{F0: (v54 + "foo bar"), F1: 2147483647}
                    v4 += (v4 / 1.0)
                    var v63 int = 127
                    _ = v62
                    _ = v63
                }
                _ = fi61
            }
        default:
            var v64 float32 = 0.0
            _ = v64
        }
    }
    for fi65 := 0; fi65 < 12; fi65++ {
        var v66 int = int(v52)
        v5 = v5
        fmt.Println(v2, v4, v2)
        _ = fi65
        _ = v66
    }
    v52 = (v52 / ((1) | 1))
    delete(v3, 789)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    fmt.Println("v54:", v54)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v51
    _ = v52
    _ = v53
    _ = v54
}
