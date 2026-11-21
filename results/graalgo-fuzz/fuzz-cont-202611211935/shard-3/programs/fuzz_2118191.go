package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
func genFn0(p0 uint32, p1 int8, p2 int) bool {
    v7 := []int{}
    p2 = (p2 & (p2 - 322))
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return (421 <= 41)
}
func genFn1() int16 {
    var v8 bool = true
    var v9 uint = 446
    _ = v8
    _ = v9
    return 201
}

func main() {
    var v1 int32 = 0
    var v2 int = 0
    v3 := []uint32{0}
    var v4 int = v2
    var v5 uint16 = 1
    var v6 string = "hello"
    for fi10 := 0; fi10 < 1; fi10++ {
        v2 = (v2 ^ v4)
        _ = fi10
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v11 := GenS1{F0: -1}
    switch "foo bar" {
    case "foo bar":
        var v12 int64 = 812
        _ = v12
    case "a":
        v13 := GenS1{F0: genFn1()}
        _ = v13
    default:
        fmt.Println(v6, v6, v6)
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v14 uint64 = 1
    for fi15 := 0; fi15 < 8; fi15++ {
        if len(v3) > 0 {
            v3[0] = 2
        }
        v3 = append(v3, 1)
        _ = fi15
    }
    for fi16 := 0; fi16 < 7; fi16++ {
        var v17 uint64 = 4294967295
        v17 = 0
        _ = fi16
        _ = v17
    }
    v2 = v2
    fmt.Println(v5, v5)
    v18 := GenS1{F0: 436}
    v1 = (v1 * v1)
    var v19 uint32 = 65535
    fmt.Println(v14)
    for fi20 := 0; fi20 < 16; fi20++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi20
    }
    if false {
        var v21 float64 = (0.1 - ((3.14159 / (float64(1.5) * 0.1)) + 3.14159))
        v22 := GenS1{F0: 0}
        _ = v21
        _ = v22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v14
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
