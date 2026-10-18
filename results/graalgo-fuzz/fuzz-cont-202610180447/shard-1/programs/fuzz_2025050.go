package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int64; F1 int8 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int32; F1 bool; F2 uint16 }
type GenS6 struct { F0 int8; F1 uint64 }
type GenS7 struct { F0 int; F1 uint16; F2 int16 }
type GenS8 struct { F0 uint16; F1 int8 }
func genFn0() int32 {
    var v6 uint16 = 65535
    v6 += (v6 * 501)
    v7 := GenS1{F0: 2147483647}
    if false {
        for fi8 := 0; fi8 < 8; fi8++ {
            if true {
                for fi9 := 0; fi9 < 2; fi9++ {
                    v10 := GenS1{F0: 895}
                    _ = fi9
                    _ = v10
                }
            }
            _ = fi8
        }
        switch "foo bar" {
        case "foo bar":
            for fi11 := 0; fi11 < 15; fi11++ {
                v12 := GenS1{F0: 42}
                _ = fi11
                _ = v12
            }
        }
        switch 32768 {
        case 4294967295:
            if ((((true || true) || true) || ((127 == 9223372036854775807) && false)) || true) {
                fmt.Println(v6)
                v6 -= 346
            }
        case 2147483647:
            switch (("foo bar" + "foo bar") + "\n\t") {
            case "":
                var v13 uint8 = 127
                _ = v13
            case "a":
                v14 := []int{42, 635}
                _ = v14
            case "\n\t":
                var v15 int64 = 568
                _ = v15
            }
        }
    } else {
        for fi16 := 0; fi16 < 7; fi16++ {
            v6 -= 1
            v17 := []int{-1, (fi16 / (((fi16 / ((127) | 1))) | 1)), (fi16 * 9223372036854775807), -1}
            v18 := GenS2{F0: (false || (1 != 42))}
            _ = fi16
            _ = v17
            _ = v18
        }
        var v19 uint = 843
        var v20 int32 = -1
        _ = v19
        _ = v20
    }
    _ = v6
    _ = v7
    return 0
}

