package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 bool; F1 uint32 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int64 }
func genFn0() int16 {
    var v7 uint8 = 128
    v7 = (v7 << ((0) & 63))
    v7 = v7
    _ = v7
    return 516
}
func genFn1(p0 uint32, p1 uint) uint32 {
    v8 := GenS1{F0: 128}
    p0 = 0
    fmt.Println(p1)
    p1 = p1
    var v9 int32 = 2147483647
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return p0
}

func main() {
    var v1 uint = 192
    v2 := []string{"", "", ""}
    v3 := make(map[string]string)
    var v4 uint = (v1 >> ((v1) & 63))
    v5 := []string{"hello", ("a" + "a"), ("a" + "a"), ((("a" + ("" + "foo bar")) + "\n\t") + ("" + "hello"))}
    var v6 bool = false
    var v10 int = 0
    for fi11 := 0; fi11 < 19; fi11++ {
        var v12 int8 = 0
        _ = fi11
        _ = v12
    }
    switch ("" + "a") {
    case "hello":
        var v13 int16 = -181
        _ = v13
    case "foo bar":
        var v14 float32 = (3.14159 - 3.14159)
        _ = v14
    case "\n\t":
        fmt.Println(v10, v1)
    default:
        if v6 {
            if len(v2) > 0 {
                v2[0] = "foo bar"
            }
        }
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v15 int16 = 32767
    v16 := GenS2{F0: true, F1: 673}
    v17 := GenS3{F0: 2147483647}
    v15 += genFn0()
    var v18 int32 = 0
    var v19 int8 = -96
    v20 := GenS4{F0: -1}
    v1 -= v4
    v21 := GenS2{F0: (v6 && (!(!v6))), F1: 4294967295}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
