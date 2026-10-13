package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint; F2 int64 }
type GenS2 struct { F0 uint16; F1 uint8; F2 uint64 }
func genFn0() int16 {
    if (((false || false) && (370 != 65535)) || (-1 == 0)) {
        var v4 int = 42
        for fi5 := 0; fi5 < 9; fi5++ {
            var v6 string = "a"
            _ = fi5
            _ = v6
        }
        _ = v4
    } else {
        var v7 uint = 1
        if true {
            v7 -= (v7 + (v7 + v7))
            for fi8 := 0; fi8 < 17; fi8++ {
                v9 := GenS1{F0: false, F1: (v7 * (v7 - v7)), F2: 315}
                _ = fi8
                _ = v9
            }
            v10 := GenS1{F0: (!false), F1: v7, F2: -493}
            _ = v10
        } else {
            for fi11 := 0; fi11 < 7; fi11++ {
                fmt.Println(fi11, v7)
                var v12 float64 = (1.5 * (float64(3.14159) * ((0.1 + 1.5) / 3.14159)))
                if (!(false || true)) {
                    var v13 float64 = (v12 - ((0.0 + v12) * v12))
                    v7 += v7
                    v14 := []int{fi11, fi11, fi11}
                    _ = v13
                    _ = v14
                }
                _ = fi11
                _ = v12
            }
            for fi15 := 0; fi15 < 15; fi15++ {
                var v16 uint32 = 48
                var v17 uint = 4294967295
                _ = fi15
                _ = v16
                _ = v17
            }
            var v18 uint16 = 65535
            _ = v18
        }
        _ = v7
    }
    v19 := GenS1{F0: false, F1: 559, F2: 297}
    var v20 int8 = 42
    var v21 uint = 255
    _ = v19
    _ = v20
    _ = v21
    return -1
}
func genFn1() uint32 {
    switch 42 {
    case 128:
        var v22 uint8 = 255
        _ = v22
    case 483:
        var v23 bool = true
        _ = v23
    case 42:
        var v24 uint32 = 869
        _ = v24
    default:
        var v25 string = "\n\t"
        _ = v25
    }
    v26 := []int{282, 0}
    _ = len(v26)
    _ = v26
    return 79
}

func main() {
    var v1 int16 = 0
    var v2 string = ""
    v3 := []string{}
    switch ("a" + v2) {
    case "hello":
        for fi27 := 0; fi27 < 19; fi27++ {
            fmt.Println(v1)
            _ = fi27
        }
    case "":
        var v28 float32 = (1e10 * 1.5)
        _ = v28
    case "foo bar":
        _ = len(v3)
    default:
        v2 = v2
    }
    var v29 int32 = -2147483648
    fmt.Println(v1, v29, v2)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v29 -= 522
    var v30 bool = (!false)
    var v31 uint = 243
    var v32 string = v2
    var v33 uint32 = 1
    for fi34 := 0; fi34 < 7; fi34++ {
        v35 := GenS1{F0: (((v32 + v2) + v32) > ""), F1: 84, F2: int64(v33)}
        _ = fi34
        _ = v35
    }
    v1 -= int16(v33)
    v32 = "hello"
    if v30 {
        var v36 uint8 = 42
        var v37 int32 = (v29 / ((int32(v1)) | 1))
        if (v37 != 42) {
            var v38 int32 = (v37 & 42)
            _ = v38
        } else {
            for fi39 := 0; fi39 < 3; fi39++ {
                var v40 int16 = genFn0()
                _ = fi39
                _ = v40
            }
            var v41 int32 = v29
            if v30 {
                switch 127 {
                case 1:
                    var v42 uint = v31
                    _ = v42
                case 128:
                    var v43 int32 = (v37 << (((v31 << (((v31 | (v31 + v31))) & 63))) & 63))
                    _ = v43
                default:
                    v41 = v37
                }
            } else {
                if v30 {
                    if v30 {
                        fmt.Println(v31)
                        switch 9223372036854775807 {
                        case 871:
                            v30 = ((((3.14159 - 3.14159) * 0.1) * -2.25) > 0.1)
                        default:
                            v3 = append(v3, v2)
                        }
                    } else {
                        v36 = (v36 << ((v31) & 63))
                        v36 = (v36 + v36)
                        fmt.Println(v37, v37)
                    }
                    v31 -= v31
                    v44 := GenS1{F0: true, F1: v31, F2: 1}
                    _ = v44
                } else {
                    v45 := GenS2{F0: 65535, F1: v36, F2: 171}
                    fmt.Println(v31)
                    var v46 uint64 = 0
                    _ = v45
                    _ = v46
                }
                v47 := GenS2{F0: 65535, F1: 96, F2: 4294967295}
                fmt.Println(v31)
                _ = v47
            }
            _ = v41
        }
        _ = v36
        _ = v37
    }
    v31 += (v31 % (((v31 >> ((709) & 63))) | 1))
    var v48 uint64 = 2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v48:", v48)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v48
}
