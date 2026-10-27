package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint32; F1 int16 }
func genFn0(p0 uint32) uint16 {
    var v4 uint = 4294967295
    if true {
        var v5 uint64 = 466
        for fi6 := 0; fi6 < 19; fi6++ {
            v7 := GenS1{F0: (^(v5))}
            _ = fi6
            _ = v7
        }
        v8 := GenS1{F0: v5}
        _ = v5
        _ = v8
    } else {
        v9 := GenS2{F0: ("hello" + "a")}
        _ = v9
    }
    for fi10 := 0; fi10 < 4; fi10++ {
        var v11 float64 = (0.1 - 0.0)
        _ = fi10
        _ = v11
    }
    v12 := GenS1{F0: 42}
    _ = p0
    _ = v12
    _ = v4
    return 468
}
func genFn1() uint32 {
    if (1 >= 775) {
        var v13 int32 = 689
        var v14 int8 = 25
        v13 -= v13
        _ = v13
        _ = v14
    } else {
        var v15 float64 = -2.25
        var v16 int64 = 0
        v16 -= v16
        _ = v15
        _ = v16
    }
    var v17 uint64 = 0
    v18 := []int{128, -2147483648, -528}
    fmt.Println(v17)
    var v19 uint16 = 0
    _ = v17
    _ = v18
    _ = v19
    return 2147483647
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[uint32]bool)
    v3 := make(map[uint32]uint32)
    switch ("foo bar" + ("a" + "foo bar")) {
    case "hello":
        var v20 uint64 = 4294967295
        _ = v20
    case "\n\t":
        var v21 float64 = 0.1
        _ = v21
    case "":
        var v22 int8 = 127
        _ = v22
    }
    var v23 float32 = (0.0 + (float32(float64(0.0)) * ((0.0 + 3.14159) - (3.14159 + 0.1))))
    v24 := GenS3{F0: 664, F1: 32767}
    v23 -= ((1.5 - (3.14159 + 1e10)) * -2.25)
    var v25 uint = 18446744073709551615
    for fi26 := 0; fi26 < 3; fi26++ {
        v23 -= (0.1 * (v23 + 0.1))
        if (3.14159 <= 0.1) {
            var v27 int16 = 16
            v28 := []int{(fi26 << ((4) & 63)), int(v27), fi26}
            v29 := GenS1{F0: 18446744073709551615}
            _ = v27
            _ = v28
            _ = v29
        }
        v25 = (v25 & (v25 | (v25 * 0)))
        _ = fi26
    }
    var v30 string = ""
    v31 := GenS1{F0: 188}
    fmt.Println(v30, v30, v30)
    for fi32 := 0; fi32 < 4; fi32++ {
        v33 := []int{42, (fi32 & (^(fi32))), (-(fi32)), (-(fi32))}
        var v34 int8 = 4
        v34 -= -1
        _ = fi32
        _ = v33
        _ = v34
    }
    v25 = 1
    var v35 int16 = 0
    v36 := GenS1{F0: 473}
    var v37 int32 = 1
    fmt.Println(v35)
    fmt.Println(v37, v25, v23)
    _ = len(v2)
    v38 := []int{114, -253}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v35:", v35)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v38:", len(v38), v38)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v30
    _ = v31
    _ = v35
    _ = v36
    _ = v37
    _ = v38
}
