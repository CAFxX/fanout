package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint }

func main() {
    v1 := []string{"foo bar", ("" + (("\n\t" + "hello") + "foo bar")), ("foo bar" + "hello")}
    v2 := []uint32{42, 243, 652, 0}
    v3 := make(map[string]uint32)
    v4 := make(map[string]int)
    v5 := make(map[uint32]int)
    v6 := make(map[uint32]uint32)
    var v7 uint8 = 2
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println(v7)
    fmt.Println(v7)
    switch "a" {
    case "\n\t":
        var v8 int8 = -128
        _ = v8
    case "a":
        v7 = (v7 | 89)
    }
    v7 = v7
    for fi9 := 0; fi9 < 4; fi9++ {
        v7 = 2
        _ = fi9
    }
    v10 := GenS1{F0: 32767, F1: 255}
    switch "a" {
    case "\n\t":
        var v11 int8 = -75
        _ = v11
    case "":
        v4["hello"] = -239
    case "hello":
        switch "hello" {
        case "":
            delete(v6, 2147483648)
        }
    }
    if false {
        v7 = v7
    }
    var v12 int16 = 37
    var v13 int32 = 142
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
