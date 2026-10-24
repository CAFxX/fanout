package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }

func main() {
    v1 := make(map[int]string)
    v2 := make(map[string]int)
    var v3 bool = true
    v4 := make(map[int]int)
    var v5 int8 = -113
    v6 := []uint32{}
    var v7 float32 = 1.5
    var v8 int32 = 861
    v5 = (-(v5))
    if (!(1.5 == 1e10)) {
        if v3 {
            v8 = (v8 ^ (v8 | (v8 | (v8 >> ((358) & 63)))))
            for fi9 := 0; fi9 < 9; fi9++ {
                for fi10 := 0; fi10 < 5; fi10++ {
                    _ = len(v6)
                    _ = fi10
                }
                _ = fi9
            }
        } else {
            for fi11 := 0; fi11 < 2; fi11++ {
                v5 = v5
                v7 = v7
                _ = fi11
            }
            delete(v2, "\n\t")
        }
    }
    switch "foo bar" {
    case "a":
        if v3 {
            for fi12 := 0; fi12 < 4; fi12++ {
                v5 += -113
                for fi13 := 0; fi13 < 8; fi13++ {
                    var v14 float64 = (1e10 - (-2.25 / (0.1 * 0.1)))
                    _ = fi13
                    _ = v14
                }
                _ = fi12
            }
            v5 -= (^(v5))
        } else {
            v15 := v1[32768]
            v16 := GenS1{F0: -32768}
            _ = v15
            _ = v16
        }
    default:
        var v17 float32 = float32(-2.25)
        _ = v17
    }
    v5 += (v5 - v5)
    var v18 uint64 = 0
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    var v19 uint64 = v18
    var v20 int64 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
