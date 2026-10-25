package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool; F2 bool }

func main() {
    var v1 float32 = ((-2.25 - 3.14159) * float32(1.5))
    var v2 int = 9223372036854775807
    v3 := []string{(("" + ("" + "foo bar")) + ""), "a", "", ("" + "\n\t")}
    var v4 uint8 = 127
    var v5 uint64 = 2
    fmt.Println(v4, v1)
    v2 = 4294967295
    v2 -= (-(v2))
    v6 := GenS1{F0: 32767, F1: ((v2 > 128) && (199 < 341)), F2: true}
    if len(v3) > 0 {
        v3[0] = ""
    }
    var v7 uint32 = 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
