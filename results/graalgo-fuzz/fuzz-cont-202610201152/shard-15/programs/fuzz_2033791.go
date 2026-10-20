package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64 }
type GenS2 struct { F0 uint; F1 int32; F2 int16 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint16; F1 int }
func genFn0(p0 int32) bool {
    fmt.Println(p0)
    var v7 uint = 848
    fmt.Println(v7, p0)
    _ = p0
    _ = v7
    return ((!true) && true)
}
func genFn1(p0 bool, p1 bool, p2 uint64) uint16 {
    switch int(p2) {
    case 2147483647:
        v8 := GenS1{F0: 956, F1: p2}
        _ = v8
    case 9223372036854775807:
        var v9 string = "foo bar"
        _ = v9
    }
    fmt.Println(p1, p0, p2)
    for fi10 := 0; fi10 < 17; fi10++ {
        switch "foo bar" {
        case "a":
            p2 = (p2 >> ((797) & 63))
        default:
            var v11 bool = (255 == 0)
            _ = v11
        }
        _ = fi10
    }
    _ = p0
    _ = p1
    _ = p2
    return 42
}

func main() {
    var v1 int8 = 114
    var v2 string = "a"
    v3 := make(map[int]string)
    var v4 uint32 = 1
    var v5 uint64 = 0
    var v6 uint16 = 2
    switch "" {
    case "\n\t":
        v2 = v2
    }
    for fi12 := 0; fi12 < 13; fi12++ {
        switch 127 {
        case 0:
            v13 := GenS1{F0: 0, F1: 42}
            _ = v13
        }
        v2 = v2
        _ = fi12
    }
    if false {
        v14 := GenS1{F0: -1, F1: v5}
        _ = v14
    }
    v15 := GenS2{F0: 1, F1: -2147483648, F2: 835}
    v6 = (v6 << ((1) & 63))
    switch v2 {
    case "\n\t":
        v5 = (v5 + v5)
    case "hello":
        if false {
            for fi16 := 0; fi16 < 20; fi16++ {
                var v17 bool = false
                _ = fi16
                _ = v17
            }
            var v18 int8 = v1
            _ = v18
        } else {
            var v19 uint8 = 128
            v19 = v19
            _ = v19
        }
    case "foo bar":
        var v20 string = ""
        _ = v20
    }
    var v21 int64 = 690
    v1 = (v1 % (((v1 << ((4294967295) & 63))) | 1))
    v6 = genFn1(true, ((v1 / ((v1) | 1)) == (v1 * (v1 / ((int8(v4)) | 1)))), (v5 << ((4294967295) & 63)))
    v21 = v21
    v21 = (v21 << ((uint(v6)) & 63))
    for fi22 := 0; fi22 < 2; fi22++ {
        fmt.Println(v2, v6, v1)
        var v23 float64 = ((1e10 - 1e10) * 1e10)
        var v24 uint8 = 255
        _ = fi22
        _ = v23
        _ = v24
    }
    if (v2 != v2) {
        for fi25 := 0; fi25 < 2; fi25++ {
            if (-1 == int32(v21)) {
                for fi26 := 0; fi26 < 16; fi26++ {
                    var v27 int64 = 42
                    _ = fi26
                    _ = v27
                }
                v1 = 1
                if true {
                    v28 := []int{}
                    _ = v28
                } else {
                    _ = len(v3)
                }
            }
            var v29 uint8 = 0
            v4 = v4
            _ = fi25
            _ = v29
        }
    } else {
        v30 := v3[-288]
        _ = v30
    }
    for fi31 := 0; fi31 < 8; fi31++ {
        switch (^(fi31)) {
        case 519:
            v2 = "hello"
        case 32767:
            for fi32 := 0; fi32 < 4; fi32++ {
                if (v2 <= "\n\t") {
                    var v33 int32 = 0
                    var v34 float32 = (0.1 / -2.25)
                    _ = v33
                    _ = v34
                }
                fmt.Println(fi31, v2, v4)
                _ = fi32
            }
        case 127:
            v35 := GenS3{F0: (fi31 ^ 32767)}
            _ = v35
        default:
            if true {
                var v36 uint8 = 39
                v37 := []int{-868, (fi31 >> ((42) & 63))}
                _ = v36
                _ = v37
            } else {
                v38 := GenS1{F0: 493, F1: v5}
                switch v2 {
                case "":
                    if true {
                        var v39 uint8 = 2
                        _ = v39
                    } else {
                        if (1e10 != -2.25) {
                            var v40 uint64 = v5
                            v4 = (v4 & 42)
                            _ = v40
                        } else {
                            v21 = v21
                            switch v2 {
                            case "hello":
                                fmt.Println(v5, v1)
                            case "\n\t":
                                v6 += (v6 - (^(v6)))
                            case "foo bar":
                                v4 -= (v4 * 42)
                            }
                            switch (fi31 % ((fi31) | 1)) {
                            case 128:
                                for fi41 := 0; fi41 < 8; fi41++ {
                                    var v42 int8 = v1
                                    _ = fi41
                                    _ = v42
                                }
                            }
                        }
                        switch 128 {
                        case 9223372036854775807:
                            fmt.Println(v6, v2)
                        case -826:
                            v43 := []int{32767, fi31, -665}
                            _ = v43
                        }
                        v44 := v3[(fi31 << ((uint(v5)) & 63))]
                        _ = v44
                    }
                default:
                    for fi45 := 0; fi45 < 15; fi45++ {
                        var v46 float32 = float32((1.5 + 1e10))
                        _ = fi45
                        _ = v46
                    }
                }
                v47 := GenS4{F0: (^(v6)), F1: fi31}
                _ = v38
                _ = v47
            }
        }
        for fi48 := 0; fi48 < 6; fi48++ {
            var v49 int8 = v1
            if (v1 == v49) {
                var v50 string = (v2 + (v2 + "\n\t"))
                v6 = (^(v6))
                switch (v50 + v2) {
                case "hello":
                    var v51 float32 = 1e10
                    _ = v51
                case "a":
                    v4 -= v4
                }
                v21 += (-(v21))
                v50 = v50
                v4 = (v4 * (v4 / (((v4 & (v4 * (v4 * v4)))) | 1)))
                var v52 uint8 = 0
                var v53 uint = 255
                _ = v50
                _ = v52
                _ = v53
            } else {
                if (v6 <= v6) {
                    if true {
                        var v54 uint = 42
                        fmt.Println(v1)
                        if false {
                            fmt.Println(v4, fi48)
                            if ((v6 | genFn1(true, (uint16(v5) <= v6), v5)) > (v6 >> ((617) & 63))) {
                                fmt.Println(fi48)
                                v54 += (v54 + (v54 * 4294967295))
                                v21 += 224
                            }
                            v2 = v2
                        }
                        v55 := []int{(fi48 >> ((uint(v6)) & 63)), (fi48 % (((fi48 ^ fi31)) | 1))}
                        if ((!true) && false) {
                            _ = len(v3)
                        }
                        if (!((^(v6)) == v6)) {
                            var v56 uint16 = 573
                            var v57 int32 = int32(fi48)
                            v5 = (v5 % ((v5) | 1))
                            v58 := GenS3{F0: (fi31 | 669)}
                            _ = v56
                            _ = v57
                            _ = v58
                        } else {
                            fmt.Println(v1)
                            v54 += v54
                            if (((3.14159 - (0.0 * (-2.25 * 1.5))) != 1e10) || false) {
                                var v59 string = v2
                                fmt.Println(v49)
                                _ = len(v55)
                                _ = v59
                            }
                        }
                        v2 = v2
                        if false {
                            switch fi31 {
                            case 905:
                                fmt.Println(v21, v5)
                            case 129:
                                v60 := GenS4{F0: uint16(v5), F1: 994}
                                _ = v60
                            case 127:
                                fmt.Println(v54)
                            }
                            if true {
                                v61 := GenS4{F0: (v6 & genFn1(false, (v1 == v1), (v5 >> ((4294967295) & 63)))), F1: fi31}
                                v62 := GenS2{F0: (^(v54)), F1: -1, F2: 1000}
                                var v63 float32 = float32((3.14159 - float64((1.5 - float32((0.0 + float64(((((3.14159 - 0.1) / 1e10) - 0.1) / -2.25))))))))
                                v64 := GenS4{F0: genFn1(false, true, v5), F1: (fi48 << ((65535) & 63))}
                                _ = v61
                                _ = v62
                                _ = v63
                                _ = v64
                            } else {
                                v54 += (v54 & v54)
                                var v65 uint = 65535
                                v6 -= 65535
                                _ = v65
                            }
                            v66 := GenS4{F0: (v6 - v6), F1: (fi48 >> (((v54 << ((v54) & 63))) & 63))}
                            _ = v66
                        }
                        if false {
                            fmt.Println(fi31, fi31, fi48)
                            var v67 int16 = 32767
                            _ = v67
                        }
                        _ = v54
                        _ = v55
                    } else {
                        v21 = -782
                    }
                }
                v68 := GenS4{F0: (v6 - (v6 & (v6 / ((120) | 1)))), F1: fi31}
                _ = v68
            }
            v69 := v3[(-(fi31))]
            _ = fi48
            _ = v49
            _ = v69
        }
        _ = fi31
    }
    fmt.Println(v21, v1, v21)
    v70 := []int{}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v70:", len(v70), v70)
    _ = v1
    _ = v15
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v70
}
