package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8; F2 uint32 }

func main() {
    var v1 float64 = float64(3.14159)
    var v2 string = "a"
    v3 := make(map[uint32]string)
    var v4 int = 859
    var v5 int = 819
    switch v2 {
    case "\n\t":
        var v6 float64 = v1
        _ = v6
    case "":
        delete(v3, 42)
    }
    v7 := GenS1{F0: 0, F1: 127, F2: 2}
    var v8 float64 = (float64((3.14159 + (1e10 - (-2.25 * -2.25)))) + v1)
    switch v5 {
    case 127:
        v4 = (v4 | (v5 >> ((42) & 63)))
    case 0:
        for fi9 := 0; fi9 < 14; fi9++ {
            v10 := v3[2147483647]
            _ = fi9
            _ = v10
        }
    case 32767:
        delete(v3, 1)
    default:
        fmt.Println(v8)
    }
    var v11 uint64 = 899
    v4 = v5
    switch (v4 ^ v4) {
    case -714:
        switch (v2 + (("a" + ("" + "hello")) + "a")) {
        case "\n\t":
            v2 = (((v2 + v2) + v2) + v2)
        default:
            var v12 float64 = v1
            _ = v12
        }
    default:
        v13 := []int{(v4 & v4)}
        _ = v13
    }
    var v14 string = "foo bar"
    v15 := GenS1{F0: 42, F1: 100, F2: 1}
    if true {
        var v16 uint16 = 187
        _ = v16
    } else {
        var v17 int8 = 42
        var v18 float64 = v8
        if false {
            v19 := GenS1{F0: -1, F1: 100, F2: 65535}
            _ = v19
        }
        _ = v17
        _ = v18
    }
    v1 = float64(((-2.25 - (-2.25 / float32(v8))) + 0.1))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
