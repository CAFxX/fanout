package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64 }
type GenS2 struct { F0 int32; F1 int16; F2 uint32 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int64; F1 bool; F2 uint32 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 int8; F1 bool }
func genFn0() uint {
    for fi7 := 0; fi7 < 12; fi7++ {
        switch 558 {
        case 32768:
            v8 := []int{(fi7 | -332), 32768}
            _ = v8
        case -1:
            var v9 bool = false
            _ = v9
        default:
            v10 := GenS1{F0: -1, F1: 42}
            _ = v10
        }
        v11 := []int{(fi7 & 42), fi7, fi7}
        _ = fi7
        _ = v11
    }
    if (1 != 4294967295) {
        var v12 uint8 = 2
        var v13 string = "a"
        _ = v12
        _ = v13
    } else {
        if true {
            var v14 string = "hello"
            var v15 uint8 = 42
            for fi16 := 0; fi16 < 20; fi16++ {
                for fi17 := 0; fi17 < 1; fi17++ {
                    var v18 int32 = int32(v15)
                    v14 = (v14 + v14)
                    var v19 string = v14
                    switch 0 {
                    case -585:
                        v20 := GenS1{F0: (fi16 & (-(fi17))), F1: 9223372036854775807}
                        _ = v20
                    case -340:
                        var v21 uint8 = (v15 << ((1) & 63))
                        _ = v21
                    }
                    _ = fi17
                    _ = v18
                    _ = v19
                }
                for fi22 := 0; fi22 < 5; fi22++ {
                    if true {
                        var v23 bool = false
                        var v24 float64 = 0.0
                        var v25 int = (fi22 * fi22)
                        _ = v23
                        _ = v24
                        _ = v25
                    }
                    var v26 uint8 = (^(v15))
                    _ = fi22
                    _ = v26
                }
                _ = fi16
            }
            _ = v14
            _ = v15
        }
        var v27 string = ""
        _ = v27
    }
    var v28 int8 = -120
    var v29 float32 = 0.1
    v30 := []int{32768, 128}
    _ = v28
    _ = v29
    _ = v30
    return 361
}
func genFn1(p0 int32, p1 uint32) bool {
    p1 += p1
    var v31 uint16 = uint16(p1)
    var v32 uint64 = 2
    for fi33 := 0; fi33 < 6; fi33++ {
        v34 := GenS2{F0: 562, F1: 1000, F2: (p1 & (p1 / ((2147483648) | 1)))}
        v35 := []int{fi33, 1, fi33, fi33}
        _ = fi33
        _ = v34
        _ = v35
    }
    _ = p0
    _ = p1
    _ = v31
    _ = v32
    return true
}

