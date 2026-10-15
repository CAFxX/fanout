package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8; F2 uint8 }
type GenS2 struct { F0 int16; F1 uint64 }
type GenS3 struct { F0 int8; F1 int16 }
func genFn0(p0 uint8, p1 string, p2 int64) uint8 {
    switch p1 {
    case "a":
        if true {
            for fi7 := 0; fi7 < 1; fi7++ {
                v8 := []int{fi7}
                p2 = p2
                _ = fi7
                _ = v8
            }
            p0 = (p0 + p0)
        } else {
            if true {
                switch 9223372036854775807 {
                case 1:
                    fmt.Println(p2, p1, p1)
                case -292:
                    var v9 int = -586
                    _ = v9
                default:
                    fmt.Println(p2, p0)
                }
            } else {
                p1 = p1
                v10 := GenS1{F0: 1, F1: -1, F2: (^(p0))}
                var v11 int8 = 127
                _ = v10
                _ = v11
            }
            var v12 uint = 4294967295
            _ = v12
        }
    case "\n\t":
        var v13 int64 = p2
        _ = v13
    }
    p1 = p1
    _ = p0
    _ = p1
    _ = p2
    return (p0 ^ (p0 ^ 255))
}

func main() {
    var v1 float64 = (1e10 - (3.14159 / 1.0))
    var v2 uint16 = 42
    var v3 uint8 = 42
    var v4 int64 = -901
    v5 := []int{}
    v6 := make(map[int]uint32)
    switch "foo bar" {
    case "hello":
        var v14 int64 = -51
        _ = v14
    }
    fmt.Println(v3, v1)
    _ = len(v5)
    for fi15 := 0; fi15 < 5; fi15++ {
        if false {
            var v16 int32 = 1
            v17 := GenS1{F0: 4294967295, F1: 0, F2: v3}
            _ = v16
            _ = v17
        } else {
            if false {
                v18 := GenS1{F0: 2, F1: int8(fi15), F2: v3}
                if len(v5) > 0 {
                    v5[0] = fi15
                }
                _ = v18
            }
        }
        _ = fi15
    }
    var v19 float32 = (1e10 + 0.0)
    v2 = (v2 + (v2 ^ (v2 / ((175) | 1))))
    var v20 uint32 = 1
    v21 := GenS2{F0: 42, F1: 4294967295}
    v4 = (v4 & -9223372036854775808)
    var v22 bool = false
    if true {
        var v23 float64 = v1
        if v22 {
            var v24 int64 = 0
            var v25 string = (("\n\t" + "hello") + "")
            v5 = append(v5, 1)
            _ = v24
            _ = v25
        } else {
            var v26 uint16 = v2
            fmt.Println(v1)
            _ = v26
        }
        v27 := GenS2{F0: 0, F1: 18446744073709551615}
        _ = v23
        _ = v27
    }
    v20 = 999
    if v22 {
        for fi28 := 0; fi28 < 4; fi28++ {
            fmt.Println(v1, v20, fi28)
            if v22 {
                if (v22 || (v22 && false)) {
                    var v29 uint32 = v20
                    v30 := GenS2{F0: 271, F1: 42}
                    var v31 float32 = float32(((v1 / float64(float32(v1))) / v1))
                    _ = v29
                    _ = v30
                    _ = v31
                }
                v20 -= (v20 / (((v20 ^ v20)) | 1))
            }
            v32 := v6[fi28]
            _ = fi28
            _ = v32
        }
        for fi33 := 0; fi33 < 9; fi33++ {
            for fi34 := 0; fi34 < 18; fi34++ {
                v20 = v20
                var v35 string = ("hello" + "")
                if (int64(v3) >= v4) {
                    v3 += genFn0((v3 - 42), v35, (v4 | v4))
                    var v36 int = 0
                    _ = v36
                }
                _ = fi34
                _ = v35
            }
            _ = fi33
        }
        if (v22 && ((v3 != 19) && (v1 > v1))) {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        }
    } else {
        fmt.Println(v4)
        v22 = v22
    }
    _ = len(v6)
    fmt.Println(v20, v1)
    if v22 {
        var v37 int16 = 1000
        var v38 int64 = v4
        _ = v37
        _ = v38
    } else {
        if v22 {
            if len(v5) > 0 {
                v5[0] = 32768
            }
        }
        v39 := GenS3{F0: int8(v2), F1: 1000}
        for fi40 := 0; fi40 < 17; fi40++ {
            switch int(v2) {
            case 1:
                var v41 int32 = -387
                _ = v41
            default:
                fmt.Println(fi40, v20)
            }
            if (-2.25 > (float32((((0.1 - (((float64(v19) * v1) + v1) + v1)) + (v1 * v1)) + (v1 * v1))) * 1e10)) {
                var v42 int = (-(fi40))
                _ = v42
            }
            _ = fi40
        }
        _ = v39
    }
    var v43 uint = 0
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v43
    _ = v5
    _ = v6
}
