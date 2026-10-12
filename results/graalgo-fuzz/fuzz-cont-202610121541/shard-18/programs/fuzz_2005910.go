package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8 }
type GenS2 struct { F0 int16; F1 int8; F2 uint32 }

func main() {
    var v1 uint32 = 24
    var v2 int32 = 2147483647
    var v3 string = ""
    v4 := make(map[string]int)
    for fi5 := 0; fi5 < 7; fi5++ {
        v3 = v3
        v2 += -2147483648
        _ = fi5
    }
    v6 := []int{-347, -1, 2147483647, 127}
    for fi7 := 0; fi7 < 3; fi7++ {
        fmt.Println(v3, v2, v3)
        fmt.Println(v1, v2)
        var v8 uint64 = 2
        _ = fi7
        _ = v8
    }
    var v9 int8 = -25
    v10 := GenS1{F0: false, F1: 63}
    switch ("" + (v3 + "\n\t")) {
    case "foo bar":
        for fi11 := 0; fi11 < 20; fi11++ {
            v12 := GenS2{F0: 32767, F1: v9, F2: v1}
            v1 = v1
            _ = fi11
            _ = v12
        }
    case "\n\t":
        for fi13 := 0; fi13 < 3; fi13++ {
            v1 -= v1
            fmt.Println(v9, v9)
            v3 = "\n\t"
            _ = fi13
        }
    default:
        v6 = append(v6, 408)
    }
    var v14 uint8 = 255
    switch v3 {
    case "foo bar":
        v14 = (v14 + v14)
    case "\n\t":
        v6 = append(v6, 42)
    case "hello":
        for fi15 := 0; fi15 < 12; fi15++ {
            v9 += v9
            _ = fi15
        }
    default:
        var v16 float32 = -2.25
        _ = v16
    }
    v2 = v2
    var v17 uint = 65535
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v14
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v9
}
