package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32 }
type GenS2 struct { F0 string; F1 int; F2 bool }

func main() {
    var v1 int64 = 563
    var v2 int = 128
    var v3 int64 = (v1 >> ((65535) & 63))
    var v4 uint16 = 32767
    v5 := []string{"foo bar", "\n\t", ("\n\t" + "a"), ("a" + "hello")}
    v1 += (v1 % ((-198) | 1))
    if false {
        if (0.1 != 3.14159) {
            for fi6 := 0; fi6 < 15; fi6++ {
                var v7 int64 = v3
                _ = len(v5)
                _ = fi6
                _ = v7
            }
            v4 -= 255
            _ = len(v5)
        }
        v8 := GenS1{F0: -32768, F1: 42}
        _ = v8
    } else {
        switch "" {
        case "hello":
            var v9 uint16 = (^(v4))
            _ = v9
        default:
            for fi10 := 0; fi10 < 2; fi10++ {
                if (!(true && true)) {
                    var v11 int8 = 1
                    if len(v5) > 0 {
                        v5[0] = "\n\t"
                    }
                    _ = v11
                } else {
                    v5 = append(v5, "\n\t")
                }
                _ = fi10
            }
        }
        var v12 int32 = -1
        if (1.5 == (3.14159 - float64(((0.1 * (-2.25 - ((((3.14159 - 0.1) + 1.5) / float32(((3.14159 - ((-2.25 + 0.1) + 1e10)) + 1e10))) - 3.14159))) * ((((1e10 - 1e10) * 3.14159) + float32(1.5)) * ((3.14159 - 0.0) / 1.5)))))) {
            for fi13 := 0; fi13 < 8; fi13++ {
                v3 = v1
                _ = fi13
            }
            var v14 uint = 2
            _ = v14
        }
        _ = v12
    }
    if (42 >= 4294967295) {
        if len(v5) > 0 {
            v5[0] = "foo bar"
        }
    }
    if len(v5) > 0 {
        v5[0] = "hello"
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v15 := GenS2{F0: ("\n\t" + (("\n\t" + (("" + ((("foo bar" + "foo bar") + "\n\t") + "\n\t")) + "a")) + "\n\t")), F1: v2, F2: (-1 > 168)}
    v16 := GenS1{F0: 71, F1: 2147483647}
    var v17 uint8 = 133
    switch ("\n\t" + "hello") {
    case "":
        if false {
            var v18 int16 = -1
            _ = v18
        } else {
            v1 = 1
        }
    case "a":
        v3 += v3
    case "\n\t":
        _ = len(v5)
    default:
        var v19 int64 = (-(v3))
        _ = v19
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
