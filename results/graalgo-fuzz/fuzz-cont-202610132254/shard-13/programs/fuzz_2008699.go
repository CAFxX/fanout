package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16; F2 uint32 }
type GenS2 struct { F0 int16; F1 uint8; F2 uint64 }
type GenS3 struct { F0 int32; F1 uint8; F2 uint64 }
type GenS4 struct { F0 uint16; F1 int64; F2 int16 }
type GenS5 struct { F0 string; F1 uint32; F2 uint }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 bool }
func genFn0(p0 int8, p1 int8, p2 int64) int64 {
    for fi6 := 0; fi6 < 19; fi6++ {
        var v7 uint = 300
        _ = fi6
        _ = v7
    }
    v8 := []int{558, 1, 2147483647}
    switch "foo bar" {
    case "hello":
        p2 += (p2 / ((1) | 1))
    case "a":
        fmt.Println(p2)
    }
    if (true || true) {
        if len(v8) > 0 {
            v8[0] = 282
        }
    }
    var v9 float64 = 1.5
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return 750
}
func genFn1(p0 uint16, p1 uint32, p2 uint8) uint8 {
    if (1e10 != float64(1e10)) {
        v10 := GenS1{F0: 4294967295, F1: 1, F2: (p1 / ((777) | 1))}
        _ = v10
    } else {
        for fi11 := 0; fi11 < 3; fi11++ {
            var v12 uint = 450
            switch (("" + "hello") + "foo bar") {
            case "hello":
                var v13 int16 = 639
                _ = v13
            case "foo bar":
                for fi14 := 0; fi14 < 10; fi14++ {
                    v15 := GenS1{F0: v12, F1: -1, F2: (p1 - p1)}
                    var v16 int16 = 32767
                    if (!true) {
                        if (18446744073709551615 <= 353) {
                            v17 := GenS1{F0: v12, F1: 32767, F2: 2147483647}
                            v12 = 511
                            _ = v17
                        } else {
                            var v18 uint8 = p2
                            _ = v18
                        }
                    } else {
                        var v19 uint64 = 4294967295
                        if false {
                            v16 += v16
                        } else {
                            v20 := GenS2{F0: 1000, F1: (p2 << ((v12) & 63)), F2: v19}
                            if true {
                                var v21 int8 = 0
                                fmt.Println(p1, p0, v12)
                                _ = v21
                            }
                            v22 := []int{716, fi14, (-(fi14))}
                            _ = v20
                            _ = v22
                        }
                        _ = v19
                    }
                    if ((true && (((v16 - (v16 ^ v16)) == v16) || (false || true))) || ((v16 << ((4294967295) & 63)) != (v16 & (v16 - (v16 | (v16 >> (((v12 >> ((v12) & 63))) & 63))))))) {
                        var v23 uint = v12
                        _ = v23
                    }
                    v12 += (v12 >> ((v12) & 63))
                    var v24 float32 = 1e10
                    _ = fi14
                    _ = v15
                    _ = v16
                    _ = v24
                }
            case "\n\t":
                fmt.Println(v12, v12, v12)
            default:
                p1 += p1
            }
            var v25 int16 = 303
            _ = fi11
            _ = v12
            _ = v25
        }
        var v26 uint16 = p0
        _ = v26
    }
    if (true && false) {
        for fi27 := 0; fi27 < 12; fi27++ {
            v28 := GenS1{F0: 742, F1: 0, F2: (p1 << ((0) & 63))}
            v29 := []int{fi27, 32768, (fi27 % (((fi27 + -9223372036854775808)) | 1))}
            v30 := GenS3{F0: int32(p0), F1: 127, F2: 0}
            _ = fi27
            _ = v28
            _ = v29
            _ = v30
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return p2
}

func main() {
    var v1 uint16 = 65535
    var v2 int32 = -1
    var v3 int16 = 1
    v4 := make(map[string]string)
    var v5 bool = false
    v2 -= v2
    fmt.Println(v2, v1, v5)
    for fi31 := 0; fi31 < 7; fi31++ {
        fmt.Println(v3)
        var v32 int = fi31
        _ = fi31
        _ = v32
    }
    v33 := []int{42}
    v34 := v4["\n\t"]
    for fi35 := 0; fi35 < 19; fi35++ {
        v36 := GenS3{F0: (v2 & (-(v2))), F1: 89, F2: 863}
        _ = fi35
        _ = v36
    }
    v34 = v34
    switch v34 {
    case "foo bar":
        v34 = v34
    case "\n\t":
        if (1e10 == 0.1) {
            if v5 {
                v37 := GenS4{F0: (^(v1)), F1: 0, F2: 367}
                switch ("foo bar" + v34) {
                case "\n\t":
                    if (v5 || true) {
                        var v38 uint = 4294967295
                        _ = v38
                    }
                case "hello":
                    v5 = (-1 != 0)
                default:
                    v39 := GenS5{F0: "a", F1: 4294967295, F2: 18446744073709551615}
                    _ = v39
                }
                _ = v37
            }
            v40 := GenS6{F0: 335}
            _ = v40
        } else {
            v41 := GenS7{F0: (v5 && (4294967295 >= 4294967295))}
            v34 = v34
            _ = v41
        }
    case "a":
        v3 = v3
    }
    v1 -= v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v5
}
