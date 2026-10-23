package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint8; F1 int; F2 int }

func main() {
    var v1 uint32 = 1
    v2 := []string{"foo bar", ("a" + "hello"), "\n\t", "a"}
    var v3 uint = 453
    var v4 uint32 = v1
    var v5 uint32 = v1
    var v6 uint8 = 128
    v7 := GenS1{F0: uint(v6)}
    switch ((("a" + "a") + (((((("foo bar" + (("foo bar" + "a") + "hello")) + "\n\t") + (((("hello" + "\n\t") + "a") + "") + ("foo bar" + "foo bar"))) + "a") + "a") + "\n\t")) + (("a" + ("hello" + ("foo bar" + "\n\t"))) + "hello")) {
    case "foo bar":
        fmt.Println(v4)
    case "":
        fmt.Println(v1, v1, v3)
    case "a":
        v5 -= (v4 + v4)
    }
    for fi8 := 0; fi8 < 11; fi8++ {
        switch "a" {
        case "foo bar":
            v1 = v4
        case "\n\t":
            var v9 int32 = 2147483647
            _ = v9
        case "hello":
            v10 := GenS1{F0: v3}
            _ = v10
        default:
            for fi11 := 0; fi11 < 9; fi11++ {
                if (!true) {
                    if (false && false) {
                        fmt.Println(v5, v1)
                    }
                } else {
                    var v12 float64 = 0.0
                    v4 = (v1 + (^(v4)))
                    _ = v12
                }
                _ = fi11
            }
        }
        switch ((("hello" + ("a" + "foo bar")) + "\n\t") + "hello") {
        case "hello":
            var v13 int16 = 1000
            _ = v13
        default:
            for fi14 := 0; fi14 < 13; fi14++ {
                var v15 uint32 = 2147483648
                switch "\n\t" {
                case "hello":
                    v15 -= (v1 + v4)
                default:
                    v16 := GenS1{F0: (v3 >> ((2) & 63))}
                    _ = v16
                }
                if true {
                    if true {
                        if true {
                            var v17 uint = v3
                            _ = v17
                        } else {
                            if len(v2) > 0 {
                                v2[0] = ""
                            }
                        }
                    }
                }
                _ = fi14
                _ = v15
            }
        }
        var v18 uint64 = 0
        _ = fi8
        _ = v18
    }
    var v19 int8 = 127
    v20 := GenS2{F0: v6, F1: 387, F2: -800}
    v5 = uint32(v3)
    if true {
        if false {
            var v21 int8 = (v19 ^ v19)
            v2 = append(v2, "")
            _ = v21
        } else {
            var v22 uint16 = 38
            _ = v22
        }
        for fi23 := 0; fi23 < 20; fi23++ {
            v2 = append(v2, "a")
            _ = fi23
        }
        var v24 int = 127
        _ = v24
    }
    v25 := GenS1{F0: v3}
    var v26 string = ""
    v6 = v6
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
