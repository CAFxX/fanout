package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint; F2 uint32 }
type GenS2 struct { F0 string; F1 uint64 }
type GenS3 struct { F0 uint64; F1 uint32; F2 int32 }
type GenS4 struct { F0 bool; F1 int }
type GenS5 struct { F0 uint32; F1 uint32 }
type GenS6 struct { F0 int8; F1 string; F2 int8 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 int32; F1 string }
type GenS9 struct { F0 uint8 }
type GenS10 struct { F0 uint16; F1 uint; F2 uint32 }

func main() {
    v1 := []uint32{2, 0, 743, 42}
    var v2 bool = true
    v3 := make(map[string]int)
    var v4 string = "\n\t"
    v5 := GenS1{F0: 538, F1: 1, F2: 2}
    v6 := GenS1{F0: 2147483648, F1: 707, F2: 4294967295}
    if len(v1) > 0 {
        v1[0] = 435
    }
    var v7 int64 = 1
    v2 = v2
    switch (v4 + (v4 + v4)) {
    case "a":
        for fi8 := 0; fi8 < 18; fi8++ {
            v9 := GenS2{F0: "\n\t", F1: 18446744073709551615}
            _ = fi8
            _ = v9
        }
    case "\n\t":
        fmt.Println(v2, v4, v4)
    case "foo bar":
        v10 := GenS1{F0: 65535, F1: 2, F2: 495}
        _ = v10
    }
    v7 -= -755
    for fi11 := 0; fi11 < 17; fi11++ {
        fmt.Println(v2)
        fmt.Println(v4)
        _ = fi11
    }
    for fi12 := 0; fi12 < 13; fi12++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        var v13 uint32 = 313
        var v14 uint = 480
        _ = fi12
        _ = v13
        _ = v14
    }
    fmt.Println(v2, v4)
    v15 := GenS1{F0: 587, F1: 4294967295, F2: 4294967295}
    v4 = v4
    v2 = (!((v2 || (((float64((0.1 + 0.0)) < 1e10) || (23 >= 2)) || false)) && (true || v2)))
    switch 32767 {
    case 718:
        v7 = (v7 * (v7 / ((v7) | 1)))
    case -10:
        v16 := GenS3{F0: 4294967295, F1: 4294967295, F2: 0}
        _ = v16
    }
    v7 -= v7
    for fi17 := 0; fi17 < 18; fi17++ {
        switch "\n\t" {
        case "hello":
            v7 = -9223372036854775808
        case "":
            if (!true) {
                for fi18 := 0; fi18 < 18; fi18++ {
                    fmt.Println(v4, v7)
                    var v19 uint8 = 255
                    fmt.Println(v19, v7)
                    _ = fi18
                    _ = v19
                }
                for fi20 := 0; fi20 < 5; fi20++ {
                    if v2 {
                        _ = len(v1)
                        if ((fi20 > 2147483647) && ((v2 && (255 <= 255)) && (712 != 967))) {
                            v2 = (!false)
                            switch (fi17 / ((fi20) | 1)) {
                            case -1:
                                var v21 uint8 = 160
                                _ = v21
                            case 4294967295:
                                var v22 int = 9223372036854775807
                                _ = v22
                            }
                            switch "a" {
                            case "":
                                if v2 {
                                    v7 = -98
                                    var v23 uint16 = 1
                                    v24 := GenS3{F0: 621, F1: 2147483647, F2: -600}
                                    _ = v23
                                    _ = v24
                                }
                            case "\n\t":
                                v1 = append(v1, 4294967295)
                            case "hello":
                                v7 = 1
                            default:
                                if v2 {
                                    v25 := GenS4{F0: v2, F1: (^(fi20))}
                                    _ = v25
                                } else {
                                    switch v4 {
                                    case "hello":
                                        switch "foo bar" {
                                        case "hello":
                                            v26 := GenS5{F0: 2147483647, F1: 2147483647}
                                            _ = v26
                                        default:
                                            v27 := GenS6{F0: -128, F1: v4, F2: 1}
                                            _ = v27
                                        }
                                    case "\n\t":
                                        var v28 int = -2147483648
                                        _ = v28
                                    case "a":
                                        if v2 {
                                            fmt.Println(fi20, v2)
                                        }
                                    default:
                                        v29 := GenS7{F0: 519}
                                        _ = v29
                                    }
                                    v3[((v4 + ((v4 + v4) + v4)) + v4)] = (fi20 / ((fi17) | 1))
                                }
                                if (!v2) {
                                    v4 = ""
                                } else {
                                    v2 = (v2 && v2)
                                    var v30 uint32 = 42
                                    _ = v30
                                }
                                var v31 float32 = (float32(0.1) + (float32(0.0) / 3.14159))
                                _ = v31
                            }
                        } else {
                            delete(v3, v4)
                            var v32 int16 = -32768
                            _ = v32
                        }
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    }
                    if len(v1) > 0 {
                        v1[0] = 886
                    }
                    var v33 string = (v4 + v4)
                    _ = fi20
                    _ = v33
                }
                if true {
                    if (839 > 0) {
                        switch (("" + ("foo bar" + (("foo bar" + v4) + (v4 + ("hello" + v4))))) + v4) {
                        case "hello":
                            for fi34 := 0; fi34 < 1; fi34++ {
                                var v35 int = (-(fi34))
                                v36 := GenS2{F0: (("a" + "") + "foo bar"), F1: 2}
                                var v37 int8 = -1
                                _ = fi34
                                _ = v35
                                _ = v36
                                _ = v37
                            }
                        }
                        var v38 float64 = (0.1 + ((0.0 / ((0.1 - 1.5) + ((3.14159 / (float64(float32(float64((0.0 + (3.14159 / (-2.25 + 0.1)))))) / (0.1 + 3.14159))) - 0.1))) - float64((-2.25 / ((-2.25 / float32(float64(float32(1e10)))) - (((0.1 - (0.0 - float32((1.5 / 1e10)))) * 0.1) - (-2.25 + (1e10 - 1.5))))))))
                        _ = v38
                    } else {
                        v39 := GenS8{F0: -912, F1: ""}
                        if len(v1) > 0 {
                            v1[0] = 2147483648
                        }
                        _ = v39
                    }
                }
            } else {
                switch "foo bar" {
                case "\n\t":
                    v40 := GenS9{F0: 37}
                    _ = v40
                }
                for fi41 := 0; fi41 < 10; fi41++ {
                    switch ((v4 + ((v4 + ("hello" + ("hello" + (v4 + (("" + (v4 + ("a" + v4))) + (("\n\t" + v4) + v4)))))) + "a")) + v4) {
                    case "hello":
                        var v42 int = -522
                        fmt.Println(v7, fi17, v4)
                        _ = v42
                    case "":
                        switch v4 {
                        case "\n\t":
                            v7 = v7
                        }
                    default:
                        if (0 == 2147483648) {
                            switch v4 {
                            case "":
                                var v43 uint8 = 0
                                _ = v43
                            case "foo bar":
                                var v44 int64 = v7
                                _ = v44
                            }
                        } else {
                            v7 += (-(v7))
                            v4 = "foo bar"
                        }
                    }
                    v45 := GenS10{F0: 32767, F1: 8, F2: 2147483647}
                    var v46 uint16 = 32767
                    _ = fi41
                    _ = v45
                    _ = v46
                }
            }
        case "foo bar":
            var v47 uint16 = 908
            _ = v47
        }
        v1 = append(v1, 493)
        _ = fi17
    }
    var v48 bool = true
    v2 = v2
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v5
    _ = v6
    _ = v7
}
