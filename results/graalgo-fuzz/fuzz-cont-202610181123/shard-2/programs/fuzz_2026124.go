package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int8; F1 uint }

func main() {
    var v1 bool = false
    var v2 uint16 = 0
    v3 := []uint32{501, 0, 842}
    _ = len(v3)
    v4 := GenS1{F0: "a"}
    fmt.Println(v2)
    if ((!(v1 && v1)) || (842 >= 32767)) {
        fmt.Println(v2, v1)
        if (false && v1) {
            if (v1 && v1) {
                var v5 string = ("" + "hello")
                fmt.Println(v1, v2)
                _ = v5
            } else {
                v2 += (^(v2))
                v1 = (0 > 743)
                v2 = (^(v2))
            }
            v1 = ((1.5 - -2.25) < -2.25)
            fmt.Println(v1, v2)
        } else {
            if false {
                v6 := GenS1{F0: ""}
                v7 := GenS2{F0: -127, F1: 255}
                switch "" {
                case "\n\t":
                    switch -809 {
                    case 4294967295:
                        var v8 int8 = -1
                        _ = v8
                    }
                default:
                    if (1 == 2) {
                        var v9 string = "hello"
                        v10 := GenS2{F0: 42, F1: 18446744073709551615}
                        switch "" {
                        case "foo bar":
                            var v11 int64 = 9223372036854775807
                            _ = v11
                        }
                        _ = v10
                        _ = v9
                    }
                }
                _ = v6
                _ = v7
            }
        }
    }
    fmt.Println(v2, v1)
    for fi12 := 0; fi12 < 7; fi12++ {
        if true {
            switch (fi12 * fi12) {
            case 4294967295:
                for fi13 := 0; fi13 < 6; fi13++ {
                    v1 = ((1.5 + 1e10) <= -2.25)
                    var v14 uint8 = 2
                    _ = fi13
                    _ = v14
                }
            case 818:
                fmt.Println(v2, v1)
            case -9223372036854775808:
                fmt.Println(v2, fi12, v2)
            }
            var v15 uint8 = 246
            _ = v15
        } else {
            v16 := GenS1{F0: ("" + (("" + "\n\t") + "hello"))}
            _ = v16
        }
        for fi17 := 0; fi17 < 19; fi17++ {
            switch "hello" {
            case "":
                if ((fi17 | (-(fi17))) != 132) {
                    var v18 int16 = 762
                    switch fi12 {
                    case 32768:
                        switch ((("a" + "hello") + ((("foo bar" + "\n\t") + "hello") + "hello")) + "\n\t") {
                        case "foo bar":
                            var v19 float64 = (0.1 + 3.14159)
                            _ = v19
                        }
                    case 1:
                        v3 = append(v3, 2)
                    case 128:
                        if false {
                            v18 += (-(v18))
                        } else {
                            _ = len(v3)
                            fmt.Println(fi12)
                            _ = len(v3)
                        }
                    }
                    _ = v18
                } else {
                    v2 -= 374
                    switch (fi12 & (fi12 & 128)) {
                    case 32768:
                        if v1 {
                            switch 42 {
                            case -186:
                                switch "hello" {
                                case "a":
                                    if (855 >= 0) {
                                        v3 = append(v3, 648)
                                        var v20 uint16 = v2
                                        v3 = append(v3, 2147483648)
                                        _ = v20
                                    }
                                case "hello":
                                    v3 = append(v3, 42)
                                case "\n\t":
                                    var v21 uint = 1
                                    _ = v21
                                }
                            }
                        }
                    }
                    v1 = v1
                }
                var v22 int = 9223372036854775807
                fmt.Println(fi17, fi17, fi12)
                _ = v22
            case "a":
                var v23 int32 = 2147483647
                _ = v23
            case "foo bar":
                v2 = v2
            }
            _ = fi17
        }
        _ = fi12
    }
    for fi24 := 0; fi24 < 9; fi24++ {
        v25 := GenS1{F0: (("hello" + "\n\t") + ("\n\t" + "a"))}
        if (1 <= 9223372036854775807) {
            v26 := GenS1{F0: ("hello" + "a")}
            if v1 {
                fmt.Println(fi24, v2, v2)
                v27 := GenS1{F0: ""}
                _ = v27
            } else {
                if (-2147483648 >= 2147483647) {
                    if false {
                        switch int(v2) {
                        case -413:
                            var v28 uint32 = 2
                            _ = v28
                        case 42:
                            if ((((fi24 % (((fi24 << ((2) & 63))) | 1)) >= (^(fi24))) && (32767 != 689)) || v1) {
                                v2 = v2
                            }
                        case 9223372036854775807:
                            var v29 string = (("" + "a") + "\n\t")
                            _ = v29
                        default:
                            var v30 string = ("" + "\n\t")
                            _ = v30
                        }
                    }
                    var v31 int = (^(fi24))
                    _ = v31
                }
                v2 += v2
                v3 = append(v3, 2147483647)
            }
            _ = v26
        }
        _ = fi24
        _ = v25
    }
    switch 32767 {
    case 9223372036854775807:
        v2 = v2
    case 1:
        if len(v3) > 0 {
            v3[0] = 993
        }
    }
    for fi32 := 0; fi32 < 20; fi32++ {
        _ = len(v3)
        var v33 float32 = 0.1
        _ = fi32
        _ = v33
    }
    if false {
        v2 = (v2 & 1)
        v2 -= (v2 % ((v2) | 1))
        var v34 int = 4294967295
        _ = v34
    }
    v35 := GenS1{F0: "foo bar"}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v4
}
