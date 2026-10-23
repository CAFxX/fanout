package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32; F2 uint16 }
type GenS2 struct { F0 bool; F1 int16; F2 uint16 }
func genFn0(p0 uint64, p1 int) uint {
    v5 := GenS1{F0: "hello", F1: 182, F2: 255}
    v6 := []int{p1, (p1 | (^(p1))), (p1 / (((p1 / ((1) | 1))) | 1)), p1}
    for fi7 := 0; fi7 < 17; fi7++ {
        var v8 float64 = 3.14159
        if (4294967295 < 813) {
            for fi9 := 0; fi9 < 15; fi9++ {
                v6 = append(v6, p1)
                p0 = p0
                fmt.Println(fi7)
                var v10 uint64 = p0
                if true {
                    fmt.Println(v10, p0)
                    fmt.Println(v10, fi7, v10)
                    v6 = append(v6, fi9)
                }
                fmt.Println(fi9, p0)
                _ = fi9
                _ = v10
            }
            switch "foo bar" {
            case "\n\t":
                if (uint32(p0) >= 65535) {
                    v6 = append(v6, (fi7 & (p1 | 2147483647)))
                }
            default:
                fmt.Println(p0, p0)
            }
        }
        v11 := GenS1{F0: "a", F1: 0, F2: 32767}
        _ = fi7
        _ = v11
        _ = v8
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 4294967295
}
func genFn1() uint16 {
    var v12 float32 = 1.5
    v13 := []int{127, 490}
    if true {
        v12 = ((v12 - v12) * float32(1.5))
    } else {
        switch 815 {
        case -9223372036854775808:
            _ = len(v13)
        }
        var v14 bool = (-896 >= -32768)
        _ = v14
    }
    if (128 == 42) {
        if false {
            var v15 int32 = -2147483648
            fmt.Println(v12, v12)
            _ = v15
        }
        v12 = v12
    } else {
        v12 += v12
    }
    _ = v12
    _ = v13
    return 32767
}

func main() {
    var v1 string = "hello"
    v2 := make(map[uint32]int)
    var v3 uint16 = 42
    var v4 float32 = (-2.25 * 1.5)
    v1 = ("a" + v1)
    if false {
        fmt.Println(v3)
    } else {
        for fi16 := 0; fi16 < 14; fi16++ {
            var v17 int64 = -9223372036854775808
            var v18 float32 = v4
            v1 = v1
            _ = fi16
            _ = v17
            _ = v18
        }
        delete(v2, 65535)
        v3 = 32767
    }
    for fi19 := 0; fi19 < 2; fi19++ {
        fmt.Println(v1, v3)
        _ = len(v2)
        v20 := GenS1{F0: v1, F1: 0, F2: (v3 ^ 0)}
        _ = fi19
        _ = v20
    }
    if true {
        fmt.Println(v3, v3, v4)
        fmt.Println(v1)
    } else {
        v4 = (v4 - ((float32(float64(-2.25)) * v4) * v4))
    }
    switch v1 {
    case "\n\t":
        v21 := GenS1{F0: "foo bar", F1: 65535, F2: genFn1()}
        _ = v21
    case "a":
        v22 := GenS1{F0: "", F1: 99, F2: (v3 / (((^(v3))) | 1))}
        _ = v22
    case "foo bar":
        var v23 string = ("hello" + v1)
        _ = v23
    }
    var v24 int = -2147483648
    v25 := GenS2{F0: (-238 <= 0), F1: 0, F2: v3}
    switch (v24 * 223) {
    case 127:
        v26 := GenS2{F0: false, F1: -345, F2: (^(v3))}
        _ = v26
    case 4294967295:
        switch "a" {
        case "a":
            fmt.Println(v4)
        default:
            var v27 int16 = 1
            _ = v27
        }
    default:
        v3 += genFn1()
    }
    v4 += 0.1
    v3 = 396
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
}
