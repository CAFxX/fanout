package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint16; F1 uint16 }

func main() {
    var v1 float32 = 1e10
    v2 := make(map[uint32]int)
    var v3 int64 = 1
    v4 := []string{"\n\t"}
    var v5 uint16 = 895
    var v6 uint32 = 0
    var v7 uint64 = 319
    if (("a" == (("a" + (((("hello" + "") + ("\n\t" + (("hello" + "") + ("\n\t" + ("a" + "\n\t"))))) + "hello") + "a")) + "")) || true) {
        v3 -= (v3 / (((v3 | (^(v3)))) | 1))
    } else {
        var v8 uint64 = 1
        v9 := GenS1{F0: -540}
        fmt.Println(v6, v1, v1)
        _ = v8
        _ = v9
    }
    var v10 uint64 = (v7 | v7)
    v11 := GenS1{F0: 2147483647}
    fmt.Println(v5, v3, v3)
    fmt.Println(v10, v6, v5)
    v6 -= v6
    var v12 int = 0
    var v13 int = (^(v12))
    if false {
        var v14 int = -214
        delete(v2, v6)
        v15 := GenS2{F0: v5, F1: (v5 - uint16(v6))}
        _ = v14
        _ = v15
    } else {
        v1 -= (3.14159 / -2.25)
        if (!(0.1 == (((0.1 - -2.25) / (3.14159 * (3.14159 - (1.5 + 0.1)))) - ((float64(3.14159) * -2.25) / (((1.5 / float64(v1)) / 3.14159) + -2.25))))) {
            var v16 uint16 = 32767
            _ = len(v4)
            v17 := GenS2{F0: (v16 / (((v5 % ((v5) | 1))) | 1)), F1: v16}
            _ = v16
            _ = v17
        } else {
            v18 := GenS2{F0: (v5 << ((370) & 63)), F1: (^(v5))}
            _ = v18
        }
    }
    v12 += 32768
    for fi19 := 0; fi19 < 19; fi19++ {
        v3 = 42
        if (94 > 320) {
            var v20 uint32 = uint32(v7)
            _ = v20
        }
        v21 := GenS2{F0: (v5 << ((1) & 63)), F1: v5}
        _ = fi19
        _ = v21
    }
    v1 = 1e10
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
