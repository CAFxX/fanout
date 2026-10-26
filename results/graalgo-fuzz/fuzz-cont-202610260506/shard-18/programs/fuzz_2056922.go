package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }

func main() {
    v1 := make(map[int]int)
    var v2 uint = 255
    v3 := make(map[uint32]int)
    var v4 uint8 = 128
    v5 := []string{"foo bar"}
    var v6 float64 = 1.5
    v3[583] = 127
    v4 -= v4
    var v7 bool = true
    var v8 uint32 = 12
    v9 := GenS1{F0: 32767}
    v3[(v8 / (((^(v8))) | 1))] = 32768
    var v10 uint32 = 863
    switch (("" + ("foo bar" + "")) + "\n\t") {
    case "":
        var v11 uint = v2
        _ = v11
    case "a":
        switch ("foo bar" + "foo bar") {
        case "hello":
            v7 = (0 != 523)
        }
    case "hello":
        var v12 uint8 = (v4 & (v4 - (^(v4))))
        _ = v12
    }
    fmt.Println(v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
