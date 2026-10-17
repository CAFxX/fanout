package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint8; F1 uint32 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint8; F1 int }
type GenS5 struct { F0 int8 }
func genFn0(p0 int32) string {
    var v5 uint = 18446744073709551615
    if true {
        var v6 uint64 = 0
        var v7 float32 = (0.1 * 1e10)
        _ = v6
        _ = v7
    } else {
        switch "hello" {
        case "hello":
            if ((-82 < p0) || (1 > 4294967295)) {
                p0 += (-(p0))
                if (32767 == -1) {
                    if ((0.0 != float32((1.5 - -2.25))) && (p0 == (p0 + (-(p0))))) {
                        var v8 int = -2147483648
                        p0 = (p0 % (((p0 ^ (^(p0)))) | 1))
                        for fi9 := 0; fi9 < 9; fi9++ {
                            for fi10 := 0; fi10 < 12; fi10++ {
                                var v11 int32 = p0
                                var v12 int16 = -1
                                _ = fi10
                                _ = v11
                                _ = v12
                            }
                            fmt.Println(fi9, v5, p0)
                            v13 := GenS1{F0: (fi9 | (^(v8)))}
                            _ = fi9
                            _ = v13
                        }
                        _ = v8
                    }
                    p0 = (p0 / ((-522) | 1))
                }
                var v14 uint8 = 128
                _ = v14
            }
        case "":
            p0 += p0
        case "\n\t":
            for fi15 := 0; fi15 < 6; fi15++ {
                v16 := []int{-2147483648, fi15, 1, (fi15 + fi15)}
                v17 := GenS1{F0: (fi15 / (((fi15 / ((32767) | 1))) | 1))}
                fmt.Println(fi15, p0, v5)
                _ = fi15
                _ = v16
                _ = v17
            }
        }
        var v18 float32 = (((((float32(0.0) + -2.25) - 1.5) / (-2.25 / (0.1 * -2.25))) - ((float32(-2.25) - 1e10) + ((1.5 * (0.0 - 1e10)) / 1.5))) - ((1.5 / -2.25) - (1e10 - float32(-2.25))))
        if false {
            v19 := GenS2{F0: 2, F1: 42}
            v5 -= (v5 - (v5 & v5))
            _ = v19
        } else {
            for fi20 := 0; fi20 < 19; fi20++ {
                switch ((("foo bar" + "\n\t") + "foo bar") + "") {
                case "a":
                    if true {
                        v21 := []int{127, -887, (fi20 << (((v5 ^ (v5 + (^(v5))))) & 63)), (-(fi20))}
                        fmt.Println(fi20, fi20, v5)
                        _ = len(v21)
                        _ = v21
                    }
                }
                v5 = v5
                _ = fi20
            }
            var v22 uint = 42
            _ = v22
        }
        _ = v18
    }
    for fi23 := 0; fi23 < 14; fi23++ {
        for fi24 := 0; fi24 < 19; fi24++ {
            v25 := GenS3{F0: 0}
            v26 := []int{}
            if false {
                fmt.Println(fi24, p0)
                var v27 int16 = -32768
                _ = v27
            } else {
                v28 := GenS2{F0: 128, F1: 2}
                var v29 uint = (v5 ^ 65535)
                fmt.Println(fi23, p0, v29)
                _ = v28
                _ = v29
            }
            _ = fi24
            _ = v25
            _ = v26
        }
        for fi30 := 0; fi30 < 4; fi30++ {
            v31 := []int{}
            p0 = p0
            _ = len(v31)
            p0 -= (p0 * p0)
            var v32 int8 = -128
            _ = fi30
            _ = v31
            _ = v32
        }
        _ = fi23
    }
    fmt.Println(v5, p0)
    _ = p0
    _ = v5
    return "foo bar"
}

func main() {
    var v1 float64 = (1.5 * -2.25)
    v2 := []uint32{626, 2147483647}
    var v3 float32 = float32(v1)
    var v4 uint8 = 0
    for fi33 := 0; fi33 < 7; fi33++ {
        v3 -= 0.0
        _ = fi33
    }
    if (18446744073709551615 > 1) {
        v34 := GenS3{F0: -50}
        if ((v4 ^ (v4 * 1)) <= 0) {
            var v35 int8 = 40
            _ = v35
        }
        _ = v34
    }
    var v36 uint16 = 0
    v37 := GenS1{F0: 127}
    var v38 uint64 = 4294967295
    v39 := GenS4{F0: (v4 - v4), F1: -1}
    v40 := GenS5{F0: 1}
    if len(v2) > 0 {
        v2[0] = 210
    }
    v41 := GenS2{F0: v4, F1: 2147483647}
    v36 = (v36 & 32767)
    v1 = 0.1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
}
