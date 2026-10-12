package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32 }
func genFn0(p0 string, p1 int64, p2 int8) uint {
    v6 := []int{}
    fmt.Println(p0, p0)
    switch p0 {
    case "a":
        for fi7 := 0; fi7 < 11; fi7++ {
            var v8 string = p0
            if (!true) {
                var v9 int32 = 0
                if (-178 <= 1) {
                    var v10 int8 = (p2 | 42)
                    _ = v10
                }
                v11 := GenS1{F0: 759, F1: -956}
                _ = v11
                _ = v9
            }
            _ = fi7
            _ = v8
        }
    default:
        v6 = append(v6, 9223372036854775807)
    }
    switch "" {
    case "\n\t":
        var v12 uint64 = 193
        _ = v12
    case "":
        _ = len(v6)
    default:
        p0 = p0
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 65535
}

func main() {
    v1 := []int{128, 9223372036854775807}
    var v2 uint16 = 65535
    v3 := make(map[int]string)
    v4 := []int{}
    var v5 uint32 = uint32(v2)
    for fi13 := 0; fi13 < 5; fi13++ {
        v5 -= 573
        var v14 float32 = 1.5
        _ = fi13
        _ = v14
    }
    _ = len(v1)
    if false {
        var v15 uint64 = 0
        _ = v15
    }
    for fi16 := 0; fi16 < 3; fi16++ {
        switch "foo bar" {
        case "":
            fmt.Println(v2, v2)
        }
        var v17 int = (fi16 << ((2) & 63))
        v2 += 1
        _ = fi16
        _ = v17
    }
    v5 = (v5 / (((^(v5))) | 1))
    v2 += (v2 ^ 42)
    v18 := v3[-546]
    var v19 int32 = 1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
