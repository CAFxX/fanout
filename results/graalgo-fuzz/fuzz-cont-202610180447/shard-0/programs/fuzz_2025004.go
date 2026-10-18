package main

import (
    "fmt"
)

type GenS1 struct { F0 int }

func main() {
    v1 := []int{2147483647, -1}
    v2 := []int{4294967295, -9223372036854775808, -280}
    v3 := make(map[string]bool)
    var v4 int8 = 1
    v5 := make(map[string]bool)
    var v6 uint = 255
    var v7 bool = (2 > 2)
    v4 += (v4 & (v4 | 0))
    for fi8 := 0; fi8 < 20; fi8++ {
        v9 := GenS1{F0: -225}
        _ = fi8
        _ = v9
    }
    if false {
        _ = len(v2)
    }
    _ = len(v5)
    fmt.Println(v4, v7, v6)
    v10 := GenS1{F0: 128}
    fmt.Println(v4, v6)
    if v7 {
        v4 = (v4 / ((-1) | 1))
    }
    var v11 uint64 = 2
    v4 -= v4
    v6 += 18446744073709551615
    if v7 {
        if true {
            if false {
                var v12 float32 = ((3.14159 / (float32(-2.25) + -2.25)) / (3.14159 / (0.0 - 1e10)))
                _ = v12
            } else {
                _ = len(v1)
            }
        } else {
            var v13 uint64 = (v11 << (((v6 * (v6 | (v6 - v6)))) & 63))
            _ = v13
        }
    } else {
        var v14 int32 = 42
        v15 := GenS1{F0: -2147483648}
        for fi16 := 0; fi16 < 3; fi16++ {
            v17 := GenS1{F0: fi16}
            if v7 {
                var v18 float32 = (float32(1e10) - (float32(1.5) + 1e10))
                for fi19 := 0; fi19 < 1; fi19++ {
                    switch "hello" {
                    case "\n\t":
                        var v20 bool = (771 <= 0)
                        _ = v20
                    }
                    var v21 string = ""
                    _ = fi19
                    _ = v21
                }
                _ = v18
            }
            _ = fi16
            _ = v17
        }
        _ = v14
        _ = v15
    }
    v11 = 0
    var v22 int32 = 42
    for fi23 := 0; fi23 < 4; fi23++ {
        v4 = 100
        _ = fi23
    }
    v24 := GenS1{F0: -174}
    var v25 uint16 = 1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v22
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
