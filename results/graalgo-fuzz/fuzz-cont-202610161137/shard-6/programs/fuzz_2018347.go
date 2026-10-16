package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint16; F1 uint16; F2 uint32 }
type GenS3 struct { F0 string; F1 uint16; F2 string }
type GenS4 struct { F0 uint64; F1 int64; F2 int32 }
func genFn0(p0 string) int16 {
    v7 := GenS1{F0: 32767}
    var v8 int32 = 158
    var v9 int16 = -965
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return (v9 - 1)
}
func genFn1(p0 int8) bool {
    fmt.Println(p0)
    var v10 string = ""
    for fi11 := 0; fi11 < 3; fi11++ {
        v12 := GenS2{F0: 1, F1: 166, F2: 2147483647}
        _ = fi11
        _ = v12
    }
    var v13 bool = (17 < 255)
    _ = p0
    _ = v10
    _ = v13
    return (137 < 127)
}

func main() {
    var v1 uint64 = 677
    var v2 bool = false
    v3 := []int{-2147483648, 224}
    v4 := make(map[string]bool)
    var v5 int32 = -890
    var v6 uint32 = 2147483647
    switch "" {
    case "":
        v1 = (v1 | v1)
    }
    for fi14 := 0; fi14 < 12; fi14++ {
        v15 := v4[""]
        _ = fi14
        _ = v15
    }
    v2 = (("a" + ("foo bar" + "hello")) != "foo bar")
    fmt.Println(v2, v6)
    if (-949 < -245) {
        v16 := GenS3{F0: ("foo bar" + "foo bar"), F1: 0, F2: "foo bar"}
        v5 = (^(v5))
        var v17 uint8 = 128
        _ = v16
        _ = v17
    }
    for fi18 := 0; fi18 < 14; fi18++ {
        v19 := GenS4{F0: 0, F1: 1, F2: (v5 / (((v5 >> ((18446744073709551615) & 63))) | 1))}
        v20 := GenS4{F0: v1, F1: int64(fi18), F2: (v5 | (v5 + v5))}
        _ = fi18
        _ = v19
        _ = v20
    }
    v21 := GenS4{F0: 0, F1: 9223372036854775807, F2: -2147483648}
    v1 = v1
    switch -1 {
    case -1:
        delete(v4, "hello")
    }
    for fi22 := 0; fi22 < 14; fi22++ {
        v23 := v4[("foo bar" + "foo bar")]
        _ = fi22
        _ = v23
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
