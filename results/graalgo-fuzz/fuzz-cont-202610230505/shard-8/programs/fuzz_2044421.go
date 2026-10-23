package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int; F2 uint }
type GenS2 struct { F0 string }

func main() {
    var v1 uint8 = 255
    var v2 int64 = 0
    v3 := []uint32{4294967295, 2147483648, 767, 1}
    v4 := []string{}
    var v5 uint16 = 338
    var v6 bool = false
    v7 := GenS1{F0: (v1 >> ((65535) & 63)), F1: 127, F2: 4294967295}
    var v8 string = ""
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi9 := 0; fi9 < 9; fi9++ {
        var v10 int32 = 774
        v11 := GenS2{F0: ((v8 + v8) + ((v8 + v8) + v8))}
        switch "a" {
        case "hello":
            fmt.Println(v5)
        default:
            fmt.Println(v8)
        }
        _ = fi9
        _ = v10
        _ = v11
    }
    if len(v3) > 0 {
        v3[0] = 0
    }
    if len(v4) > 0 {
        v4[0] = v8
    }
    if (((("" + ((v8 + ("a" + "foo bar")) + v8)) + ("hello" + v8)) + v8) > (("a" + v8) + ((v8 + ((v8 + ("" + (v8 + (v8 + v8)))) + v8)) + (v8 + v8)))) {
        switch -2147483648 {
        case 2147483647:
            if v6 {
                v1 += v1
                var v12 float64 = 0.1
                _ = v12
            }
        }
        fmt.Println(v8, v1, v8)
    }
    fmt.Println(v6)
    v5 += (^(v5))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
