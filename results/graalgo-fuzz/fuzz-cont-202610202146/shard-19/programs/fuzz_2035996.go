package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16 }

func main() {
    var v1 string = "hello"
    v2 := make(map[uint32]string)
    v3 := []uint32{}
    v4 := []string{(v1 + "hello"), "hello", v1}
    v5 := make(map[uint32]uint32)
    switch ((v1 + (v1 + (v1 + "foo bar"))) + "\n\t") {
    case "\n\t":
        var v6 int = 9223372036854775807
        _ = v6
    case "a":
        fmt.Println(v1)
    }
    if len(v4) > 0 {
        v4[0] = (("" + ("a" + "")) + "")
    }
    if false {
        v1 = "a"
        var v7 string = ((("hello" + (("" + v1) + "hello")) + v1) + "")
        _ = v7
    }
    v8 := GenS1{F0: (((1e10 * float64(1e10)) / -2.25) < ((1e10 - (1.5 * 1e10)) - 0.0)), F1: 2}
    _ = len(v3)
    var v9 uint8 = 154
    v9 -= (v9 / ((0) | 1))
    var v10 int32 = 0
    if true {
        for fi11 := 0; fi11 < 4; fi11++ {
            var v12 int8 = 127
            _ = fi11
            _ = v12
        }
        var v13 int16 = -1
        _ = v13
    } else {
        v2[702] = v1
        switch "foo bar" {
        case "hello":
            var v14 string = ((((v1 + ("hello" + v1)) + (v1 + "a")) + v1) + (v1 + "hello"))
            _ = v14
        case "foo bar":
            v4 = append(v4, "")
        }
        var v15 bool = (100 == 127)
        _ = v15
    }
    delete(v5, 2147483647)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
