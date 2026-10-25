package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16 }
type GenS2 struct { F0 int; F1 int }
type GenS3 struct { F0 int; F1 uint64 }
func genFn0(p0 int16) int16 {
    var v7 int64 = -9223372036854775808
    var v8 int = -41
    v8 = 128
    v7 = v7
    _ = p0
    _ = v7
    _ = v8
    return (-(p0))
}

func main() {
    var v1 string = "\n\t"
    v2 := []string{}
    v3 := make(map[int]uint32)
    v4 := make(map[uint32]string)
    var v5 int64 = -1
    var v6 int64 = v5
    fmt.Println(v6)
    _ = len(v2)
    for fi9 := 0; fi9 < 6; fi9++ {
        v1 = ""
        _ = fi9
    }
    switch 32768 {
    case 47:
        var v10 uint = 1
        _ = v10
    case 199:
        for fi11 := 0; fi11 < 3; fi11++ {
            var v12 float64 = (((1e10 + 3.14159) / (3.14159 + 3.14159)) + 3.14159)
            v5 = v5
            _ = fi11
            _ = v12
        }
    default:
        fmt.Println(v6)
    }
    for fi13 := 0; fi13 < 14; fi13++ {
        for fi14 := 0; fi14 < 20; fi14++ {
            var v15 int = fi13
            _ = fi14
            _ = v15
        }
        fmt.Println(v5, fi13, v5)
        for fi16 := 0; fi16 < 1; fi16++ {
            if (!(!false)) {
                fmt.Println(fi13)
                fmt.Println(v1, fi16)
                v17 := v3[fi13]
                var v18 int = fi16
                var v19 uint16 = 350
                _ = v17
                _ = v18
                _ = v19
            } else {
                fmt.Println(v5, v5)
                _ = len(v2)
                fmt.Println(fi16, fi13)
            }
            _ = fi16
        }
        _ = fi13
    }
    for fi20 := 0; fi20 < 8; fi20++ {
        v6 = -1
        v21 := GenS1{F0: (v5 | -9223372036854775808), F1: 479}
        if false {
            var v22 float64 = float64(1e10)
            v22 = (v22 / v22)
            v22 -= v22
            _ = v22
        }
        _ = fi20
        _ = v21
    }
    switch -783 {
    case 127:
        v6 = (^(v6))
    default:
        switch 696 {
        case 32768:
            switch (v1 + v1) {
            case "hello":
                switch "a" {
                case "hello":
                    if ((1 < 4294967295) && true) {
                        if true {
                            var v23 int = 32768
                            for fi24 := 0; fi24 < 8; fi24++ {
                                v25 := GenS1{F0: v5, F1: 2}
                                v6 += int64(fi24)
                                _ = fi24
                                _ = v25
                            }
                            v26 := GenS1{F0: (-(v5)), F1: 146}
                            _ = v23
                            _ = v26
                        } else {
                            switch 127 {
                            case -2147483648:
                                fmt.Println(v5, v6)
                            case 448:
                                if true {
                                    v2 = append(v2, v1)
                                }
                            }
                            _ = len(v3)
                        }
                        v6 = (v6 / (((v5 | v5)) | 1))
                        for fi27 := 0; fi27 < 13; fi27++ {
                            var v28 uint16 = 42
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            _ = fi27
                            _ = v28
                        }
                    }
                case "":
                    fmt.Println(v6, v6)
                case "foo bar":
                    v6 += v5
                }
            case "\n\t":
                switch 910 {
                case 135:
                    v5 += (v6 << ((234) & 63))
                case -2147483648:
                    fmt.Println(v1)
                default:
                    v29 := GenS1{F0: -488, F1: 1}
                    _ = v29
                }
            default:
                var v30 uint8 = 1
                _ = v30
            }
        case -26:
            var v31 int = -2147483648
            _ = v31
        case 0:
            switch -2147483648 {
            case 69:
                if (v6 < (v6 >> ((2) & 63))) {
                    v32 := GenS1{F0: (-(v6)), F1: 42}
                    var v33 int32 = 130
                    var v34 int = 363
                    _ = v32
                    _ = v33
                    _ = v34
                } else {
                    if len(v2) > 0 {
                        v2[0] = ""
                    }
                }
            case 2147483647:
                for fi35 := 0; fi35 < 17; fi35++ {
                    v1 = v1
                    _ = fi35
                }
            case 128:
                for fi36 := 0; fi36 < 11; fi36++ {
                    v37 := GenS2{F0: 128, F1: (^(fi36))}
                    for fi38 := 0; fi38 < 7; fi38++ {
                        v39 := GenS3{F0: (fi38 / ((1) | 1)), F1: 217}
                        v6 = v5
                        var v40 int16 = -32768
                        _ = fi38
                        _ = v39
                        _ = v40
                    }
                    fmt.Println(v1)
                    _ = fi36
                    _ = v37
                }
            }
        default:
            delete(v4, 4294967295)
        }
    }
    fmt.Println(v1, v6, v5)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v6 = (v5 & (^(v6)))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
