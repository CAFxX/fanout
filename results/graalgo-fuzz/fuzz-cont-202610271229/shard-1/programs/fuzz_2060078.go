package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string }

func main() {
    v1 := []uint32{}
    var v2 uint64 = 0
    var v3 int8 = -1
    var v4 int8 = (-(v3))
    var v5 uint8 = 250
    var v6 uint64 = (v2 & v2)
    var v7 uint = 65535
    switch ("a" + "foo bar") {
    case "":
        var v8 string = "foo bar"
        _ = v8
    case "hello":
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    v9 := GenS1{F0: "hello", F1: (("\n\t" + "\n\t") + ("a" + ("foo bar" + ("\n\t" + ""))))}
    var v10 int64 = 1
    var v11 uint64 = 1
    if (true || (!(!true))) {
        var v12 uint64 = v2
        var v13 uint16 = 255
        var v14 float64 = float64(3.14159)
        _ = v12
        _ = v13
        _ = v14
    }
    v4 = v4
    var v15 uint32 = 207
    var v16 int64 = 75
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
