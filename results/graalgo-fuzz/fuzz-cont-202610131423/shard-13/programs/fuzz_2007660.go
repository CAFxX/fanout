package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint8; F1 int32; F2 int64 }
func genFn0() string {
    v7 := []int{32767, 42, 32768, 42}
    switch "a" {
    case "hello":
        var v8 string = ("" + "\n\t")
        _ = v8
    case "a":
        for fi9 := 0; fi9 < 12; fi9++ {
            var v10 uint64 = 42
            _ = fi9
            _ = v10
        }
    case "foo bar":
        var v11 uint64 = 744
        _ = v11
    }
    if false {
        if (!false) {
            if len(v7) > 0 {
                v7[0] = -1
            }
            var v12 uint = 2
            _ = v12
        }
        if false {
            v13 := GenS1{F0: ("a" + "hello")}
            _ = v13
        }
        _ = len(v7)
    }
    if (((0.0 * -2.25) + float32(0.0)) != 0.0) {
        var v14 int = -344
        v14 = (-(v14))
        for fi15 := 0; fi15 < 16; fi15++ {
            v7 = append(v7, (v14 << ((65535) & 63)))
            v16 := GenS2{F0: 2, F1: -833, F2: 42}
            _ = fi15
            _ = v16
        }
        _ = v14
    }
    switch 1 {
    case -868:
        var v17 uint8 = 248
        _ = v17
    }
    _ = v7
    return "\n\t"
}

func main() {
    var v1 int = 42
    var v2 uint64 = 4294967295
    v3 := make(map[int]string)
    var v4 int64 = 950
    v5 := make(map[uint32]int)
    var v6 int16 = 1
    var v18 int32 = -938
    for fi19 := 0; fi19 < 9; fi19++ {
        var v20 int8 = 42
        _ = fi19
        _ = v20
    }
    var v21 int16 = (v6 & 42)
    var v22 uint32 = 2147483647
    fmt.Println(v4, v21, v2)
    var v23 uint = 4294967295
    fmt.Println(v22, v18)
    fmt.Println(v6)
    var v24 float64 = 1.5
    var v25 uint16 = 749
    var v26 uint16 = v25
    v4 += -1
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
