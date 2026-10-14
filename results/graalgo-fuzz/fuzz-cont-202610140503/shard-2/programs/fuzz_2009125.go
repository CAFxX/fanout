package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16; F2 uint8 }
type GenS2 struct { F0 bool }
func genFn0(p0 int64) uint32 {
    v7 := GenS1{F0: 128, F1: 496, F2: 127}
    p0 += (p0 % (((^(p0))) | 1))
    _ = p0
    _ = v7
    return 672
}

func main() {
    var v1 uint32 = 0
    var v2 uint8 = 186
    v3 := []uint32{(v1 >> ((2) & 63)), 699}
    var v4 uint64 = 2
    v5 := []int{}
    v6 := make(map[string]int)
    var v8 bool = true
    if len(v5) > 0 {
        v5[0] = 355
    }
    v4 = (v4 - (v4 >> ((42) & 63)))
    switch -1 {
    case 699:
        var v9 uint64 = 18446744073709551615
        _ = v9
    case 128:
        if (34 < -77) {
            if ("a" < "a") {
                v4 -= (v4 >> ((314) & 63))
                v10 := GenS2{F0: ((((1e10 * -2.25) - 1.5) + float64(-2.25)) < ((0.1 + (float64(float32((3.14159 - 1e10))) / 3.14159)) + (-2.25 - 1e10)))}
                _ = v10
            } else {
                var v11 uint16 = 2
                for fi12 := 0; fi12 < 16; fi12++ {
                    var v13 int = fi12
                    for fi14 := 0; fi14 < 15; fi14++ {
                        var v15 int64 = 1
                        var v16 uint32 = 65535
                        _ = fi14
                        _ = v15
                        _ = v16
                    }
                    _ = fi12
                    _ = v13
                }
                _ = v11
            }
            v1 += 4294967295
        }
    }
    fmt.Println(v1, v1)
    var v17 int16 = 32767
    v17 += v17
    for fi18 := 0; fi18 < 5; fi18++ {
        v19 := GenS1{F0: fi18, F1: -417, F2: 0}
        var v20 uint64 = (v4 * v4)
        if v8 {
            var v21 uint = 65535
            _ = v21
        }
        _ = fi18
        _ = v19
        _ = v20
    }
    v4 = 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
