package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32; F2 uint8 }
type GenS2 struct { F0 uint32; F1 int; F2 bool }

func main() {
    var v1 int32 = -2147483648
    var v2 int = 164
    var v3 int32 = (-(v1))
    var v4 uint8 = 183
    fmt.Println(v2)
    v3 = 1
    v5 := []int{v2, (v2 >> ((0) & 63))}
    fmt.Println(v2)
    var v6 uint64 = 2
    v1 -= v3
    v1 += v1
    if (!("\n\t" < "\n\t")) {
        if len(v5) > 0 {
            v5[0] = v2
        }
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    }
    for fi7 := 0; fi7 < 1; fi7++ {
        v3 -= v3
        _ = fi7
    }
    v8 := GenS1{F0: (0 != -1), F1: (v3 ^ (v3 * v3)), F2: uint8(v6)}
    fmt.Println(v4, v3, v2)
    v9 := GenS1{F0: (!false), F1: v1, F2: 42}
    if true {
        if true {
            v10 := GenS2{F0: 1, F1: (v2 * 32767), F2: (v2 >= -1)}
            var v11 bool = (255 == 1)
            _ = v10
            _ = v11
        }
    } else {
        switch 2147483647 {
        case 913:
            fmt.Println(v1)
        default:
            var v12 string = "foo bar"
            _ = v12
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
