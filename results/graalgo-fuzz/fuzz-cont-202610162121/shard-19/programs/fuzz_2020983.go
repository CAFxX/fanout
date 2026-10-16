package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint64; F1 string }
func genFn0(p0 bool) uint64 {
    if false {
        fmt.Println(p0)
        for fi5 := 0; fi5 < 3; fi5++ {
            fmt.Println(p0, p0)
            v6 := []int{-795, 32768}
            var v7 int64 = 9223372036854775807
            _ = fi5
            _ = v6
            _ = v7
        }
        p0 = p0
    } else {
        fmt.Println(p0)
    }
    p0 = p0
    p0 = (42 >= 127)
    v8 := []int{0, 42, 32768, 32768}
    var v9 bool = (-2.25 > 1.5)
    _ = p0
    _ = v8
    _ = v9
    return 471
}
func genFn1(p0 uint) int {
    for fi10 := 0; fi10 < 17; fi10++ {
        var v11 bool = (2147483647 <= fi10)
        var v12 int = -882
        _ = fi10
        _ = v11
        _ = v12
    }
    v13 := []int{32768, 32767, 2147483647}
    p0 -= (^(p0))
    _ = p0
    _ = v13
    return 2147483647
}

func main() {
    var v1 bool = (false && (!(2147483647 == 2147483648)))
    v2 := make(map[string]string)
    v3 := []string{("\n\t" + "hello"), ("a" + "\n\t")}
    var v4 bool = ((!(9223372036854775807 >= 0)) || (("" == ("foo bar" + "foo bar")) && (2 >= 1)))
    if v1 {
        v14 := GenS1{F0: 0}
        var v15 string = ("\n\t" + ("" + "hello"))
        _ = v14
        _ = v15
    }
    fmt.Println(v1, v4)
    switch 128 {
    case 2147483647:
        if v1 {
            var v16 int64 = -492
            fmt.Println(v1)
            _ = v16
        } else {
            v1 = (-34 >= 127)
            if (v4 && (100 == -128)) {
                var v17 int8 = -128
                var v18 bool = true
                _ = v17
                _ = v18
            } else {
                if (65535 >= 4294967295) {
                    v4 = (-2.25 >= 1.5)
                    v3 = append(v3, (("foo bar" + "\n\t") + "a"))
                    var v19 uint16 = 1
                    _ = v19
                } else {
                    var v20 int64 = -900
                    var v21 uint64 = 454
                    _ = v20
                    _ = v21
                }
                var v22 bool = ((!(0.1 <= 0.0)) || v4)
                _ = v22
            }
            fmt.Println(v4)
        }
    }
    for fi23 := 0; fi23 < 4; fi23++ {
        v24 := v2["hello"]
        _ = fi23
        _ = v24
    }
    fmt.Println(v1, v1)
    if v4 {
        var v25 int = 4294967295
        _ = v25
    }
    if len(v3) > 0 {
        v3[0] = ("\n\t" + (("a" + "foo bar") + ""))
    }
    switch "a" {
    case "foo bar":
        _ = len(v3)
    case "\n\t":
        v26 := GenS2{F0: 136, F1: ("foo bar" + (("foo bar" + ("a" + "hello")) + ("foo bar" + ("foo bar" + "\n\t"))))}
        _ = v26
    }
    if (-1 == 9223372036854775807) {
        for fi27 := 0; fi27 < 5; fi27++ {
            var v28 uint64 = genFn0(((("foo bar" + "") + "hello") == ("a" + ("\n\t" + ("hello" + "a")))))
            var v29 uint32 = 2147483647
            _ = fi27
            _ = v28
            _ = v29
        }
    } else {
        v30 := v2["hello"]
        var v31 uint64 = 745
        _ = v30
        _ = v31
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
}
