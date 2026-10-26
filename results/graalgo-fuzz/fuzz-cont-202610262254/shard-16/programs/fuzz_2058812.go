package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16 }
func genFn0(p0 int32, p1 uint8) int {
    fmt.Println(p1, p1)
    var v5 int64 = 1
    var v6 bool = (false && false)
    v7 := GenS1{F0: -1, F1: -32768}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return 81
}

func main() {
    var v1 bool = true
    v2 := make(map[uint32]string)
    v3 := []string{}
    var v4 bool = ("foo bar" >= "a")
    v8 := GenS1{F0: 0, F1: 89}
    var v9 uint16 = 255
    v4 = false
    var v10 int64 = -686
    var v11 int32 = 0
    var v12 int64 = 9223372036854775807
    v1 = (!false)
    for fi13 := 0; fi13 < 3; fi13++ {
        v9 = (v9 * (v9 * 1))
        v2[2147483648] = ""
        var v14 uint = 42
        _ = fi13
        _ = v14
    }
    for fi15 := 0; fi15 < 17; fi15++ {
        for fi16 := 0; fi16 < 13; fi16++ {
            fmt.Println(v1)
            _ = fi16
        }
        v12 -= v12
        var v17 uint64 = 42
        _ = fi15
        _ = v17
    }
    var v18 int = 9223372036854775807
    var v19 int16 = -880
    if true {
        v1 = (42 <= v19)
    } else {
        v4 = true
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
