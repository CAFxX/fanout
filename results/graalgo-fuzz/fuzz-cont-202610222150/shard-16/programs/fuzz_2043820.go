package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
func genFn0(p0 bool) int8 {
    fmt.Println(p0)
    p0 = true
    v5 := GenS1{F0: 1}
    _ = p0
    _ = v5
    return 100
}

func main() {
    var v1 uint = 77
    var v2 uint32 = 4294967295
    var v3 int = 32767
    var v4 uint = v1
    v6 := []int{32768}
    v4 -= v4
    for fi7 := 0; fi7 < 15; fi7++ {
        switch "\n\t" {
        case "a":
            var v8 uint16 = 49
            _ = v8
        case "foo bar":
            v3 = v3
        default:
            fmt.Println(fi7, v2)
        }
        var v9 int16 = -586
        _ = fi7
        _ = v9
    }
    v10 := GenS1{F0: 607}
    _ = len(v6)
    v11 := GenS1{F0: -75}
    v3 -= v3
    var v12 uint = 990
    var v13 int16 = -301
    var v14 bool = false
    var v15 uint32 = 1
    fmt.Println(v13, v15)
    for fi16 := 0; fi16 < 2; fi16++ {
        v17 := GenS1{F0: 1}
        _ = fi16
        _ = v17
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v6
}
