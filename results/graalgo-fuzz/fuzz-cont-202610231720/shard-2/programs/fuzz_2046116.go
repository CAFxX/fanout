package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16 }

func main() {
    var v1 uint16 = 102
    var v2 string = "a"
    var v3 int32 = -2147483648
    v4 := make(map[uint32]uint32)
    v5 := []int{128, 0, 128, -364}
    v3 = -506
    fmt.Println(v2, v2, v1)
    if ((42 > 2147483648) && ("foo bar" >= v2)) {
        var v6 uint = 1
        _ = v6
    }
    for fi7 := 0; fi7 < 11; fi7++ {
        var v8 float32 = (-2.25 + float32(3.14159))
        v1 = (^(v1))
        _ = fi7
        _ = v8
    }
    var v9 uint32 = 65535
    var v10 int = 1
    v11 := GenS1{F0: 742, F1: 1000}
    var v12 uint32 = v9
    var v13 uint8 = 2
    var v14 int8 = 100
    v4[65535] = (v9 ^ (^(v9)))
    delete(v4, (v9 & 58))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
