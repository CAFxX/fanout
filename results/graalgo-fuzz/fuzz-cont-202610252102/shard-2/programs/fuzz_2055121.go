package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32 }
type GenS2 struct { F0 int32; F1 int8 }
type GenS3 struct { F0 int8; F1 bool; F2 uint8 }
type GenS4 struct { F0 int32; F1 int64; F2 string }
type GenS5 struct { F0 uint8 }
func genFn0(p0 string, p1 int64, p2 uint32) uint64 {
    v4 := GenS1{F0: 1, F1: 0}
    if (-9223372036854775808 < 592) {
        v5 := []int{}
        _ = v5
    }
    v6 := []int{32768, 397, 42, 32768}
    var v7 int32 = 0
    p0 = ""
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v6
    _ = v7
    return 18446744073709551615
}
func genFn1(p0 uint16) int64 {
    var v8 int32 = 2147483647
    v9 := []int{1}
    var v10 int32 = (v8 / (((^(v8))) | 1))
    v8 += (v10 / (((v8 + v10)) | 1))
    if len(v9) > 0 {
        v9[0] = 329
    }
    _ = p0
    _ = v10
    _ = v8
    _ = v9
    return int64(p0)
}

func main() {
    v1 := make(map[int]string)
    var v2 int8 = 1
    var v3 uint32 = 2
    var v11 bool = false
    v12 := GenS2{F0: 386, F1: v2}
    v13 := GenS2{F0: 2147483647, F1: (-(v2))}
    for fi14 := 0; fi14 < 20; fi14++ {
        if (0 == 481) {
            v15 := GenS1{F0: (v2 << ((42) & 63)), F1: -1}
            v3 += v3
            _ = v15
        }
        _ = fi14
    }
    for fi16 := 0; fi16 < 9; fi16++ {
        switch "foo bar" {
        case "foo bar":
            for fi17 := 0; fi17 < 5; fi17++ {
                var v18 int16 = 42
                _ = fi17
                _ = v18
            }
        default:
            var v19 uint8 = 100
            _ = v19
        }
        _ = fi16
    }
    var v20 float32 = float32(0.1)
    for fi21 := 0; fi21 < 1; fi21++ {
        var v22 uint32 = (v3 % (((^(v3))) | 1))
        for fi23 := 0; fi23 < 10; fi23++ {
            var v24 int16 = -1
            _ = fi23
            _ = v24
        }
        _ = fi21
        _ = v22
    }
    v25 := []int{-214, 2147483647}
    fmt.Println(v11, v3)
    fmt.Println(v2)
    _ = len(v1)
    delete(v1, -1)
    for fi26 := 0; fi26 < 12; fi26++ {
        if true {
            if v11 {
                switch 128 {
                case 223:
                    var v27 bool = (v20 == -2.25)
                    _ = v27
                case 32768:
                    switch "a" {
                    case "hello":
                        var v28 float32 = float32(float64(float32((float64((-2.25 * (v20 / ((float32(1.5) * 0.0) * v20)))) / (0.1 - 1.5)))))
                        _ = v28
                    }
                case 128:
                    v2 -= v2
                default:
                    if v11 {
                        v29 := GenS3{F0: (v2 & v2), F1: v11, F2: 42}
                        v11 = v11
                        _ = v29
                    }
                }
                var v30 int32 = 0
                v20 += (float32(-2.25) * (v20 * (0.1 * 0.0)))
                _ = v30
            } else {
                var v31 int8 = v2
                v25 = append(v25, 568)
                v32 := GenS3{F0: (v31 & (v31 << ((1000) & 63))), F1: v11, F2: 42}
                _ = v31
                _ = v32
            }
            if ((-2.25 >= (0.0 / (-2.25 - (1.5 / (float64(0.0) - ((1e10 / -2.25) * -2.25)))))) || (32767 == 657)) {
                v33 := GenS4{F0: 0, F1: 1, F2: "foo bar"}
                fmt.Println(v11)
                if len(v25) > 0 {
                    v25[0] = fi26
                }
                _ = v33
            } else {
                if v11 {
                    for fi34 := 0; fi34 < 18; fi34++ {
                        v3 += v3
                        v20 = v20
                        _ = fi34
                    }
                    v35 := GenS5{F0: uint8(v3)}
                    _ = v35
                } else {
                    var v36 int32 = -2147483648
                    v36 -= v36
                    var v37 uint32 = 42
                    _ = v36
                    _ = v37
                }
                var v38 uint8 = 127
                var v39 int64 = 257
                _ = v38
                _ = v39
            }
        }
        v2 = v2
        _ = fi26
    }
    var v40 uint64 = 4294967295
    for fi41 := 0; fi41 < 20; fi41++ {
        for fi42 := 0; fi42 < 12; fi42++ {
            var v43 int64 = genFn1(848)
            v44 := GenS5{F0: 176}
            v20 += v20
            v45 := GenS4{F0: 632, F1: (v43 * genFn1(80)), F2: (("a" + "") + "a")}
            _ = fi42
            _ = v43
            _ = v44
            _ = v45
        }
        fmt.Println(v40, fi41, v20)
        _ = fi41
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v3:", v3)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v20
    _ = v25
    _ = v3
    _ = v40
}
