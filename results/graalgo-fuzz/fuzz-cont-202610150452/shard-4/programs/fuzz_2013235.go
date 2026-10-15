package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 string; F1 uint64; F2 int }
type GenS4 struct { F0 int16; F1 uint8; F2 int8 }
type GenS5 struct { F0 uint8; F1 uint32 }
func genFn0(p0 int) uint32 {
    v7 := []int{(p0 ^ p0), p0}
    var v8 uint32 = 2147483648
    p0 += (p0 * p0)
    p0 += 32768
    _ = p0
    _ = v7
    _ = v8
    return (v8 + v8)
}
func genFn1() bool {
    var v9 float32 = 3.14159
    for fi10 := 0; fi10 < 20; fi10++ {
        if (!true) {
            if (0 < 574) {
                if (2147483647 > -540) {
                    fmt.Println(v9, v9)
                    switch (("foo bar" + "hello") + ("" + "")) {
                    case "":
                        v11 := GenS1{F0: (fi10 - (fi10 << ((61) & 63)))}
                        _ = v11
                    }
                }
            }
            v12 := []int{(^(fi10))}
            fmt.Println(fi10)
            _ = v12
        }
        var v13 uint64 = 1
        v14 := GenS1{F0: fi10}
        _ = fi10
        _ = v13
        _ = v14
    }
    _ = v9
    return (true || (!true))
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[int]string)
    var v3 int32 = 0
    v4 := make(map[string]bool)
    v5 := make(map[string]string)
    var v6 int = -1
    fmt.Println(v3, v6)
    var v15 int64 = -9223372036854775808
    var v16 uint8 = 127
    var v17 uint = 892
    var v18 bool = false
    v19 := GenS1{F0: v6}
    v20 := []int{v6, v6, (^(v6)), (v6 - (v6 / ((v6) | 1)))}
    v21 := GenS1{F0: 9223372036854775807}
    v22 := GenS2{F0: v15}
    switch (v6 - v6) {
    case -397:
        v23 := GenS3{F0: ("a" + (("foo bar" + "hello") + "")), F1: 42, F2: v6}
        _ = v23
    case -2147483648:
        v24 := GenS4{F0: -32768, F1: v16, F2: 1}
        _ = v24
    case 128:
        var v25 float32 = float32((0.1 - 0.1))
        _ = v25
    }
    for fi26 := 0; fi26 < 4; fi26++ {
        v3 -= -977
        v27 := v1["foo bar"]
        v28 := GenS5{F0: v16, F1: 42}
        _ = fi26
        _ = v27
        _ = v28
    }
    fmt.Println(v17, v16)
    for fi29 := 0; fi29 < 15; fi29++ {
        switch ("\n\t" + ("" + "\n\t")) {
        case "foo bar":
            v1[("" + (("foo bar" + "foo bar") + "foo bar"))] = ("" + "hello")
        case "a":
            var v30 uint16 = 255
            _ = v30
        case "\n\t":
            if (2 > 42) {
                v15 = v15
                for fi31 := 0; fi31 < 10; fi31++ {
                    if len(v20) > 0 {
                        v20[0] = fi29
                    }
                    fmt.Println(v6, v18, v3)
                    if ((fi31 & 128) != (fi31 << (((^(v17))) & 63))) {
                        var v32 float64 = 3.14159
                        _ = v32
                    }
                    _ = fi31
                }
                for fi33 := 0; fi33 < 4; fi33++ {
                    var v34 float64 = 3.14159
                    if false {
                        var v35 int16 = 32767
                        _ = v35
                    } else {
                        if v18 {
                            if (v3 < v3) {
                                fmt.Println(v6)
                            }
                            v36 := GenS1{F0: (fi33 & (fi29 + 2147483647))}
                            v15 = 9223372036854775807
                            _ = v36
                        }
                        fmt.Println(v6)
                        delete(v2, (v6 * 127))
                    }
                    v6 += (v6 * fi29)
                    _ = fi33
                    _ = v34
                }
            } else {
                v37 := GenS2{F0: v15}
                var v38 float32 = 0.1
                if v18 {
                    fmt.Println(v17)
                    var v39 int32 = (v3 << ((2) & 63))
                    _ = v39
                }
                _ = v37
                _ = v38
            }
        default:
            v3 -= v3
        }
        v40 := v1[""]
        var v41 int = (v6 & fi29)
        _ = fi29
        _ = v40
        _ = v41
    }
    var v42 int64 = v15
    v6 = (v6 >> (((v17 - 255)) & 63))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v42
    _ = v5
    _ = v6
}
