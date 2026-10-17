package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint }
type GenS2 struct { F0 uint16; F1 uint; F2 uint16 }
type GenS3 struct { F0 uint64; F1 uint32 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint16; F1 uint64; F2 string }
type GenS6 struct { F0 uint8; F1 uint32; F2 int64 }
func genFn0(p0 int64, p1 uint) int8 {
    for fi6 := 0; fi6 < 20; fi6++ {
        if true {
            p0 = (p0 % (((p0 * p0)) | 1))
        }
        fmt.Println(fi6, fi6)
        _ = fi6
    }
    if (false || ((4294967295 != 4294967295) || true)) {
        if true {
            v7 := GenS1{F0: -414, F1: (p1 & 42)}
            var v8 bool = false
            _ = v7
            _ = v8
        }
    }
    p1 += p1
    _ = p0
    _ = p1
    return 127
}
func genFn1(p0 int64, p1 uint32) uint8 {
    var v9 float32 = (0.0 + (1e10 - 1e10))
    v10 := GenS2{F0: 274, F1: 2, F2: 65535}
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return 2
}

func main() {
    var v1 float32 = 0.0
    var v2 string = ""
    var v3 int32 = 2147483647
    v4 := make(map[uint32]uint32)
    var v5 int16 = 537
    switch 4294967295 {
    case 32768:
        for fi11 := 0; fi11 < 10; fi11++ {
            for fi12 := 0; fi12 < 2; fi12++ {
                var v13 int64 = 1
                _ = fi12
                _ = v13
            }
            _ = fi11
        }
    case 1:
        v14 := v4[2]
        _ = v14
    }
    v15 := []int{-757, -1, 42}
    fmt.Println(v2, v1, v5)
    v16 := GenS3{F0: 1, F1: 4294967295}
    if (true || false) {
        v17 := GenS1{F0: 1, F1: 1}
        _ = v17
    }
    if false {
        if false {
            fmt.Println(v1, v3, v5)
            for fi18 := 0; fi18 < 4; fi18++ {
                v15 = append(v15, (fi18 / (((fi18 ^ 9223372036854775807)) | 1)))
                if len(v15) >= 2 {
                    v15 = v15[1:len(v15)]
                }
                v19 := GenS3{F0: 18446744073709551615, F1: 65535}
                _ = fi18
                _ = v19
            }
        }
        var v20 int = 32767
        var v21 uint16 = 42
        _ = v20
        _ = v21
    } else {
        if len(v15) > 0 {
            v15[0] = -939
        }
        for fi22 := 0; fi22 < 16; fi22++ {
            var v23 uint32 = 2
            for fi24 := 0; fi24 < 16; fi24++ {
                v2 = v2
                var v25 float32 = float32(3.14159)
                if (-2.25 != -2.25) {
                    if true {
                        if (!(true && false)) {
                            v26 := GenS3{F0: 740, F1: 37}
                            v27 := GenS4{F0: 1}
                            var v28 int32 = (v3 >> ((65535) & 63))
                            _ = v26
                            _ = v27
                            _ = v28
                        }
                    }
                    v5 -= v5
                }
                _ = fi24
                _ = v25
            }
            _ = fi22
            _ = v23
        }
    }
    if (-9223372036854775808 < -9223372036854775808) {
        fmt.Println(v3, v3, v2)
    }
    v29 := GenS5{F0: 32767, F1: 902, F2: ("foo bar" + "\n\t")}
    fmt.Println(v5, v3, v1)
    v30 := GenS6{F0: 2, F1: 903, F2: -9223372036854775808}
    v1 = v1
    v1 = (0.0 * (v1 / (v1 - v1)))
    for fi31 := 0; fi31 < 7; fi31++ {
        if ((genFn1(797, 42) <= 255) || ((1 >= 65535) || (-119 <= -1))) {
            if len(v15) >= 2 {
                v15 = v15[1:len(v15)]
            }
            var v32 int = fi31
            delete(v4, 620)
            _ = v32
        } else {
            switch -1 {
            case -609:
                v2 = v2
            }
        }
        var v33 uint64 = 631
        _ = fi31
        _ = v33
    }
    if false {
        var v34 float64 = (1e10 - float64(1e10))
        var v35 uint8 = 188
        _ = v34
        _ = v35
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
