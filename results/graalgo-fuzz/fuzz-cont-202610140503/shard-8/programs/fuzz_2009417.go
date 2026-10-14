package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int32; F1 uint; F2 int16 }

func main() {
    v1 := make(map[string]int)
    var v2 uint32 = 4294967295
    var v3 int16 = 42
    var v4 uint16 = 2
    v5 := make(map[string]string)
    v6 := GenS1{F0: (v2 - v2)}
    var v7 float64 = 3.14159
    v5["a"] = ("foo bar" + "\n\t")
    v8 := GenS2{F0: 1, F1: uint(v4), F2: (v3 * 305)}
    for fi9 := 0; fi9 < 11; fi9++ {
        v4 = (v4 << ((0) & 63))
        if true {
            var v10 uint = 42
            fmt.Println(v7)
            _ = v10
        }
        _ = fi9
    }
    var v11 uint = 2
    var v12 int8 = int8(v4)
    v11 += (^(v11))
    if true {
        v13 := GenS1{F0: v2}
        _ = v13
    } else {
        fmt.Println(v3, v11)
        if true {
            switch ("foo bar" + "") {
            case "\n\t":
                var v14 string = ("hello" + "a")
                _ = v14
            }
            for fi15 := 0; fi15 < 10; fi15++ {
                var v16 bool = ((!false) || (1 <= 43))
                _ = fi15
                _ = v16
            }
        } else {
            v17 := []int{}
            v1["\n\t"] = 1
            _ = v17
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
