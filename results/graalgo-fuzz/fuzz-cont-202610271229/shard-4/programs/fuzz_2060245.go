package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int16; F1 int8; F2 int32 }
type GenS3 struct { F0 uint; F1 int }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 string; F1 int8 }
type GenS6 struct { F0 uint; F1 int16 }
func genFn0(p0 uint8) uint {
    p0 -= (p0 + p0)
    v6 := GenS1{F0: 571}
    for fi7 := 0; fi7 < 5; fi7++ {
        for fi8 := 0; fi8 < 5; fi8++ {
            var v9 float64 = 1.5
            switch "" {
            case "":
                v10 := []int{fi7, fi7, -861}
                _ = v10
            }
            switch "\n\t" {
            case "hello":
                if (true && (1.5 < 1e10)) {
                    p0 -= 0
                    v11 := []int{(fi8 | fi7)}
                    switch "" {
                    case "hello":
                        var v12 uint32 = 2147483648
                        _ = v12
                    case "":
                        v13 := GenS2{F0: -466, F1: 1, F2: 243}
                        _ = v13
                    }
                    _ = v11
                } else {
                    var v14 uint64 = 42
                    v15 := []int{}
                    _ = v14
                    _ = v15
                }
                switch 32767 {
                case -9223372036854775808:
                    p0 = (^(p0))
                default:
                    v16 := []int{fi7, fi7, (fi7 | (fi8 + 9223372036854775807)), (fi8 + fi8)}
                    _ = v16
                }
            case "":
                if true {
                    var v17 float64 = v9
                    v18 := []int{fi8, fi8, (fi8 * fi8), -9223372036854775808}
                    _ = v17
                    _ = v18
                } else {
                    v19 := []int{-746, fi8}
                    if false {
                        fmt.Println(fi7, v9)
                    }
                    _ = len(v19)
                    _ = v19
                }
            case "foo bar":
                var v20 string = "hello"
                _ = v20
            default:
                v9 -= ((v9 + (1.5 / (v9 + (v9 / 1.5)))) * v9)
            }
            _ = fi8
            _ = v9
        }
        for fi21 := 0; fi21 < 14; fi21++ {
            v22 := GenS3{F0: 255, F1: fi7}
            if false {
                p0 = p0
                p0 -= p0
                var v23 float64 = -2.25
                _ = v23
            }
            _ = fi21
            _ = v22
        }
        _ = fi7
    }
    v24 := []int{32767, -9223372036854775808, 2147483647, -9223372036854775808}
    _ = p0
    _ = v24
    _ = v6
    return 255
}
func genFn1(p0 string) uint64 {
    v25 := []int{-520, 1, -865, 461}
    _ = len(v25)
    for fi26 := 0; fi26 < 8; fi26++ {
        p0 = p0
        if ((42 >= 100) || false) {
            fmt.Println(fi26, p0)
            switch -9223372036854775808 {
            case 127:
                p0 = "\n\t"
            case 0:
                v27 := GenS4{F0: 621}
                _ = v27
            case -2147483648:
                _ = len(v25)
            }
        }
        _ = fi26
    }
    _ = p0
    _ = v25
    return 747
}

func main() {
    var v1 int32 = 0
    v2 := []int{32768}
    v3 := make(map[uint32]bool)
    var v4 bool = (4294967295 != 4294967295)
    var v5 int16 = 32767
    var v28 uint64 = genFn1("a")
    fmt.Println(v28, v1, v28)
    fmt.Println(v5, v28)
    if v4 {
        var v29 int32 = (^(v1))
        _ = len(v2)
        _ = v29
    } else {
        _ = len(v2)
        fmt.Println(v28, v5, v1)
    }
    if false {
        v30 := GenS2{F0: v5, F1: 19, F2: v1}
        _ = v30
    } else {
        v1 -= (-(v1))
        switch "" {
        case "foo bar":
            if v4 {
                var v31 uint32 = 65535
                var v32 uint32 = 415
                _ = v31
                _ = v32
            } else {
                v28 = genFn1("\n\t")
            }
        default:
            fmt.Println(v1, v28, v4)
        }
    }
    fmt.Println(v4)
    if (v4 || (v28 != 0)) {
        v1 -= v1
        var v33 uint8 = 129
        fmt.Println(v5, v28)
        _ = v33
    }
    for fi34 := 0; fi34 < 19; fi34++ {
        if (-295 == 582) {
            v35 := GenS3{F0: 96, F1: fi34}
            for fi36 := 0; fi36 < 12; fi36++ {
                if ((v4 && ((fi34 * fi36) >= fi34)) && v4) {
                    v3[950] = v4
                    v37 := GenS4{F0: 1}
                    _ = v37
                } else {
                    v5 = (v5 * (v5 / ((1) | 1)))
                    v1 = (-(v1))
                    var v38 uint64 = v28
                    _ = v38
                }
                v4 = true
                fmt.Println(v4, fi36, fi34)
                _ = fi36
            }
            v39 := GenS4{F0: 321}
            _ = v35
            _ = v39
        }
        if (!v4) {
            switch "hello" {
            case "hello":
                v40 := GenS5{F0: "\n\t", F1: -1}
                _ = v40
            case "foo bar":
                var v41 uint64 = 4294967295
                _ = v41
            case "":
                v42 := GenS6{F0: 255, F1: (-(v5))}
                _ = v42
            }
        }
        _ = fi34
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
