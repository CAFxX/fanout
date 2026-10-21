package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string; F2 uint8 }
type GenS2 struct { F0 uint8 }

func main() {
    var v1 float64 = (-2.25 - 0.0)
    v2 := make(map[uint32]bool)
    v3 := make(map[int]string)
    var v4 string = "a"
    var v5 uint16 = 65535
    fmt.Println(v4, v1, v4)
    switch v4 {
    case "":
        var v6 float32 = (3.14159 + (1.5 / ((-2.25 + (0.1 * float32(1.5))) - 0.0)))
        _ = v6
    case "hello":
        v7 := GenS1{F0: -1, F1: v4, F2: 255}
        _ = v7
    }
    v4 = (v4 + ("\n\t" + ("foo bar" + "foo bar")))
    for fi8 := 0; fi8 < 16; fi8++ {
        v9 := GenS2{F0: 1}
        v10 := GenS1{F0: 42, F1: v4, F2: 255}
        if (128 <= 42) {
            var v11 float32 = 0.1
            _ = v11
        }
        _ = fi8
        _ = v10
        _ = v9
    }
    v12 := GenS1{F0: 2147483647, F1: "foo bar", F2: 117}
    v4 = ""
    var v13 int8 = 1
    for fi14 := 0; fi14 < 15; fi14++ {
        for fi15 := 0; fi15 < 13; fi15++ {
            v5 += (v5 & 0)
            fmt.Println(fi15, v13)
            v4 = ""
            _ = fi15
        }
        var v16 int8 = v13
        var v17 uint = 65535
        _ = fi14
        _ = v16
        _ = v17
    }
    v18 := v2[2147483648]
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
