package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }

func main() {
    v1 := make(map[string]int)
    var v2 bool = true
    var v3 float32 = (-2.25 * 1.5)
    var v4 uint = 0
    fmt.Println(v4, v3, v2)
    var v5 float32 = -2.25
    var v6 uint64 = 4294967295
    var v7 int8 = 100
    if false {
        v8 := GenS1{F0: 533}
        var v9 uint32 = 65535
        _ = v8
        _ = v9
    }
    v2 = v2
    v7 = v7
    var v10 float32 = v5
    v2 = v2
    v4 += (v4 * (v4 & (v4 & v4)))
    v11 := []int{9223372036854775807}
    for fi12 := 0; fi12 < 4; fi12++ {
        if ((2 > (v6 >> ((18446744073709551615) & 63))) && ((v5 <= v10) && (fi12 >= 127))) {
            var v13 int = -443
            _ = v13
        }
        _ = fi12
    }
    var v14 bool = (65535 >= 590)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
