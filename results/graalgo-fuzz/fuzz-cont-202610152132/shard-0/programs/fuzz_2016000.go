package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16 }
type GenS2 struct { F0 int16; F1 int; F2 uint16 }
type GenS3 struct { F0 uint32; F1 bool; F2 int64 }
type GenS4 struct { F0 uint }
func genFn0(p0 uint64) uint32 {
    v4 := GenS1{F0: (p0 % ((p0) | 1)), F1: 0}
    if true {
        var v5 uint = 1
        p0 = (p0 << ((v5) & 63))
        _ = v5
    }
    if ("" != "a") {
        v6 := []int{}
        p0 = 2
        _ = v6
    } else {
        p0 = p0
        for fi7 := 0; fi7 < 19; fi7++ {
            var v8 float64 = -2.25
            _ = fi7
            _ = v8
        }
    }
    v9 := GenS2{F0: 42, F1: 1, F2: 2}
    v10 := GenS3{F0: 602, F1: (!true), F2: 1}
    _ = p0
    _ = v10
    _ = v4
    _ = v9
    return 660
}
func genFn1(p0 uint8, p1 int32) uint8 {
    v11 := GenS2{F0: 276, F1: 1, F2: 65535}
    p1 += 1
    var v12 uint16 = 42
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    return 133
}

func main() {
    v1 := []uint32{}
    v2 := []uint32{158, 65535, 42}
    var v3 float32 = 3.14159
    if false {
        var v13 int32 = -2147483648
        fmt.Println(v13, v3)
        var v14 float32 = 0.1
        _ = v13
        _ = v14
    } else {
        var v15 uint32 = 2
        for fi16 := 0; fi16 < 13; fi16++ {
            if (525 != 0) {
                v17 := GenS1{F0: 0, F1: 42}
                var v18 int16 = 407
                _ = len(v2)
                _ = v17
                _ = v18
            } else {
                var v19 int8 = 1
                v15 += 4294967295
                var v20 int32 = 1
                _ = v19
                _ = v20
            }
            _ = fi16
        }
        _ = v15
    }
    fmt.Println(v3)
    var v21 string = "foo bar"
    if (true && (!(-9223372036854775808 == 0))) {
        v22 := GenS3{F0: genFn0(2), F1: false, F2: -1}
        v23 := GenS1{F0: 56, F1: 42}
        if len(v1) > 0 {
            v1[0] = 2147483648
        }
        _ = v22
        _ = v23
    }
    var v24 uint16 = 255
    v24 += v24
    var v25 int8 = 0
    switch 872 {
    case 1:
        switch 0 {
        case -2147483648:
            v26 := GenS2{F0: 32767, F1: 32768, F2: v24}
            _ = v26
        case 4294967295:
            if false {
                if (!(true && true)) {
                    switch 1 {
                    case -9223372036854775808:
                        for fi27 := 0; fi27 < 20; fi27++ {
                            if true {
                                v3 = v3
                                if (fi27 < fi27) {
                                    v21 = (("\n\t" + v21) + v21)
                                } else {
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    var v28 float32 = ((v3 + (float32(3.14159) + (v3 - v3))) / ((float32(1e10) * v3) * -2.25))
                                    _ = v28
                                }
                            }
                            _ = fi27
                        }
                    }
                }
                var v29 uint32 = 0
                v24 = (v24 / ((v24) | 1))
                _ = v29
            } else {
                v25 = 123
                v24 += 0
                v30 := GenS2{F0: -1, F1: -237, F2: 1}
                _ = v30
            }
        case -1:
            v21 = "hello"
        }
    case 32768:
        var v31 uint = 255
        _ = v31
    default:
        for fi32 := 0; fi32 < 1; fi32++ {
            v24 = 42
            _ = fi32
        }
    }
    fmt.Println(v25, v25, v21)
    if true {
        var v33 uint32 = 424
        _ = v33
    } else {
        if true {
            v34 := GenS2{F0: 0, F1: 1, F2: 255}
            _ = v34
        }
        for fi35 := 0; fi35 < 19; fi35++ {
            fmt.Println(v3, v3, v24)
            v25 = 42
            _ = fi35
        }
    }
    v2 = append(v2, 363)
    for fi36 := 0; fi36 < 11; fi36++ {
        var v37 uint8 = 159
        v38 := GenS4{F0: 65535}
        _ = fi36
        _ = v37
        _ = v38
    }
    switch 127 {
    case 1:
        v21 = v21
    case 32768:
        var v39 bool = false
        _ = v39
    case 541:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    v40 := GenS1{F0: 0, F1: (v24 / (((v24 & v24)) | 1))}
    var v41 bool = true
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v2
    _ = v21
    _ = v24
    _ = v25
    _ = v3
    _ = v40
    _ = v41
}
