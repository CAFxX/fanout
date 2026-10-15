package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }

func main() {
    var v1 int8 = -1
    var v2 float64 = 1e10
    v3 := make(map[string]int)
    var v4 int8 = (v1 | v1)
    var v5 uint = 42
    var v6 int32 = 2147483647
    v7 := []int{128, -2147483648, -9223372036854775808}
    var v8 bool = true
    v5 += v5
    for fi9 := 0; fi9 < 18; fi9++ {
        switch "a" {
        case "hello":
            v8 = (v8 && false)
        case "\n\t":
            var v10 uint8 = 255
            _ = v10
        }
        _ = fi9
    }
    fmt.Println(v8, v2, v8)
    v11 := GenS1{F0: 1}
    var v12 int = 2147483647
    v8 = ((v8 || v8) && true)
    if v8 {
        v12 -= 32767
        if v8 {
            var v13 int32 = v6
            switch 4294967295 {
            case 0:
                v14 := GenS1{F0: v13}
                _ = v14
            case 32768:
                v13 -= (v6 - (v13 % (((-(v13))) | 1)))
            case -9223372036854775808:
                v15 := GenS1{F0: v13}
                _ = v15
            }
            _ = v13
        } else {
            var v16 bool = v8
            v17 := GenS1{F0: (v6 % (((v6 * (v6 % ((v6) | 1)))) | 1))}
            fmt.Println(v8)
            _ = v16
            _ = v17
        }
    }
    for fi18 := 0; fi18 < 12; fi18++ {
        _ = len(v3)
        _ = fi18
    }
    if v8 {
        v3[""] = v12
    }
    var v19 uint16 = 927
    fmt.Println(v5, v12)
    var v20 string = "a"
    fmt.Println(v4, v19)
    v19 += v19
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
