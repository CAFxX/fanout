package main

import (
    "fmt"
)


func main() {
    var v1 int64 = 921
    var v2 int = 42
    var v3 uint8 = 1
    var v4 float32 = -2.25
    v5 := make(map[uint32]string)
    var v6 uint32 = 65535
    var v7 uint32 = 65535
    v8 := v5[v7]
    if (65535 <= 32767) {
        v9 := []int{(v2 & (-(v2))), -1, 1}
        var v10 uint64 = 42
        _ = v10
        _ = v9
    } else {
        v11 := []int{v2, (-(v2)), v2}
        _ = v11
    }
    switch v8 {
    case "a":
        var v12 int16 = 32767
        _ = v12
    case "hello":
        var v13 int64 = v1
        _ = v13
    default:
        if (!(1 == 731)) {
            fmt.Println(v6, v2)
            v6 = v7
            var v14 int = (v2 * (v2 - -2147483648))
            _ = v14
        }
    }
    var v15 int64 = 42
    v15 -= -566
    var v16 float32 = (v4 * 0.0)
    for fi17 := 0; fi17 < 16; fi17++ {
        fmt.Println(v1, fi17)
        _ = fi17
    }
    var v18 int8 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v15
    _ = v16
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
