package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool }

func main() {
    v1 := []uint32{0}
    var v2 float64 = 3.14159
    var v3 bool = true
    var v4 float64 = v2
    var v5 uint64 = 4294967295
    if v3 {
        if (v5 != (v5 & 42)) {
            _ = len(v1)
        }
        var v6 uint = 804
        v7 := GenS1{F0: "foo bar", F1: v3}
        _ = v6
        _ = v7
    }
    if v3 {
        v8 := GenS1{F0: (("" + "\n\t") + "hello"), F1: false}
        _ = v8
    }
    v9 := GenS1{F0: "hello", F1: false}
    if ((127 != -5) && v3) {
        var v10 uint16 = 1
        var v11 string = "hello"
        v11 = (v11 + ((v11 + "hello") + v11))
        _ = v10
        _ = v11
    }
    var v12 int = 127
    v5 -= (v5 >> ((119) & 63))
    if (!v3) {
        v1 = append(v1, 2147483648)
        fmt.Println(v5, v3, v4)
        v12 += 127
    }
    fmt.Println(v2, v4)
    var v13 string = "foo bar"
    v13 = v13
    if (841 >= 1) {
        fmt.Println(v3, v5)
        var v14 float64 = 0.1
        switch v13 {
        case "hello":
            v13 = v13
        case "":
            v15 := GenS1{F0: (v13 + v13), F1: v3}
            _ = v15
        case "\n\t":
            var v16 uint32 = 184
            _ = v16
        }
        _ = v14
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
