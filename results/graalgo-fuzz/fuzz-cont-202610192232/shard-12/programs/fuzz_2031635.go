package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint; F2 uint32 }
type GenS2 struct { F0 uint; F1 uint8; F2 int8 }
type GenS3 struct { F0 int64 }
func genFn0(p0 uint) int8 {
    v4 := GenS1{F0: 580, F1: 894, F2: 2147483647}
    if true {
        p0 = p0
        var v5 uint16 = 676
        for fi6 := 0; fi6 < 19; fi6++ {
            var v7 uint8 = 235
            if (fi6 > (fi6 ^ 127)) {
                if (2 == 42) {
                    v8 := GenS2{F0: (p0 % (((p0 + p0)) | 1)), F1: (v7 + (v7 | 255)), F2: 3}
                    var v9 uint = p0
                    _ = v8
                    _ = v9
                }
                var v10 int32 = -2147483648
                _ = v10
            }
            _ = fi6
            _ = v7
        }
        _ = v5
    }
    _ = p0
    _ = v4
    return 1
}
func genFn1(p0 int32, p1 uint16, p2 bool) uint64 {
    switch -2147483648 {
    case 42:
        for fi11 := 0; fi11 < 17; fi11++ {
            var v12 int8 = -128
            for fi13 := 0; fi13 < 14; fi13++ {
                var v14 uint = 0
                var v15 int32 = (p0 / (((p0 & (p0 % ((-1) | 1)))) | 1))
                v16 := GenS2{F0: (^(v14)), F1: 255, F2: -1}
                var v17 int8 = 64
                var v18 uint = v14
                _ = fi13
                _ = v14
                _ = v15
                _ = v16
                _ = v17
                _ = v18
            }
            _ = fi11
            _ = v12
        }
    case 1:
        p1 -= (p1 / (((p1 / ((p1) | 1))) | 1))
    default:
        var v19 uint = 18446744073709551615
        _ = v19
    }
    p0 += p0
    _ = p0
    _ = p1
    _ = p2
    return 4294967295
}

func main() {
    var v1 uint8 = 128
    var v2 int = -326
    var v3 uint32 = 65535
    fmt.Println(v2, v3)
    switch 853 {
    case 127:
        var v20 uint8 = 156
        _ = v20
    }
    for fi21 := 0; fi21 < 20; fi21++ {
        v22 := []int{(fi21 + 1), (v2 << ((65535) & 63)), (fi21 + v2), v2}
        fmt.Println(v2, v1, v3)
        switch 128 {
        case -1:
            _ = len(v22)
        case 127:
            var v23 string = "\n\t"
            _ = v23
        }
        _ = fi21
        _ = v22
    }
    if (!false) {
        var v24 float32 = (float32(1e10) * float32(float64(1.5)))
        if true {
            var v25 int64 = 1
            v26 := []int{(-(v2))}
            switch ("foo bar" + "foo bar") {
            case "a":
                fmt.Println(v1)
            default:
                switch "foo bar" {
                case "":
                    var v27 int64 = (-(v25))
                    _ = v27
                }
            }
            _ = v25
            _ = v26
        }
        for fi28 := 0; fi28 < 9; fi28++ {
            if true {
                if true {
                    v3 += v3
                }
                v29 := []int{32768, (^(fi28)), (v2 / (((^(v2))) | 1)), 9223372036854775807}
                v3 = (v3 - (v3 / (((v3 * v3)) | 1)))
                _ = v29
            } else {
                v30 := GenS1{F0: 65535, F1: 469, F2: (v3 >> ((42) & 63))}
                _ = v30
            }
            var v31 int32 = 1
            _ = fi28
            _ = v31
        }
        _ = v24
    }
    v32 := []int{}
    for fi33 := 0; fi33 < 12; fi33++ {
        v3 = (v3 / ((2147483648) | 1))
        v34 := GenS3{F0: 9223372036854775807}
        _ = fi33
        _ = v34
    }
    for fi35 := 0; fi35 < 4; fi35++ {
        fmt.Println(v3, fi35, v3)
        fmt.Println(v2, v3)
        _ = fi35
    }
    switch v2 {
    case -2147483648:
        v1 = 255
    case -239:
        _ = len(v32)
    case 128:
        fmt.Println(v2, v1, v1)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v32:", len(v32), v32)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
}
