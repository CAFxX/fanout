package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool; F2 int8 }
type GenS2 struct { F0 uint32; F1 uint; F2 int16 }

func main() {
    var v1 int8 = 34
    var v2 int16 = -32768
    v3 := make(map[uint32]string)
    v4 := make(map[int]bool)
    v5 := make(map[int]uint32)
    if true {
        for fi6 := 0; fi6 < 1; fi6++ {
            v7 := GenS1{F0: 42, F1: ((!("" >= ("" + (("" + "foo bar") + "\n\t")))) || ("foo bar" >= "")), F2: v1}
            v8 := GenS1{F0: 705, F1: (true && ((v1 & 99) != (v1 + v1))), F2: (^(v1))}
            fmt.Println(v2, fi6)
            _ = fi6
            _ = v7
            _ = v8
        }
        var v9 uint = 124
        if true {
            fmt.Println(v9, v9)
            if (32767 > v2) {
                v10 := GenS2{F0: 2147483648, F1: 65535, F2: (v2 / ((42) | 1))}
                _ = v10
            }
            v11 := GenS1{F0: 65535, F1: (710 >= 0), F2: (^(v1))}
            _ = v11
        }
        _ = v9
    }
    v1 = v1
    v2 = v2
    if ((float32((float64(1.5) + 0.1)) >= 1.5) || (!false)) {
        var v12 uint16 = 279
        _ = v12
    }
    v1 = (v1 % ((v1) | 1))
    delete(v4, -2147483648)
    fmt.Println(v2)
    if (!true) {
        fmt.Println(v2, v1)
        v1 -= (-(v1))
        v5[-754] = 978
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
