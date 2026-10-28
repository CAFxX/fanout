package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32; F2 string }

func main() {
    var v1 int16 = 32767
    v2 := make(map[uint32]string)
    var v3 int8 = 127
    if false {
        var v4 string = "hello"
        v5 := GenS1{F0: 32767, F1: 942, F2: v4}
        _ = v4
        _ = v5
    } else {
        v1 = (^(v1))
        var v6 int = 32768
        _ = v6
    }
    v2[4294967295] = (("" + "\n\t") + "hello")
    v7 := GenS1{F0: 4294967295, F1: -1, F2: "foo bar"}
    var v8 uint16 = 42
    if true {
        switch (("hello" + ("" + "\n\t")) + "\n\t") {
        case "":
            switch 360 {
            case 32768:
                v3 += v3
            default:
                v9 := []int{9223372036854775807}
                _ = v9
            }
        default:
            fmt.Println(v8, v1, v1)
        }
        var v10 int64 = 0
        var v11 int64 = v10
        _ = v10
        _ = v11
    }
    v8 = v8
    var v12 uint64 = 4294967295
    var v13 bool = (v12 <= (v12 ^ v12))
    fmt.Println(v12, v13, v12)
    v14 := GenS1{F0: 1, F1: 1, F2: "a"}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v7
    _ = v8
}
