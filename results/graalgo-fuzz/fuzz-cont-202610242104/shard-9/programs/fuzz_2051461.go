package main

import (
    "fmt"
)

type GenS1 struct { F0 int }

func main() {
    v1 := []int{}
    var v2 string = "hello"
    var v3 float64 = 0.0
    var v4 float64 = ((0.0 - v3) * (v3 + -2.25))
    v5 := []int{127}
    var v6 int16 = 42
    switch v2 {
    case "":
        if (1 < 42) {
            for fi7 := 0; fi7 < 14; fi7++ {
                var v8 uint16 = 0
                var v9 uint16 = 1
                _ = fi7
                _ = v8
                _ = v9
            }
            for fi10 := 0; fi10 < 5; fi10++ {
                var v11 int64 = -9223372036854775808
                fmt.Println(v6, fi10, v3)
                _ = fi10
                _ = v11
            }
        }
    default:
        if true {
            var v12 uint8 = 28
            for fi13 := 0; fi13 < 20; fi13++ {
                v14 := GenS1{F0: fi13}
                _ = len(v1)
                _ = fi13
                _ = v14
            }
            _ = v12
        }
    }
    var v15 uint32 = 42
    for fi16 := 0; fi16 < 8; fi16++ {
        var v17 uint8 = 127
        var v18 string = (("a" + ("a" + v2)) + v2)
        var v19 uint8 = uint8(v15)
        _ = fi16
        _ = v17
        _ = v18
        _ = v19
    }
    fmt.Println(v6, v15, v15)
    var v20 uint8 = 237
    var v21 string = ""
    fmt.Println(v2)
    var v22 int8 = 127
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
