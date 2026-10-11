package main

import (
    "fmt"
)

type GenS1 struct { F0 int }

func main() {
    var v1 int8 = -69
    var v2 float64 = (3.14159 * 0.1)
    var v3 bool = false
    var v4 int = 696
    fmt.Println(v4, v1, v4)
    v2 += (float64(1.5) + v2)
    var v5 int32 = 487
    var v6 uint16 = 530
    v7 := GenS1{F0: (v4 * v4)}
    for fi8 := 0; fi8 < 4; fi8++ {
        v1 += 100
        _ = fi8
    }
    switch 4294967295 {
    case 42:
        for fi9 := 0; fi9 < 14; fi9++ {
            var v10 float64 = (v2 * float64((3.14159 / 0.1)))
            v11 := []int{fi9, (fi9 / (((fi9 * (fi9 ^ v4))) | 1)), fi9}
            _ = fi9
            _ = v10
            _ = v11
        }
    case -448:
        v12 := []int{v4, v4}
        _ = v12
    default:
        v13 := GenS1{F0: 2147483647}
        _ = v13
    }
    v3 = (v3 || (v3 && v3))
    if (("a" < (("a" + "\n\t") + "")) || true) {
        var v14 uint16 = v6
        _ = v14
    }
    v5 = (v5 | v5)
    v1 = 115
    v4 += 127
    v15 := []int{(v4 / ((32767) | 1)), (v4 << ((507) & 63)), 1}
    switch "hello" {
    case "hello":
        fmt.Println(v2)
    case "a":
        v15 = append(v15, v4)
    case "\n\t":
        if (("\n\t" + ((("a" + "") + "") + "foo bar")) >= "hello") {
            v6 = (v6 + 1)
            fmt.Println(v5, v2)
            var v16 uint32 = 65535
            _ = v16
        } else {
            v1 = v1
        }
    }
    for fi17 := 0; fi17 < 10; fi17++ {
        var v18 uint64 = 18446744073709551615
        switch "hello" {
        case "foo bar":
            fmt.Println(v2, v5, v3)
        case "hello":
            var v19 uint64 = (v18 << ((255) & 63))
            _ = v19
        }
        _ = fi17
        _ = v18
    }
    var v20 int32 = 323
    if (!((!v3) && false)) {
        v15 = append(v15, v4)
        if false {
            var v21 int8 = v1
            _ = v21
        }
        if len(v15) > 0 {
            v15[0] = (v4 - v4)
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
