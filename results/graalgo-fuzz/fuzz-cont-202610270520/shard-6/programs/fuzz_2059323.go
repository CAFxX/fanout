package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint8; F1 uint; F2 bool }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint32; F1 string; F2 int16 }
func genFn0() uint8 {
    if (false || true) {
        switch 127 {
        case 127:
            v7 := GenS1{F0: ((("a" + "hello") + "a") + "a")}
            _ = v7
        }
        if true {
            switch (("foo bar" + "") + "") {
            case "hello":
                v8 := GenS1{F0: "foo bar"}
                _ = v8
            case "foo bar":
                var v9 string = ""
                _ = v9
            default:
                var v10 int8 = 0
                _ = v10
            }
            v11 := []int{0}
            _ = v11
        }
        v12 := []int{4294967295, -934}
        _ = v12
    } else {
        for fi13 := 0; fi13 < 20; fi13++ {
            var v14 uint32 = 77
            _ = fi13
            _ = v14
        }
        var v15 uint = 18446744073709551615
        fmt.Println(v15)
        _ = v15
    }
    var v16 uint8 = 70
    var v17 float64 = -2.25
    v16 = v16
    if ("" != "hello") {
        switch -1 {
        case 228:
            v18 := []int{-9223372036854775808, 74, -394, -2147483648}
            _ = v18
        }
        v17 = ((0.1 * 1e10) / ((v17 / (v17 * v17)) * 0.0))
        v17 = v17
    } else {
        if ((v16 ^ v16) < 0) {
            v19 := []int{}
            for fi20 := 0; fi20 < 11; fi20++ {
                var v21 int64 = -9223372036854775808
                v16 += v16
                _ = fi20
                _ = v21
            }
            var v22 uint8 = (v16 * 90)
            _ = v19
            _ = v22
        } else {
            for fi23 := 0; fi23 < 10; fi23++ {
                v24 := GenS2{F0: 1}
                fmt.Println(fi23, v16, fi23)
                _ = fi23
                _ = v24
            }
            v25 := []int{}
            _ = v25
        }
    }
    _ = v16
    _ = v17
    return v16
}
func genFn1(p0 int64) uint64 {
    p0 += (p0 / ((p0) | 1))
    if true {
        for fi26 := 0; fi26 < 17; fi26++ {
            for fi27 := 0; fi27 < 1; fi27++ {
                if ((("\n\t" == "hello") || false) || true) {
                    v28 := GenS3{F0: 62, F1: 702, F2: false}
                    var v29 int32 = 42
                    _ = v28
                    _ = v29
                }
                _ = fi27
            }
            var v30 uint8 = 183
            v31 := []int{}
            _ = fi26
            _ = v30
            _ = v31
        }
        p0 = (p0 / ((p0) | 1))
    }
    switch ("foo bar" + (("\n\t" + "") + "")) {
    case "":
        p0 += (^(p0))
    case "\n\t":
        v32 := GenS3{F0: 127, F1: 65535, F2: false}
        _ = v32
    case "a":
        v33 := []int{9223372036854775807, 766, 0, 9223372036854775807}
        _ = v33
    }
    v34 := []int{239, 32767, 2147483647}
    _ = p0
    _ = v34
    return 2
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[int]string)
    var v3 float64 = 0.0
    v4 := []uint32{310}
    var v5 string = "foo bar"
    var v6 int32 = 2147483647
    for fi35 := 0; fi35 < 18; fi35++ {
        v36 := GenS3{F0: 2, F1: 0, F2: (1 == 18446744073709551615)}
        _ = fi35
        _ = v36
    }
    for fi37 := 0; fi37 < 3; fi37++ {
        var v38 int64 = 1
        _ = fi37
        _ = v38
    }
    switch 811 {
    case -643:
        v5 = (v5 + v5)
    case -1:
        v39 := GenS1{F0: ("foo bar" + ("hello" + v5))}
        _ = v39
    default:
        switch v5 {
        case "a":
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        case "\n\t":
            if false {
                var v40 float32 = -2.25
                _ = v40
            }
        case "foo bar":
            var v41 uint32 = 4294967295
            _ = v41
        default:
            v3 -= float64((1e10 * float32(v3)))
        }
    }
    var v42 uint64 = 2
    for fi43 := 0; fi43 < 1; fi43++ {
        fmt.Println(fi43, v5)
        _ = fi43
    }
    v42 += 18446744073709551615
    for fi44 := 0; fi44 < 18; fi44++ {
        var v45 int8 = 127
        v4 = append(v4, 0)
        v46 := GenS4{F0: 9223372036854775807}
        _ = fi44
        _ = v45
        _ = v46
    }
    var v47 int8 = -128
    fmt.Println(v5, v47)
    switch v5 {
    case "hello":
        v42 = v42
    case "":
        switch 2147483647 {
        case 127:
            if false {
                if (!((-2.25 + 0.1) <= (float32(v3) - (((1.5 + -2.25) / -2.25) + float32(((v3 * v3) * 3.14159)))))) {
                    switch -951 {
                    case 90:
                        fmt.Println(v47, v3, v6)
                    case 42:
                        v3 -= (1e10 - v3)
                    default:
                        if len(v4) > 0 {
                            v4[0] = 2147483648
                        }
                    }
                    var v48 uint16 = 449
                    if (((((((v5 + (v5 + (v5 + (v5 + "hello")))) + (v5 + "")) + v5) + v5) + v5) <= v5) && false) {
                        v3 += ((v3 * (v3 / float64(((((-2.25 + -2.25) * float32((v3 / (v3 - -2.25)))) / 0.1) + (0.0 * 1e10))))) * float64(0.0))
                    }
                    _ = v48
                }
            } else {
                var v49 float64 = (-2.25 / (v3 - v3))
                v4 = append(v4, 65535)
                _ = v49
            }
        default:
            var v50 uint16 = 0
            _ = v50
        }
    default:
        if false {
            for fi51 := 0; fi51 < 7; fi51++ {
                v52 := GenS4{F0: 9223372036854775807}
                _ = fi51
                _ = v52
            }
        }
    }
    v53 := GenS5{F0: 2147483648}
    fmt.Println(v5, v42)
    v54 := GenS6{F0: 696, F1: v5, F2: 0}
    v2[4294967295] = v5
    v3 = (v3 * (v3 * v3))
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    for fi55 := 0; fi55 < 7; fi55++ {
        v6 += (-(v6))
        _ = len(v4)
        v6 = v6
        _ = fi55
    }
    fmt.Println(v3, v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v42:", v42)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v47
    _ = v5
    _ = v53
    _ = v54
    _ = v6
}
