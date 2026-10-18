package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint8 }

func main() {
    v1 := []uint32{0, 4294967295, 0}
    var v2 uint64 = 0
    var v3 uint16 = 32767
    var v4 bool = true
    var v5 float32 = 0.0
    var v6 int8 = 42
    v5 -= 1.5
    var v7 uint64 = v2
    var v8 uint = 114
    if (20 > (^(v6))) {
        v9 := GenS1{F0: 42}
        fmt.Println(v2, v3)
        if (42 < 2147483647) {
            if v4 {
                for fi10 := 0; fi10 < 16; fi10++ {
                    if len(v1) > 0 {
                        v1[0] = 1
                    }
                    _ = fi10
                }
                for fi11 := 0; fi11 < 4; fi11++ {
                    var v12 string = (((("hello" + "") + "\n\t") + "\n\t") + "foo bar")
                    _ = fi11
                    _ = v12
                }
                v1 = append(v1, 2147483648)
            }
            var v13 bool = v4
            var v14 uint8 = 0
            _ = v13
            _ = v14
        }
        _ = v9
    }
    v15 := GenS1{F0: 1}
    v1 = append(v1, 42)
    switch -2147483648 {
    case 4294967295:
        var v16 uint16 = 238
        _ = v16
    case 190:
        var v17 float32 = ((v5 * ((v5 / (v5 / v5)) * v5)) - (v5 + (v5 + v5)))
        _ = v17
    }
    v18 := GenS1{F0: 1}
    var v19 int64 = 473
    var v20 int64 = (v19 % (((v19 * v19)) | 1))
    if true {
        var v21 int8 = (^(v6))
        var v22 int64 = 42
        _ = v21
        _ = v22
    }
    if (-92 < 42) {
        if (646 > 2) {
            if v4 {
                var v23 float32 = (float32(0.1) + 0.1)
                _ = v23
            } else {
                _ = len(v1)
                v24 := GenS2{F0: v2}
                _ = v24
            }
        } else {
            v8 = (v8 ^ 1)
            v7 += (v2 | (v7 - v2))
        }
        var v25 float64 = (float64(v5) + 1e10)
        _ = v25
    }
    v26 := GenS3{F0: uint8(v3)}
    if true {
        v27 := GenS1{F0: 1}
        v4 = v4
        _ = v27
    }
    v5 = v5
    fmt.Println(v8)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
