package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool }
type GenS2 struct { F0 int; F1 bool; F2 uint64 }
type GenS3 struct { F0 uint }
func genFn0(p0 int8) uint {
    v5 := GenS1{F0: 1, F1: true}
    var v6 int = 0
    for fi7 := 0; fi7 < 4; fi7++ {
        v8 := GenS2{F0: v6, F1: false, F2: 18446744073709551615}
        _ = fi7
        _ = v8
    }
    p0 = 127
    _ = p0
    _ = v5
    _ = v6
    return 42
}
func genFn1(p0 uint8, p1 uint64, p2 uint32) bool {
    var v9 uint64 = p1
    p1 = (v9 & p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return (42 == 1)
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[string]bool)
    var v3 float64 = 1.5
    var v4 int64 = -9223372036854775808
    fmt.Println(v3)
    v10 := []int{1, -2147483648}
    var v11 uint16 = 32767
    if true {
        for fi12 := 0; fi12 < 4; fi12++ {
            var v13 int = fi12
            _ = fi12
            _ = v13
        }
        v3 += 3.14159
        for fi14 := 0; fi14 < 4; fi14++ {
            v3 = (-2.25 * 0.0)
            v11 += (v11 / ((65535) | 1))
            _ = fi14
        }
    }
    if false {
        v15 := v1[-194]
        for fi16 := 0; fi16 < 5; fi16++ {
            fmt.Println(fi16, v3, v15)
            var v17 string = (("\n\t" + (("foo bar" + ("foo bar" + "")) + "\n\t")) + ("a" + ("\n\t" + ("foo bar" + "hello"))))
            _ = len(v10)
            _ = fi16
            _ = v17
        }
        _ = v15
    } else {
        for fi18 := 0; fi18 < 11; fi18++ {
            for fi19 := 0; fi19 < 7; fi19++ {
                var v20 int32 = -256
                if (v4 > v4) {
                    v21 := GenS3{F0: genFn0(126)}
                    var v22 bool = (!false)
                    switch "\n\t" {
                    case "":
                        var v23 int16 = 1
                        _ = v23
                    }
                    _ = v21
                    _ = v22
                }
                v20 = (v20 - v20)
                fmt.Println(v3, v3)
                v24 := v1[fi19]
                _ = fi19
                _ = v20
                _ = v24
            }
            _ = fi18
        }
        v3 -= v3
    }
    for fi25 := 0; fi25 < 16; fi25++ {
        v4 = (v4 | (v4 * (v4 + (v4 * v4))))
        var v26 string = "hello"
        _ = fi25
        _ = v26
    }
    v3 += v3
    fmt.Println(v11, v3)
    if len(v10) >= 2 {
        v10 = v10[1:len(v10)]
    }
    fmt.Println(v3)
    if false {
        v27 := GenS1{F0: v4, F1: (91 < 211)}
        if len(v10) >= 2 {
            v10 = v10[1:len(v10)]
        }
        v11 = 58
        _ = v27
    } else {
        var v28 uint = 4294967295
        _ = v28
    }
    for fi29 := 0; fi29 < 3; fi29++ {
        fmt.Println(v4)
        for fi30 := 0; fi30 < 18; fi30++ {
            if (false || (false && false)) {
                v4 -= v4
                var v31 int = 32768
                _ = v31
            } else {
                v2["hello"] = true
            }
            v4 += v4
            v3 -= (v3 * float64(-2.25))
            _ = fi30
        }
        _ = fi29
    }
    v4 = v4
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
}
