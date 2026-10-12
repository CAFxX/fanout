package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string; F2 int8 }
type GenS2 struct { F0 int64 }
func genFn0(p0 string) uint16 {
    var v6 float32 = 1e10
    var v7 int16 = 32767
    v8 := GenS1{F0: 127, F1: ("foo bar" + p0), F2: 32}
    v9 := GenS1{F0: -128, F1: p0, F2: -128}
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 42
}
func genFn1(p0 uint32) int {
    v10 := []int{-1, 9223372036854775807, 0, -639}
    var v11 int16 = -32768
    _ = p0
    _ = v10
    _ = v11
    return -561
}

func main() {
    var v1 uint32 = 4294967295
    v2 := make(map[string]int)
    var v3 uint = 574
    var v4 int16 = 0
    var v5 uint32 = v1
    for fi12 := 0; fi12 < 5; fi12++ {
        switch int(v5) {
        case -829:
            if (fi12 >= fi12) {
                var v13 uint16 = 785
                _ = v13
            }
        case -625:
            switch (-(fi12)) {
            case 1:
                var v14 float64 = 1e10
                _ = v14
            }
        case 32768:
            _ = len(v2)
        default:
            for fi15 := 0; fi15 < 3; fi15++ {
                v16 := GenS2{F0: 9223372036854775807}
                _ = len(v2)
                var v17 uint8 = 42
                _ = fi15
                _ = v16
                _ = v17
            }
        }
        _ = fi12
    }
    v18 := GenS1{F0: -57, F1: ("\n\t" + "hello"), F2: 100}
    fmt.Println(v1)
    v19 := GenS1{F0: 1, F1: ((("a" + "a") + "") + "foo bar"), F2: 100}
    v2["foo bar"] = genFn1(424)
    v1 = v1
    v2[("foo bar" + (("\n\t" + ("hello" + (("" + (("hello" + (("foo bar" + "") + ("foo bar" + "hello"))) + ("hello" + (("\n\t" + "foo bar") + "foo bar")))) + ""))) + "\n\t"))] = 222
    switch 127 {
    case 455:
        var v20 int64 = -73
        _ = v20
    case 42:
        if (!(!(959 != 349))) {
            v5 = (v1 / ((4294967295) | 1))
        }
    default:
        v2[((("a" + ((("hello" + "\n\t") + "\n\t") + ("" + "a"))) + "a") + ("hello" + "\n\t"))] = 2147483647
    }
    v5 = (v1 & (v1 & v5))
    v21 := []int{0, -567, 914}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
