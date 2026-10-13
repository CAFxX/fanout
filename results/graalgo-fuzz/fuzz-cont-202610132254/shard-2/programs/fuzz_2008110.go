package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16; F2 int8 }
func genFn0() uint32 {
    var v5 uint32 = 2147483647
    var v6 uint8 = 42
    for fi7 := 0; fi7 < 18; fi7++ {
        fmt.Println(v6, v5)
        _ = fi7
    }
    var v8 float32 = -2.25
    for fi9 := 0; fi9 < 8; fi9++ {
        fmt.Println(v8, v8)
        for fi10 := 0; fi10 < 6; fi10++ {
            v11 := []int{(^(fi9)), fi9, (fi9 / ((fi10) | 1)), (fi9 + fi10)}
            fmt.Println(v5, v5, fi9)
            switch (("foo bar" + "hello") + ("" + "hello")) {
            case "a":
                fmt.Println(fi10)
            case "hello":
                v11 = append(v11, (fi9 - 127))
            }
            _ = fi10
            _ = v11
        }
        _ = fi9
    }
    _ = v5
    _ = v6
    _ = v8
    return (v5 * 2)
}

func main() {
    v1 := make(map[int]int)
    var v2 int16 = -888
    var v3 float32 = ((0.1 + 1.5) - -2.25)
    v4 := []uint32{0, 1}
    v3 = v3
    for fi12 := 0; fi12 < 3; fi12++ {
        for fi13 := 0; fi13 < 5; fi13++ {
            fmt.Println(v2)
            _ = fi13
        }
        _ = fi12
    }
    v14 := GenS1{F0: v2, F1: v2, F2: 127}
    var v15 int64 = -957
    v2 += (v2 / ((1000) | 1))
    delete(v1, -2147483648)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    for fi16 := 0; fi16 < 12; fi16++ {
        v2 = (v2 * v2)
        _ = fi16
    }
    if false {
        fmt.Println(v2, v15, v3)
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    } else {
        v3 = v3
        var v17 int8 = -16
        _ = v17
    }
    switch "foo bar" {
    case "hello":
        fmt.Println(v15, v3, v3)
    case "":
        var v18 int8 = 127
        _ = v18
    case "\n\t":
        v3 = v3
    }
    v2 = v2
    var v19 uint64 = 42
    var v20 int8 = 1
    var v21 int64 = (v15 % ((-9223372036854775808) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
}
