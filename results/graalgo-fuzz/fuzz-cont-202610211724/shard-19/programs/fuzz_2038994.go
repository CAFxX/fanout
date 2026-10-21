package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint; F1 uint32 }
func genFn0(p0 int16) int64 {
    var v4 int8 = 22
    p0 = p0
    if true {
        v5 := []int{0, 32767, 2147483647}
        _ = v5
    } else {
        fmt.Println(p0)
    }
    v6 := GenS1{F0: ""}
    _ = p0
    _ = v4
    _ = v6
    return 9223372036854775807
}
func genFn1(p0 bool, p1 uint32, p2 string) int {
    if (-440 >= 682) {
        v7 := []int{1, 32767, 310}
        fmt.Println(p0, p0, p1)
        fmt.Println(p1, p1, p2)
        _ = v7
    }
    fmt.Println(p2)
    v8 := []int{9223372036854775807, 37, 128}
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 9223372036854775807
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[uint32]bool)
    v3 := []int{1, 679}
    for fi9 := 0; fi9 < 16; fi9++ {
        fmt.Println(fi9)
        var v10 bool = true
        var v11 int64 = 462
        _ = fi9
        _ = v10
        _ = v11
    }
    v12 := GenS1{F0: "\n\t"}
    var v13 uint16 = 32767
    v14 := v1["\n\t"]
    var v15 bool = (9223372036854775807 != 42)
    fmt.Println(v13, v15, v14)
    for fi16 := 0; fi16 < 13; fi16++ {
        v1[((v14 + v14) + v14)] = (((v14 + (v14 + "")) + (((v14 + ("a" + (v14 + (v14 + "foo bar")))) + v14) + "foo bar")) + (v14 + v14))
        v17 := v1["a"]
        v18 := GenS2{F0: 0, F1: 42}
        _ = fi16
        _ = v17
        _ = v18
    }
    v19 := GenS1{F0: (("a" + v14) + (v14 + ""))}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v19
    _ = v2
    _ = v3
}
