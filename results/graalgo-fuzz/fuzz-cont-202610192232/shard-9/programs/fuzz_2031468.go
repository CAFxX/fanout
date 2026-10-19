package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64 }
type GenS2 struct { F0 uint32; F1 int }
type GenS3 struct { F0 uint16; F1 uint64; F2 int32 }
type GenS4 struct { F0 uint8; F1 int64; F2 uint }
type GenS5 struct { F0 int16; F1 int8 }
type GenS6 struct { F0 uint8; F1 int32; F2 int8 }
type GenS7 struct { F0 int8 }
func genFn0() uint32 {
    var v4 uint8 = 2
    switch "\n\t" {
    case "\n\t":
        var v5 float32 = 1.5
        _ = v5
    case "":
        if true {
            var v6 uint64 = 759
            var v7 uint32 = 4294967295
            _ = v6
            _ = v7
        }
    }
    fmt.Println(v4)
    _ = v4
    return 2
}
func genFn1(p0 int) uint64 {
    v8 := GenS1{F0: (-(p0)), F1: 65}
    for fi9 := 0; fi9 < 20; fi9++ {
        var v10 uint = 538
        _ = fi9
        _ = v10
    }
    v11 := []int{-2147483648, (p0 << ((18446744073709551615) & 63)), 9223372036854775807, (p0 * p0)}
    for fi12 := 0; fi12 < 13; fi12++ {
        switch (^(p0)) {
        case 42:
            if len(v11) > 0 {
                v11[0] = (-(fi12))
            }
        case 1:
            p0 -= fi12
        case 0:
            if len(v11) >= 2 {
                v11 = v11[1:len(v11)]
            }
        }
        v13 := GenS2{F0: 2, F1: p0}
        _ = fi12
        _ = v13
    }
    v14 := GenS2{F0: 848, F1: (p0 * 116)}
    _ = p0
    _ = v11
    _ = v14
    _ = v8
    return 42
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[int]string)
    var v3 uint16 = 42
    v15 := GenS1{F0: 2147483647, F1: 462}
    for fi16 := 0; fi16 < 5; fi16++ {
        switch (fi16 % (((fi16 ^ 0)) | 1)) {
        case -2147483648:
            var v17 int = (fi16 / (((-(fi16))) | 1))
            _ = v17
        case -9223372036854775808:
            for fi18 := 0; fi18 < 5; fi18++ {
                if true {
                    v3 = (v3 << ((65535) & 63))
                    _ = len(v2)
                    if false {
                        fmt.Println(fi16, v3)
                        v3 -= (v3 * v3)
                        var v19 string = "\n\t"
                        _ = v19
                    } else {
                        v20 := GenS3{F0: v3, F1: genFn1(fi18), F2: 1}
                        v3 = 42
                        _ = v20
                    }
                } else {
                    if (1 != 100) {
                        _ = len(v1)
                        v3 = v3
                    } else {
                        switch "foo bar" {
                        case "\n\t":
                            if true {
                                var v21 float32 = 1.5
                                v22 := []int{fi16}
                                fmt.Println(fi18, fi18, fi16)
                                _ = v21
                                _ = v22
                            } else {
                                v3 -= v3
                            }
                        case "a":
                            var v23 bool = (true || (!true))
                            _ = v23
                        }
                        if (v3 >= (v3 & v3)) {
                            if true {
                                if (217 == 0) {
                                    if ((float32((1e10 / 1e10)) + (((-2.25 * 1e10) - 1.5) + -2.25)) < float32(0.1)) {
                                        var v24 uint = 1
                                        var v25 uint32 = 42
                                        var v26 int8 = int8(fi16)
                                        switch -9223372036854775808 {
                                        case 127:
                                            v3 = v3
                                        default:
                                            v3 -= (v3 / (((v3 % ((403) | 1))) | 1))
                                        }
                                        fmt.Println(v24)
                                        _ = v24
                                        _ = v25
                                        _ = v26
                                    } else {
                                        var v27 int32 = 42
                                        _ = v27
                                    }
                                }
                            } else {
                                v28 := GenS4{F0: 2, F1: -1, F2: 4294967295}
                                _ = v28
                            }
                        }
                    }
                    v3 += (v3 % (((v3 * v3)) | 1))
                    fmt.Println(fi18, fi18, v3)
                    switch -808 {
                    case 0:
                        v29 := []int{}
                        _ = v29
                    case 42:
                        if (0 == 42) {
                            var v30 uint16 = v3
                            v31 := []int{42, (fi18 - (fi18 + fi18)), (^(fi18)), fi16}
                            _ = v30
                            _ = v31
                        }
                    }
                    var v32 uint = 42
                    _ = v32
                }
                var v33 uint8 = 0
                var v34 float32 = ((3.14159 + -2.25) + 3.14159)
                _ = fi18
                _ = v33
                _ = v34
            }
        default:
            switch ("a" + (("a" + (("foo bar" + (("\n\t" + ("a" + "")) + "foo bar")) + "hello")) + "foo bar")) {
            case "\n\t":
                var v35 uint16 = v3
                _ = v35
            default:
                v36 := GenS3{F0: (v3 ^ 0), F1: 985, F2: 42}
                _ = v36
            }
        }
        var v37 uint8 = 0
        _ = fi16
        _ = v37
    }
    if true {
        var v38 int8 = 127
        v3 = (v3 % ((v3) | 1))
        v3 += (^(v3))
        _ = v38
    }
    for fi39 := 0; fi39 < 15; fi39++ {
        v40 := GenS5{F0: 32767, F1: -71}
        v41 := GenS4{F0: 1, F1: 9223372036854775807, F2: 18446744073709551615}
        v3 -= (v3 & 42)
        _ = fi39
        _ = v40
        _ = v41
    }
    for fi42 := 0; fi42 < 16; fi42++ {
        v43 := GenS2{F0: 4294967295, F1: fi42}
        _ = fi42
        _ = v43
    }
    v3 = v3
    v44 := GenS6{F0: 237, F1: 489, F2: -2}
    for fi45 := 0; fi45 < 17; fi45++ {
        v3 += v3
        for fi46 := 0; fi46 < 19; fi46++ {
            var v47 int8 = 0
            if (-128 >= v47) {
                var v48 int16 = 32767
                if (!true) {
                    if true {
                        switch 9223372036854775807 {
                        case 334:
                            v49 := GenS6{F0: 1, F1: 0, F2: v47}
                            _ = v49
                        case 2147483647:
                            switch "" {
                            case "a":
                                var v50 uint = 642
                                _ = v50
                            case "\n\t":
                                if (114 < 1) {
                                    v47 = (v47 ^ -128)
                                }
                            case "":
                                fmt.Println(v47, v47, v47)
                                switch "hello" {
                                case "\n\t":
                                    var v51 uint8 = 1
                                    fmt.Println(v47, v3, v48)
                                    var v52 uint8 = (v51 << ((833) & 63))
                                    _ = len(v2)
                                    _ = v51
                                    _ = v52
                                case "hello":
                                    fmt.Println(fi45, v48, fi46)
                                }
                                var v53 int16 = v48
                                v3 -= v3
                                _ = v53
                            }
                        default:
                            fmt.Println(fi46, fi45, fi46)
                        }
                        v47 -= v47
                        v47 = (v47 % (((v47 | v47)) | 1))
                    }
                    var v54 int64 = -9223372036854775808
                    _ = v54
                }
                _ = v48
            } else {
                v55 := GenS7{F0: (v47 ^ v47)}
                _ = len(v2)
                _ = v55
            }
            v3 -= v3
            _ = fi46
            _ = v47
        }
        v2[(fi45 & -2147483648)] = "foo bar"
        _ = fi45
    }
    if (!(2147483648 > 2)) {
        v56 := v1[genFn0()]
        _ = v56
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v44
}
