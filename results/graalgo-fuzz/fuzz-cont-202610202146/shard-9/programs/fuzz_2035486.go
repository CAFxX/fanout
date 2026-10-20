package main

import (
    "fmt"
)

func genFn0() uint8 {
    var v7 int16 = -32768
    var v8 uint32 = 762
    v9 := []int{}
    var v10 int32 = 899
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return 2
}

func main() {
    var v1 string = ("\n\t" + "a")
    var v2 int8 = 100
    var v3 float64 = 1.5
    v4 := make(map[uint32]string)
    var v5 uint64 = 0
    v6 := make(map[uint32]bool)
    var v11 uint8 = 42
    if false {
        var v12 int16 = -1
        v2 -= 1
        _ = v12
    } else {
        v1 = ("hello" + v1)
        var v13 int32 = 571
        v2 = v2
        _ = v13
    }
    v11 = (v11 >> ((4294967295) & 63))
    switch -980 {
    case -577:
        var v14 uint32 = 410
        _ = v14
    }
    fmt.Println(v5, v5, v11)
    var v15 int64 = -9223372036854775808
    var v16 uint = 65535
    v16 = v16
    var v17 string = v1
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
