package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int32; F1 bool; F2 int32 }
func genFn0(p0 uint16) uint64 {
    p0 -= 42
    var v7 uint32 = 2147483647
    _ = p0
    _ = v7
    return 4294967295
}

func main() {
    var v1 uint64 = 1
    var v2 string = ((("" + "foo bar") + "\n\t") + "hello")
    var v3 float64 = 3.14159
    var v4 int = 9223372036854775807
    v5 := make(map[string]bool)
    v6 := make(map[uint32]string)
    if (!(!true)) {
        var v8 uint = 18446744073709551615
        _ = v8
    } else {
        var v9 int64 = 764
        _ = v9
    }
    fmt.Println(v2, v3, v3)
    v10 := GenS1{F0: -32768}
    v11 := []int{(v4 | v4), (v4 & v4)}
    if len(v11) >= 2 {
        v11 = v11[1:len(v11)]
    }
    _ = len(v11)
    var v12 string = ("foo bar" + "foo bar")
    v11 = append(v11, -1)
    var v13 int8 = -35
    var v14 string = ((v2 + v12) + "")
    var v15 float64 = v3
    var v16 int32 = 1
    v17 := GenS2{F0: v16, F1: (!true), F2: (^(v16))}
    fmt.Println(v13)
    var v18 uint64 = genFn0(520)
    if false {
        if len(v11) >= 2 {
            v11 = v11[1:len(v11)]
        }
        v19 := GenS1{F0: 0}
        _ = v19
    } else {
        if false {
            if len(v11) >= 2 {
                v11 = v11[1:len(v11)]
            }
        }
        var v20 string = v2
        _ = v20
    }
    fmt.Println(v18)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
