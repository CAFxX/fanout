package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32 }
type GenS2 struct { F0 int64; F1 int64; F2 string }
type GenS3 struct { F0 bool; F1 uint8 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint8; F1 bool; F2 int32 }
type GenS6 struct { F0 uint32; F1 int16 }
type GenS7 struct { F0 int64; F1 int; F2 uint64 }
func genFn0(p0 int8) bool {
    if false {
        v6 := []int{42}
        _ = v6
    } else {
        if ((214 != 63) && ((2 >= 277) || true)) {
            var v7 string = ""
            for fi8 := 0; fi8 < 2; fi8++ {
                if false {
                    for fi9 := 0; fi9 < 4; fi9++ {
                        v10 := GenS1{F0: 1, F1: 2147483647}
                        var v11 uint32 = 42
                        _ = fi9
                        _ = v10
                        _ = v11
                    }
                    v12 := []int{127}
                    _ = v12
                } else {
                    v13 := []int{}
                    _ = len(v13)
                    _ = v13
                }
                if (679 < -1) {
                    v14 := GenS2{F0: 166, F1: -1, F2: ""}
                    _ = v14
                }
                fmt.Println(v7, v7, fi8)
                _ = fi8
            }
            _ = v7
        }
        p0 -= p0
    }
    var v15 int16 = 1000
    v16 := GenS1{F0: 127, F1: 2147483647}
    _ = p0
    _ = v15
    _ = v16
    return true
}

func main() {
    var v1 string = "a"
    var v2 uint16 = 255
    var v3 int32 = -1
    var v4 uint8 = 42
    var v5 float64 = 3.14159
    v17 := GenS3{F0: false, F1: uint8(v2)}
    v18 := GenS3{F0: false, F1: 42}
    for fi19 := 0; fi19 < 5; fi19++ {
        if true {
            v5 = v5
        } else {
            if (-9223372036854775808 > 747) {
                fmt.Println(fi19)
            }
            v20 := GenS4{F0: v2}
            _ = v20
        }
        for fi21 := 0; fi21 < 13; fi21++ {
            v4 = v4
            _ = fi21
        }
        fmt.Println(v5, v3)
        _ = fi19
    }
    v22 := []int{-115}
    if ((0 <= 874) && true) {
        if (!(!(-128 != -1))) {
            for fi23 := 0; fi23 < 16; fi23++ {
                for fi24 := 0; fi24 < 6; fi24++ {
                    switch v1 {
                    case "hello":
                        v4 += v4
                    default:
                        v1 = v1
                    }
                    v25 := GenS5{F0: v4, F1: (v4 > 127), F2: -1}
                    _ = fi24
                    _ = v25
                }
                v4 -= uint8(v2)
                _ = fi23
            }
        } else {
            v3 += (v3 ^ (v3 & v3))
            var v26 int = int(v3)
            _ = v26
        }
        v27 := GenS6{F0: 2147483648, F1: -32768}
        _ = len(v22)
        _ = v27
    } else {
        fmt.Println(v3)
    }
    v28 := GenS5{F0: (v4 << ((551) & 63)), F1: (!false), F2: -2147483648}
    if (255 == 128) {
        if len(v22) > 0 {
            v22[0] = 1
        }
        if false {
            v29 := GenS5{F0: (^(v4)), F1: ((!(4294967295 > 4294967295)) && ((-2.25 < 1e10) && (!false))), F2: (-(v3))}
            v4 -= v4
            _ = v29
        } else {
            fmt.Println(v2)
            var v30 uint16 = (v2 >> ((407) & 63))
            _ = v30
        }
        if false {
            v3 += v3
            if true {
                if false {
                    if len(v22) >= 2 {
                        v22 = v22[1:len(v22)]
                    }
                    var v31 float64 = (v5 - ((v5 + (v5 / 1.5)) / float64(1e10)))
                    var v32 bool = (1 >= -880)
                    _ = v31
                    _ = v32
                } else {
                    for fi33 := 0; fi33 < 19; fi33++ {
                        for fi34 := 0; fi34 < 14; fi34++ {
                            v2 += (v2 << ((1) & 63))
                            _ = fi34
                        }
                        fmt.Println(v2)
                        _ = fi33
                    }
                    if (!(-1 != -1)) {
                        var v35 uint32 = 421
                        _ = v35
                    }
                    v3 = (v3 - -934)
                }
            }
            if true {
                var v36 int = 705
                fmt.Println(v36, v2)
                _ = v36
            }
        } else {
            if (v1 > "\n\t") {
                if len(v22) >= 2 {
                    v22 = v22[1:len(v22)]
                }
            }
            v37 := GenS1{F0: (^(v4)), F1: v3}
            _ = v37
        }
    }
    switch 127 {
    case -1:
        v2 += 32767
    case 128:
        v38 := GenS7{F0: 9223372036854775807, F1: -1, F2: 1}
        _ = v38
    }
    if false {
        var v39 uint16 = 2
        switch v1 {
        case "hello":
            for fi40 := 0; fi40 < 12; fi40++ {
                fmt.Println(v5, v2, v39)
                _ = fi40
            }
        default:
            v1 = "a"
        }
        var v41 uint64 = uint64(v39)
        _ = v39
        _ = v41
    }
    var v42 uint8 = (v4 * (^(v4)))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v22
    _ = v28
    _ = v3
    _ = v4
    _ = v42
    _ = v5
}
