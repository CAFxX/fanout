package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string; F2 int }
type GenS2 struct { F0 uint32; F1 uint64; F2 uint32 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int32; F1 int64; F2 uint }
func genFn0(p0 int32) int {
    for fi6 := 0; fi6 < 8; fi6++ {
        v7 := GenS1{F0: 2, F1: "hello", F2: (fi6 | (fi6 ^ -2147483648))}
        p0 = (p0 * 287)
        v8 := []int{(fi6 << ((519) & 63))}
        _ = fi6
        _ = v7
        _ = v8
    }
    switch -2147483648 {
    case 9223372036854775807:
        for fi9 := 0; fi9 < 4; fi9++ {
            v10 := GenS2{F0: 2147483647, F1: 4294967295, F2: 507}
            var v11 float64 = 0.1
            if false {
                switch "a" {
                case "\n\t":
                    for fi12 := 0; fi12 < 7; fi12++ {
                        var v13 uint = 18446744073709551615
                        var v14 uint32 = 449
                        _ = fi12
                        _ = v13
                        _ = v14
                    }
                }
                var v15 int32 = p0
                _ = v15
            }
            _ = fi9
            _ = v10
            _ = v11
        }
    case 127:
        p0 = (p0 / ((-1) | 1))
    }
    _ = p0
    return 777
}

func main() {
    v1 := []int{4294967295, 2147483647}
    v2 := []int{-758, 660, 512, -8}
    v3 := make(map[string]string)
    v4 := []uint32{}
    var v5 int64 = 42
    var v16 uint64 = 348
    if (62 >= 797) {
        v5 += (v5 | v5)
        var v17 int8 = 20
        _ = v17
    } else {
        fmt.Println(v16)
        for fi18 := 0; fi18 < 3; fi18++ {
            var v19 float32 = (((float32(0.0) * (((0.1 / (-2.25 * 1.5)) / 1.0) - (1.5 / (3.14159 + (-2.25 * 3.14159))))) / 3.14159) + 1e10)
            var v20 string = "foo bar"
            _ = fi18
            _ = v19
            _ = v20
        }
    }
    v21 := GenS2{F0: 2147483648, F1: (v16 - (v16 % (((v16 >> ((65535) & 63))) | 1))), F2: 2}
    v16 -= v16
    for fi22 := 0; fi22 < 3; fi22++ {
        v23 := GenS3{F0: ((!false) && ((true || (true || false)) || false))}
        switch "foo bar" {
        case "":
            if (255 > 42) {
                v24 := GenS3{F0: (210 != 1)}
                _ = v24
            } else {
                v5 = v5
            }
        case "hello":
            v25 := GenS2{F0: 2147483648, F1: 4294967295, F2: 0}
            _ = v25
        case "a":
            var v26 uint64 = 538
            _ = v26
        }
        _ = fi22
        _ = v23
    }
    v5 = v5
    if (true || true) {
        for fi27 := 0; fi27 < 7; fi27++ {
            v28 := GenS4{F0: -2147483648}
            var v29 uint = 309
            for fi30 := 0; fi30 < 12; fi30++ {
                var v31 uint8 = 42
                v16 = (v16 * v16)
                _ = fi30
                _ = v31
            }
            _ = fi27
            _ = v28
            _ = v29
        }
        v32 := GenS2{F0: 2, F1: v16, F2: 727}
        _ = v32
    } else {
        if ((!(1 > -346)) && false) {
            for fi33 := 0; fi33 < 15; fi33++ {
                v34 := GenS5{F0: int32(v5), F1: 1, F2: 2}
                var v35 int16 = 0
                _ = fi33
                _ = v34
                _ = v35
            }
            _ = len(v3)
        }
        var v36 float32 = 1.5
        switch 32768 {
        case 42:
            fmt.Println(v36, v36, v5)
        case 744:
            var v37 int8 = -1
            _ = v37
        case 308:
            v16 -= v16
        default:
            _ = len(v1)
        }
        _ = v36
    }
    var v38 bool = (false || false)
    for fi39 := 0; fi39 < 4; fi39++ {
        switch (fi39 - 127) {
        case -2147483648:
            v2 = append(v2, 2147483647)
        }
        v16 = 694
        v40 := GenS5{F0: -2147483648, F1: v5, F2: 2}
        _ = fi39
        _ = v40
    }
    for fi41 := 0; fi41 < 9; fi41++ {
        switch (fi41 | fi41) {
        case 32767:
            switch fi41 {
            case -2147483648:
                _ = len(v1)
            case 4294967295:
                delete(v3, ("foo bar" + "a"))
            }
        case -2147483648:
            var v42 string = (("hello" + "hello") + "a")
            _ = v42
        default:
            _ = len(v3)
        }
        _ = fi41
    }
    v16 -= 2
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v2
    _ = v21
    _ = v3
    _ = v38
    _ = v4
    _ = v5
}
