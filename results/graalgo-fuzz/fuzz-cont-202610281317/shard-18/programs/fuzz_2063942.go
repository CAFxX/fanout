package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64; F2 uint8 }

func main() {
    v1 := []string{((("hello" + ("" + "hello")) + "\n\t") + "hello"), ("\n\t" + ("foo bar" + "foo bar")), "a", "hello"}
    v2 := make(map[uint32]uint32)
    v3 := []string{"hello"}
    var v4 uint64 = 42
    v5 := []string{"", "foo bar", ("" + "a"), (("a" + "foo bar") + "\n\t")}
    var v6 int = -80
    v4 = v4
    var v7 uint64 = (v4 % ((v4) | 1))
    for fi8 := 0; fi8 < 18; fi8++ {
        var v9 float32 = 3.14159
        v6 = -81
        v9 = v9
        _ = fi8
        _ = v9
    }
    for fi10 := 0; fi10 < 12; fi10++ {
        v4 -= 15
        v2[4294967295] = 4294967295
        var v11 uint8 = 42
        _ = fi10
        _ = v11
    }
    v6 -= 489
    _ = len(v3)
    if ((428 == 32767) && false) {
        if false {
            v6 -= 9223372036854775807
            var v12 uint8 = 112
            _ = v12
        }
        switch (v6 - v6) {
        case 127:
            fmt.Println(v4, v4)
        default:
            if false {
                switch v6 {
                case 127:
                    v4 -= (v4 % (((v7 / ((v4) | 1))) | 1))
                case 0:
                    v6 -= 294
                default:
                    v13 := GenS1{F0: 1, F1: (v4 | 18446744073709551615), F2: 0}
                    _ = v13
                }
            }
        }
        var v14 int8 = 0
        _ = v14
    } else {
        v4 = (v7 << ((0) & 63))
    }
    v15 := GenS1{F0: 127, F1: (v7 ^ (v7 >> ((0) & 63))), F2: 211}
    var v16 float64 = -2.25
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
