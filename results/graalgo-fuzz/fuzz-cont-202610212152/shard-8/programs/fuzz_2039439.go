package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32 }
type GenS2 struct { F0 string }

func main() {
    var v1 float64 = (0.1 * -2.25)
    v2 := make(map[uint32]bool)
    var v3 int = -1
    if false {
        for fi4 := 0; fi4 < 14; fi4++ {
            var v5 float64 = float64(((1.5 * -2.25) - ((float32((-2.25 - v1)) + 1.5) + ((1.5 * float32(((v1 * float64(0.1)) / (v1 / (float64(((3.14159 / 3.14159) * 1.5)) - v1))))) / 1e10))))
            v6 := []int{2147483647}
            fmt.Println(v5, v5, v5)
            _ = fi4
            _ = v5
            _ = v6
        }
    } else {
        if true {
            fmt.Println(v1, v1)
        }
    }
    switch (("hello" + "a") + "hello") {
    case "hello":
        v7 := GenS1{F0: 65535, F1: 0}
        _ = v7
    case "":
        v8 := GenS1{F0: 32767, F1: -2147483648}
        _ = v8
    }
    var v9 uint32 = 195
    _ = len(v2)
    v1 = (v1 * 3.14159)
    var v10 float64 = v1
    v11 := []int{}
    v10 += 0.0
    v12 := v2[(^(v9))]
    if len(v11) > 0 {
        v11[0] = v3
    }
    if len(v11) >= 2 {
        v11 = v11[1:len(v11)]
    }
    for fi13 := 0; fi13 < 1; fi13++ {
        for fi14 := 0; fi14 < 15; fi14++ {
            v1 += 1.5
            var v15 float64 = v10
            _ = fi14
            _ = v15
        }
        _ = len(v2)
        v2[(v9 >> ((255) & 63))] = v12
        _ = fi13
    }
    v1 = v1
    switch ("" + "\n\t") {
    case "a":
        var v16 int64 = int64(v9)
        _ = v16
    case "":
        v9 = (v9 >> ((1) & 63))
    case "foo bar":
        _ = len(v11)
    }
    v17 := GenS2{F0: ("foo bar" + "\n\t")}
    v18 := v2[(v9 + 1)]
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v9
}
