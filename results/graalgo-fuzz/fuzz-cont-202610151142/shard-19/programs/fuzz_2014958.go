package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32 }
type GenS2 struct { F0 uint8; F1 int64 }
type GenS3 struct { F0 uint32; F1 int32; F2 int32 }
type GenS4 struct { F0 bool; F1 uint; F2 uint16 }

func main() {
    var v1 int32 = 42
    var v2 int16 = -503
    v3 := make(map[uint32]bool)
    switch "\n\t" {
    case "hello":
        var v4 uint16 = 999
        _ = v4
    case "":
        var v5 bool = true
        _ = v5
    default:
        fmt.Println(v1, v1)
    }
    v1 = v1
    v1 -= v1
    v6 := GenS1{F0: 0, F1: (v1 & v1)}
    v3[2147483647] = (255 != 2)
    for fi7 := 0; fi7 < 14; fi7++ {
        for fi8 := 0; fi8 < 9; fi8++ {
            v9 := GenS1{F0: 4294967295, F1: v1}
            if ((float64(1e10) / 1.0) <= (3.14159 - 3.14159)) {
                v1 += (v1 / (((v1 * 2147483647)) | 1))
                if true {
                    var v10 uint16 = 0
                    delete(v3, 686)
                    var v11 int16 = (v2 << ((498) & 63))
                    _ = v10
                    _ = v11
                }
            }
            fmt.Println(v2, fi7)
            _ = fi8
            _ = v9
        }
        delete(v3, 306)
        _ = fi7
    }
    if (("a" + ("hello" + "hello")) > "\n\t") {
        var v12 int8 = 1
        var v13 uint16 = 1
        _ = v12
        _ = v13
    } else {
        for fi14 := 0; fi14 < 11; fi14++ {
            v15 := v3[680]
            var v16 int16 = 32767
            _ = fi14
            _ = v15
            _ = v16
        }
    }
    switch 1 {
    case 93:
        for fi17 := 0; fi17 < 3; fi17++ {
            v2 -= v2
            v18 := GenS2{F0: 255, F1: int64(v1)}
            v19 := []int{(fi17 << ((4294967295) & 63)), 42}
            _ = fi17
            _ = v18
            _ = v19
        }
    case 0:
        var v20 uint8 = 234
        _ = v20
    case 528:
        v21 := GenS3{F0: 600, F1: v1, F2: 122}
        _ = v21
    }
    for fi22 := 0; fi22 < 9; fi22++ {
        for fi23 := 0; fi23 < 8; fi23++ {
            fmt.Println(v2)
            v24 := GenS1{F0: 947, F1: -2147483648}
            _ = fi23
            _ = v24
        }
        _ = fi22
    }
    for fi25 := 0; fi25 < 14; fi25++ {
        v1 -= v1
        _ = fi25
    }
    var v26 int = -173
    if (v2 < (v2 * (v2 / (((v2 << ((1) & 63))) | 1)))) {
        v2 = (v2 >> ((4294967295) & 63))
        v1 += 2147483647
    } else {
        if true {
            for fi27 := 0; fi27 < 14; fi27++ {
                v28 := []int{v26, (v26 ^ v26), 32767, (v26 % (((v26 + (fi27 | (-(fi27))))) | 1))}
                _ = fi27
                _ = v28
            }
            for fi29 := 0; fi29 < 6; fi29++ {
                fmt.Println(v1, v26, v1)
                for fi30 := 0; fi30 < 9; fi30++ {
                    if true {
                        var v31 uint32 = 0
                        v1 -= (v1 >> ((2) & 63))
                        _ = v31
                    } else {
                        fmt.Println(v2, v26)
                        v2 += (v2 - -1)
                    }
                    if true {
                        v2 += (v2 + (v2 >> ((2) & 63)))
                        fmt.Println(fi30, v26, v26)
                        _ = len(v3)
                    } else {
                        v26 = (fi30 | -19)
                    }
                    var v32 bool = (false || false)
                    _ = fi30
                    _ = v32
                }
                _ = fi29
            }
        }
        v3[1] = ((false || false) || true)
    }
    var v33 uint = 18446744073709551615
    v34 := v3[2147483647]
    v3[2] = v34
    if false {
        v34 = (1e10 < 1.5)
        fmt.Println(v33, v34, v26)
        v33 = v33
    } else {
        for fi35 := 0; fi35 < 12; fi35++ {
            v1 = 0
            var v36 int16 = (v2 >> (((v33 % ((v33) | 1))) & 63))
            if v34 {
                var v37 int = (v26 & (fi35 >> ((v33) & 63)))
                var v38 int8 = -128
                _ = v37
                _ = v38
            } else {
                var v39 uint64 = 0
                var v40 int32 = v1
                _ = v39
                _ = v40
            }
            _ = fi35
            _ = v36
        }
    }
    if false {
        fmt.Println(v1)
        for fi41 := 0; fi41 < 17; fi41++ {
            v26 = -2147483648
            if ((!(!v34)) || (!((false && ((0.0 - 0.0) >= ((1.5 * 3.14159) - 0.0))) || v34))) {
                fmt.Println(v2)
                var v42 uint = (v33 + v33)
                _ = v42
            } else {
                fmt.Println(fi41, v26, v33)
                for fi43 := 0; fi43 < 9; fi43++ {
                    v44 := GenS4{F0: v34, F1: 65535, F2: 110}
                    _ = fi43
                    _ = v44
                }
            }
            _ = fi41
        }
        v1 -= (v1 / ((v1) | 1))
    } else {
        var v45 int16 = 0
        _ = v45
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v33
    _ = v34
    _ = v6
}
