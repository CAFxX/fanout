package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint16; F1 uint32; F2 bool }
type GenS3 struct { F0 uint8; F1 bool }
type GenS4 struct { F0 uint64; F1 uint32; F2 int64 }
type GenS5 struct { F0 int32; F1 int16 }
type GenS6 struct { F0 int64; F1 int32 }
type GenS7 struct { F0 bool }
func genFn0() uint64 {
    var v5 uint = 4294967295
    for fi6 := 0; fi6 < 5; fi6++ {
        switch "a" {
        case "a":
            v7 := []int{fi6, fi6}
            _ = v7
        }
        var v8 int8 = -78
        var v9 uint16 = 1
        _ = fi6
        _ = v8
        _ = v9
    }
    _ = v5
    return 0
}
func genFn1(p0 int8, p1 int16) uint16 {
    for fi10 := 0; fi10 < 20; fi10++ {
        switch fi10 {
        case 32768:
            v11 := GenS1{F0: p1}
            _ = v11
        }
        p0 = p0
        var v12 int32 = -2147483648
        _ = fi10
        _ = v12
    }
    fmt.Println(p0, p0)
    v13 := []int{42, 32767, 32768}
    switch 835 {
    case 4294967295:
        for fi14 := 0; fi14 < 4; fi14++ {
            for fi15 := 0; fi15 < 18; fi15++ {
                var v16 uint8 = 175
                p0 = (^(p0))
                var v17 int8 = -1
                _ = fi15
                _ = v16
                _ = v17
            }
            var v18 uint = 0
            _ = fi14
            _ = v18
        }
    default:
        fmt.Println(p0)
    }
    for fi19 := 0; fi19 < 16; fi19++ {
        var v20 uint64 = 4294967295
        for fi21 := 0; fi21 < 9; fi21++ {
            var v22 uint32 = 1
            switch "" {
            case "a":
                v23 := GenS2{F0: 65535, F1: (v22 | v22), F2: true}
                _ = v23
            }
            switch fi21 {
            case 4294967295:
                p0 -= (p0 - (p0 - p0))
            case 42:
                v24 := GenS3{F0: 159, F1: (255 > 128)}
                if true {
                    fmt.Println(v22, p1)
                    var v25 uint32 = v22
                    if (0 == 2) {
                        v26 := GenS2{F0: 376, F1: 641, F2: false}
                        switch "a" {
                        case "hello":
                            if true {
                                if true {
                                    fmt.Println(v25, v22)
                                }
                                if true {
                                    var v27 float32 = (1.5 / (1e10 - (1e10 - 1.5)))
                                    if true {
                                        fmt.Println(p1)
                                        fmt.Println(v22)
                                        fmt.Println(fi19, p1, v25)
                                        var v28 bool = true
                                        _ = v28
                                    }
                                    _ = v27
                                } else {
                                    v29 := GenS4{F0: (v20 / (((v20 * v20)) | 1)), F1: v25, F2: 1}
                                    var v30 uint64 = (^(v20))
                                    _ = v29
                                    _ = v30
                                }
                                if (!false) {
                                    _ = len(v13)
                                    var v31 int16 = int16(v22)
                                    var v32 float32 = ((0.1 / -2.25) + 1.5)
                                    _ = v31
                                    _ = v32
                                } else {
                                    if len(v13) >= 2 {
                                        v13 = v13[1:len(v13)]
                                    }
                                    if len(v13) >= 2 {
                                        v13 = v13[1:len(v13)]
                                    }
                                }
                            }
                        }
                        if (463 <= v20) {
                            v25 -= v22
                            if false {
                                var v33 uint16 = 255
                                _ = v33
                            } else {
                                switch (fi21 - fi21) {
                                case 32767:
                                    fmt.Println(v20, fi21, v25)
                                case -1:
                                    if false {
                                        var v34 uint8 = 127
                                        _ = v34
                                    } else {
                                        var v35 int16 = p1
                                        _ = v35
                                    }
                                default:
                                    p1 += (-(p1))
                                }
                                v25 -= (^(v25))
                                var v36 uint16 = 793
                                switch 32768 {
                                case 42:
                                    fmt.Println(v36, v36)
                                case 872:
                                    v37 := GenS4{F0: (v20 / ((168) | 1)), F1: (^(v25)), F2: 1}
                                    _ = v37
                                }
                                v38 := GenS2{F0: v36, F1: 2, F2: false}
                                v25 -= 594
                                v20 += uint64(v22)
                                _ = v36
                                _ = v38
                            }
                            _ = len(v13)
                        }
                        fmt.Println(p0, v20)
                        v22 += (^(v22))
                        _ = v26
                    }
                    _ = v25
                } else {
                    switch fi19 {
                    case 128:
                        if (!false) {
                            var v39 int = 32768
                            _ = v39
                        }
                    case 9223372036854775807:
                        switch "" {
                        case "foo bar":
                            v40 := GenS1{F0: p1}
                            _ = v40
                        default:
                            var v41 int = (fi19 % ((4294967295) | 1))
                            _ = v41
                        }
                    case 1:
                        var v42 uint8 = 128
                        _ = v42
                    default:
                        fmt.Println(v20, p1, v20)
                    }
                }
                if false {
                    v22 = 42
                }
                if false {
                    var v43 bool = (2 > 128)
                    v44 := GenS2{F0: 2, F1: (v22 & (v22 << ((4294967295) & 63))), F2: v43}
                    v45 := GenS1{F0: (p1 % ((-756) | 1))}
                    v46 := GenS5{F0: -286, F1: p1}
                    _ = v43
                    _ = v44
                    _ = v45
                    _ = v46
                } else {
                    v20 = (v20 << ((255) & 63))
                    if len(v13) >= 2 {
                        v13 = v13[1:len(v13)]
                    }
                }
                v20 = v20
                _ = v24
            default:
                var v47 bool = (!false)
                _ = v47
            }
            _ = fi21
            _ = v22
        }
        _ = fi19
        _ = v20
    }
    _ = p0
    _ = p1
    _ = v13
    return 255
}

