package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint16; F1 int; F2 uint16 }
func genFn0(p0 int32, p1 int64, p2 int) uint64 {
    p1 = -9223372036854775808
    var v5 float32 = 0.1
    var v6 float32 = (v5 - v5)
    var v7 string = "a"
    p0 += (p0 << ((450) & 63))
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return 18446744073709551615
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[int]uint32)
    v3 := make(map[string]string)
    var v4 bool = (4294967295 > -63)
    v8 := []int{2147483647, 32768, -61, -143}
    v9 := GenS1{F0: 831}
    v10 := v2[-822]
    v4 = v4
    v8 = append(v8, 42)
    v10 -= (^(v10))
    for fi11 := 0; fi11 < 7; fi11++ {
        var v12 uint32 = v10
        for fi13 := 0; fi13 < 4; fi13++ {
            switch fi11 {
            case 128:
                var v14 float64 = ((float64(-2.25) + -2.25) / (float64(((0.1 - 0.0) - (1e10 - 1.5))) + 1.5))
                _ = v14
            }
            _ = fi13
        }
        _ = fi11
        _ = v12
    }
    for fi15 := 0; fi15 < 9; fi15++ {
        switch (("a" + ("\n\t" + ("hello" + ("hello" + ("" + "foo bar"))))) + "foo bar") {
        case "":
            v10 -= v10
        case "\n\t":
            var v16 int = fi15
            _ = v16
        case "foo bar":
            if len(v8) >= 2 {
                v8 = v8[1:len(v8)]
            }
        }
        var v17 uint8 = 128
        for fi18 := 0; fi18 < 11; fi18++ {
            v10 = 4294967295
            switch fi15 {
            case -9223372036854775808:
                if false {
                    var v19 uint64 = 653
                    v20 := GenS2{F0: 255, F1: fi15, F2: 32767}
                    _ = v19
                    _ = v20
                }
            default:
                fmt.Println(v4, v4, v4)
            }
            _ = fi18
        }
        _ = fi15
        _ = v17
    }
    v21 := v1[v10]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
