package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8 }
type GenS2 struct { F0 uint; F1 string }
func genFn0(p0 int64, p1 uint64, p2 string) bool {
    var v4 uint8 = 2
    v4 -= 32
    p1 += p1
    p1 -= p1
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return false
}
func genFn1() uint {
    var v5 int8 = 127
    v6 := GenS1{F0: 1, F1: 255}
    v7 := GenS2{F0: 2, F1: "\n\t"}
    v8 := GenS2{F0: 65535, F1: "hello"}
    var v9 int64 = 0
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 255
}

func main() {
    v1 := make(map[string]string)
    var v2 uint8 = 127
    var v3 uint8 = 11
    var v10 int64 = 396
    v3 = v2
    v11 := []int{int(v3), 1, 4294967295, 32767}
    v12 := GenS2{F0: 0, F1: ("" + "a")}
    v13 := GenS2{F0: 328, F1: ("hello" + "hello")}
    var v14 int64 = v10
    var v15 uint64 = 18446744073709551615
    v2 -= (v2 * 128)
    v10 += 1
    v3 -= v3
    var v16 uint = genFn1()
    var v17 int32 = 2147483647
    v2 = (v2 | v2)
    var v18 uint16 = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
}
