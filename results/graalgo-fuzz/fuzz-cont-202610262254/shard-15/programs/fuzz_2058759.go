package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool; F2 int8 }
type GenS2 struct { F0 string; F1 uint64 }
type GenS3 struct { F0 bool }
func genFn0(p0 uint64, p1 int32, p2 int32) uint32 {
    fmt.Println(p1)
    v6 := GenS1{F0: 1000, F1: true, F2: 65}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 2
}
func genFn1(p0 uint64, p1 string, p2 uint) uint {
    var v7 float32 = (3.14159 + 1.5)
    var v8 bool = (711 != 0)
    for fi9 := 0; fi9 < 3; fi9++ {
        v10 := GenS2{F0: p1, F1: 727}
        _ = fi9
        _ = v10
    }
    if v8 {
        if ((615 == 1) && true) {
            v11 := GenS2{F0: p1, F1: 838}
            _ = v11
        }
        v12 := GenS2{F0: ((p1 + ("a" + "hello")) + p1), F1: 18446744073709551615}
        _ = v12
    }
    v13 := []int{-2147483648, -2147483648}
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v7
    _ = v8
    return 810
}

func main() {
    v1 := []string{(("hello" + ("hello" + ("a" + ("a" + "\n\t")))) + "foo bar"), ("hello" + ""), "", ""}
    v2 := []uint32{993, 148, 741}
    var v3 uint64 = 42
    v4 := make(map[string]uint32)
    var v5 int64 = 0
    v14 := GenS3{F0: true}
    v3 += 18446744073709551615
    _ = len(v1)
    for fi15 := 0; fi15 < 1; fi15++ {
        v16 := GenS3{F0: false}
        _ = fi15
        _ = v16
    }
    v17 := GenS1{F0: 485, F1: ((!false) && (1 == 1)), F2: -128}
    if false {
        v5 = v5
        for fi18 := 0; fi18 < 6; fi18++ {
            v19 := GenS2{F0: "\n\t", F1: v3}
            for fi20 := 0; fi20 < 19; fi20++ {
                var v21 uint32 = genFn0(42, 42, -581)
                v21 = v21
                v21 += genFn0(v3, 2147483647, 69)
                _ = fi20
                _ = v21
            }
            _ = fi18
            _ = v19
        }
        v3 = (v3 * v3)
    }
    var v22 uint = genFn1(v3, "foo bar", 65535)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v17
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
