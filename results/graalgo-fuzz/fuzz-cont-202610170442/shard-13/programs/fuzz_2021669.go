package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int; F2 int8 }
type GenS2 struct { F0 string }
func genFn0(p0 int16) uint16 {
    fmt.Println(p0)
    var v4 int = -1
    _ = p0
    _ = v4
    return 65535
}

func main() {
    v1 := make(map[uint32]int)
    var v2 bool = false
    v3 := make(map[string]int)
    var v5 int8 = 100
    v5 = v5
    v6 := GenS1{F0: 1000, F1: 9223372036854775807, F2: (-(v5))}
    v5 += 69
    var v7 int32 = 512
    v8 := GenS1{F0: 42, F1: -140, F2: 1}
    _ = len(v1)
    for fi9 := 0; fi9 < 14; fi9++ {
        if v2 {
            v5 += v5
            var v10 string = "\n\t"
            var v11 int64 = 9223372036854775807
            _ = v10
            _ = v11
        } else {
            for fi12 := 0; fi12 < 17; fi12++ {
                fmt.Println(v5)
                v13 := GenS1{F0: 1, F1: -910, F2: v5}
                _ = fi12
                _ = v13
            }
            switch "hello" {
            case "foo bar":
                v14 := GenS2{F0: "\n\t"}
                _ = v14
            default:
                v2 = false
            }
        }
        _ = fi9
    }
    v15 := []int{-898, -248}
    for fi16 := 0; fi16 < 15; fi16++ {
        v7 = 1
        var v17 int8 = (v5 | (v5 >> ((1) & 63)))
        v18 := GenS2{F0: "\n\t"}
        _ = fi16
        _ = v17
        _ = v18
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
