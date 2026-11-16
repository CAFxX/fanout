package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
func genFn0(p0 int16) uint {
    switch "\n\t" {
    case "":
        var v5 uint32 = 358
        _ = v5
    default:
        var v6 int8 = -102
        _ = v6
    }
    for fi7 := 0; fi7 < 2; fi7++ {
        v8 := []int{(-(fi7))}
        p0 -= (-(p0))
        var v9 int8 = -1
        _ = fi7
        _ = v8
        _ = v9
    }
    p0 = (p0 | (p0 - p0))
    v10 := GenS1{F0: (!false)}
    switch "foo bar" {
    case "\n\t":
        p0 -= p0
    case "hello":
        v11 := GenS1{F0: (!(false || false))}
        _ = v11
    case "":
        v12 := GenS1{F0: (65535 != 417)}
        _ = v12
    }
    _ = p0
    _ = v10
    return 1
}

func main() {
    v1 := []int{}
    var v2 int8 = -78
    v3 := []uint32{}
    var v4 string = "\n\t"
    v2 -= (v2 | v2)
    if len(v3) > 0 {
        v3[0] = 1
    }
    v4 = (v4 + ((("foo bar" + ((v4 + v4) + ("\n\t" + (((v4 + v4) + ((v4 + "a") + (("hello" + (v4 + ((v4 + v4) + v4))) + (v4 + v4)))) + ((v4 + v4) + v4))))) + "foo bar") + ""))
    var v13 uint16 = 0
    for fi14 := 0; fi14 < 10; fi14++ {
        fmt.Println(v2)
        var v15 uint = 18446744073709551615
        _ = fi14
        _ = v15
    }
    var v16 int16 = 0
    v17 := GenS1{F0: true}
    var v18 uint32 = 0
    var v19 string = v4
    v4 = ((v19 + v19) + (v19 + v19))
    v1 = append(v1, 1)
    v13 = v13
    switch 458 {
    case 9223372036854775807:
        v18 = v18
    default:
        v19 = "hello"
    }
    var v20 bool = false
    var v21 uint32 = v18
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
}
