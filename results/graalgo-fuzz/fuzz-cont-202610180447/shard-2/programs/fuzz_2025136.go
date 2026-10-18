package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint8 }

func main() {
    v1 := make(map[uint32]string)
    var v2 uint64 = 209
    var v3 int8 = 1
    v3 += v3
    if false {
        v4 := GenS1{F0: 240}
        _ = v4
    }
    for fi5 := 0; fi5 < 18; fi5++ {
        var v6 float64 = ((1.5 - 1e10) - 0.0)
        v7 := GenS1{F0: 127}
        _ = fi5
        _ = v6
        _ = v7
    }
    fmt.Println(v2, v3)
    switch ("\n\t" + "") {
    case "\n\t":
        for fi8 := 0; fi8 < 6; fi8++ {
            v9 := GenS1{F0: 0}
            v10 := GenS2{F0: 5}
            v3 = v3
            _ = fi8
            _ = v10
            _ = v9
        }
    default:
        var v11 int = 4294967295
        _ = v11
    }
    fmt.Println(v3, v2)
    v2 -= v2
    fmt.Println(v3, v3)
    v12 := []int{-107, 32767, 127, int(v2)}
    var v13 float64 = (0.0 + (0.0 - -2.25))
    var v14 uint32 = 2
    for fi15 := 0; fi15 < 1; fi15++ {
        var v16 uint64 = v2
        v13 += (v13 - v13)
        if (1 <= 517) {
            if (!((((!true) && (-124 <= 0)) || (!(!(!false)))) || true)) {
                var v17 float64 = (1e10 / float64((1e10 + 1e10)))
                var v18 int = fi15
                _ = v17
                _ = v18
            }
            var v19 int16 = -32768
            var v20 uint8 = 127
            _ = v19
            _ = v20
        } else {
            v16 -= v2
            fmt.Println(v14, v2, fi15)
        }
        _ = fi15
        _ = v16
    }
    v12 = append(v12, 2147483647)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
}
