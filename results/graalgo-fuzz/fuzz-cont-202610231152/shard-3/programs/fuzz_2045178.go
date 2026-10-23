package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int; F1 int16; F2 string }
func genFn0() uint8 {
    v5 := []int{84, 42, 32767, 32768}
    var v6 uint64 = 130
    if len(v5) > 0 {
        v5[0] = 1
    }
    fmt.Println(v6)
    _ = v5
    _ = v6
    return 143
}

func main() {
    v1 := []string{}
    var v2 int64 = -250
    var v3 uint8 = 127
    var v4 bool = (-2.25 <= (1.5 + 0.0))
    var v7 int8 = 115
    v7 = (v7 | v7)
    if v4 {
        v8 := GenS1{F0: -32768}
        for fi9 := 0; fi9 < 6; fi9++ {
            _ = len(v1)
            _ = fi9
        }
        _ = v8
    }
    switch ("foo bar" + "a") {
    case "hello":
        v7 = v7
    case "foo bar":
        switch "foo bar" {
        case "foo bar":
            v7 = v7
        case "":
            v7 += (v7 * (-(v7)))
        }
    default:
        fmt.Println(v2, v7)
    }
    for fi10 := 0; fi10 < 14; fi10++ {
        v11 := GenS1{F0: int16(fi10)}
        v12 := GenS1{F0: 0}
        _ = fi10
        _ = v11
        _ = v12
    }
    for fi13 := 0; fi13 < 10; fi13++ {
        fmt.Println(v7)
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi13
    }
    v14 := GenS1{F0: 32767}
    v3 += v3
    v2 = -1
    switch 0 {
    case 32767:
        v15 := GenS1{F0: -32768}
        _ = v15
    case 870:
        v16 := GenS1{F0: 1000}
        _ = v16
    }
    v17 := GenS2{F0: 754, F1: 371, F2: (("\n\t" + "\n\t") + ("foo bar" + "foo bar"))}
    switch "hello" {
    case "foo bar":
        var v18 bool = v4
        _ = v18
    case "":
        v19 := GenS1{F0: -246}
        _ = v19
    case "hello":
        v20 := GenS1{F0: 220}
        _ = v20
    }
    v3 = (v3 / ((genFn0()) | 1))
    v7 += v7
    v21 := GenS2{F0: 32767, F1: 32767, F2: (("\n\t" + (("foo bar" + "a") + "")) + ((((((("\n\t" + "hello") + (("hello" + "hello") + "\n\t")) + "") + ((("hello" + (("" + "") + ("hello" + ""))) + "a") + "")) + "\n\t") + ((("foo bar" + "a") + (("\n\t" + "hello") + (("" + "\n\t") + ("foo bar" + "a")))) + "")) + "foo bar"))}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v17
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v7
}
