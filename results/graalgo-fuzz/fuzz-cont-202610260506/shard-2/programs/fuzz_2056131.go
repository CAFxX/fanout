package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64; F2 uint32 }
type GenS2 struct { F0 int8; F1 string }
func genFn0(p0 uint64, p1 int8) int {
    var v6 bool = true
    v7 := []int{2147483647, 42}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 32768
}
func genFn1(p0 int16, p1 uint64, p2 string) int {
    fmt.Println(p1, p0)
    v8 := GenS1{F0: true, F1: 42, F2: 42}
    if true {
        v9 := GenS1{F0: false, F1: 9223372036854775807, F2: 2}
        if false {
            fmt.Println(p0, p0)
        }
        v10 := GenS2{F0: 42, F1: p2}
        _ = v10
        _ = v9
    }
    fmt.Println(p1)
    v11 := []int{42, 32768, genFn0((p1 | 0), 42)}
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v8
    return 128
}

func main() {
    v1 := make(map[uint32]string)
    var v2 uint = 838
    var v3 float64 = 1e10
    var v4 int32 = -937
    v5 := make(map[int]uint32)
    v12 := GenS1{F0: (false || (128 == 128)), F1: 42, F2: 65535}
    var v13 int = int(v4)
    v14 := GenS2{F0: 0, F1: "hello"}
    var v15 float64 = v3
    v4 = (v4 * v4)
    for fi16 := 0; fi16 < 20; fi16++ {
        v3 = float64(0.0)
        var v17 string = "a"
        fmt.Println(v2)
        _ = fi16
        _ = v17
    }
    v4 -= 2147483647
    v4 -= v4
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
