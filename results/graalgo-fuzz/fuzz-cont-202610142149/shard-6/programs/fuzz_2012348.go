package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 bool; F1 int; F2 bool }
func genFn0(p0 uint64, p1 int64) uint16 {
    fmt.Println(p0, p1)
    p1 = (p1 / (((-(p1))) | 1))
    _ = p0
    _ = p1
    return 65535
}
func genFn1() int32 {
    switch "foo bar" {
    case "a":
        var v7 int16 = 1000
        _ = v7
    case "foo bar":
        v8 := []int{32768, 2147483647}
        _ = v8
    case "":
        var v9 uint64 = 18446744073709551615
        _ = v9
    }
    var v10 int = -849
    fmt.Println(v10)
    fmt.Println(v10)
    _ = v10
    return 1
}

func main() {
    v1 := []uint32{0}
    v2 := make(map[int]uint32)
    var v3 int64 = 42
    v4 := []int{}
    v5 := []int{2147483647, 1, 32768, -1}
    v6 := []string{"\n\t"}
    var v11 int64 = v3
    v12 := GenS1{F0: 0}
    var v13 int16 = 821
    _ = len(v2)
    _ = len(v2)
    v14 := GenS2{F0: false, F1: 1, F2: true}
    _ = len(v6)
    if ((false || false) || true) {
        v15 := GenS1{F0: 0}
        v11 -= (v11 - 329)
        _ = v15
    } else {
        switch "\n\t" {
        case "a":
            if true {
                if false {
                    var v16 int64 = (v11 % ((v3) | 1))
                    v11 += v11
                    _ = v16
                }
            }
        case "foo bar":
            if true {
                v2[32767] = 4294967295
            }
        case "hello":
            switch "hello" {
            case "\n\t":
                v17 := GenS1{F0: 1}
                _ = v17
            case "":
                if (!true) {
                    fmt.Println(v11)
                }
            case "hello":
                v3 = (v3 / (((v11 << ((1) & 63))) | 1))
            }
        }
        v13 -= v13
    }
    var v18 float64 = 0.0
    if false {
        fmt.Println(v3, v3, v3)
        v18 -= -2.25
        switch "foo bar" {
        case "a":
            for fi19 := 0; fi19 < 3; fi19++ {
                for fi20 := 0; fi20 < 11; fi20++ {
                    v21 := v2[(^(fi19))]
                    if false {
                        var v22 string = "\n\t"
                        _ = v22
                    }
                    var v23 int = (fi19 ^ (fi20 - -9223372036854775808))
                    fmt.Println(v18, v23, v23)
                    _ = fi20
                    _ = v21
                    _ = v23
                }
                _ = fi19
            }
        default:
            if true {
                v5 = append(v5, -9223372036854775808)
            }
        }
    } else {
        if false {
            var v24 uint = 534
            v3 += v11
            _ = v24
        } else {
            if false {
                for fi25 := 0; fi25 < 11; fi25++ {
                    v26 := GenS1{F0: 65535}
                    _ = fi25
                    _ = v26
                }
                var v27 int64 = v3
                fmt.Println(v3)
                _ = v27
            }
        }
    }
    for fi28 := 0; fi28 < 19; fi28++ {
        v13 += 42
        var v29 uint = 2
        _ = fi28
        _ = v29
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
