package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
func genFn0() bool {
    var v6 int64 = 42
    v7 := []int{602, 1, 603}
    _ = v6
    _ = v7
    return (4294967295 >= 65535)
}

func main() {
    v1 := []int{}
    var v2 string = ""
    v3 := make(map[uint32]string)
    v4 := []uint32{1, 2147483648, 2, 0}
    var v5 string = (v2 + "")
    v8 := GenS1{F0: (1 == 0)}
    for fi9 := 0; fi9 < 3; fi9++ {
        for fi10 := 0; fi10 < 6; fi10++ {
            v5 = (v2 + "hello")
            var v11 uint64 = 2
            var v12 uint32 = 2
            _ = fi10
            _ = v11
            _ = v12
        }
        _ = fi9
    }
    fmt.Println(v5)
    var v13 uint = 18446744073709551615
    v14 := GenS1{F0: true}
    for fi15 := 0; fi15 < 12; fi15++ {
        v3[0] = "hello"
        _ = fi15
    }
    for fi16 := 0; fi16 < 16; fi16++ {
        var v17 int = 42
        v13 -= 139
        _ = fi16
        _ = v17
    }
    if len(v1) > 0 {
        v1[0] = 42
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
