package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8; F2 bool }
type GenS2 struct { F0 uint32; F1 uint64 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint64; F1 bool; F2 bool }
type GenS5 struct { F0 uint16 }
func genFn0(p0 uint32, p1 int32) uint16 {
    fmt.Println(p1, p1)
    var v6 uint32 = 511
    var v7 int32 = (p1 << ((2) & 63))
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 1
}
func genFn1(p0 int16, p1 uint, p2 int16) uint32 {
    v8 := GenS1{F0: p0, F1: 29, F2: ((true && true) || (false && true))}
    p2 = 450
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 2
}

func main() {
    var v1 uint64 = 4294967295
    var v2 uint16 = 705
    v3 := []int{-285, -281, 42}
    v4 := []uint32{}
    v5 := make(map[string]int)
    for fi9 := 0; fi9 < 20; fi9++ {
        var v10 float64 = 0.1
        var v11 uint8 = 42
        _ = fi9
        _ = v10
        _ = v11
    }
    var v12 uint8 = 255
    v13 := GenS1{F0: 0, F1: 1, F2: (!true)}
    fmt.Println(v12, v2, v1)
    switch -908 {
    case -897:
        v12 = (v12 >> ((494) & 63))
    case 331:
        var v14 int8 = 0
        _ = v14
    case -540:
        for fi15 := 0; fi15 < 2; fi15++ {
            v16 := GenS2{F0: 2, F1: v1}
            var v17 string = ("a" + ("hello" + ("\n\t" + "a")))
            v2 += (v2 % ((genFn0(0, 0)) | 1))
            _ = fi15
            _ = v16
            _ = v17
        }
    default:
        fmt.Println(v1, v12)
    }
    var v18 uint32 = genFn1(1000, 695, 1000)
    for fi19 := 0; fi19 < 8; fi19++ {
        if true {
            v18 += genFn1(-1, 18446744073709551615, 1000)
        } else {
            var v20 int8 = 1
            for fi21 := 0; fi21 < 1; fi21++ {
                fmt.Println(fi19, v12, v12)
                _ = len(v5)
                fmt.Println(v18, v1)
                _ = fi21
            }
            _ = v20
        }
        var v22 bool = true
        _ = fi19
        _ = v22
    }
    for fi23 := 0; fi23 < 11; fi23++ {
        v1 = (v1 | uint64(v12))
        switch ((("\n\t" + "") + "a") + ("" + "\n\t")) {
        case "\n\t":
            v18 = (v18 & (v18 % ((genFn1(-631, 256, 1)) | 1)))
        case "hello":
            for fi24 := 0; fi24 < 13; fi24++ {
                v25 := GenS2{F0: (^(v18)), F1: v1}
                _ = len(v4)
                if ((float32((((((-2.25 - (float64((0.1 * 1.5)) - 3.14159)) / 3.14159) * 3.14159) * 1.5) / 1.0)) + 3.14159) != float32(0.0)) {
                    v1 -= v1
                    if len(v4) > 0 {
                        v4[0] = (v18 % ((genFn1(-32768, 1, -1)) | 1))
                    }
                    var v26 float32 = float32((float64((float32(-2.25) * 0.1)) * float64((3.14159 + 1e10))))
                    _ = v26
                } else {
                    if false {
                        v1 -= (^(v1))
                    }
                    var v27 int = (fi23 - (fi23 << ((1) & 63)))
                    v27 = fi23
                    _ = v27
                }
                v28 := GenS3{F0: 641}
                if true {
                    v29 := GenS4{F0: v1, F1: true, F2: true}
                    var v30 int64 = 146
                    _ = v29
                    _ = v30
                }
                v2 -= 2
                var v31 int64 = -127
                v18 = (v18 + 4294967295)
                fmt.Println(v1, v12)
                var v32 uint64 = (v1 >> ((939) & 63))
                v31 -= v31
                _ = fi24
                _ = v25
                _ = v28
                _ = v31
                _ = v32
            }
        default:
            if ((true || false) && (!(65535 == uint(v12)))) {
                v1 = v1
                fmt.Println(v18)
                for fi33 := 0; fi33 < 20; fi33++ {
                    fmt.Println(v2, fi23, fi23)
                    var v34 uint32 = (^(v18))
                    var v35 uint64 = v1
                    _ = fi33
                    _ = v34
                    _ = v35
                }
            } else {
                var v36 float64 = (0.0 / (1.5 + ((0.0 / ((0.0 - ((3.14159 + ((((float64(0.0) * float64((0.0 + 0.1))) + 1e10) + 3.14159) - 0.1)) + (3.14159 - 1.5))) / (float64(1.5) * 1.5))) + 1e10)))
                var v37 int16 = 32767
                _ = v36
                _ = v37
            }
        }
        v18 -= v18
        _ = fi23
    }
    for fi38 := 0; fi38 < 11; fi38++ {
        fmt.Println(v1, v1)
        v39 := GenS5{F0: (^(v2))}
        _ = fi38
        _ = v39
    }
    var v40 uint8 = 42
    if (false || false) {
        v41 := GenS2{F0: (v18 >> ((503) & 63)), F1: (v1 | v1)}
        _ = v41
    } else {
        var v42 int = 1
        fmt.Println(v12, v42, v12)
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = v42
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v12
    _ = v13
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v5
}
