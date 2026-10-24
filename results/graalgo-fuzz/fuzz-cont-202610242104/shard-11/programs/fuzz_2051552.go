package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8; F2 int }
type GenS2 struct { F0 int8; F1 string }
func genFn0(p0 uint, p1 int8) int {
    switch (("a" + "\n\t") + "hello") {
    case "hello":
        var v5 int32 = 42
        _ = v5
    }
    v6 := []int{32767, 32768, 350}
    var v7 uint32 = 0
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    switch 0 {
    case 128:
        var v8 uint8 = 128
        _ = v8
    case 127:
        for fi9 := 0; fi9 < 18; fi9++ {
            p1 -= 1
            var v10 int = fi9
            var v11 float64 = -2.25
            _ = fi9
            _ = v10
            _ = v11
        }
    default:
        if true {
            p1 = 0
        } else {
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            p1 += p1
            v12 := GenS1{F0: (^(p0)), F1: int8(p0), F2: 9223372036854775807}
            _ = v12
        }
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 32767
}
func genFn1() int16 {
    var v13 uint64 = 0
    var v14 uint32 = 4294967295
    switch "foo bar" {
    case "\n\t":
        switch 127 {
        case 9223372036854775807:
            v14 = 652
        case 32767:
            var v15 int16 = int16(v13)
            _ = v15
        }
    case "":
        v14 -= 2147483647
    case "a":
        for fi16 := 0; fi16 < 19; fi16++ {
            v17 := []int{(fi16 | 32767)}
            _ = fi16
            _ = v17
        }
    default:
        if (!false) {
            v14 -= v14
            var v18 float32 = (1.5 + 0.0)
            var v19 bool = (true && true)
            _ = v18
            _ = v19
        }
    }
    v20 := GenS1{F0: 560, F1: 45, F2: 4294967295}
    _ = v13
    _ = v14
    _ = v20
    return 0
}

func main() {
    var v1 int64 = 0
    v2 := make(map[string]bool)
    v3 := make(map[string]int)
    v4 := make(map[uint32]uint32)
    v1 = v1
    var v21 uint32 = 0
    if (0 > 255) {
        v22 := GenS2{F0: int8(v21), F1: "a"}
        v4[v21] = v21
        _ = v22
    } else {
        fmt.Println(v1)
    }
    v23 := []int{}
    delete(v3, "foo bar")
    fmt.Println(v21, v1)
    var v24 int8 = -1
    if len(v23) > 0 {
        v23[0] = 912
    }
    v24 -= (v24 >> ((55) & 63))
    var v25 int64 = (v1 * -815)
    switch 127 {
    case 746:
        fmt.Println(v1, v25, v24)
    }
    var v26 uint64 = 654
    v27 := v3[("foo bar" + (("foo bar" + "hello") + ""))]
    var v28 bool = false
    fmt.Println(v25, v1)
    v1 = (v1 << ((881) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v21
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
}
