package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8 }

func main() {
    var v1 string = ""
    var v2 string = (("a" + ("" + "hello")) + v1)
    v3 := []uint32{}
    v4 := make(map[uint32]int)
    var v5 uint8 = 128
    v3 = append(v3, 65535)
    for fi6 := 0; fi6 < 11; fi6++ {
        var v7 int16 = -32768
        _ = fi6
        _ = v7
    }
    v1 = (v1 + ("\n\t" + v2))
    if (1 > 1) {
        var v8 int32 = 0
        _ = v8
    }
    var v9 bool = (1 > 2147483648)
    switch "hello" {
    case "":
        for fi10 := 0; fi10 < 14; fi10++ {
            for fi11 := 0; fi11 < 20; fi11++ {
                var v12 uint16 = 32767
                var v13 string = v1
                _ = fi11
                _ = v12
                _ = v13
            }
            if len(v3) > 0 {
                v3[0] = 1
            }
            _ = fi10
        }
    case "foo bar":
        var v14 string = v2
        _ = v14
    default:
        var v15 bool = (0 < 234)
        _ = v15
    }
    for fi16 := 0; fi16 < 7; fi16++ {
        _ = len(v4)
        if v9 {
            var v17 uint64 = 4294967295
            var v18 uint16 = 32767
            _ = v17
            _ = v18
        }
        v19 := GenS1{F0: 1, F1: v5}
        _ = fi16
        _ = v19
    }
    for fi20 := 0; fi20 < 20; fi20++ {
        var v21 int16 = 964
        var v22 uint32 = 199
        v9 = v9
        _ = fi20
        _ = v21
        _ = v22
    }
    if ((2147483647 > -740) && v9) {
        var v23 float32 = -2.25
        v23 += 1e10
        _ = v23
    }
    fmt.Println(v5)
    v9 = ((127 <= 100) || v9)
    for fi24 := 0; fi24 < 9; fi24++ {
        _ = len(v3)
        _ = fi24
    }
    var v25 bool = v9
    v26 := GenS1{F0: 255, F1: v5}
    if ((216 != (v5 % ((v5) | 1))) && (v25 || (!true))) {
        if len(v3) > 0 {
            v3[0] = 1
        }
        v1 = v1
    } else {
        if len(v3) > 0 {
            v3[0] = 1
        }
        var v27 int64 = 42
        var v28 int8 = 42
        _ = v27
        _ = v28
    }
    v2 = "\n\t"
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
