package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string }

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 uint16 = 42
    var v3 uint32 = 476
    var v4 int = 32767
    switch v4 {
    case -59:
        fmt.Println(v2, v1)
    }
    v5 := []int{(v4 / ((v4) | 1)), (v4 & (v4 % ((v4) | 1)))}
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    for fi6 := 0; fi6 < 4; fi6++ {
        var v7 float32 = 0.1
        var v8 uint = 2
        var v9 uint = 1
        _ = fi6
        _ = v7
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 11; fi10++ {
        var v11 uint16 = (v2 * (v2 / ((v2) | 1)))
        v12 := GenS1{F0: ("hello" + ("foo bar" + "hello")), F1: "a"}
        var v13 bool = (!false)
        _ = fi10
        _ = v11
        _ = v12
        _ = v13
    }
    v4 += int(v2)
    for fi14 := 0; fi14 < 19; fi14++ {
        fmt.Println(v1)
        _ = fi14
    }
    fmt.Println(v3)
    _ = len(v5)
    if (!(!false)) {
        fmt.Println(v1, v1)
    } else {
        v5 = append(v5, 0)
        var v15 uint16 = 1
        _ = v15
    }
    var v16 uint8 = 180
    for fi17 := 0; fi17 < 19; fi17++ {
        v1 -= (^(v1))
        fmt.Println(v2, v16)
        _ = len(v5)
        _ = fi17
    }
    if false {
        if false {
            var v18 int32 = int32(v16)
            _ = len(v5)
            _ = v18
        }
    } else {
        switch 0 {
        case 127:
            var v19 uint64 = (v1 % ((v1) | 1))
            _ = v19
        }
        var v20 uint = 2
        _ = v20
    }
    fmt.Println(v2)
    switch "" {
    case "hello":
        v16 = 127
    case "a":
        var v21 uint64 = (v1 << ((2) & 63))
        _ = v21
    case "\n\t":
        v4 -= (v4 % ((v4) | 1))
    }
    var v22 string = ""
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v16
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
