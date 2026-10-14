package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8; F2 uint32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint; F1 uint64 }
type GenS4 struct { F0 int64; F1 uint }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint32; F1 uint64; F2 uint64 }
func genFn0(p0 int16) uint64 {
    for fi7 := 0; fi7 < 14; fi7++ {
        for fi8 := 0; fi8 < 10; fi8++ {
            if false {
                v9 := GenS1{F0: 1, F1: -128, F2: 2147483648}
                v10 := GenS1{F0: 65535, F1: -110, F2: 2147483647}
                v11 := GenS2{F0: -128}
                v12 := []int{244, (fi8 * 620), (fi7 - 4294967295), fi8}
                v13 := GenS2{F0: -128}
                _ = v10
                _ = v11
                _ = v12
                _ = v13
                _ = v9
            }
            v14 := GenS3{F0: 4294967295, F1: 18446744073709551615}
            _ = fi8
            _ = v14
        }
        for fi15 := 0; fi15 < 16; fi15++ {
            if true {
                v16 := GenS3{F0: 123, F1: 92}
                if true {
                    var v17 bool = false
                    p0 -= p0
                    var v18 uint8 = 255
                    _ = v17
                    _ = v18
                }
                if true {
                    if (-734 >= -620) {
                        var v19 int8 = -128
                        p0 -= -1
                        if (!(!true)) {
                            var v20 int32 = 42
                            _ = v20
                        }
                        v21 := []int{(fi15 & (fi7 / (((fi7 - (fi7 ^ 128))) | 1))), 9223372036854775807, fi7, (fi15 * fi7)}
                        v22 := GenS3{F0: 82, F1: 0}
                        _ = v19
                        _ = v21
                        _ = v22
                    } else {
                        fmt.Println(p0, fi7, p0)
                        var v23 float32 = (0.0 + float32((0.1 + float64(float32(3.14159)))))
                        _ = v23
                    }
                    if false {
                        switch "foo bar" {
                        case "foo bar":
                            v24 := []int{int(p0), fi7}
                            _ = v24
                        }
                    }
                    v25 := GenS4{F0: 0, F1: 0}
                    p0 = (p0 * (-(p0)))
                    p0 -= (-(p0))
                    _ = v25
                }
                _ = v16
            } else {
                p0 = (p0 % (((p0 * p0)) | 1))
                fmt.Println(p0, fi15)
                var v26 uint64 = 4294967295
                _ = v26
            }
            v27 := GenS3{F0: 365, F1: 2}
            _ = fi15
            _ = v27
        }
        if false {
            var v28 int64 = -710
            var v29 int64 = v28
            _ = v28
            _ = v29
        } else {
            p0 = (p0 % ((p0) | 1))
        }
        _ = fi7
    }
    var v30 int32 = 1
    v31 := []int{-164, 127, 9223372036854775807}
    _ = p0
    _ = v30
    _ = v31
    return 229
}
func genFn1() int64 {
    v32 := GenS2{F0: 0}
    var v33 uint8 = 128
    v33 -= 1
    fmt.Println(v33)
    _ = v32
    _ = v33
    return 1
}

func main() {
    var v1 uint16 = 1
    var v2 uint64 = 4294967295
    var v3 uint8 = 2
    var v4 int32 = -1
    v5 := make(map[int]string)
    v6 := make(map[string]bool)
    fmt.Println(v3, v2, v4)
    for fi34 := 0; fi34 < 11; fi34++ {
        v2 = (^(v2))
        v1 = 32767
        _ = fi34
    }
    v4 = v4
    for fi35 := 0; fi35 < 17; fi35++ {
        delete(v5, -411)
        if true {
            for fi36 := 0; fi36 < 9; fi36++ {
                v37 := GenS5{F0: (v3 & 2)}
                if (true || false) {
                    v2 = 1
                    var v38 int = 42
                    _ = v38
                }
                _ = fi36
                _ = v37
            }
            if ((1000 != -32768) || true) {
                fmt.Println(v2, v4, v1)
                var v39 int = fi35
                var v40 uint16 = (v1 - 923)
                _ = v39
                _ = v40
            }
        }
        _ = fi35
    }
    v41 := GenS6{F0: 292, F1: (v2 & v2), F2: 42}
    var v42 uint16 = (v1 >> ((18446744073709551615) & 63))
    fmt.Println(v42, v3, v4)
    var v43 int16 = -32768
    v1 = 65535
    var v44 int8 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v44
    _ = v5
    _ = v6
}
