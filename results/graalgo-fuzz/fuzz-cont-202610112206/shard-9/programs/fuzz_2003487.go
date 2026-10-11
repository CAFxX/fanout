package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16 }
type GenS2 struct { F0 int64; F1 uint8 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint16; F1 uint64 }
type GenS5 struct { F0 uint8; F1 int16 }
type GenS6 struct { F0 uint64 }
func genFn0(p0 bool, p1 uint64) int8 {
    for fi7 := 0; fi7 < 5; fi7++ {
        v8 := GenS1{F0: -1, F1: 1000}
        _ = fi7
        _ = v8
    }
    var v9 uint16 = 225
    v9 += v9
    var v10 bool = p0
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return 100
}
func genFn1(p0 uint16, p1 uint) int8 {
    if true {
        p0 = (p0 * (^(p0)))
        for fi11 := 0; fi11 < 11; fi11++ {
            for fi12 := 0; fi12 < 1; fi12++ {
                var v13 int32 = 42
                var v14 float32 = 0.1
                _ = fi12
                _ = v13
                _ = v14
            }
            v15 := []int{(-(fi11)), (fi11 % ((fi11) | 1)), fi11}
            _ = fi11
            _ = v15
        }
        if ((0.1 == 1e10) && true) {
            p1 = p1
            fmt.Println(p1, p0)
            v16 := GenS2{F0: 9223372036854775807, F1: 238}
            _ = v16
        } else {
            for fi17 := 0; fi17 < 17; fi17++ {
                switch "" {
                case "":
                    v18 := GenS2{F0: 0, F1: 1}
                    _ = v18
                default:
                    v19 := []int{fi17}
                    _ = v19
                }
                var v20 float64 = 0.1
                _ = fi17
                _ = v20
            }
            if (75 < 1) {
                v21 := GenS1{F0: 137, F1: 1}
                if false {
                    v22 := []int{9223372036854775807}
                    if false {
                        p1 -= p1
                        fmt.Println(p0, p0)
                        v23 := GenS1{F0: 0, F1: 870}
                        _ = v23
                    }
                    _ = v22
                } else {
                    p0 -= p0
                    var v24 int = 128
                    _ = v24
                }
                v25 := GenS3{F0: 4294967295}
                _ = v21
                _ = v25
            } else {
                var v26 int16 = 32767
                _ = v26
            }
        }
    }
    v27 := GenS1{F0: 2147483647, F1: 0}
    if true {
        v28 := []int{-1, 32768}
        _ = v28
    } else {
        v29 := []int{4294967295, 127}
        v30 := GenS4{F0: 235, F1: 18446744073709551615}
        _ = v29
        _ = v30
    }
    v31 := GenS5{F0: 128, F1: -1000}
    _ = p0
    _ = p1
    _ = v27
    _ = v31
    return -1
}

func main() {
    v1 := make(map[string]int)
    var v2 uint = 0
    var v3 float64 = (-2.25 + float64(((float32((0.0 * 0.0)) - (float32(1.5) * ((0.1 / ((0.0 - 3.14159) - 0.0)) - 1.5))) / 1e10)))
    var v4 uint8 = 0
    var v5 string = (("a" + "") + ("hello" + "\n\t"))
    v6 := make(map[int]uint32)
    for fi32 := 0; fi32 < 11; fi32++ {
        v5 = v5
        var v33 string = ((v5 + ("foo bar" + v5)) + v5)
        switch fi32 {
        case 51:
            fmt.Println(v3)
        case -1:
            v5 = v5
        }
        _ = fi32
        _ = v33
    }
    fmt.Println(v4, v3, v4)
    v6[32768] = 2147483648
    var v34 uint = (v2 >> ((v2) & 63))
    v35 := GenS6{F0: 2}
    if true {
        fmt.Println(v2, v34)
        if false {
            for fi36 := 0; fi36 < 12; fi36++ {
                var v37 string = "hello"
                fmt.Println(v2)
                _ = fi36
                _ = v37
            }
        }
    } else {
        v38 := v1[("foo bar" + (v5 + "a"))]
        v39 := GenS4{F0: 621, F1: 1}
        v5 = v5
        _ = v38
        _ = v39
    }
    v40 := []int{2147483647, -1}
    v5 = "\n\t"
    switch -614 {
    case 0:
        switch (v5 + "") {
        case "":
            v2 += v34
        default:
            switch v5 {
            case "foo bar":
                fmt.Println(v5, v3, v3)
            case "hello":
                v41 := GenS1{F0: 917, F1: 1}
                _ = v41
            case "":
                _ = len(v40)
            default:
                if (-9223372036854775808 >= 0) {
                    if len(v40) > 0 {
                        v40[0] = 4294967295
                    }
                    v34 = (v2 + (v34 - (v2 >> (((v2 >> ((1) & 63))) & 63))))
                }
            }
        }
    }
    v42 := GenS1{F0: 42, F1: 370}
    v3 = (float64(1e10) / (v3 * 0.1))
    if true {
        if len(v40) >= 2 {
            v40 = v40[1:len(v40)]
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v40:", len(v40), v40)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v4
    _ = v40
    _ = v42
    _ = v5
    _ = v6
}
