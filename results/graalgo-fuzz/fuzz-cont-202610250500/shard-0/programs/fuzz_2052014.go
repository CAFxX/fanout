package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 bool; F1 int32 }

func main() {
    var v1 uint = 18446744073709551615
    v2 := make(map[string]string)
    var v3 string = ("" + "a")
    var v4 int32 = 0
    var v5 uint16 = 42
    v6 := GenS1{F0: v3}
    v7 := []int{42, -594, 0, -9223372036854775808}
    v8 := GenS2{F0: (v5 != 660), F1: 0}
    var v9 uint64 = 1
    if false {
        var v10 int32 = (v4 >> ((v1) & 63))
        _ = v10
    }
    fmt.Println(v5)
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    var v11 int = -2147483648
    for fi12 := 0; fi12 < 2; fi12++ {
        if (false && (-1 == 1000)) {
            var v13 uint = v1
            v2[(("hello" + "hello") + ("\n\t" + v3))] = v3
            _ = v13
        }
        var v14 int64 = -1
        for fi15 := 0; fi15 < 4; fi15++ {
            switch v3 {
            case "\n\t":
                if len(v7) > 0 {
                    v7[0] = 127
                }
            default:
                v11 = v11
            }
            _ = fi15
        }
        _ = fi12
        _ = v14
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
