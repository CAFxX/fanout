package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }

func main() {
    v1 := make(map[string]bool)
    var v2 uint32 = 65535
    var v3 uint8 = 255
    v4 := make(map[uint32]bool)
    var v5 float64 = 3.14159
    v6 := []int{}
    v2 = 2147483648
    _ = len(v6)
    fmt.Println(v3)
    fmt.Println(v2, v3)
    fmt.Println(v3, v3)
    for fi7 := 0; fi7 < 14; fi7++ {
        switch "\n\t" {
        case "a":
            var v8 int16 = 32767
            _ = v8
        case "":
            var v9 float64 = (0.1 / (float64(3.14159) * (v5 + v5)))
            _ = v9
        case "foo bar":
            var v10 int16 = -32768
            _ = v10
        }
        _ = fi7
    }
    var v11 string = ("\n\t" + "\n\t")
    v12 := GenS1{F0: (uint(v3) == 371)}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
