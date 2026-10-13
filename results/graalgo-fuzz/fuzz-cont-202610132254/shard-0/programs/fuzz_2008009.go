package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint; F2 uint64 }
type GenS2 struct { F0 uint64; F1 uint64 }
type GenS3 struct { F0 uint; F1 int; F2 bool }
type GenS4 struct { F0 int8; F1 string }
func genFn0(p0 uint8) bool {
    var v7 int32 = 0
    if ((!true) && false) {
        var v8 float64 = 1.5
        fmt.Println(v7)
        p0 += p0
        _ = v8
    } else {
        var v9 int16 = 0
        var v10 int32 = v7
        _ = v10
        _ = v9
    }
    p0 += p0
    v7 -= -305
    _ = p0
    _ = v7
    return false
}

func main() {
    v1 := []string{("\n\t" + "hello")}
    var v2 int16 = -735
    var v3 int16 = v2
    var v4 int32 = -2147483648
    v5 := []int{1, 127, 4294967295, -2147483648}
    var v6 int64 = 42
    var v11 uint = 18446744073709551615
    var v12 float64 = float64((float32(1e10) + -2.25))
    var v13 int32 = -510
    if false {
        if (true || false) {
            _ = len(v5)
            var v14 uint32 = 2147483648
            v4 = 826
            _ = v14
        }
        var v15 string = "\n\t"
        var v16 int16 = (v3 & int16(v6))
        _ = v15
        _ = v16
    } else {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        for fi17 := 0; fi17 < 13; fi17++ {
            v5 = append(v5, fi17)
            switch "foo bar" {
            case "\n\t":
                for fi18 := 0; fi18 < 3; fi18++ {
                    fmt.Println(v13, fi17)
                    v19 := GenS1{F0: 0, F1: v11, F2: 2}
                    v4 = (v13 * 42)
                    _ = fi18
                    _ = v19
                }
            case "foo bar":
                v20 := GenS1{F0: 2, F1: (v11 * v11), F2: 4294967295}
                _ = v20
            }
            if true {
                v21 := GenS1{F0: (v11 % ((v11) | 1)), F1: v11, F2: 322}
                _ = v21
            }
            _ = fi17
        }
        v3 = (v3 % (((v3 * v3)) | 1))
    }
    fmt.Println(v2)
    switch "hello" {
    case "a":
        v22 := GenS1{F0: v11, F1: (v11 * v11), F2: 1}
        _ = v22
    case "foo bar":
        v12 += v12
    case "\n\t":
        fmt.Println(v4, v3)
    }
    v4 = v13
    v4 += v4
    if (!((v13 > (v4 % ((v4) | 1))) && ((v11 % (((v11 * (v11 >> ((1) & 63)))) | 1)) < (v11 / (((v11 - 187)) | 1))))) {
        if ((v4 | 2147483647) != (v4 << (((v11 ^ (v11 - v11))) & 63))) {
            if (!false) {
                switch "" {
                case "hello":
                    fmt.Println(v4, v2)
                case "a":
                    v1 = append(v1, "foo bar")
                case "\n\t":
                    fmt.Println(v12)
                }
                switch 430 {
                case 1:
                    var v23 int = -2147483648
                    _ = v23
                case 127:
                    v12 += 1e10
                case -9223372036854775808:
                    v2 = (-(v3))
                default:
                    switch -757 {
                    case -1:
                        switch "foo bar" {
                        case "":
                            var v24 int8 = 0
                            _ = v24
                        case "\n\t":
                            if len(v5) > 0 {
                                v5[0] = 673
                            }
                        case "hello":
                            if ((v3 | -563) > 986) {
                                switch "" {
                                case "\n\t":
                                    v12 = v12
                                case "a":
                                    v25 := GenS1{F0: 255, F1: (^(v11)), F2: 42}
                                    _ = v25
                                case "hello":
                                    v26 := GenS1{F0: 18446744073709551615, F1: v11, F2: 42}
                                    _ = v26
                                }
                                for fi27 := 0; fi27 < 4; fi27++ {
                                    v2 -= 1000
                                    v28 := GenS2{F0: 267, F1: 42}
                                    if len(v5) > 0 {
                                        v5[0] = (^(fi27))
                                    }
                                    _ = fi27
                                    _ = v28
                                }
                                if (false || true) {
                                    var v29 uint64 = 42
                                    v30 := GenS2{F0: v29, F1: v29}
                                    switch "" {
                                    case "foo bar":
                                        v6 = v6
                                    case "a":
                                        v29 = v29
                                    case "":
                                        v29 = (v29 % ((0) | 1))
                                    }
                                    _ = v29
                                    _ = v30
                                }
                            } else {
                                v31 := GenS1{F0: v11, F1: v11, F2: uint64(v11)}
                                for fi32 := 0; fi32 < 15; fi32++ {
                                    v3 -= (v3 << ((v11) & 63))
                                    v12 = float64(3.14159)
                                    v33 := GenS2{F0: 42, F1: 1}
                                    _ = fi32
                                    _ = v33
                                }
                                v13 += v4
                                _ = v31
                            }
                        default:
                            if (true && true) {
                                fmt.Println(v11, v3, v4)
                                v34 := GenS3{F0: (^(v11)), F1: int(v11), F2: (949 <= 468)}
                                var v35 int32 = (v4 ^ 1)
                                _ = v34
                                _ = v35
                            }
                        }
                    case 9223372036854775807:
                        v11 -= 1
                    case -613:
                        if (true && false) {
                            v36 := GenS4{F0: -1, F1: "foo bar"}
                            fmt.Println(v2, v3, v2)
                            _ = v36
                        }
                    }
                }
            } else {
                switch 32767 {
                case 549:
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                case -422:
                    var v37 int = 1
                    _ = v37
                case 105:
                    switch "a" {
                    case "a":
                        switch -596 {
                        case -1:
                            if (!false) {
                                v38 := GenS2{F0: 18446744073709551615, F1: 2}
                                _ = v38
                            } else {
                                v2 = v3
                                if (v13 <= (v13 % ((int32(v11)) | 1))) {
                                    v6 = v6
                                    for fi39 := 0; fi39 < 16; fi39++ {
                                        v4 -= (v4 + (v13 | v4))
                                        _ = fi39
                                    }
                                    for fi40 := 0; fi40 < 11; fi40++ {
                                        fmt.Println(v13, v4)
                                        _ = fi40
                                    }
                                }
                                if true {
                                    if ((0.0 - 1e10) > 0.0) {
                                        v41 := GenS4{F0: -124, F1: "a"}
                                        _ = v41
                                    }
                                }
                            }
                        case 128:
                            v5 = append(v5, 4294967295)
                        case 9223372036854775807:
                            v4 -= -1
                        }
                    case "hello":
                        var v42 uint8 = 239
                        _ = v42
                    }
                }
            }
            if (-29 >= 127) {
                if len(v1) > 0 {
                    v1[0] = ("\n\t" + ((("\n\t" + ("" + ("hello" + (("foo bar" + "") + "foo bar")))) + ("a" + "foo bar")) + "foo bar"))
                }
            } else {
                v43 := GenS4{F0: 0, F1: "foo bar"}
                _ = v43
            }
        } else {
            if len(v5) > 0 {
                v5[0] = 258
            }
            if ((18446744073709551615 > uint64(v11)) && (true || false)) {
                v3 += int16(v6)
            }
            v3 = 117
        }
    } else {
        v44 := GenS3{F0: (v11 / ((997) | 1)), F1: 150, F2: (!false)}
        v12 -= ((float64(1e10) / (v12 / float64(((0.1 * (float32(1e10) - -2.25)) - 1e10)))) / ((v12 - v12) + ((v12 / v12) - ((v12 - 1e10) / v12))))
        var v45 uint16 = 673
        _ = v44
        _ = v45
    }
    for fi46 := 0; fi46 < 18; fi46++ {
        var v47 int16 = v3
        _ = fi46
        _ = v47
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
