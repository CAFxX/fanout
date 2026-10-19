package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 bool; F1 int16 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint64; F1 int }
type GenS6 struct { F0 string; F1 uint16 }
type GenS7 struct { F0 uint16 }
func genFn0(p0 uint8, p1 int16, p2 int8) uint64 {
    p2 = p2
    p2 = p2
    v6 := []int{947}
    var v7 bool = true
    var v8 float32 = (float32(float64(0.1)) - (float32(0.1) * 0.0))
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return 67
}
func genFn1(p0 string, p1 int16) int8 {
    if ((!(123 > 225)) || true) {
        switch p0 {
        case "a":
            for fi9 := 0; fi9 < 4; fi9++ {
                v10 := GenS1{F0: false, F1: (("\n\t" + "a") + p0)}
                _ = fi9
                _ = v10
            }
        case "":
            v11 := GenS1{F0: (42 != -2147483648), F1: (p0 + (p0 + p0))}
            _ = v11
        case "hello":
            p0 = (p0 + "foo bar")
        }
    } else {
        if true {
            v12 := GenS1{F0: (42 != 2), F1: p0}
            v13 := GenS1{F0: true, F1: p0}
            v14 := GenS1{F0: (0.1 >= (1.5 * float64(3.14159))), F1: (("a" + p0) + p0)}
            _ = v12
            _ = v13
            _ = v14
        }
    }
    v15 := GenS1{F0: false, F1: p0}
    var v16 float64 = (((float64(1e10) * 3.14159) - ((0.0 * -2.25) / 1.0)) - (1e10 / 1.5))
    var v17 int8 = -128
    _ = p0
    _ = p1
    _ = v15
    _ = v16
    _ = v17
    return 42
}

func main() {
    var v1 uint = 854
    var v2 bool = true
    var v3 string = "a"
    v4 := make(map[int]bool)
    v5 := []uint32{1, 65535}
    fmt.Println(v2)
    v18 := GenS2{F0: 0}
    var v19 uint16 = 65535
    for fi20 := 0; fi20 < 4; fi20++ {
        v2 = (v19 >= (v19 >> ((255) & 63)))
        switch (v3 + v3) {
        case "hello":
            _ = len(v5)
        case "foo bar":
            switch (v3 + v3) {
            case "foo bar":
                fmt.Println(v1, fi20)
            case "":
                if (int32(v19) == 2147483647) {
                    v2 = (v1 < (v1 * 1))
                    if (v2 && (((v19 * v19) >= (v19 | (v19 / (((v19 + 951)) | 1)))) || v2)) {
                        v21 := GenS2{F0: 957}
                        var v22 uint16 = v19
                        _ = v21
                        _ = v22
                    }
                    v2 = false
                }
            case "a":
                v3 = "foo bar"
            default:
                v3 = v3
            }
        }
        _ = fi20
    }
    _ = len(v5)
    v2 = false
    if true {
        var v23 uint8 = 128
        _ = v23
    } else {
        for fi24 := 0; fi24 < 5; fi24++ {
            v1 -= v1
            fmt.Println(v2, v2)
            _ = fi24
        }
        v19 = 560
    }
    v2 = v2
    switch 127 {
    case 0:
        var v25 string = "\n\t"
        _ = v25
    case 128:
        var v26 int8 = genFn1(v3, 1)
        _ = v26
    default:
        switch 1 {
        case 0:
            v19 += (v19 * 2)
        }
    }
    if (true && false) {
        var v27 uint8 = 36
        _ = v27
    }
    var v28 bool = (0 >= 18446744073709551615)
    var v29 float32 = (float32(((-2.25 / 3.14159) + 1.5)) + float32(1.5))
    var v30 string = v3
    v31 := GenS3{F0: v2, F1: 0}
    switch v3 {
    case "":
        v32 := GenS4{F0: (128 <= 1)}
        _ = v32
    default:
        for fi33 := 0; fi33 < 13; fi33++ {
            v34 := GenS5{F0: 2, F1: 592}
            for fi35 := 0; fi35 < 17; fi35++ {
                v36 := GenS6{F0: ((v3 + "foo bar") + "\n\t"), F1: v19}
                var v37 bool = v2
                if true {
                    if v2 {
                        v30 = v3
                        v1 = v1
                        v4[(fi35 * fi35)] = false
                    } else {
                        var v38 int16 = 42
                        v5 = append(v5, 2147483647)
                        v37 = (v38 != 42)
                        _ = v38
                    }
                } else {
                    v2 = (v1 != (v1 % ((v1) | 1)))
                    v3 = "foo bar"
                    _ = len(v5)
                    v19 += uint16(v1)
                }
                _ = fi35
                _ = v36
                _ = v37
            }
            switch ((((v30 + v3) + "\n\t") + "") + "hello") {
            case "foo bar":
                for fi39 := 0; fi39 < 10; fi39++ {
                    fmt.Println(fi33, v30)
                    if (false || v28) {
                        v40 := GenS1{F0: v28, F1: v30}
                        _ = len(v5)
                        _ = v40
                    }
                    var v41 uint8 = 255
                    v29 = v29
                    _ = fi39
                    _ = v41
                }
            case "\n\t":
                var v42 float64 = 1.5
                _ = v42
            case "":
                var v43 int = -2147483648
                _ = v43
            }
            _ = fi33
            _ = v34
        }
    }
    v4[0] = (!v2)
    for fi44 := 0; fi44 < 6; fi44++ {
        var v45 uint8 = 127
        fmt.Println(v2, v19, v30)
        switch 1 {
        case 20:
            v2 = false
        case 4294967295:
            v46 := GenS7{F0: v19}
            _ = v46
        }
        _ = fi44
        _ = v45
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
}
