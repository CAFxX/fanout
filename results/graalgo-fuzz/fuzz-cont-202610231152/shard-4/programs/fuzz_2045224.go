package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint32; F1 int16 }
type GenS3 struct { F0 int64; F1 bool }
type GenS4 struct { F0 uint16; F1 uint32; F2 uint32 }
type GenS5 struct { F0 uint16 }
func genFn0() int32 {
    var v7 uint16 = 32767
    fmt.Println(v7)
    var v8 int64 = 42
    if (!true) {
        v7 += v7
        v9 := []int{4294967295}
        _ = v9
    } else {
        fmt.Println(v8)
    }
    _ = v7
    _ = v8
    return 0
}
func genFn1(p0 uint16, p1 bool) uint64 {
    p0 += p0
    var v10 uint = 0
    for fi11 := 0; fi11 < 9; fi11++ {
        v12 := GenS1{F0: (^(v10))}
        v13 := GenS2{F0: 742, F1: -1}
        _ = fi11
        _ = v12
        _ = v13
    }
    _ = p0
    _ = p1
    _ = v10
    return 42
}

func main() {
    var v1 uint = 0
    var v2 int64 = -946
    var v3 uint8 = 2
    var v4 float32 = float32(float64(1.5))
    var v5 uint = (v1 & 1)
    var v6 uint8 = (v3 & (v3 % (((^(v3))) | 1)))
    var v14 bool = true
    for fi15 := 0; fi15 < 5; fi15++ {
        if v14 {
            if (v4 == ((0.0 * v4) * ((float32(1e10) * -2.25) * v4))) {
                fmt.Println(v1, v14, v6)
                if v14 {
                    v16 := GenS1{F0: v1}
                    v17 := GenS1{F0: (^(v1))}
                    _ = v16
                    _ = v17
                } else {
                    var v18 int32 = 851
                    _ = v18
                }
            }
            v2 += (v2 / ((v2) | 1))
            fmt.Println(v5, v6, v3)
        } else {
            v19 := []int{2147483647, 0, (fi15 * fi15), (fi15 & fi15)}
            v5 = (v1 & v5)
            for fi20 := 0; fi20 < 1; fi20++ {
                v4 -= (((float32(float64((v4 + ((1e10 * float32(0.1)) * float32(((3.14159 / 1.0) - (float64(float32(-2.25)) / -2.25))))))) + 1e10) / 1e10) * v4)
                var v21 uint = 166
                fmt.Println(v14, v2, v3)
                _ = fi20
                _ = v21
            }
            _ = v19
        }
        v22 := GenS2{F0: 818, F1: 32767}
        switch "foo bar" {
        case "hello":
            if (false || false) {
                fmt.Println(fi15, v1, v2)
            }
        case "foo bar":
            v23 := []int{42}
            _ = v23
        default:
            if v14 {
                if v14 {
                    fmt.Println(v3, v4, v5)
                    var v24 int64 = v2
                    v25 := GenS2{F0: 0, F1: -896}
                    _ = v24
                    _ = v25
                }
                if (42 <= int16(v3)) {
                    var v26 int8 = 1
                    _ = v26
                } else {
                    if v14 {
                        if ((v4 + 0.1) > v4) {
                            if v14 {
                                for fi27 := 0; fi27 < 18; fi27++ {
                                    v28 := GenS3{F0: (v2 | 9223372036854775807), F1: ((!true) && false)}
                                    _ = fi27
                                    _ = v28
                                }
                            }
                            var v29 int32 = 42
                            _ = v29
                        }
                    }
                }
                v30 := []int{}
                _ = v30
            } else {
                switch 4294967295 {
                case 128:
                    switch "a" {
                    case "foo bar":
                        var v31 int64 = 1
                        _ = v31
                    case "":
                        v32 := []int{32768}
                        _ = v32
                    case "hello":
                        v1 = v1
                    default:
                        fmt.Println(v14, v3)
                    }
                case 9223372036854775807:
                    switch fi15 {
                    case -747:
                        for fi33 := 0; fi33 < 10; fi33++ {
                            var v34 int64 = (v2 | v2)
                            var v35 string = ""
                            _ = fi33
                            _ = v34
                            _ = v35
                        }
                    case -469:
                        v14 = ((v2 <= (v2 | 640)) && v14)
                    default:
                        v36 := GenS1{F0: v5}
                        _ = v36
                    }
                case 32768:
                    var v37 int32 = 1
                    _ = v37
                default:
                    var v38 string = "foo bar"
                    _ = v38
                }
                var v39 string = ("hello" + ((("a" + "a") + ("hello" + "hello")) + "foo bar"))
                for fi40 := 0; fi40 < 15; fi40++ {
                    v41 := []int{fi15, fi15}
                    _ = fi40
                    _ = v41
                }
                _ = v39
            }
        }
        _ = fi15
        _ = v22
    }
    v42 := GenS4{F0: 42, F1: 42, F2: 65535}
    v43 := []int{32767, 32768}
    v1 = (v5 | 0)
    v43 = append(v43, 32767)
    var v44 int64 = v2
    for fi45 := 0; fi45 < 15; fi45++ {
        v46 := GenS4{F0: 32767, F1: 2147483648, F2: 4294967295}
        if (v4 != v4) {
            v47 := GenS4{F0: 255, F1: 447, F2: 422}
            _ = v47
        }
        _ = fi45
        _ = v46
    }
    var v48 float32 = v4
    if (!(v14 || v14)) {
        v49 := GenS5{F0: 65535}
        switch -1 {
        case 1:
            v14 = v14
        case -1:
            for fi50 := 0; fi50 < 9; fi50++ {
                v5 = (v5 + v1)
                fmt.Println(v6, v48)
                var v51 uint8 = 207
                _ = fi50
                _ = v51
            }
        case 154:
            if len(v43) > 0 {
                v43[0] = 127
            }
        }
        if v14 {
            v44 = -9223372036854775808
            if (v14 && ((v6 / ((v3) | 1)) < (v6 >> (((v5 % ((746) | 1))) & 63)))) {
                if false {
                    if false {
                        switch 42 {
                        case 127:
                            fmt.Println(v3)
                        default:
                            v1 = 160
                        }
                        fmt.Println(v6, v6)
                    }
                    var v52 int64 = (v2 % (((v2 % ((v44) | 1))) | 1))
                    _ = v52
                }
            } else {
                _ = len(v43)
                v3 = 255
                var v53 float32 = 1e10
                _ = v53
            }
        }
        _ = v49
    } else {
        if len(v43) >= 2 {
            v43 = v43[1:len(v43)]
        }
    }
    v54 := GenS4{F0: 312, F1: 439, F2: 2147483648}
    var v55 string = ("\n\t" + "a")
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v43:", len(v43), v43)
    fmt.Println("v44:", v44)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    fmt.Println("v55:", v55)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v44
    _ = v48
    _ = v5
    _ = v54
    _ = v55
    _ = v6
}
