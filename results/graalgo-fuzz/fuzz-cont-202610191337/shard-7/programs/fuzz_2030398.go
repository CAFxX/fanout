package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string; F2 uint64 }
type GenS2 struct { F0 uint64; F1 uint8 }
type GenS3 struct { F0 int; F1 int32 }
type GenS4 struct { F0 int8; F1 uint8; F2 bool }
type GenS5 struct { F0 bool; F1 bool }

func main() {
    v1 := make(map[string]bool)
    var v2 uint16 = 1
    var v3 float64 = (-2.25 + 0.1)
    var v4 uint8 = 255
    var v5 uint32 = 2147483647
    v6 := []int{}
    fmt.Println(v4)
    for fi7 := 0; fi7 < 12; fi7++ {
        switch ("hello" + "") {
        case "\n\t":
            v2 = v2
        default:
            var v8 int32 = 2147483647
            _ = v8
        }
        switch fi7 {
        case -910:
            _ = len(v1)
        case 32767:
            switch 494 {
            case -9223372036854775808:
                for fi9 := 0; fi9 < 7; fi9++ {
                    if true {
                        switch -999 {
                        case 127:
                            if (936 < 0) {
                                v4 -= v4
                            } else {
                                if true {
                                    v10 := GenS1{F0: 64, F1: "a", F2: 0}
                                    v5 = 84
                                    _ = v10
                                } else {
                                    v11 := GenS1{F0: (v4 << ((uint(v5)) & 63)), F1: ("hello" + "hello"), F2: 18446744073709551615}
                                    var v12 int16 = 900
                                    _ = v11
                                    _ = v12
                                }
                                if false {
                                    var v13 int64 = int64(fi7)
                                    _ = v13
                                } else {
                                    _ = len(v1)
                                    v3 += ((((v3 / 0.1) - (((v3 + (0.0 + (((-2.25 + v3) - -2.25) + (v3 / (v3 - 0.1))))) - ((v3 * (0.1 * v3)) + 3.14159)) + v3)) + 0.1) * (-2.25 * float64((3.14159 - 1.5))))
                                    v14 := GenS1{F0: v4, F1: "foo bar", F2: 4294967295}
                                    _ = v14
                                }
                                if false {
                                    var v15 int64 = 1
                                    v15 += (v15 - (v15 ^ v15))
                                    _ = v15
                                } else {
                                    var v16 int32 = 1
                                    var v17 float64 = v3
                                    v4 -= uint8(v2)
                                    _ = v16
                                    _ = v17
                                }
                            }
                        case 126:
                            if false {
                                fmt.Println(v3, fi7)
                                var v18 uint16 = (v2 - 32767)
                                _ = v18
                            }
                        default:
                            fmt.Println(fi9, v4, v3)
                        }
                        var v19 bool = (false || (48 == 65535))
                        _ = v19
                    } else {
                        if (fi9 >= fi7) {
                            var v20 uint64 = 632
                            _ = len(v6)
                            v2 += v2
                            _ = v20
                        }
                        if false {
                            if true {
                                v6 = append(v6, (fi9 + (fi7 & fi7)))
                            }
                            var v21 float32 = 0.0
                            _ = v21
                        } else {
                            v6 = append(v6, fi9)
                            v2 = v2
                        }
                        var v22 float64 = (float64((1.5 - 0.1)) - v3)
                        _ = v22
                    }
                    _ = fi9
                }
            case 128:
                switch 418 {
                case -976:
                    v3 += 3.14159
                default:
                    switch ("" + "foo bar") {
                    case "hello":
                        v23 := GenS1{F0: (^(v4)), F1: ((("" + "foo bar") + (((("\n\t" + "\n\t") + "a") + ((("a" + "a") + (("a" + ("foo bar" + (((("\n\t" + "\n\t") + "\n\t") + ("\n\t" + "hello")) + "a"))) + "")) + "a")) + "")) + "hello"), F2: 968}
                        _ = v23
                    case "\n\t":
                        v24 := GenS1{F0: v4, F1: "a", F2: 254}
                        _ = v24
                    case "a":
                        switch "\n\t" {
                        case "foo bar":
                            fmt.Println(fi7)
                        case "a":
                            if (true || true) {
                                fmt.Println(v4, v5, v5)
                                v25 := GenS2{F0: uint64(v5), F1: (v4 | 128)}
                                _ = v25
                            }
                        case "\n\t":
                            for fi26 := 0; fi26 < 7; fi26++ {
                                var v27 uint16 = v2
                                v28 := GenS1{F0: v4, F1: ("a" + "hello"), F2: 1}
                                _ = fi26
                                _ = v27
                                _ = v28
                            }
                        default:
                            var v29 uint16 = uint16(v4)
                            _ = v29
                        }
                    }
                }
            case 9223372036854775807:
                var v30 uint32 = (^(v5))
                _ = v30
            default:
                fmt.Println(v4, v2)
            }
        case 42:
            if true {
                var v31 uint8 = 200
                _ = v31
            }
        }
        fmt.Println(v4, fi7)
        _ = fi7
    }
    var v32 int64 = 966
    _ = len(v1)
    var v33 uint16 = uint16(v4)
    for fi34 := 0; fi34 < 1; fi34++ {
        v35 := GenS2{F0: 4294967295, F1: (v4 + v4)}
        v36 := GenS2{F0: 2, F1: v4}
        if (false || (((v4 * (v4 | v4)) > (^(v4))) && (false && true))) {
            _ = len(v1)
            for fi37 := 0; fi37 < 17; fi37++ {
                v38 := GenS3{F0: -71, F1: 42}
                v1["hello"] = (false && ((("\n\t" < "foo bar") || ((true || (42 < 0)) && false)) && false))
                v2 = v33
                var v39 int = fi34
                _ = fi37
                _ = v38
                _ = v39
            }
        } else {
            var v40 uint32 = (v5 + 0)
            switch fi34 {
            case 102:
                v1[("foo bar" + "foo bar")] = ((!true) || true)
            case -9223372036854775808:
                if (true && (!false)) {
                    v41 := GenS1{F0: v4, F1: "a", F2: 2}
                    _ = v41
                } else {
                    var v42 float32 = (1.5 * 3.14159)
                    _ = v42
                }
            default:
                v43 := GenS1{F0: (v4 | v4), F1: "hello", F2: 0}
                _ = v43
            }
            var v44 int32 = 0
            _ = v40
            _ = v44
        }
        _ = fi34
        _ = v35
        _ = v36
    }
    v45 := GenS4{F0: 0, F1: v4, F2: true}
    v3 = v3
    if true {
        for fi46 := 0; fi46 < 12; fi46++ {
            for fi47 := 0; fi47 < 19; fi47++ {
                v2 -= v2
                _ = fi47
            }
            switch (fi46 * 128) {
            case 42:
                v48 := GenS4{F0: 0, F1: 128, F2: (true && true)}
                _ = v48
            case -465:
                switch 42 {
                case 2147483647:
                    v4 = v4
                }
            default:
                v5 += v5
            }
            v49 := GenS5{F0: (fi46 == int(v32)), F1: (127 > -50)}
            _ = fi46
            _ = v49
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v45
    _ = v5
    _ = v6
}
