package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32; F2 int }
func genFn0(p0 int) int64 {
    v6 := GenS1{F0: (("hello" + "\n\t") + "\n\t"), F1: 2147483647, F2: 1}
    fmt.Println(p0)
    v7 := []int{(-(p0)), 4294967295}
    if ((p0 >= 1) || (((((((0.1 * -2.25) * 1.5) / 1.0) / (((1e10 * -2.25) + 1.5) * 3.14159)) + 3.14159) - (3.14159 - float32(1e10))) <= 0.0)) {
        v8 := GenS1{F0: "hello", F1: -408, F2: p0}
        var v9 float64 = ((1e10 - (3.14159 / ((float64(-2.25) - float64(((3.14159 / 1e10) - 1e10))) + 3.14159))) * (0.0 + 3.14159))
        _ = v8
        _ = v9
    }
    _ = p0
    _ = v6
    _ = v7
    return 42
}
func genFn1(p0 uint) uint8 {
    fmt.Println(p0)
    switch (("\n\t" + "foo bar") + (("" + "") + "foo bar")) {
    case "foo bar":
        for fi10 := 0; fi10 < 17; fi10++ {
            var v11 int64 = 1
            var v12 int16 = 42
            v12 = (v12 % ((-32768) | 1))
            _ = fi10
            _ = v11
            _ = v12
        }
    default:
        var v13 float32 = 0.0
        _ = v13
    }
    var v14 int16 = 0
    _ = p0
    _ = v14
    return 255
}

func main() {
    var v1 uint16 = 255
    var v2 int = 177
    v3 := []int{}
    v4 := []string{}
    var v5 uint16 = (v1 << ((1) & 63))
    switch v2 {
    case 9223372036854775807:
        v1 += 327
    }
    v5 -= (v1 * (v5 - 42))
    fmt.Println(v2, v1)
    v5 += v5
    var v15 uint8 = 240
    var v16 int8 = 42
    fmt.Println(v15, v2, v5)
    v16 = v16
    if (v16 == (^(v16))) {
        var v17 int32 = -1
        var v18 uint8 = (v15 << ((0) & 63))
        _ = v17
        _ = v18
    } else {
        var v19 int32 = -1
        var v20 uint16 = (v1 | (v5 - (^(v1))))
        _ = v19
        _ = v20
    }
    var v21 uint32 = 65535
    for fi22 := 0; fi22 < 16; fi22++ {
        v23 := GenS1{F0: "hello", F1: 2147483647, F2: fi22}
        for fi24 := 0; fi24 < 12; fi24++ {
            switch "" {
            case "":
                v4 = append(v4, "foo bar")
            }
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = fi24
        }
        v25 := GenS1{F0: "", F1: 42, F2: 846}
        _ = fi22
        _ = v23
        _ = v25
    }
    var v26 uint16 = (^(v5))
    fmt.Println(v16, v5, v26)
    v27 := GenS1{F0: "foo bar", F1: 70, F2: (v2 - int(v16))}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v21
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
