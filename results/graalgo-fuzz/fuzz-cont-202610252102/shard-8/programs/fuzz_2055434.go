package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int }
type GenS2 struct { F0 int; F1 string; F2 uint32 }
type GenS3 struct { F0 uint16; F1 uint8 }
type GenS4 struct { F0 int }
func genFn0(p0 int64, p1 uint64, p2 string) int {
    fmt.Println(p1, p2)
    if (132 != 724) {
        var v5 float64 = ((3.14159 - float64(((0.0 / (1e10 / 1.0)) / 1e10))) * 0.0)
        p1 -= p1
        _ = v5
    } else {
        var v6 int32 = -789
        _ = v6
    }
    if (255 > 42) {
        v7 := []int{-257, 0}
        if false {
            _ = len(v7)
        } else {
            if (9223372036854775807 < p0) {
                fmt.Println(p0)
            }
            var v8 float64 = float64(0.0)
            _ = v8
        }
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    return 9223372036854775807
}

func main() {
    var v1 bool = (((-1 <= 1) && true) && ((!true) || (3.14159 != -2.25)))
    v2 := make(map[uint32]bool)
    v3 := make(map[uint32]string)
    var v4 int8 = 42
    if v1 {
        var v9 string = ""
        _ = v9
    } else {
        var v10 uint64 = 2
        if false {
            if v1 {
                v11 := GenS1{F0: ("" + ("\n\t" + "\n\t")), F1: 128}
                v12 := []int{-837, 128, genFn0(660, v10, "foo bar"), 2147483647}
                _ = v11
                _ = v12
            } else {
                v10 = 0
                if (-2147483648 < 42) {
                    v10 = v10
                }
                for fi13 := 0; fi13 < 18; fi13++ {
                    v1 = v1
                    _ = fi13
                }
            }
        } else {
            v4 += v4
            v10 = (v10 | v10)
        }
        v14 := GenS1{F0: "", F1: 2147483647}
        _ = v10
        _ = v14
    }
    var v15 string = ""
    v15 = "a"
    if v1 {
        v16 := GenS2{F0: -2147483648, F1: "a", F2: 2}
        var v17 float64 = 0.1
        var v18 float64 = (v17 * 0.0)
        _ = v16
        _ = v17
        _ = v18
    }
    for fi19 := 0; fi19 < 13; fi19++ {
        if ((fi19 % ((fi19) | 1)) <= genFn0(-1, 793, v15)) {
            if v1 {
                var v20 int = fi19
                v21 := GenS2{F0: v20, F1: ("foo bar" + v15), F2: 0}
                _ = v20
                _ = v21
            } else {
                fmt.Println(fi19, v15)
                var v22 int32 = 304
                _ = v22
            }
        } else {
            var v23 uint8 = 242
            _ = v23
        }
        v24 := []int{(fi19 / ((genFn0(-9223372036854775808, 931, ("" + "hello"))) | 1)), (fi19 << ((590) & 63)), fi19, fi19}
        _ = fi19
        _ = v24
    }
    var v25 int = 1
    var v26 uint8 = 0
    switch v15 {
    case "hello":
        fmt.Println(v25, v25, v1)
    case "\n\t":
        v27 := GenS3{F0: 266, F1: 192}
        _ = v27
    case "a":
        v28 := v2[316]
        _ = v28
    }
    v29 := GenS4{F0: v25}
    switch (v25 * 128) {
    case 32768:
        for fi30 := 0; fi30 < 20; fi30++ {
            fmt.Println(v26)
            _ = fi30
        }
    case -548:
        v15 = ("hello" + ("foo bar" + v15))
    case 21:
        fmt.Println(v26, v25)
    }
    fmt.Println(v15, v1, v1)
    if false {
        v31 := []int{v25, 396}
        v26 = (v26 << ((4294967295) & 63))
        var v32 float64 = ((float64(0.0) + ((0.0 / 1.5) * (0.1 + 1.5))) / -2.25)
        _ = v31
        _ = v32
    }
    v4 = (v4 + (-(v4)))
    v15 = "\n\t"
    if v1 {
        v33 := GenS2{F0: (v25 << ((4294967295) & 63)), F1: "foo bar", F2: 910}
        _ = v33
    }
    v26 = v26
    var v34 string = ((v15 + v15) + "hello")
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v2
    _ = v25
    _ = v26
    _ = v29
    _ = v3
    _ = v34
    _ = v4
}
