package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool; F2 int16 }
type GenS2 struct { F0 bool; F1 uint16 }
type GenS3 struct { F0 uint64; F1 uint; F2 uint8 }

func main() {
    var v1 int = 267
    v2 := make(map[int]int)
    var v3 int8 = 2
    var v4 uint64 = 1
    v5 := make(map[string]bool)
    v6 := []string{"foo bar", ("\n\t" + ("foo bar" + ("a" + "a")))}
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    var v7 bool = (("hello" + "") < ("" + "a"))
    fmt.Println(v7, v7)
    for fi8 := 0; fi8 < 15; fi8++ {
        v3 = -1
        v3 += -1
        if true {
            for fi9 := 0; fi9 < 1; fi9++ {
                if (false && v7) {
                    v5[("a" + "")] = (!(!v7))
                    if len(v6) > 0 {
                        v6[0] = "\n\t"
                    }
                } else {
                    fmt.Println(v1, fi9, fi8)
                }
                if false {
                    if (!v7) {
                        switch "hello" {
                        case "hello":
                            v4 += v4
                        case "":
                            switch ("foo bar" + "") {
                            case "a":
                                var v10 int64 = 0
                                _ = v10
                            case "hello":
                                v11 := GenS1{F0: -255, F1: ((v7 && (!v7)) || (!v7)), F2: 1000}
                                _ = v11
                            case "foo bar":
                                fmt.Println(fi9)
                            default:
                                var v12 uint = 0
                                _ = v12
                            }
                            v13 := GenS1{F0: 9223372036854775807, F1: ((^(v4)) <= v4), F2: 32767}
                            _ = v13
                        default:
                            if v7 {
                                v4 = v4
                                v4 += 930
                            } else {
                                _ = len(v5)
                                v14 := GenS1{F0: 9223372036854775807, F1: v7, F2: 562}
                                _ = v14
                            }
                        }
                    } else {
                        var v15 uint8 = 98
                        _ = v15
                    }
                    v16 := GenS1{F0: 9223372036854775807, F1: v7, F2: 933}
                    _ = v16
                }
                _ = fi9
            }
            if true {
                for fi17 := 0; fi17 < 2; fi17++ {
                    v1 = (v1 | (v1 * fi8))
                    _ = fi17
                }
                switch (v1 / ((fi8) | 1)) {
                case 2147483647:
                    delete(v2, 127)
                default:
                    v1 -= 127
                }
                for fi18 := 0; fi18 < 10; fi18++ {
                    delete(v5, "foo bar")
                    _ = fi18
                }
            } else {
                v19 := v2[v1]
                _ = v19
            }
            var v20 int = (v1 % ((722) | 1))
            _ = v20
        } else {
            if (722 < 1) {
                v6 = append(v6, ("a" + (("" + ("hello" + ("" + (("\n\t" + "a") + "")))) + "")))
                if (v7 || false) {
                    v3 += 127
                    for fi21 := 0; fi21 < 10; fi21++ {
                        var v22 uint8 = 146
                        v23 := v2[2147483647]
                        v22 += 209
                        if (-2147483648 < (fi21 & v23)) {
                            fmt.Println(fi21, v7, v23)
                            var v24 uint = 2
                            _ = v24
                        }
                        v4 += v4
                        switch (v23 % (((fi8 / ((-558) | 1))) | 1)) {
                        case 42:
                            v3 = (v3 / ((v3) | 1))
                        case -2147483648:
                            fmt.Println(v7)
                        default:
                            var v25 string = ("foo bar" + "foo bar")
                            v3 = v3
                            v26 := GenS2{F0: true, F1: 32767}
                            v4 += (v4 ^ v4)
                            v4 = v4
                            v6 = append(v6, v25)
                            _ = v25
                            _ = v26
                        }
                        _ = fi21
                        _ = v22
                        _ = v23
                    }
                    v27 := GenS1{F0: 42, F1: ("foo bar" <= ""), F2: -656}
                    _ = v27
                } else {
                    if (!((float32(3.14159) > 1e10) && (!v7))) {
                        for fi28 := 0; fi28 < 16; fi28++ {
                            if len(v6) > 0 {
                                v6[0] = ""
                            }
                            v3 += 67
                            if (42 == 65535) {
                                v4 = 69
                                v7 = (1000 >= int16(v4))
                            }
                            v1 += 127
                            _ = fi28
                        }
                    }
                }
                var v29 uint64 = v4
                _ = v29
            }
        }
        _ = fi8
    }
    var v30 int8 = v3
    switch "\n\t" {
    case "\n\t":
        if (9223372036854775807 == -9223372036854775808) {
            for fi31 := 0; fi31 < 17; fi31++ {
                if ((!v7) && (!v7)) {
                    if (!v7) {
                        v7 = (-128 <= v3)
                    } else {
                        switch v1 {
                        case 32768:
                            if true {
                                v32 := v2[v1]
                                _ = v32
                            } else {
                                var v33 bool = (!v7)
                                var v34 uint32 = 2
                                if len(v6) > 0 {
                                    v6[0] = ("hello" + ("a" + "a"))
                                }
                                _ = v33
                                _ = v34
                            }
                        case -1:
                            switch (("hello" + "") + "hello") {
                            case "foo bar":
                                for fi35 := 0; fi35 < 16; fi35++ {
                                    fmt.Println(v30)
                                    var v36 int16 = -975
                                    v4 += v4
                                    _ = fi35
                                    _ = v36
                                }
                            case "a":
                                v37 := GenS1{F0: -290, F1: v7, F2: 1000}
                                _ = v37
                            case "hello":
                                var v38 int32 = 42
                                _ = v38
                            }
                        default:
                            fmt.Println(v4)
                        }
                        fmt.Println(v1, v4)
                        v7 = ((-2.25 / 1e10) < 0.1)
                    }
                    v1 += 32767
                    if true {
                        for fi39 := 0; fi39 < 17; fi39++ {
                            var v40 uint16 = 2
                            _ = fi39
                            _ = v40
                        }
                    }
                } else {
                    v7 = ((false && ((((1.5 / 1.0) * float32(0.1)) - (1e10 + 3.14159)) != 0.1)) && v7)
                }
                _ = fi31
            }
            switch (v1 / ((v1) | 1)) {
            case 128:
                fmt.Println(v30)
            case 2147483647:
                switch (v1 * (v1 & (v1 >> ((1) & 63)))) {
                case 32768:
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                default:
                    var v41 int64 = -1
                    _ = v41
                }
            case -2147483648:
                fmt.Println(v30, v7)
            default:
                switch "hello" {
                case "a":
                    v7 = v7
                case "foo bar":
                    for fi42 := 0; fi42 < 14; fi42++ {
                        v30 += 26
                        for fi43 := 0; fi43 < 8; fi43++ {
                            var v44 int8 = v30
                            _ = fi43
                            _ = v44
                        }
                        v4 += (v4 >> ((411) & 63))
                        _ = fi42
                    }
                }
            }
        }
    case "hello":
        v4 = v4
    default:
        v1 = v1
    }
    v30 -= v30
    var v45 string = ((("a" + ("hello" + ("\n\t" + ("hello" + "foo bar")))) + "hello") + ("a" + ("foo bar" + "hello")))
    v4 -= 18446744073709551615
    var v46 uint8 = 5
    for fi47 := 0; fi47 < 18; fi47++ {
        for fi48 := 0; fi48 < 10; fi48++ {
            fmt.Println(v45, v4)
            var v49 float32 = 0.0
            switch fi47 {
            case 9223372036854775807:
                v50 := GenS3{F0: v4, F1: 65, F2: v46}
                if len(v6) > 0 {
                    v6[0] = (((((("a" + v45) + ((("foo bar" + "hello") + v45) + "foo bar")) + v45) + v45) + ((v45 + v45) + v45)) + v45)
                }
                fmt.Println(v4)
                v1 = v1
                _ = v50
            case 1:
                if ((v7 || v7) && true) {
                    fmt.Println(v7, v46)
                }
            }
            _ = fi48
            _ = v49
        }
        _ = fi47
    }
    v7 = v7
    var v51 uint32 = 65535
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v51:", v51)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v45
    _ = v46
    _ = v5
    _ = v51
    _ = v6
    _ = v7
}
