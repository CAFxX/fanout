package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32 }

func main() {
    v1 := []string{"", "a", ""}
    v2 := make(map[uint32]string)
    var v3 uint = 4294967295
    if (0 == 2) {
        v3 += (v3 | (v3 % ((v3) | 1)))
    }
    var v4 int16 = 1
    var v5 float64 = (((3.14159 / 1.0) * ((-2.25 - (0.0 / 1e10)) - ((0.0 / float64(float32(0.1))) - 3.14159))) - 0.1)
    switch ("foo bar" + "hello") {
    case "\n\t":
        fmt.Println(v4, v5, v4)
    case "":
        v6 := GenS1{F0: 65535, F1: 65535}
        _ = v6
    }
    v5 = v5
    var v7 uint64 = 4294967295
    switch (("hello" + "a") + "\n\t") {
    case "\n\t":
        var v8 bool = (!(!false))
        _ = v8
    default:
        switch 640 {
        case -9223372036854775808:
            var v9 int64 = 0
            _ = v9
        default:
            for fi10 := 0; fi10 < 15; fi10++ {
                switch 0 {
                case 0:
                    v5 += (1.5 + -2.25)
                case -216:
                    fmt.Println(v7)
                case 2147483647:
                    var v11 int8 = 42
                    _ = v11
                }
                v12 := GenS1{F0: 2147483648, F1: 2147483647}
                var v13 uint = v3
                _ = fi10
                _ = v12
                _ = v13
            }
        }
    }
    switch ("hello" + ("foo bar" + ("\n\t" + "a"))) {
    case "\n\t":
        v5 = (3.14159 / v5)
    default:
        v4 = v4
    }
    if false {
        v3 = 18446744073709551615
        v7 += v7
    } else {
        v1 = append(v1, "foo bar")
    }
    var v14 uint64 = (v7 & (v7 | v7))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
