package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int }
type GenS2 struct { F0 uint16; F1 string; F2 int32 }
type GenS3 struct { F0 uint64; F1 string }
type GenS4 struct { F0 int16; F1 int }
type GenS5 struct { F0 int16; F1 bool }
type GenS6 struct { F0 uint8; F1 uint16 }
type GenS7 struct { F0 int64; F1 int16 }
func genFn0(p0 uint) int {
    if (!(true || (!false))) {
        v4 := []int{-689, 0}
        var v5 uint8 = 1
        var v6 int = -2147483648
        _ = v4
        _ = v5
        _ = v6
    } else {
        p0 = (p0 | p0)
        var v7 int16 = 0
        var v8 uint8 = 128
        _ = v7
        _ = v8
    }
    for fi9 := 0; fi9 < 17; fi9++ {
        var v10 int16 = 1000
        _ = fi9
        _ = v10
    }
    p0 = p0
    for fi11 := 0; fi11 < 14; fi11++ {
        for fi12 := 0; fi12 < 13; fi12++ {
            var v13 bool = (!false)
            _ = fi12
            _ = v13
        }
        _ = fi11
    }
    _ = p0
    return -1
}
func genFn1(p0 uint8, p1 uint32) uint32 {
    v14 := GenS1{F0: (0 == 0), F1: 232}
    if (!(p1 >= (p1 | (^(p1))))) {
        var v15 int8 = -1
        var v16 uint32 = p1
        v17 := []int{4294967295}
        _ = v15
        _ = v16
        _ = v17
    }
    for fi18 := 0; fi18 < 8; fi18++ {
        v19 := []int{genFn0(209), genFn0(18446744073709551615), fi18, (fi18 % ((fi18) | 1))}
        switch "a" {
        case "a":
            for fi20 := 0; fi20 < 10; fi20++ {
                var v21 int64 = -1
                switch fi20 {
                case -9223372036854775808:
                    var v22 uint8 = p0
                    _ = v22
                case 2147483647:
                    if false {
                        var v23 uint = 69
                        _ = v23
                    }
                case 1:
                    var v24 uint32 = p1
                    if ((v21 >> ((863) & 63)) == v21) {
                        var v25 bool = true
                        var v26 bool = (1.5 >= float32((-2.25 + -2.25)))
                        var v27 bool = (!v26)
                        _ = v25
                        _ = v26
                        _ = v27
                    } else {
                        fmt.Println(p0, fi18)
                    }
                    var v28 int16 = -32768
                    _ = v24
                    _ = v28
                }
                v29 := GenS1{F0: false, F1: 1}
                _ = fi20
                _ = v21
                _ = v29
            }
        case "\n\t":
            p0 = p0
        case "hello":
            v30 := GenS2{F0: 42, F1: ("foo bar" + ("foo bar" + "hello")), F2: 2147483647}
            _ = v30
        default:
            fmt.Println(fi18, p1)
        }
        var v31 int = (fi18 >> ((18446744073709551615) & 63))
        _ = fi18
        _ = v19
        _ = v31
    }
    _ = p0
    _ = p1
    _ = v14
    return p1
}

func main() {
    v1 := []uint32{65535}
    v2 := make(map[uint32]bool)
    v3 := []string{("\n\t" + ("\n\t" + "\n\t")), "a", "a", "a"}
    if true {
        var v32 uint16 = 42
        _ = v32
    }
    v33 := GenS2{F0: 2, F1: "foo bar", F2: 0}
    var v34 bool = true
    v34 = (v34 || false)
    var v35 uint16 = 32767
    if (333 < (^(v35))) {
        v36 := GenS3{F0: 18446744073709551615, F1: "a"}
        var v37 int8 = -128
        _ = v36
        _ = v37
    } else {
        for fi38 := 0; fi38 < 10; fi38++ {
            if false {
                var v39 int32 = -2147483648
                _ = v39
            }
            v2[genFn1(81, genFn1(127, 2))] = (-272 == 1)
            for fi40 := 0; fi40 < 9; fi40++ {
                if v34 {
                    v35 = (v35 ^ v35)
                    if ((4294967295 == 18446744073709551615) || true) {
                        v34 = true
                        var v41 float64 = 0.0
                        _ = v41
                    }
                    _ = len(v3)
                }
                var v42 string = "a"
                switch -9223372036854775808 {
                case 127:
                    if (!((1 > -176) || (1 != 2147483647))) {
                        v43 := GenS4{F0: 32767, F1: 128}
                        if v34 {
                            if false {
                                v42 = v42
                                fmt.Println(v34)
                                v34 = true
                            } else {
                                var v44 int32 = 2147483647
                                var v45 int64 = 0
                                v44 -= (v44 ^ v44)
                                _ = v44
                                _ = v45
                            }
                            v34 = v34
                            v42 = ""
                        }
                        _ = v43
                    } else {
                        _ = len(v1)
                    }
                }
                v46 := GenS5{F0: 783, F1: (42 < 100)}
                v47 := v2[178]
                var v48 float64 = 0.1
                _ = fi40
                _ = v42
                _ = v46
                _ = v47
                _ = v48
            }
            _ = fi38
        }
        v49 := GenS6{F0: 42, F1: v35}
        var v50 int32 = 1
        _ = v49
        _ = v50
    }
    v51 := GenS7{F0: 42, F1: 42}
    var v52 int64 = -9223372036854775808
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v51
    _ = v52
}
