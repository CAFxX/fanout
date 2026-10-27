package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 string; F1 int16; F2 int8 }
type GenS3 struct { F0 uint16; F1 uint8; F2 uint }
type GenS4 struct { F0 bool }
func genFn0(p0 bool, p1 int16) int64 {
    var v7 int = -1
    if (42 >= 36) {
        v8 := GenS1{F0: -783}
        var v9 int32 = -27
        _ = v8
        _ = v9
    } else {
        fmt.Println(v7, v7)
        var v10 int8 = 22
        _ = v10
    }
    switch v7 {
    case 258:
        switch (v7 & (-(v7))) {
        case 32768:
            var v11 bool = p0
            _ = v11
        case 42:
            v12 := []int{0, 4294967295}
            _ = v12
        case -94:
            v13 := []int{}
            _ = v13
        default:
            v7 = (-(v7))
        }
    default:
        v7 = v7
    }
    var v14 uint8 = 255
    _ = p0
    _ = p1
    _ = v14
    _ = v7
    return 555
}

func main() {
    var v1 int32 = 42
    v2 := make(map[int]uint32)
    v3 := []int{-486, 42, 0, -748}
    v4 := make(map[string]uint32)
    v5 := make(map[string]uint32)
    var v6 int = -9223372036854775808
    fmt.Println(v1)
    var v15 float64 = 0.1
    if (1e10 != (3.14159 - (float64((0.1 - 3.14159)) - 0.0))) {
        for fi16 := 0; fi16 < 20; fi16++ {
            var v17 float32 = (1.5 + 1e10)
            var v18 string = ("hello" + "\n\t")
            _ = fi16
            _ = v17
            _ = v18
        }
        v5["\n\t"] = 636
    } else {
        v6 += (-(v6))
        fmt.Println(v15, v15, v6)
        var v19 uint8 = 1
        _ = v19
    }
    for fi20 := 0; fi20 < 12; fi20++ {
        var v21 bool = false
        var v22 uint = 42
        var v23 int = (fi20 * fi20)
        _ = fi20
        _ = v21
        _ = v22
        _ = v23
    }
    for fi24 := 0; fi24 < 9; fi24++ {
        v25 := GenS1{F0: v1}
        _ = fi24
        _ = v25
    }
    for fi26 := 0; fi26 < 9; fi26++ {
        var v27 float32 = 3.14159
        _ = fi26
        _ = v27
    }
    if false {
        v28 := GenS2{F0: "foo bar", F1: 42, F2: -25}
        _ = v28
    } else {
        v6 = (v6 * v6)
        for fi29 := 0; fi29 < 1; fi29++ {
            var v30 uint16 = 42
            _ = fi29
            _ = v30
        }
    }
    v6 -= (v6 & (v6 * (v6 << ((65535) & 63))))
    fmt.Println(v1)
    fmt.Println(v15)
    var v31 uint = 42
    if false {
        var v32 string = "a"
        var v33 bool = false
        v34 := GenS3{F0: 2, F1: 1, F2: 0}
        _ = v32
        _ = v33
        _ = v34
    } else {
        var v35 uint16 = 0
        _ = v35
    }
    for fi36 := 0; fi36 < 9; fi36++ {
        var v37 uint = v31
        if false {
            if (("" + "foo bar") == ("hello" + "hello")) {
                if true {
                    v38 := GenS4{F0: false}
                    _ = v38
                }
                var v39 uint64 = 0
                var v40 uint8 = 78
                _ = v39
                _ = v40
            }
        }
        _ = fi36
        _ = v37
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
