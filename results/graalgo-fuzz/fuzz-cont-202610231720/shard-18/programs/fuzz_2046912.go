package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int8 }
func genFn0(p0 int32, p1 bool) uint32 {
    switch "a" {
    case "\n\t":
        switch 32767 {
        case -52:
            var v5 int16 = 1
            _ = v5
        case -999:
            var v6 string = "a"
            _ = v6
        default:
            v7 := []int{-258, 1, 128, -1}
            _ = v7
        }
    case "a":
        switch ("a" + ("\n\t" + "hello")) {
        case "a":
            v8 := []int{}
            _ = v8
        default:
            var v9 uint64 = 18446744073709551615
            _ = v9
        }
    }
    fmt.Println(p1, p0)
    fmt.Println(p1)
    v10 := []int{-9223372036854775808, -9223372036854775808, 128}
    v10 = append(v10, 1)
    _ = p0
    _ = p1
    _ = v10
    return 42
}

func main() {
    v1 := []int{}
    var v2 string = "a"
    var v3 float32 = 1e10
    var v4 uint = 752
    v3 = v3
    var v11 float32 = 3.14159
    fmt.Println(v11)
    v4 -= 404
    var v12 int16 = 42
    if true {
        _ = len(v1)
        fmt.Println(v4, v4, v11)
        for fi13 := 0; fi13 < 11; fi13++ {
            v2 = v2
            v14 := GenS1{F0: 65535}
            _ = fi13
            _ = v14
        }
    } else {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        var v15 bool = (-115 >= 0)
        fmt.Println(v4, v2)
        _ = v15
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    switch 127 {
    case 127:
        var v16 float64 = 1e10
        _ = v16
    case 0:
        switch 9223372036854775807 {
        case -97:
            fmt.Println(v2, v12, v4)
        case -207:
            v4 = 251
        case 1:
            var v17 string = v2
            _ = v17
        default:
            var v18 string = (v2 + (v2 + ""))
            _ = v18
        }
    }
    var v19 uint8 = 1
    v20 := GenS2{F0: 42}
    v12 = v12
    v2 = v2
    fmt.Println(v19, v3, v4)
    _ = len(v1)
    for fi21 := 0; fi21 < 5; fi21++ {
        var v22 string = ((v2 + "a") + "foo bar")
        v23 := GenS3{F0: 1}
        _ = fi21
        _ = v22
        _ = v23
    }
    v24 := GenS1{F0: 512}
    switch 450 {
    case -1:
        v25 := GenS4{F0: 100}
        _ = v25
    case 9223372036854775807:
        var v26 float64 = 3.14159
        _ = v26
    default:
        var v27 int32 = -1
        _ = v27
    }
    fmt.Println(v19, v3, v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v19
    _ = v2
    _ = v20
    _ = v24
    _ = v3
    _ = v4
}
