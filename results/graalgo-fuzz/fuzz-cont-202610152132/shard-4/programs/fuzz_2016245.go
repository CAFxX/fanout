package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint32; F1 int64; F2 bool }

func main() {
    var v1 bool = false
    var v2 int8 = 42
    var v3 float64 = 3.14159
    var v4 string = "hello"
    var v5 string = (((v4 + ((v4 + (v4 + ("" + "a"))) + v4)) + (v4 + v4)) + v4)
    var v6 int = 127
    v7 := []int{(v6 - 32768), v6}
    fmt.Println(v3)
    v2 += v2
    v1 = true
    v5 = v4
    v4 = v5
    v8 := GenS1{F0: 65535}
    v9 := GenS2{F0: 1, F1: 1, F2: (9223372036854775807 == 821)}
    for fi10 := 0; fi10 < 10; fi10++ {
        switch (v5 + v4) {
        case "":
            var v11 int16 = 1000
            _ = v11
        case "a":
            v7 = append(v7, (fi10 * (v6 / ((v6) | 1))))
        }
        _ = fi10
    }
    fmt.Println(v3)
    if v1 {
        v1 = (!false)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
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
