package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint16; F1 uint }
type GenS3 struct { F0 uint32; F1 int16; F2 uint16 }
type GenS4 struct { F0 string; F1 uint32; F2 int }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 int }

func main() {
    v1 := make(map[int]string)
    var v2 float32 = (1e10 + 3.14159)
    v3 := make(map[int]uint32)
    if ((false && (1 != 18446744073709551615)) && false) {
        v4 := GenS1{F0: 94}
        fmt.Println(v2)
        if (350 < 42) {
            delete(v1, 796)
            var v5 bool = false
            var v6 int32 = -2147483648
            _ = v5
            _ = v6
        }
        _ = v4
    }
    for fi7 := 0; fi7 < 6; fi7++ {
        v2 -= (float32(1e10) / (((v2 - (v2 - float32(((1.5 + -2.25) * -2.25)))) - (v2 * v2)) + -2.25))
        _ = fi7
    }
    var v8 int16 = 584
    var v9 uint16 = 65535
    var v10 string = "a"
    switch -512 {
    case -203:
        var v11 bool = false
        _ = v11
    case 837:
        switch v10 {
        case "a":
            var v12 uint8 = 42
            _ = v12
        case "foo bar":
            v8 += (v8 & (v8 << ((65535) & 63)))
        case "":
            v9 = 108
        default:
            v13 := GenS1{F0: -126}
            _ = v13
        }
    default:
        var v14 float64 = 1e10
        _ = v14
    }
    for fi15 := 0; fi15 < 20; fi15++ {
        v16 := GenS1{F0: 56}
        if (!false) {
            for fi17 := 0; fi17 < 3; fi17++ {
                if true {
                    switch (fi17 ^ fi17) {
                    case 78:
                        v18 := []int{662}
                        _ = v18
                    case 127:
                        v19 := GenS2{F0: (v9 / ((v9) | 1)), F1: 18446744073709551615}
                        _ = v19
                    default:
                        if false {
                            fmt.Println(fi17)
                            fmt.Println(v8, v10, fi15)
                            v20 := GenS1{F0: 85}
                            var v21 int = (fi15 ^ -304)
                            if true {
                                if false {
                                    if false {
                                        v8 = 94
                                    } else {
                                        var v22 uint32 = 2147483647
                                        var v23 int16 = (v8 | v8)
                                        v9 = v9
                                        _ = v22
                                        _ = v23
                                    }
                                    if (1e10 == (1e10 - (0.1 / float64(0.1)))) {
                                        var v24 uint8 = 77
                                        fmt.Println(fi15)
                                        _ = v24
                                    } else {
                                        v9 += (v9 >> ((1) & 63))
                                        var v25 uint8 = 128
                                        _ = v25
                                    }
                                    v26 := []int{-2147483648}
                                    _ = v26
                                }
                                _ = len(v1)
                                var v27 int32 = 1
                                var v28 uint16 = (v9 % ((32767) | 1))
                                _ = v27
                                _ = v28
                            }
                            _ = v20
                            _ = v21
                        } else {
                            var v29 string = "foo bar"
                            _ = v29
                        }
                    }
                } else {
                    v8 += (v8 + v8)
                }
                v30 := GenS3{F0: 2, F1: 271, F2: v9}
                if (32767 < v9) {
                    delete(v3, (^(fi17)))
                    var v31 uint16 = 2
                    _ = v31
                }
                v32 := []int{fi15, fi15, -395}
                var v33 uint8 = 128
                var v34 uint64 = 176
                switch (fi15 % (((fi17 * (fi15 & (fi15 ^ -9223372036854775808)))) | 1)) {
                case -772:
                    v35 := GenS3{F0: 0, F1: v8, F2: v9}
                    _ = v35
                }
                _ = fi17
                _ = v30
                _ = v32
                _ = v33
                _ = v34
            }
        } else {
            v36 := GenS4{F0: v10, F1: 290, F2: 0}
            var v37 int16 = 32767
            for fi38 := 0; fi38 < 15; fi38++ {
                var v39 uint8 = 2
                if false {
                    v40 := GenS4{F0: v10, F1: 4294967295, F2: (fi38 + (fi15 + fi15))}
                    var v41 float32 = (v2 - ((v2 - (((v2 - (float32(-2.25) / 0.1)) * 1.5) + (v2 - (v2 / v2)))) * 1e10))
                    _ = v40
                    _ = v41
                } else {
                    v2 = v2
                    v42 := GenS3{F0: 2147483647, F1: 32767, F2: (v9 / (((v9 ^ v9)) | 1))}
                    var v43 uint8 = v39
                    _ = v42
                    _ = v43
                }
                _ = fi38
                _ = v39
            }
            _ = v36
            _ = v37
        }
        if (v9 <= (v9 * (v9 * v9))) {
            fmt.Println(v8)
        }
        _ = fi15
        _ = v16
    }
    for fi44 := 0; fi44 < 20; fi44++ {
        fmt.Println(v8, v9)
        var v45 uint8 = 128
        _ = fi44
        _ = v45
    }
    var v46 uint = 4294967295
    switch -9223372036854775808 {
    case -49:
        if ((true || false) && (65535 >= 4294967295)) {
            var v47 int8 = -128
            _ = v47
        }
    default:
        for fi48 := 0; fi48 < 20; fi48++ {
            fmt.Println(v8)
            var v49 int = -2147483648
            _ = fi48
            _ = v49
        }
    }
    if true {
        var v50 uint = v46
        _ = v50
    } else {
        var v51 uint = v46
        _ = len(v3)
        var v52 uint = v46
        _ = v51
        _ = v52
    }
    fmt.Println(v46)
    switch 417 {
    case 1:
        for fi53 := 0; fi53 < 15; fi53++ {
            v54 := GenS5{F0: 65535}
            _ = fi53
            _ = v54
        }
    case 4294967295:
        switch (v10 + ((v10 + "foo bar") + ("hello" + v10))) {
        case "hello":
            v55 := GenS6{F0: 127}
            _ = v55
        default:
            v2 = 1.5
        }
    default:
        var v56 uint32 = 782
        _ = v56
    }
    var v57 string = v10
    switch v57 {
    case "a":
        fmt.Println(v2)
    default:
        v46 += (v46 ^ (^(v46)))
    }
    v9 -= (v9 >> ((v46) & 63))
    for fi58 := 0; fi58 < 16; fi58++ {
        v46 += v46
        v59 := v3[fi58]
        _ = fi58
        _ = v59
    }
    fmt.Println(v9, v9)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v46:", v46)
    fmt.Println("v57:", v57)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v46
    _ = v57
    _ = v8
    _ = v9
}
