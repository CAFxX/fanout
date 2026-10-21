package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint8; F1 uint64; F2 uint8 }
type GenS3 struct { F0 string; F1 string; F2 int8 }
type GenS4 struct { F0 uint8; F1 uint32 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 uint32; F1 uint }
type GenS7 struct { F0 uint; F1 uint32 }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 string }
func genFn0() int {
    var v4 uint = 776
    if true {
        if (false || ((!false) && false)) {
            for fi5 := 0; fi5 < 20; fi5++ {
                v6 := []int{fi5, 128}
                v7 := GenS1{F0: 2}
                _ = fi5
                _ = v6
                _ = v7
            }
            if true {
                v4 = v4
                v8 := GenS1{F0: 243}
                _ = v8
            }
            v4 = v4
        } else {
            v9 := []int{-9223372036854775808, 2147483647, -1}
            _ = v9
        }
        var v10 int16 = 399
        switch "a" {
        case "\n\t":
            for fi11 := 0; fi11 < 18; fi11++ {
                v12 := []int{-579, (fi11 + 42), (fi11 >> ((v4) & 63))}
                v13 := GenS2{F0: 124, F1: 42, F2: 0}
                if (float64(((0.0 * 0.1) * 3.14159)) <= (1.5 - 3.14159)) {
                    if (!((fi11 - fi11) > 9223372036854775807)) {
                        v14 := GenS3{F0: "foo bar", F1: "\n\t", F2: -98}
                        v15 := GenS2{F0: 240, F1: 2, F2: 1}
                        _ = v14
                        _ = v15
                    }
                } else {
                    if false {
                        var v16 float64 = 1e10
                        var v17 uint32 = 2147483647
                        v18 := GenS2{F0: 1, F1: 695, F2: 127}
                        _ = v16
                        _ = v17
                        _ = v18
                    } else {
                        switch "a" {
                        case "hello":
                            if len(v12) > 0 {
                                v12[0] = -2147483648
                            }
                        default:
                            v19 := GenS4{F0: 151, F1: 570}
                            _ = v19
                        }
                        switch "" {
                        case "hello":
                            fmt.Println(v10, fi11)
                        case "foo bar":
                            if true {
                                var v20 int32 = -1
                                v4 += v4
                                var v21 uint = (v4 >> (((v4 % ((v4) | 1))) & 63))
                                _ = v20
                                _ = v21
                            } else {
                                var v22 uint32 = 116
                                switch (fi11 & 0) {
                                case 277:
                                    fmt.Println(v4)
                                }
                                _ = v22
                            }
                        default:
                            if ((fi11 + (fi11 | -2147483648)) == -507) {
                                var v23 int8 = 1
                                _ = v23
                            }
                        }
                        var v24 int16 = 933
                        _ = v24
                    }
                    v25 := GenS5{F0: 917}
                    v4 += 0
                    _ = v25
                }
                _ = fi11
                _ = v12
                _ = v13
            }
        case "foo bar":
            for fi26 := 0; fi26 < 2; fi26++ {
                fmt.Println(fi26)
                _ = fi26
            }
        case "":
            for fi27 := 0; fi27 < 9; fi27++ {
                v10 += -163
                v28 := GenS5{F0: 1}
                _ = fi27
                _ = v28
            }
        }
        _ = v10
    } else {
        for fi29 := 0; fi29 < 4; fi29++ {
            v30 := GenS1{F0: 127}
            _ = fi29
            _ = v30
        }
        v31 := []int{}
        switch ("" + "") {
        case "":
            v4 = v4
        }
        _ = v31
    }
    _ = v4
    return -2147483648
}
func genFn1(p0 int64, p1 uint64) string {
    v32 := []int{2147483647, 32767}
    v33 := GenS4{F0: 172, F1: 2147483647}
    _ = p0
    _ = p1
    _ = v32
    _ = v33
    return "hello"
}

