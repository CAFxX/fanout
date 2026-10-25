package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint16; F1 uint16; F2 int64 }
type GenS4 struct { F0 uint; F1 int8; F2 string }
type GenS5 struct { F0 bool; F1 uint; F2 uint }
func genFn0(p0 uint32, p1 uint, p2 int16) string {
    v6 := GenS1{F0: 127}
    var v7 uint16 = 187
    for fi8 := 0; fi8 < 14; fi8++ {
        v9 := []int{}
        switch (fi8 + (fi8 + fi8)) {
        case -1:
            if false {
                for fi10 := 0; fi10 < 4; fi10++ {
                    v11 := GenS2{F0: 105}
                    v12 := GenS1{F0: 116}
                    _ = fi10
                    _ = v11
                    _ = v12
                }
                var v13 int32 = -2147483648
                _ = v13
            }
        case -2147483648:
            fmt.Println(p2, fi8)
        case 128:
            fmt.Println(p0)
        default:
            fmt.Println(p0)
        }
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return "hello"
}

func main() {
    var v1 int32 = 707
    v2 := []int{0, -2147483648, 1}
    var v3 int32 = v1
    v4 := []string{}
    var v5 uint32 = 65535
    v5 -= v5
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v3 -= v1
    var v14 bool = (!true)
    v14 = ("foo bar" == "a")
    v15 := GenS1{F0: 128}
    for fi16 := 0; fi16 < 18; fi16++ {
        v1 = 42
        switch fi16 {
        case 42:
            fmt.Println(v3, v3, fi16)
        case 487:
            switch fi16 {
            case 32768:
                v17 := GenS3{F0: 255, F1: uint16(v5), F2: 42}
                _ = v17
            case 0:
                v18 := GenS1{F0: 57}
                _ = v18
            default:
                switch "foo bar" {
                case "hello":
                    fmt.Println(v3)
                case "a":
                    if (((1.5 / (3.14159 / ((1.5 - 3.14159) - float32(0.1)))) + float32((0.0 + 1e10))) != ((1e10 + 3.14159) / (0.0 + (1e10 + 0.0)))) {
                        for fi19 := 0; fi19 < 15; fi19++ {
                            v1 = (v3 / (((v1 * (v3 / ((v3) | 1)))) | 1))
                            switch (fi19 | fi19) {
                            case 128:
                                fmt.Println(v3, fi16, v3)
                            }
                            switch 32768 {
                            case 1:
                                v1 = (v3 ^ -2147483648)
                            default:
                                v5 -= v5
                                if (true || (-154 < (fi19 * (fi16 << ((4294967295) & 63))))) {
                                    var v20 uint8 = 42
                                    _ = v20
                                } else {
                                    if ((("hello" + "hello") + "foo bar") <= ("\n\t" + "foo bar")) {
                                        v2 = append(v2, (-(fi19)))
                                        switch "foo bar" {
                                        case "hello":
                                            fmt.Println(v14)
                                        }
                                    } else {
                                        if v14 {
                                            var v21 string = ""
                                            if (241 != 0) {
                                                v22 := GenS2{F0: 61}
                                                if len(v2) > 0 {
                                                    v2[0] = (fi16 & (fi19 + (fi19 + 0)))
                                                }
                                                _ = v22
                                            }
                                            _ = v21
                                        } else {
                                            fmt.Println(v3, v5, v5)
                                            v5 += (^(v5))
                                            if (!v14) {
                                                var v23 int = fi16
                                                _ = v23
                                            } else {
                                                fmt.Println(v5, v14)
                                            }
                                        }
                                        v14 = (2 > 127)
                                        fmt.Println(v14)
                                    }
                                    var v24 bool = (!true)
                                    _ = v24
                                }
                            }
                            _ = fi19
                        }
                        for fi25 := 0; fi25 < 4; fi25++ {
                            v1 = int32(v5)
                            v14 = (v14 || v14)
                            _ = len(v4)
                            _ = fi25
                        }
                    } else {
                        v14 = true
                        if len(v4) > 0 {
                            v4[0] = (("foo bar" + "") + (("\n\t" + ("hello" + (("hello" + "") + ("a" + "\n\t")))) + "a"))
                        }
                        for fi26 := 0; fi26 < 9; fi26++ {
                            _ = len(v2)
                            _ = fi26
                        }
                    }
                default:
                    if false {
                        v27 := GenS3{F0: 255, F1: 255, F2: 976}
                        _ = v27
                    } else {
                        for fi28 := 0; fi28 < 16; fi28++ {
                            v3 += v1
                            switch "" {
                            case "a":
                                fmt.Println(fi16, fi16)
                            case "":
                                var v29 uint8 = 50
                                _ = v29
                            case "hello":
                                fmt.Println(v1)
                            default:
                                if false {
                                    switch (fi16 + 0) {
                                    case -115:
                                        switch 32767 {
                                        case 2147483647:
                                            if len(v2) > 0 {
                                                v2[0] = fi28
                                            }
                                        case 32768:
                                            v30 := GenS4{F0: 2, F1: -41, F2: "\n\t"}
                                            _ = v30
                                        }
                                    case -22:
                                        if len(v4) > 0 {
                                            v4[0] = "hello"
                                        }
                                    default:
                                        if v14 {
                                            v3 = (^(v3))
                                            v31 := GenS2{F0: 127}
                                            var v32 int16 = -1
                                            _ = v31
                                            _ = v32
                                        }
                                    }
                                }
                            }
                            _ = fi28
                        }
                        for fi33 := 0; fi33 < 19; fi33++ {
                            var v34 uint64 = 0
                            _ = fi33
                            _ = v34
                        }
                        if len(v4) > 0 {
                            v4[0] = "hello"
                        }
                    }
                }
            }
        default:
            if (true && v14) {
                v14 = true
                if len(v2) > 0 {
                    v2[0] = (fi16 / (((-(fi16))) | 1))
                }
            }
        }
        _ = fi16
    }
    v14 = false
    if (("" + "hello") != "hello") {
        v5 = v5
    } else {
        switch "\n\t" {
        case "hello":
            v1 += (^(v1))
        default:
            var v35 int32 = 42
            _ = v35
        }
        v5 -= 84
    }
    switch "a" {
    case "foo bar":
        if (!(42 >= 2)) {
            switch 4294967295 {
            case 32767:
                var v36 float32 = float32(3.14159)
                _ = v36
            case 2147483647:
                if len(v4) > 0 {
                    v4[0] = "foo bar"
                }
            }
            switch "\n\t" {
            case "":
                v37 := GenS1{F0: 214}
                _ = v37
            case "a":
                v38 := GenS4{F0: 0, F1: 127, F2: ""}
                _ = v38
            case "foo bar":
                var v39 uint16 = 125
                _ = v39
            }
            fmt.Println(v1, v3)
        }
    default:
        var v40 float32 = 1.5
        _ = v40
    }
    _ = len(v4)
    if (v5 >= v5) {
        if ((((!v14) || true) && true) && v14) {
            _ = len(v4)
            fmt.Println(v3)
            v2 = append(v2, 547)
        } else {
            var v41 int64 = 0
            v3 = v1
            switch 32768 {
            case 904:
                for fi42 := 0; fi42 < 2; fi42++ {
                    var v43 string = "hello"
                    _ = fi42
                    _ = v43
                }
            case 32767:
                var v44 uint = 255
                _ = v44
            default:
                v41 = (v41 & v41)
            }
            _ = v41
        }
        for fi45 := 0; fi45 < 12; fi45++ {
            v46 := GenS5{F0: v14, F1: 255, F2: 255}
            fmt.Println(v5)
            _ = fi45
            _ = v46
        }
    } else {
        fmt.Println(v1)
        if (!(-9223372036854775808 == 1)) {
            fmt.Println(v5, v3, v5)
        }
        _ = len(v2)
    }
    var v47 float32 = ((0.0 * float32(1e10)) - -2.25)
    var v48 int8 = 127
    for fi49 := 0; fi49 < 10; fi49++ {
        v48 -= v48
        switch (fi49 / ((fi49) | 1)) {
        case -2147483648:
            var v50 float64 = 1.5
            _ = v50
        case -691:
            var v51 int16 = 32767
            _ = v51
        }
        if false {
            var v52 int16 = 1000
            for fi53 := 0; fi53 < 8; fi53++ {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                v47 = 1e10
                switch "" {
                case "\n\t":
                    v48 = 1
                case "":
                    v48 = 42
                case "foo bar":
                    if v14 {
                        v54 := GenS1{F0: 42}
                        _ = v54
                    } else {
                        v1 = (v1 % ((v3) | 1))
                    }
                }
                _ = fi53
            }
            _ = v52
        }
        _ = fi49
    }
    v3 = 2147483647
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v48
    _ = v5
}
