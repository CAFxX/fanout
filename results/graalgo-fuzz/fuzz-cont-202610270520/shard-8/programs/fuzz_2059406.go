package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string; F2 bool }

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[uint32]bool)
    var v3 uint32 = 1
    v4 := []string{("foo bar" + "hello"), ("a" + "hello")}
    v3 -= v3
    fmt.Println(v3)
    switch ("a" + ("a" + "a")) {
    case "a":
        fmt.Println(v3)
    case "hello":
        var v5 uint16 = 97
        _ = v5
    case "":
        if true {
            if true {
                v3 += v3
            } else {
                if false {
                    var v6 bool = (true && true)
                    switch "foo bar" {
                    case "hello":
                        var v7 float64 = (float64(((float32((1e10 - 1.5)) / 1.5) * ((-2.25 - 1.5) - (float32(3.14159) + 0.0)))) / 1.0)
                        _ = v7
                    case "\n\t":
                        fmt.Println(v3, v3)
                    }
                    _ = v6
                } else {
                    var v8 bool = true
                    _ = v8
                }
                var v9 uint32 = v3
                if (2147483647 <= -2147483648) {
                    v3 = (v3 ^ v9)
                    if len(v4) > 0 {
                        v4[0] = (("" + (((("\n\t" + "foo bar") + "a") + ("foo bar" + "foo bar")) + ("hello" + "\n\t"))) + "hello")
                    }
                    var v10 uint16 = 0
                    _ = v10
                } else {
                    if (!("\n\t" == "a")) {
                        v3 -= v3
                        fmt.Println(v9)
                        var v11 uint8 = 172
                        _ = v11
                    }
                }
                _ = v9
            }
        }
    default:
        v12 := GenS1{F0: -32768, F1: "a", F2: (32768 != 42)}
        _ = v12
    }
    fmt.Println(v3)
    for fi13 := 0; fi13 < 19; fi13++ {
        v14 := GenS1{F0: -32768, F1: "foo bar", F2: (0 == 1)}
        _ = fi13
        _ = v14
    }
    var v15 float64 = 1e10
    v3 -= 4294967295
    for fi16 := 0; fi16 < 3; fi16++ {
        if false {
            v4 = append(v4, "")
        }
        var v17 uint8 = 187
        _ = fi16
        _ = v17
    }
    var v18 int8 = 42
    var v19 uint8 = 255
    v18 += (v18 & -7)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
