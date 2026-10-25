package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int16; F1 uint; F2 uint }
type GenS3 struct { F0 int32; F1 uint; F2 uint16 }
func genFn0(p0 int8) bool {
    if (0 < 28) {
        var v7 int8 = p0
        var v8 int8 = (v7 / (((v7 * v7)) | 1))
        for fi9 := 0; fi9 < 7; fi9++ {
            v8 = (v7 + p0)
            switch fi9 {
            case 1:
                fmt.Println(fi9, fi9, v8)
            case 4294967295:
                for fi10 := 0; fi10 < 3; fi10++ {
                    v8 += -128
                    _ = fi10
                }
            case -1:
                v11 := GenS1{F0: 0}
                _ = v11
            default:
                v12 := []int{fi9, (fi9 >> ((4294967295) & 63)), 2147483647}
                _ = v12
            }
            _ = fi9
        }
        _ = v7
        _ = v8
    } else {
        switch (("\n\t" + "hello") + ("\n\t" + "foo bar")) {
        case "foo bar":
            var v13 uint32 = 1
            _ = v13
        case "a":
            v14 := GenS2{F0: -104, F1: 4294967295, F2: 18446744073709551615}
            _ = v14
        case "":
            p0 += -101
        default:
            var v15 int64 = 82
            _ = v15
        }
    }
    var v16 uint8 = 2
    v16 = (^(v16))
    v16 = v16
    v17 := GenS1{F0: -184}
    _ = p0
    _ = v16
    _ = v17
    return false
}

func main() {
    v1 := make(map[string]uint32)
    var v2 uint64 = 18446744073709551615
    var v3 float32 = 1e10
    var v4 int64 = 9223372036854775807
    var v5 float64 = 0.0
    var v6 float32 = float32((1e10 + float64((3.14159 * float32(-2.25)))))
    var v18 int8 = 42
    v19 := GenS1{F0: -2147483648}
    v2 -= v2
    if (!(true && true)) {
        if true {
            if false {
                var v20 int32 = 0
                var v21 float32 = -2.25
                switch 42 {
                case 128:
                    var v22 uint = 536
                    _ = v22
                default:
                    fmt.Println(v21, v2)
                }
                _ = v20
                _ = v21
            } else {
                delete(v1, "hello")
                for fi23 := 0; fi23 < 20; fi23++ {
                    var v24 string = ("" + (("" + "") + "\n\t"))
                    var v25 uint16 = 0
                    fmt.Println(v2, v24, fi23)
                    _ = fi23
                    _ = v24
                    _ = v25
                }
                v26 := GenS2{F0: 0, F1: 2, F2: 616}
                _ = v26
            }
            if (!((true && false) || true)) {
                if false {
                    if (false && false) {
                        v5 -= v5
                        v18 = 0
                        v5 = (v5 * 0.1)
                    }
                } else {
                    for fi27 := 0; fi27 < 10; fi27++ {
                        v28 := []int{(fi27 + 2147483647), fi27}
                        _ = fi27
                        _ = v28
                    }
                    v29 := []int{32768, 4294967295, 686, 4294967295}
                    _ = v29
                }
                for fi30 := 0; fi30 < 10; fi30++ {
                    var v31 int16 = -1
                    var v32 uint64 = (v2 ^ v2)
                    fmt.Println(v3)
                    _ = fi30
                    _ = v31
                    _ = v32
                }
                v5 = 0.1
            }
        } else {
            v5 = v5
            if (-32768 <= -949) {
                v33 := []int{845, 0, 1}
                _ = v33
            }
            v34 := []int{127, 81}
            _ = v34
        }
        var v35 int = 32767
        v3 += 1.5
        _ = v35
    }
    if (false || false) {
        var v36 int64 = v4
        v2 += (v2 % ((2) | 1))
        var v37 float64 = v5
        _ = v36
        _ = v37
    }
    v38 := []int{-9223372036854775808, 1, 128}
    var v39 bool = (v18 != (v18 ^ (v18 & (v18 + v18))))
    v40 := GenS3{F0: -842, F1: 764, F2: 0}
    var v41 bool = v39
    for fi42 := 0; fi42 < 7; fi42++ {
        switch 127 {
        case -1:
            v3 -= v6
        case 2147483647:
            var v43 int32 = 42
            _ = v43
        }
        if (v41 && v39) {
            v5 -= 3.14159
            v39 = v39
            v1["foo bar"] = 34
        }
        _ = fi42
    }
    v1["a"] = 683
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v38:", len(v38), v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v6
}
