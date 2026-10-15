package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32 }
type GenS2 struct { F0 uint8; F1 uint }
type GenS3 struct { F0 uint16; F1 string }
type GenS4 struct { F0 string; F1 bool }
type GenS5 struct { F0 uint8; F1 int64 }

func main() {
    var v1 bool = ((!false) && true)
    var v2 int8 = 100
    var v3 int64 = 1
    var v4 uint16 = 32767
    v5 := []uint32{497, 744, 2147483647}
    var v6 uint16 = (^(v4))
    for fi7 := 0; fi7 < 20; fi7++ {
        var v8 bool = v1
        _ = fi7
        _ = v8
    }
    var v9 float32 = (0.1 - ((1e10 - 3.14159) + 3.14159))
    if false {
        v4 = (v4 << ((795) & 63))
    }
    var v10 uint16 = v6
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    fmt.Println(v9, v4, v3)
    if v1 {
        if false {
            for fi11 := 0; fi11 < 20; fi11++ {
                v12 := GenS1{F0: ("hello" + "\n\t"), F1: 1}
                var v13 int8 = (v2 + -65)
                if (984 == 1) {
                    var v14 uint8 = 216
                    v15 := GenS2{F0: (v14 % ((42) | 1)), F1: 616}
                    _ = v14
                    _ = v15
                }
                _ = fi11
                _ = v12
                _ = v13
            }
            var v16 int16 = 0
            _ = v16
        } else {
            switch 2147483647 {
            case 32768:
                v6 -= 42
            }
            var v17 int64 = (v3 >> ((6) & 63))
            _ = v17
        }
        v1 = (522 <= 1)
        fmt.Println(v3, v9, v2)
    }
    var v18 int32 = 42
    if (v10 >= 65535) {
        if (!(("a" + "") == "hello")) {
            v19 := GenS2{F0: 128, F1: 42}
            v1 = v1
            _ = v19
        } else {
            _ = len(v5)
        }
    } else {
        var v20 float64 = (3.14159 * -2.25)
        if v1 {
            if (948 != v18) {
                fmt.Println(v3)
                var v21 uint16 = v10
                _ = v21
            }
            var v22 bool = (!v1)
            _ = v22
        } else {
            switch "" {
            case "a":
                v4 -= 42
            case "hello":
                v23 := GenS3{F0: 2, F1: ""}
                _ = v23
            case "":
                v3 = (v3 - v3)
            }
            switch "hello" {
            case "":
                var v24 int64 = 9223372036854775807
                _ = v24
            }
            switch "foo bar" {
            case "\n\t":
                if true {
                    v2 += (v2 >> ((99) & 63))
                    v25 := GenS2{F0: 83, F1: 65535}
                    v2 = (-(v2))
                    _ = v25
                } else {
                    if true {
                        var v26 uint32 = 772
                        switch "\n\t" {
                        case "hello":
                            if (4294967295 < 255) {
                                _ = len(v5)
                            }
                        case "\n\t":
                            v10 = (v4 + v4)
                        default:
                            _ = len(v5)
                        }
                        var v27 int64 = v3
                        _ = v26
                        _ = v27
                    }
                    v28 := GenS2{F0: 128, F1: 18446744073709551615}
                    _ = v28
                }
            case "":
                switch "\n\t" {
                case "foo bar":
                    v3 = -9223372036854775808
                case "a":
                    v9 += v9
                }
            case "hello":
                if true {
                    v29 := GenS1{F0: "hello", F1: 0}
                    fmt.Println(v1, v1)
                    fmt.Println(v4, v4)
                    _ = v29
                } else {
                    v5 = append(v5, 576)
                }
            default:
                for fi30 := 0; fi30 < 8; fi30++ {
                    for fi31 := 0; fi31 < 2; fi31++ {
                        var v32 uint16 = v4
                        _ = len(v5)
                        _ = fi31
                        _ = v32
                    }
                    _ = fi30
                }
            }
        }
        if v1 {
            v33 := GenS1{F0: "", F1: (v18 - (v18 / (((v18 >> ((534) & 63))) | 1)))}
            var v34 uint64 = 1
            v3 = v3
            _ = v33
            _ = v34
        }
        _ = v20
    }
    v5 = append(v5, 1)
    if (!true) {
        if len(v5) > 0 {
            v5[0] = 2147483648
        }
        for fi35 := 0; fi35 < 14; fi35++ {
            v18 -= v18
            _ = fi35
        }
    } else {
        var v36 float64 = 0.0
        v37 := GenS4{F0: "hello", F1: ((!true) && (("hello" != "\n\t") || v1))}
        for fi38 := 0; fi38 < 19; fi38++ {
            v3 -= -9223372036854775808
            v2 += 127
            v39 := GenS4{F0: ("hello" + ""), F1: (v1 || v1)}
            _ = fi38
            _ = v39
        }
        _ = v36
        _ = v37
    }
    v40 := GenS5{F0: 205, F1: (v3 / (((v3 * v3)) | 1))}
    for fi41 := 0; fi41 < 6; fi41++ {
        if true {
            switch (fi41 + 9223372036854775807) {
            case 1:
                var v42 int32 = 1
                _ = v42
            case 2147483647:
                switch "foo bar" {
                case "\n\t":
                    v10 -= 32767
                case "":
                    if len(v5) > 0 {
                        v5[0] = 1
                    }
                case "foo bar":
                    var v43 uint = 2
                    _ = v43
                }
            }
            var v44 int32 = v18
            _ = v44
        }
        v45 := GenS3{F0: 31, F1: "\n\t"}
        _ = fi41
        _ = v45
    }
    fmt.Println(v4, v1, v10)
    v46 := GenS1{F0: (("a" + "") + ("foo bar" + "foo bar")), F1: (v18 << ((65535) & 63))}
    var v47 float32 = (3.14159 + (v9 - 1e10))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v46
    _ = v47
    _ = v5
    _ = v6
    _ = v9
}
