package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint32 }

func main() {
    var v1 int8 = 127
    var v2 uint64 = 219
    var v3 float64 = 1e10
    var v4 uint32 = 42
    v5 := make(map[uint32]string)
    var v6 string = "hello"
    for fi7 := 0; fi7 < 3; fi7++ {
        v8 := GenS1{F0: 1}
        v9 := GenS1{F0: -870}
        _ = fi7
        _ = v8
        _ = v9
    }
    var v10 int16 = 0
    var v11 int16 = (v10 + (-(v10)))
    v6 = "\n\t"
    var v12 int32 = -1
    var v13 uint = 4294967295
    v14 := GenS1{F0: v10}
    v10 = 475
    v15 := []int{9223372036854775807, 0}
    var v16 float32 = (3.14159 * 3.14159)
    var v17 int8 = (v1 - (^(v1)))
    for fi18 := 0; fi18 < 8; fi18++ {
        fmt.Println(v2)
        v15 = append(v15, (fi18 % (((fi18 | (-(fi18)))) | 1)))
        v19 := GenS2{F0: v4}
        _ = fi18
        _ = v19
    }
    var v20 int64 = 1
    switch (v6 + "hello") {
    case "hello":
        v12 = (-(v12))
    case "":
        v2 = (v2 / ((v2) | 1))
    case "\n\t":
        if true {
            v4 = v4
        } else {
            var v21 string = ((v6 + "") + v6)
            if len(v15) >= 2 {
                v15 = v15[1:len(v15)]
            }
            if len(v15) >= 2 {
                v15 = v15[1:len(v15)]
            }
            _ = v21
        }
    default:
        var v22 int32 = v12
        _ = v22
    }
    for fi23 := 0; fi23 < 14; fi23++ {
        v16 += (float32((3.14159 - v3)) * (1e10 * v16))
        v24 := GenS1{F0: (v11 - (v10 | v11))}
        _ = fi23
        _ = v24
    }
    switch (("hello" + v6) + "\n\t") {
    case "a":
        v25 := GenS1{F0: v11}
        _ = v25
    case "hello":
        var v26 string = v6
        _ = v26
    case "":
        var v27 uint16 = 0
        _ = v27
    default:
        var v28 uint = (v13 % (((^(v13))) | 1))
        _ = v28
    }
    v20 -= (^(v20))
    v29 := GenS1{F0: v10}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
