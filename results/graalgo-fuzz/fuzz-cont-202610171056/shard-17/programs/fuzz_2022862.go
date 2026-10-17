package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 uint16 }
type GenS2 struct { F0 uint64 }

func main() {
    v1 := []string{("a" + ("\n\t" + ("foo bar" + "a"))), (("a" + "") + ("a" + ("a" + "foo bar"))), ("\n\t" + "hello"), (("" + "hello") + "foo bar")}
    var v2 float32 = (1.5 + 1.5)
    var v3 uint16 = 610
    var v4 uint8 = 42
    var v5 bool = ((0 == -1) && true)
    v6 := GenS1{F0: v3, F1: (^(v4)), F2: (v3 >> ((4294967295) & 63))}
    if false {
        fmt.Println(v3, v3)
        var v7 bool = v5
        v7 = v5
        _ = v7
    }
    v3 -= (v3 - (v3 & v3))
    var v8 float32 = float32(3.14159)
    if v5 {
        for fi9 := 0; fi9 < 13; fi9++ {
            v10 := GenS2{F0: 1}
            var v11 bool = v5
            _ = fi9
            _ = v10
            _ = v11
        }
        fmt.Println(v8)
    } else {
        v2 += v8
        fmt.Println(v5)
        for fi12 := 0; fi12 < 17; fi12++ {
            v3 = v3
            v13 := GenS2{F0: 215}
            for fi14 := 0; fi14 < 1; fi14++ {
                v5 = v5
                _ = fi14
            }
            _ = fi12
            _ = v13
        }
    }
    v2 = ((v8 * 1e10) * (v8 * ((v2 / 1.0) * 0.0)))
    switch 2147483647 {
    case 620:
        v15 := GenS2{F0: 712}
        _ = v15
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
