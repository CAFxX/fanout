package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int8; F1 int8; F2 string }
type GenS3 struct { F0 uint8; F1 int16; F2 string }

func main() {
    v1 := make(map[int]uint32)
    var v2 uint32 = 0
    var v3 bool = ((!true) || (!(((986 > 2) && true) && (32768 >= 4294967295))))
    v4 := make(map[uint32]int)
    v5 := make(map[string]int)
    if (-612 > -9223372036854775808) {
        v6 := GenS1{F0: -1}
        fmt.Println(v2, v2)
        _ = v6
    }
    v2 = v2
    fmt.Println(v3)
    v7 := GenS2{F0: int8(v2), F1: 1, F2: "a"}
    var v8 float64 = -2.25
    switch 502 {
    case -228:
        v9 := GenS3{F0: 60, F1: 627, F2: "a"}
        _ = v9
    case 638:
        var v10 bool = (((!(v3 && v3)) && v3) && (42 != 1000))
        _ = v10
    }
    if (0 <= 615) {
        var v11 bool = (!(!(!(v3 && (-101 != 0)))))
        _ = v11
    } else {
        var v12 bool = (v2 == v2)
        if v3 {
            for fi13 := 0; fi13 < 20; fi13++ {
                fmt.Println(fi13, fi13)
                var v14 uint8 = 255
                _ = fi13
                _ = v14
            }
        }
        _ = v12
    }
    fmt.Println(v3)
    _ = len(v5)
    var v15 int64 = -213
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
