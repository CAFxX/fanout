package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int }
type GenS2 struct { F0 bool; F1 bool; F2 int64 }
type GenS3 struct { F0 uint16; F1 uint }
func genFn0(p0 string, p1 int8, p2 bool) uint32 {
    var v7 string = p0
    v8 := GenS1{F0: -128, F1: 682}
    for fi9 := 0; fi9 < 13; fi9++ {
        if ((!(42 < 4294967295)) && p2) {
            v10 := []int{-809, -1, (fi9 * fi9)}
            fmt.Println(p2)
            _ = v10
        } else {
            v11 := []int{}
            fmt.Println(v7, fi9, p1)
            p2 = (p2 && ((!p2) && true))
            _ = v11
        }
        _ = fi9
    }
    p0 = ((("" + ((v7 + (p0 + "\n\t")) + "\n\t")) + v7) + (("hello" + ((v7 + p0) + (p0 + (v7 + v7)))) + p0))
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 42
}

func main() {
    var v1 uint8 = 230
    v2 := make(map[string]string)
    var v3 uint32 = 2147483648
    v4 := []string{(("hello" + "") + "\n\t"), "a"}
    var v5 uint16 = 0
    var v6 int32 = -1
    switch 32767 {
    case -2147483648:
        fmt.Println(v6)
    case -1:
        v3 = 2147483647
    }
    v3 += 0
    v12 := GenS2{F0: true, F1: false, F2: -94}
    fmt.Println(v1, v6, v5)
    fmt.Println(v6)
    delete(v2, ("foo bar" + "hello"))
    for fi13 := 0; fi13 < 8; fi13++ {
        v3 -= 2147483647
        v14 := GenS3{F0: 65535, F1: 18446744073709551615}
        v4 = append(v4, "foo bar")
        _ = fi13
        _ = v14
    }
    if (!(0 < 0)) {
        var v15 float32 = -2.25
        _ = v15
    } else {
        v5 = (v5 + v5)
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
