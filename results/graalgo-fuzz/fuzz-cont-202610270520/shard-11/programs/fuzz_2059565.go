package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 bool; F1 uint }

func main() {
    var v1 int = 127
    var v2 uint32 = 4294967295
    v3 := []string{"foo bar"}
    switch ((("foo bar" + "") + ("a" + (("a" + "a") + "hello"))) + "\n\t") {
    case "\n\t":
        if ((v2 >> ((798) & 63)) > (v2 % (((^(v2))) | 1))) {
            for fi4 := 0; fi4 < 12; fi4++ {
                var v5 float64 = 0.0
                _ = fi4
                _ = v5
            }
        }
    }
    switch "foo bar" {
    case "":
        if true {
            switch ("" + "hello") {
            case "":
                switch "a" {
                case "a":
                    if (0 < 203) {
                        v1 -= (v1 | (v1 * v1))
                        switch 0 {
                        case 471:
                            v2 -= v2
                        case -9223372036854775808:
                            for fi6 := 0; fi6 < 16; fi6++ {
                                var v7 float32 = (0.0 + (-2.25 / -2.25))
                                _ = fi6
                                _ = v7
                            }
                        case 599:
                            v2 = 160
                        }
                    }
                }
            }
            switch "hello" {
            case "a":
                var v8 float32 = (1e10 / -2.25)
                _ = v8
            case "\n\t":
                if (583 != 488) {
                    var v9 uint64 = 891
                    fmt.Println(v9)
                    v10 := GenS1{F0: -109, F1: 0}
                    _ = v10
                    _ = v9
                } else {
                    v2 = v2
                    if len(v3) > 0 {
                        v3[0] = ("hello" + "\n\t")
                    }
                }
            default:
                fmt.Println(v1, v1)
            }
        } else {
            switch v1 {
            case 32768:
                v1 = 9223372036854775807
            case -9223372036854775808:
                switch v1 {
                case 330:
                    v11 := GenS2{F0: -2147483648}
                    _ = v11
                }
            case 423:
                if (false && (int16(v1) >= -289)) {
                    if (65535 != 1) {
                        var v12 int64 = -1
                        v2 -= 65535
                        _ = v12
                    }
                    for fi13 := 0; fi13 < 15; fi13++ {
                        switch fi13 {
                        case 32767:
                            v1 = (v1 & 244)
                        case 42:
                            v2 = v2
                        }
                        if len(v3) > 0 {
                            v3[0] = ("\n\t" + ("" + "foo bar"))
                        }
                        _ = fi13
                    }
                    v14 := GenS1{F0: -128, F1: 1}
                    _ = v14
                }
            }
            v15 := GenS2{F0: 696}
            if (!false) {
                var v16 uint8 = 160
                fmt.Println(v1)
                _ = v16
            } else {
                var v17 string = (((("a" + "") + "a") + "\n\t") + ("foo bar" + "foo bar"))
                fmt.Println(v1, v2)
                v1 -= (-(v1))
                _ = v17
            }
            _ = v15
        }
    case "a":
        v18 := GenS3{F0: 255}
        _ = v18
    case "foo bar":
        fmt.Println(v2, v1)
    }
    v1 = -378
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi19 := 0; fi19 < 10; fi19++ {
        fmt.Println(v1, v2)
        var v20 int = (fi19 >> ((1) & 63))
        _ = fi19
        _ = v20
    }
    v3 = append(v3, ("hello" + "a"))
    for fi21 := 0; fi21 < 15; fi21++ {
        v1 += v1
        var v22 uint32 = (v2 & 2147483647)
        v22 -= 2147483648
        _ = fi21
        _ = v22
    }
    if len(v3) > 0 {
        v3[0] = "a"
    }
    v23 := GenS4{F0: false, F1: 0}
    v1 += v1
    if false {
        var v24 float32 = 1.5
        _ = v24
    } else {
        if len(v3) > 0 {
            v3[0] = "a"
        }
    }
    v2 = (v2 ^ 0)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
}
