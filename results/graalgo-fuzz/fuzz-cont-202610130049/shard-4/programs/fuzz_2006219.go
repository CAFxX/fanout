package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
func genFn0(p0 int8, p1 string) int32 {
    var v4 bool = ((793 == 127) && false)
    fmt.Println(v4)
    _ = p0
    _ = p1
    _ = v4
    return 42
}

func main() {
    v1 := []int{32767, 9223372036854775807, 35, 32767}
    var v2 uint = 4294967295
    var v3 float64 = 1e10
    if (!(1 < 255)) {
        _ = len(v1)
    }
    v5 := GenS1{F0: 193}
    fmt.Println(v2, v2)
    for fi6 := 0; fi6 < 19; fi6++ {
        v2 = 1
        _ = fi6
    }
    v3 = 0.0
    if true {
        v7 := GenS1{F0: 65535}
        _ = v7
    } else {
        v1 = append(v1, -695)
        var v8 int32 = 42
        _ = v8
    }
    var v9 int32 = 42
    if true {
        fmt.Println(v2)
    }
    for fi10 := 0; fi10 < 7; fi10++ {
        var v11 bool = (32767 != 32767)
        _ = fi10
        _ = v11
    }
    if (true || (127 == 42)) {
        fmt.Println(v9, v2)
        v9 += v9
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v5
    _ = v9
}