func main() {
    var v1 uint = 502
    v2 := make(map[int]int)
    var v3 int64 = 9223372036854775807
    var v4 float64 = 0.1
    var v5 int16 = 1
    v6 := []uint32{574, 528, 2, 4294967295}
    delete(v2, 4294967295)
    if len(v6) > 0 {
        v6[0] = 65535
    }
    if (127 != 0) {
        var v36 int8 = -128
        if true {
            for fi37 := 0; fi37 < 13; fi37++ {
                switch (fi37 | fi37) {
                case 32767:
                    for fi38 := 0; fi38 < 9; fi38++ {
                        v6 = append(v6, 4294967295)
                        v36 -= int8(fi38)
                        var v39 bool = false
                        _ = fi38
                        _ = v39
                    }
                case 0:
                    _ = len(v2)
                case 4294967295:
                    for fi40 := 0; fi40 < 20; fi40++ {
                        v41 := GenS3{F0: false}
                        var v42 int = (fi37 + (fi40 % (((fi40 % ((1) | 1))) | 1)))
                        _ = fi40
                        _ = v41
                        _ = v42
                    }
                default:
                    v4 -= 0.0
                }
                switch (fi37 << ((649) & 63)) {
                case 173:
                    switch ("foo bar" + "foo bar") {
                    case "foo bar":
                        if true {
                            v43 := GenS3{F0: false}
                            v3 -= v3
                            _ = v43
                        } else {
                            v3 += (v3 | v3)
                        }
                    case "a":
                        v44 := GenS1{F0: (fi37 * 1), F1: v3}
                        _ = v44
                    case "":
                        var v45 uint64 = 2
                        _ = v45
                    }
                case 0:
                    fmt.Println(v5, v3, v1)
                }
                if false {
                    v46 := GenS4{F0: (v3 << ((65535) & 63)), F1: false, F2: 4294967295}
                    _ = len(v2)
                    _ = v46
                } else {
                    v47 := v2[32767]
                    fmt.Println(v5)
                    v2[817] = 646
                    _ = v47
                }
                _ = fi37
            }
        }
        fmt.Println(v36, v5, v36)
        _ = v36
    } else {
        switch "a" {
        case "\n\t":
            if true {
                fmt.Println(v4)
                if (v4 == float64((1e10 + (3.14159 / 1.5)))) {
                    var v48 string = ("\n\t" + "foo bar")
                    _ = v48
                } else {
                    var v49 int = 4294967295
                    v49 = (v49 * (v49 + (v49 | -1)))
                    var v50 uint8 = 0
                    _ = v49
                    _ = v50
                }
            }
        case "hello":
            if (!(1 <= 2)) {
                if true {
                    var v51 string = ""
                    v5 = v5
                    _ = v51
                } else {
                    v52 := GenS4{F0: (v3 % ((9223372036854775807) | 1)), F1: false, F2: 384}
                    _ = v52
                }
                v1 += (v1 / ((636) | 1))
                v4 += v4
            }
        case "foo bar":
            switch "" {
            case "":
                v4 = v4
            default:
                var v53 uint64 = 1
                _ = v53
            }
        default:
            fmt.Println(v1, v1, v1)
        }
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    v5 -= v5
    var v54 uint8 = 2
    if (true || (!(1 <= -616))) {
        v6 = append(v6, 2147483648)
        v55 := GenS4{F0: v3, F1: (!false), F2: 2147483647}
        fmt.Println(v5)
        _ = v55
    } else {
        switch "foo bar" {
        case "\n\t":
            v3 = v3
        case "hello":
            var v56 float64 = 3.14159
            _ = v56
        default:
            v57 := GenS1{F0: 4294967295, F1: v3}
            _ = v57
        }
        delete(v2, -2147483648)
        var v58 int32 = 178
        _ = v58
    }
    var v59 uint32 = 1
    var v60 uint32 = uint32(v1)
    v4 = (float64((((((float32(1e10) / -2.25) - ((float32(1.5) + -2.25) - 0.0)) * float32((v4 + (v4 + (v4 / v4))))) * 0.0) / 1.0)) * float64(0.1))
    v6 = append(v6, (v60 | (v60 / (((v60 + 4294967295)) | 1))))
    v54 = (v54 >> ((v1) & 63))
    var v61 uint64 = 565
    for fi62 := 0; fi62 < 19; fi62++ {
        v3 -= (v3 >> ((v1) & 63))
        var v63 uint8 = v54
        v64 := GenS5{F0: v3}
        _ = fi62
        _ = v63
        _ = v64
    }
    var v65 uint8 = v54
    for fi66 := 0; fi66 < 1; fi66++ {
        fmt.Println(v3, fi66, v5)
        v65 = v65
        fmt.Println(v59)
        _ = fi66
    }
    v67 := GenS6{F0: 90, F1: true}
    v68 := GenS5{F0: 259}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v54:", v54)
    fmt.Println("v59:", v59)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v60:", v60)
    fmt.Println("v61:", v61)
    fmt.Println("v65:", v65)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v54
    _ = v59
    _ = v6
    _ = v60
    _ = v61
    _ = v65
    _ = v67
    _ = v68
}
