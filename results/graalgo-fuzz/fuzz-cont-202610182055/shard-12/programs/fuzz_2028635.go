package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64; F2 int8 }

func main() {
    var v1 uint8 = 2
    var v2 int8 = 1
    var v3 uint8 = 2
    v4 := []string{"\n\t", "a", "\n\t", "\n\t"}
    if len(v4) > 0 {
        v4[0] = ((((("" + ("a" + ("foo bar" + "hello"))) + "hello") + ("\n\t" + "\n\t")) + "\n\t") + ((("a" + ("\n\t" + "")) + "hello") + ("hello" + "foo bar")))
    }
    v3 -= v3
    var v5 float64 = 1e10
    fmt.Println(v3, v5)
    for fi6 := 0; fi6 < 7; fi6++ {
        v7 := GenS1{F0: 65535, F1: 42, F2: (v2 / (((^(v2))) | 1))}
        v1 -= 255
        _ = fi6
        _ = v7
    }
    v5 -= 3.14159
    switch ("" + "hello") {
    case "hello":
        for fi8 := 0; fi8 < 8; fi8++ {
            for fi9 := 0; fi9 < 14; fi9++ {
                var v10 float32 = 1e10
                _ = fi9
                _ = v10
            }
            v5 = v5
            for fi11 := 0; fi11 < 10; fi11++ {
                fmt.Println(v2, v3, v5)
                var v12 int32 = -2147483648
                v1 = 192
                _ = fi11
                _ = v12
            }
            _ = fi8
        }
    case "a":
        var v13 uint32 = 0
        _ = v13
    default:
        var v14 uint64 = 18446744073709551615
        _ = v14
    }
    v15 := GenS1{F0: 304, F1: 755, F2: 72}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
