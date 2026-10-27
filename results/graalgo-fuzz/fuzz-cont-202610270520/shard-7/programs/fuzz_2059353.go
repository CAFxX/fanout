package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int32; F1 int64 }
type GenS5 struct { F0 uint16; F1 uint16; F2 int8 }
func genFn0() uint {
    for fi7 := 0; fi7 < 13; fi7++ {
        for fi8 := 0; fi8 < 8; fi8++ {
            v9 := GenS1{F0: 328}
            var v10 uint = 645
            _ = fi8
            _ = v10
            _ = v9
        }
        var v11 int16 = 0
        _ = fi7
        _ = v11
    }
    v12 := []int{-42, 36, 9223372036854775807}
    for fi13 := 0; fi13 < 12; fi13++ {
        var v14 float32 = 1.5
        for fi15 := 0; fi15 < 6; fi15++ {
            v16 := GenS2{F0: 845}
            var v17 uint64 = 0
            var v18 bool = (-667 != -1)
            v14 += float32(0.0)
            var v19 uint8 = uint8(v17)
            var v20 int = (fi15 | (fi13 * -2147483648))
            _ = fi15
            _ = v16
            _ = v17
            _ = v18
            _ = v19
            _ = v20
        }
        _ = fi13
        _ = v14
    }
    if (!true) {
        var v21 uint = 18446744073709551615
        _ = v21
    } else {
        var v22 int16 = -837
        _ = v22
    }
    switch 128 {
    case -9223372036854775808:
        var v23 int16 = 42
        _ = v23
    case 9223372036854775807:
        v24 := GenS2{F0: 9223372036854775807}
        _ = v24
    }
    _ = v12
    return 7
}

func main() {
    var v1 int = 452
    var v2 int8 = 42
    var v3 float64 = 0.0
    var v4 int32 = 0
    var v5 uint = 65535
    v6 := make(map[uint32]string)
    switch ("" + "foo bar") {
    case "":
        v25 := GenS3{F0: v5}
        _ = v25
    case "\n\t":
        v1 = v1
    default:
        switch "a" {
        case "a":
            if (false && true) {
                if false {
                    var v26 float32 = 0.0
                    for fi27 := 0; fi27 < 5; fi27++ {
                        var v28 float32 = (float32(float64(float32(v3))) / v26)
                        var v29 uint16 = 247
                        _ = fi27
                        _ = v28
                        _ = v29
                    }
                    v3 = (float64(((-2.25 * (v26 + v26)) - 0.1)) + (v3 * -2.25))
                    _ = v26
                }
                for fi30 := 0; fi30 < 6; fi30++ {
                    v31 := GenS4{F0: 1, F1: -790}
                    switch fi30 {
                    case -511:
                        v6[2] = "a"
                    default:
                        v2 = 100
                    }
                    v32 := GenS5{F0: 1, F1: 65535, F2: v2}
                    _ = fi30
                    _ = v31
                    _ = v32
                }
                var v33 string = ("foo bar" + "foo bar")
                _ = v33
            } else {
                var v34 uint = v5
                v35 := GenS4{F0: (v4 ^ v4), F1: 0}
                _ = v34
                _ = v35
            }
        default:
            if false {
                var v36 uint = 42
                _ = v36
            } else {
                v37 := GenS3{F0: 0}
                _ = v37
            }
        }
    }
    for fi38 := 0; fi38 < 12; fi38++ {
        v39 := v6[1]
        v1 = 127
        _ = fi38
        _ = v39
    }
    var v40 int16 = 809
    var v41 uint32 = 2147483647
    v6[(v41 - v41)] = "\n\t"
    fmt.Println(v5, v40)
    var v42 uint16 = 32767
    v43 := []int{v1}
    if len(v43) >= 2 {
        v43 = v43[1:len(v43)]
    }
    for fi44 := 0; fi44 < 17; fi44++ {
        v40 += 42
        var v45 string = "\n\t"
        _ = fi44
        _ = v45
    }
    if len(v43) >= 2 {
        v43 = v43[1:len(v43)]
    }
    for fi46 := 0; fi46 < 4; fi46++ {
        v47 := GenS2{F0: 1}
        _ = fi46
        _ = v47
    }
    v4 += (v4 - (v4 % ((-20) | 1)))
    var v48 uint64 = 347
    var v49 float32 = 3.14159
    fmt.Println(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("slicelen v43:", len(v43), v43)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v48
    _ = v49
    _ = v5
    _ = v6
}
