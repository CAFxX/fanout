package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool; F2 int64 }
type GenS2 struct { F0 uint; F1 string }

func main() {
    v1 := make(map[uint32]string)
    var v2 int16 = -1
    var v3 int32 = 677
    var v4 int8 = 42
    for fi5 := 0; fi5 < 15; fi5++ {
        v3 += -475
        _ = fi5
    }
    v6 := GenS1{F0: (-(v3)), F1: false, F2: -577}
    v4 -= (-(v4))
    v7 := []int{}
    v7 = append(v7, 2147483647)
    delete(v1, 2)
    switch "" {
    case "foo bar":
        v4 += v4
    case "hello":
        v8 := GenS2{F0: 342, F1: ("\n\t" + "")}
        _ = v8
    case "a":
        var v9 uint32 = 4294967295
        _ = v9
    }
    for fi10 := 0; fi10 < 10; fi10++ {
        var v11 int = (fi10 / ((fi10) | 1))
        if (42 > 842) {
            var v12 uint32 = 540
            v11 = (-(fi10))
            switch "a" {
            case "\n\t":
                var v13 bool = (v12 <= (v12 * v12))
                _ = v13
            case "":
                var v14 uint16 = 65535
                _ = v14
            case "hello":
                v4 = v4
            default:
                for fi15 := 0; fi15 < 7; fi15++ {
                    var v16 float32 = 1.5
                    _ = fi15
                    _ = v16
                }
            }
            _ = v12
        }
        v4 += v4
        _ = fi10
        _ = v11
    }
    var v17 uint8 = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
}
