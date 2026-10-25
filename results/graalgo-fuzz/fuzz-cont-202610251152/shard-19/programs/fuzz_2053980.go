package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16 }
type GenS2 struct { F0 string; F1 uint8; F2 int64 }
type GenS3 struct { F0 uint8; F1 int16; F2 int }
type GenS4 struct { F0 int64; F1 int64 }
func genFn0(p0 uint64) int8 {
    if (!(!true)) {
        p0 = p0
        v5 := GenS1{F0: -2147483648, F1: 410}
        var v6 uint32 = 2147483648
        _ = v5
        _ = v6
    }
    var v7 uint64 = 0
    if (42 >= 94) {
        switch -87 {
        case -2147483648:
            v7 = v7
        case -1:
            v8 := GenS2{F0: "foo bar", F1: 128, F2: 721}
            _ = v8
        }
    }
    for fi9 := 0; fi9 < 4; fi9++ {
        if (65535 == 0) {
            if false {
                var v10 int = (fi9 | fi9)
                for fi11 := 0; fi11 < 10; fi11++ {
                    v10 = v10
                    _ = fi11
                }
                var v12 uint8 = 143
                _ = v10
                _ = v12
            } else {
                v13 := []int{fi9, (fi9 + 9223372036854775807)}
                fmt.Println(p0)
                v7 = v7
                _ = v13
            }
        } else {
            v7 -= p0
            p0 -= (^(p0))
            var v14 bool = (false && true)
            _ = v14
        }
        v15 := []int{(fi9 / ((fi9) | 1)), (fi9 / ((fi9) | 1)), int(v7), (-(fi9))}
        p0 = p0
        _ = fi9
        _ = v15
    }
    _ = p0
    _ = v7
    return 100
}
func genFn1(p0 int8, p1 int32) int {
    var v16 float64 = float64(-2.25)
    v17 := []int{-1, 4294967295, -810, -1}
    v16 = (1.5 / (v16 / (v16 * (v16 + float64(-2.25)))))
    if (!(((0 < 139) || true) && (338 <= 0))) {
        for fi18 := 0; fi18 < 8; fi18++ {
            for fi19 := 0; fi19 < 4; fi19++ {
                if false {
                    _ = len(v17)
                }
                v20 := GenS3{F0: 148, F1: 42, F2: 128}
                v16 += (0.0 - v16)
                _ = fi19
                _ = v20
            }
            _ = fi18
        }
        if len(v17) >= 2 {
            v17 = v17[1:len(v17)]
        }
    } else {
        fmt.Println(p1)
    }
    _ = p0
    _ = p1
    _ = v16
    _ = v17
    return -1
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[string]uint32)
    var v3 int = 9223372036854775807
    var v4 bool = (!false)
    v3 = (v3 ^ (v3 % ((genFn1(genFn0(392), -1)) | 1)))
    var v21 bool = (1000 == 1)
    v3 += v3
    var v22 int16 = -1
    var v23 uint64 = 664
    var v24 uint = 4294967295
    var v25 uint64 = v23
    v26 := GenS1{F0: v3, F1: 675}
    v27 := []int{v3}
    var v28 float32 = (((3.14159 / 1.0) * 1.5) / (1e10 - -2.25))
    v29 := GenS2{F0: ("\n\t" + "a"), F1: 128, F2: -465}
    fmt.Println(v28, v24)
    v28 -= (v28 - ((v28 - float32(0.1)) * (float32((1.5 / 3.14159)) + (0.1 - -2.25))))
    if false {
        v27 = append(v27, (v3 | 630))
    } else {
        var v30 string = ""
        for fi31 := 0; fi31 < 13; fi31++ {
            var v32 uint8 = 42
            var v33 float64 = ((float64((((v28 + v28) + v28) - -2.25)) + 0.0) + (3.14159 + -2.25))
            for fi34 := 0; fi34 < 4; fi34++ {
                v30 = ""
                _ = fi34
            }
            _ = fi31
            _ = v32
            _ = v33
        }
        _ = v30
    }
    for fi35 := 0; fi35 < 10; fi35++ {
        var v36 string = "foo bar"
        v37 := GenS4{F0: 0, F1: -365}
        v27 = append(v27, (^(v3)))
        _ = fi35
        _ = v36
        _ = v37
    }
    for fi38 := 0; fi38 < 9; fi38++ {
        if (!v21) {
            for fi39 := 0; fi39 < 1; fi39++ {
                if len(v27) >= 2 {
                    v27 = v27[1:len(v27)]
                }
                var v40 float64 = (0.1 / 3.14159)
                _ = fi39
                _ = v40
            }
            var v41 int64 = -1
            _ = v41
        } else {
            var v42 bool = true
            switch -9223372036854775808 {
            case -969:
                var v43 string = "a"
                _ = v43
            case 0:
                if ((1 >= -9223372036854775808) || (!v21)) {
                    v42 = v42
                    var v44 uint = v24
                    _ = v44
                } else {
                    var v45 int32 = -390
                    switch (^(fi38)) {
                    case 32767:
                        var v46 float32 = (v28 * float32((0.0 / -2.25)))
                        _ = v46
                    default:
                        var v47 int16 = (v22 ^ v22)
                        _ = v47
                    }
                    _ = v45
                }
            }
            for fi48 := 0; fi48 < 4; fi48++ {
                var v49 string = (("hello" + "foo bar") + ("a" + "foo bar"))
                _ = fi48
                _ = v49
            }
            _ = v42
        }
        _ = fi38
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
}
