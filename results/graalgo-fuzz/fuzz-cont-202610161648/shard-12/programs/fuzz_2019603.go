package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16 }
type GenS2 struct { F0 int32; F1 int; F2 uint16 }
type GenS3 struct { F0 int16; F1 uint8; F2 int64 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint8; F1 int64 }
type GenS6 struct { F0 int; F1 uint }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 uint64; F1 uint32; F2 uint32 }
type GenS9 struct { F0 int32; F1 int }
func genFn0(p0 int16) uint8 {
    p0 = 606
    v7 := []int{-9223372036854775808, 32768, 53}
    v8 := GenS1{F0: -1, F1: 1000}
    _ = p0
    _ = v7
    _ = v8
    return 42
}
func genFn1(p0 int16) int64 {
    var v9 bool = true
    for fi10 := 0; fi10 < 18; fi10++ {
        var v11 uint16 = 1
        _ = fi10
        _ = v11
    }
    v12 := GenS2{F0: 927, F1: 9223372036854775807, F2: 32767}
    var v13 float32 = float32(float64(float32(float64(0.0))))
    p0 -= (-(p0))
    _ = p0
    _ = v12
    _ = v13
    _ = v9
    return -375
}

func main() {
    var v1 bool = (0 >= 743)
    var v2 string = (("foo bar" + "a") + "a")
    var v3 uint = 926
    var v4 float64 = 1.5
    var v5 int64 = -9223372036854775808
    var v6 uint16 = 65535
    v4 += v4
    for fi14 := 0; fi14 < 12; fi14++ {
        switch (fi14 - fi14) {
        case 32767:
            for fi15 := 0; fi15 < 7; fi15++ {
                var v16 int8 = int8(v3)
                v17 := GenS3{F0: 388, F1: 202, F2: v5}
                _ = fi15
                _ = v16
                _ = v17
            }
        case -998:
            v18 := []int{2147483647, -2147483648, fi14}
            _ = v18
        default:
            v19 := []int{fi14, fi14}
            _ = v19
        }
        if ((("foo bar" != v2) && ((599 >= 1) || (v3 < (v3 - (v3 >> (((v3 / ((uint(v6)) | 1))) & 63)))))) || v1) {
            if false {
                v20 := GenS4{F0: genFn0(-1)}
                var v21 bool = false
                _ = v20
                _ = v21
            }
            if v1 {
                v22 := GenS5{F0: 222, F1: -814}
                for fi23 := 0; fi23 < 2; fi23++ {
                    var v24 int16 = -776
                    _ = fi23
                    _ = v24
                }
                if false {
                    if v1 {
                        if v1 {
                            var v25 int32 = -1
                            for fi26 := 0; fi26 < 15; fi26++ {
                                v3 -= (v3 - (^(v3)))
                                _ = fi26
                            }
                            _ = v25
                        } else {
                            for fi27 := 0; fi27 < 16; fi27++ {
                                fmt.Println(fi27, fi27)
                                fmt.Println(v2)
                                _ = fi27
                            }
                            fmt.Println(v6, v1)
                        }
                        v28 := GenS6{F0: (fi14 | (fi14 / ((-220) | 1))), F1: (v3 >> ((v3) & 63))}
                        v29 := GenS3{F0: 296, F1: genFn0(42), F2: 487}
                        _ = v28
                        _ = v29
                    }
                    if true {
                        fmt.Println(v2)
                    }
                    var v30 string = (v2 + (("hello" + v2) + "foo bar"))
                    _ = v30
                }
                _ = v22
            }
            var v31 int64 = (v5 ^ (v5 - (v5 + (v5 % ((v5) | 1)))))
            _ = v31
        }
        var v32 int8 = 90
        _ = fi14
        _ = v32
    }
    var v33 int32 = -1
    v2 = v2
    var v34 uint16 = (^(v6))
    for fi35 := 0; fi35 < 19; fi35++ {
        v36 := GenS7{F0: 722}
        v1 = (fi35 != fi35)
        var v37 int64 = 1
        _ = fi35
        _ = v36
        _ = v37
    }
    if ("\n\t" > "\n\t") {
        if v1 {
            fmt.Println(v6, v34, v6)
            v38 := []int{844, -2147483648}
            _ = v38
        }
        v33 = 0
        v33 = (v33 * (v33 & -414))
    } else {
        for fi39 := 0; fi39 < 4; fi39++ {
            var v40 int16 = 522
            _ = fi39
            _ = v40
        }
        switch 2147483647 {
        case 4294967295:
            v41 := []int{}
            _ = v41
        case -2147483648:
            for fi42 := 0; fi42 < 17; fi42++ {
                v6 = v6
                if false {
                    var v43 uint16 = v6
                    var v44 int = fi42
                    fmt.Println(v34)
                    _ = v43
                    _ = v44
                } else {
                    if (-32768 >= 0) {
                        fmt.Println(v33, fi42, v6)
                        v45 := GenS8{F0: 442, F1: 2147483648, F2: 32}
                        v46 := GenS9{F0: v33, F1: (fi42 - 0)}
                        _ = v45
                        _ = v46
                    }
                }
                v47 := []int{fi42, (fi42 & fi42), (fi42 << ((307) & 63)), fi42}
                _ = fi42
                _ = v47
            }
        case -1:
            var v48 uint32 = 4294967295
            _ = v48
        }
    }
    var v49 float32 = ((0.1 * 1.5) / 1.5)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v49
    _ = v5
    _ = v6
}
