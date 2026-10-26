package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32 }

func main() {
    var v1 uint64 = 0
    var v2 string = "\n\t"
    var v3 string = v2
    var v4 uint8 = 127
    v5 := make(map[int]string)
    v3 = v2
    fmt.Println(v3, v4)
    var v6 uint64 = (v1 - 4294967295)
    v7 := GenS1{F0: 42, F1: -678}
    v2 = ("" + "hello")
    if false {
        if false {
            fmt.Println(v6)
            v6 -= v1
            v8 := GenS1{F0: 118, F1: -456}
            _ = v8
        } else {
            for fi9 := 0; fi9 < 14; fi9++ {
                v1 = v1
                _ = fi9
            }
            var v10 string = ""
            v2 = "foo bar"
            _ = v10
        }
        fmt.Println(v6, v2)
    }
    var v11 int16 = 1
    v1 -= 42
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
