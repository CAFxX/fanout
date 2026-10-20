package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32 }

func main() {
    var v1 uint = 1
    v2 := []uint32{3}
    v3 := []string{"\n\t"}
    v4 := GenS1{F0: (601 < 495), F1: 268}
    if len(v3) > 0 {
        v3[0] = ("a" + ("" + "\n\t"))
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println(v1)
    for fi5 := 0; fi5 < 19; fi5++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        v1 -= v1
        _ = fi5
    }
    fmt.Println(v1)
    var v6 int = 1
    v7 := GenS1{F0: true, F1: 2147483648}
    switch "" {
    case "":
        var v8 string = ("" + ("hello" + ("hello" + "foo bar")))
        _ = v8
    }
    fmt.Println(v6, v1)
    switch "a" {
    case "a":
        v6 = v6
    case "":
        if (!(1000 == -142)) {
            var v9 float64 = 0.1
            var v10 int32 = -1
            v11 := GenS1{F0: false, F1: uint32(v1)}
            _ = v10
            _ = v11
            _ = v9
        }
    case "hello":
        v1 += (^(v1))
    default:
        _ = len(v3)
    }
    switch v6 {
    case 42:
        var v12 uint8 = 255
        _ = v12
    }
    var v13 float64 = float64(0.1)
    switch v6 {
    case 2147483647:
        var v14 string = "hello"
        _ = v14
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
}
