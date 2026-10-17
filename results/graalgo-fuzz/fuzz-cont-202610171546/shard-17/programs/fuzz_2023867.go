package main

import (
    "fmt"
)

type GenS1 struct { F0 string }

func main() {
    var v1 uint16 = 206
    var v2 uint8 = 1
    var v3 float32 = (0.1 - 0.0)
    v4 := []uint32{1, 65535, 782, 0}
    var v5 bool = (-9223372036854775808 == 1)
    v6 := make(map[uint32]uint32)
    var v7 int = -9223372036854775808
    v3 -= 1e10
    for fi8 := 0; fi8 < 17; fi8++ {
        v9 := GenS1{F0: ("hello" + "\n\t")}
        _ = fi8
        _ = v9
    }
    var v10 int8 = -128
    if (!(("hello" != "hello") && v5)) {
        var v11 bool = v5
        var v12 int16 = -32768
        var v13 int = v7
        _ = v11
        _ = v12
        _ = v13
    } else {
        if (-1 > -1) {
            var v14 uint32 = 670
            _ = v14
        } else {
            v15 := v6[65535]
            var v16 int32 = -1
            _ = v15
            _ = v16
        }
        fmt.Println(v1, v7, v5)
    }
    v17 := GenS1{F0: (("" + "a") + (("foo bar" + "") + "hello"))}
    if false {
        var v18 float32 = 1e10
        _ = v18
    } else {
        v19 := GenS1{F0: "foo bar"}
        v3 = v3
        _ = v19
    }
    var v20 int16 = -32768
    v10 -= 0
    _ = len(v6)
    var v21 int64 = -781
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v17
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