func main() {
    var v1 string = "hello"
    var v2 int64 = 1
    var v3 uint32 = 756
    v34 := []int{1, 2147483647}
    if true {
        v35 := GenS6{F0: v3, F1: 303}
        for fi36 := 0; fi36 < 4; fi36++ {
            var v37 int = -9223372036854775808
            for fi38 := 0; fi38 < 10; fi38++ {
                if true {
                    v1 = ((v1 + (((("foo bar" + (v1 + v1)) + v1) + (v1 + "foo bar")) + "\n\t")) + ("" + "hello"))
                }
                var v39 string = v1
                _ = fi38
                _ = v39
            }
            _ = fi36
            _ = v37
        }
        if false {
            v1 = "hello"
            for fi40 := 0; fi40 < 12; fi40++ {
                var v41 uint32 = v3
                if false {
                    var v42 uint8 = 127
                    _ = v42
                }
                var v43 uint32 = (v3 << ((301) & 63))
                _ = fi40
                _ = v41
                _ = v43
            }
            v44 := GenS7{F0: 0, F1: 2}
            _ = v44
        } else {
            v1 = v1
            v2 += v2
        }
        _ = v35
    } else {
        var v45 int = 947
        v45 = (v45 ^ 4294967295)
        _ = v45
    }
    var v46 uint8 = 127
    var v47 int32 = -102
    v48 := GenS6{F0: 42, F1: 443}
    var v49 int64 = v2
    for fi50 := 0; fi50 < 9; fi50++ {
        switch v1 {
        case "\n\t":
            _ = len(v34)
        }
        var v51 float64 = 1.5
        if true {
            var v52 uint8 = (v46 ^ (v46 << ((65535) & 63)))
            for fi53 := 0; fi53 < 5; fi53++ {
                if (2 == 18446744073709551615) {
                    var v54 int16 = -32768
                    _ = v54
                } else {
                    var v55 uint16 = 42
                    var v56 int = 788
                    _ = v55
                    _ = v56
                }
                _ = fi53
            }
            _ = v52
        }
        _ = fi50
        _ = v51
    }
    var v57 uint32 = (v3 / ((39) | 1))
    if (0.0 <= ((-2.25 * ((1.5 * 1.5) - 0.1)) * 1.5)) {
        if len(v34) > 0 {
            v34[0] = 1
        }
        v2 -= v49
    }
    fmt.Println(v57, v57)
    for fi58 := 0; fi58 < 1; fi58++ {
        v34 = append(v34, fi58)
        for fi59 := 0; fi59 < 10; fi59++ {
            v60 := GenS4{F0: v46, F1: v57}
            v46 -= 127
            _ = fi59
            _ = v60
        }
        for fi61 := 0; fi61 < 3; fi61++ {
            v62 := GenS8{F0: 1}
            if false {
                v34 = append(v34, int(v2))
                fmt.Println(v1, v1)
            } else {
                fmt.Println(v47, fi61, v1)
                if (!true) {
                    if len(v34) >= 2 {
                        v34 = v34[1:len(v34)]
                    }
                    var v63 int16 = 1
                    _ = len(v34)
                    if true {
                        var v64 uint8 = v46
                        _ = v64
                    } else {
                        fmt.Println(v47, v49)
                        var v65 uint = 42
                        fmt.Println(v49)
                        _ = v65
                    }
                    v57 += 2
                    v66 := GenS7{F0: 53, F1: v3}
                    _ = v63
                    _ = v66
                }
            }
            _ = fi61
            _ = v62
        }
        _ = fi58
    }
    v67 := GenS9{F0: (v1 + ("foo bar" + v1))}
    switch 127 {
    case 32767:
        switch v1 {
        case "\n\t":
            if len(v34) >= 2 {
                v34 = v34[1:len(v34)]
            }
        }
    case 0:
        if len(v34) >= 2 {
            v34 = v34[1:len(v34)]
        }
    case -9223372036854775808:
        var v68 uint32 = v57
        _ = v68
    default:
        v69 := GenS9{F0: v1}
        _ = v69
    }
    v49 += v2
    v2 = -1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v49:", v49)
    fmt.Println("v57:", v57)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v46
    _ = v47
    _ = v48
    _ = v49
    _ = v57
    _ = v67
}
