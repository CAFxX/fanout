package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64; F2 int16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint; F1 uint32; F2 string }
type GenS4 struct { F0 uint64; F1 uint }
type GenS5 struct { F0 uint16; F1 bool; F2 string }
type GenS6 struct { F0 uint8; F1 uint32; F2 uint16 }
func genFn0(p0 uint, p1 uint8) int8 {
    var v5 uint32 = 0
    if true {
        switch ("foo bar" + "hello") {
        case "foo bar":
            switch "foo bar" {
            case "\n\t":
                if true {
                    for fi6 := 0; fi6 < 9; fi6++ {
                        if false {
                            v7 := GenS1{F0: 255, F1: 1, F2: -1}
                            for fi8 := 0; fi8 < 15; fi8++ {
                                fmt.Println(v5, p1, v5)
                                _ = fi8
                            }
                            v9 := GenS1{F0: (^(p0)), F1: 2, F2: -1}
                            _ = v7
                            _ = v9
                        } else {
                            for fi10 := 0; fi10 < 9; fi10++ {
                                var v11 float32 = -2.25
                                var v12 float32 = v11
                                v13 := GenS1{F0: p0, F1: 18446744073709551615, F2: 0}
                                _ = fi10
                                _ = v11
                                _ = v12
                                _ = v13
                            }
                            v5 += 4294967295
                        }
                        _ = fi6
                    }
                    var v14 int = 4294967295
                    switch "" {
                    case "":
                        if true {
                            var v15 int8 = -50
                            _ = v15
                        }
                    case "foo bar":
                        for fi16 := 0; fi16 < 9; fi16++ {
                            v17 := []int{(^(v14))}
                            _ = fi16
                            _ = v17
                        }
                    case "a":
                        var v18 int64 = 42
                        _ = v18
                    default:
                        v19 := []int{-932, (-(v14))}
                        _ = v19
                    }
                    _ = v14
                }
            case "a":
                if ((p1 + 0) == p1) {
                    v20 := []int{127, 553, -132}
                    _ = v20
                } else {
                    for fi21 := 0; fi21 < 16; fi21++ {
                        switch "foo bar" {
                        case "\n\t":
                            p0 -= (^(p0))
                        case "hello":
                            for fi22 := 0; fi22 < 7; fi22++ {
                                v23 := []int{32768, fi21, (fi22 | fi21), (-(fi21))}
                                v24 := GenS1{F0: p0, F1: 1, F2: -32768}
                                switch fi21 {
                                case 1:
                                    if false {
                                        var v25 int16 = 0
                                        _ = v25
                                    }
                                case -177:
                                    if len(v23) > 0 {
                                        v23[0] = (fi22 | (-(fi22)))
                                    }
                                }
                                if true {
                                    var v26 uint8 = (p1 >> ((18446744073709551615) & 63))
                                    p1 -= v26
                                    p1 -= (p1 % (((^(v26))) | 1))
                                    _ = v26
                                }
                                _ = fi22
                                _ = v23
                                _ = v24
                            }
                        case "a":
                            p0 = (p0 & (p0 * p0))
                        }
                        v5 -= v5
                        fmt.Println(p0, p1, fi21)
                        _ = fi21
                    }
                }
            case "foo bar":
                fmt.Println(p1)
            }
        case "a":
            p0 = uint(v5)
        }
        p0 = 950
        var v27 uint16 = 302
        _ = v27
    }
    _ = p0
    _ = p1
    _ = v5
    return 100
}

func main() {
    var v1 string = ""
    v2 := []int{-59}
    v3 := make(map[uint32]uint32)
    var v4 int8 = -123
    v28 := GenS2{F0: 1000}
    var v29 int = 0
    v30 := GenS3{F0: 936, F1: 42, F2: v1}
    if false {
        var v31 int32 = 1
        if (false && false) {
            var v32 uint8 = 128
            _ = v32
        }
        _ = v31
    }
    var v33 uint8 = 0
    if true {
        if false {
            for fi34 := 0; fi34 < 3; fi34++ {
                v4 = v4
                v29 = (fi34 + v29)
                _ = fi34
            }
            v29 += v29
            v35 := GenS1{F0: 1, F1: 46, F2: 1}
            _ = v35
        } else {
            v1 = v1
            var v36 int16 = 0
            _ = v36
        }
        v4 -= (v4 & v4)
        v37 := GenS4{F0: 18446744073709551615, F1: 900}
        _ = v37
    }
    v38 := GenS5{F0: 203, F1: (false || false), F2: "\n\t"}
    fmt.Println(v4)
    v4 += (v4 ^ 1)
    switch (v29 * (v29 / ((32767) | 1))) {
    case -666:
        switch "foo bar" {
        case "foo bar":
            v39 := GenS1{F0: 18446744073709551615, F1: 0, F2: 1}
            _ = v39
        case "hello":
            v4 += (v4 >> ((0) & 63))
        case "":
            v33 = 255
        }
    case -1:
        fmt.Println(v29)
    case 0:
        var v40 int16 = 673
        _ = v40
    }
    switch (v29 >> ((1) & 63)) {
    case 2147483647:
        _ = len(v2)
    }
    for fi41 := 0; fi41 < 16; fi41++ {
        for fi42 := 0; fi42 < 6; fi42++ {
            v43 := GenS2{F0: 499}
            if false {
                v4 += (^(v4))
                fmt.Println(v33)
                var v44 int = fi42
                v45 := v3[4294967295]
                v46 := GenS4{F0: 1, F1: 67}
                _ = v44
                _ = v45
                _ = v46
            } else {
                delete(v3, 2147483648)
                var v47 float64 = 0.1
                v48 := v3[0]
                _ = v47
                _ = v48
            }
            if true {
                var v49 uint32 = 460
                var v50 uint = 1
                switch v1 {
                case "":
                    var v51 uint = v50
                    _ = v51
                case "hello":
                    v49 = v49
                case "a":
                    v29 -= fi42
                }
                _ = v49
                _ = v50
            } else {
                var v52 int8 = (-(v4))
                if true {
                    var v53 int64 = 9223372036854775807
                    _ = v53
                }
                _ = v52
            }
            _ = fi42
            _ = v43
        }
        _ = fi41
    }
    v33 = 127
    fmt.Println(v4)
    v54 := GenS6{F0: 255, F1: 118, F2: 2}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v33
    _ = v38
    _ = v4
    _ = v54
}
