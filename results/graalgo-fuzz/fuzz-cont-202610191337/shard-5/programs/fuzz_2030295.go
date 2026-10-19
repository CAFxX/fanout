package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint; F1 int16; F2 string }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint16; F1 bool; F2 bool }
func genFn0(p0 int16) int16 {
    p0 += (p0 & p0)
    p0 -= p0
    _ = p0
    return (-(p0))
}

func main() {
    var v1 uint8 = 21
    v2 := []int{4294967295, 975}
    v3 := make(map[string]uint32)
    var v4 int16 = 1
    var v5 int64 = 137
    v6 := GenS1{F0: true}
    v7 := GenS1{F0: (655 <= 1)}
    v8 := GenS2{F0: 255, F1: (v4 - genFn0((v4 * v4))), F2: ("hello" + "hello")}
    v9 := GenS3{F0: 4294967295}
    v4 = 0
    v10 := GenS3{F0: 2}
    var v11 float64 = -2.25
    v12 := GenS4{F0: 32767, F1: true, F2: true}
    fmt.Println(v11)
    for fi13 := 0; fi13 < 20; fi13++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        v3[(((("hello" + "a") + "") + "a") + "")] = 0
        _ = fi13
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
