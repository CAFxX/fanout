package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint16; F1 int32; F2 int32 }
type GenS3 struct { F0 string; F1 uint16 }
type GenS4 struct { F0 int64; F1 int }
type GenS5 struct { F0 uint16; F1 int32; F2 bool }
type GenS6 struct { F0 uint64; F1 uint8 }
func genFn0(p0 int) uint64 {
    for fi5 := 0; fi5 < 4; fi5++ {
        var v6 int64 = 1
        _ = fi5
        _ = v6
    }
    var v7 uint8 = 171
    _ = p0
    _ = v7
    return 18446744073709551615
}
func genFn1(p0 uint16) int32 {
    if true {
        p0 = 1
    } else {
        for fi8 := 0; fi8 < 17; fi8++ {
            p0 += p0
            if true {
                if false {
                    var v9 int = 127
                    _ = v9
                } else {
                    var v10 uint = 406
                    fmt.Println(fi8, p0)
                    v10 = v10
                    _ = v10
                }
                var v11 uint64 = 931
                if false {
                    p0 = p0
                    var v12 uint8 = 16
                    _ = v12
                }
                _ = v11
            }
            _ = fi8
        }
    }
    fmt.Println(p0)
    for fi13 := 0; fi13 < 10; fi13++ {
        v14 := GenS1{F0: -1}
        p0 += p0
        if (100 < 127) {
            switch "" {
            case "a":
                var v15 string = "\n\t"
                _ = v15
            case "foo bar":
                fmt.Println(fi13, fi13)
            default:
                p0 = p0
            }
            if (137 <= 1) {
                v16 := []int{}
                _ = v16
            }
        } else {
            fmt.Println(fi13)
        }
        _ = fi13
        _ = v14
    }
    v17 := GenS2{F0: (p0 | p0), F1: -944, F2: 42}
    var v18 uint8 = 1
    _ = p0
    _ = v17
    _ = v18
    return 0
}

func main() {
    v1 := []uint32{2147483648, 0}
    var v2 int64 = 9223372036854775807
    var v3 int16 = 1000
    v4 := []string{("a" + "\n\t"), ("\n\t" + "foo bar"), "", "hello"}
    v19 := GenS3{F0: "hello", F1: 65535}
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v3 += -247
    var v20 float64 = 0.0
    v2 -= 42
    v2 = v2
    if true {
        if true {
            var v21 int32 = -748
            _ = v21
        } else {
            var v22 uint = 2
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = v22
        }
        switch -9223372036854775808 {
        case -244:
            _ = len(v1)
        case 0:
            fmt.Println(v3, v20, v20)
        }
    }
    switch 734 {
    case 32768:
        if (0 == 671) {
            v3 = (v3 + v3)
            var v23 int8 = -1
            _ = v23
        }
    case 47:
        if false {
            v24 := GenS2{F0: 65535, F1: 1, F2: -46}
            _ = v24
        }
    default:
        v25 := GenS3{F0: "\n\t", F1: 255}
        _ = v25
    }
    if (128 == 255) {
        for fi26 := 0; fi26 < 2; fi26++ {
            var v27 float64 = ((((1.5 - 0.1) - v20) * float64(1e10)) * -2.25)
            v28 := GenS4{F0: v2, F1: (fi26 ^ fi26)}
            _ = fi26
            _ = v27
            _ = v28
        }
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    } else {
        fmt.Println(v20, v3, v20)
        for fi29 := 0; fi29 < 15; fi29++ {
            v30 := GenS2{F0: 730, F1: 42, F2: 724}
            v20 = (-2.25 * v20)
            for fi31 := 0; fi31 < 11; fi31++ {
                fmt.Println(v3)
                v20 -= v20
                _ = fi31
            }
            _ = fi29
            _ = v30
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    if (!false) {
        var v32 int64 = (v2 - (v2 / (((v2 - (v2 / ((v2) | 1)))) | 1)))
        if (1 >= 65535) {
            v33 := GenS2{F0: 42, F1: -39, F2: 2147483647}
            _ = v33
        }
        v34 := GenS2{F0: 972, F1: -156, F2: 42}
        _ = v32
        _ = v34
    }
    for fi35 := 0; fi35 < 4; fi35++ {
        var v36 uint16 = 32767
        _ = fi35
        _ = v36
    }
    switch 2147483647 {
    case 4294967295:
        switch "" {
        case "a":
            var v37 float32 = 0.0
            _ = v37
        }
    case 977:
        v3 = v3
    case -109:
        fmt.Println(v3, v20, v2)
    default:
        v3 = 0
    }
    var v38 uint16 = 255
    switch ((("a" + "hello") + "a") + "a") {
    case "foo bar":
        if true {
            switch -1 {
            case -125:
                v4 = append(v4, "hello")
            case 2147483647:
                v39 := GenS5{F0: v38, F1: 685, F2: true}
                _ = v39
            }
            for fi40 := 0; fi40 < 3; fi40++ {
                var v41 uint32 = 8
                _ = len(v4)
                if ((v3 | 1000) >= v3) {
                    switch "foo bar" {
                    case "a":
                        var v42 int16 = -32768
                        _ = v42
                    case "":
                        v3 += -1
                    }
                    v38 += v38
                    v43 := GenS5{F0: v38, F1: 2147483647, F2: false}
                    _ = v43
                }
                _ = fi40
                _ = v41
            }
            switch 1 {
            case 32768:
                v44 := GenS1{F0: 0}
                _ = v44
            }
        } else {
            v1 = append(v1, 2)
            for fi45 := 0; fi45 < 10; fi45++ {
                if true {
                    for fi46 := 0; fi46 < 7; fi46++ {
                        if true {
                            _ = len(v1)
                            var v47 int16 = -36
                            v1 = append(v1, 42)
                            _ = v47
                        }
                        switch ("a" + "") {
                        case "":
                            if (false && (2 < genFn0((fi46 - fi45)))) {
                                var v48 bool = ((fi45 | fi45) < (fi46 * (fi46 * (fi45 & fi45))))
                                v2 -= 1
                                v49 := GenS6{F0: 0, F1: 142}
                                _ = v48
                                _ = v49
                            }
                        case "a":
                            switch fi46 {
                            case 1:
                                var v50 int32 = 2147483647
                                _ = v50
                            case 127:
                                if false {
                                    fmt.Println(v38, fi45)
                                    fmt.Println(v3, v2, fi46)
                                }
                            case -9223372036854775808:
                                var v51 int8 = 0
                                _ = v51
                            default:
                                fmt.Println(v38, fi45)
                            }
                        case "\n\t":
                            var v52 float64 = (v20 / (0.1 - float64(-2.25)))
                            _ = v52
                        default:
                            v53 := GenS6{F0: 486, F1: 42}
                            _ = v53
                        }
                        fmt.Println(v38, fi46, fi45)
                        v20 = -2.25
                        _ = fi46
                    }
                }
                _ = fi45
            }
        }
    }
    v20 = (v20 / v20)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v38
    _ = v4
}
