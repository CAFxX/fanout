package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }

func main() {
    v1 := make(map[int]int)
    var v2 int8 = -1
    v3 := make(map[string]uint32)
    var v4 int32 = 369
    v5 := []uint32{}
    var v6 uint16 = 0
    v6 = (v6 >> ((2) & 63))
    var v7 int16 = 0
    var v8 int64 = -971
    var v9 uint16 = v6
    fmt.Println(v6, v7)
    v10 := v3["foo bar"]
    v4 -= (-(v4))
    if true {
        var v11 uint32 = v10
        _ = v11
    } else {
        if false {
            var v12 float32 = 1e10
            var v13 int8 = v2
            if len(v5) > 0 {
                v5[0] = v10
            }
            _ = v12
            _ = v13
        }
    }
    switch 32768 {
    case 127:
        v10 = v10
    }
    if (uint(v6) >= 995) {
        if len(v5) > 0 {
            v5[0] = (v10 | (^(v10)))
        }
        var v14 int32 = (v4 ^ (-(v4)))
        v15 := GenS1{F0: (v2 + (v2 % ((-1) | 1)))}
        _ = v14
        _ = v15
    } else {
        if len(v5) > 0 {
            v5[0] = (v10 ^ v10)
        }
    }
    for fi16 := 0; fi16 < 19; fi16++ {
        if (((v6 % ((v6) | 1)) <= v6) || false) {
            fmt.Println(v2, fi16)
            v17 := GenS1{F0: v2}
            switch 9223372036854775807 {
            case 4294967295:
                switch 2147483647 {
                case 1:
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                case 128:
                    v8 -= -9223372036854775808
                }
            }
            _ = v17
        }
        _ = fi16
    }
    v10 += v10
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
