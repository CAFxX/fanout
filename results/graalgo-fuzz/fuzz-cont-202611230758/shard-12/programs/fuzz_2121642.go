package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8 }
type GenS2 struct { F0 uint32; F1 int32 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 bool; F1 uint64 }
type GenS5 struct { F0 int8; F1 uint16; F2 uint64 }
func genFn0(p0 uint16, p1 int8, p2 bool) uint64 {
    v4 := []int{-9223372036854775808, 42, 9223372036854775807}
    var v5 uint32 = 0
    var v6 int16 = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    return 4294967295
}
func genFn1(p0 int64) int {
    var v7 uint = 0
    v8 := GenS1{F0: 0, F1: -128}
    _ = p0
    _ = v7
    _ = v8
    return 4294967295
}

func main() {
    var v1 uint32 = 2147483648
    var v2 float64 = -2.25
    var v3 uint16 = 65535
    fmt.Println(v2, v3, v3)
    v3 -= 32767
    fmt.Println(v3, v2, v3)
    if (128 >= 71) {
        if ((v3 | v3) > (^(v3))) {
            if true {
                if false {
                    switch "a" {
                    case "foo bar":
                        if false {
                            v3 = v3
                        } else {
                            v9 := GenS2{F0: v1, F1: 2147483647}
                            _ = v9
                        }
                    }
                } else {
                    v10 := GenS3{F0: 98}
                    if true {
                        fmt.Println(v2, v1)
                        var v11 int64 = -432
                        _ = v11
                    } else {
                        v12 := []int{1, 638, -182, 4294967295}
                        v1 += 2147483648
                        var v13 int32 = 42
                        _ = v12
                        _ = v13
                    }
                    v14 := GenS1{F0: 0, F1: 42}
                    _ = v10
                    _ = v14
                }
            } else {
                var v15 float64 = v2
                switch -767 {
                case 128:
                    v3 -= 1
                case 32768:
                    v16 := GenS2{F0: v1, F1: -1}
                    _ = v16
                case 570:
                    fmt.Println(v3, v2)
                }
                v17 := []int{}
                _ = v15
                _ = v17
            }
            v1 = (v1 / ((v1) | 1))
            v18 := GenS4{F0: ((32767 == -713) || (!false)), F1: 441}
            _ = v18
        } else {
            for fi19 := 0; fi19 < 9; fi19++ {
                v20 := GenS4{F0: false, F1: 704}
                var v21 int64 = 0
                _ = fi19
                _ = v20
                _ = v21
            }
            fmt.Println(v3)
        }
        var v22 int16 = 446
        _ = v22
    }
    if (false && (!(0 >= 42))) {
        for fi23 := 0; fi23 < 1; fi23++ {
            v24 := GenS5{F0: 1, F1: 32767, F2: 2}
            var v25 int = (fi23 | 32767)
            _ = fi23
            _ = v24
            _ = v25
        }
        var v26 uint32 = 422
        fmt.Println(v1)
        _ = v26
    } else {
        for fi27 := 0; fi27 < 7; fi27++ {
            var v28 bool = (false || true)
            v29 := GenS3{F0: -228}
            v30 := GenS2{F0: 65535, F1: int32(fi27)}
            _ = fi27
            _ = v28
            _ = v29
            _ = v30
        }
        v3 = 32767
        v3 = 255
    }
    switch "\n\t" {
    case "":
        v31 := []int{127, -1, -1, 508}
        _ = v31
    case "hello":
        v32 := GenS4{F0: (v2 <= float64(((1e10 / 1e10) * -2.25))), F1: 2}
        _ = v32
    default:
        v3 += 255
    }
    for fi33 := 0; fi33 < 12; fi33++ {
        v2 += 0.1
        var v34 float32 = 1.5
        _ = fi33
        _ = v34
    }
    v3 = 42
    switch -2147483648 {
    case 32768:
        var v35 int8 = 42
        _ = v35
    case 1:
        v36 := GenS4{F0: false, F1: 432}
        _ = v36
    case -393:
        fmt.Println(v1, v3, v2)
    }
    var v37 string = "foo bar"
    var v38 uint16 = (v3 % (((v3 - (^(v3)))) | 1))
    var v39 bool = true
    for fi40 := 0; fi40 < 12; fi40++ {
        v38 = v3
        v41 := GenS4{F0: v39, F1: 2}
        _ = fi40
        _ = v41
    }
    switch (v37 + v37) {
    case "":
        if false {
            var v42 uint8 = 105
            _ = v42
        }
    case "foo bar":
        var v43 uint16 = (v3 | (v3 % (((v3 - 0)) | 1)))
        _ = v43
    }
    var v44 int8 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v44
}
