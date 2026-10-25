package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 string; F1 uint32 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 string; F1 uint16; F2 string }
type GenS5 struct { F0 uint; F1 bool; F2 bool }
type GenS6 struct { F0 uint16; F1 int8; F2 uint64 }
type GenS7 struct { F0 uint16; F1 int8 }
func genFn0(p0 int) uint32 {
    p0 -= -527
    var v5 float32 = 3.14159
    _ = p0
    _ = v5
    return 42
}
func genFn1() uint {
    switch 2147483647 {
    case -300:
        switch "foo bar" {
        case "hello":
            if true {
                var v6 float32 = (0.0 * (1e10 + 0.0))
                for fi7 := 0; fi7 < 6; fi7++ {
                    v6 = 0.1
                    _ = fi7
                }
                if true {
                    v6 = v6
                }
                _ = v6
            }
        case "":
            v8 := GenS1{F0: 0}
            _ = v8
        default:
            if (!(!(568 == 1))) {
                v9 := GenS2{F0: "foo bar", F1: 2147483647}
                v10 := GenS3{F0: ("\n\t" + ("foo bar" + "foo bar"))}
                _ = v10
                _ = v9
            }
        }
    }
    if true {
        if (2 >= 18446744073709551615) {
            var v11 int16 = 1
            _ = v11
        }
    }
    if false {
        var v12 int8 = -52
        _ = v12
    } else {
        v13 := []int{208}
        if true {
            v14 := GenS2{F0: "a", F1: 0}
            for fi15 := 0; fi15 < 1; fi15++ {
                var v16 int8 = -128
                _ = fi15
                _ = v16
            }
            v17 := GenS3{F0: ""}
            _ = v14
            _ = v17
        } else {
            for fi18 := 0; fi18 < 7; fi18++ {
                v13 = append(v13, (fi18 + 219))
                _ = fi18
            }
            var v19 uint64 = 0
            v19 += (v19 % ((v19) | 1))
            _ = v19
        }
        _ = v13
    }
    var v20 int8 = -1
    v21 := GenS4{F0: ("foo bar" + ""), F1: 720, F2: "foo bar"}
    _ = v20
    _ = v21
    return 42
}

