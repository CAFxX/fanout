package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32 }
type GenS2 struct { F0 int16; F1 uint8; F2 int8 }
type GenS3 struct { F0 uint32; F1 int32; F2 string }
type GenS4 struct { F0 int32 }
func genFn0(p0 uint8) int {
    fmt.Println(p0)
    fmt.Println(p0)
    if (42 == -1) {
        for fi6 := 0; fi6 < 7; fi6++ {
            v7 := GenS1{F0: 127, F1: 1}
            _ = fi6
            _ = v7
        }
        for fi8 := 0; fi8 < 11; fi8++ {
            if (32767 < 174) {
                var v9 uint16 = 513
                var v10 int64 = 9223372036854775807
                var v11 float32 = 0.1
                _ = v10
                _ = v11
                _ = v9
            }
            _ = fi8
        }
    }
    p0 += 241
    _ = p0
    return 0
}
func genFn1(p0 uint8) int8 {
    v12 := GenS2{F0: 32767, F1: 42, F2: 127}
    switch 128 {
    case 1:
        if true {
            v13 := GenS2{F0: 0, F1: 2, F2: -1}
            for fi14 := 0; fi14 < 19; fi14++ {
                for fi15 := 0; fi15 < 12; fi15++ {
                    if (1 == 0) {
                        var v16 float32 = ((float32(-2.25) - (1.5 + (float32(-2.25) / 1.0))) - 3.14159)
                        _ = v16
                    }
                    _ = fi15
                }
                for fi17 := 0; fi17 < 13; fi17++ {
                    p0 = p0
                    p0 = (p0 >> ((0) & 63))
                    _ = fi17
                }
                v18 := []int{fi14, 2147483647, 32768, (fi14 | (fi14 | fi14))}
                _ = fi14
                _ = v18
            }
            _ = v13
        }
    case -153:
        var v19 float64 = 1.5
        _ = v19
    case 2147483647:
        v20 := []int{}
        _ = v20
    default:
        p0 = (p0 * p0)
    }
    var v21 string = "foo bar"
    fmt.Println(v21)
    _ = p0
    _ = v12
    _ = v21
    return 1
}

func main() {
    var v1 int64 = 9223372036854775807
    v2 := make(map[uint32]bool)
    v3 := make(map[int]int)
    var v4 string = ((("hello" + (((("" + ("" + "a")) + "a") + "") + "a")) + "\n\t") + ("a" + "foo bar"))
    var v5 bool = false
    fmt.Println(v1)
    if (42 <= 4294967295) {
        if v5 {
            fmt.Println(v1, v1)
        }
        v5 = (9223372036854775807 > -218)
        v1 -= v1
    }
    switch -563 {
    case 9223372036854775807:
        v22 := []int{}
        _ = v22
    case -805:
        v23 := GenS3{F0: 2147483647, F1: 2147483647, F2: "hello"}
        _ = v23
    case -87:
        var v24 float64 = 1.5
        _ = v24
    default:
        if (v5 || (!v5)) {
            var v25 uint32 = 0
            v26 := []int{127, genFn0(127)}
            _ = v25
            _ = v26
        } else {
            v1 += -1
        }
    }
    fmt.Println(v1, v5, v4)
    for fi27 := 0; fi27 < 11; fi27++ {
        fmt.Println(v5)
        _ = fi27
    }
    var v28 float32 = 1.5
    var v29 int = -2147483648
    v30 := GenS4{F0: -1}
    v29 = (v29 + v29)
    for fi31 := 0; fi31 < 20; fi31++ {
        var v32 uint = 18446744073709551615
        v3[(fi31 / ((int(v32)) | 1))] = (fi31 + (^(v29)))
        fmt.Println(v29, v5)
        _ = fi31
        _ = v32
    }
    delete(v2, 2147483648)
    var v33 int = (^(v29))
    fmt.Println(v28, v1)
    if v5 {
        v34 := v2[95]
        v33 = 0
        _ = v34
    } else {
        var v35 float64 = 1e10
        _ = v35
    }
    v4 = ((v4 + v4) + "a")
    var v36 int32 = 2147483647
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v33
    _ = v36
    _ = v4
    _ = v5
}
