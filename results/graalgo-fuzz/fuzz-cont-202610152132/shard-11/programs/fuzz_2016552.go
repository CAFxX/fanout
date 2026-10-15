package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }

func main() {
    var v1 float64 = 0.1
    var v2 string = ("foo bar" + ("a" + "hello"))
    var v3 uint = 65535
    var v4 float64 = 0.1
    v1 = (v4 + (v1 * v4))
    v5 := GenS1{F0: 42}
    switch v2 {
    case "hello":
        v2 = ("a" + "")
    case "foo bar":
        v6 := []int{}
        _ = v6
    default:
        v7 := GenS1{F0: 1}
        _ = v7
    }
    if false {
        v8 := GenS1{F0: 100}
        _ = v8
    }
    v3 = 42
    var v9 int64 = -1
    v2 = ((v2 + "") + v2)
    v9 = (v9 >> (((v3 >> ((2) & 63))) & 63))
    switch 32767 {
    case 775:
        for fi10 := 0; fi10 < 13; fi10++ {
            var v11 int32 = -1
            _ = fi10
            _ = v11
        }
    }
    var v12 uint32 = 46
    if (false || (true && (81 >= -128))) {
        v13 := []int{}
        v3 = 1
        var v14 float64 = v4
        _ = v13
        _ = v14
    }
    v3 = v3
    if (true && (!(4294967295 < 133))) {
        v15 := GenS1{F0: -112}
        var v16 int64 = (v9 / ((v9) | 1))
        v9 = v16
        _ = v15
        _ = v16
    } else {
        fmt.Println(v2, v4)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
