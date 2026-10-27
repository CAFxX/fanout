package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int }
func genFn0(p0 uint16, p1 int, p2 uint8) int64 {
    v7 := []int{32767, (^(p1))}
    if (65535 > 0) {
        _ = len(v7)
        if false {
            for fi8 := 0; fi8 < 9; fi8++ {
                for fi9 := 0; fi9 < 4; fi9++ {
                    v10 := GenS1{F0: 18446744073709551615}
                    v11 := GenS1{F0: 42}
                    if true {
                        var v12 uint8 = p2
                        var v13 uint = 4294967295
                        _ = v12
                        _ = v13
                    }
                    _ = fi9
                    _ = v10
                    _ = v11
                }
                _ = len(v7)
                _ = fi8
            }
        }
        fmt.Println(p1, p2, p0)
    }
    v14 := GenS1{F0: 18446744073709551615}
    var v15 int64 = -9223372036854775808
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    _ = v15
    _ = v7
    return v15
}

func main() {
    v1 := []uint32{2147483647, 42}
    var v2 uint8 = 29
    var v3 uint8 = 1
    var v4 int32 = 0
    var v5 float32 = 0.1
    var v6 int16 = -930
    fmt.Println(v6, v3)
    var v16 uint8 = v2
    v17 := GenS1{F0: 313}
    v4 += v4
    var v18 uint32 = 1
    var v19 uint8 = v3
    v19 = 2
    v20 := GenS2{F0: 128}
    if len(v1) > 0 {
        v1[0] = v18
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
