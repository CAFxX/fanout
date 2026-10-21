package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int; F2 uint16 }
type GenS2 struct { F0 int64; F1 uint32 }
type GenS3 struct { F0 string; F1 uint }
type GenS4 struct { F0 int16; F1 string; F2 bool }
type GenS5 struct { F0 uint64; F1 uint; F2 uint64 }

func main() {
    v1 := []string{"a", "a", ""}
    var v2 int = 128
    var v3 float64 = float64(1.5)
    var v4 float64 = v3
    var v5 int8 = 42
    var v6 int64 = -9223372036854775808
    for fi7 := 0; fi7 < 18; fi7++ {
        var v8 string = "hello"
        _ = fi7
        _ = v8
    }
    v9 := GenS1{F0: -1, F1: (v2 | (v2 * v2)), F2: 255}
    var v10 int = (v2 & v2)
    for fi11 := 0; fi11 < 2; fi11++ {
        var v12 string = (("\n\t" + "foo bar") + "foo bar")
        if len(v1) > 0 {
            v1[0] = ""
        }
        _ = fi11
        _ = v12
    }
    var v13 string = ""
    if (v2 == (v2 - (v10 >> ((2) & 63)))) {
        switch v2 {
        case 2147483647:
            _ = len(v1)
        case 42:
            fmt.Println(v6)
        case 0:
            var v14 int8 = v5
            _ = v14
        }
    } else {
        _ = len(v1)
    }
    if false {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        var v15 float32 = (0.1 / 0.1)
        if (1 >= 1) {
            var v16 int64 = (v6 / ((-9223372036854775808) | 1))
            switch v13 {
            case "hello":
                for fi17 := 0; fi17 < 7; fi17++ {
                    if (float64(v15) < 1.5) {
                        v18 := GenS2{F0: v6, F1: 1}
                        var v19 uint16 = 0
                        var v20 int16 = 1
                        _ = v18
                        _ = v19
                        _ = v20
                    } else {
                        fmt.Println(v16, fi17)
                    }
                    if true {
                        if (655 < -860) {
                            fmt.Println(fi17, v5)
                            var v21 uint8 = 1
                            _ = v21
                        }
                        if true {
                            if (true || (!true)) {
                                var v22 int = fi17
                                if (!true) {
                                    v3 = v4
                                    v22 += fi17
                                    v5 = (^(v5))
                                } else {
                                    switch -1 {
                                    case 0:
                                        v23 := GenS2{F0: v6, F1: 2147483647}
                                        _ = v23
                                    case -148:
                                        if false {
                                            fmt.Println(v10)
                                            var v24 uint32 = 1
                                            _ = v24
                                        }
                                    default:
                                        fmt.Println(fi17, v3, v16)
                                    }
                                }
                                v25 := GenS2{F0: 1, F1: 65535}
                                _ = v22
                                _ = v25
                            } else {
                                fmt.Println(v16, v16, fi17)
                                v4 += -2.25
                                switch "\n\t" {
                                case "hello":
                                    var v26 bool = (!(!false))
                                    _ = v26
                                case "":
                                    switch v2 {
                                    case 9223372036854775807:
                                        if len(v1) > 0 {
                                            v1[0] = ((v13 + v13) + (("\n\t" + "foo bar") + v13))
                                        }
                                    default:
                                        var v27 uint16 = 724
                                        _ = v27
                                    }
                                }
                            }
                            if true {
                                var v28 bool = true
                                _ = v28
                            } else {
                                for fi29 := 0; fi29 < 6; fi29++ {
                                    v13 = (v13 + v13)
                                    _ = fi29
                                }
                            }
                        } else {
                            v4 = v4
                            for fi30 := 0; fi30 < 13; fi30++ {
                                fmt.Println(v4)
                                _ = fi30
                            }
                            v31 := GenS3{F0: v13, F1: 973}
                            _ = v31
                        }
                        switch (v2 + fi17) {
                        case -1:
                            switch 1 {
                            case 9223372036854775807:
                                for fi32 := 0; fi32 < 8; fi32++ {
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    _ = fi32
                                }
                            case -1:
                                if ((int16(v6) == 118) || (42 <= 578)) {
                                    v2 += (fi17 & fi17)
                                    v33 := GenS4{F0: -1, F1: v13, F2: (!(!false))}
                                    v3 += v4
                                    _ = v33
                                } else {
                                    var v34 string = ("foo bar" + "foo bar")
                                    _ = v34
                                }
                            case 938:
                                if true {
                                    if true {
                                        if len(v1) > 0 {
                                            v1[0] = v13
                                        }
                                    } else {
                                        var v35 uint16 = 255
                                        _ = v35
                                    }
                                }
                            }
                        case -9223372036854775808:
                            var v36 int32 = 42
                            _ = v36
                        case 34:
                            var v37 float32 = (1.5 * 1e10)
                            _ = v37
                        default:
                            var v38 float64 = v3
                            _ = v38
                        }
                    }
                    v10 -= int(v5)
                    _ = fi17
                }
            }
            _ = v16
        }
        _ = v15
    }
    for fi39 := 0; fi39 < 18; fi39++ {
        var v40 float64 = ((v3 * v3) / (((1e10 / v3) / v4) - -2.25))
        var v41 string = v13
        v5 = -1
        _ = fi39
        _ = v40
        _ = v41
    }
    v10 = (v2 >> ((184) & 63))
    var v42 int8 = (v5 % ((v5) | 1))
    _ = len(v1)
    var v43 string = ("foo bar" + v13)
    for fi44 := 0; fi44 < 13; fi44++ {
        var v45 int64 = v6
        _ = fi44
        _ = v45
    }
    v42 = (v42 + v5)
    switch v43 {
    case "\n\t":
        fmt.Println(v5, v43, v4)
    case "":
        switch "" {
        case "":
            fmt.Println(v13)
        case "hello":
            var v46 int32 = 2147483647
            _ = v46
        default:
            var v47 bool = true
            _ = v47
        }
    }
    v6 += v6
    var v48 bool = true
    if (-1 > (^(v6))) {
        var v49 uint64 = 42
        v3 -= float64(0.0)
        _ = v49
    } else {
        var v50 uint16 = 255
        v51 := GenS5{F0: 4294967295, F1: 18446744073709551615, F2: 0}
        _ = v50
        _ = v51
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v48
    _ = v5
    _ = v6
    _ = v9
}
