package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint; F2 bool }
type GenS2 struct { F0 int64 }

func main() {
    var v1 uint32 = 51
    v2 := make(map[string]int)
    var v3 uint8 = 127
    var v4 int8 = -1
    if true {
        if false {
            var v5 bool = (42 > 42)
            fmt.Println(v1, v3, v4)
            v6 := GenS1{F0: "a", F1: 452, F2: true}
            _ = v5
            _ = v6
        } else {
            var v7 int8 = v4
            v3 = (v3 * v3)
            _ = v7
        }
        fmt.Println(v1, v4, v1)
        v2["hello"] = -876
    }
    switch -1 {
    case 2147483647:
        switch "a" {
        case "\n\t":
            if false {
                var v8 uint16 = 1
                var v9 float32 = (0.0 / ((0.0 / 1.0) - (1e10 + (float32(0.1) / float32(-2.25)))))
                _ = v8
                _ = v9
            } else {
                v10 := GenS1{F0: "a", F1: 65535, F2: false}
                _ = v10
            }
        case "hello":
            for fi11 := 0; fi11 < 1; fi11++ {
                v12 := v2["a"]
                for fi13 := 0; fi13 < 14; fi13++ {
                    fmt.Println(fi13, fi13, v1)
                    v4 -= v4
                    var v14 int32 = 770
                    _ = fi13
                    _ = v14
                }
                _ = fi11
                _ = v12
            }
        default:
            var v15 uint16 = 730
            _ = v15
        }
    }
    v16 := GenS2{F0: 191}
    fmt.Println(v4, v3)
    v17 := GenS2{F0: 520}
    for fi18 := 0; fi18 < 3; fi18++ {
        v19 := GenS1{F0: "", F1: 18446744073709551615, F2: ("\n\t" >= "")}
        var v20 string = "\n\t"
        _ = fi18
        _ = v19
        _ = v20
    }
    v4 -= v4
    v4 -= (-(v4))
    fmt.Println(v4, v4, v3)
    var v21 int32 = -2147483648
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v21
    _ = v3
    _ = v4
}
