package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int; F1 int16 }

func main() {
    var v1 float32 = 1e10
    var v2 float64 = 0.1
    var v3 int8 = 42
    var v4 uint16 = 683
    var v5 uint16 = v4
    var v6 int8 = (v3 % (((^(v3))) | 1))
    var v7 int8 = 100
    v8 := GenS1{F0: 748}
    v9 := []int{713, 127, 32767, -683}
    var v10 string = ("foo bar" + ("hello" + "hello"))
    v5 += (v5 ^ v5)
    var v11 uint8 = 255
    var v12 string = v10
    fmt.Println(v4)
    if len(v9) >= 2 {
        v9 = v9[1:len(v9)]
    }
    _ = len(v9)
    var v13 uint = 594
    var v14 uint8 = 23
    for fi15 := 0; fi15 < 14; fi15++ {
        var v16 float64 = v2
        var v17 uint = 1
        _ = fi15
        _ = v16
        _ = v17
    }
    v18 := GenS2{F0: 42}
    switch 32768 {
    case 128:
        v5 -= 32767
    case 42:
        v19 := GenS3{F0: 26, F1: 42}
        _ = v19
    case 2147483647:
        var v20 string = v10
        _ = v20
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
