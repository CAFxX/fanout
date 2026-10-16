package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint16; F1 int8 }
type GenS3 struct { F0 uint; F1 int8; F2 bool }
type GenS4 struct { F0 bool; F1 uint; F2 uint8 }
type GenS5 struct { F0 int }
func genFn0(p0 int, p1 int) uint {
    v6 := []int{}
    if (255 <= 1) {
        if true {
            var v7 int16 = -32768
            if len(v6) > 0 {
                v6[0] = (^(p0))
            }
            if (-2147483648 == -608) {
                if false {
                    var v8 int8 = -122
                    v9 := GenS1{F0: v7}
                    _ = v8
                    _ = v9
                } else {
                    for fi10 := 0; fi10 < 7; fi10++ {
                        var v11 int = p1
                        fmt.Println(v11, v11, p0)
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                        _ = fi10
                        _ = v11
                    }
                    v12 := GenS1{F0: (v7 - (v7 & (^(v7))))}
                    _ = v12
                }
                v13 := GenS2{F0: 47, F1: 11}
                _ = v13
            } else {
                if ((!(!((((1e10 * (float64((1.5 / 0.1)) + float64(0.1))) / 0.1) * 0.0) == (0.1 + -2.25)))) || true) {
                    v14 := GenS3{F0: 18446744073709551615, F1: -128, F2: true}
                    var v15 string = ("a" + "")
                    _ = v14
                    _ = v15
                } else {
                    switch "" {
                    case "hello":
                        p0 = p1
                    default:
                        var v16 int = (p1 ^ -1)
                        _ = v16
                    }
                }
                if false {
                    var v17 uint64 = 2
                    if false {
                        p1 -= (p0 - 919)
                    } else {
                        p0 = p1
                    }
                    for fi18 := 0; fi18 < 6; fi18++ {
                        v6 = append(v6, (p1 % ((-1) | 1)))
                        _ = fi18
                    }
                    _ = v17
                }
            }
            _ = v7
        }
        fmt.Println(p0, p0)
    } else {
        if true {
            v19 := GenS4{F0: true, F1: 255, F2: 248}
            fmt.Println(p0, p1)
            _ = v19
        }
        fmt.Println(p1)
    }
    if false {
        var v20 float32 = ((((float32(0.1) - ((0.0 + float32(0.1)) - float32((((1e10 * 1e10) + 1.5) * (1e10 - 1e10))))) * 0.0) - 0.1) + 0.1)
        _ = v20
    } else {
        if len(v6) > 0 {
            v6[0] = p0
        }
        var v21 int64 = 463
        _ = v21
    }
    for fi22 := 0; fi22 < 1; fi22++ {
        var v23 int32 = 2147483647
        _ = fi22
        _ = v23
    }
    _ = p0
    _ = p1
    _ = v6
    return 18446744073709551615
}
func genFn1(p0 int, p1 uint64, p2 uint32) bool {
    switch ((("foo bar" + "hello") + ("\n\t" + "hello")) + "foo bar") {
    case "a":
        var v24 int16 = 1
        _ = v24
    }
    p1 = p1
    var v25 uint = 65535
    var v26 float32 = (-2.25 * 0.1)
    v25 += v25
    _ = p0
    _ = p1
    _ = p2
    _ = v25
    _ = v26
    return true
}

func main() {
    var v1 bool = false
    v2 := make(map[string]bool)
    var v3 uint32 = 2147483647
    var v4 int32 = -1
    var v5 int32 = (v4 ^ 303)
    switch "hello" {
    case "foo bar":
        v2["foo bar"] = v1
    case "":
        var v27 float32 = (-2.25 * 3.14159)
        _ = v27
    case "\n\t":
        v28 := GenS5{F0: -405}
        _ = v28
    default:
        v29 := []int{834, 4294967295}
        _ = v29
    }
    v30 := GenS2{F0: 0, F1: 70}
    var v31 int16 = -633
    for fi32 := 0; fi32 < 16; fi32++ {
        fmt.Println(v3)
        _ = fi32
    }
    v31 = (v31 - -32768)
    v5 = (v5 - -2147483648)
    _ = len(v2)
    fmt.Println(v3)
    v33 := GenS2{F0: 65535, F1: 69}
    v34 := v2[("\n\t" + ((("a" + "\n\t") + ((("" + "") + ("\n\t" + "a")) + "\n\t")) + ("a" + "")))]
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v33
    _ = v34
    _ = v4
    _ = v5
}
