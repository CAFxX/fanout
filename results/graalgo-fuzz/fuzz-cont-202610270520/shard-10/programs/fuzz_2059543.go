package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string }
type GenS2 struct { F0 string; F1 uint16; F2 bool }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint64 }

func main() {
    v1 := []uint32{956, 464}
    var v2 int = 32767
    v3 := make(map[uint32]uint32)
    v4 := make(map[int]uint32)
    v2 -= v2
    v2 += (v2 >> ((18446744073709551615) & 63))
    _ = len(v1)
    if (127 == 0) {
        fmt.Println(v2)
        if ((true && true) || true) {
            var v5 float32 = 1.5
            v6 := GenS1{F0: 2, F1: (("hello" + (("foo bar" + ("foo bar" + "\n\t")) + ("foo bar" + ("" + "")))) + "")}
            _ = v5
            _ = v6
        } else {
            v7 := v4[(v2 ^ (v2 | 459))]
            for fi8 := 0; fi8 < 2; fi8++ {
                var v9 uint8 = 127
                switch fi8 {
                case -1:
                    v10 := GenS2{F0: "a", F1: 614, F2: true}
                    _ = v10
                }
                if (v9 <= 2) {
                    v2 = 127
                    v2 = fi8
                } else {
                    var v11 bool = true
                    _ = v11
                }
                _ = fi8
                _ = v9
            }
            fmt.Println(v7, v2)
            _ = v7
        }
    } else {
        _ = len(v1)
        v12 := GenS2{F0: ("a" + "\n\t"), F1: 255, F2: ((0.0 / 3.14159) < -2.25)}
        for fi13 := 0; fi13 < 16; fi13++ {
            v2 += (v2 / (((fi13 * fi13)) | 1))
            switch (fi13 / (((v2 % ((-9223372036854775808) | 1))) | 1)) {
            case 443:
                for fi14 := 0; fi14 < 6; fi14++ {
                    var v15 int16 = int16(fi13)
                    _ = fi14
                    _ = v15
                }
            case 32768:
                for fi16 := 0; fi16 < 17; fi16++ {
                    if true {
                        var v17 uint32 = 4294967295
                        var v18 int64 = 42
                        if false {
                            fmt.Println(v17, fi13, v2)
                        }
                        _ = v17
                        _ = v18
                    } else {
                        fmt.Println(fi16)
                    }
                    switch "\n\t" {
                    case "foo bar":
                        fmt.Println(v2)
                    default:
                        v19 := GenS1{F0: 488, F1: "a"}
                        _ = v19
                    }
                    _ = fi16
                }
            case -816:
                switch (("\n\t" + "foo bar") + "a") {
                case "":
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                case "\n\t":
                    v2 -= 906
                case "foo bar":
                    var v20 float32 = (0.1 * float32(3.14159))
                    _ = v20
                default:
                    switch (fi13 | v2) {
                    case -9223372036854775808:
                        if true {
                            var v21 bool = false
                            var v22 bool = false
                            var v23 float64 = ((0.0 / 1.0) + float64(1e10))
                            _ = v21
                            _ = v22
                            _ = v23
                        } else {
                            if (1.5 > (-2.25 + -2.25)) {
                                for fi24 := 0; fi24 < 5; fi24++ {
                                    v25 := GenS2{F0: "foo bar", F1: 0, F2: (!true)}
                                    _ = fi24
                                    _ = v25
                                }
                            }
                            v2 -= 42
                        }
                    default:
                        v26 := GenS3{F0: 850}
                        _ = v26
                    }
                }
            }
            _ = fi13
        }
        _ = v12
    }
    v2 = -2147483648
    switch "foo bar" {
    case "\n\t":
        var v27 uint16 = 65535
        _ = v27
    default:
        v28 := GenS3{F0: 18446744073709551615}
        _ = v28
    }
    var v29 int = v2
    v1 = append(v1, 203)
    delete(v4, (^(v29)))
    v2 = 2147483647
    v30 := GenS4{F0: 9223372036854775807}
    var v31 int = (v29 << ((4294967295) & 63))
    switch 9223372036854775807 {
    case -1:
        if (0 != 0) {
            switch "hello" {
            case "\n\t":
                for fi32 := 0; fi32 < 2; fi32++ {
                    var v33 uint8 = 178
                    _ = fi32
                    _ = v33
                }
            }
            for fi34 := 0; fi34 < 16; fi34++ {
                for fi35 := 0; fi35 < 2; fi35++ {
                    v2 = v31
                    var v36 int = -891
                    _ = fi35
                    _ = v36
                }
                if true {
                    var v37 bool = ((0.1 - (1.5 / 1e10)) != 3.14159)
                    _ = v37
                }
                _ = fi34
            }
            var v38 uint64 = 599
            _ = v38
        } else {
            var v39 float32 = (3.14159 + 0.0)
            switch (("\n\t" + "foo bar") + "hello") {
            case "hello":
                var v40 uint16 = 1
                _ = v40
            case "\n\t":
                v29 = v2
            }
            for fi41 := 0; fi41 < 3; fi41++ {
                if (!(-128 > -28)) {
                    switch v2 {
                    case 2147483647:
                        for fi42 := 0; fi42 < 6; fi42++ {
                            var v43 string = "hello"
                            _ = fi42
                            _ = v43
                        }
                    case -1:
                        var v44 int8 = 100
                        _ = v44
                    }
                    v29 = (v31 - v2)
                } else {
                    v45 := GenS5{F0: 18446744073709551615}
                    var v46 bool = (((!false) && (false && (v39 > -2.25))) && (!(9223372036854775807 != 42)))
                    var v47 uint = 883
                    _ = v45
                    _ = v46
                    _ = v47
                }
                var v48 int = (v29 - fi41)
                for fi49 := 0; fi49 < 5; fi49++ {
                    v48 += 1
                    delete(v3, 2)
                    _ = fi49
                }
                _ = fi41
                _ = v48
            }
            _ = v39
        }
    case 127:
        if (true && true) {
            _ = len(v4)
        }
    case 42:
        var v50 uint32 = 4294967295
        _ = v50
    }
    var v51 float64 = float64((1e10 * 1e10))
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v51:", v51)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v51
}
