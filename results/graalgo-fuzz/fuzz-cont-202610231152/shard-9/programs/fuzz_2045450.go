package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool }
type GenS2 struct { F0 int32; F1 uint; F2 uint }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int32; F1 int64 }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 int; F1 uint8 }
type GenS9 struct { F0 int8 }
func genFn0(p0 uint64) int16 {
    for fi4 := 0; fi4 < 1; fi4++ {
        v5 := []int{(fi4 % ((624) | 1)), (fi4 ^ fi4), fi4}
        if len(v5) > 0 {
            v5[0] = (fi4 | (fi4 / (((-(fi4))) | 1)))
        }
        _ = fi4
        _ = v5
    }
    v6 := []int{4294967295, int(p0), 1, 9223372036854775807}
    for fi7 := 0; fi7 < 1; fi7++ {
        v8 := GenS1{F0: 897, F1: (!(3.14159 > 1.5))}
        if len(v6) > 0 {
            v6[0] = (fi7 & 9223372036854775807)
        }
        _ = fi7
        _ = v8
    }
    _ = p0
    _ = v6
    return 1
}
func genFn1(p0 uint64) uint32 {
    var v9 uint64 = p0
    v9 -= (v9 + 2)
    var v10 int = -1
    for fi11 := 0; fi11 < 20; fi11++ {
        var v12 int32 = -1
        _ = fi11
        _ = v12
    }
    if (!false) {
        switch ("hello" + ("hello" + ("\n\t" + "foo bar"))) {
        case "foo bar":
            v13 := []int{(-(v10))}
            _ = v13
        }
        v10 += (v10 / (((-(v10))) | 1))
    } else {
        if false {
            v9 = (p0 - v9)
            if (-624 == 42) {
                var v14 int64 = 0
                _ = v14
            }
        }
        if (1 >= -32768) {
            v15 := GenS2{F0: 264, F1: 0, F2: 2}
            v16 := []int{(v10 * v10), v10}
            if true {
                var v17 bool = false
                _ = v17
            } else {
                switch 2147483647 {
                case -9223372036854775808:
                    for fi18 := 0; fi18 < 14; fi18++ {
                        switch "foo bar" {
                        case "\n\t":
                            switch (fi18 / ((v10) | 1)) {
                            case 32768:
                                var v19 uint = 0
                                _ = v19
                            case -1:
                                var v20 uint32 = 42
                                _ = v20
                            default:
                                v10 = (-(fi18))
                            }
                        case "":
                            v10 = v10
                        }
                        v10 -= fi18
                        v21 := GenS3{F0: (fi18 & v10)}
                        _ = fi18
                        _ = v21
                    }
                case 9223372036854775807:
                    if true {
                        var v22 int32 = 0
                        _ = v22
                    } else {
                        v23 := GenS3{F0: (v10 / ((v10) | 1))}
                        _ = v23
                    }
                case 906:
                    if false {
                        if len(v16) >= 2 {
                            v16 = v16[1:len(v16)]
                        }
                        for fi24 := 0; fi24 < 4; fi24++ {
                            v10 = v10
                            fmt.Println(v9, v9)
                            _ = fi24
                        }
                        v9 -= 18446744073709551615
                    }
                }
                _ = len(v16)
            }
            _ = v15
            _ = v16
        }
    }
    _ = p0
    _ = v10
    _ = v9
    return 4294967295
}

