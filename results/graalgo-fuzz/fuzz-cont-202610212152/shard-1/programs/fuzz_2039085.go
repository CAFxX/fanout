package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool }
type GenS2 struct { F0 bool }

func main() {
    var v1 int64 = 452
    v2 := make(map[int]int)
    v3 := []string{}
    switch (("hello" + (((((("hello" + ("hello" + "a")) + "") + ("a" + "\n\t")) + "\n\t") + "a") + "foo bar")) + ("" + "hello")) {
    case "foo bar":
        v4 := GenS1{F0: 32767, F1: false}
        _ = v4
    case "a":
        var v5 string = "hello"
        _ = v5
    case "hello":
        var v6 string = ""
        _ = v6
    default:
        var v7 float64 = 0.0
        _ = v7
    }
    var v8 uint64 = 178
    v1 += 834
    for fi9 := 0; fi9 < 15; fi9++ {
        switch (fi9 % (((-(fi9))) | 1)) {
        case 32768:
            v1 -= (v1 << ((473) & 63))
        case 127:
            v8 = v8
        case 42:
            v1 = v1
        default:
            fmt.Println(v8)
        }
        for fi10 := 0; fi10 < 18; fi10++ {
            if (true && (uint(v8) > 255)) {
                var v11 float64 = float64(0.1)
                var v12 int32 = 0
                _ = v11
                _ = v12
            }
            var v13 uint64 = 960
            _ = fi10
            _ = v13
        }
        _ = fi9
    }
    v8 = v8
    if ((("foo bar" + ("foo bar" + "a")) + ("a" + "hello")) <= "\n\t") {
        v14 := v2[128]
        _ = v14
    }
    for fi15 := 0; fi15 < 12; fi15++ {
        v8 = (v8 - v8)
        _ = fi15
    }
    if false {
        v16 := GenS2{F0: true}
        fmt.Println(v1, v8)
        _ = len(v3)
        _ = v16
    }
    var v17 uint8 = 201
    if len(v3) > 0 {
        v3[0] = ("hello" + ("hello" + "\n\t"))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v8
}
