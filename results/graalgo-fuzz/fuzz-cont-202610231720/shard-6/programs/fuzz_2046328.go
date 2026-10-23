package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }

func main() {
    v1 := make(map[uint32]int)
    var v2 uint32 = 65535
    var v3 bool = (!false)
    var v4 float32 = 1e10
    v5 := make(map[uint32]int)
    v6 := make(map[uint32]int)
    if v3 {
        var v7 int8 = 127
        v7 = v7
        _ = v7
    } else {
        switch "\n\t" {
        case "a":
            switch ("" + ("" + "\n\t")) {
            case "foo bar":
                v2 = (v2 | v2)
            case "hello":
                if v3 {
                    var v8 float64 = 3.14159
                    v8 += v8
                    var v9 int = 295
                    _ = v8
                    _ = v9
                }
            }
        }
        fmt.Println(v3, v4)
    }
    var v10 bool = (1 < v2)
    var v11 uint64 = 4294967295
    v12 := []int{}
    var v13 uint8 = 127
    switch ("\n\t" + "hello") {
    case "foo bar":
        v12 = append(v12, 32768)
    case "\n\t":
        var v14 int32 = 42
        _ = v14
    case "a":
        v10 = v3
    default:
        if len(v12) > 0 {
            v12[0] = 32768
        }
    }
    v3 = (-9223372036854775808 <= 0)
    if v10 {
        if v10 {
            if len(v12) > 0 {
                v12[0] = 2147483647
            }
            v2 = 801
        } else {
            switch 42 {
            case 127:
                var v15 uint16 = 65535
                _ = v15
            case -940:
                if (255 == 1) {
                    v16 := GenS1{F0: -2147483648}
                    _ = v16
                } else {
                    var v17 uint64 = (v11 >> ((924) & 63))
                    _ = v17
                }
            default:
                var v18 uint8 = (^(v13))
                _ = v18
            }
        }
        v3 = v3
    }
    _ = len(v6)
    _ = len(v12)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
