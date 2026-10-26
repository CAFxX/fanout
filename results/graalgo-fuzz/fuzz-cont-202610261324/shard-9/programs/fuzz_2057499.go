package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int }
type GenS2 struct { F0 bool }

func main() {
    var v1 uint8 = 127
    var v2 float32 = float32(3.14159)
    var v3 bool = false
    v4 := []uint32{65535, 2}
    v5 := []string{("\n\t" + "a"), (("" + (("foo bar" + "\n\t") + "\n\t")) + "hello")}
    v6 := []int{2147483647, 32767, -2147483648}
    if v3 {
        fmt.Println(v3)
        fmt.Println(v1, v1, v2)
    } else {
        var v7 int8 = -128
        switch ("hello" + ("\n\t" + "foo bar")) {
        case "\n\t":
            var v8 int32 = 86
            _ = v8
        case "":
            v1 += 2
        case "a":
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
        default:
            var v9 uint = 7
            _ = v9
        }
        _ = v7
    }
    v1 += (v1 / (((v1 + 0)) | 1))
    var v10 int32 = -394
    for fi11 := 0; fi11 < 7; fi11++ {
        fmt.Println(v1, v3, v2)
        _ = fi11
    }
    v12 := GenS1{F0: (v3 && false), F1: -1}
    var v13 uint8 = v1
    v13 = (v13 | 255)
    v14 := GenS2{F0: true}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
