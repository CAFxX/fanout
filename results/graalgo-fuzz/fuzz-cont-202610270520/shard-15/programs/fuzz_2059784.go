package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool }

func main() {
    v1 := make(map[string]int)
    v2 := make(map[uint32]int)
    v3 := []int{128, 128, 495, -837}
    var v4 int8 = 127
    for fi5 := 0; fi5 < 18; fi5++ {
        v4 = v4
        _ = fi5
    }
    for fi6 := 0; fi6 < 2; fi6++ {
        fmt.Println(v4)
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi6
    }
    v4 -= v4
    var v7 uint = 65535
    if true {
        if false {
            var v8 uint8 = 185
            _ = v8
        }
        v9 := GenS1{F0: 0, F1: (!false)}
        for fi10 := 0; fi10 < 15; fi10++ {
            var v11 float32 = 3.14159
            _ = fi10
            _ = v11
        }
        _ = v9
    } else {
        if ((2147483647 <= 367) || (("" + "\n\t") >= ("\n\t" + "foo bar"))) {
            var v12 int8 = 100
            var v13 int64 = 42
            v4 = -1
            _ = v12
            _ = v13
        }
    }
    v4 = (v4 / ((-25) | 1))
    for fi14 := 0; fi14 < 2; fi14++ {
        switch "" {
        case "":
            var v15 uint8 = 49
            _ = v15
        case "hello":
            v4 = (v4 * v4)
        default:
            var v16 uint32 = 1
            _ = v16
        }
        _ = fi14
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}
