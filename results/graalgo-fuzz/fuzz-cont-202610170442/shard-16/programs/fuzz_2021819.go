package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16; F2 uint64 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int; F1 int8 }
func genFn0() int16 {
    var v7 int8 = -48
    fmt.Println(v7)
    _ = v7
    return 567
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[int]bool)
    var v3 int32 = -1
    v4 := []int{4294967295}
    var v5 float32 = 3.14159
    var v6 int32 = (v3 ^ -1)
    v8 := GenS1{F0: (v6 + 0), F1: -660, F2: 201}
    fmt.Println(v3)
    v9 := GenS2{F0: 42}
    v3 -= 2147483647
    var v10 bool = ((0 == 65535) || false)
    v5 -= 0.1
    v11 := GenS2{F0: -760}
    var v12 int8 = 1
    fmt.Println(v5, v6)
    var v13 uint32 = 65535
    var v14 uint32 = v13
    v15 := GenS3{F0: 1, F1: (v12 - -128)}
    var v16 string = "a"
    _ = len(v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
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
    _ = v8
    _ = v9
}
