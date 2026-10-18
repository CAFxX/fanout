package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string }
func genFn0(p0 int16) uint64 {
    p0 += (p0 + p0)
    v4 := GenS1{F0: 606, F1: "foo bar"}
    _ = p0
    _ = v4
    return 42
}
func genFn1(p0 uint64, p1 int32, p2 string) uint64 {
    switch p2 {
    case "hello":
        if (!(129 <= 1)) {
            var v5 int32 = (p1 | p1)
            v6 := GenS1{F0: 772, F1: p2}
            var v7 uint = 2
            _ = v5
            _ = v6
            _ = v7
        }
    case "foo bar":
        switch "" {
        case "foo bar":
            v8 := []int{-9223372036854775808, 0, -9223372036854775808, 49}
            _ = v8
        case "\n\t":
            p2 = p2
        }
    case "a":
        var v9 int8 = 1
        _ = v9
    }
    var v10 uint8 = 1
    if false {
        if false {
            if true {
                var v11 string = p2
                v10 += 255
                fmt.Println(p1)
                _ = v11
            }
        }
        var v12 int8 = 45
        fmt.Println(p2)
        _ = v12
    } else {
        v13 := []int{42, 0, -1, 42}
        _ = v13
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return (p0 << ((255) & 63))
}

func main() {
    v1 := make(map[int]string)
    var v2 float32 = ((1e10 - 1.5) * ((0.0 + float32(1.5)) * ((1e10 * 3.14159) / 1.5)))
    v3 := []uint32{4294967295, 283}
    v2 -= (float32(((3.14159 / 0.1) + float64(float32(((((1e10 + -2.25) * 0.1) * 3.14159) / -2.25))))) + (v2 + ((v2 / v2) - -2.25)))
    if (!false) {
        var v14 string = "a"
        var v15 float32 = -2.25
        _ = v14
        _ = v15
    }
    v2 = v2
    var v16 uint32 = 175
    v17 := GenS1{F0: v16, F1: ("a" + ("foo bar" + "foo bar"))}
    v2 = 0.0
    var v18 int64 = 33
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi19 := 0; fi19 < 7; fi19++ {
        if true {
            var v20 uint16 = 1
            v18 -= (v18 - 0)
            v1[fi19] = ("" + "foo bar")
            _ = v20
        }
        _ = fi19
    }
    for fi21 := 0; fi21 < 15; fi21++ {
        v3 = append(v3, (v16 / ((v16) | 1)))
        var v22 uint = 255
        for fi23 := 0; fi23 < 7; fi23++ {
            var v24 int32 = int32(fi23)
            v25 := GenS1{F0: (v16 | v16), F1: ""}
            _ = fi23
            _ = v24
            _ = v25
        }
        _ = fi21
        _ = v22
    }
    for fi26 := 0; fi26 < 8; fi26++ {
        var v27 uint8 = 195
        _ = fi26
        _ = v27
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
}
