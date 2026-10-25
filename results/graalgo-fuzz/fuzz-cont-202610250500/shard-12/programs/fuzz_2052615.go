package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int8; F1 string; F2 int32 }
type GenS4 struct { F0 uint; F1 uint8; F2 uint32 }
func genFn0(p0 string) uint {
    for fi4 := 0; fi4 < 20; fi4++ {
        v5 := GenS1{F0: 2, F1: 0}
        _ = fi4
        _ = v5
    }
    v6 := []int{-1}
    _ = p0
    _ = v6
    return 0
}

func main() {
    v1 := make(map[int]uint32)
    var v2 uint8 = 255
    var v3 uint = 651
    v3 -= 0
    if false {
        delete(v1, -2147483648)
        var v7 int16 = -952
        var v8 string = ("" + "hello")
        _ = v7
        _ = v8
    }
    if (false || true) {
        fmt.Println(v3)
        for fi9 := 0; fi9 < 20; fi9++ {
            v2 = 68
            v2 = v2
            v2 = 219
            _ = fi9
        }
        v3 = (v3 / (((v3 + genFn0((("a" + "foo bar") + "")))) | 1))
    } else {
        v10 := GenS2{F0: (v2 % ((v2) | 1))}
        _ = v10
    }
    v3 = 4294967295
    fmt.Println(v2)
    delete(v1, 128)
    var v11 int = 4294967295
    var v12 int64 = -9223372036854775808
    if true {
        var v13 float32 = (float32(0.1) / (0.0 + 1e10))
        _ = v13
    } else {
        for fi14 := 0; fi14 < 12; fi14++ {
            v3 = (v3 << (((v3 & v3)) & 63))
            fmt.Println(v12)
            v11 -= (fi14 >> (((v3 - 255)) & 63))
            _ = fi14
        }
        var v15 int8 = -1
        _ = v15
    }
    for fi16 := 0; fi16 < 19; fi16++ {
        v17 := GenS2{F0: (v2 >> ((genFn0((("a" + "a") + "foo bar"))) & 63))}
        var v18 string = ""
        var v19 bool = false
        _ = fi16
        _ = v17
        _ = v18
        _ = v19
    }
    switch v11 {
    case 42:
        v11 -= (v11 << ((255) & 63))
    default:
        for fi20 := 0; fi20 < 6; fi20++ {
            switch ("hello" + ("\n\t" + "")) {
            case "hello":
                for fi21 := 0; fi21 < 9; fi21++ {
                    var v22 int32 = 1
                    v1[fi21] = 65535
                    var v23 int64 = (v12 | (v12 + (^(v12))))
                    var v24 float32 = 0.0
                    v24 = v24
                    _ = fi21
                    _ = v22
                    _ = v23
                    _ = v24
                }
            default:
                v25 := GenS3{F0: 82, F1: "foo bar", F2: 1}
                _ = v25
            }
            for fi26 := 0; fi26 < 16; fi26++ {
                v12 += v12
                v11 = v11
                _ = fi26
            }
            _ = fi20
        }
    }
    switch 42 {
    case 946:
        fmt.Println(v2, v11, v3)
    case 4294967295:
        if (!(v2 != 127)) {
            for fi27 := 0; fi27 < 20; fi27++ {
                v28 := []int{127}
                for fi29 := 0; fi29 < 12; fi29++ {
                    var v30 int = fi29
                    _ = fi29
                    _ = v30
                }
                _ = fi27
                _ = v28
            }
        } else {
            for fi31 := 0; fi31 < 8; fi31++ {
                for fi32 := 0; fi32 < 4; fi32++ {
                    v2 = (v2 ^ v2)
                    fmt.Println(v3, fi32)
                    switch (fi32 * (fi32 & (fi31 + 0))) {
                    case -616:
                        var v33 int8 = 100
                        _ = v33
                    case 1:
                        switch fi32 {
                        case 2147483647:
                            v3 = 255
                        default:
                            v2 += (v2 * 82)
                        }
                    default:
                        if false {
                            v34 := []int{1, (fi31 * int(v12))}
                            if (v12 <= v12) {
                                var v35 int8 = 127
                                fmt.Println(v3)
                                _ = v35
                            } else {
                                fmt.Println(fi31)
                                if (true || true) {
                                    delete(v1, (fi32 << ((v3) & 63)))
                                    var v36 bool = (true || true)
                                    fmt.Println(v2, fi32, v3)
                                    _ = v36
                                }
                            }
                            _ = v34
                        }
                    }
                    fmt.Println(v11, v3)
                    _ = fi32
                }
                if (!true) {
                    v2 = v2
                    for fi37 := 0; fi37 < 19; fi37++ {
                        var v38 uint16 = 255
                        switch fi31 {
                        case -1:
                            if (!(!(v12 != (v12 % ((v12) | 1))))) {
                                delete(v1, (-(fi31)))
                            }
                        }
                        _ = fi37
                        _ = v38
                    }
                } else {
                    if true {
                        switch ("" + ((("foo bar" + "") + "hello") + "")) {
                        case "foo bar":
                            v39 := GenS1{F0: 0, F1: -703}
                            _ = v39
                        case "a":
                            v2 = v2
                        default:
                            var v40 uint8 = 127
                            _ = v40
                        }
                    } else {
                        if true {
                            v41 := []int{v11, (v11 - (v11 ^ (fi31 * fi31))), v11}
                            _ = v41
                        }
                        var v42 uint32 = 794
                        var v43 float64 = 1e10
                        _ = v42
                        _ = v43
                    }
                    v3 = (v3 / ((v3) | 1))
                }
                var v44 uint8 = v2
                _ = fi31
                _ = v44
            }
            delete(v1, v11)
        }
    }
    var v45 uint16 = 157
    v46 := GenS1{F0: 204, F1: 951}
    fmt.Println(v12, v12)
    for fi47 := 0; fi47 < 17; fi47++ {
        v48 := GenS2{F0: v2}
        _ = fi47
        _ = v48
    }
    v49 := GenS4{F0: genFn0(("" + "")), F1: v2, F2: 2147483647}
    switch 0 {
    case -835:
        fmt.Println(v2, v45, v12)
    case 127:
        var v50 uint8 = 2
        _ = v50
    case -503:
        v51 := []int{}
        _ = v51
    default:
        var v52 int16 = 1000
        _ = v52
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v45
    _ = v46
    _ = v49
}
