package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16 }
type GenS2 struct { F0 string; F1 uint8 }
func genFn0(p0 bool, p1 uint16) uint32 {
    fmt.Println(p0, p1)
    v4 := GenS1{F0: -543, F1: 1000}
    var v5 int8 = 42
    if true {
        p1 = p1
        var v6 int64 = 42
        fmt.Println(p1, v6, p1)
        _ = v6
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 65535
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[int]bool)
    v3 := make(map[int]string)
    v7 := GenS1{F0: 42, F1: 723}
    if ((0 < -2147483648) || (4294967295 > 4294967295)) {
        var v8 float64 = float64(((0.0 - (1e10 / 3.14159)) / (float32(1e10) + 1e10)))
        var v9 bool = true
        if false {
            var v10 int8 = -128
            v11 := GenS1{F0: 42, F1: 42}
            _ = v10
            _ = v11
        }
        _ = v8
        _ = v9
    } else {
        for fi12 := 0; fi12 < 18; fi12++ {
            v13 := []int{fi12, 9223372036854775807}
            var v14 float32 = (3.14159 + (1.5 - (0.0 - (0.0 + 0.1))))
            v15 := GenS2{F0: (("a" + "hello") + ""), F1: 127}
            _ = fi12
            _ = v13
            _ = v14
            _ = v15
        }
        var v16 int32 = -2147483648
        _ = v16
    }
    var v17 float32 = 0.1
    v18 := []int{0, 9223372036854775807, 9223372036854775807}
    var v19 int16 = -1
    v17 = v17
    v17 = v17
    switch 32767 {
    case 2147483647:
        v17 += (-2.25 / v17)
    case 182:
        fmt.Println(v19, v19)
    }
    _ = len(v18)
    fmt.Println(v17, v17)
    v20 := GenS2{F0: "a", F1: 1}
    for fi21 := 0; fi21 < 6; fi21++ {
        fmt.Println(v19)
        var v22 uint64 = 18446744073709551615
        switch (fi21 / ((-2147483648) | 1)) {
        case -993:
            fmt.Println(v17, v17)
        default:
            for fi23 := 0; fi23 < 14; fi23++ {
                var v24 int64 = 0
                v19 += v19
                _ = fi23
                _ = v24
            }
        }
        _ = fi21
        _ = v22
    }
    v25 := GenS2{F0: "foo bar", F1: 227}
    switch "hello" {
    case "foo bar":
        fmt.Println(v19)
    case "hello":
        switch 9223372036854775807 {
        case -113:
            switch (((("a" + (("\n\t" + "a") + "\n\t")) + "a") + "foo bar") + "") {
            case "":
                v26 := GenS1{F0: -2147483648, F1: (v19 | (-(v19)))}
                _ = v26
            case "a":
                if len(v18) >= 2 {
                    v18 = v18[1:len(v18)]
                }
            }
        case 128:
            var v27 int = 128
            _ = v27
        case 127:
            v2[552] = (-595 <= 4294967295)
        default:
            for fi28 := 0; fi28 < 7; fi28++ {
                for fi29 := 0; fi29 < 8; fi29++ {
                    fmt.Println(v19, v19)
                    _ = fi29
                }
                v30 := GenS2{F0: "a", F1: 2}
                _ = fi28
                _ = v30
            }
        }
    }
    for fi31 := 0; fi31 < 5; fi31++ {
        _ = len(v3)
        if len(v18) > 0 {
            v18[0] = fi31
        }
        var v32 uint16 = 255
        _ = fi31
        _ = v32
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v25
    _ = v3
    _ = v7
}
