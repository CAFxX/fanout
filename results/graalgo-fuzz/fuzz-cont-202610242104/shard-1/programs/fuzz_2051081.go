package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint16; F1 uint16 }

func main() {
    var v1 int64 = -757
    var v2 uint = 2
    var v3 int = 2147483647
    v4 := []uint32{433, 42, 657}
    v5 := make(map[string]uint32)
    var v6 uint16 = 32767
    for fi7 := 0; fi7 < 11; fi7++ {
        var v8 string = "hello"
        var v9 float32 = ((3.14159 + 1.5) / 0.1)
        for fi10 := 0; fi10 < 1; fi10++ {
            v11 := GenS1{F0: -32768}
            fmt.Println(v3, v8, v8)
            switch fi7 {
            case -764:
                v8 = ("foo bar" + v8)
            case 4294967295:
                v2 += (v2 % ((2) | 1))
            default:
                var v12 int16 = 251
                _ = v12
            }
            _ = fi10
            _ = v11
        }
        _ = fi7
        _ = v8
        _ = v9
    }
    var v13 float32 = 0.0
    v14 := GenS1{F0: 158}
    v5[(("a" + "foo bar") + ("foo bar" + ("\n\t" + "hello")))] = uint32(v2)
    v2 = 731
    var v15 bool = false
    v15 = ((((v2 >> (((^(v2))) & 63)) == (^(v2))) || false) || true)
    if v15 {
        v16 := GenS2{F0: v6, F1: (v6 ^ v6)}
        v4 = append(v4, 246)
        _ = v16
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
