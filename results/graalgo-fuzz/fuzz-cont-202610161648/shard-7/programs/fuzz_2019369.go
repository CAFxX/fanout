package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint64; F1 int16; F2 uint16 }
type GenS4 struct { F0 string; F1 bool; F2 int32 }
type GenS5 struct { F0 int16; F1 int8; F2 uint }
type GenS6 struct { F0 int32; F1 bool; F2 uint }

func main() {
    var v1 float32 = (3.14159 + (0.0 - (float32((((float64((1e10 - (1e10 + float32(0.1)))) * 0.0) + 0.0) - float64(0.1))) * ((3.14159 + 1e10) / 1e10))))
    var v2 bool = (-1 <= -1)
    var v3 int = 128
    v4 := []string{"foo bar", "", "a"}
    var v5 uint64 = 116
    v5 = 461
    var v6 int64 = 9223372036854775807
    switch 0 {
    case 42:
        if (32767 != 1) {
            var v7 int8 = 25
            switch v3 {
            case -2147483648:
                switch v3 {
                case 128:
                    fmt.Println(v6, v5, v5)
                }
            case 16:
                v8 := GenS1{F0: v6, F1: -48}
                _ = v8
            }
            _ = v7
        } else {
            var v9 int = v3
            if false {
                var v10 uint64 = v5
                if len(v4) > 0 {
                    v4[0] = "a"
                }
                _ = v10
            } else {
                switch "\n\t" {
                case "foo bar":
                    v11 := GenS2{F0: 50}
                    _ = v11
                case "a":
                    switch "hello" {
                    case "\n\t":
                        if (2147483648 != 4294967295) {
                            var v12 int64 = v6
                            if ((^(v12)) == v6) {
                                v13 := GenS3{F0: (v5 << ((2) & 63)), F1: int16(v6), F2: 6}
                                _ = v13
                            }
                            v2 = (((^(v5)) >= v5) && v2)
                            _ = v12
                        } else {
                            fmt.Println(v3, v1)
                            for fi14 := 0; fi14 < 18; fi14++ {
                                var v15 float64 = 0.0
                                fmt.Println(fi14)
                                if v2 {
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                    v3 = v3
                                } else {
                                    if true {
                                        for fi16 := 0; fi16 < 2; fi16++ {
                                            v17 := GenS1{F0: 0, F1: 32767}
                                            _ = fi16
                                            _ = v17
                                        }
                                    }
                                }
                                _ = fi14
                                _ = v15
                            }
                        }
                    case "":
                        if ((!false) && (v2 || (0.0 <= (v1 * v1)))) {
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            v6 = 42
                        }
                    default:
                        var v18 float64 = (-2.25 - ((3.14159 - 1.5) - 3.14159))
                        _ = v18
                    }
                default:
                    if v2 {
                        v19 := GenS3{F0: 4294967295, F1: 42, F2: 2}
                        v20 := GenS4{F0: "foo bar", F1: ((-679 < 1) || ((!(127 >= 127)) || v2)), F2: 1}
                        _ = v19
                        _ = v20
                    } else {
                        switch "a" {
                        case "a":
                            switch v3 {
                            case -2147483648:
                                var v21 uint8 = 2
                                _ = v21
                            case -141:
                                for fi22 := 0; fi22 < 13; fi22++ {
                                    fmt.Println(v5)
                                    for fi23 := 0; fi23 < 3; fi23++ {
                                        v9 += (fi23 % ((v9) | 1))
                                        _ = fi23
                                    }
                                    _ = fi22
                                }
                            case 2147483647:
                                var v24 int64 = (v6 - v6)
                                _ = v24
                            default:
                                v25 := GenS5{F0: 1, F1: -17, F2: 131}
                                _ = v25
                            }
                        }
                        var v26 bool = (v2 && v2)
                        _ = v26
                    }
                }
                fmt.Println(v3, v9, v5)
            }
            v1 -= v1
            _ = v9
        }
    case -2147483648:
        if len(v4) > 0 {
            v4[0] = ("" + "foo bar")
        }
    default:
        var v27 uint = 42
        _ = v27
    }
    _ = len(v4)
    if false {
        if (false && (v5 < v5)) {
            _ = len(v4)
        } else {
            fmt.Println(v2, v5)
            var v28 uint = 65535
            if v2 {
                v4 = append(v4, ("foo bar" + "hello"))
            }
            _ = v28
        }
        var v29 int16 = 1
        _ = v29
    }
    v4 = append(v4, "\n\t")
    v4 = append(v4, ("\n\t" + "a"))
    var v30 uint64 = v5
    var v31 uint64 = v30
    var v32 uint32 = 803
    switch v3 {
    case 377:
        switch "" {
        case "":
            v33 := GenS3{F0: v30, F1: 1000, F2: 522}
            _ = v33
        case "\n\t":
            _ = len(v4)
        case "foo bar":
            var v34 bool = v2
            _ = v34
        default:
            for fi35 := 0; fi35 < 3; fi35++ {
                v2 = (true && ((87 == -1) && v2))
                _ = fi35
            }
        }
    default:
        for fi36 := 0; fi36 < 18; fi36++ {
            var v37 bool = (!(v2 && (-404 != 952)))
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            if (-87 < -120) {
                v38 := GenS6{F0: 856, F1: (-2.25 >= ((((3.14159 + 3.14159) - 1e10) + 0.0) / ((3.14159 / 1e10) / 1e10))), F2: 65535}
                _ = v38
            }
            _ = fi36
            _ = v37
        }
    }
    var v39 float32 = v1
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    for fi40 := 0; fi40 < 17; fi40++ {
        var v41 uint8 = 128
        fmt.Println(v5)
        _ = fi40
        _ = v41
    }
    fmt.Println(v32, v2)
    if len(v4) > 0 {
        v4[0] = "\n\t"
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
