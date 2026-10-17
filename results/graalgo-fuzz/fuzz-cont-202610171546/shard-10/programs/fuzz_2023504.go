package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64 }
type GenS2 struct { F0 string; F1 uint16; F2 int16 }
type GenS3 struct { F0 uint8; F1 int64 }
func genFn0(p0 int64, p1 uint) int16 {
    v4 := GenS1{F0: (p0 << ((42) & 63)), F1: 2}
    p1 = p1
    v5 := GenS2{F0: "", F1: 487, F2: 32767}
    for fi6 := 0; fi6 < 2; fi6++ {
        p1 -= (p1 + p1)
        _ = fi6
    }
    if (-447 > -32768) {
        for fi7 := 0; fi7 < 19; fi7++ {
            var v8 bool = (2 < 1)
            fmt.Println(p1)
            _ = fi7
            _ = v8
        }
    } else {
        p0 -= (p0 / ((p0) | 1))
        switch 634 {
        case 42:
            if (-1 > -425) {
                for fi9 := 0; fi9 < 3; fi9++ {
                    if (((1.5 + (0.1 - 0.1)) == float64(3.14159)) || true) {
                        fmt.Println(fi9, p1)
                        p1 -= (^(p1))
                        fmt.Println(p0, p0, p0)
                    }
                    if false {
                        p0 -= (^(p0))
                    }
                    _ = fi9
                }
                var v10 int8 = 0
                _ = v10
            }
        case 213:
            v11 := GenS3{F0: 1, F1: (p0 - p0)}
            _ = v11
        case 9223372036854775807:
            var v12 uint = 0
            _ = v12
        default:
            p0 = p0
        }
        v13 := []int{32767, 1, -2147483648}
        _ = v13
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 0
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint = 4294967295
    v3 := make(map[string]int)
    var v14 uint = (v2 << (((v2 * 18446744073709551615)) & 63))
    var v15 float32 = 3.14159
    v16 := v1[2]
    v1[65535] = (2 < 42)
    if v16 {
        v17 := GenS3{F0: 89, F1: 0}
        v18 := GenS3{F0: 0, F1: -1}
        v19 := GenS3{F0: 121, F1: -9223372036854775808}
        _ = v17
        _ = v18
        _ = v19
    } else {
        v16 = v16
        var v20 uint32 = 2147483647
        v21 := []int{2147483647}
        _ = v20
        _ = v21
    }
    if (-2147483648 <= 0) {
        v22 := GenS3{F0: 255, F1: -9223372036854775808}
        _ = v22
    } else {
        var v23 uint32 = 2
        var v24 int64 = -9223372036854775808
        v2 = (v2 & v14)
        _ = v23
        _ = v24
    }
    var v25 int = 127
    for fi26 := 0; fi26 < 19; fi26++ {
        fmt.Println(fi26, v14)
        for fi27 := 0; fi27 < 16; fi27++ {
            v3["hello"] = (fi26 >> ((v2) & 63))
            v28 := v3["foo bar"]
            _ = fi27
            _ = v28
        }
        if (2 > 2) {
            for fi29 := 0; fi29 < 19; fi29++ {
                v30 := []int{0, 9223372036854775807}
                _ = fi29
                _ = v30
            }
            switch "" {
            case "\n\t":
                for fi31 := 0; fi31 < 3; fi31++ {
                    v16 = (2 < 2147483647)
                    v32 := []int{fi26, -1, (fi26 % ((v25) | 1))}
                    _ = fi31
                    _ = v32
                }
            case "hello":
                v25 += v25
            case "":
                v14 = 1
            }
            v3["foo bar"] = (v25 & (^(v25)))
        }
        _ = fi26
    }
    for fi33 := 0; fi33 < 19; fi33++ {
        var v34 int = v25
        _ = fi33
        _ = v34
    }
    v1[181] = v16
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v25
    _ = v3
}
