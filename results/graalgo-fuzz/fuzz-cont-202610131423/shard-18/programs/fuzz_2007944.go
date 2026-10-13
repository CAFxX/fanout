package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16; F2 uint8 }
func genFn0(p0 int32, p1 int64) uint16 {
    var v5 float64 = -2.25
    v5 += v5
    fmt.Println(v5)
    _ = p0
    _ = p1
    _ = v5
    return 308
}
func genFn1(p0 uint32) uint16 {
    switch -2147483648 {
    case 469:
        p0 -= 0
    case 42:
        switch ("\n\t" + "\n\t") {
        case "foo bar":
            var v6 uint = 0
            _ = v6
        case "hello":
            p0 += (p0 & 405)
        case "a":
            var v7 int16 = 800
            _ = v7
        }
    default:
        var v8 int64 = -918
        _ = v8
    }
    v9 := []int{-463, -687, -9223372036854775808}
    _ = p0
    _ = v9
    return 0
}

func main() {
    var v1 int = -67
    v2 := []int{-9223372036854775808}
    var v3 uint16 = 182
    var v4 float64 = 3.14159
    var v10 int8 = 127
    var v11 uint16 = v3
    var v12 int8 = v10
    v13 := GenS1{F0: 437, F1: -32768, F2: 26}
    switch 32768 {
    case 4294967295:
        var v14 int16 = -1
        _ = v14
    }
    for fi15 := 0; fi15 < 5; fi15++ {
        v2 = append(v2, v1)
        _ = fi15
    }
    var v16 uint8 = 96
    v10 -= v12
    var v17 int32 = -1
    var v18 uint8 = uint8(v11)
    fmt.Println(v18, v11, v11)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
