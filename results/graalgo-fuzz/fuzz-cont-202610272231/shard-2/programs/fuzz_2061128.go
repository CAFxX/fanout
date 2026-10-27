package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string; F2 int64 }

func main() {
    var v1 uint8 = 255
    var v2 string = "\n\t"
    var v3 uint16 = 255
    v4 := []uint32{}
    if false {
        if true {
            var v5 uint = 700
            _ = v5
        }
        v6 := GenS1{F0: 87, F1: v2, F2: -9223372036854775808}
        _ = v6
    } else {
        var v7 float32 = (3.14159 + -2.25)
        if true {
            fmt.Println(v3, v2, v2)
            var v8 string = "a"
            _ = v8
        }
        switch int(v3) {
        case 9223372036854775807:
            switch "foo bar" {
            case "":
                fmt.Println(v3)
            default:
                var v9 int = 4294967295
                _ = v9
            }
        }
        _ = v7
    }
    for fi10 := 0; fi10 < 9; fi10++ {
        v3 += v3
        _ = fi10
    }
    var v11 bool = (0 != 0)
    var v12 uint16 = 444
    v13 := GenS1{F0: 0, F1: v2, F2: 9223372036854775807}
    v14 := GenS1{F0: -1, F1: "", F2: -9223372036854775808}
    v15 := GenS1{F0: 127, F1: v2, F2: 0}
    v4 = append(v4, 42)
    v1 += v1
    if true {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        switch -158 {
        case 127:
            var v16 int16 = 1
            _ = v16
        case 128:
            var v17 uint16 = (^(v12))
            _ = v17
        default:
            v11 = false
        }
    } else {
        fmt.Println(v2)
        v1 = 149
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
}
