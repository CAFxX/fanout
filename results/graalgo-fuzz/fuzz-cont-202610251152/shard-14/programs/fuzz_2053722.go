package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 bool; F1 uint64; F2 uint16 }
type GenS3 struct { F0 uint32; F1 uint64; F2 uint8 }
type GenS4 struct { F0 uint8; F1 int8 }
type GenS5 struct { F0 int32; F1 uint64 }
func genFn0(p0 int16) uint16 {
    fmt.Println(p0)
    var v6 int16 = (p0 | -1)
    v7 := []int{323, -1}
    _ = p0
    _ = v6
    _ = v7
    return 255
}
func genFn1(p0 uint8) uint16 {
    for fi8 := 0; fi8 < 17; fi8++ {
        if true {
            switch fi8 {
            case 32768:
                v9 := GenS1{F0: 255}
                _ = v9
            case 127:
                var v10 int8 = -128
                _ = v10
            case -9223372036854775808:
                if false {
                    var v11 uint8 = p0
                    _ = v11
                } else {
                    v12 := GenS1{F0: (p0 * p0)}
                    var v13 uint8 = p0
                    _ = v12
                    _ = v13
                }
            }
            switch "hello" {
            case "foo bar":
                p0 += 0
            case "a":
                v14 := []int{(fi8 | (fi8 * -9223372036854775808)), 1, (fi8 | 1), fi8}
                _ = v14
            case "\n\t":
                var v15 string = "hello"
                _ = v15
            default:
                fmt.Println(p0, fi8)
            }
        }
        for fi16 := 0; fi16 < 11; fi16++ {
            p0 += p0
            p0 -= 128
            _ = fi16
        }
        var v17 int16 = 1
        _ = fi8
        _ = v17
    }
    if true {
        p0 = p0
    }
    for fi18 := 0; fi18 < 15; fi18++ {
        fmt.Println(fi18)
        for fi19 := 0; fi19 < 16; fi19++ {
            v20 := GenS1{F0: (^(p0))}
            if false {
                var v21 string = "hello"
                v21 = v21
                _ = v21
            } else {
                v22 := []int{(fi19 - -2147483648)}
                _ = v22
            }
            var v23 uint64 = 18446744073709551615
            _ = fi19
            _ = v20
            _ = v23
        }
        var v24 uint64 = 0
        _ = fi18
        _ = v24
    }
    v25 := []int{-95, 779, 32767}
    var v26 int = 128
    _ = p0
    _ = v25
    _ = v26
    return 32767
}

func main() {
    var v1 int16 = 496
    var v2 int8 = 37
    var v3 uint = 184
    var v4 int = 917
    v5 := []string{("hello" + ""), "\n\t"}
    for fi27 := 0; fi27 < 20; fi27++ {
        for fi28 := 0; fi28 < 8; fi28++ {
            v29 := GenS2{F0: (!false), F1: 210, F2: 65535}
            _ = len(v5)
            _ = fi28
            _ = v29
        }
        for fi30 := 0; fi30 < 20; fi30++ {
            var v31 bool = (!false)
            var v32 string = "a"
            _ = fi30
            _ = v31
            _ = v32
        }
        _ = len(v5)
        _ = fi27
    }
    var v33 float64 = 3.14159
    v3 = v3
    for fi34 := 0; fi34 < 10; fi34++ {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        var v35 uint32 = 2
        _ = fi34
        _ = v35
    }
    fmt.Println(v1, v1)
    var v36 uint64 = 641
    v4 += 1
    var v37 uint64 = v36
    switch v4 {
    case 2147483647:
        fmt.Println(v37, v1)
    case -222:
        switch 375 {
        case 433:
            if (true && (806 > 2147483648)) {
                if (32767 == (v4 >> ((18446744073709551615) & 63))) {
                    v3 = 42
                    v33 = v33
                }
            } else {
                var v38 int32 = 2147483647
                var v39 uint32 = 258
                var v40 bool = (false && false)
                _ = v38
                _ = v39
                _ = v40
            }
        default:
            v2 += v2
        }
    default:
        _ = len(v5)
    }
    for fi41 := 0; fi41 < 1; fi41++ {
        if (((v33 + v33) - (((1.5 * v33) / v33) + (v33 + (1e10 * 1.5)))) >= ((v33 - ((((((0.1 - (((1e10 + ((1.5 - ((((((v33 / v33) / 1e10) - 1.5) - -2.25) + v33) * v33)) + 0.1)) + (1e10 - 1.5)) - (v33 + (v33 / v33)))) / v33) / 1.5) / 1e10) * 3.14159) * (1.5 * v33))) * (((float64(1.5) - 0.1) + float64(0.1)) / 1.0))) {
            var v42 int16 = v1
            switch v4 {
            case 2147483647:
                v33 -= (v33 + (1.5 + v33))
            }
            v43 := GenS3{F0: 101, F1: uint64(v3), F2: 128}
            _ = v42
            _ = v43
        }
        v5 = append(v5, "a")
        _ = fi41
    }
    for fi44 := 0; fi44 < 18; fi44++ {
        var v45 int8 = (v2 >> (((^(v3))) & 63))
        _ = fi44
        _ = v45
    }
    v36 = v36
    v1 += v1
    fmt.Println(v4)
    var v46 int16 = v1
    if true {
        var v47 uint16 = 65535
        _ = v47
    } else {
        v4 -= (v4 | v4)
        if (uint16(v3) < 65535) {
            var v48 int32 = 0
            v46 = 1000
            v49 := GenS4{F0: 1, F1: (v2 % ((int8(v37)) | 1))}
            _ = v48
            _ = v49
        }
        v1 = v1
    }
    v3 += (v3 ^ v3)
    for fi50 := 0; fi50 < 13; fi50++ {
        for fi51 := 0; fi51 < 1; fi51++ {
            v5 = append(v5, "a")
            v52 := GenS2{F0: ((255 != 297) || false), F1: 0, F2: 2}
            var v53 int8 = v2
            v36 = (v37 << ((v3) & 63))
            v54 := GenS5{F0: -460, F1: (^(v36))}
            _ = fi51
            _ = v52
            _ = v53
            _ = v54
        }
        var v55 uint64 = 4294967295
        if (2 >= (^(v37))) {
            var v56 string = "a"
            v57 := GenS1{F0: 115}
            _ = v56
            _ = v57
        }
        _ = fi50
        _ = v55
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v36
    _ = v37
    _ = v4
    _ = v46
    _ = v5
}
