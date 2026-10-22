package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int; F1 bool; F2 int }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int64; F1 uint64 }

func main() {
    var v1 uint32 = 0
    var v2 float32 = (1e10 / 1.5)
    v3 := []uint32{v1, v1, (^(v1)), v1}
    v4 := make(map[uint32]int)
    var v5 int64 = 985
    v2 = (1.5 - (0.1 / v2))
    v6 := v4[0]
    switch ("" + "") {
    case "\n\t":
        v7 := GenS1{F0: 2, F1: "\n\t"}
        _ = v7
    }
    v6 = (v6 * 42)
    switch v6 {
    case 2147483647:
        v8 := GenS2{F0: 1000}
        _ = v8
    case 0:
        for fi9 := 0; fi9 < 20; fi9++ {
            var v10 float32 = v2
            _ = fi9
            _ = v10
        }
    case 128:
        var v11 uint16 = 255
        _ = v11
    default:
        for fi12 := 0; fi12 < 19; fi12++ {
            v3 = append(v3, (v1 << ((4294967295) & 63)))
            v13 := v4[65535]
            _ = fi12
            _ = v13
        }
    }
    if true {
        switch (("\n\t" + "\n\t") + "hello") {
        case "":
            v1 = v1
        case "a":
            fmt.Println(v2, v6)
        case "foo bar":
            v6 += v6
        default:
            v14 := GenS1{F0: 361, F1: "hello"}
            _ = v14
        }
    } else {
        v1 -= (v1 + v1)
    }
    v3 = append(v3, v1)
    for fi15 := 0; fi15 < 4; fi15++ {
        if true {
            switch "" {
            case "\n\t":
                switch "a" {
                case "a":
                    fmt.Println(fi15, v5, v6)
                case "hello":
                    v16 := GenS3{F0: 65535}
                    _ = v16
                }
            case "":
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            case "hello":
                var v17 float64 = 3.14159
                _ = v17
            default:
                switch ("a" + "\n\t") {
                case "foo bar":
                    if false {
                        var v18 string = ("a" + ("foo bar" + ("hello" + ("hello" + "\n\t"))))
                        switch v6 {
                        case -9223372036854775808:
                            var v19 uint64 = 1
                            _ = v19
                        }
                        for fi20 := 0; fi20 < 5; fi20++ {
                            v21 := GenS4{F0: v6, F1: (false && (v18 >= (v18 + "a"))), F2: (v6 % ((32767) | 1))}
                            _ = fi20
                            _ = v21
                        }
                        _ = v18
                    } else {
                        v5 = 0
                    }
                default:
                    if (!((3.14159 / 1.5) <= ((float64(v2) - 1.5) / 1.0))) {
                        fmt.Println(v1, fi15)
                        switch "" {
                        case "\n\t":
                            var v22 bool = false
                            _ = v22
                        case "hello":
                            if ("foo bar" <= "\n\t") {
                                fmt.Println(v6, fi15)
                                v23 := GenS4{F0: (v6 ^ (v6 * fi15)), F1: (2147483647 >= 42), F2: 9223372036854775807}
                                _ = v23
                            }
                        default:
                            for fi24 := 0; fi24 < 1; fi24++ {
                                v25 := GenS5{F0: -128}
                                _ = fi24
                                _ = v25
                            }
                        }
                        _ = len(v3)
                    } else {
                        if true {
                            v1 = v1
                        }
                        fmt.Println(v1)
                        v6 += v6
                    }
                }
            }
        }
        v2 = v2
        _ = len(v3)
        _ = fi15
    }
    v26 := GenS6{F0: -697, F1: 0}
    var v27 int = v6
    fmt.Println(v5)
    if len(v3) > 0 {
        v3[0] = (^(v1))
    }
    for fi28 := 0; fi28 < 2; fi28++ {
        var v29 int8 = 42
        var v30 uint32 = v1
        _ = fi28
        _ = v29
        _ = v30
    }
    if false {
        v31 := GenS4{F0: v27, F1: true, F2: (^(v27))}
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        v5 = -1
        _ = v31
    }
    var v32 float64 = -2.25
    switch v6 {
    case 9223372036854775807:
        for fi33 := 0; fi33 < 20; fi33++ {
            v1 += (^(v1))
            if ((v1 % ((2) | 1)) == (v1 >> ((18446744073709551615) & 63))) {
                fmt.Println(fi33, v5, fi33)
            }
            _ = fi33
        }
    case 0:
        if false {
            for fi34 := 0; fi34 < 3; fi34++ {
                if false {
                    fmt.Println(v32, v27, fi34)
                } else {
                    if ((3.14159 + 0.0) <= -2.25) {
                        fmt.Println(v5)
                        v35 := GenS3{F0: 819}
                        v2 = 0.1
                        _ = v35
                    }
                }
                _ = fi34
            }
            switch (v6 << ((4294967295) & 63)) {
            case 2147483647:
                var v36 bool = true
                _ = v36
            default:
                for fi37 := 0; fi37 < 18; fi37++ {
                    var v38 float32 = v2
                    _ = fi37
                    _ = v38
                }
            }
            for fi39 := 0; fi39 < 2; fi39++ {
                v6 += (v6 | -9223372036854775808)
                _ = len(v3)
                if true {
                    switch (("hello" + ("" + "foo bar")) + (("a" + ("hello" + ((("foo bar" + "\n\t") + ("" + ("hello" + ""))) + "a"))) + (("hello" + ("a" + "foo bar")) + "hello"))) {
                    case "\n\t":
                        var v40 string = (("\n\t" + "hello") + "hello")
                        _ = v40
                    case "a":
                        var v41 float32 = float32(v32)
                        _ = v41
                    default:
                        var v42 int64 = int64(v1)
                        _ = v42
                    }
                    for fi43 := 0; fi43 < 5; fi43++ {
                        v2 = 1e10
                        _ = fi43
                    }
                }
                _ = fi39
            }
        }
    }
    for fi44 := 0; fi44 < 11; fi44++ {
        var v45 uint = 255
        switch (fi44 * v27) {
        case 42:
            if ("a" < "foo bar") {
                if false {
                    delete(v4, (v1 % (((v1 / (((v1 / ((v1) | 1))) | 1))) | 1)))
                    _ = len(v4)
                    var v46 uint32 = uint32(v45)
                    _ = v46
                } else {
                    if true {
                        switch (("foo bar" + "") + ("foo bar" + ("foo bar" + "a"))) {
                        case "":
                            delete(v4, v1)
                        default:
                            var v47 float64 = ((((1.5 / 1.0) + ((v32 / (float64(float32(v32)) + 1.5)) / v32)) / v32) / ((v32 - ((-2.25 * v32) - (0.0 + v32))) - (v32 / (1.5 * v32))))
                            _ = v47
                        }
                        for fi48 := 0; fi48 < 9; fi48++ {
                            var v49 uint64 = 186
                            _ = fi48
                            _ = v49
                        }
                    } else {
                        switch 4294967295 {
                        case 32767:
                            fmt.Println(fi44, v27)
                        case 42:
                            v50 := GenS1{F0: v45, F1: "\n\t"}
                            _ = v50
                        case 226:
                            fmt.Println(v6, v32, v2)
                        }
                        fmt.Println(v32)
                        v51 := v4[(v1 ^ (v1 + v1))]
                        _ = v51
                    }
                    v32 = (0.0 - (-2.25 + ((float64(-2.25) - v32) / (((1e10 / (v32 + v32)) / v32) - float64((0.1 / (v2 + 1e10)))))))
                }
                var v52 int8 = 42
                switch (v6 & (v6 >> ((v45) & 63))) {
                case 2147483647:
                    v1 -= v1
                default:
                    fmt.Println(v32, v6, v6)
                }
                _ = v52
            }
        case 229:
            v27 += (v27 + fi44)
        case -527:
            for fi53 := 0; fi53 < 11; fi53++ {
                switch fi53 {
                case 2147483647:
                    if (872 == 740) {
                        var v54 float32 = 0.0
                        v54 += 3.14159
                        v45 -= (v45 + (^(v45)))
                        _ = v54
                    } else {
                        switch "" {
                        case "foo bar":
                            v3 = append(v3, v1)
                        default:
                            if len(v3) > 0 {
                                v3[0] = (v1 & v1)
                            }
                        }
                    }
                case -614:
                    v27 = v27
                case 1:
                    v27 = 617
                }
                v1 -= 4294967295
                v55 := GenS2{F0: -1}
                _ = fi53
                _ = v55
            }
        }
        _ = fi44
        _ = v45
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
