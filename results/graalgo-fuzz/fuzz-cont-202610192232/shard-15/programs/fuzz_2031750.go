package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string }
type GenS2 struct { F0 int64; F1 int16 }

func main() {
    var v1 float32 = 0.1
    var v2 int8 = 48
    v3 := make(map[uint32]int)
    var v4 uint32 = 884
    v5 := GenS1{F0: 0, F1: "hello"}
    v6 := v3[(v4 << ((uint(v4)) & 63))]
    switch "foo bar" {
    case "hello":
        fmt.Println(v2)
    case "a":
        var v7 uint32 = (v4 / ((523) | 1))
        _ = v7
    default:
        var v8 float64 = ((0.1 + 1e10) + ((3.14159 + float64(v1)) / (1e10 / 0.1)))
        _ = v8
    }
    var v9 uint64 = 2
    if true {
        switch ("hello" + "a") {
        case "a":
            v1 -= (1.5 + (v1 + v1))
        }
        switch 0 {
        case -249:
            if (-508 <= 42) {
                if true {
                    var v10 float32 = v1
                    _ = v10
                }
                v9 += 1
            }
        case 32767:
            v9 -= v9
        case 128:
            var v11 uint = 65535
            _ = v11
        }
    }
    switch ((("" + ("" + ("foo bar" + ("hello" + ("" + "hello"))))) + ("a" + ("" + "foo bar"))) + "a") {
    case "hello":
        var v12 float64 = float64(((v1 * (v1 + (0.0 * ((float32(0.0) / ((v1 / v1) / -2.25)) * v1)))) * 1.5))
        _ = v12
    case "\n\t":
        v4 = 65535
    default:
        v13 := GenS2{F0: 0, F1: -165}
        _ = v13
    }
    switch (("\n\t" + "") + "") {
    case "foo bar":
        v4 = (v4 + (v4 & v4))
    default:
        for fi14 := 0; fi14 < 13; fi14++ {
            v15 := v3[v4]
            var v16 string = ("" + "a")
            _ = fi14
            _ = v15
            _ = v16
        }
    }
    for fi17 := 0; fi17 < 3; fi17++ {
        v2 = (-(v2))
        _ = fi17
    }
    var v18 int64 = -487
    var v19 float32 = v1
    for fi20 := 0; fi20 < 18; fi20++ {
        v18 += v18
        for fi21 := 0; fi21 < 8; fi21++ {
            v2 = v2
            _ = fi21
        }
        _ = fi20
    }
    var v22 int64 = -710
    v1 = v1
    v23 := v3[567]
    v24 := []int{(v23 >> ((2) & 63)), 127, (v23 % ((v6) | 1))}
    v25 := GenS1{F0: 54, F1: "a"}
    var v26 int = 127
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
