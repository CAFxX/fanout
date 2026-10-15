package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int; F1 int }
func genFn0(p0 int16, p1 bool) int64 {
    fmt.Println(p0, p1)
    v6 := []int{-696, 608}
    _ = p0
    _ = p1
    _ = v6
    return -895
}

func main() {
    var v1 uint = 2
    var v2 int32 = 637
    v3 := make(map[int]uint32)
    v4 := make(map[uint32]int)
    var v5 int64 = 985
    switch "foo bar" {
    case "a":
        v5 = (v5 + genFn0(42, (true && true)))
    default:
        v7 := GenS1{F0: 750}
        _ = v7
    }
    for fi8 := 0; fi8 < 2; fi8++ {
        var v9 int64 = (v5 >> (((v1 / ((669) | 1))) & 63))
        var v10 int = -25
        _ = fi8
        _ = v10
        _ = v9
    }
    fmt.Println(v1)
    if false {
        if true {
            var v11 string = ("hello" + "a")
            v12 := GenS1{F0: v1}
            _ = v11
            _ = v12
        }
        if (v1 != v1) {
            fmt.Println(v1)
        } else {
            v13 := GenS1{F0: (v1 >> ((v1) & 63))}
            if (false || false) {
                v14 := GenS2{F0: 643, F1: 32767}
                switch 32767 {
                case 9223372036854775807:
                    var v15 int = -376
                    _ = v15
                case -1:
                    for fi16 := 0; fi16 < 13; fi16++ {
                        v1 -= v1
                        _ = fi16
                    }
                }
                switch (("a" + "") + (("hello" + (("hello" + ("foo bar" + ("\n\t" + (("\n\t" + "a") + "hello")))) + "a")) + "foo bar")) {
                case "foo bar":
                    v2 = (v2 + v2)
                case "":
                    var v17 int16 = -32768
                    _ = v17
                default:
                    v4[1] = 817
                }
                _ = v14
            }
            _ = v13
        }
        var v18 string = "hello"
        _ = v18
    }
    fmt.Println(v1)
    v1 = v1
    var v19 uint32 = 4294967295
    for fi20 := 0; fi20 < 9; fi20++ {
        for fi21 := 0; fi21 < 13; fi21++ {
            v22 := []int{32767, (^(fi20))}
            _ = len(v22)
            _ = fi21
            _ = v22
        }
        _ = fi20
    }
    var v23 string = "\n\t"
    v24 := GenS1{F0: (v1 % ((v1) | 1))}
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
