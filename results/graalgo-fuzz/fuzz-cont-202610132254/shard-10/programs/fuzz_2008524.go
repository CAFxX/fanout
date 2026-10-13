package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8 }
type GenS2 struct { F0 uint16 }

func main() {
    var v1 float32 = float32(-2.25)
    var v2 string = ""
    var v3 int = 32768
    var v4 uint = 947
    v5 := make(map[int]string)
    v6 := GenS1{F0: 0, F1: 42}
    if false {
        v7 := []int{}
        v8 := GenS1{F0: 255, F1: 42}
        _ = v7
        _ = v8
    } else {
        _ = len(v5)
        fmt.Println(v2)
        delete(v5, (v3 | (v3 ^ (v3 + 172))))
    }
    fmt.Println(v4, v3, v4)
    for fi9 := 0; fi9 < 3; fi9++ {
        fmt.Println(fi9, v1, v1)
        v1 = v1
        v2 = (v2 + (v2 + ((v2 + v2) + "\n\t")))
        _ = fi9
    }
    switch "" {
    case "hello":
        switch v3 {
        case 127:
            var v10 int8 = -1
            _ = v10
        case -1:
            var v11 int32 = 0
            _ = v11
        case -910:
            for fi12 := 0; fi12 < 16; fi12++ {
                switch 32768 {
                case 32767:
                    v4 += 698
                }
                _ = fi12
            }
        }
    default:
        if false {
            fmt.Println(v1)
            var v13 int8 = -1
            _ = v13
        }
    }
    fmt.Println(v3, v2, v1)
    if (42 > 42) {
        v1 += v1
        v14 := GenS2{F0: 42}
        _ = v14
    }
    v5[v3] = "hello"
    delete(v5, (v3 * v3))
    var v15 string = "a"
    var v16 int64 = -1
    var v17 uint8 = 42
    fmt.Println(v3)
    v1 = v1
    v18 := GenS2{F0: 65535}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
