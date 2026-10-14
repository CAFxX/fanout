package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint }
func genFn0(p0 string) bool {
    switch "a" {
    case "":
        v7 := GenS1{F0: -307}
        _ = v7
    case "\n\t":
        v8 := GenS2{F0: 1}
        _ = v8
    case "hello":
        p0 = "hello"
    default:
        switch -298 {
        case 127:
            p0 = (p0 + p0)
        case 32768:
            p0 = ((p0 + "\n\t") + "foo bar")
        default:
            var v9 bool = true
            _ = v9
        }
    }
    v10 := GenS2{F0: 1}
    var v11 int16 = 42
    _ = p0
    _ = v10
    _ = v11
    return (!((2147483648 != 65535) || false))
}

func main() {
    var v1 bool = (42 <= 2)
    v2 := make(map[int]int)
    v3 := []uint32{2}
    v4 := []string{(("foo bar" + "foo bar") + ""), ""}
    var v5 uint = 338
    var v6 string = "hello"
    var v12 uint16 = 133
    for fi13 := 0; fi13 < 11; fi13++ {
        v14 := GenS1{F0: fi13}
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = fi13
        _ = v14
    }
    v15 := GenS3{F0: (v5 % ((v5) | 1))}
    fmt.Println(v12)
    fmt.Println(v5, v12)
    v16 := GenS2{F0: 2}
    switch "a" {
    case "":
        fmt.Println(v1, v5)
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi17 := 0; fi17 < 5; fi17++ {
        if len(v4) > 0 {
            v4[0] = "hello"
        }
        fmt.Println(fi17)
        _ = fi17
    }
    var v18 uint = v5
    if (209 <= 127) {
        var v19 uint = v18
        v6 = v6
        _ = v19
    }
    v6 = v6
    v20 := GenS2{F0: 42}
    var v21 int32 = 130
    var v22 float64 = 0.1
    var v23 uint32 = 558
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v15
    _ = v16
    _ = v18
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
