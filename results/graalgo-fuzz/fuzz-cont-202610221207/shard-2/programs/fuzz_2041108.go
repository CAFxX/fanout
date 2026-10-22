package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 string }

func main() {
    v1 := make(map[uint32]bool)
    var v2 float32 = 3.14159
    var v3 int64 = 42
    v4 := []string{"\n\t", "a", "hello"}
    var v5 int = -1
    var v6 bool = (v5 < 4294967295)
    for fi7 := 0; fi7 < 4; fi7++ {
        fmt.Println(fi7, v3)
        var v8 int = -2147483648
        _ = fi7
        _ = v8
    }
    _ = len(v1)
    if len(v4) > 0 {
        v4[0] = ("foo bar" + "hello")
    }
    if (1.5 <= 0.0) {
        v5 -= v5
        if v6 {
            v9 := GenS1{F0: ("\n\t" + ("hello" + ("" + ("hello" + "a")))), F1: 18446744073709551615}
            v10 := GenS1{F0: "foo bar", F1: 2}
            v11 := GenS1{F0: ("\n\t" + "hello"), F1: 42}
            _ = v10
            _ = v11
            _ = v9
        } else {
            if (false && (!((!false) || v6))) {
                var v12 float32 = (float32(((1e10 + 3.14159) + ((((1.5 - (0.0 / -2.25)) * 1e10) * (1e10 - -2.25)) / 1.0))) + (((0.0 * 0.1) * -2.25) + (v2 - ((v2 / v2) - v2))))
                v13 := GenS1{F0: ("" + "a"), F1: 1}
                switch "foo bar" {
                case "\n\t":
                    if v6 {
                        _ = len(v1)
                    }
                case "hello":
                    v2 -= v2
                case "":
                    fmt.Println(v3, v6)
                default:
                    fmt.Println(v12, v3, v12)
                }
                _ = v12
                _ = v13
            } else {
                switch "\n\t" {
                case "a":
                    v14 := v1[65535]
                    _ = v14
                case "hello":
                    var v15 uint64 = 4294967295
                    _ = v15
                }
            }
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        }
    }
    v6 = true
    switch ("foo bar" + "a") {
    case "hello":
        var v16 int64 = v3
        _ = v16
    case "":
        _ = len(v4)
    case "\n\t":
        if v6 {
            v17 := GenS2{F0: -92}
            _ = v17
        } else {
            for fi18 := 0; fi18 < 19; fi18++ {
                v6 = true
                if (!(!false)) {
                    if len(v4) > 0 {
                        v4[0] = ""
                    }
                }
                _ = fi18
            }
            v19 := GenS3{F0: "\n\t"}
            v4 = append(v4, "a")
            _ = v19
        }
    }
    if (2147483647 > -700) {
        if (true && v6) {
            for fi20 := 0; fi20 < 9; fi20++ {
                _ = len(v1)
                _ = fi20
            }
        }
        v21 := GenS3{F0: ("" + ("" + "foo bar"))}
        v22 := GenS2{F0: 77}
        _ = v21
        _ = v22
    } else {
        v2 += float32(-2.25)
    }
    v23 := GenS1{F0: ("\n\t" + "foo bar"), F1: 4294967295}
    _ = len(v4)
    var v24 int64 = v3
    v2 -= ((((0.1 - 1.5) - float32(3.14159)) + v2) / v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
