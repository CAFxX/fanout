package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool }
type GenS2 struct { F0 uint32; F1 int64 }

func main() {
    var v1 uint32 = 772
    var v2 int = 128
    var v3 int16 = -32768
    v4 := []int{v2, v2}
    v1 += v1
    v3 = 171
    for fi5 := 0; fi5 < 19; fi5++ {
        v4 = append(v4, (^(v2)))
        _ = fi5
    }
    v4 = append(v4, v2)
    var v6 uint32 = 2147483647
    v7 := GenS1{F0: 336, F1: true}
    switch 127 {
    case 86:
        if false {
            v8 := GenS2{F0: (^(v1)), F1: -885}
            for fi9 := 0; fi9 < 11; fi9++ {
                var v10 int16 = 32767
                fmt.Println(v1, v6)
                _ = fi9
                _ = v10
            }
            _ = v8
        } else {
            v11 := GenS1{F0: 4294967295, F1: false}
            var v12 int8 = -1
            if len(v4) > 0 {
                v4[0] = v2
            }
            _ = v11
            _ = v12
        }
    }
    if (9223372036854775807 <= 812) {
        v2 += -1
        var v13 uint64 = 2
        _ = v13
    }
    var v14 int8 = 100
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
}
