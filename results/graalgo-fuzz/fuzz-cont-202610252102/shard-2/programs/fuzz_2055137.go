package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64; F2 int8 }
type GenS2 struct { F0 uint; F1 int8 }
type GenS3 struct { F0 int; F1 uint64; F2 bool }
type GenS4 struct { F0 int; F1 int64 }
type GenS5 struct { F0 uint16; F1 bool; F2 int }
func genFn0(p0 int64, p1 uint, p2 uint8) int8 {
    v4 := GenS1{F0: 324, F1: 42, F2: int8(p1)}
    v5 := []int{967, 4294967295, 32768, int(p2)}
    var v6 uint32 = 0
    p0 = -9223372036854775808
    var v7 float32 = (float32(-2.25) + (3.14159 * 1.5))
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return -128
}
func genFn1(p0 int8, p1 uint8) uint8 {
    var v8 int8 = (p0 >> ((184) & 63))
    fmt.Println(p1, p0, p0)
    if (306 > 0) {
        var v9 int = 0
        switch "foo bar" {
        case "\n\t":
            p0 = (p0 & (p0 ^ genFn0(-520, 357, p1)))
        }
        _ = v9
    }
    v10 := []int{}
    v11 := GenS1{F0: -2147483648, F1: 973, F2: genFn0(541, 0, (p1 | (p1 % ((p1) | 1))))}
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v8
    return (^(p1))
}

func main() {
    var v1 string = "\n\t"
    var v2 int16 = 0
    v3 := make(map[int]int)
    v12 := GenS2{F0: 1, F1: 127}
    for fi13 := 0; fi13 < 7; fi13++ {
        fmt.Println(v1)
        _ = fi13
    }
    v1 = ("hello" + (((v1 + v1) + ("foo bar" + v1)) + v1))
    var v14 uint16 = 2
    v1 = "\n\t"
    v1 = v1
    var v15 uint32 = 539
    switch 4294967295 {
    case 2147483647:
        var v16 uint64 = 973
        _ = v16
    case -9223372036854775808:
        var v17 bool = (486 >= (v15 & (v15 * (v15 - uint32(v14)))))
        _ = v17
    }
    var v18 int64 = 0
    for fi19 := 0; fi19 < 9; fi19++ {
        switch fi19 {
        case -13:
            var v20 uint64 = 89
            _ = v20
        case 0:
            v21 := GenS2{F0: 1, F1: 42}
            _ = v21
        case 32768:
            v14 -= 463
        default:
            v18 = int64(fi19)
        }
        if true {
            var v22 float64 = 1.5
            var v23 uint64 = 18446744073709551615
            _ = v22
            _ = v23
        }
        if true {
            v24 := GenS2{F0: 958, F1: 127}
            _ = v24
        } else {
            var v25 uint = 68
            for fi26 := 0; fi26 < 16; fi26++ {
                var v27 uint32 = (v15 ^ (v15 << ((v25) & 63)))
                _ = fi26
                _ = v27
            }
            _ = v25
        }
        _ = fi19
    }
    for fi28 := 0; fi28 < 12; fi28++ {
        if false {
            var v29 uint = 73
            _ = v29
        } else {
            for fi30 := 0; fi30 < 3; fi30++ {
                if false {
                    _ = len(v3)
                } else {
                    v31 := GenS1{F0: 127, F1: 0, F2: 18}
                    _ = v31
                }
                v1 = "\n\t"
                v14 -= v14
                var v32 uint = 565
                if false {
                    if ((!true) || (!true)) {
                        if true {
                            fmt.Println(v18)
                            switch (v1 + v1) {
                            case "foo bar":
                                if (0 <= -39) {
                                    v1 = v1
                                    switch v1 {
                                    case "hello":
                                        v33 := GenS3{F0: -2147483648, F1: 181, F2: true}
                                        _ = v33
                                    case "\n\t":
                                        var v34 float32 = -2.25
                                        if false {
                                            var v35 uint16 = uint16(v32)
                                            v36 := GenS4{F0: -9223372036854775808, F1: v18}
                                            var v37 float64 = 3.14159
                                            _ = v35
                                            _ = v36
                                            _ = v37
                                        } else {
                                            v18 += (v18 << (((^(v32))) & 63))
                                        }
                                        _ = v34
                                    }
                                    v38 := GenS3{F0: fi28, F1: 2, F2: ((fi28 / ((int(v15)) | 1)) <= 2147483647)}
                                    _ = v38
                                } else {
                                    v39 := v3[(fi28 % ((176) | 1))]
                                    v40 := GenS4{F0: fi28, F1: (v18 | (v18 % ((42) | 1)))}
                                    _ = v39
                                    _ = v40
                                }
                            case "hello":
                                v41 := v3[(fi30 >> ((v32) & 63))]
                                _ = v41
                            case "":
                                v14 -= v14
                            }
                            fmt.Println(v32)
                        } else {
                            v42 := GenS1{F0: fi28, F1: uint64(v14), F2: 127}
                            v3[(fi28 % ((681) | 1))] = fi28
                            var v43 uint8 = 2
                            _ = v42
                            _ = v43
                        }
                        v44 := GenS3{F0: fi28, F1: 2, F2: false}
                        _ = v44
                    } else {
                        switch v1 {
                        case "\n\t":
                            if true {
                                var v45 int64 = v18
                                v46 := GenS4{F0: (^(fi30)), F1: v18}
                                _ = v45
                                _ = v46
                            }
                        case "":
                            v2 += 32767
                        }
                        if false {
                            if true {
                                v47 := GenS5{F0: v14, F1: false, F2: -9223372036854775808}
                                _ = v47
                            } else {
                                var v48 int = fi30
                                _ = v48
                            }
                        } else {
                            var v49 int64 = 0
                            var v50 bool = false
                            _ = v49
                            _ = v50
                        }
                    }
                } else {
                    if true {
                        var v51 uint16 = (v14 * v14)
                        _ = len(v3)
                        fmt.Println(v1, v18, v32)
                        var v52 float64 = (float64(3.14159) * (float64(3.14159) * (-2.25 * -2.25)))
                        v51 = 42
                        _ = v51
                        _ = v52
                    }
                    var v53 bool = (0.1 >= 0.0)
                    v54 := GenS1{F0: fi30, F1: 2, F2: 35}
                    _ = v53
                    _ = v54
                }
                fmt.Println(v15, v32, v15)
                _ = fi30
                _ = v32
            }
            v14 = v14
            v15 -= v15
        }
        _ = fi28
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v12
    _ = v14
    _ = v15
    _ = v18
    _ = v2
    _ = v3
}
