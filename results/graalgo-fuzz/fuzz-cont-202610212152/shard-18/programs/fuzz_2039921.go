package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint; F2 bool }
type GenS2 struct { F0 int16; F1 bool; F2 uint }
func genFn0(p0 int16) int32 {
    p0 = p0
    var v4 bool = ((65535 > 23) && true)
    switch -2147483648 {
    case -2147483648:
        switch 0 {
        case 467:
            v5 := GenS1{F0: -431, F1: 482, F2: (p0 >= 32767)}
            _ = v5
        case 631:
            p0 -= (p0 + -160)
        case -1:
            fmt.Println(v4, v4)
        }
    case -455:
        p0 -= (-(p0))
    case -647:
        switch ((("" + "a") + "\n\t") + ("" + "\n\t")) {
        case "a":
            var v6 int = 965
            _ = v6
        case "hello":
            if true {
                p0 += -105
            } else {
                fmt.Println(p0)
                var v7 bool = (v4 && (p0 >= p0))
                _ = v7
            }
        }
    default:
        v8 := []int{-1, 0, 0}
        _ = v8
    }
    fmt.Println(p0, p0)
    _ = p0
    _ = v4
    return 0
}
func genFn1(p0 uint) bool {
    p0 = 4294967295
    switch (("foo bar" + ("\n\t" + ("a" + ("foo bar" + ("" + "hello"))))) + ("" + "foo bar")) {
    case "":
        v9 := []int{int(p0), 9223372036854775807, 2147483647, 4294967295}
        _ = v9
    case "foo bar":
        if false {
            if true {
                p0 += (p0 - p0)
            } else {
                p0 += p0
                fmt.Println(p0)
            }
        }
    }
    var v10 int8 = 127
    v10 = v10
    _ = p0
    _ = v10
    return false
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[int]int)
    var v3 int8 = 0
    var v11 int8 = v3
    var v12 uint64 = 17
    v3 = -1
    v13 := GenS2{F0: 251, F1: (!true), F2: 255}
    v11 -= v11
    var v14 uint32 = 2
    v14 = (v14 / ((v14) | 1))
    var v15 bool = (0.0 == 1.5)
    v15 = ((3.14159 == -2.25) || false)
    v16 := GenS1{F0: genFn0(0), F1: 255, F2: (!false)}
    v15 = (709 == 368)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
}
