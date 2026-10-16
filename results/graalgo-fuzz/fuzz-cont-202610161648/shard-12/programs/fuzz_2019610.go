package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int; F2 int16 }
type GenS2 struct { F0 int; F1 int8 }
type GenS3 struct { F0 int16; F1 uint32; F2 uint8 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint32; F1 uint16 }
func genFn0(p0 uint64, p1 uint16) int {
    switch 474 {
    case 42:
        fmt.Println(p0, p0)
    case 32767:
        v7 := []int{-732, 0}
        _ = v7
    case 128:
        if false {
            for fi8 := 0; fi8 < 14; fi8++ {
                var v9 uint16 = 221
                v10 := GenS1{F0: 4294967295, F1: 870, F2: -177}
                _ = fi8
                _ = v10
                _ = v9
            }
        } else {
            v11 := []int{-736, 439}
            fmt.Println(p0)
            p1 = p1
            _ = v11
        }
    default:
        var v12 bool = false
        _ = v12
    }
    p0 += (^(p0))
    var v13 uint = 2
    if true {
        v14 := GenS2{F0: 699, F1: 127}
        var v15 int8 = 100
        var v16 uint8 = 128
        _ = v14
        _ = v15
        _ = v16
    }
    if (false || true) {
        var v17 int16 = 1
        fmt.Println(v13)
        fmt.Println(p0, p0)
        _ = v17
    } else {
        var v18 int16 = -32768
        if true {
            fmt.Println(p1, v13)
            p1 += p1
        } else {
            v19 := GenS3{F0: v18, F1: 829, F2: 111}
            _ = v19
        }
        p0 = (p0 & p0)
        _ = v18
    }
    _ = p0
    _ = p1
    _ = v13
    return -196
}
func genFn1() int8 {
    if (!false) {
        for fi20 := 0; fi20 < 17; fi20++ {
            v21 := []int{genFn0(4294967295, 1), fi20, (^(fi20))}
            var v22 int64 = -9223372036854775808
            _ = fi20
            _ = v21
            _ = v22
        }
    }
    var v23 float64 = 0.1
    v24 := GenS1{F0: 388, F1: genFn0(18446744073709551615, 2), F2: 32767}
    if true {
        v23 = (v23 - (v23 / v23))
        v23 -= 1.5
    }
    var v25 uint16 = 255
    _ = v23
    _ = v24
    _ = v25
    return 100
}

func main() {
    v1 := []string{("hello" + "\n\t")}
    var v2 uint8 = 128
    var v3 uint = 4294967295
    var v4 string = "\n\t"
    var v5 string = ((v4 + v4) + v4)
    var v6 uint32 = 228
    v26 := GenS4{F0: 100}
    if false {
        var v27 float32 = (float32((-2.25 + -2.25)) / ((float32(0.1) - (0.1 - 1.5)) + float32(0.1)))
        if false {
            v28 := GenS4{F0: 100}
            _ = v28
        } else {
            switch v5 {
            case "a":
                if false {
                    v6 = (v6 * (v6 ^ v6))
                    fmt.Println(v3, v2, v3)
                } else {
                    fmt.Println(v6, v27)
                }
            default:
                for fi29 := 0; fi29 < 10; fi29++ {
                    v1 = append(v1, (v5 + ("" + v4)))
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    v5 = ("foo bar" + "")
                    _ = fi29
                }
            }
            v1 = append(v1, "hello")
        }
        _ = v27
    } else {
        for fi30 := 0; fi30 < 13; fi30++ {
            var v31 string = v5
            _ = fi30
            _ = v31
        }
        v32 := GenS2{F0: 32767, F1: -128}
        var v33 int64 = 1
        _ = v32
        _ = v33
    }
    var v34 uint32 = 4294967295
    v2 = v2
    for fi35 := 0; fi35 < 14; fi35++ {
        var v36 uint = v3
        _ = fi35
        _ = v36
    }
    v37 := GenS5{F0: (v6 + v6), F1: 1}
    var v38 uint32 = 2147483648
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println(v5, v3, v38)
    v2 += 178
    v39 := GenS5{F0: (^(v38)), F1: 255}
    var v40 int8 = 42
    v6 -= 956
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v34
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
