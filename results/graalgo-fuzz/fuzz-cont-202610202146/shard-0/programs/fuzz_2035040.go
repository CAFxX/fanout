package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool }

func main() {
    var v1 uint = 18446744073709551615
    v2 := make(map[uint32]uint32)
    var v3 float32 = (1e10 / 1.0)
    v4 := make(map[string]string)
    var v5 uint = 18446744073709551615
    var v6 uint8 = uint8(v5)
    var v7 int64 = -1
    for fi8 := 0; fi8 < 7; fi8++ {
        v5 += v5
        var v9 uint16 = 843
        var v10 uint8 = 107
        _ = fi8
        _ = v10
        _ = v9
    }
    var v11 string = "foo bar"
    var v12 uint16 = 32767
    if false {
        delete(v4, (v11 + ((("hello" + ("hello" + v11)) + v11) + (v11 + v11))))
        fmt.Println(v5, v1, v6)
    }
    var v13 int8 = -128
    for fi14 := 0; fi14 < 4; fi14++ {
        v15 := []int{(-(fi14)), fi14, 42, 42}
        v16 := GenS1{F0: "a", F1: (2 >= 1)}
        _ = fi14
        _ = v15
        _ = v16
    }
    v11 = ""
    switch ("" + (v11 + ("foo bar" + ((v11 + "") + v11)))) {
    case "foo bar":
        var v17 uint = v1
        _ = v17
    case "":
        var v18 float64 = (((((-2.25 / float64(3.14159)) - (3.14159 * 3.14159)) / (0.1 / (1.5 * (1.5 - 0.1)))) + ((-2.25 / ((float64(1.5) + 0.1) / float64(-2.25))) * 1.5)) - 0.0)
        _ = v18
    }
    v5 += (v1 & uint(v12))
    if true {
        var v19 uint8 = v6
        v11 = ("\n\t" + v11)
        v4[(("" + ((v11 + (v11 + v11)) + (v11 + v11))) + "foo bar")] = ((("\n\t" + ("a" + v11)) + "foo bar") + v11)
        _ = v19
    } else {
        v20 := GenS1{F0: v11, F1: true}
        _ = v20
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
