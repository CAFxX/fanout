package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint32 }

func main() {
    v1 := make(map[int]bool)
    var v2 uint = 1
    var v3 int32 = 547
    v4 := v1[496]
    var v5 int64 = 1
    var v6 uint = (v2 % ((4294967295) | 1))
    for fi7 := 0; fi7 < 7; fi7++ {
        v8 := GenS1{F0: 499}
        _ = fi7
        _ = v8
    }
    delete(v1, 1)
    var v9 float32 = (1e10 / 1e10)
    var v10 int16 = 308
    switch -9223372036854775808 {
    case -2147483648:
        var v11 uint64 = 2
        _ = v11
    case 131:
        if v4 {
            switch ((((("" + "\n\t") + "") + "hello") + "\n\t") + (("a" + ("\n\t" + ("\n\t" + "\n\t"))) + "hello")) {
            case "foo bar":
                var v12 int8 = -128
                _ = v12
            case "a":
                var v13 uint8 = 255
                _ = v13
            default:
                v14 := GenS1{F0: 65535}
                _ = v14
            }
            fmt.Println(v6)
        }
    default:
        fmt.Println(v5, v10)
    }
    for fi15 := 0; fi15 < 9; fi15++ {
        var v16 int32 = v3
        v17 := GenS2{F0: 2147483647}
        _ = fi15
        _ = v16
        _ = v17
    }
    if false {
        v10 = -100
        v18 := GenS2{F0: 1}
        var v19 uint = 724
        _ = v18
        _ = v19
    }
    fmt.Println(v5, v5, v9)
    v20 := []int{168}
    var v21 int8 = 100
    var v22 uint32 = 610
    var v23 uint64 = 339
    if len(v20) > 0 {
        v20[0] = 2147483647
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
