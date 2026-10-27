package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int8; F1 uint8; F2 uint16 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 string; F1 bool; F2 uint }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 bool; F1 int16; F2 uint16 }

func main() {
    v1 := make(map[int]uint32)
    var v2 uint16 = 255
    var v3 uint64 = 118
    v4 := []string{"", ("hello" + ("a" + "hello")), ("" + ("a" + "\n\t")), "foo bar"}
    if false {
        var v5 uint8 = 2
        _ = v5
    } else {
        v3 += (v3 & (v3 % (((v3 | 1)) | 1)))
        v3 = v3
        _ = len(v1)
    }
    for fi6 := 0; fi6 < 11; fi6++ {
        if ((!(((-2.25 > ((1e10 * 1e10) * (3.14159 - (1e10 / (float32(3.14159) / 0.1))))) && false) && true)) || false) {
            v3 -= v3
            switch ("\n\t" + ("a" + "hello")) {
            case "hello":
                switch "a" {
                case "foo bar":
                    v3 += v3
                case "":
                    v2 = v2
                }
            case "foo bar":
                for fi7 := 0; fi7 < 1; fi7++ {
                    if (false || (!false)) {
                        v3 += (v3 + (v3 + v3))
                    } else {
                        fmt.Println(fi6)
                        v8 := GenS1{F0: 1}
                        _ = v8
                    }
                    if false {
                        if false {
                            if (uint(v2) < 42) {
                                var v9 float64 = 0.1
                                _ = v9
                            }
                        } else {
                            v10 := GenS2{F0: 64}
                            _ = v10
                        }
                        v2 -= (^(v2))
                        switch "foo bar" {
                        case "\n\t":
                            v2 -= 0
                        case "a":
                            v1[(fi6 % ((fi7) | 1))] = 424
                        case "":
                            var v11 int32 = 677
                            fmt.Println(fi6, v3)
                            var v12 int16 = 153
                            v13 := GenS2{F0: 0}
                            if (!((v2 % (((v2 >> ((uint(v3)) & 63))) | 1)) < uint16(v3))) {
                                if false {
                                    fmt.Println(v2, v12, fi7)
                                }
                                var v14 int16 = 1
                                _ = v14
                            }
                            _ = v11
                            _ = v12
                            _ = v13
                        }
                        fmt.Println(v3, v2)
                        switch "" {
                        case "hello":
                            var v15 int16 = 32767
                            v16 := GenS3{F0: 4294967295}
                            _ = v15
                            _ = v16
                        case "":
                            var v17 int32 = -2147483648
                            _ = v17
                        }
                        var v18 uint = 946
                        _ = v18
                    } else {
                        fmt.Println(fi6)
                        var v19 int32 = -184
                        _ = v19
                    }
                    _ = fi7
                }
            case "":
                var v20 float32 = 0.1
                _ = v20
            }
            v3 = (v3 * (v3 >> ((65535) & 63)))
        }
        fmt.Println(fi6, fi6)
        switch "hello" {
        case "foo bar":
            v2 = (v2 ^ v2)
        case "a":
            var v21 uint64 = (v3 + (v3 * 1))
            _ = v21
        }
        _ = fi6
    }
    switch ((("a" + "") + ("a" + "hello")) + ("\n\t" + "a")) {
    case "\n\t":
        if false {
            v22 := GenS2{F0: 74}
            _ = v22
        } else {
            v3 = (v3 - v3)
            var v23 int32 = 42
            _ = v23
        }
    case "foo bar":
        switch ("\n\t" + (("foo bar" + "") + "")) {
        case "foo bar":
            v24 := v1[-1]
            _ = v24
        case "\n\t":
            if true {
                if (!false) {
                    v2 = (^(v2))
                } else {
                    v25 := GenS4{F0: 42, F1: 0, F2: (v2 / (((v2 % ((v2) | 1))) | 1))}
                    fmt.Println(v2, v2)
                    _ = v25
                }
            }
        case "a":
            _ = len(v1)
        default:
            if true {
                _ = len(v1)
            } else {
                fmt.Println(v3, v2)
                var v26 int16 = 32767
                _ = v26
            }
        }
    }
    v27 := GenS3{F0: 503}
    fmt.Println(v2, v3)
    fmt.Println(v2, v2)
    for fi28 := 0; fi28 < 18; fi28++ {
        if (42 < 18446744073709551615) {
            _ = len(v1)
            for fi29 := 0; fi29 < 10; fi29++ {
                switch 32767 {
                case -637:
                    var v30 int16 = -861
                    _ = v30
                case -2147483648:
                    var v31 int8 = 0
                    _ = v31
                }
                _ = fi29
            }
            var v32 uint16 = 429
            _ = v32
        }
        for fi33 := 0; fi33 < 4; fi33++ {
            switch fi28 {
            case 2147483647:
                v34 := GenS3{F0: 65535}
                _ = v34
            }
            _ = fi33
        }
        delete(v1, fi28)
        _ = fi28
    }
    switch "" {
    case "\n\t":
        switch "a" {
        case "":
            for fi35 := 0; fi35 < 20; fi35++ {
                var v36 float32 = (3.14159 / 0.1)
                v2 = v2
                var v37 float32 = v36
                _ = fi35
                _ = v36
                _ = v37
            }
        case "hello":
            v38 := GenS5{F0: "a"}
            _ = v38
        }
    case "foo bar":
        v39 := GenS6{F0: ("a" + ("a" + "")), F1: (!(-2.25 > 3.14159)), F2: 2}
        _ = v39
    case "":
        var v40 int32 = 0
        _ = v40
    default:
        if true {
            var v41 uint16 = (v2 / (((v2 << ((42) & 63))) | 1))
            for fi42 := 0; fi42 < 18; fi42++ {
                if (!false) {
                    if true {
                        fmt.Println(v2, fi42, v41)
                        v41 = 65535
                        delete(v1, (^(fi42)))
                    }
                }
                v2 = 595
                _ = fi42
            }
            _ = v41
        } else {
            v43 := GenS7{F0: 1}
            _ = v43
        }
    }
    switch 4294967295 {
    case 32768:
        var v44 uint = 433
        _ = v44
    }
    var v45 uint8 = 2
    for fi46 := 0; fi46 < 7; fi46++ {
        v47 := GenS8{F0: true, F1: 711, F2: (v2 * v2)}
        if len(v4) > 0 {
            v4[0] = "hello"
        }
        switch fi46 {
        case 2147483647:
            v2 = v2
        case 32767:
            _ = len(v4)
        }
        _ = fi46
        _ = v47
    }
    v3 += (v3 & v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v45
}
