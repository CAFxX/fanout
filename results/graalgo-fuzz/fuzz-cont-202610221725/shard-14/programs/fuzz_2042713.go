package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64; F2 uint16 }
func genFn0(p0 bool) uint8 {
    if (100 == 42) {
        fmt.Println(p0)
        var v7 int32 = 42
        switch "" {
        case "foo bar":
            var v8 bool = (2 == 192)
            _ = v8
        case "hello":
            p0 = ((552 != 18446744073709551615) && (!true))
        default:
            var v9 uint32 = 0
            _ = v9
        }
        _ = v7
    } else {
        fmt.Println(p0)
    }
    v10 := []int{-910}
    var v11 uint32 = 75
    _ = len(v10)
    p0 = (556 > 337)
    _ = p0
    _ = v10
    _ = v11
    return 127
}

func main() {
    var v1 uint = 42
    var v2 uint32 = 93
    v3 := make(map[uint32]string)
    var v4 bool = (v1 > 0)
    v5 := []string{(("" + "foo bar") + ("" + "a"))}
    var v6 string = ""
    var v12 uint32 = v2
    v6 = "foo bar"
    var v13 int8 = 100
    v14 := GenS1{F0: 728, F1: 351, F2: 421}
    for fi15 := 0; fi15 < 18; fi15++ {
        var v16 float32 = -2.25
        if len(v5) > 0 {
            v5[0] = v6
        }
        _ = len(v3)
        _ = fi15
        _ = v16
    }
    delete(v3, v2)
    switch 209 {
    case 951:
        var v17 int64 = 9223372036854775807
        _ = v17
    default:
        v5 = append(v5, ((("hello" + "a") + (v6 + v6)) + (v6 + v6)))
    }
    fmt.Println(v4)
    fmt.Println(v13)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
