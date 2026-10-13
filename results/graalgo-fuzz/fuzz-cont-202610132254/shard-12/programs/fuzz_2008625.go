package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool }
type GenS2 struct { F0 int64; F1 uint8 }
func genFn0(p0 uint16) bool {
    if true {
        v6 := GenS1{F0: 4294967295, F1: (1000 != -981)}
        _ = v6
    }
    switch "" {
    case "hello":
        p0 = (p0 * (^(p0)))
    case "\n\t":
        v7 := []int{127, 2147483647, 589}
        _ = v7
    }
    p0 -= (p0 - (p0 | (p0 << ((1) & 63))))
    _ = p0
    return false
}
func genFn1(p0 uint8, p1 int8) bool {
    v8 := GenS1{F0: 18446744073709551615, F1: (15 > 805)}
    fmt.Println(p1)
    p0 -= (^(p0))
    _ = p0
    _ = p1
    _ = v8
    return (!false)
}

func main() {
    v1 := make(map[int]int)
    var v2 uint32 = 65535
    var v3 uint8 = 127
    var v4 uint32 = (v2 ^ v2)
    var v5 float64 = ((((1.5 / (0.1 * 1.5)) * -2.25) - 0.1) * (3.14159 - 3.14159))
    v3 -= v3
    v9 := []int{-1, 149, -2147483648, 1}
    switch -508 {
    case 2147483647:
        var v10 string = ""
        _ = v10
    default:
        var v11 uint16 = 2
        _ = v11
    }
    for fi12 := 0; fi12 < 19; fi12++ {
        v4 = v4
        v13 := GenS1{F0: 42, F1: (false || (fi12 < (fi12 - 2147483647)))}
        _ = fi12
        _ = v13
    }
    v14 := GenS1{F0: 856, F1: false}
    var v15 uint8 = 255
    var v16 uint8 = v15
    if ((!false) && false) {
        if (0 <= 2) {
            var v17 int64 = -577
            v2 = v2
            if ((v3 >> ((255) & 63)) != 127) {
                var v18 int = 183
                v5 = (v5 + 0.0)
                _ = v18
            } else {
                if ("a" == (("hello" + (("" + ("hello" + (("" + ((("" + "") + "foo bar") + "\n\t")) + ("hello" + "hello")))) + "foo bar")) + "a")) {
                    v3 += v15
                    var v19 string = (((((("foo bar" + "foo bar") + "") + "foo bar") + "\n\t") + (("\n\t" + "hello") + ((("hello" + "hello") + "hello") + ""))) + "hello")
                    v16 = 1
                    _ = v19
                } else {
                    fmt.Println(v5, v3)
                }
                fmt.Println(v4, v17)
                if false {
                    v20 := GenS2{F0: v17, F1: v3}
                    if true {
                        v4 -= (v4 & (v2 - v2))
                        var v21 float32 = (0.1 - -2.25)
                        _ = v21
                    } else {
                        if (42 == 1) {
                            var v22 uint = 65535
                            _ = v22
                        } else {
                            v15 -= (v3 / (((^(v16))) | 1))
                        }
                    }
                    _ = v20
                } else {
                    if false {
                        if true {
                            v4 = 65535
                            fmt.Println(v4, v16)
                        }
                    } else {
                        fmt.Println(v4, v2, v4)
                        v23 := v1[-9223372036854775808]
                        _ = v23
                    }
                    var v24 float32 = float32(v5)
                    _ = v24
                }
            }
            _ = v17
        }
    } else {
        var v25 string = "foo bar"
        if (9223372036854775807 == 42) {
            v15 += v15
            fmt.Println(v15, v3, v4)
            _ = len(v9)
        } else {
            v9 = append(v9, -328)
        }
        _ = v25
    }
    v9 = append(v9, 9223372036854775807)
    var v26 bool = true
    switch "" {
    case "a":
        v2 += 2
    case "hello":
        var v27 int16 = 32767
        _ = v27
    case "\n\t":
        for fi28 := 0; fi28 < 5; fi28++ {
            for fi29 := 0; fi29 < 20; fi29++ {
                v16 -= (v3 | (v3 / ((v3) | 1)))
                var v30 uint32 = v4
                _ = fi29
                _ = v30
            }
            _ = fi28
        }
    }
    var v31 int64 = 9223372036854775807
    v3 -= 2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v26
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v9
}
