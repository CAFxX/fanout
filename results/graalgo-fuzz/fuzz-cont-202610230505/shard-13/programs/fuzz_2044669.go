package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64 }
type GenS2 struct { F0 uint8 }

func main() {
    var v1 string = "foo bar"
    var v2 int = 128
    var v3 string = v1
    var v4 int16 = 112
    switch 823 {
    case 127:
        fmt.Println(v3, v3)
    default:
        if true {
            v2 = (^(v2))
            v5 := []int{}
            fmt.Println(v1, v2)
            _ = v5
        }
    }
    var v6 uint64 = 1
    v7 := []int{v2, (-(v2)), v2, v2}
    _ = len(v7)
    v8 := GenS1{F0: 42, F1: 9223372036854775807}
    switch "a" {
    case "a":
        v1 = "a"
    default:
        for fi9 := 0; fi9 < 5; fi9++ {
            var v10 string = "\n\t"
            var v11 float64 = 0.1
            _ = fi9
            _ = v10
            _ = v11
        }
    }
    v3 = ((v1 + "") + v3)
    var v12 uint = 0
    _ = len(v7)
    v13 := GenS2{F0: 255}
    v6 = v6
    v3 = (v3 + v3)
    var v14 int16 = (-(v4))
    if true {
        if false {
            for fi15 := 0; fi15 < 9; fi15++ {
                for fi16 := 0; fi16 < 10; fi16++ {
                    v14 -= v14
                    var v17 float32 = -2.25
                    switch v3 {
                    case "":
                        var v18 uint8 = 191
                        _ = v18
                    case "hello":
                        var v19 uint32 = 42
                        _ = v19
                    case "foo bar":
                        if len(v7) >= 2 {
                            v7 = v7[1:len(v7)]
                        }
                    default:
                        v20 := GenS1{F0: 4294967295, F1: 546}
                        _ = v20
                    }
                    _ = fi16
                    _ = v17
                }
                v21 := GenS2{F0: 128}
                _ = fi15
                _ = v21
            }
        }
        var v22 int = v2
        var v23 int8 = 0
        _ = v22
        _ = v23
    }
    var v24 uint8 = 255
    for fi25 := 0; fi25 < 2; fi25++ {
        if false {
            v7 = append(v7, int(v12))
            var v26 string = "foo bar"
            _ = v26
        }
        _ = len(v7)
        _ = fi25
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v8
}
