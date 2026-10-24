package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool }
func genFn0(p0 int64, p1 uint64) int8 {
    fmt.Println(p0)
    switch "a" {
    case "hello":
        p0 -= p0
    case "\n\t":
        v5 := GenS1{F0: 501, F1: true}
        _ = v5
    }
    p1 += p1
    var v6 float64 = 0.1
    _ = p0
    _ = p1
    _ = v6
    return 0
}
func genFn1() string {
    v7 := []int{}
    _ = len(v7)
    v8 := GenS1{F0: 42, F1: false}
    _ = v7
    _ = v8
    return (("" + "") + "foo bar")
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int16 = -505
    v3 := make(map[string]bool)
    v4 := make(map[int]int)
    v1[2] = 42
    var v9 bool = (!(-9223372036854775808 < -9223372036854775808))
    delete(v3, ("foo bar" + ("a" + "hello")))
    for fi10 := 0; fi10 < 15; fi10++ {
        fmt.Println(v2)
        switch "\n\t" {
        case "a":
            delete(v1, 2)
        case "foo bar":
            for fi11 := 0; fi11 < 6; fi11++ {
                if v9 {
                    var v12 uint32 = 42
                    _ = v12
                } else {
                    v9 = ((-31 != -128) || ((v9 || (((!v9) && (!(2 <= 695))) || (255 != 2))) || true))
                }
                var v13 bool = v9
                _ = fi11
                _ = v13
            }
        default:
            v1[42] = 4294967295
        }
        var v14 uint64 = 565
        _ = fi10
        _ = v14
    }
    v15 := GenS1{F0: 2147483648, F1: v9}
    v9 = (v9 || v9)
    for fi16 := 0; fi16 < 7; fi16++ {
        v9 = (301 <= 468)
        v2 -= v2
        v9 = v9
        _ = fi16
    }
    v17 := GenS1{F0: 2, F1: (!((!(781 < 565)) && v9))}
    fmt.Println(v9, v2)
    v18 := GenS1{F0: 489, F1: false}
    var v19 int64 = -727
    v9 = v9
    v9 = false
    if false {
        for fi20 := 0; fi20 < 9; fi20++ {
            v9 = (4294967295 >= 2)
            _ = fi20
        }
        fmt.Println(v19)
    } else {
        v4[235] = -1
        for fi21 := 0; fi21 < 10; fi21++ {
            fmt.Println(fi21)
            if v9 {
                if v9 {
                    var v22 uint32 = 2147483647
                    v23 := GenS1{F0: v22, F1: true}
                    for fi24 := 0; fi24 < 2; fi24++ {
                        if (89 != -128) {
                            v2 -= -387
                        }
                        v25 := GenS1{F0: 437, F1: v9}
                        _ = fi24
                        _ = v25
                    }
                    _ = v22
                    _ = v23
                } else {
                    if false {
                        v26 := []int{fi21, (fi21 * 2147483647), fi21, (fi21 % ((0) | 1))}
                        _ = v26
                    } else {
                        v27 := GenS1{F0: 4294967295, F1: (v9 || (3.14159 < 0.1))}
                        _ = v27
                    }
                }
                v28 := GenS1{F0: 469, F1: v9}
                _ = v28
            } else {
                v29 := GenS1{F0: 42, F1: v9}
                var v30 uint = 18446744073709551615
                _ = v29
                _ = v30
            }
            _ = fi21
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v15
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
