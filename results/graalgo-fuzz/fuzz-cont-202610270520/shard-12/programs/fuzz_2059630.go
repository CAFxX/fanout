package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32; F2 int64 }

func main() {
    var v1 int32 = 1
    v2 := []string{"foo bar"}
    v3 := make(map[uint32]string)
    var v4 int16 = -662
    v5 := GenS1{F0: 1, F1: (v1 % ((v1) | 1)), F2: -9223372036854775808}
    v6 := GenS1{F0: 0, F1: -870, F2: -880}
    if (1 == 1) {
        if false {
            fmt.Println(v4, v4)
            var v7 int8 = -1
            _ = v7
        }
    }
    _ = len(v3)
    v4 += v4
    switch -881 {
    case 2147483647:
        var v8 string = "a"
        _ = v8
    case 32767:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    case -46:
        var v9 int32 = 982
        _ = v9
    default:
        v1 = v1
    }
    var v10 int8 = -44
    var v11 int = -2147483648
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
