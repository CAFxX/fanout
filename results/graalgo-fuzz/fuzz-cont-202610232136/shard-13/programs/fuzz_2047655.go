package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64 }
type GenS2 struct { F0 uint32; F1 uint8 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint; F1 uint64 }
type GenS5 struct { F0 bool; F1 int }
func genFn0(p0 uint16, p1 int64) uint16 {
    p1 = p1
    switch "a" {
    case "\n\t":
        var v6 int16 = 42
        _ = v6
    case "hello":
        v7 := GenS1{F0: (p0 % ((p0) | 1)), F1: p1}
        _ = v7
    default:
        var v8 int8 = 100
        _ = v8
    }
    v9 := GenS1{F0: p0, F1: (^(p1))}
    for fi10 := 0; fi10 < 12; fi10++ {
        v11 := GenS1{F0: 65535, F1: 611}
        v12 := []int{(fi10 % (((fi10 & -192)) | 1))}
        if false {
            if false {
                var v13 uint = 228
                if len(v12) >= 2 {
                    v12 = v12[1:len(v12)]
                }
                _ = v13
            } else {
                if (!false) {
                    var v14 uint64 = 4294967295
                    if len(v12) >= 2 {
                        v12 = v12[1:len(v12)]
                    }
                    _ = v14
                } else {
                    p0 -= (p0 >> ((2) & 63))
                    v15 := GenS2{F0: 1, F1: 237}
                    var v16 bool = (2 <= 915)
                    _ = v15
                    _ = v16
                }
                v12 = append(v12, fi10)
                _ = len(v12)
            }
            fmt.Println(p0, p1, fi10)
            if (true || ((!((p1 + (p1 - -446)) < p1)) && false)) {
                var v17 uint16 = 65535
                _ = v17
            } else {
                switch (fi10 << ((4294967295) & 63)) {
                case -9223372036854775808:
                    var v18 int8 = 109
                    _ = v18
                }
                v19 := GenS2{F0: 2147483648, F1: 2}
                var v20 uint8 = 255
                _ = v19
                _ = v20
            }
        }
        _ = fi10
        _ = v11
        _ = v12
    }
    var v21 uint8 = 186
    _ = p0
    _ = p1
    _ = v21
    _ = v9
    return (p0 ^ (p0 - (p0 % ((p0) | 1))))
}

func main() {
    var v1 int8 = 100
    var v2 uint = 65535
    var v3 float64 = 1.5
    v4 := []uint32{2147483647, 2147483648, 500}
    var v5 uint8 = 128
    switch 9223372036854775807 {
    case 4294967295:
        if ((v1 + v1) > (v1 ^ v1)) {
            if true {
                fmt.Println(v2)
            }
            var v22 uint16 = 0
            _ = v22
        }
    case 0:
        for fi23 := 0; fi23 < 17; fi23++ {
            fmt.Println(v2, v5)
            for fi24 := 0; fi24 < 13; fi24++ {
                fmt.Println(v5)
                switch -291 {
                case -1:
                    v4 = append(v4, 1)
                default:
                    if true {
                        v25 := GenS1{F0: 785, F1: 42}
                        if len(v4) > 0 {
                            v4[0] = uint32(v5)
                        }
                        var v26 int = -1
                        _ = v25
                        _ = v26
                    }
                }
                _ = fi24
            }
            var v27 float64 = 0.0
            _ = fi23
            _ = v27
        }
    case 108:
        fmt.Println(v2)
    default:
        var v28 uint8 = (^(v5))
        _ = v28
    }
    var v29 uint = 443
    v29 = (v2 >> (((v29 >> ((v29) & 63))) & 63))
    for fi30 := 0; fi30 < 15; fi30++ {
        for fi31 := 0; fi31 < 8; fi31++ {
            v32 := GenS1{F0: 255, F1: int64(v2)}
            _ = fi31
            _ = v32
        }
        v1 = (-(v1))
        for fi33 := 0; fi33 < 4; fi33++ {
            var v34 uint32 = 2147483648
            if (!(!true)) {
                if false {
                    fmt.Println(v3)
                    var v35 int8 = (v1 ^ v1)
                    v35 += v1
                    var v36 uint64 = uint64(v34)
                    _ = v35
                    _ = v36
                }
                v37 := GenS3{F0: 42}
                _ = v37
            } else {
                v38 := GenS4{F0: uint(v34), F1: 1}
                _ = v38
            }
            _ = fi33
            _ = v34
        }
        _ = fi30
    }
    var v39 string = ""
    for fi40 := 0; fi40 < 6; fi40++ {
        var v41 uint16 = 119
        for fi42 := 0; fi42 < 11; fi42++ {
            var v43 float64 = v3
            _ = fi42
            _ = v43
        }
        _ = fi40
        _ = v41
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v29 = v29
    var v44 uint64 = 0
    if false {
        var v45 string = (v39 + v39)
        fmt.Println(v29, v5)
        _ = v45
    } else {
        if (!true) {
            for fi46 := 0; fi46 < 7; fi46++ {
                fmt.Println(v39, v39, v2)
                for fi47 := 0; fi47 < 8; fi47++ {
                    v29 = 0
                    v29 = (v2 | 4294967295)
                    v48 := GenS1{F0: genFn0(811, -9223372036854775808), F1: 42}
                    _ = fi47
                    _ = v48
                }
                _ = fi46
            }
        }
    }
    v49 := GenS5{F0: false, F1: 128}
    v44 += v44
    for fi50 := 0; fi50 < 11; fi50++ {
        var v51 uint8 = (v5 >> ((2) & 63))
        var v52 int8 = (v1 - v1)
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = fi50
        _ = v51
        _ = v52
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v39
    _ = v4
    _ = v44
    _ = v49
    _ = v5
}
