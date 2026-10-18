package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint16; F1 int8; F2 uint8 }

func main() {
    v1 := make(map[int]string)
    v2 := []string{"\n\t"}
    var v3 int16 = 0
    var v4 uint32 = 2147483647
    var v5 float32 = 1.5
    var v6 int16 = v3
    v7 := GenS1{F0: 128}
    var v8 int64 = 937
    for fi9 := 0; fi9 < 13; fi9++ {
        var v10 int8 = 1
        var v11 int8 = int8(fi9)
        v3 += int16(v10)
        _ = fi9
        _ = v10
        _ = v11
    }
    var v12 float64 = 1.5
    v13 := GenS2{F0: 344, F1: 22, F2: 255}
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v14 int8 = 100
    var v15 int = 902
    v16 := GenS2{F0: 65535, F1: v14, F2: 255}
    var v17 int8 = (v14 - (v14 << ((691) & 63)))
    var v18 int32 = -146
    var v19 int = v15
    var v20 string = ("hello" + "foo bar")
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
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
