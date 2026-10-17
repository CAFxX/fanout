package main

import (
    "fmt"
)

func genFn0(p0 bool, p1 uint) int32 {
    v4 := []int{2147483647, 127, -889}
    p0 = p0
    var v5 int16 = 402
    v5 = (v5 - -122)
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return -714
}
func genFn1() int32 {
    for fi6 := 0; fi6 < 18; fi6++ {
        var v7 uint32 = 2147483647
        var v8 int16 = 148
        _ = fi6
        _ = v7
        _ = v8
    }
    if true {
        switch 42 {
        case 128:
            v9 := []int{-2147483648, 0}
            _ = v9
        }
        var v10 uint = 1
        _ = v10
    }
    var v11 uint8 = 128
    v11 += (^(v11))
    _ = v11
    return 485
}

func main() {
    var v1 uint64 = 0
    var v2 string = ("foo bar" + "a")
    var v3 uint32 = 150
    var v12 uint8 = 1
    switch -775 {
    case 128:
        var v13 float32 = 1e10
        _ = v13
    default:
        var v14 string = (v2 + v2)
        _ = v14
    }
    v3 -= (v3 ^ v3)
    fmt.Println(v1, v12)
    var v15 uint16 = 1
    v16 := []int{-9223372036854775808, -1}
    var v17 uint8 = v12
    var v18 string = "\n\t"
    if len(v16) >= 2 {
        v16 = v16[1:len(v16)]
    }
    var v19 float64 = 1e10
    fmt.Println(v1, v1)
    fmt.Println(v17, v15)
    if (1 >= 1) {
        v12 += (v17 ^ v12)
        v1 -= (v1 * v1)
    }
    v3 = 65535
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
}
