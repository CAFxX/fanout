package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint16; F1 uint }
type GenS3 struct { F0 bool; F1 string; F2 int32 }
type GenS4 struct { F0 uint32 }
func genFn0(p0 int8, p1 int16, p2 uint) string {
    if false {
        var v4 uint64 = 18446744073709551615
        p2 -= 1
        _ = v4
    } else {
        var v5 uint16 = 0
        var v6 int32 = 2147483647
        _ = v5
        _ = v6
    }
    switch 42 {
    case 32767:
        if true {
            var v7 int16 = (p1 * p1)
            fmt.Println(p2, p2)
            _ = v7
        } else {
            p1 -= p1
            if (p1 != (p1 | (p1 ^ p1))) {
                for fi8 := 0; fi8 < 10; fi8++ {
                    p1 -= p1
                    v9 := []int{}
                    if len(v9) > 0 {
                        v9[0] = (-(fi8))
                    }
                    _ = fi8
                    _ = v9
                }
                p1 = p1
            }
            v10 := []int{421, 32768}
            _ = v10
        }
    case 128:
        for fi11 := 0; fi11 < 17; fi11++ {
            p0 -= 127
            p2 = p2
            switch ("foo bar" + ("foo bar" + "a")) {
            case "\n\t":
                p0 = (p0 % ((100) | 1))
            case "foo bar":
                var v12 int32 = 2147483647
                _ = v12
            }
            _ = fi11
        }
    }
    for fi13 := 0; fi13 < 4; fi13++ {
        p2 += (p2 * (p2 + (p2 % ((p2) | 1))))
        _ = fi13
    }
    if true {
        v14 := GenS1{F0: p1}
        var v15 uint16 = 42
        for fi16 := 0; fi16 < 20; fi16++ {
            var v17 float32 = (1.5 - (0.0 / 1e10))
            p1 -= p1
            _ = fi16
            _ = v17
        }
        _ = v14
        _ = v15
    }
    _ = p0
    _ = p1
    _ = p2
    return ""
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[string]bool)
    v3 := make(map[string]int)
    var v18 int32 = 42
    _ = len(v3)
    switch "foo bar" {
    case "hello":
        v19 := []int{136, 128}
        _ = v19
    default:
        v1[563] = 835
    }
    if false {
        v20 := GenS2{F0: 689, F1: 18446744073709551615}
        _ = v20
    } else {
        for fi21 := 0; fi21 < 10; fi21++ {
            var v22 int16 = 1
            _ = fi21
            _ = v22
        }
    }
    var v23 int = -1
    for fi24 := 0; fi24 < 10; fi24++ {
        v25 := GenS3{F0: ((true && false) && true), F1: "hello", F2: 2147483647}
        delete(v2, "foo bar")
        _ = fi24
        _ = v25
    }
    v23 -= v23
    if true {
        var v26 int = v23
        if ((false || ("" != (("foo bar" + "") + "hello"))) && false) {
            v23 -= v23
            for fi27 := 0; fi27 < 1; fi27++ {
                if true {
                    v18 = v18
                    v26 = fi27
                }
                _ = fi27
            }
            v18 += (v18 % ((-980) | 1))
        }
        v28 := GenS2{F0: 32767, F1: 2}
        _ = v26
        _ = v28
    } else {
        fmt.Println(v18, v23)
        v29 := GenS4{F0: 0}
        _ = v29
    }
    v30 := []int{}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v30:", len(v30), v30)
    _ = v1
    _ = v18
    _ = v2
    _ = v23
    _ = v3
    _ = v30
}
