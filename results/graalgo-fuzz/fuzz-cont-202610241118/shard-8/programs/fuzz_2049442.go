package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 bool; F1 int16; F2 int }
type GenS4 struct { F0 int; F1 int16 }
func genFn0(p0 uint8, p1 int16) string {
    for fi6 := 0; fi6 < 1; fi6++ {
        if true {
            p1 -= (^(p1))
        }
        for fi7 := 0; fi7 < 9; fi7++ {
            v8 := GenS1{F0: 2}
            p1 = 32767
            _ = fi7
            _ = v8
        }
        p0 = (p0 - p0)
        _ = fi6
    }
    for fi9 := 0; fi9 < 4; fi9++ {
        v10 := []int{206, fi9}
        _ = fi9
        _ = v10
    }
    _ = p0
    _ = p1
    return "\n\t"
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint64 = 2
    var v3 int64 = 9223372036854775807
    v4 := make(map[string]uint32)
    v5 := make(map[string]uint32)
    v11 := []int{}
    var v12 uint16 = 255
    var v13 uint16 = 377
    for fi14 := 0; fi14 < 14; fi14++ {
        v3 = (v3 + -1)
        _ = fi14
    }
    v15 := GenS1{F0: 1}
    var v16 float32 = ((1e10 / 1.0) + 0.0)
    if (0 > -2147483648) {
        v17 := GenS1{F0: uint64(v12)}
        var v18 int16 = 1
        fmt.Println(v18, v2)
        _ = v17
        _ = v18
    }
    if false {
        switch -1 {
        case 32767:
            v3 -= 42
        case 1:
            v19 := GenS2{F0: 493}
            _ = v19
        }
    } else {
        v20 := GenS3{F0: (127 >= int8(v12)), F1: 42, F2: 0}
        var v21 uint32 = 4294967295
        var v22 uint8 = 87
        _ = v20
        _ = v21
        _ = v22
    }
    switch (("a" + "\n\t") + ("\n\t" + "")) {
    case "foo bar":
        var v23 int64 = v3
        _ = v23
    case "\n\t":
        if (v3 >= (v3 + (v3 >> ((577) & 63)))) {
            v24 := v5["\n\t"]
            var v25 int = 32767
            var v26 int8 = 100
            _ = v24
            _ = v25
            _ = v26
        } else {
            v3 -= v3
            fmt.Println(v13, v16, v12)
        }
    }
    v13 += 491
    v27 := v5[("foo bar" + "\n\t")]
    v2 = (v2 - v2)
    v27 = uint32(v2)
    v28 := GenS4{F0: -209, F1: 1000}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v15
    _ = v16
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
