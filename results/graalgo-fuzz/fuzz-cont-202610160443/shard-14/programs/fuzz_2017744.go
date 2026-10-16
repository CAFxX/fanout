package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16 }
type GenS2 struct { F0 int64; F1 string; F2 uint }
type GenS3 struct { F0 int; F1 int16; F2 uint64 }
type GenS4 struct { F0 int8; F1 int16; F2 int16 }
type GenS5 struct { F0 uint; F1 int8 }
type GenS6 struct { F0 uint64; F1 uint32 }

func main() {
    var v1 uint16 = 883
    v2 := []string{"\n\t", "foo bar", "foo bar"}
    var v3 uint16 = (v1 - 748)
    var v4 int8 = -80
    v1 += (v3 | (v3 | (v1 + v3)))
    var v5 int64 = 0
    var v6 uint32 = 0
    switch ("\n\t" + "foo bar") {
    case "a":
        var v7 int8 = 100
        _ = v7
    default:
        switch "foo bar" {
        case "\n\t":
            var v8 int64 = (v5 + -622)
            _ = v8
        case "":
            for fi9 := 0; fi9 < 15; fi9++ {
                v3 = v1
                v10 := GenS1{F0: 474, F1: (v1 / ((v3) | 1))}
                _ = fi9
                _ = v10
            }
        case "hello":
            if false {
                for fi11 := 0; fi11 < 17; fi11++ {
                    _ = len(v2)
                    switch 128 {
                    case -2147483648:
                        for fi12 := 0; fi12 < 20; fi12++ {
                            var v13 string = ""
                            _ = fi12
                            _ = v13
                        }
                    case -440:
                        if false {
                            var v14 uint = 1
                            fmt.Println(v3)
                            var v15 int32 = 402
                            _ = v14
                            _ = v15
                        }
                    case 768:
                        if false {
                            fmt.Println(fi11, v3)
                            fmt.Println(v5, fi11)
                            fmt.Println(fi11)
                        } else {
                            v6 = (v6 % (((v6 << ((133) & 63))) | 1))
                        }
                    default:
                        switch fi11 {
                        case 42:
                            v1 += v1
                        case 1:
                            var v16 int8 = -104
                            _ = v16
                        default:
                            v1 = 32767
                        }
                    }
                    _ = fi11
                }
                fmt.Println(v1)
                var v17 uint8 = 127
                _ = v17
            } else {
                if ((!false) || true) {
                    var v18 int = -947
                    _ = v18
                } else {
                    switch "hello" {
                    case "a":
                        v3 = v1
                    case "":
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    case "hello":
                        if false {
                            for fi19 := 0; fi19 < 11; fi19++ {
                                for fi20 := 0; fi20 < 16; fi20++ {
                                    switch 1 {
                                    case 1:
                                        v5 -= (v5 / (((v5 << ((1) & 63))) | 1))
                                    case 6:
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                    case 9223372036854775807:
                                        v4 = v4
                                    }
                                    if true {
                                        fmt.Println(v6)
                                        _ = len(v2)
                                        if true {
                                            fmt.Println(v5)
                                        } else {
                                            switch "" {
                                            case "foo bar":
                                                v4 = (v4 & v4)
                                            case "\n\t":
                                                fmt.Println(v3, v1)
                                            }
                                            if ((v4 + (v4 / (((v4 & (v4 | -1))) | 1))) < 0) {
                                                fmt.Println(v6, v5, v5)
                                                if len(v2) > 0 {
                                                    v2[0] = "hello"
                                                }
                                            }
                                            if true {
                                                if ((true && ("" != "")) && true) {
                                                    if true {
                                                        if ((false || (false || (!true))) && (18446744073709551615 == 215)) {
                                                            fmt.Println(v1, v1)
                                                            var v21 float64 = 3.14159
                                                            v22 := GenS1{F0: 258, F1: 0}
                                                            _ = v21
                                                            _ = v22
                                                        }
                                                    }
                                                    if false {
                                                        v2 = append(v2, ("\n\t" + ("a" + (("a" + "hello") + "a"))))
                                                    }
                                                    switch "foo bar" {
                                                    case "foo bar":
                                                        fmt.Println(fi20)
                                                    default:
                                                        var v23 int16 = 42
                                                        _ = v23
                                                    }
                                                }
                                                if ((!true) || true) {
                                                    var v24 uint16 = (v1 - (v3 << ((255) & 63)))
                                                    _ = v24
                                                }
                                                switch "" {
                                                case "a":
                                                    v25 := GenS1{F0: 173, F1: v3}
                                                    _ = v25
                                                }
                                            }
                                        }
                                        if (32768 == fi19) {
                                            v3 = 321
                                            var v26 float64 = 1.5
                                            _ = v26
                                        }
                                        fmt.Println(fi20, v4, fi19)
                                    }
                                    v4 += int8(v5)
                                    if (!true) {
                                        fmt.Println(fi20)
                                        var v27 float32 = 0.1
                                        _ = v27
                                    } else {
                                        if len(v2) > 0 {
                                            v2[0] = "foo bar"
                                        }
                                    }
                                    switch "a" {
                                    case "":
                                        v28 := GenS2{F0: v5, F1: "a", F2: 42}
                                        _ = v28
                                    case "\n\t":
                                        var v29 string = ("" + ((("" + "a") + "hello") + ("\n\t" + ((("foo bar" + ("hello" + "\n\t")) + "\n\t") + "a"))))
                                        _ = v29
                                    case "hello":
                                        if true {
                                            var v30 float64 = 1.5
                                            _ = v30
                                        }
                                    default:
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                    }
                                    v3 = (v1 % ((v1) | 1))
                                    _ = fi20
                                }
                                _ = fi19
                            }
                            var v31 string = ("foo bar" + "a")
                            var v32 uint = 65535
                            _ = v31
                            _ = v32
                        } else {
                            if true {
                                fmt.Println(v5, v5, v6)
                                switch (("\n\t" + ("hello" + "")) + "a") {
                                case "a":
                                    _ = len(v2)
                                case "":
                                    for fi33 := 0; fi33 < 20; fi33++ {
                                        var v34 int64 = v5
                                        if len(v2) > 0 {
                                            v2[0] = "foo bar"
                                        }
                                        _ = fi33
                                        _ = v34
                                    }
                                case "\n\t":
                                    v4 = int8(v5)
                                default:
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                }
                            } else {
                                var v35 uint = 1
                                _ = v35
                            }
                        }
                    }
                    v3 = 32767
                    for fi36 := 0; fi36 < 20; fi36++ {
                        var v37 string = "a"
                        switch -149 {
                        case 32767:
                            v38 := GenS3{F0: 42, F1: 1, F2: 42}
                            _ = v38
                        case -9223372036854775808:
                            for fi39 := 0; fi39 < 4; fi39++ {
                                v40 := GenS2{F0: 42, F1: "a", F2: 42}
                                var v41 int8 = (v4 * v4)
                                _ = fi39
                                _ = v40
                                _ = v41
                            }
                        }
                        if true {
                            v6 += (v6 << ((65535) & 63))
                        } else {
                            if (!true) {
                                fmt.Println(v4)
                            }
                        }
                        _ = fi36
                        _ = v37
                    }
                }
                var v42 uint64 = 18446744073709551615
                if false {
                    var v43 uint8 = uint8(v1)
                    _ = v43
                }
                _ = v42
            }
        }
    }
    var v44 int32 = -2147483648
    if false {
        fmt.Println(v4)
        if ((9223372036854775807 != 614) && true) {
            var v45 uint32 = uint32(v3)
            _ = v45
        }
        v4 = v4
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v46 := GenS4{F0: 1, F1: 0, F2: -1}
    for fi47 := 0; fi47 < 1; fi47++ {
        for fi48 := 0; fi48 < 14; fi48++ {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            var v49 bool = ((v4 | (-(v4))) != v4)
            v6 = (v6 | 2147483647)
            var v50 int = 220
            fmt.Println(fi47)
            _ = fi48
            _ = v49
            _ = v50
        }
        fmt.Println(v3)
        _ = fi47
    }
    var v51 uint64 = 0
    switch "foo bar" {
    case "a":
        v52 := GenS5{F0: 255, F1: v4}
        _ = v52
    }
    _ = len(v2)
    v4 = 42
    v53 := GenS6{F0: (v51 % (((v51 - (v51 | 715))) | 1)), F1: (v6 | (^(v6)))}
    v5 += (^(v5))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v46
    _ = v5
    _ = v51
    _ = v53
    _ = v6
}
