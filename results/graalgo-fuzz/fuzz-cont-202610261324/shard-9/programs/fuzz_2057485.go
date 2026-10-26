package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8; F2 int32 }
type GenS2 struct { F0 int32; F1 uint64 }
func genFn0() uint {
    if (621 >= 0) {
        var v5 uint = 895
        fmt.Println(v5)
        switch ("hello" + ("hello" + "foo bar")) {
        case "a":
            fmt.Println(v5)
        case "":
            v5 = v5
        default:
            if true {
                fmt.Println(v5)
            } else {
                switch 4294967295 {
                case 298:
                    v5 += (v5 + (v5 | v5))
                default:
                    v6 := []int{1, 0, -908, 2147483647}
                    _ = v6
                }
            }
        }
        _ = v5
    }
    var v7 string = "foo bar"
    v8 := []int{}
    fmt.Println(v7)
    _ = v7
    _ = v8
    return 42
}

func main() {
    var v1 int = 599
    var v2 bool = (!false)
    var v3 uint16 = 1
    v4 := make(map[string]string)
    var v9 uint = 2
    var v10 uint16 = v3
    switch ((("" + "\n\t") + "foo bar") + ("a" + ("\n\t" + "a"))) {
    case "a":
        var v11 float32 = -2.25
        _ = v11
    case "":
        var v12 uint32 = 432
        _ = v12
    case "hello":
        fmt.Println(v1, v2, v2)
    default:
        v13 := GenS1{F0: (v1 * v1), F1: -1, F2: -1}
        _ = v13
    }
    v14 := GenS2{F0: 1, F1: 314}
    var v15 int32 = 1
    fmt.Println(v9)
    var v16 uint64 = 478
    var v17 int32 = (v15 - int32(v3))
    fmt.Println(v10, v17, v3)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
