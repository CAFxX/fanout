package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint16; F1 uint8 }

func main() {
    v1 := []uint32{2}
    var v2 int64 = -9223372036854775808
    v3 := []uint32{2, 2147483648}
    for fi4 := 0; fi4 < 9; fi4++ {
        v5 := GenS1{F0: (("hello" + ("" + "\n\t")) + ((("foo bar" + "foo bar") + "\n\t") + "foo bar"))}
        v3 = append(v3, 239)
        _ = fi4
        _ = v5
    }
    if (255 != 255) {
        v2 -= (^(v2))
        var v6 uint16 = 65535
        fmt.Println(v6, v6)
        _ = v6
    }
    switch 997 {
    case -1:
        switch -857 {
        case 42:
            v1 = append(v1, 372)
        case 469:
            var v7 float64 = 0.1
            _ = v7
        default:
            switch "a" {
            case "a":
                if (false || true) {
                    v2 += -9223372036854775808
                } else {
                    v8 := GenS1{F0: "hello"}
                    _ = v8
                }
            case "\n\t":
                if len(v3) > 0 {
                    v3[0] = 840
                }
            case "foo bar":
                v9 := GenS1{F0: (("" + "") + "hello")}
                _ = v9
            default:
                v10 := GenS1{F0: ""}
                _ = v10
            }
        }
    case 0:
        fmt.Println(v2)
    case 9223372036854775807:
        for fi11 := 0; fi11 < 2; fi11++ {
            v2 = (^(v2))
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = fi11
        }
    default:
        var v12 int16 = 42
        _ = v12
    }
    fmt.Println(v2)
    var v13 uint16 = 0
    if (("foo bar" + "") != "hello") {
        v14 := GenS1{F0: ("" + "hello")}
        v15 := GenS1{F0: "foo bar"}
        _ = v14
        _ = v15
    }
    for fi16 := 0; fi16 < 3; fi16++ {
        var v17 uint16 = (v13 + v13)
        if (true || true) {
            if false {
                var v18 float64 = ((3.14159 * (0.1 + 1.5)) + 0.0)
                if (v13 <= 65535) {
                    v13 += 65535
                    var v19 int = (fi16 & fi16)
                    _ = v19
                } else {
                    if (v17 >= 42) {
                        fmt.Println(v13)
                        var v20 int64 = (-(v2))
                        v1 = append(v1, 65535)
                        _ = v20
                    }
                    for fi21 := 0; fi21 < 20; fi21++ {
                        var v22 float32 = 0.1
                        _ = fi21
                        _ = v22
                    }
                }
                var v23 uint32 = 1
                _ = v18
                _ = v23
            } else {
                switch "" {
                case "foo bar":
                    v17 = 32767
                case "hello":
                    v1 = append(v1, 833)
                default:
                    fmt.Println(v13)
                }
                for fi24 := 0; fi24 < 9; fi24++ {
                    v17 -= 32767
                    var v25 int = fi16
                    _ = fi24
                    _ = v25
                }
            }
            _ = len(v1)
        } else {
            var v26 string = "a"
            var v27 bool = (127 >= 224)
            _ = v26
            _ = v27
        }
        _ = fi16
        _ = v17
    }
    _ = len(v1)
    fmt.Println(v13, v13)
    v13 -= (^(v13))
    v28 := GenS1{F0: "\n\t"}
    fmt.Println(v2)
    fmt.Println(v13, v2)
    switch "" {
    case "":
        v3 = append(v3, 1)
    case "a":
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    default:
        v13 = v13
    }
    v29 := GenS2{F0: (^(v13)), F1: 10}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v13
    _ = v2
    _ = v28
    _ = v29
    _ = v3
}
