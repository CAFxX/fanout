package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int }

func main() {
    var v1 uint8 = 2
    var v2 bool = ("\n\t" != (("foo bar" + ("a" + ("foo bar" + ""))) + ""))
    var v3 uint64 = 1
    v4 := make(map[string]string)
    var v5 int8 = 42
    v6 := GenS1{F0: 658, F1: -322}
    var v7 int32 = 42
    var v8 uint64 = 4294967295
    fmt.Println(v5)
    var v9 uint = 4294967295
    var v10 uint64 = (^(v3))
    fmt.Println(v9, v7)
    v9 = (v9 | (v9 / ((4294967295) | 1)))
    v8 += 1
    switch -2147483648 {
    case 32767:
        if v2 {
            v11 := []int{32767}
            switch "hello" {
            case "foo bar":
                var v12 bool = v2
                _ = v12
            case "\n\t":
                for fi13 := 0; fi13 < 11; fi13++ {
                    v3 = (v8 | 2)
                    for fi14 := 0; fi14 < 1; fi14++ {
                        var v15 float64 = 1.5
                        var v16 float64 = v15
                        var v17 float32 = 1.5
                        _ = fi14
                        _ = v15
                        _ = v16
                        _ = v17
                    }
                    v1 = (v1 / ((v1) | 1))
                    _ = fi13
                }
            case "hello":
                var v18 int16 = int16(v9)
                _ = v18
            }
            _ = v11
        } else {
            var v19 int8 = (-(v5))
            v10 += 42
            v2 = (-32768 <= -1)
            _ = v19
        }
    default:
        v1 += (^(v1))
    }
    _ = len(v4)
    v5 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
