package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16; F2 int }
func genFn0(p0 string, p1 int8) uint32 {
    var v7 int16 = 1000
    switch ((p0 + p0) + p0) {
    case "a":
        v7 = -32768
    case "hello":
        var v8 float64 = 1.5
        _ = v8
    }
    v9 := GenS1{F0: 4294967295, F1: -463, F2: -865}
    _ = p0
    _ = p1
    _ = v7
    _ = v9
    return 248
}

func main() {
    v1 := make(map[int]int)
    var v2 int16 = 32767
    v3 := make(map[string]string)
    var v4 int16 = (v2 | (v2 << ((1) & 63)))
    var v5 int64 = 0
    var v6 int32 = 2147483647
    v10 := []int{}
    v11 := GenS1{F0: -709, F1: v2, F2: -830}
    if (392 != 400) {
        var v12 uint32 = 0
        v5 = v5
        v12 -= v12
        _ = v12
    }
    var v13 uint64 = 1
    switch "hello" {
    case "":
        var v14 int32 = (v6 & v6)
        _ = v14
    case "\n\t":
        var v15 int16 = (v4 * v4)
        _ = v15
    default:
        switch "foo bar" {
        case "a":
            v4 = v2
        case "foo bar":
            fmt.Println(v4)
        }
    }
    for fi16 := 0; fi16 < 15; fi16++ {
        var v17 uint32 = 453
        if false {
            for fi18 := 0; fi18 < 10; fi18++ {
                if len(v10) > 0 {
                    v10[0] = fi16
                }
                v19 := GenS1{F0: -1, F1: 1000, F2: -2147483648}
                _ = fi18
                _ = v19
            }
            if len(v10) > 0 {
                v10[0] = -2147483648
            }
            var v20 int8 = -1
            _ = v20
        }
        _ = fi16
        _ = v17
    }
    v6 += v6
    for fi21 := 0; fi21 < 3; fi21++ {
        v22 := GenS1{F0: (^(fi21)), F1: (v4 >> ((65535) & 63)), F2: -2147483648}
        var v23 string = ((("hello" + ("a" + ("" + (("a" + "a") + (("a" + (("a" + "foo bar") + ("a" + "hello"))) + ("hello" + "hello")))))) + "\n\t") + "hello")
        fmt.Println(v4, v13, v6)
        _ = fi21
        _ = v22
        _ = v23
    }
    v1[-1] = 1
    if true {
        var v24 int8 = 42
        var v25 string = "foo bar"
        _ = v24
        _ = v25
    } else {
        for fi26 := 0; fi26 < 9; fi26++ {
            for fi27 := 0; fi27 < 19; fi27++ {
                var v28 int = fi26
                var v29 int16 = 1000
                if false {
                    var v30 int32 = 0
                    fmt.Println(v30, fi27)
                    _ = v30
                }
                _ = fi27
                _ = v28
                _ = v29
            }
            v5 = 1
            _ = fi26
        }
        fmt.Println(v2, v13)
        for fi31 := 0; fi31 < 11; fi31++ {
            v10 = append(v10, 1)
            _ = fi31
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
