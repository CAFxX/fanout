package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64 }
type GenS2 struct { F0 uint32; F1 uint16 }
type GenS3 struct { F0 int8; F1 uint32 }
type GenS4 struct { F0 int16; F1 uint }
type GenS5 struct { F0 uint; F1 uint32; F2 uint32 }
type GenS6 struct { F0 bool; F1 bool; F2 string }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 int8; F1 string; F2 uint16 }

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[string]uint32)
    v3 := make(map[string]bool)
    v4 := []int{4294967295, 433, -132, 2147483647}
    var v5 uint32 = 4294967295
    switch ((("" + "foo bar") + "\n\t") + "a") {
    case "":
        if (-64 == 110) {
            switch 991 {
            case 4294967295:
                v6 := GenS1{F0: 42, F1: 363}
                _ = v6
            case -1:
                switch -9223372036854775808 {
                case -181:
                    if (false || false) {
                        v5 -= 4294967295
                        v2[""] = (v5 + (v5 & 4294967295))
                    } else {
                        switch "foo bar" {
                        case "":
                            v3[""] = false
                        case "hello":
                            v5 += v5
                        case "a":
                            var v7 int = 0
                            _ = v7
                        }
                    }
                case -2147483648:
                    for fi8 := 0; fi8 < 15; fi8++ {
                        if true {
                            var v9 uint64 = 2
                            _ = v9
                        } else {
                            if (!false) {
                                for fi10 := 0; fi10 < 17; fi10++ {
                                    v5 -= (v5 / (((^(v5))) | 1))
                                    switch fi10 {
                                    case -1:
                                        var v11 uint32 = (v5 >> ((255) & 63))
                                        _ = v11
                                    case -9223372036854775808:
                                        if (fi10 >= fi10) {
                                            var v12 int16 = 32767
                                            var v13 int16 = (v12 - v12)
                                            v14 := v2[""]
                                            _ = v12
                                            _ = v13
                                            _ = v14
                                        }
                                        var v15 int8 = -38
                                        _ = v15
                                    }
                                    if false {
                                        var v16 uint64 = 18446744073709551615
                                        var v17 int32 = 0
                                        _ = v16
                                        _ = v17
                                    }
                                    _ = fi10
                                }
                                var v18 uint16 = 32767
                                if false {
                                    v5 = v5
                                } else {
                                    fmt.Println(v5, fi8, v18)
                                    var v19 uint32 = v5
                                    fmt.Println(v19)
                                    _ = v19
                                }
                                _ = v18
                            }
                            fmt.Println(fi8, v5)
                        }
                        if (!(!((42 <= 255) || ("foo bar" > "foo bar")))) {
                            var v20 int32 = 42
                            for fi21 := 0; fi21 < 15; fi21++ {
                                v22 := GenS1{F0: 1000, F1: 880}
                                if (!(0 >= -1)) {
                                    switch 9223372036854775807 {
                                    case 2147483647:
                                        v20 = v20
                                    case 1:
                                        fmt.Println(fi8, fi8)
                                    default:
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                    }
                                }
                                _ = fi21
                                _ = v22
                            }
                            var v23 uint32 = v5
                            _ = v20
                            _ = v23
                        } else {
                            _ = len(v4)
                            for fi24 := 0; fi24 < 5; fi24++ {
                                v25 := GenS1{F0: 1, F1: 42}
                                v26 := GenS1{F0: 687, F1: 18446744073709551615}
                                _ = len(v3)
                                _ = fi24
                                _ = v25
                                _ = v26
                            }
                            if (true && false) {
                                for fi27 := 0; fi27 < 19; fi27++ {
                                    switch (("hello" + "") + "hello") {
                                    case "\n\t":
                                        v28 := GenS2{F0: 0, F1: 1}
                                        _ = v28
                                    }
                                    if len(v4) > 0 {
                                        v4[0] = (^(fi27))
                                    }
                                    _ = fi27
                                }
                                if (0.1 >= 1e10) {
                                    v29 := GenS3{F0: 77, F1: (v5 * v5)}
                                    _ = v29
                                } else {
                                    var v30 int16 = 0
                                    if (1 > 1) {
                                        v5 = (v5 & (^(v5)))
                                        if (!(!true)) {
                                            if (127 == 1) {
                                                if false {
                                                    for fi31 := 0; fi31 < 10; fi31++ {
                                                        v30 -= 32767
                                                        var v32 uint16 = 0
                                                        v32 = v32
                                                        _ = fi31
                                                        _ = v32
                                                    }
                                                    fmt.Println(v5)
                                                    for fi33 := 0; fi33 < 1; fi33++ {
                                                        v5 += v5
                                                        _ = fi33
                                                    }
                                                } else {
                                                    var v34 int16 = -32768
                                                    var v35 uint64 = 18446744073709551615
                                                    v36 := GenS1{F0: v30, F1: 18446744073709551615}
                                                    _ = v34
                                                    _ = v35
                                                    _ = v36
                                                }
                                            } else {
                                                if true {
                                                    if (1 != -3) {
                                                        fmt.Println(v30)
                                                    } else {
                                                        for fi37 := 0; fi37 < 3; fi37++ {
                                                            v38 := GenS4{F0: 32767, F1: 2}
                                                            _ = fi37
                                                            _ = v38
                                                        }
                                                        var v39 int64 = 42
                                                        _ = v39
                                                    }
                                                    v40 := GenS1{F0: (v30 >> ((2) & 63)), F1: 1}
                                                    _ = v40
                                                }
                                                v41 := GenS5{F0: 0, F1: (v5 << ((412) & 63)), F2: 2147483648}
                                                _ = v41
                                            }
                                        } else {
                                            fmt.Println(fi8)
                                            for fi42 := 0; fi42 < 13; fi42++ {
                                                switch "foo bar" {
                                                case "a":
                                                    switch fi42 {
                                                    case 4294967295:
                                                        v30 += v30
                                                    case 0:
                                                        if false {
                                                            v43 := GenS6{F0: (!false), F1: (1e10 == 0.0), F2: "hello"}
                                                            _ = len(v4)
                                                            _ = v43
                                                        }
                                                    case 9223372036854775807:
                                                        fmt.Println(v30, v30)
                                                    default:
                                                        if ("a" < ("hello" + "hello")) {
                                                            var v44 int16 = 681
                                                            if len(v4) >= 2 {
                                                                v4 = v4[1:len(v4)]
                                                            }
                                                            _ = v44
                                                        } else {
                                                            v45 := GenS7{F0: (v30 >> ((533) & 63))}
                                                            var v46 uint16 = 255
                                                            _ = v45
                                                            _ = v46
                                                        }
                                                    }
                                                case "hello":
                                                    v47 := v1[fi8]
                                                    _ = v47
                                                default:
                                                    var v48 uint8 = 1
                                                    _ = v48
                                                }
                                                v30 = 1
                                                _ = fi42
                                            }
                                        }
                                        v1[32767] = false
                                    }
                                    _ = v30
                                }
                                v5 = v5
                            }
                        }
                        fmt.Println(fi8, fi8)
                        _ = fi8
                    }
                case -1000:
                    switch "hello" {
                    case "a":
                        v5 = (v5 - v5)
                    }
                default:
                    if (!(0.1 == (1e10 / 1.5))) {
                        v49 := GenS7{F0: -32768}
                        var v50 bool = false
                        _ = v49
                        _ = v50
                    }
                }
            case 9223372036854775807:
                v51 := GenS8{F0: 102, F1: "", F2: 0}
                _ = v51
            }
            var v52 bool = true
            _ = v52
        }
    }
    v5 = v5
    if (125 == 255) {
        v5 += v5
        var v53 int32 = -780
        _ = v53
    }
    var v54 uint32 = v5
    v5 = v5
    v5 -= 4294967295
    if false {
        switch 0 {
        case 127:
            v54 = (^(v54))
        case -9223372036854775808:
            var v55 uint64 = 18446744073709551615
            _ = v55
        }
        for fi56 := 0; fi56 < 1; fi56++ {
            if (3.14159 <= (1.5 + -2.25)) {
                switch (fi56 & fi56) {
                case 0:
                    v57 := GenS8{F0: 42, F1: "\n\t", F2: 683}
                    _ = v57
                case 1:
                    v58 := GenS2{F0: 4294967295, F1: 42}
                    _ = v58
                case -550:
                    _ = len(v1)
                default:
                    for fi59 := 0; fi59 < 19; fi59++ {
                        fmt.Println(v54, fi56)
                        var v60 uint16 = 1
                        _ = fi59
                        _ = v60
                    }
                }
                if (("a" + ("hello" + "hello")) == ("" + "a")) {
                    _ = len(v1)
                    delete(v3, ("" + ""))
                    if false {
                        v5 = 65535
                    } else {
                        v4 = append(v4, (fi56 - 127))
                    }
                }
            }
            _ = fi56
        }
        v54 -= v54
    }
    for fi61 := 0; fi61 < 14; fi61++ {
        delete(v3, "")
        _ = fi61
    }
    fmt.Println(v54, v5)
    var v62 int64 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v54:", v54)
    fmt.Println("v62:", v62)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v54
    _ = v62
}
