package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64 }

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[string]string)
    v3 := []string{"\n\t", "foo bar", ""}
    var v4 uint64 = 42
    var v5 uint = 582
    var v6 string = ("\n\t" + "hello")
    var v7 uint32 = 1
    v3 = append(v3, (("\n\t" + (((("a" + v6) + "\n\t") + "foo bar") + (v6 + v6))) + v6))
    v4 = v4
    var v8 uint64 = (^(v4))
    v9 := GenS1{F0: 127, F1: (v4 >> ((v5) & 63))}
    fmt.Println(v8)
    v10 := v2["\n\t"]
    switch -2147483648 {
    case -651:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    case 42:
        if len(v3) > 0 {
            v3[0] = "foo bar"
        }
    case -1:
        v10 = v6
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
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
