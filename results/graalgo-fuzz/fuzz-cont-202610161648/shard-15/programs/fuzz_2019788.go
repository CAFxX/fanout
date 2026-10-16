package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int8; F1 uint64; F2 string }
func genFn0(p0 uint64) uint64 {
    v4 := GenS1{F0: ("a" + "a")}
    for fi5 := 0; fi5 < 19; fi5++ {
        p0 += p0
        p0 -= p0
        if false {
            for fi6 := 0; fi6 < 13; fi6++ {
                v7 := []int{-9223372036854775808, int(p0)}
                _ = fi6
                _ = v7
            }
            switch "hello" {
            case "a":
                fmt.Println(p0, p0)
            }
            v8 := GenS1{F0: ""}
            _ = v8
        } else {
            var v9 uint32 = 1
            var v10 int32 = 573
            _ = v10
            _ = v9
        }
        _ = fi5
    }
    v11 := GenS2{F0: 127, F1: (p0 / (((p0 ^ 42)) | 1)), F2: (("" + ("" + "\n\t")) + "")}
    _ = p0
    _ = v11
    _ = v4
    return p0
}
func genFn1() bool {
    var v12 bool = (42 != 0)
    var v13 int64 = 1
    _ = v12
    _ = v13
    return true
}

func main() {
    v1 := make(map[int]int)
    v2 := []uint32{2147483647, 2147483648, 433, 1}
    v3 := make(map[string]string)
    if len(v2) > 0 {
        v2[0] = 0
    }
    var v14 int8 = -1
    v14 -= v14
    if true {
        v14 -= v14
    }
    var v15 float64 = 0.0
    for fi16 := 0; fi16 < 9; fi16++ {
        for fi17 := 0; fi17 < 13; fi17++ {
            switch (-(fi17)) {
            case 32767:
                v18 := GenS2{F0: 42, F1: genFn0(18446744073709551615), F2: "\n\t"}
                _ = v18
            case 9223372036854775807:
                var v19 int32 = -652
                _ = v19
            }
            v15 = float64(0.1)
            var v20 int16 = 1000
            _ = fi17
            _ = v20
        }
        v21 := GenS1{F0: "foo bar"}
        _ = fi16
        _ = v21
    }
    for fi22 := 0; fi22 < 9; fi22++ {
        var v23 int16 = -1
        var v24 uint32 = 818
        var v25 int = -9223372036854775808
        _ = fi22
        _ = v23
        _ = v24
        _ = v25
    }
    v15 += v15
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
}
