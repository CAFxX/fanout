package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8; F2 bool }

func main() {
    var v1 float64 = -2.25
    v2 := []uint32{}
    v3 := make(map[int]bool)
    var v4 bool = (!(105 < 2147483647))
    var v5 uint = 42
    var v6 uint = (v5 | 0)
    delete(v3, -1)
    v4 = (false || (363 != 990))
    switch ((("foo bar" + "a") + "a") + ("hello" + "")) {
    case "a":
        for fi7 := 0; fi7 < 9; fi7++ {
            v8 := GenS1{F0: -29, F1: 168, F2: true}
            var v9 string = "\n\t"
            _ = fi7
            _ = v8
            _ = v9
        }
    case "":
        v6 += v6
    }
    var v10 float64 = (3.14159 * -2.25)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v5 -= (v6 & (v5 << ((v6) & 63)))
    v6 -= (v5 ^ 65535)
    v1 -= (v10 / (v1 * float64(3.14159)))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
