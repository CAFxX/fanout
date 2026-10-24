package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int }
type GenS2 struct { F0 uint32; F1 uint16; F2 int32 }
type GenS3 struct { F0 uint16; F1 uint64; F2 uint8 }
type GenS4 struct { F0 int32; F1 int16; F2 uint16 }
type GenS5 struct { F0 int32 }
func genFn0() uint8 {
    var v7 int64 = -9223372036854775808
    if false {
        v7 = -728
    } else {
        if (65535 <= 255) {
            v8 := []int{-2147483648}
            _ = v8
        }
    }
    fmt.Println(v7)
    var v9 int64 = v7
    fmt.Println(v7, v9)
    _ = v7
    _ = v9
    return 0
}
func genFn1(p0 int64) uint64 {
    for fi10 := 0; fi10 < 2; fi10++ {
        p0 += -1
        p0 = (p0 << ((828) & 63))
        _ = fi10
    }
    switch -910 {
    case -1:
        p0 = p0
    default:
        v11 := GenS1{F0: 795, F1: -2147483648}
        _ = v11
    }
    var v12 float64 = -2.25
    var v13 string = "foo bar"
    var v14 uint8 = 2
    _ = p0
    _ = v12
    _ = v13
    _ = v14
    return 305
}

func main() {
    var v1 uint32 = 346
    v2 := make(map[int]string)
    var v3 float32 = (((0.0 - 0.0) * 3.14159) / -2.25)
    var v4 uint32 = (^(v1))
    var v5 uint32 = v4
    v6 := make(map[int]int)
    switch -651 {
    case -745:
        for fi15 := 0; fi15 < 18; fi15++ {
            var v16 int64 = -264
            _ = fi15
            _ = v16
        }
    default:
        v5 = 0
    }
    v1 -= 2
    v17 := []int{32768}
    for fi18 := 0; fi18 < 4; fi18++ {
        v5 = 2147483648
        _ = len(v2)
        fmt.Println(v4, v1, v4)
        _ = fi18
    }
    v19 := v2[712]
    for fi20 := 0; fi20 < 16; fi20++ {
        v21 := GenS2{F0: v5, F1: 748, F2: 682}
        switch v19 {
        case "hello":
            for fi22 := 0; fi22 < 14; fi22++ {
                var v23 int = fi22
                _ = fi22
                _ = v23
            }
        case "":
            for fi24 := 0; fi24 < 19; fi24++ {
                if (fi20 != (fi20 ^ (fi20 * fi20))) {
                    switch (^(fi24)) {
                    case 127:
                        switch (v19 + ("\n\t" + v19)) {
                        case "\n\t":
                            var v25 uint64 = genFn1(-1)
                            _ = v25
                        case "a":
                            var v26 uint8 = 127
                            _ = v26
                        default:
                            v5 += (v1 | (v1 | v5))
                        }
                    }
                    v27 := GenS3{F0: 2, F1: 2, F2: 2}
                    v5 += (v4 % ((v1) | 1))
                    _ = v27
                }
                fmt.Println(v3, fi20, v4)
                _ = len(v17)
                _ = fi24
            }
        default:
            fmt.Println(fi20)
        }
        _ = fi20
        _ = v21
    }
    v28 := GenS4{F0: 2147483647, F1: 42, F2: 42}
    v2[-9223372036854775808] = v19
    v5 = (^(v5))
    v1 = (^(v1))
    v29 := GenS3{F0: 431, F1: 640, F2: 211}
    for fi30 := 0; fi30 < 7; fi30++ {
        var v31 int8 = -50
        v32 := GenS5{F0: -1}
        _ = fi30
        _ = v31
        _ = v32
    }
    if true {
        if (int64(v5) < -9223372036854775808) {
            var v33 int64 = -9223372036854775808
            var v34 int = 0
            _ = v33
            _ = v34
        }
        var v35 float32 = v3
        _ = v35
    } else {
        var v36 int16 = 32767
        var v37 float64 = 0.0
        delete(v2, -1)
        _ = v36
        _ = v37
    }
    switch 128 {
    case 32767:
        var v38 int = 9223372036854775807
        _ = v38
    case 32768:
        v39 := GenS4{F0: 1, F1: -471, F2: 545}
        _ = v39
    case -9223372036854775808:
        v40 := GenS3{F0: 66, F1: 685, F2: 0}
        _ = v40
    }
    if false {
        for fi41 := 0; fi41 < 3; fi41++ {
            if len(v17) >= 2 {
                v17 = v17[1:len(v17)]
            }
            _ = fi41
        }
    }
    for fi42 := 0; fi42 < 19; fi42++ {
        if (v19 < "a") {
            var v43 float64 = 3.14159
            v43 = v43
            v43 = (v43 * (v43 - v43))
            _ = v43
        } else {
            v44 := v6[fi42]
            v45 := GenS5{F0: -291}
            _ = v44
            _ = v45
        }
        _ = fi42
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v17
    _ = v19
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
