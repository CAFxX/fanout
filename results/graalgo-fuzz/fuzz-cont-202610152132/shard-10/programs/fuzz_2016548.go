package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8; F2 int8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint32; F1 uint; F2 int }
type GenS4 struct { F0 int16 }
func genFn0() int32 {
    for fi4 := 0; fi4 < 15; fi4++ {
        var v5 uint8 = 128
        _ = fi4
        _ = v5
    }
    var v6 uint32 = 65535
    for fi7 := 0; fi7 < 12; fi7++ {
        for fi8 := 0; fi8 < 10; fi8++ {
            fmt.Println(fi7)
            v9 := []int{(fi8 ^ fi8)}
            if (1e10 <= 0.0) {
                var v10 uint = 18446744073709551615
                v11 := GenS1{F0: (v6 >> ((v10) & 63)), F1: 127, F2: -112}
                if (true && (true || (!(-2.25 == 0.1)))) {
                    switch "foo bar" {
                    case "":
                        var v12 uint8 = 220
                        _ = v12
                    case "foo bar":
                        switch "hello" {
                        case "hello":
                            if false {
                                v13 := GenS1{F0: 1, F1: 12, F2: 0}
                                v14 := GenS1{F0: (v6 ^ v6), F1: 6, F2: 100}
                                _ = v13
                                _ = v14
                            }
                        case "foo bar":
                            v6 -= 735
                        default:
                            if (81 != 155) {
                                var v15 int32 = 1
                                _ = v15
                            }
                        }
                    case "a":
                        fmt.Println(fi7)
                        var v16 int8 = 100
                        v16 = (v16 | -1)
                        var v17 uint16 = 1
                        v17 += v17
                        _ = v16
                        _ = v17
                    default:
                        var v18 int64 = 42
                        _ = v18
                    }
                    var v19 bool = false
                    if (v19 && v19) {
                        fmt.Println(fi7, v19, v19)
                        if (612 >= 679) {
                            if true {
                                v20 := GenS2{F0: true}
                                var v21 int = (fi8 | 0)
                                _ = v20
                                _ = v21
                            } else {
                                var v22 uint = 2
                                if (4294967295 == v22) {
                                    var v23 int8 = 0
                                    if (!((42 > 198) && ((-(fi7)) <= 32768))) {
                                        var v24 string = ("hello" + ("" + "a"))
                                        fmt.Println(fi7, v10)
                                        _ = v24
                                    } else {
                                        if v19 {
                                            if (!(!((v23 % ((42) | 1)) == v23))) {
                                                var v25 bool = v19
                                                _ = v25
                                            }
                                            _ = len(v9)
                                            v26 := GenS3{F0: (v6 | v6), F1: (v10 / ((v10) | 1)), F2: (fi8 << ((v22) & 63))}
                                            _ = v26
                                        }
                                    }
                                    _ = v23
                                }
                                v19 = ((((3.14159 - -2.25) + ((3.14159 / (1.5 - 1e10)) + 1e10)) * -2.25) > (0.1 - 0.0))
                                _ = v22
                            }
                        } else {
                            v10 = v10
                            v10 = 2
                            if true {
                                var v27 uint = (v10 + (v10 * (v10 % ((v10) | 1))))
                                if false {
                                    v19 = v19
                                } else {
                                    var v28 uint16 = 42
                                    _ = v28
                                }
                                _ = v27
                            }
                            v29 := GenS4{F0: 32767}
                            if v19 {
                                v19 = ((fi8 == -1) && (0 >= -1))
                                if (32767 < 1000) {
                                    var v30 float32 = 0.1
                                    v31 := GenS1{F0: v6, F1: 42, F2: 100}
                                    var v32 uint32 = v6
                                    _ = v30
                                    _ = v31
                                    _ = v32
                                }
                                var v33 int = fi7
                                _ = v33
                            } else {
                                _ = len(v9)
                                var v34 uint16 = 556
                                _ = v34
                            }
                            _ = v29
                        }
                        fmt.Println(fi7, fi7, v10)
                    }
                    var v35 float64 = 1.5
                    if len(v9) > 0 {
                        v9[0] = (fi8 & fi8)
                    }
                    _ = v19
                    _ = v35
                } else {
                    switch "foo bar" {
                    case "hello":
                        var v36 int16 = 0
                        _ = v36
                    case "\n\t":
                        if (!true) {
                            if len(v9) >= 2 {
                                v9 = v9[1:len(v9)]
                            }
                            fmt.Println(fi8, v10, v6)
                            switch (^(fi8)) {
                            case 32767:
                                if ((v6 / (((v6 & (v6 / ((v6) | 1)))) | 1)) > v6) {
                                    v37 := GenS1{F0: v6, F1: -1, F2: 42}
                                    _ = v37
                                }
                            case 886:
                                v10 = 0
                            case -219:
                                switch (^(fi7)) {
                                case 254:
                                    _ = len(v9)
                                default:
                                    v10 = 18446744073709551615
                                }
                            default:
                                fmt.Println(fi8, fi7, v6)
                            }
                        }
                    case "a":
                        var v38 int = 99
                        _ = v38
                    }
                }
                _ = v10
                _ = v11
            }
            _ = fi8
            _ = v9
        }
        _ = fi7
    }
    v39 := GenS1{F0: (v6 | 2147483647), F1: 35, F2: 0}
    _ = v39
    _ = v6
    return 2147483647
}

func main() {
    v1 := []int{0, -9223372036854775808, -1, 2147483647}
    var v2 int8 = 127
    var v3 float32 = -2.25
    var v40 uint64 = 1
    var v41 uint64 = v40
    if true {
        if (-1 == 0) {
            if len(v1) > 0 {
                v1[0] = 707
            }
            if len(v1) > 0 {
                v1[0] = 128
            }
        }
        var v42 uint = 65535
        fmt.Println(v3)
        _ = v42
    } else {
        var v43 float64 = ((1e10 / (((1e10 / 1e10) - (((0.1 / (-2.25 - 0.0)) / 1.0) - 0.1)) - 1.5)) - 1.5)
        _ = v43
    }
    for fi44 := 0; fi44 < 3; fi44++ {
        v2 = v2
        var v45 uint16 = 32767
        _ = fi44
        _ = v45
    }
    for fi46 := 0; fi46 < 7; fi46++ {
        for fi47 := 0; fi47 < 16; fi47++ {
            var v48 float64 = 1e10
            var v49 uint64 = (v41 * v41)
            _ = fi47
            _ = v48
            _ = v49
        }
        _ = fi46
    }
    if len(v1) > 0 {
        v1[0] = -363
    }
    var v50 uint64 = 0
    v41 += v41
    for fi51 := 0; fi51 < 1; fi51++ {
        switch fi51 {
        case 42:
            v3 -= -2.25
        default:
            v52 := GenS3{F0: 921, F1: 42, F2: fi51}
            _ = v52
        }
        v41 = v50
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi51
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v2
    _ = v3
    _ = v40
    _ = v41
    _ = v50
}
