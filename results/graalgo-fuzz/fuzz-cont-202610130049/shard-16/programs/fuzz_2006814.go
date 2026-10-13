package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8 }
type GenS2 struct { F0 string }

func main() {
    var v1 string = ("hello" + "")
    var v2 uint64 = 984
    var v3 int8 = 0
    var v4 uint16 = 42
    var v5 bool = false
    var v6 string = v1
    v1 = (((v1 + ((v6 + (v1 + v6)) + (v6 + (v1 + v6)))) + v1) + v1)
    v7 := []int{128}
    v3 = (v3 - 1)
    if (!v5) {
        var v8 int64 = int64(v4)
        _ = v8
    }
    fmt.Println(v3)
    var v9 float32 = -2.25
    if len(v7) > 0 {
        v7[0] = -243
    }
    v10 := GenS1{F0: 189, F1: 127}
    var v11 uint16 = (v4 / ((32767) | 1))
    fmt.Println(v11, v4, v6)
    for fi12 := 0; fi12 < 6; fi12++ {
        for fi13 := 0; fi13 < 8; fi13++ {
            var v14 uint = 65535
            _ = fi13
            _ = v14
        }
        _ = fi12
    }
    switch (v1 + "") {
    case "\n\t":
        if len(v7) > 0 {
            v7[0] = 128
        }
    case "hello":
        v1 = v6
    case "a":
        if v5 {
            if (255 < 127) {
                var v15 uint16 = 0
                v16 := GenS2{F0: "foo bar"}
                _ = v15
                _ = v16
            } else {
                fmt.Println(v6, v5, v6)
            }
            var v17 uint16 = 42
            _ = v17
        }
    default:
        fmt.Println(v9)
    }
    for fi18 := 0; fi18 < 9; fi18++ {
        var v19 int16 = -238
        _ = fi18
        _ = v19
    }
    switch 32768 {
    case 0:
        switch 127 {
        case 0:
            v11 = (v4 + v11)
        case 620:
            fmt.Println(v9, v5)
        default:
            _ = len(v7)
        }
    default:
        var v20 int32 = 42
        _ = v20
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
