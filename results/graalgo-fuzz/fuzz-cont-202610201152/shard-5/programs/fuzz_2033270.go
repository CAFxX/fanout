package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32; F2 int8 }
type GenS2 struct { F0 int64; F1 bool; F2 uint }
func genFn0(p0 int8, p1 uint16) int8 {
    var v6 uint16 = 2
    for fi7 := 0; fi7 < 10; fi7++ {
        v6 -= (p1 & p1)
        if true {
            v8 := []int{}
            v9 := GenS1{F0: "foo bar", F1: 0, F2: (p0 | p0)}
            v6 = 32767
            _ = v8
            _ = v9
        }
        _ = fi7
    }
    switch ((((("foo bar" + ("" + (("\n\t" + "\n\t") + ""))) + (("" + "hello") + "foo bar")) + "\n\t") + "hello") + "") {
    case "":
        p1 = (v6 ^ v6)
    case "foo bar":
        v10 := []int{-440, 1, 4294967295}
        _ = v10
    case "\n\t":
        var v11 uint64 = 497
        _ = v11
    }
    _ = p0
    _ = p1
    _ = v6
    return (^(p0))
}

func main() {
    var v1 int = 1
    var v2 int8 = 100
    v3 := make(map[int]uint32)
    var v4 uint = 255
    var v5 int8 = (v2 ^ v2)
    var v12 int32 = -2147483648
    v13 := GenS2{F0: -966, F1: true, F2: (v4 ^ v4)}
    var v14 string = ("hello" + "hello")
    if (!(2 == 255)) {
        v2 -= v5
        v15 := GenS2{F0: 651, F1: (((float64((-2.25 / 1.0)) / 1.0) <= 0.1) && false), F2: 4294967295}
        _ = v15
    } else {
        v2 -= 127
        var v16 string = v14
        if false {
            var v17 int8 = (v2 * v5)
            v4 = v4
            _ = v17
        }
        _ = v16
    }
    var v18 uint64 = 42
    for fi19 := 0; fi19 < 18; fi19++ {
        if (42 == -9223372036854775808) {
            if true {
                for fi20 := 0; fi20 < 8; fi20++ {
                    var v21 int32 = 559
                    if (v14 > "") {
                        v22 := GenS1{F0: "a", F1: 42, F2: v2}
                        if false {
                            var v23 uint16 = 1
                            _ = v23
                        } else {
                            v14 = v14
                            v3[fi20] = 42
                            v5 = v5
                        }
                        v24 := GenS1{F0: v14, F1: (v21 - 565), F2: (v5 / ((0) | 1))}
                        _ = v22
                        _ = v24
                    }
                    v4 = (v4 + 255)
                    v1 = fi20
                    _ = fi20
                    _ = v21
                }
            }
        }
        _ = fi19
    }
    if ((v18 - v18) == v18) {
        v25 := GenS2{F0: -421, F1: (!true), F2: (^(v4))}
        if false {
            v2 -= 1
            var v26 int8 = (v5 ^ (^(v2)))
            var v27 uint8 = 42
            _ = v26
            _ = v27
        } else {
            for fi28 := 0; fi28 < 8; fi28++ {
                v29 := GenS1{F0: v14, F1: v12, F2: v2}
                v18 = v18
                var v30 uint8 = 1
                _ = fi28
                _ = v29
                _ = v30
            }
        }
        _ = v25
    }
    v2 += int8(v12)
    v2 -= 0
    v1 = v1
    var v31 int8 = -128
    var v32 bool = false
    v14 = ((v14 + (("hello" + ("a" + v14)) + v14)) + "foo bar")
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v18
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
}
