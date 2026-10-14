package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8 }
type GenS2 struct { F0 bool; F1 uint; F2 uint64 }
type GenS3 struct { F0 int; F1 int }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int64; F1 int16 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 uint32 }
type GenS9 struct { F0 bool; F1 int16; F2 uint16 }

func main() {
    v1 := make(map[string]string)
    v2 := []uint32{2147483648, 1}
    var v3 string = (("" + "hello") + ((("foo bar" + (("hello" + ("a" + ((("" + "hello") + "\n\t") + "a"))) + "")) + "a") + "\n\t"))
    var v4 int8 = 100
    for fi5 := 0; fi5 < 14; fi5++ {
        var v6 uint64 = 4294967295
        _ = fi5
        _ = v6
    }
    if false {
        var v7 uint64 = 42
        fmt.Println(v3, v4, v3)
        v7 += 243
        _ = v7
    } else {
        for fi8 := 0; fi8 < 14; fi8++ {
            delete(v1, "\n\t")
            if true {
                v9 := GenS1{F0: -458, F1: v4}
                if (true || true) {
                    for fi10 := 0; fi10 < 16; fi10++ {
                        fmt.Println(fi8, fi10)
                        _ = fi10
                    }
                    if (322 < 469) {
                        v11 := GenS1{F0: -258, F1: (v4 % ((-19) | 1))}
                        switch v3 {
                        case "":
                            v12 := GenS1{F0: -19, F1: v4}
                            _ = v12
                        case "\n\t":
                            v3 = (v3 + v3)
                        default:
                            for fi13 := 0; fi13 < 18; fi13++ {
                                var v14 uint8 = 187
                                _ = fi13
                                _ = v14
                            }
                        }
                        _ = v11
                    } else {
                        var v15 uint16 = 0
                        _ = v15
                    }
                } else {
                    switch ((v3 + v3) + (("a" + v3) + "")) {
                    case "":
                        v3 = v3
                    case "a":
                        v4 -= v4
                    case "hello":
                        fmt.Println(fi8)
                    }
                }
                _ = v9
            } else {
                for fi16 := 0; fi16 < 8; fi16++ {
                    v3 = (((v3 + (v3 + "hello")) + v3) + v3)
                    v4 -= -128
                    _ = fi16
                }
                v17 := GenS1{F0: 355, F1: (v4 % (((v4 * (v4 * 42))) | 1))}
                _ = v17
            }
            v3 = ((("a" + v3) + (v3 + v3)) + "\n\t")
            _ = fi8
        }
        fmt.Println(v4)
        var v18 float32 = ((1e10 / -2.25) / (1e10 - ((3.14159 - 0.0) * (-2.25 * 1e10))))
        _ = v18
    }
    for fi19 := 0; fi19 < 2; fi19++ {
        for fi20 := 0; fi20 < 6; fi20++ {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            v21 := GenS2{F0: (!true), F1: 0, F2: 18446744073709551615}
            if false {
                v4 -= 62
            }
            _ = fi20
            _ = v21
        }
        v22 := GenS2{F0: (255 > 309), F1: 4294967295, F2: 2}
        _ = fi19
        _ = v22
    }
    for fi23 := 0; fi23 < 4; fi23++ {
        v4 -= v4
        v24 := GenS3{F0: (fi23 % ((fi23) | 1)), F1: fi23}
        _ = fi23
        _ = v24
    }
    _ = len(v1)
    for fi25 := 0; fi25 < 3; fi25++ {
        switch (fi25 % (((fi25 / ((893) | 1))) | 1)) {
        case -246:
            if false {
                v3 = v3
                switch 2147483647 {
                case -2147483648:
                    var v26 bool = (("" + v3) >= (v3 + (v3 + ((v3 + v3) + v3))))
                    _ = v26
                case 127:
                    for fi27 := 0; fi27 < 19; fi27++ {
                        fmt.Println(fi27)
                        _ = fi27
                    }
                case 2147483647:
                    var v28 float32 = 0.1
                    _ = v28
                }
                v4 = (^(v4))
            }
        case 128:
            if ((fi25 ^ 783) != 0) {
                var v29 uint8 = 178
                v4 = (-(v4))
                _ = v29
            } else {
                for fi30 := 0; fi30 < 11; fi30++ {
                    v3 = (v3 + "a")
                    var v31 int = fi25
                    var v32 float64 = 1.5
                    if true {
                        if true {
                            v32 = 0.0
                        }
                    } else {
                        switch v3 {
                        case "a":
                            v3 = v3
                        case "\n\t":
                            var v33 uint = 255
                            _ = v33
                        case "hello":
                            var v34 uint = 214
                            _ = v34
                        }
                        if true {
                            var v35 int = (v31 | 128)
                            if (2 <= 463) {
                                v3 = "\n\t"
                            } else {
                                if false {
                                    v36 := GenS2{F0: (1 > -603), F1: 1, F2: 2}
                                    _ = v36
                                } else {
                                    v3 = "a"
                                }
                            }
                            var v37 string = v3
                            v3 = "a"
                            v38 := GenS4{F0: (fi25 >> ((0) & 63))}
                            _ = v35
                            _ = v37
                            _ = v38
                        } else {
                            v39 := GenS3{F0: (fi30 | (fi25 + 2147483647)), F1: fi25}
                            var v40 int32 = 107
                            v3 = "hello"
                            if false {
                                v1["\n\t"] = (("hello" + ("a" + v3)) + v3)
                                var v41 uint = 601
                                _ = v41
                            }
                            fmt.Println(fi25, v32, v31)
                            _ = v39
                            _ = v40
                        }
                    }
                    switch (^(fi30)) {
                    case 48:
                        switch 197 {
                        case 342:
                            var v42 uint8 = 0
                            _ = v42
                        case 1:
                            var v43 int8 = v4
                            _ = v43
                        default:
                            _ = len(v1)
                        }
                    default:
                        v44 := GenS4{F0: 2147483647}
                        _ = v44
                    }
                    _ = fi30
                    _ = v31
                    _ = v32
                }
            }
        case 127:
            v4 -= v4
        default:
            for fi45 := 0; fi45 < 8; fi45++ {
                v3 = ("hello" + "a")
                v46 := GenS5{F0: 0, F1: 42}
                _ = fi45
                _ = v46
            }
        }
        v4 = 127
        for fi47 := 0; fi47 < 20; fi47++ {
            var v48 float64 = float64(0.1)
            _ = fi47
            _ = v48
        }
        _ = fi25
    }
    if false {
        v4 = 1
        fmt.Println(v3, v4)
        v49 := GenS6{F0: 42}
        _ = v49
    } else {
        _ = len(v1)
        for fi50 := 0; fi50 < 7; fi50++ {
            switch v3 {
            case "hello":
                for fi51 := 0; fi51 < 11; fi51++ {
                    fmt.Println(v4)
                    if len(v2) > 0 {
                        v2[0] = 944
                    }
                    _ = fi51
                }
            case "foo bar":
                var v52 float32 = (float32(0.0) * (1.5 + float32(1.5)))
                _ = v52
            case "\n\t":
                v3 = v3
            }
            v53 := GenS7{F0: 42}
            _ = fi50
            _ = v53
        }
        var v54 bool = false
        _ = v54
    }
    v55 := GenS4{F0: 32768}
    for fi56 := 0; fi56 < 20; fi56++ {
        if ((fi56 | fi56) < fi56) {
            for fi57 := 0; fi57 < 16; fi57++ {
                fmt.Println(fi56)
                v4 -= (^(v4))
                _ = fi57
            }
            v3 = "foo bar"
        }
        _ = fi56
    }
    var v58 float32 = ((0.1 / 1.0) * 3.14159)
    v4 = (v4 | 127)
    v59 := v1[(v3 + "")]
    switch v59 {
    case "hello":
        v4 += v4
    case "\n\t":
        fmt.Println(v3, v58, v4)
    case "foo bar":
        if false {
            v4 -= v4
            fmt.Println(v3)
            if (0.1 < 0.0) {
                var v60 uint8 = 127
                _ = v60
            } else {
                var v61 int16 = 1
                switch (v59 + (((("hello" + "foo bar") + v59) + ("foo bar" + v59)) + "a")) {
                case "foo bar":
                    v1[v59] = v59
                case "a":
                    for fi62 := 0; fi62 < 9; fi62++ {
                        v63 := GenS8{F0: 42}
                        _ = fi62
                        _ = v63
                    }
                default:
                    delete(v1, v3)
                }
                if len(v2) > 0 {
                    v2[0] = 623
                }
                _ = v61
            }
        }
    default:
        for fi64 := 0; fi64 < 20; fi64++ {
            v65 := GenS9{F0: true, F1: -777, F2: 2}
            _ = len(v2)
            switch 128 {
            case 42:
                switch fi64 {
                case 9223372036854775807:
                    v66 := GenS9{F0: true, F1: 226, F2: 151}
                    _ = v66
                case 726:
                    v3 = v3
                case -9223372036854775808:
                    v58 -= v58
                }
            }
            _ = fi64
            _ = v65
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v58:", v58)
    fmt.Println("v59:", v59)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v55
    _ = v58
    _ = v59
}
