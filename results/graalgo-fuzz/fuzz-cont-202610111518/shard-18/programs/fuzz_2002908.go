package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint8; F1 uint16 }
type GenS3 struct { F0 uint8; F1 int32; F2 uint64 }
type GenS4 struct { F0 uint; F1 uint64 }
type GenS5 struct { F0 uint64; F1 bool; F2 uint32 }
type GenS6 struct { F0 string; F1 uint; F2 uint64 }
func genFn0(p0 int8) bool {
    for fi4 := 0; fi4 < 18; fi4++ {
        v5 := []int{}
        if false {
            fmt.Println(p0)
            var v6 uint = 0
            _ = v6
        } else {
            p0 += -1
            var v7 bool = false
            _ = v7
        }
        v8 := GenS1{F0: "foo bar"}
        _ = fi4
        _ = v5
        _ = v8
    }
    for fi9 := 0; fi9 < 19; fi9++ {
        var v10 int16 = 466
        v11 := GenS1{F0: ("" + "a")}
        _ = fi9
        _ = v10
        _ = v11
    }
    _ = p0
    return (925 < 533)
}
func genFn1(p0 string, p1 uint64, p2 int64) int8 {
    fmt.Println(p0, p2)
    var v12 int16 = -800
    v13 := []int{934, int(p1), 32768, 535}
    if len(v13) >= 2 {
        v13 = v13[1:len(v13)]
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v13
    return 126
}

func main() {
    var v1 string = ""
    v2 := make(map[int]bool)
    v3 := []int{-2147483648, 4294967295}
    v14 := GenS2{F0: 2, F1: 32767}
    for fi15 := 0; fi15 < 11; fi15++ {
        var v16 int16 = 0
        _ = fi15
        _ = v16
    }
    if true {
        v17 := v2[9223372036854775807]
        _ = v17
    }
    if true {
        var v18 uint = 644
        if (!true) {
            var v19 bool = ((v18 * 1) <= (^(v18)))
            var v20 int16 = 0
            _ = v19
            _ = v20
        }
        var v21 int64 = 0
        _ = v18
        _ = v21
    } else {
        var v22 string = (v1 + ((v1 + v1) + "a"))
        fmt.Println(v1)
        _ = v22
    }
    var v23 uint64 = 1
    var v24 bool = true
    fmt.Println(v23, v23, v23)
    var v25 float64 = float64((-2.25 * (3.14159 - ((((1e10 - 0.1) * (-2.25 / 0.1)) / (0.1 / (0.0 + (3.14159 / (0.1 / 0.1))))) + (-2.25 - 1.5)))))
    var v26 string = "\n\t"
    v1 = v26
    var v27 int = 0
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v28 := v2[(-(v27))]
    v29 := GenS3{F0: 1, F1: -475, F2: (v23 & v23)}
    for fi30 := 0; fi30 < 4; fi30++ {
        for fi31 := 0; fi31 < 12; fi31++ {
            if len(v3) > 0 {
                v3[0] = (^(fi31))
            }
            v32 := GenS4{F0: 896, F1: v23}
            switch fi30 {
            case 42:
                var v33 int16 = 53
                _ = v33
            case 32768:
                if (v28 && (v25 >= (v25 * -2.25))) {
                    v3 = append(v3, v27)
                } else {
                    var v34 uint8 = 107
                    _ = v34
                }
            default:
                v35 := v2[-544]
                _ = v35
            }
            _ = fi31
            _ = v32
        }
        if v28 {
            var v36 string = v26
            v37 := GenS5{F0: (v23 / (((v23 - 210)) | 1)), F1: (!v28), F2: 570}
            if (-432 != -9223372036854775808) {
                v38 := GenS5{F0: v23, F1: v24, F2: 915}
                _ = v38
            } else {
                v26 = ("\n\t" + v1)
                var v39 int32 = -756
                fmt.Println(v25, v36)
                _ = v39
            }
            _ = v36
            _ = v37
        } else {
            v28 = (-1 != genFn1("\n\t", (v23 / (((v23 * v23)) | 1)), 7))
        }
        v40 := GenS4{F0: 42, F1: (v23 % ((v23) | 1))}
        _ = fi30
        _ = v40
    }
    for fi41 := 0; fi41 < 17; fi41++ {
        _ = len(v3)
        var v42 uint16 = 232
        v43 := GenS6{F0: "\n\t", F1: 2, F2: (v23 << ((255) & 63))}
        _ = fi41
        _ = v42
        _ = v43
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v14
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
}
