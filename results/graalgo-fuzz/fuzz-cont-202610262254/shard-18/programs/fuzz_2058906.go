package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool }
type GenS2 struct { F0 string }
type GenS3 struct { F0 string; F1 uint64 }
type GenS4 struct { F0 int64; F1 string; F2 uint32 }
type GenS5 struct { F0 uint8; F1 int16 }
type GenS6 struct { F0 uint8; F1 uint64 }
type GenS7 struct { F0 bool; F1 uint }
type GenS8 struct { F0 int8; F1 int16; F2 uint8 }
type GenS9 struct { F0 int }

func main() {
    var v1 uint8 = 0
    v2 := make(map[int]int)
    var v3 int = -9223372036854775808
    var v4 float32 = (1e10 * float32(float64(3.14159)))
    var v5 int32 = 599
    var v6 int = v3
    v1 = 128
    v4 -= ((v4 / 1.5) + ((-2.25 * v4) - (v4 / float32(-2.25))))
    for fi7 := 0; fi7 < 6; fi7++ {
        v8 := GenS1{F0: (false || true), F1: false}
        switch "a" {
        case "hello":
            v9 := GenS2{F0: "\n\t"}
            _ = v9
        default:
            var v10 bool = true
            _ = v10
        }
        if (!(!(fi7 < fi7))) {
            if true {
                for fi11 := 0; fi11 < 1; fi11++ {
                    switch "a" {
                    case "hello":
                        if false {
                            v2[506] = (v6 & v6)
                            v12 := []int{(v6 * (v3 ^ (fi11 * fi7))), fi7, (fi11 / ((9223372036854775807) | 1)), (v6 % ((fi7) | 1))}
                            _ = v12
                        } else {
                            if false {
                                if (!false) {
                                    var v13 uint32 = 1
                                    if (float32((((3.14159 / -2.25) / 1.0) + -2.25)) > v4) {
                                        if false {
                                            v14 := GenS1{F0: false, F1: (v5 != v5)}
                                            var v15 uint32 = (v13 | v13)
                                            v16 := GenS3{F0: "foo bar", F1: 0}
                                            _ = v14
                                            _ = v15
                                            _ = v16
                                        }
                                        v17 := GenS4{F0: 1, F1: "hello", F2: 65535}
                                        v5 += (^(v5))
                                        var v18 float64 = 3.14159
                                        v19 := []int{fi7, 32767, (fi7 ^ (v3 + 42))}
                                        var v20 bool = (0 > 42)
                                        _ = v17
                                        _ = v18
                                        _ = v19
                                        _ = v20
                                    } else {
                                        v21 := GenS5{F0: 127, F1: 1000}
                                        _ = v21
                                    }
                                    _ = v13
                                }
                                if (902 <= 4294967295) {
                                    if (1e10 < ((v4 / 1.0) * (v4 + (1.5 + v4)))) {
                                        v22 := GenS5{F0: v1, F1: -1}
                                        var v23 int8 = 127
                                        fmt.Println(v3, v23, fi11)
                                        _ = v22
                                        _ = v23
                                    } else {
                                        var v24 int64 = int64(v1)
                                        fmt.Println(fi7, v24, fi7)
                                        var v25 float32 = (v4 + 3.14159)
                                        _ = v24
                                        _ = v25
                                    }
                                    switch "a" {
                                    case "\n\t":
                                        v1 = (v1 / ((v1) | 1))
                                    case "hello":
                                        var v26 uint8 = 2
                                        _ = v26
                                    case "a":
                                        v27 := GenS6{F0: v1, F1: 934}
                                        _ = v27
                                    default:
                                        if ((false && ((0.1 * 0.1) < (1e10 + (float64(-2.25) + (0.0 - 3.14159))))) || true) {
                                            if (25 <= 971) {
                                                v6 = (fi11 | v6)
                                                v28 := GenS4{F0: 9223372036854775807, F1: "hello", F2: uint32(v1)}
                                                v29 := GenS5{F0: (^(v1)), F1: 42}
                                                _ = v28
                                                _ = v29
                                            }
                                            if false {
                                                var v30 int8 = -105
                                                _ = v30
                                            }
                                            v3 = (fi7 ^ v3)
                                        }
                                    }
                                    v3 -= 32768
                                }
                                v5 += (^(v5))
                            } else {
                                if true {
                                    var v31 int64 = -9223372036854775808
                                    if (true && true) {
                                        v31 = (v31 | v31)
                                        var v32 int32 = 0
                                        var v33 int = (fi11 + fi7)
                                        v32 += (v32 + (v5 & -1))
                                        delete(v2, (^(v33)))
                                        _ = v32
                                        _ = v33
                                    }
                                    _ = v31
                                }
                                v34 := GenS6{F0: 172, F1: 2}
                                v35 := []int{fi11, (fi11 + v6)}
                                _ = v34
                                _ = v35
                            }
                            v6 -= (v6 & fi7)
                        }
                    }
                    var v36 int32 = 2147483647
                    v36 = 0
                    var v37 int16 = -1
                    _ = fi11
                    _ = v36
                    _ = v37
                }
            }
            v38 := []int{(fi7 & -717), v6, fi7}
            v1 -= (v1 ^ 127)
            _ = v38
        } else {
            if (!((574 != 32767) && true)) {
                var v39 float64 = float64(v4)
                _ = v39
            }
        }
        _ = fi7
        _ = v8
    }
    v6 -= v6
    var v40 uint32 = 2147483647
    delete(v2, -2147483648)
    switch v6 {
    case -2147483648:
        if ("a" < ("" + (("" + "a") + "\n\t"))) {
            var v41 int64 = 580
            _ = v41
        }
    }
    if true {
        for fi42 := 0; fi42 < 3; fi42++ {
            v5 = -1
            switch fi42 {
            case -2147483648:
                if (false || false) {
                    v3 = v3
                    v6 = v6
                }
            case 2147483647:
                v2[(fi42 | v6)] = (v3 ^ 71)
            case -512:
                var v43 int16 = -32768
                _ = v43
            default:
                if (-32768 <= -32768) {
                    var v44 uint16 = 149
                    v3 = v3
                    _ = v44
                }
            }
            _ = fi42
        }
        if (594 != (v40 % ((2147483647) | 1))) {
            v45 := GenS6{F0: v1, F1: 1}
            v46 := []int{v6, 2147483647, -2147483648, v6}
            _ = v45
            _ = v46
        } else {
            if (255 == 42) {
                v47 := []int{0, v3}
                v48 := v2[128]
                _ = v47
                _ = v48
            }
            v5 -= (v5 + (v5 << ((243) & 63)))
        }
        if (18446744073709551615 == 337) {
            v40 = v40
            v3 = 32768
        } else {
            v5 -= -524
        }
    }
    for fi49 := 0; fi49 < 10; fi49++ {
        v50 := GenS7{F0: (42 > 100), F1: 288}
        for fi51 := 0; fi51 < 9; fi51++ {
            v40 -= 77
            v52 := GenS2{F0: ""}
            _ = fi51
            _ = v52
        }
        _ = fi49
        _ = v50
    }
    switch 0 {
    case -1:
        var v53 uint16 = 65535
        _ = v53
    case 243:
        if false {
            v54 := v2[v6]
            _ = v54
        }
    case 4294967295:
        v55 := GenS8{F0: -58, F1: -1, F2: (v1 << ((4294967295) & 63))}
        _ = v55
    }
    if (!((!true) || false)) {
        if ("a" > ("foo bar" + "")) {
            var v56 string = "\n\t"
            if (0 > (v40 << ((718) & 63))) {
                v56 = v56
            }
            _ = v56
        }
        v57 := GenS9{F0: v6}
        _ = v57
    }
    delete(v2, (v3 * v6))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
