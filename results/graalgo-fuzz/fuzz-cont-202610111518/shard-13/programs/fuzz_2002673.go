package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32 }
type GenS2 struct { F0 int64 }

func main() {
    v1 := make(map[uint32]string)
    var v2 float32 = (1e10 * ((float32(1.5) * 1e10) * 3.14159))
    var v3 uint32 = 613
    v4 := make(map[int]string)
    var v5 uint8 = 255
    v6 := []int{32768}
    for fi7 := 0; fi7 < 2; fi7++ {
        switch "foo bar" {
        case "hello":
            var v8 int = fi7
            _ = v8
        case "\n\t":
            fmt.Println(v2, fi7)
        case "foo bar":
            var v9 uint16 = 65535
            _ = v9
        default:
            if (((992 == 1) && true) || true) {
                v3 += v3
            }
        }
        v5 = 2
        fmt.Println(v2, v5, v5)
        _ = fi7
    }
    v5 += v5
    if false {
        v6 = append(v6, 9223372036854775807)
        switch 291 {
        case 42:
            fmt.Println(v5, v3, v5)
        case -829:
            v3 = v3
        default:
            var v10 float32 = -2.25
            _ = v10
        }
        var v11 uint8 = (v5 % (((v5 << ((434) & 63))) | 1))
        _ = v11
    }
    v12 := GenS1{F0: 32768, F1: -1}
    switch 325 {
    case -9223372036854775808:
        for fi13 := 0; fi13 < 10; fi13++ {
            fmt.Println(fi13, fi13, v5)
            switch ("foo bar" + "a") {
            case "hello":
                v14 := GenS2{F0: 558}
                _ = v14
            case "":
                v5 = v5
            case "foo bar":
                var v15 uint16 = 255
                _ = v15
            }
            _ = fi13
        }
    }
    fmt.Println(v2)
    v3 = v3
    var v16 uint8 = v5
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v12
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
