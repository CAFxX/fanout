package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16; F2 int }
type GenS2 struct { F0 uint16; F1 string; F2 int16 }
func genFn0(p0 uint32, p1 bool, p2 string) int16 {
    if p1 {
        v6 := GenS1{F0: (p0 >> ((1) & 63)), F1: 1000, F2: -325}
        _ = v6
    } else {
        if p1 {
            for fi7 := 0; fi7 < 16; fi7++ {
                if p1 {
                    switch 42 {
                    case -9223372036854775808:
                        p0 = 826
                    case 32767:
                        var v8 uint = 0
                        _ = v8
                    case 4294967295:
                        v9 := []int{}
                        _ = v9
                    }
                }
                _ = fi7
            }
        } else {
            var v10 int8 = -10
            v11 := GenS2{F0: 32767, F1: "a", F2: 32767}
            p2 = "a"
            _ = v10
            _ = v11
        }
    }
    if p1 {
        v12 := []int{}
        _ = v12
    } else {
        v13 := []int{0, 465, 42, 32767}
        _ = v13
    }
    var v14 float64 = -2.25
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    return 34
}
func genFn1() bool {
    var v15 int32 = -1
    if false {
        var v16 int64 = -405
        if false {
            v17 := []int{-220, 32767, -333}
            fmt.Println(v15, v16)
            _ = v17
        } else {
            fmt.Println(v16, v15)
        }
        _ = v16
    } else {
        v15 = v15
    }
    var v18 string = "hello"
    fmt.Println(v18)
    _ = v15
    _ = v18
    return (1 <= -1)
}

func main() {
    var v1 float64 = 3.14159
    v2 := []string{"\n\t", "\n\t", ("a" + ("" + "\n\t")), (("foo bar" + "hello") + ("\n\t" + ""))}
    v3 := make(map[string]bool)
    var v4 float32 = -2.25
    v5 := make(map[int]bool)
    v4 = v4
    if true {
        if true {
            if (!((-826 == 529) || (42 <= 1))) {
                v4 -= v4
                for fi19 := 0; fi19 < 11; fi19++ {
                    v4 -= v4
                    var v20 float32 = v4
                    v4 -= float32(3.14159)
                    _ = fi19
                    _ = v20
                }
            }
            v21 := v3["foo bar"]
            _ = v21
        }
    }
    v1 = -2.25
    var v22 uint8 = 186
    _ = len(v3)
    if (9223372036854775807 < 846) {
        _ = len(v2)
        v1 += (v1 + v1)
        var v23 uint16 = 32767
        _ = v23
    }
    var v24 string = ("" + ((("hello" + "a") + "") + "foo bar"))
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v25 int16 = 67
    v4 += (((v4 - v4) - (1e10 + (1.5 + v4))) + ((-2.25 - (0.0 * (v4 * v4))) + 0.0))
    v22 = 252
    v24 = "foo bar"
    if (!true) {
        var v26 bool = false
        if v26 {
            var v27 uint16 = 0
            _ = v27
        }
        var v28 uint16 = 32767
        _ = v26
        _ = v28
    } else {
        v29 := GenS2{F0: 32767, F1: v24, F2: (v25 << ((0) & 63))}
        fmt.Println(v4, v25)
        _ = v29
    }
    v4 -= float32(v1)
    v24 = (v24 + v24)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v22
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
