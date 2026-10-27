package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int8; F1 uint16; F2 int32 }
type GenS3 struct { F0 uint64; F1 uint8 }
type GenS4 struct { F0 int64; F1 int8; F2 uint32 }
type GenS5 struct { F0 int64; F1 uint8 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 int8 }

func main() {
    v1 := []string{((("foo bar" + ((("a" + "") + "") + "")) + "foo bar") + "hello"), "a", "foo bar", (("hello" + ("hello" + "\n\t")) + "hello")}
    var v2 int8 = 42
    var v3 int8 = (-(v2))
    fmt.Println(v2)
    _ = len(v1)
    if ((v2 | v3) > v3) {
        fmt.Println(v3, v3)
        if len(v1) > 0 {
            v1[0] = ""
        }
    }
    for fi4 := 0; fi4 < 14; fi4++ {
        if len(v1) > 0 {
            v1[0] = "foo bar"
        }
        v5 := GenS1{F0: 0}
        switch "hello" {
        case "hello":
            if true {
                v2 = (v3 / ((20) | 1))
            }
        case "\n\t":
            v3 -= 98
        case "foo bar":
            v3 = (^(v3))
        default:
            switch (fi4 << ((0) & 63)) {
            case -1:
                switch fi4 {
                case 725:
                    if (0 <= -350) {
                        v3 -= (v3 - -1)
                    }
                }
            }
        }
        _ = fi4
        _ = v5
    }
    v2 -= -1
    for fi6 := 0; fi6 < 4; fi6++ {
        if true {
            v3 = 100
            v7 := GenS2{F0: (-(v3)), F1: 922, F2: 0}
            _ = v7
        }
        v8 := GenS1{F0: 0}
        v9 := GenS1{F0: 132}
        _ = fi6
        _ = v8
        _ = v9
    }
    v10 := GenS1{F0: 42}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v11 float32 = (3.14159 - -2.25)
    v2 = (v2 - (v2 << ((1) & 63)))
    switch "hello" {
    case "\n\t":
        switch ("hello" + "\n\t") {
        case "":
            v2 -= v3
        default:
            v3 = v3
        }
    case "":
        if true {
            fmt.Println(v2, v3)
            for fi12 := 0; fi12 < 14; fi12++ {
                _ = len(v1)
                if (true || (true && (-1 != 0))) {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                } else {
                    fmt.Println(v11)
                }
                _ = fi12
            }
        } else {
            v13 := GenS3{F0: 18446744073709551615, F1: 128}
            if false {
                if false {
                    var v14 int32 = -1
                    var v15 uint8 = 127
                    _ = v14
                    _ = v15
                } else {
                    if false {
                        fmt.Println(v11)
                    }
                }
                switch -184 {
                case -568:
                    if true {
                        v16 := GenS3{F0: 430, F1: 0}
                        _ = v16
                    } else {
                        for fi17 := 0; fi17 < 9; fi17++ {
                            var v18 uint64 = 18446744073709551615
                            switch fi17 {
                            case 42:
                                switch ("a" + "foo bar") {
                                case "hello":
                                    for fi19 := 0; fi19 < 13; fi19++ {
                                        if true {
                                            if len(v1) > 0 {
                                                v1[0] = ("a" + "\n\t")
                                            }
                                            _ = len(v1)
                                            v11 = v11
                                            if len(v1) > 0 {
                                                v1[0] = "foo bar"
                                            }
                                            v2 = (v2 >> ((1) & 63))
                                            fmt.Println(v2, v18)
                                            v20 := GenS3{F0: (^(v18)), F1: 255}
                                            _ = len(v1)
                                            if false {
                                                var v21 uint = 540
                                                _ = v21
                                            } else {
                                                switch "hello" {
                                                case "\n\t":
                                                    switch (^(fi19)) {
                                                    case 1:
                                                        var v22 uint32 = 4294967295
                                                        _ = v22
                                                    case -15:
                                                        v2 = v3
                                                    case 4294967295:
                                                        var v23 bool = true
                                                        _ = v23
                                                    }
                                                case "hello":
                                                    fmt.Println(v3, v11)
                                                case "":
                                                    v2 = 80
                                                }
                                            }
                                            _ = v20
                                        }
                                        v3 = 42
                                        v11 = (0.0 / v11)
                                        _ = fi19
                                    }
                                case "":
                                    if true {
                                        v3 -= (v3 / ((v2) | 1))
                                    } else {
                                        v11 = v11
                                        var v24 string = "foo bar"
                                        _ = v24
                                    }
                                case "\n\t":
                                    v25 := GenS4{F0: -1, F1: (v3 % ((v2) | 1)), F2: 1}
                                    _ = v25
                                }
                            case 110:
                                switch "" {
                                case "a":
                                    switch ("hello" + "") {
                                    case "":
                                        if len(v1) > 0 {
                                            v1[0] = ("foo bar" + "foo bar")
                                        }
                                    }
                                case "foo bar":
                                    if true {
                                        v26 := GenS5{F0: 1, F1: 102}
                                        _ = v26
                                    }
                                case "\n\t":
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                }
                            case -831:
                                switch "hello" {
                                case "":
                                    if (false || true) {
                                        for fi27 := 0; fi27 < 1; fi27++ {
                                            v2 = v2
                                            _ = fi27
                                        }
                                        var v28 uint = 18446744073709551615
                                        fmt.Println(v18, v11, v28)
                                        _ = v28
                                    }
                                case "foo bar":
                                    if (241 < 124) {
                                        v2 += -1
                                    } else {
                                        v29 := GenS3{F0: (v18 >> ((661) & 63)), F1: 1}
                                        _ = v29
                                    }
                                }
                            default:
                                for fi30 := 0; fi30 < 14; fi30++ {
                                    var v31 uint8 = 179
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    var v32 uint64 = 2
                                    _ = fi30
                                    _ = v31
                                    _ = v32
                                }
                            }
                            _ = fi17
                            _ = v18
                        }
                        for fi33 := 0; fi33 < 16; fi33++ {
                            v34 := GenS2{F0: -1, F1: 98, F2: 42}
                            if ("\n\t" < "hello") {
                                _ = len(v1)
                            }
                            if (1 == 18446744073709551615) {
                                if (-9223372036854775808 > 0) {
                                    fmt.Println(v3, fi33)
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                }
                                v11 = v11
                                if (-9223372036854775808 <= 42) {
                                    if (v2 == v3) {
                                        v3 = v2
                                    } else {
                                        v35 := GenS6{F0: 0}
                                        _ = v35
                                    }
                                } else {
                                    var v36 uint8 = 136
                                    fmt.Println(v3, v3, v3)
                                    if (2 <= 501) {
                                        v37 := GenS7{F0: 1}
                                        var v38 uint32 = 65535
                                        for fi39 := 0; fi39 < 7; fi39++ {
                                            fmt.Println(fi33, fi33, v11)
                                            _ = fi39
                                        }
                                        _ = v37
                                        _ = v38
                                    }
                                    _ = v36
                                }
                            } else {
                                fmt.Println(v2, v3)
                            }
                            _ = fi33
                            _ = v34
                        }
                    }
                case -621:
                    var v40 uint = 18446744073709551615
                    _ = v40
                case 456:
                    switch "" {
                    case "a":
                        for fi41 := 0; fi41 < 15; fi41++ {
                            fmt.Println(v3)
                            _ = fi41
                        }
                    default:
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    }
                }
            }
            _ = v13
        }
    case "foo bar":
        var v42 string = "a"
        _ = v42
    default:
        v43 := GenS2{F0: v2, F1: 527, F2: -591}
        _ = v43
    }
    v1 = append(v1, "a")
    v2 -= 42
    v3 += 37
    v44 := GenS8{F0: (v3 - 127)}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v44
}
