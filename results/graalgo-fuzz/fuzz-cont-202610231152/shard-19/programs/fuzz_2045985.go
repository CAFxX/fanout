package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint8; F1 uint16 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint; F1 int16; F2 uint32 }
func genFn0(p0 uint8) int64 {
    v4 := []int{804, 32768}
    v5 := GenS1{F0: (p0 << ((255) & 63))}
    fmt.Println(p0)
    _ = p0
    _ = v4
    _ = v5
    return 0
}
func genFn1(p0 int32, p1 bool) uint8 {
    for fi6 := 0; fi6 < 10; fi6++ {
        if (42 > 2147483648) {
            fmt.Println(p1, p0, fi6)
            v7 := []int{fi6, fi6, fi6}
            _ = v7
        } else {
            p1 = p1
        }
        v8 := GenS2{F0: 164, F1: 788}
        v9 := []int{fi6, fi6, (fi6 >> ((42) & 63)), fi6}
        _ = fi6
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 17; fi10++ {
        p1 = (-2.25 != 0.1)
        p0 = -331
        _ = fi10
    }
    p1 = p1
    p1 = (true && (!((p0 / ((p0) | 1)) > p0)))
    p1 = p1
    _ = p0
    _ = p1
    return 0
}

func main() {
    var v1 int32 = -2147483648
    var v2 int = 42
    var v3 float64 = (1e10 + 3.14159)
    switch v2 {
    case 9223372036854775807:
        v11 := []int{-9223372036854775808, -129, v2, v2}
        _ = v11
    }
    v1 = 0
    if (1 == 945) {
        switch -1 {
        case -9223372036854775808:
            fmt.Println(v3, v3, v3)
        case -977:
            v12 := []int{}
            _ = v12
        }
        var v13 float64 = v3
        _ = v13
    } else {
        var v14 float32 = (3.14159 / 1.0)
        _ = v14
    }
    switch ("" + "a") {
    case "foo bar":
        var v15 int16 = 0
        _ = v15
    case "a":
        switch (("hello" + "a") + ("foo bar" + "hello")) {
        case "\n\t":
            var v16 int64 = 0
            _ = v16
        case "hello":
            switch "foo bar" {
            case "":
                v17 := GenS2{F0: 0, F1: 990}
                _ = v17
            default:
                if true {
                    v18 := []int{v2, (v2 | v2), -439}
                    _ = v18
                }
            }
        default:
            v2 += v2
        }
    }
    if (0.1 < 0.0) {
        var v19 float32 = 0.1
        v20 := []int{v2, (v2 | (v2 - (v2 * 510))), v2}
        fmt.Println(v2)
        _ = v19
        _ = v20
    } else {
        var v21 int8 = 42
        v22 := []int{v2, (v2 + v2), (v2 / ((128) | 1))}
        v23 := GenS3{F0: genFn1((v1 ^ v1), true)}
        _ = v21
        _ = v22
        _ = v23
    }
    v24 := []int{(v2 % ((v2) | 1)), -648}
    switch (v2 ^ v2) {
    case 127:
        fmt.Println(v2)
    }
    for fi25 := 0; fi25 < 12; fi25++ {
        v2 -= (v2 << ((65535) & 63))
        var v26 uint32 = 278
        if (true && ((-1 != 0) && (false && true))) {
            if (v26 <= v26) {
                v1 -= (v1 - v1)
                if false {
                    v24 = append(v24, fi25)
                }
                var v27 uint = 42
                _ = v27
            }
        }
        _ = fi25
        _ = v26
    }
    if false {
        if (1 >= 1) {
            var v28 float32 = -2.25
            switch "hello" {
            case "foo bar":
                var v29 uint32 = 0
                _ = v29
            default:
                v3 += v3
            }
            _ = v28
        }
        v2 -= (-(v2))
    }
    v30 := GenS1{F0: 42}
    fmt.Println(v1)
    if (true && (!true)) {
        if true {
            if (!(65535 == 255)) {
                if ((((!(!false)) || true) && true) && (32767 <= 0)) {
                    for fi31 := 0; fi31 < 9; fi31++ {
                        fmt.Println(fi31)
                        if (genFn0(1) <= 42) {
                            if false {
                                switch (v2 ^ 4294967295) {
                                case 4294967295:
                                    v32 := GenS1{F0: 0}
                                    _ = v32
                                case -2147483648:
                                    v2 = v2
                                default:
                                    _ = len(v24)
                                }
                                var v33 float64 = v3
                                _ = v33
                            } else {
                                _ = len(v24)
                                _ = len(v24)
                            }
                        } else {
                            switch "a" {
                            case "\n\t":
                                if false {
                                    v2 = fi31
                                    for fi34 := 0; fi34 < 14; fi34++ {
                                        fmt.Println(fi34)
                                        var v35 uint = 65535
                                        if len(v24) > 0 {
                                            v24[0] = fi31
                                        }
                                        if true {
                                            if (true && false) {
                                                v35 = 176
                                            }
                                        }
                                        v3 = ((v3 / float64(1.5)) + -2.25)
                                        _ = fi34
                                        _ = v35
                                    }
                                }
                            default:
                                var v36 uint64 = 0
                                _ = v36
                            }
                        }
                        v37 := GenS1{F0: 255}
                        _ = fi31
                        _ = v37
                    }
                    var v38 int8 = -5
                    _ = v38
                } else {
                    switch v2 {
                    case -9223372036854775808:
                        v39 := GenS1{F0: 2}
                        _ = v39
                    default:
                        for fi40 := 0; fi40 < 7; fi40++ {
                            v2 = fi40
                            if len(v24) >= 2 {
                                v24 = v24[1:len(v24)]
                            }
                            _ = fi40
                        }
                    }
                    var v41 float32 = (1.5 - 0.1)
                    v2 = -2147483648
                    _ = v41
                }
            } else {
                v42 := GenS3{F0: 196}
                fmt.Println(v2, v3)
                _ = v42
            }
        }
        for fi43 := 0; fi43 < 19; fi43++ {
            v1 = v1
            _ = fi43
        }
        fmt.Println(v2, v3)
    } else {
        for fi44 := 0; fi44 < 6; fi44++ {
            v45 := GenS4{F0: 42, F1: -554, F2: 582}
            if (false || false) {
                var v46 bool = ((-1 != 100) || (229 >= v1))
                v47 := GenS1{F0: 128}
                v1 = (v1 % ((1) | 1))
                _ = v46
                _ = v47
            }
            if len(v24) > 0 {
                v24[0] = 97
            }
            _ = fi44
            _ = v45
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v30
}
