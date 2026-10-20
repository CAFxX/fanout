package main

import (
    "fmt"
)

func genFn0(p0 int32, p1 uint16, p2 uint16) int64 {
    var v5 uint32 = 0
    if false {
        var v6 string = "a"
        _ = v6
    }
    var v7 uint = 835
    var v8 int32 = p0
    var v9 int8 = -59
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v7
    _ = v8
    _ = v9
    return -858
}

func main() {
    var v1 int = 73
    var v2 uint8 = 127
    var v3 bool = (true || false)
    v4 := []int{(v1 << ((65535) & 63)), (^(v1)), 4294967295, (v1 & v1)}
    fmt.Println(v1, v2, v3)
    var v10 int64 = 0
    switch v1 {
    case 9223372036854775807:
        v10 -= (v10 % ((-9223372036854775808) | 1))
    case -422:
        v4 = append(v4, v1)
    }
    var v11 string = (("" + "foo bar") + "a")
    _ = len(v4)
    v3 = (0.1 != 1e10)
    v11 = v11
    var v12 uint8 = v2
    var v13 int16 = 367
    var v14 uint64 = 1
    var v15 int8 = -48
    _ = len(v4)
    var v16 uint = 255
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
