package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8 }

func main() {
    var v1 uint32 = 2147483648
    var v2 int = 45
    var v3 bool = (255 > 351)
    v4 := make(map[int]uint32)
    var v5 int8 = -128
    var v6 uint16 = 32767
    var v7 uint8 = 42
    fmt.Println(v2, v5, v5)
    var v8 int32 = 42
    fmt.Println(v5, v5)
    v7 = v7
    v3 = v3
    var v9 string = ("foo bar" + (("" + ("\n\t" + "\n\t")) + "hello"))
    switch v9 {
    case "":
        for fi10 := 0; fi10 < 5; fi10++ {
            fmt.Println(v5, v6)
            var v11 uint16 = v6
            _ = fi10
            _ = v11
        }
    default:
        v12 := GenS1{F0: 818, F1: (v7 * (v7 | v7))}
        _ = v12
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