func main() {
    v1 := []int{42, 0, -299, 32767}
    var v2 float32 = 1e10
    var v3 uint32 = 2
    if false {
        _ = len(v1)
    }
    fmt.Println(v3)
    if false {
        var v25 int16 = 97
        switch "a" {
        case "foo bar":
            if (genFn1(771) > 9) {
                var v26 uint64 = 4294967295
                v27 := GenS4{F0: -1}
                v1 = append(v1, 4294967295)
                _ = v26
                _ = v27
            }
        case "":
            if false {
                fmt.Println(v2, v3)
                v3 += (v3 / ((734) | 1))
            } else {
                if ((uint8(v3) < 42) && (2 > 0)) {
                    fmt.Println(v3)
                    for fi28 := 0; fi28 < 10; fi28++ {
                        if (128 > 128) {
                            var v29 uint8 = 1
                            if ((v25 * (^(v25))) != genFn0(42)) {
                                switch fi28 {
                                case 73:
                                    fmt.Println(v3, v29, v3)
                                case 0:
                                    for fi30 := 0; fi30 < 16; fi30++ {
                                        v31 := GenS4{F0: (v25 % (((v25 | v25)) | 1))}
                                        var v32 uint32 = (v3 * genFn1(18446744073709551615))
                                        _ = fi30
                                        _ = v31
                                        _ = v32
                                    }
                                }
                            }
                            var v33 uint = 955
                            _ = v29
                            _ = v33
                        } else {
                            v25 -= v25
                            if false {
                                fmt.Println(v3, v3)
                                if false {
                                    var v34 float64 = (1e10 / ((-2.25 - 3.14159) / (1e10 * (0.0 + ((3.14159 / float64(v2)) + 0.0)))))
                                    v2 += (1.5 + v2)
                                    var v35 uint32 = 0
                                    _ = v34
                                    _ = v35
                                }
                                if ((((!(!true)) || true) || true) || (v2 < ((float32((1.5 / -2.25)) * (float32((3.14159 + -2.25)) + v2)) / v2))) {
                                    var v36 int64 = 9223372036854775807
                                    _ = v36
                                } else {
                                    v2 -= v2
                                    for fi37 := 0; fi37 < 13; fi37++ {
                                        v3 = (v3 & (v3 - 2))
                                        var v38 uint = 4294967295
                                        _ = fi37
                                        _ = v38
                                    }
                                    v39 := GenS5{F0: (v25 * v25)}
                                    _ = v39
                                }
                            }
                        }
                        _ = fi28
                    }
                    v40 := GenS5{F0: (v25 - (^(v25)))}
                    _ = v40
                } else {
                    fmt.Println(v25, v3)
                    for fi41 := 0; fi41 < 9; fi41++ {
                        v42 := GenS6{F0: -130, F1: 60}
                        v43 := GenS7{F0: 127}
                        _ = fi41
                        _ = v42
                        _ = v43
                    }
                }
                v2 += ((3.14159 * ((v2 + 1e10) / 0.1)) - float32(3.14159))
                for fi44 := 0; fi44 < 13; fi44++ {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    if (107 > 1) {
                        var v45 int = fi44
                        for fi46 := 0; fi46 < 10; fi46++ {
                            v3 += v3
                            switch fi46 {
                            case 32767:
                                v1 = append(v1, (fi44 ^ (v45 >> ((755) & 63))))
                            case -1:
                                v25 = 563
                            case 2147483647:
                                v1 = append(v1, (^(fi46)))
                            }
                            if true {
                                v3 -= (v3 ^ (v3 | v3))
                                _ = len(v1)
                            }
                            _ = fi46
                        }
                        _ = v45
                    }
                    _ = fi44
                }
            }
        case "hello":
            if len(v1) > 0 {
                v1[0] = 32767
            }
        default:
            fmt.Println(v25, v2, v3)
        }
        fmt.Println(v3, v2, v25)
        _ = v25
    }
    if false {
        for fi47 := 0; fi47 < 7; fi47++ {
            var v48 float32 = v2
            v49 := GenS5{F0: 1}
            switch ("foo bar" + "\n\t") {
            case "a":
                v1 = append(v1, int(v3))
            case "\n\t":
                v1 = append(v1, fi47)
            case "":
                var v50 uint = 42
                _ = v50
            }
            _ = fi47
            _ = v48
            _ = v49
        }
        switch -2147483648 {
        case 1:
            var v51 int8 = 127
            _ = v51
        case 2147483647:
            switch 467 {
            case 121:
                fmt.Println(v2, v2)
            default:
                _ = len(v1)
            }
        case 9223372036854775807:
            v2 = v2
        default:
            v52 := GenS8{F0: -2147483648, F1: 255}
            _ = v52
        }
        v3 -= genFn1(953)
    } else {
        v2 -= 0.0
        v3 -= (v3 >> ((842) & 63))
        var v53 int8 = -35
        _ = v53
    }
    v54 := GenS9{F0: -128}
    v3 = genFn1(0)
    var v55 uint16 = 65535
    for fi56 := 0; fi56 < 15; fi56++ {
        if true {
            v55 += v55
            if (128 < 86) {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = len(v1)
                v57 := GenS8{F0: (fi56 >> ((18446744073709551615) & 63)), F1: 42}
                _ = v57
            }
        } else {
            var v58 bool = true
            _ = v58
        }
        fmt.Println(v2, fi56, v3)
        switch "" {
        case "hello":
            for fi59 := 0; fi59 < 19; fi59++ {
                fmt.Println(v55, fi56)
                _ = fi59
            }
        default:
            v3 = v3
        }
        _ = fi56
    }
    if (!(241 != 1)) {
        var v60 int16 = 1000
        for fi61 := 0; fi61 < 13; fi61++ {
            v1 = append(v1, (fi61 + fi61))
            _ = fi61
        }
        var v62 uint64 = 42
        _ = v60
        _ = v62
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v55:", v55)
    _ = v1
    _ = v2
    _ = v3
    _ = v54
    _ = v55
}
