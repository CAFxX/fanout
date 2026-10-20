package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int8; F1 int32; F2 int }
type GenS3 struct { F0 uint32; F1 int32; F2 uint }
func genFn0(p0 bool, p1 int32) uint8 {
    v5 := []int{0, 32767, 0}
    switch 128 {
    case 4294967295:
        for fi6 := 0; fi6 < 10; fi6++ {
            var v7 uint = 42
            p0 = (p0 && (p1 >= p1))
            v7 += 0
            _ = fi6
            _ = v7
        }
    default:
        var v8 uint64 = 18446744073709551615
        _ = v8
    }
    var v9 uint32 = 346
    if (!((((-2.25 - 0.0) + 3.14159) * ((3.14159 * float64(-2.25)) / 1e10)) <= 1.5)) {
        fmt.Println(p0)
        fmt.Println(p0, v9, p0)
    } else {
        var v10 uint8 = 255
        switch 42 {
        case 448:
            if ((65535 != 526) || ((1e10 * ((0.1 - (1e10 - (1.5 - 1e10))) + -2.25)) <= 0.0)) {
                v5 = append(v5, 9223372036854775807)
                v5 = append(v5, 4294967295)
            } else {
                for fi11 := 0; fi11 < 17; fi11++ {
                    fmt.Println(p1)
                    v5 = append(v5, 32768)
                    v9 += (^(v9))
                    _ = fi11
                }
                for fi12 := 0; fi12 < 19; fi12++ {
                    if (!false) {
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        if (-2147483648 > (fi12 & (fi12 << ((255) & 63)))) {
                            var v13 int64 = 1
                            p0 = (false && (-641 > 604))
                            _ = v13
                        }
                    }
                    p0 = (!(378 != -1))
                    switch "a" {
                    case "foo bar":
                        v5 = append(v5, (fi12 % ((int(p1)) | 1)))
                    case "a":
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                    default:
                        v14 := GenS1{F0: v10}
                        _ = v14
                    }
                    _ = fi12
                }
            }
        }
        _ = v10
    }
    for fi15 := 0; fi15 < 4; fi15++ {
        var v16 uint = 255
        _ = fi15
        _ = v16
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v9
    return 1
}
func genFn1(p0 int16, p1 int, p2 bool) bool {
    var v17 uint8 = 255
    v17 += (v17 | (v17 >> ((983) & 63)))
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    return (4294967295 != 603)
}

func main() {
    var v1 uint32 = 231
    var v2 uint32 = 0
    var v3 int16 = 0
    var v4 int16 = -32768
    var v18 string = "a"
    v19 := []int{32768}
    v3 = 42
    v20 := GenS2{F0: -100, F1: -2147483648, F2: -91}
    if (!((v2 >= (v1 | (v1 ^ v1))) && true)) {
        for fi21 := 0; fi21 < 16; fi21++ {
            fmt.Println(v3, v3, v18)
            _ = fi21
        }
    }
    for fi22 := 0; fi22 < 18; fi22++ {
        _ = len(v19)
        fmt.Println(v2)
        for fi23 := 0; fi23 < 1; fi23++ {
            var v24 float32 = 1.5
            _ = fi23
            _ = v24
        }
        _ = fi22
    }
    var v25 bool = false
    v19 = append(v19, 0)
    var v26 int8 = 71
    for fi27 := 0; fi27 < 9; fi27++ {
        var v28 float64 = (0.1 / (1e10 - (0.0 + ((float64(0.1) / ((1e10 * 0.1) + (1.5 - ((((0.0 + (1.5 / 1e10)) + (0.0 + 0.0)) / (-2.25 - 1e10)) * 0.1)))) - (1.5 + (((3.14159 * (1.5 * -2.25)) / 1.0) + 3.14159))))))
        v29 := GenS3{F0: (v2 >> ((0) & 63)), F1: 0, F2: 65535}
        _ = fi27
        _ = v28
        _ = v29
    }
    if ((!(v26 != (v26 / ((-128) | 1)))) && (!(!false))) {
        if (-2.25 > -2.25) {
            var v30 bool = false
            _ = v30
        }
        switch 0 {
        case 524:
            switch "hello" {
            case "a":
                for fi31 := 0; fi31 < 20; fi31++ {
                    var v32 uint16 = 2
                    _ = fi31
                    _ = v32
                }
            case "foo bar":
                var v33 uint64 = 18446744073709551615
                _ = v33
            }
        case 2147483647:
            v19 = append(v19, -2147483648)
        }
    }
    if true {
        var v34 uint8 = 92
        _ = v34
    }
    var v35 int16 = (v3 / ((v4) | 1))
    var v36 bool = v25
    v37 := GenS2{F0: (v26 + v26), F1: -165, F2: 36}
    fmt.Println(v4)
    fmt.Println(v35, v26)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v25
    _ = v26
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v4
}
