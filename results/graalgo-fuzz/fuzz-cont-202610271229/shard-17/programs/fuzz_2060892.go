package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint; F1 uint }
type GenS4 struct { F0 string; F1 uint64 }

func main() {
    var v1 float64 = 0.1
    var v2 int16 = 32767
    var v3 int32 = 42
    var v4 int8 = -128
    var v5 float32 = ((0.0 - 3.14159) / (0.0 / float32(v1)))
    v6 := GenS1{F0: 240, F1: ("\n\t" < "\n\t")}
    fmt.Println(v2, v3)
    switch 32767 {
    case 961:
        var v7 int32 = (v3 >> ((913) & 63))
        _ = v7
    case 195:
        for fi8 := 0; fi8 < 15; fi8++ {
            var v9 int8 = v4
            v2 = 1
            var v10 int8 = 0
            _ = fi8
            _ = v10
            _ = v9
        }
    }
    v2 = v2
    v11 := GenS2{F0: ("foo bar" + "\n\t")}
    v12 := []int{1, 42, -395}
    v13 := GenS2{F0: ""}
    switch "" {
    case "foo bar":
        v14 := GenS1{F0: 1, F1: true}
        _ = v14
    }
    if (!true) {
        _ = len(v12)
        v15 := GenS3{F0: 630, F1: 42}
        v16 := GenS1{F0: 42, F1: true}
        _ = v15
        _ = v16
    } else {
        v3 += -191
    }
    v2 -= v2
    if false {
        v1 = v1
    }
    if (false && true) {
        _ = len(v12)
        v5 += v5
    }
    v3 = v3
    v17 := GenS3{F0: 765, F1: 4294967295}
    v18 := GenS4{F0: ("a" + ""), F1: 4294967295}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
