package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int16; F1 int8; F2 bool }
func genFn0(p0 uint32, p1 uint16, p2 uint64) uint32 {
    p1 = (^(p1))
    for fi4 := 0; fi4 < 1; fi4++ {
        v5 := []int{fi4, (^(fi4)), -545}
        _ = fi4
        _ = v5
    }
    _ = p0
    _ = p1
    _ = p2
    return 492
}
func genFn1(p0 int64) uint16 {
    v6 := GenS1{F0: (!true)}
    v7 := []int{}
    var v8 uint32 = 0
    var v9 int32 = 327
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 65535
}

func main() {
    var v1 float32 = 3.14159
    v2 := make(map[string]string)
    var v3 int16 = 0
    v10 := []int{128}
    var v11 int8 = 100
    fmt.Println(v1)
    if true {
        v10 = append(v10, 648)
        fmt.Println(v3)
    }
    fmt.Println(v1, v11, v3)
    v12 := GenS2{F0: (v3 >> ((2) & 63)), F1: 100, F2: (0 < 4294967295)}
    if false {
        v3 += v3
        switch "foo bar" {
        case "a":
            var v13 string = (((((("" + "\n\t") + ("" + "\n\t")) + "") + "\n\t") + ("\n\t" + "foo bar")) + ("" + "hello"))
            _ = v13
        case "\n\t":
            var v14 int8 = (v11 >> ((1) & 63))
            _ = v14
        case "hello":
            v2["a"] = "\n\t"
        }
    } else {
        fmt.Println(v3, v11, v11)
        if false {
            var v15 float64 = float64(v1)
            _ = v15
        }
    }
    var v16 float32 = ((v1 / ((v1 / v1) - v1)) * (v1 - (float32(0.0) / (v1 / ((v1 / 3.14159) / (v1 + ((v1 * v1) * (v1 / v1))))))))
    for fi17 := 0; fi17 < 15; fi17++ {
        if (0.1 >= 1.5) {
            var v18 float64 = (float64((3.14159 * 0.1)) * (-2.25 / ((1.5 * 0.0) - 3.14159)))
            if (v16 <= float32(0.0)) {
                v16 = v1
                switch "hello" {
                case "hello":
                    _ = len(v10)
                default:
                    switch "foo bar" {
                    case "\n\t":
                        if len(v10) >= 2 {
                            v10 = v10[1:len(v10)]
                        }
                    case "foo bar":
                        var v19 int8 = (-(v11))
                        _ = v19
                    }
                }
            }
            _ = v18
        }
        if len(v10) > 0 {
            v10[0] = fi17
        }
        _ = fi17
    }
    v2[("hello" + "\n\t")] = "foo bar"
    var v20 uint8 = 2
    _ = len(v10)
    var v21 int = 2147483647
    if true {
        fmt.Println(v1, v21, v21)
    }
    if true {
        v20 = (v20 ^ 255)
        fmt.Println(v3, v16)
    } else {
        v10 = append(v10, (v21 % ((v21) | 1)))
    }
    switch 32767 {
    case -9223372036854775808:
        var v22 int8 = 100
        _ = v22
    }
    v16 = v16
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v16
    _ = v2
    _ = v20
    _ = v21
    _ = v3
}
