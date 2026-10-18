package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool }
type GenS2 struct { F0 int64; F1 uint64; F2 uint16 }
type GenS3 struct { F0 int64; F1 uint8; F2 int8 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 string; F1 int; F2 uint }
type GenS6 struct { F0 string; F1 uint16 }
func genFn0(p0 int16, p1 uint64, p2 int64) uint8 {
    var v5 uint32 = 65535
    var v6 int = -181
    fmt.Println(p0, v6, p0)
    v7 := GenS1{F0: true, F1: true}
    v8 := GenS1{F0: true, F1: (v5 != 0)}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return 2
}
func genFn1() uint8 {
    var v9 string = (("foo bar" + "\n\t") + ("a" + "\n\t"))
    switch -1 {
    case 2147483647:
        var v10 bool = true
        _ = v10
    case -670:
        var v11 uint64 = 18446744073709551615
        _ = v11
    case 32767:
        fmt.Println(v9)
    }
    v9 = (v9 + (v9 + (v9 + "hello")))
    _ = v9
    return 50
}

func main() {
    var v1 int16 = -1
    var v2 uint = 4294967295
    var v3 uint16 = 32767
    v4 := []uint32{0, 0, 188}
    for fi12 := 0; fi12 < 14; fi12++ {
        var v13 int32 = -380
        fmt.Println(v13, fi12, fi12)
        _ = fi12
        _ = v13
    }
    v14 := GenS2{F0: -821, F1: 2, F2: (v3 << ((v2) & 63))}
    switch ("" + "hello") {
    case "":
        if (-9223372036854775808 >= 1) {
            if false {
                v1 = v1
                switch ("hello" + (("hello" + ("\n\t" + "\n\t")) + ("foo bar" + (((("foo bar" + "") + "a") + ("" + ("hello" + ("" + ("foo bar" + "foo bar"))))) + "a")))) {
                case "hello":
                    v15 := GenS1{F0: (127 < 1), F1: false}
                    _ = v15
                case "\n\t":
                    fmt.Println(v2)
                case "a":
                    v2 -= v2
                default:
                    switch 9223372036854775807 {
                    case 127:
                        _ = len(v4)
                    case 4294967295:
                        switch 1 {
                        case -1:
                            for fi16 := 0; fi16 < 12; fi16++ {
                                var v17 uint8 = 240
                                for fi18 := 0; fi18 < 3; fi18++ {
                                    if ((1 < 96) || false) {
                                        var v19 int32 = -366
                                        _ = v19
                                    } else {
                                        v2 = (v2 * 18446744073709551615)
                                        var v20 uint32 = 1
                                        v21 := GenS3{F0: -971, F1: v17, F2: -1}
                                        _ = v20
                                        _ = v21
                                    }
                                    _ = fi18
                                }
                                _ = fi16
                                _ = v17
                            }
                        case 127:
                            var v22 uint64 = 643
                            _ = v22
                        case 42:
                            if true {
                                var v23 float32 = 1.5
                                for fi24 := 0; fi24 < 11; fi24++ {
                                    fmt.Println(fi24, v3)
                                    v23 -= -2.25
                                    var v25 bool = false
                                    _ = fi24
                                    _ = v25
                                }
                                var v26 int32 = -2147483648
                                _ = v23
                                _ = v26
                            }
                        }
                    default:
                        switch 42 {
                        case 9223372036854775807:
                            v1 -= v1
                        }
                    }
                }
            } else {
                v1 -= (v1 - -132)
                v3 -= (v3 ^ v3)
            }
        }
    case "hello":
        v1 += v1
    default:
        switch "hello" {
        case "\n\t":
            for fi27 := 0; fi27 < 20; fi27++ {
                v28 := GenS1{F0: false, F1: (!false)}
                for fi29 := 0; fi29 < 3; fi29++ {
                    var v30 int8 = -1
                    var v31 float32 = 3.14159
                    _ = fi29
                    _ = v30
                    _ = v31
                }
                switch "hello" {
                case "hello":
                    v3 += (v3 & v3)
                case "":
                    if ("foo bar" <= "a") {
                        var v32 int64 = -599
                        _ = v32
                    } else {
                        v3 -= 2
                    }
                default:
                    for fi33 := 0; fi33 < 11; fi33++ {
                        switch 4294967295 {
                        case 128:
                            var v34 int64 = 0
                            _ = v34
                        case 1:
                            if len(v4) > 0 {
                                v4[0] = 1
                            }
                        case -1:
                            if true {
                                var v35 string = "hello"
                                v4 = append(v4, 2147483648)
                                var v36 uint32 = 1
                                _ = v35
                                _ = v36
                            } else {
                                v37 := GenS1{F0: (fi33 != (fi33 | int(v2))), F1: false}
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                switch 4294967295 {
                                case 32768:
                                    fmt.Println(fi33)
                                }
                                _ = v37
                            }
                        }
                        _ = fi33
                    }
                }
                _ = fi27
                _ = v28
            }
        case "a":
            if (false && true) {
                for fi38 := 0; fi38 < 8; fi38++ {
                    fmt.Println(v2)
                    _ = fi38
                }
                for fi39 := 0; fi39 < 18; fi39++ {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    fmt.Println(v1)
                    _ = fi39
                }
            }
        case "foo bar":
            if len(v4) > 0 {
                v4[0] = 1
            }
        }
    }
    fmt.Println(v2)
    switch ("a" + "hello") {
    case "hello":
        v4 = append(v4, 29)
    case "foo bar":
        v40 := GenS3{F0: -1, F1: 255, F2: -1}
        _ = v40
    case "":
        v2 = 4294967295
    default:
        v41 := GenS4{F0: (v1 & -217)}
        _ = v41
    }
    v1 -= v1
    var v42 int = 604
    v2 = 42
    fmt.Println(v3)
    if true {
        var v43 uint64 = 930
        _ = v43
    } else {
        var v44 int16 = (v1 * v1)
        _ = v44
    }
    for fi45 := 0; fi45 < 13; fi45++ {
        fmt.Println(fi45, v42, v1)
        _ = fi45
    }
    var v46 uint64 = 18446744073709551615
    for fi47 := 0; fi47 < 14; fi47++ {
        v48 := GenS4{F0: 32767}
        _ = fi47
        _ = v48
    }
    if false {
        if true {
            v4 = append(v4, 4294967295)
            v4 = append(v4, 16)
        }
        var v49 uint8 = 0
        _ = v49
    }
    v4 = append(v4, 4294967295)
    if (true || ((v2 ^ 18446744073709551615) == (v2 / (((v2 / ((uint(v3)) | 1))) | 1)))) {
        v1 = v1
    } else {
        for fi50 := 0; fi50 < 8; fi50++ {
            var v51 int64 = 783
            var v52 uint16 = v3
            _ = fi50
            _ = v51
            _ = v52
        }
        if true {
            v53 := GenS3{F0: 856, F1: 1, F2: 116}
            v54 := GenS4{F0: 42}
            _ = v53
            _ = v54
        } else {
            fmt.Println(v46, v2, v3)
        }
        for fi55 := 0; fi55 < 5; fi55++ {
            for fi56 := 0; fi56 < 6; fi56++ {
                if false {
                    if ((true && false) && false) {
                        var v57 bool = (!(0.0 <= 0.0))
                        _ = v57
                    }
                    var v58 uint = 65535
                    _ = v58
                }
                fmt.Println(fi55)
                var v59 float32 = -2.25
                _ = fi56
                _ = v59
            }
            for fi60 := 0; fi60 < 4; fi60++ {
                v3 = v3
                v61 := GenS4{F0: (-(v1))}
                switch (fi55 % ((32768) | 1)) {
                case 119:
                    var v62 uint32 = 2147483647
                    _ = v62
                case -9223372036854775808:
                    if false {
                        v3 += 2
                        fmt.Println(v42)
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        fmt.Println(v46, v2, v3)
                    }
                default:
                    _ = len(v4)
                }
                v63 := GenS5{F0: "a", F1: 567, F2: v2}
                fmt.Println(v1, fi60)
                _ = fi60
                _ = v61
                _ = v63
            }
            _ = fi55
        }
    }
    v64 := GenS6{F0: "foo bar", F1: (v3 / ((v3) | 1))}
    v42 = 2147483647
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v42:", v42)
    fmt.Println("v46:", v46)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v46
    _ = v64
}
