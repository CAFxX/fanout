package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32 }

func main() {
    v1 := make(map[string]string)
    v2 := []string{"", "hello", "hello"}
    var v3 uint16 = 1
    v4 := GenS1{F0: (4294967295 != 593), F1: uint32(v3)}
    v5 := GenS1{F0: (128 >= 140), F1: 2147483648}
    var v6 int32 = -1
    var v7 float64 = float64(1e10)
    var v8 uint = 65535
    var v9 int64 = 9223372036854775807
    v2 = append(v2, "a")
    v2 = append(v2, ("\n\t" + "foo bar"))
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
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
