package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int16 }

func main() {
    v1 := make(map[int]int)
    var v2 int16 = 1000
    v3 := []uint32{4294967295, 133, 1}
    switch (("" + ("\n\t" + ("\n\t" + ""))) + "foo bar") {
    case "":
        var v4 int32 = int32(v2)
        _ = v4
    case "foo bar":
        var v5 int8 = 42
        _ = v5
    default:
        if len(v3) > 0 {
            v3[0] = 4294967295
        }
    }
    v2 -= v2
    fmt.Println(v2)
    var v6 int = -888
    if true {
        v6 = (-(v6))
        switch 2147483647 {
        case 1:
            var v7 int8 = -6
            _ = v7
        case -288:
            v8 := GenS1{F0: 9223372036854775807}
            _ = v8
        case 436:
            for fi9 := 0; fi9 < 11; fi9++ {
                var v10 uint32 = 669
                _ = fi9
                _ = v10
            }
        default:
            var v11 int16 = 0
            _ = v11
        }
        var v12 bool = (true || (!true))
        _ = v12
    } else {
        for fi13 := 0; fi13 < 2; fi13++ {
            v2 -= v2
            v6 += (v6 >> ((2) & 63))
            v14 := GenS2{F0: v2}
            _ = fi13
            _ = v14
        }
    }
    var v15 int32 = 2147483647
    var v16 uint16 = 277
    for fi17 := 0; fi17 < 19; fi17++ {
        var v18 float64 = float64(float32(0.0))
        v19 := GenS2{F0: v2}
        v18 -= (v18 + 0.1)
        _ = fi17
        _ = v18
        _ = v19
    }
    v15 += -1
    v6 -= (v6 | (v6 >> ((42) & 63)))
    var v20 int = (v6 ^ 4294967295)
    v2 -= v2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v20
    _ = v3
    _ = v6
}
