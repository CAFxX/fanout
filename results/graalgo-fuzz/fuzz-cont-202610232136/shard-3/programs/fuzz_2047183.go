package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint }
type GenS2 struct { F0 string; F1 uint64 }
func genFn0(p0 int16, p1 uint) string {
    p1 += (p1 << (((p1 << (((^(p1))) & 63))) & 63))
    switch 723 {
    case 127:
        v6 := []int{0, 4294967295, int(p1)}
        _ = v6
    default:
        var v7 int64 = 1
        _ = v7
    }
    v8 := []int{2147483647}
    switch 0 {
    case 32768:
        if (0 > 0) {
            for fi9 := 0; fi9 < 9; fi9++ {
                if len(v8) >= 2 {
                    v8 = v8[1:len(v8)]
                }
                var v10 uint64 = 0
                fmt.Println(fi9, p1, p1)
                _ = fi9
                _ = v10
            }
            p0 += p0
        } else {
            p1 = (^(p1))
        }
    case 0:
        for fi11 := 0; fi11 < 10; fi11++ {
            var v12 uint16 = 0
            var v13 int32 = 42
            _ = fi11
            _ = v12
            _ = v13
        }
    }
    _ = p0
    _ = p1
    _ = v8
    return "a"
}
func genFn1(p0 int) int {
    p0 = (p0 >> ((255) & 63))
    if true {
        v14 := []int{(p0 + 32768), p0, (p0 << ((0) & 63))}
        _ = v14
    }
    v15 := []int{p0, (p0 / ((128) | 1))}
    switch ("a" + ((("hello" + ("" + "hello")) + "\n\t") + "foo bar")) {
    case "hello":
        for fi16 := 0; fi16 < 5; fi16++ {
            p0 = 42
            if false {
                p0 -= p0
                v15 = append(v15, (fi16 << ((1) & 63)))
                p0 = (^(p0))
            } else {
                if true {
                    var v17 int64 = -642
                    fmt.Println(fi16, v17, fi16)
                    _ = v17
                }
                _ = len(v15)
            }
            v15 = append(v15, fi16)
            _ = fi16
        }
    case "a":
        p0 += p0
    }
    v18 := GenS1{F0: p0, F1: 42}
    _ = p0
    _ = v15
    _ = v18
    return (-(p0))
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int32 = -555
    var v3 bool = (-1 <= -49)
    var v4 uint32 = 4294967295
    var v5 float64 = 0.0
    v3 = v3
    var v19 int = 9223372036854775807
    v20 := GenS1{F0: 127, F1: 855}
    v21 := GenS2{F0: "\n\t", F1: 0}
    var v22 float64 = (1.5 / v5)
    fmt.Println(v2)
    switch (-(v19)) {
    case 9223372036854775807:
        v5 = -2.25
    }
    v22 = v5
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
