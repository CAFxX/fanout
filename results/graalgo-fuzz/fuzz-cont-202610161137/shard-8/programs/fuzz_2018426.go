package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16 }
type GenS2 struct { F0 bool; F1 int16; F2 int64 }
type GenS3 struct { F0 uint }
func genFn0(p0 int16, p1 int16) bool {
    for fi5 := 0; fi5 < 19; fi5++ {
        switch 4294967295 {
        case 2147483647:
            for fi6 := 0; fi6 < 19; fi6++ {
                var v7 int16 = p0
                _ = fi6
                _ = v7
            }
        case -49:
            var v8 uint8 = 128
            _ = v8
        }
        _ = fi5
    }
    v9 := []int{-2147483648, -367}
    fmt.Println(p1)
    switch 0 {
    case 9223372036854775807:
        var v10 string = "foo bar"
        _ = v10
    default:
        for fi11 := 0; fi11 < 20; fi11++ {
            if true {
                v12 := GenS1{F0: 1, F1: 0}
                v13 := GenS2{F0: true, F1: (p1 >> ((2) & 63)), F2: 42}
                v9 = append(v9, fi11)
                _ = v12
                _ = v13
            } else {
                fmt.Println(p0)
                var v14 uint8 = 2
                _ = v14
            }
            if false {
                var v15 bool = true
                v16 := GenS1{F0: 103, F1: 1}
                fmt.Println(v15, p0)
                _ = v15
                _ = v16
            } else {
                fmt.Println(p1, p1)
                p0 = (p1 >> ((65535) & 63))
                v17 := GenS1{F0: 1, F1: 255}
                _ = v17
            }
            _ = fi11
        }
    }
    _ = p0
    _ = p1
    _ = v9
    return true
}
func genFn1(p0 uint64, p1 uint) uint64 {
    if false {
        var v18 int16 = -32768
        _ = v18
    } else {
        for fi19 := 0; fi19 < 4; fi19++ {
            v20 := GenS3{F0: uint(p0)}
            for fi21 := 0; fi21 < 14; fi21++ {
                var v22 string = "\n\t"
                v23 := []int{fi21, -780, (^(fi19))}
                v24 := GenS1{F0: 838, F1: 182}
                _ = fi21
                _ = v22
                _ = v23
                _ = v24
            }
            var v25 uint32 = 751
            _ = fi19
            _ = v20
            _ = v25
        }
    }
    if true {
        p0 += 18446744073709551615
    } else {
        p0 = 42
        p0 = p0
        fmt.Println(p0, p0)
    }
    _ = p0
    _ = p1
    return 463
}

func main() {
    v1 := []int{32767, -1}
    v2 := make(map[string]string)
    var v3 uint16 = 1
    v4 := []uint32{107}
    v3 -= v3
    v3 -= v3
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println(v3)
    var v26 bool = (!((2 > 2) || (128 >= -769)))
    v27 := v2[(("\n\t" + "a") + (("\n\t" + "foo bar") + "\n\t"))]
    var v28 string = ("hello" + "a")
    switch 32767 {
    case 127:
        v29 := GenS1{F0: 674, F1: 2}
        _ = v29
    case -9223372036854775808:
        var v30 uint16 = 42
        _ = v30
    default:
        v27 = v28
    }
    v31 := GenS3{F0: 255}
    for fi32 := 0; fi32 < 12; fi32++ {
        v2["a"] = v28
        var v33 int = -992
        if (!(false && v26)) {
            v34 := GenS1{F0: (^(v3)), F1: (v3 * v3)}
            _ = v34
        }
        _ = fi32
        _ = v33
    }
    for fi35 := 0; fi35 < 17; fi35++ {
        for fi36 := 0; fi36 < 13; fi36++ {
            fmt.Println(v3, v27)
            _ = fi36
        }
        fmt.Println(v26)
        var v37 uint8 = 255
        _ = fi35
        _ = v37
    }
    v26 = ((((-2.25 - -2.25) * float64(float32((1.5 / 1e10)))) - -2.25) >= 0.1)
    var v38 uint = 18446744073709551615
    var v39 uint8 = 12
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v31
    _ = v38
    _ = v39
    _ = v4
}
