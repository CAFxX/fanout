package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int16; F1 uint8 }
type GenS3 struct { F0 uint8; F1 int32; F2 int16 }
type GenS4 struct { F0 bool }
func genFn0() uint16 {
    var v6 float64 = 3.14159
    v7 := GenS1{F0: ""}
    fmt.Println(v6)
    var v8 float64 = v6
    _ = v6
    _ = v7
    _ = v8
    return 88
}
func genFn1(p0 uint8) int64 {
    for fi9 := 0; fi9 < 6; fi9++ {
        v10 := []int{(fi9 | (fi9 / (((fi9 * fi9)) | 1))), fi9}
        fmt.Println(p0)
        _ = fi9
        _ = v10
    }
    p0 = p0
    v11 := []int{-2147483648, int(p0), -637}
    fmt.Println(p0)
    _ = p0
    _ = v11
    return 284
}

func main() {
    var v1 bool = true
    var v2 uint16 = 0
    v3 := []uint32{42, uint32(v2)}
    var v4 uint16 = v2
    v5 := []string{("foo bar" + "\n\t")}
    for fi12 := 0; fi12 < 17; fi12++ {
        v2 -= (^(v4))
        _ = fi12
    }
    switch "" {
    case "a":
        var v13 uint16 = genFn0()
        _ = v13
    case "hello":
        fmt.Println(v1, v1, v1)
    case "foo bar":
        var v14 uint64 = 169
        _ = v14
    }
    if (!v1) {
        switch "hello" {
        case "foo bar":
            v4 += (v4 ^ v4)
        default:
            v4 -= (v2 >> ((18446744073709551615) & 63))
        }
        var v15 uint8 = 0
        var v16 float32 = float32(0.0)
        _ = v15
        _ = v16
    }
    var v17 uint = 0
    switch 42 {
    case 0:
        for fi18 := 0; fi18 < 15; fi18++ {
            if v1 {
                var v19 float32 = -2.25
                _ = v19
            } else {
                v2 = genFn0()
            }
            if ((v2 % (((^(v4))) | 1)) >= (v4 >> (((v17 / (((v17 - (v17 / ((v17) | 1)))) | 1))) & 63))) {
                v17 -= 1
                var v20 uint = uint(v2)
                _ = v20
            } else {
                fmt.Println(v1, v17, v2)
            }
            _ = fi18
        }
    case -2147483648:
        for fi21 := 0; fi21 < 19; fi21++ {
            _ = len(v5)
            v17 -= (v17 - (^(v17)))
            for fi22 := 0; fi22 < 11; fi22++ {
                v1 = false
                _ = fi22
            }
            _ = fi21
        }
    }
    for fi23 := 0; fi23 < 18; fi23++ {
        var v24 uint16 = 65535
        _ = fi23
        _ = v24
    }
    v3 = append(v3, 293)
    if (!v1) {
        var v25 uint8 = 1
        switch "hello" {
        case "":
            if (v1 && v1) {
                for fi26 := 0; fi26 < 6; fi26++ {
                    v27 := GenS2{F0: 32767, F1: 128}
                    var v28 int64 = 9223372036854775807
                    _ = fi26
                    _ = v27
                    _ = v28
                }
                v3 = append(v3, 578)
            } else {
                for fi29 := 0; fi29 < 13; fi29++ {
                    for fi30 := 0; fi30 < 2; fi30++ {
                        var v31 float32 = 1.5
                        var v32 int8 = 1
                        _ = fi30
                        _ = v31
                        _ = v32
                    }
                    _ = fi29
                }
                v4 = (v2 | (v4 - v2))
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            }
        default:
            if true {
                if v1 {
                    fmt.Println(v1, v17, v4)
                    if (true || v1) {
                        v1 = true
                    }
                } else {
                    v4 = (^(v4))
                    var v33 int64 = genFn1(v25)
                    _ = v33
                }
            } else {
                switch ("\n\t" + "foo bar") {
                case "hello":
                    v3 = append(v3, 65535)
                case "":
                    fmt.Println(v2, v17)
                }
                fmt.Println(v4)
            }
        }
        _ = v25
    } else {
        v1 = (((-1 < 100) || ((-9223372036854775808 <= 42) || true)) || v1)
        var v34 uint = uint(v4)
        _ = v34
    }
    if (!(2 != v17)) {
        var v35 int16 = 119
        _ = v35
    }
    if (true && (!v1)) {
        var v36 int8 = 6
        _ = v36
    }
    for fi37 := 0; fi37 < 12; fi37++ {
        for fi38 := 0; fi38 < 3; fi38++ {
            if (0 < 0) {
                if false {
                    v5 = append(v5, (("" + ("foo bar" + (("hello" + "\n\t") + ""))) + "\n\t"))
                }
                var v39 int8 = 127
                _ = v39
            }
            var v40 uint8 = 88
            _ = fi38
            _ = v40
        }
        fmt.Println(v17, v17, v4)
        v2 += v4
        _ = fi37
    }
    for fi41 := 0; fi41 < 9; fi41++ {
        v42 := GenS3{F0: 46, F1: 1, F2: -1}
        v1 = true
        _ = len(v3)
        _ = fi41
        _ = v42
    }
    var v43 int = 42
    for fi44 := 0; fi44 < 15; fi44++ {
        for fi45 := 0; fi45 < 6; fi45++ {
            fmt.Println(v4)
            v46 := GenS4{F0: v1}
            _ = fi45
            _ = v46
        }
        _ = fi44
    }
    var v47 int64 = -238
    if (v1 || (((^(v4)) < v2) && (-2.25 != -2.25))) {
        v47 += (v47 % (((v47 / (((^(v47))) | 1))) | 1))
        for fi48 := 0; fi48 < 5; fi48++ {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi48
        }
        fmt.Println(v1)
    } else {
        v43 = v43
        fmt.Println(v4, v4)
    }
    v43 = int(v4)
    fmt.Println(v43, v4)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v47:", v47)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v47
    _ = v5
}
