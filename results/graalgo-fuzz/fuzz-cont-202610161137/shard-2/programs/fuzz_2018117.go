package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16 }

func main() {
    var v1 uint8 = 106
    var v2 int32 = 42
    var v3 float32 = 1.5
    var v4 int32 = (^(v2))
    if false {
        v2 -= -1
        switch -496 {
        case 42:
            fmt.Println(v4)
        case 0:
            v5 := []int{42}
            _ = v5
        case 2147483647:
            v4 = v2
        }
    }
    v6 := GenS1{F0: ("\n\t" + "hello"), F1: 0}
    fmt.Println(v4, v2, v3)
    for fi7 := 0; fi7 < 3; fi7++ {
        var v8 int8 = -1
        fmt.Println(v2, v2, v1)
        _ = fi7
        _ = v8
    }
    if false {
        var v9 int64 = 0
        v1 += (v1 >> ((65535) & 63))
        _ = v9
    }
    for fi10 := 0; fi10 < 2; fi10++ {
        v4 = v2
        v11 := []int{}
        _ = fi10
        _ = v11
    }
    var v12 uint16 = 255
    if false {
        var v13 int32 = 2147483647
        _ = v13
    } else {
        v4 = v2
    }
    for fi14 := 0; fi14 < 4; fi14++ {
        v15 := []int{-448, fi14, 9223372036854775807, (-(fi14))}
        var v16 int32 = (v2 + -969)
        _ = fi14
        _ = v15
        _ = v16
    }
    v2 += (v2 - v4)
    if (42 < -76) {
        for fi17 := 0; fi17 < 14; fi17++ {
            v18 := []int{(fi17 >> ((2) & 63)), 0}
            _ = fi17
            _ = v18
        }
    }
    fmt.Println(v1, v1, v1)
    var v19 int64 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v6
}
