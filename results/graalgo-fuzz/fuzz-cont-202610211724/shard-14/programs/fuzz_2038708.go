package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int; F2 uint8 }
type GenS2 struct { F0 bool; F1 int64; F2 int32 }
type GenS3 struct { F0 string; F1 bool }
type GenS4 struct { F0 int8; F1 int8 }
type GenS5 struct { F0 int; F1 int32 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 uint8; F1 uint16 }
func genFn0() int16 {
    v4 := []int{}
    var v5 uint16 = 449
    var v6 uint8 = 0
    var v7 uint8 = 1
    v5 = 65535
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return -32768
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[string]int)
    var v3 int16 = 265
    var v8 float32 = 3.14159
    for fi9 := 0; fi9 < 1; fi9++ {
        fmt.Println(v3, v3)
        var v10 uint = 18446744073709551615
        fmt.Println(v8)
        _ = fi9
        _ = v10
    }
    for fi11 := 0; fi11 < 8; fi11++ {
        switch ("foo bar" + "hello") {
        case "hello":
            v3 += 32767
        default:
            _ = len(v2)
        }
        for fi12 := 0; fi12 < 16; fi12++ {
            var v13 string = ("a" + "\n\t")
            v8 -= (v8 + float32((0.1 / 1e10)))
            _ = fi12
            _ = v13
        }
        _ = fi11
    }
    var v14 float32 = -2.25
    v8 = 3.14159
    switch -693 {
    case 9223372036854775807:
        v15 := GenS1{F0: 1, F1: 42, F2: 128}
        _ = v15
    }
    switch 42 {
    case -1:
        for fi16 := 0; fi16 < 8; fi16++ {
            v8 -= v8
            for fi17 := 0; fi17 < 16; fi17++ {
                v18 := GenS2{F0: false, F1: 1, F2: 808}
                fmt.Println(v8, fi17, v3)
                _ = fi17
                _ = v18
            }
            v19 := v2["hello"]
            _ = fi16
            _ = v19
        }
    case -55:
        for fi20 := 0; fi20 < 6; fi20++ {
            if false {
                for fi21 := 0; fi21 < 9; fi21++ {
                    var v22 uint = 0
                    v22 -= (v22 / (((v22 & (^(v22)))) | 1))
                    v22 += (v22 - v22)
                    var v23 uint = v22
                    _ = fi21
                    _ = v22
                    _ = v23
                }
                fmt.Println(v3)
                if (!true) {
                    v3 = v3
                }
            } else {
                var v24 float64 = ((1e10 * 3.14159) + float64(v14))
                fmt.Println(v8, fi20, v14)
                _ = v24
            }
            switch "\n\t" {
            case "\n\t":
                v25 := GenS3{F0: "a", F1: false}
                _ = v25
            case "a":
                fmt.Println(v3, v8)
            case "hello":
                for fi26 := 0; fi26 < 14; fi26++ {
                    fmt.Println(fi26)
                    if ((false || ((1.5 / 1.0) != (3.14159 * (0.0 * 1e10)))) || true) {
                        var v27 int = (fi20 / ((fi26) | 1))
                        v28 := GenS4{F0: 68, F1: 0}
                        v3 += v3
                        _ = v27
                        _ = v28
                    } else {
                        if (true || (v3 > v3)) {
                            v8 = v8
                            v29 := GenS1{F0: 42, F1: fi20, F2: 225}
                            v14 = v8
                            _ = v29
                        } else {
                            v30 := GenS5{F0: fi26, F1: 42}
                            _ = v30
                        }
                        v3 -= (v3 | (v3 + 1000))
                        fmt.Println(v3, v8, fi26)
                        v31 := GenS6{F0: 2}
                        if (!(v3 != int16(fi26))) {
                            fmt.Println(fi26)
                            v32 := GenS7{F0: 128, F1: 141}
                            _ = v32
                        }
                        _ = v31
                    }
                    _ = fi26
                }
            default:
                var v33 float64 = 3.14159
                _ = v33
            }
            _ = fi20
        }
    case 374:
        for fi34 := 0; fi34 < 8; fi34++ {
            delete(v1, 2147483648)
            fmt.Println(v14)
            fmt.Println(v14)
            _ = fi34
        }
    }
    if ((1e10 + (0.1 + v14)) >= float32(3.14159)) {
        v35 := []int{32767, -9223372036854775808, 42, 528}
        _ = v35
    }
    v36 := []int{127}
    for fi37 := 0; fi37 < 4; fi37++ {
        if (255 >= 116) {
            for fi38 := 0; fi38 < 16; fi38++ {
                v3 = v3
                _ = fi38
            }
        } else {
            v3 = genFn0()
        }
        var v39 int32 = 896
        _ = fi37
        _ = v39
    }
    _ = len(v36)
    for fi40 := 0; fi40 < 18; fi40++ {
        v41 := GenS2{F0: true, F1: 9223372036854775807, F2: 0}
        var v42 float32 = v14
        _ = fi40
        _ = v41
        _ = v42
    }
    var v43 int16 = (v3 * (v3 % ((42) | 1)))
    _ = len(v1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v36:", len(v36), v36)
    fmt.Println("v43:", v43)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v36
    _ = v43
    _ = v8
}
