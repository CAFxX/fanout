package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int8; F1 string; F2 bool }
type GenS3 struct { F0 int }
func genFn0(p0 int16, p1 string, p2 int) uint16 {
    p2 -= (p2 / ((-2147483648) | 1))
    p1 = p1
    if (0.1 != (float64(1e10) / 1e10)) {
        var v7 float64 = (-2.25 - -2.25)
        v8 := GenS1{F0: 245}
        v9 := []int{p2, p2, 127}
        _ = v7
        _ = v8
        _ = v9
    }
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = p2
    return 0
}
func genFn1(p0 int64) int64 {
    p0 += (p0 * (p0 & p0))
    v10 := []int{32768, -281, -31, 578}
    var v11 int64 = p0
    var v12 int64 = p0
    var v13 float32 = 0.0
    _ = p0
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    return v12
}

func main() {
    v1 := make(map[string]int)
    var v2 bool = false
    var v3 uint64 = 0
    var v4 uint = 4294967295
    v5 := []int{-722}
    v6 := []int{}
    v4 -= 358
    var v14 float32 = 0.1
    for fi15 := 0; fi15 < 20; fi15++ {
        var v16 uint32 = 2147483648
        v2 = v2
        var v17 uint64 = v3
        _ = fi15
        _ = v16
        _ = v17
    }
    var v18 int64 = 9223372036854775807
    v1["\n\t"] = 127
    v1[((((("" + ("\n\t" + ((((((("foo bar" + "") + ("\n\t" + ("\n\t" + ((("\n\t" + "") + ("hello" + "a")) + (("foo bar" + "foo bar") + ("foo bar" + "foo bar")))))) + "") + "\n\t") + "hello") + ("\n\t" + ("a" + "hello"))) + "hello"))) + "foo bar") + "hello") + "hello") + ("hello" + "foo bar"))] = 32768
    v19 := GenS1{F0: v4}
    v20 := GenS2{F0: 1, F1: "foo bar", F2: (("foo bar" >= (("\n\t" + ("" + "a")) + (("hello" + "") + "hello"))) && ((!(true && v2)) || true))}
    v2 = (1.5 < 0.1)
    switch int(v4) {
    case -1:
        v18 -= (v18 % (((v18 - (v18 ^ (v18 - 536)))) | 1))
    default:
        v21 := GenS3{F0: 42}
        _ = v21
    }
    v14 = 0.0
    if len(v5) > 0 {
        v5[0] = 0
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v14
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
