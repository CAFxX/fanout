package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int32; F1 int8 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint64 }
func genFn0() uint64 {
    if (128 <= 1) {
        switch -381 {
        case 127:
            if true {
                var v5 float64 = 0.0
                if (1.5 < 1e10) {
                    v5 -= v5
                    var v6 int16 = 32767
                    _ = v6
                } else {
                    for fi7 := 0; fi7 < 5; fi7++ {
                        var v8 float64 = 0.0
                        _ = fi7
                        _ = v8
                    }
                    if (1 <= 32767) {
                        v5 -= ((v5 * v5) * ((((v5 + (v5 - 0.1)) + 3.14159) + (v5 / v5)) + 1.5))
                        v9 := []int{0, -1, -989}
                        v10 := GenS1{F0: 861, F1: 100}
                        _ = v10
                        _ = v9
                    }
                    v11 := GenS2{F0: 1}
                    _ = v11
                }
                v5 += (v5 / 0.1)
                _ = v5
            } else {
                var v12 int = -176
                v12 = 32768
                _ = v12
            }
        case 32768:
            var v13 uint64 = 1
            _ = v13
        case 4294967295:
            var v14 uint64 = 0
            _ = v14
        }
        v15 := GenS2{F0: 65535}
        var v16 int64 = 1
        _ = v15
        _ = v16
    }
    v17 := []int{128, 673, 32768}
    var v18 bool = (!(42 > -607))
    _ = v17
    _ = v18
    return 192
}
func genFn1(p0 int64, p1 uint, p2 int64) uint8 {
    if (((^(p1)) == 2) && (!(true && (0.0 < 0.0)))) {
        v19 := GenS3{F0: 42, F1: 127}
        _ = v19
    }
    if (true && false) {
        for fi20 := 0; fi20 < 11; fi20++ {
            fmt.Println(p1, fi20, p1)
            var v21 float64 = 1.5
            switch "foo bar" {
            case "a":
                v22 := GenS1{F0: 65535, F1: 1}
                _ = v22
            case "hello":
                if false {
                    v21 = (v21 * v21)
                    for fi23 := 0; fi23 < 15; fi23++ {
                        fmt.Println(v21, fi23, p0)
                        _ = fi23
                    }
                    var v24 float32 = 1.5
                    _ = v24
                }
            }
            _ = fi20
            _ = v21
        }
        p0 = p0
        v25 := GenS4{F0: 1}
        _ = v25
    }
    switch "" {
    case "foo bar":
        for fi26 := 0; fi26 < 11; fi26++ {
            var v27 int16 = 42
            v28 := []int{fi26, (fi26 << ((p1) & 63)), (fi26 << (((p1 * (p1 % ((p1) | 1)))) & 63))}
            var v29 int = (fi26 * (fi26 % ((int(p2)) | 1)))
            _ = fi26
            _ = v27
            _ = v28
            _ = v29
        }
    default:
        p0 += p0
    }
    p1 += p1
    v30 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v30
    return 42
}

func main() {
    var v1 string = "a"
    v2 := make(map[string]int)
    var v3 int8 = 22
    var v4 int64 = 0
    v31 := GenS5{F0: 433}
    var v32 string = "hello"
    for fi33 := 0; fi33 < 17; fi33++ {
        v4 -= 9223372036854775807
        switch (-(fi33)) {
        case -385:
            v34 := []int{-761, (^(fi33))}
            _ = v34
        }
        _ = fi33
    }
    for fi35 := 0; fi35 < 7; fi35++ {
        var v36 uint = 65535
        _ = fi35
        _ = v36
    }
    v32 = (v32 + ((v1 + v32) + (v32 + v1)))
    for fi37 := 0; fi37 < 12; fi37++ {
        var v38 uint32 = 0
        _ = fi37
        _ = v38
    }
    v32 = v32
    v39 := GenS1{F0: 65535, F1: (v3 << ((4294967295) & 63))}
    var v40 int16 = -1
    v3 = v3
    var v41 int16 = (v40 & 943)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v39
    _ = v4
    _ = v40
    _ = v41
}
