package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string; F2 uint64 }
type GenS2 struct { F0 int16; F1 bool; F2 uint16 }
type GenS3 struct { F0 uint; F1 int16 }
type GenS4 struct { F0 int32; F1 uint64 }
type GenS5 struct { F0 uint8; F1 bool }
type GenS6 struct { F0 uint64; F1 int16 }
type GenS7 struct { F0 string; F1 bool; F2 int8 }
type GenS8 struct { F0 int64; F1 int64 }

func main() {
    var v1 uint64 = 4294967295
    v2 := []int{32768, 128}
    var v3 int = 23
    var v4 int64 = -1
    switch "foo bar" {
    case "hello":
        v1 += v1
    }
    v4 -= -1
    for fi5 := 0; fi5 < 19; fi5++ {
        for fi6 := 0; fi6 < 13; fi6++ {
            if (1000 != 1) {
                switch "\n\t" {
                case "hello":
                    v4 = (^(v4))
                default:
                    if false {
                        if (9 >= 376) {
                            if (-954 == 1) {
                                if (!false) {
                                    v7 := GenS1{F0: 42, F1: "", F2: (v1 << ((0) & 63))}
                                    var v8 uint64 = v1
                                    var v9 uint8 = 2
                                    _ = v7
                                    _ = v8
                                    _ = v9
                                } else {
                                    v10 := GenS1{F0: 912, F1: ("" + "hello"), F2: 283}
                                    var v11 uint64 = v1
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    _ = v10
                                    _ = v11
                                }
                                var v12 uint16 = 919
                                _ = v12
                            } else {
                                v13 := GenS2{F0: 390, F1: (52 == 237), F2: 255}
                                v3 = 32768
                                v14 := GenS3{F0: 4294967295, F1: -32768}
                                _ = v13
                                _ = v14
                            }
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        }
                        switch "\n\t" {
                        case "":
                            v1 += (v1 ^ v1)
                        case "foo bar":
                            var v15 string = ("hello" + "hello")
                            _ = v15
                        }
                    } else {
                        if ((3.14159 / 3.14159) < 1.5) {
                            v3 -= v3
                        } else {
                            switch "\n\t" {
                            case "":
                                if ("a" >= "hello") {
                                    if (!true) {
                                        v3 = fi5
                                        switch ((("foo bar" + "foo bar") + (("\n\t" + "a") + "a")) + "a") {
                                        case "":
                                            v3 -= (fi6 >> ((2) & 63))
                                        case "foo bar":
                                            if false {
                                                v2 = append(v2, (v3 & (fi6 + (fi6 & (-(v3))))))
                                            } else {
                                                switch "\n\t" {
                                                case "":
                                                    var v16 int64 = v4
                                                    _ = v16
                                                case "a":
                                                    switch fi6 {
                                                    case 42:
                                                        v17 := GenS4{F0: -576, F1: 2}
                                                        _ = v17
                                                    default:
                                                        if len(v2) > 0 {
                                                            v2[0] = fi6
                                                        }
                                                    }
                                                }
                                                if len(v2) > 0 {
                                                    v2[0] = v3
                                                }
                                                v4 = (^(v4))
                                                var v18 int64 = v4
                                                v19 := GenS2{F0: 1, F1: false, F2: 0}
                                                _ = v18
                                                _ = v19
                                            }
                                        case "\n\t":
                                            if (-2147483648 > 1) {
                                                v3 = fi5
                                                v1 = (^(v1))
                                                var v20 int = fi6
                                                _ = v20
                                            }
                                        }
                                        v21 := GenS1{F0: 4294967295, F1: (((("\n\t" + "\n\t") + "a") + "foo bar") + "a"), F2: v1}
                                        _ = v21
                                    }
                                }
                                v22 := GenS5{F0: 1, F1: true}
                                if true {
                                    var v23 int16 = 0
                                    v2 = append(v2, fi6)
                                    switch ("a" + "foo bar") {
                                    case "a":
                                        var v24 uint32 = 2
                                        _ = v24
                                    case "hello":
                                        if len(v2) > 0 {
                                            v2[0] = (fi5 + (fi5 | fi5))
                                        }
                                    default:
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                    }
                                    _ = v23
                                } else {
                                    v1 += 1
                                }
                                _ = v22
                            case "a":
                                var v25 int = fi5
                                v26 := GenS4{F0: 2147483647, F1: (v1 >> ((18446744073709551615) & 63))}
                                _ = v25
                                _ = v26
                            case "foo bar":
                                var v27 uint16 = 633
                                _ = v27
                            default:
                                var v28 float64 = (((0.1 * 1e10) * (1e10 - (0.1 * 3.14159))) + 0.1)
                                _ = v28
                            }
                            var v29 int8 = -92
                            v2 = append(v2, (v3 >> ((0) & 63)))
                            _ = v29
                        }
                        if false {
                            switch (^(v3)) {
                            case -337:
                                v3 += fi6
                            case 887:
                                if len(v2) > 0 {
                                    v2[0] = (^(v3))
                                }
                                var v30 uint64 = 42
                                _ = v30
                            }
                        }
                    }
                }
                v31 := GenS6{F0: v1, F1: -352}
                _ = v31
            }
            if (!false) {
                if true {
                    if true {
                        var v32 string = ""
                        var v33 int32 = -2147483648
                        var v34 uint64 = 2
                        _ = v32
                        _ = v33
                        _ = v34
                    }
                    switch "a" {
                    case "hello":
                        var v35 string = "a"
                        _ = v35
                    }
                    if false {
                        v4 = v4
                        if false {
                            v1 = v1
                            fmt.Println(v4, v4, v3)
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        } else {
                            var v36 uint32 = 2147483648
                            _ = v36
                        }
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    }
                    v37 := GenS7{F0: "foo bar", F1: (40 <= 1), F2: -128}
                    var v38 uint64 = (^(v1))
                    _ = v37
                    _ = v38
                } else {
                    v4 -= v4
                    fmt.Println(fi6, fi6, v3)
                    v39 := GenS1{F0: 65535, F1: ("a" + ("" + "foo bar")), F2: v1}
                    _ = v39
                }
                if ("hello" <= (("foo bar" + (("a" + "a") + ("foo bar" + "\n\t"))) + "")) {
                    fmt.Println(v4)
                    var v40 uint = 88
                    _ = v40
                }
            }
            _ = fi6
        }
        for fi41 := 0; fi41 < 12; fi41++ {
            v42 := GenS6{F0: (v1 >> ((708) & 63)), F1: 32767}
            _ = fi41
            _ = v42
        }
        _ = fi5
    }
    v43 := GenS8{F0: (v4 * (v4 - v4)), F1: 1}
    fmt.Println(v1, v1)
    var v44 int64 = -9223372036854775808
    var v45 uint16 = 334
    v45 = (v45 % ((v45) | 1))
    if true {
        v1 -= (v1 >> ((29) & 63))
        switch 128 {
        case 2147483647:
            var v46 float64 = 1e10
            _ = v46
        default:
            v3 = 763
        }
        v3 = v3
    }
    if true {
        v44 = (v4 + (-(v4)))
        v1 += (^(v1))
    } else {
        var v47 int8 = 36
        _ = v47
    }
    switch v3 {
    case 0:
        var v48 int8 = 100
        _ = v48
    case -9223372036854775808:
        var v49 float64 = 1e10
        _ = v49
    }
    var v50 string = ""
    var v51 uint32 = 65535
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v45
    _ = v50
    _ = v51
}
