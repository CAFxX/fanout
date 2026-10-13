package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string; F2 string }
type GenS2 struct { F0 int8; F1 int16; F2 int16 }
type GenS3 struct { F0 int16; F1 uint8 }
func genFn0(p0 string, p1 int32) uint8 {
    switch "foo bar" {
    case "hello":
        fmt.Println(p1, p0)
    case "\n\t":
        var v7 int16 = 42
        _ = v7
    default:
        p1 = 0
    }
    v8 := GenS1{F0: 255, F1: p0, F2: "foo bar"}
    fmt.Println(p0, p0)
    v9 := GenS1{F0: 2, F1: "hello", F2: p0}
    var v10 uint16 = 2
    _ = p0
    _ = p1
    _ = v10
    _ = v8
    _ = v9
    return 37
}
func genFn1() int32 {
    var v11 int = 216
    v12 := []int{1, v11, v11}
    var v13 float64 = ((0.0 * 3.14159) / (3.14159 * (0.1 * 0.1)))
    if len(v12) >= 2 {
        v12 = v12[1:len(v12)]
    }
    var v14 int32 = 42
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    return (v14 - (v14 + (v14 / ((v14) | 1))))
}

func main() {
    v1 := []string{"hello", "\n\t", "a", "hello"}
    v2 := make(map[string]int)
    var v3 uint16 = 255
    var v4 int = 128
    var v5 string = "hello"
    v6 := make(map[int]uint32)
    var v15 int32 = -2147483648
    var v16 uint32 = 65535
    if false {
        if len(v1) > 0 {
            v1[0] = ((((v5 + "a") + v5) + "\n\t") + v5)
        }
        v3 -= 32767
    }
    switch v4 {
    case 2147483647:
        v17 := GenS1{F0: v3, F1: v5, F2: ((((("\n\t" + ((v5 + v5) + ("foo bar" + "\n\t"))) + "\n\t") + ("foo bar" + v5)) + "foo bar") + (v5 + v5))}
        _ = v17
    }
    var v18 int16 = -32768
    v19 := GenS1{F0: 0, F1: ("a" + ""), F2: v5}
    v20 := GenS1{F0: (v3 ^ (^(v3))), F1: ((("a" + v5) + "a") + v5), F2: "foo bar"}
    v18 = (v18 + v18)
    var v21 int = v4
    v18 += int16(v21)
    var v22 bool = false
    for fi23 := 0; fi23 < 2; fi23++ {
        v24 := GenS2{F0: -43, F1: v18, F2: 42}
        v5 = ((v5 + v5) + v5)
        for fi25 := 0; fi25 < 9; fi25++ {
            v6[(fi23 % ((fi25) | 1))] = 0
            v26 := v2[(v5 + (v5 + ("a" + v5)))]
            v16 = (v16 ^ (v16 * v16))
            v16 += 461
            _ = fi25
            _ = v26
        }
        _ = fi23
        _ = v24
    }
    var v27 float32 = 1e10
    v28 := GenS3{F0: -533, F1: 135}
    var v29 int = (v4 - v4)
    for fi30 := 0; fi30 < 4; fi30++ {
        fmt.Println(v22)
        fmt.Println(v16, v29)
        _ = fi30
    }
    var v31 int32 = int32(v21)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v15
    _ = v16
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
