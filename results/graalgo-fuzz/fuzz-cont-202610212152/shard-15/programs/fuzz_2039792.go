package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16 }
type GenS2 struct { F0 string; F1 uint32; F2 bool }
type GenS3 struct { F0 int8; F1 uint32 }
type GenS4 struct { F0 uint32; F1 uint8; F2 bool }
type GenS5 struct { F0 uint64; F1 int32; F2 bool }
func genFn0() uint64 {
    var v5 int8 = 122
    v5 -= 1
    v6 := GenS1{F0: (false && true), F1: -699}
    _ = v5
    _ = v6
    return 42
}
func genFn1(p0 int64, p1 uint8, p2 uint64) int8 {
    for fi7 := 0; fi7 < 9; fi7++ {
        var v8 uint64 = (p2 ^ p2)
        v9 := []int{fi7}
        _ = fi7
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 19; fi10++ {
        fmt.Println(p0, p0, fi10)
        _ = fi10
    }
    var v11 uint16 = 0
    fmt.Println(p2, p0, p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return 1
}

func main() {
    var v1 int = 408
    var v2 bool = false
    var v3 int = (-(v1))
    var v4 string = ("hello" + ("" + "hello"))
    v1 -= -208
    switch v4 {
    case "\n\t":
        v12 := []int{v1, 32768, v1, -9223372036854775808}
        _ = v12
    case "a":
        v13 := GenS2{F0: v4, F1: 2147483647, F2: v2}
        _ = v13
    default:
        v14 := GenS3{F0: -1, F1: 1}
        _ = v14
    }
    v4 = v4
    switch 42 {
    case -861:
        var v15 uint64 = 18446744073709551615
        _ = v15
    case -2147483648:
        v16 := []int{v3}
        _ = v16
    case 32768:
        v3 -= -9223372036854775808
    default:
        switch (v4 + v4) {
        case "":
            if true {
                if false {
                    if (18446744073709551615 < 0) {
                        var v17 uint64 = 116
                        _ = v17
                    }
                }
                for fi18 := 0; fi18 < 14; fi18++ {
                    if v2 {
                        switch (^(v1)) {
                        case -321:
                            v1 += (fi18 << ((169) & 63))
                        case 4294967295:
                            var v19 bool = (!v2)
                            _ = v19
                        case -2147483648:
                            switch (v1 ^ fi18) {
                            case 0:
                                for fi20 := 0; fi20 < 13; fi20++ {
                                    v21 := []int{-1, (-(fi20)), v1}
                                    var v22 float32 = (-2.25 / 0.1)
                                    _ = fi20
                                    _ = v21
                                    _ = v22
                                }
                            case 208:
                                v2 = false
                            default:
                                for fi23 := 0; fi23 < 11; fi23++ {
                                    v24 := []int{}
                                    v25 := GenS4{F0: 42, F1: 255, F2: ((!v2) && (!(!true)))}
                                    var v26 uint16 = 42
                                    _ = fi23
                                    _ = v24
                                    _ = v25
                                    _ = v26
                                }
                            }
                        }
                    }
                    var v27 int32 = 42
                    var v28 bool = ("\n\t" >= "\n\t")
                    _ = fi18
                    _ = v27
                    _ = v28
                }
            }
        case "\n\t":
            var v29 int = (v1 / ((v1) | 1))
            _ = v29
        }
    }
    var v30 float32 = 1e10
    fmt.Println(v30, v2)
    if (v2 || v2) {
        var v31 uint = 1
        var v32 uint16 = 0
        var v33 string = v4
        _ = v31
        _ = v32
        _ = v33
    } else {
        if (v3 <= v1) {
            v1 -= v1
            if (3.14159 == 1.5) {
                fmt.Println(v4, v3)
                v4 = ("a" + ((((v4 + v4) + v4) + "\n\t") + v4))
                for fi34 := 0; fi34 < 4; fi34++ {
                    for fi35 := 0; fi35 < 6; fi35++ {
                        fmt.Println(fi35)
                        _ = fi35
                    }
                    switch (v4 + v4) {
                    case "\n\t":
                        if (v30 != 0.0) {
                            if v2 {
                                v36 := GenS2{F0: v4, F1: 65535, F2: (-32768 >= 196)}
                                _ = v36
                            }
                            for fi37 := 0; fi37 < 5; fi37++ {
                                var v38 float32 = 1.5
                                v30 = 0.1
                                _ = fi37
                                _ = v38
                            }
                            var v39 float64 = (((0.0 - -2.25) - -2.25) * 0.1)
                            _ = v39
                        } else {
                            v40 := []int{fi34, fi34, v3, (v1 - fi34)}
                            _ = v40
                        }
                    default:
                        var v41 int32 = -2147483648
                        _ = v41
                    }
                    _ = fi34
                }
            } else {
                if v2 {
                    var v42 int32 = 2147483647
                    v42 += (v42 ^ v42)
                    v3 += 1
                    _ = v42
                } else {
                    v43 := GenS5{F0: 4294967295, F1: -541, F2: v2}
                    v1 = v1
                    _ = v43
                }
            }
            v2 = true
        }
        for fi44 := 0; fi44 < 10; fi44++ {
            v30 -= v30
            v45 := []int{(fi44 | 138), v3}
            _ = fi44
            _ = v45
        }
        v4 = v4
    }
    v30 = v30
    var v46 int16 = 1
    var v47 int16 = v46
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v46
    _ = v47
}
