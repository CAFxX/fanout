package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint32; F1 uint }
type GenS4 struct { F0 bool; F1 uint; F2 int16 }
type GenS5 struct { F0 uint8; F1 uint16; F2 uint }
type GenS6 struct { F0 int32; F1 int64; F2 int64 }
type GenS7 struct { F0 int }
func genFn0() int64 {
    v4 := GenS1{F0: 18446744073709551615, F1: "hello"}
    if true {
        var v5 uint64 = 42
        _ = v5
    }
    for fi6 := 0; fi6 < 8; fi6++ {
        v7 := GenS2{F0: 6}
        for fi8 := 0; fi8 < 18; fi8++ {
            var v9 bool = (!(245 > 42))
            switch "foo bar" {
            case "foo bar":
                v9 = v9
            case "\n\t":
                v10 := []int{}
                _ = v10
            case "a":
                if (0 == 0) {
                    v9 = false
                    var v11 float64 = 1e10
                    var v12 int16 = 32767
                    _ = v11
                    _ = v12
                }
            default:
                if ((fi8 >> ((42) & 63)) <= 32768) {
                    v13 := GenS3{F0: 4294967295, F1: 65535}
                    _ = v13
                } else {
                    if true {
                        v9 = ("foo bar" == "foo bar")
                    }
                    v14 := GenS4{F0: v9, F1: 65535, F2: 488}
                    var v15 string = "hello"
                    _ = v14
                    _ = v15
                }
            }
            v9 = false
            v16 := []int{}
            v16 = append(v16, (fi6 + 0))
            _ = fi8
            _ = v16
            _ = v9
        }
        _ = fi6
        _ = v7
    }
    var v17 bool = true
    v17 = true
    _ = v17
    _ = v4
    return 9223372036854775807
}
func genFn1(p0 bool, p1 int32) int {
    var v18 float64 = -2.25
    v19 := []int{1, -727, int(p1), 4294967295}
    p0 = (32767 > -9223372036854775808)
    if len(v19) >= 2 {
        v19 = v19[1:len(v19)]
    }
    _ = p0
    _ = p1
    _ = v18
    _ = v19
    return 4294967295
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[int]bool)
    var v3 bool = false
    v3 = v3
    var v20 float32 = (0.0 * -2.25)
    v1[127] = "foo bar"
    v20 += v20
    v20 = float32((((1e10 / 1e10) - -2.25) + float64(v20)))
    fmt.Println(v20)
    v3 = v3
    if v3 {
        v21 := []int{}
        _ = v21
    }
    if true {
        v3 = (!(42 == 4294967295))
    }
    delete(v2, -9223372036854775808)
    v3 = v3
    switch 4294967295 {
    case -842:
        v22 := []int{0, -1}
        _ = v22
    case 462:
        v23 := []int{2147483647, -9223372036854775808, -493}
        _ = v23
    default:
        var v24 uint32 = 883
        _ = v24
    }
    switch "" {
    case "\n\t":
        if (("hello" + ("a" + "a")) >= "hello") {
            if (42 != 42) {
                switch 0 {
                case 127:
                    v20 -= 1e10
                default:
                    for fi25 := 0; fi25 < 2; fi25++ {
                        switch fi25 {
                        case 724:
                            v26 := GenS3{F0: 110, F1: 0}
                            _ = v26
                        case -410:
                            v3 = (!v3)
                        case 127:
                            var v27 float32 = 3.14159
                            _ = v27
                        default:
                            var v28 int = 127
                            _ = v28
                        }
                        v29 := GenS5{F0: 128, F1: 1, F2: 4294967295}
                        v30 := []int{fi25}
                        _ = fi25
                        _ = v29
                        _ = v30
                    }
                }
                fmt.Println(v20)
            } else {
                var v31 int64 = -9223372036854775808
                _ = v31
            }
            var v32 uint16 = 255
            v32 = (^(v32))
            _ = v32
        } else {
            v33 := GenS6{F0: 927, F1: -1, F2: genFn0()}
            var v34 int64 = 9223372036854775807
            fmt.Println(v3, v20)
            _ = v33
            _ = v34
        }
    case "":
        var v35 uint32 = 2
        _ = v35
    case "foo bar":
        fmt.Println(v3)
    default:
        for fi36 := 0; fi36 < 14; fi36++ {
            v37 := GenS2{F0: 883}
            fmt.Println(v20)
            var v38 int8 = 0
            _ = fi36
            _ = v37
            _ = v38
        }
    }
    var v39 float64 = 0.1
    for fi40 := 0; fi40 < 8; fi40++ {
        v41 := GenS7{F0: (fi40 | 320)}
        _ = fi40
        _ = v41
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v39
}
