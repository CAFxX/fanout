package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64 }
type GenS2 struct { F0 bool; F1 uint8 }
func genFn0(p0 uint64) int32 {
    for fi6 := 0; fi6 < 19; fi6++ {
        if (1 > -1) {
            for fi7 := 0; fi7 < 19; fi7++ {
                v8 := []int{(-(fi6)), (fi7 ^ 9223372036854775807), fi6, (fi7 << ((0) & 63))}
                _ = fi7
                _ = v8
            }
            switch ("foo bar" + "") {
            case "hello":
                if (!true) {
                    fmt.Println(fi6, p0)
                    var v9 int32 = -848
                    var v10 int = (^(fi6))
                    _ = v10
                    _ = v9
                } else {
                    fmt.Println(p0)
                }
            case "":
                var v11 int = (fi6 % ((fi6) | 1))
                _ = v11
            }
        }
        _ = fi6
    }
    v12 := GenS1{F0: 0, F1: 2}
    _ = p0
    _ = v12
    return 42
}
func genFn1() uint32 {
    var v13 uint32 = 2147483647
    for fi14 := 0; fi14 < 1; fi14++ {
        var v15 uint64 = 1
        v15 -= (v15 * v15)
        switch ("foo bar" + "a") {
        case "\n\t":
            if (!false) {
                if ((v13 & v13) > (v13 * 0)) {
                    for fi16 := 0; fi16 < 20; fi16++ {
                        v17 := []int{}
                        var v18 float32 = 1.5
                        _ = fi16
                        _ = v17
                        _ = v18
                    }
                    var v19 bool = false
                    _ = v19
                } else {
                    v15 = 2
                    var v20 int32 = 42
                    v21 := []int{fi14, (fi14 * (fi14 << ((uint(v13)) & 63)))}
                    _ = v20
                    _ = v21
                }
                fmt.Println(v13, fi14, v13)
            }
        default:
            var v22 int = (-(fi14))
            _ = v22
        }
        _ = fi14
        _ = v15
    }
    if false {
        var v23 int32 = int32(v13)
        v24 := GenS1{F0: 32767, F1: 18446744073709551615}
        var v25 float32 = 3.14159
        _ = v23
        _ = v24
        _ = v25
    } else {
        for fi26 := 0; fi26 < 2; fi26++ {
            var v27 uint64 = 0
            if false {
                for fi28 := 0; fi28 < 2; fi28++ {
                    var v29 int16 = -1
                    v30 := []int{}
                    switch "a" {
                    case "foo bar":
                        var v31 bool = false
                        _ = v31
                    }
                    var v32 int64 = 0
                    v33 := GenS1{F0: 1, F1: (v27 | (^(v27)))}
                    var v34 float32 = 1e10
                    _ = fi28
                    _ = v29
                    _ = v30
                    _ = v32
                    _ = v33
                    _ = v34
                }
                v35 := []int{}
                _ = v35
            } else {
                v36 := []int{fi26, fi26}
                _ = v36
            }
            switch fi26 {
            case 976:
                var v37 int = (fi26 + (fi26 & fi26))
                _ = v37
            case 9223372036854775807:
                if (-868 == genFn0((v27 << ((0) & 63)))) {
                    v38 := []int{(fi26 / ((128) | 1)), fi26}
                    var v39 int = (fi26 - 2147483647)
                    _ = v38
                    _ = v39
                } else {
                    var v40 uint32 = 4294967295
                    var v41 int32 = -1
                    v42 := []int{(-(fi26)), fi26, (fi26 / ((-304) | 1)), fi26}
                    _ = v40
                    _ = v41
                    _ = v42
                }
            default:
                v43 := []int{}
                _ = v43
            }
            _ = fi26
            _ = v27
        }
    }
    _ = v13
    return v13
}

func main() {
    v1 := make(map[int]string)
    var v2 bool = true
    var v3 int8 = 127
    var v4 uint64 = 0
    var v5 int32 = 42
    v44 := GenS2{F0: (v2 && ((v2 || v2) || (0 == 4294967295))), F1: uint8(v4)}
    v45 := GenS1{F0: 0, F1: v4}
    var v46 uint = 212
    var v47 uint8 = 128
    _ = len(v1)
    switch -2147483648 {
    case 108:
        _ = len(v1)
    case 2147483647:
        v4 = v4
    default:
        v48 := []int{}
        _ = v48
    }
    v49 := []int{}
    v46 = (v46 & v46)
    var v50 float32 = (((float32((0.1 + 1.5)) + (float32(1e10) / 3.14159)) / (float32((3.14159 / 1.0)) * 3.14159)) + 1e10)
    for fi51 := 0; fi51 < 15; fi51++ {
        var v52 uint8 = v47
        _ = fi51
        _ = v52
    }
    if v2 {
        var v53 uint64 = (v4 / ((2) | 1))
        switch 4294967295 {
        case -2147483648:
            var v54 string = ("foo bar" + "foo bar")
            _ = v54
        case 2147483647:
            v50 = 0.0
        }
        fmt.Println(v5)
        _ = v53
    }
    var v55 uint8 = 122
    for fi56 := 0; fi56 < 9; fi56++ {
        var v57 uint32 = 279
        _ = fi56
        _ = v57
    }
    v47 = uint8(v4)
    v4 += (v4 - 436)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("slicelen v49:", len(v49), v49)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v55:", v55)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v46
    _ = v47
    _ = v49
    _ = v5
    _ = v50
    _ = v55
}
