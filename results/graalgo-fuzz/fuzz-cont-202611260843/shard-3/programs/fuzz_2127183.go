package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8 }
func genFn0(p0 string, p1 bool) uint {
    switch -756 {
    case -70:
        var v7 int = -870
        _ = v7
    case 32768:
        v8 := []int{2147483647, 678}
        _ = v8
    default:
        fmt.Println(p1, p0)
    }
    p1 = (826 >= 42)
    fmt.Println(p0)
    _ = p0
    _ = p1
    return 18446744073709551615
}
func genFn1() uint32 {
    var v9 string = ("" + "\n\t")
    v10 := []int{-1, 17}
    fmt.Println(v9)
    if (32767 > 0) {
        for fi11 := 0; fi11 < 1; fi11++ {
            v9 = ""
            var v12 int64 = 9223372036854775807
            var v13 int = 309
            _ = fi11
            _ = v12
            _ = v13
        }
    }
    switch -38 {
    case -319:
        if len(v10) > 0 {
            v10[0] = 127
        }
    case 32768:
        var v14 string = (v9 + "a")
        _ = v14
    case 42:
        fmt.Println(v9)
    default:
        if len(v10) >= 2 {
            v10 = v10[1:len(v10)]
        }
    }
    _ = v10
    _ = v9
    return 955
}

func main() {
    var v1 uint8 = 45
    var v2 uint = 1
    var v3 uint8 = (v1 * (v1 ^ v1))
    var v4 uint32 = 4294967295
    var v5 float64 = (0.1 - 0.0)
    v6 := make(map[uint32]bool)
    if false {
        v5 = 3.14159
        v1 = v1
    } else {
        var v15 uint32 = 0
        _ = v15
    }
    var v16 string = ("hello" + "\n\t")
    switch v16 {
    case "\n\t":
        v5 = v5
    default:
        if true {
            v17 := v6[4294967295]
            var v18 int8 = 127
            if (-1 <= 42) {
                v1 = (v3 & 1)
                fmt.Println(v3)
            } else {
                delete(v6, v4)
            }
            _ = v17
            _ = v18
        } else {
            switch 0 {
            case 843:
                v19 := []int{}
                _ = v19
            default:
                v5 -= v5
            }
            for fi20 := 0; fi20 < 1; fi20++ {
                delete(v6, (v4 << (((v2 * uint(v4))) & 63)))
                _ = fi20
            }
        }
    }
    if false {
        var v21 uint8 = v3
        _ = v21
    } else {
        var v22 float32 = float32(0.0)
        _ = v22
    }
    v23 := []int{-9223372036854775808, int(v1)}
    if len(v23) >= 2 {
        v23 = v23[1:len(v23)]
    }
    v1 -= (v3 * v1)
    if ((((float32(((v5 + 3.14159) + ((v5 + float64(((float32(1e10) * -2.25) - 1e10))) / (-2.25 * v5)))) + 1.5) - 0.1) <= (((float32(v5) - float32(float64(((0.1 + 0.0) / (float32(1e10) + -2.25))))) + (3.14159 / 1.0)) * 1.5)) && (float64(float32(-2.25)) <= v5)) {
        for fi24 := 0; fi24 < 11; fi24++ {
            v23 = append(v23, (fi24 + fi24))
            _ = fi24
        }
        for fi25 := 0; fi25 < 7; fi25++ {
            for fi26 := 0; fi26 < 8; fi26++ {
                v16 = v16
                v1 = v3
                _ = fi26
            }
            _ = fi25
        }
    } else {
        v3 += 255
        if (("\n\t" + "hello") > v16) {
            fmt.Println(v16)
            if len(v23) > 0 {
                v23[0] = -2147483648
            }
        } else {
            if true {
                if len(v23) >= 2 {
                    v23 = v23[1:len(v23)]
                }
            }
        }
    }
    var v27 int8 = 1
    v5 = v5
    v28 := GenS1{F0: "foo bar", F1: v27}
    var v29 int64 = 1
    fmt.Println(v16, v5)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v16
    _ = v2
    _ = v23
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
