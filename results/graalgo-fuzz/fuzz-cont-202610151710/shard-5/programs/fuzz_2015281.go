package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8 }

func main() {
    v1 := make(map[string]int)
    var v2 uint64 = 582
    var v3 string = ("foo bar" + "foo bar")
    var v4 int32 = 2147483647
    v5 := []uint32{707, 2, 65535}
    var v6 int32 = (v4 & v4)
    if false {
        switch (v3 + v3) {
        case "":
            var v7 int64 = -1
            _ = v7
        }
        if false {
            fmt.Println(v2, v4)
            for fi8 := 0; fi8 < 20; fi8++ {
                var v9 uint16 = 0
                var v10 uint64 = v2
                var v11 int16 = -11
                _ = fi8
                _ = v10
                _ = v11
                _ = v9
            }
        } else {
            var v12 uint16 = 1
            var v13 uint64 = v2
            _ = v12
            _ = v13
        }
    } else {
        fmt.Println(v6, v6, v6)
        fmt.Println(v3, v3)
    }
    v1["a"] = 0
    for fi14 := 0; fi14 < 9; fi14++ {
        var v15 uint64 = (v2 / ((1) | 1))
        fmt.Println(v15, v2, v15)
        var v16 float32 = 1e10
        _ = fi14
        _ = v15
        _ = v16
    }
    v2 += (v2 << ((1) & 63))
    for fi17 := 0; fi17 < 2; fi17++ {
        for fi18 := 0; fi18 < 3; fi18++ {
            v6 = v6
            _ = fi18
        }
        var v19 int32 = v4
        v6 = (v19 + 361)
        _ = fi17
        _ = v19
    }
    v20 := GenS1{F0: 255, F1: 42}
    fmt.Println(v4, v6)
    v6 += (^(v4))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
