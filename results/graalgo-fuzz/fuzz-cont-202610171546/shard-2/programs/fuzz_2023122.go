package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8 }
type GenS2 struct { F0 uint; F1 int64; F2 uint16 }
type GenS3 struct { F0 bool }
func genFn0() string {
    if false {
        v7 := []int{395, 127}
        v8 := GenS1{F0: true, F1: 42}
        _ = v7
        _ = v8
    }
    var v9 int16 = 0
    v10 := []int{32767, 2147483647, 42}
    _ = v10
    _ = v9
    return "hello"
}
func genFn1(p0 uint8) uint8 {
    p0 = (p0 % (((^(p0))) | 1))
    if true {
        var v11 uint = 4294967295
        fmt.Println(p0, p0)
        _ = v11
    }
    var v12 uint16 = uint16(p0)
    p0 = (p0 % ((42) | 1))
    _ = p0
    _ = v12
    return p0
}

func main() {
    v1 := make(map[uint32]string)
    var v2 uint64 = 847
    v3 := []string{"a", "hello", ("hello" + "hello")}
    v4 := []uint32{}
    var v5 uint8 = 255
    var v6 bool = (42 > 65535)
    var v13 string = "\n\t"
    fmt.Println(v6)
    for fi14 := 0; fi14 < 17; fi14++ {
        fmt.Println(v6, v13)
        var v15 uint16 = 361
        _ = fi14
        _ = v15
    }
    v6 = v6
    v16 := GenS2{F0: 255, F1: -219, F2: 548}
    v4 = append(v4, 783)
    var v17 int64 = 532
    v2 = (v2 - 1)
    v18 := GenS2{F0: 65535, F1: v17, F2: 1}
    fmt.Println(v6)
    var v19 string = ("a" + "foo bar")
    if (!true) {
        var v20 int64 = (v17 * (-(v17)))
        var v21 int64 = v17
        v2 += 2
        _ = v20
        _ = v21
    }
    v22 := GenS3{F0: v6}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
