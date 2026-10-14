package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int; F1 bool }
type GenS3 struct { F0 int32; F1 uint64 }
type GenS4 struct { F0 uint64; F1 int8; F2 bool }

func main() {
    var v1 uint8 = 255
    var v2 uint32 = uint32(v1)
    var v3 uint8 = 0
    var v4 int64 = 9223372036854775807
    v5 := []uint32{(v2 >> ((2) & 63)), v2, (v2 | (v2 / ((0) | 1))), (v2 | 621)}
    if false {
        switch 313 {
        case 32768:
            fmt.Println(v4)
        case 2147483647:
            var v6 int = 2147483647
            _ = v6
        case 128:
            fmt.Println(v4, v4)
        default:
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        }
    } else {
        if len(v5) > 0 {
            v5[0] = v2
        }
        for fi7 := 0; fi7 < 12; fi7++ {
            if len(v5) > 0 {
                v5[0] = v2
            }
            v8 := GenS1{F0: 32767}
            if true {
                _ = len(v5)
                v4 = 42
                for fi9 := 0; fi9 < 12; fi9++ {
                    switch fi9 {
                    case 0:
                        var v10 int = 32767
                        _ = v10
                    case -9223372036854775808:
                        switch "\n\t" {
                        case "\n\t":
                            var v11 float32 = float32(3.14159)
                            _ = v11
                        default:
                            v1 = v1
                        }
                    }
                    var v12 int64 = 16
                    var v13 uint64 = 0
                    _ = fi9
                    _ = v12
                    _ = v13
                }
            } else {
                v4 = v4
                v3 += v1
                fmt.Println(v2, v3, v3)
            }
            _ = fi7
            _ = v8
        }
    }
    v1 = 255
    var v14 int = -392
    for fi15 := 0; fi15 < 1; fi15++ {
        if false {
            _ = len(v5)
            if true {
                v16 := GenS2{F0: (v14 * 4294967295), F1: (true || false)}
                if false {
                    v3 = 25
                    switch "a" {
                    case "":
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                    case "\n\t":
                        v17 := GenS3{F0: -815, F1: 2}
                        _ = v17
                    case "foo bar":
                        var v18 int64 = (^(v4))
                        _ = v18
                    }
                } else {
                    if true {
                        for fi19 := 0; fi19 < 15; fi19++ {
                            if (!(!(!(18446744073709551615 <= 0)))) {
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                                var v20 int8 = 42
                                _ = v20
                            }
                            _ = fi19
                        }
                    } else {
                        v21 := GenS3{F0: 304, F1: 1}
                        var v22 string = ("foo bar" + ("hello" + (("foo bar" + "hello") + "a")))
                        var v23 float64 = (-2.25 + float64((0.1 / 3.14159)))
                        _ = v21
                        _ = v22
                        _ = v23
                    }
                    if false {
                        var v24 uint64 = 691
                        _ = v24
                    }
                }
                var v25 uint8 = v3
                _ = v16
                _ = v25
            } else {
                if true {
                    if true {
                        var v26 uint8 = 127
                        _ = v26
                    } else {
                        fmt.Println(fi15)
                        if (false && (442 <= 200)) {
                            if (v4 == v4) {
                                var v27 bool = (2 <= 404)
                                switch "a" {
                                case "":
                                    v5 = append(v5, (v2 | (^(v2))))
                                case "foo bar":
                                    if (3.14159 >= (1.5 / (((0.1 * 1e10) + 3.14159) + float64((-2.25 / float32(0.1)))))) {
                                        v14 = 4294967295
                                        var v28 int = int(v4)
                                        _ = v28
                                    }
                                default:
                                    var v29 int = (^(fi15))
                                    _ = v29
                                }
                                _ = v27
                            } else {
                                var v30 int = 401
                                var v31 int = (-(v30))
                                v5 = append(v5, (v2 << ((3) & 63)))
                                _ = v30
                                _ = v31
                            }
                            if true {
                                for fi32 := 0; fi32 < 17; fi32++ {
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                    _ = fi32
                                }
                            } else {
                                v5 = append(v5, 2147483647)
                            }
                        } else {
                            v14 += fi15
                        }
                    }
                    var v33 int32 = 42
                    _ = v33
                } else {
                    switch (fi15 % ((32768) | 1)) {
                    case -1:
                        if true {
                            v2 = uint32(v1)
                            v34 := GenS3{F0: 2147483647, F1: 902}
                            _ = v34
                        }
                    case -825:
                        v1 += (v3 * v3)
                    }
                    if (1 <= 2) {
                        v3 = 2
                        var v35 uint16 = 2
                        _ = v35
                    }
                    var v36 uint64 = 2
                    _ = v36
                }
                var v37 uint8 = 1
                _ = v37
            }
            v38 := GenS4{F0: 2, F1: -128, F2: true}
            _ = v38
        } else {
            v2 += (v2 & v2)
            var v39 float32 = 3.14159
            _ = v39
        }
        _ = fi15
    }
    fmt.Println(v3)
    fmt.Println(v1, v14, v3)
    v2 = uint32(v1)
    v1 = v3
    switch -9223372036854775808 {
    case 128:
        var v40 uint8 = v1
        _ = v40
    default:
        var v41 bool = true
        _ = v41
    }
    if (int32(v4) != 2147483647) {
        for fi42 := 0; fi42 < 8; fi42++ {
            if (-1 == 912) {
                v43 := GenS2{F0: (fi42 >> ((255) & 63)), F1: ((false || true) || (1e10 > (0.1 * -2.25)))}
                for fi44 := 0; fi44 < 3; fi44++ {
                    var v45 uint32 = (v2 * (^(v2)))
                    fmt.Println(fi42, v1)
                    v46 := GenS1{F0: -1}
                    _ = fi44
                    _ = v45
                    _ = v46
                }
                var v47 uint8 = v3
                _ = v43
                _ = v47
            }
            v4 = v4
            v1 -= (v3 + (^(v3)))
            _ = fi42
        }
    }
    for fi48 := 0; fi48 < 14; fi48++ {
        v5 = append(v5, (v2 & v2))
        _ = fi48
    }
    v5 = append(v5, (v2 | (^(v2))))
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
