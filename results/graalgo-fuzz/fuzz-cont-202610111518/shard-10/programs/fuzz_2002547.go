package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16; F2 string }
type GenS2 struct { F0 uint32; F1 int8; F2 uint64 }
type GenS3 struct { F0 int32; F1 bool; F2 uint16 }
func genFn0(p0 string) int {
    v5 := GenS1{F0: p0, F1: -13, F2: p0}
    p0 = ("a" + p0)
    var v6 uint = 101
    if true {
        p0 = p0
        p0 = "hello"
        v7 := GenS1{F0: "foo bar", F1: 286, F2: p0}
        _ = v7
    } else {
        var v8 float64 = 0.0
        _ = v8
    }
    _ = p0
    _ = v5
    _ = v6
    return 120
}
func genFn1() int64 {
    v9 := GenS1{F0: "a", F1: 32767, F2: "hello"}
    var v10 int64 = -927
    v11 := GenS1{F0: "hello", F1: 42, F2: ("" + ("\n\t" + "hello"))}
    _ = v10
    _ = v11
    _ = v9
    return v10
}

func main() {
    v1 := make(map[int]bool)
    v2 := []uint32{4294967295, 2147483648, 2147483648}
    var v3 uint16 = 32767
    var v4 int = 127
    var v12 bool = true
    for fi13 := 0; fi13 < 7; fi13++ {
        fmt.Println(v12)
        _ = fi13
    }
    var v14 int32 = 0
    v15 := GenS2{F0: 885, F1: 127, F2: 596}
    v16 := v1[-941]
    var v17 int = 32767
    var v18 int8 = -128
    for fi19 := 0; fi19 < 3; fi19++ {
        _ = len(v2)
        if false {
            for fi20 := 0; fi20 < 19; fi20++ {
                v21 := GenS1{F0: ("a" + ""), F1: 430, F2: ("" + "")}
                if (0 != 2) {
                    _ = len(v1)
                    v17 -= (^(v4))
                } else {
                    var v22 uint32 = 65535
                    var v23 uint64 = 4294967295
                    _ = v22
                    _ = v23
                }
                v24 := GenS3{F0: (-(v14)), F1: (2 >= 62), F2: 32767}
                var v25 uint16 = v3
                _ = fi20
                _ = v21
                _ = v24
                _ = v25
            }
            fmt.Println(v14)
            if (true || v16) {
                var v26 bool = false
                v3 += (v3 & v3)
                _ = v26
            }
        } else {
            v18 -= (^(v18))
            _ = len(v2)
        }
        _ = fi19
    }
    fmt.Println(v4, v17)
    v17 = genFn0(((("" + "") + "\n\t") + ("a" + "a")))
    var v27 int64 = -293
    var v28 uint64 = 777
    _ = len(v1)
    _ = len(v1)
    var v29 string = "a"
    var v30 int16 = -862
    v2 = append(v2, 42)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
}
