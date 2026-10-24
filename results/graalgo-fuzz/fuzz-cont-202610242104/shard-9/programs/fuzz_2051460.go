package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32; F2 bool }
type GenS2 struct { F0 bool; F1 int16; F2 uint8 }
type GenS3 struct { F0 bool; F1 int8 }
type GenS4 struct { F0 int8; F1 uint8 }
type GenS5 struct { F0 uint16 }
func genFn0() int8 {
    switch "foo bar" {
    case "hello":
        var v7 int64 = 0
        _ = v7
    case "\n\t":
        var v8 float64 = (-2.25 / 1.0)
        _ = v8
    default:
        v9 := GenS1{F0: 18446744073709551615, F1: 812, F2: true}
        _ = v9
    }
    var v10 int = 0
    _ = v10
    return 79
}
func genFn1(p0 int8, p1 int) int {
    fmt.Println(p0, p1)
    v11 := []int{(-(p1)), p1, (p1 - -9223372036854775808), (-(p1))}
    var v12 uint = 737
    fmt.Println(v12)
    p1 -= p1
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    return p1
}

func main() {
    var v1 uint = 254
    v2 := make(map[string]bool)
    var v3 int32 = -1
    v4 := make(map[int]int)
    v5 := make(map[int]bool)
    var v6 uint = (v1 - 255)
    v13 := GenS1{F0: 2, F1: 2147483648, F2: (!((v6 % (((v6 - (v6 | v6))) | 1)) < v6))}
    v1 += (v6 >> ((v6) & 63))
    var v14 int64 = -1
    var v15 bool = ((float32(3.14159) * 3.14159) < 1e10)
    for fi16 := 0; fi16 < 15; fi16++ {
        var v17 int = (^(fi16))
        _ = fi16
        _ = v17
    }
    switch 0 {
    case 1:
        var v18 uint8 = 118
        _ = v18
    case 32768:
        v19 := GenS1{F0: 18446744073709551615, F1: 0, F2: false}
        _ = v19
    default:
        if ((^(v14)) <= (v14 - v14)) {
            v20 := GenS1{F0: 617, F1: 65535, F2: (!false)}
            _ = v20
        } else {
            var v21 int16 = 0
            v22 := GenS1{F0: 722, F1: 65535, F2: true}
            switch genFn1(0, 127) {
            case 0:
                for fi23 := 0; fi23 < 17; fi23++ {
                    v24 := []int{}
                    var v25 float64 = float64(-2.25)
                    v3 = (v3 + v3)
                    _ = fi23
                    _ = v24
                    _ = v25
                }
            case 976:
                switch -45 {
                case 0:
                    v15 = (1.5 <= ((-2.25 / 1.0) * (float32(0.0) * ((float32(1.5) + 1.5) * (0.1 - 1.5)))))
                case 127:
                    var v26 uint64 = 148
                    _ = v26
                case 961:
                    v27 := GenS2{F0: v15, F1: 42, F2: 3}
                    _ = v27
                }
            }
            _ = v21
            _ = v22
        }
    }
    for fi28 := 0; fi28 < 20; fi28++ {
        var v29 string = (((("a" + "") + "foo bar") + "foo bar") + ("a" + "a"))
        v14 = -1
        if v15 {
            v29 = "\n\t"
            v29 = (v29 + ("foo bar" + "a"))
            if v15 {
                var v30 int16 = -32768
                fmt.Println(v15, v15, v3)
                var v31 float32 = -2.25
                _ = v30
                _ = v31
            }
        } else {
            var v32 float32 = ((-2.25 - ((((3.14159 * float32(float64((0.1 / 0.1)))) / 1.5) - 1.5) / (0.1 - 3.14159))) - 0.0)
            switch (fi28 * fi28) {
            case 0:
                v33 := GenS1{F0: 0, F1: 372, F2: v15}
                _ = v33
            case -203:
                fmt.Println(v32, v6, v14)
            case 32767:
                v29 = v29
            }
            _ = v32
        }
        _ = fi28
        _ = v29
    }
    for fi34 := 0; fi34 < 15; fi34++ {
        if ((!v15) || false) {
            _ = len(v5)
        }
        fmt.Println(v3)
        v35 := GenS2{F0: v15, F1: 42, F2: 24}
        _ = fi34
        _ = v35
    }
    v36 := []int{}
    for fi37 := 0; fi37 < 11; fi37++ {
        v36 = append(v36, (fi37 * (fi37 >> (((v6 + v6)) & 63))))
        for fi38 := 0; fi38 < 13; fi38++ {
            var v39 uint8 = 2
            fmt.Println(v3, v3, v39)
            var v40 uint8 = 161
            _ = fi38
            _ = v39
            _ = v40
        }
        _ = fi37
    }
    for fi41 := 0; fi41 < 12; fi41++ {
        if true {
            v42 := GenS3{F0: true, F1: 0}
            _ = v42
        } else {
            var v43 int8 = 0
            v44 := GenS1{F0: 73, F1: 460, F2: (!v15)}
            for fi45 := 0; fi45 < 4; fi45++ {
                var v46 uint64 = 0
                _ = fi45
                _ = v46
            }
            _ = v43
            _ = v44
        }
        _ = fi41
    }
    for fi47 := 0; fi47 < 4; fi47++ {
        if v15 {
            for fi48 := 0; fi48 < 9; fi48++ {
                switch 42 {
                case 500:
                    var v49 uint16 = 743
                    _ = v49
                case -239:
                    switch "\n\t" {
                    case "foo bar":
                        v50 := GenS3{F0: true, F1: 46}
                        _ = v50
                    case "\n\t":
                        var v51 string = ""
                        _ = v51
                    default:
                        if (122 <= 42) {
                            var v52 int16 = 1000
                            v52 -= v52
                            v53 := GenS4{F0: 127, F1: 2}
                            if (835 > 318) {
                                switch "a" {
                                case "foo bar":
                                    v54 := GenS5{F0: 233}
                                    _ = v54
                                case "":
                                    v36 = append(v36, (fi47 - (fi47 ^ fi48)))
                                default:
                                    v5[(fi47 * (fi48 & (fi47 % ((237) | 1))))] = v15
                                }
                                var v55 float64 = -2.25
                                fmt.Println(fi48)
                                var v56 bool = v15
                                fmt.Println(fi47, v15, v3)
                                _ = v55
                                _ = v56
                            } else {
                                fmt.Println(v1, fi48, fi47)
                                v57 := GenS3{F0: (42 <= 42), F1: 1}
                                v6 -= (v6 - (v6 & 1))
                                v3 -= v3
                                _ = v57
                            }
                            _ = v52
                            _ = v53
                        } else {
                            var v58 int = genFn1(127, fi48)
                            _ = v58
                        }
                    }
                case 32767:
                    v3 = -1000
                default:
                    _ = len(v4)
                }
                v14 -= v14
                _ = fi48
            }
            var v59 float32 = 0.1
            v14 = (v14 + v14)
            _ = v59
        }
        for fi60 := 0; fi60 < 18; fi60++ {
            var v61 float32 = 0.0
            v62 := GenS1{F0: 4294967295, F1: 656, F2: v15}
            _ = fi60
            _ = v61
            _ = v62
        }
        _ = fi47
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v36:", len(v36), v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
