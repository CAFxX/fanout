package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int; F2 int32 }
type GenS2 struct { F0 int64; F1 string; F2 int }
func genFn0() uint {
    var v7 int64 = 1
    v8 := []int{128, 32767, -614}
    v7 -= v7
    _ = v7
    _ = v8
    return 65535
}
func genFn1(p0 int64) uint64 {
    var v9 int8 = 100
    switch -312 {
    case 32767:
        for fi10 := 0; fi10 < 9; fi10++ {
            v11 := GenS1{F0: v9, F1: fi10, F2: -877}
            v12 := []int{}
            _ = fi10
            _ = v11
            _ = v12
        }
    case 1:
        if false {
            fmt.Println(v9, p0)
            fmt.Println(v9, p0)
        } else {
            if (!true) {
                fmt.Println(p0)
                fmt.Println(p0, p0)
                for fi13 := 0; fi13 < 3; fi13++ {
                    v14 := GenS1{F0: v9, F1: -845, F2: 2147483647}
                    _ = fi13
                    _ = v14
                }
            } else {
                var v15 float32 = 1.5
                var v16 uint16 = 32767
                _ = v15
                _ = v16
            }
        }
    default:
        v17 := GenS1{F0: v9, F1: 128, F2: -836}
        _ = v17
    }
    if true {
        var v18 int8 = v9
        _ = v18
    }
    if false {
        for fi19 := 0; fi19 < 6; fi19++ {
            var v20 int64 = (p0 / (((p0 >> ((18446744073709551615) & 63))) | 1))
            _ = fi19
            _ = v20
        }
    } else {
        fmt.Println(v9)
    }
    v21 := GenS1{F0: (v9 ^ v9), F1: 1, F2: -340}
    _ = p0
    _ = v21
    _ = v9
    return 4294967295
}

func main() {
    var v1 uint = 65535
    v2 := make(map[uint32]uint32)
    v3 := make(map[uint32]bool)
    var v4 int64 = 454
    v5 := make(map[uint32]string)
    v6 := []string{("\n\t" + "foo bar"), ("\n\t" + ("a" + ""))}
    v1 = 42
    for fi22 := 0; fi22 < 8; fi22++ {
        switch ((("foo bar" + "\n\t") + "foo bar") + ("" + "\n\t")) {
        case "\n\t":
            v6 = append(v6, "a")
        }
        for fi23 := 0; fi23 < 13; fi23++ {
            if len(v6) > 0 {
                v6[0] = "\n\t"
            }
            fmt.Println(v1, v1)
            v24 := GenS1{F0: 56, F1: fi22, F2: -647}
            _ = fi23
            _ = v24
        }
        _ = fi22
    }
    var v25 float64 = (((1.5 - 1.5) / (-2.25 + 1.5)) / (0.0 + ((0.1 - ((0.1 * -2.25) + (0.0 - -2.25))) / 1.0)))
    switch 77 {
    case 4294967295:
        if true {
            v4 -= (v4 << ((124) & 63))
        } else {
            var v26 uint = 0
            _ = v26
        }
    default:
        v27 := v3[uint32(v1)]
        _ = v27
    }
    switch "hello" {
    case "hello":
        for fi28 := 0; fi28 < 15; fi28++ {
            var v29 float32 = (-2.25 / (3.14159 * 3.14159))
            v6 = append(v6, "a")
            fmt.Println(fi28, fi28, v1)
            _ = fi28
            _ = v29
        }
    case "a":
        var v30 int16 = -1
        _ = v30
    case "foo bar":
        v31 := GenS1{F0: -128, F1: -1, F2: -2147483648}
        _ = v31
    }
    v32 := GenS1{F0: 21, F1: 32768, F2: -2147483648}
    switch "foo bar" {
    case "foo bar":
        if true {
            v33 := GenS1{F0: 42, F1: 2147483647, F2: -123}
            var v34 int = 127
            v35 := GenS1{F0: -75, F1: v34, F2: -2147483648}
            _ = v33
            _ = v34
            _ = v35
        } else {
            var v36 uint = genFn0()
            _ = v36
        }
    default:
        fmt.Println(v1)
    }
    v37 := GenS2{F0: (v4 + (v4 % (((v4 + v4)) | 1))), F1: "", F2: 42}
    _ = len(v6)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v32
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
