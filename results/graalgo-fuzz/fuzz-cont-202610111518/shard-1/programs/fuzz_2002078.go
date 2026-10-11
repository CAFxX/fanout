package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32 }
type GenS2 struct { F0 bool; F1 bool }
type GenS3 struct { F0 string; F1 uint }
type GenS4 struct { F0 int32; F1 bool; F2 bool }
func genFn0(p0 uint8) uint {
    for fi5 := 0; fi5 < 6; fi5++ {
        p0 = (p0 % (((p0 >> ((255) & 63))) | 1))
        p0 += p0
        _ = fi5
    }
    var v6 uint = 1
    if false {
        v7 := []int{}
        for fi8 := 0; fi8 < 5; fi8++ {
            v6 = 746
            _ = fi8
        }
        _ = v7
    }
    var v9 uint = (v6 * (v6 / ((v6) | 1)))
    if (p0 >= p0) {
        for fi10 := 0; fi10 < 7; fi10++ {
            var v11 uint64 = 18446744073709551615
            _ = fi10
            _ = v11
        }
        var v12 bool = (true && false)
        var v13 uint64 = 0
        _ = v12
        _ = v13
    } else {
        if (2 >= 0) {
            if (255 <= 212) {
                v14 := GenS1{F0: p0, F1: 2}
                v15 := GenS1{F0: 124, F1: 4294967295}
                _ = v14
                _ = v15
            } else {
                v16 := []int{9223372036854775807, 530, 127}
                _ = v16
            }
        } else {
            var v17 int64 = -574
            switch -1 {
            case 42:
                for fi18 := 0; fi18 < 17; fi18++ {
                    if true {
                        v17 = 42
                    }
                    v19 := GenS1{F0: p0, F1: 4294967295}
                    _ = fi18
                    _ = v19
                }
            }
            p0 -= p0
            _ = v17
        }
    }
    _ = p0
    _ = v6
    _ = v9
    return v9
}

func main() {
    var v1 float32 = 1.5
    var v2 int16 = 32767
    var v3 int8 = 42
    var v4 bool = true
    for fi20 := 0; fi20 < 11; fi20++ {
        if (!false) {
            switch ("hello" + "foo bar") {
            case "\n\t":
                var v21 uint64 = 42
                _ = v21
            case "hello":
                var v22 int32 = -2147483648
                _ = v22
            default:
                if v4 {
                    if false {
                        v23 := []int{2147483647, fi20, (fi20 ^ 128)}
                        v24 := GenS2{F0: false, F1: false}
                        _ = v23
                        _ = v24
                    }
                    var v25 uint32 = 0
                    _ = v25
                } else {
                    for fi26 := 0; fi26 < 3; fi26++ {
                        fmt.Println(v1, v1, v3)
                        if (127 > 255) {
                            v27 := []int{fi20, fi26}
                            _ = v27
                        }
                        v4 = v4
                        _ = fi26
                    }
                    v28 := GenS3{F0: ("\n\t" + ("foo bar" + "hello")), F1: 4294967295}
                    fmt.Println(v4)
                    _ = v28
                }
            }
            for fi29 := 0; fi29 < 8; fi29++ {
                v3 = -40
                v4 = (((((v1 + v1) - v1) * v1) + v1) >= ((float32(0.0) * v1) / (1.5 / (((v1 - (v1 + v1)) * 1e10) * v1))))
                v2 -= v2
                _ = fi29
            }
            var v30 int16 = v2
            _ = v30
        }
        _ = fi20
    }
    var v31 int64 = 228
    switch 2147483647 {
    case -2147483648:
        for fi32 := 0; fi32 < 19; fi32++ {
            var v33 uint32 = 2
            v34 := []int{fi32, (fi32 ^ (fi32 * 9223372036854775807))}
            if len(v34) >= 2 {
                v34 = v34[1:len(v34)]
            }
            _ = fi32
            _ = v33
            _ = v34
        }
    case 32768:
        v35 := GenS3{F0: "a", F1: 65535}
        _ = v35
    }
    v36 := GenS1{F0: 2, F1: 2147483648}
    var v37 int32 = 2147483647
    if true {
        var v38 uint32 = 387
        v2 += v2
        _ = v38
    }
    var v39 int = 2147483647
    v2 = (v2 * (v2 | v2))
    var v40 int64 = (v31 & 9223372036854775807)
    fmt.Println(v39, v3)
    v31 = (^(v40))
    v41 := []int{128}
    for fi42 := 0; fi42 < 11; fi42++ {
        if (v1 <= -2.25) {
            v31 -= v40
            v43 := GenS4{F0: v37, F1: ((v39 | (v39 * fi42)) < (fi42 | 4294967295)), F2: (!(568 == 255))}
            _ = v43
        } else {
            var v44 uint32 = 73
            _ = v44
        }
        for fi45 := 0; fi45 < 15; fi45++ {
            v40 = v31
            v46 := GenS1{F0: 103, F1: 42}
            _ = fi45
            _ = v46
        }
        fmt.Println(v3, v4)
        _ = fi42
    }
    v40 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v37:", v37)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("slicelen v41:", len(v41), v41)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v36
    _ = v37
    _ = v39
    _ = v4
    _ = v40
    _ = v41
}
