package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }

func main() {
    var v1 uint8 = 1
    v2 := []uint32{}
    var v3 int64 = 9223372036854775807
    if false {
        v1 -= (v1 - (v1 ^ v1))
        switch ("\n\t" + "\n\t") {
        case "hello":
            v3 -= v3
        }
    }
    v3 += (^(v3))
    var v4 float32 = 3.14159
    if (!false) {
        var v5 uint64 = 2
        _ = len(v2)
        _ = v5
    }
    for fi6 := 0; fi6 < 10; fi6++ {
        v2 = append(v2, 275)
        _ = fi6
    }
    var v7 float64 = float64(0.1)
    v7 = (v7 + (3.14159 + (((v7 + float64(v4)) / v7) + 1.5)))
    if (0.0 == v4) {
        v8 := GenS1{F0: v3}
        if ((v1 ^ 255) == 1) {
            if false {
                fmt.Println(v1)
                var v9 uint = 255
                var v10 string = ("a" + "hello")
                _ = v10
                _ = v9
            } else {
                var v11 float32 = v4
                _ = v11
            }
        }
        _ = v8
    }
    for fi12 := 0; fi12 < 6; fi12++ {
        var v13 float32 = v4
        _ = fi12
        _ = v13
    }
    fmt.Println(v3, v3)
    v4 += (v4 + v4)
    switch "a" {
    case "":
        for fi14 := 0; fi14 < 16; fi14++ {
            v4 = 3.14159
            _ = fi14
        }
    case "foo bar":
        v4 = 0.0
    }
    v3 = v3
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}
