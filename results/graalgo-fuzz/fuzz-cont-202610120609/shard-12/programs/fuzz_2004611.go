package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string }
type GenS2 struct { F0 uint16 }
func genFn0(p0 int32, p1 uint64, p2 uint8) int64 {
    p1 = 18446744073709551615
    v5 := GenS1{F0: 127, F1: "foo bar"}
    var v6 int64 = -554
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return (v6 + (v6 & v6))
}
func genFn1(p0 int, p1 int32, p2 int) bool {
    var v7 uint64 = 91
    p1 = p1
    if ("" == "hello") {
        var v8 float32 = 0.0
        _ = v8
    } else {
        var v9 uint32 = 42
        if true {
            var v10 int32 = (p1 >> ((65535) & 63))
            var v11 float32 = 3.14159
            _ = v10
            _ = v11
        } else {
            v12 := GenS1{F0: (p0 - -9223372036854775808), F1: ((("hello" + "foo bar") + ((("a" + ("\n\t" + "a")) + "\n\t") + ("\n\t" + ""))) + "\n\t")}
            v13 := []int{(p0 * 127), (p2 / ((p2) | 1))}
            _ = v12
            _ = v13
        }
        v14 := GenS1{F0: p2, F1: "a"}
        _ = v14
        _ = v9
    }
    v15 := GenS1{F0: -723, F1: "a"}
    p1 += p1
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    _ = v7
    return true
}

func main() {
    var v1 uint32 = 65535
    var v2 bool = false
    var v3 uint64 = 42
    var v4 float32 = (-2.25 - -2.25)
    fmt.Println(v4, v1)
    v16 := []int{-9223372036854775808, -711, 127, -9223372036854775808}
    var v17 float64 = 0.0
    for fi18 := 0; fi18 < 15; fi18++ {
        var v19 uint8 = 255
        v4 = v4
        _ = fi18
        _ = v19
    }
    v20 := GenS1{F0: 0, F1: "a"}
    var v21 uint = 2
    switch (("a" + "foo bar") + "") {
    case "foo bar":
        var v22 uint64 = 18446744073709551615
        _ = v22
    }
    fmt.Println(v3, v4, v2)
    fmt.Println(v21)
    var v23 uint8 = 1
    v3 += 904
    v24 := GenS2{F0: 32767}
    v25 := GenS2{F0: 1}
    if v2 {
        var v26 uint32 = 1
        _ = v26
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v21
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
}
