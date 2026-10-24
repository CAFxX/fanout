package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16; F2 uint16 }

func main() {
    v1 := []int{}
    var v2 int = 0
    v3 := make(map[string]int)
    var v4 int = 128
    for fi5 := 0; fi5 < 10; fi5++ {
        fmt.Println(fi5, v2, v4)
        for fi6 := 0; fi6 < 10; fi6++ {
            fmt.Println(v4)
            _ = fi6
        }
        _ = fi5
    }
    for fi7 := 0; fi7 < 19; fi7++ {
        for fi8 := 0; fi8 < 11; fi8++ {
            switch v2 {
            case 9223372036854775807:
                delete(v3, (("foo bar" + "") + "hello"))
            case 42:
                v4 = fi7
            case 127:
                v2 = (v2 / (((v2 ^ fi7)) | 1))
            }
            v9 := GenS1{F0: 129, F1: -5, F2: 1}
            fmt.Println(v4)
            _ = fi8
            _ = v9
        }
        _ = fi7
    }
    delete(v3, "")
    var v10 float32 = 0.0
    var v11 int16 = 1000
    var v12 string = "hello"
    v12 = v12
    var v13 uint8 = 127
    v12 = v12
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
}
