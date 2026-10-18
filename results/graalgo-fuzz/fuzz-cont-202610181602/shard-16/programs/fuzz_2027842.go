package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint16; F1 bool; F2 uint8 }
func genFn0() string {
    if (!false) {
        var v6 float64 = (1e10 - (1.5 - 3.14159))
        _ = v6
    }
    v7 := []int{-824, 463, -895, 32768}
    v7 = append(v7, 637)
    var v8 int = 220
    _ = v7
    _ = v8
    return ("foo bar" + ("" + "a"))
}

func main() {
    v1 := make(map[string]string)
    var v2 uint = 255
    var v3 bool = false
    var v4 int32 = 327
    var v5 uint64 = uint64(v2)
    var v9 int8 = 37
    fmt.Println(v2, v2)
    v10 := v1["a"]
    v11 := GenS1{F0: -32768}
    v12 := GenS1{F0: 1}
    fmt.Println(v5, v4, v4)
    if v3 {
        fmt.Println(v4, v9)
        fmt.Println(v10, v4)
        for fi13 := 0; fi13 < 3; fi13++ {
            for fi14 := 0; fi14 < 6; fi14++ {
                switch v10 {
                case "hello":
                    v1[v10] = (v10 + v10)
                }
                _ = fi14
            }
            _ = fi13
        }
    } else {
        fmt.Println(v4)
        if ((v10 + v10) >= (v10 + (v10 + "a"))) {
            v4 = (v4 / ((v4) | 1))
            delete(v1, "\n\t")
            v15 := GenS2{F0: 665, F1: true, F2: 44}
            _ = v15
        } else {
            for fi16 := 0; fi16 < 5; fi16++ {
                _ = len(v1)
                switch v10 {
                case "hello":
                    var v17 int32 = v4
                    _ = v17
                }
                var v18 uint8 = 177
                _ = fi16
                _ = v18
            }
            var v19 int8 = -128
            var v20 int32 = v4
            _ = v19
            _ = v20
        }
    }
    v1[v10] = (v10 + "foo bar")
    var v21 uint16 = 2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
