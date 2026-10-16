package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8 }

func main() {
    v1 := make(map[string]bool)
    v2 := []int{}
    var v3 uint32 = 42
    var v4 uint64 = 18446744073709551615
    var v5 uint16 = 0
    switch "foo bar" {
    case "":
        for fi6 := 0; fi6 < 16; fi6++ {
            var v7 float32 = float32(float64(1.5))
            v3 = 2147483648
            _ = fi6
            _ = v7
        }
    }
    v4 += 0
    var v8 string = (("a" + "hello") + "foo bar")
    v3 += v3
    if (0 <= -868) {
        v5 = (v5 + uint16(v3))
        switch 42 {
        case 652:
            v9 := GenS1{F0: v4, F1: 0}
            _ = v9
        case 9223372036854775807:
            v3 -= v3
        case 127:
            _ = len(v2)
        default:
            v4 = (^(v4))
        }
    }
    if (true && false) {
        if false {
            v8 = (v8 + v8)
        } else {
            fmt.Println(v8)
        }
    }
    fmt.Println(v5)
    if true {
        if false {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        } else {
            v10 := GenS1{F0: v4, F1: 1}
            if (1 < 2147483647) {
                var v11 int64 = 9223372036854775807
                var v12 float32 = 0.1
                switch 32768 {
                case 344:
                    v4 = v4
                case 32768:
                    for fi13 := 0; fi13 < 18; fi13++ {
                        for fi14 := 0; fi14 < 9; fi14++ {
                            v15 := v1[v8]
                            _ = fi14
                            _ = v15
                        }
                        _ = fi13
                    }
                }
                _ = v11
                _ = v12
            }
            _ = v10
        }
        v4 += v4
        delete(v1, v8)
    } else {
        var v16 int16 = -1
        if len(v2) > 0 {
            v2[0] = -9223372036854775808
        }
        var v17 int8 = -65
        _ = v16
        _ = v17
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
