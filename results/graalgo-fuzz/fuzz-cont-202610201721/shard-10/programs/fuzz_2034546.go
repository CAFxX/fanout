package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8 }
type GenS2 struct { F0 uint8; F1 uint32; F2 uint64 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint }
func genFn0(p0 int8, p1 uint16, p2 bool) int64 {
    fmt.Println(p1, p0, p1)
    v4 := []int{-702, 42}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return -1
}

func main() {
    var v1 uint8 = 221
    v2 := make(map[int]string)
    var v3 bool = (!true)
    v5 := GenS1{F0: 1, F1: (v1 >> ((18446744073709551615) & 63))}
    v6 := []int{4294967295, 1, 128}
    v7 := GenS2{F0: 1, F1: 511, F2: 335}
    v8 := GenS2{F0: v1, F1: 2, F2: 2}
    switch (("hello" + (("foo bar" + "hello") + "hello")) + ("\n\t" + ("a" + ("a" + ("\n\t" + "foo bar"))))) {
    case "foo bar":
        fmt.Println(v1)
    case "hello":
        for fi9 := 0; fi9 < 4; fi9++ {
            v3 = v3
            var v10 bool = (100 >= 1)
            v11 := GenS1{F0: 2, F1: (^(v1))}
            _ = fi9
            _ = v10
            _ = v11
        }
    default:
        for fi12 := 0; fi12 < 3; fi12++ {
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            v2[(-(fi12))] = ("foo bar" + "hello")
            _ = fi12
        }
    }
    v13 := GenS3{F0: v1}
    delete(v2, 209)
    v1 = 42
    var v14 int32 = 2147483647
    if (2 == 2) {
        var v15 uint64 = 2
        _ = v15
    } else {
        for fi16 := 0; fi16 < 9; fi16++ {
            v14 -= 0
            delete(v2, (^(fi16)))
            v17 := GenS4{F0: 0}
            _ = fi16
            _ = v17
        }
        fmt.Println(v14, v3, v14)
    }
    var v18 int32 = 42
    var v19 uint16 = 0
    for fi20 := 0; fi20 < 1; fi20++ {
        switch 1 {
        case 4294967295:
            _ = len(v6)
        default:
            var v21 uint64 = 1
            _ = v21
        }
        _ = fi20
    }
    if v3 {
        v19 -= (v19 ^ v19)
        var v22 uint32 = 2
        v18 -= 1
        _ = v22
    } else {
        switch -913 {
        case 654:
            v3 = v3
        case 32768:
            v2[127] = "foo bar"
        default:
            v18 = v18
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
