package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 bool; F1 uint16 }
type GenS4 struct { F0 uint8; F1 int64; F2 int16 }
type GenS5 struct { F0 int8; F1 int8; F2 int }
type GenS6 struct { F0 int }
type GenS7 struct { F0 bool; F1 uint16; F2 int32 }
type GenS8 struct { F0 string }
func genFn0() uint64 {
    v6 := []int{}
    if true {
        var v7 uint32 = 867
        v7 = (v7 * v7)
        v7 = (^(v7))
        _ = v7
    } else {
        v8 := GenS1{F0: 255}
        _ = v8
    }
    var v9 uint16 = 1
    if false {
        var v10 uint16 = 2
        v9 += (^(v9))
        _ = v10
    } else {
        v9 = (v9 << ((18446744073709551615) & 63))
        for fi11 := 0; fi11 < 8; fi11++ {
            var v12 string = "a"
            v9 -= (v9 << ((42) & 63))
            _ = fi11
            _ = v12
        }
        for fi13 := 0; fi13 < 15; fi13++ {
            v9 -= v9
            _ = fi13
        }
    }
    _ = v6
    _ = v9
    return 0
}
func genFn1() int16 {
    if true {
        for fi14 := 0; fi14 < 5; fi14++ {
            if true {
                for fi15 := 0; fi15 < 10; fi15++ {
                    v16 := GenS1{F0: 255}
                    if true {
                        v17 := []int{fi15, (fi14 / ((fi14) | 1)), (fi14 * -74)}
                        _ = v17
                    } else {
                        v18 := []int{}
                        switch "hello" {
                        case "hello":
                            v19 := GenS1{F0: 463}
                            _ = v19
                        default:
                            v20 := GenS2{F0: -92}
                            _ = v20
                        }
                        if len(v18) >= 2 {
                            v18 = v18[1:len(v18)]
                        }
                        _ = v18
                    }
                    v21 := []int{(fi15 ^ (fi14 & (-(fi14)))), (-(fi15)), fi14}
                    if (-1 > -1) {
                        var v22 uint = 445
                        var v23 uint16 = 42
                        _ = v22
                        _ = v23
                    } else {
                        var v24 int16 = 42
                        _ = v24
                    }
                    if false {
                        var v25 uint8 = 128
                        var v26 uint16 = 1
                        var v27 bool = false
                        _ = v25
                        _ = v26
                        _ = v27
                    } else {
                        var v28 int64 = -35
                        v29 := GenS2{F0: -92}
                        _ = v28
                        _ = v29
                    }
                    _ = fi15
                    _ = v16
                    _ = v21
                }
            }
            if false {
                switch fi14 {
                case 0:
                    v30 := []int{fi14}
                    _ = v30
                }
                for fi31 := 0; fi31 < 7; fi31++ {
                    var v32 float32 = 1.5
                    var v33 string = "hello"
                    v34 := []int{0, 899, fi31, (fi14 * (fi14 % (((fi31 * (fi14 + -1))) | 1)))}
                    _ = fi31
                    _ = v32
                    _ = v33
                    _ = v34
                }
                switch "foo bar" {
                case "a":
                    v35 := []int{fi14}
                    _ = v35
                case "hello":
                    for fi36 := 0; fi36 < 7; fi36++ {
                        var v37 int64 = 0
                        v38 := []int{fi36, fi14, (fi14 / (((^(fi14))) | 1))}
                        _ = fi36
                        _ = v37
                        _ = v38
                    }
                }
            }
            _ = fi14
        }
    } else {
        for fi39 := 0; fi39 < 4; fi39++ {
            v40 := GenS3{F0: true, F1: 2}
            _ = fi39
            _ = v40
        }
    }
    var v41 int8 = 1
    v41 = 0
    _ = v41
    return -10
}

