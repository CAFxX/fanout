package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint; F1 uint64 }

func main() {
    var v1 uint64 = 4294967295
    var v2 int64 = -754
    v3 := make(map[string]int)
    var v4 float64 = 0.0
    v5 := make(map[uint32]string)
    var v6 uint16 = uint16(v1)
    for fi7 := 0; fi7 < 6; fi7++ {
        if false {
            var v8 float32 = (-2.25 - 1.5)
            v4 -= v4
            _ = v8
        }
        var v9 int64 = (v2 >> ((65535) & 63))
        if (true && ((false || false) && false)) {
            for fi10 := 0; fi10 < 18; fi10++ {
                var v11 uint32 = 65535
                _ = fi10
                _ = v11
            }
            v2 += (v2 - (v9 % ((v9) | 1)))
            v12 := GenS1{F0: (v9 ^ v9)}
            _ = v12
        } else {
            var v13 string = ((("\n\t" + (("a" + "\n\t") + "hello")) + ("foo bar" + "a")) + ("a" + ("a" + "hello")))
            _ = v13
        }
        _ = fi7
        _ = v9
    }
    if (4294967295 > 0) {
        var v14 uint8 = 128
        v14 -= (v14 * 128)
        v15 := GenS1{F0: (v2 >> ((uint(v14)) & 63))}
        _ = v14
        _ = v15
    } else {
        var v16 uint16 = (^(v6))
        fmt.Println(v4, v1, v4)
        var v17 int16 = 1000
        _ = v16
        _ = v17
    }
    v6 -= 844
    fmt.Println(v2, v1)
    v18 := GenS2{F0: 255, F1: (v1 + 4294967295)}
    v2 -= (v2 % ((42) | 1))
    var v19 uint16 = v6
    v5[uint32(v6)] = "foo bar"
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
