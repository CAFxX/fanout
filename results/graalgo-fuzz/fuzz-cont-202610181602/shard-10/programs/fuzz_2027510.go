package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16; F2 int8 }
func genFn0() int8 {
    var v5 int8 = 79
    var v6 int = -1
    for fi7 := 0; fi7 < 14; fi7++ {
        for fi8 := 0; fi8 < 1; fi8++ {
            fmt.Println(fi7)
            _ = fi8
        }
        v5 -= (-(v5))
        _ = fi7
    }
    _ = v5
    _ = v6
    return v5
}
func genFn1(p0 int16, p1 int64, p2 int8) uint8 {
    var v9 float64 = 0.1
    var v10 uint8 = 1
    fmt.Println(p0, v9)
    var v11 float32 = 1.5
    var v12 string = ("hello" + ("a" + ""))
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v12
    _ = v9
    return (v10 + 42)
}

func main() {
    v1 := []string{}
    var v2 uint32 = 4294967295
    var v3 int64 = 789
    var v4 uint64 = 2
    fmt.Println(v2)
    v1 = append(v1, "\n\t")
    var v13 bool = (false && false)
    fmt.Println(v13, v4, v3)
    if false {
        if v13 {
            fmt.Println(v4)
            for fi14 := 0; fi14 < 2; fi14++ {
                if false {
                    v2 += (v2 - v2)
                    for fi15 := 0; fi15 < 12; fi15++ {
                        var v16 float32 = (3.14159 * 3.14159)
                        var v17 int64 = v3
                        _ = fi15
                        _ = v16
                        _ = v17
                    }
                }
                for fi18 := 0; fi18 < 2; fi18++ {
                    fmt.Println(fi14)
                    v4 -= v4
                    if len(v1) > 0 {
                        v1[0] = "foo bar"
                    }
                    _ = fi18
                }
                switch fi14 {
                case 1:
                    if (-461 == 780) {
                        v13 = true
                        v3 += v3
                    } else {
                        v1 = append(v1, "a")
                    }
                case 42:
                    v13 = v13
                default:
                    var v19 uint8 = 127
                    _ = v19
                }
                _ = fi14
            }
            if (true || (genFn0() != -1)) {
                if len(v1) > 0 {
                    v1[0] = ("\n\t" + "foo bar")
                }
                if (!(genFn0() < 92)) {
                    if (true || v13) {
                        v2 = uint32(v4)
                        fmt.Println(v4, v3, v13)
                        v20 := GenS1{F0: 255, F1: 1000, F2: -45}
                        _ = v20
                    } else {
                        v13 = v13
                    }
                    var v21 float64 = 0.0
                    v22 := GenS1{F0: 42, F1: -594, F2: 91}
                    _ = v21
                    _ = v22
                }
                fmt.Println(v3)
            } else {
                v23 := GenS1{F0: genFn1(1, 9223372036854775807, -128), F1: 0, F2: 48}
                v13 = ((true || (true && v13)) && (false && v13))
                for fi24 := 0; fi24 < 10; fi24++ {
                    if len(v1) > 0 {
                        v1[0] = ""
                    }
                    var v25 int16 = -1
                    _ = fi24
                    _ = v25
                }
                _ = v23
            }
        } else {
            fmt.Println(v4)
            v1 = append(v1, "\n\t")
            var v26 int8 = -1
            _ = v26
        }
        var v27 int64 = (v3 + v3)
        _ = len(v1)
        _ = v27
    }
    v28 := GenS1{F0: 2, F1: 672, F2: 42}
    var v29 string = "hello"
    var v30 int16 = -32768
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
}
