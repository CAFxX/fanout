package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint8; F1 int16 }
func genFn0(p0 int16) uint8 {
    var v5 string = ""
    var v6 int = 662
    for fi7 := 0; fi7 < 15; fi7++ {
        var v8 uint64 = 659
        v9 := GenS1{F0: -1}
        _ = fi7
        _ = v8
        _ = v9
    }
    _ = p0
    _ = v5
    _ = v6
    return 0
}

func main() {
    var v1 float32 = -2.25
    v2 := []uint32{2147483647}
    v3 := []string{"foo bar", "a", "\n\t", ""}
    var v4 string = ("a" + "")
    var v10 uint32 = 4294967295
    switch v4 {
    case "foo bar":
        v11 := GenS2{F0: 0, F1: 29}
        _ = v11
    case "a":
        for fi12 := 0; fi12 < 10; fi12++ {
            v1 += 0.0
            _ = fi12
        }
    }
    fmt.Println(v10)
    v10 = v10
    if len(v2) > 0 {
        v2[0] = (v10 >> ((65535) & 63))
    }
    v4 = (v4 + "")
    var v13 int = 577
    var v14 int16 = 711
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
}
