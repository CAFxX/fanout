package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool; F2 string }
type GenS2 struct { F0 int32; F1 int32 }
type GenS3 struct { F0 bool; F1 int16 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int16; F1 int16 }
func genFn0(p0 uint32, p1 int) uint8 {
    v6 := []int{}
    if (true && true) {
        var v7 int8 = 0
        if (65535 < p0) {
            var v8 string = ("a" + (("hello" + "\n\t") + (("foo bar" + (("a" + "foo bar") + "\n\t")) + (("hello" + "foo bar") + ((("a" + (("a" + ("a" + ("foo bar" + "\n\t"))) + "")) + "") + ("" + "foo bar"))))))
            p1 -= (p1 >> ((559) & 63))
            _ = v8
        } else {
            v9 := GenS1{F0: 42, F1: (-401 <= -1), F2: ""}
            v10 := GenS2{F0: 2147483647, F1: 2147483647}
            _ = v10
            _ = v9
        }
        _ = v7
    } else {
        v11 := GenS3{F0: ((42 <= -9223372036854775808) && (false || true)), F1: 42}
        _ = v11
    }
    var v12 bool = true
    if len(v6) > 0 {
        v6[0] = 127
    }
    fmt.Println(v12, p1)
    _ = p0
    _ = p1
    _ = v12
    _ = v6
    return uint8(p0)
}
func genFn1(p0 int16) int8 {
    if true {
        v13 := []int{9223372036854775807, 472}
        _ = v13
    }
    switch "foo bar" {
    case "foo bar":
        p0 -= -166
    case "":
        if ((-(p0)) > p0) {
            p0 += p0
        }
    }
    if (467 == 1) {
        var v14 int64 = -9223372036854775808
        _ = v14
    } else {
        v15 := []int{}
        _ = v15
    }
    if (true && true) {
        if true {
            var v16 int64 = 242
            fmt.Println(p0)
            _ = v16
        }
    } else {
        v17 := GenS4{F0: -152}
        if true {
            var v18 int32 = 0
            _ = v18
        } else {
            v19 := GenS3{F0: (((("hello" + ("foo bar" + "a")) + "") + "a") <= ""), F1: (p0 & (p0 ^ -1))}
            v20 := []int{}
            var v21 int = 474
            _ = v19
            _ = v20
            _ = v21
        }
        v22 := []int{0, 128, 1}
        _ = v17
        _ = v22
    }
    var v23 int8 = 69
    _ = p0
    _ = v23
    return v23
}

func main() {
    var v1 uint8 = 34
    v2 := make(map[uint32]bool)
    v3 := make(map[int]int)
    var v4 bool = true
    v5 := make(map[int]uint32)
    for fi24 := 0; fi24 < 19; fi24++ {
        fmt.Println(v4)
        _ = fi24
    }
    var v25 int32 = -621
    fmt.Println(v4, v1, v4)
    v1 = v1
    var v26 uint8 = v1
    if true {
        v25 = (v25 | (-(v25)))
    } else {
        v3[-35] = 32767
    }
    fmt.Println(v1, v25, v4)
    for fi27 := 0; fi27 < 19; fi27++ {
        var v28 int16 = 508
        v1 += (v26 ^ (v26 / ((v1) | 1)))
        _ = fi27
        _ = v28
    }
    v29 := GenS3{F0: v4, F1: 600}
    fmt.Println(v1, v25, v1)
    v30 := []int{}
    var v31 int = -240
    v32 := GenS5{F0: 1000, F1: 1}
    var v33 uint16 = 0
    v4 = (1 < 100)
    var v34 int32 = (v25 >> ((1) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
}
