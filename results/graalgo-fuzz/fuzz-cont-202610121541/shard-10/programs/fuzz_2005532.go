package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16; F2 int }
type GenS2 struct { F0 string }
func genFn0() uint8 {
    if false {
        var v5 uint64 = 755
        if false {
            var v6 float64 = (0.1 * (3.14159 * -2.25))
            var v7 int64 = 491
            _ = v6
            _ = v7
        }
        _ = v5
    }
    var v8 uint8 = 0
    var v9 string = ""
    v9 = v9
    _ = v8
    _ = v9
    return v8
}
func genFn1(p0 int) bool {
    switch p0 {
    case 928:
        if false {
            v10 := []int{-903}
            _ = v10
        }
    case 4294967295:
        var v11 uint64 = 0
        _ = v11
    case 42:
        v12 := GenS1{F0: false, F1: 65535, F2: (p0 - (p0 - (p0 + 2147483647)))}
        _ = v12
    }
    var v13 int = p0
    v13 += (p0 ^ p0)
    for fi14 := 0; fi14 < 2; fi14++ {
        v15 := []int{fi14, (p0 | -2147483648)}
        p0 += 0
        _ = fi14
        _ = v15
    }
    _ = p0
    _ = v13
    return (127 == (p0 & (v13 + p0)))
}

func main() {
    var v1 int = 32767
    var v2 uint16 = 255
    var v3 uint16 = v2
    var v4 uint8 = 250
    fmt.Println(v3, v3)
    v16 := GenS1{F0: (1 > 4294967295), F1: v3, F2: v1}
    var v17 string = "a"
    v18 := []int{v1, v1}
    v3 += (v2 + v2)
    switch v1 {
    case 32768:
        fmt.Println(v3, v2, v4)
    }
    _ = len(v18)
    fmt.Println(v4)
    var v19 int8 = -1
    v19 += (-(v19))
    fmt.Println(v1, v3, v1)
    for fi20 := 0; fi20 < 4; fi20++ {
        var v21 uint32 = 2
        for fi22 := 0; fi22 < 16; fi22++ {
            switch v17 {
            case "\n\t":
                v23 := GenS1{F0: ((v4 % ((253) | 1)) != v4), F1: v2, F2: (v1 >> ((2) & 63))}
                _ = v23
            }
            var v24 float32 = 0.1
            _ = fi22
            _ = v24
        }
        var v25 int64 = -877
        _ = fi20
        _ = v21
        _ = v25
    }
    if (v1 > v1) {
        if true {
            if len(v18) >= 2 {
                v18 = v18[1:len(v18)]
            }
            v26 := GenS2{F0: (v17 + (v17 + "\n\t"))}
            fmt.Println(v2, v3)
            _ = v26
        } else {
            if true {
                switch v1 {
                case 127:
                    fmt.Println(v4, v1, v1)
                case 630:
                    v2 = (v3 >> ((4294967295) & 63))
                case 2147483647:
                    if len(v18) >= 2 {
                        v18 = v18[1:len(v18)]
                    }
                default:
                    switch v1 {
                    case -455:
                        var v27 uint = 18446744073709551615
                        _ = v27
                    }
                }
                for fi28 := 0; fi28 < 9; fi28++ {
                    v1 += -9223372036854775808
                    v17 = ((((v17 + v17) + (v17 + ((v17 + v17) + v17))) + v17) + ((v17 + v17) + "\n\t"))
                    _ = fi28
                }
                v1 = v1
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
