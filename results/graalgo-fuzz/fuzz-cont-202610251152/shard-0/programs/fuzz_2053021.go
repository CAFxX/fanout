package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16 }
type GenS2 struct { F0 uint8; F1 int8; F2 uint16 }
type GenS3 struct { F0 uint16; F1 int; F2 uint16 }
type GenS4 struct { F0 int64 }
func genFn0(p0 uint8) uint32 {
    fmt.Println(p0)
    for fi4 := 0; fi4 < 8; fi4++ {
        fmt.Println(fi4, p0)
        p0 = p0
        fmt.Println(p0)
        _ = fi4
    }
    v5 := []int{2147483647, 931, 950}
    v6 := GenS1{F0: (false && false), F1: 1000}
    for fi7 := 0; fi7 < 13; fi7++ {
        for fi8 := 0; fi8 < 5; fi8++ {
            v9 := GenS2{F0: (p0 + 42), F1: 42, F2: 65535}
            _ = fi8
            _ = v9
        }
        _ = fi7
    }
    _ = p0
    _ = v5
    _ = v6
    return 42
}
func genFn1(p0 int32) uint32 {
    if (42 <= 42) {
        for fi10 := 0; fi10 < 2; fi10++ {
            var v11 uint8 = 0
            _ = fi10
            _ = v11
        }
    } else {
        v12 := GenS1{F0: (565 == 4294967295), F1: -32768}
        if false {
            var v13 uint = 255
            v14 := []int{}
            _ = v13
            _ = v14
        } else {
            p0 += p0
            for fi15 := 0; fi15 < 16; fi15++ {
                p0 -= -317
                var v16 uint8 = 146
                v17 := []int{}
                _ = fi15
                _ = v16
                _ = v17
            }
        }
        _ = v12
    }
    if true {
        v18 := []int{42, 811, 127}
        _ = v18
    } else {
        p0 += p0
        for fi19 := 0; fi19 < 4; fi19++ {
            v20 := GenS3{F0: 2, F1: fi19, F2: 42}
            switch -9223372036854775808 {
            case 42:
                if false {
                    p0 = p0
                    var v21 uint8 = 128
                    var v22 uint8 = (v21 | v21)
                    _ = v21
                    _ = v22
                }
            default:
                p0 += p0
            }
            switch fi19 {
            case 9223372036854775807:
                v23 := []int{521, (-(fi19)), fi19}
                _ = v23
            case -27:
                if ("a" == "hello") {
                    for fi24 := 0; fi24 < 16; fi24++ {
                        var v25 int16 = 1
                        var v26 float32 = 1.5
                        var v27 bool = false
                        _ = fi24
                        _ = v25
                        _ = v26
                        _ = v27
                    }
                    var v28 uint = 466
                    switch fi19 {
                    case 678:
                        v29 := []int{(fi19 * fi19), fi19}
                        _ = v29
                    case 4294967295:
                        var v30 int64 = -9223372036854775808
                        _ = v30
                    case 9223372036854775807:
                        v28 = v28
                    }
                    _ = v28
                }
            default:
                for fi31 := 0; fi31 < 6; fi31++ {
                    fmt.Println(fi19)
                    _ = fi31
                }
            }
            _ = fi19
            _ = v20
        }
    }
    p0 = -1
    fmt.Println(p0)
    _ = p0
    return 551
}

func main() {
    v1 := []uint32{2147483647, 238, 65535, 42}
    var v2 bool = (false && true)
    var v3 int = 4294967295
    if (!v2) {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        v2 = v2
        fmt.Println(v3, v2)
    }
    if true {
        fmt.Println(v3, v3)
    } else {
        for fi32 := 0; fi32 < 18; fi32++ {
            var v33 int16 = 980
            var v34 uint8 = 235
            v2 = false
            _ = fi32
            _ = v33
            _ = v34
        }
    }
    if false {
        for fi35 := 0; fi35 < 14; fi35++ {
            if true {
                var v36 int8 = -1
                v3 = v3
                _ = v36
            }
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            if len(v1) > 0 {
                v1[0] = 2147483647
            }
            _ = fi35
        }
    }
    var v37 uint8 = 191
    fmt.Println(v37, v37)
    v1 = append(v1, 1)
    v3 += (v3 / ((-67) | 1))
    var v38 string = "\n\t"
    var v39 int32 = 0
    v37 = (v37 / ((v37) | 1))
    var v40 float32 = 3.14159
    for fi41 := 0; fi41 < 2; fi41++ {
        for fi42 := 0; fi42 < 13; fi42++ {
            v2 = (900 >= 255)
            v43 := GenS1{F0: (!v2), F1: 42}
            _ = fi42
            _ = v43
        }
        if v2 {
            if v2 {
                if false {
                    _ = len(v1)
                    var v44 uint8 = 0
                    _ = v44
                } else {
                    if len(v1) > 0 {
                        v1[0] = 463
                    }
                    var v45 uint8 = (^(v37))
                    v37 -= (v37 >> ((1) & 63))
                    _ = v45
                }
            } else {
                for fi46 := 0; fi46 < 16; fi46++ {
                    v47 := GenS2{F0: v37, F1: 120, F2: 42}
                    v48 := GenS4{F0: -628}
                    _ = fi46
                    _ = v47
                    _ = v48
                }
                var v49 uint32 = 193
                _ = v49
            }
            v50 := GenS3{F0: 65535, F1: (v3 >> ((2) & 63)), F2: 65535}
            for fi51 := 0; fi51 < 13; fi51++ {
                if false {
                    var v52 int64 = 42
                    var v53 uint64 = 1
                    v54 := GenS1{F0: true, F1: 32767}
                    _ = v52
                    _ = v53
                    _ = v54
                } else {
                    var v55 uint = 4294967295
                    var v56 uint64 = 100
                    _ = v55
                    _ = v56
                }
                if false {
                    _ = len(v1)
                } else {
                    switch "hello" {
                    case "foo bar":
                        fmt.Println(v40, fi41, v37)
                    case "\n\t":
                        var v57 uint = 2
                        _ = v57
                    case "hello":
                        fmt.Println(v3)
                    default:
                        var v58 float32 = ((v40 - 3.14159) / 1.0)
                        _ = v58
                    }
                    switch (v3 / (((-(v3))) | 1)) {
                    case 168:
                        var v59 int64 = -562
                        _ = v59
                    }
                    var v60 uint = 24
                    _ = v60
                }
                _ = fi51
            }
            _ = v50
        } else {
            for fi61 := 0; fi61 < 1; fi61++ {
                v62 := GenS4{F0: 1}
                if false {
                    var v63 float64 = float64(v40)
                    fmt.Println(fi61, v37)
                    fmt.Println(fi61, fi61)
                    switch (v38 + "hello") {
                    case "hello":
                        var v64 int = (fi61 << ((2) & 63))
                        _ = v64
                    case "":
                        fmt.Println(v40, v39, fi41)
                        if v2 {
                            v3 = fi61
                        }
                    case "a":
                        var v65 uint64 = 42
                        _ = v65
                    default:
                        var v66 string = "a"
                        _ = v66
                    }
                    var v67 uint8 = (v37 ^ (^(v37)))
                    _ = v63
                    _ = v67
                } else {
                    var v68 uint32 = 0
                    _ = v68
                }
                var v69 uint = 670
                var v70 uint = v69
                var v71 uint = (v69 << ((uint(v37)) & 63))
                _ = fi61
                _ = v62
                _ = v69
                _ = v70
                _ = v71
            }
        }
        _ = fi41
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v40
}
