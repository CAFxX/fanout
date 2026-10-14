package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32; F2 uint16 }

func main() {
    var v1 int = 68
    var v2 uint16 = 2
    var v3 string = "foo bar"
    var v4 int32 = 418
    var v5 uint64 = 1
    var v6 string = v3
    var v7 int64 = 0
    v8 := []int{}
    v9 := GenS1{F0: 255, F1: v4, F2: v2}
    if false {
        for fi10 := 0; fi10 < 3; fi10++ {
            var v11 int16 = 909
            var v12 float32 = 1.5
            _ = fi10
            _ = v11
            _ = v12
        }
        if (127 > 127) {
            var v13 int16 = -1
            _ = v13
        }
    } else {
        v3 = v3
    }
    if len(v8) >= 2 {
        v8 = v8[1:len(v8)]
    }
    if false {
        if len(v8) > 0 {
            v8[0] = (v1 + v1)
        }
    }
    var v14 uint16 = 255
    if true {
        switch 988 {
        case -9223372036854775808:
            var v15 uint16 = v14
            _ = v15
        }
    } else {
        switch v6 {
        case "a":
            v6 = ("hello" + v3)
        case "hello":
            var v16 string = ""
            _ = v16
        default:
            v2 = v2
        }
        v2 = (v14 / ((9) | 1))
        var v17 uint32 = 2147483647
        _ = v17
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
