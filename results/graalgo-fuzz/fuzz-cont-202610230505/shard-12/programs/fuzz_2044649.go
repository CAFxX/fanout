package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
func genFn0(p0 string, p1 uint64) uint64 {
    for fi6 := 0; fi6 < 12; fi6++ {
        v7 := []int{-419, 4294967295}
        var v8 uint32 = 1
        fmt.Println(p0)
        _ = fi6
        _ = v7
        _ = v8
    }
    if true {
        v9 := []int{-2147483648, 845, -616}
        v9 = append(v9, -113)
        if false {
            p1 = 0
            fmt.Println(p0, p0)
        }
        _ = v9
    }
    _ = p0
    _ = p1
    return p1
}

func main() {
    var v1 bool = ((("" + "a") + ("\n\t" + ("" + ("a" + ("\n\t" + "\n\t"))))) < "a")
    var v2 int16 = 42
    var v3 float64 = (3.14159 - (-2.25 * ((1.5 * 1e10) * 0.0)))
    v4 := make(map[string]int)
    var v5 int16 = (^(v2))
    var v10 int32 = -1
    v2 -= (v5 / ((42) | 1))
    if v1 {
        var v11 int32 = (v10 / ((v10) | 1))
        _ = v11
    } else {
        v12 := []int{-338}
        _ = v12
    }
    v10 = v10
    var v13 string = "a"
    v13 = v13
    var v14 uint = 308
    v15 := GenS1{F0: 65535}
    v16 := []int{}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
