package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32 }
type GenS2 struct { F0 uint8; F1 uint64; F2 bool }
type GenS3 struct { F0 bool; F1 uint16; F2 int }
func genFn0(p0 bool, p1 string) uint32 {
    fmt.Println(p1)
    if (42 >= -1) {
        for fi6 := 0; fi6 < 6; fi6++ {
            p0 = false
            if (2147483648 < 55) {
                var v7 int32 = -1
                if (true || true) {
                    for fi8 := 0; fi8 < 17; fi8++ {
                        switch fi8 {
                        case 640:
                            fmt.Println(fi6, fi6)
                        default:
                            p0 = p0
                        }
                        p0 = (-128 > -121)
                        var v9 int16 = -32768
                        _ = fi8
                        _ = v9
                    }
                }
                v10 := GenS1{F0: 0, F1: 1}
                _ = v10
                _ = v7
            } else {
                for fi11 := 0; fi11 < 2; fi11++ {
                    if p0 {
                        var v12 uint8 = 112
                        _ = v12
                    }
                    _ = fi11
                }
                v13 := []int{}
                fmt.Println(p1)
                _ = v13
            }
            var v14 float64 = 1.5
            _ = fi6
            _ = v14
        }
        if true {
            if (p0 || p0) {
                v15 := []int{}
                _ = v15
            } else {
                var v16 int8 = 42
                v17 := []int{}
                p1 = p1
                _ = v16
                _ = v17
            }
            v18 := GenS2{F0: 128, F1: 328, F2: (true || (!false))}
            switch "a" {
            case "a":
                var v19 bool = true
                _ = v19
            case "foo bar":
                p0 = ((p0 && (1.5 > float32((3.14159 * ((3.14159 + 0.0) * float64(((3.14159 / 0.1) * 3.14159))))))) || false)
            case "hello":
                v20 := GenS3{F0: p0, F1: 2, F2: -1}
                _ = v20
            }
            _ = v18
        } else {
            fmt.Println(p1)
        }
    }
    _ = p0
    _ = p1
    return 2
}
func genFn1(p0 uint16, p1 bool, p2 bool) uint64 {
    p0 = (p0 / (((p0 - 255)) | 1))
    fmt.Println(p0, p0)
    _ = p0
    _ = p1
    _ = p2
    return 0
}

func main() {
    var v1 int16 = 489
    v2 := make(map[uint32]uint32)
    var v3 uint8 = 1
    var v4 uint16 = 65535
    var v5 float64 = 1e10
    v5 = ((1.5 - v5) - (v5 * (-2.25 / (1.5 - 1e10))))
    var v21 uint8 = (^(v3))
    v22 := GenS3{F0: false, F1: v4, F2: 4294967295}
    v23 := []int{}
    var v24 int64 = 0
    v25 := GenS1{F0: 65535, F1: 2147483647}
    v1 += -1
    for fi26 := 0; fi26 < 18; fi26++ {
        v5 -= (float64(0.0) * 1e10)
        _ = fi26
    }
    var v27 float32 = 1e10
    fmt.Println(v1)
    v24 = v24
    var v28 float32 = 1.5
    delete(v2, genFn0(true, (((("hello" + ("hello" + "foo bar")) + "foo bar") + "foo bar") + "")))
    if false {
        if (!(0 > 18446744073709551615)) {
            v3 += v21
            v27 += v27
            v29 := GenS1{F0: 42, F1: 0}
            _ = v29
        } else {
            fmt.Println(v1, v27, v28)
            if (1 > 1) {
                v24 -= v24
                v3 = v21
                var v30 float32 = ((1e10 - v27) * 0.1)
                _ = v30
            }
            fmt.Println(v21, v5, v1)
        }
    } else {
        if len(v23) > 0 {
            v23[0] = 32768
        }
        for fi31 := 0; fi31 < 5; fi31++ {
            v4 += v4
            fmt.Println(v3)
            v4 += 1
            _ = fi31
        }
    }
    switch 4294967295 {
    case -2147483648:
        fmt.Println(v28, v5, v24)
    case 4294967295:
        v23 = append(v23, 32768)
    case 32768:
        v32 := v2[genFn0(true, "\n\t")]
        _ = v32
    }
    var v33 bool = ((v4 & (v4 / ((v4) | 1))) < (^(v4)))
    for fi34 := 0; fi34 < 20; fi34++ {
        switch (fi34 & 32767) {
        case 1:
            v35 := GenS1{F0: v4, F1: int32(v24)}
            _ = v35
        case 4294967295:
            v36 := GenS3{F0: v33, F1: v4, F2: (fi34 >> ((55) & 63))}
            _ = v36
        case -9223372036854775808:
            var v37 float32 = v27
            _ = v37
        default:
            for fi38 := 0; fi38 < 14; fi38++ {
                v3 = v21
                v39 := GenS1{F0: v4, F1: -2147483648}
                var v40 int8 = 1
                _ = fi38
                _ = v39
                _ = v40
            }
        }
        v1 += 0
        _ = fi34
    }
    for fi41 := 0; fi41 < 18; fi41++ {
        fmt.Println(v4, v24, v28)
        var v42 int = (fi41 << ((4294967295) & 63))
        _ = fi41
        _ = v42
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v24:", v24)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v27
    _ = v28
    _ = v3
    _ = v33
    _ = v4
    _ = v5
}
