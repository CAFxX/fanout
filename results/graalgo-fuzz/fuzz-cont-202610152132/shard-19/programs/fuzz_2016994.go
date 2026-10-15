package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string; F2 uint16 }
type GenS2 struct { F0 string; F1 int16 }
func genFn0(p0 int16, p1 string) uint16 {
    var v5 int32 = -2147483648
    v6 := GenS1{F0: 1, F1: ((p1 + (p1 + ("foo bar" + p1))) + p1), F2: 583}
    for fi7 := 0; fi7 < 4; fi7++ {
        v8 := []int{32768, 0, (-(fi7))}
        fmt.Println(p0)
        fmt.Println(p0, p0, v5)
        _ = fi7
        _ = v8
    }
    switch (p1 + ((p1 + p1) + (("\n\t" + ("" + "hello")) + p1))) {
    case "hello":
        v9 := []int{1, 128}
        _ = v9
    case "\n\t":
        var v10 uint = 506
        _ = v10
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 888
}
func genFn1(p0 int, p1 int32) bool {
    v11 := GenS1{F0: 127, F1: "", F2: genFn0(1, "a")}
    p1 -= (p1 >> ((2) & 63))
    p1 = (-(p1))
    p1 -= p1
    _ = p0
    _ = p1
    _ = v11
    return false
}

func main() {
    var v1 uint8 = 128
    var v2 uint16 = 2
    v3 := make(map[string]bool)
    var v4 int64 = 1
    v12 := []int{1, -1, 541, -174}
    if len(v12) >= 2 {
        v12 = v12[1:len(v12)]
    }
    v13 := GenS2{F0: ((("" + ("hello" + ("foo bar" + "foo bar"))) + "foo bar") + ("a" + "a")), F1: 1000}
    if len(v12) > 0 {
        v12[0] = 128
    }
    var v14 int32 = 1
    var v15 int = -1
    v14 -= v14
    _ = len(v12)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
}
