package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint; F2 string }
type GenS2 struct { F0 uint32 }
func genFn0(p0 bool, p1 int8) int {
    v6 := []int{32767}
    v6 = append(v6, -723)
    _ = p0
    _ = p1
    _ = v6
    return 127
}
func genFn1() int64 {
    var v7 uint32 = 830
    v7 -= v7
    _ = v7
    return -9223372036854775808
}

func main() {
    var v1 int64 = 1
    var v2 float64 = 3.14159
    v3 := make(map[uint32]uint32)
    v4 := make(map[int]bool)
    var v5 int = 32768
    var v8 int64 = -525
    var v9 float64 = ((float64(((1.5 - 1.5) * (float32((1.5 / 3.14159)) + (0.0 + (-2.25 + 1e10))))) / (1.5 * v2)) * v2)
    var v10 string = ("foo bar" + "hello")
    fmt.Println(v9, v8)
    v1 = (v1 / (((-(v8))) | 1))
    switch (v5 / ((434) | 1)) {
    case 9223372036854775807:
        _ = len(v3)
    case 4294967295:
        var v11 uint64 = 1
        _ = v11
    case 127:
        if ((float64(float32(float64((float32(v9) / 1e10)))) * (v9 / ((v9 / (v9 / v2)) * (v9 / -2.25)))) <= v2) {
            switch v10 {
            case "foo bar":
                v5 -= v5
            case "\n\t":
                fmt.Println(v2)
            case "a":
                v8 += (v8 * v1)
            }
        } else {
            var v12 uint16 = 2
            _ = v12
        }
    }
    v5 += (v5 % (((v5 & (v5 + v5))) | 1))
    v13 := []int{-53, genFn0((false && (!true)), 127), v5}
    if false {
        if true {
            v2 = (v9 * (v2 - v9))
        } else {
            v2 = float64(-2.25)
            for fi14 := 0; fi14 < 18; fi14++ {
                v5 = genFn0((!true), -128)
                _ = fi14
            }
        }
        v8 += 1
    }
    for fi15 := 0; fi15 < 16; fi15++ {
        for fi16 := 0; fi16 < 11; fi16++ {
            v4[fi15] = (2 > 0)
            _ = fi16
        }
        switch genFn0(false, -128) {
        case 42:
            for fi17 := 0; fi17 < 13; fi17++ {
                v5 = v5
                var v18 int32 = -508
                _ = fi17
                _ = v18
            }
        }
        _ = fi15
    }
    v19 := GenS1{F0: 32767, F1: 65535, F2: "a"}
    if (166 > 127) {
        v8 = -9223372036854775808
        switch v10 {
        case "a":
            if ((1 > 4294967295) && (genFn0(((1 != 75) || false), 0) <= (v5 + (v5 << ((65535) & 63))))) {
                var v20 int = v5
                v21 := GenS2{F0: 42}
                if len(v13) > 0 {
                    v13[0] = (v5 & (v20 >> ((650) & 63)))
                }
                _ = v20
                _ = v21
            }
        case "hello":
            for fi22 := 0; fi22 < 16; fi22++ {
                v2 += 1.5
                _ = fi22
            }
        }
    }
    if true {
        v8 = genFn1()
    }
    for fi23 := 0; fi23 < 17; fi23++ {
        v24 := GenS1{F0: 368, F1: 770, F2: ("" + (v10 + v10))}
        v4[v5] = false
        _ = fi23
        _ = v24
    }
    if (21 != 188) {
        var v25 bool = (2 >= 42)
        for fi26 := 0; fi26 < 19; fi26++ {
            fmt.Println(v1, v10)
            v1 -= v8
            _ = fi26
        }
        v4[(v5 - (-(v5)))] = (!(3.14159 != (0.1 / 1.0)))
        _ = v25
    } else {
        for fi27 := 0; fi27 < 2; fi27++ {
            fmt.Println(v8, v9)
            fmt.Println(v8, v5, v8)
            v8 = 0
            _ = fi27
        }
    }
    fmt.Println(v2, v10)
    v5 = -639
    switch v10 {
    case "hello":
        v3[2147483648] = 355
    case "\n\t":
        v28 := GenS1{F0: -1, F1: 255, F2: ("" + "hello")}
        _ = v28
    case "foo bar":
        fmt.Println(v8)
    default:
        var v29 string = v10
        _ = v29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v13
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
