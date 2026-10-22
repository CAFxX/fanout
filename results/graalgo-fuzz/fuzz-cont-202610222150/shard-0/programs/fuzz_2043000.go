package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
func genFn0() string {
    if false {
        v7 := []int{32767, -774}
        var v8 int = -868
        var v9 uint = 255
        _ = v7
        _ = v8
        _ = v9
    } else {
        v10 := GenS1{F0: 42}
        var v11 uint16 = 303
        _ = v10
        _ = v11
    }
    if true {
        var v12 uint64 = 18446744073709551615
        _ = v12
    } else {
        var v13 int32 = -911
        v14 := GenS1{F0: 1}
        _ = v13
        _ = v14
    }
    var v15 int = 127
    v15 = -479
    _ = v15
    return ("" + "foo bar")
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []int{1}
    v3 := make(map[string]uint32)
    v4 := []string{}
    var v5 uint = 0
    var v6 uint32 = 2147483648
    v16 := GenS1{F0: v6}
    v17 := v3["hello"]
    fmt.Println(v5)
    v6 = (v6 % ((v17) | 1))
    v2 = append(v2, 127)
    v5 -= (v5 & v5)
    v17 = v17
    for fi18 := 0; fi18 < 7; fi18++ {
        if true {
            v5 = v5
        }
        _ = fi18
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
