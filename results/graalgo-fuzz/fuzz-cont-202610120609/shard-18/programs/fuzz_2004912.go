package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8; F2 string }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int8 }
func genFn0(p0 uint16, p1 int64, p2 int32) string {
    v7 := []int{42, 127}
    var v8 float64 = 0.0
    var v9 uint32 = 2147483647
    for fi10 := 0; fi10 < 5; fi10++ {
        v11 := GenS1{F0: "", F1: -21, F2: "\n\t"}
        if false {
            switch fi10 {
            case -562:
                var v12 float32 = float32((v8 / (((v8 * ((v8 / (v8 - (v8 / ((float64(0.0) + 1e10) - v8)))) * (1e10 / 1.0))) * (1e10 / v8)) + float64(1e10))))
                _ = v12
            case -9223372036854775808:
                if len(v7) >= 2 {
                    v7 = v7[1:len(v7)]
                }
            }
            for fi13 := 0; fi13 < 17; fi13++ {
                var v14 float32 = float32((-2.25 / ((float64(1.5) - ((v8 - ((-2.25 / 0.1) / (float64((3.14159 - float32(((v8 + v8) - -2.25)))) + ((v8 - (v8 / (float64(1e10) * v8))) + -2.25)))) / (v8 + (0.1 - v8)))) - float64(1.5))))
                p2 += (p2 / ((int32(p0)) | 1))
                p0 -= p0
                _ = fi13
                _ = v14
            }
            v9 = v9
        }
        _ = fi10
        _ = v11
    }
    v15 := GenS1{F0: "", F1: -78, F2: "a"}
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    _ = v7
    _ = v8
    _ = v9
    return "hello"
}
func genFn1(p0 uint64) uint64 {
    fmt.Println(p0)
    p0 = p0
    _ = p0
    return 18446744073709551615
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 string = "a"
    v3 := make(map[int]uint32)
    v4 := make(map[string]string)
    v5 := make(map[int]string)
    v6 := []int{42, 1, 1}
    switch ("foo bar" + (v2 + v2)) {
    case "\n\t":
        var v16 uint16 = 65535
        _ = v16
    }
    for fi17 := 0; fi17 < 9; fi17++ {
        if false {
            v18 := GenS1{F0: ("a" + v2), F1: -128, F2: (("" + "a") + "")}
            _ = v18
        }
        var v19 uint = 1
        _ = fi17
        _ = v19
    }
    v2 = (((v2 + v2) + "") + "a")
    v5[128] = v2
    v2 = "a"
    for fi20 := 0; fi20 < 6; fi20++ {
        v21 := GenS1{F0: v2, F1: 123, F2: ("a" + v2)}
        _ = fi20
        _ = v21
    }
    var v22 uint = 1
    if false {
        v22 += v22
        if true {
            v23 := v3[-9223372036854775808]
            fmt.Println(v2)
            var v24 uint64 = 2
            _ = v23
            _ = v24
        } else {
            switch v2 {
            case "foo bar":
                fmt.Println(v2, v2)
            case "a":
                v6 = append(v6, 695)
            default:
                v25 := GenS2{F0: 1}
                _ = v25
            }
        }
        v26 := v3[4294967295]
        _ = v26
    }
    v27 := GenS2{F0: 127}
    switch v2 {
    case "":
        v22 -= 1
    default:
        v22 = v22
    }
    for fi28 := 0; fi28 < 14; fi28++ {
        fmt.Println(v22)
        for fi29 := 0; fi29 < 15; fi29++ {
            v30 := GenS3{F0: -1}
            fmt.Println(v2, v2)
            _ = len(v6)
            _ = fi29
            _ = v30
        }
        if ((2 == 42) && true) {
            v2 = v2
            v2 = v2
        } else {
            var v31 float64 = -2.25
            v32 := GenS4{F0: -1}
            _ = v31
            _ = v32
        }
        _ = fi28
    }
    var v33 int16 = -599
    var v34 int8 = 42
    var v35 int16 = v33
    fmt.Println(v35, v22)
    var v36 bool = true
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v27
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
