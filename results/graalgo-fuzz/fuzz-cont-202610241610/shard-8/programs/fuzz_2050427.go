package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64 }
type GenS2 struct { F0 int32 }
func genFn0(p0 string, p1 int8, p2 int64) uint {
    for fi5 := 0; fi5 < 20; fi5++ {
        for fi6 := 0; fi6 < 2; fi6++ {
            fmt.Println(p2)
            _ = fi6
        }
        fmt.Println(fi5, fi5, p0)
        _ = fi5
    }
    for fi7 := 0; fi7 < 1; fi7++ {
        p0 = (p0 + p0)
        if true {
            v8 := GenS1{F0: 0, F1: 0}
            fmt.Println(p0)
            p0 = p0
            _ = v8
        }
        _ = fi7
    }
    for fi9 := 0; fi9 < 15; fi9++ {
        v10 := []int{-365, fi9}
        if len(v10) >= 2 {
            v10 = v10[1:len(v10)]
        }
        p1 -= p1
        _ = fi9
        _ = v10
    }
    _ = p0
    _ = p1
    _ = p2
    return 42
}

func main() {
    var v1 uint32 = 83
    var v2 float32 = (0.0 / (0.1 / (1.5 - ((-2.25 + (3.14159 + (float32((1.5 + 0.1)) + 1e10))) + -2.25))))
    var v3 float64 = ((3.14159 * 0.0) + (float64(((float32((((float64(v2) * -2.25) * (0.1 * (((-2.25 + -2.25) * 1e10) - (1e10 - -2.25)))) - 1.5)) * (((v2 / 1.0) - 1.5) + ((v2 - v2) + v2))) - 1e10)) + (1.5 - (1e10 + (-2.25 + 1.5)))))
    var v4 float32 = 1e10
    v11 := []int{-9223372036854775808, 4294967295, 127, -632}
    var v12 int = 32767
    if (-1 <= 1) {
        if false {
            for fi13 := 0; fi13 < 15; fi13++ {
                v14 := GenS2{F0: 42}
                fmt.Println(v4)
                fmt.Println(v4, v2)
                _ = fi13
                _ = v14
            }
        } else {
            if (((!(((("\n\t" + "") + ("hello" + "\n\t")) + "a") < "hello")) || ((^(v1)) >= (v1 << ((446) & 63)))) && (3.14159 >= ((1.5 + v4) + -2.25))) {
                if false {
                    if len(v11) > 0 {
                        v11[0] = (v12 + 0)
                    }
                }
            }
        }
        v2 -= v4
        v15 := GenS1{F0: 42, F1: 0}
        _ = v15
    } else {
        fmt.Println(v1, v3, v4)
        switch v12 {
        case 42:
            v1 += (v1 & (v1 * (v1 << ((4294967295) & 63))))
        case 4294967295:
            if (32767 != -32768) {
                if len(v11) > 0 {
                    v11[0] = (v12 | (v12 - v12))
                }
            } else {
                var v16 int64 = -1
                if len(v11) > 0 {
                    v11[0] = (v12 ^ v12)
                }
                fmt.Println(v1, v2)
                _ = v16
            }
        case -372:
            var v17 int = 594
            _ = v17
        }
    }
    v12 = 450
    var v18 float64 = 3.14159
    v1 += 408
    for fi19 := 0; fi19 < 6; fi19++ {
        if len(v11) > 0 {
            v11[0] = -2147483648
        }
        var v20 int32 = 1
        _ = fi19
        _ = v20
    }
    _ = len(v11)
    fmt.Println(v3)
    var v21 string = "\n\t"
    switch (v12 ^ 127) {
    case 0:
        var v22 uint8 = 255
        _ = v22
    case 546:
        v3 -= v18
    default:
        fmt.Println(v3, v21, v3)
    }
    v23 := GenS2{F0: -55}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v18
    _ = v2
    _ = v21
    _ = v23
    _ = v3
    _ = v4
}
