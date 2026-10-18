package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16 }
type GenS2 struct { F0 bool; F1 uint; F2 uint }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int32; F1 uint32; F2 int64 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 string; F1 uint16; F2 uint16 }
type GenS7 struct { F0 uint16; F1 uint }
type GenS8 struct { F0 int16; F1 int32; F2 uint64 }
type GenS9 struct { F0 uint32 }

func main() {
    v1 := []int{42, 443}
    var v2 int64 = -9223372036854775808
    v3 := make(map[uint32]bool)
    var v4 int8 = -128
    v5 := GenS1{F0: 128, F1: -310}
    var v6 uint8 = 0
    fmt.Println(v4)
    v3[65535] = false
    v4 -= 0
    if ("a" <= "foo bar") {
        if true {
            var v7 int16 = 42
            var v8 int8 = (v4 & (^(v4)))
            for fi9 := 0; fi9 < 16; fi9++ {
                v8 += -29
                if true {
                    v2 += v2
                    v10 := GenS1{F0: 4294967295, F1: (v7 ^ int16(v2))}
                    _ = v10
                } else {
                    switch ("foo bar" + "foo bar") {
                    case "\n\t":
                        for fi11 := 0; fi11 < 16; fi11++ {
                            switch "foo bar" {
                            case "a":
                                v12 := v3[2]
                                _ = v12
                            }
                            v13 := GenS1{F0: (fi9 / ((fi11) | 1)), F1: (v7 + -1)}
                            v8 = (v4 + (v4 << ((42) & 63)))
                            _ = fi11
                            _ = v13
                        }
                    }
                    var v14 uint8 = v6
                    var v15 uint64 = 1
                    _ = v14
                    _ = v15
                }
                _ = fi9
            }
            _ = v7
            _ = v8
        } else {
            if (true && (1000 < -32768)) {
                if true {
                    for fi16 := 0; fi16 < 1; fi16++ {
                        v17 := GenS1{F0: fi16, F1: 32767}
                        _ = fi16
                        _ = v17
                    }
                    if false {
                        switch "\n\t" {
                        case "\n\t":
                            for fi18 := 0; fi18 < 5; fi18++ {
                                if (-1 < -485) {
                                    v19 := GenS1{F0: (fi18 * (fi18 % ((fi18) | 1))), F1: 0}
                                    for fi20 := 0; fi20 < 8; fi20++ {
                                        v6 = (^(v6))
                                        var v21 uint64 = 18446744073709551615
                                        _ = fi20
                                        _ = v21
                                    }
                                    switch "" {
                                    case "\n\t":
                                        switch fi18 {
                                        case 0:
                                            v6 += v6
                                        case -205:
                                            v4 = (v4 * v4)
                                        case -468:
                                            v1 = append(v1, (fi18 * -9223372036854775808))
                                        }
                                    case "foo bar":
                                        v22 := GenS1{F0: fi18, F1: -810}
                                        _ = v22
                                    case "hello":
                                        var v23 bool = ((v6 ^ (v6 - v6)) >= (v6 << ((1) & 63)))
                                        _ = v23
                                    default:
                                        v24 := GenS2{F0: (85 != 4294967295), F1: 42, F2: 18446744073709551615}
                                        _ = v24
                                    }
                                    _ = v19
                                }
                                fmt.Println(v4, v2)
                                if (1.5 == 0.0) {
                                    if ((!false) && ((^(v6)) < v6)) {
                                        switch (fi18 | (fi18 | fi18)) {
                                        case 127:
                                            if (!false) {
                                                if false {
                                                    for fi25 := 0; fi25 < 17; fi25++ {
                                                        v26 := GenS3{F0: (fi25 - 722)}
                                                        var v27 int = fi25
                                                        var v28 uint8 = v6
                                                        _ = fi25
                                                        _ = v26
                                                        _ = v27
                                                        _ = v28
                                                    }
                                                    if true {
                                                        v29 := GenS3{F0: (-(fi18))}
                                                        var v30 int32 = 27
                                                        _ = v29
                                                        _ = v30
                                                    }
                                                }
                                            }
                                        }
                                    }
                                }
                                _ = fi18
                            }
                        case "":
                            v31 := GenS4{F0: 122, F1: 2147483647, F2: (v2 << ((uint(v6)) & 63))}
                            _ = v31
                        }
                        v32 := v3[42]
                        _ = v32
                    }
                }
            }
            v33 := GenS2{F0: (true && true), F1: 0, F2: 42}
            fmt.Println(v4)
            _ = v33
        }
        if (!false) {
            var v34 int = -9223372036854775808
            _ = v34
        }
    } else {
        v2 = -468
        var v35 int8 = -128
        switch 2147483647 {
        case 829:
            for fi36 := 0; fi36 < 7; fi36++ {
                v3[65535] = (1 != 100)
                v37 := GenS5{F0: v6}
                fmt.Println(fi36, v6)
                _ = fi36
                _ = v37
            }
        case 299:
            switch -152 {
            case -2147483648:
                v38 := GenS6{F0: "\n\t", F1: 255, F2: 18}
                _ = v38
            }
        default:
            v39 := GenS7{F0: 42, F1: 1}
            _ = v39
        }
        _ = v35
    }
    if true {
        for fi40 := 0; fi40 < 13; fi40++ {
            for fi41 := 0; fi41 < 2; fi41++ {
                v42 := v3[65535]
                _ = fi41
                _ = v42
            }
            var v43 float64 = ((1.5 - 0.1) * (0.1 * 1.5))
            _ = fi40
            _ = v43
        }
    } else {
        for fi44 := 0; fi44 < 13; fi44++ {
            switch "foo bar" {
            case "\n\t":
                v45 := GenS6{F0: "a", F1: 0, F2: 2}
                _ = v45
            case "a":
                if (false || (!false)) {
                    switch fi44 {
                    case 1:
                        var v46 uint = 4294967295
                        _ = v46
                    default:
                        fmt.Println(fi44, v6)
                    }
                }
            case "foo bar":
                v47 := GenS8{F0: 1, F1: -2147483648, F2: 926}
                _ = v47
            default:
                var v48 float64 = 3.14159
                _ = v48
            }
            _ = fi44
        }
    }
    _ = len(v3)
    v4 = 100
    switch "a" {
    case "a":
        var v49 bool = ((3.14159 != (-2.25 + 3.14159)) && true)
        _ = v49
    case "hello":
        _ = len(v1)
    case "\n\t":
        fmt.Println(v6, v4, v2)
    }
    fmt.Println(v4)
    v2 = (v2 >> ((236) & 63))
    v50 := GenS9{F0: 2147483648}
    var v51 int = 127
    if (-2.25 >= float32(((0.0 * 1.5) / 1.5))) {
        v52 := GenS8{F0: int16(v6), F1: 2147483647, F2: 1}
        for fi53 := 0; fi53 < 10; fi53++ {
            v1 = append(v1, v51)
            _ = fi53
        }
        _ = v52
    }
    if true {
        fmt.Println(v51, v51, v2)
        var v54 uint8 = (v6 + (v6 | (v6 / ((1) | 1))))
        _ = v54
    } else {
        var v55 uint64 = 195
        v51 -= (v51 & (v51 - int(v6)))
        _ = v55
    }
    v3[42] = (3.14159 != 0.0)
    var v56 uint8 = v6
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v51:", v51)
    fmt.Println("v56:", v56)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v51
    _ = v56
    _ = v6
}
