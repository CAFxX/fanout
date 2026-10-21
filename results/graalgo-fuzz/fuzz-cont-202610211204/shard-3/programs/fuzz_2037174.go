package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint; F2 uint16 }
type GenS2 struct { F0 int8; F1 uint8; F2 uint }
type GenS3 struct { F0 int32; F1 uint64 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 bool; F1 uint16 }
type GenS7 struct { F0 int64 }

func main() {
    var v1 uint = 4294967295
    var v2 int = 2147483647
    var v3 int16 = -1
    v4 := []int{}
    if false {
        switch ("hello" + ((("foo bar" + ("foo bar" + "hello")) + ("a" + ("a" + "foo bar"))) + "")) {
        case "":
            v5 := GenS1{F0: 128, F1: 2, F2: 42}
            _ = v5
        case "\n\t":
            var v6 uint = 0
            _ = v6
        default:
            v3 = v3
        }
    } else {
        switch v2 {
        case 0:
            fmt.Println(v1, v3)
        case 9223372036854775807:
            var v7 float32 = 1.5
            _ = v7
        default:
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        }
    }
    v8 := GenS1{F0: -92, F1: v1, F2: 700}
    switch "" {
    case "":
        for fi9 := 0; fi9 < 12; fi9++ {
            for fi10 := 0; fi10 < 16; fi10++ {
                var v11 int32 = -2147483648
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                _ = fi10
                _ = v11
            }
            var v12 string = "a"
            for fi13 := 0; fi13 < 12; fi13++ {
                if len(v4) > 0 {
                    v4[0] = 0
                }
                _ = fi13
            }
            _ = fi9
            _ = v12
        }
    case "\n\t":
        fmt.Println(v1)
    default:
        v4 = append(v4, (v2 - 1))
    }
    var v14 uint16 = 65535
    var v15 uint16 = 328
    switch "hello" {
    case "foo bar":
        if ((42 >= 880) || (!(1.5 >= 3.14159))) {
            for fi16 := 0; fi16 < 4; fi16++ {
                _ = len(v4)
                switch 401 {
                case 385:
                    var v17 int8 = 12
                    _ = v17
                case 9223372036854775807:
                    if false {
                        v15 = (v14 * uint16(v1))
                        v18 := GenS2{F0: -1, F1: 1, F2: 648}
                        _ = v18
                    }
                case -718:
                    if ((v1 - 65535) == (v1 / ((v1) | 1))) {
                        v3 -= (v3 * v3)
                        v15 += (v15 | 32767)
                        fmt.Println(fi16)
                    } else {
                        if true {
                            _ = len(v4)
                            if (-1 == -843) {
                                switch (v2 + 173) {
                                case 2147483647:
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                }
                                var v19 int64 = 1
                                v20 := GenS1{F0: fi16, F1: (v1 % (((v1 << ((318) & 63))) | 1)), F2: v14}
                                _ = v19
                                _ = v20
                            } else {
                                _ = len(v4)
                            }
                            switch "foo bar" {
                            case "\n\t":
                                v14 += v14
                            case "hello":
                                if (18446744073709551615 <= 634) {
                                    v21 := GenS3{F0: 42, F1: 883}
                                    if (false || ((true || false) && (628 < 783))) {
                                        fmt.Println(fi16)
                                    }
                                    _ = v21
                                } else {
                                    switch "hello" {
                                    case "\n\t":
                                        if (!(0 <= 307)) {
                                            v14 += v14
                                        }
                                    default:
                                        switch "" {
                                        case "foo bar":
                                            v15 = 0
                                        case "":
                                            switch v2 {
                                            case 32768:
                                                if ((fi16 ^ (fi16 / ((v2) | 1))) != int(v3)) {
                                                    var v22 uint8 = 255
                                                    for fi23 := 0; fi23 < 3; fi23++ {
                                                        v24 := GenS4{F0: 0}
                                                        var v25 uint16 = (v14 << ((v1) & 63))
                                                        fmt.Println(v22)
                                                        if len(v4) > 0 {
                                                            v4[0] = int(v1)
                                                        }
                                                        _ = fi23
                                                        _ = v24
                                                        _ = v25
                                                    }
                                                    _ = v22
                                                }
                                            default:
                                                if true {
                                                    var v26 uint16 = 65535
                                                    for fi27 := 0; fi27 < 7; fi27++ {
                                                        if false {
                                                            v4 = append(v4, 32768)
                                                            var v28 float32 = (3.14159 / -2.25)
                                                            _ = v28
                                                        }
                                                        switch (("a" + ("foo bar" + ("\n\t" + "foo bar"))) + ("\n\t" + "hello")) {
                                                        case "foo bar":
                                                            var v29 uint64 = 42
                                                            _ = v29
                                                        default:
                                                            if false {
                                                                if len(v4) > 0 {
                                                                    v4[0] = (fi27 % (((fi16 / (((fi16 * (fi16 ^ (v2 - (v2 + 2147483647))))) | 1))) | 1))
                                                                }
                                                                v26 -= (v14 / ((v15) | 1))
                                                                v30 := GenS5{F0: v1}
                                                                _ = v30
                                                            } else {
                                                                v31 := GenS5{F0: (v1 >> ((196) & 63))}
                                                                _ = v31
                                                            }
                                                        }
                                                        var v32 uint = v1
                                                        v33 := GenS6{F0: true, F1: 209}
                                                        _ = fi27
                                                        _ = v32
                                                        _ = v33
                                                    }
                                                    _ = v26
                                                }
                                            }
                                        case "hello":
                                            var v34 uint16 = v14
                                            _ = v34
                                        default:
                                            fmt.Println(fi16)
                                        }
                                    }
                                }
                            case "":
                                switch (fi16 & v2) {
                                case 32768:
                                    for fi35 := 0; fi35 < 20; fi35++ {
                                        fmt.Println(v14)
                                        v36 := GenS1{F0: (fi35 / ((fi35) | 1)), F1: (v1 | 255), F2: v15}
                                        _ = fi35
                                        _ = v36
                                    }
                                default:
                                    var v37 string = "foo bar"
                                    _ = v37
                                }
                            }
                        }
                        fmt.Println(v15)
                    }
                default:
                    fmt.Println(v1)
                }
                if true {
                    var v38 int = fi16
                    var v39 uint8 = 0
                    _ = v38
                    _ = v39
                } else {
                    for fi40 := 0; fi40 < 3; fi40++ {
                        v15 += (v14 & (v14 + 1))
                        fmt.Println(fi16, fi40)
                        if true {
                            if false {
                                var v41 string = "\n\t"
                                v4 = append(v4, (fi16 + fi40))
                                v4 = append(v4, fi40)
                                _ = v41
                            } else {
                                fmt.Println(fi16, v1, fi40)
                                var v42 uint16 = (v15 / ((255) | 1))
                                switch 888 {
                                case 42:
                                    v43 := GenS7{F0: int64(v3)}
                                    v1 -= (v1 >> ((v1) & 63))
                                    v14 += (v42 % ((65535) | 1))
                                    _ = v43
                                case 128:
                                    fmt.Println(v14, v3, v14)
                                case 172:
                                    if len(v4) > 0 {
                                        v4[0] = (fi40 * (v2 - 127))
                                    }
                                default:
                                    var v44 float64 = ((1e10 + 3.14159) - (float64(float32(1.5)) * (0.0 - (3.14159 / -2.25))))
                                    _ = v44
                                }
                                _ = v42
                            }
                        }
                        _ = fi40
                    }
                }
                _ = fi16
            }
        }
    default:
        if false {
            v1 = (v1 >> ((v1) & 63))
        }
    }
    v2 = 9223372036854775807
    for fi45 := 0; fi45 < 2; fi45++ {
        v15 -= (v14 << ((v1) & 63))
        var v46 uint32 = 4294967295
        v4 = append(v4, fi45)
        _ = fi45
        _ = v46
    }
    var v47 string = ("foo bar" + ("" + "foo bar"))
    _ = len(v4)
    v1 = 0
    v14 = v14
    for fi48 := 0; fi48 < 13; fi48++ {
        v1 = 4294967295
        _ = fi48
    }
    if false {
        v2 = (v2 ^ (v2 << ((42) & 63)))
        v47 = (((v47 + ("\n\t" + ((("foo bar" + v47) + "\n\t") + (v47 + "hello")))) + "a") + v47)
        var v49 string = "hello"
        _ = v49
    } else {
        v15 = (v15 & (v15 / ((v15) | 1)))
        var v50 bool = true
        _ = v50
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v47:", v47)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v8
}
