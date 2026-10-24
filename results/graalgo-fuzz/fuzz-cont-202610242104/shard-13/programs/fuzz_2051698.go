package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 bool; F1 uint64; F2 string }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int32; F1 uint8 }
type GenS5 struct { F0 int8; F1 bool; F2 int }

func main() {
    var v1 uint8 = 255
    var v2 bool = ((false || true) && (0 < -2147483648))
    v3 := make(map[string]bool)
    v4 := []int{-9223372036854775808, 0}
    var v5 uint8 = (^(v1))
    switch "\n\t" {
    case "foo bar":
        if true {
            if (400 <= 1) {
                switch -152 {
                case 238:
                    v6 := GenS1{F0: v2}
                    _ = v6
                case -1:
                    for fi7 := 0; fi7 < 11; fi7++ {
                        v8 := GenS2{F0: true, F1: 4294967295, F2: "a"}
                        fmt.Println(v2, v5)
                        _ = fi7
                        _ = v8
                    }
                case -205:
                    if (-553 == 2147483647) {
                        fmt.Println(v2, v5, v1)
                        delete(v3, "\n\t")
                        var v9 uint8 = (v1 - 2)
                        _ = v9
                    } else {
                        v10 := v3["\n\t"]
                        v4 = append(v4, -1)
                        switch "" {
                        case "":
                            if v10 {
                                v1 = 146
                            }
                        }
                        _ = v10
                    }
                }
                v11 := GenS2{F0: (true && true), F1: 18446744073709551615, F2: "hello"}
                v12 := GenS3{F0: 9223372036854775807}
                _ = v11
                _ = v12
            }
            v4 = append(v4, 400)
            v13 := GenS4{F0: 2147483647, F1: 2}
            _ = v13
        } else {
            fmt.Println(v1, v5, v5)
        }
    case "\n\t":
        _ = len(v4)
    case "":
        var v14 uint = 4294967295
        _ = v14
    default:
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    var v15 uint64 = 2
    fmt.Println(v1)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    switch 127 {
    case -2147483648:
        var v16 string = "hello"
        _ = v16
    case 42:
        for fi17 := 0; fi17 < 7; fi17++ {
            var v18 int16 = -32768
            var v19 uint8 = (v1 ^ v1)
            _ = fi17
            _ = v18
            _ = v19
        }
    default:
        var v20 string = "foo bar"
        _ = v20
    }
    _ = len(v4)
    for fi21 := 0; fi21 < 19; fi21++ {
        var v22 uint16 = 0
        v23 := v3[""]
        switch fi21 {
        case 682:
            v22 = v22
        case 1:
            fmt.Println(v5)
        default:
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        }
        _ = fi21
        _ = v22
        _ = v23
    }
    for fi24 := 0; fi24 < 12; fi24++ {
        fmt.Println(v5)
        v25 := GenS3{F0: 42}
        _ = fi24
        _ = v25
    }
    var v26 float64 = (1e10 + 3.14159)
    for fi27 := 0; fi27 < 11; fi27++ {
        for fi28 := 0; fi28 < 9; fi28++ {
            var v29 float32 = -2.25
            _ = fi28
            _ = v29
        }
        _ = fi27
    }
    v30 := GenS5{F0: 100, F1: true, F2: 917}
    if (-455 <= 42) {
        for fi31 := 0; fi31 < 17; fi31++ {
            var v32 int32 = 1
            fmt.Println(v1)
            _ = fi31
            _ = v32
        }
        if (!true) {
            if (2 >= 18446744073709551615) {
                for fi33 := 0; fi33 < 15; fi33++ {
                    if ("" <= "\n\t") {
                        v26 += v26
                        var v34 int8 = 127
                        if (759 >= 0) {
                            var v35 int = fi33
                            var v36 uint32 = 1
                            _ = v35
                            _ = v36
                        }
                        _ = v34
                    }
                    v4 = append(v4, fi33)
                    var v37 bool = v2
                    _ = fi33
                    _ = v37
                }
                v15 -= uint64(v1)
            }
            _ = len(v3)
            var v38 int = 42
            _ = v38
        }
    }
    if ((65535 != 663) || (!false)) {
        if (390 >= 255) {
            switch 1 {
            case 1:
                v2 = v2
            case -9223372036854775808:
                v1 = 127
            default:
                if v2 {
                    for fi39 := 0; fi39 < 2; fi39++ {
                        if (-9223372036854775808 < (fi39 - -33)) {
                            if len(v4) > 0 {
                                v4[0] = fi39
                            }
                            switch ((((("a" + "foo bar") + "") + (("hello" + "a") + ("hello" + "foo bar"))) + "\n\t") + "") {
                            case "\n\t":
                                if (!false) {
                                    for fi40 := 0; fi40 < 12; fi40++ {
                                        v26 -= (float64(1e10) - ((-2.25 / v26) * (v26 * v26)))
                                        v26 = v26
                                        _ = fi40
                                    }
                                    fmt.Println(fi39)
                                    v41 := GenS5{F0: -128, F1: v2, F2: (fi39 & fi39)}
                                    _ = v41
                                } else {
                                    v42 := v3["\n\t"]
                                    var v43 uint16 = 32767
                                    _ = v42
                                    _ = v43
                                }
                            case "hello":
                                for fi44 := 0; fi44 < 4; fi44++ {
                                    if true {
                                        v1 += v5
                                        _ = len(v3)
                                        fmt.Println(v5, v26, v15)
                                    } else {
                                        fmt.Println(v2, fi44)
                                        var v45 int16 = 0
                                        if v2 {
                                            v46 := GenS3{F0: 1}
                                            _ = len(v4)
                                            v5 += 42
                                            _ = v46
                                        }
                                        _ = v45
                                    }
                                    v1 += (^(v5))
                                    var v47 string = ((("hello" + "a") + "hello") + ((("foo bar" + ("" + "foo bar")) + (("\n\t" + "\n\t") + "hello")) + ("a" + ("\n\t" + "\n\t"))))
                                    switch fi39 {
                                    case -2147483648:
                                        var v48 uint32 = 310
                                        _ = v48
                                    case 104:
                                        if v2 {
                                            var v49 float64 = (v26 / ((-2.25 * ((v26 + v26) - 0.1)) * v26))
                                            _ = v49
                                        }
                                    default:
                                        v50 := v3["hello"]
                                        _ = v50
                                    }
                                    _ = fi44
                                    _ = v47
                                }
                            }
                            var v51 float64 = -2.25
                            _ = v51
                        } else {
                            for fi52 := 0; fi52 < 16; fi52++ {
                                var v53 uint64 = v15
                                _ = fi52
                                _ = v53
                            }
                            v1 -= (v5 ^ (^(v1)))
                            v5 += 255
                        }
                        _ = fi39
                    }
                }
            }
            v54 := GenS3{F0: 1}
            _ = v54
        }
    } else {
        var v55 int16 = -699
        _ = len(v3)
        _ = v55
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v2
    _ = v26
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
