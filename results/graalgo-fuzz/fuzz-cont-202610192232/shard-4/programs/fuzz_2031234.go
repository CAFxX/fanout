package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16; F2 uint16 }
type GenS2 struct { F0 uint8; F1 int8; F2 int32 }
type GenS3 struct { F0 uint32; F1 int16 }
type GenS4 struct { F0 uint64; F1 int }
type GenS5 struct { F0 int64 }
func genFn0(p0 int) bool {
    fmt.Println(p0)
    p0 = p0
    var v7 int64 = -125
    _ = p0
    _ = v7
    return true
}
func genFn1(p0 uint64, p1 string, p2 int) int64 {
    v8 := []int{(^(p2)), 32768, (-(p2)), p2}
    var v9 int64 = int64(p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return 42
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[string]uint32)
    v3 := []int{2147483647, 884, 32767, 4294967295}
    var v4 uint16 = 32767
    v5 := []uint32{42, 2147483648, 574, 842}
    var v6 bool = (!((v4 * v4) >= (v4 << ((32) & 63))))
    var v10 int32 = 591
    var v11 uint16 = v4
    if (!(v6 || v6)) {
        fmt.Println(v6, v11)
        switch "hello" {
        case "hello":
            switch 127 {
            case -2147483648:
                var v12 uint32 = 65535
                _ = v12
            default:
                for fi13 := 0; fi13 < 5; fi13++ {
                    v14 := GenS1{F0: v4, F1: 555, F2: 62}
                    if (!v6) {
                        var v15 uint8 = 164
                        for fi16 := 0; fi16 < 2; fi16++ {
                            if (v10 > (v10 << ((4294967295) & 63))) {
                                if v6 {
                                    if (!((122 <= 18446744073709551615) || (!(fi16 > fi13)))) {
                                        if (777 >= 402) {
                                            v11 = (^(v4))
                                            v17 := v2[("hello" + "\n\t")]
                                            if (v6 && v6) {
                                                var v18 string = ""
                                                var v19 int16 = -883
                                                _ = len(v3)
                                                _ = v18
                                                _ = v19
                                            } else {
                                                switch ("a" + "hello") {
                                                case "":
                                                    var v20 float64 = 0.0
                                                    _ = v20
                                                case "hello":
                                                    fmt.Println(fi16, v11, v15)
                                                case "a":
                                                    var v21 float32 = 1e10
                                                    _ = v21
                                                }
                                            }
                                            _ = v17
                                        }
                                        fmt.Println(fi13, v15, v11)
                                    }
                                    v6 = (42 == 721)
                                    _ = len(v5)
                                    fmt.Println(fi16)
                                }
                            }
                            if v6 {
                                fmt.Println(v11)
                                v10 = int32(v15)
                                if v6 {
                                    v15 -= 173
                                    switch (fi13 - fi13) {
                                    case 32767:
                                        _ = len(v1)
                                    case -814:
                                        delete(v2, (("hello" + "foo bar") + "foo bar"))
                                    case 4294967295:
                                        var v22 string = "\n\t"
                                        _ = v22
                                    }
                                }
                            }
                            _ = fi16
                        }
                        var v23 float32 = (0.1 / float32(((-2.25 * (3.14159 * 1.5)) / float64(1.5))))
                        _ = v15
                        _ = v23
                    }
                    _ = fi13
                    _ = v14
                }
            }
        case "a":
            v24 := GenS2{F0: 42, F1: -1, F2: 1}
            _ = v24
        }
    }
    var v25 float32 = 0.0
    var v26 float64 = -2.25
    v27 := GenS3{F0: 220, F1: 1}
    v25 = (v25 - (v25 / (((v25 - float32(((v26 / 1e10) - v26))) / v25) * 3.14159)))
    var v28 int16 = -32768
    v5 = append(v5, 1)
    v28 += (v28 << ((2) & 63))
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v29 float32 = 0.1
    if (!v6) {
        if v6 {
            for fi30 := 0; fi30 < 17; fi30++ {
                v6 = ((9223372036854775807 == -401) || true)
                var v31 string = "foo bar"
                _ = fi30
                _ = v31
            }
            fmt.Println(v6, v25, v25)
            v32 := GenS1{F0: (v4 | 154), F1: (-(v28)), F2: v11}
            _ = v32
        } else {
            v25 = (-2.25 + (-2.25 - v29))
            v33 := GenS3{F0: 0, F1: -32768}
            v34 := GenS3{F0: 0, F1: (v28 / ((v28) | 1))}
            _ = v33
            _ = v34
        }
        var v35 uint8 = 128
        v10 -= (v10 % ((-1) | 1))
        _ = v35
    }
    fmt.Println(v4)
    v36 := GenS4{F0: 908, F1: 127}
    if (!(!(v6 || (!("hello" < "\n\t"))))) {
        v37 := GenS5{F0: 723}
        var v38 int32 = v10
        _ = v37
        _ = v38
    } else {
        var v39 int16 = v28
        v11 += 555
        _ = v39
    }
    var v40 uint64 = 1
    fmt.Println(v11, v29, v6)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v36
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
