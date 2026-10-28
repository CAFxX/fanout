package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int32; F1 uint32 }
type GenS3 struct { F0 uint16 }
func genFn0() bool {
    v7 := GenS1{F0: ""}
    var v8 int = 32767
    v8 = -295
    v9 := GenS1{F0: ""}
    _ = v7
    _ = v8
    _ = v9
    return true
}

func main() {
    v1 := make(map[string]string)
    var v2 int8 = 100
    v3 := make(map[uint32]bool)
    var v4 float64 = 0.0
    v5 := []uint32{659}
    var v6 uint = 582
    v6 = v6
    var v10 int32 = 0
    v11 := GenS2{F0: (v10 & (v10 % ((442) | 1))), F1: 2147483648}
    v6 -= v6
    var v12 string = "hello"
    v13 := GenS3{F0: 424}
    var v14 uint8 = 2
    switch v12 {
    case "hello":
        v10 -= 1
    }
    if (false && (v4 < v4)) {
        _ = len(v5)
    }
    v4 -= v4
    var v15 int8 = (v2 / ((v2) | 1))
    var v16 string = (v12 + v12)
    for fi17 := 0; fi17 < 1; fi17++ {
        v10 += v10
        v15 = v2
        _ = fi17
    }
    v18 := v3[42]
    v18 = v18
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
