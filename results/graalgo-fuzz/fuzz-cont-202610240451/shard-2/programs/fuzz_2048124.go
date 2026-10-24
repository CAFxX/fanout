package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64 }
type GenS2 struct { F0 string; F1 int32 }

func main() {
    var v1 uint = 4294967295
    var v2 uint32 = 2147483647
    v3 := make(map[uint32]uint32)
    var v4 uint32 = (v2 % (((^(v2))) | 1))
    if true {
        v5 := GenS1{F0: 1, F1: 580}
        v6 := GenS2{F0: "a", F1: -1}
        if ("foo bar" >= "\n\t") {
            v2 -= (v2 / ((v2) | 1))
        }
        _ = v5
        _ = v6
    } else {
        var v7 int8 = -128
        _ = v7
    }
    var v8 uint64 = 878
    if ((!true) && false) {
        fmt.Println(v4, v2)
    } else {
        for fi9 := 0; fi9 < 17; fi9++ {
            for fi10 := 0; fi10 < 9; fi10++ {
                if false {
                    v2 = (v2 / ((v4) | 1))
                }
                if (2147483647 <= -2147483648) {
                    var v11 int16 = 42
                    _ = v11
                }
                var v12 uint = v1
                _ = fi10
                _ = v12
            }
            v8 = 0
            _ = fi9
        }
        v2 += v4
        for fi13 := 0; fi13 < 19; fi13++ {
            var v14 int64 = -594
            _ = fi13
            _ = v14
        }
    }
    v15 := []int{4294967295}
    v3[v4] = uint32(v8)
    var v16 uint = v1
    var v17 int8 = 127
    for fi18 := 0; fi18 < 12; fi18++ {
        var v19 uint32 = 2147483648
        switch (("hello" + (("" + ("" + "\n\t")) + ("a" + ("" + "hello")))) + "\n\t") {
        case "foo bar":
            fmt.Println(v4)
        default:
            if ((v16 == 175) || false) {
                fmt.Println(v8)
                fmt.Println(v16)
            }
        }
        _ = fi18
        _ = v19
    }
    fmt.Println(v4, v17, v17)
    switch "foo bar" {
    case "":
        var v20 int16 = 1000
        _ = v20
    }
    fmt.Println(v4, v4, v17)
    var v21 int8 = v17
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v8
}
