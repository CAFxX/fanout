package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int; F1 uint; F2 int64 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint64; F1 string; F2 bool }
type GenS5 struct { F0 int32; F1 int8 }
func genFn0() int {
    for fi4 := 0; fi4 < 3; fi4++ {
        v5 := GenS1{F0: (fi4 << ((224) & 63))}
        _ = fi4
        _ = v5
    }
    var v6 int = -9223372036854775808
    for fi7 := 0; fi7 < 8; fi7++ {
        v6 += (fi7 + (fi7 & v6))
        var v8 uint = 18446744073709551615
        _ = fi7
        _ = v8
    }
    if true {
        switch (v6 & v6) {
        case 123:
            for fi9 := 0; fi9 < 8; fi9++ {
                fmt.Println(fi9)
                _ = fi9
            }
        case 1:
            switch v6 {
            case 127:
                var v10 int32 = -508
                _ = v10
            case 173:
                switch "\n\t" {
                case "hello":
                    for fi11 := 0; fi11 < 1; fi11++ {
                        fmt.Println(v6, v6)
                        v6 += 2147483647
                        if (-1 == 386) {
                            v12 := GenS2{F0: 128, F1: 4294967295, F2: int64(fi11)}
                            v6 = fi11
                            v13 := GenS1{F0: (v6 << ((1) & 63))}
                            _ = v12
                            _ = v13
                        } else {
                            v14 := []int{v6, (v6 - (fi11 ^ (v6 * 32767))), 4294967295}
                            var v15 uint16 = 32767
                            _ = v14
                            _ = v15
                        }
                        _ = fi11
                    }
                case "\n\t":
                    fmt.Println(v6)
                default:
                    v6 -= 9223372036854775807
                }
            }
        case 4294967295:
            v6 = 32768
        default:
            v16 := []int{(v6 * (^(v6))), v6}
            _ = v16
        }
    }
    _ = v6
    return 128
}
func genFn1() uint {
    var v17 float32 = (3.14159 + float32(1e10))
    for fi18 := 0; fi18 < 9; fi18++ {
        for fi19 := 0; fi19 < 8; fi19++ {
            if false {
                var v20 int16 = -165
                _ = v20
            } else {
                v17 = v17
                if false {
                    fmt.Println(fi19, fi19)
                }
            }
            fmt.Println(v17)
            var v21 int16 = -261
            _ = fi19
            _ = v21
        }
        v22 := []int{}
        for fi23 := 0; fi23 < 20; fi23++ {
            var v24 int = genFn0()
            if ((-2.25 + float64(0.0)) > 0.0) {
                if false {
                    switch (-(fi18)) {
                    case 42:
                        if len(v22) >= 2 {
                            v22 = v22[1:len(v22)]
                        }
                    }
                    fmt.Println(v24)
                }
                fmt.Println(v24)
                v24 = (fi18 >> ((1) & 63))
                v17 = v17
            }
            _ = fi23
            _ = v24
        }
        _ = fi18
        _ = v22
    }
    v25 := GenS1{F0: 4294967295}
    v17 = 0.1
    v26 := GenS1{F0: 1}
    _ = v17
    _ = v25
    _ = v26
    return 4294967295
}

