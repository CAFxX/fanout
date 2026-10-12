package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64; F2 uint8 }
type GenS2 struct { F0 uint8; F1 int32; F2 uint32 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint32; F1 bool; F2 string }
func genFn0(p0 int) uint64 {
    v4 := []int{4294967295, (p0 ^ (-(p0)))}
    var v5 float64 = ((1.5 * ((1.5 + -2.25) / 1.5)) + 0.0)
    _ = p0
    _ = v4
    _ = v5
    return 29
}
func genFn1(p0 int32, p1 string, p2 uint32) int8 {
    var v6 string = ("" + "hello")
    switch 4294967295 {
    case 9223372036854775807:
        var v7 int64 = 9223372036854775807
        _ = v7
    case 4294967295:
        if ((4294967295 <= 391) || (!(false || false))) {
            v6 = p1
            v8 := GenS1{F0: 292, F1: 1, F2: 128}
            p1 = p1
            _ = v8
        }
    case -892:
        v9 := []int{947, 42, -2147483648}
        _ = v9
    default:
        p0 -= p0
    }
    p1 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return -128
}

func main() {
    var v1 bool = true
    v2 := []uint32{65535, 389, 0}
    v3 := make(map[string]uint32)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    _ = len(v2)
    for fi10 := 0; fi10 < 8; fi10++ {
        fmt.Println(v1, fi10)
        _ = fi10
    }
    if v1 {
        var v11 int16 = 1
        _ = v11
    }
    v12 := GenS1{F0: 42, F1: 370, F2: 1}
    switch "a" {
    case "\n\t":
        if true {
            var v13 float32 = ((float32(1.5) / float32(1.5)) / 1e10)
            if v1 {
                v13 -= v13
                switch ("" + "") {
                case "":
                    var v14 int = -263
                    _ = v14
                case "foo bar":
                    v13 -= v13
                case "a":
                    switch (("hello" + "hello") + "\n\t") {
                    case "hello":
                        v15 := GenS1{F0: 18446744073709551615, F1: 273, F2: 69}
                        _ = v15
                    case "foo bar":
                        v13 += float32(float64(0.1))
                    case "\n\t":
                        for fi16 := 0; fi16 < 7; fi16++ {
                            var v17 int16 = -109
                            v1 = v1
                            v13 = 3.14159
                            _ = fi16
                            _ = v17
                        }
                    }
                }
                for fi18 := 0; fi18 < 12; fi18++ {
                    v1 = v1
                    _ = fi18
                }
            } else {
                var v19 int32 = -2147483648
                _ = v19
            }
            fmt.Println(v1, v1)
            _ = v13
        }
    case "a":
        var v20 int = 127
        _ = v20
    case "":
        v21 := GenS2{F0: 127, F1: -2147483648, F2: 2147483648}
        _ = v21
    default:
        for fi22 := 0; fi22 < 12; fi22++ {
            for fi23 := 0; fi23 < 6; fi23++ {
                var v24 uint64 = 1
                v1 = v1
                v25 := GenS2{F0: 128, F1: 2147483647, F2: 2147483647}
                _ = fi23
                _ = v24
                _ = v25
            }
            if (false && v1) {
                var v26 int8 = genFn1(636, "foo bar", 765)
                for fi27 := 0; fi27 < 16; fi27++ {
                    var v28 int8 = 36
                    if len(v2) > 0 {
                        v2[0] = 2
                    }
                    v28 = (v26 << ((531) & 63))
                    var v29 int16 = -1
                    _ = fi27
                    _ = v28
                    _ = v29
                }
                _ = v26
            }
            _ = fi22
        }
    }
    v30 := GenS1{F0: 1, F1: 0, F2: 42}
    var v31 uint8 = 0
    v32 := GenS3{F0: 134}
    for fi33 := 0; fi33 < 7; fi33++ {
        var v34 float64 = 0.1
        var v35 bool = true
        _ = fi33
        _ = v34
        _ = v35
    }
    _ = len(v2)
    v36 := GenS4{F0: 604, F1: (true || (float32(((1e10 * 1e10) / 0.1)) != 1e10)), F2: "hello"}
    if v1 {
        v37 := v3[("foo bar" + ("a" + "\n\t"))]
        if len(v2) > 0 {
            v2[0] = 42
        }
        _ = v37
    }
    v31 -= (v31 - (v31 ^ (v31 - v31)))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v36
}
