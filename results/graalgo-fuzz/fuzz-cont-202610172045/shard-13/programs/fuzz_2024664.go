package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8; F2 int }
func genFn0(p0 uint, p1 int16, p2 bool) string {
    v7 := []int{701, -1}
    fmt.Println(p0)
    v8 := GenS1{F0: (p0 | 18446744073709551615), F1: 42, F2: 503}
    p0 += 866
    for fi9 := 0; fi9 < 3; fi9++ {
        for fi10 := 0; fi10 < 8; fi10++ {
            _ = len(v7)
            p2 = p2
            _ = fi10
        }
        _ = fi9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return ("a" + "a")
}
func genFn1(p0 int64, p1 int16) int8 {
    for fi11 := 0; fi11 < 3; fi11++ {
        switch "foo bar" {
        case "foo bar":
            var v12 uint16 = 2
            _ = v12
        }
        var v13 uint = 42
        _ = fi11
        _ = v13
    }
    var v14 uint32 = 42
    v15 := []int{}
    if true {
        p1 += (p1 / ((p1) | 1))
        if (!(1e10 == 0.1)) {
            var v16 int16 = (p1 / ((p1) | 1))
            v17 := GenS1{F0: 0, F1: 0, F2: -58}
            _ = v16
            _ = v17
        }
    }
    if len(v15) >= 2 {
        v15 = v15[1:len(v15)]
    }
    _ = p0
    _ = p1
    _ = v14
    _ = v15
    return 100
}

func main() {
    var v1 uint8 = 183
    v2 := make(map[string]bool)
    v3 := []string{("\n\t" + (("" + "foo bar") + "\n\t")), ("\n\t" + ("\n\t" + "\n\t"))}
    var v4 uint32 = 413
    var v5 int8 = 127
    var v6 uint16 = 1
    var v18 uint64 = 4294967295
    fmt.Println(v4, v1, v18)
    v4 = (^(v4))
    fmt.Println(v4, v6, v1)
    _ = len(v2)
    v5 -= v5
    fmt.Println(v4)
    if len(v3) > 0 {
        v3[0] = ("foo bar" + ((((("a" + "") + "") + "a") + ("foo bar" + ("a" + "\n\t"))) + "hello"))
    }
    v3 = append(v3, "")
    var v19 int8 = 100
    if false {
        fmt.Println(v1, v4)
        var v20 int = 2147483647
        switch "foo bar" {
        case "a":
            delete(v2, "")
        }
        _ = v20
    } else {
        switch "" {
        case "hello":
            v5 = 100
        case "a":
            var v21 int64 = 42
            _ = v21
        }
        var v22 float64 = ((((1e10 + 0.1) / 0.1) / 1.5) - ((1e10 / (-2.25 + -2.25)) / 1.5))
        _ = v22
    }
    v1 -= v1
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
