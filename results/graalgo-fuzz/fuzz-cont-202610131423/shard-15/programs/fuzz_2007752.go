package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64; F2 uint32 }
type GenS2 struct { F0 int32; F1 bool; F2 int16 }
type GenS3 struct { F0 int64; F1 int8 }
type GenS4 struct { F0 int64; F1 uint8; F2 int16 }
type GenS5 struct { F0 bool; F1 int16; F2 uint64 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 int64; F1 uint }

func main() {
    v1 := []uint32{42, 65535, 0, 65535}
    v2 := []uint32{42}
    v3 := make(map[string]string)
    v4 := []uint32{2147483647, 496, 42, 1}
    var v5 int = 1
    if true {
        v5 = v5
        switch (("a" + ("" + ("" + "a"))) + "foo bar") {
        case "hello":
            v5 += (^(v5))
        default:
            v6 := GenS1{F0: "foo bar", F1: 2, F2: 205}
            _ = v6
        }
        switch (v5 % ((v5) | 1)) {
        case 516:
            var v7 uint = 18446744073709551615
            _ = v7
        case 771:
            v3["hello"] = ("foo bar" + ("foo bar" + "a"))
        }
    }
    for fi8 := 0; fi8 < 3; fi8++ {
        if (278 == 904) {
            switch "" {
            case "foo bar":
                if (42 < -1) {
                    fmt.Println(v5, v5)
                    fmt.Println(v5)
                } else {
                    switch fi8 {
                    case 0:
                        v5 = (-(fi8))
                    case 2147483647:
                        switch v5 {
                        case 2147483647:
                            var v9 int = 9223372036854775807
                            _ = v9
                        }
                    case -52:
                        v5 += 1
                    }
                    v10 := v3["a"]
                    _ = v10
                }
            case "a":
                v5 = v5
            case "\n\t":
                var v11 uint8 = 0
                _ = v11
            default:
                if (true || true) {
                    switch "hello" {
                    case "foo bar":
                        v12 := GenS2{F0: -639, F1: (2 == 216), F2: -32768}
                        _ = v12
                    }
                } else {
                    v5 -= 880
                    switch "hello" {
                    case "hello":
                        fmt.Println(v5)
                    case "\n\t":
                        for fi13 := 0; fi13 < 2; fi13++ {
                            var v14 uint32 = 2147483647
                            _ = fi13
                            _ = v14
                        }
                    case "a":
                        var v15 string = ("a" + "foo bar")
                        _ = v15
                    default:
                        switch 42 {
                        case 128:
                            if (0 > 2) {
                                var v16 int = v5
                                var v17 int64 = 0
                                _ = v16
                                _ = v17
                            } else {
                                var v18 int16 = -827
                                _ = v18
                            }
                        case 180:
                            v5 -= -172
                        case -9223372036854775808:
                            for fi19 := 0; fi19 < 16; fi19++ {
                                v5 = fi8
                                var v20 uint8 = 42
                                switch ("a" + "hello") {
                                case "foo bar":
                                    v20 -= v20
                                case "":
                                    switch -406 {
                                    case -9223372036854775808:
                                        if (!((int16(v20) != -230) || false)) {
                                            v1 = append(v1, 4294967295)
                                        } else {
                                            switch -2147483648 {
                                            case 0:
                                                var v21 float64 = float64((3.14159 / float32(float64(((0.1 / 3.14159) * 3.14159)))))
                                                _ = v21
                                            default:
                                                v22 := GenS3{F0: 9223372036854775807, F1: 120}
                                                _ = v22
                                            }
                                            var v23 int8 = -1
                                            _ = v23
                                        }
                                    case 32768:
                                        var v24 float64 = 0.0
                                        _ = v24
                                    case 2147483647:
                                        var v25 float32 = -2.25
                                        _ = v25
                                    default:
                                        switch (^(v5)) {
                                        case 397:
                                            var v26 int8 = 0
                                            _ = v26
                                        case -9223372036854775808:
                                            fmt.Println(v5)
                                        default:
                                            v20 += v20
                                        }
                                    }
                                case "a":
                                    var v27 int32 = 42
                                    _ = v27
                                }
                                _ = fi19
                                _ = v20
                            }
                        default:
                            switch ("hello" + "a") {
                            case "\n\t":
                                for fi28 := 0; fi28 < 3; fi28++ {
                                    switch (fi8 * (v5 << ((688) & 63))) {
                                    case 32768:
                                        fmt.Println(fi8)
                                    case 4294967295:
                                        var v29 uint64 = 1
                                        _ = v29
                                    case 584:
                                        if false {
                                            v30 := GenS4{F0: 218, F1: 42, F2: 1}
                                            _ = v30
                                        } else {
                                            if true {
                                                if true {
                                                    var v31 int8 = -128
                                                    if false {
                                                        fmt.Println(v5, v5)
                                                    } else {
                                                        v32 := GenS5{F0: (!true), F1: -32768, F2: 738}
                                                        v5 -= (fi28 * 2147483647)
                                                        _ = v32
                                                    }
                                                    var v33 bool = ((0.0 - -2.25) > (3.14159 / (1.5 / 0.1)))
                                                    _ = v31
                                                    _ = v33
                                                } else {
                                                    v5 += (fi28 & (fi8 & v5))
                                                }
                                            }
                                        }
                                    default:
                                        var v34 int32 = 1
                                        _ = v34
                                    }
                                    _ = fi28
                                }
                            case "foo bar":
                                var v35 string = ("foo bar" + ("\n\t" + "foo bar"))
                                _ = v35
                            default:
                                var v36 bool = (-1 < -191)
                                _ = v36
                            }
                        }
                    }
                }
            }
            v37 := GenS1{F0: "a", F1: 0, F2: 864}
            if (!true) {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                for fi38 := 0; fi38 < 14; fi38++ {
                    var v39 int16 = 32767
                    fmt.Println(fi38, v39)
                    _ = fi38
                    _ = v39
                }
                var v40 int64 = -749
                _ = v40
            } else {
                switch "\n\t" {
                case "hello":
                    fmt.Println(v5, v5)
                }
            }
            _ = v37
        }
        _ = fi8
    }
    switch (^(v5)) {
    case 0:
        v41 := GenS6{F0: 624}
        _ = v41
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v42 float64 = 1e10
    for fi43 := 0; fi43 < 6; fi43++ {
        if false {
            if false {
                for fi44 := 0; fi44 < 12; fi44++ {
                    fmt.Println(fi43)
                    _ = fi44
                }
            }
            var v45 float64 = 1e10
            v46 := GenS1{F0: "foo bar", F1: 4294967295, F2: 603}
            _ = v45
            _ = v46
        }
        switch "foo bar" {
        case "foo bar":
            for fi47 := 0; fi47 < 13; fi47++ {
                fmt.Println(v5, v5)
                v3["hello"] = (("\n\t" + (("" + ("foo bar" + "\n\t")) + ("\n\t" + "foo bar"))) + "\n\t")
                _ = fi47
            }
        case "\n\t":
            if (!(65535 < 255)) {
                if false {
                    delete(v3, "foo bar")
                }
                v42 -= (v42 / -2.25)
                v5 += v5
            }
        }
        _ = fi43
    }
    for fi48 := 0; fi48 < 15; fi48++ {
        fmt.Println(v42, v5)
        _ = fi48
    }
    if true {
        if (0 >= 65535) {
            v42 = v42
        } else {
            v42 = float64(0.0)
            fmt.Println(v5)
        }
    } else {
        switch ("hello" + "foo bar") {
        case "":
            var v49 int = v5
            _ = v49
        case "\n\t":
            _ = len(v4)
        }
        switch "a" {
        case "foo bar":
            v3["\n\t"] = ((("" + "") + "foo bar") + "foo bar")
        case "\n\t":
            for fi50 := 0; fi50 < 9; fi50++ {
                v42 -= ((v42 + 1.5) / -2.25)
                _ = len(v3)
                _ = fi50
            }
        case "hello":
            delete(v3, (("" + "hello") + ""))
        }
        v5 -= v5
    }
    fmt.Println(v42)
    v51 := GenS7{F0: 1, F1: 0}
    _ = len(v3)
    v5 = -9223372036854775808
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v5
    _ = v51
}
