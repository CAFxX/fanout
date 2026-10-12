package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint; F2 int }
type GenS2 struct { F0 int32 }

func main() {
    var v1 uint16 = 732
    var v2 float32 = float32(1.5)
    var v3 string = ("hello" + "a")
    var v4 uint64 = 725
    v5 := GenS1{F0: false, F1: 0, F2: 127}
    v6 := []int{}
    fmt.Println(v4, v4)
    v2 += float32(3.14159)
    v7 := GenS2{F0: 0}
    var v8 uint = 2
    var v9 int16 = 0
    switch 32767 {
    case -189:
        var v10 string = (("a" + "a") + ("\n\t" + ""))
        _ = v10
    case -909:
        v1 += (v1 >> (((^(v8))) & 63))
    case -1:
        var v11 uint8 = 2
        _ = v11
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
