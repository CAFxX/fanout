package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16 }
type GenS2 struct { F0 int; F1 int }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int64; F1 int32 }
type GenS5 struct { F0 uint }

func main() {
    var v1 float32 = (-2.25 + -2.25)
    var v2 float32 = (3.14159 + 0.1)
    var v3 int = 0
    v4 := make(map[string]bool)
    v5 := []string{("\n\t" + "a"), "hello", ("\n\t" + "")}
    var v6 int = 1
    v2 -= 0.1
    v7 := GenS1{F0: 42, F1: 32767}
    for fi8 := 0; fi8 < 3; fi8++ {
        fmt.Println(fi8, fi8, v2)
        switch (("foo bar" + ("" + ("a" + ("foo bar" + "hello")))) + ("" + "\n\t")) {
        case "\n\t":
            if true {
                v9 := GenS2{F0: fi8, F1: fi8}
                switch "a" {
                case "hello":
                    fmt.Println(v1, fi8, fi8)
                case "foo bar":
                    var v10 uint32 = 65535
                    _ = v10
                }
                var v11 bool = false
                _ = v11
                _ = v9
            } else {
                v1 = float32(float64(v1))
                var v12 uint64 = 4294967295
                var v13 uint16 = 32767
                _ = v12
                _ = v13
            }
        default:
            if (("a" != "hello") && false) {
                v3 = (v3 / ((fi8) | 1))
            } else {
                switch "a" {
                case "\n\t":
                    for fi14 := 0; fi14 < 6; fi14++ {
                        fmt.Println(v3)
                        fmt.Println(fi8, v2, fi14)
                        _ = len(v4)
                        _ = fi14
                    }
                case "foo bar":
                    var v15 float32 = 3.14159
                    _ = v15
                case "hello":
                    if true {
                        v3 = fi8
                        if true {
                            fmt.Println(v2, v2)
                        } else {
                            if false {
                                var v16 int = (fi8 | 127)
                                for fi17 := 0; fi17 < 3; fi17++ {
                                    var v18 uint64 = 0
                                    var v19 float64 = 0.1
                                    v20 := GenS3{F0: 1}
                                    _ = fi17
                                    _ = v18
                                    _ = v19
                                    _ = v20
                                }
                                _ = v16
                            }
                            for fi21 := 0; fi21 < 3; fi21++ {
                                var v22 string = (("" + "a") + ("hello" + "a"))
                                var v23 float32 = (v2 / v2)
                                _ = fi21
                                _ = v22
                                _ = v23
                            }
                            _ = len(v4)
                        }
                    }
                default:
                    if false {
                        if len(v5) > 0 {
                            v5[0] = "\n\t"
                        }
                        v6 -= (-(fi8))
                    } else {
                        v24 := GenS3{F0: 1}
                        fmt.Println(fi8, v2)
                        _ = v24
                    }
                }
                var v25 uint16 = 2
                v26 := GenS2{F0: (v6 | v3), F1: (v6 % ((v6) | 1))}
                _ = v25
                _ = v26
            }
        }
        _ = fi8
    }
    fmt.Println(v3, v6)
    if true {
        if len(v5) > 0 {
            v5[0] = ""
        }
        var v27 uint64 = 4294967295
        _ = v27
    } else {
        v28 := GenS2{F0: (v6 / (((v3 / ((128) | 1))) | 1)), F1: (^(v6))}
        v1 += ((float32(1e10) * (v1 * -2.25)) / (v2 - 0.0))
        if false {
            v29 := GenS2{F0: -801, F1: v6}
            _ = v29
        }
        _ = v28
    }
    for fi30 := 0; fi30 < 15; fi30++ {
        v5 = append(v5, "a")
        for fi31 := 0; fi31 < 4; fi31++ {
            v32 := GenS3{F0: 2147483647}
            if (-123 <= -1) {
                v3 = fi30
            } else {
                fmt.Println(fi31, fi31, fi30)
            }
            v2 = float32(1.5)
            v2 += (v2 * float32((0.0 / (3.14159 - ((0.0 - 0.1) - -2.25)))))
            switch 32768 {
            case -556:
                delete(v4, "\n\t")
            case 1:
                v1 = (-2.25 - (3.14159 + v2))
            default:
                v33 := GenS4{F0: 79, F1: 380}
                _ = v33
            }
            _ = fi31
            _ = v32
        }
        if (317 < 312) {
            for fi34 := 0; fi34 < 7; fi34++ {
                var v35 float32 = ((v2 + (v2 * v1)) / ((3.14159 / v2) * v1))
                var v36 float64 = float64(v1)
                _ = fi34
                _ = v35
                _ = v36
            }
            if (2 < 1) {
                var v37 int32 = 1
                var v38 int64 = -1
                _ = v37
                _ = v38
            } else {
                if len(v5) > 0 {
                    v5[0] = "foo bar"
                }
            }
        }
        _ = fi30
    }
    var v39 bool = true
    var v40 uint32 = 646
    var v41 int64 = 0
    if (255 <= 136) {
        if (((v39 || (v39 || (709 < 42))) || v39) || (437 < 2)) {
            _ = len(v4)
        } else {
            switch v6 {
            case 42:
                v41 += v41
            case -9223372036854775808:
                var v42 int = v3
                _ = v42
            case 32768:
                fmt.Println(v2)
            default:
                var v43 uint8 = 128
                _ = v43
            }
            if (-2.25 > (-2.25 * (v2 + v2))) {
                v44 := GenS3{F0: 902}
                v45 := GenS5{F0: 687}
                v6 -= -2147483648
                _ = v44
                _ = v45
            } else {
                var v46 bool = (35 >= -81)
                v1 += (1e10 - (0.0 / 1e10))
                var v47 float32 = 1.5
                _ = v46
                _ = v47
            }
            fmt.Println(v39)
        }
        var v48 int64 = (v41 ^ -1)
        _ = v48
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v6
    _ = v7
}
