package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64; F2 int8 }

func main() {
    var v1 int64 = -582
    var v2 uint32 = 0
    v3 := []uint32{(v2 * (v2 - 229)), v2}
    var v4 string = "hello"
    var v5 int16 = -1
    _ = len(v3)
    switch 960 {
    case 4294967295:
        v3 = append(v3, v2)
    case 127:
        v5 += (v5 ^ v5)
    }
    fmt.Println(v1)
    var v6 uint64 = 4294967295
    switch "foo bar" {
    case "hello":
        v5 = 42
    case "foo bar":
        v7 := GenS1{F0: 255, F1: (v1 - int64(v2)), F2: 0}
        _ = v7
    case "\n\t":
        var v8 int8 = 127
        _ = v8
    }
    for fi9 := 0; fi9 < 1; fi9++ {
        if true {
            v2 += (^(v2))
            var v10 uint32 = v2
            _ = v10
        } else {
            v11 := GenS1{F0: 1, F1: 0, F2: -5}
            var v12 uint16 = 288
            for fi13 := 0; fi13 < 13; fi13++ {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                if len(v3) > 0 {
                    v3[0] = v2
                }
                _ = fi13
            }
            _ = v11
            _ = v12
        }
        _ = len(v3)
        v14 := GenS1{F0: 4294967295, F1: (^(v1)), F2: int8(v6)}
        _ = fi9
        _ = v14
    }
    var v15 int64 = v1
    var v16 int32 = -2147483648
    fmt.Println(v16, v1, v2)
    fmt.Println(v4)
    var v17 float32 = ((0.0 + ((0.0 + float32(((1e10 / float64(1e10)) + -2.25))) - ((-2.25 + -2.25) - (0.1 + 1.5)))) / 0.1)
    _ = len(v3)
    var v18 bool = false
    fmt.Println(v16, v18, v1)
    switch v4 {
    case "foo bar":
        v16 -= v16
    default:
        for fi19 := 0; fi19 < 19; fi19++ {
            v4 = "hello"
            var v20 int8 = 0
            _ = fi19
            _ = v20
        }
    }
    v21 := GenS1{F0: 652, F1: v1, F2: 100}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
