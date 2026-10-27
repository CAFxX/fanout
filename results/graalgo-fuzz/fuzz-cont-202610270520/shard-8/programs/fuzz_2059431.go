package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint32; F1 int32; F2 uint32 }
type GenS4 struct { F0 int64; F1 uint16; F2 int16 }
type GenS5 struct { F0 uint32; F1 int }
type GenS6 struct { F0 uint32; F1 int32 }
type GenS7 struct { F0 int32; F1 uint16 }
type GenS8 struct { F0 int64; F1 uint64; F2 uint64 }
func genFn0(p0 int) bool {
    v5 := []int{(p0 % ((p0) | 1)), p0}
    var v6 string = ("foo bar" + "hello")
    _ = p0
    _ = v5
    _ = v6
    return (-32768 < 794)
}
func genFn1(p0 int16) uint {
    if true {
        var v7 float64 = 1e10
        v7 -= float64((1.5 * 0.0))
        v8 := []int{}
        _ = v7
        _ = v8
    }
    var v9 uint8 = 128
    if false {
        if false {
            v10 := GenS1{F0: 0, F1: 127}
            if false {
                v11 := []int{}
                var v12 uint32 = 206
                var v13 int64 = 0
                _ = v11
                _ = v12
                _ = v13
            } else {
                var v14 uint = 969
                v15 := GenS1{F0: 0, F1: -86}
                _ = v14
                _ = v15
            }
            _ = v10
        }
        if true {
            v16 := []int{4294967295, 42, -1, -1}
            var v17 int16 = -1
            _ = len(v16)
            _ = v16
            _ = v17
        }
        fmt.Println(v9, p0)
    } else {
        var v18 uint16 = 343
        v18 -= 144
        _ = v18
    }
    var v19 int64 = -1
    _ = p0
    _ = v19
    _ = v9
    return 255
}

