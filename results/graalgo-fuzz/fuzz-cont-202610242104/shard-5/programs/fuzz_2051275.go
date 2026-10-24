package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int16; F1 int }
type GenS3 struct { F0 int64; F1 string }
type GenS4 struct { F0 string; F1 int64; F2 string }
type GenS5 struct { F0 uint64; F1 uint64 }
type GenS6 struct { F0 string; F1 uint8; F2 bool }
type GenS7 struct { F0 int8; F1 uint8 }
type GenS8 struct { F0 int64; F1 uint8; F2 uint16 }
type GenS9 struct { F0 bool; F1 uint64 }
type GenS10 struct { F0 uint; F1 uint; F2 int8 }
type GenS11 struct { F0 string }
type GenS12 struct { F0 int16 }
func genFn0(p0 int, p1 uint16) int16 {
    v7 := GenS1{F0: 65535}
    v8 := GenS2{F0: 0, F1: (p0 >> ((65535) & 63))}
    for fi9 := 0; fi9 < 2; fi9++ {
        var v10 int8 = 127
        _ = fi9
        _ = v10
    }
    switch (p0 ^ 1) {
    case -103:
        v11 := GenS2{F0: -32768, F1: p0}
        _ = v11
    case 1:
        p1 += (p1 | (p1 % (((p1 | 1)) | 1)))
    default:
        var v12 int64 = -1
        _ = v12
    }
    for fi13 := 0; fi13 < 1; fi13++ {
        p1 -= p1
        v14 := []int{p0, (p0 - -2147483648)}
        p0 -= p0
        _ = fi13
        _ = v14
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return -318
}
func genFn1(p0 uint8) int {
    switch 365 {
    case -1:
        for fi15 := 0; fi15 < 20; fi15++ {
            fmt.Println(p0, fi15)
            _ = fi15
        }
    case -2147483648:
        for fi16 := 0; fi16 < 10; fi16++ {
            v17 := []int{fi16, (fi16 | (fi16 & (fi16 ^ (fi16 - 790))))}
            p0 = p0
            switch fi16 {
            case 127:
                p0 -= 101
            case 401:
                p0 -= p0
            case -1:
                p0 = p0
            default:
                p0 -= 179
            }
            _ = fi16
            _ = v17
        }
    }
    v18 := GenS1{F0: 769}
    _ = p0
    _ = v18
    return 771
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[int]uint32)
    var v3 int8 = 100
    var v4 int8 = v3
    var v5 float64 = float64((-2.25 - 1.5))
    var v6 int8 = v4
    v19 := GenS3{F0: -1, F1: "foo bar"}
    switch 42 {
    case 32767:
        v20 := GenS3{F0: 42, F1: ""}
        _ = v20
    case -238:
        v6 += v4
    case 2147483647:
        for fi21 := 0; fi21 < 6; fi21++ {
            for fi22 := 0; fi22 < 3; fi22++ {
                v6 -= v6
                fmt.Println(v4)
                var v23 float32 = 0.1
                v3 = v4
                v3 = -128
                _ = fi22
                _ = v23
            }
            for fi24 := 0; fi24 < 10; fi24++ {
                v25 := GenS3{F0: 1, F1: ("\n\t" + "foo bar")}
                v26 := GenS2{F0: 42, F1: -458}
                _ = fi24
                _ = v25
                _ = v26
            }
            var v27 int8 = (v6 << ((2) & 63))
            _ = fi21
            _ = v27
        }
    default:
        var v28 bool = true
        _ = v28
    }
    v29 := GenS4{F0: "a", F1: 792, F2: "\n\t"}
    v30 := []int{2147483647, 584}
    switch "" {
    case "foo bar":
        delete(v1, 1)
    case "a":
        v1[449] = 570
    }
    v4 = (v4 % ((-128) | 1))
    switch "\n\t" {
    case "\n\t":
        for fi31 := 0; fi31 < 4; fi31++ {
            _ = len(v30)
            if true {
                switch ("" + ("foo bar" + (("\n\t" + ("foo bar" + "a")) + ("\n\t" + "\n\t")))) {
                case "hello":
                    v32 := GenS5{F0: 484, F1: 18446744073709551615}
                    _ = v32
                case "foo bar":
                    v5 = float64(-2.25)
                default:
                    switch (fi31 & -1) {
                    case 4294967295:
                        v33 := GenS6{F0: "a", F1: 127, F2: false}
                        _ = v33
                    case 32767:
                        v5 -= v5
                    case 42:
                        v3 -= (-(v6))
                    }
                }
                var v34 int16 = -312
                _ = v34
            }
            _ = fi31
        }
    default:
        _ = len(v1)
    }
    switch -1 {
    case -500:
        switch "\n\t" {
        case "\n\t":
            if (false || ((1000 == -32768) || ((203 == 127) || (((true || false) && false) || false)))) {
                v35 := GenS7{F0: -123, F1: 1}
                _ = v35
            } else {
                v36 := GenS8{F0: -1, F1: 0, F2: 42}
                switch -9223372036854775808 {
                case 32768:
                    v37 := GenS3{F0: 0, F1: ""}
                    _ = v37
                case 2147483647:
                    fmt.Println(v4, v4, v5)
                default:
                    switch genFn1(2) {
                    case 9223372036854775807:
                        v38 := GenS9{F0: (128 > 127), F1: 4294967295}
                        _ = v38
                    case 32767:
                        for fi39 := 0; fi39 < 12; fi39++ {
                            var v40 int8 = v3
                            switch "\n\t" {
                            case "foo bar":
                                fmt.Println(v5)
                            case "hello":
                                v6 += (^(v3))
                            default:
                                if (("" + "\n\t") <= ("foo bar" + ("\n\t" + "hello"))) {
                                    _ = len(v30)
                                    v41 := GenS10{F0: 255, F1: 42, F2: v3}
                                    _ = v41
                                }
                            }
                            _ = fi39
                            _ = v40
                        }
                    case -2147483648:
                        fmt.Println(v5, v3, v5)
                    default:
                        switch 4294967295 {
                        case -90:
                            switch -973 {
                            case 32767:
                                for fi42 := 0; fi42 < 20; fi42++ {
                                    v5 += float64(0.0)
                                    _ = fi42
                                }
                            case 127:
                                if true {
                                    v3 = v6
                                    _ = len(v30)
                                    fmt.Println(v4, v3)
                                } else {
                                    var v43 uint64 = 4294967295
                                    fmt.Println(v5, v5, v3)
                                    _ = v43
                                }
                            case -868:
                                if ((127 <= (v3 / ((42) | 1))) && true) {
                                    delete(v2, 0)
                                } else {
                                    v5 = (v5 * v5)
                                    var v44 int32 = 0
                                    v45 := GenS11{F0: "foo bar"}
                                    _ = v44
                                    _ = v45
                                }
                            default:
                                for fi46 := 0; fi46 < 13; fi46++ {
                                    if false {
                                        v47 := GenS8{F0: 42, F1: 244, F2: 42}
                                        var v48 float64 = (((0.0 / 3.14159) * float64(1e10)) / v5)
                                        _ = v47
                                        _ = v48
                                    } else {
                                        if true {
                                            fmt.Println(v6, v5)
                                        } else {
                                            v3 = 42
                                            delete(v1, 809)
                                            switch fi46 {
                                            case 615:
                                                var v49 int8 = -45
                                                _ = v49
                                            default:
                                                if (fi46 >= (fi46 ^ (fi46 & fi46))) {
                                                    v5 += 3.14159
                                                    for fi50 := 0; fi50 < 9; fi50++ {
                                                        var v51 int16 = 42
                                                        _ = fi50
                                                        _ = v51
                                                    }
                                                    var v52 int64 = 9223372036854775807
                                                    _ = v52
                                                }
                                            }
                                        }
                                        if ((-1 >= 859) && (2 < 500)) {
                                            for fi53 := 0; fi53 < 20; fi53++ {
                                                var v54 float64 = 1.5
                                                v55 := v1[922]
                                                _ = fi53
                                                _ = v54
                                                _ = v55
                                            }
                                            v56 := GenS4{F0: "\n\t", F1: 0, F2: ""}
                                            _ = v56
                                        }
                                    }
                                    v30 = append(v30, (-(fi46)))
                                    _ = fi46
                                }
                            }
                        case 757:
                            var v57 uint = 118
                            _ = v57
                        default:
                            var v58 int32 = 42
                            _ = v58
                        }
                    }
                }
                _ = v36
            }
        default:
            var v59 uint8 = 128
            _ = v59
        }
    case 1:
        var v60 uint64 = 0
        _ = v60
    case 127:
        var v61 float32 = 1e10
        _ = v61
    default:
        var v62 uint32 = 1
        _ = v62
    }
    if len(v30) >= 2 {
        v30 = v30[1:len(v30)]
    }
    v63 := GenS12{F0: 0}
    v6 = (v3 >> ((551) & 63))
    var v64 float32 = -2.25
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v64:", v64)
    _ = v1
    _ = v19
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
    _ = v63
    _ = v64
}
