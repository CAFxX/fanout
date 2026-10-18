package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 string }

func main() {
    var v1 uint64 = 2
    var v2 int16 = -1
    var v3 bool = false
    var v4 int64 = -9223372036854775808
    var v5 float32 = (0.0 / -2.25)
    v3 = true
    v5 -= (0.1 * (v5 + (v5 / 0.1)))
    for fi6 := 0; fi6 < 8; fi6++ {
        v1 = 42
        _ = fi6
    }
    v7 := GenS1{F0: 1}
    v8 := []int{128, 4294967295, 32768}
    var v9 uint8 = 2
    var v10 bool = false
    var v11 int = 1
    v12 := GenS2{F0: ""}
    for fi13 := 0; fi13 < 17; fi13++ {
        var v14 int64 = 42
        _ = fi13
        _ = v14
    }
    var v15 bool = v3
    v1 += v1
    fmt.Println(v10, v15)
    v16 := GenS1{F0: -9223372036854775808}
    var v17 uint64 = v1
    v18 := GenS2{F0: "\n\t"}
    v19 := GenS2{F0: ("\n\t" + "hello")}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
