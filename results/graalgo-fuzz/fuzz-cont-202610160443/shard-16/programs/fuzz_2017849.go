package main

import (
    "fmt"
)

type GenS1 struct { F0 string }

func main() {
    var v1 float32 = 0.0
    v2 := make(map[uint32]string)
    v3 := []uint32{0}
    var v4 int8 = -1
    var v5 int16 = 1000
    var v6 uint64 = 4294967295
    for fi7 := 0; fi7 < 5; fi7++ {
        var v8 bool = (!false)
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi7
        _ = v8
    }
    switch 888 {
    case -494:
        var v9 uint = 650
        _ = v9
    default:
        var v10 uint8 = 42
        _ = v10
    }
    if (-1 > -2147483648) {
        for fi11 := 0; fi11 < 12; fi11++ {
            var v12 int = 2147483647
            _ = fi11
            _ = v12
        }
        fmt.Println(v1)
        for fi13 := 0; fi13 < 4; fi13++ {
            if (!(2147483648 > 65535)) {
                var v14 uint16 = 65535
                _ = v14
            } else {
                v15 := GenS1{F0: ("" + ("a" + "\n\t"))}
                fmt.Println(v6, fi13)
                _ = v15
            }
            var v16 int64 = 42
            var v17 uint8 = 100
            _ = fi13
            _ = v16
            _ = v17
        }
    }
    v1 = v1
    var v18 int64 = -121
    v19 := v2[2147483648]
    v20 := v2[639]
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
