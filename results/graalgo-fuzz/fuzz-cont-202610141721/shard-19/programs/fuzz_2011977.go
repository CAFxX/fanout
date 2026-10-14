package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32; F2 bool }
type GenS2 struct { F0 uint; F1 string; F2 string }
type GenS3 struct { F0 int16; F1 uint8; F2 uint64 }
type GenS4 struct { F0 int8; F1 int8; F2 int }
type GenS5 struct { F0 int32; F1 int8; F2 string }
type GenS6 struct { F0 uint16 }
func genFn0() uint {
    for fi4 := 0; fi4 < 5; fi4++ {
        var v5 string = ("foo bar" + (("\n\t" + "") + ""))
        _ = fi4
        _ = v5
    }
    var v6 int16 = 32767
    var v7 uint32 = 2147483648
    _ = v6
    _ = v7
    return 255
}

func main() {
    var v1 uint8 = 2
    var v2 uint16 = 677
    v3 := []string{"\n\t", "\n\t", ("hello" + "hello"), "\n\t"}
    if ((!false) && true) {
        if len(v3) > 0 {
            v3[0] = "\n\t"
        }
        for fi8 := 0; fi8 < 3; fi8++ {
            for fi9 := 0; fi9 < 3; fi9++ {
                fmt.Println(fi8, v2)
                switch ("" + "\n\t") {
                case "":
                    v2 -= (v2 % ((42) | 1))
                case "foo bar":
                    fmt.Println(fi9, v1, v1)
                case "a":
                    fmt.Println(v1, fi9)
                default:
                    v2 = (^(v2))
                }
                v1 += (v1 >> ((0) & 63))
                if (!(1 < -558)) {
                    v3 = append(v3, ("\n\t" + "\n\t"))
                    fmt.Println(fi8, fi9)
                    v10 := GenS1{F0: "\n\t", F1: 2, F2: ("foo bar" <= ("\n\t" + "a"))}
                    _ = v10
                }
                _ = fi9
            }
            switch (fi8 - fi8) {
            case -1:
                switch "hello" {
                case "foo bar":
                    if len(v3) > 0 {
                        v3[0] = "\n\t"
                    }
                }
            case 184:
                if (((genFn0() == 0) || false) && false) {
                    for fi11 := 0; fi11 < 1; fi11++ {
                        v2 -= 251
                        switch "hello" {
                        case "foo bar":
                            v2 = 371
                            v1 += v1
                        case "hello":
                            if ("hello" != "\n\t") {
                                v12 := GenS1{F0: "hello", F1: 2, F2: true}
                                v13 := GenS1{F0: "hello", F1: 386, F2: (1.5 >= (-2.25 / 3.14159))}
                                fmt.Println(v2)
                                _ = v12
                                _ = v13
                            }
                        case "a":
                            if (!(-1 > fi8)) {
                                var v14 uint32 = 2147483648
                                _ = len(v3)
                                _ = v14
                            }
                        }
                        _ = fi11
                    }
                    var v15 int8 = 6
                    var v16 int16 = -532
                    _ = v15
                    _ = v16
                }
            default:
                var v17 float32 = 3.14159
                _ = v17
            }
            _ = fi8
        }
    } else {
        v18 := GenS2{F0: uint(v2), F1: "hello", F2: ("foo bar" + ("hello" + ("foo bar" + "")))}
        fmt.Println(v2, v2)
        _ = v18
    }
    var v19 string = ""
    if true {
        _ = len(v3)
        v20 := GenS1{F0: "\n\t", F1: 818, F2: ((v2 >> ((2) & 63)) != (v2 ^ (v2 * v2)))}
        _ = v20
    }
    switch -413 {
    case 42:
        var v21 uint64 = 0
        _ = v21
    case -9223372036854775808:
        fmt.Println(v19, v19, v19)
    default:
        v22 := GenS1{F0: v19, F1: 0, F2: false}
        _ = v22
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v2 = (v2 - (v2 >> ((341) & 63)))
    for fi23 := 0; fi23 < 11; fi23++ {
        fmt.Println(v2)
        _ = fi23
    }
    if true {
        switch "" {
        case "\n\t":
            v24 := GenS1{F0: "\n\t", F1: 65535, F2: false}
            _ = v24
        case "a":
            var v25 float32 = (1e10 / -2.25)
            _ = v25
        case "":
            var v26 uint = 0
            _ = v26
        default:
            v27 := GenS3{F0: 1000, F1: 198, F2: 894}
            _ = v27
        }
        if (2 != 18446744073709551615) {
            if (53 > 286) {
                if ((false && false) && true) {
                    if (0.1 < float64(-2.25)) {
                        var v28 int = 1
                        _ = v28
                    }
                    var v29 string = (v19 + (("hello" + v19) + (("hello" + "\n\t") + v19)))
                    _ = v29
                }
            } else {
                switch ((v19 + v19) + "foo bar") {
                case "":
                    v30 := GenS3{F0: -1, F1: v1, F2: 0}
                    _ = v30
                case "a":
                    for fi31 := 0; fi31 < 1; fi31++ {
                        _ = len(v3)
                        v1 -= (v1 | v1)
                        _ = fi31
                    }
                default:
                    if (0.1 != 3.14159) {
                        if (1 <= -2147483648) {
                            if true {
                                if (v19 != v19) {
                                    if (-128 >= 127) {
                                        v19 = "\n\t"
                                    }
                                    var v32 string = ((v19 + v19) + "")
                                    v33 := GenS4{F0: 81, F1: 42, F2: 42}
                                    _ = v32
                                    _ = v33
                                }
                                if (!true) {
                                    for fi34 := 0; fi34 < 12; fi34++ {
                                        fmt.Println(v19, fi34)
                                        v19 = v19
                                        v35 := GenS4{F0: 5, F1: -1, F2: fi34}
                                        _ = fi34
                                        _ = v35
                                    }
                                } else {
                                    var v36 int8 = 0
                                    fmt.Println(v1, v1, v19)
                                    _ = v36
                                }
                            }
                        } else {
                            v19 = ("foo bar" + v19)
                        }
                        v1 += v1
                        v1 -= v1
                    }
                }
            }
            _ = len(v3)
        } else {
            _ = len(v3)
            v37 := GenS2{F0: 2, F1: v19, F2: v19}
            v1 = (v1 % ((42) | 1))
            _ = v37
        }
    } else {
        _ = len(v3)
        v38 := GenS4{F0: 127, F1: 127, F2: 673}
        switch ("a" + v19) {
        case "":
            v39 := GenS5{F0: 42, F1: 127, F2: ("\n\t" + ("\n\t" + v19))}
            _ = v39
        case "foo bar":
            var v40 uint8 = 222
            _ = v40
        default:
            for fi41 := 0; fi41 < 4; fi41++ {
                var v42 string = ("" + ((v19 + "") + ("hello" + v19)))
                _ = fi41
                _ = v42
            }
        }
        _ = v38
    }
    if false {
        if (true && (!(false || (0 >= 2)))) {
            if (!false) {
                if false {
                    for fi43 := 0; fi43 < 11; fi43++ {
                        if true {
                            var v44 uint32 = 65535
                            var v45 uint32 = (v44 ^ (v44 - (v44 - 847)))
                            _ = v44
                            _ = v45
                        }
                        _ = fi43
                    }
                    v19 = v19
                }
                switch "" {
                case "\n\t":
                    var v46 uint16 = (v2 & (v2 | (v2 - (v2 & v2))))
                    _ = v46
                case "hello":
                    v47 := GenS6{F0: v2}
                    _ = v47
                default:
                    if (2 != 42) {
                        v2 += v2
                        var v48 int = 823
                        v49 := GenS3{F0: -5, F1: (v1 >> ((42) & 63)), F2: 0}
                        _ = v48
                        _ = v49
                    }
                }
            }
        } else {
            var v50 uint8 = (v1 ^ v1)
            _ = v50
        }
        v51 := GenS2{F0: 681, F1: (("a" + v19) + ""), F2: "\n\t"}
        _ = v51
    }
    var v52 bool = false
    for fi53 := 0; fi53 < 15; fi53++ {
        fmt.Println(fi53, fi53, v1)
        _ = fi53
    }
    v2 += (v2 ^ 208)
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v52
}
