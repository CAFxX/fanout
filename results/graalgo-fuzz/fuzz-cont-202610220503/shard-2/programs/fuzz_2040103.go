package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint; F2 uint16 }
type GenS2 struct { F0 int8; F1 bool; F2 uint8 }
type GenS3 struct { F0 int; F1 uint8 }
type GenS4 struct { F0 uint8; F1 int; F2 string }
func genFn0(p0 uint64, p1 int64, p2 uint) uint32 {
    p2 = 0
    if (1 <= 0) {
        var v4 int16 = 42
        _ = v4
    } else {
        if false {
            p2 = uint(p0)
            p2 = p2
        }
        v5 := GenS1{F0: "a", F1: 542, F2: 758}
        v6 := []int{}
        _ = v5
        _ = v6
    }
    p0 = 18446744073709551615
    for fi7 := 0; fi7 < 20; fi7++ {
        var v8 uint32 = 42
        fmt.Println(v8, fi7)
        if ((-2.25 - 1e10) != float64(float32(0.0))) {
            for fi9 := 0; fi9 < 17; fi9++ {
                v8 += 846
                fmt.Println(p1)
                switch "\n\t" {
                case "hello":
                    v10 := GenS1{F0: (("foo bar" + "foo bar") + "hello"), F1: p2, F2: 228}
                    _ = v10
                case "foo bar":
                    var v11 int16 = -1
                    _ = v11
                }
                _ = fi9
            }
            for fi12 := 0; fi12 < 1; fi12++ {
                if true {
                    if false {
                        v8 = v8
                        var v13 bool = (((3.14159 + -2.25) > (1e10 * 0.1)) || true)
                        _ = v13
                    }
                }
                _ = fi12
            }
            v8 -= (v8 | (v8 / (((v8 % (((v8 - v8)) | 1))) | 1)))
        } else {
            for fi14 := 0; fi14 < 13; fi14++ {
                v15 := []int{}
                if len(v15) >= 2 {
                    v15 = v15[1:len(v15)]
                }
                _ = fi14
                _ = v15
            }
            switch -2147483648 {
            case -9223372036854775808:
                var v16 uint8 = 128
                _ = v16
            case 282:
                fmt.Println(p0)
            }
            fmt.Println(v8, p0)
        }
        _ = fi7
        _ = v8
    }
    fmt.Println(p1, p1, p2)
    _ = p0
    _ = p1
    _ = p2
    return 2
}
func genFn1() int64 {
    v17 := GenS2{F0: 1, F1: false, F2: 127}
    for fi18 := 0; fi18 < 10; fi18++ {
        var v19 int = fi18
        _ = fi18
        _ = v19
    }
    _ = v17
    return -462
}

func main() {
    var v1 uint32 = 526
    v2 := make(map[uint32]string)
    var v3 int64 = 42
    var v20 int8 = 127
    var v21 int16 = 32767
    v22 := GenS2{F0: int8(v1), F1: false, F2: 182}
    v23 := []int{}
    v24 := GenS3{F0: 9223372036854775807, F1: 127}
    var v25 uint8 = 42
    var v26 uint8 = (v25 << ((1) & 63))
    var v27 uint16 = 523
    v28 := GenS4{F0: 49, F1: 4294967295, F2: "foo bar"}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
}
