package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32; F2 bool }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint32; F1 uint64; F2 int32 }
type GenS4 struct { F0 bool; F1 uint16; F2 int16 }
type GenS5 struct { F0 bool; F1 uint8; F2 bool }
type GenS6 struct { F0 uint64; F1 uint64; F2 uint64 }
type GenS7 struct { F0 uint; F1 string }
func genFn0() uint8 {
    var v4 uint8 = 2
    for fi5 := 0; fi5 < 3; fi5++ {
        for fi6 := 0; fi6 < 2; fi6++ {
            var v7 uint64 = 195
            v8 := GenS1{F0: 0, F1: 0, F2: (0.0 <= 0.1)}
            _ = fi6
            _ = v7
            _ = v8
        }
        _ = fi5
    }
    _ = v4
    return (v4 - (^(v4)))
}
func genFn1() int8 {
    v9 := []int{4294967295, 2147483647}
    if false {
        switch -512 {
        case 45:
            var v10 uint16 = 42
            _ = v10
        case 684:
            for fi11 := 0; fi11 < 16; fi11++ {
                fmt.Println(fi11)
                for fi12 := 0; fi12 < 7; fi12++ {
                    v13 := GenS2{F0: 1}
                    var v14 float64 = (((-2.25 * (0.1 - (1e10 + 0.0))) - 1e10) + (-2.25 / 1.0))
                    v15 := GenS2{F0: 4294967295}
                    _ = fi12
                    _ = v13
                    _ = v14
                    _ = v15
                }
                var v16 float64 = float64(0.0)
                _ = fi11
                _ = v16
            }
        case 128:
            var v17 int = 4294967295
            _ = v17
        default:
            var v18 int16 = -977
            _ = v18
        }
        if len(v9) >= 2 {
            v9 = v9[1:len(v9)]
        }
    } else {
        for fi19 := 0; fi19 < 13; fi19++ {
            var v20 float32 = 3.14159
            if (374 > 42) {
                if (true && true) {
                    v21 := GenS1{F0: -128, F1: 42, F2: (65535 != 734)}
                    _ = v21
                } else {
                    switch "hello" {
                    case "foo bar":
                        if len(v9) > 0 {
                            v9[0] = (fi19 + (fi19 << ((4294967295) & 63)))
                        }
                    case "hello":
                        fmt.Println(fi19, fi19)
                    }
                    v22 := GenS1{F0: 42, F1: 0, F2: (false && ((204 == 694) && (!false)))}
                    _ = v22
                }
                fmt.Println(v20, v20)
                if (!true) {
                    v23 := GenS3{F0: 1, F1: 838, F2: -2147483648}
                    var v24 uint16 = 564
                    v24 = 953
                    _ = v23
                    _ = v24
                } else {
                    fmt.Println(v20, v20)
                    var v25 int16 = 763
                    _ = v25
                }
            } else {
                for fi26 := 0; fi26 < 5; fi26++ {
                    _ = len(v9)
                    _ = fi26
                }
            }
            var v27 int = 9223372036854775807
            _ = fi19
            _ = v20
            _ = v27
        }
    }
    _ = v9
    return 127
}

func main() {
    var v1 uint8 = 127
    v2 := make(map[string]bool)
    var v3 int = 167
    var v28 uint64 = 0
    var v29 int32 = 2147483647
    v30 := GenS3{F0: 4294967295, F1: (v28 / (((v28 & 1)) | 1)), F2: (v29 | v29)}
    _ = len(v2)
    v3 += v3
    v3 -= (v3 & 0)
    for fi31 := 0; fi31 < 2; fi31++ {
        v32 := GenS2{F0: 4294967295}
        var v33 float32 = (-2.25 - (((1.5 + 3.14159) * 1e10) * -2.25))
        _ = fi31
        _ = v32
        _ = v33
    }
    var v34 string = (("foo bar" + ("" + "a")) + "a")
    for fi35 := 0; fi35 < 9; fi35++ {
        fmt.Println(v3, v29)
        var v36 int8 = 0
        if true {
            v37 := []int{}
            for fi38 := 0; fi38 < 9; fi38++ {
                v39 := GenS4{F0: (!false), F1: 498, F2: 774}
                v34 = "a"
                _ = len(v2)
                var v40 uint32 = 4294967295
                var v41 int = 9223372036854775807
                _ = fi38
                _ = v39
                _ = v40
                _ = v41
            }
            _ = v37
        }
        _ = fi35
        _ = v36
    }
    v42 := GenS5{F0: (2147483647 <= 0), F1: v1, F2: true}
    var v43 int32 = 798
    var v44 uint32 = 286
    var v45 int32 = int32(v28)
    v46 := GenS3{F0: 1, F1: (v28 * 4294967295), F2: (-(v43))}
    if true {
        for fi47 := 0; fi47 < 1; fi47++ {
            fmt.Println(v43, v28)
            _ = fi47
        }
        var v48 float32 = (((0.1 - -2.25) + (((-2.25 * 1.5) / -2.25) + 3.14159)) * 1.5)
        _ = v48
    } else {
        v49 := GenS6{F0: (v28 * (^(v28))), F1: (v28 ^ (v28 | v28)), F2: (v28 << ((4294967295) & 63))}
        v28 += (^(v28))
        var v50 float32 = 1e10
        _ = v49
        _ = v50
    }
    delete(v2, v34)
    if (v34 >= v34) {
        v51 := GenS7{F0: 65535, F1: v34}
        for fi52 := 0; fi52 < 11; fi52++ {
            v3 += fi52
            _ = fi52
        }
        _ = v51
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v34
    _ = v42
    _ = v43
    _ = v44
    _ = v45
    _ = v46
}
