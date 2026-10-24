package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 string; F1 int }
type GenS4 struct { F0 uint8; F1 uint64; F2 string }
type GenS5 struct { F0 uint32; F1 int32; F2 uint64 }
type GenS6 struct { F0 uint8; F1 uint32 }

func main() {
    v1 := make(map[string]int)
    v2 := []uint32{2147483647}
    var v3 int16 = 1
    var v4 bool = true
    v5 := v1[""]
    var v6 uint16 = 42
    var v7 float32 = (((0.1 / 1.0) - 1e10) / 1e10)
    fmt.Println(v6, v5, v4)
    for fi8 := 0; fi8 < 4; fi8++ {
        fmt.Println(v3, fi8)
        _ = len(v1)
        switch v5 {
        case 4294967295:
            for fi9 := 0; fi9 < 3; fi9++ {
                fmt.Println(v7)
                if true {
                    if ("hello" < (((("foo bar" + "hello") + "a") + (("foo bar" + "hello") + "")) + ("" + ((("" + "hello") + ("\n\t" + "\n\t")) + "\n\t")))) {
                        switch "foo bar" {
                        case "a":
                            v10 := GenS1{F0: (!v4)}
                            _ = v10
                        }
                        v11 := GenS1{F0: v4}
                        _ = v11
                    }
                } else {
                    v12 := GenS2{F0: 0}
                    switch "foo bar" {
                    case "a":
                        if len(v2) > 0 {
                            v2[0] = 23
                        }
                    }
                    _ = v12
                }
                _ = fi9
            }
        default:
            v13 := GenS2{F0: 818}
            _ = v13
        }
        _ = fi8
    }
    v14 := GenS3{F0: "a", F1: (v5 ^ 1)}
    switch v5 {
    case 847:
        v1[("foo bar" + ("" + (("hello" + ("a" + ("foo bar" + ("a" + (("hello" + (((("a" + "a") + "foo bar") + ("" + "foo bar")) + ("hello" + "foo bar"))) + (("hello" + "") + "a")))))) + ("" + ("a" + "a")))))] = v5
    case 447:
        v15 := GenS1{F0: ((("" >= "a") && v4) || (140 != uint8(v6)))}
        _ = v15
    case 2147483647:
        var v16 float64 = (0.0 * -2.25)
        _ = v16
    default:
        v17 := GenS1{F0: (42 >= 2147483648)}
        _ = v17
    }
    if false {
        v18 := GenS1{F0: true}
        _ = v18
    } else {
        var v19 uint64 = 0
        var v20 int8 = -82
        _ = v19
        _ = v20
    }
    v7 = v7
    if (42 >= 0) {
        v7 = (v7 - (v7 - v7))
        var v21 float32 = v7
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = v21
    } else {
        for fi22 := 0; fi22 < 5; fi22++ {
            v3 -= v3
            var v23 int64 = -9223372036854775808
            for fi24 := 0; fi24 < 9; fi24++ {
                fmt.Println(v7, fi22)
                switch "foo bar" {
                case "foo bar":
                    switch "foo bar" {
                    case "":
                        v25 := GenS4{F0: 185, F1: 1, F2: ""}
                        _ = v25
                    case "foo bar":
                        fmt.Println(v3)
                    default:
                        v1["a"] = fi22
                    }
                case "hello":
                    delete(v1, "hello")
                case "a":
                    v23 = (-(v23))
                }
                _ = fi24
            }
            _ = fi22
            _ = v23
        }
        var v26 uint8 = 2
        _ = v26
    }
    if true {
        v4 = ((-364 <= (v5 ^ (v5 & 32767))) || (!((false || true) && (42 < 255))))
    }
    var v27 int = 2147483647
    for fi28 := 0; fi28 < 20; fi28++ {
        switch ("a" + "foo bar") {
        case "hello":
            v7 = (v7 * (float32(1.5) / 0.1))
        default:
            switch (v27 % (((fi28 << ((4294967295) & 63))) | 1)) {
            case 354:
                var v29 float32 = float32(0.0)
                _ = v29
            case -846:
                var v30 uint16 = v6
                _ = v30
            case 9223372036854775807:
                v31 := GenS2{F0: 623}
                _ = v31
            default:
                for fi32 := 0; fi32 < 18; fi32++ {
                    fmt.Println(v3, fi32)
                    v5 = (fi28 & (fi28 + fi32))
                    fmt.Println(fi28)
                    _ = fi32
                }
            }
        }
        for fi33 := 0; fi33 < 17; fi33++ {
            _ = len(v1)
            _ = fi33
        }
        _ = fi28
    }
    delete(v1, "hello")
    if len(v2) > 0 {
        v2[0] = uint32(v6)
    }
    if (v4 || (v4 || ((-1 == 510) || true))) {
        switch (v27 % ((v27) | 1)) {
        case 4294967295:
            v6 += v6
        case -92:
            if v4 {
                fmt.Println(v27)
                if (!v4) {
                    v34 := GenS5{F0: 0, F1: -856, F2: 42}
                    switch ("\n\t" + "\n\t") {
                    case "\n\t":
                        fmt.Println(v6, v7)
                    case "foo bar":
                        if (!v4) {
                            for fi35 := 0; fi35 < 2; fi35++ {
                                if (127 != 255) {
                                    for fi36 := 0; fi36 < 19; fi36++ {
                                        if ((!v4) || false) {
                                            var v37 uint32 = 4294967295
                                            v27 -= (v5 - (fi35 >> ((18446744073709551615) & 63)))
                                            _ = v37
                                        } else {
                                            v5 -= (v5 | (fi35 % ((fi36) | 1)))
                                            v38 := GenS1{F0: v4}
                                            if ((^(v5)) <= 127) {
                                                switch ((("\n\t" + "\n\t") + "\n\t") + "") {
                                                case "foo bar":
                                                    fmt.Println(v3, v5)
                                                case "":
                                                    var v39 float32 = v7
                                                    _ = v39
                                                }
                                                delete(v1, ("\n\t" + "a"))
                                                _ = len(v1)
                                            }
                                            _ = v38
                                        }
                                        _ = fi36
                                    }
                                    var v40 float32 = ((0.1 / float32((-2.25 / ((((float64(3.14159) - 1.5) / 0.1) / -2.25) / 1e10)))) - (((v7 / (1e10 / v7)) * ((v7 / 1e10) * (v7 - (((v7 / ((v7 - ((v7 * (float32(float64(v7)) * 1.5)) + 0.0)) / 3.14159)) + 3.14159) + (float32(3.14159) * float32(1e10)))))) + float32(1e10)))
                                    _ = v40
                                }
                                v27 = fi35
                                _ = fi35
                            }
                            var v41 int16 = (v3 % ((976) | 1))
                            _ = v41
                        }
                    case "hello":
                        for fi42 := 0; fi42 < 14; fi42++ {
                            v27 -= 2147483647
                            _ = fi42
                        }
                    }
                    fmt.Println(v6, v3, v5)
                    _ = v34
                } else {
                    var v43 int32 = 42
                    _ = v43
                }
                v44 := GenS3{F0: "a", F1: (v27 + (v5 >> ((310) & 63)))}
                _ = v44
            }
        }
        v45 := GenS2{F0: 4294967295}
        _ = v45
    } else {
        v46 := GenS6{F0: 127, F1: 611}
        var v47 int8 = 100
        fmt.Println(v7)
        _ = v46
        _ = v47
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
