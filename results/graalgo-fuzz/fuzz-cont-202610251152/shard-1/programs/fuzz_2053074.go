package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int8; F1 bool; F2 string }
func genFn0(p0 uint, p1 int) int32 {
    for fi6 := 0; fi6 < 12; fi6++ {
        v7 := GenS1{F0: p0}
        _ = fi6
        _ = v7
    }
    fmt.Println(p0, p1)
    var v8 int16 = 1000
    _ = p0
    _ = p1
    _ = v8
    return 0
}
func genFn1(p0 uint16, p1 int16) int32 {
    v9 := GenS1{F0: 492}
    for fi10 := 0; fi10 < 8; fi10++ {
        var v11 uint32 = 2147483648
        var v12 string = (("" + ("a" + "foo bar")) + (("foo bar" + "\n\t") + "hello"))
        _ = fi10
        _ = v11
        _ = v12
    }
    p1 += p1
    _ = p0
    _ = p1
    _ = v9
    return 42
}

func main() {
    var v1 uint = 2
    var v2 uint64 = 1
    var v3 float64 = 3.14159
    v4 := []string{"foo bar", "\n\t", "a", ("hello" + "")}
    var v5 uint64 = (v2 + (^(v2)))
    v3 = v3
    v3 = v3
    var v13 uint8 = 255
    fmt.Println(v2, v13, v5)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println(v13)
    v14 := GenS1{F0: v1}
    v15 := GenS2{F0: -38, F1: true, F2: ((("a" + "hello") + "a") + "a")}
    var v16 bool = (true && (v3 <= v3))
    switch 128 {
    case 9223372036854775807:
        var v17 int8 = 76
        _ = v17
    }
    fmt.Println(v13)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
