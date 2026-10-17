package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64; F2 uint }
type GenS2 struct { F0 uint8; F1 bool; F2 int }
type GenS3 struct { F0 uint; F1 uint32; F2 uint8 }
type GenS4 struct { F0 bool; F1 uint8 }
type GenS5 struct { F0 int32; F1 uint8; F2 uint8 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 int32; F1 int }
type GenS8 struct { F0 int32; F1 int }
type GenS9 struct { F0 int32; F1 int64; F2 uint64 }
func genFn0() uint32 {
    var v6 int64 = 9223372036854775807
    v7 := GenS1{F0: 227, F1: 42, F2: 65535}
    for fi8 := 0; fi8 < 15; fi8++ {
        switch 32767 {
        case -1:
            if true {
                var v9 float64 = ((1.5 - (0.0 - (1.5 + 1e10))) / 0.1)
                var v10 int16 = -32768
                v11 := GenS2{F0: 0, F1: true, F2: (fi8 >> ((2) & 63))}
                _ = v10
                _ = v11
                _ = v9
            } else {
                var v12 float32 = 3.14159
                _ = v12
            }
        case 318:
            var v13 int8 = -43
            _ = v13
        default:
            v14 := []int{fi8, (^(fi8)), fi8, fi8}
            _ = v14
        }
        v15 := []int{32768, fi8}
        _ = fi8
        _ = v15
    }
    var v16 uint32 = 510
    _ = v16
    _ = v6
    _ = v7
    return v16
}

