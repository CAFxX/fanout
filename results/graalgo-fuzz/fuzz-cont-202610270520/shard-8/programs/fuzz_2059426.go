package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int32; F1 int16 }
type GenS3 struct { F0 uint8; F1 int8 }
func genFn0(p0 uint16, p1 int16) int32 {
    v5 := []int{903}
    for fi6 := 0; fi6 < 7; fi6++ {
        v7 := GenS1{F0: (fi6 % ((fi6) | 1))}
        if true {
            if (-9223372036854775808 <= 42) {
                fmt.Println(p1, p1)
                if (false || true) {
                    fmt.Println(fi6, p1)
                    v8 := GenS1{F0: fi6}
                    _ = v8
                }
            }
            if ((756 != 1) && ((-(fi6)) > 2147483647)) {
                _ = len(v5)
            } else {
                var v9 uint = 4294967295
                switch (fi6 / ((fi6) | 1)) {
                case 2147483647:
                    if true {
                        v10 := GenS2{F0: int32(fi6), F1: p1}
                        _ = v10
                    }
                }
                switch ("" + ("hello" + ("" + "a"))) {
                case "\n\t":
                    var v11 uint = 8
                    _ = v11
                }
                _ = v9
            }
        }
        _ = fi6
        _ = v7
    }
    v12 := GenS2{F0: -1, F1: p1}
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v13 := GenS1{F0: 9223372036854775807}
    _ = p0
    _ = p1
    _ = v12
    _ = v13
    _ = v5
    return -901
}
func genFn1(p0 int, p1 int64, p2 bool) int {
    var v14 uint8 = 242
    for fi15 := 0; fi15 < 1; fi15++ {
        if false {
            var v16 float32 = 0.0
            var v17 int16 = -284
            _ = v16
            _ = v17
        } else {
            switch (p0 / (((^(fi15))) | 1)) {
            case 4294967295:
                v18 := []int{(p0 / ((128) | 1)), (-(p0)), fi15}
                _ = v18
            default:
                if p2 {
                    var v19 int16 = -220
                    _ = v19
                }
            }
            var v20 uint = 65535
            _ = v20
        }
        _ = fi15
    }
    v21 := []int{p0, (p0 - p0)}
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    _ = v21
    return 42
}

func main() {
    v1 := []uint32{2147483647, 42}
    var v2 int16 = -32768
    var v3 uint = 42
    var v4 uint64 = 4294967295
    for fi22 := 0; fi22 < 1; fi22++ {
        for fi23 := 0; fi23 < 17; fi23++ {
            v3 += v3
            switch (-(fi23)) {
            case 328:
                v2 = (v2 + (v2 ^ v2))
            default:
                var v24 uint64 = v4
                _ = v24
            }
            if ((42 != 296) && true) {
                var v25 int = fi22
                v3 += v3
                v4 += v4
                _ = v25
            }
            _ = fi23
        }
        v3 -= v3
        _ = fi22
    }
    for fi26 := 0; fi26 < 17; fi26++ {
        _ = len(v1)
        _ = fi26
    }
    var v27 uint16 = 42
    fmt.Println(v2, v2, v2)
    v3 += v3
    if len(v1) > 0 {
        v1[0] = 42
    }
    var v28 uint32 = 45
    for fi29 := 0; fi29 < 12; fi29++ {
        v30 := GenS3{F0: 2, F1: 42}
        fmt.Println(fi29, fi29)
        _ = fi29
        _ = v30
    }
    var v31 uint16 = v27
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v31
    _ = v4
}
