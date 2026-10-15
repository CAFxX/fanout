package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint }
type GenS2 struct { F0 int8 }
func genFn0() int16 {
    var v7 uint16 = 32767
    var v8 int32 = 1
    v8 += 0
    _ = v7
    _ = v8
    return -1
}
func genFn1(p0 bool) int {
    switch 32768 {
    case -763:
        fmt.Println(p0)
    case -1:
        fmt.Println(p0)
    case 128:
        var v9 uint8 = 128
        _ = v9
    default:
        var v10 uint = 18446744073709551615
        _ = v10
    }
    if (2147483647 > 842) {
        p0 = true
    } else {
        p0 = (!true)
    }
    _ = p0
    return 9223372036854775807
}

func main() {
    var v1 int32 = 905
    var v2 uint8 = 79
    var v3 string = (("hello" + "\n\t") + "foo bar")
    var v4 int = -164
    v5 := make(map[uint32]int)
    var v6 bool = false
    v1 = (v1 ^ (^(v1)))
    v11 := GenS1{F0: ((!false) && false), F1: 1}
    for fi12 := 0; fi12 < 8; fi12++ {
        if true {
            v1 = v1
            fmt.Println(v6, v2, v6)
            v13 := []int{genFn1((fi12 <= v4)), v4, -460, fi12}
            _ = v13
        } else {
            for fi14 := 0; fi14 < 12; fi14++ {
                v15 := GenS1{F0: ((0 == 42) || v6), F1: 42}
                _ = fi14
                _ = v15
            }
            v4 = (fi12 + fi12)
            var v16 uint32 = 554
            _ = v16
        }
        _ = fi12
    }
    v17 := GenS1{F0: v6, F1: 1}
    v18 := []int{genFn1(v6), (^(v4))}
    fmt.Println(v4)
    v19 := GenS1{F0: v6, F1: 2}
    if true {
        var v20 uint = 0
        _ = v20
    } else {
        for fi21 := 0; fi21 < 19; fi21++ {
            switch v3 {
            case "\n\t":
                if (!(42 <= 42)) {
                    var v22 uint = 1
                    if v6 {
                        for fi23 := 0; fi23 < 7; fi23++ {
                            v24 := GenS1{F0: v6, F1: (v22 % ((1) | 1))}
                            _ = len(v5)
                            _ = fi23
                            _ = v24
                        }
                        v18 = append(v18, v4)
                    }
                    fmt.Println(v6, v22)
                    _ = v22
                }
            case "":
                switch fi21 {
                case -501:
                    delete(v5, 2147483647)
                case 0:
                    var v25 int8 = 46
                    _ = v25
                case -644:
                    switch (v3 + ((v3 + (v3 + v3)) + v3)) {
                    case "hello":
                        fmt.Println(fi21, fi21)
                    default:
                        v6 = true
                    }
                }
            case "a":
                if len(v18) > 0 {
                    v18[0] = genFn1(false)
                }
            default:
                v2 = (v2 - (v2 ^ v2))
            }
            _ = fi21
        }
        v6 = v6
    }
    var v26 int16 = 91
    var v27 uint16 = 255
    v4 = -2147483648
    v18 = append(v18, v4)
    if true {
        for fi28 := 0; fi28 < 17; fi28++ {
            fmt.Println(v3, v6, fi28)
            var v29 float64 = 0.0
            _ = fi28
            _ = v29
        }
    } else {
        var v30 int8 = int8(v26)
        _ = v30
    }
    var v31 uint32 = 2147483648
    v32 := GenS2{F0: 0}
    var v33 int32 = v1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
