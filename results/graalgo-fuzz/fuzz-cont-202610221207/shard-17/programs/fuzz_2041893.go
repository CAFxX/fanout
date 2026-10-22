package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64; F2 int }
type GenS2 struct { F0 int; F1 bool; F2 string }
type GenS3 struct { F0 int; F1 bool; F2 int }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int64; F1 uint16 }

func main() {
    var v1 uint8 = 22
    var v2 uint8 = (v1 >> ((18446744073709551615) & 63))
    v3 := []string{"hello", ("hello" + ("a" + "hello")), "a", ("a" + "")}
    var v4 uint32 = 65535
    v5 := []string{}
    var v6 int64 = 151
    switch "hello" {
    case "":
        v4 += v4
    case "a":
        switch 4294967295 {
        case 127:
            if false {
                var v7 int32 = 1
                _ = v7
            }
        case -2147483648:
            _ = len(v3)
        case 4294967295:
            v1 = (v2 % ((1) | 1))
        }
    default:
        var v8 int64 = (v6 << ((uint(v4)) & 63))
        _ = v8
    }
    var v9 int = -9223372036854775808
    var v10 int32 = 0
    if len(v3) > 0 {
        v3[0] = (("hello" + "hello") + "")
    }
    v11 := GenS1{F0: false, F1: 266, F2: v9}
    var v12 uint16 = uint16(v4)
    fmt.Println(v4, v10)
    switch int(v2) {
    case 4294967295:
        if true {
            for fi13 := 0; fi13 < 12; fi13++ {
                fmt.Println(v9)
                v12 -= v12
                for fi14 := 0; fi14 < 13; fi14++ {
                    var v15 int64 = 42
                    v12 = 65535
                    if ((v12 | 42) > (v12 >> ((255) & 63))) {
                        var v16 int = 128
                        _ = v16
                    } else {
                        if len(v5) > 0 {
                            v5[0] = "\n\t"
                        }
                    }
                    _ = fi14
                    _ = v15
                }
                _ = fi13
            }
            if true {
                switch (v9 >> ((255) & 63)) {
                case -300:
                    var v17 uint64 = 18446744073709551615
                    _ = v17
                case -9:
                    fmt.Println(v12, v1)
                case -444:
                    v4 = (v4 & (v4 % ((v4) | 1)))
                default:
                    for fi18 := 0; fi18 < 18; fi18++ {
                        if (!(!false)) {
                            v19 := GenS1{F0: (3.14159 == 1.5), F1: uint64(v12), F2: v9}
                            v20 := GenS1{F0: false, F1: 34, F2: v9}
                            _ = v19
                            _ = v20
                        }
                        _ = fi18
                    }
                }
            } else {
                switch "\n\t" {
                case "hello":
                    if len(v3) > 0 {
                        v3[0] = ("foo bar" + (("foo bar" + "foo bar") + "a"))
                    }
                case "foo bar":
                    for fi21 := 0; fi21 < 18; fi21++ {
                        for fi22 := 0; fi22 < 4; fi22++ {
                            v23 := GenS2{F0: int(v10), F1: false, F2: "a"}
                            v24 := GenS3{F0: fi22, F1: false, F2: v9}
                            if ((!(false || (v4 <= v4))) || false) {
                                if true {
                                    v25 := GenS4{F0: (v12 / ((v12) | 1))}
                                    v9 = fi22
                                    _ = v25
                                }
                                fmt.Println(v2, v6)
                                v5 = append(v5, ("foo bar" + "a"))
                            }
                            _ = fi22
                            _ = v23
                            _ = v24
                        }
                        v26 := GenS5{F0: v6, F1: v12}
                        fmt.Println(fi21, v9, v4)
                        _ = fi21
                        _ = v26
                    }
                case "\n\t":
                    fmt.Println(v10)
                default:
                    _ = len(v3)
                }
                v9 += v9
                var v27 int8 = -23
                _ = v27
            }
        }
    case 1:
        var v28 int16 = 0
        _ = v28
    }
    var v29 uint16 = (v12 << ((178) & 63))
    var v30 int32 = -412
    fmt.Println(v10)
    var v31 string = "a"
    switch v9 {
    case -225:
        v9 += v9
    case -301:
        v29 += v29
    case 127:
        var v32 float64 = -2.25
        _ = v32
    default:
        var v33 int8 = 0
        _ = v33
    }
    if (42 < 409) {
        _ = len(v3)
        fmt.Println(v31, v9, v12)
        if false {
            switch (v9 * 42) {
            case -629:
                var v34 uint64 = 18446744073709551615
                _ = v34
            case 4294967295:
                var v35 float32 = 1.5
                _ = v35
            default:
                var v36 string = ("" + v31)
                _ = v36
            }
            switch (v9 & v9) {
            case -2147483648:
                if (true && true) {
                    for fi37 := 0; fi37 < 15; fi37++ {
                        v30 -= (v30 % (((v30 & -2147483648)) | 1))
                        _ = fi37
                    }
                    if false {
                        v31 = v31
                        fmt.Println(v6, v31)
                        switch (v31 + v31) {
                        case "foo bar":
                            v2 = v1
                        case "\n\t":
                            var v38 uint16 = (v29 + v12)
                            _ = v38
                        case "hello":
                            if true {
                                if (-382 != 42) {
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                    v39 := GenS4{F0: (^(v29))}
                                    _ = v39
                                }
                                fmt.Println(v30, v2)
                                v40 := GenS3{F0: (v9 - v9), F1: false, F2: v9}
                                _ = v40
                            } else {
                                if len(v5) > 0 {
                                    v5[0] = v31
                                }
                            }
                        }
                    }
                    switch "\n\t" {
                    case "\n\t":
                        for fi41 := 0; fi41 < 7; fi41++ {
                            switch 9223372036854775807 {
                            case 310:
                                var v42 int16 = 286
                                _ = v42
                            case 968:
                                v43 := GenS3{F0: (fi41 * (v9 >> ((255) & 63))), F1: (false || (((((v4 / ((v4) | 1)) <= 42) && (42 > 1)) && true) || (true && true))), F2: (v9 & v9)}
                                _ = v43
                            default:
                                v10 = 42
                            }
                            var v44 uint8 = v1
                            fmt.Println(v44)
                            _ = fi41
                            _ = v44
                        }
                    case "foo bar":
                        if (!(!(false && false))) {
                            v6 -= int64(v30)
                            if false {
                                var v45 int64 = int64(v30)
                                _ = v45
                            }
                        }
                    }
                } else {
                    if len(v3) > 0 {
                        v3[0] = ((((v31 + "\n\t") + v31) + v31) + (v31 + v31))
                    }
                    if true {
                        fmt.Println(v29)
                    }
                }
            case -220:
                v29 = v12
            }
        } else {
            v30 -= (-(v10))
            var v46 uint64 = 209
            var v47 int32 = v30
            _ = v46
            _ = v47
        }
    }
    v2 = (v2 / (((v2 << ((2) & 63))) | 1))
    v29 = 399
    switch v31 {
    case "foo bar":
        v48 := GenS4{F0: 65535}
        _ = v48
    }
    var v49 uint = 255
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v49
    _ = v5
    _ = v6
    _ = v9
}
