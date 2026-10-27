package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }

func main() {
    var v1 int64 = 0
    var v2 uint64 = 0
    var v3 bool = false
    var v4 float32 = 0.0
    var v5 bool = true
    v6 := make(map[uint32]int)
    delete(v6, 2147483648)
    if (!false) {
        var v7 uint64 = 0
        _ = v7
    } else {
        v4 = v4
    }
    if (255 < 153) {
        for fi8 := 0; fi8 < 14; fi8++ {
            v2 = v2
            v1 = (-(v1))
            _ = fi8
        }
    }
    v1 = (v1 | 560)
    if true {
        v3 = false
        var v9 int8 = 1
        v10 := GenS1{F0: 556}
        _ = v10
        _ = v9
    }
    for fi11 := 0; fi11 < 13; fi11++ {
        for fi12 := 0; fi12 < 18; fi12++ {
            if (4294967295 >= fi11) {
                switch fi11 {
                case -2147483648:
                    v3 = (fi11 != (fi12 << ((42) & 63)))
                }
            } else {
                var v13 int8 = -1
                _ = v13
            }
            v4 = float32(1e10)
            v14 := GenS1{F0: 230}
            _ = fi12
            _ = v14
        }
        _ = fi11
    }
    if false {
        var v15 int64 = (v1 | (v1 % (((v1 & v1)) | 1)))
        _ = v15
    }
    v16 := v6[1]
    fmt.Println(v3, v16)
    var v17 uint8 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
