package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint }
type GenS2 struct { F0 uint; F1 uint32 }
type GenS3 struct { F0 int64; F1 uint16; F2 int64 }
type GenS4 struct { F0 string; F1 int; F2 uint8 }
type GenS5 struct { F0 int8; F1 int32; F2 uint16 }
func genFn0(p0 uint16, p1 uint32) int {
    var v7 uint32 = p1
    p1 = p1
    var v8 uint16 = p0
    if (0.0 != (1e10 / (float32(float64(3.14159)) / 1e10))) {
        var v9 string = "\n\t"
        _ = v9
    } else {
        if false {
            if false {
                var v10 uint = 142
                switch "" {
                case "a":
                    for fi11 := 0; fi11 < 6; fi11++ {
                        var v12 int64 = 774
                        v13 := []int{149}
                        v14 := GenS1{F0: 9223372036854775807, F1: (v10 >> ((255) & 63))}
                        _ = fi11
                        _ = v12
                        _ = v13
                        _ = v14
                    }
                }
                _ = v10
            }
        }
        if false {
            v8 -= (p0 | 42)
        }
    }
    for fi15 := 0; fi15 < 19; fi15++ {
        var v16 uint = 65535
        v17 := []int{int(p1), (fi15 / ((-999) | 1))}
        _ = fi15
        _ = v16
        _ = v17
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 2147483647
}

func main() {
    v1 := []uint32{511, 2147483647, 2147483648, 49}
    v2 := make(map[string]int)
    var v3 int16 = -32768
    var v4 uint16 = 0
    v5 := make(map[int]uint32)
    v6 := []int{127, 2147483647, 128}
    if false {
        for fi18 := 0; fi18 < 17; fi18++ {
            switch "\n\t" {
            case "hello":
                v19 := GenS1{F0: fi18, F1: 65535}
                _ = v19
            }
            var v20 uint8 = 255
            var v21 int16 = (-(v3))
            _ = fi18
            _ = v20
            _ = v21
        }
        v22 := GenS1{F0: 32767, F1: 65535}
        var v23 uint16 = (^(v4))
        _ = v22
        _ = v23
    } else {
        v24 := GenS1{F0: 32767, F1: 1}
        _ = v24
    }
    for fi25 := 0; fi25 < 4; fi25++ {
        v4 += 134
        _ = fi25
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    if len(v1) > 0 {
        v1[0] = 439
    }
    for fi26 := 0; fi26 < 17; fi26++ {
        fmt.Println(fi26)
        _ = fi26
    }
    var v27 int16 = 32767
    if false {
        v27 = (v27 % ((-1) | 1))
        fmt.Println(v27, v4, v3)
    }
    if true {
        var v28 int32 = 42
        _ = v28
    } else {
        v27 = (v27 * (v27 ^ (v27 / ((v3) | 1))))
        for fi29 := 0; fi29 < 15; fi29++ {
            v30 := GenS1{F0: (^(fi29)), F1: 2}
            _ = fi29
            _ = v30
        }
    }
    if (((1e10 / (float64(1.5) * -2.25)) - (float64(3.14159) / ((1.5 - 1e10) * (float64((0.1 - 3.14159)) - 0.1)))) < float64(1.5)) {
        v27 += 1
        v31 := GenS2{F0: 42, F1: 839}
        _ = v31
    } else {
        v32 := GenS1{F0: 32768, F1: 2}
        fmt.Println(v4, v27, v27)
        _ = v32
    }
    if ((-(v3)) <= v27) {
        v33 := GenS3{F0: -9223372036854775808, F1: v4, F2: 42}
        for fi34 := 0; fi34 < 17; fi34++ {
            for fi35 := 0; fi35 < 10; fi35++ {
                var v36 uint64 = 934
                var v37 int16 = (-(v3))
                v38 := GenS2{F0: uint(v4), F1: 2}
                _ = fi35
                _ = v36
                _ = v37
                _ = v38
            }
            v3 = 898
            _ = fi34
        }
        _ = v33
    }
    for fi39 := 0; fi39 < 16; fi39++ {
        v4 += v4
        _ = len(v6)
        _ = fi39
    }
    var v40 bool = true
    switch 90 {
    case -244:
        for fi41 := 0; fi41 < 5; fi41++ {
            var v42 uint = 1
            _ = fi41
            _ = v42
        }
    case 4294967295:
        var v43 int64 = 9223372036854775807
        _ = v43
    }
    fmt.Println(v3)
    if v40 {
        var v44 int32 = -562
        v5[0] = 0
        v45 := GenS4{F0: (("foo bar" + "\n\t") + ""), F1: 892, F2: 0}
        _ = v44
        _ = v45
    } else {
        v46 := GenS5{F0: -128, F1: -788, F2: 2}
        _ = v46
    }
    v3 = (v27 % ((v3) | 1))
    v47 := GenS2{F0: 255, F1: 472}
    v3 -= v27
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v40
    _ = v47
    _ = v5
    _ = v6
}
