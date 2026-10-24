package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int16 }
func genFn0(p0 int64) string {
    var v7 string = ((("" + "") + "") + ("a" + "hello"))
    v7 = ("hello" + v7)
    for fi8 := 0; fi8 < 12; fi8++ {
        v9 := []int{fi8, -9223372036854775808, (fi8 ^ fi8), fi8}
        var v10 uint16 = 2
        var v11 uint = 0
        _ = fi8
        _ = v10
        _ = v11
        _ = v9
    }
    switch (("foo bar" + "hello") + (v7 + v7)) {
    case "\n\t":
        p0 += 0
    case "hello":
        var v12 float64 = -2.25
        _ = v12
    case "foo bar":
        var v13 int8 = 55
        _ = v13
    default:
        v14 := GenS1{F0: 2}
        _ = v14
    }
    _ = p0
    _ = v7
    return (((v7 + "foo bar") + v7) + "a")
}

func main() {
    var v1 int16 = -1
    v2 := []uint32{}
    var v3 int16 = v1
    var v4 bool = (-274 > 122)
    v5 := make(map[int]int)
    var v6 uint32 = 4294967295
    v15 := GenS2{F0: 127}
    var v16 bool = v4
    if false {
        switch -834 {
        case 170:
            v17 := GenS2{F0: 1}
            _ = v17
        case -948:
            var v18 int8 = 0
            _ = v18
        }
        switch 844 {
        case 128:
            v19 := GenS2{F0: 128}
            _ = v19
        case 42:
            fmt.Println(v4, v1)
        default:
            if true {
                v20 := GenS3{F0: v3}
                _ = v20
            }
        }
    }
    v6 -= (^(v6))
    fmt.Println(v1, v3, v6)
    var v21 int16 = v3
    if v16 {
        var v22 uint32 = (v6 & (v6 + v6))
        for fi23 := 0; fi23 < 3; fi23++ {
            if len(v2) > 0 {
                v2[0] = (v6 % ((748) | 1))
            }
            var v24 int8 = 0
            _ = fi23
            _ = v24
        }
        _ = v22
    }
    var v25 uint16 = 65535
    v16 = ((((3.14159 - 3.14159) / float64((float32(1.5) - (0.1 / 1.5)))) - (3.14159 / 3.14159)) == 0.1)
    var v26 uint32 = (v6 - v6)
    fmt.Println(v21, v1, v26)
    v27 := GenS1{F0: v6}
    v28 := GenS1{F0: (v26 | (v6 % ((114) | 1)))}
    v1 -= 89
    v25 -= v25
    fmt.Println(v25, v4)
    v29 := GenS1{F0: v6}
    fmt.Println(v3, v21)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v21
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
