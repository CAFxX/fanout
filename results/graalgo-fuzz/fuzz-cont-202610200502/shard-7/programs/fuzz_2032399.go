package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint }

func main() {
    var v1 float32 = float32(float64(1.5))
    var v2 bool = false
    var v3 float32 = -2.25
    var v4 uint16 = 563
    var v5 float64 = -2.25
    var v6 uint64 = 1
    for fi7 := 0; fi7 < 11; fi7++ {
        v8 := GenS1{F0: 42}
        _ = fi7
        _ = v8
    }
    var v9 uint32 = 1
    switch 128 {
    case -9223372036854775808:
        v3 = v3
    }
    v10 := []int{32767, -857, 32767}
    if v2 {
        v2 = (185 >= 128)
        var v11 uint16 = v4
        if len(v10) >= 2 {
            v10 = v10[1:len(v10)]
        }
        _ = v11
    } else {
        v6 = v6
    }
    fmt.Println(v2, v9)
    v4 = (v4 + (v4 >> ((255) & 63)))
    v12 := GenS1{F0: 376}
    var v13 bool = (v2 && v2)
    if len(v10) >= 2 {
        v10 = v10[1:len(v10)]
    }
    v5 = ((((v5 - ((v5 * v5) + (v5 - (float64((3.14159 + v1)) - 0.0)))) * v5) * 1e10) - 3.14159)
    var v14 int64 = -510
    var v15 int64 = v14
    var v16 int32 = -525
    v17 := GenS2{F0: (v14 << ((2) & 63))}
    v18 := GenS3{F0: 4294967295}
    if v13 {
        for fi19 := 0; fi19 < 9; fi19++ {
            var v20 uint64 = (v6 + v6)
            _ = fi19
            _ = v20
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
