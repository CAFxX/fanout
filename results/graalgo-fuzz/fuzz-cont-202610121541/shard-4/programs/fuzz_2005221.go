package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int }
type GenS2 struct { F0 uint8; F1 int16 }
type GenS3 struct { F0 uint32; F1 int; F2 string }
func genFn0(p0 uint, p1 bool) uint8 {
    p0 = (p0 ^ p0)
    fmt.Println(p0, p0)
    fmt.Println(p0, p1)
    if (p1 && (-1 >= 100)) {
        fmt.Println(p1, p0)
        v7 := GenS1{F0: 2, F1: 42}
        p0 = (p0 % ((42) | 1))
        _ = v7
    } else {
        v8 := []int{-2147483648, 158, 9223372036854775807, 0}
        p0 += p0
        var v9 bool = p1
        _ = v8
        _ = v9
    }
    _ = p0
    _ = p1
    return 0
}
func genFn1() uint64 {
    var v10 uint32 = 0
    fmt.Println(v10)
    for fi11 := 0; fi11 < 17; fi11++ {
        var v12 int8 = 100
        if (!true) {
            var v13 float64 = float64(0.0)
            v10 += (v10 / ((688) | 1))
            v14 := []int{(fi11 - fi11)}
            _ = v13
            _ = v14
        }
        fmt.Println(v12, v12)
        _ = fi11
        _ = v12
    }
    _ = v10
    return 4294967295
}

func main() {
    v1 := []string{("\n\t" + (("\n\t" + "") + ""))}
    var v2 uint = 255
    var v3 uint8 = 2
    v4 := make(map[string]uint32)
    var v5 uint = 255
    v6 := make(map[int]int)
    v15 := GenS2{F0: v3, F1: 605}
    for fi16 := 0; fi16 < 16; fi16++ {
        var v17 int16 = 1
        fmt.Println(v5)
        if (((9223372036854775807 == 9223372036854775807) && (false || false)) || (0 != 723)) {
            v17 += (v17 ^ v17)
        }
        _ = fi16
        _ = v17
    }
    var v18 int = -138
    for fi19 := 0; fi19 < 11; fi19++ {
        v3 = 2
        fmt.Println(v2, v3, v3)
        _ = fi19
    }
    v5 -= v5
    for fi20 := 0; fi20 < 5; fi20++ {
        v18 += v18
        v18 = v18
        _ = fi20
    }
    var v21 int = 0
    v6[(v21 >> ((255) & 63))] = (-(v21))
    if true {
        var v22 string = ("a" + "")
        _ = v22
    } else {
        switch "a" {
        case "\n\t":
            v21 += 2147483647
        case "a":
            if false {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
        }
        for fi23 := 0; fi23 < 5; fi23++ {
            if (false && (2 != 2)) {
                if (-1 == 1) {
                    if (!false) {
                        if (-1 < 42) {
                            v24 := GenS3{F0: 120, F1: v21, F2: ("foo bar" + "foo bar")}
                            _ = v24
                        }
                        var v25 float32 = 0.0
                        _ = v25
                    }
                    v18 = v21
                } else {
                    var v26 int16 = 792
                    switch "\n\t" {
                    case "a":
                        v5 += (v5 % ((v5) | 1))
                    case "foo bar":
                        var v27 uint8 = v3
                        _ = v27
                    }
                    v3 -= genFn0((v5 | v2), true)
                    _ = v26
                }
            } else {
                var v28 float32 = ((((float32(-2.25) - -2.25) / 1.5) - 1.5) - ((((3.14159 - 3.14159) + 1.5) * 1e10) / 1.0))
                _ = v28
            }
            v18 = 2147483647
            _ = fi23
        }
    }
    v21 = -1
    var v29 int32 = 1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v15
    _ = v18
    _ = v2
    _ = v21
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