func main() {
    var v1 float64 = (float64(1e10) - 1.5)
    var v2 int = 0
    var v3 int64 = -630
    v4 := []int{-970, v2, 132}
    v5 := make(map[string]uint32)
    for fi21 := 0; fi21 < 11; fi21++ {
        _ = len(v5)
        var v22 uint32 = 447
        v2 = fi21
        _ = fi21
        _ = v22
    }
    for fi23 := 0; fi23 < 9; fi23++ {
        v3 = v3
        _ = len(v4)
        v3 -= v3
        _ = fi23
    }
    v2 = 4294967295
    for fi24 := 0; fi24 < 18; fi24++ {
        _ = len(v5)
        v3 = v3
        var v25 uint32 = 42
        _ = fi24
        _ = v25
    }
    for fi26 := 0; fi26 < 20; fi26++ {
        for fi27 := 0; fi27 < 1; fi27++ {
            fmt.Println(v3)
            var v28 bool = false
            switch (^(v2)) {
            case 128:
                var v29 float64 = v1
                _ = v29
            case -2147483648:
                v30 := v5["\n\t"]
                v28 = (-1 < 77)
                _ = v30
            case 2147483647:
                var v31 int32 = 1
                _ = v31
            }
            _ = fi27
            _ = v28
        }
        for fi32 := 0; fi32 < 20; fi32++ {
            v3 = -591
            switch fi32 {
            case 997:
                v2 -= fi32
            case -326:
                if ((v3 * v3) == v3) {
                    if false {
                        var v33 uint8 = 2
                        v4 = append(v4, fi32)
                        fmt.Println(v33)
                        _ = v33
                    }
                    if false {
                        if true {
                            var v34 float32 = 0.1
                            _ = v34
                        }
                        var v35 int64 = 1
                        _ = v35
                    } else {
                        var v36 uint32 = 612
                        v37 := GenS2{F0: false}
                        _ = v36
                        _ = v37
                    }
                    if false {
                        v38 := GenS2{F0: (((v3 & v3) <= (-(v3))) && (606 < 998))}
                        switch (^(fi26)) {
                        case -1:
                            var v39 uint32 = 992
                            _ = v39
                        case 9223372036854775807:
                            v1 = 3.14159
                        case 32768:
                            var v40 int64 = v3
                            _ = v40
                        default:
                            v41 := GenS2{F0: (1000 <= -798)}
                            _ = v41
                        }
                        _ = v38
                    } else {
                        _ = len(v4)
                    }
                    var v42 bool = false
                    _ = v42
                }
            case -216:
                var v43 int16 = -932
                _ = v43
            }
            if true {
                var v44 int32 = -272
                var v45 uint32 = 2
                _ = v44
                _ = v45
            } else {
                var v46 float64 = float64((0.0 - float32(0.0)))
                v47 := GenS1{F0: -1}
                _ = v46
                _ = v47
            }
            _ = fi32
        }
        v48 := GenS2{F0: true}
        _ = fi26
        _ = v48
    }
    v49 := GenS3{F0: (v3 & v3), F1: 100}
    fmt.Println(v3)
    var v50 uint8 = 42
    v51 := GenS3{F0: v3, F1: 0}
    if len(v4) > 0 {
        v4[0] = int(v3)
    }
    v52 := v5["hello"]
    v53 := GenS3{F0: v3, F1: 100}
    fmt.Println(v1, v1)
    for fi54 := 0; fi54 < 16; fi54++ {
        _ = len(v5)
        for fi55 := 0; fi55 < 13; fi55++ {
            if false {
                if false {
                    v56 := GenS4{F0: (("" + "") + ("foo bar" + ((("\n\t" + "") + "\n\t") + ("a" + "foo bar"))))}
                    v57 := GenS3{F0: (v3 % (((v3 / ((-143) | 1))) | 1)), F1: 42}
                    if (!(!false)) {
                        if true {
                            v58 := GenS5{F0: 94, F1: false, F2: 382}
                            _ = v58
                        } else {
                            switch v2 {
                            case -2147483648:
                                v2 = (fi55 % ((fi54) | 1))
                            case -9223372036854775808:
                                switch "\n\t" {
                                case "\n\t":
                                    v52 += 65535
                                default:
                                    _ = len(v4)
                                }
                            }
                        }
                        if len(v4) > 0 {
                            v4[0] = (-(v2))
                        }
                        switch (fi55 + v2) {
                        case -348:
                            var v59 bool = false
                            _ = v59
                        }
                    } else {
                        if (350 == -1) {
                            v1 += v1
                            if true {
                                v60 := GenS6{F0: 100, F1: 540}
                                _ = v60
                            } else {
                                v61 := GenS7{F0: (fi54 % ((9223372036854775807) | 1)), F1: 42, F2: int16(v52)}
                                v5[("foo bar" + (("hello" + "foo bar") + "hello"))] = (v52 * v52)
                                if (!true) {
                                    var v62 bool = (!true)
                                    v52 = (v52 % (((^(v52))) | 1))
                                    switch "a" {
                                    case "foo bar":
                                        var v63 uint8 = (v50 ^ 165)
                                        _ = v63
                                    }
                                    fmt.Println(v1, v2, v52)
                                    var v64 uint32 = 42
                                    var v65 string = (("foo bar" + "hello") + "\n\t")
                                    v64 -= (^(v52))
                                    _ = v62
                                    _ = v64
                                    _ = v65
                                }
                                var v66 bool = true
                                var v67 string = "a"
                                _ = v61
                                _ = v66
                                _ = v67
                            }
                        }
                        v68 := GenS7{F0: 2147483647, F1: 882, F2: -72}
                        _ = v68
                    }
                    _ = v56
                    _ = v57
                }
                v50 = v50
                v1 += (v1 * ((v1 / 1.0) * v1))
                v5["\n\t"] = v52
            } else {
                var v69 int32 = 42
                var v70 string = ("\n\t" + "")
                _ = v69
                _ = v70
            }
            var v71 uint = 576
            _ = fi55
            _ = v71
        }
        v72 := GenS8{F0: 255, F1: 42}
        _ = fi54
        _ = v72
    }
    var v73 int = (v2 - (v2 & v2))
    var v74 int32 = -2147483648
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v50:", v50)
    fmt.Println("v52:", v52)
    fmt.Println("v73:", v73)
    fmt.Println("v74:", v74)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v50
    _ = v51
    _ = v52
    _ = v53
    _ = v73
    _ = v74
}
