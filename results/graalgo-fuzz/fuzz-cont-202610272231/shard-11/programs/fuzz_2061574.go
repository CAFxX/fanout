package main

import (
    "fmt"
)


func main() {
    var v1 int8 = 0
    v2 := make(map[uint32]string)
    var v3 float64 = 0.0
    v4 := []int{127}
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    switch 9223372036854775807 {
    case -1:
        v1 = v1
    default:
        for fi5 := 0; fi5 < 19; fi5++ {
            for fi6 := 0; fi6 < 16; fi6++ {
                if (42 == 0) {
                    v7 := v2[42]
                    if len(v4) > 0 {
                        v4[0] = (fi5 & (fi6 >> ((797) & 63)))
                    }
                    var v8 float32 = float32((0.1 - v3))
                    _ = v7
                    _ = v8
                } else {
                    var v9 int32 = -1
                    _ = v9
                }
                fmt.Println(fi5, fi5)
                _ = fi6
            }
            _ = fi5
        }
    }
    v3 -= 1e10
    var v10 int16 = 32767
    v1 = -105
    v4 = append(v4, 1)
    if true {
        v3 = ((v3 / ((v3 / v3) / v3)) - ((((v3 / v3) - v3) - v3) * ((-2.25 - v3) * (v3 * v3))))
        var v11 int8 = (v1 * v1)
        v11 = (v11 ^ (v1 | (v1 & v11)))
        _ = v11
    }
    for fi12 := 0; fi12 < 13; fi12++ {
        fmt.Println(v10)
        _ = fi12
    }
    v1 -= v1
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
}
