package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int32; F1 int16; F2 uint8 }
type GenS3 struct { F0 string }

func main() {
    var v1 uint32 = 33
    var v2 uint8 = 1
    var v3 uint = 4294967295
    v4 := []int{0, 2147483647}
    if len(v4) > 0 {
        v4[0] = 975
    }
    switch "a" {
    case "hello":
        if true {
            var v5 int32 = int32(v3)
            _ = v5
        }
    case "":
        v6 := GenS1{F0: -447}
        _ = v6
    case "foo bar":
        fmt.Println(v2)
    }
    v2 -= 1
    fmt.Println(v1)
    if len(v4) > 0 {
        v4[0] = -54
    }
    for fi7 := 0; fi7 < 5; fi7++ {
        v8 := GenS2{F0: -580, F1: -32768, F2: 255}
        _ = fi7
        _ = v8
    }
    v9 := GenS3{F0: ("a" + "\n\t")}
    var v10 uint = 834
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
