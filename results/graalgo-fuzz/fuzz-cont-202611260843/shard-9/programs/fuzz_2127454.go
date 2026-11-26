package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint; F1 uint16 }
type GenS4 struct { F0 int16; F1 uint; F2 string }
type GenS5 struct { F0 uint32; F1 uint8 }
type GenS6 struct { F0 string }
func genFn0(p0 int, p1 uint8, p2 uint) int32 {
    switch "hello" {
    case "\n\t":
        v4 := GenS1{F0: 1, F1: (^(p2))}
        _ = v4
    case "hello":
        switch p0 {
        case -2147483648:
            if (!(!(!(!(0 < -1))))) {
                switch (p0 / ((p0) | 1)) {
                case -2147483648:
                    p2 = (p2 >> (((p2 + p2)) & 63))
                case 42:
                    v5 := []int{(-(p0)), (p0 / ((p0) | 1))}
                    _ = v5
                case 622:
                    v6 := GenS2{F0: 42}
                    _ = v6
                }
                var v7 float64 = float64((0.0 - ((-2.25 - float32(0.1)) / (3.14159 / (-2.25 / 1.0)))))
                _ = v7
            }
        case 0:
            var v8 int64 = -382
            _ = v8
        }
    default:
        var v9 float64 = 3.14159
        _ = v9
    }
    p2 -= (p2 >> ((0) & 63))
    var v10 bool = (true && true)
    if v10 {
        v11 := GenS1{F0: 439, F1: p2}
        v12 := GenS3{F0: p2, F1: 0}
        _ = v11
        _ = v12
    } else {
        if v10 {
            switch "a" {
            case "a":
                v13 := []int{(p0 - 42), (p0 % (((p0 << ((p2) & 63))) | 1))}
                _ = v13
            case "foo bar":
                v14 := []int{32768}
                _ = v14
            default:
                v15 := []int{(p0 - p0), p0, p0}
                _ = v15
            }
            var v16 int32 = 1
            v10 = (((v16 >= (v16 % ((42) | 1))) && (!(v10 || v10))) && false)
            _ = v16
        }
        var v17 int64 = -1
        _ = v17
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return -206
}
func genFn1(p0 uint16, p1 uint, p2 uint64) string {
    for fi18 := 0; fi18 < 10; fi18++ {
        p2 = p2
        if (42 == p2) {
            if false {
                var v19 uint64 = p2
                switch 32 {
                case 96:
                    var v20 int8 = -128
                    _ = v20
                default:
                    for fi21 := 0; fi21 < 15; fi21++ {
                        p1 -= (p1 << ((2) & 63))
                        v22 := []int{fi18, fi18, int(p0), 127}
                        var v23 float32 = 3.14159
                        _ = fi21
                        _ = v22
                        _ = v23
                    }
                }
                for fi24 := 0; fi24 < 8; fi24++ {
                    if (((p0 / ((0) | 1)) != 65535) || (false || true)) {
                        if (4294967295 != 138) {
                            var v25 int64 = 0
                            _ = v25
                        } else {
                            fmt.Println(fi18, fi24)
                        }
                    } else {
                        var v26 string = "\n\t"
                        v27 := GenS4{F0: -1, F1: p1, F2: (v26 + (v26 + ("a" + "hello")))}
                        var v28 float64 = (float64(1.5) * 1e10)
                        _ = v26
                        _ = v27
                        _ = v28
                    }
                    switch int(p1) {
                    case 0:
                        v29 := GenS1{F0: uint32(p0), F1: (p1 / ((p1) | 1))}
                        _ = v29
                    }
                    _ = fi24
                }
                _ = v19
            } else {
                var v30 uint16 = 552
                p0 = 32767
                p2 += (^(p2))
                _ = v30
            }
            var v31 int64 = -1
            _ = v31
        }
        _ = fi18
    }
    p1 = (p1 << (((p1 ^ p1)) & 63))
    _ = p0
    _ = p1
    _ = p2
    return ""
}

func main() {
    var v1 bool = (570 > 1)
    v2 := []uint32{65535, 2147483648, 65535}
    v3 := make(map[int]bool)
    var v32 int = 4294967295
    var v33 int = -2147483648
    v34 := GenS5{F0: 0, F1: 127}
    v2 = append(v2, 42)
    if (!("foo bar" >= "\n\t")) {
        v33 -= v33
    } else {
        v33 = v32
        v1 = v1
        v33 = -428
    }
    v35 := GenS2{F0: genFn0(-2147483648, 92, 65535)}
    v33 -= v33
    delete(v3, (v32 - v33))
    if (true && (194 == 137)) {
        v36 := GenS2{F0: 42}
        v37 := GenS6{F0: ""}
        var v38 string = "\n\t"
        _ = v36
        _ = v37
        _ = v38
    }
    fmt.Println(v32)
    var v39 int64 = 0
    var v40 int16 = 229
    var v41 int64 = (-(v39))
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    _ = len(v2)
    var v42 int64 = (^(v39))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v39:", v39)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v39
    _ = v40
    _ = v41
    _ = v42
}
