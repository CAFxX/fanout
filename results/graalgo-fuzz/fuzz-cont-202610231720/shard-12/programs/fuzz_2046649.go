package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int8; F1 uint32 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int32; F1 int; F2 int64 }
type GenS6 struct { F0 uint64; F1 uint }
func genFn0() int16 {
    v7 := GenS1{F0: 1}
    for fi8 := 0; fi8 < 3; fi8++ {
        for fi9 := 0; fi9 < 2; fi9++ {
            switch "\n\t" {
            case "\n\t":
                if (false || ("" < "hello")) {
                    var v10 int = (fi8 ^ (fi8 << ((515) & 63)))
                    _ = v10
                }
                var v11 float32 = 0.0
                _ = v11
            case "":
                var v12 float64 = (1e10 * (0.0 - ((-2.25 / 1.0) - ((0.0 * 0.0) * 1.5))))
                _ = v12
            case "foo bar":
                v13 := []int{2147483647}
                _ = v13
            default:
                var v14 int64 = 0
                _ = v14
            }
            _ = fi9
        }
        switch "foo bar" {
        case "a":
            v15 := GenS1{F0: 164}
            _ = v15
        case "":
            v16 := []int{fi8, (fi8 - 720)}
            _ = v16
        case "foo bar":
            var v17 uint16 = 46
            _ = v17
        default:
            var v18 float64 = 0.1
            _ = v18
        }
        _ = fi8
    }
    if false {
        var v19 uint32 = 166
        if true {
            if (311 >= 128) {
                v20 := []int{-1, -1}
                switch ("\n\t" + "hello") {
                case "hello":
                    v19 = 4294967295
                case "":
                    v19 += 65535
                case "foo bar":
                    v19 = 958
                }
                v21 := GenS1{F0: 58}
                _ = v20
                _ = v21
            }
            if (965 < 294) {
                v19 += 65535
                v22 := []int{32768, -9223372036854775808, -592, 4294967295}
                var v23 uint32 = (^(v19))
                _ = v22
                _ = v23
            } else {
                if false {
                    var v24 float32 = ((((3.14159 * 3.14159) - 1e10) * -2.25) + -2.25)
                    var v25 bool = ((2 <= 2) || false)
                    _ = v24
                    _ = v25
                } else {
                    v26 := GenS2{F0: 100, F1: v19}
                    v19 -= v19
                    v27 := GenS3{F0: 42}
                    _ = v26
                    _ = v27
                }
            }
        }
        _ = v19
    } else {
        var v28 bool = false
        _ = v28
    }
    var v29 float32 = (1.5 * ((0.0 * 3.14159) - (1e10 / -2.25)))
    _ = v29
    _ = v7
    return 1
}
func genFn1(p0 uint32) int32 {
    for fi30 := 0; fi30 < 13; fi30++ {
        for fi31 := 0; fi31 < 11; fi31++ {
            var v32 int16 = -1
            switch "" {
            case "a":
                p0 = p0
            case "":
                v33 := []int{388, (fi30 | 32767), fi30}
                _ = v33
            default:
                v34 := GenS4{F0: 6}
                _ = v34
            }
            _ = fi31
            _ = v32
        }
        _ = fi30
    }
    v35 := GenS5{F0: 42, F1: 32767, F2: 1}
    for fi36 := 0; fi36 < 10; fi36++ {
        p0 += (p0 + 2147483648)
        v37 := GenS4{F0: 21}
        v38 := GenS6{F0: 18446744073709551615, F1: uint(p0)}
        _ = fi36
        _ = v37
        _ = v38
    }
    var v39 float64 = 1e10
    _ = p0
    _ = v35
    _ = v39
    return 0
}

func main() {
    var v1 int8 = 100
    v2 := make(map[uint32]uint32)
    v3 := make(map[string]int)
    v4 := make(map[string]string)
    var v5 int32 = -510
    var v6 int64 = 113
    fmt.Println(v5)
    v5 += (v5 >> ((943) & 63))
    var v40 int = -2147483648
    for fi41 := 0; fi41 < 13; fi41++ {
        v42 := GenS4{F0: (v1 % ((int8(v6)) | 1))}
        var v43 int16 = -32768
        _ = fi41
        _ = v42
        _ = v43
    }
    v1 -= int8(v40)
    fmt.Println(v1)
    for fi44 := 0; fi44 < 2; fi44++ {
        v45 := GenS5{F0: (v5 << ((668) & 63)), F1: fi44, F2: 42}
        v46 := []int{(-(fi44))}
        v4["a"] = "hello"
        _ = fi44
        _ = v45
        _ = v46
    }
    var v47 uint16 = 1
    v48 := GenS6{F0: 18446744073709551615, F1: 347}
    if false {
        var v49 int = (v40 * 4294967295)
        v5 += v5
        _ = v49
    } else {
        var v50 uint64 = 2
        v6 -= 269
        var v51 int32 = (v5 & (v5 / ((genFn1(1)) | 1)))
        _ = v50
        _ = v51
    }
    var v52 uint16 = v47
    v53 := []int{(-(v40))}
    if (!false) {
        v54 := GenS6{F0: 18446744073709551615, F1: 441}
        _ = v54
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("v52:", v52)
    fmt.Println("slicelen v53:", len(v53), v53)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v47
    _ = v48
    _ = v5
    _ = v52
    _ = v53
    _ = v6
}
