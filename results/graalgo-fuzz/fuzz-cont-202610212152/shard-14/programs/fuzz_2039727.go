package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16; F2 bool }

func main() {
    var v1 uint = 502
    var v2 uint64 = 1
    var v3 uint64 = v2
    var v4 int16 = 780
    var v5 uint8 = 128
    v5 += 127
    v6 := []int{-9223372036854775808, -893}
    for fi7 := 0; fi7 < 3; fi7++ {
        v4 = (v4 * v4)
        _ = fi7
    }
    var v8 string = (("" + "") + "a")
    switch v8 {
    case "foo bar":
        var v9 string = v8
        _ = v9
    case "a":
        v10 := GenS1{F0: -2147483648, F1: 202, F2: (0 <= (v5 % (((^(v5))) | 1)))}
        _ = v10
    }
    v3 -= (v3 | 4294967295)
    if (2147483648 < 795) {
        var v11 float32 = 0.1
        _ = v11
    } else {
        switch 862 {
        case 32768:
            v12 := GenS1{F0: 0, F1: v4, F2: true}
            _ = v12
        case -24:
            var v13 uint32 = 56
            _ = v13
        }
        v5 += (v5 << ((v1) & 63))
        fmt.Println(v4, v8, v4)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
