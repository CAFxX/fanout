package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8; F2 int16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int8; F1 string; F2 uint8 }
type GenS5 struct { F0 uint64; F1 uint; F2 int16 }
type GenS6 struct { F0 uint32; F1 uint8 }
func genFn0(p0 uint32) string {
    if (1e10 >= -2.25) {
        v6 := GenS1{F0: 2147483647, F1: -79, F2: 1000}
        v7 := []int{32767, 598, 128, -976}
        _ = v6
        _ = v7
    } else {
        var v8 int8 = 76
        v9 := GenS2{F0: -234}
        for fi10 := 0; fi10 < 20; fi10++ {
            v11 := GenS1{F0: 578, F1: v8, F2: -32768}
            v12 := GenS2{F0: -1}
            _ = fi10
            _ = v11
            _ = v12
        }
        _ = v8
        _ = v9
    }
    v13 := GenS2{F0: 32767}
    v14 := GenS3{F0: (4294967295 < 2)}
    _ = p0
    _ = v13
    _ = v14
    return ""
}

func main() {
    var v1 int16 = 1
    v2 := make(map[string]uint32)
    v3 := make(map[uint32]int)
    var v4 float64 = 0.1
    var v5 uint64 = 4294967295
    v15 := []int{362, -479, 42}
    v5 += v5
    v15 = append(v15, 42)
    var v16 int = 42
    var v17 uint8 = 228
    if len(v15) >= 2 {
        v15 = v15[1:len(v15)]
    }
    for fi18 := 0; fi18 < 7; fi18++ {
        switch (fi18 | fi18) {
        case 374:
            for fi19 := 0; fi19 < 17; fi19++ {
                v17 = v17
                _ = fi19
            }
        case 44:
            for fi20 := 0; fi20 < 2; fi20++ {
                v21 := GenS2{F0: v1}
                _ = fi20
                _ = v21
            }
        }
        if false {
            v16 += 847
            v22 := GenS4{F0: 100, F1: "", F2: (v17 - v17)}
            for fi23 := 0; fi23 < 19; fi23++ {
                v17 += 127
                if false {
                    var v24 uint16 = 65535
                    _ = v24
                }
                switch fi23 {
                case -108:
                    var v25 string = "a"
                    _ = v25
                default:
                    v1 = (v1 | 1)
                    fmt.Println(v1, v5, v17)
                }
                switch (^(fi23)) {
                case -65:
                    var v26 int64 = 9223372036854775807
                    v27 := GenS4{F0: int8(v26), F1: (((("foo bar" + "a") + "foo bar") + "a") + "a"), F2: 134}
                    if false {
                        if len(v15) >= 2 {
                            v15 = v15[1:len(v15)]
                        }
                        if len(v15) > 0 {
                            v15[0] = (v16 >> ((2) & 63))
                        }
                        v28 := GenS5{F0: (v5 << ((42) & 63)), F1: 65535, F2: 0}
                        _ = v28
                    } else {
                        var v29 uint16 = 205
                        if (!(v26 == (v26 & v26))) {
                            v4 += -2.25
                            var v30 int16 = -497
                            var v31 int32 = 0
                            _ = v30
                            _ = v31
                        }
                        v16 = fi18
                        _ = v29
                    }
                    _ = len(v15)
                    _ = v26
                    _ = v27
                case 128:
                    v4 -= (1e10 - ((float64(float32(v4)) - ((v4 + 1e10) * 0.1)) * 3.14159))
                    switch "foo bar" {
                    case "\n\t":
                        fmt.Println(fi23)
                    case "":
                        v17 = 127
                    default:
                        if (v5 > 0) {
                            v32 := GenS2{F0: (v1 % (((v1 % ((v1) | 1))) | 1))}
                            var v33 bool = false
                            v1 += v1
                            _ = v32
                            _ = v33
                        }
                    }
                    if (81 >= 321) {
                        if ((false && true) && true) {
                            if (v5 >= (v5 - 18446744073709551615)) {
                                var v34 string = (("" + "a") + "a")
                                v35 := v2[v34]
                                _ = v34
                                _ = v35
                            }
                            var v36 int8 = -1
                            _ = v36
                        } else {
                            v5 = (^(v5))
                        }
                        v37 := GenS5{F0: 467, F1: 4294967295, F2: (v1 ^ 1)}
                        _ = v37
                    } else {
                        _ = len(v2)
                        v4 -= 0.1
                        var v38 int32 = -1
                        _ = v38
                    }
                case 47:
                    v39 := GenS1{F0: 1, F1: -128, F2: v1}
                    _ = v39
                }
                _ = fi23
            }
            _ = v22
        } else {
            v40 := GenS6{F0: 536, F1: v17}
            var v41 uint8 = v17
            _ = v40
            _ = v41
        }
        fmt.Println(v1)
        _ = fi18
    }
    var v42 int8 = 127
    v15 = append(v15, (v16 - (v16 & v16)))
    v42 += (v42 % ((1) | 1))
    v17 += v17
    _ = len(v15)
    for fi43 := 0; fi43 < 6; fi43++ {
        fmt.Println(v42, v4, v42)
        for fi44 := 0; fi44 < 2; fi44++ {
            fmt.Println(fi43, v4, v42)
            _ = fi44
        }
        switch (fi43 % ((0) | 1)) {
        case 32767:
            var v45 bool = ("hello" < "")
            _ = v45
        case 1:
            v46 := GenS4{F0: 0, F1: "", F2: (v17 ^ v17)}
            _ = v46
        default:
            for fi47 := 0; fi47 < 13; fi47++ {
                if false {
                    v1 = -32768
                } else {
                    if true {
                        fmt.Println(fi47, v4, v5)
                        v42 -= (v42 ^ v42)
                        if (65535 > 544) {
                            v4 += v4
                            var v48 float64 = v4
                            v15 = append(v15, fi47)
                            _ = v48
                        } else {
                            var v49 int32 = -2147483648
                            fmt.Println(v4, v17)
                            v17 = (v17 & v17)
                            _ = v49
                        }
                        switch fi47 {
                        case -107:
                            v5 -= v5
                        case 127:
                            var v50 uint8 = (^(v17))
                            _ = v50
                        case -1:
                            var v51 int32 = -708
                            _ = v51
                        }
                    }
                }
                v52 := GenS6{F0: 540, F1: v17}
                if true {
                    v5 = (v5 | v5)
                }
                _ = fi47
                _ = v52
            }
        }
        _ = fi43
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v5
}
