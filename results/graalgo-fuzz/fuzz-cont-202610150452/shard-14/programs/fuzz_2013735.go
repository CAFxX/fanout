package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 string; F1 bool; F2 uint64 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int8; F1 int }
func genFn0(p0 int16, p1 uint32) int8 {
    p1 += p1
    var v4 bool = (true || true)
    fmt.Println(v4)
    v4 = v4
    p1 = (p1 & 2147483647)
    _ = p0
    _ = p1
    _ = v4
    return 120
}

func main() {
    var v1 int64 = 0
    var v2 int = -9223372036854775808
    var v3 float64 = 1.5
    if true {
        var v5 int64 = (v1 ^ 389)
        for fi6 := 0; fi6 < 9; fi6++ {
            var v7 bool = (100 != 42)
            v8 := []int{-9223372036854775808, (v2 ^ v2), fi6}
            fmt.Println(v5, v2, v7)
            _ = fi6
            _ = v7
            _ = v8
        }
        v9 := []int{-9223372036854775808, (v2 & 127), v2}
        _ = v5
        _ = v9
    }
    var v10 float64 = v3
    v11 := []int{(v2 / (((^(v2))) | 1)), v2, -1, v2}
    for fi12 := 0; fi12 < 11; fi12++ {
        for fi13 := 0; fi13 < 3; fi13++ {
            var v14 bool = true
            switch v2 {
            case 42:
                v15 := GenS1{F0: 2}
                _ = v15
            case -448:
                var v16 bool = (true && v14)
                _ = v16
            default:
                var v17 string = (("" + "foo bar") + ((("hello" + "foo bar") + "") + ("foo bar" + "\n\t")))
                _ = v17
            }
            _ = fi13
            _ = v14
        }
        if len(v11) >= 2 {
            v11 = v11[1:len(v11)]
        }
        if false {
            v18 := GenS1{F0: 1}
            _ = v18
        }
        _ = fi12
    }
    v3 -= v10
    var v19 bool = (false || true)
    for fi20 := 0; fi20 < 16; fi20++ {
        v21 := GenS2{F0: "hello", F1: v19, F2: 18446744073709551615}
        switch ((("" + "") + "\n\t") + "") {
        case "a":
            v3 = v3
        }
        _ = fi20
        _ = v21
    }
    switch "\n\t" {
    case "a":
        v22 := GenS3{F0: 1}
        _ = v22
    case "hello":
        v23 := GenS4{F0: genFn0(32767, 4294967295), F1: (v2 % ((586) | 1))}
        _ = v23
    default:
        for fi24 := 0; fi24 < 6; fi24++ {
            var v25 int = (v2 ^ (fi24 ^ fi24))
            _ = fi24
            _ = v25
        }
    }
    v2 += v2
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v19
    _ = v2
    _ = v3
}
