package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16 }
func genFn0(p0 uint64) int32 {
    if true {
        p0 = p0
        var v4 uint16 = 755
        _ = v4
    } else {
        var v5 int8 = -31
        p0 = p0
        var v6 int16 = 0
        _ = v5
        _ = v6
    }
    v7 := []int{}
    var v8 string = ("a" + ("hello" + ("a" + "\n\t")))
    _ = p0
    _ = v7
    _ = v8
    return -2147483648
}
func genFn1(p0 uint32, p1 bool) uint {
    v9 := []int{33}
    if ((p0 % (((p0 + (p0 * p0))) | 1)) >= 2147483648) {
        switch "foo bar" {
        case "":
            _ = len(v9)
        }
    }
    for fi10 := 0; fi10 < 13; fi10++ {
        p1 = p1
        for fi11 := 0; fi11 < 14; fi11++ {
            p1 = true
            p0 = p0
            p0 += (p0 / (((p0 - 0)) | 1))
            _ = fi11
        }
        _ = fi10
    }
    v9 = append(v9, 544)
    _ = p0
    _ = p1
    _ = v9
    return 540
}

func main() {
    var v1 int = 128
    v2 := make(map[string]uint32)
    v3 := []string{"foo bar"}
    v1 += (v1 | (v1 / ((v1) | 1)))
    v1 += -831
    var v12 bool = (1 != 213)
    v1 = (v1 % ((v1) | 1))
    if v12 {
        fmt.Println(v1, v12)
        var v13 string = "a"
        _ = v13
    }
    if v12 {
        v12 = ((4294967295 >= 42) || v12)
    } else {
        var v14 int32 = -1
        if (2 <= 210) {
            v15 := GenS1{F0: 9223372036854775807, F1: 65535}
            _ = v15
        } else {
            v16 := v2[("\n\t" + "a")]
            var v17 uint16 = 255
            _ = v16
            _ = v17
        }
        _ = v14
    }
    v2["foo bar"] = 629
    var v18 int = 1
    v1 = (-(v1))
    if (343 >= 255) {
        v19 := GenS1{F0: 42, F1: 43}
        v12 = v12
        delete(v2, "")
        _ = v19
    }
    for fi20 := 0; fi20 < 13; fi20++ {
        v21 := v2["a"]
        _ = fi20
        _ = v21
    }
    for fi22 := 0; fi22 < 13; fi22++ {
        if v12 {
            var v23 float64 = 1.5
            var v24 uint32 = 4294967295
            switch "a" {
            case "":
                v1 = v1
            }
            _ = v23
            _ = v24
        }
        v1 = (fi22 | (^(fi22)))
        v25 := GenS1{F0: -993, F1: 698}
        _ = fi22
        _ = v25
    }
    var v26 float32 = float32(((1.5 + (-2.25 / 1.5)) - (1e10 / 1.5)))
    v26 += (((v26 * v26) / v26) / 1e10)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v12
    _ = v18
    _ = v2
    _ = v26
    _ = v3
}
