package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64; F2 int32 }
type GenS2 struct { F0 uint64; F1 uint32; F2 int32 }

func main() {
    var v1 bool = true
    var v2 string = "a"
    v3 := make(map[uint32]bool)
    v4 := make(map[uint32]bool)
    v5 := make(map[string]string)
    var v6 bool = (2147483648 != 4294967295)
    v7 := GenS1{F0: 11, F1: 1, F2: 0}
    v8 := GenS1{F0: 2, F1: 2, F2: 42}
    delete(v3, 42)
    v9 := GenS1{F0: 65535, F1: 1, F2: 0}
    if true {
        if v1 {
            v10 := GenS2{F0: 1, F1: 4294967295, F2: 42}
            _ = v10
        } else {
            fmt.Println(v2)
            switch 32768 {
            case 42:
                var v11 uint64 = 1
                _ = v11
            case 780:
                var v12 uint = 255
                _ = v12
            case 2147483647:
                var v13 float64 = 0.0
                _ = v13
            }
        }
        if false {
            if v1 {
                var v14 float32 = float32(0.1)
                _ = v14
            }
        }
        _ = len(v3)
    } else {
        var v15 string = "hello"
        v6 = true
        _ = v15
    }
    var v16 float64 = 0.0
    v17 := GenS1{F0: 2147483648, F1: 729, F2: -493}
    fmt.Println(v6)
    switch "\n\t" {
    case "a":
        v4[2147483648] = ((100 != 42) && v1)
    case "\n\t":
        v18 := []int{-330, 0, 4294967295, 32767}
        _ = v18
    }
    if v1 {
        for fi19 := 0; fi19 < 8; fi19++ {
            v1 = (42 >= 1)
            _ = fi19
        }
        fmt.Println(v16, v16, v2)
    } else {
        var v20 int8 = 77
        _ = v20
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
