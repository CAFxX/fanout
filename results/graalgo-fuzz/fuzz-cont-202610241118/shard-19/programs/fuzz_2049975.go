package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool }
type GenS2 struct { F0 int; F1 bool; F2 uint16 }
type GenS3 struct { F0 uint; F1 uint64; F2 int32 }
type GenS4 struct { F0 string; F1 uint; F2 int32 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 uint16; F1 uint64; F2 uint16 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 uint; F1 int8 }

func main() {
    v1 := []string{"a", ""}
    var v2 uint16 = 32767
    v3 := []uint32{42}
    v4 := make(map[uint32]bool)
    v5 := GenS1{F0: 65535, F1: (445 >= 0)}
    v6 := GenS1{F0: 2147483647, F1: ("" >= "hello")}
    for fi7 := 0; fi7 < 15; fi7++ {
        for fi8 := 0; fi8 < 5; fi8++ {
            v2 += (v2 % ((v2) | 1))
            _ = fi8
        }
        _ = len(v1)
        _ = fi7
    }
    var v9 bool = ((-128 == -128) && true)
    switch "a" {
    case "":
        for fi10 := 0; fi10 < 14; fi10++ {
            if v9 {
                for fi11 := 0; fi11 < 2; fi11++ {
                    if v9 {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        _ = len(v1)
                        if v9 {
                            v9 = false
                            if (v9 || (!v9)) {
                                v12 := GenS1{F0: 65535, F1: v9}
                                _ = v12
                            }
                            if false {
                                v1 = append(v1, ((("\n\t" + "hello") + ("a" + "foo bar")) + "hello"))
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                var v13 float64 = (((3.14159 * ((1.5 / 1.0) - 3.14159)) + 0.0) / ((0.1 / 0.1) * 3.14159))
                                if ((((1.5 + 1e10) * -2.25) == ((3.14159 + (float32((v13 / v13)) + ((1.5 * (1e10 * (-2.25 / 1e10))) - 1e10))) - -2.25)) && (9223372036854775807 <= -9223372036854775808)) {
                                    v4[2] = (!(127 >= 28))
                                }
                                v9 = (0 >= 255)
                                _ = v13
                            }
                        } else {
                            if len(v1) > 0 {
                                v1[0] = ("a" + ("hello" + ""))
                            }
                            var v14 uint32 = 0
                            switch (-(fi10)) {
                            case -9223372036854775808:
                                var v15 uint8 = 1
                                v1 = append(v1, (("hello" + "a") + "foo bar"))
                                _ = v15
                            default:
                                v16 := GenS2{F0: (^(fi11)), F1: false, F2: (v2 / ((v2) | 1))}
                                _ = v16
                            }
                            _ = len(v3)
                            if v9 {
                                fmt.Println(v14, fi11, fi11)
                                v14 = v14
                            } else {
                                var v17 bool = (2 == 658)
                                _ = v17
                            }
                            _ = v14
                        }
                    }
                    _ = fi11
                }
            } else {
                for fi18 := 0; fi18 < 12; fi18++ {
                    v19 := GenS3{F0: 255, F1: 0, F2: -1}
                    _ = fi18
                    _ = v19
                }
                v9 = ((^(v2)) < v2)
            }
            switch (("\n\t" + "a") + "\n\t") {
            case "foo bar":
                v20 := GenS2{F0: 127, F1: (!(v2 == v2)), F2: v2}
                _ = v20
            case "hello":
                var v21 int64 = 1
                _ = v21
            case "a":
                v1 = append(v1, "")
            }
            _ = fi10
        }
    }
    v22 := GenS4{F0: ("hello" + "\n\t"), F1: 37, F2: 639}
    var v23 uint64 = 0
    var v24 float64 = 1.5
    v25 := GenS1{F0: 754, F1: v9}
    fmt.Println(v2, v24)
    for fi26 := 0; fi26 < 2; fi26++ {
        fmt.Println(v24)
        if v9 {
            fmt.Println(v9)
            v27 := GenS2{F0: -2147483648, F1: (!v9), F2: v2}
            for fi28 := 0; fi28 < 11; fi28++ {
                v23 += (v23 >> ((18446744073709551615) & 63))
                _ = fi28
            }
            _ = v27
        } else {
            switch (fi26 / ((-740) | 1)) {
            case 32767:
                var v29 float32 = 3.14159
                _ = v29
            case 4294967295:
                if (2 == 2147483647) {
                    switch "foo bar" {
                    case "\n\t":
                        if v9 {
                            for fi30 := 0; fi30 < 13; fi30++ {
                                v23 += (^(v23))
                                _ = fi30
                            }
                            for fi31 := 0; fi31 < 2; fi31++ {
                                if (42 <= uint(v23)) {
                                    if ((v23 * 977) == 1) {
                                        v32 := GenS1{F0: uint32(v2), F1: false}
                                        v1 = append(v1, "\n\t")
                                        _ = v32
                                    }
                                    var v33 uint = 877
                                    fmt.Println(v23, v9, fi26)
                                    _ = v33
                                }
                                v2 += (v2 * v2)
                                _ = fi31
                            }
                        }
                    default:
                        if (v9 || (42 >= 1)) {
                            v34 := GenS4{F0: ("\n\t" + ""), F1: 1, F2: 2147483647}
                            _ = v34
                        }
                    }
                    v35 := GenS5{F0: 0}
                    _ = v35
                } else {
                    fmt.Println(fi26, fi26)
                }
            case 32768:
                for fi36 := 0; fi36 < 13; fi36++ {
                    v9 = (((0.0 - (0.0 + 1e10)) - 0.0) >= 1e10)
                    v9 = v9
                    fmt.Println(v9, v2)
                    var v37 uint = 0
                    if v9 {
                        v23 += (v23 << ((v37) & 63))
                        v38 := GenS6{F0: v2, F1: v23, F2: v2}
                        _ = v38
                    } else {
                        v9 = true
                        v4[2147483648] = (((42 != 42) && (42 < 818)) || (-12 > -1))
                        v39 := GenS7{F0: ("hello" + ("hello" + ("" + "foo bar")))}
                        _ = v39
                    }
                    _ = fi36
                    _ = v37
                }
            }
            if (127 <= 100) {
                fmt.Println(v23, fi26)
                v40 := v4[uint32(v23)]
                _ = v40
            } else {
                fmt.Println(v9)
                if (42 >= 0) {
                    switch "hello" {
                    case "":
                        v23 += v23
                    default:
                        for fi41 := 0; fi41 < 20; fi41++ {
                            if (0 <= 37) {
                                v42 := GenS8{F0: -101}
                                var v43 string = "a"
                                fmt.Println(fi26, v43)
                                fmt.Println(fi26)
                                _ = v42
                                _ = v43
                            }
                            if ((0 > (v2 | v2)) && (v9 || ((v23 % ((v23) | 1)) < v23))) {
                                if v9 {
                                    fmt.Println(v2, fi26)
                                    switch "a" {
                                    case "a":
                                        v44 := GenS4{F0: ("\n\t" + (((("hello" + "foo bar") + "a") + ("a" + ("\n\t" + "foo bar"))) + "")), F1: 18446744073709551615, F2: -2147483648}
                                        _ = v44
                                    case "foo bar":
                                        if v9 {
                                            delete(v4, 2147483648)
                                            var v45 int8 = 0
                                            _ = v45
                                        } else {
                                            v46 := GenS9{F0: 255, F1: -128}
                                            v2 = v2
                                            v24 -= v24
                                            _ = v46
                                        }
                                    }
                                    _ = len(v4)
                                } else {
                                    v2 += (^(v2))
                                }
                            } else {
                                v2 += v2
                                if (v23 > (v23 * v23)) {
                                    if (851 == 2147483648) {
                                        var v47 uint = 471
                                        _ = v47
                                    }
                                } else {
                                    var v48 uint8 = 0
                                    _ = v48
                                }
                                v2 = v2
                            }
                            v49 := v4[42]
                            _ = fi41
                            _ = v49
                        }
                    }
                    fmt.Println(v9)
                }
            }
            var v50 bool = v9
            _ = v50
        }
        var v51 uint64 = (v23 * v23)
        _ = fi26
        _ = v51
    }
    if v9 {
        for fi52 := 0; fi52 < 4; fi52++ {
            fmt.Println(v9)
            if true {
                v53 := GenS3{F0: 2, F1: v23, F2: 42}
                _ = v53
            } else {
                v23 = (v23 + 1)
                var v54 uint16 = uint16(v23)
                if len(v1) > 0 {
                    v1[0] = ("hello" + "a")
                }
                _ = v54
            }
            _ = fi52
        }
        fmt.Println(v23, v24, v2)
        var v55 uint16 = v2
        _ = v55
    } else {
        for fi56 := 0; fi56 < 10; fi56++ {
            v57 := GenS6{F0: (v2 / (((v2 * 32767)) | 1)), F1: (v23 % ((v23) | 1)), F2: v2}
            for fi58 := 0; fi58 < 5; fi58++ {
                delete(v4, 2)
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = fi58
            }
            fmt.Println(fi56, v2, fi56)
            _ = fi56
            _ = v57
        }
        var v59 int64 = 42
        _ = v59
    }
    v1 = append(v1, ("foo bar" + ""))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