func main() {
    var v1 int = 127
    var v2 uint16 = 32767
    var v3 int8 = -128
    var v4 float32 = (3.14159 + 0.1)
    var v5 uint = 31
    v2 += 0
    v17 := []int{}
    _ = len(v17)
    var v18 int64 = -9223372036854775808
    for fi19 := 0; fi19 < 14; fi19++ {
        var v20 string = ""
        _ = fi19
        _ = v20
    }
    for fi21 := 0; fi21 < 15; fi21++ {
        var v22 int = (fi21 / ((fi21) | 1))
        if true {
            switch (v1 << ((v5) & 63)) {
            case 2147483647:
                v23 := GenS3{F0: v5, F1: 4294967295, F2: 1}
                _ = v23
            case 4294967295:
                switch int(v5) {
                case 42:
                    switch "\n\t" {
                    case "hello":
                        var v24 string = ""
                        _ = v24
                    case "a":
                        v4 = 0.1
                    default:
                        var v25 float32 = -2.25
                        _ = v25
                    }
                case 1:
                    var v26 bool = ((v5 >> ((v5) & 63)) >= v5)
                    _ = v26
                }
            case -1:
                v22 = v1
            default:
                switch (v22 & fi21) {
                case -99:
                    var v27 int16 = 1
                    _ = v27
                default:
                    v28 := GenS2{F0: 1, F1: (7 != 2), F2: (v1 / ((fi21) | 1))}
                    _ = v28
                }
            }
            for fi29 := 0; fi29 < 10; fi29++ {
                var v30 uint16 = v2
                _ = fi29
                _ = v30
            }
            v17 = append(v17, 0)
        } else {
            v18 = (v18 | (^(v18)))
            var v31 uint = v5
            var v32 int64 = 1
            _ = v31
            _ = v32
        }
        var v33 uint16 = v2
        _ = fi21
        _ = v22
        _ = v33
    }
    v18 -= v18
    var v34 bool = false
    if len(v17) > 0 {
        v17[0] = (-(v1))
    }
    for fi35 := 0; fi35 < 8; fi35++ {
        for fi36 := 0; fi36 < 3; fi36++ {
            v1 = fi36
            if v34 {
                v3 = int8(v5)
            } else {
                v37 := GenS4{F0: false, F1: 1}
                _ = v37
            }
            v1 = int(v18)
            _ = fi36
        }
        if v34 {
            var v38 bool = (((^(v5)) > v5) && false)
            v39 := GenS5{F0: -2147483648, F1: 0, F2: 42}
            for fi40 := 0; fi40 < 10; fi40++ {
                var v41 uint32 = 2147483647
                var v42 int64 = (-(v18))
                _ = fi40
                _ = v41
                _ = v42
            }
            _ = v38
            _ = v39
        } else {
            switch (v1 / (((fi35 & 0)) | 1)) {
            case 2147483647:
                fmt.Println(v2)
            }
        }
        _ = fi35
    }
    if len(v17) > 0 {
        v17[0] = v1
    }
    if v34 {
        var v43 float32 = (v4 * v4)
        _ = v43
    } else {
        for fi44 := 0; fi44 < 6; fi44++ {
            v17 = append(v17, 32768)
            v45 := GenS1{F0: v5, F1: 1, F2: (v5 & v5)}
            _ = fi44
            _ = v45
        }
    }
    for fi46 := 0; fi46 < 1; fi46++ {
        switch (("a" + "hello") + "foo bar") {
        case "\n\t":
            v47 := GenS6{F0: 2147483648}
            _ = v47
        case "hello":
            v3 = (v3 - v3)
        }
        v48 := GenS7{F0: 0, F1: v1}
        switch (fi46 >> ((18446744073709551615) & 63)) {
        case 32767:
            for fi49 := 0; fi49 < 20; fi49++ {
                switch ("a" + "foo bar") {
                case "hello":
                    if (0 > -1) {
                        if v34 {
                            var v50 bool = (v18 != v18)
                            v2 += v2
                            if v50 {
                                v2 -= v2
                                var v51 uint8 = 42
                                if false {
                                    var v52 float64 = 1e10
                                    v53 := GenS8{F0: 0, F1: (fi49 | (fi49 & fi49))}
                                    v17 = append(v17, (fi49 - (fi46 / ((fi46) | 1))))
                                    _ = v52
                                    _ = v53
                                } else {
                                    var v54 float32 = v4
                                    _ = v54
                                }
                                _ = v51
                            }
                            var v55 int16 = 1
                            _ = v50
                            _ = v55
                        }
                        if v34 {
                            v1 = v1
                            var v56 uint8 = 128
                            switch v1 {
                            case 128:
                                var v57 float64 = 0.1
                                _ = v57
                            case 9223372036854775807:
                                switch "a" {
                                case "\n\t":
                                    v17 = append(v17, (^(fi46)))
                                case "":
                                    v58 := GenS2{F0: 48, F1: (((float32(1e10) >= v4) || v34) || ((v5 > 42) || (v34 || v34))), F2: v1}
                                    _ = v58
                                case "a":
                                    if len(v17) > 0 {
                                        v17[0] = (fi46 << ((uint(v2)) & 63))
                                    }
                                default:
                                    fmt.Println(v3, v1)
                                }
                            case 32768:
                                if v34 {
                                    var v59 uint64 = uint64(v56)
                                    var v60 int = fi46
                                    _ = len(v17)
                                    v60 = v60
                                    fmt.Println(v34)
                                    _ = v59
                                    _ = v60
                                }
                            default:
                                v1 -= (^(fi49))
                                v61 := GenS9{F0: 2147483647, F1: (v18 | v18), F2: 437}
                                if len(v17) >= 2 {
                                    v17 = v17[1:len(v17)]
                                }
                                _ = v61
                            }
                            _ = v56
                        } else {
                            if len(v17) >= 2 {
                                v17 = v17[1:len(v17)]
                            }
                            v3 -= (v3 >> ((210) & 63))
                        }
                        var v62 int32 = -2147483648
                        _ = v62
                    }
                }
                var v63 float32 = 3.14159
                _ = fi49
                _ = v63
            }
        case -435:
            for fi64 := 0; fi64 < 3; fi64++ {
                var v65 string = ("foo bar" + ("hello" + ("\n\t" + "a")))
                var v66 uint = (v5 >> ((4294967295) & 63))
                _ = fi64
                _ = v65
                _ = v66
            }
        }
        _ = fi46
        _ = v48
    }
    v34 = (!v34)
    v67 := GenS4{F0: v34, F1: 127}
    _ = len(v17)
    if v34 {
        fmt.Println(v1, v5)
        v34 = (798 <= 91)
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v67
}
