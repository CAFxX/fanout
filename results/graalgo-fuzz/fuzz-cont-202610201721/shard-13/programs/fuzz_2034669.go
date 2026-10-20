package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8 }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[int]bool)
    var v3 int32 = 2147483647
    var v4 float32 = (3.14159 + (0.1 / -2.25))
    var v5 float32 = float32(0.1)
    v6 := make(map[int]int)
    fmt.Println(v5)
    v5 -= 1e10
    fmt.Println(v4, v5)
    var v7 float32 = 0.1
    var v8 string = ""
    if true {
        if false {
            v4 += v4
            v9 := v1[1]
            _ = len(v2)
            _ = v9
        } else {
            var v10 int = 202
            if ((1 > 184) || false) {
                var v11 int = v10
                var v12 bool = ((!true) || (0 > 100))
                delete(v2, (v10 / ((v10) | 1)))
                _ = v11
                _ = v12
            }
            if (2 >= 65535) {
                var v13 int64 = 42
                _ = v13
            } else {
                v14 := []int{}
                _ = v14
            }
            _ = v10
        }
    }
    for fi15 := 0; fi15 < 5; fi15++ {
        if false {
            for fi16 := 0; fi16 < 19; fi16++ {
                var v17 int8 = 42
                var v18 int16 = -622
                var v19 int32 = v3
                _ = fi16
                _ = v17
                _ = v18
                _ = v19
            }
            var v20 int64 = 951
            _ = v20
        } else {
            var v21 uint64 = 838
            v3 = (v3 % ((v3) | 1))
            v7 += 0.0
            _ = v21
        }
        _ = fi15
    }
    v22 := GenS1{F0: 0, F1: -64}
    v4 = 1.5
    v23 := GenS1{F0: -1, F1: 1}
    if false {
        v24 := GenS1{F0: 740, F1: 1}
        _ = v24
    }
    for fi25 := 0; fi25 < 2; fi25++ {
        v8 = "hello"
        _ = fi25
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
