package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool }
type GenS2 struct { F0 uint64; F1 uint16; F2 int64 }

func main() {
    var v1 uint8 = 42
    var v2 uint32 = 4294967295
    v3 := make(map[string]bool)
    var v4 uint8 = 42
    v5 := GenS1{F0: (v4 * (^(v4))), F1: false}
    if true {
        var v6 int64 = 9223372036854775807
        _ = v6
    } else {
        var v7 int8 = 1
        v8 := []int{-1, -1, 127, int(v1)}
        if ((0.1 * 3.14159) != ((-2.25 / float32(float64(0.1))) - 1.5)) {
            v9 := GenS2{F0: uint64(v2), F1: 42, F2: -1}
            if (("\n\t" + "hello") > "") {
                v1 = v4
                var v10 int32 = -2147483648
                _ = v10
            }
            for fi11 := 0; fi11 < 5; fi11++ {
                if true {
                    var v12 int = fi11
                    fmt.Println(v1)
                    _ = v12
                } else {
                    if len(v8) >= 2 {
                        v8 = v8[1:len(v8)]
                    }
                    var v13 int64 = 42
                    if ((2 != 18446744073709551615) || (!(42 <= 255))) {
                        v4 -= v1
                        v2 += 2
                        switch "\n\t" {
                        case "foo bar":
                            for fi14 := 0; fi14 < 8; fi14++ {
                                var v15 uint = 165
                                var v16 int64 = (v13 / ((v13) | 1))
                                var v17 uint8 = (v1 & v1)
                                _ = fi14
                                _ = v15
                                _ = v16
                                _ = v17
                            }
                        case "hello":
                            for fi18 := 0; fi18 < 13; fi18++ {
                                fmt.Println(v2, v1, v4)
                                var v19 bool = (!false)
                                _ = fi18
                                _ = v19
                            }
                        case "\n\t":
                            if len(v8) > 0 {
                                v8[0] = 921
                            }
                        default:
                            for fi20 := 0; fi20 < 6; fi20++ {
                                v13 = v13
                                var v21 int64 = v13
                                var v22 bool = true
                                _ = fi20
                                _ = v21
                                _ = v22
                            }
                        }
                    } else {
                        for fi23 := 0; fi23 < 7; fi23++ {
                            if (!(!false)) {
                                fmt.Println(fi23)
                                v24 := GenS1{F0: (v4 % (((v1 % ((v4) | 1))) | 1)), F1: ((^(v7)) > v7)}
                                var v25 uint8 = v1
                                _ = v24
                                _ = v25
                            }
                            switch (fi23 & 0) {
                            case 804:
                                fmt.Println(fi11)
                            case 2147483647:
                                switch -9223372036854775808 {
                                case 4294967295:
                                    _ = len(v3)
                                default:
                                    var v26 float32 = 1.5
                                    _ = v26
                                }
                                v8 = append(v8, -1)
                            }
                            if false {
                                var v27 uint64 = 80
                                var v28 int = fi23
                                v29 := GenS1{F0: (v1 & (v1 << ((4294967295) & 63))), F1: true}
                                _ = v27
                                _ = v28
                                _ = v29
                            } else {
                                v30 := GenS1{F0: (v4 | 2), F1: (42 > 205)}
                                if true {
                                    if false {
                                        _ = len(v3)
                                    } else {
                                        var v31 int32 = -110
                                        _ = v31
                                    }
                                    var v32 int8 = v7
                                    _ = v32
                                } else {
                                    var v33 uint16 = 32767
                                    switch "" {
                                    case "a":
                                        fmt.Println(v7, v2)
                                    case "foo bar":
                                        delete(v3, "hello")
                                    case "":
                                        delete(v3, "a")
                                    default:
                                        v4 = (v4 & v1)
                                    }
                                    if false {
                                        v34 := GenS2{F0: 230, F1: 0, F2: v13}
                                        var v35 string = ("" + "hello")
                                        _ = v34
                                        _ = v35
                                    } else {
                                        v13 = int64(fi11)
                                    }
                                    _ = v33
                                }
                                _ = v30
                            }
                            _ = fi23
                        }
                        if true {
                            var v36 float32 = float32(0.1)
                            _ = v36
                        } else {
                            fmt.Println(v4, v1, v2)
                        }
                    }
                    _ = v13
                }
                var v37 int = 2147483647
                _ = fi11
                _ = v37
            }
            _ = v9
        }
        _ = v7
        _ = v8
    }
    var v38 uint = 0
    v4 += 0
    v1 = (v1 + v4)
    for fi39 := 0; fi39 < 7; fi39++ {
        for fi40 := 0; fi40 < 14; fi40++ {
            var v41 uint32 = (^(v2))
            v2 -= 42
            v42 := GenS1{F0: (v1 + v1), F1: false}
            _ = fi40
            _ = v41
            _ = v42
        }
        v43 := GenS1{F0: v4, F1: (1 > 42)}
        _ = fi39
        _ = v43
    }
    for fi44 := 0; fi44 < 18; fi44++ {
        v4 += uint8(v2)
        for fi45 := 0; fi45 < 18; fi45++ {
            var v46 int = (fi44 ^ 4294967295)
            switch 0 {
            case -846:
                if (((v46 >> ((v38) & 63)) >= (v46 << (((^(v38))) & 63))) || false) {
                    var v47 float32 = -2.25
                    _ = v47
                }
            case 42:
                if true {
                    v3[("foo bar" + ((("\n\t" + ("" + "")) + ("\n\t" + "a")) + "a"))] = false
                }
            }
            v4 = v1
            _ = fi45
            _ = v46
        }
        v1 -= uint8(v2)
        _ = fi44
    }
    v2 -= (v2 * (v2 << (((v38 % ((uint(v2)) | 1))) & 63)))
    var v48 int = -1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v48
    _ = v5
}
