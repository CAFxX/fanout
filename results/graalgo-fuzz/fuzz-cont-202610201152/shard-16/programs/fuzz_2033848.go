package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64; F2 int32 }
type GenS2 struct { F0 int; F1 int16 }
type GenS3 struct { F0 uint16; F1 uint; F2 int16 }
type GenS4 struct { F0 int8 }

func main() {
    var v1 int8 = 0
    v2 := []string{(((("a" + ("\n\t" + ("hello" + (((("\n\t" + ("a" + "\n\t")) + "\n\t") + ((("\n\t" + "hello") + ("\n\t" + ("a" + ("foo bar" + "foo bar")))) + "hello")) + "")))) + "hello") + "foo bar") + "hello"), "foo bar"}
    v3 := make(map[int]bool)
    v4 := make(map[uint32]int)
    var v5 bool = (!true)
    var v6 int = -7
    v5 = false
    v1 += v1
    for fi7 := 0; fi7 < 3; fi7++ {
        for fi8 := 0; fi8 < 20; fi8++ {
            v1 -= v1
            v6 = fi7
            _ = fi8
        }
        for fi9 := 0; fi9 < 1; fi9++ {
            if (0 != -32768) {
                fmt.Println(v6)
                if v5 {
                    fmt.Println(fi7, v1, v6)
                    switch (v6 / ((fi9) | 1)) {
                    case 32767:
                        v10 := GenS1{F0: v1, F1: -9223372036854775808, F2: 1}
                        _ = v10
                    case 230:
                        v2 = append(v2, "\n\t")
                    case 127:
                        v11 := GenS2{F0: fi9, F1: 42}
                        _ = v11
                    }
                    switch fi9 {
                    case 727:
                        v12 := GenS2{F0: (fi9 & (-(fi9))), F1: -32768}
                        _ = v12
                    case 32767:
                        var v13 float32 = 0.1
                        _ = v13
                    case 9223372036854775807:
                        switch ("hello" + "") {
                        case "foo bar":
                            v14 := GenS2{F0: (fi9 & fi9), F1: -32768}
                            _ = v14
                        case "hello":
                            switch (fi9 * v6) {
                            case 42:
                                v4[4294967295] = 452
                            case 9223372036854775807:
                                var v15 bool = true
                                _ = v15
                            default:
                                fmt.Println(v6, fi7)
                                var v16 float32 = 3.14159
                                var v17 bool = (((-9223372036854775808 <= -9223372036854775808) && true) || (0.1 < v16))
                                _ = v16
                                _ = v17
                            }
                        }
                        v1 = (v1 * (v1 << ((65535) & 63)))
                    }
                }
                fmt.Println(v1)
            }
            if (!(!(0 == 673))) {
                v2 = append(v2, "\n\t")
            } else {
                var v18 bool = (!((0.0 < ((((1.5 * 0.0) + ((0.1 * (1.5 * 3.14159)) + (((0.0 * 0.0) * 3.14159) + 0.0))) - 0.1) + 0.1)) && false))
                _ = v18
            }
            _ = fi9
        }
        _ = fi7
    }
    for fi19 := 0; fi19 < 9; fi19++ {
        for fi20 := 0; fi20 < 8; fi20++ {
            v21 := GenS1{F0: (v1 + (v1 | (v1 / ((v1) | 1)))), F1: 1, F2: -672}
            switch -887 {
            case -1:
                v3[(v6 >> ((922) & 63))] = v5
            case -850:
                v5 = (-34 == fi19)
            case 568:
                switch fi20 {
                case 32767:
                    v1 += 4
                case 122:
                    var v22 int8 = (v1 - (v1 / ((-6) | 1)))
                    _ = v22
                }
            default:
                var v23 int16 = 1
                _ = v23
            }
            var v24 bool = (v5 && (!(-545 <= 2147483647)))
            var v25 uint8 = 82
            _ = fi20
            _ = v21
            _ = v24
            _ = v25
        }
        _ = fi19
    }
    for fi26 := 0; fi26 < 15; fi26++ {
        for fi27 := 0; fi27 < 20; fi27++ {
            v5 = v5
            delete(v4, 65535)
            v6 = (v6 - v6)
            v5 = (!(true && (false && false)))
            _ = fi27
        }
        switch v6 {
        case 32768:
            var v28 int16 = 42
            _ = v28
        }
        switch "a" {
        case "hello":
            v5 = ("" <= "foo bar")
        case "\n\t":
            var v29 int32 = 0
            _ = v29
        default:
            v30 := GenS3{F0: 32767, F1: 65535, F2: -1}
            _ = v30
        }
        _ = fi26
    }
    if v5 {
        v31 := GenS3{F0: 534, F1: 0, F2: 1}
        _ = v31
    } else {
        for fi32 := 0; fi32 < 11; fi32++ {
            v33 := GenS4{F0: 1}
            v5 = (-331 >= 1)
            var v34 uint64 = 0
            _ = fi32
            _ = v33
            _ = v34
        }
        var v35 uint16 = 2
        fmt.Println(v1, v5, v5)
        _ = v35
    }
    v5 = (386 >= 65535)
    switch 42 {
    case 4294967295:
        v6 = (v6 & (v6 << ((18446744073709551615) & 63)))
    case 2147483647:
        v36 := GenS2{F0: (v6 - v6), F1: -62}
        _ = v36
    case 32767:
        v37 := GenS1{F0: 1, F1: -1, F2: -1}
        _ = v37
    default:
        if len(v2) > 0 {
            v2[0] = "\n\t"
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
