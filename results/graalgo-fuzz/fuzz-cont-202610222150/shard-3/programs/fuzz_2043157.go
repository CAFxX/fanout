package main

import (
    "fmt"
)

type GenS1 struct { F0 string }

func main() {
    var v1 float64 = ((-2.25 + (0.1 + 3.14159)) * 1e10)
    var v2 uint16 = 0
    v3 := make(map[uint32]string)
    v1 += (v1 / 0.1)
    v2 = 221
    switch -1 {
    case 32767:
        v1 = -2.25
    case -2147483648:
        v4 := GenS1{F0: ("\n\t" + "hello")}
        _ = v4
    case 128:
        if (3.14159 > 0.0) {
            var v5 int16 = -32768
            _ = v5
        }
    default:
        v2 = 42
    }
    v1 -= v1
    var v6 uint16 = 42
    fmt.Println(v2, v6, v2)
    v7 := GenS1{F0: ("foo bar" + "")}
    v2 = 2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v6
    _ = v7
}
