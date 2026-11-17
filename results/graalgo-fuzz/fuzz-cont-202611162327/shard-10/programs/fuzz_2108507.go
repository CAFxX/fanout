package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8 }

func main() {
    var v1 uint16 = 255
    v2 := make(map[int]string)
    v3 := []uint32{927}
    var v4 string = "a"
    switch v4 {
    case "a":
        v5 := v2[32768]
        _ = v5
    case "":
        var v6 uint32 = 1
        _ = v6
    }
    var v7 string = v4
    if len(v3) > 0 {
        v3[0] = 65535
    }
    for fi8 := 0; fi8 < 11; fi8++ {
        v7 = ((("" + (v4 + "a")) + (((v4 + (((v4 + v4) + "") + v7)) + (v7 + (v7 + ""))) + v4)) + v4)
        _ = fi8
    }
    v7 = v7
    var v9 uint32 = 415
    v10 := GenS1{F0: 32767, F1: 1}
    v11 := GenS1{F0: 1, F1: -44}
    if false {
        var v12 uint16 = (v1 * (^(v1)))
        _ = v12
    }
    fmt.Println(v9)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v9
}
