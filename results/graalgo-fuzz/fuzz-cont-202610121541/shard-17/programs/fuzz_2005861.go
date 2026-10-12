package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int }
func genFn0(p0 int) int16 {
    v5 := []int{-1, p0, (p0 / (((p0 | -1)) | 1)), 263}
    fmt.Println(p0)
    var v6 uint8 = 127
    _ = p0
    _ = v5
    _ = v6
    return -359
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[uint32]int)
    v3 := make(map[int]uint32)
    v4 := []string{"\n\t", "foo bar", ("\n\t" + "foo bar"), "\n\t"}
    switch 42 {
    case 370:
        var v7 int64 = 0
        _ = v7
    default:
        var v8 uint32 = 161
        _ = v8
    }
    var v9 uint32 = 565
    for fi10 := 0; fi10 < 18; fi10++ {
        var v11 int32 = 42
        v11 -= v11
        if false {
            var v12 uint8 = 116
            _ = len(v4)
            _ = v12
        }
        _ = fi10
        _ = v11
    }
    v13 := GenS1{F0: "\n\t", F1: 32767}
    v1[(("" + "a") + "a")] = 79
    for fi14 := 0; fi14 < 13; fi14++ {
        v9 -= 4294967295
        v9 = 4294967295
        _ = fi14
    }
    for fi15 := 0; fi15 < 8; fi15++ {
        var v16 uint64 = 653
        _ = fi15
        _ = v16
    }
    for fi17 := 0; fi17 < 19; fi17++ {
        var v18 int32 = 0
        _ = fi17
        _ = v18
    }
    var v19 int8 = -1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
