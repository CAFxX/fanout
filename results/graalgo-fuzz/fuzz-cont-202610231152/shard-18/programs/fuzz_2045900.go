package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int8; F1 uint32; F2 uint16 }
type GenS3 struct { F0 uint }
func genFn0(p0 string) uint64 {
    switch "foo bar" {
    case "\n\t":
        v7 := []int{-53, -1, 998}
        _ = v7
    case "a":
        var v8 int8 = 42
        _ = v8
    case "foo bar":
        switch p0 {
        case "a":
            var v9 int32 = -2147483648
            _ = v9
        case "foo bar":
            p0 = "hello"
        default:
            fmt.Println(p0)
        }
    }
    p0 = p0
    _ = p0
    return 4294967295
}

func main() {
    v1 := []int{}
    v2 := []uint32{0, 2147483647, 0}
    var v3 bool = true
    var v4 float64 = -2.25
    v5 := make(map[uint32]string)
    v6 := make(map[int]string)
    fmt.Println(v4, v3)
    v3 = (42 < 1)
    v10 := GenS1{F0: 0}
    v4 += ((0.0 + 0.1) - v4)
    var v11 int = -506
    v11 += v11
    if (970 < 910) {
        v4 += ((((v4 * v4) + (v4 / (1.5 - v4))) - 0.1) * 0.0)
    }
    v12 := GenS1{F0: -128}
    for fi13 := 0; fi13 < 12; fi13++ {
        var v14 uint = 945
        if ((32767 <= 0) || (float32((0.0 / v4)) == (0.0 * 1e10))) {
            v15 := v5[503]
            v3 = ((0.0 - 1.5) >= 0.1)
            v4 -= 3.14159
            _ = v15
        } else {
            v16 := GenS2{F0: 100, F1: 763, F2: 65535}
            var v17 int16 = 1000
            _ = v16
            _ = v17
        }
        switch ("\n\t" + "\n\t") {
        case "hello":
            fmt.Println(v14, v14, v3)
        case "\n\t":
            v18 := GenS3{F0: v14}
            _ = v18
        case "foo bar":
            switch (v11 * v11) {
            case 96:
                if (!v3) {
                    if (("foo bar" + "hello") > "hello") {
                        for fi19 := 0; fi19 < 16; fi19++ {
                            v20 := GenS3{F0: 0}
                            var v21 uint = (v14 | v14)
                            _ = fi19
                            _ = v20
                            _ = v21
                        }
                    }
                    switch 2147483647 {
                    case 32767:
                        var v22 int = v11
                        _ = v22
                    case -286:
                        v23 := GenS3{F0: v14}
                        _ = v23
                    case -426:
                        var v24 uint16 = uint16(v14)
                        _ = v24
                    }
                    v25 := GenS1{F0: -128}
                    _ = v25
                }
            case 42:
                var v26 int16 = -1
                _ = v26
            case 932:
                var v27 bool = v3
                _ = v27
            default:
                v11 -= (fi13 - (^(v11)))
            }
        default:
            var v28 int16 = 834
            _ = v28
        }
        _ = fi13
        _ = v14
    }
    var v29 float32 = 0.0
    var v30 float64 = v4
    var v31 uint16 = 32767
    var v32 int = v11
    for fi33 := 0; fi33 < 7; fi33++ {
        var v34 int64 = 0
        _ = fi33
        _ = v34
    }
    switch "a" {
    case "foo bar":
        v4 += 0.1
    case "":
        fmt.Println(v30)
    }
    var v35 bool = (v3 && v3)
    for fi36 := 0; fi36 < 14; fi36++ {
        switch (v11 * v32) {
        case 124:
            v37 := GenS2{F0: -1, F1: 436, F2: (v31 / ((20) | 1))}
            _ = v37
        case 42:
            v2 = append(v2, 2147483647)
        case -1:
            v29 -= 3.14159
        }
        var v38 bool = (("foo bar" + "foo bar") >= ("" + "foo bar"))
        fmt.Println(v35, v4, v31)
        _ = fi36
        _ = v38
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
