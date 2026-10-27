package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int32; F1 uint32 }
func genFn0(p0 bool, p1 int32) int16 {
    p1 = p1
    v5 := GenS1{F0: 0}
    _ = p0
    _ = p1
    _ = v5
    return 32767
}
func genFn1() int16 {
    var v6 bool = false
    for fi7 := 0; fi7 < 16; fi7++ {
        fmt.Println(v6, fi7)
        switch ("\n\t" + "\n\t") {
        case "":
            var v8 int64 = 1
            _ = v8
        case "a":
            v9 := GenS1{F0: 0}
            _ = v9
        }
        if v6 {
            var v10 float64 = float64(float32(0.0))
            fmt.Println(v10, fi7)
            _ = v10
        }
        _ = fi7
    }
    _ = v6
    return 0
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[uint32]int)
    v3 := []int{-753, 32768}
    var v4 int32 = -94
    v4 += (v4 * v4)
    if true {
        var v11 uint32 = 16
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = v11
    }
    v12 := GenS1{F0: -83}
    fmt.Println(v4)
    var v13 bool = (!(0.0 <= 0.1))
    v14 := GenS2{F0: (v4 << ((42) & 63)), F1: 0}
    switch ("a" + "\n\t") {
    case "hello":
        v15 := GenS2{F0: 0, F1: 2147483647}
        _ = v15
    case "":
        v4 = v4
    }
    var v16 uint16 = 2
    switch "\n\t" {
    case "\n\t":
        if v13 {
            fmt.Println(v13, v4, v4)
        }
    case "":
        fmt.Println(v13)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
