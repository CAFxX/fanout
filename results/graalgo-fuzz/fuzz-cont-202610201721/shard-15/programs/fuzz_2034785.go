package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64; F2 int8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int64; F1 uint32 }
type GenS4 struct { F0 uint16; F1 uint32; F2 uint8 }
type GenS5 struct { F0 uint16; F1 uint16; F2 uint8 }
type GenS6 struct { F0 int; F1 uint64 }
func genFn0(p0 bool, p1 int64) uint8 {
    p0 = p0
    v7 := []int{-1, 32767, -1, -110}
    for fi8 := 0; fi8 < 2; fi8++ {
        var v9 uint32 = 1
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v7
    return 42
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int8 = 100
    v3 := make(map[int]int)
    var v4 int8 = (v2 | 42)
    v5 := make(map[int]int)
    var v6 uint8 = 1
    for fi10 := 0; fi10 < 9; fi10++ {
        var v11 string = "foo bar"
        v12 := GenS1{F0: 65535, F1: 64, F2: (v2 - v2)}
        for fi13 := 0; fi13 < 12; fi13++ {
            switch (fi10 % (((fi10 - fi13)) | 1)) {
            case 32768:
                var v14 uint8 = v6
                v4 -= (v2 % (((v2 << ((1) & 63))) | 1))
                _ = v14
            case -1:
                if (!(-525 >= 32767)) {
                    if true {
                        switch (v11 + ((v11 + "") + "\n\t")) {
                        case "hello":
                            fmt.Println(v2, v11, v4)
                        case "\n\t":
                            v2 -= v2
                        }
                    } else {
                        v15 := GenS2{F0: true}
                        if ((!true) || (768 < 4294967295)) {
                            delete(v3, (fi13 + fi13))
                        } else {
                            fmt.Println(fi13)
                        }
                        var v16 int8 = int8(v6)
                        var v17 uint8 = genFn0(true, -1)
                        var v18 bool = ((false || true) && true)
                        _ = v15
                        _ = v16
                        _ = v17
                        _ = v18
                    }
                } else {
                    switch 32768 {
                    case 2147483647:
                        delete(v3, (fi10 | fi13))
                    case -957:
                        v2 -= v4
                    default:
                        switch (-(fi13)) {
                        case 42:
                            v2 -= (v4 ^ (v4 % (((v4 - (^(v4)))) | 1)))
                        }
                    }
                    v19 := []int{(^(fi10))}
                    if false {
                        var v20 int64 = -9223372036854775808
                        _ = v20
                    }
                    _ = v19
                }
            default:
                var v21 int16 = -218
                _ = v21
            }
            fmt.Println(fi10)
            var v22 string = ""
            _ = fi13
            _ = v22
        }
        _ = fi10
        _ = v11
        _ = v12
    }
    v6 += v6
    v23 := []int{9223372036854775807, 1}
    v24 := GenS3{F0: 776, F1: 2}
    var v25 int32 = 0
    switch "hello" {
    case "\n\t":
        var v26 uint8 = v6
        _ = v26
    case "hello":
        var v27 float64 = (float64(-2.25) + 0.1)
        _ = v27
    default:
        v6 = (v6 / ((225) | 1))
    }
    v28 := GenS4{F0: 65535, F1: 770, F2: (v6 & (^(v6)))}
    if true {
        if (!false) {
            fmt.Println(v25)
        }
    } else {
        v29 := GenS5{F0: 32767, F1: uint16(v6), F2: 237}
        v25 -= (v25 * (v25 >> ((65535) & 63)))
        var v30 int32 = (v25 >> ((2) & 63))
        _ = v29
        _ = v30
    }
    for fi31 := 0; fi31 < 17; fi31++ {
        for fi32 := 0; fi32 < 7; fi32++ {
            var v33 uint = 65535
            if false {
                var v34 float32 = (1e10 - 0.0)
                if false {
                    switch "a" {
                    case "foo bar":
                        var v35 float64 = 0.1
                        _ = v35
                    case "hello":
                        switch 128 {
                        case -2147483648:
                            var v36 int8 = -128
                            _ = v36
                        default:
                            v34 -= v34
                        }
                    }
                    if (true || true) {
                        if ((true || (!((((0.1 - ((1.5 + float64(((float32(((0.0 * 0.1) * 1e10)) - (v34 - float32(-2.25))) - (((float32(1.5) + v34) - float32(0.0)) - v34)))) - 3.14159)) * 3.14159) - float64((v34 + (v34 + v34)))) > float64(float32(-2.25))))) && (240 > 255)) {
                            var v37 string = (("a" + ("foo bar" + "a")) + "")
                            v4 -= (v4 & (v4 | (v4 * -84)))
                            if false {
                                fmt.Println(fi32, v34)
                                v33 += v33
                            }
                            _ = v37
                        }
                        v4 = (v4 >> ((v33) & 63))
                        var v38 float64 = (1.5 / (float64((float32(-2.25) / 1.0)) * 1.5))
                        var v39 bool = false
                        switch "foo bar" {
                        case "hello":
                            var v40 string = "a"
                            _ = v40
                        case "foo bar":
                            var v41 float32 = ((((-2.25 - v34) - v34) - (((v34 * 1.5) - 1e10) * (v34 * 3.14159))) / 1.0)
                            _ = v41
                        case "a":
                            fmt.Println(v39, fi31, fi32)
                        }
                        _ = v38
                        _ = v39
                    }
                } else {
                    v4 = 100
                }
                var v42 uint8 = 1
                _ = v34
                _ = v42
            }
            v6 -= v6
            _ = fi32
            _ = v33
        }
        for fi43 := 0; fi43 < 13; fi43++ {
            v44 := GenS1{F0: 4294967295, F1: 4294967295, F2: -27}
            _ = fi43
            _ = v44
        }
        _ = fi31
    }
    v45 := GenS3{F0: 42, F1: 741}
    if true {
        for fi46 := 0; fi46 < 16; fi46++ {
            v47 := GenS6{F0: fi46, F1: 2}
            for fi48 := 0; fi48 < 9; fi48++ {
                v6 = v6
                fmt.Println(v25, fi48)
                fmt.Println(v4)
                _ = fi48
            }
            _ = fi46
            _ = v47
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v28
    _ = v3
    _ = v4
    _ = v45
    _ = v5
    _ = v6
}
