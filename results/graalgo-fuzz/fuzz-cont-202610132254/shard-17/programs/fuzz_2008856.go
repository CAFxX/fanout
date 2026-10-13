package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int; F2 uint8 }
type GenS2 struct { F0 int; F1 bool; F2 uint8 }
type GenS3 struct { F0 bool; F1 uint8 }
type GenS4 struct { F0 uint8; F1 uint32 }
type GenS5 struct { F0 uint64; F1 int32; F2 string }
type GenS6 struct { F0 uint64; F1 int64 }
func genFn0(p0 int32, p1 int) int8 {
    var v4 float32 = (0.1 + 0.0)
    for fi5 := 0; fi5 < 20; fi5++ {
        v6 := GenS1{F0: 1, F1: fi5, F2: 128}
        var v7 float64 = 0.1
        _ = fi5
        _ = v6
        _ = v7
    }
    switch "hello" {
    case "":
        v4 = v4
    case "foo bar":
        for fi8 := 0; fi8 < 5; fi8++ {
            v9 := []int{fi8, p1, -2147483648}
            if false {
                var v10 int32 = (p0 * p0)
                var v11 int = (fi8 / ((p1) | 1))
                _ = v10
                _ = v11
            }
            p0 = (p0 << ((0) & 63))
            _ = fi8
            _ = v9
        }
    default:
        fmt.Println(p0, p1)
    }
    switch 32767 {
    case 22:
        v12 := GenS1{F0: 2, F1: (p1 | (p1 ^ p1)), F2: 74}
        _ = v12
    default:
        switch p1 {
        case -124:
            if true {
                v13 := []int{p1, 937, (p1 >> ((42) & 63)), (-(p1))}
                _ = v13
            } else {
                var v14 float32 = float32(1e10)
                for fi15 := 0; fi15 < 17; fi15++ {
                    v14 += v4
                    _ = fi15
                }
                _ = v14
            }
        default:
            switch 9223372036854775807 {
            case 1:
                if (0 != 1) {
                    for fi16 := 0; fi16 < 7; fi16++ {
                        p1 -= (^(fi16))
                        v17 := GenS2{F0: (p1 >> ((242) & 63)), F1: false, F2: 40}
                        _ = fi16
                        _ = v17
                    }
                    v18 := []int{}
                    p0 += (p0 % (((p0 + (p0 | (p0 ^ (p0 % (((p0 * p0)) | 1)))))) | 1))
                    _ = v18
                } else {
                    v19 := []int{p1}
                    _ = v19
                }
            case 9223372036854775807:
                v20 := []int{}
                _ = v20
            default:
                fmt.Println(p1)
            }
        }
    }
    v4 = (float32((0.0 + 1.5)) - 1e10)
    _ = p0
    _ = p1
    _ = v4
    return 100
}

func main() {
    var v1 int32 = 42
    var v2 int32 = v1
    v3 := make(map[int]uint32)
    if (true || true) {
        var v21 float64 = float64(0.0)
        v22 := []int{-9223372036854775808}
        for fi23 := 0; fi23 < 7; fi23++ {
            v24 := GenS2{F0: fi23, F1: (2147483647 >= 0), F2: 1}
            v25 := GenS2{F0: (fi23 * 32767), F1: false, F2: 1}
            _ = fi23
            _ = v24
            _ = v25
        }
        _ = v21
        _ = v22
    } else {
        v26 := GenS3{F0: false, F1: 255}
        v1 = -2147483648
        v2 = 835
        _ = v26
    }
    for fi27 := 0; fi27 < 15; fi27++ {
        var v28 int = fi27
        fmt.Println(fi27, v28)
        _ = fi27
        _ = v28
    }
    if true {
        v2 = 0
        var v29 uint64 = 18446744073709551615
        for fi30 := 0; fi30 < 15; fi30++ {
            v31 := v3[-392]
            var v32 bool = (416 != 39)
            _ = fi30
            _ = v31
            _ = v32
        }
        _ = v29
    }
    if true {
        v33 := GenS4{F0: 42, F1: 2147483648}
        fmt.Println(v2, v1)
        v34 := GenS5{F0: 42, F1: -432, F2: ""}
        _ = v33
        _ = v34
    }
    fmt.Println(v2, v1)
    v1 = v1
    if (!(1 > 341)) {
        v2 = 201
        for fi35 := 0; fi35 < 20; fi35++ {
            v1 = (v2 + 1)
            var v36 uint8 = 127
            for fi37 := 0; fi37 < 4; fi37++ {
                if true {
                    if false {
                        v36 += (^(v36))
                        fmt.Println(v36, v1, v2)
                    } else {
                        var v38 float64 = float64(1e10)
                        _ = v38
                    }
                    switch "foo bar" {
                    case "":
                        fmt.Println(fi35, fi35)
                    case "a":
                        v39 := GenS3{F0: true, F1: (v36 * 127)}
                        _ = v39
                    }
                    v36 += v36
                }
                _ = fi37
            }
            _ = fi35
            _ = v36
        }
    }
    if (!(("a" + "hello") < "foo bar")) {
        for fi40 := 0; fi40 < 12; fi40++ {
            var v41 string = "\n\t"
            _ = fi40
            _ = v41
        }
        v42 := GenS6{F0: 0, F1: 823}
        for fi43 := 0; fi43 < 8; fi43++ {
            v2 = v1
            fmt.Println(fi43, v2, v1)
            v1 = (v1 ^ (^(v1)))
            _ = fi43
        }
        _ = v42
    } else {
        for fi44 := 0; fi44 < 2; fi44++ {
            v2 = v1
            v2 = v2
            v1 = (v1 & 2147483647)
            _ = fi44
        }
        for fi45 := 0; fi45 < 15; fi45++ {
            v1 += v1
            v46 := GenS1{F0: 18446744073709551615, F1: fi45, F2: 1}
            _ = fi45
            _ = v46
        }
    }
    var v47 uint64 = 4294967295
    v1 -= v1
    v47 = 1
    for fi48 := 0; fi48 < 18; fi48++ {
        fmt.Println(v2, fi48)
        var v49 uint = 532
        _ = fi48
        _ = v49
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v47:", v47)
    _ = v1
    _ = v2
    _ = v3
    _ = v47
}
