package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool; F2 uint }

func main() {
    v1 := []int{-9223372036854775808, -1, 128}
    var v2 int32 = 42
    v3 := make(map[int]uint32)
    var v4 uint64 = 42
    var v5 int8 = -128
    var v6 float32 = 3.14159
    if len(v1) > 0 {
        v1[0] = 127
    }
    for fi7 := 0; fi7 < 5; fi7++ {
        var v8 int = 32767
        _ = fi7
        _ = v8
    }
    if (!true) {
        v6 += (v6 - v6)
    } else {
        var v9 uint32 = 65535
        _ = v9
    }
    v1 = append(v1, 0)
    switch "foo bar" {
    case "":
        if (true || (!false)) {
            delete(v3, 32767)
        }
    case "hello":
        v10 := GenS1{F0: 2147483647, F1: (2 < 126), F2: 2}
        _ = v10
    case "foo bar":
        fmt.Println(v6)
    default:
        switch 0 {
        case 127:
            var v11 float32 = v6
            _ = v11
        }
    }
    var v12 int64 = -9223372036854775808
    if false {
        for fi13 := 0; fi13 < 11; fi13++ {
            switch "" {
            case "":
                switch fi13 {
                case 922:
                    fmt.Println(v5, fi13, v2)
                case 0:
                    v1 = append(v1, 0)
                }
            }
            _ = fi13
        }
        if false {
            if true {
                var v14 string = "foo bar"
                var v15 uint64 = 42
                var v16 int32 = v2
                _ = v14
                _ = v15
                _ = v16
            } else {
                if false {
                    v17 := GenS1{F0: 2147483647, F1: (true || (-9223372036854775808 != -1)), F2: 42}
                    _ = v17
                }
            }
        }
        v12 = (v12 / (((v12 / ((v12) | 1))) | 1))
    }
    for fi18 := 0; fi18 < 5; fi18++ {
        var v19 uint8 = 254
        _ = fi18
        _ = v19
    }
    fmt.Println(v12, v2, v12)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
