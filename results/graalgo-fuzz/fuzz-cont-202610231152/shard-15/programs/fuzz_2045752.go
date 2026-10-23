package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64; F2 uint8 }
type GenS2 struct { F0 uint16; F1 uint16 }
type GenS3 struct { F0 uint32; F1 uint }
type GenS4 struct { F0 int8; F1 uint16 }
func genFn0() int8 {
    for fi4 := 0; fi4 < 10; fi4++ {
        for fi5 := 0; fi5 < 17; fi5++ {
            var v6 uint32 = 894
            fmt.Println(v6, fi4, v6)
            v7 := GenS1{F0: 238, F1: 18446744073709551615, F2: 2}
            _ = fi5
            _ = v6
            _ = v7
        }
        _ = fi4
    }
    if true {
        var v8 uint64 = 0
        _ = v8
    } else {
        var v9 float64 = (float64(((1.5 - 1.5) - 1e10)) / 0.1)
        v9 += 3.14159
        fmt.Println(v9)
        _ = v9
    }
    if (42 <= 1) {
        for fi10 := 0; fi10 < 12; fi10++ {
            var v11 uint = 0
            v12 := []int{1, (fi10 - fi10), fi10, (fi10 | fi10)}
            if (42 <= 42) {
                var v13 uint16 = 32767
                if (-128 > 0) {
                    fmt.Println(v13)
                }
                switch "foo bar" {
                case "":
                    v13 += (v13 / ((v13) | 1))
                case "\n\t":
                    var v14 int32 = 1
                    _ = v14
                case "a":
                    fmt.Println(fi10, fi10, v11)
                default:
                    fmt.Println(v13, v11, fi10)
                }
                _ = v13
            }
            _ = fi10
            _ = v11
            _ = v12
        }
        if (false || false) {
            v15 := []int{32768, 128, 547}
            if (false || true) {
                if (4294967295 < 1) {
                    var v16 uint64 = 673
                    if false {
                        var v17 int8 = 42
                        if true {
                            fmt.Println(v16, v17)
                            v17 = (v17 % (((v17 - (-(v17)))) | 1))
                        } else {
                            fmt.Println(v17, v16)
                        }
                        v17 = (v17 ^ 100)
                        _ = v17
                    } else {
                        var v18 int32 = -1
                        _ = v18
                    }
                    _ = v16
                } else {
                    var v19 int16 = -1
                    v19 -= v19
                    _ = v19
                }
                var v20 int32 = 42
                var v21 float32 = (float32((-2.25 / 0.1)) / 3.14159)
                _ = v20
                _ = v21
            } else {
                var v22 uint8 = 192
                if len(v15) > 0 {
                    v15[0] = 32768
                }
                var v23 uint8 = 1
                _ = v22
                _ = v23
            }
            var v24 uint = 42
            _ = v15
            _ = v24
        }
        var v25 uint32 = 1
        _ = v25
    } else {
        var v26 float32 = 3.14159
        _ = v26
    }
    var v27 uint8 = 127
    v28 := []int{}
    _ = v27
    _ = v28
    return 100
}
func genFn1(p0 uint8) int16 {
    v29 := GenS1{F0: (p0 - (p0 ^ p0)), F1: 0, F2: (p0 ^ (^(p0)))}
    var v30 int = 9223372036854775807
    for fi31 := 0; fi31 < 18; fi31++ {
        switch (fi31 + (fi31 ^ (-(fi31)))) {
        case 1:
            fmt.Println(fi31)
        default:
            for fi32 := 0; fi32 < 17; fi32++ {
                v30 = fi31
                _ = fi32
            }
        }
        _ = fi31
    }
    _ = p0
    _ = v29
    _ = v30
    return 0
}

func main() {
    var v1 int16 = 42
    v2 := []uint32{}
    v3 := []int{310}
    v1 = v1
    v33 := GenS2{F0: 373, F1: 537}
    var v34 int32 = 1
    switch "a" {
    case "\n\t":
        for fi35 := 0; fi35 < 11; fi35++ {
            for fi36 := 0; fi36 < 6; fi36++ {
                switch "hello" {
                case "a":
                    fmt.Println(v1)
                case "\n\t":
                    var v37 int32 = (v34 - (v34 + v34))
                    _ = v37
                case "hello":
                    v3 = append(v3, fi36)
                default:
                    v34 = v34
                    var v38 int8 = -1
                    _ = v38
                }
                v39 := GenS3{F0: 2147483647, F1: 18446744073709551615}
                var v40 string = "hello"
                _ = fi36
                _ = v39
                _ = v40
            }
            _ = fi35
        }
    case "foo bar":
        v34 = (v34 << ((42) & 63))
    }
    if (!(518 > 372)) {
        if false {
            fmt.Println(v1, v1)
            v41 := GenS2{F0: 42, F1: 2}
            switch -500 {
            case 4294967295:
                v34 = v34
            case 32767:
                v34 -= -1
            }
            _ = v41
        }
        var v42 int = 558
        _ = v42
    } else {
        var v43 int64 = 1
        _ = v43
    }
    var v44 int8 = -79
    if len(v2) > 0 {
        v2[0] = 65535
    }
    v34 += 2147483647
    for fi45 := 0; fi45 < 17; fi45++ {
        var v46 uint8 = 128
        v47 := GenS4{F0: int8(v46), F1: 0}
        _ = fi45
        _ = v46
        _ = v47
    }
    for fi48 := 0; fi48 < 9; fi48++ {
        var v49 float64 = (((1.5 + 3.14159) - 3.14159) - 3.14159)
        for fi50 := 0; fi50 < 20; fi50++ {
            var v51 float64 = (float64(0.1) / v49)
            v1 = v1
            switch fi50 {
            case 75:
                var v52 bool = ("a" >= "")
                _ = v52
            }
            var v53 int = fi48
            switch ("foo bar" + "foo bar") {
            case "a":
                v44 = -109
            case "\n\t":
                v54 := GenS2{F0: 0, F1: 42}
                _ = v54
            default:
                v1 -= int16(v34)
            }
            _ = fi50
            _ = v51
            _ = v53
        }
        _ = fi48
        _ = v49
    }
    if (42 != 0) {
        if false {
            v55 := GenS3{F0: 1, F1: 1}
            v44 = (v44 * v44)
            _ = v55
        } else {
            var v56 float32 = 1e10
            _ = v56
        }
    }
    if (2147483647 >= 1) {
        var v57 uint8 = 1
        var v58 int16 = (v1 ^ genFn1((v57 % ((v57) | 1))))
        v59 := GenS2{F0: 863, F1: 0}
        _ = v57
        _ = v58
        _ = v59
    }
    if (!true) {
        switch int(v44) {
        case -9223372036854775808:
            fmt.Println(v1, v1)
        case 2147483647:
            var v60 bool = (!(1e10 != 0.0))
            _ = v60
        case -758:
            _ = len(v2)
        default:
            _ = len(v2)
        }
        v1 = (v1 | genFn1(127))
        v61 := GenS2{F0: 384, F1: 996}
        _ = v61
    }
    var v62 uint8 = 42
    v1 += (v1 * v1)
    v44 = 0
    var v63 int8 = (v44 & v44)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("v44:", v44)
    fmt.Println("v62:", v62)
    fmt.Println("v63:", v63)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v44
    _ = v62
    _ = v63
}
