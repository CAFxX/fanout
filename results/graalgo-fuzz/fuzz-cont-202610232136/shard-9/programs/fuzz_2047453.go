package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32; F2 uint32 }
type GenS2 struct { F0 bool; F1 uint }
func genFn0() uint16 {
    var v5 int8 = 42
    switch "\n\t" {
    case "foo bar":
        fmt.Println(v5)
    case "\n\t":
        if true {
            for fi6 := 0; fi6 < 10; fi6++ {
                switch ("a" + ("a" + "foo bar")) {
                case "foo bar":
                    switch 32767 {
                    case 128:
                        v7 := GenS1{F0: 255, F1: -86, F2: 321}
                        _ = v7
                    case 1:
                        switch (-(fi6)) {
                        case -900:
                            v8 := GenS1{F0: 18446744073709551615, F1: 0, F2: 2}
                            _ = v8
                        case 2147483647:
                            var v9 int8 = (-(v5))
                            _ = v9
                        case 127:
                            var v10 int8 = (v5 & v5)
                            _ = v10
                        default:
                            for fi11 := 0; fi11 < 11; fi11++ {
                                v5 += v5
                                v12 := []int{fi11, (fi6 / (((fi6 & 32767)) | 1))}
                                _ = fi11
                                _ = v12
                            }
                        }
                    default:
                        v5 = v5
                    }
                case "":
                    v5 -= v5
                }
                _ = fi6
            }
        }
    }
    fmt.Println(v5)
    for fi13 := 0; fi13 < 2; fi13++ {
        fmt.Println(fi13, v5)
        if false {
            v14 := GenS1{F0: 18446744073709551615, F1: 42, F2: 2147483647}
            for fi15 := 0; fi15 < 13; fi15++ {
                v5 = v5
                v5 = v5
                switch "" {
                case "\n\t":
                    v16 := GenS2{F0: false, F1: 4294967295}
                    _ = v16
                case "hello":
                    v17 := []int{fi15, fi15, fi13}
                    _ = v17
                }
                _ = fi15
            }
            _ = v14
        }
        v18 := []int{(fi13 / ((fi13) | 1)), fi13, 2147483647, (fi13 / ((fi13) | 1))}
        _ = fi13
        _ = v18
    }
    var v19 uint = 2
    _ = v19
    _ = v5
    return 0
}
func genFn1(p0 int8, p1 int32) int64 {
    p1 += p1
    var v20 uint = 426
    v21 := GenS1{F0: v20, F1: p1, F2: 347}
    if true {
        v22 := []int{32768, 1, -44}
        p0 = p0
        switch "foo bar" {
        case "":
            v20 = (v20 | 0)
        case "hello":
            switch 470 {
            case 2147483647:
                if len(v22) > 0 {
                    v22[0] = -697
                }
            case 683:
                var v23 int32 = (p1 * (p1 | (p1 & p1)))
                _ = v23
            default:
                if (!(136 != 0)) {
                    var v24 uint = v20
                    var v25 float32 = (float32(0.0) / 1.5)
                    _ = len(v22)
                    _ = v24
                    _ = v25
                } else {
                    if len(v22) >= 2 {
                        v22 = v22[1:len(v22)]
                    }
                }
            }
        default:
            v26 := GenS2{F0: (4294967295 <= 42), F1: (v20 | (v20 | (v20 | 0)))}
            _ = v26
        }
        _ = v22
    }
    _ = p0
    _ = p1
    _ = v20
    _ = v21
    return 506
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[uint32]string)
    v3 := make(map[uint32]bool)
    var v4 int16 = 1
    var v27 int16 = v4
    v28 := []int{-550, -958}
    v4 = v4
    if (!(!(970 > 32767))) {
        switch "" {
        case "foo bar":
            _ = len(v1)
        case "":
            var v29 uint = 255
            _ = v29
        }
        _ = len(v28)
        v27 -= 672
    }
    switch "foo bar" {
    case "a":
        var v30 float64 = (0.0 - -2.25)
        _ = v30
    }
    if true {
        for fi31 := 0; fi31 < 20; fi31++ {
            v32 := GenS2{F0: ((65535 > 255) || false), F1: 4294967295}
            var v33 uint = 1
            _ = fi31
            _ = v32
            _ = v33
        }
        delete(v3, 4294967295)
    }
    switch "a" {
    case "a":
        var v34 uint8 = 42
        _ = v34
    case "hello":
        var v35 float64 = ((float64(0.0) * (-2.25 + (-2.25 - (0.1 * (1.5 - 1.5))))) + (3.14159 - 0.1))
        _ = v35
    case "\n\t":
        if false {
            v4 = v4
        }
    default:
        if true {
            if false {
                fmt.Println(v4, v27)
                var v36 int = 476
                _ = v36
            }
        }
    }
    if len(v28) >= 2 {
        v28 = v28[1:len(v28)]
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
}
