package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64 }
type GenS2 struct { F0 int64; F1 uint64; F2 int }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int32; F1 uint8; F2 string }
func genFn0(p0 uint, p1 bool, p2 uint8) int64 {
    fmt.Println(p2, p2)
    if p1 {
        fmt.Println(p2, p1)
        fmt.Println(p2)
        p2 = 1
    } else {
        var v5 uint64 = 18446744073709551615
        _ = v5
    }
    _ = p0
    _ = p1
    _ = p2
    return -9223372036854775808
}

func main() {
    var v1 float64 = (1e10 + (3.14159 + float64(0.1)))
    var v2 int16 = 0
    v3 := make(map[uint32]int)
    var v4 string = ("" + "hello")
    v3[609] = 78
    v6 := GenS1{F0: 9223372036854775807, F1: 0}
    v1 = v1
    var v7 int = -738
    var v8 float64 = v1
    for fi9 := 0; fi9 < 15; fi9++ {
        var v10 float64 = float64(3.14159)
        _ = fi9
        _ = v10
    }
    v11 := GenS1{F0: genFn0(255, (4294967295 < 255), 255), F1: -682}
    switch v4 {
    case "hello":
        for fi12 := 0; fi12 < 3; fi12++ {
            fmt.Println(v2, v2)
            _ = fi12
        }
    case "foo bar":
        v2 = (v2 * v2)
    }
    v4 = v4
    for fi13 := 0; fi13 < 6; fi13++ {
        if (false || (933 < 4294967295)) {
            var v14 int32 = 2147483647
            if false {
                var v15 int16 = -432
                _ = v15
            }
            _ = v14
        } else {
            v8 = 1e10
        }
        v1 = 1e10
        _ = fi13
    }
    v16 := []int{v7, v7, (v7 * v7), v7}
    if ((-2.25 + ((0.1 / 1e10) / (0.1 + (3.14159 * (1.5 / float32(float64(1.5))))))) > (0.1 + (0.0 + (1e10 * (((-2.25 / -2.25) * 1.5) / 0.1))))) {
        for fi17 := 0; fi17 < 13; fi17++ {
            v18 := GenS1{F0: -1, F1: 140}
            for fi19 := 0; fi19 < 17; fi19++ {
                if (1 == 2) {
                    var v20 float32 = 0.0
                    var v21 uint16 = 1
                    switch v4 {
                    case "hello":
                        var v22 int64 = int64(fi19)
                        _ = v22
                    case "foo bar":
                        v4 = "a"
                    }
                    _ = v20
                    _ = v21
                } else {
                    v2 = v2
                    fmt.Println(v2)
                }
                v4 = ("" + v4)
                fmt.Println(fi19, fi17)
                v4 = ""
                _ = fi19
            }
            _ = fi17
            _ = v18
        }
        var v23 uint16 = 42
        _ = v23
    }
    if false {
        v8 += v8
        v24 := GenS1{F0: 403, F1: -1}
        v7 = (v7 ^ v7)
        _ = v24
    }
    for fi25 := 0; fi25 < 12; fi25++ {
        var v26 uint64 = 736
        if (true || false) {
            v7 -= fi25
        } else {
            var v27 bool = false
            switch v4 {
            case "\n\t":
                var v28 int = -9223372036854775808
                _ = v28
            default:
                v4 = v4
            }
            _ = v27
        }
        if ((v26 + 18446744073709551615) < v26) {
            switch v4 {
            case "":
                v8 += 0.1
            case "hello":
                var v29 float64 = v1
                _ = v29
            case "a":
                if false {
                    v30 := GenS2{F0: genFn0(2, (96 >= -642), 209), F1: (v26 * (v26 * (v26 - (v26 & (^(v26)))))), F2: 4294967295}
                    if len(v16) > 0 {
                        v16[0] = (fi25 * (^(v7)))
                    }
                    _ = v30
                } else {
                    v31 := GenS3{F0: 107}
                    v32 := GenS4{F0: -743, F1: 128, F2: ("a" + ("foo bar" + "foo bar"))}
                    _ = len(v3)
                    _ = v31
                    _ = v32
                }
            }
            fmt.Println(v2, v1)
            if (true && (905 <= 2)) {
                if (42 == 65535) {
                    var v33 uint = 42
                    v2 -= v2
                    _ = v33
                } else {
                    v34 := v3[42]
                    _ = v34
                }
            } else {
                for fi35 := 0; fi35 < 5; fi35++ {
                    v7 = fi35
                    if (255 != 42) {
                        var v36 uint32 = 157
                        _ = v36
                    }
                    if (42 < (fi25 - (fi35 + 24))) {
                        var v37 float32 = -2.25
                        var v38 uint32 = 2
                        if (!true) {
                            fmt.Println(v38, v4)
                        }
                        _ = v37
                        _ = v38
                    }
                    _ = fi35
                }
                v2 = (v2 - v2)
                var v39 uint32 = 478
                _ = v39
            }
        }
        _ = fi25
        _ = v26
    }
    for fi40 := 0; fi40 < 19; fi40++ {
        var v41 bool = false
        _ = fi40
        _ = v41
    }
    var v42 float32 = 0.1
    fmt.Println(v7, v42, v8)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v6
    _ = v7
    _ = v8
}