func main() {
    var v1 float64 = 1.5
    var v2 int64 = 42
    var v3 int64 = (v2 - v2)
    var v4 int32 = 432
    v20 := GenS2{F0: 0}
    for fi21 := 0; fi21 < 6; fi21++ {
        switch fi21 {
        case 128:
            v22 := []int{(fi21 / ((fi21) | 1)), (fi21 - fi21), 4294967295}
            _ = v22
        case 90:
            for fi23 := 0; fi23 < 20; fi23++ {
                var v24 int8 = -128
                var v25 uint64 = 2
                _ = fi23
                _ = v24
                _ = v25
            }
        default:
            v26 := GenS1{F0: 173, F1: -101}
            _ = v26
        }
        _ = fi21
    }
    var v27 uint = 4294967295
    v28 := GenS2{F0: 2}
    var v29 uint64 = 2
    switch "foo bar" {
    case "":
        if true {
            for fi30 := 0; fi30 < 13; fi30++ {
                if true {
                    var v31 int8 = 100
                    _ = v31
                }
                if true {
                    v32 := GenS2{F0: (v29 * 174)}
                    v33 := []int{}
                    switch "" {
                    case "":
                        var v34 uint32 = 2147483648
                        _ = v34
                    case "hello":
                        var v35 float32 = 0.0
                        _ = v35
                    default:
                        if len(v33) >= 2 {
                            v33 = v33[1:len(v33)]
                        }
                    }
                    _ = v32
                    _ = v33
                } else {
                    for fi36 := 0; fi36 < 2; fi36++ {
                        var v37 int = fi36
                        if false {
                            fmt.Println(v2, v2)
                            v38 := GenS3{F0: uint32(v29), F1: v4, F2: 41}
                            v39 := GenS3{F0: 4294967295, F1: 2147483647, F2: 65535}
                            _ = v38
                            _ = v39
                        }
                        v40 := GenS3{F0: 0, F1: v4, F2: 4294967295}
                        _ = fi36
                        _ = v37
                        _ = v40
                    }
                }
                _ = fi30
            }
            var v41 bool = (true && ((106 <= 100) || true))
            _ = v41
        } else {
            var v42 uint32 = 65535
            v2 = (v2 >> ((v27) & 63))
            for fi43 := 0; fi43 < 13; fi43++ {
                v2 += (v2 << ((v27) & 63))
                v44 := GenS4{F0: v3, F1: 32767, F2: 1000}
                _ = fi43
                _ = v44
            }
            _ = v42
        }
    case "\n\t":
        var v45 uint8 = 1
        _ = v45
    default:
        for fi46 := 0; fi46 < 14; fi46++ {
            v47 := GenS4{F0: v3, F1: 445, F2: -4}
            switch fi46 {
            case -275:
                if true {
                    v48 := GenS5{F0: 849, F1: fi46}
                    if (false && (2 >= 1)) {
                        v4 = v4
                        if false {
                            v49 := []int{}
                            v50 := GenS5{F0: 2147483647, F1: fi46}
                            _ = v49
                            _ = v50
                        }
                    } else {
                        if true {
                            fmt.Println(v1, v27, v3)
                        } else {
                            var v51 int32 = (-(v4))
                            _ = v51
                        }
                        var v52 uint = (v27 | genFn1(1000))
                        _ = v52
                    }
                    for fi53 := 0; fi53 < 7; fi53++ {
                        var v54 uint = genFn1(1000)
                        _ = fi53
                        _ = v54
                    }
                    _ = v48
                } else {
                    v55 := []int{fi46, 0, fi46}
                    _ = v55
                }
            case -2147483648:
                for fi56 := 0; fi56 < 8; fi56++ {
                    v29 -= (v29 >> (((v27 * v27)) & 63))
                    if true {
                        v57 := []int{}
                        _ = v57
                    } else {
                        if true {
                            v58 := []int{4294967295, (fi46 ^ fi56)}
                            v27 = (v27 - v27)
                            _ = len(v58)
                            _ = v58
                        }
                        var v59 float64 = float64(-2.25)
                        var v60 uint8 = 0
                        switch "foo bar" {
                        case "":
                            if ((v29 != (v29 * uint64(v60))) || false) {
                                v61 := GenS6{F0: 4294967295, F1: -2147483648}
                                v62 := GenS1{F0: v4, F1: 0}
                                _ = v61
                                _ = v62
                            } else {
                                v63 := GenS4{F0: (^(v2)), F1: 42, F2: -164}
                                var v64 float64 = v59
                                v65 := GenS7{F0: (-(v4)), F1: 42}
                                _ = v63
                                _ = v64
                                _ = v65
                            }
                        case "a":
                            v66 := []int{(fi46 * (^(fi56))), fi56}
                            _ = v66
                        default:
                            fmt.Println(v3, v27, v2)
                        }
                        _ = v59
                        _ = v60
                    }
                    _ = fi56
                }
            }
            fmt.Println(v3, v3, v29)
            _ = fi46
            _ = v47
        }
    }
    fmt.Println(v29, v27)
    v67 := GenS1{F0: v4, F1: 100}
    var v68 float64 = v1
    v69 := []int{2147483647, int(v4)}
    if false {
        var v70 int16 = -481
        _ = v70
    } else {
        if len(v69) > 0 {
            v69[0] = 128
        }
        if len(v69) >= 2 {
            v69 = v69[1:len(v69)]
        }
        v68 = (3.14159 + -2.25)
    }
    var v71 uint32 = 2147483648
    var v72 uint = 595
    for fi73 := 0; fi73 < 1; fi73++ {
        for fi74 := 0; fi74 < 18; fi74++ {
            v75 := GenS5{F0: (v71 / ((1) | 1)), F1: (fi74 * 128)}
            _ = fi74
            _ = v75
        }
        _ = fi73
    }
    switch "\n\t" {
    case "\n\t":
        switch 9223372036854775807 {
        case -9223372036854775808:
            v76 := GenS8{F0: (v2 % (((v2 & v3)) | 1)), F1: (v29 << (((v27 + (v72 + v27))) & 63)), F2: (^(v29))}
            _ = v76
        case 128:
            for fi77 := 0; fi77 < 15; fi77++ {
                v78 := GenS1{F0: (v4 << ((v72) & 63)), F1: -84}
                _ = fi77
                _ = v78
            }
        default:
            v68 = (v68 * (v68 + -2.25))
        }
    case "foo bar":
        switch 127 {
        case 4294967295:
            var v79 uint8 = 127
            _ = v79
        }
    case "a":
        v4 = v4
    default:
        switch "foo bar" {
        case "foo bar":
            var v80 int16 = 448
            _ = v80
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v68:", v68)
    fmt.Println("slicelen v69:", len(v69), v69)
    fmt.Println("v71:", v71)
    fmt.Println("v72:", v72)
    _ = v1
    _ = v2
    _ = v20
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v67
    _ = v68
    _ = v69
    _ = v71
    _ = v72
}
