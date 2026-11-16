package main

import (
    "fmt"
)

type GenS1 struct { F0 int }

func main() {
    var v1 float32 = (3.14159 - (-2.25 - -2.25))
    v2 := []int{215, 42, -277, 0}
    var v3 bool = (!(42 == 127))
    v4 := []uint32{188}
    for fi5 := 0; fi5 < 19; fi5++ {
        v2 = append(v2, (fi5 - fi5))
        for fi6 := 0; fi6 < 10; fi6++ {
            var v7 uint8 = 241
            _ = fi6
            _ = v7
        }
        if len(v2) > 0 {
            v2[0] = fi5
        }
        _ = fi5
    }
    v3 = (!(42 != 4294967295))
    var v8 uint16 = 206
    if (!v3) {
        v2 = append(v2, -270)
    }
    var v9 uint8 = 223
    v1 = v1
    var v10 uint16 = (v8 % (((v8 >> ((829) & 63))) | 1))
    if (!v3) {
        var v11 int32 = 2147483647
        var v12 bool = (42 >= 64)
        var v13 int8 = 66
        _ = v11
        _ = v12
        _ = v13
    } else {
        if true {
            v14 := GenS1{F0: 127}
            _ = v14
        }
        v15 := GenS1{F0: 128}
        _ = v15
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
