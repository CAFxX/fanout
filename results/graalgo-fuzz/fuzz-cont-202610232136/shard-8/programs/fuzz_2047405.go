package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int8 }

func main() {
    v1 := make(map[uint32]int)
    var v2 bool = (0.1 == 0.0)
    var v3 uint16 = 284
    v4 := []int{-450, 163, 829}
    var v5 int64 = -1
    var v6 uint8 = 1
    fmt.Println(v6)
    fmt.Println(v3)
    v6 = 127
    v2 = v2
    v2 = ((1e10 * ((3.14159 / -2.25) * 0.1)) <= (((-2.25 - float32(0.1)) * (0.1 / 3.14159)) - ((-2.25 + (((-2.25 / (1.5 - ((1e10 / (float32((0.0 - float64((0.1 * (((((0.1 * 0.0) - 0.1) / 1.0) + float32(1e10)) * 3.14159))))) * (float32(3.14159) * 0.1))) - 1.5))) * 0.0) - (float32(1.5) * 1.5))) + float32(0.0))))
    v7 := GenS1{F0: -125}
    v8 := GenS2{F0: -128}
    v3 = v3
    for fi9 := 0; fi9 < 8; fi9++ {
        var v10 string = ("a" + "a")
        fmt.Println(v5, v5, v5)
        _ = fi9
        _ = v10
    }
    fmt.Println(v6, v6, v2)
    if false {
        for fi11 := 0; fi11 < 19; fi11++ {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = fi11
        }
    }
    _ = len(v4)
    v3 += v3
    fmt.Println(v6)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
