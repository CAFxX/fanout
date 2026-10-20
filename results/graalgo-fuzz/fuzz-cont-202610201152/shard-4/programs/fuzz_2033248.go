package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint; F1 uint8; F2 int }

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 float32 = ((0.1 + (1.5 - 0.1)) * -2.25)
    var v3 int = 32767
    v4 := make(map[int]string)
    for fi5 := 0; fi5 < 13; fi5++ {
        var v6 int32 = 42
        var v7 float64 = 0.1
        switch v3 {
        case 398:
            v3 = -9223372036854775808
        default:
            for fi8 := 0; fi8 < 12; fi8++ {
                v1 = 730
                v6 = int32(v1)
                _ = fi8
            }
        }
        _ = fi5
        _ = v6
        _ = v7
    }
    switch "foo bar" {
    case "\n\t":
        for fi9 := 0; fi9 < 14; fi9++ {
            for fi10 := 0; fi10 < 16; fi10++ {
                var v11 string = "a"
                _ = fi10
                _ = v11
            }
            _ = fi9
        }
    case "hello":
        fmt.Println(v2, v1, v1)
    default:
        var v12 string = ("\n\t" + ("foo bar" + "foo bar"))
        _ = v12
    }
    if ((((("foo bar" + "") <= "") || false) || true) && false) {
        fmt.Println(v3, v3)
    }
    fmt.Println(v3, v1)
    v2 += (3.14159 * v2)
    for fi13 := 0; fi13 < 14; fi13++ {
        var v14 uint64 = (v1 * (v1 << ((0) & 63)))
        if false {
            switch (fi13 - v3) {
            case -199:
                if true {
                    v15 := GenS1{F0: 42}
                    _ = v15
                } else {
                    fmt.Println(fi13, fi13)
                    v3 -= fi13
                }
            case 9223372036854775807:
                v1 = (v14 & (v14 % (((v14 - v14)) | 1)))
            case -9223372036854775808:
                v4[fi13] = ("\n\t" + ((("hello" + ("" + "")) + "") + ""))
            default:
                delete(v4, 128)
            }
        }
        _ = fi13
        _ = v14
    }
    v3 -= v3
    v16 := GenS2{F0: 255, F1: 127, F2: (v3 | v3)}
    v4[v3] = (("a" + "\n\t") + "a")
    v17 := GenS1{F0: -63}
    if true {
        var v18 int16 = int16(v1)
        v18 += v18
        _ = v18
    }
    delete(v4, 32768)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
