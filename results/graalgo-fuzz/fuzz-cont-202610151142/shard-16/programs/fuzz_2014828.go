package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string; F2 string }
type GenS2 struct { F0 uint8; F1 uint16; F2 uint32 }

func main() {
    var v1 uint64 = 2
    v2 := make(map[string]bool)
    v3 := make(map[uint32]int)
    v4 := make(map[int]int)
    var v5 int32 = -382
    v6 := GenS1{F0: 2, F1: "a", F2: ("hello" + "hello")}
    switch ("hello" + "\n\t") {
    case "a":
        var v7 int64 = 42
        _ = v7
    default:
        var v8 uint64 = (v1 + v1)
        _ = v8
    }
    var v9 float64 = ((1.5 / (1.5 - 0.1)) - 3.14159)
    if false {
        v10 := v2["foo bar"]
        switch -2147483648 {
        case 2147483647:
            fmt.Println(v1)
        default:
            v11 := GenS2{F0: 177, F1: 255, F2: 0}
            _ = v11
        }
        var v12 bool = ("a" >= "")
        _ = v10
        _ = v12
    }
    v13 := GenS1{F0: 2, F1: "hello", F2: ""}
    if (9223372036854775807 == 4294967295) {
        var v14 uint32 = 4294967295
        _ = v14
    }
    if true {
        for fi15 := 0; fi15 < 10; fi15++ {
            v16 := []int{fi15}
            _ = fi15
            _ = v16
        }
        var v17 int64 = -9223372036854775808
        _ = v17
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
