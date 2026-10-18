package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64 }
type GenS2 struct { F0 uint16; F1 bool; F2 int }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 int16; F1 bool; F2 uint32 }

func main() {
    v1 := []string{"\n\t"}
    v2 := make(map[uint32]string)
    v3 := []uint32{192, 2147483647, 42}
    v4 := []string{}
    var v5 uint64 = 2
    v6 := GenS1{F0: 23, F1: 880}
    for fi7 := 0; fi7 < 20; fi7++ {
        fmt.Println(v5, fi7)
        var v8 int16 = -300
        _ = fi7
        _ = v8
    }
    v9 := GenS2{F0: 65535, F1: (false || false), F2: 32768}
    if (!false) {
        v10 := GenS3{F0: 42}
        v5 -= 151
        fmt.Println(v5)
        _ = v10
    }
    if false {
        delete(v2, 65535)
    }
    switch 128 {
    case -690:
        switch -907 {
        case -801:
            v5 += (v5 & v5)
        case 42:
            v5 = v5
        case 786:
            for fi11 := 0; fi11 < 16; fi11++ {
                switch "foo bar" {
                case "a":
                    if false {
                        for fi12 := 0; fi12 < 13; fi12++ {
                            if false {
                                var v13 float64 = 3.14159
                                _ = v13
                            }
                            v5 -= (v5 * v5)
                            _ = fi12
                        }
                        var v14 uint64 = v5
                        v15 := GenS1{F0: 0, F1: v5}
                        _ = v14
                        _ = v15
                    } else {
                        switch ("hello" + "foo bar") {
                        case "hello":
                            fmt.Println(v5)
                        default:
                            for fi16 := 0; fi16 < 16; fi16++ {
                                fmt.Println(fi16)
                                var v17 int16 = 246
                                var v18 float64 = (1.5 / (0.0 - -2.25))
                                switch 42 {
                                case 0:
                                    var v19 string = "foo bar"
                                    _ = v19
                                case 672:
                                    v2[0] = "\n\t"
                                case -1:
                                    var v20 float32 = 1e10
                                    _ = v20
                                default:
                                    if false {
                                        v21 := GenS1{F0: 2, F1: v5}
                                        _ = v21
                                    } else {
                                        var v22 int64 = 648
                                        var v23 uint = 255
                                        _ = v22
                                        _ = v23
                                    }
                                }
                                _ = fi16
                                _ = v17
                                _ = v18
                            }
                        }
                    }
                case "hello":
                    v24 := GenS4{F0: v5}
                    _ = v24
                case "":
                    v5 -= v5
                default:
                    if len(v4) > 0 {
                        v4[0] = ""
                    }
                }
                if (!false) {
                    if ((^(v5)) < (v5 % ((v5) | 1))) {
                        _ = len(v2)
                        switch -76 {
                        case 127:
                            var v25 uint16 = 255
                            _ = v25
                        case 2147483647:
                            v5 += v5
                        }
                    }
                    for fi26 := 0; fi26 < 12; fi26++ {
                        if ((4294967295 >= 42) || (1e10 >= -2.25)) {
                            fmt.Println(v5, fi11)
                            _ = len(v2)
                        }
                        v2[uint32(v5)] = ("a" + (("a" + ("\n\t" + ("foo bar" + "foo bar"))) + "a"))
                        v5 += (v5 << ((91) & 63))
                        _ = fi26
                    }
                    if (42 < uint16(v5)) {
                        delete(v2, 65535)
                        if (-2.25 != 0.1) {
                            if (-1 > -1) {
                                if true {
                                    for fi27 := 0; fi27 < 8; fi27++ {
                                        if (!false) {
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                            v5 = v5
                                        }
                                        _ = fi27
                                    }
                                } else {
                                    v5 = v5
                                }
                            }
                            if true {
                                var v28 float32 = 1e10
                                _ = v28
                            } else {
                                switch (fi11 / (((fi11 ^ (fi11 & fi11))) | 1)) {
                                case 42:
                                    switch "a" {
                                    case "\n\t":
                                        var v29 int8 = 85
                                        _ = v29
                                    case "":
                                        v30 := GenS4{F0: v5}
                                        _ = v30
                                    default:
                                        if len(v4) > 0 {
                                            v4[0] = ("hello" + "\n\t")
                                        }
                                    }
                                case 128:
                                    v5 = 18446744073709551615
                                default:
                                    switch (fi11 | 128) {
                                    case 32767:
                                        if (2 != 1) {
                                            for fi31 := 0; fi31 < 14; fi31++ {
                                                if (!true) {
                                                    if true {
                                                        var v32 float32 = 1e10
                                                        v32 += (-2.25 / 1.5)
                                                        if true {
                                                            v5 = (v5 >> ((65535) & 63))
                                                            var v33 int = fi31
                                                            _ = v33
                                                        }
                                                        _ = v32
                                                    } else {
                                                        switch (fi11 - (fi11 - fi31)) {
                                                        case 32767:
                                                            fmt.Println(fi11, v5, fi31)
                                                        }
                                                    }
                                                    v34 := GenS1{F0: 2147483647, F1: v5}
                                                    _ = v34
                                                } else {
                                                    if (0.0 >= 1e10) {
                                                        if (false && ((!(2147483647 >= -1)) || (true || false))) {
                                                            v5 -= v5
                                                            var v35 uint32 = 42
                                                            _ = v35
                                                        }
                                                        delete(v2, 42)
                                                    } else {
                                                        if false {
                                                            v36 := GenS5{F0: false}
                                                            _ = v36
                                                        }
                                                        var v37 int16 = 812
                                                        _ = v37
                                                    }
                                                }
                                                v5 = (v5 >> ((0) & 63))
                                                _ = fi31
                                            }
                                            if (255 >= 184) {
                                                v5 = 970
                                                var v38 uint = 255
                                                _ = v38
                                            }
                                            v39 := v2[1]
                                            _ = v39
                                        } else {
                                            if (2 > (v5 << ((4294967295) & 63))) {
                                                v40 := GenS3{F0: 912}
                                                v5 += (v5 + v5)
                                                var v41 int64 = 305
                                                _ = v40
                                                _ = v41
                                            }
                                            fmt.Println(v5, v5)
                                        }
                                    case 127:
                                        var v42 string = "a"
                                        _ = v42
                                    }
                                }
                                if len(v4) > 0 {
                                    v4[0] = ((("\n\t" + "a") + ("\n\t" + ((("" + "") + "hello") + ""))) + "")
                                }
                                v5 -= (v5 - (v5 << ((655) & 63)))
                            }
                        }
                        if len(v4) > 0 {
                            v4[0] = "hello"
                        }
                    } else {
                        v5 = 680
                        switch fi11 {
                        case 4294967295:
                            v43 := GenS2{F0: 255, F1: ((v5 | (v5 - v5)) == 804), F2: fi11}
                            _ = v43
                        case 439:
                            var v44 int64 = 9223372036854775807
                            _ = v44
                        }
                    }
                }
                _ = fi11
            }
        }
    case 128:
        for fi45 := 0; fi45 < 13; fi45++ {
            var v46 uint16 = 1
            _ = fi45
            _ = v46
        }
    case 999:
        v47 := GenS3{F0: 1}
        _ = v47
    }
    fmt.Println(v5)
    if (127 == 127) {
        v5 -= (v5 + (^(v5)))
        v48 := GenS5{F0: (0 > 0)}
        var v49 int16 = 32767
        _ = v48
        _ = v49
    } else {
        v50 := GenS6{F0: -32768, F1: false, F2: 1}
        v5 = (v5 + 4294967295)
        _ = v50
    }
    fmt.Println(v5)
    var v51 uint64 = v5
    var v52 uint = 844
    var v53 uint32 = 2147483648
    delete(v2, v53)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v51
    _ = v52
    _ = v53
    _ = v6
    _ = v9
}
