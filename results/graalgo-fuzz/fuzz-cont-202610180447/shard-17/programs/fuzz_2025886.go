package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }

func main() {
    var v1 int8 = 100
    var v2 int32 = -1
    var v3 int = 2147483647
    var v4 uint = 42
    var v5 int64 = 42
    v6 := make(map[uint32]uint32)
    _ = len(v6)
    fmt.Println(v1, v4, v5)
    v2 -= v2
    var v7 float32 = (-2.25 - (-2.25 * (1.5 + 1.5)))
    var v8 string = "foo bar"
    var v9 string = v8
    v3 = v3
    for fi10 := 0; fi10 < 10; fi10++ {
        var v11 uint16 = 42
        _ = fi10
        _ = v11
    }
    v12 := GenS1{F0: -1}
    for fi13 := 0; fi13 < 1; fi13++ {
        v4 = 4294967295
        v14 := []int{}
        _ = fi13
        _ = v14
    }
    fmt.Println(v8, v4)
    v15 := []int{(v3 & v3), (v3 >> (((v4 - v4)) & 63)), v3, -9223372036854775808}
    for fi16 := 0; fi16 < 8; fi16++ {
        if false {
            for fi17 := 0; fi17 < 10; fi17++ {
                fmt.Println(v9, v5, v3)
                var v18 string = v8
                _ = fi17
                _ = v18
            }
            var v19 float32 = v7
            _ = v19
        }
        if true {
            for fi20 := 0; fi20 < 19; fi20++ {
                v9 = v9
                switch (fi16 * (^(fi20))) {
                case -351:
                    var v21 uint64 = 0
                    _ = v21
                default:
                    delete(v6, 2147483647)
                }
                _ = fi20
            }
        }
        _ = fi16
    }
    var v22 int = v3
    fmt.Println(v8, v4, v1)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v15
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
