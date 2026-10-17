package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool; F2 string }
type GenS2 struct { F0 int16; F1 int8; F2 uint }
type GenS3 struct { F0 string; F1 int8 }
type GenS4 struct { F0 uint16; F1 uint16 }
type GenS5 struct { F0 int32; F1 int32; F2 uint64 }
type GenS6 struct { F0 bool }
func genFn0() string {
    v7 := GenS1{F0: 694, F1: (false || true), F2: "a"}
    v8 := GenS1{F0: 2, F1: (!false), F2: "foo bar"}
    if false {
        v9 := []int{1, -122, -793}
        _ = v9
    }
    var v10 int64 = 0
    switch -162 {
    case 0:
        v11 := GenS2{F0: 0, F1: 0, F2: 65535}
        _ = v11
    case 32767:
        var v12 int64 = (v10 & v10)
        _ = v12
    case 20:
        v10 += v10
    }
    _ = v10
    _ = v7
    _ = v8
    return "foo bar"
}

func main() {
    var v1 uint16 = 1
    var v2 float64 = -2.25
    var v3 uint64 = 736
    v4 := []int{-1, -9223372036854775808, -9223372036854775808}
    var v5 float32 = ((0.0 / float32((v2 / (1.5 + (((v2 * 3.14159) / ((v2 / float64(3.14159)) - 1e10)) - v2))))) / (3.14159 * 3.14159))
    var v6 uint8 = 255
    v4 = append(v4, 32768)
    for fi13 := 0; fi13 < 9; fi13++ {
        for fi14 := 0; fi14 < 20; fi14++ {
            v15 := GenS3{F0: "foo bar", F1: -1}
            if (!true) {
                v2 -= v2
                switch -1 {
                case 4294967295:
                    if true {
                        var v16 int16 = -32768
                        _ = v16
                    } else {
                        if true {
                            if ((fi13 | -9223372036854775808) > -9223372036854775808) {
                                var v17 string = ("a" + "hello")
                                v18 := GenS1{F0: v3, F1: false, F2: (v17 + v17)}
                                _ = v17
                                _ = v18
                            } else {
                                v6 += (v6 << ((4294967295) & 63))
                                _ = len(v4)
                                fmt.Println(v2, v2, fi13)
                            }
                        } else {
                            fmt.Println(fi14, v1)
                            if true {
                                var v19 int16 = 0
                                v20 := GenS2{F0: (v19 << ((4294967295) & 63)), F1: -98, F2: 65535}
                                _ = v19
                                _ = v20
                            }
                        }
                        v21 := GenS3{F0: "hello", F1: -128}
                        var v22 int64 = -9223372036854775808
                        _ = v21
                        _ = v22
                    }
                case -1:
                    v23 := GenS1{F0: (^(v3)), F1: (!true), F2: ((((("foo bar" + ("foo bar" + ("\n\t" + "\n\t"))) + "a") + "a") + ("\n\t" + "foo bar")) + "foo bar")}
                    _ = v23
                }
                v3 = v3
            } else {
                switch fi14 {
                case -1:
                    v2 -= (v2 - (3.14159 + ((v2 + (v2 + (v2 / v2))) - float64(v5))))
                }
                v5 = (v5 + float32(-2.25))
                var v24 uint32 = 0
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                _ = v24
            }
            v25 := GenS4{F0: 100, F1: (v1 + 76)}
            _ = fi14
            _ = v15
            _ = v25
        }
        var v26 int = fi13
        var v27 uint32 = 4294967295
        _ = fi13
        _ = v26
        _ = v27
    }
    v3 -= (v3 | (^(v3)))
    var v28 uint32 = 0
    v29 := GenS2{F0: -130, F1: -61, F2: 850}
    var v30 uint8 = uint8(v3)
    var v31 float32 = ((v5 / v5) / v5)
    var v32 int8 = 100
    v33 := GenS4{F0: v1, F1: v1}
    fmt.Println(v3, v6)
    v34 := GenS5{F0: 0, F1: 42, F2: v3}
    if (-9223372036854775808 >= 32767) {
        for fi35 := 0; fi35 < 20; fi35++ {
            fmt.Println(v28, v31, v2)
            if ((278 >= 1) || false) {
                var v36 uint = 2
                var v37 int32 = -517
                v38 := GenS6{F0: true}
                _ = v36
                _ = v37
                _ = v38
            }
            _ = fi35
        }
        v3 += (v3 / (((v3 >> ((42) & 63))) | 1))
        v39 := GenS2{F0: 1, F1: v32, F2: 65535}
        _ = v39
    }
    fmt.Println(v2, v5, v3)
    switch 0 {
    case 4294967295:
        fmt.Println(v30, v5, v1)
    }
    v40 := GenS1{F0: (v3 / ((uint64(v6)) | 1)), F1: (!false), F2: "a"}
    v41 := GenS5{F0: -1, F1: -1, F2: 51}
    v42 := GenS4{F0: v1, F1: (v1 + v1)}
    v32 -= -128
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v5
    _ = v6
}
