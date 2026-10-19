package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16; F2 bool }
type GenS2 struct { F0 int32; F1 uint8; F2 uint64 }
type GenS3 struct { F0 int32; F1 int8 }

func main() {
    var v1 int32 = 40
    v2 := []string{(("a" + "foo bar") + (("\n\t" + "a") + ("" + "\n\t"))), "a", "\n\t"}
    var v3 int = 32768
    var v4 uint8 = 0
    var v5 int64 = 0
    var v6 string = "hello"
    switch (v3 / ((int(v5)) | 1)) {
    case -1:
        v7 := GenS1{F0: 2, F1: 1, F2: false}
        _ = v7
    case 32767:
        var v8 float32 = 1.5
        _ = v8
    default:
        v9 := GenS2{F0: v1, F1: (v4 / (((v4 / ((v4) | 1))) | 1)), F2: 42}
        _ = v9
    }
    v4 = 42
    v10 := GenS2{F0: (v1 * v1), F1: (v4 | (v4 >> ((2) & 63))), F2: 4294967295}
    v11 := GenS1{F0: 42, F1: int16(v4), F2: false}
    switch v3 {
    case 9223372036854775807:
        v2 = append(v2, (v6 + v6))
    case -418:
        fmt.Println(v1, v5, v5)
    default:
        if len(v2) > 0 {
            v2[0] = v6
        }
    }
    var v12 float32 = 0.1
    switch v3 {
    case 42:
        switch v3 {
        case -9223372036854775808:
            for fi13 := 0; fi13 < 14; fi13++ {
                v14 := GenS2{F0: (^(v1)), F1: (^(v4)), F2: 1}
                switch fi13 {
                case 502:
                    v5 += 716
                case 0:
                    var v15 int8 = 1
                    _ = v15
                case -2147483648:
                    if ((!(true || (v5 >= (-(v5))))) || (151 == v4)) {
                        v16 := GenS3{F0: (v1 << ((42) & 63)), F1: 1}
                        if (!(!false)) {
                            v1 += 474
                        }
                        _ = v16
                    }
                default:
                    v12 = (v12 * (v12 + v12))
                }
                var v17 int16 = 1
                _ = fi13
                _ = v14
                _ = v17
            }
        case -847:
            if true {
                var v18 int64 = v5
                _ = v18
            } else {
                var v19 float64 = (1.5 / float64(v12))
                v3 = 4294967295
                _ = v19
            }
        default:
            for fi20 := 0; fi20 < 8; fi20++ {
                fmt.Println(fi20)
                var v21 uint = 65535
                fmt.Println(v4)
                _ = fi20
                _ = v21
            }
        }
    case -702:
        for fi22 := 0; fi22 < 16; fi22++ {
            for fi23 := 0; fi23 < 8; fi23++ {
                switch (v3 / ((127) | 1)) {
                case -9223372036854775808:
                    var v24 uint = 42
                    _ = v24
                case 0:
                    v25 := GenS3{F0: (v1 | v1), F1: -118}
                    _ = v25
                case 354:
                    v12 = v12
                default:
                    var v26 int = 781
                    _ = v26
                }
                var v27 string = (v6 + (("" + (v6 + v6)) + (v6 + (((v6 + "foo bar") + v6) + ("\n\t" + "")))))
                fmt.Println(fi22, v12, v4)
                var v28 float64 = -2.25
                _ = fi23
                _ = v27
                _ = v28
            }
            _ = fi22
        }
    default:
        v12 += 0.1
    }
    v29 := GenS2{F0: 53, F1: 128, F2: 0}
    var v30 string = (("foo bar" + (v6 + ("\n\t" + ("hello" + v6)))) + ("" + v6))
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if false {
        fmt.Println(v1)
    } else {
        v31 := GenS2{F0: (v1 ^ int32(v5)), F1: v4, F2: 18446744073709551615}
        _ = v31
    }
    v1 -= v1
    switch "a" {
    case "a":
        if (!false) {
            var v32 int32 = (v1 << ((2) & 63))
            switch v6 {
            case "foo bar":
                switch (v3 + (v3 * v3)) {
                case 128:
                    v32 = (v1 - 1)
                default:
                    v1 = v1
                }
            case "":
                for fi33 := 0; fi33 < 13; fi33++ {
                    var v34 uint = 532
                    _ = fi33
                    _ = v34
                }
            case "hello":
                for fi35 := 0; fi35 < 13; fi35++ {
                    v36 := GenS2{F0: (v1 & v32), F1: v4, F2: 4294967295}
                    switch v30 {
                    case "hello":
                        if (1 > 1) {
                            fmt.Println(v12, v30)
                            _ = len(v2)
                        } else {
                            v37 := GenS1{F0: 2, F1: 1000, F2: true}
                            var v38 int64 = v5
                            _ = v37
                            _ = v38
                        }
                    case "":
                        fmt.Println(v5, v5)
                    case "a":
                        v5 += (v5 | -305)
                    default:
                        for fi39 := 0; fi39 < 10; fi39++ {
                            v6 = v30
                            _ = fi39
                        }
                    }
                    var v40 bool = (uint(v4) >= 65535)
                    _ = fi35
                    _ = v36
                    _ = v40
                }
            }
            v12 = ((1e10 * 1.5) - ((v12 / v12) / float32(((1e10 + -2.25) * 1e10))))
            _ = v32
        } else {
            var v41 int16 = 32767
            _ = v41
        }
    case "foo bar":
        _ = len(v2)
    }
    var v42 uint = 363
    v43 := GenS1{F0: 0, F1: 42, F2: true}
    if false {
        switch "foo bar" {
        case "":
            if len(v2) > 0 {
                v2[0] = v6
            }
        case "foo bar":
            if (true && (true && (-281 <= v1))) {
                for fi44 := 0; fi44 < 18; fi44++ {
                    var v45 float64 = -2.25
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    var v46 bool = true
                    _ = fi44
                    _ = v45
                    _ = v46
                }
                if (!((v5 >= (v5 + 1)) && (42 == 0))) {
                    var v47 float64 = 3.14159
                    var v48 int16 = 32767
                    for fi49 := 0; fi49 < 9; fi49++ {
                        fmt.Println(v5, v12, v5)
                        var v50 int32 = (v1 * (v1 % ((v1) | 1)))
                        var v51 int32 = v50
                        _ = fi49
                        _ = v50
                        _ = v51
                    }
                    _ = v47
                    _ = v48
                } else {
                    switch (v3 + (v3 / ((v3) | 1))) {
                    case 761:
                        v52 := GenS2{F0: v1, F1: v4, F2: 42}
                        _ = v52
                    case 740:
                        var v53 float64 = 1.5
                        _ = v53
                    case 127:
                        if (!true) {
                            v4 = (v4 & 128)
                        } else {
                            v5 = (v5 >> (((^(v42))) & 63))
                            v42 -= v42
                        }
                    }
                    var v54 uint16 = 650
                    _ = v54
                }
            } else {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                switch v6 {
                case "":
                    var v55 int = v3
                    _ = v55
                case "foo bar":
                    fmt.Println(v3, v30)
                case "a":
                    for fi56 := 0; fi56 < 4; fi56++ {
                        if len(v2) > 0 {
                            v2[0] = "hello"
                        }
                        _ = fi56
                    }
                }
                var v57 int32 = v1
                _ = v57
            }
        default:
            v30 = (((v6 + v30) + "a") + v6)
        }
        var v58 int16 = int16(v3)
        _ = v58
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v42
    _ = v43
    _ = v5
    _ = v6
}
