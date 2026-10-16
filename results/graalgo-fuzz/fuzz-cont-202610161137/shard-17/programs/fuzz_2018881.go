package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int32; F1 int8; F2 uint32 }
type GenS4 struct { F0 int16; F1 int32 }

func main() {
    var v1 float32 = 3.14159
    var v2 int8 = 100
    v3 := make(map[uint32]uint32)
    var v4 float64 = (1e10 - 1e10)
    var v5 int32 = 0
    v6 := []string{"\n\t", "hello", "foo bar"}
    switch ("" + ("hello" + ("a" + (("\n\t" + ((("foo bar" + "") + "\n\t") + "\n\t")) + "")))) {
    case "\n\t":
        switch "a" {
        case "a":
            v7 := GenS1{F0: -9223372036854775808, F1: (v5 | (v5 | -2147483648))}
            _ = v7
        case "\n\t":
            _ = len(v6)
        }
    case "foo bar":
        var v8 bool = (793 <= 0)
        _ = v8
    }
    var v9 int32 = v5
    var v10 int8 = v2
    if (4294967295 > 0) {
        if (true || ((!(true || (18446744073709551615 > 2))) && (-9223372036854775808 != 0))) {
            switch "a" {
            case "":
                switch "hello" {
                case "a":
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                default:
                    if true {
                        v11 := v3[1]
                        fmt.Println(v4, v5)
                        _ = v11
                    }
                }
            case "hello":
                if true {
                    var v12 uint64 = 18446744073709551615
                    if (20 > 42) {
                        v2 -= v2
                    } else {
                        switch 1 {
                        case 127:
                            var v13 int8 = (v10 & (^(v10)))
                            _ = v13
                        case -954:
                            var v14 bool = (!(!(!(-2147483648 == 287))))
                            _ = v14
                        }
                        for fi15 := 0; fi15 < 2; fi15++ {
                            var v16 uint16 = 1
                            var v17 uint = 65535
                            for fi18 := 0; fi18 < 14; fi18++ {
                                fmt.Println(fi18, v12, v4)
                                v9 = -324
                                fmt.Println(v10, v17, v10)
                                v16 = (v16 * 2)
                                _ = fi18
                            }
                            _ = fi15
                            _ = v16
                            _ = v17
                        }
                    }
                    var v19 uint16 = 32767
                    _ = v12
                    _ = v19
                }
            }
            switch "a" {
            case "hello":
                switch ("hello" + ((("a" + "foo bar") + "hello") + "\n\t")) {
                case "foo bar":
                    fmt.Println(v2)
                }
            case "a":
                for fi20 := 0; fi20 < 11; fi20++ {
                    v21 := GenS2{F0: 361}
                    _ = fi20
                    _ = v21
                }
            case "\n\t":
                var v22 string = ""
                _ = v22
            }
            v2 = (v10 * v10)
        } else {
            var v23 uint16 = 0
            var v24 float32 = float32(v4)
            if (9223372036854775807 > 201) {
                v24 += (float32(v4) - v1)
                if (v23 > v23) {
                    v25 := GenS1{F0: 32768, F1: (-(v9))}
                    _ = v25
                } else {
                    var v26 float32 = -2.25
                    _ = len(v6)
                    for fi27 := 0; fi27 < 20; fi27++ {
                        fmt.Println(v1, v2, fi27)
                        for fi28 := 0; fi28 < 12; fi28++ {
                            var v29 uint = 0
                            var v30 uint32 = 2147483647
                            switch int(v9) {
                            case -966:
                                v31 := GenS3{F0: (v9 & -145), F1: v10, F2: 4294967295}
                                _ = v31
                            case 32767:
                                if (871 > 0) {
                                    var v32 uint16 = (v23 | (^(v23)))
                                    if (((3.14159 / v4) * (1.5 - 1e10)) > ((v4 * v4) / ((((0.1 / ((v4 / (float64(v1) * v4)) + ((v4 + -2.25) * (-2.25 * (1e10 + (v4 * v4)))))) / v4) * (float64(v1) - float64(v1))) / float64(v26)))) {
                                        switch ("hello" + "hello") {
                                        case "":
                                            var v33 float32 = (-2.25 * v24)
                                            _ = v33
                                        default:
                                            switch (^(fi27)) {
                                            case 653:
                                                v4 = v4
                                            case -2147483648:
                                                var v34 uint = (v29 - v29)
                                                _ = v34
                                            case 128:
                                                v10 += (-(v10))
                                            default:
                                                var v35 uint32 = (v30 * v30)
                                                _ = v35
                                            }
                                        }
                                        var v36 int64 = 42
                                        fmt.Println(fi27, fi27)
                                        fmt.Println(v24, v32, fi27)
                                        _ = v36
                                    } else {
                                        var v37 uint16 = (v23 + (v23 ^ (^(v23))))
                                        _ = v37
                                    }
                                    v5 = (v5 & v5)
                                    _ = v32
                                }
                            case 4294967295:
                                var v38 int64 = -574
                                _ = v38
                            default:
                                v24 += 0.1
                            }
                            _ = fi28
                            _ = v29
                            _ = v30
                        }
                        _ = fi27
                    }
                    _ = v26
                }
            }
            _ = v23
            _ = v24
        }
    } else {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
    }
    if (42 < -383) {
        v39 := v3[2]
        _ = v39
    }
    for fi40 := 0; fi40 < 14; fi40++ {
        for fi41 := 0; fi41 < 13; fi41++ {
            if true {
                switch (^(fi40)) {
                case -1:
                    v10 -= v2
                case 1:
                    v6 = append(v6, (("a" + "") + ("foo bar" + "\n\t")))
                case 32767:
                    var v42 int64 = 1
                    _ = v42
                }
                var v43 string = "foo bar"
                _ = v43
            } else {
                fmt.Println(v10, v10)
                var v44 uint64 = 684
                if true {
                    var v45 bool = true
                    v10 = -53
                    _ = v45
                } else {
                    v2 = v10
                    v1 += v1
                }
                delete(v3, 480)
                _ = v44
            }
            var v46 string = "\n\t"
            if true {
                v1 += 0.0
                var v47 uint16 = 912
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                _ = v47
            } else {
                var v48 int32 = (v9 ^ v9)
                var v49 int = fi40
                _ = v48
                _ = v49
            }
            var v50 int16 = -352
            fmt.Println(v46, v1)
            _ = fi41
            _ = v46
            _ = v50
        }
        _ = fi40
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    var v51 uint = 2
    if false {
        if true {
            if false {
                v10 -= (^(v10))
            }
            for fi52 := 0; fi52 < 20; fi52++ {
                var v53 int64 = -9223372036854775808
                v53 += (v53 | 42)
                for fi54 := 0; fi54 < 14; fi54++ {
                    v1 = (v1 + v1)
                    _ = fi54
                }
                _ = fi52
                _ = v53
            }
            switch "\n\t" {
            case "a":
                var v55 float64 = ((0.0 / v4) * 1.5)
                _ = v55
            case "\n\t":
                fmt.Println(v51, v1, v10)
            case "":
                var v56 int8 = (v2 * (v2 + v2))
                _ = v56
            default:
                if (((-(v2)) != v2) && (("a" + ("" + ("foo bar" + "foo bar"))) > "foo bar")) {
                    v4 = v4
                    var v57 int = -1
                    _ = v57
                }
            }
        }
        fmt.Println(v9, v1)
        v58 := v3[2]
        _ = v58
    }
    v10 = (v2 * 1)
    v5 = (v5 * v5)
    if (int64(v10) == -9223372036854775808) {
        var v59 float64 = float64((v1 / v1))
        v59 = ((v4 / (v4 + 0.0)) * (((1.5 / -2.25) * 0.0) + 1.5))
        if true {
            v60 := GenS4{F0: -648, F1: (v9 / ((v5) | 1))}
            for fi61 := 0; fi61 < 1; fi61++ {
                var v62 int = 1
                v63 := GenS1{F0: int(v51), F1: -1}
                _ = fi61
                _ = v62
                _ = v63
            }
            _ = v60
        }
        _ = v59
    }
    v64 := GenS2{F0: 18446744073709551615}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v51
    _ = v6
    _ = v64
    _ = v9
}
