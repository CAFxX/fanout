package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 bool; F1 int64 }
func genFn0() int {
    var v5 int = -40
    v5 = (v5 / ((1) | 1))
    _ = v5
    return v5
}

func main() {
    v1 := []int{-1}
    var v2 uint = 417
    var v3 float32 = -2.25
    v4 := make(map[uint32]uint32)
    var v6 int8 = 46
    for fi7 := 0; fi7 < 12; fi7++ {
        delete(v4, 862)
        _ = fi7
    }
    _ = len(v1)
    v8 := GenS1{F0: uint32(v2)}
    for fi9 := 0; fi9 < 2; fi9++ {
        fmt.Println(v6, v3, v6)
        fmt.Println(v3, v2)
        var v10 uint = v2
        _ = fi9
        _ = v10
    }
    fmt.Println(v2)
    v11 := GenS1{F0: 1}
    var v12 bool = false
    v6 += 42
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v13 := GenS1{F0: 723}
    var v14 int32 = 1
    if v12 {
        v12 = ("\n\t" > ("foo bar" + (("" + "foo bar") + ("\n\t" + "a"))))
    }
    for fi15 := 0; fi15 < 15; fi15++ {
        v16 := GenS1{F0: 65535}
        for fi17 := 0; fi17 < 7; fi17++ {
            fmt.Println(fi15)
            v18 := GenS2{F0: v12, F1: -9223372036854775808}
            _ = fi17
            _ = v18
        }
        _ = fi15
        _ = v16
    }
    var v19 float32 = v3
    v2 += v2
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v8
}
