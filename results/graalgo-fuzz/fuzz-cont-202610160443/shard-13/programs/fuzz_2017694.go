package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int8; F1 uint64 }

func main() {
    var v1 bool = (42 == -9223372036854775808)
    v2 := []uint32{42, 2147483647, 1}
    var v3 int64 = -713
    var v4 uint32 = 4294967295
    v5 := GenS1{F0: "a"}
    if (884 <= 42) {
        var v6 int16 = 32767
        for fi7 := 0; fi7 < 9; fi7++ {
            v1 = true
            _ = fi7
        }
        fmt.Println(v4, v6, v1)
        _ = v6
    }
    switch "a" {
    case "\n\t":
        if (!(false || (1 <= 194))) {
            switch "a" {
            case "hello":
                v8 := GenS1{F0: "a"}
                _ = v8
            default:
                var v9 uint32 = (^(v4))
                _ = v9
            }
            v4 += (v4 << ((42) & 63))
            v10 := GenS1{F0: "foo bar"}
            _ = v10
        }
    }
    var v11 uint = 0
    v11 = (v11 / ((2) | 1))
    if true {
        switch 971 {
        case -894:
            var v12 uint64 = 4294967295
            _ = v12
        default:
            switch 175 {
            case 127:
                switch "hello" {
                case "a":
                    switch "" {
                    case "a":
                        var v13 float64 = float64(-2.25)
                        _ = v13
                    case "hello":
                        v1 = v1
                    case "\n\t":
                        for fi14 := 0; fi14 < 15; fi14++ {
                            var v15 uint = 4294967295
                            var v16 int16 = -1
                            _ = fi14
                            _ = v15
                            _ = v16
                        }
                    }
                case "\n\t":
                    if v1 {
                        if len(v2) > 0 {
                            v2[0] = 650
                        }
                        fmt.Println(v3)
                    }
                default:
                    var v17 float32 = 0.1
                    _ = v17
                }
            }
        }
        v3 = 9223372036854775807
    } else {
        var v18 bool = v1
        if len(v2) > 0 {
            v2[0] = 1
        }
        _ = v18
    }
    switch ("foo bar" + "") {
    case "foo bar":
        var v19 bool = (!(1000 <= 1000))
        _ = v19
    default:
        v4 += v4
    }
    v4 = v4
    for fi20 := 0; fi20 < 17; fi20++ {
        if len(v2) > 0 {
            v2[0] = v4
        }
        fmt.Println(v4, v1, v4)
        var v21 int8 = 42
        _ = fi20
        _ = v21
    }
    switch "\n\t" {
    case "hello":
        var v22 uint8 = 2
        _ = v22
    default:
        v23 := GenS2{F0: 0, F1: 2}
        _ = v23
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi24 := 0; fi24 < 20; fi24++ {
        _ = len(v2)
        var v25 int = fi24
        fmt.Println(v1)
        _ = fi24
        _ = v25
    }
    v11 = v11
    if v1 {
        var v26 uint = 4294967295
        if ((v3 << ((v11) & 63)) >= -9223372036854775808) {
            var v27 bool = true
            v28 := GenS1{F0: ("\n\t" + ("foo bar" + "\n\t"))}
            v29 := GenS1{F0: ((("foo bar" + "a") + ((("a" + ("\n\t" + "a")) + "foo bar") + ("\n\t" + ((("hello" + "a") + "a") + "foo bar")))) + "\n\t")}
            _ = v27
            _ = v28
            _ = v29
        } else {
            var v30 float64 = (1.5 * (0.0 / 1.5))
            if (uint8(v26) == 2) {
                var v31 uint16 = 65535
                _ = v31
            }
            _ = v30
        }
        _ = v26
    }
    v11 = (v11 - v11)
    if len(v2) > 0 {
        v2[0] = 0
    }
    var v32 int8 = -128
    switch ((("\n\t" + (("a" + "\n\t") + "hello")) + "") + "foo bar") {
    case "a":
        fmt.Println(v1, v1)
    case "":
        switch 127 {
        case 202:
            v11 = 18446744073709551615
        default:
            v33 := GenS1{F0: "hello"}
            _ = v33
        }
    case "\n\t":
        if (0.1 >= ((0.0 + 1e10) + 0.1)) {
            if (v1 && (!v1)) {
                fmt.Println(v4, v4)
            }
        } else {
            if false {
                v2 = append(v2, (^(v4)))
                for fi34 := 0; fi34 < 12; fi34++ {
                    var v35 int8 = (^(v32))
                    var v36 uint64 = 765
                    _ = fi34
                    _ = v35
                    _ = v36
                }
            } else {
                v37 := GenS1{F0: "\n\t"}
                _ = v37
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v5
}
