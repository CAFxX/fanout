package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int8; F1 uint64 }
type GenS3 struct { F0 uint32 }

func main() {
    var v1 uint64 = 2
    var v2 float64 = (3.14159 - (0.1 * 1e10))
    v3 := make(map[uint32]bool)
    v4 := []uint32{}
    var v5 uint = 395
    var v6 int32 = -111
    var v7 uint16 = 32767
    v8 := GenS1{F0: "foo bar"}
    v9 := GenS1{F0: "\n\t"}
    v10 := GenS1{F0: ("a" + ("hello" + ""))}
    for fi11 := 0; fi11 < 9; fi11++ {
        for fi12 := 0; fi12 < 5; fi12++ {
            var v13 int64 = 1
            var v14 int32 = -1
            if false {
                v15 := GenS2{F0: 0, F1: v1}
                _ = v15
            }
            v16 := GenS1{F0: "\n\t"}
            _ = fi12
            _ = v13
            _ = v14
            _ = v16
        }
        var v17 uint64 = v1
        switch ("foo bar" + ("\n\t" + ("hello" + "\n\t"))) {
        case "hello":
            var v18 uint32 = 42
            _ = v18
        default:
            _ = len(v3)
        }
        _ = fi11
        _ = v17
    }
    v19 := GenS3{F0: uint32(v7)}
    var v20 int64 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
