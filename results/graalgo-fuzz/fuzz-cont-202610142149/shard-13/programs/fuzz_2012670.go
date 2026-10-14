package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64; F2 int }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint8; F1 uint32 }
type GenS4 struct { F0 uint64; F1 uint64 }
func genFn0(p0 int32, p1 uint16) int32 {
    var v6 int64 = -1
    if (!(!true)) {
        var v7 int8 = 100
        switch ("" + ("hello" + (("" + "") + "a"))) {
        case "foo bar":
            v8 := GenS1{F0: (p1 % (((p1 / ((p1) | 1))) | 1)), F1: -847, F2: 869}
            _ = v8
        case "hello":
            v7 -= 0
        }
        var v9 int8 = (-(v7))
        _ = v7
        _ = v9
    }
    v10 := GenS2{F0: 2147483647}
    for fi11 := 0; fi11 < 1; fi11++ {
        if false {
            p1 = p1
            p1 = (p1 ^ p1)
        } else {
            p1 = p1
            fmt.Println(v6, fi11, v6)
            v6 = 1
        }
        if false {
            fmt.Println(p1)
            fmt.Println(p1, fi11, v6)
        } else {
            var v12 uint32 = 2147483647
            _ = v12
        }
        switch fi11 {
        case -2147483648:
            if false {
                for fi13 := 0; fi13 < 13; fi13++ {
                    v14 := []int{fi11, fi13}
                    switch (((("hello" + "\n\t") + (("\n\t" + "") + "hello")) + "") + "hello") {
                    case "foo bar":
                        var v15 uint32 = 196
                        _ = v15
                    default:
                        v16 := GenS2{F0: 0}
                        _ = v16
                    }
                    _ = fi13
                    _ = v14
                }
            }
        case 1:
            v17 := []int{(fi11 >> ((65535) & 63)), fi11}
            _ = v17
        case 630:
            if (!true) {
                p0 -= (p0 | p0)
            }
        default:
            v18 := []int{-2147483648, fi11}
            _ = v18
        }
        _ = fi11
    }
    fmt.Println(v6, p0)
    _ = p0
    _ = p1
    _ = v10
    _ = v6
    return (p0 - (p0 + (p0 & 470)))
}

func main() {
    v1 := make(map[string]uint32)
    var v2 float64 = (0.0 / -2.25)
    v3 := make(map[int]bool)
    v4 := []int{-1, 32767, 9223372036854775807}
    var v5 string = (("" + ("foo bar" + "")) + "foo bar")
    for fi19 := 0; fi19 < 10; fi19++ {
        for fi20 := 0; fi20 < 7; fi20++ {
            v5 = v5
            var v21 uint32 = 1
            fmt.Println(v21, fi19)
            _ = fi20
            _ = v21
        }
        v22 := GenS3{F0: 164, F1: 0}
        _ = fi19
        _ = v22
    }
    var v23 uint = 0
    v1[(v5 + v5)] = 2
    fmt.Println(v23, v2, v2)
    v24 := GenS4{F0: 42, F1: 608}
    v5 = ""
    var v25 uint16 = 976
    v25 -= (v25 - v25)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
