package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }

func main() {
    v1 := make(map[string]uint32)
    var v2 int64 = 0
    var v3 uint64 = 0
    v4 := []int{}
    v3 += v3
    v2 = v2
    var v5 int = 9223372036854775807
    for fi6 := 0; fi6 < 1; fi6++ {
        v3 = v3
        for fi7 := 0; fi7 < 16; fi7++ {
            if true {
                v8 := v1[("foo bar" + "")]
                if false {
                    var v9 uint32 = 980
                    _ = v9
                } else {
                    v10 := GenS1{F0: 14}
                    v1["foo bar"] = (v8 + v8)
                    _ = v10
                }
                _ = v8
            }
            _ = fi7
        }
        _ = fi6
    }
    var v11 bool = ((0.0 / 1e10) > 1e10)
    var v12 uint32 = 2147483647
    if (-1 < 42) {
        var v13 uint64 = (v3 + 106)
        switch -1 {
        case 128:
            v14 := GenS1{F0: 42}
            _ = v14
        case 127:
            v11 = (!(!true))
        default:
            _ = len(v1)
        }
        var v15 uint8 = 135
        _ = v13
        _ = v15
    }
    var v16 uint64 = (v3 * v3)
    fmt.Println(v16, v12)
    var v17 int32 = -1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
