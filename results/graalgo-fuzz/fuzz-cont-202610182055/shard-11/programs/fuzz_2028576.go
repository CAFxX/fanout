package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int }
func genFn0(p0 uint, p1 uint8, p2 uint64) uint64 {
    fmt.Println(p2, p0, p2)
    if (p1 >= (^(p1))) {
        v4 := []int{}
        p2 = 0
        for fi5 := 0; fi5 < 15; fi5++ {
            for fi6 := 0; fi6 < 8; fi6++ {
                p2 += (p2 - (p2 & p2))
                var v7 uint32 = uint32(p0)
                var v8 int8 = 113
                _ = fi6
                _ = v7
                _ = v8
            }
            var v9 uint64 = 42
            _ = fi5
            _ = v9
        }
        _ = v4
    }
    _ = p0
    _ = p1
    _ = p2
    return (^(p2))
}

func main() {
    var v1 float64 = (0.1 / 0.1)
    var v2 float64 = (0.1 - ((v1 / 1e10) / float64((((-2.25 * ((0.0 + -2.25) / 3.14159)) * float32(v1)) - (1e10 * (1.5 * (-2.25 + 1.5)))))))
    v3 := []int{42, 450}
    if (65535 >= 100) {
        for fi10 := 0; fi10 < 12; fi10++ {
            v1 += (-2.25 - v1)
            _ = fi10
        }
        for fi11 := 0; fi11 < 6; fi11++ {
            switch (fi11 ^ (-(fi11))) {
            case -124:
                v1 -= (1.5 + 0.1)
            case -1:
                v2 += 3.14159
            default:
                var v12 float64 = v2
                _ = v12
            }
            _ = fi11
        }
    } else {
        v1 = 1e10
    }
    if len(v3) > 0 {
        v3[0] = 42
    }
    for fi13 := 0; fi13 < 10; fi13++ {
        var v14 bool = (396 <= 42)
        v15 := GenS1{F0: 4294967295, F1: true}
        _ = fi13
        _ = v14
        _ = v15
    }
    fmt.Println(v1)
    for fi16 := 0; fi16 < 14; fi16++ {
        switch fi16 {
        case 128:
            if (((("\n\t" + "") + "\n\t") + (("\n\t" + "a") + "")) < "") {
                var v17 uint16 = 32767
                _ = v17
            } else {
                v18 := GenS1{F0: 65535, F1: true}
                var v19 uint16 = 2
                var v20 int = fi16
                _ = v18
                _ = v19
                _ = v20
            }
        default:
            v21 := GenS2{F0: 1}
            _ = v21
        }
        switch ("" + "hello") {
        case "a":
            v22 := GenS3{F0: 635}
            _ = v22
        case "hello":
            if false {
                switch (fi16 / ((fi16) | 1)) {
                case 9223372036854775807:
                    v1 = v2
                }
                for fi23 := 0; fi23 < 12; fi23++ {
                    if (2147483647 == -2147483648) {
                        v3 = append(v3, (fi23 & fi16))
                    } else {
                        var v24 int8 = 0
                        v3 = append(v3, 32768)
                        v1 = v1
                        if true {
                            if true {
                                v3 = append(v3, fi16)
                            }
                            if false {
                                v3 = append(v3, (fi16 | fi23))
                                v3 = append(v3, -1)
                            } else {
                                var v25 uint32 = 668
                                _ = v25
                            }
                        }
                        _ = v24
                    }
                    if ((0 < -32768) && true) {
                        v26 := GenS3{F0: 9223372036854775807}
                        fmt.Println(v2, fi23)
                        _ = v26
                    } else {
                        v2 = v2
                        v1 = (0.0 + (float64(0.0) * v1))
                        switch "foo bar" {
                        case "hello":
                            if (!false) {
                                v27 := GenS2{F0: 0}
                                v1 = (v2 - 1.5)
                                if (!false) {
                                    var v28 float32 = (((-2.25 * 1.5) - ((((3.14159 - (((float32(1e10) + float32((v1 * v1))) - 1e10) - float32(v1))) / (-2.25 - 1e10)) * ((3.14159 - 0.1) * (0.0 * 0.1))) * ((((-2.25 / float32(v2)) * -2.25) * ((float32(1e10) * (0.1 * (0.1 + float32(((v1 / 1.0) - 1.5))))) - 1e10)) + 0.1))) / -2.25)
                                    fmt.Println(fi16)
                                    v2 = (3.14159 / float64(((float32(((-2.25 / 1.5) / 0.1)) * float32((v2 - v2))) - (v28 - v28))))
                                    _ = v28
                                }
                                _ = v27
                            }
                        case "\n\t":
                            if (false && (false || true)) {
                                v29 := GenS1{F0: 255, F1: (-128 >= -128)}
                                v1 = ((float64(1e10) + 1e10) / float64(-2.25))
                                _ = v29
                            }
                        default:
                            var v30 uint32 = 258
                            _ = v30
                        }
                    }
                    v31 := GenS2{F0: 772}
                    _ = fi23
                    _ = v31
                }
                var v32 int64 = 9223372036854775807
                _ = v32
            }
        }
        _ = fi16
    }
    v2 -= v2
    switch -9223372036854775808 {
    case -630:
        fmt.Println(v2, v2)
    default:
        for fi33 := 0; fi33 < 8; fi33++ {
            v2 = v2
            v34 := GenS4{F0: (fi33 ^ fi33)}
            v35 := GenS4{F0: (fi33 * (fi33 + (-(fi33))))}
            _ = fi33
            _ = v34
            _ = v35
        }
    }
    var v36 uint = 18446744073709551615
    v1 = (((v1 * 1.5) * 1.5) + (1.5 - 1e10))
    v2 = -2.25
    v3 = append(v3, 42)
    var v37 int16 = 741
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
}
