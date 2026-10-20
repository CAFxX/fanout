package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8 }
func genFn0() int16 {
    var v5 int8 = 0
    fmt.Println(v5)
    switch "hello" {
    case "":
        if true {
            v6 := []int{-696, 0}
            var v7 uint = 65535
            v5 -= (v5 / ((-1) | 1))
            _ = v6
            _ = v7
        }
    case "foo bar":
        v8 := []int{225}
        _ = v8
    case "\n\t":
        var v9 int64 = 1
        _ = v9
    default:
        var v10 float32 = (-2.25 / 1.5)
        _ = v10
    }
    var v11 float32 = 1e10
    var v12 uint64 = 307
    _ = v11
    _ = v12
    _ = v5
    return -333
}

func main() {
    var v1 int32 = -174
    var v2 uint = 427
    v3 := make(map[string]int)
    var v4 bool = true
    var v13 float64 = 3.14159
    var v14 bool = false
    v15 := []int{32768}
    switch "\n\t" {
    case "a":
        v16 := GenS1{F0: v2, F1: 128}
        _ = v16
    case "hello":
        v1 = v1
    }
    if len(v15) > 0 {
        v15[0] = int(v2)
    }
    for fi17 := 0; fi17 < 17; fi17++ {
        fmt.Println(v4, v1)
        v14 = v4
        _ = fi17
    }
    fmt.Println(v14)
    var v18 int8 = -91
    var v19 float32 = -2.25
    switch ("hello" + "foo bar") {
    case "foo bar":
        v4 = (1 <= 65535)
    case "\n\t":
        if ((v13 >= v13) || false) {
            v20 := GenS1{F0: 4294967295, F1: 128}
            switch "" {
            case "a":
                if v14 {
                    switch 32767 {
                    case 2147483647:
                        v21 := GenS1{F0: v2, F1: 0}
                        _ = v21
                    case 549:
                        fmt.Println(v19, v14)
                    default:
                        fmt.Println(v4, v18, v14)
                    }
                    var v22 float32 = 0.0
                    _ = v22
                } else {
                    if false {
                        v23 := GenS1{F0: v2, F1: 128}
                        v15 = append(v15, 4294967295)
                        _ = v23
                    }
                    v24 := GenS1{F0: v2, F1: 128}
                    v3["foo bar"] = 4294967295
                    _ = v24
                }
            case "\n\t":
                v18 -= v18
            }
            var v25 uint32 = 2147483647
            _ = v20
            _ = v25
        }
    }
    if len(v15) > 0 {
        v15[0] = 42
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
