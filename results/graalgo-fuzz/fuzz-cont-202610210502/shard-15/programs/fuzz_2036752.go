package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool }
func genFn0(p0 int64, p1 uint64, p2 bool) bool {
    var v5 int16 = -32768
    for fi6 := 0; fi6 < 5; fi6++ {
        var v7 bool = p2
        _ = fi6
        _ = v7
    }
    if (!(p2 || (254 < 200))) {
        p2 = false
        switch -38 {
        case 9223372036854775807:
            var v8 int = 0
            _ = v8
        default:
            if p2 {
                var v9 int32 = 1
                _ = v9
            }
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return p2
}
func genFn1(p0 uint8, p1 string) int {
    if false {
        for fi10 := 0; fi10 < 8; fi10++ {
            v11 := GenS1{F0: -2147483648, F1: true}
            _ = fi10
            _ = v11
        }
        var v12 string = p1
        _ = v12
    }
    fmt.Println(p0)
    _ = p0
    _ = p1
    return -2147483648
}

func main() {
    v1 := []int{-9223372036854775808, 42}
    v2 := make(map[uint32]bool)
    var v3 int = -768
    var v4 uint8 = 169
    var v13 uint32 = 2
    v4 = v4
    var v14 uint32 = (v13 % (((^(v13))) | 1))
    for fi15 := 0; fi15 < 13; fi15++ {
        _ = len(v1)
        _ = fi15
    }
    if (v3 != (^(v3))) {
        var v16 bool = false
        _ = v16
    } else {
        if true {
            v17 := GenS1{F0: -2147483648, F1: true}
            _ = v17
        }
        var v18 uint64 = 18446744073709551615
        if false {
            for fi19 := 0; fi19 < 18; fi19++ {
                switch "hello" {
                case "hello":
                    v13 = v13
                case "a":
                    if true {
                        var v20 uint32 = 533
                        v21 := GenS1{F0: 42, F1: false}
                        fmt.Println(v14, v13)
                        _ = v20
                        _ = v21
                    }
                }
                if true {
                    for fi22 := 0; fi22 < 11; fi22++ {
                        if false {
                            v18 -= uint64(v4)
                        } else {
                            v23 := GenS1{F0: -356, F1: true}
                            _ = v23
                        }
                        _ = fi22
                    }
                    v13 = (v14 * (v14 + uint32(v18)))
                    switch genFn1((v4 * uint8(v18)), ("hello" + "a")) {
                    case 2147483647:
                        var v24 int16 = 1
                        _ = v24
                    case -1:
                        v4 += v4
                    default:
                        v3 += genFn1(1, ("a" + "\n\t"))
                    }
                }
                _ = fi19
            }
        }
        _ = v18
    }
    switch ((("" + ("\n\t" + (("a" + "a") + "hello"))) + ("" + "foo bar")) + "") {
    case "foo bar":
        var v25 uint64 = 1
        _ = v25
    case "hello":
        var v26 uint = 18446744073709551615
        _ = v26
    default:
        v4 = (v4 & (v4 - v4))
    }
    if true {
        v27 := GenS1{F0: 42, F1: false}
        _ = v27
    }
    v28 := GenS1{F0: -659, F1: (2 >= 18446744073709551615)}
    fmt.Println(v4)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v28
    _ = v3
    _ = v4
}
