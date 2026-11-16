package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint64; F1 int8 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint16; F1 uint8 }
type GenS5 struct { F0 int; F1 bool }
type GenS6 struct { F0 uint; F1 uint8; F2 uint32 }
type GenS7 struct { F0 int64; F1 int8 }
func genFn0() uint8 {
    for fi4 := 0; fi4 < 11; fi4++ {
        var v5 int32 = 802
        switch fi4 {
        case 2147483647:
            v6 := []int{}
            _ = v6
        case -800:
            for fi7 := 0; fi7 < 17; fi7++ {
                switch 32768 {
                case 128:
                    v8 := GenS1{F0: 483}
                    _ = v8
                case 4294967295:
                    switch "hello" {
                    case "a":
                        v5 = 42
                    case "hello":
                        if (!true) {
                            fmt.Println(fi7, fi7, v5)
                            var v9 float32 = 0.0
                            var v10 int32 = (v5 >> ((255) & 63))
                            _ = v10
                            _ = v9
                        }
                    case "foo bar":
                        var v11 uint8 = 1
                        _ = v11
                    default:
                        v5 = v5
                    }
                case -2147483648:
                    v12 := GenS2{F0: 0, F1: 127}
                    _ = v12
                }
                var v13 uint16 = 2
                _ = fi7
                _ = v13
            }
        default:
            v14 := []int{fi4, (fi4 - (fi4 << ((2) & 63))), (fi4 & 72), (fi4 >> ((1) & 63))}
            _ = v14
        }
        v5 -= v5
        _ = fi4
        _ = v5
    }
    v15 := GenS3{F0: 0}
    if false {
        switch -174 {
        case 682:
            if true {
                switch 2147483647 {
                case 189:
                    var v16 int8 = 0
                    _ = v16
                case 128:
                    if false {
                        var v17 int64 = -9223372036854775808
                        _ = v17
                    }
                case 32767:
                    v18 := []int{-288, 0}
                    _ = v18
                }
            }
        case 9223372036854775807:
            var v19 uint32 = 2147483648
            _ = v19
        case -643:
            var v20 uint64 = 0
            _ = v20
        default:
            var v21 uint16 = 2
            _ = v21
        }
        var v22 float64 = (0.0 + 1e10)
        switch 127 {
        case 615:
            fmt.Println(v22)
        case 32768:
            v23 := []int{0}
            _ = v23
        default:
            var v24 int16 = 42
            _ = v24
        }
        _ = v22
    }
    v25 := []int{32768, 2147483647, -77, 128}
    v26 := GenS3{F0: 42}
    _ = v15
    _ = v25
    _ = v26
    return 1
}

func main() {
    var v1 bool = false
    var v2 string = ("a" + "a")
    v3 := make(map[uint32]int)
    for fi27 := 0; fi27 < 14; fi27++ {
        if (128 != 847) {
            fmt.Println(v2)
        } else {
            fmt.Println(fi27, fi27)
            v28 := GenS4{F0: 0, F1: 24}
            _ = v28
        }
        _ = fi27
    }
    v2 = "foo bar"
    fmt.Println(v2, v1)
    v29 := GenS5{F0: -9223372036854775808, F1: true}
    v1 = v1
    switch ((v2 + (v2 + (((v2 + v2) + v2) + v2))) + ("\n\t" + v2)) {
    case "foo bar":
        fmt.Println(v2)
    case "\n\t":
        v30 := []int{32768}
        _ = v30
    case "hello":
        v31 := GenS1{F0: 569}
        _ = v31
    default:
        if ((18446744073709551615 <= 18446744073709551615) && true) {
            var v32 int64 = 1
            if (-44 < 42) {
                for fi33 := 0; fi33 < 4; fi33++ {
                    if (612 != 65535) {
                        switch 32767 {
                        case -825:
                            if v1 {
                                var v34 uint8 = 0
                                _ = v34
                            }
                        default:
                            v2 = (v2 + v2)
                        }
                    }
                    if (4294967295 <= 603) {
                        var v35 int32 = 393
                        _ = v35
                    } else {
                        if (!v1) {
                            if (!(-56 < -1)) {
                                var v36 int64 = v32
                                v1 = v1
                                _ = v36
                            }
                        }
                        _ = len(v3)
                    }
                    delete(v3, 42)
                    _ = fi33
                }
                v32 = (v32 & (v32 ^ 281))
                var v37 float64 = 0.1
                _ = v37
            } else {
                var v38 bool = true
                v39 := []int{-667, 924}
                v38 = (v38 && v1)
                _ = v38
                _ = v39
            }
            _ = v32
        }
    }
    v1 = (false && (((1.5 * 1e10) + 3.14159) != -2.25))
    var v40 int32 = 2147483647
    for fi41 := 0; fi41 < 14; fi41++ {
        fmt.Println(v2, v2, v1)
        v1 = true
        var v42 uint = 4294967295
        _ = fi41
        _ = v42
    }
    switch 664 {
    case -1:
        var v43 int = -855
        _ = v43
    case 9223372036854775807:
        for fi44 := 0; fi44 < 20; fi44++ {
            v2 = (v2 + v2)
            fmt.Println(fi44, v40)
            var v45 bool = v1
            _ = fi44
            _ = v45
        }
    }
    v46 := GenS6{F0: 4294967295, F1: 233, F2: 4294967295}
    _ = len(v3)
    if v1 {
        if true {
            v47 := v3[2]
            _ = v47
        } else {
            if true {
                var v48 bool = v1
                _ = v48
            }
            v49 := GenS7{F0: -1, F1: -1}
            _ = v49
        }
        v2 = v2
        fmt.Println(v2)
    }
    v3[2147483648] = -10
    v50 := []int{32767, -465}
    var v51 uint8 = genFn0()
    var v52 int = 64
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v40:", v40)
    fmt.Println("slicelen v50:", len(v50), v50)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v40
    _ = v46
    _ = v50
    _ = v51
    _ = v52
}
