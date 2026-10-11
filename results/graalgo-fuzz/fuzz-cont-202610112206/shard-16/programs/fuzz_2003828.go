package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint32; F1 int }
type GenS3 struct { F0 bool; F1 string; F2 uint }

func main() {
    var v1 uint32 = 2147483648
    v2 := []uint32{(v1 * (v1 & 2147483648)), v1, v1}
    v3 := []string{"hello", "\n\t", ("" + ("a" + (("foo bar" + "\n\t") + "hello"))), "foo bar"}
    var v4 uint32 = (^(v1))
    v5 := make(map[string]bool)
    v6 := GenS1{F0: 816}
    switch -9223372036854775808 {
    case 958:
        for fi7 := 0; fi7 < 17; fi7++ {
            v8 := GenS1{F0: -1}
            for fi9 := 0; fi9 < 11; fi9++ {
                fmt.Println(fi9, fi7)
                _ = fi9
            }
            v1 = v1
            _ = fi7
            _ = v8
        }
    case 884:
        switch -396 {
        case 0:
            var v10 bool = false
            _ = v10
        }
    case 2147483647:
        for fi11 := 0; fi11 < 7; fi11++ {
            if (-507 > 42) {
                delete(v5, "\n\t")
                v12 := GenS1{F0: -9223372036854775808}
                if ((!false) && (true && ((0.1 - -2.25) > ((1e10 + 1.5) / -2.25)))) {
                    _ = len(v2)
                    var v13 int64 = -9223372036854775808
                    _ = v13
                } else {
                    for fi14 := 0; fi14 < 9; fi14++ {
                        v4 = v1
                        _ = fi14
                    }
                }
                _ = v12
            } else {
                delete(v5, "a")
                if len(v2) > 0 {
                    v2[0] = (^(v1))
                }
            }
            _ = fi11
        }
    }
    var v15 uint = 560
    for fi16 := 0; fi16 < 5; fi16++ {
        var v17 bool = true
        _ = fi16
        _ = v17
    }
    if false {
        v18 := GenS2{F0: 1, F1: 32767}
        _ = v18
    }
    if false {
        if (v4 <= (^(v4))) {
            switch -1 {
            case 931:
                switch 127 {
                case -9223372036854775808:
                    v5["a"] = (true || false)
                }
            default:
                v15 = 1
            }
            v4 = 314
            if len(v3) > 0 {
                v3[0] = "\n\t"
            }
        } else {
            var v19 int = 0
            if ((255 <= 42) && (1 >= 42)) {
                v5["\n\t"] = false
                if false {
                    fmt.Println(v19, v15, v19)
                }
            } else {
                v20 := GenS2{F0: (v1 * v4), F1: v19}
                if true {
                    switch "foo bar" {
                    case "a":
                        var v21 float64 = 3.14159
                        _ = v21
                    case "":
                        v22 := GenS2{F0: (v1 & uint32(v15)), F1: v19}
                        _ = v22
                    case "foo bar":
                        var v23 int8 = 100
                        _ = v23
                    default:
                        var v24 uint64 = 18446744073709551615
                        _ = v24
                    }
                    var v25 int32 = 839
                    _ = v25
                } else {
                    v26 := GenS3{F0: true, F1: ("a" + "\n\t"), F2: 0}
                    v19 = (-(v19))
                    switch "hello" {
                    case "":
                        if len(v3) > 0 {
                            v3[0] = "hello"
                        }
                    case "hello":
                        switch "foo bar" {
                        case "hello":
                            switch "a" {
                            case "foo bar":
                                var v27 uint8 = 159
                                _ = v27
                            default:
                                v3 = append(v3, "")
                            }
                        case "foo bar":
                            var v28 uint32 = uint32(v15)
                            _ = v28
                        case "\n\t":
                            v1 = (v4 >> (((v15 % ((v15) | 1))) & 63))
                        default:
                            v5["hello"] = (843 < 32767)
                        }
                    case "a":
                        v15 = (v15 % (((v15 >> (((v15 << (((v15 & (^(v15)))) & 63))) & 63))) | 1))
                    default:
                        var v29 bool = (1 <= 1)
                        _ = v29
                    }
                    _ = v26
                }
                _ = v20
            }
            v5[""] = (!true)
            _ = v19
        }
        var v30 uint32 = (v1 / (((v4 / ((v1) | 1))) | 1))
        _ = v30
    } else {
        var v31 int8 = 1
        fmt.Println(v15, v31, v31)
        for fi32 := 0; fi32 < 11; fi32++ {
            for fi33 := 0; fi33 < 11; fi33++ {
                if ((42 < 70) && ((720 >= (v15 << ((v15) & 63))) || true)) {
                    v34 := GenS3{F0: ((1.5 > float32((((-2.25 / 1.5) * -2.25) + 3.14159))) || ((fi32 != 32768) || false)), F1: "foo bar", F2: v15}
                    var v35 float32 = 1.5
                    var v36 float32 = 1e10
                    _ = v34
                    _ = v35
                    _ = v36
                }
                _ = fi33
            }
            fmt.Println(v1, v15, v31)
            for fi37 := 0; fi37 < 4; fi37++ {
                v15 += v15
                _ = fi37
            }
            _ = fi32
        }
        _ = v31
    }
    fmt.Println(v4)
    v1 = 1
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v38 int64 = 0
    fmt.Println(v15, v4, v4)
    v15 = v15
    v1 = v4
    for fi39 := 0; fi39 < 13; fi39++ {
        v4 -= (v1 >> (((^(v15))) & 63))
        for fi40 := 0; fi40 < 11; fi40++ {
            var v41 int16 = -497
            _ = fi40
            _ = v41
        }
        var v42 uint16 = 853
        _ = fi39
        _ = v42
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
