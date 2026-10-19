package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
func genFn0(p0 uint32, p1 uint, p2 int64) int {
    fmt.Println(p1)
    if true {
        var v7 int64 = p2
        v8 := GenS1{F0: 1}
        for fi9 := 0; fi9 < 7; fi9++ {
            var v10 int16 = -778
            _ = fi9
            _ = v10
        }
        _ = v7
        _ = v8
    } else {
        var v11 int = 11
        _ = v11
    }
    p1 = (^(p1))
    v12 := []int{32767, int(p1), 42}
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return 128
}

func main() {
    var v1 int = 312
    var v2 bool = (!false)
    v3 := make(map[string]string)
    var v4 int = v1
    v5 := []int{}
    var v6 int64 = -931
    var v13 int32 = 42
    var v14 float64 = (-2.25 + 1e10)
    for fi15 := 0; fi15 < 8; fi15++ {
        var v16 string = "hello"
        _ = fi15
        _ = v16
    }
    if len(v5) > 0 {
        v5[0] = (v1 + v4)
    }
    var v17 bool = v2
    for fi18 := 0; fi18 < 13; fi18++ {
        v6 = (v6 << ((18446744073709551615) & 63))
        _ = fi18
    }
    v17 = v2
    fmt.Println(v13)
    var v19 uint = 0
    v20 := GenS1{F0: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v17
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
