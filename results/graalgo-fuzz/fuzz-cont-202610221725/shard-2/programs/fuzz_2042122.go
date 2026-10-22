package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }

func main() {
    var v1 int32 = 0
    var v2 uint64 = 435
    var v3 uint8 = 127
    v4 := make(map[string]uint32)
    var v5 float32 = 3.14159
    v6 := GenS1{F0: 57}
    v1 = (v1 - v1)
    var v7 int32 = v1
    var v8 uint8 = v3
    v9 := GenS1{F0: -375}
    v3 = (v8 << ((88) & 63))
    switch -9223372036854775808 {
    case 9223372036854775807:
        fmt.Println(v3)
    case 2147483647:
        fmt.Println(v1)
    }
    v10 := []int{127, -2147483648, 2147483647, -1}
    for fi11 := 0; fi11 < 12; fi11++ {
        fmt.Println(v2)
        _ = fi11
    }
    if false {
        fmt.Println(v8, v5)
        v12 := v4[(("a" + "foo bar") + (("hello" + ("hello" + "hello")) + ""))]
        _ = v12
    }
    fmt.Println(v1)
    if ((v8 + v3) >= v8) {
        v5 = v5
    } else {
        var v13 uint32 = uint32(v3)
        var v14 int64 = -9223372036854775808
        _ = v13
        _ = v14
    }
    v15 := GenS1{F0: 1}
    v8 = (v3 >> ((2) & 63))
    var v16 float32 = (v5 - ((v5 + v5) * v5))
    v2 = 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
