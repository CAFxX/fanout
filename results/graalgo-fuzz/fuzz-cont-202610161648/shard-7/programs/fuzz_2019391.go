package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32 }
type GenS2 struct { F0 int16; F1 bool }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int; F1 string }
func genFn0(p0 int64, p1 string) uint16 {
    v5 := GenS1{F0: "a", F1: -2147483648}
    if (-128 > 127) {
        var v6 int32 = 615
        switch (p1 + p1) {
        case "\n\t":
            switch 4294967295 {
            case 886:
                v7 := []int{-402}
                _ = v7
            case 4294967295:
                p1 = ("" + (p1 + ("hello" + (((p1 + (p1 + p1)) + (p1 + (p1 + ((("hello" + p1) + p1) + "\n\t")))) + (((p1 + ((p1 + p1) + "foo bar")) + p1) + p1)))))
            case 2147483647:
                p1 = p1
            }
        }
        v8 := GenS1{F0: p1, F1: (v6 * 2147483647)}
        _ = v6
        _ = v8
    }
    var v9 float64 = 3.14159
    switch -405 {
    case -162:
        fmt.Println(p1, p0, p0)
    case -9223372036854775808:
        p0 = 9223372036854775807
    default:
        v10 := GenS2{F0: -32768, F1: false}
        _ = v10
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v9
    return 353
}

func main() {
    v1 := make(map[int]string)
    var v2 uint16 = 113
    var v3 string = ("a" + "")
    v4 := []string{v3, ((v3 + "") + ("hello" + ("a" + (("\n\t" + v3) + v3)))), "\n\t", v3}
    v4 = append(v4, ("hello" + v3))
    if (!false) {
        var v11 int32 = -650
        fmt.Println(v3, v3)
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = v11
    } else {
        for fi12 := 0; fi12 < 14; fi12++ {
            for fi13 := 0; fi13 < 12; fi13++ {
                var v14 bool = (18446744073709551615 >= 1)
                v1[(fi12 + (fi12 ^ 4294967295))] = (v3 + (((v3 + v3) + (v3 + "hello")) + "foo bar"))
                switch fi13 {
                case 42:
                    v2 = genFn0(0, v3)
                }
                _ = fi13
                _ = v14
            }
            if len(v4) > 0 {
                v4[0] = v3
            }
            _ = fi12
        }
        var v15 int32 = 42
        _ = v15
    }
    if (509 != 42) {
        fmt.Println(v2, v3)
        var v16 int64 = 1
        _ = v16
    }
    v4 = append(v4, v3)
    v17 := GenS3{F0: 255}
    v18 := GenS4{F0: -1, F1: v3}
    v4 = append(v4, ("a" + v3))
    var v19 uint64 = 42
    v19 -= (v19 * (v19 | v19))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
