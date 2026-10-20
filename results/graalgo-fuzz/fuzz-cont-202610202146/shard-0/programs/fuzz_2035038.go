package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 bool; F1 int }
func genFn0() string {
    var v5 string = ""
    v6 := GenS1{F0: 0}
    v5 = ""
    _ = v5
    _ = v6
    return ((v5 + "hello") + (v5 + ("hello" + ("hello" + v5))))
}

func main() {
    var v1 bool = (128 == 2)
    var v2 uint32 = 1
    v3 := make(map[int]bool)
    v4 := make(map[string]bool)
    if true {
        for fi7 := 0; fi7 < 16; fi7++ {
            var v8 int32 = -237
            fmt.Println(fi7)
            _ = fi7
            _ = v8
        }
        fmt.Println(v2, v2)
        fmt.Println(v2)
    } else {
        v2 += (^(v2))
        v1 = true
        var v9 uint16 = 0
        _ = v9
    }
    v10 := GenS1{F0: 615}
    v1 = v1
    var v11 uint = 0
    var v12 float64 = (-2.25 - -2.25)
    v13 := GenS2{F0: (!((1.5 / (v12 * (v12 / v12))) <= (1e10 * float64(1e10)))), F1: -2147483648}
    switch "hello" {
    case "a":
        var v14 uint8 = 255
        _ = v14
    case "":
        v1 = v1
    default:
        v15 := GenS1{F0: v11}
        _ = v15
    }
    var v16 int16 = -528
    v12 += 3.14159
    delete(v4, (("a" + "hello") + (((("a" + ("a" + "hello")) + ("\n\t" + "")) + "") + "")))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
