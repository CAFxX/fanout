package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
func genFn0() int8 {
    var v7 bool = ("\n\t" > "\n\t")
    if v7 {
        v7 = (v7 && true)
        v8 := GenS1{F0: 42}
        _ = v8
    } else {
        fmt.Println(v7)
        for fi9 := 0; fi9 < 15; fi9++ {
            v7 = (618 != 814)
            var v10 uint32 = 954
            _ = fi9
            _ = v10
        }
    }
    var v11 float64 = -2.25
    _ = v11
    _ = v7
    return -111
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint16 = 255
    v3 := make(map[uint32]uint32)
    v4 := []int{}
    var v5 uint16 = (v2 | v2)
    var v6 int64 = -1
    v6 = int64(v2)
    fmt.Println(v6)
    var v12 string = (("a" + "") + "\n\t")
    v5 = 2
    var v13 int64 = (v6 + v6)
    delete(v1, 65535)
    if false {
        v14 := GenS1{F0: 48}
        v6 = v6
        _ = v14
    } else {
        var v15 uint = 0
        if false {
            v16 := GenS1{F0: 1}
            _ = v16
        }
        _ = v15
    }
    var v17 uint16 = (v5 % ((v2) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