func main() {
    var v1 string = ("foo bar" + "")
    v2 := []int{128, 170, 9223372036854775807, 9223372036854775807}
    var v3 uint8 = 11
    var v4 uint8 = 0
    v48 := GenS6{F0: 42, F1: 0}
    var v49 bool = (0 <= 1)
    var v50 uint8 = v4
    for fi51 := 0; fi51 < 7; fi51++ {
        if v49 {
            if v49 {
                v49 = (v49 && v49)
            }
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        for fi52 := 0; fi52 < 20; fi52++ {
            var v53 uint16 = 42
            var v54 int16 = 789
            _ = fi52
            _ = v53
            _ = v54
        }
        _ = fi51
    }
    v55 := GenS1{F0: -1}
    v56 := GenS3{F0: (v4 << ((413) & 63)), F1: (3.14159 != float64(1.5))}
    v57 := GenS1{F0: 1}
    for fi58 := 0; fi58 < 8; fi58++ {
        fmt.Println(v49)
        _ = fi58
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if ((0.1 + (0.0 * 1e10)) < 1e10) {
        switch 395 {
        case 9223372036854775807:
            v49 = false
        case 127:
            fmt.Println(v4, v49, v49)
        case 2147483647:
            var v59 uint = 255
            _ = v59
        }
        if v49 {
            v4 += (v50 - 2)
        }
    } else {
        var v60 float64 = 1.5
        _ = v60
    }
    v61 := GenS2{F0: 65535, F1: 0, F2: ((float32(float64(1.5)) * 3.14159) != 3.14159)}
    var v62 int16 = 7
    v3 = (v4 + (v4 | (v50 - v50)))
    v63 := GenS7{F0: (2 == 65535)}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    fmt.Println("v50:", v50)
    fmt.Println("v62:", v62)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v49
    _ = v50
    _ = v55
    _ = v56
    _ = v57
    _ = v61
    _ = v62
    _ = v63
}
