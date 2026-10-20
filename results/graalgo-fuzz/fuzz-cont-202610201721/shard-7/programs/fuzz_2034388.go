package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint }

func main() {
    v1 := []string{(("" + "") + ""), ((("hello" + (("foo bar" + ("\n\t" + "")) + "a")) + "\n\t") + (("a" + "\n\t") + "")), ("a" + "hello"), "\n\t"}
    var v2 int32 = -869
    var v3 int32 = v2
    v3 = v2
    fmt.Println(v3)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v4 uint64 = 768
    switch "a" {
    case "\n\t":
        _ = len(v1)
    default:
        switch 385 {
        case 128:
            if (false || ("a" != (("foo bar" + "foo bar") + ""))) {
                if ((255 >= 53) || false) {
                    var v5 int64 = 42
                    var v6 uint8 = 42
                    _ = v5
                    _ = v6
                }
                switch "a" {
                case "a":
                    var v7 float64 = 1e10
                    _ = v7
                case "\n\t":
                    v1 = append(v1, "a")
                case "":
                    for fi8 := 0; fi8 < 16; fi8++ {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        _ = fi8
                    }
                default:
                    switch -2147483648 {
                    case -9223372036854775808:
                        if (false && false) {
                            v9 := GenS1{F0: -2147483648, F1: 42}
                            _ = len(v1)
                            for fi10 := 0; fi10 < 2; fi10++ {
                                v11 := GenS1{F0: 2147483647, F1: 4294967295}
                                _ = fi10
                                _ = v11
                            }
                            _ = v9
                        } else {
                            fmt.Println(v4, v4)
                        }
                    case -739:
                        if false {
                            var v12 uint64 = (v4 * 1)
                            var v13 int32 = (v2 << ((433) & 63))
                            var v14 int64 = 9223372036854775807
                            _ = v12
                            _ = v13
                            _ = v14
                        } else {
                            var v15 float32 = ((0.1 / 0.1) - ((1.5 * (3.14159 * (3.14159 - 0.1))) + 0.0))
                            v16 := GenS1{F0: 127, F1: 65535}
                            v2 = 42
                            _ = v15
                            _ = v16
                        }
                    case 2147483647:
                        switch 621 {
                        case 9223372036854775807:
                            v2 = v2
                        case -1:
                            var v17 int64 = 1
                            _ = v17
                        default:
                            v4 += (^(v4))
                        }
                    }
                }
            }
        default:
            var v18 int32 = v2
            _ = v18
        }
    }
    v3 += (v2 / (((v3 - (^(v2)))) | 1))
    v4 += (v4 & 42)
    if true {
        if (1 < 42) {
            v3 = v3
            v19 := GenS1{F0: -9223372036854775808, F1: 2}
            switch 1 {
            case 347:
                v3 += v2
            case 9223372036854775807:
                _ = len(v1)
            case 4294967295:
                fmt.Println(v3, v2)
            default:
                var v20 int16 = -310
                _ = v20
            }
            _ = v19
        } else {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            switch 9223372036854775807 {
            case 128:
                fmt.Println(v4)
            case 744:
                v2 = 2147483647
            }
            switch (("a" + ("hello" + "\n\t")) + "hello") {
            case "hello":
                v2 -= v2
            case "a":
                v21 := GenS1{F0: -928, F1: 654}
                _ = v21
            }
        }
    } else {
        var v22 uint = 18446744073709551615
        if (249 == -2147483648) {
            v2 += (v2 * -832)
            v1 = append(v1, ("" + ""))
        }
        v4 = v4
        _ = v22
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
}
