package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16; F2 int }

func main() {
    var v1 int32 = 1
    var v2 string = "hello"
    var v3 bool = (!(!true))
    var v4 uint8 = 145
    var v5 uint = 42
    var v6 int = 128
    v4 += v4
    var v7 int = v6
    var v8 int = 127
    var v9 uint16 = 42
    v5 -= 255
    var v10 int64 = -9223372036854775808
    v8 = 310
    var v11 int16 = int16(v4)
    var v12 int64 = -9223372036854775808
    fmt.Println(v6, v6)
    v2 = v2
    var v13 bool = false
    if v3 {
        switch v6 {
        case 1:
            fmt.Println(v5)
        }
    }
    v9 += v9
    for fi14 := 0; fi14 < 10; fi14++ {
        v15 := GenS1{F0: (-(v6)), F1: v11, F2: -828}
        v5 = (v5 & (v5 + uint(v9)))
        _ = fi14
        _ = v15
    }
    var v16 uint8 = v4
    v12 += 42
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