func main() {
    v1 := []int{-1, -38, 4294967295}
    var v2 int32 = -197
    var v3 uint8 = 226
    var v4 int16 = 42
    fmt.Println(v4, v2)
    fmt.Println(v2, v2, v2)
    if ((0.0 > ((1.5 - float32(1e10)) * 3.14159)) && ((("a" + ((("hello" + ("hello" + ("foo bar" + (("a" + "a") + (((("hello" + "hello") + "foo bar") + "") + ("hello" + ("a" + ("foo bar" + "a")))))))) + (("foo bar" + "\n\t") + "")) + "\n\t")) + "foo bar") != "a")) {
        fmt.Println(v3, v3, v4)
        for fi22 := 0; fi22 < 8; fi22++ {
            var v23 uint = 600
            v4 = v4
            var v24 uint32 = 4294967295
            _ = fi22
            _ = v23
            _ = v24
        }
        fmt.Println(v4)
    } else {
        var v25 uint32 = 4294967295
        _ = v25
    }
    v26 := GenS5{F0: genFn1(), F1: (true && (-2.25 >= 1e10)), F2: true}
    var v27 uint64 = 653
    for fi28 := 0; fi28 < 15; fi28++ {
        var v29 string = ("hello" + ((("foo bar" + "\n\t") + "hello") + ("foo bar" + "")))
        _ = fi28
        _ = v29
    }
    fmt.Println(v2)
    switch "foo bar" {
    case "":
        fmt.Println(v2, v2)
    case "a":
        if false {
            switch 0 {
            case -22:
                for fi30 := 0; fi30 < 17; fi30++ {
                    switch fi30 {
                    case 128:
                        for fi31 := 0; fi31 < 5; fi31++ {
                            if false {
                                v32 := GenS5{F0: genFn1(), F1: (!false), F2: (v27 <= 2)}
                                v3 += (v3 << ((genFn1()) & 63))
                                var v33 uint8 = v3
                                _ = v32
                                _ = v33
                            }
                            v3 = v3
                            _ = fi31
                        }
                    }
                    _ = fi30
                }
            case 1:
                v2 += (v2 + -2147483648)
            }
        }
    case "\n\t":
        var v34 string = "a"
        _ = v34
    default:
        fmt.Println(v27, v4, v4)
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println(v2, v2)
    switch "a" {
    case "":
        switch "a" {
        case "foo bar":
            switch "foo bar" {
            case "hello":
                for fi35 := 0; fi35 < 15; fi35++ {
                    _ = len(v1)
                    var v36 uint16 = 255
                    v36 += (v36 >> ((genFn1()) & 63))
                    _ = fi35
                    _ = v36
                }
            case "\n\t":
                if ((v2 << ((4294967295) & 63)) != v2) {
                    switch "foo bar" {
                    case "\n\t":
                        fmt.Println(v2, v3, v3)
                    case "foo bar":
                        for fi37 := 0; fi37 < 17; fi37++ {
                            v2 = -168
                            if (65535 <= 2) {
                                var v38 uint8 = (v3 << ((340) & 63))
                                v3 -= v3
                                _ = v38
                            } else {
                                v27 += (v27 % ((v27) | 1))
                                v39 := GenS6{F0: 766, F1: 42, F2: (v27 << ((813) & 63))}
                                fmt.Println(v2)
                                _ = v39
                            }
                            _ = fi37
                        }
                    case "a":
                        v4 -= (^(v4))
                    }
                } else {
                    var v40 string = ("a" + ("hello" + ""))
                    v41 := GenS1{F0: genFn0(-2147483648)}
                    _ = v40
                    _ = v41
                }
            }
        }
    }
    if true {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        switch "" {
        case "foo bar":
            if (!false) {
                fmt.Println(v27)
            } else {
                v2 = v2
                switch 231 {
                case -2147483648:
                    var v42 uint = uint(v27)
                    _ = v42
                default:
                    if false {
                        v3 = 2
                    }
                }
            }
        case "a":
            if ("\n\t" >= "") {
                for fi43 := 0; fi43 < 20; fi43++ {
                    v44 := GenS7{F0: 255, F1: 127}
                    var v45 bool = false
                    fmt.Println(v2)
                    _ = fi43
                    _ = v44
                    _ = v45
                }
                var v46 int16 = v4
                var v47 uint16 = 255
                _ = v46
                _ = v47
            } else {
                v1 = append(v1, 9223372036854775807)
                if (true || ((0.1 + 0.0) != 0.0)) {
                    _ = len(v1)
                } else {
                    var v48 uint16 = 32767
                    if true {
                        fmt.Println(v3, v3, v27)
                    }
                    for fi49 := 0; fi49 < 5; fi49++ {
                        v3 = v3
                        _ = fi49
                    }
                    _ = v48
                }
                v4 += (v4 << ((18446744073709551615) & 63))
            }
        case "\n\t":
            for fi50 := 0; fi50 < 17; fi50++ {
                v27 = v27
                var v51 int16 = 1
                for fi52 := 0; fi52 < 4; fi52++ {
                    v53 := GenS1{F0: 207}
                    _ = len(v1)
                    fmt.Println(v51, v2, fi52)
                    _ = fi52
                    _ = v53
                }
                _ = fi50
                _ = v51
            }
        default:
            var v54 int8 = -8
            _ = v54
        }
        v27 += v27
    }
    _ = len(v1)
    v27 -= v27
    var v55 float64 = 1.5
    fmt.Println(v27, v27)
    if true {
        var v56 int = 42
        v57 := GenS5{F0: 65535, F1: true, F2: ((v2 >= -687) || true)}
        v27 += (v27 ^ v27)
        _ = v56
        _ = v57
    }
    for fi58 := 0; fi58 < 4; fi58++ {
        v27 -= (v27 >> ((18446744073709551615) & 63))
        _ = fi58
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v55:", v55)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v55
}