func main() {
    var v1 uint32 = 65535
    var v2 int8 = 39
    var v3 uint16 = 555
    var v4 int64 = -940
    var v5 float64 = (float64(1.5) - 0.0)
    v42 := []int{-1, -892, 127}
    for fi43 := 0; fi43 < 17; fi43++ {
        if false {
            if true {
                fmt.Println(v1)
                for fi44 := 0; fi44 < 14; fi44++ {
                    switch "hello" {
                    case "":
                        switch "hello" {
                        case "a":
                            if (1000 != 239) {
                                switch ("\n\t" + "a") {
                                case "\n\t":
                                    if ((v3 | 15) == (v3 % ((v3) | 1))) {
                                        v4 = v4
                                        switch "foo bar" {
                                        case "\n\t":
                                            var v45 uint64 = 18446744073709551615
                                            _ = v45
                                        case "":
                                            var v46 int16 = -482
                                            _ = v46
                                        }
                                        if false {
                                            var v47 string = "foo bar"
                                            _ = len(v42)
                                            if ((!false) || true) {
                                                v48 := GenS3{F0: ((v2 <= v2) && false), F1: (^(v3))}
                                                _ = len(v42)
                                                v1 = (^(v1))
                                                _ = v48
                                            }
                                            _ = v47
                                        } else {
                                            if len(v42) > 0 {
                                                v42[0] = 42
                                            }
                                            v5 += v5
                                            _ = len(v42)
                                            var v49 uint16 = v3
                                            v50 := GenS4{F0: 127, F1: 0, F2: 32767}
                                            fmt.Println(v5, v49)
                                            _ = v49
                                            _ = v50
                                        }
                                    }
                                case "foo bar":
                                    if false {
                                        var v51 uint16 = v3
                                        v51 = (v51 >> ((297) & 63))
                                        switch (("" + "foo bar") + "") {
                                        case "a":
                                            v51 = 1
                                        case "foo bar":
                                            var v52 int = fi44
                                            _ = v52
                                        }
                                        _ = v51
                                    } else {
                                        v53 := GenS5{F0: (-(v2)), F1: 42, F2: fi44}
                                        if true {
                                            v54 := GenS6{F0: -1}
                                            var v55 int = (fi43 % ((608) | 1))
                                            var v56 uint32 = v1
                                            v57 := GenS5{F0: 100, F1: v2, F2: fi44}
                                            _ = v54
                                            _ = v55
                                            _ = v56
                                            _ = v57
                                        } else {
                                            var v58 float64 = v5
                                            v59 := GenS5{F0: (v2 % (((v2 * v2)) | 1)), F1: v2, F2: (fi44 % ((fi43) | 1))}
                                            _ = v58
                                            _ = v59
                                        }
                                        var v60 float64 = (v5 - -2.25)
                                        _ = v53
                                        _ = v60
                                    }
                                case "hello":
                                    var v61 bool = false
                                    _ = v61
                                }
                                _ = len(v42)
                                var v62 int32 = 0
                                fmt.Println(fi43)
                                v62 = (^(v62))
                                _ = v62
                            }
                        default:
                            if (32767 >= 575) {
                                v63 := GenS7{F0: false, F1: 948, F2: 1}
                                _ = v63
                            } else {
                                v64 := GenS5{F0: (v2 | v2), F1: (^(v2)), F2: fi43}
                                _ = v64
                            }
                        }
                    }
                    v3 = v3
                    _ = fi44
                }
                v1 += (v1 >> ((255) & 63))
            }
        }
        if len(v42) > 0 {
            v42[0] = fi43
        }
        v65 := GenS3{F0: false, F1: v3}
        _ = fi43
        _ = v65
    }
    if (false || (1000 < genFn1())) {
        var v66 int8 = (v2 + (v2 + -29))
        v66 = v66
        var v67 int8 = v2
        _ = v66
        _ = v67
    } else {
        _ = len(v42)
        switch 128 {
        case 4294967295:
            var v68 bool = (!false)
            _ = v68
        default:
            var v69 uint = 898
            _ = v69
        }
    }
    var v70 uint8 = 0
    v5 -= (float64(float32(v5)) + (((((float64(-2.25) / -2.25) * (v5 - v5)) / v5) / v5) * float64((0.0 + (float32(0.0) + -2.25)))))
    if len(v42) >= 2 {
        v42 = v42[1:len(v42)]
    }
    v71 := GenS6{F0: 283}
    for fi72 := 0; fi72 < 15; fi72++ {
        if len(v42) > 0 {
            v42[0] = -1
        }
        var v73 uint8 = v70
        _ = fi72
        _ = v73
    }
    var v74 uint64 = 42
    _ = len(v42)
    v75 := GenS8{F0: "\n\t"}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v42:", len(v42), v42)
    fmt.Println("v5:", v5)
    fmt.Println("v70:", v70)
    fmt.Println("v74:", v74)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v5
    _ = v70
    _ = v71
    _ = v74
    _ = v75
}
