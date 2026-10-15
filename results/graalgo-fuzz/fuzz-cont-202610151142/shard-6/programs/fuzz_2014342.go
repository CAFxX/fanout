package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32 }
type GenS2 struct { F0 uint16; F1 uint32; F2 string }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint64; F1 string; F2 uint32 }
func genFn0(p0 uint32, p1 string) uint {
    for fi4 := 0; fi4 < 6; fi4++ {
        v5 := []int{0, fi4}
        var v6 float32 = (1e10 + (1.5 + ((1.5 - (0.0 / ((((0.1 + 3.14159) / 1e10) / 0.1) / (0.1 / 0.1)))) + 1.5)))
        var v7 float32 = 0.1
        _ = fi4
        _ = v5
        _ = v6
        _ = v7
    }
    switch 480 {
    case 127:
        for fi8 := 0; fi8 < 2; fi8++ {
            p0 = p0
            p1 = p1
            if (11 <= 1) {
                fmt.Println(p1, p0)
            } else {
                var v9 uint8 = 249
                _ = v9
            }
            _ = fi8
        }
    case 1:
        var v10 int = 128
        _ = v10
    }
    v11 := GenS1{F0: -32768, F1: p0}
    switch p1 {
    case "\n\t":
        var v12 int8 = 0
        _ = v12
    case "":
        v13 := []int{42}
        _ = v13
    case "a":
        var v14 float64 = (-2.25 + 0.1)
        _ = v14
    }
    v15 := []int{32768, 127, -703, -2147483648}
    _ = p0
    _ = p1
    _ = v11
    _ = v15
    return 42
}
func genFn1(p0 int64, p1 uint64, p2 int32) int {
    var v16 float64 = -2.25
    p1 = (p1 / ((263) | 1))
    var v17 int64 = p0
    if true {
        v18 := []int{-1, -9223372036854775808, -1, -9223372036854775808}
        switch "" {
        case "hello":
            v19 := GenS2{F0: 65535, F1: 2, F2: "foo bar"}
            _ = v19
        case "foo bar":
            var v20 uint8 = 217
            _ = v20
        case "":
            var v21 float64 = (v16 - 0.1)
            _ = v21
        default:
            v18 = append(v18, 2147483647)
        }
        if len(v18) >= 2 {
            v18 = v18[1:len(v18)]
        }
        _ = v18
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    _ = v17
    return -149
}

func main() {
    var v1 int16 = 927
    var v2 float64 = ((-2.25 * 1e10) - -2.25)
    var v3 uint = 5
    for fi22 := 0; fi22 < 17; fi22++ {
        if false {
            switch "" {
            case "foo bar":
                if (true || ((1 < 42) || ((!(9223372036854775807 != -9223372036854775808)) && (2 >= 731)))) {
                    if (1.5 <= 0.0) {
                        v1 -= v1
                        var v23 int32 = 0
                        _ = v23
                    }
                    var v24 uint = (^(v3))
                    _ = v24
                } else {
                    var v25 float32 = ((0.0 - (1.5 / -2.25)) / 1.5)
                    _ = v25
                }
            case "hello":
                v3 -= 4294967295
            case "":
                fmt.Println(v3)
            default:
                var v26 int8 = 1
                _ = v26
            }
            if false {
                for fi27 := 0; fi27 < 11; fi27++ {
                    v3 = genFn0(0, "hello")
                    v28 := []int{fi22, (fi22 >> ((v3) & 63))}
                    _ = fi27
                    _ = v28
                }
                var v29 int64 = -9223372036854775808
                var v30 float64 = ((v2 + (v2 - v2)) + v2)
                _ = v29
                _ = v30
            }
        }
        var v31 uint16 = 100
        _ = fi22
        _ = v31
    }
    switch 51 {
    case -1:
        v32 := GenS3{F0: 18446744073709551615}
        _ = v32
    case 32767:
        if true {
            v2 = 0.1
        }
    default:
        var v33 uint = v3
        _ = v33
    }
    v2 -= float64(1e10)
    v34 := []int{37}
    if true {
        if true {
            for fi35 := 0; fi35 < 19; fi35++ {
                v36 := GenS4{F0: 4294967295, F1: "foo bar", F2: 1}
                v2 += v2
                _ = fi35
                _ = v36
            }
            v34 = append(v34, 2147483647)
        } else {
            v3 = v3
        }
    }
    var v37 bool = (2147483647 >= 750)
    fmt.Println(v2, v1, v2)
    v2 += (1e10 * (1.5 + 1.5))
    if len(v34) > 0 {
        v34[0] = 4294967295
    }
    fmt.Println(v3, v3, v1)
    if (!true) {
        switch ((("hello" + "a") + "a") + "a") {
        case "hello":
            for fi38 := 0; fi38 < 11; fi38++ {
                var v39 int16 = (^(v1))
                _ = fi38
                _ = v39
            }
        }
        var v40 string = (("\n\t" + "") + ("hello" + "foo bar"))
        fmt.Println(v3)
        _ = v40
    }
    var v41 int8 = 127
    var v42 uint64 = 0
    v42 -= v42
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("v37:", v37)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v37
    _ = v41
    _ = v42
}
