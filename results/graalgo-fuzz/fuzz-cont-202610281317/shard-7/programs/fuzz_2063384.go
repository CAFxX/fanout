package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint8; F1 uint8; F2 uint }
type GenS3 struct { F0 int64; F1 string }
type GenS4 struct { F0 int8; F1 uint8 }
type GenS5 struct { F0 uint64; F1 int8; F2 uint8 }
type GenS6 struct { F0 uint8; F1 string }
func genFn0(p0 int16) string {
    if (true && true) {
        v7 := []int{-1, 32768, -2147483648, -36}
        _ = v7
    } else {
        for fi8 := 0; fi8 < 6; fi8++ {
            var v9 int16 = (p0 + p0)
            v10 := GenS1{F0: 32767}
            _ = fi8
            _ = v10
            _ = v9
        }
    }
    for fi11 := 0; fi11 < 9; fi11++ {
        if (!(!((!(false || true)) || false))) {
            var v12 uint32 = 417
            fmt.Println(p0, p0)
            _ = v12
        } else {
            if false {
                if (2 != 4294967295) {
                    var v13 uint8 = 192
                    v14 := []int{(fi11 | fi11), 127}
                    _ = v13
                    _ = v14
                }
                v15 := []int{fi11, (fi11 ^ (-(fi11))), 32767}
                _ = v15
            }
            v16 := GenS1{F0: 32767}
            _ = v16
        }
        _ = fi11
    }
    if (false || false) {
        p0 += -303
    } else {
        p0 = 1000
        if ((3.14159 * ((1.5 + (1.5 * 1e10)) + (1e10 / float32((1.5 * ((0.1 - float64(1.5)) - 0.0)))))) > 0.1) {
            for fi17 := 0; fi17 < 12; fi17++ {
                p0 = (p0 + (-(p0)))
                switch "\n\t" {
                case "":
                    v18 := []int{}
                    _ = v18
                default:
                    v19 := []int{127, (^(fi17)), -2147483648, (-(fi17))}
                    _ = v19
                }
                var v20 float64 = (((float64(1e10) * float64(float32((1e10 * ((0.0 / 3.14159) * 3.14159))))) - (((3.14159 + 1.5) - 0.0) - 1e10)) + (float64(float32(-2.25)) + 1e10))
                _ = fi17
                _ = v20
            }
            var v21 uint32 = 453
            _ = v21
        } else {
            v22 := GenS1{F0: 0}
            _ = v22
        }
        var v23 bool = (false && true)
        _ = v23
    }
    v24 := GenS1{F0: 9223372036854775807}
    _ = p0
    _ = v24
    return "a"
}

func main() {
    var v1 uint16 = 778
    var v2 float32 = ((float32(0.0) + 3.14159) + 0.0)
    v3 := []int{-2147483648}
    var v4 uint = 18446744073709551615
    v5 := make(map[string]bool)
    v6 := make(map[string]bool)
    v1 = (v1 ^ v1)
    v4 += (v4 % ((v4) | 1))
    v2 += ((v2 - v2) + v2)
    v25 := GenS2{F0: 137, F1: 2, F2: v4}
    var v26 uint16 = v1
    for fi27 := 0; fi27 < 10; fi27++ {
        fmt.Println(fi27, v2)
        fmt.Println(v4, fi27)
        fmt.Println(v2)
        _ = fi27
    }
    if (2 <= v26) {
        var v28 int32 = 1
        _ = len(v3)
        if false {
            if ((v1 * (v1 - (v1 << ((v4) & 63)))) >= v26) {
                v29 := GenS3{F0: 42, F1: "\n\t"}
                _ = v29
            } else {
                v30 := GenS4{F0: 66, F1: 61}
                v31 := GenS2{F0: 2, F1: 2, F2: 582}
                for fi32 := 0; fi32 < 8; fi32++ {
                    v2 += (v2 / ((v2 - v2) * float32(1.5)))
                    _ = fi32
                }
                _ = v30
                _ = v31
            }
            if true {
                for fi33 := 0; fi33 < 20; fi33++ {
                    for fi34 := 0; fi34 < 2; fi34++ {
                        switch (fi33 - (^(fi34))) {
                        case 128:
                            switch "a" {
                            case "":
                                v2 = (((((float32(0.0) * (float32(0.1) * 1.5)) - v2) / 1e10) * -2.25) - ((v2 * ((0.0 + (3.14159 * 0.1)) / float32(-2.25))) / v2))
                            case "foo bar":
                                v1 = (v26 * 2)
                            default:
                                var v35 float32 = ((v2 - (v2 - 1.5)) + v2)
                                _ = v35
                            }
                        case 1:
                            delete(v6, "")
                        }
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        switch (fi33 * (^(fi34))) {
                        case 520:
                            v36 := GenS5{F0: 473, F1: 1, F2: 255}
                            _ = v36
                        }
                        _ = fi34
                    }
                    _ = fi33
                }
            }
            switch ((("a" + ("\n\t" + "a")) + "a") + "foo bar") {
            case "a":
                v26 = v1
            case "":
                var v37 int16 = -32768
                _ = v37
            case "hello":
                v1 = v26
            default:
                var v38 uint8 = 2
                _ = v38
            }
        }
        _ = v28
    }
    var v39 int32 = -1
    var v40 float64 = 0.0
    switch -2147483648 {
    case 270:
        fmt.Println(v1)
    case 778:
        v40 -= 0.0
    case -514:
        fmt.Println(v4, v4, v4)
    }
    var v41 int64 = -440
    v4 -= 42
    var v42 int8 = int8(v26)
    switch 32768 {
    case 32768:
        v39 += 286
    case 2147483647:
        v2 = (float32(v40) + v2)
    default:
        fmt.Println(v40, v39)
    }
    v43 := GenS6{F0: 100, F1: "\n\t"}
    fmt.Println(v41, v26)
    v41 = int64(v42)
    v26 = 65535
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v5
    _ = v6
}
