package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8 }
func genFn0(p0 int32, p1 int8, p2 uint8) int {
    var v4 uint32 = 2
    var v5 uint32 = v4
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 2147483647
}

func main() {
    var v1 int8 = 43
    v2 := []uint32{2147483647, 624, 845}
    v3 := []string{(("a" + "hello") + "\n\t")}
    switch -9223372036854775808 {
    case 32768:
        switch 1 {
        case 2147483647:
            var v6 uint8 = 255
            _ = v6
        default:
            for fi7 := 0; fi7 < 12; fi7++ {
                for fi8 := 0; fi8 < 11; fi8++ {
                    if true {
                        if len(v3) > 0 {
                            v3[0] = (("" + "foo bar") + ("a" + ""))
                        }
                    } else {
                        var v9 uint = 42
                        if len(v2) > 0 {
                            v2[0] = 1
                        }
                        _ = v9
                    }
                    if (42 > 0) {
                        var v10 int32 = -600
                        fmt.Println(v10, v1, v1)
                        v11 := GenS1{F0: 127, F1: (v1 << ((255) & 63))}
                        _ = v10
                        _ = v11
                    }
                    v2 = append(v2, 441)
                    _ = fi8
                }
                _ = fi7
            }
        }
    case 1:
        for fi12 := 0; fi12 < 11; fi12++ {
            _ = len(v2)
            v1 += (v1 - (v1 | (v1 >> ((893) & 63))))
            _ = fi12
        }
    }
    for fi13 := 0; fi13 < 19; fi13++ {
        _ = len(v3)
        v14 := GenS1{F0: v1, F1: v1}
        _ = fi13
        _ = v14
    }
    fmt.Println(v1)
    if false {
        var v15 int32 = -1
        var v16 int64 = 0
        _ = v15
        _ = v16
    }
    v1 = v1
    var v17 uint64 = 19
    fmt.Println(v1, v17)
    v1 += 42
    if len(v3) > 0 {
        v3[0] = "foo bar"
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
}
