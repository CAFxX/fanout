package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int8; F1 uint16; F2 string }
type GenS3 struct { F0 string; F1 bool; F2 bool }
func genFn0(p0 uint16, p1 uint, p2 int32) uint {
    v5 := GenS1{F0: p2}
    p2 = p2
    var v6 uint = (^(p1))
    var v7 float32 = 1e10
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return p1
}
func genFn1(p0 int8, p1 int16, p2 uint16) int {
    var v8 int = -889
    for fi9 := 0; fi9 < 9; fi9++ {
        v10 := []int{v8, v8, (fi9 >> ((0) & 63))}
        var v11 uint = uint(p2)
        _ = fi9
        _ = v10
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return (-(v8))
}

func main() {
    v1 := make(map[int]bool)
    var v2 float32 = (-2.25 + float32(3.14159))
    var v3 uint8 = 1
    var v4 uint16 = 32767
    if false {
        v2 = float32((0.1 / 1.5))
    } else {
        if ((float64((3.14159 / (float32(-2.25) * v2))) / (1.5 / (1.5 * -2.25))) <= (1.5 - 3.14159)) {
            for fi12 := 0; fi12 < 14; fi12++ {
                if false {
                    _ = len(v1)
                    switch "a" {
                    case "a":
                        v13 := []int{9223372036854775807, fi12}
                        _ = v13
                    default:
                        if true {
                            if (1 != 4294967295) {
                                v14 := v1[(fi12 ^ 128)]
                                v2 = -2.25
                                _ = v14
                            } else {
                                v3 += (v3 | 2)
                            }
                            var v15 int = (fi12 - (fi12 ^ fi12))
                            _ = v15
                        }
                    }
                } else {
                    fmt.Println(v3)
                    if true {
                        v2 += (v2 + v2)
                        if true {
                            delete(v1, fi12)
                            var v16 uint8 = 0
                            _ = v16
                        } else {
                            var v17 uint64 = 1
                            fmt.Println(v17, v3)
                            v17 -= (^(v17))
                            _ = v17
                        }
                        var v18 int32 = 42
                        _ = v18
                    }
                }
                _ = fi12
            }
        }
        var v19 int32 = -1
        for fi20 := 0; fi20 < 11; fi20++ {
            v21 := GenS1{F0: v19}
            v2 += v2
            _ = fi20
            _ = v21
        }
        _ = v19
    }
    fmt.Println(v3)
    var v22 int64 = -1
    for fi23 := 0; fi23 < 11; fi23++ {
        for fi24 := 0; fi24 < 5; fi24++ {
            var v25 float32 = (v2 / 1.0)
            v26 := GenS1{F0: 239}
            _ = fi24
            _ = v25
            _ = v26
        }
        fmt.Println(v4, v4)
        for fi27 := 0; fi27 < 1; fi27++ {
            v28 := GenS2{F0: 1, F1: 65535, F2: ""}
            _ = fi27
            _ = v28
        }
        _ = fi23
    }
    switch int(v22) {
    case -902:
        var v29 int16 = -32768
        _ = v29
    }
    if ((1.5 / (0.1 + float64(0.1))) == -2.25) {
        var v30 bool = false
        _ = v30
    }
    var v31 uint8 = (^(v3))
    fmt.Println(v3)
    v2 = float32((0.1 + ((0.1 + 1.5) + float64(0.1))))
    fmt.Println(v2, v22)
    v32 := GenS1{F0: 42}
    fmt.Println(v22, v2, v2)
    for fi33 := 0; fi33 < 3; fi33++ {
        if (true || (!true)) {
            v34 := []int{(fi33 | genFn1(127, 499, 383)), (fi33 & 4294967295), (fi33 >> ((42) & 63))}
            _ = v34
        } else {
            var v35 uint64 = 976
            var v36 int32 = 128
            v37 := GenS1{F0: (^(v36))}
            _ = v35
            _ = v36
            _ = v37
        }
        v38 := []int{fi33, (fi33 * fi33), (fi33 % (((fi33 * fi33)) | 1))}
        _ = fi33
        _ = v38
    }
    v39 := GenS2{F0: -1, F1: 65535, F2: ("a" + "\n\t")}
    var v40 int8 = 1
    switch "\n\t" {
    case "\n\t":
        for fi41 := 0; fi41 < 8; fi41++ {
            var v42 string = "\n\t"
            v43 := GenS3{F0: "", F1: (!false), F2: false}
            _ = fi41
            _ = v42
            _ = v43
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v31
    _ = v32
    _ = v39
    _ = v4
    _ = v40
}
