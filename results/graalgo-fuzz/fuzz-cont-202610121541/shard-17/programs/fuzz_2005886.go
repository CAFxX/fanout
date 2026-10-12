package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int }

func main() {
    v1 := []uint32{4294967295, 0, 42, 65535}
    var v2 uint = 0
    var v3 int = 598
    v4 := make(map[string]bool)
    if (!(1 >= 0)) {
        v5 := GenS1{F0: 102, F1: (v3 - (-(v3)))}
        v3 -= v3
        var v6 string = "\n\t"
        _ = v5
        _ = v6
    }
    for fi7 := 0; fi7 < 7; fi7++ {
        v3 = fi7
        _ = fi7
    }
    switch v3 {
    case 128:
        for fi8 := 0; fi8 < 3; fi8++ {
            var v9 float32 = 0.0
            _ = fi8
            _ = v9
        }
    case 127:
        fmt.Println(v3, v3)
    default:
        v1 = append(v1, 0)
    }
    _ = len(v1)
    var v10 int64 = -1
    if (!false) {
        v11 := v4[(("" + "foo bar") + "hello")]
        _ = v11
    } else {
        if true {
            v3 = (v3 | 806)
            var v12 int64 = v10
            var v13 uint32 = 2
            _ = v12
            _ = v13
        }
        var v14 float32 = float32(1.5)
        _ = v14
    }
    var v15 float32 = 3.14159
    var v16 bool = true
    v10 = (v10 & -1)
    v1 = append(v1, 1)
    var v17 uint16 = 177
    var v18 int16 = -1
    switch (("" + "\n\t") + "\n\t") {
    case "":
        for fi19 := 0; fi19 < 9; fi19++ {
            for fi20 := 0; fi20 < 3; fi20++ {
                var v21 uint8 = 10
                if v16 {
                    var v22 int64 = 1
                    var v23 uint16 = v17
                    _ = v22
                    _ = v23
                }
                var v24 int32 = -2147483648
                var v25 uint8 = (v21 ^ (v21 + (v21 | 0)))
                _ = fi20
                _ = v21
                _ = v24
                _ = v25
            }
            _ = fi19
        }
    default:
        if true {
            if ((v2 << ((v2) & 63)) >= (v2 ^ (v2 % ((v2) | 1)))) {
                var v26 int64 = int64(v18)
                _ = v26
            }
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
