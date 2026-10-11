package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64; F2 uint16 }
type GenS2 struct { F0 int16; F1 bool }
func genFn0(p0 uint8, p1 uint) int64 {
    p0 -= 2
    v5 := []int{-9223372036854775808, int(p1), 406}
    v6 := GenS1{F0: ("a" + ("\n\t" + ("hello" + "a"))), F1: -9223372036854775808, F2: 65535}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 163
}
func genFn1(p0 uint, p1 uint32) string {
    for fi7 := 0; fi7 < 16; fi7++ {
        var v8 uint64 = 18446744073709551615
        _ = fi7
        _ = v8
    }
    v9 := []int{-620, 0}
    if (p0 <= 247) {
        v10 := GenS1{F0: "", F1: 260, F2: 1}
        _ = v10
    }
    var v11 float32 = -2.25
    var v12 float64 = -2.25
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    _ = v9
    return ("a" + "foo bar")
}

func main() {
    var v1 uint8 = 1
    v2 := []uint32{2147483647, 1, 65535}
    v3 := make(map[uint32]bool)
    v4 := make(map[int]uint32)
    if (90 != 2) {
        v1 += 42
    }
    var v13 int32 = -1
    v13 -= v13
    v13 -= v13
    if false {
        if true {
            v3[2147483648] = (("hello" + (("a" + "a") + "foo bar")) > ((("a" + "foo bar") + ("foo bar" + "")) + "a"))
        } else {
            var v14 uint16 = 32767
            v14 += v14
            _ = v14
        }
    } else {
        switch "\n\t" {
        case "a":
            if false {
                _ = len(v3)
                for fi15 := 0; fi15 < 10; fi15++ {
                    for fi16 := 0; fi16 < 20; fi16++ {
                        v13 -= v13
                        var v17 int16 = 1000
                        _ = fi16
                        _ = v17
                    }
                    var v18 bool = (-1 != 0)
                    _ = fi15
                    _ = v18
                }
                switch -1 {
                case 2147483647:
                    v19 := GenS1{F0: (("foo bar" + "foo bar") + "a"), F1: genFn0(v1, 1), F2: 996}
                    _ = v19
                default:
                    v13 = (v13 | (^(v13)))
                }
            }
        case "hello":
            var v20 string = "a"
            _ = v20
        case "foo bar":
            var v21 int32 = v13
            _ = v21
        }
        for fi22 := 0; fi22 < 6; fi22++ {
            if (-1 < -542) {
                switch "a" {
                case "foo bar":
                    v1 = v1
                case "hello":
                    v23 := GenS2{F0: 1, F1: false}
                    _ = v23
                case "\n\t":
                    for fi24 := 0; fi24 < 20; fi24++ {
                        _ = len(v4)
                        v1 += (v1 ^ (v1 >> ((204) & 63)))
                        v25 := GenS2{F0: -32768, F1: (42 <= 2)}
                        _ = fi24
                        _ = v25
                    }
                }
                v1 -= (^(v1))
            }
            _ = fi22
        }
        fmt.Println(v13, v1)
    }
    v26 := GenS2{F0: 255, F1: true}
    fmt.Println(v1, v13)
    var v27 float64 = float64(1.5)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v13
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
}
