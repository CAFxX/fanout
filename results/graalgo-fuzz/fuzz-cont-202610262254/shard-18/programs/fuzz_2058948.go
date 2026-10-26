package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64 }

func main() {
    v1 := make(map[int]string)
    v2 := make(map[uint32]bool)
    var v3 uint16 = 2
    v4 := make(map[int]string)
    v5 := []int{2147483647, 177}
    v6 := make(map[string]int)
    v7 := GenS1{F0: 1, F1: -203}
    _ = len(v6)
    delete(v4, 483)
    delete(v1, -694)
    var v8 float32 = 0.1
    fmt.Println(v8)
    switch -442 {
    case -34:
        v5 = append(v5, 252)
    case -1:
        for fi9 := 0; fi9 < 11; fi9++ {
            if true {
                var v10 int8 = 127
                _ = v10
            }
            v8 = v8
            if (false && (-461 == 9223372036854775807)) {
                fmt.Println(v3)
                var v11 uint64 = 344
                _ = v11
            }
            _ = fi9
        }
    case -2147483648:
        v3 = 32767
    }
    var v12 int8 = 85
    if len(v5) > 0 {
        v5[0] = 128
    }
    var v13 float64 = 1e10
    fmt.Println(v13, v12)
    fmt.Println(v8)
    delete(v1, 32767)
    var v14 bool = (!(v3 > v3))
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v15 int = 127
    var v16 int = 187
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
