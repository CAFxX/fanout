package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32 }
type GenS2 struct { F0 int32; F1 int64 }
type GenS3 struct { F0 int; F1 bool }
type GenS4 struct { F0 int8; F1 int32; F2 uint64 }
type GenS5 struct { F0 string; F1 int64; F2 int16 }
type GenS6 struct { F0 int8; F1 uint }

func main() {
    var v1 int8 = 29
    var v2 int16 = int16(v1)
    var v3 int32 = 0
    var v4 bool = false
    for fi5 := 0; fi5 < 15; fi5++ {
        var v6 float64 = 3.14159
        _ = fi5
        _ = v6
    }
    if true {
        switch int(v1) {
        case 128:
            var v7 int = 32767
            _ = v7
        case 1:
            fmt.Println(v3, v4, v3)
        case 4294967295:
            for fi8 := 0; fi8 < 19; fi8++ {
                if v4 {
                    var v9 uint32 = 96
                    var v10 int16 = (v2 * -32768)
                    _ = v10
                    _ = v9
                }
                v11 := []int{(fi8 & (fi8 % ((-9223372036854775808) | 1))), (fi8 >> ((0) & 63)), (fi8 * 275), -2147483648}
                v3 += (^(v3))
                _ = fi8
                _ = v11
            }
        default:
            if true {
                v1 = v1
                v1 = (v1 ^ v1)
                for fi12 := 0; fi12 < 16; fi12++ {
                    if (v3 <= v3) {
                        var v13 int8 = (v1 - (v1 / (((v1 % ((v1) | 1))) | 1)))
                        _ = v13
                    }
                    _ = fi12
                }
            }
        }
        var v14 float32 = 0.0
        v15 := GenS1{F0: 819, F1: v3}
        _ = v14
        _ = v15
    } else {
        for fi16 := 0; fi16 < 4; fi16++ {
            var v17 float64 = 3.14159
            _ = fi16
            _ = v17
        }
    }
    v18 := []int{4294967295, 42, 0, 724}
    v4 = v4
    v19 := GenS1{F0: 2, F1: v3}
    for fi20 := 0; fi20 < 7; fi20++ {
        switch (("a" + "a") + ("\n\t" + "hello")) {
        case "\n\t":
            if (0 <= 42) {
                switch (fi20 ^ (-(fi20))) {
                case 915:
                    if (v4 || (float32(1e10) == 0.0)) {
                        var v21 int16 = (v2 | (v2 | (v2 * (v2 ^ (v2 - (v2 << ((255) & 63)))))))
                        fmt.Println(fi20, v4, v3)
                        var v22 bool = false
                        _ = v21
                        _ = v22
                    }
                }
                v2 -= (v2 << ((0) & 63))
                v23 := GenS1{F0: 827, F1: v3}
                _ = v23
            } else {
                for fi24 := 0; fi24 < 7; fi24++ {
                    var v25 bool = v4
                    _ = fi24
                    _ = v25
                }
            }
        }
        for fi26 := 0; fi26 < 1; fi26++ {
            var v27 int8 = (-(v1))
            v4 = v4
            _ = fi26
            _ = v27
        }
        _ = fi20
    }
    _ = len(v18)
    for fi28 := 0; fi28 < 18; fi28++ {
        v29 := GenS2{F0: v3, F1: -9223372036854775808}
        _ = fi28
        _ = v29
    }
    v1 = v1
    switch "" {
    case "":
        if v4 {
            if true {
                fmt.Println(v3, v4)
                if len(v18) >= 2 {
                    v18 = v18[1:len(v18)]
                }
                v3 = (v3 >> ((18446744073709551615) & 63))
            } else {
                if (v3 < 2147483647) {
                    if len(v18) >= 2 {
                        v18 = v18[1:len(v18)]
                    }
                } else {
                    if false {
                        var v30 uint8 = 255
                        v1 -= (v1 % (((^(v1))) | 1))
                        _ = v30
                    } else {
                        v2 = (v2 << ((1) & 63))
                    }
                    v1 = -128
                }
                _ = len(v18)
                v2 -= (v2 >> ((0) & 63))
            }
        } else {
            for fi31 := 0; fi31 < 9; fi31++ {
                v32 := GenS1{F0: 897, F1: (-(v3))}
                if (("\n\t" != ("foo bar" + "foo bar")) && (0.1 != ((3.14159 * ((0.0 - 1.5) * 0.1)) * float64(0.1)))) {
                    v3 += (v3 ^ v3)
                    v2 = v2
                    if (!(!true)) {
                        v33 := GenS3{F0: fi31, F1: (0.0 < 3.14159)}
                        fmt.Println(v1, v2, fi31)
                        if ((v3 | (-(v3))) >= 1) {
                            for fi34 := 0; fi34 < 2; fi34++ {
                                v3 = -1
                                if (v4 || (-2147483648 != (fi31 << ((18446744073709551615) & 63)))) {
                                    var v35 int8 = v1
                                    v4 = v4
                                    _ = v35
                                } else {
                                    switch ("hello" + "a") {
                                    case "foo bar":
                                        if false {
                                            v18 = append(v18, int(v3))
                                            if (244 <= 255) {
                                                fmt.Println(v3, v2, fi31)
                                                if (v2 < -466) {
                                                    v2 += v2
                                                } else {
                                                    v2 = 1
                                                }
                                            } else {
                                                v36 := GenS1{F0: 65535, F1: v3}
                                                _ = v36
                                            }
                                            v18 = append(v18, (fi31 & (fi31 * fi34)))
                                        }
                                    }
                                    var v37 int32 = (v3 ^ v3)
                                    _ = v37
                                }
                                _ = fi34
                            }
                            var v38 string = ("" + "hello")
                            _ = v38
                        }
                        _ = v33
                    } else {
                        if (!((156 != 1) && (("hello" < "\n\t") || v4))) {
                            v4 = false
                            v39 := GenS4{F0: (v1 | (v1 >> ((1) & 63))), F1: v3, F2: 4294967295}
                            _ = v39
                        }
                    }
                }
                v40 := GenS1{F0: 1, F1: 2147483647}
                _ = fi31
                _ = v32
                _ = v40
            }
        }
    case "foo bar":
        v3 += (^(v3))
    case "hello":
        var v41 int32 = 330
        _ = v41
    default:
        v42 := GenS2{F0: (v3 % (((v3 | 138)) | 1)), F1: 9223372036854775807}
        _ = v42
    }
    if (!(((-2.25 * ((0.0 + -2.25) - (float64(0.0) / 1.0))) + 3.14159) <= 1.5)) {
        var v43 uint64 = 422
        _ = v43
    } else {
        v1 = 42
        v1 += -128
    }
    v44 := GenS5{F0: "hello", F1: 42, F2: v2}
    for fi45 := 0; fi45 < 7; fi45++ {
        switch fi45 {
        case 2147483647:
            if len(v18) > 0 {
                v18[0] = fi45
            }
        default:
            v46 := GenS5{F0: "foo bar", F1: -424, F2: (v2 & v2)}
            _ = v46
        }
        v47 := GenS1{F0: 65535, F1: 0}
        var v48 float64 = (float64((((-2.25 / 1.5) - ((1e10 * 3.14159) + 1.5)) - (float32(-2.25) - 3.14159))) + 1e10)
        _ = fi45
        _ = v47
        _ = v48
    }
    switch 128 {
    case 827:
        v49 := GenS6{F0: 100, F1: 255}
        _ = v49
    default:
        for fi50 := 0; fi50 < 8; fi50++ {
            if (2 <= 625) {
                v2 -= 0
                var v51 uint16 = 255
                _ = v51
            }
            if true {
                v1 = v1
                if true {
                    for fi52 := 0; fi52 < 15; fi52++ {
                        fmt.Println(v3, fi52, fi50)
                        var v53 uint32 = 65535
                        _ = fi52
                        _ = v53
                    }
                    v54 := GenS3{F0: (-(fi50)), F1: v4}
                    if (-566 > -9223372036854775808) {
                        v2 -= (v2 | 32767)
                    } else {
                        fmt.Println(v4, v2, v1)
                        for fi55 := 0; fi55 < 7; fi55++ {
                            var v56 int32 = v3
                            fmt.Println(v3, v2)
                            var v57 uint16 = 32767
                            _ = fi55
                            _ = v56
                            _ = v57
                        }
                        fmt.Println(v2)
                    }
                    _ = v54
                }
            }
            var v58 uint64 = 4294967295
            _ = fi50
            _ = v58
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v44
}
