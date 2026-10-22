package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint }

func main() {
    v1 := []string{"foo bar", ("" + "\n\t"), ("hello" + "\n\t")}
    v2 := make(map[int]int)
    v3 := make(map[uint32]bool)
    var v4 uint32 = 1
    v5 := []uint32{2, v4}
    switch "foo bar" {
    case "hello":
        v4 -= (^(v4))
    case "foo bar":
        var v6 float32 = 1e10
        _ = v6
    case "":
        for fi7 := 0; fi7 < 5; fi7++ {
            var v8 uint32 = v4
            _ = len(v1)
            _ = fi7
            _ = v8
        }
    default:
        for fi9 := 0; fi9 < 8; fi9++ {
            if true {
                for fi10 := 0; fi10 < 11; fi10++ {
                    v11 := v2[4294967295]
                    _ = fi10
                    _ = v11
                }
                v12 := GenS1{F0: 42, F1: 2}
                v4 = v4
                _ = v12
            }
            fmt.Println(fi9, v4)
            _ = fi9
        }
    }
    v4 = (^(v4))
    for fi13 := 0; fi13 < 5; fi13++ {
        for fi14 := 0; fi14 < 20; fi14++ {
            var v15 uint32 = v4
            var v16 bool = (1e10 < 1e10)
            _ = fi14
            _ = v15
            _ = v16
        }
        var v17 float32 = 3.14159
        _ = fi13
        _ = v17
    }
    var v18 uint16 = 716
    var v19 uint = 2
    var v20 int32 = 109
    switch 161 {
    case 1:
        var v21 int = 653
        _ = v21
    case -1:
        v19 = 4294967295
    case 32768:
        v22 := GenS1{F0: 1, F1: (^(v19))}
        _ = v22
    default:
        var v23 int64 = -480
        _ = v23
    }
    v19 -= (v19 - 535)
    v20 = int32(v18)
    var v24 int64 = int64(v18)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
