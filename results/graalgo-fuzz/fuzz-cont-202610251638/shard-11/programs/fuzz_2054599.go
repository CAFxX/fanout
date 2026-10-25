package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }

func main() {
    var v1 uint8 = 255
    var v2 uint16 = 1
    var v3 string = ""
    var v4 float64 = 0.0
    switch 0 {
    case 2147483647:
        v5 := []int{}
        _ = v5
    case -1:
        v6 := GenS1{F0: 42}
        _ = v6
    case -2147483648:
        v2 = v2
    default:
        v7 := GenS1{F0: 42}
        _ = v7
    }
    var v8 int = 128
    v9 := []int{2147483647, (-(v8))}
    var v10 int64 = 685
    v9 = append(v9, int(v2))
    v2 += (v2 | (v2 >> ((4294967295) & 63)))
    var v11 float32 = (0.1 * -2.25)
    fmt.Println(v10, v4, v1)
    var v12 string = v3
    fmt.Println(v10, v4, v11)
    if len(v9) > 0 {
        v9[0] = 9223372036854775807
    }
    v13 := GenS1{F0: 127}
    if true {
        var v14 uint = 83
        var v15 uint = (v14 - v14)
        switch (v8 + 127) {
        case -1:
            if true {
                v16 := GenS1{F0: -1}
                var v17 int8 = 127
                _ = v16
                _ = v17
            } else {
                v18 := GenS1{F0: 1}
                fmt.Println(v8, v15)
                v9 = append(v9, v8)
                _ = v18
            }
        case 127:
            v3 = "\n\t"
        }
        _ = v14
        _ = v15
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
