package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64 }
type GenS2 struct { F0 string; F1 int16 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int16; F1 string; F2 bool }
func genFn0(p0 int32) uint16 {
    var v5 string = "\n\t"
    v6 := GenS1{F0: true, F1: 594}
    var v7 int32 = (^(p0))
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return 86
}
func genFn1(p0 uint64, p1 uint, p2 int64) string {
    var v8 float64 = (0.1 / 1.5)
    var v9 int32 = 0
    v10 := GenS2{F0: "a", F1: 0}
    var v11 uint64 = (p0 ^ 499)
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v8
    _ = v9
    return "hello"
}

func main() {
    var v1 uint16 = 2
    var v2 int16 = -32768
    v3 := []uint32{4294967295, 2}
    v4 := make(map[string]uint32)
    var v12 uint = 656
    var v13 int8 = 24
    v14 := GenS3{F0: true}
    v15 := GenS4{F0: (^(v2)), F1: ("a" + "a"), F2: false}
    v2 += v2
    v1 = v1
    v3 = append(v3, 65535)
    v2 = 0
    var v16 uint32 = 172
    v12 += 452
    var v17 uint64 = 320
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
