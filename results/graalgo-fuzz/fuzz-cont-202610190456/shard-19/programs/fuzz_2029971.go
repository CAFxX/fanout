package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int64; F1 uint8 }
type GenS3 struct { F0 uint64; F1 uint32; F2 uint8 }
func genFn0(p0 int64, p1 int32, p2 uint) uint32 {
    var v7 int16 = 42
    for fi8 := 0; fi8 < 1; fi8++ {
        fmt.Println(p0, p2)
        p0 = (p0 << (((^(p2))) & 63))
        switch 32767 {
        case -2147483648:
            for fi9 := 0; fi9 < 17; fi9++ {
                var v10 int64 = (p0 / ((289) | 1))
                _ = fi9
                _ = v10
            }
        case 674:
            p2 = (^(p2))
        case 128:
            switch (fi8 + (fi8 + fi8)) {
            case -2147483648:
                fmt.Println(p2, fi8)
            case -24:
                for fi11 := 0; fi11 < 7; fi11++ {
                    switch (fi8 % ((fi8) | 1)) {
                    case 215:
                        v12 := []int{int(p2), fi8, 719}
                        _ = v12
                    case -160:
                        if false {
                            p1 -= p1
                            p1 += (p1 >> ((p2) & 63))
                            p2 -= p2
                        }
                    default:
                        v13 := []int{fi8, -40, fi11}
                        _ = v13
                    }
                    _ = fi11
                }
            case 0:
                v14 := GenS1{F0: -1}
                _ = v14
            default:
                var v15 int = (-(fi8))
                _ = v15
            }
        }
        _ = fi8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 42
}
func genFn1(p0 uint) int8 {
    for fi16 := 0; fi16 < 10; fi16++ {
        if false {
            var v17 uint8 = 255
            p0 += 18446744073709551615
            _ = v17
        }
        _ = fi16
    }
    var v18 uint = p0
    v19 := []int{4294967295, -999, -268}
    _ = p0
    _ = v18
    _ = v19
    return -8
}

func main() {
    v1 := make(map[string]string)
    var v2 uint64 = 69
    var v3 string = "foo bar"
    v4 := make(map[uint32]string)
    v5 := make(map[uint32]uint32)
    var v6 int64 = 9223372036854775807
    var v20 int8 = 0
    for fi21 := 0; fi21 < 19; fi21++ {
        var v22 int8 = (v20 | genFn1(0))
        _ = fi21
        _ = v22
    }
    switch 681 {
    case -1:
        switch "" {
        case "a":
            v6 = v6
        case "hello":
            switch v3 {
            case "foo bar":
                switch -9223372036854775808 {
                case -2147483648:
                    v3 = "\n\t"
                }
            case "a":
                v20 = -89
            default:
                var v23 uint = 2
                _ = v23
            }
        default:
            var v24 uint = 4294967295
            _ = v24
        }
    case 42:
        var v25 int8 = (v20 % ((v20) | 1))
        _ = v25
    default:
        for fi26 := 0; fi26 < 12; fi26++ {
            for fi27 := 0; fi27 < 16; fi27++ {
                var v28 int64 = v6
                switch v3 {
                case "hello":
                    v20 -= 100
                case "":
                    v29 := GenS1{F0: 0}
                    _ = v29
                }
                _ = fi27
                _ = v28
            }
            v30 := GenS2{F0: 45, F1: 127}
            var v31 bool = ((^(v20)) > v20)
            _ = fi26
            _ = v30
            _ = v31
        }
    }
    var v32 int32 = -156
    v3 = "foo bar"
    switch 128 {
    case -1:
        v33 := GenS1{F0: (v20 & (v20 * (v20 ^ (v20 - (v20 & v20)))))}
        _ = v33
    case 522:
        switch 2147483647 {
        case 127:
            v3 = v3
        case 1:
            v6 += 42
        default:
            for fi34 := 0; fi34 < 14; fi34++ {
                var v35 int32 = (v32 << ((65535) & 63))
                v20 = genFn1(65535)
                _ = fi34
                _ = v35
            }
        }
    default:
        v36 := GenS2{F0: (-(v6)), F1: 255}
        _ = v36
    }
    switch -484 {
    case -9223372036854775808:
        v37 := []int{9223372036854775807, -2147483648, 0}
        _ = v37
    }
    var v38 int = -9223372036854775808
    v2 += (v2 & 18446744073709551615)
    v20 -= v20
    v39 := GenS1{F0: (v20 & 14)}
    fmt.Println(v3, v38, v2)
    _ = len(v1)
    if false {
        var v40 uint16 = 943
        var v41 uint = 339
        _ = v40
        _ = v41
    } else {
        switch "\n\t" {
        case "a":
            var v42 int8 = v20
            _ = v42
        default:
            v32 = v32
        }
        if (1000 != -32768) {
            if true {
                v6 += v6
            }
            v43 := []int{int(v6)}
            if ((!(v20 >= 1)) && false) {
                fmt.Println(v20, v38)
            }
            _ = v43
        } else {
            v4[genFn0(v6, (v32 % ((v32) | 1)), uint(v2))] = ""
            v38 += (v38 | (v38 | 32768))
        }
        switch v3 {
        case "":
            v44 := GenS1{F0: (v20 / (((v20 / (((v20 ^ genFn1(527))) | 1))) | 1))}
            _ = v44
        default:
            if false {
                var v45 int64 = (v6 ^ v6)
                var v46 uint32 = 2147483647
                _ = v45
                _ = v46
            }
        }
    }
    fmt.Println(v20, v32)
    var v47 int32 = -1
    var v48 float32 = (((1e10 + 3.14159) - -2.25) + (((3.14159 / 1.0) * 1.5) + 3.14159))
    switch v3 {
    case "hello":
        _ = len(v4)
    case "":
        if (float32((1.5 - 3.14159)) >= float32(float64((((v48 - -2.25) * v48) - float32(3.14159))))) {
            for fi49 := 0; fi49 < 3; fi49++ {
                _ = len(v1)
                _ = fi49
            }
            var v50 int64 = v6
            if (181 == 4294967295) {
                for fi51 := 0; fi51 < 17; fi51++ {
                    var v52 uint8 = 255
                    fmt.Println(v38, v52)
                    _ = fi51
                    _ = v52
                }
            } else {
                v38 = -2147483648
                fmt.Println(v3, v47, v2)
            }
            _ = v50
        }
    case "\n\t":
        v38 = v38
    default:
        for fi53 := 0; fi53 < 5; fi53++ {
            var v54 uint = 1
            v55 := []int{(v38 + int(v20)), 292, fi53}
            v56 := GenS3{F0: (v2 & v2), F1: 42, F2: 76}
            _ = fi53
            _ = v54
            _ = v55
            _ = v56
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v32
    _ = v38
    _ = v39
    _ = v4
    _ = v47
    _ = v48
    _ = v5
    _ = v6
}
