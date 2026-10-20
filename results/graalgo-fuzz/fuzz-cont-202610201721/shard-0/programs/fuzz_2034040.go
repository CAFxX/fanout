package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }

func main() {
    var v1 uint = 65535
    var v2 int32 = 904
    var v3 int16 = 1000
    var v4 bool = false
    v5 := []uint32{}
    var v6 int32 = v2
    if v4 {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        v5 = append(v5, 1)
        if true {
            for fi7 := 0; fi7 < 20; fi7++ {
                for fi8 := 0; fi8 < 11; fi8++ {
                    v5 = append(v5, 65535)
                    _ = fi8
                }
                _ = fi7
            }
            v2 += 2147483647
            v1 += (v1 * v1)
        }
    } else {
        var v9 int64 = 0
        _ = v9
    }
    v3 = (v3 * v3)
    v4 = true
    v6 = (v6 << (((^(v1))) & 63))
    v3 -= v3
    var v10 float64 = 0.0
    var v11 uint16 = 32767
    for fi12 := 0; fi12 < 1; fi12++ {
        if len(v5) > 0 {
            v5[0] = 991
        }
        _ = fi12
    }
    v13 := GenS1{F0: 397}
    fmt.Println(v4)
    var v14 string = "a"
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
