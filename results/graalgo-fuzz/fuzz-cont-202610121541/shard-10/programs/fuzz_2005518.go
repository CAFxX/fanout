package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32; F2 int32 }
type GenS2 struct { F0 int }
func genFn0(p0 uint, p1 string, p2 int32) bool {
    v4 := []int{}
    for fi5 := 0; fi5 < 6; fi5++ {
        var v6 float32 = 3.14159
        v7 := GenS1{F0: 255, F1: 666, F2: 42}
        _ = fi5
        _ = v6
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return (42 == 127)
}
func genFn1(p0 uint32, p1 int32, p2 int) int {
    switch p2 {
    case -2147483648:
        p0 = p0
    default:
        v8 := GenS1{F0: 255, F1: 0, F2: p1}
        _ = v8
    }
    v9 := []int{p2, p2, 128}
    var v10 uint8 = 0
    var v11 uint8 = (v10 & (v10 % (((v10 % ((0) | 1))) | 1)))
    var v12 float64 = 0.0
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v12
    _ = v9
    return -49
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint = 4294967295
    v3 := []int{9223372036854775807, -9223372036854775808, 333, 9223372036854775807}
    v1[("hello" + "foo bar")] = false
    fmt.Println(v2)
    v2 += v2
    v13 := GenS1{F0: 0, F1: 2147483648, F2: 284}
    v1[("a" + "a")] = (!true)
    v3 = append(v3, 145)
    var v14 uint = v2
    v14 -= v2
    v15 := GenS2{F0: 42}
    if false {
        var v16 uint16 = 1
        for fi17 := 0; fi17 < 12; fi17++ {
            var v18 uint64 = 18446744073709551615
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            switch (fi17 >> ((v14) & 63)) {
            case -669:
                var v19 uint = uint(v18)
                _ = v19
            case 2147483647:
                for fi20 := 0; fi20 < 11; fi20++ {
                    if len(v3) > 0 {
                        v3[0] = fi17
                    }
                    _ = fi20
                }
            }
            _ = fi17
            _ = v18
        }
        _ = v16
    }
    v2 = v14
    switch -9223372036854775808 {
    case 128:
        for fi21 := 0; fi21 < 8; fi21++ {
            var v22 uint8 = 255
            _ = fi21
            _ = v22
        }
    case 947:
        v2 = 0
    }
    var v23 string = "hello"
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v23
    _ = v3
}
