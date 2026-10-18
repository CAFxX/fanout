package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32 }
type GenS2 struct { F0 int8; F1 bool }
type GenS3 struct { F0 uint8; F1 int8 }
type GenS4 struct { F0 int32 }
func genFn0() int16 {
    var v7 int64 = -1
    for fi8 := 0; fi8 < 6; fi8++ {
        var v9 int64 = v7
        v9 = v7
        v10 := []int{32767, fi8}
        _ = fi8
        _ = v10
        _ = v9
    }
    _ = v7
    return -1
}

func main() {
    var v1 float32 = 1e10
    var v2 bool = true
    var v3 float64 = (1e10 + 0.1)
    v4 := []string{}
    var v5 int32 = 776
    v6 := []uint32{1, 2, 963}
    for fi11 := 0; fi11 < 4; fi11++ {
        switch fi11 {
        case -1:
            var v12 int = -9223372036854775808
            _ = v12
        case -33:
            var v13 int32 = (^(v5))
            _ = v13
        }
        _ = fi11
    }
    v3 = (((((1.5 * float64(v1)) * v3) - v3) - (v3 - (1.5 / (0.1 / 1.5)))) * (v3 / v3))
    if ((1 >= 18446744073709551615) || (v2 || (!v2))) {
        v14 := GenS1{F0: -1, F1: 107}
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        fmt.Println(v3)
        _ = v14
    } else {
        if (v2 || true) {
            fmt.Println(v3, v1)
            v5 = v5
            v3 = (1.5 / v3)
        }
        var v15 int16 = genFn0()
        fmt.Println(v2, v3, v15)
        _ = v15
    }
    v1 -= (v1 / ((v1 * 1e10) - float32(v3)))
    v16 := GenS2{F0: -128, F1: v2}
    for fi17 := 0; fi17 < 6; fi17++ {
        v5 -= 1
        v18 := GenS3{F0: 1, F1: 93}
        switch "hello" {
        case "a":
            var v19 uint64 = 18446744073709551615
            _ = v19
        case "hello":
            _ = len(v4)
        case "foo bar":
            for fi20 := 0; fi20 < 6; fi20++ {
                _ = len(v6)
                v2 = v2
                if (1 == 32767) {
                    fmt.Println(fi17, fi20)
                    v4 = append(v4, ("hello" + ""))
                    v21 := GenS3{F0: 255, F1: -1}
                    v22 := GenS4{F0: -352}
                    _ = v21
                    _ = v22
                }
                _ = fi20
            }
        }
        _ = fi17
        _ = v18
    }
    v3 += v3
    if (-907 != -903) {
        v6 = append(v6, 2147483647)
        for fi23 := 0; fi23 < 14; fi23++ {
            var v24 bool = false
            v3 -= v3
            var v25 int = fi23
            _ = fi23
            _ = v24
            _ = v25
        }
        for fi26 := 0; fi26 < 18; fi26++ {
            if v2 {
                var v27 uint32 = 501
                var v28 int64 = 0
                if len(v6) > 0 {
                    v6[0] = (v27 * 65535)
                }
                _ = v27
                _ = v28
            } else {
                for fi29 := 0; fi29 < 7; fi29++ {
                    var v30 int64 = 9223372036854775807
                    _ = fi29
                    _ = v30
                }
                v31 := GenS1{F0: 1000, F1: 65535}
                fmt.Println(v1, fi26, v2)
                _ = v31
            }
            var v32 uint64 = 2
            var v33 int64 = -1
            _ = fi26
            _ = v32
            _ = v33
        }
    }
    var v34 int64 = -9223372036854775808
    v5 = 42
    for fi35 := 0; fi35 < 7; fi35++ {
        fmt.Println(v2, v34)
        v34 -= v34
        _ = fi35
    }
    var v36 uint32 = 4294967295
    for fi37 := 0; fi37 < 12; fi37++ {
        for fi38 := 0; fi38 < 18; fi38++ {
            v1 = (((0.0 - 0.0) * v1) - (v1 / 1.5))
            _ = len(v4)
            _ = fi38
        }
        _ = fi37
    }
    v1 += v1
    var v39 uint8 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v36:", v36)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v34
    _ = v36
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
