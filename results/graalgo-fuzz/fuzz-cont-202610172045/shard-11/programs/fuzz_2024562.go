package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint; F2 uint8 }
type GenS2 struct { F0 string; F1 int }
type GenS3 struct { F0 int; F1 int; F2 uint8 }

func main() {
    v1 := []int{127, 32767, -2147483648}
    var v2 string = "foo bar"
    v3 := make(map[uint32]int)
    v4 := []string{("foo bar" + "a"), v2}
    var v5 uint16 = 575
    var v6 uint = 1
    for fi7 := 0; fi7 < 3; fi7++ {
        fmt.Println(fi7, v2, fi7)
        v6 = v6
        _ = fi7
    }
    if ((42 == 1) || ((-696 != 42) && false)) {
        var v8 int64 = 0
        for fi9 := 0; fi9 < 19; fi9++ {
            _ = len(v3)
            _ = fi9
        }
        _ = v8
    }
    var v10 uint64 = 4294967295
    var v11 int64 = 121
    fmt.Println(v5, v10, v6)
    var v12 bool = false
    v10 = (v10 / (((v10 << (((v6 & v6)) & 63))) | 1))
    v3[1] = 4294967295
    var v13 uint16 = (v5 | v5)
    delete(v3, 0)
    for fi14 := 0; fi14 < 8; fi14++ {
        for fi15 := 0; fi15 < 2; fi15++ {
            v13 = (v5 % ((v13) | 1))
            fmt.Println(v6, v11)
            v10 += v10
            v10 = (v10 / ((v10) | 1))
            _ = fi15
        }
        v6 = v6
        _ = fi14
    }
    if ((83 != 56) || (!v12)) {
        v16 := GenS1{F0: 42, F1: 321, F2: 161}
        _ = v16
    }
    v11 = (v11 % ((273) | 1))
    if (71 < 1) {
        fmt.Println(v11)
        if ((false && (1 != 958)) || v12) {
            _ = len(v3)
        } else {
            var v17 uint = v6
            var v18 uint = v6
            v19 := GenS1{F0: 18446744073709551615, F1: (v6 << ((v18) & 63)), F2: 94}
            _ = v17
            _ = v18
            _ = v19
        }
        for fi20 := 0; fi20 < 14; fi20++ {
            v12 = (1 != 53)
            for fi21 := 0; fi21 < 5; fi21++ {
                delete(v3, 18)
                switch int(v11) {
                case -2147483648:
                    if (v12 || (-1 > 42)) {
                        delete(v3, 0)
                        var v22 int8 = 111
                        var v23 uint8 = 42
                        _ = v22
                        _ = v23
                    } else {
                        var v24 bool = v12
                        v25 := GenS2{F0: (v2 + ""), F1: fi20}
                        v13 = (^(v5))
                        _ = v24
                        _ = v25
                    }
                default:
                    var v26 int64 = v11
                    _ = v26
                }
                v27 := GenS1{F0: v6, F1: 280, F2: 128}
                fmt.Println(fi21, fi21)
                v4 = append(v4, (v2 + v2))
                fmt.Println(fi20, v2)
                switch v2 {
                case "a":
                    v28 := GenS1{F0: uint(v10), F1: v6, F2: 77}
                    _ = v28
                case "\n\t":
                    var v29 int64 = (v11 ^ v11)
                    _ = v29
                case "hello":
                    v3[364] = (fi20 % (((fi21 / ((fi20) | 1))) | 1))
                default:
                    v11 = -1
                }
                if true {
                    v6 += v6
                }
                switch fi21 {
                case 32768:
                    var v30 string = v2
                    _ = v30
                case -9223372036854775808:
                    switch v2 {
                    case "hello":
                        var v31 uint8 = 255
                        _ = v31
                    case "":
                        switch fi20 {
                        case 0:
                            switch 4294967295 {
                            case 0:
                                var v32 float64 = 3.14159
                                _ = v32
                            case 32767:
                                v2 = "hello"
                            default:
                                var v33 int = (fi20 - (^(fi20)))
                                _ = v33
                            }
                        case -1:
                            var v34 uint8 = 45
                            _ = v34
                        case 32768:
                            if v12 {
                                var v35 float64 = -2.25
                                v11 = 1
                                _ = v35
                            } else {
                                switch fi20 {
                                case -1:
                                    switch v2 {
                                    case "hello":
                                        switch fi21 {
                                        case 1:
                                            v36 := GenS1{F0: (v6 >> ((v6) & 63)), F1: v6, F2: 42}
                                            if (42 <= (v10 >> ((v6) & 63))) {
                                                v11 -= (v11 ^ 1)
                                            } else {
                                                var v37 uint8 = 128
                                                _ = len(v1)
                                                _ = v37
                                            }
                                            v38 := v3[1]
                                            if v12 {
                                                v6 = v6
                                                switch ("" + "foo bar") {
                                                case "":
                                                    v13 = (^(v5))
                                                case "a":
                                                    if (65535 != v6) {
                                                        fmt.Println(v5, v38, v12)
                                                        fmt.Println(fi21, v6)
                                                    }
                                                }
                                                v38 = fi21
                                                var v39 int32 = 690
                                                _ = v39
                                            } else {
                                                var v40 int = fi21
                                                _ = v40
                                            }
                                            v12 = v12
                                            v41 := GenS2{F0: "\n\t", F1: v38}
                                            _ = v36
                                            _ = v38
                                            _ = v41
                                        default:
                                            if v12 {
                                                fmt.Println(fi20, v5, v11)
                                                var v42 int16 = -264
                                                _ = v42
                                            }
                                        }
                                    case "\n\t":
                                        v11 = (^(v11))
                                        v43 := GenS3{F0: int(v6), F1: (^(fi20)), F2: 1}
                                        _ = v43
                                    }
                                case 32767:
                                    fmt.Println(v6, v10)
                                }
                                switch (("" + v2) + v2) {
                                case "\n\t":
                                    v10 = v10
                                case "foo bar":
                                    _ = len(v3)
                                    fmt.Println(v5, v5)
                                }
                                v3[uint32(v13)] = fi21
                            }
                        }
                    case "\n\t":
                        switch v2 {
                        case "foo bar":
                            v44 := GenS2{F0: v2, F1: fi21}
                            _ = v44
                        }
                    default:
                        v12 = v12
                    }
                default:
                    var v45 int8 = -121
                    _ = v45
                }
                _ = fi21
                _ = v27
            }
            var v46 int8 = 1
            _ = fi20
            _ = v46
        }
    } else {
        var v47 int8 = 0
        for fi48 := 0; fi48 < 2; fi48++ {
            if false {
                var v49 uint8 = 39
                for fi50 := 0; fi50 < 10; fi50++ {
                    var v51 uint32 = 558
                    _ = fi50
                    _ = v51
                }
                fmt.Println(v10)
                _ = v49
            }
            var v52 int8 = -9
            if (!true) {
                for fi53 := 0; fi53 < 7; fi53++ {
                    v6 = v6
                    _ = fi53
                }
            }
            _ = fi48
            _ = v52
        }
        delete(v3, 4294967295)
        _ = v47
    }
    v54 := GenS2{F0: v2, F1: 1}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v54
    _ = v6
}
