package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32; F2 uint64 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint64; F1 uint8 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint8; F1 int32 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 string }
func genFn0(p0 uint8, p1 uint) int32 {
    var v6 uint64 = 1
    var v7 uint32 = 42
    switch 32767 {
    case -94:
        v6 = (v6 % (((v6 + v6)) | 1))
    case 308:
        v8 := []int{32767, 4294967295, 9223372036854775807, 362}
        _ = v8
    default:
        v9 := GenS1{F0: p0, F1: v7, F2: (v6 + 1)}
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return -1
}
func genFn1() uint64 {
    var v10 int64 = 0
    v11 := GenS2{F0: 42}
    _ = v10
    _ = v11
    return 0
}

func main() {
    var v1 int8 = 42
    v2 := make(map[uint32]int)
    var v3 uint = 18446744073709551615
    v4 := make(map[uint32]bool)
    var v5 uint8 = 97
    fmt.Println(v3)
    if false {
        var v12 uint8 = v5
        var v13 bool = ((true || false) || false)
        _ = v12
        _ = v13
    }
    if ((((1.5 - (0.0 * 1e10)) * (1e10 + 0.1)) + -2.25) <= (0.0 - (1e10 - (float64(0.1) - 3.14159)))) {
        var v14 int64 = -9223372036854775808
        fmt.Println(v14)
        _ = v14
    } else {
        if true {
            if (!true) {
                v3 += v3
                _ = len(v2)
                v1 -= v1
            }
        }
    }
    var v15 int32 = 2147483647
    v3 = 65535
    v16 := v4[136]
    fmt.Println(v3, v1, v1)
    v17 := GenS3{F0: 983, F1: (^(v5))}
    for fi18 := 0; fi18 < 9; fi18++ {
        var v19 int16 = 0
        var v20 uint = 786
        if (v16 && true) {
            v21 := GenS4{F0: 32768}
            if v16 {
                v4[2] = false
                var v22 uint32 = 2147483647
                var v23 uint = v3
                _ = v22
                _ = v23
            } else {
                var v24 uint32 = 2147483648
                for fi25 := 0; fi25 < 2; fi25++ {
                    var v26 int64 = -1
                    v3 -= v20
                    _ = fi25
                    _ = v26
                }
                _ = v24
            }
            _ = v21
        } else {
            v5 = (v5 & (v5 * v5))
            for fi27 := 0; fi27 < 5; fi27++ {
                v3 += v20
                v28 := []int{}
                v29 := GenS5{F0: 255, F1: v15}
                _ = fi27
                _ = v28
                _ = v29
            }
        }
        _ = fi18
        _ = v19
        _ = v20
    }
    v3 -= (^(v3))
    if (v5 <= (v5 + (v5 * v5))) {
        for fi30 := 0; fi30 < 20; fi30++ {
            var v31 string = ("\n\t" + ("a" + ""))
            v32 := GenS6{F0: 2147483647}
            var v33 string = v31
            _ = fi30
            _ = v31
            _ = v32
            _ = v33
        }
        v34 := GenS7{F0: "foo bar"}
        v35 := []int{}
        _ = v34
        _ = v35
    }
    delete(v4, 42)
    v36 := []int{-502, int(v1), 94}
    v16 = true
    var v37 float64 = 1e10
    for fi38 := 0; fi38 < 10; fi38++ {
        var v39 float64 = (v37 / (float64((float32((1e10 - float64(1.5))) - 0.1)) - v37))
        var v40 float64 = float64(((float32((v39 / v39)) / ((-2.25 + 1.5) + float32(v37))) + 0.1))
        _ = fi38
        _ = v39
        _ = v40
    }
    if (((v1 >> ((v3) & 63)) == v1) && v16) {
        var v41 float32 = ((1.5 + (0.1 / 1e10)) * -2.25)
        if (((v1 + v1) <= -74) || ((v15 - v15) >= (^(v15)))) {
            fmt.Println(v5, v37, v37)
            if (32767 < 148) {
                if (!("foo bar" >= "foo bar")) {
                    if len(v36) > 0 {
                        v36[0] = 32767
                    }
                }
            }
            if (int8(v3) < v1) {
                for fi42 := 0; fi42 < 18; fi42++ {
                    fmt.Println(v37, v15)
                    _ = fi42
                }
                var v43 float32 = v41
                _ = v43
            }
        }
        var v44 float32 = 1e10
        _ = v41
        _ = v44
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v36:", len(v36), v36)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v5
}
