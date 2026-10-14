package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint; F1 string }
type GenS3 struct { F0 uint64; F1 uint16; F2 uint }

func main() {
    v1 := make(map[string]bool)
    var v2 float64 = 0.0
    var v3 uint = 0
    var v4 bool = true
    v5 := []uint32{}
    _ = len(v1)
    fmt.Println(v4, v2)
    fmt.Println(v2)
    v2 = v2
    switch 32768 {
    case 32767:
        var v6 string = "hello"
        _ = v6
    }
    v4 = (1 != -937)
    if true {
        if false {
            v2 += float64((0.1 + 3.14159))
        }
        var v7 int32 = 0
        _ = v7
    }
    v8 := GenS1{F0: "a"}
    var v9 uint32 = 0
    var v10 float64 = v2
    var v11 int8 = -106
    var v12 float64 = (v2 / 3.14159)
    v11 += (-(v11))
    if (!(true || false)) {
        for fi13 := 0; fi13 < 4; fi13++ {
            v14 := GenS2{F0: (v3 % (((v3 % (((^(v3))) | 1))) | 1)), F1: "hello"}
            for fi15 := 0; fi15 < 14; fi15++ {
                v16 := GenS3{F0: 2, F1: 0, F2: (v3 * (v3 | (v3 + v3)))}
                var v17 int = fi13
                var v18 string = "foo bar"
                _ = fi15
                _ = v16
                _ = v17
                _ = v18
            }
            v11 = -1
            _ = fi13
            _ = v14
        }
    }
    v19 := GenS3{F0: 876, F1: 255, F2: v3}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
