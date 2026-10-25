package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }

func main() {
    var v1 int = -757
    var v2 float64 = 0.0
    var v3 uint = 89
    var v4 int64 = 0
    var v5 int32 = 2147483647
    v1 = v1
    for fi6 := 0; fi6 < 5; fi6++ {
        v7 := []int{fi6, 2147483647, (v1 << ((0) & 63)), fi6}
        var v8 bool = (-101 >= v1)
        _ = fi6
        _ = v7
        _ = v8
    }
    var v9 int8 = 1
    v2 = v2
    if (v1 > v1) {
        v10 := []int{(v1 - v1), (v1 + v1)}
        v11 := GenS1{F0: 861}
        v9 += (v9 >> (((^(v3))) & 63))
        _ = v10
        _ = v11
    }
    v12 := []int{}
    if len(v12) > 0 {
        v12[0] = -518
    }
    v2 = float64(1.5)
    var v13 string = ""
    for fi14 := 0; fi14 < 10; fi14++ {
        var v15 int = fi14
        _ = fi14
        _ = v15
    }
    for fi16 := 0; fi16 < 3; fi16++ {
        if true {
            v17 := GenS1{F0: 0}
            var v18 float64 = (float64(1e10) - v2)
            _ = v17
            _ = v18
        }
        for fi19 := 0; fi19 < 9; fi19++ {
            _ = len(v12)
            _ = fi19
        }
        _ = fi16
    }
    fmt.Println(v5)
    v20 := GenS1{F0: 95}
    _ = len(v12)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
