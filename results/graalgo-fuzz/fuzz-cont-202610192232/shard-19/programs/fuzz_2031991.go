package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint; F2 uint8 }
type GenS2 struct { F0 int16; F1 int32; F2 int }
type GenS3 struct { F0 int; F1 uint8; F2 string }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 bool; F1 uint }

func main() {
    var v1 bool = (32767 >= 42)
    v2 := []int{2147483647}
    var v3 bool = (v1 && (v1 && true))
    var v4 int32 = 798
    v1 = true
    _ = len(v2)
    if (127 <= -2147483648) {
        v5 := GenS1{F0: -1, F1: 108, F2: 2}
        for fi6 := 0; fi6 < 11; fi6++ {
            v1 = (255 > 642)
            _ = fi6
        }
        if (v1 || v1) {
            fmt.Println(v4, v1)
            v7 := GenS1{F0: -846, F1: 18446744073709551615, F2: 255}
            _ = v7
        }
        _ = v5
    } else {
        var v8 int = 2147483647
        switch "foo bar" {
        case "foo bar":
            _ = len(v2)
        case "":
            var v9 float32 = 1e10
            _ = v9
        case "a":
            switch (v8 >> ((65535) & 63)) {
            case 9223372036854775807:
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            case -1:
                v1 = v1
            default:
                if (0.0 != 3.14159) {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    fmt.Println(v3, v1, v4)
                }
            }
        default:
            v2 = append(v2, (-(v8)))
        }
        v10 := GenS2{F0: 42, F1: (v4 ^ v4), F2: v8}
        _ = v10
        _ = v8
    }
    v11 := GenS3{F0: -943, F1: 62, F2: "a"}
    v2 = append(v2, -285)
    var v12 int = 127
    fmt.Println(v1, v12, v4)
    for fi13 := 0; fi13 < 3; fi13++ {
        switch 128 {
        case 1:
            if true {
                _ = len(v2)
                if (v3 || false) {
                    for fi14 := 0; fi14 < 3; fi14++ {
                        v15 := GenS4{F0: (true && (true || (-8 > 100)))}
                        if true {
                            var v16 uint64 = 1
                            _ = len(v2)
                            _ = v16
                        } else {
                            var v17 float32 = (1e10 / float32(((((float64(float32(3.14159)) * ((3.14159 + 1e10) + (3.14159 / 1.0))) + 1e10) * float64(1.5)) + 1e10)))
                            _ = len(v2)
                            _ = v17
                        }
                        switch ("a" + ("foo bar" + "\n\t")) {
                        case "a":
                            var v18 uint = 42
                            _ = v18
                        case "foo bar":
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        case "hello":
                            var v19 int16 = -1
                            _ = v19
                        default:
                            var v20 int8 = -128
                            _ = v20
                        }
                        _ = fi14
                        _ = v15
                    }
                    v4 = -163
                    var v21 float32 = (-2.25 - 0.0)
                    _ = v21
                }
            } else {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                if (("" < "foo bar") && false) {
                    var v22 uint16 = 537
                    for fi23 := 0; fi23 < 8; fi23++ {
                        _ = len(v2)
                        _ = fi23
                    }
                    _ = v22
                }
                v24 := GenS4{F0: ((((((1.5 / 1.5) - (0.0 * 1.5)) * 1.5) * (3.14159 / 0.1)) * -2.25) < (1e10 - -2.25))}
                _ = v24
            }
        case 4294967295:
            v25 := GenS3{F0: int(v4), F1: 81, F2: "\n\t"}
            _ = v25
        default:
            var v26 uint = 885
            _ = v26
        }
        if (((("hello" + ((("foo bar" + "\n\t") + "hello") + ("\n\t" + "a"))) + "\n\t") == ("" + "foo bar")) || (v4 != 2147483647)) {
            v12 -= (-(fi13))
            switch (v12 + (fi13 - 32768)) {
            case 4294967295:
                var v27 uint = 624
                _ = v27
            case 1:
                var v28 uint16 = 267
                _ = v28
            default:
                v29 := GenS5{F0: (!("hello" == "hello")), F1: 0}
                _ = v29
            }
            v1 = (4294967295 == 42)
        } else {
            fmt.Println(v3)
            var v30 bool = ("foo bar" > "")
            _ = v30
        }
        _ = fi13
    }
    for fi31 := 0; fi31 < 14; fi31++ {
        var v32 string = ("a" + "")
        v1 = v1
        _ = fi31
        _ = v32
    }
    v1 = (!(false && (!(-924 <= 1))))
    _ = len(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
}
