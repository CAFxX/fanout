package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32; F2 int }
type GenS2 struct { F0 int16; F1 int }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint32; F1 uint; F2 uint64 }
type GenS5 struct { F0 bool; F1 int64 }
type GenS6 struct { F0 string; F1 uint64; F2 uint16 }
func genFn0(p0 bool, p1 int32, p2 int) int64 {
    if (true && (p0 && true)) {
        for fi4 := 0; fi4 < 16; fi4++ {
            p0 = p0
            p1 += 42
            v5 := []int{}
            _ = fi4
            _ = v5
        }
        for fi6 := 0; fi6 < 8; fi6++ {
            var v7 int32 = -419
            _ = fi6
            _ = v7
        }
        switch "\n\t" {
        case "foo bar":
            v8 := GenS1{F0: p1, F1: (p1 << ((298) & 63)), F2: p2}
            _ = v8
        case "a":
            v9 := GenS2{F0: -1, F1: 32768}
            _ = v9
        case "hello":
            p0 = p0
        default:
            var v10 int8 = 1
            _ = v10
        }
    }
    var v11 float64 = (float64(-2.25) - (1e10 / 0.1))
    v12 := GenS3{F0: 100}
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    return 9223372036854775807
}

func main() {
    v1 := make(map[uint32]string)
    var v2 float64 = 1e10
    var v3 int16 = 1000
    v2 -= 3.14159
    v3 += (v3 * v3)
    if (!false) {
        v13 := []int{32768, 9223372036854775807, 2147483647, 952}
        _ = v13
    }
    var v14 int = 32768
    for fi15 := 0; fi15 < 14; fi15++ {
        var v16 uint32 = 4294967295
        _ = fi15
        _ = v16
    }
    switch v14 {
    case -660:
        v3 -= -1
    case 128:
        switch "a" {
        case "hello":
            if ((0 < 4294967295) && (v2 < (v2 * v2))) {
                v2 -= -2.25
                var v17 uint64 = 0
                _ = v17
            } else {
                delete(v1, 1)
            }
        default:
            v18 := GenS4{F0: 671, F1: 65535, F2: 2}
            _ = v18
        }
    }
    v19 := []int{127, v14, (v14 & v14), v14}
    for fi20 := 0; fi20 < 3; fi20++ {
        v21 := GenS4{F0: 4294967295, F1: 4294967295, F2: 4294967295}
        fmt.Println(fi20, v3, v14)
        _ = fi20
        _ = v21
    }
    v22 := GenS2{F0: v3, F1: v14}
    for fi23 := 0; fi23 < 20; fi23++ {
        fmt.Println(v3, fi23, v14)
        fmt.Println(v3)
        for fi24 := 0; fi24 < 18; fi24++ {
            if true {
                v25 := GenS5{F0: false, F1: 0}
                _ = v25
            } else {
                var v26 float32 = -2.25
                if false {
                    fmt.Println(v3, v26)
                    switch "" {
                    case "\n\t":
                        switch ("" + "\n\t") {
                        case "hello":
                            var v27 string = ""
                            _ = v27
                        case "\n\t":
                            if false {
                                var v28 bool = (218 == 42)
                                _ = v28
                            } else {
                                switch ("hello" + "") {
                                case "hello":
                                    var v29 int = (-(v14))
                                    _ = v29
                                }
                                v14 = 553
                            }
                        }
                    case "foo bar":
                        fmt.Println(v3, v26)
                    }
                    v14 -= 4294967295
                }
                var v30 float32 = 1.5
                _ = v26
                _ = v30
            }
            v14 = v14
            _ = fi24
        }
        _ = fi23
    }
    v14 += 9223372036854775807
    if ((v14 << ((1) & 63)) < 127) {
        v3 = 517
        v14 -= (-(v14))
        fmt.Println(v2, v3)
    }
    if (!(47 > 127)) {
        switch "hello" {
        case "a":
            if len(v19) > 0 {
                v19[0] = (v14 ^ v14)
            }
        default:
            v19 = append(v19, (^(v14)))
        }
        var v31 string = (("" + "hello") + ((("\n\t" + "a") + "") + "\n\t"))
        switch v14 {
        case -9223372036854775808:
            v32 := GenS3{F0: 58}
            _ = v32
        case 4294967295:
            switch 2147483647 {
            case 32767:
                if true {
                    fmt.Println(v14)
                    v31 = "foo bar"
                }
            case -98:
                v14 = (v14 / (((v14 << ((42) & 63))) | 1))
            default:
                fmt.Println(v2, v3)
            }
        case -402:
            delete(v1, 42)
        }
        _ = v31
    } else {
        if true {
            if (true && false) {
                v3 = v3
            } else {
                v33 := GenS6{F0: "foo bar", F1: 183, F2: 458}
                v34 := GenS3{F0: int8(v14)}
                fmt.Println(v14, v3)
                _ = v33
                _ = v34
            }
            var v35 int64 = int64(v14)
            _ = v35
        }
    }
    for fi36 := 0; fi36 < 19; fi36++ {
        v2 += v2
        if false {
            fmt.Println(v2, fi36, v3)
            v37 := v1[2]
            v14 -= v14
            _ = v37
        }
        _ = fi36
    }
    v38 := v1[2147483647]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    _ = v1
    _ = v14
    _ = v19
    _ = v2
    _ = v22
    _ = v3
    _ = v38
}
