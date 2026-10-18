package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8 }
type GenS2 struct { F0 uint64; F1 uint8; F2 uint32 }
type GenS3 struct { F0 int8; F1 uint16; F2 bool }
type GenS4 struct { F0 uint32; F1 int; F2 uint }
type GenS5 struct { F0 bool; F1 int64 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 int; F1 int16 }
type GenS8 struct { F0 uint; F1 int64; F2 int8 }

func main() {
    v1 := []string{"a", "\n\t"}
    var v2 uint64 = 177
    v3 := []string{"", ("" + "\n\t")}
    v4 := make(map[uint32]string)
    var v5 string = "\n\t"
    var v6 string = ((v5 + v5) + v5)
    var v7 float64 = (1.5 - (((-2.25 * 0.0) * float64((1.5 - float32(3.14159)))) - ((1.5 / -2.25) - (3.14159 * 1e10))))
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v5 = "\n\t"
    v6 = v5
    for fi8 := 0; fi8 < 4; fi8++ {
        for fi9 := 0; fi9 < 18; fi9++ {
            if (false && true) {
                var v10 int64 = 9223372036854775807
                switch v5 {
                case "a":
                    switch "a" {
                    case "hello":
                        v6 = ("\n\t" + ("foo bar" + "hello"))
                    case "a":
                        switch fi8 {
                        case 32768:
                            v11 := GenS1{F0: 1, F1: 30}
                            _ = v11
                        case -393:
                            var v12 uint16 = 715
                            _ = v12
                        case -9223372036854775808:
                            var v13 float32 = (-2.25 * (((0.0 + (-2.25 + 1e10)) + 3.14159) / 1e10))
                            _ = v13
                        }
                    }
                case "\n\t":
                    if false {
                        var v14 int32 = 0
                        _ = v14
                    } else {
                        fmt.Println(fi8, v6)
                        v15 := GenS2{F0: v2, F1: 82, F2: 2}
                        v5 = v5
                        _ = v15
                    }
                case "foo bar":
                    switch ((v6 + ("hello" + "a")) + (((v5 + v6) + (v6 + v5)) + "a")) {
                    case "hello":
                        fmt.Println(v6, v2, v7)
                        if false {
                            v10 = v10
                        }
                    }
                }
                _ = v10
            } else {
                v2 -= v2
                v2 += v2
                var v16 int32 = 225
                _ = v16
            }
            _ = fi9
        }
        var v17 int32 = -2147483648
        _ = fi8
        _ = v17
    }
    var v18 string = v5
    if (2147483647 < 0) {
        if true {
            for fi19 := 0; fi19 < 8; fi19++ {
                v20 := GenS3{F0: -1, F1: 886, F2: ((("foo bar" + (((v5 + "foo bar") + v5) + v18)) + "") != "hello")}
                v21 := GenS4{F0: 2147483648, F1: (-(fi19)), F2: 4294967295}
                _ = fi19
                _ = v20
                _ = v21
            }
            if len(v3) > 0 {
                v3[0] = v18
            }
            v22 := v4[4294967295]
            _ = v22
        } else {
            if (432 == 42) {
                v18 = v18
                var v23 int32 = -2147483648
                fmt.Println(v6, v18, v5)
                _ = v23
            }
            v4[1] = v5
        }
        v24 := GenS5{F0: (false && false), F1: -9223372036854775808}
        _ = v24
    }
    v25 := GenS6{F0: -971}
    var v26 uint16 = 788
    v27 := GenS7{F0: 127, F1: 1000}
    v26 = (^(v26))
    for fi28 := 0; fi28 < 6; fi28++ {
        var v29 uint16 = 2
        var v30 uint8 = 1
        _ = fi28
        _ = v29
        _ = v30
    }
    fmt.Println(v2, v7)
    var v31 uint = 18446744073709551615
    if false {
        v7 = -2.25
        if false {
            if true {
                var v32 uint32 = 753
                if false {
                    if false {
                        var v33 bool = (!false)
                        if ((((v18 + "hello") + ("hello" + ("hello" + (v6 + (v6 + v5))))) + "hello") == (("hello" + "hello") + "hello")) {
                            var v34 bool = ((float32(((v7 + (0.0 / (v7 / v7))) + (float64((3.14159 * 0.0)) / (v7 - ((v7 * v7) + float64(0.1)))))) - float32((v7 * v7))) < 1.5)
                            var v35 int64 = 1
                            _ = v34
                            _ = v35
                        } else {
                            v32 -= (v32 >> ((v31) & 63))
                            v26 = 794
                        }
                        fmt.Println(v6, v33, v32)
                        _ = v33
                    }
                    v31 += v31
                    v36 := GenS8{F0: (v31 << (((^(v31))) & 63)), F1: 4, F2: 100}
                    _ = v36
                } else {
                    fmt.Println(v7, v26)
                    if true {
                        for fi37 := 0; fi37 < 13; fi37++ {
                            if len(v1) > 0 {
                                v1[0] = ("hello" + v5)
                            }
                            var v38 int64 = 9223372036854775807
                            _ = fi37
                            _ = v38
                        }
                        v39 := GenS7{F0: -2147483648, F1: -534}
                        v40 := v4[(v32 | (v32 >> ((65535) & 63)))]
                        _ = v39
                        _ = v40
                    }
                }
                var v41 uint8 = uint8(v32)
                _ = v32
                _ = v41
            } else {
                switch 32768 {
                case -940:
                    _ = len(v3)
                case -9223372036854775808:
                    var v42 int64 = 42
                    _ = v42
                default:
                    v7 = (v7 / ((v7 - v7) - float64(3.14159)))
                }
                var v43 int8 = -128
                _ = v43
            }
            v44 := v4[uint32(v2)]
            _ = v44
        }
        for fi45 := 0; fi45 < 5; fi45++ {
            switch 128 {
            case 368:
                for fi46 := 0; fi46 < 13; fi46++ {
                    var v47 float64 = ((float64(0.0) * float64(float32(0.1))) + float64((1e10 * 0.1)))
                    fmt.Println(v7, v26, v6)
                    _ = fi46
                    _ = v47
                }
            case -488:
                v48 := GenS2{F0: v2, F1: 42, F2: 870}
                _ = v48
            case 42:
                fmt.Println(v2, v5)
            default:
                for fi49 := 0; fi49 < 1; fi49++ {
                    v18 = ("\n\t" + v18)
                    _ = fi49
                }
            }
            _ = fi45
        }
    } else {
        switch 484 {
        case 2147483647:
            fmt.Println(v6)
        case 4294967295:
            fmt.Println(v31, v2)
        }
        for fi50 := 0; fi50 < 3; fi50++ {
            var v51 uint64 = v2
            _ = fi50
            _ = v51
        }
    }
    v52 := GenS1{F0: v31, F1: 252}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v18
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v52
    _ = v6
    _ = v7
}
