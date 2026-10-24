package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }

func main() {
    var v1 int64 = 1
    var v2 float32 = 3.14159
    var v3 int8 = -111
    var v4 uint16 = 0
    var v5 int64 = (-(v1))
    v6 := make(map[uint32]string)
    fmt.Println(v2, v4)
    var v7 uint16 = (v4 << ((1) & 63))
    var v8 int32 = 0
    switch 9223372036854775807 {
    case -1:
        var v9 float32 = v2
        _ = v9
    }
    v4 = v4
    v4 += v7
    v2 = v2
    var v10 string = ("" + ("\n\t" + "foo bar"))
    if false {
        var v11 uint16 = v4
        _ = v11
    }
    v12 := GenS1{F0: uint32(v7)}
    var v13 int32 = v8
    v14 := GenS1{F0: 2147483647}
    if true {
        v1 -= v1
    }
    if true {
        fmt.Println(v1)
        if (-1 < 42) {
            v15 := GenS1{F0: 164}
            _ = v15
        }
        v16 := GenS1{F0: 42}
        _ = v16
    }
    v17 := []int{1, -468, 597, 42}
    v3 = v3
    if len(v17) > 0 {
        v17[0] = 32768
    }
    switch int(v5) {
    case -2147483648:
        switch (v10 + v10) {
        case "hello":
            v18 := GenS1{F0: 2147483647}
            _ = v18
        default:
            v10 = "\n\t"
        }
    default:
        v1 += v1
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v14
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