func main() {
    var v1 int16 = 1000
    var v2 bool = (!true)
    var v3 uint64 = 4294967295
    for fi27 := 0; fi27 < 20; fi27++ {
        if true {
            v28 := GenS1{F0: (fi27 >> ((228) & 63))}
            v29 := []int{genFn0(), fi27, fi27, genFn0()}
            switch fi27 {
            case -2147483648:
                for fi30 := 0; fi30 < 11; fi30++ {
                    if len(v29) > 0 {
                        v29[0] = 9223372036854775807
                    }
                    switch 1 {
                    case 127:
                        v31 := GenS3{F0: genFn0()}
                        _ = v31
                    }
                    _ = fi30
                }
            case 32767:
                v2 = (!(65535 == 489))
            default:
                var v32 string = ("a" + "")
                _ = v32
            }
            _ = v28
            _ = v29
        }
        if v2 {
            switch (fi27 & fi27) {
            case 32768:
                var v33 string = "hello"
                _ = v33
            case 127:
                v3 -= 18446744073709551615
            }
            fmt.Println(v3, v3)
        }
        fmt.Println(v2)
        _ = fi27
    }
    v2 = (!(v2 || v2))
    if true {
        switch 128 {
        case 914:
            v34 := []int{32767}
            _ = v34
        }
        v35 := []int{}
        _ = v35
    } else {
        if (v2 && ((2147483648 >= uint32(v3)) && ((v3 < v3) && (v2 || (3.14159 == (-2.25 - 1e10)))))) {
            var v36 int = 42
            _ = v36
        } else {
            for fi37 := 0; fi37 < 8; fi37++ {
                if (!v2) {
                    switch genFn0() {
                    case 127:
                        if v2 {
                            if v2 {
                                var v38 float32 = (-2.25 + 3.14159)
                                _ = v38
                            }
                            v39 := []int{genFn0(), (fi37 / (((fi37 << ((0) & 63))) | 1))}
                            _ = v39
                        }
                    }
                }
                if (1 <= 223) {
                    v3 += (^(v3))
                    v40 := GenS1{F0: (-(fi37))}
                    var v41 int = fi37
                    _ = v40
                    _ = v41
                } else {
                    if v2 {
                        var v42 int16 = v1
                        var v43 int32 = 1
                        _ = v42
                        _ = v43
                    }
                    v3 = 1
                }
                _ = fi37
            }
            fmt.Println(v3, v1)
            v3 += v3
        }
    }
    for fi44 := 0; fi44 < 12; fi44++ {
        fmt.Println(v2, fi44)
        var v45 uint64 = (v3 % ((144) | 1))
        _ = fi44
        _ = v45
    }
    if v2 {
        switch "hello" {
        case "hello":
            for fi46 := 0; fi46 < 3; fi46++ {
                v1 -= (v1 & v1)
                _ = fi46
            }
        default:
            v47 := []int{2147483647, 760, 2147483647, 263}
            _ = v47
        }
    } else {
        var v48 int32 = 511
        if v2 {
            var v49 uint8 = 203
            for fi50 := 0; fi50 < 5; fi50++ {
                v51 := GenS3{F0: genFn0()}
                fmt.Println(v2, v49, v3)
                fmt.Println(v49, v1, v2)
                _ = fi50
                _ = v51
            }
            _ = v49
        }
        _ = v48
    }
    var v52 int8 = 100
    if (18446744073709551615 != 255) {
        for fi53 := 0; fi53 < 4; fi53++ {
            var v54 bool = true
            v54 = (v54 && (0.1 >= float64(-2.25)))
            v55 := []int{int(v52), -2147483648, fi53}
            _ = fi53
            _ = v54
            _ = v55
        }
        fmt.Println(v2, v3, v3)
        var v56 string = ("" + "\n\t")
        _ = v56
    } else {
        v57 := GenS2{F0: -1, F1: 42, F2: 1}
        _ = v57
    }
    for fi58 := 0; fi58 < 9; fi58++ {
        v1 += 1
        v59 := GenS4{F0: (v3 - 104), F1: "\n\t", F2: false}
        _ = fi58
        _ = v59
    }
    if true {
        for fi60 := 0; fi60 < 8; fi60++ {
            var v61 int = -924
            _ = fi60
            _ = v61
        }
        v62 := GenS2{F0: -1, F1: 255, F2: 9223372036854775807}
        var v63 uint8 = 127
        _ = v62
        _ = v63
    } else {
        v1 -= v1
    }
    for fi64 := 0; fi64 < 7; fi64++ {
        if (42 == 128) {
            v65 := []int{fi64, genFn0()}
            switch (fi64 / (((fi64 ^ int(v52))) | 1)) {
            case 127:
                v66 := GenS1{F0: (fi64 % ((fi64) | 1))}
                _ = v66
            case 1:
                if (!v2) {
                    if len(v65) >= 2 {
                        v65 = v65[1:len(v65)]
                    }
                    switch "hello" {
                    case "a":
                        fmt.Println(v2, v1)
                    case "foo bar":
                        for fi67 := 0; fi67 < 15; fi67++ {
                            var v68 int8 = (v52 - v52)
                            if len(v65) > 0 {
                                v65[0] = (fi64 ^ int(v52))
                            }
                            _ = fi67
                            _ = v68
                        }
                    case "\n\t":
                        v2 = (v2 || (!((1e10 * (1e10 * (0.1 - -2.25))) != (((-2.25 / 1.0) / float64((3.14159 + 1.5))) * 1e10))))
                    default:
                        if true {
                            var v69 int32 = 223
                            v69 -= v69
                            v70 := GenS5{F0: (^(v69)), F1: v52}
                            _ = v69
                            _ = v70
                        }
                    }
                    v71 := GenS4{F0: 2, F1: ("a" + "\n\t"), F2: true}
                    _ = v71
                } else {
                    var v72 int = (^(fi64))
                    var v73 uint64 = v3
                    var v74 float64 = (0.0 + 0.1)
                    _ = v72
                    _ = v73
                    _ = v74
                }
            case -9223372036854775808:
                if (("a" + "") > "\n\t") {
                    var v75 string = "a"
                    _ = v75
                } else {
                    var v76 float64 = (3.14159 * float64(-2.25))
                    _ = v76
                }
            default:
                _ = len(v65)
            }
            _ = v65
        } else {
            v3 += 1
            v1 = (-(v1))
            var v77 float32 = 1.5
            _ = v77
        }
        v52 = (-(v52))
        _ = fi64
    }
    fmt.Println(v52, v1, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v2
    _ = v3
    _ = v52
}
