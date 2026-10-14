package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16 }
type GenS2 struct { F0 int; F1 string }
type GenS3 struct { F0 int8; F1 int8 }
type GenS4 struct { F0 uint64; F1 int32; F2 uint64 }
type GenS5 struct { F0 uint; F1 int32; F2 uint16 }
type GenS6 struct { F0 int32; F1 string; F2 int }
func genFn0(p0 int64) uint16 {
    v4 := GenS1{F0: 0, F1: 992}
    var v5 int16 = 243
    _ = p0
    _ = v4
    _ = v5
    return 65535
}

func main() {
    var v1 int32 = 2147483647
    var v2 uint32 = 975
    var v3 float64 = 1.5
    for fi6 := 0; fi6 < 16; fi6++ {
        for fi7 := 0; fi7 < 11; fi7++ {
            if ((^(v1)) > (v1 / ((v1) | 1))) {
                v8 := GenS2{F0: (fi7 - fi6), F1: "a"}
                v1 -= (v1 >> ((42) & 63))
                var v9 int8 = 1
                v10 := []int{}
                fmt.Println(v2, v2, v3)
                _ = v10
                _ = v8
                _ = v9
            } else {
                if (407 >= -1) {
                    var v11 int32 = v1
                    _ = v11
                } else {
                    v1 -= (v1 >> ((953) & 63))
                }
                switch 1 {
                case 42:
                    switch ("hello" + "hello") {
                    case "":
                        fmt.Println(v3, v1)
                    case "foo bar":
                        if (!false) {
                            v12 := GenS3{F0: 118, F1: 0}
                            var v13 int8 = 127
                            _ = v12
                            _ = v13
                        }
                    case "\n\t":
                        v14 := []int{fi7, 4294967295}
                        _ = v14
                    default:
                        v15 := GenS3{F0: -1, F1: 45}
                        _ = v15
                    }
                case 128:
                    v16 := GenS2{F0: fi6, F1: ""}
                    _ = v16
                case 2147483647:
                    if true {
                        var v17 float64 = float64(3.14159)
                        var v18 int64 = 42
                        _ = v17
                        _ = v18
                    } else {
                        v1 = -2147483648
                        v1 = 42
                        var v19 int8 = -128
                        _ = v19
                    }
                }
            }
            v20 := GenS4{F0: 42, F1: v1, F2: 4294967295}
            var v21 int8 = 0
            _ = fi7
            _ = v20
            _ = v21
        }
        _ = fi6
    }
    for fi22 := 0; fi22 < 11; fi22++ {
        for fi23 := 0; fi23 < 13; fi23++ {
            v2 -= (v2 & v2)
            _ = fi23
        }
        v1 += -427
        _ = fi22
    }
    var v24 int32 = 42
    if true {
        v2 = v2
    } else {
        var v25 int32 = (v24 % (((^(v24))) | 1))
        if ((!(!false)) || (32767 <= -241)) {
            var v26 int = 4294967295
            _ = v26
        }
        if ((((1.5 - float32(float64(1.5))) >= float32(0.0)) || false) || (!(("hello" > "a") && true))) {
            var v27 string = "\n\t"
            _ = v27
        } else {
            v28 := []int{2147483647, 0, 9223372036854775807}
            var v29 uint32 = v2
            _ = v28
            _ = v29
        }
        _ = v25
    }
    v30 := GenS5{F0: 0, F1: (v24 << ((438) & 63)), F2: 65535}
    for fi31 := 0; fi31 < 18; fi31++ {
        var v32 uint32 = v2
        _ = fi31
        _ = v32
    }
    switch 4294967295 {
    case -102:
        var v33 int8 = -49
        _ = v33
    default:
        v34 := GenS6{F0: v24, F1: ("\n\t" + ""), F2: 127}
        _ = v34
    }
    if true {
        if ((42 >= genFn0(1)) || (!("hello" < "\n\t"))) {
            v1 = v1
            for fi35 := 0; fi35 < 19; fi35++ {
                for fi36 := 0; fi36 < 19; fi36++ {
                    fmt.Println(fi35)
                    _ = fi36
                }
                _ = fi35
            }
        }
    }
    v37 := []int{-1}
    if (555 > 42) {
        if false {
            switch ("hello" + ("\n\t" + (("foo bar" + "\n\t") + "foo bar"))) {
            case "a":
                fmt.Println(v2, v2, v24)
            default:
                var v38 float32 = (((-2.25 - 1e10) - 3.14159) / (0.0 - 0.1))
                _ = v38
            }
            if (9223372036854775807 <= 9223372036854775807) {
                if len(v37) > 0 {
                    v37[0] = 0
                }
            } else {
                v37 = append(v37, 590)
                v37 = append(v37, -52)
            }
        } else {
            v24 -= v1
            switch "" {
            case "foo bar":
                if len(v37) > 0 {
                    v37[0] = 128
                }
            default:
                v1 -= (v1 % ((0) | 1))
            }
        }
        var v39 int = 0
        _ = v39
    } else {
        v37 = append(v37, 0)
        v40 := GenS1{F0: -1, F1: 2}
        var v41 int32 = v24
        _ = v40
        _ = v41
    }
    v3 = v3
    if true {
        if len(v37) >= 2 {
            v37 = v37[1:len(v37)]
        }
    }
    v24 -= v24
    var v42 float32 = float32((1.5 + (v3 - 3.14159)))
    if (!(!(990 > 65535))) {
        fmt.Println(v3, v24, v3)
        var v43 int16 = 1000
        _ = v43
    }
    v37 = append(v37, -326)
    v2 -= (^(v2))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v37:", len(v37), v37)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v30
    _ = v37
    _ = v42
}
