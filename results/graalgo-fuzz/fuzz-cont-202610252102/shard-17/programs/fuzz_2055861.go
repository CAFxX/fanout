package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int; F2 int64 }

func main() {
    var v1 float32 = 0.0
    var v2 int32 = -1
    v3 := make(map[int]uint32)
    fmt.Println(v1)
    for fi4 := 0; fi4 < 7; fi4++ {
        v5 := GenS1{F0: -20, F1: (fi4 & (fi4 & -2147483648)), F2: 1}
        _ = fi4
        _ = v5
    }
    fmt.Println(v1, v2)
    var v6 int = -738
    v7 := GenS1{F0: -32768, F1: (v6 - v6), F2: 42}
    for fi8 := 0; fi8 < 12; fi8++ {
        v1 = float32(float64((v1 / v1)))
        v1 = ((v1 - v1) / float32((-2.25 / (float64(v1) * 1.5))))
        _ = fi8
    }
    var v9 bool = (!("hello" >= "\n\t"))
    v1 += (v1 * (v1 / 0.1))
    switch ("foo bar" + "") {
    case "hello":
        if (!(!(3.14159 != (1.5 + v1)))) {
            v2 = 42
        }
    case "\n\t":
        v3[9223372036854775807] = 336
    case "foo bar":
        v9 = ("foo bar" >= "")
    }
    v9 = false
    for fi10 := 0; fi10 < 18; fi10++ {
        switch ("foo bar" + ((("a" + ("\n\t" + "\n\t")) + "\n\t") + "\n\t")) {
        case "\n\t":
            var v11 uint32 = 0
            _ = v11
        }
        _ = fi10
    }
    v9 = (47 <= 127)
    var v12 float64 = -2.25
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v9
}
