package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint16; F1 uint16; F2 uint32 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint16; F1 bool; F2 int8 }
type GenS5 struct { F0 uint16; F1 uint32; F2 uint32 }
type GenS6 struct { F0 uint; F1 int8; F2 uint64 }
type GenS7 struct { F0 uint16; F1 uint64; F2 uint8 }
type GenS8 struct { F0 int32; F1 bool }
type GenS9 struct { F0 uint32 }
type GenS10 struct { F0 uint64; F1 int64 }
type GenS11 struct { F0 int8; F1 bool }
type GenS12 struct { F0 uint; F1 int64 }
type GenS13 struct { F0 string }
type GenS14 struct { F0 uint64 }
func genFn0(p0 int16) int16 {
    if false {
        for fi4 := 0; fi4 < 11; fi4++ {
            var v5 bool = (79 > 493)
            _ = fi4
            _ = v5
        }
        v6 := GenS1{F0: 42}
        _ = v6
    } else {
        if false {
            if true {
                v7 := []int{}
                _ = v7
            }
        } else {
            p0 = 32767
            switch 42 {
            case 457:
                v8 := []int{128}
                _ = v8
            }
        }
        v9 := []int{0}
        _ = v9
    }
    for fi10 := 0; fi10 < 4; fi10++ {
        if false {
            var v11 uint64 = 113
            v12 := GenS1{F0: (v11 | v11)}
            _ = v11
            _ = v12
        }
        var v13 uint16 = 255
        _ = fi10
        _ = v13
    }
    if false {
        var v14 string = (("foo bar" + ("hello" + "foo bar")) + ("hello" + "a"))
        for fi15 := 0; fi15 < 7; fi15++ {
            switch v14 {
            case "a":
                for fi16 := 0; fi16 < 10; fi16++ {
                    v17 := []int{(fi16 ^ 0), (^(fi15))}
                    switch (fi16 << ((735) & 63)) {
                    case -1:
                        if true {
                            v18 := GenS2{F0: 1, F1: 42, F2: 986}
                            v14 = v14
                            switch fi16 {
                            case 127:
                                var v19 uint8 = 2
                                if false {
                                    fmt.Println(v19, fi15)
                                }
                                p0 = (p0 * (p0 - -1))
                                _ = v19
                            }
                            _ = v18
                        }
                    case 4294967295:
                        var v20 bool = false
                        _ = v20
                    }
                    v14 = v14
                    _ = fi16
                    _ = v17
                }
            case "foo bar":
                v21 := GenS3{F0: 255}
                _ = v21
            case "hello":
                v14 = (v14 + ("a" + v14))
            }
            _ = fi15
        }
        _ = v14
    }
    if false {
        switch 9223372036854775807 {
        case 32768:
            switch 4294967295 {
            case 453:
                v22 := GenS4{F0: 241, F1: true, F2: 127}
                _ = v22
            default:
                for fi23 := 0; fi23 < 2; fi23++ {
                    v24 := GenS5{F0: 196, F1: 0, F2: 0}
                    switch 32768 {
                    case 42:
                        v25 := GenS6{F0: 4294967295, F1: -128, F2: 0}
                        _ = v25
                    case 1:
                        v26 := []int{(fi23 >> ((2) & 63))}
                        _ = v26
                    case 32767:
                        for fi27 := 0; fi27 < 11; fi27++ {
                            if false {
                                v28 := []int{(fi27 * fi27), fi23}
                                if true {
                                    if len(v28) >= 2 {
                                        v28 = v28[1:len(v28)]
                                    }
                                    p0 -= p0
                                    switch (fi27 + (fi27 << ((240) & 63))) {
                                    case 128:
                                        fmt.Println(p0, fi23)
                                    case 127:
                                        v29 := GenS7{F0: 558, F1: 2, F2: 1}
                                        _ = v29
                                    case 1:
                                        if (76 < 2) {
                                            p0 += (p0 & (p0 * p0))
                                            if (-9223372036854775808 != -1) {
                                                p0 = (^(p0))
                                            }
                                            fmt.Println(fi23, fi23)
                                        } else {
                                            var v30 int16 = p0
                                            if (false && true) {
                                                v31 := GenS8{F0: -214, F1: false}
                                                if (!false) {
                                                    var v32 int32 = -592
                                                    _ = v32
                                                }
                                                _ = v31
                                            }
                                            switch "" {
                                            case "hello":
                                                var v33 int = (fi23 % ((fi23) | 1))
                                                _ = v33
                                            case "a":
                                                fmt.Println(p0)
                                            }
                                            if false {
                                                p0 += -32768
                                            }
                                            var v34 bool = true
                                            _ = v30
                                            _ = v34
                                        }
                                    default:
                                        var v35 uint64 = 2
                                        _ = v35
                                    }
                                }
                                v36 := GenS5{F0: 255, F1: 65535, F2: 2}
                                _ = v28
                                _ = v36
                            } else {
                                var v37 uint8 = 128
                                _ = v37
                            }
                            if false {
                                var v38 uint32 = 559
                                fmt.Println(fi27, fi23, fi23)
                                if false {
                                    fmt.Println(fi27, p0, p0)
                                    fmt.Println(p0, v38, fi27)
                                    var v39 bool = true
                                    if ((p0 << ((4294967295) & 63)) != int16(v38)) {
                                        v40 := []int{(fi23 / ((fi23) | 1)), (fi23 % ((fi23) | 1))}
                                        v41 := GenS9{F0: (^(v38))}
                                        _ = v40
                                        _ = v41
                                    } else {
                                        v42 := GenS8{F0: -2147483648, F1: v39}
                                        _ = v42
                                    }
                                    _ = v39
                                } else {
                                    v38 -= v38
                                    v43 := []int{fi27, fi27, (fi23 << ((2) & 63))}
                                    if (fi27 < -321) {
                                        var v44 uint8 = 128
                                        if len(v43) >= 2 {
                                            v43 = v43[1:len(v43)]
                                        }
                                        _ = v44
                                    } else {
                                        var v45 string = ""
                                        var v46 string = "hello"
                                        v47 := GenS5{F0: 124, F1: (v38 | (v38 | 653)), F2: (^(v38))}
                                        _ = v45
                                        _ = v46
                                        _ = v47
                                    }
                                    _ = v43
                                }
                                _ = v38
                            }
                            p0 -= (p0 | -979)
                            v48 := GenS10{F0: 18446744073709551615, F1: 652}
                            switch 32768 {
                            case 2147483647:
                                var v49 int8 = -128
                                _ = v49
                            case -1:
                                var v50 float32 = 1e10
                                _ = v50
                            default:
                                var v51 uint8 = 0
                                _ = v51
                            }
                            _ = fi27
                            _ = v48
                        }
                    default:
                        v52 := GenS8{F0: -557, F1: ((2147483648 != 2147483648) || false)}
                        _ = v52
                    }
                    _ = fi23
                    _ = v24
                }
            }
        case 2147483647:
            p0 = p0
        case -539:
            for fi53 := 0; fi53 < 18; fi53++ {
                switch "foo bar" {
                case "":
                    for fi54 := 0; fi54 < 15; fi54++ {
                        p0 = (p0 + 1000)
                        var v55 uint16 = 255
                        _ = fi54
                        _ = v55
                    }
                case "a":
                    var v56 int64 = -135
                    _ = v56
                default:
                    for fi57 := 0; fi57 < 4; fi57++ {
                        p0 -= 963
                        _ = fi57
                    }
                }
                var v58 uint32 = 441
                _ = fi53
                _ = v58
            }
        }
        var v59 float64 = -2.25
        _ = v59
    } else {
        fmt.Println(p0)
    }
    v60 := GenS11{F0: 1, F1: (true && false)}
    _ = p0
    _ = v60
    return (p0 | (p0 * (p0 / (((p0 * p0)) | 1))))
}

func main() {
    var v1 uint16 = 65535
    var v2 uint16 = v1
    v3 := []int{9223372036854775807}
    var v61 float64 = 0.1
    v62 := GenS6{F0: 42, F1: 7, F2: 18446744073709551615}
    var v63 int32 = 0
    v61 = (v61 + (v61 / v61))
    v1 = (v2 + 1)
    v64 := GenS12{F0: 301, F1: 0}
    v3 = append(v3, 758)
    _ = len(v3)
    v65 := GenS4{F0: (^(v1)), F1: false, F2: -1}
    var v66 uint = 307
    var v67 uint = (v66 + v66)
    v68 := GenS13{F0: ("" + "a")}
    v69 := GenS14{F0: 18446744073709551615}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v61:", v61)
    fmt.Println("v63:", v63)
    fmt.Println("v66:", v66)
    fmt.Println("v67:", v67)
    _ = v1
    _ = v2
    _ = v3
    _ = v61
    _ = v62
    _ = v63
    _ = v64
    _ = v65
    _ = v66
    _ = v67
    _ = v68
    _ = v69
}
