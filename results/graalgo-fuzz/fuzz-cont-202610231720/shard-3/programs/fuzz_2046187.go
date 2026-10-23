package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint32; F1 int32; F2 uint8 }

func main() {
    var v1 string = "foo bar"
    v2 := make(map[int]uint32)
    var v3 int = -2147483648
    v4 := make(map[string]string)
    var v5 uint8 = 117
    if true {
        var v6 float32 = 1e10
        fmt.Println(v6, v6)
        if true {
            var v7 int32 = -1
            v8 := GenS1{F0: -57}
            v9 := GenS1{F0: 127}
            _ = v7
            _ = v8
            _ = v9
        }
        _ = v6
    } else {
        var v10 uint64 = uint64(v5)
        v1 = ""
        if (9223372036854775807 < (v3 >> ((1) & 63))) {
            if true {
                v11 := GenS1{F0: 100}
                v12 := GenS2{F0: 42, F1: -942, F2: v5}
                _ = v11
                _ = v12
            } else {
                v13 := []int{(v3 << ((65535) & 63)), v3, int(v10)}
                _ = v13
            }
            v3 -= (v3 << ((42) & 63))
        }
        _ = v10
    }
    var v14 uint64 = 946
    v15 := v2[700]
    fmt.Println(v15, v1)
    var v16 int16 = 32767
    v15 = (v15 + v15)
    var v17 int64 = 0
    delete(v4, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
