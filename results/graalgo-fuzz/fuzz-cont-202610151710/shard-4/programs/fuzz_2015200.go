package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool; F2 uint8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint64; F1 bool }

func main() {
    v1 := []uint32{2147483648, 65535, 42}
    v2 := make(map[string]int)
    var v3 uint64 = 18446744073709551615
    var v4 uint32 = 2147483647
    v5 := make(map[int]bool)
    if len(v1) > 0 {
        v1[0] = (v4 ^ (v4 & v4))
    }
    var v6 string = "a"
    if false {
        v6 = "foo bar"
        v7 := GenS1{F0: 4294967295, F1: true, F2: 255}
        for fi8 := 0; fi8 < 13; fi8++ {
            v3 = (v3 & v3)
            if true {
                for fi9 := 0; fi9 < 18; fi9++ {
                    switch v6 {
                    case "\n\t":
                        v4 = v4
                    }
                    v4 = (v4 / ((v4) | 1))
                    var v10 uint64 = 42
                    v11 := GenS1{F0: v4, F1: false, F2: 0}
                    _ = fi9
                    _ = v10
                    _ = v11
                }
                v12 := GenS1{F0: (v4 >> ((908) & 63)), F1: true, F2: 128}
                switch "foo bar" {
                case "foo bar":
                    var v13 int = fi8
                    _ = v13
                case "\n\t":
                    fmt.Println(fi8, v3)
                default:
                    fmt.Println(v4, v3)
                }
                _ = v12
            } else {
                _ = len(v5)
                fmt.Println(v6)
            }
            v14 := GenS1{F0: (v4 / ((2147483647) | 1)), F1: (!true), F2: 128}
            _ = fi8
            _ = v14
        }
        _ = v7
    }
    v6 = ("hello" + "")
    for fi15 := 0; fi15 < 14; fi15++ {
        if false {
            fmt.Println(v6, v4)
        } else {
            switch v6 {
            case "\n\t":
                fmt.Println(fi15, v4)
            case "hello":
                if true {
                    if false {
                        if false {
                            var v16 int8 = 42
                            _ = v16
                        } else {
                            fmt.Println(fi15, v6)
                            if ((v4 != (v4 * v4)) || (-2147483648 < -1)) {
                                if (false && true) {
                                    fmt.Println(v4, v3)
                                    if (!false) {
                                        var v17 uint16 = 588
                                        _ = v17
                                    }
                                    if true {
                                        if ((false && true) || (((0.1 + (0.1 + 3.14159)) + 1.5) == 0.0)) {
                                            v6 = v6
                                            v18 := GenS1{F0: v4, F1: ((v4 >> ((1) & 63)) == v4), F2: 242}
                                            _ = v18
                                        } else {
                                            v6 = "\n\t"
                                        }
                                        _ = len(v5)
                                    }
                                } else {
                                    fmt.Println(v3)
                                    for fi19 := 0; fi19 < 3; fi19++ {
                                        var v20 int32 = 2147483647
                                        _ = fi19
                                        _ = v20
                                    }
                                }
                                v21 := GenS2{F0: 1}
                                _ = v21
                            } else {
                                var v22 int64 = 9223372036854775807
                                var v23 bool = true
                                v3 += (v3 >> ((2) & 63))
                                _ = v22
                                _ = v23
                            }
                            v3 -= v3
                        }
                    } else {
                        for fi24 := 0; fi24 < 14; fi24++ {
                            v6 = v6
                            v6 = "a"
                            _ = len(v2)
                            _ = fi24
                        }
                        var v25 string = "\n\t"
                        _ = v25
                    }
                } else {
                    var v26 uint64 = (^(v3))
                    var v27 float64 = 0.1
                    _ = len(v1)
                    _ = v26
                    _ = v27
                }
            case "":
                v28 := GenS3{F0: v6}
                _ = v28
            }
            v3 = (v3 << ((18446744073709551615) & 63))
        }
        v29 := GenS4{F0: (fi15 << ((18446744073709551615) & 63))}
        var v30 float64 = (1e10 * -2.25)
        _ = fi15
        _ = v29
        _ = v30
    }
    v31 := GenS4{F0: 1}
    v5[-2147483648] = ((v4 / ((v4) | 1)) != (v4 >> ((232) & 63)))
    for fi32 := 0; fi32 < 10; fi32++ {
        for fi33 := 0; fi33 < 19; fi33++ {
            fmt.Println(v6, v3)
            switch (fi33 << ((1) & 63)) {
            case 32767:
                v34 := GenS4{F0: 128}
                _ = v34
            case -183:
                _ = len(v1)
            default:
                if true {
                    var v35 int16 = 32767
                    delete(v2, v6)
                    if (v4 < 2147483648) {
                        v1 = append(v1, (v4 - v4))
                        v36 := GenS4{F0: (fi32 * 968)}
                        delete(v5, (fi33 * (fi33 >> ((295) & 63))))
                        var v37 int16 = (v35 << ((255) & 63))
                        var v38 bool = (3.14159 == 0.0)
                        _ = v36
                        _ = v37
                        _ = v38
                    } else {
                        v35 -= v35
                    }
                    _ = v35
                } else {
                    v6 = v6
                }
            }
            var v39 int32 = -355
            v6 = "foo bar"
            _ = fi33
            _ = v39
        }
        _ = fi32
    }
    var v40 int8 = -128
    v1 = append(v1, 0)
    var v41 int8 = v40
    if true {
        v40 += v40
    } else {
        v42 := GenS1{F0: (v4 % ((v4) | 1)), F1: true, F2: 127}
        v40 += v41
        switch -1 {
        case 912:
            v43 := v2["\n\t"]
            _ = v43
        case 32768:
            var v44 string = "hello"
            _ = v44
        }
        _ = v42
    }
    v1 = append(v1, (v4 | v4))
    var v45 int8 = -7
    v46 := GenS5{F0: (v3 - 18446744073709551615), F1: (1 < 0)}
    for fi47 := 0; fi47 < 4; fi47++ {
        var v48 uint16 = 255
        _ = fi47
        _ = v48
    }
    v49 := GenS5{F0: v3, F1: true}
    for fi50 := 0; fi50 < 16; fi50++ {
        if false {
            fmt.Println(v3, fi50, v6)
            if (42 == 2) {
                v51 := GenS2{F0: (^(v3))}
                for fi52 := 0; fi52 < 15; fi52++ {
                    if (("foo bar" != ((v6 + v6) + ("" + v6))) || true) {
                        fmt.Println(v41, v3, v45)
                    }
                    _ = fi52
                }
                var v53 int64 = 1
                _ = v51
                _ = v53
            }
        } else {
            var v54 string = v6
            if (42 > -681) {
                var v55 int64 = 601
                v45 += 70
                _ = v55
            }
            _ = v54
        }
        if true {
            v45 = v45
        }
        _ = fi50
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v45:", v45)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v40
    _ = v41
    _ = v45
    _ = v46
    _ = v49
    _ = v5
    _ = v6
}
