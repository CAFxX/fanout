package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16 }
type GenS2 struct { F0 int8; F1 int16; F2 uint16 }

func main() {
    var v1 uint8 = 0
    var v2 string = "hello"
    var v3 uint32 = 0
    var v4 uint = 540
    var v5 uint32 = v3
    v6 := []uint32{v5, v3, (v5 ^ (^(v3)))}
    v2 = (("\n\t" + v2) + ((v2 + "hello") + "foo bar"))
    if ((v1 / (((v1 << (((^(v4))) & 63))) | 1)) == 128) {
        v2 = (((v2 + v2) + (v2 + (("foo bar" + ("a" + v2)) + v2))) + ("\n\t" + ("foo bar" + ("\n\t" + ""))))
        v3 = v3
    } else {
        v4 += (v4 >> ((v4) & 63))
        _ = len(v6)
        fmt.Println(v1, v1, v3)
    }
    v7 := GenS1{F0: (65535 != 393), F1: 0}
    var v8 bool = (true && (((0.0 + ((1.5 / (((1e10 - 1.5) + 3.14159) * -2.25)) * (((-2.25 / -2.25) * 1e10) * (1.5 / 3.14159)))) / 0.1) > 1e10))
    if v8 {
        v6 = append(v6, v3)
        var v9 float32 = (((3.14159 * float32(0.0)) + (0.1 + 0.1)) * 0.1)
        if (v8 || (!v8)) {
            v10 := GenS2{F0: -112, F1: -1, F2: 505}
            var v11 string = v2
            _ = v10
            _ = v11
        }
        _ = v9
    }
    var v12 uint8 = 159
    v13 := GenS2{F0: 127, F1: -32768, F2: 42}
    v14 := GenS1{F0: v8, F1: 64}
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    _ = len(v6)
    var v15 int8 = 100
    switch v2 {
    case "\n\t":
        v8 = v8
    case "a":
        var v16 bool = v8
        _ = v16
    case "":
        switch 32767 {
        case 9223372036854775807:
            v5 += v3
        case -731:
            _ = len(v6)
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
