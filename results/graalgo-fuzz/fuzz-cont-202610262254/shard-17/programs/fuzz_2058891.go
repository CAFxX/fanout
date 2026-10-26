package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool; F2 int64 }

func main() {
    var v1 int32 = -632
    v2 := make(map[uint32]uint32)
    v3 := make(map[int]string)
    v4 := []string{("" + ("\n\t" + "\n\t"))}
    var v5 int = 32767
    var v6 string = "a"
    _ = len(v4)
    v7 := GenS1{F0: 450, F1: (!(19 >= 42)), F2: 42}
    v6 = v6
    fmt.Println(v5)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v8 uint = 523
    v1 = v1
    if true {
        v4 = append(v4, (v6 + v6))
    }
    delete(v3, v5)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
