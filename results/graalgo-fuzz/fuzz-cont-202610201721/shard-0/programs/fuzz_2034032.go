package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64; F2 uint }
func genFn0(p0 uint16, p1 uint16, p2 bool) int32 {
    p2 = (!("hello" != (("a" + "hello") + ("a" + "hello"))))
    v5 := []int{1, 0}
    var v6 int8 = -1
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return int32(p0)
}
func genFn1() int8 {
    for fi7 := 0; fi7 < 15; fi7++ {
        var v8 uint = 91
        var v9 bool = true
        _ = fi7
        _ = v8
        _ = v9
    }
    v10 := GenS1{F0: 739, F1: 973, F2: 18446744073709551615}
    var v11 string = ("a" + "foo bar")
    switch v11 {
    case "hello":
        for fi12 := 0; fi12 < 11; fi12++ {
            v13 := []int{(fi12 % ((fi12) | 1)), fi12, (fi12 ^ fi12)}
            fmt.Println(v11)
            var v14 float64 = (0.1 + 1e10)
            _ = fi12
            _ = v13
            _ = v14
        }
    case "":
        switch "a" {
        case "\n\t":
            v11 = ""
        case "":
            if true {
                v11 = "hello"
                var v15 int32 = 301
                v16 := []int{}
                _ = v15
                _ = v16
            } else {
                switch 2147483647 {
                case -647:
                    v11 = (v11 + "\n\t")
                case 4294967295:
                    var v17 float32 = float32((float64(-2.25) - (-2.25 - 1.5)))
                    _ = v17
                }
                v18 := GenS1{F0: 8, F1: 695, F2: 738}
                _ = v18
            }
        }
    }
    _ = v10
    _ = v11
    return -30
}

func main() {
    var v1 int64 = 0
    var v2 int16 = -1
    v3 := []string{"", ("\n\t" + "\n\t")}
    var v4 int = 2147483647
    fmt.Println(v4, v4, v1)
    var v19 bool = false
    v20 := GenS1{F0: 65535, F1: 1, F2: 18446744073709551615}
    if (!(!v19)) {
        var v21 int8 = 121
        _ = v21
    }
    for fi22 := 0; fi22 < 6; fi22++ {
        v2 = v2
        _ = fi22
    }
    v23 := GenS1{F0: 2, F1: v1, F2: 42}
    var v24 float64 = 3.14159
    var v25 uint32 = 4294967295
    var v26 int8 = 42
    v24 = v24
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
}
