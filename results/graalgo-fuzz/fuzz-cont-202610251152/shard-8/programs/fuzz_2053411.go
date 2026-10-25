package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string }
type GenS2 struct { F0 int }
type GenS3 struct { F0 bool; F1 int32 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int16; F1 uint64 }
type GenS6 struct { F0 int64; F1 string }
func genFn0(p0 uint32, p1 uint32) uint32 {
    var v5 uint32 = (p1 << ((784) & 63))
    v6 := GenS1{F0: -770, F1: (("foo bar" + ("a" + "")) + "\n\t")}
    v5 = 1
    fmt.Println(p0)
    for fi7 := 0; fi7 < 8; fi7++ {
        var v8 uint8 = 128
        for fi9 := 0; fi9 < 14; fi9++ {
            switch "foo bar" {
            case "foo bar":
                var v10 uint = 0
                _ = v10
            case "hello":
                fmt.Println(fi9)
            case "":
                var v11 int8 = 27
                _ = v11
            }
            _ = fi9
        }
        _ = fi7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return p0
}
func genFn1(p0 string, p1 uint8) int {
    var v12 float64 = -2.25
    switch 32767 {
    case 128:
        if (561 >= 32767) {
            var v13 bool = true
            fmt.Println(p1, v12)
            var v14 int16 = 477
            _ = v13
            _ = v14
        }
    case 1:
        v15 := []int{-1}
        _ = v15
    case -1:
        var v16 int16 = -335
        _ = v16
    }
    v17 := []int{}
    _ = len(v17)
    _ = p0
    _ = p1
    _ = v12
    _ = v17
    return 32768
}

func main() {
    var v1 int16 = -18
    var v2 float64 = 1.5
    var v3 string = (("\n\t" + ("a" + ("a" + "foo bar"))) + "hello")
    var v4 float64 = v2
    switch v3 {
    case "\n\t":
        v18 := GenS2{F0: genFn1("", 0)}
        _ = v18
    }
    v19 := []int{genFn1(("a" + (v3 + (v3 + v3))), 2), 42, 396}
    if len(v19) > 0 {
        v19[0] = 203
    }
    for fi20 := 0; fi20 < 18; fi20++ {
        for fi21 := 0; fi21 < 13; fi21++ {
            if false {
                fmt.Println(v2, v4, fi21)
                if len(v19) > 0 {
                    v19[0] = fi20
                }
                if true {
                    v3 = v3
                    v19 = append(v19, 127)
                    switch v3 {
                    case "foo bar":
                        var v22 int32 = 42
                        _ = v22
                    default:
                        fmt.Println(v4, fi21, fi21)
                    }
                } else {
                    v23 := GenS2{F0: -993}
                    _ = v23
                }
            }
            v4 = v2
            switch (fi21 & fi20) {
            case 2147483647:
                if (42 >= 1) {
                    var v24 int32 = 1
                    fmt.Println(fi21)
                    _ = v24
                } else {
                    v25 := GenS2{F0: (^(fi20))}
                    v4 -= (v4 / -2.25)
                    _ = v25
                }
            case 32768:
                if (4294967295 >= 65535) {
                    fmt.Println(fi21, fi20)
                    v4 += (-2.25 - v2)
                }
            }
            _ = fi21
        }
        v2 += 0.1
        v26 := GenS3{F0: true, F1: 447}
        _ = fi20
        _ = v26
    }
    for fi27 := 0; fi27 < 6; fi27++ {
        if (614 >= 4294967295) {
            var v28 string = v3
            v28 = "foo bar"
            var v29 bool = (int8(fi27) < 0)
            _ = v28
            _ = v29
        } else {
            v3 = v3
            var v30 int16 = (v1 / ((v1) | 1))
            _ = v30
        }
        fmt.Println(v4, v1, v3)
        for fi31 := 0; fi31 < 4; fi31++ {
            v2 = float64((0.1 + ((0.0 + 3.14159) + (((float32(v2) + (float32(-2.25) * ((1e10 - 1.5) / float32(3.14159)))) + 1e10) / 0.1))))
            v2 = (((v4 * (float64((float32((-2.25 * v2)) / (1e10 / 1.0))) - 1e10)) + v4) / v2)
            var v32 uint32 = 955
            _ = fi31
            _ = v32
        }
        _ = fi27
    }
    var v33 int32 = -1
    for fi34 := 0; fi34 < 2; fi34++ {
        for fi35 := 0; fi35 < 17; fi35++ {
            fmt.Println(v2, v3, fi35)
            if (!((v1 + 1000) > v1)) {
                switch fi34 {
                case 127:
                    v36 := GenS4{F0: 100}
                    _ = v36
                case 32768:
                    var v37 int = fi34
                    _ = v37
                default:
                    v4 += v4
                }
                fmt.Println(v3, v1)
            } else {
                v1 -= v1
                fmt.Println(v33, v4)
                if true {
                    v1 = v1
                    v38 := GenS5{F0: (v1 * v1), F1: 4294967295}
                    _ = v38
                } else {
                    if true {
                        v39 := GenS2{F0: fi34}
                        fmt.Println(v4, fi35)
                        _ = v39
                    } else {
                        v33 = 0
                        v40 := GenS6{F0: 1, F1: (v3 + v3)}
                        _ = v40
                    }
                }
                v4 = 1e10
                v33 = v33
                fmt.Println(v4, v4)
                v3 = v3
                v33 = -598
                var v41 int32 = -2147483648
                _ = v41
            }
            _ = fi35
        }
        v3 = "\n\t"
        _ = fi34
    }
    _ = len(v19)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v33
    _ = v4
}
