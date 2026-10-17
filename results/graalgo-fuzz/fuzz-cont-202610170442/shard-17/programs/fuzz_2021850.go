package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16 }
type GenS2 struct { F0 uint16 }
func genFn0(p0 uint) int16 {
    switch "hello" {
    case "\n\t":
        if false {
            var v6 float32 = 3.14159
            v7 := GenS1{F0: 1, F1: 420}
            _ = v6
            _ = v7
        } else {
            if (!true) {
                var v8 int64 = -565
                _ = v8
            } else {
                v9 := []int{2147483647, 169, -1, -186}
                _ = v9
            }
            v10 := []int{255, 2147483647, 218}
            _ = v10
        }
    }
    v11 := []int{}
    _ = p0
    _ = v11
    return 32767
}

func main() {
    var v1 int32 = 42
    var v2 uint = 4294967295
    var v3 uint8 = 128
    var v4 int = 128
    var v5 uint32 = 2147483647
    var v12 uint64 = 161
    var v13 int8 = -1
    if false {
        v14 := []int{v4}
        for fi15 := 0; fi15 < 7; fi15++ {
            fmt.Println(v13)
            _ = fi15
        }
        fmt.Println(v13)
        _ = v14
    }
    v16 := []int{v4}
    for fi17 := 0; fi17 < 13; fi17++ {
        v18 := GenS1{F0: 32767, F1: -284}
        var v19 uint32 = v5
        _ = fi17
        _ = v18
        _ = v19
    }
    v20 := GenS1{F0: -1, F1: 32767}
    if len(v16) > 0 {
        v16[0] = (v4 - 42)
    }
    v21 := GenS2{F0: 322}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v16
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
