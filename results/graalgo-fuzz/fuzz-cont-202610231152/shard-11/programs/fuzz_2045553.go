package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }

func main() {
    v1 := make(map[uint32]string)
    var v2 int32 = -1
    var v3 uint32 = 235
    var v4 bool = false
    var v5 uint = 512
    var v6 uint = v5
    v7 := []int{1, -2147483648}
    v8 := GenS1{F0: 18446744073709551615}
    switch 455 {
    case 32767:
        fmt.Println(v4, v2, v5)
    }
    var v9 int32 = (^(v2))
    v2 -= -1
    if false {
        for fi10 := 0; fi10 < 10; fi10++ {
            switch fi10 {
            case -1:
                if false {
                    var v11 uint8 = 1
                    v5 += 699
                    var v12 int8 = 42
                    _ = v11
                    _ = v12
                }
            case -181:
                var v13 int8 = 42
                _ = v13
            default:
                delete(v1, 1)
            }
            v6 = (v6 << ((v5) & 63))
            v6 += 42
            _ = fi10
        }
    }
    switch "" {
    case "hello":
        switch "a" {
        case "\n\t":
            var v14 int = -123
            _ = v14
        case "":
            v7 = append(v7, 128)
        case "a":
            var v15 uint64 = 2
            _ = v15
        default:
            var v16 int32 = 1
            _ = v16
        }
    case "a":
        v17 := GenS1{F0: uint64(v6)}
        _ = v17
    default:
        v2 = (v9 % ((1) | 1))
    }
    if true {
        for fi18 := 0; fi18 < 9; fi18++ {
            v4 = false
            for fi19 := 0; fi19 < 13; fi19++ {
                v5 = 1
                _ = fi19
            }
            _ = fi18
        }
    }
    var v20 int8 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
