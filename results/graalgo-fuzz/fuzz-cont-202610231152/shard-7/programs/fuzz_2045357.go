package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint; F1 int64 }
type GenS3 struct { F0 string; F1 int; F2 bool }
type GenS4 struct { F0 uint32; F1 string; F2 uint32 }
type GenS5 struct { F0 string; F1 int64; F2 int }
type GenS6 struct { F0 uint16; F1 bool }
type GenS7 struct { F0 int }

func main() {
    v1 := make(map[string]uint32)
    var v2 uint = 4294967295
    v3 := []int{2147483647}
    v4 := make(map[uint32]uint32)
    switch 0 {
    case -2147483648:
        if false {
            var v5 bool = (187 > 0)
            v2 = v2
            _ = v5
        }
    case 127:
        v4[4294967295] = 2147483648
    }
    switch 2147483647 {
    case 32768:
        var v6 int32 = 2147483647
        _ = v6
    }
    switch ("\n\t" + "hello") {
    case "hello":
        if false {
            v7 := GenS1{F0: 74}
            v8 := GenS2{F0: (^(v2)), F1: 9223372036854775807}
            if false {
                v4[2] = 861
                _ = len(v3)
                v2 = 1
            }
            _ = v7
            _ = v8
        } else {
            if false {
                fmt.Println(v2)
                var v9 int = -2147483648
                if false {
                    fmt.Println(v9, v2)
                }
                _ = v9
            }
        }
    }
    if (((127 == 100) && true) && (!false)) {
        if len(v3) > 0 {
            v3[0] = 113
        }
        for fi10 := 0; fi10 < 2; fi10++ {
            v2 += 4294967295
            v2 -= v2
            v2 = v2
            _ = fi10
        }
    } else {
        fmt.Println(v2)
        switch "" {
        case "hello":
            switch -1 {
            case 889:
                switch "foo bar" {
                case "a":
                    switch 32767 {
                    case 77:
                        if (38 != 123) {
                            v2 += v2
                            if (!(!(!((-1 > 992) && true)))) {
                                if (!(!true)) {
                                    var v11 uint32 = 2
                                    _ = v11
                                } else {
                                    v12 := v4[42]
                                    var v13 string = "foo bar"
                                    var v14 bool = (!false)
                                    _ = v12
                                    _ = v13
                                    _ = v14
                                }
                                v3 = append(v3, 2147483647)
                                var v15 int8 = 42
                                _ = v15
                            }
                        } else {
                            switch -13 {
                            case -302:
                                _ = len(v4)
                            case 21:
                                v2 -= v2
                            case -2147483648:
                                switch -9223372036854775808 {
                                case -853:
                                    var v16 uint = v2
                                    _ = v16
                                case 128:
                                    for fi17 := 0; fi17 < 13; fi17++ {
                                        v2 = v2
                                        _ = fi17
                                    }
                                }
                            default:
                                v18 := GenS3{F0: "\n\t", F1: 0, F2: (1 < 42)}
                                _ = v18
                            }
                            for fi19 := 0; fi19 < 5; fi19++ {
                                for fi20 := 0; fi20 < 20; fi20++ {
                                    var v21 uint16 = 42
                                    switch "hello" {
                                    case "foo bar":
                                        var v22 uint64 = 585
                                        _ = v22
                                    case "hello":
                                        fmt.Println(v2, fi20, fi20)
                                    }
                                    _ = fi20
                                    _ = v21
                                }
                                _ = fi19
                            }
                            var v23 string = (("hello" + "hello") + (("" + "\n\t") + (((("hello" + "hello") + "foo bar") + "hello") + "")))
                            _ = v23
                        }
                    case 4294967295:
                        v24 := GenS3{F0: ("" + "foo bar"), F1: 128, F2: (1e10 > -2.25)}
                        _ = v24
                    case -860:
                        v3 = append(v3, 9223372036854775807)
                    }
                default:
                    switch ("\n\t" + (("hello" + "hello") + "foo bar")) {
                    case "\n\t":
                        var v25 uint = (v2 * (v2 << ((717) & 63)))
                        _ = v25
                    }
                }
            default:
                fmt.Println(v2)
            }
        case "foo bar":
            if false {
                v2 += (v2 * v2)
                v2 = v2
            } else {
                v2 = (v2 / ((708) | 1))
                var v26 uint32 = 2147483647
                for fi27 := 0; fi27 < 13; fi27++ {
                    v2 = v2
                    v28 := GenS4{F0: (v26 << (((v2 | (^(v2)))) & 63)), F1: ("foo bar" + "foo bar"), F2: (v26 << ((42) & 63))}
                    v29 := GenS2{F0: (v2 ^ v2), F1: -9223372036854775808}
                    _ = fi27
                    _ = v28
                    _ = v29
                }
                _ = v26
            }
        default:
            v30 := GenS5{F0: ("\n\t" + ("" + (("a" + (("a" + "\n\t") + ("hello" + "\n\t"))) + "\n\t"))), F1: -9223372036854775808, F2: 0}
            _ = v30
        }
    }
    v2 -= (v2 * v2)
    switch ("\n\t" + ((("\n\t" + ("\n\t" + (("a" + "hello") + ("" + ("\n\t" + ("a" + "")))))) + "") + "foo bar")) {
    case "a":
        v2 += (^(v2))
    case "hello":
        v31 := GenS4{F0: 2147483647, F1: ("foo bar" + ("foo bar" + ("\n\t" + "\n\t"))), F2: 0}
        _ = v31
    }
    var v32 int64 = 42
    for fi33 := 0; fi33 < 18; fi33++ {
        var v34 bool = true
        fmt.Println(v32, v34, v34)
        _ = fi33
        _ = v34
    }
    var v35 string = ""
    if true {
        v36 := GenS3{F0: (v35 + (v35 + (v35 + (v35 + "")))), F1: 32767, F2: (false || (!((true || (!true)) || (-32768 > -603))))}
        v35 = v35
        switch 9223372036854775807 {
        case -1:
            switch v35 {
            case "hello":
                v4[100] = 520
            default:
                var v37 int16 = -32768
                _ = v37
            }
        }
        _ = v36
    } else {
        v2 += (v2 * (^(v2)))
        if (true || false) {
            var v38 string = v35
            v35 = v35
            switch -9223372036854775808 {
            case 4294967295:
                for fi39 := 0; fi39 < 18; fi39++ {
                    for fi40 := 0; fi40 < 18; fi40++ {
                        switch fi40 {
                        case 4294967295:
                            v2 = v2
                        default:
                            if (false || (255 < 1)) {
                                switch (v38 + ((v35 + "\n\t") + "a")) {
                                case "foo bar":
                                    var v41 int64 = -138
                                    _ = v41
                                case "\n\t":
                                    v42 := GenS6{F0: 916, F1: (true || true)}
                                    _ = v42
                                }
                                v38 = v35
                                var v43 string = (v35 + v38)
                                _ = v43
                            } else {
                                v44 := GenS1{F0: 136}
                                _ = v44
                            }
                        }
                        v35 = v35
                        _ = fi40
                    }
                    _ = fi39
                }
            case -132:
                switch -637 {
                case 165:
                    fmt.Println(v35, v35, v32)
                }
            case 0:
                v45 := v4[42]
                _ = v45
            default:
                for fi46 := 0; fi46 < 12; fi46++ {
                    if (!((v32 / ((v32) | 1)) > v32)) {
                        delete(v4, uint32(v2))
                    }
                    v47 := GenS6{F0: 702, F1: false}
                    var v48 float64 = ((1e10 - -2.25) * 0.1)
                    _ = fi46
                    _ = v47
                    _ = v48
                }
            }
            _ = v38
        } else {
            var v49 int = 2147483647
            var v50 int16 = 1
            _ = v49
            _ = v50
        }
    }
    var v51 int = -9223372036854775808
    v52 := GenS7{F0: 42}
    if (!(978 > 0)) {
        if ((true && (!true)) && true) {
            if (false || true) {
                for fi53 := 0; fi53 < 18; fi53++ {
                    if false {
                        switch v35 {
                        case "hello":
                            switch v35 {
                            case "a":
                                for fi54 := 0; fi54 < 20; fi54++ {
                                    v55 := GenS3{F0: "foo bar", F1: v51, F2: (776 < -423)}
                                    _ = fi54
                                    _ = v55
                                }
                            case "":
                                for fi56 := 0; fi56 < 8; fi56++ {
                                    fmt.Println(v35, v51)
                                    v57 := v4[2]
                                    v51 += (v51 >> (((v2 << (((^(v2))) & 63))) & 63))
                                    _ = fi56
                                    _ = v57
                                }
                            default:
                                fmt.Println(v32, fi53)
                            }
                        case "":
                            var v58 string = (("foo bar" + v35) + v35)
                            _ = v58
                        }
                        var v59 uint64 = 327
                        if true {
                            fmt.Println(v51)
                        }
                        _ = v59
                    } else {
                        var v60 uint16 = 518
                        for fi61 := 0; fi61 < 9; fi61++ {
                            var v62 string = (v35 + (("foo bar" + ("foo bar" + ("foo bar" + ((v35 + v35) + v35)))) + "hello"))
                            if true {
                                if false {
                                    fmt.Println(v51, v51, fi53)
                                }
                            }
                            _ = fi61
                            _ = v62
                        }
                        _ = v60
                    }
                    fmt.Println(v2, v35)
                    var v63 int8 = int8(fi53)
                    _ = fi53
                    _ = v63
                }
            }
            v35 = ("" + v35)
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v51:", v51)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v35
    _ = v4
    _ = v51
    _ = v52
}
