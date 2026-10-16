package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int32; F1 int8; F2 int64 }
type GenS4 struct { F0 int64 }
func genFn0(p0 int64, p1 int32, p2 uint16) string {
    v7 := GenS1{F0: (p1 / ((2147483647) | 1))}
    v8 := GenS1{F0: (^(p1))}
    var v9 int32 = (p1 % ((0) | 1))
    for fi10 := 0; fi10 < 4; fi10++ {
        p1 += (p1 + 1)
        _ = fi10
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return "\n\t"
}
func genFn1(p0 int8, p1 int32) int16 {
    if (289 > 42) {
        v11 := []int{}
        fmt.Println(p0, p1)
        v12 := GenS1{F0: (p1 / ((p1) | 1))}
        _ = v11
        _ = v12
    } else {
        for fi13 := 0; fi13 < 1; fi13++ {
            v14 := []int{fi13}
            _ = fi13
            _ = v14
        }
        v15 := []int{32768, 61, 127}
        _ = v15
    }
    var v16 uint16 = 32767
    p1 = 0
    switch ("hello" + "\n\t") {
    case "hello":
        switch "" {
        case "a":
            for fi17 := 0; fi17 < 5; fi17++ {
                p0 = p0
                p0 += p0
                for fi18 := 0; fi18 < 10; fi18++ {
                    var v19 uint32 = 65535
                    v20 := []int{(fi17 ^ fi17), int(p0), fi18, (-(fi18))}
                    v21 := GenS1{F0: (p1 - -2147483648)}
                    _ = fi18
                    _ = v19
                    _ = v20
                    _ = v21
                }
                _ = fi17
            }
        case "foo bar":
            switch 32767 {
            case 128:
                v22 := GenS1{F0: p1}
                _ = v22
            case 2147483647:
                var v23 int64 = 69
                _ = v23
            case 1:
                switch "a" {
                case "hello":
                    p0 -= -75
                case "foo bar":
                    v16 = v16
                case "\n\t":
                    p0 -= p0
                }
            default:
                if (!(!false)) {
                    p1 -= p1
                    var v24 uint32 = 701
                    _ = v24
                } else {
                    if false {
                        for fi25 := 0; fi25 < 2; fi25++ {
                            v26 := GenS1{F0: 1}
                            _ = fi25
                            _ = v26
                        }
                    }
                    var v27 int = -2147483648
                    v28 := []int{(v27 / (((v27 << ((255) & 63))) | 1)), v27, (v27 ^ v27)}
                    _ = v27
                    _ = v28
                }
            }
        case "":
            v29 := []int{}
            _ = v29
        default:
            if true {
                v16 = 32767
                v30 := GenS2{F0: 127}
                var v31 int32 = 1
                _ = v30
                _ = v31
            } else {
                if true {
                    var v32 bool = (0 > -1)
                    _ = v32
                }
                if false {
                    p1 = 1
                    for fi33 := 0; fi33 < 14; fi33++ {
                        v34 := GenS2{F0: 73}
                        for fi35 := 0; fi35 < 6; fi35++ {
                            if ((false && (42 <= 649)) && true) {
                                fmt.Println(p0, fi35, fi33)
                            } else {
                                if (3.14159 <= float64(1e10)) {
                                    var v36 string = "hello"
                                    v37 := []int{(fi33 / ((295) | 1))}
                                    v16 = (v16 * 65535)
                                    _ = v36
                                    _ = v37
                                } else {
                                    var v38 int16 = 517
                                    fmt.Println(fi33, v38)
                                    _ = v38
                                }
                                var v39 uint16 = 42
                                _ = v39
                            }
                            _ = fi35
                        }
                        _ = fi33
                        _ = v34
                    }
                    fmt.Println(p0)
                }
            }
        }
    case "\n\t":
        switch -215 {
        case -665:
            var v40 float32 = 0.0
            _ = v40
        case 32767:
            v41 := GenS3{F0: p1, F1: p0, F2: 42}
            _ = v41
        }
    case "foo bar":
        v42 := []int{32768, 763, -1}
        _ = v42
    default:
        fmt.Println(v16)
    }
    for fi43 := 0; fi43 < 9; fi43++ {
        p1 = (^(p1))
        v44 := []int{-350, -9223372036854775808}
        _ = fi43
        _ = v44
    }
    _ = p0
    _ = p1
    _ = v16
    return int16(p1)
}

func main() {
    var v1 float32 = 0.0
    var v2 uint = 255
    var v3 int32 = 36
    var v4 float32 = v1
    v5 := make(map[uint32]int)
    var v6 int32 = v3
    fmt.Println(v6, v3, v2)
    v6 += (v6 / ((2147483647) | 1))
    v45 := GenS4{F0: -850}
    v46 := GenS1{F0: 257}
    if true {
        for fi47 := 0; fi47 < 8; fi47++ {
            v48 := GenS2{F0: 2}
            v3 += (v6 << (((^(v2))) & 63))
            var v49 float32 = 3.14159
            _ = fi47
            _ = v48
            _ = v49
        }
        for fi50 := 0; fi50 < 14; fi50++ {
            for fi51 := 0; fi51 < 16; fi51++ {
                v52 := []int{fi50, fi51}
                if (((1.5 * ((float64(0.1) + 0.0) * (1.5 + 0.1))) * 0.1) >= -2.25) {
                    switch 32767 {
                    case -9223372036854775808:
                        v53 := GenS4{F0: -1}
                        _ = v53
                    case -928:
                        v3 -= v6
                    case -2147483648:
                        v6 = (v3 | (v3 % ((2147483647) | 1)))
                    default:
                        var v54 int64 = -1
                        _ = v54
                    }
                } else {
                    _ = len(v52)
                    v55 := v5[2147483648]
                    if (!false) {
                        _ = len(v52)
                        switch "a" {
                        case "foo bar":
                            _ = len(v5)
                        }
                        fmt.Println(v3, v4)
                    } else {
                        if true {
                            v56 := v5[1]
                            v57 := v5[146]
                            switch (fi51 & 691) {
                            case 42:
                                var v58 uint32 = 800
                                _ = v58
                            case -359:
                                if len(v52) >= 2 {
                                    v52 = v52[1:len(v52)]
                                }
                            case -703:
                                var v59 float32 = v1
                                _ = v59
                            default:
                                v55 += v57
                            }
                            v60 := GenS2{F0: 30}
                            _ = v56
                            _ = v57
                            _ = v60
                        } else {
                            if false {
                                v55 += fi50
                                switch ("hello" + "foo bar") {
                                case "":
                                    fmt.Println(v2, v3)
                                case "\n\t":
                                    if true {
                                        var v61 uint8 = 128
                                        _ = v61
                                    }
                                case "hello":
                                    v4 += ((3.14159 - 1.5) * v1)
                                default:
                                    fmt.Println(v4)
                                }
                                switch (fi51 % (((v55 ^ fi51)) | 1)) {
                                case -45:
                                    var v62 int32 = -2147483648
                                    _ = v62
                                case 2147483647:
                                    v1 = float32(-2.25)
                                }
                            } else {
                                v2 -= v2
                                var v63 uint64 = 377
                                _ = v63
                            }
                        }
                        v4 += -2.25
                        v64 := GenS4{F0: 0}
                        _ = v64
                    }
                    _ = v55
                }
                _ = fi51
                _ = v52
            }
            v1 = v4
            if ((("foo bar" + "") + "foo bar") > "\n\t") {
                for fi65 := 0; fi65 < 18; fi65++ {
                    v66 := []int{int(v3), (^(fi50))}
                    var v67 float32 = float32(1e10)
                    _ = fi65
                    _ = v66
                    _ = v67
                }
            }
            _ = fi50
        }
        for fi68 := 0; fi68 < 8; fi68++ {
            fmt.Println(v3, v3)
            v1 = 0.0
            switch (("" + (("foo bar" + "a") + "foo bar")) + "\n\t") {
            case "foo bar":
                var v69 int64 = 9223372036854775807
                _ = v69
            case "\n\t":
                switch "foo bar" {
                case "":
                    _ = len(v5)
                case "foo bar":
                    if true {
                        fmt.Println(v3)
                    } else {
                        v1 = 0.0
                    }
                case "a":
                    var v70 uint64 = 42
                    _ = v70
                }
            case "":
                for fi71 := 0; fi71 < 13; fi71++ {
                    v72 := []int{}
                    v73 := GenS2{F0: uint8(v2)}
                    _ = fi71
                    _ = v72
                    _ = v73
                }
            default:
                v74 := GenS3{F0: v6, F1: 127, F2: -9223372036854775808}
                _ = v74
            }
            _ = fi68
        }
    }
    for fi75 := 0; fi75 < 17; fi75++ {
        var v76 float64 = -2.25
        _ = fi75
        _ = v76
    }
    v3 = -996
    fmt.Println(v6, v6)
    var v77 uint32 = 2147483647
    var v78 bool = true
    fmt.Println(v3)
    var v79 float64 = (1e10 * 0.0)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v77:", v77)
    fmt.Println("v78:", v78)
    fmt.Println("v79:", v79)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v5
    _ = v6
    _ = v77
    _ = v78
    _ = v79
}
