package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint; F1 uint64; F2 int64 }
type GenS3 struct { F0 int; F1 int8 }
type GenS4 struct { F0 uint16; F1 uint16; F2 int }
type GenS5 struct { F0 uint32; F1 uint64 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 uint32; F1 int }
type GenS8 struct { F0 uint16; F1 uint64; F2 uint64 }
type GenS9 struct { F0 uint64; F1 uint16 }

func main() {
    v1 := make(map[string]int)
    v2 := []uint32{1, 35, 2147483647, 356}
    v3 := make(map[int]uint32)
    var v4 int8 = -128
    var v5 uint = 255
    for fi6 := 0; fi6 < 3; fi6++ {
        v4 -= v4
        _ = fi6
    }
    if false {
        var v7 float64 = 0.0
        for fi8 := 0; fi8 < 4; fi8++ {
            v4 = (v4 >> (((v5 - v5)) & 63))
            _ = len(v1)
            var v9 float64 = (v7 - 1.5)
            _ = fi8
            _ = v9
        }
        var v10 uint = 4294967295
        _ = v10
        _ = v7
    }
    for fi11 := 0; fi11 < 16; fi11++ {
        v12 := GenS1{F0: v4}
        if false {
            var v13 uint16 = 255
            _ = v13
        }
        if (32767 > 1000) {
            var v14 string = "\n\t"
            var v15 int64 = 42
            delete(v1, "hello")
            _ = v14
            _ = v15
        }
        _ = fi11
        _ = v12
    }
    fmt.Println(v5, v4)
    v4 -= (v4 / ((v4) | 1))
    fmt.Println(v5, v4)
    fmt.Println(v4, v4)
    var v16 int16 = 1000
    for fi17 := 0; fi17 < 14; fi17++ {
        if (false || (false && true)) {
            switch "" {
            case "foo bar":
                switch "hello" {
                case "hello":
                    v18 := GenS2{F0: 4294967295, F1: 42, F2: -9223372036854775808}
                    _ = v18
                case "":
                    fmt.Println(v5, v4)
                case "foo bar":
                    if (1 > 2) {
                        v19 := GenS1{F0: v4}
                        switch "" {
                        case "a":
                            fmt.Println(v4)
                        case "foo bar":
                            if false {
                                if true {
                                    v20 := GenS1{F0: v4}
                                    var v21 int = (fi17 << (((v5 << ((903) & 63))) & 63))
                                    switch (v21 - int(v4)) {
                                    case 275:
                                        for fi22 := 0; fi22 < 12; fi22++ {
                                            v4 = v4
                                            var v23 uint64 = 19
                                            _ = fi22
                                            _ = v23
                                        }
                                    case -434:
                                        switch ("a" + ("" + "")) {
                                        case "a":
                                            var v24 uint64 = 322
                                            _ = v24
                                        case "\n\t":
                                            switch "hello" {
                                            case "a":
                                                for fi25 := 0; fi25 < 10; fi25++ {
                                                    fmt.Println(v5)
                                                    v26 := GenS3{F0: fi17, F1: -108}
                                                    v16 -= -1
                                                    _ = fi25
                                                    _ = v26
                                                }
                                            case "foo bar":
                                                if true {
                                                    delete(v1, ("\n\t" + ""))
                                                    if false {
                                                        if ((false && (1 < 363)) && true) {
                                                            v27 := GenS1{F0: (v4 << (((^(v5))) & 63))}
                                                            if (668 <= 42) {
                                                                var v28 uint32 = 748
                                                                _ = v28
                                                            } else {
                                                                var v29 uint32 = 0
                                                                fmt.Println(v29)
                                                                _ = v29
                                                            }
                                                            v30 := GenS4{F0: 1, F1: 2, F2: v21}
                                                            _ = v27
                                                            _ = v30
                                                        } else {
                                                            for fi31 := 0; fi31 < 13; fi31++ {
                                                                if true {
                                                                    if len(v2) >= 2 {
                                                                        v2 = v2[1:len(v2)]
                                                                    }
                                                                } else {
                                                                    if (false || false) {
                                                                        v16 -= 42
                                                                        v32 := GenS4{F0: 2, F1: 0, F2: fi31}
                                                                        _ = v32
                                                                    }
                                                                    v33 := GenS5{F0: 588, F1: 276}
                                                                    v4 += v4
                                                                    _ = v33
                                                                }
                                                                _ = fi31
                                                            }
                                                            v16 -= (v16 & v16)
                                                        }
                                                        for fi34 := 0; fi34 < 1; fi34++ {
                                                            if false {
                                                                fmt.Println(fi17, v16)
                                                                v35 := GenS6{F0: -2147483648}
                                                                var v36 int64 = -601
                                                                v37 := GenS7{F0: 0, F1: -9223372036854775808}
                                                                var v38 float64 = float64(float32(float64((1.5 * 3.14159))))
                                                                _ = v35
                                                                _ = v36
                                                                _ = v37
                                                                _ = v38
                                                            } else {
                                                                v16 = v16
                                                                if true {
                                                                    switch "" {
                                                                    case "\n\t":
                                                                        _ = len(v3)
                                                                    case "":
                                                                        v39 := GenS8{F0: 42, F1: 1, F2: 544}
                                                                        _ = v39
                                                                    case "a":
                                                                        v5 = (v5 | 65535)
                                                                    default:
                                                                        _ = len(v3)
                                                                    }
                                                                    delete(v3, 1)
                                                                    v16 = v16
                                                                }
                                                                if false {
                                                                    var v40 int8 = 100
                                                                    fmt.Println(v21, v4, v5)
                                                                    _ = v40
                                                                } else {
                                                                    v41 := GenS4{F0: 32767, F1: 798, F2: 42}
                                                                    if (255 >= 90) {
                                                                        v4 = (v4 & v4)
                                                                    } else {
                                                                        v4 += (-(v4))
                                                                        switch v21 {
                                                                        case 32767:
                                                                            if len(v2) >= 2 {
                                                                                v2 = v2[1:len(v2)]
                                                                            }
                                                                            switch ("\n\t" + "foo bar") {
                                                                            case "a":
                                                                                v21 = fi34
                                                                            case "\n\t":
                                                                                fmt.Println(v4, v16)
                                                                            case "foo bar":
                                                                                _ = len(v1)
                                                                            default:
                                                                                v2 = append(v2, 1)
                                                                            }
                                                                            var v42 int8 = 0
                                                                            _ = v42
                                                                        default:
                                                                            v43 := GenS7{F0: 65535, F1: 4294967295}
                                                                            _ = v43
                                                                        }
                                                                    }
                                                                    _ = v41
                                                                }
                                                            }
                                                            v16 = (v16 + (v16 ^ v16))
                                                            v16 -= -1
                                                            _ = fi34
                                                        }
                                                        for fi44 := 0; fi44 < 19; fi44++ {
                                                            var v45 uint64 = 996
                                                            _ = fi44
                                                            _ = v45
                                                        }
                                                    }
                                                    v21 += (-(v21))
                                                }
                                            default:
                                                v21 -= (v21 >> ((v5) & 63))
                                            }
                                        case "foo bar":
                                            v16 = (v16 & 1000)
                                        }
                                    case 42:
                                        for fi46 := 0; fi46 < 9; fi46++ {
                                            switch (fi46 & 32767) {
                                            case -926:
                                                if (127 != (fi46 / ((fi46) | 1))) {
                                                    fmt.Println(fi17, v16, v21)
                                                    v21 += (fi46 + fi46)
                                                    v4 += v4
                                                    v16 -= -389
                                                    v4 = 100
                                                } else {
                                                    if true {
                                                        v47 := GenS8{F0: 32767, F1: 218, F2: 18446744073709551615}
                                                        v16 = (v16 / ((v16) | 1))
                                                        var v48 uint32 = 4294967295
                                                        _ = v47
                                                        _ = v48
                                                    }
                                                    var v49 bool = (!false)
                                                    fmt.Println(fi17)
                                                    _ = v49
                                                }
                                            }
                                            v50 := GenS1{F0: -128}
                                            var v51 uint32 = 79
                                            _ = fi46
                                            _ = v50
                                            _ = v51
                                        }
                                    }
                                    _ = v20
                                    _ = v21
                                } else {
                                    var v52 int8 = v4
                                    _ = v52
                                }
                                v4 += v4
                            }
                        }
                        var v53 uint8 = 42
                        _ = v19
                        _ = v53
                    } else {
                        switch ("" + "foo bar") {
                        case "hello":
                            v54 := GenS9{F0: 42, F1: 65535}
                            _ = v54
                        default:
                            fmt.Println(v4, v16, v5)
                        }
                    }
                }
            default:
                for fi55 := 0; fi55 < 10; fi55++ {
                    v56 := GenS5{F0: 4294967295, F1: 2}
                    _ = fi55
                    _ = v56
                }
            }
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        } else {
            delete(v1, "\n\t")
        }
        _ = fi17
    }
    var v57 int = 32768
    v4 -= (v4 >> (((v5 + (v5 ^ 18446744073709551615))) & 63))
    v5 = (^(v5))
    delete(v3, v57)
    v16 -= 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v57:", v57)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v57
}
