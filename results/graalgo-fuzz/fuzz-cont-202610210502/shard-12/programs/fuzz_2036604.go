package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8; F2 int64 }

func main() {
    v1 := make(map[uint32]int)
    var v2 uint = 959
    v3 := []string{"hello"}
    var v4 uint32 = 2
    var v5 float32 = -2.25
    v2 += v2
    fmt.Println(v4, v5)
    var v6 int8 = 1
    v2 -= uint(v4)
    v2 += v2
    switch ("a" + ("\n\t" + "\n\t")) {
    case "foo bar":
        var v7 bool = true
        _ = v7
    default:
        v8 := GenS1{F0: 120, F1: 64, F2: 583}
        _ = v8
    }
    if (!true) {
        v6 += (v6 | (v6 | (v6 & v6)))
        v2 = v2
        var v9 string = ("foo bar" + "hello")
        _ = v9
    } else {
        for fi10 := 0; fi10 < 4; fi10++ {
            switch ("hello" + "hello") {
            case "hello":
                var v11 uint8 = 128
                _ = v11
            }
            _ = fi10
        }
    }
    for fi12 := 0; fi12 < 20; fi12++ {
        fmt.Println(v6)
        var v13 uint32 = v4
        _ = fi12
        _ = v13
    }
    _ = len(v3)
    for fi14 := 0; fi14 < 7; fi14++ {
        v5 = (1e10 / 1.0)
        v4 = v4
        var v15 int32 = -1
        _ = fi14
        _ = v15
    }
    v1[(v4 * v4)] = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
