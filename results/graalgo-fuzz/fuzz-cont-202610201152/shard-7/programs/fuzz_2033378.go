package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16; F2 int32 }
type GenS2 struct { F0 int16; F1 uint8 }
func genFn0() uint {
    for fi7 := 0; fi7 < 1; fi7++ {
        if (42 > 850) {
            v8 := []int{fi7, (fi7 ^ fi7)}
            _ = v8
        } else {
            fmt.Println(fi7)
            var v9 uint = 255
            v9 += (v9 | (v9 + (^(v9))))
            _ = v9
        }
        for fi10 := 0; fi10 < 1; fi10++ {
            if ("hello" > ((((("a" + "") + ("foo bar" + "a")) + "\n\t") + (((("a" + "") + "foo bar") + "") + ("hello" + ("" + "hello")))) + "foo bar")) {
                fmt.Println(fi10)
            }
            _ = fi10
        }
        if true {
            var v11 bool = (float32(-2.25) < -2.25)
            var v12 int = (^(fi7))
            _ = v11
            _ = v12
        }
        _ = fi7
    }
    v13 := GenS1{F0: 127, F1: -32768, F2: 1}
    if false {
        var v14 bool = true
        v15 := []int{32768, -2147483648}
        for fi16 := 0; fi16 < 19; fi16++ {
            v14 = v14
            _ = fi16
        }
        _ = v14
        _ = v15
    }
    v17 := GenS1{F0: 1, F1: 42, F2: -457}
    var v18 uint = 4294967295
    _ = v13
    _ = v17
    _ = v18
    return v18
}

func main() {
    v1 := []uint32{42, 852, 2147483647}
    var v2 bool = false
    v3 := make(map[int]string)
    var v4 int32 = -2147483648
    var v5 float32 = 0.0
    var v6 int32 = (v4 >> ((0) & 63))
    if (v6 <= v6) {
        for fi19 := 0; fi19 < 2; fi19++ {
            v5 -= v5
            _ = fi19
        }
    }
    v1 = append(v1, 861)
    var v20 int32 = 0
    switch 9223372036854775807 {
    case 2147483647:
        for fi21 := 0; fi21 < 1; fi21++ {
            v20 = v20
            for fi22 := 0; fi22 < 1; fi22++ {
                switch fi22 {
                case -1:
                    if (184 >= 2) {
                        v5 = (float32((float64((float32(((3.14159 + (1e10 / 1e10)) - 0.0)) * 1e10)) / 0.1)) - 3.14159)
                        var v23 float32 = ((v5 / -2.25) + v5)
                        _ = v23
                    }
                default:
                    var v24 int8 = 1
                    _ = v24
                }
                v25 := GenS1{F0: 42, F1: 1, F2: (v6 - 42)}
                _ = fi22
                _ = v25
            }
            v3[128] = "foo bar"
            _ = fi21
        }
    default:
        var v26 int8 = 42
        _ = v26
    }
    for fi27 := 0; fi27 < 2; fi27++ {
        if v2 {
            for fi28 := 0; fi28 < 13; fi28++ {
                v2 = false
                v6 -= 0
                _ = fi28
            }
        } else {
            v5 += 0.1
        }
        _ = fi27
    }
    switch "" {
    case "\n\t":
        if (v2 || true) {
            _ = len(v1)
            v3[32768] = ""
            v29 := GenS2{F0: 710, F1: 42}
            _ = v29
        } else {
            v2 = (0 > 18446744073709551615)
        }
    }
    v4 = (v6 & (v4 + v6))
    switch -9223372036854775808 {
    case -905:
        var v30 string = "hello"
        _ = v30
    case 127:
        var v31 int = -1
        _ = v31
    case -2147483648:
        fmt.Println(v20, v6)
    default:
        v32 := GenS2{F0: 379, F1: 255}
        _ = v32
    }
    var v33 bool = false
    v5 -= v5
    v34 := GenS1{F0: 128, F1: -139, F2: v4}
    fmt.Println(v6, v20, v5)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
