package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int64; F1 uint8 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int; F1 uint16; F2 int64 }
type GenS5 struct { F0 uint32; F1 int; F2 uint64 }
type GenS6 struct { F0 int32; F1 int8; F2 int }
func genFn0(p0 int8, p1 uint16, p2 int16) int {
    p1 += (p1 * p1)
    v5 := GenS1{F0: 127}
    v6 := GenS2{F0: -9223372036854775808, F1: 2}
    if (3.14159 < (((3.14159 + float64(1e10)) - (float64(3.14159) / -2.25)) - -2.25)) {
        v7 := []int{2147483647}
        var v8 uint8 = 128
        v9 := GenS3{F0: 0}
        _ = v7
        _ = v8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 1
}
func genFn1() uint16 {
    if false {
        var v10 uint8 = 2
        var v11 int64 = 233
        _ = v10
        _ = v11
    }
    var v12 uint32 = 200
    var v13 int = 32767
    switch int(v12) {
    case -885:
        fmt.Println(v12)
    case -1:
        switch (("foo bar" + ("a" + "a")) + "\n\t") {
        case "hello":
            v14 := GenS1{F0: 0}
            _ = v14
        case "\n\t":
            v12 = v12
        default:
            v12 = (v12 % (((v12 | v12)) | 1))
        }
    }
    _ = v12
    _ = v13
    return 1
}

func main() {
    var v1 uint16 = 65535
    var v2 int = 749
    var v3 int16 = int16(v1)
    var v4 uint = 0
    if (-128 != -63) {
        v15 := GenS3{F0: 601}
        switch ("" + "\n\t") {
        case "foo bar":
            switch "foo bar" {
            case "foo bar":
                var v16 string = "foo bar"
                _ = v16
            case "a":
                for fi17 := 0; fi17 < 7; fi17++ {
                    v2 += fi17
                    v18 := []int{fi17, v2}
                    fmt.Println(fi17)
                    _ = fi17
                    _ = v18
                }
            default:
                var v19 int = 0
                _ = v19
            }
        case "hello":
            var v20 float64 = (((((1e10 / (((1e10 / 3.14159) - float64(1e10)) / -2.25)) - float64(0.1)) * 1e10) * (1e10 / (((3.14159 / 1.5) / float64((float32(1.5) * 1.5))) / 1.5))) * 0.0)
            _ = v20
        case "\n\t":
            var v21 float32 = (-2.25 / 3.14159)
            _ = v21
        }
        _ = v15
    } else {
        v1 -= (^(v1))
        fmt.Println(v3)
        var v22 bool = true
        _ = v22
    }
    var v23 int = (v2 - v2)
    var v24 int32 = -386
    for fi25 := 0; fi25 < 19; fi25++ {
        var v26 uint16 = (v1 * v1)
        var v27 int64 = 9223372036854775807
        _ = fi25
        _ = v26
        _ = v27
    }
    v28 := []int{(v23 ^ genFn0(127, v1, v3)), genFn0(1, v1, -32768)}
    if true {
        if (!(65535 != 492)) {
            if false {
                _ = len(v28)
                for fi29 := 0; fi29 < 17; fi29++ {
                    fmt.Println(v1, fi29, v23)
                    _ = fi29
                }
            } else {
                v24 = v24
            }
            v4 = (^(v4))
        }
        if len(v28) >= 2 {
            v28 = v28[1:len(v28)]
        }
        switch genFn0(1, v1, (v3 / (((-(v3))) | 1))) {
        case 127:
            if (2 >= uint8(v1)) {
                v30 := GenS4{F0: v2, F1: 32767, F2: 42}
                _ = v30
            } else {
                for fi31 := 0; fi31 < 3; fi31++ {
                    if ((42 > 127) && ((!(!true)) && ((v4 < v4) && false))) {
                        for fi32 := 0; fi32 < 13; fi32++ {
                            var v33 float64 = -2.25
                            _ = fi32
                            _ = v33
                        }
                    } else {
                        if false {
                            switch "" {
                            case "":
                                v34 := GenS2{F0: 977, F1: 42}
                                _ = v34
                            case "hello":
                                v35 := GenS2{F0: -9223372036854775808, F1: 0}
                                _ = v35
                            case "a":
                                var v36 float64 = 0.0
                                _ = v36
                            }
                            for fi37 := 0; fi37 < 18; fi37++ {
                                var v38 uint32 = 838
                                v2 = genFn0(-128, (v1 | v1), (v3 - (v3 | -32768)))
                                v39 := GenS5{F0: v38, F1: v2, F2: 0}
                                v40 := GenS6{F0: v24, F1: 100, F2: fi37}
                                _ = fi37
                                _ = v38
                                _ = v39
                                _ = v40
                            }
                        }
                    }
                    v28 = append(v28, v23)
                    _ = fi31
                }
                v1 = v1
            }
        case -2147483648:
            var v41 int64 = 1
            _ = v41
        case 35:
            var v42 int = (v23 ^ v23)
            _ = v42
        default:
            for fi43 := 0; fi43 < 2; fi43++ {
                var v44 string = "\n\t"
                _ = fi43
                _ = v44
            }
        }
    } else {
        v28 = append(v28, (v2 << (((v4 - v4)) & 63)))
    }
    for fi45 := 0; fi45 < 18; fi45++ {
        for fi46 := 0; fi46 < 2; fi46++ {
            var v47 uint = (v4 / ((v4) | 1))
            _ = fi46
            _ = v47
        }
        var v48 int8 = 62
        _ = len(v28)
        _ = fi45
        _ = v48
    }
    var v49 int8 = 42
    if true {
        switch genFn0(1, (v1 | genFn1()), v3) {
        case -1:
            var v50 uint32 = 754
            _ = v50
        default:
            var v51 uint = v4
            _ = v51
        }
        for fi52 := 0; fi52 < 12; fi52++ {
            var v53 uint16 = (v1 << (((v4 - (v4 % (((v4 << (((^(v4))) & 63))) | 1)))) & 63))
            var v54 int = 127
            _ = fi52
            _ = v53
            _ = v54
        }
    }
    for fi55 := 0; fi55 < 19; fi55++ {
        for fi56 := 0; fi56 < 15; fi56++ {
            var v57 uint = (v4 - 0)
            _ = fi56
            _ = v57
        }
        if (!(0 < 42)) {
            for fi58 := 0; fi58 < 19; fi58++ {
                switch genFn0(v49, (v1 / ((v1) | 1)), v3) {
                case 4294967295:
                    var v59 int32 = 2147483647
                    _ = v59
                case 32768:
                    if true {
                        if (0.1 >= 0.1) {
                            fmt.Println(v4)
                        }
                        switch 0 {
                        case 128:
                            v60 := GenS2{F0: -46, F1: 128}
                            _ = v60
                        default:
                            var v61 int16 = (-(v3))
                            _ = v61
                        }
                        v2 += fi58
                        v1 -= (v1 + (v1 & v1))
                        var v62 uint8 = 42
                        _ = v62
                    }
                default:
                    v63 := GenS6{F0: v24, F1: (v49 ^ -1), F2: (fi55 % ((genFn0((^(v49)), v1, (v3 ^ (v3 >> ((65535) & 63))))) | 1))}
                    _ = v63
                }
                _ = fi58
            }
            v28 = append(v28, (^(v23)))
        }
        _ = fi55
    }
    var v64 int8 = (v49 | 0)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    fmt.Println("v64:", v64)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v28
    _ = v3
    _ = v4
    _ = v49
    _ = v64
}
