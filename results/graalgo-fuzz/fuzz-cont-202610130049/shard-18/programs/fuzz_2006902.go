package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 string; F1 uint32 }
type GenS3 struct { F0 int8; F1 uint64 }

func main() {
    var v1 float64 = (float64(0.0) * (3.14159 + (float64(1e10) * -2.25)))
    var v2 uint32 = 2147483647
    v3 := make(map[string]uint32)
    v4 := make(map[string]uint32)
    v5 := GenS1{F0: "hello"}
    for fi6 := 0; fi6 < 8; fi6++ {
        var v7 int16 = 1000
        _ = fi6
        _ = v7
    }
    if (!("a" <= "")) {
        v8 := GenS2{F0: "foo bar", F1: v2}
        _ = v8
    }
    var v9 uint32 = (^(v2))
    for fi10 := 0; fi10 < 16; fi10++ {
        v11 := GenS1{F0: ("hello" + "")}
        _ = fi10
        _ = v11
    }
    v9 = v2
    for fi12 := 0; fi12 < 15; fi12++ {
        switch "a" {
        case "foo bar":
            for fi13 := 0; fi13 < 7; fi13++ {
                fmt.Println(fi12, fi13, fi13)
                _ = fi13
            }
        case "hello":
            delete(v4, "\n\t")
        }
        v9 -= v9
        var v14 float32 = -2.25
        _ = fi12
        _ = v14
    }
    fmt.Println(v1)
    var v15 int64 = 1
    v16 := GenS2{F0: ("hello" + ""), F1: (v9 % ((994) | 1))}
    switch "hello" {
    case "foo bar":
        v17 := GenS3{F0: 58, F1: 318}
        _ = v17
    default:
        v9 = 65535
    }
    v18 := GenS1{F0: ("hello" + ("hello" + "foo bar"))}
    if (-9 < 4294967295) {
        v2 = v9
        v1 = v1
    } else {
        for fi19 := 0; fi19 < 17; fi19++ {
            fmt.Println(v15)
            if false {
                var v20 int32 = -197
                v20 = v20
                _ = v20
            }
            _ = fi19
        }
        v2 = 4294967295
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v15
    _ = v16
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
