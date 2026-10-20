package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32; F2 string }

func main() {
    var v1 string = ""
    v2 := []uint32{}
    v3 := make(map[uint32]bool)
    v4 := []int{127}
    fmt.Println(v1)
    fmt.Println(v1)
    if false {
        v5 := GenS1{F0: 0, F1: 0, F2: v1}
        var v6 int64 = -9223372036854775808
        if (true || (!((!true) || true))) {
            if false {
                for fi7 := 0; fi7 < 1; fi7++ {
                    var v8 int32 = 0
                    _ = fi7
                    _ = v8
                }
            }
            var v9 uint32 = 40
            _ = v9
        }
        _ = v5
        _ = v6
    }
    v10 := GenS1{F0: -32768, F1: 2147483648, F2: ((v1 + "foo bar") + "")}
    if true {
        v11 := v3[42]
        _ = v11
    }
    v12 := GenS1{F0: 0, F1: 941, F2: v1}
    var v13 uint8 = 127
    var v14 int8 = -83
    v14 = (v14 + -20)
    for fi15 := 0; fi15 < 19; fi15++ {
        var v16 uint64 = 4294967295
        _ = fi15
        _ = v16
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
}
