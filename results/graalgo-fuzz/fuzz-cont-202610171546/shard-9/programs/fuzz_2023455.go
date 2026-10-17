package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64 }
type GenS2 struct { F0 uint; F1 int8; F2 uint64 }
type GenS3 struct { F0 int64; F1 uint8; F2 uint8 }
type GenS4 struct { F0 int16; F1 int64; F2 int32 }
type GenS5 struct { F0 uint16; F1 int32; F2 int32 }
func genFn0(p0 bool, p1 string, p2 uint16) string {
    if (p0 && ("hello" == (p1 + p1))) {
        v5 := GenS1{F0: p2, F1: 78}
        _ = v5
    } else {
        fmt.Println(p1, p0)
        if p0 {
            p1 = p1
        } else {
            var v6 uint64 = 501
            _ = v6
        }
        var v7 int64 = -1
        _ = v7
    }
    v8 := GenS2{F0: 262, F1: -48, F2: 285}
    p2 += p2
    switch 32767 {
    case -2147483648:
        p0 = (p0 || p0)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return ((((p1 + (("foo bar" + (p1 + ((((p1 + ("\n\t" + "hello")) + (p1 + (p1 + (p1 + p1)))) + "hello") + p1))) + ((((((p1 + p1) + p1) + p1) + "hello") + "\n\t") + (p1 + p1)))) + p1) + "\n\t") + (p1 + p1))
}

func main() {
    var v1 float64 = (((((0.0 - -2.25) - (-2.25 + float64(1e10))) / 0.1) / (3.14159 + 1e10)) - (0.0 + 0.1))
    var v2 bool = true
    v3 := []uint32{836, 2147483647, 65535, 65535}
    var v4 int32 = 0
    fmt.Println(v2)
    for fi9 := 0; fi9 < 19; fi9++ {
        v2 = v2
        var v10 int16 = -1
        if (v4 != v4) {
            if true {
                v1 += v1
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            } else {
                fmt.Println(v1, v2, fi9)
                for fi11 := 0; fi11 < 6; fi11++ {
                    v12 := GenS3{F0: -1, F1: 81, F2: 1}
                    _ = len(v3)
                    _ = fi11
                    _ = v12
                }
            }
            if len(v3) > 0 {
                v3[0] = 1
            }
            if len(v3) > 0 {
                v3[0] = 1
            }
        }
        _ = fi9
        _ = v10
    }
    v13 := GenS4{F0: 1000, F1: -34, F2: (v4 & 124)}
    var v14 uint64 = 416
    if v2 {
        var v15 string = "hello"
        if v2 {
            if len(v3) > 0 {
                v3[0] = 42
            }
        } else {
            v14 = v14
            fmt.Println(v15, v1, v14)
        }
        switch "\n\t" {
        case "":
            fmt.Println(v14, v14)
        case "hello":
            fmt.Println(v4)
        case "foo bar":
            v16 := GenS3{F0: -1, F1: 172, F2: 127}
            _ = v16
        }
        _ = v15
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if v2 {
        v3 = append(v3, 2147483647)
        fmt.Println(v2, v4)
        v17 := GenS1{F0: 42, F1: v14}
        _ = v17
    } else {
        v2 = (2 <= 111)
    }
    v2 = true
    var v18 int = 128
    if (4294967295 > 255) {
        var v19 uint16 = 1
        v20 := GenS5{F0: 2, F1: -2147483648, F2: (v4 * 2147483647)}
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = v19
        _ = v20
    } else {
        var v21 int = 364
        var v22 uint32 = 90
        var v23 float64 = v1
        _ = v21
        _ = v22
        _ = v23
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
