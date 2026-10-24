package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int }
type GenS2 struct { F0 int; F1 int32; F2 string }
type GenS3 struct { F0 int32; F1 int32; F2 int16 }
func genFn0(p0 int8, p1 uint8, p2 int) uint {
    v7 := GenS1{F0: 442, F1: -2147483648}
    var v8 uint32 = 1
    v9 := []int{}
    var v10 int32 = -2147483648
    for fi11 := 0; fi11 < 6; fi11++ {
        var v12 int8 = -128
        _ = fi11
        _ = v12
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return 228
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[uint32]string)
    var v3 string = "hello"
    v4 := make(map[uint32]uint32)
    var v5 string = ("hello" + "a")
    var v6 uint8 = 1
    v13 := []int{128}
    if len(v13) >= 2 {
        v13 = v13[1:len(v13)]
    }
    v14 := GenS2{F0: 4294967295, F1: -875, F2: "hello"}
    fmt.Println(v3, v5)
    v4[0] = 2147483647
    if true {
        v5 = v3
    } else {
        v5 = (v3 + v5)
    }
    v15 := GenS1{F0: 4294967295, F1: 2147483647}
    v16 := GenS3{F0: -1, F1: -1, F2: 1000}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
