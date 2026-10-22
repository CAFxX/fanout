package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }

func main() {
    var v1 int8 = -71
    v2 := []int{}
    var v3 int64 = 1
    v4 := make(map[int]uint32)
    v3 -= -891
    for fi5 := 0; fi5 < 8; fi5++ {
        v3 += -9223372036854775808
        _ = fi5
    }
    var v6 float32 = ((1e10 - 0.1) - (((0.0 + -2.25) + (float32(0.0) / 0.1)) / float32(3.14159)))
    if true {
        fmt.Println(v6, v1)
        var v7 uint32 = 509
        _ = v7
    }
    switch 128 {
    case 42:
        var v8 uint64 = 1
        _ = v8
    case 642:
        v3 -= (-(v3))
    case 32768:
        var v9 int16 = 1
        _ = v9
    }
    for fi10 := 0; fi10 < 2; fi10++ {
        v4[569] = 4294967295
        for fi11 := 0; fi11 < 12; fi11++ {
            var v12 uint8 = 128
            _ = fi11
            _ = v12
        }
        _ = fi10
    }
    v6 = (1.5 / (-2.25 + 0.1))
    v13 := GenS1{F0: 42}
    switch -2147483648 {
    case 42:
        for fi14 := 0; fi14 < 1; fi14++ {
            for fi15 := 0; fi15 < 20; fi15++ {
                var v16 uint16 = 255
                _ = fi15
                _ = v16
            }
            _ = fi14
        }
    case 2147483647:
        if (!(true || (true && false))) {
            var v17 int = -433
            _ = v17
        } else {
            var v18 uint64 = 358
            fmt.Println(v18)
            _ = v18
        }
    default:
        fmt.Println(v3, v3, v3)
    }
    switch (("\n\t" + "") + "hello") {
    case "hello":
        v3 = v3
    case "":
        fmt.Println(v1, v3, v3)
    case "foo bar":
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    default:
        fmt.Println(v6, v6)
    }
    v3 = v3
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v6
}
