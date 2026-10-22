package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16 }
type GenS2 struct { F0 uint32; F1 uint32; F2 uint64 }

func main() {
    var v1 uint = 840
    v2 := []uint32{0}
    var v3 uint16 = 65535
    v4 := []int{824}
    v5 := []string{"hello", "", ("\n\t" + "a"), "a"}
    var v6 float64 = 3.14159
    var v7 int16 = 1
    if (0 == 2) {
        fmt.Println(v3, v3, v7)
        v8 := GenS1{F0: 100, F1: 2}
        _ = v8
    } else {
        fmt.Println(v7, v6)
    }
    v9 := GenS1{F0: -84, F1: v3}
    var v10 uint16 = v3
    for fi11 := 0; fi11 < 16; fi11++ {
        var v12 int8 = 100
        _ = fi11
        _ = v12
    }
    v13 := GenS2{F0: 920, F1: 2147483648, F2: 2}
    v6 = 3.14159
    if len(v4) > 0 {
        v4[0] = 705
    }
    var v14 string = ("foo bar" + "hello")
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
