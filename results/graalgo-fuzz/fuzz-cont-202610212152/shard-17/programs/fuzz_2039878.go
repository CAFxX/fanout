package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64; F2 uint32 }
type GenS2 struct { F0 uint64; F1 int32; F2 int8 }
type GenS3 struct { F0 uint64; F1 uint32 }
type GenS4 struct { F0 string; F1 uint32; F2 string }
type GenS5 struct { F0 int32; F1 string; F2 int32 }
func genFn0(p0 int, p1 uint64) string {
    switch "a" {
    case "":
        var v7 uint = 834
        _ = v7
    }
    v8 := GenS1{F0: 42, F1: p1, F2: 2147483648}
    if true {
        v9 := GenS2{F0: (p1 & 516), F1: -1, F2: 1}
        if false {
            v10 := []int{(^(p0)), 128, 127}
            v11 := GenS2{F0: p1, F1: 841, F2: 108}
            _ = v10
            _ = v11
        } else {
            v12 := GenS3{F0: (p1 - p1), F1: 65535}
            _ = v12
        }
        _ = v9
    } else {
        switch "hello" {
        case "foo bar":
            for fi13 := 0; fi13 < 6; fi13++ {
                var v14 int8 = 0
                if false {
                    var v15 int8 = v14
                    _ = v15
                } else {
                    v16 := []int{(fi13 + -174), (fi13 % (((-(fi13))) | 1)), p0, int(p1)}
                    _ = v16
                }
                _ = fi13
                _ = v14
            }
        case "":
            for fi17 := 0; fi17 < 9; fi17++ {
                if true {
                    var v18 uint = 957
                    _ = v18
                } else {
                    for fi19 := 0; fi19 < 4; fi19++ {
                        var v20 uint16 = 207
                        _ = fi19
                        _ = v20
                    }
                    var v21 uint64 = 1
                    _ = v21
                }
                _ = fi17
            }
        case "a":
            for fi22 := 0; fi22 < 1; fi22++ {
                fmt.Println(fi22, p1, p1)
                v23 := []int{}
                _ = fi22
                _ = v23
            }
        default:
            p0 = p0
        }
        for fi24 := 0; fi24 < 16; fi24++ {
            v25 := GenS1{F0: uint32(p1), F1: (p1 + p1), F2: 869}
            if false {
                v26 := GenS4{F0: (("" + ("foo bar" + "foo bar")) + "a"), F1: 1, F2: "a"}
                v27 := GenS1{F0: 2147483647, F1: (p1 | 18446744073709551615), F2: 275}
                _ = v26
                _ = v27
            }
            _ = fi24
            _ = v25
        }
    }
    _ = p0
    _ = p1
    _ = v8
    return "foo bar"
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[int]bool)
    var v3 uint64 = 42
    v4 := []uint32{0}
    v5 := []string{}
    var v6 uint64 = v3
    v3 -= v6
    if true {
        v2[-2147483648] = true
    } else {
        v28 := GenS3{F0: v3, F1: 305}
        v29 := GenS5{F0: -1, F1: "foo bar", F2: 2147483647}
        v3 = (v6 * 234)
        _ = v28
        _ = v29
    }
    v6 = (v3 | 0)
    _ = len(v2)
    switch "a" {
    case "foo bar":
        for fi30 := 0; fi30 < 12; fi30++ {
            switch "a" {
            case "a":
                var v31 uint8 = 69
                _ = v31
            case "hello":
                fmt.Println(v6, v3)
            }
            _ = fi30
        }
    case "a":
        v3 = (^(v6))
    default:
        for fi32 := 0; fi32 < 9; fi32++ {
            fmt.Println(v6, v6)
            var v33 string = ("foo bar" + "foo bar")
            var v34 int = fi32
            _ = fi32
            _ = v33
            _ = v34
        }
    }
    delete(v2, 32767)
    v6 += v3
    v6 += v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
