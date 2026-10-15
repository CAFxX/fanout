package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32; F2 bool }
type GenS2 struct { F0 int16; F1 int32 }
type GenS3 struct { F0 uint32; F1 uint32; F2 int32 }
type GenS4 struct { F0 int64; F1 uint8 }
func genFn0(p0 int, p1 string, p2 bool) uint16 {
    fmt.Println(p2, p2)
    v5 := []int{9223372036854775807, (p0 & p0), 519}
    p0 = (p0 - (^(p0)))
    p1 = ("" + "")
    var v6 float64 = 1e10
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 0
}
func genFn1() int32 {
    if false {
        v7 := GenS1{F0: -9223372036854775808, F1: 778, F2: false}
        _ = v7
    } else {
        var v8 int8 = -128
        var v9 uint8 = 130
        _ = v8
        _ = v9
    }
    var v10 uint = 18446744073709551615
    _ = v10
    return 1
}

func main() {
    v1 := []string{"a"}
    var v2 uint8 = 2
    var v3 bool = false
    var v4 uint32 = 2147483647
    switch -522 {
    case 128:
        var v11 uint = 0
        _ = v11
    case 32768:
        var v12 float32 = 1.5
        _ = v12
    case 251:
        if v3 {
            fmt.Println(v4, v4, v4)
            if true {
                var v13 int16 = 713
                var v14 int8 = -1
                _ = v13
                _ = v14
            } else {
                switch -1 {
                case 32767:
                    var v15 string = "\n\t"
                    _ = v15
                case 32768:
                    if ((0.0 / 0.1) < 3.14159) {
                        var v16 int = int(v2)
                        _ = v16
                    } else {
                        if (0 < 32767) {
                            v3 = false
                            var v17 int = 128
                            _ = v17
                        }
                    }
                case 9223372036854775807:
                    var v18 uint16 = 42
                    _ = v18
                default:
                    v19 := GenS1{F0: -411, F1: 2147483648, F2: (-194 >= 317)}
                    _ = v19
                }
            }
            v4 = (v4 << ((42) & 63))
        } else {
            v20 := GenS2{F0: 585, F1: -1}
            v1 = append(v1, "\n\t")
            v4 += v4
            _ = v20
        }
    default:
        v3 = (0.1 < (1.5 / -2.25))
    }
    var v21 int16 = 32767
    v2 = v2
    v22 := GenS3{F0: 42, F1: (^(v4)), F2: genFn1()}
    var v23 uint16 = 0
    var v24 int = -2147483648
    var v25 uint32 = 4294967295
    for fi26 := 0; fi26 < 18; fi26++ {
        switch fi26 {
        case 32768:
            if (546 <= 42) {
                var v27 int16 = (v21 >> ((480) & 63))
                _ = v27
            } else {
                v28 := GenS3{F0: 2, F1: (v4 % (((v4 - v25)) | 1)), F2: genFn1()}
                _ = v28
            }
        }
        _ = fi26
    }
    v29 := GenS2{F0: v21, F1: 30}
    var v30 int16 = 1
    fmt.Println(v2, v30, v4)
    v3 = v3
    if (-2.25 == ((float64((3.14159 / 0.1)) - 1.5) - ((float64(0.1) + float64(0.1)) / 1.5))) {
        for fi31 := 0; fi31 < 10; fi31++ {
            if v3 {
                v25 += v25
                if len(v1) > 0 {
                    v1[0] = "\n\t"
                }
                for fi32 := 0; fi32 < 9; fi32++ {
                    v33 := GenS4{F0: 1, F1: (v2 * v2)}
                    fmt.Println(fi31)
                    _ = fi32
                    _ = v33
                }
            } else {
                var v34 float32 = (-2.25 * 0.1)
                if v3 {
                    switch "foo bar" {
                    case "\n\t":
                        var v35 int32 = 1
                        _ = v35
                    default:
                        v21 = (v30 % ((v30) | 1))
                    }
                }
                _ = v34
            }
            switch v24 {
            case 32767:
                var v36 int64 = 42
                _ = v36
            case -9223372036854775808:
                var v37 string = "hello"
                _ = v37
            case 4294967295:
                var v38 string = ""
                _ = v38
            default:
                var v39 int8 = 96
                _ = v39
            }
            var v40 float64 = float64(-2.25)
            _ = fi31
            _ = v40
        }
    } else {
        if (v3 && v3) {
            var v41 uint = 816
            var v42 int = (-(v24))
            _ = v41
            _ = v42
        }
        var v43 int16 = (v21 / (((v21 << ((18446744073709551615) & 63))) | 1))
        _ = v43
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println(v3, v4, v24)
    fmt.Println(v24)
    v25 += (v25 / (((^(v25))) | 1))
    _ = len(v1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v29
    _ = v3
    _ = v30
    _ = v4
}
