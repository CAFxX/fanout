package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64; F2 uint8 }
type GenS2 struct { F0 uint8; F1 uint32; F2 uint8 }

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[int]bool)
    var v3 bool = (32768 == 585)
    var v4 string = (("hello" + ("\n\t" + "\n\t")) + (("\n\t" + ("\n\t" + "")) + ""))
    var v5 uint64 = 620
    v6 := []string{("foo bar" + v4), (v4 + ""), ("" + ("hello" + v4)), v4}
    v4 = v4
    var v7 int8 = -90
    v5 += v5
    v8 := GenS1{F0: 32768, F1: -1, F2: 127}
    var v9 string = v4
    var v10 int64 = 1
    v11 := GenS2{F0: 255, F1: 2147483648, F2: 42}
    v9 = v9
    var v12 uint16 = 993
    v2[1] = (!true)
    var v13 int8 = 1
    v3 = false
    var v14 bool = v3
    var v15 float32 = 0.0
    var v16 int16 = 961
    _ = len(v6)
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    _ = len(v1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
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
    _ = v7
    _ = v8
    _ = v9
}
