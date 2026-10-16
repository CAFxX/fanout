package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int; F2 int64 }

func main() {
    var v1 uint8 = 127
    var v2 uint = 1
    v3 := make(map[string]uint32)
    var v4 int32 = 2147483647
    v5 := []string{}
    v6 := make(map[uint32]string)
    v7 := GenS1{F0: ((400 < 1) || (!((v1 << ((v2) & 63)) < v1))), F1: int(v4), F2: 650}
    if len(v5) > 0 {
        v5[0] = ("\n\t" + ("hello" + "\n\t"))
    }
    v2 += (v2 & (v2 << ((v2) & 63)))
    var v8 uint = v2
    fmt.Println(v8, v4)
    for fi9 := 0; fi9 < 17; fi9++ {
        v10 := GenS1{F0: (!(!(!(980 < 1)))), F1: 1, F2: 1}
        _ = fi9
        _ = v10
    }
    v4 += v4
    var v11 float64 = 1e10
    var v12 uint = (v8 >> ((v8) & 63))
    var v13 string = ("\n\t" + ("\n\t" + "\n\t"))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
