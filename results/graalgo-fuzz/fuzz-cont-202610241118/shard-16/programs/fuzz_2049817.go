package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int }
type GenS2 struct { F0 int32; F1 bool }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint32; F1 uint64 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 uint32; F1 bool }
func genFn0(p0 bool, p1 uint8) uint16 {
    p0 = (!true)
    for fi7 := 0; fi7 < 13; fi7++ {
        var v8 int32 = -12
        if p0 {
            for fi9 := 0; fi9 < 5; fi9++ {
                v10 := []int{-925, fi7, fi9}
                _ = fi9
                _ = v10
            }
            v8 -= v8
        }
        _ = fi7
        _ = v8
    }
    if (-118 == 16) {
        for fi11 := 0; fi11 < 12; fi11++ {
            var v12 int32 = -512
            for fi13 := 0; fi13 < 5; fi13++ {
                fmt.Println(p1, fi13)
                fmt.Println(fi13, fi13)
                _ = fi13
            }
            _ = fi11
            _ = v12
        }
        for fi14 := 0; fi14 < 5; fi14++ {
            var v15 uint32 = 2
            for fi16 := 0; fi16 < 11; fi16++ {
                v17 := GenS1{F0: v15, F1: (fi16 % ((941) | 1))}
                _ = fi16
                _ = v17
            }
            if p0 {
                v18 := []int{int(p1), (fi14 * fi14)}
                var v19 uint16 = 888
                _ = v18
                _ = v19
            } else {
                if p0 {
                    v20 := GenS1{F0: (v15 & v15), F1: fi14}
                    var v21 string = "a"
                    var v22 float64 = (0.0 * 0.1)
                    _ = v20
                    _ = v21
                    _ = v22
                } else {
                    v23 := GenS2{F0: 42, F1: ((!false) && p0)}
                    if p0 {
                        for fi24 := 0; fi24 < 8; fi24++ {
                            var v25 float32 = (1.5 / 1.0)
                            if (2147483647 <= (fi24 & 0)) {
                                fmt.Println(p1, v25)
                                fmt.Println(fi14)
                                p1 += 228
                            }
                            _ = fi24
                            _ = v25
                        }
                        var v26 int64 = 923
                        _ = v26
                    }
                    _ = v23
                }
            }
            _ = fi14
            _ = v15
        }
    }
    _ = p0
    _ = p1
    return 255
}

func main() {
    var v1 int8 = 42
    v2 := make(map[uint32]bool)
    var v3 int = 2147483647
    var v4 string = ""
    v5 := make(map[uint32]uint32)
    var v6 string = "hello"
    for fi27 := 0; fi27 < 13; fi27++ {
        fmt.Println(v1, fi27, v4)
        fmt.Println(v4, v4, v4)
        _ = fi27
    }
    var v28 uint64 = 1
    v29 := GenS3{F0: 65535}
    v30 := GenS4{F0: 169, F1: (v28 * 0)}
    for fi31 := 0; fi31 < 7; fi31++ {
        v32 := GenS4{F0: 674, F1: (v28 % (((v28 & v28)) | 1))}
        v3 -= v3
        if (false || (fi31 != 128)) {
            fmt.Println(fi31, fi31, v28)
            v4 = "a"
            v33 := v5[uint32(v28)]
            _ = v33
        }
        _ = fi31
        _ = v32
    }
    var v34 int16 = 0
    var v35 float32 = float32((1.5 / 1.0))
    switch ((("\n\t" + "") + (("hello" + "foo bar") + v4)) + v6) {
    case "\n\t":
        for fi36 := 0; fi36 < 16; fi36++ {
            if (1 < -2147483648) {
                v37 := GenS4{F0: 2, F1: v28}
                v38 := GenS4{F0: 338, F1: v28}
                v39 := GenS5{F0: v34}
                _ = v37
                _ = v38
                _ = v39
            } else {
                for fi40 := 0; fi40 < 5; fi40++ {
                    v2[42] = false
                    _ = fi40
                }
                var v41 int16 = v34
                _ = v41
            }
            _ = fi36
        }
    case "foo bar":
        v6 = ("foo bar" + (v4 + v6))
    case "hello":
        var v42 uint32 = 0
        _ = v42
    }
    if (!(v3 >= (v3 + (v3 | 713)))) {
        var v43 uint32 = 65535
        v1 = v1
        _ = v43
    } else {
        v44 := GenS6{F0: 344, F1: true}
        _ = v44
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
