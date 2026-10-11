package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string; F2 uint16 }
type GenS2 struct { F0 int64; F1 int32; F2 int32 }
type GenS3 struct { F0 uint; F1 int64; F2 uint }
func genFn0(p0 int16) bool {
    v4 := []int{127, -444, 9223372036854775807, 32767}
    var v5 uint32 = 4294967295
    v5 = v5
    for fi6 := 0; fi6 < 12; fi6++ {
        var v7 int8 = 112
        _ = fi6
        _ = v7
    }
    fmt.Println(v5)
    _ = p0
    _ = v4
    _ = v5
    return true
}
func genFn1(p0 uint, p1 int, p2 int64) int32 {
    fmt.Println(p1)
    var v8 uint8 = 144
    for fi9 := 0; fi9 < 4; fi9++ {
        for fi10 := 0; fi10 < 8; fi10++ {
            switch ("" + ("a" + "a")) {
            case "":
                fmt.Println(v8, v8)
            }
            fmt.Println(p1, p2, fi9)
            fmt.Println(fi9, fi9, fi9)
            _ = fi10
        }
        v11 := GenS1{F0: 1, F1: "foo bar", F2: 2}
        var v12 int64 = (p2 - (p2 * (p2 / ((p2) | 1))))
        _ = fi9
        _ = v11
        _ = v12
    }
    var v13 bool = ((("hello" + ("" + "a")) + "foo bar") < "hello")
    p0 = 18446744073709551615
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v8
    return int32(p2)
}

func main() {
    var v1 uint8 = 100
    var v2 int16 = 0
    v3 := make(map[string]uint32)
    v2 -= 32767
    var v14 uint64 = 0
    fmt.Println(v1, v2, v14)
    if false {
        v1 -= v1
    }
    v15 := []int{314, 128}
    for fi16 := 0; fi16 < 20; fi16++ {
        var v17 int64 = 0
        _ = fi16
        _ = v17
    }
    for fi18 := 0; fi18 < 4; fi18++ {
        switch ("\n\t" + ("a" + "hello")) {
        case "a":
            if false {
                v19 := v3[""]
                _ = v19
            } else {
                if (false || (-69 <= -128)) {
                    v20 := GenS2{F0: -616, F1: 42, F2: -1}
                    var v21 string = "hello"
                    _ = v20
                    _ = v21
                }
            }
        case "hello":
            for fi22 := 0; fi22 < 4; fi22++ {
                var v23 uint64 = 42
                v3["hello"] = uint32(v1)
                fmt.Println(fi22, v14, v23)
                _ = fi22
                _ = v23
            }
        case "\n\t":
            for fi24 := 0; fi24 < 16; fi24++ {
                if len(v15) > 0 {
                    v15[0] = fi18
                }
                var v25 float64 = -2.25
                var v26 int = 166
                _ = fi24
                _ = v25
                _ = v26
            }
        default:
            v27 := GenS1{F0: v2, F1: "a", F2: 928}
            _ = v27
        }
        if false {
            var v28 uint32 = 278
            _ = v28
        } else {
            fmt.Println(fi18, v2)
            v29 := v3["hello"]
            var v30 int8 = -70
            _ = v29
            _ = v30
        }
        var v31 int64 = -607
        _ = fi18
        _ = v31
    }
    v14 = uint64(v1)
    _ = len(v3)
    v14 += v14
    v14 = 79
    _ = len(v15)
    if len(v15) > 0 {
        v15[0] = 32768
    }
    v32 := GenS2{F0: -9223372036854775808, F1: -2147483648, F2: 0}
    var v33 int64 = -9223372036854775808
    switch ("hello" + "hello") {
    case "a":
        switch "hello" {
        case "foo bar":
            for fi34 := 0; fi34 < 18; fi34++ {
                v2 = v2
                for fi35 := 0; fi35 < 16; fi35++ {
                    fmt.Println(fi34, v1)
                    v2 = (v2 & v2)
                    _ = fi35
                }
                var v36 int8 = 0
                _ = fi34
                _ = v36
            }
        case "hello":
            v37 := GenS3{F0: 318, F1: v33, F2: 54}
            _ = v37
        }
    case "":
        v38 := GenS1{F0: v2, F1: "foo bar", F2: 42}
        _ = v38
    case "foo bar":
        v39 := GenS1{F0: v2, F1: "\n\t", F2: 1}
        _ = v39
    }
    fmt.Println(v33, v14)
    if (-1 >= -1) {
        v1 = (v1 | (v1 << ((585) & 63)))
        _ = len(v3)
    } else {
        for fi40 := 0; fi40 < 6; fi40++ {
            if len(v15) >= 2 {
                v15 = v15[1:len(v15)]
            }
            _ = fi40
        }
        v41 := GenS3{F0: uint(v1), F1: v33, F2: 4294967295}
        _ = v41
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v32
    _ = v33
}
