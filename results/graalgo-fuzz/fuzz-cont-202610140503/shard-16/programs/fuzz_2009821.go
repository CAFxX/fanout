package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 string; F1 int8; F2 string }
type GenS3 struct { F0 uint16; F1 uint64; F2 uint32 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint8; F1 string }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 int16; F1 uint }
type GenS8 struct { F0 uint32; F1 uint32 }
type GenS9 struct { F0 int64 }
type GenS10 struct { F0 int; F1 uint32 }

func main() {
    var v1 uint16 = 0
    var v2 float32 = 3.14159
    var v3 uint64 = 4294967295
    v4 := make(map[string]int)
    v5 := []int{}
    delete(v4, "hello")
    v6 := v4["\n\t"]
    v5 = append(v5, (v6 ^ -4))
    if false {
        switch (v6 * 2147483647) {
        case 956:
            v7 := GenS1{F0: -2147483648}
            _ = v7
        case 1:
            var v8 int32 = -2147483648
            _ = v8
        default:
            var v9 int64 = 42
            _ = v9
        }
    } else {
        v6 = v6
        v3 = 1
    }
    v5 = append(v5, -2147483648)
    v6 += (v6 & 4294967295)
    v1 += uint16(v3)
    switch v6 {
    case 128:
        if true {
            if (100 == -1) {
                if (906 <= 607) {
                    fmt.Println(v2)
                    v1 -= 0
                    v5 = append(v5, v6)
                } else {
                    for fi10 := 0; fi10 < 16; fi10++ {
                        v11 := GenS2{F0: "a", F1: 1, F2: ("foo bar" + ("" + ("\n\t" + (("foo bar" + "") + "a"))))}
                        fmt.Println(v6)
                        var v12 uint16 = 2
                        _ = fi10
                        _ = v11
                        _ = v12
                    }
                }
                switch (v6 - (v6 / ((0) | 1))) {
                case 2147483647:
                    v13 := GenS3{F0: (v1 + (v1 >> ((4294967295) & 63))), F1: 1, F2: 2147483648}
                    _ = v13
                case 833:
                    if true {
                        var v14 bool = (1 != 42)
                        fmt.Println(v6)
                        v15 := GenS3{F0: (v1 ^ 0), F1: (^(v3)), F2: 1}
                        _ = v14
                        _ = v15
                    } else {
                        for fi16 := 0; fi16 < 16; fi16++ {
                            for fi17 := 0; fi17 < 19; fi17++ {
                                if (!(-1 != 1000)) {
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                }
                                var v18 float64 = -2.25
                                if false {
                                    if true {
                                        fmt.Println(fi17)
                                    }
                                    v19 := GenS4{F0: v3}
                                    if true {
                                        if ((v1 & (^(v1))) <= (^(v1))) {
                                            delete(v4, "a")
                                        }
                                        var v20 int32 = 42
                                        if true {
                                            v2 -= v2
                                            v2 -= (3.14159 + v2)
                                            if (float64(0.0) < ((v18 + (((((1e10 / v18) / (((v18 / ((v18 + v18) * v18)) + v18) / 1.0)) - (v18 / v18)) - 1.5) * v18)) - 3.14159)) {
                                                v6 = 704
                                                v2 += float32(-2.25)
                                            }
                                            if len(v5) > 0 {
                                                v5[0] = (^(fi16))
                                            }
                                            v21 := GenS4{F0: v3}
                                            _ = v21
                                        }
                                        _ = v20
                                    } else {
                                        fmt.Println(fi16, fi16)
                                        var v22 uint32 = 0
                                        v23 := GenS5{F0: 42, F1: (("foo bar" + ("foo bar" + ((("hello" + "a") + (("hello" + "a") + "")) + ""))) + "a")}
                                        if (((^(v1)) == v1) || (true && false)) {
                                            v6 -= v6
                                            v3 -= 4294967295
                                            v6 -= (v6 & 32768)
                                        }
                                        fmt.Println(v2)
                                        v24 := GenS6{F0: v3}
                                        _ = v22
                                        _ = v23
                                        _ = v24
                                    }
                                    _ = v19
                                } else {
                                    if ((-888 != -1) || false) {
                                        fmt.Println(v1, v18)
                                    }
                                }
                                _ = fi17
                                _ = v18
                            }
                            _ = fi16
                        }
                        v1 += 10
                        v1 = (^(v1))
                    }
                case 127:
                    var v25 string = ("" + "a")
                    _ = v25
                }
            }
            if true {
                var v26 int32 = -2147483648
                switch "a" {
                case "":
                    for fi27 := 0; fi27 < 7; fi27++ {
                        var v28 uint8 = 255
                        _ = fi27
                        _ = v28
                    }
                }
                v5 = append(v5, v6)
                _ = v26
            }
            v29 := v4[("" + "\n\t")]
            _ = v29
        } else {
            v1 = v1
            v6 -= 128
            if true {
                var v30 uint64 = v3
                var v31 uint16 = (v1 / (((v1 - (v1 >> ((1) & 63)))) | 1))
                var v32 string = "\n\t"
                _ = v30
                _ = v31
                _ = v32
            } else {
                switch "foo bar" {
                case "\n\t":
                    for fi33 := 0; fi33 < 3; fi33++ {
                        if ((-9223372036854775808 != 42) && ((!false) && true)) {
                            v34 := GenS7{F0: 32767, F1: 2}
                            _ = v34
                        }
                        v3 -= v3
                        _ = fi33
                    }
                case "":
                    switch v6 {
                    case 9223372036854775807:
                        if false {
                            var v35 uint64 = (v3 * (^(v3)))
                            if len(v5) > 0 {
                                v5[0] = 175
                            }
                            _ = v35
                        } else {
                            for fi36 := 0; fi36 < 17; fi36++ {
                                if ((v1 << ((249) & 63)) < v1) {
                                    for fi37 := 0; fi37 < 10; fi37++ {
                                        v3 -= 42
                                        _ = fi37
                                    }
                                    v6 = 979
                                    var v38 uint16 = 32767
                                    _ = v38
                                } else {
                                    var v39 uint32 = 65535
                                    v1 = v1
                                    _ = v39
                                }
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                                v2 = v2
                                _ = fi36
                            }
                            v1 = 124
                            for fi40 := 0; fi40 < 4; fi40++ {
                                if len(v5) > 0 {
                                    v5[0] = fi40
                                }
                                var v41 float64 = float64(v2)
                                var v42 int64 = -9223372036854775808
                                _ = fi40
                                _ = v41
                                _ = v42
                            }
                        }
                    }
                case "a":
                    for fi43 := 0; fi43 < 9; fi43++ {
                        if true {
                            var v44 bool = (1 < 530)
                            for fi45 := 0; fi45 < 3; fi45++ {
                                v46 := GenS1{F0: (fi43 << ((370) & 63))}
                                _ = fi45
                                _ = v46
                            }
                            _ = v44
                        } else {
                            v6 = 32767
                        }
                        _ = fi43
                    }
                }
                if true {
                    v47 := GenS8{F0: 1, F1: 2147483647}
                    v2 += v2
                    v6 += (v6 * (^(v6)))
                    _ = v47
                } else {
                    v48 := GenS1{F0: 4294967295}
                    if true {
                        var v49 int = (v6 & v6)
                        _ = v49
                    }
                    _ = v48
                }
                var v50 float32 = v2
                _ = v50
            }
        }
    case 609:
        fmt.Println(v2, v1, v2)
    }
    v51 := GenS9{F0: -697}
    var v52 string = ""
    v6 = (^(v6))
    fmt.Println(v52)
    v53 := GenS10{F0: (v6 >> ((4294967295) & 63)), F1: 42}
    v54 := v4[(v52 + "a")]
    var v55 int8 = -14
    if true {
        var v56 uint32 = 2
        switch (v54 << ((18446744073709551615) & 63)) {
        case -2147483648:
            var v57 float32 = ((-2.25 / ((((v2 * 1.5) - 0.1) + v2) + v2)) * 3.14159)
            _ = v57
        case 508:
            v54 = (v54 - int(v55))
        case -678:
            fmt.Println(v3, v56, v3)
        default:
            var v58 int32 = -2147483648
            _ = v58
        }
        var v59 int64 = 9223372036854775807
        _ = v56
        _ = v59
    } else {
        fmt.Println(v52)
        var v60 int16 = -744
        var v61 uint32 = 2
        _ = v60
        _ = v61
    }
    switch v6 {
    case -508:
        fmt.Println(v55, v52, v54)
    case 128:
        delete(v4, (v52 + (v52 + "")))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v52:", v52)
    fmt.Println("v54:", v54)
    fmt.Println("v55:", v55)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v51
    _ = v52
    _ = v53
    _ = v54
    _ = v55
    _ = v6
}
