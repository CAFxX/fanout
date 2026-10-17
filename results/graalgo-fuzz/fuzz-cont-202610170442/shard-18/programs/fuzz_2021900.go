package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int; F2 bool }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint8; F1 bool; F2 uint16 }
func genFn0() int32 {
    v7 := []int{}
    if (false || false) {
        var v8 int8 = 0
        v8 -= v8
        _ = v8
    }
    var v9 int = 128
    _ = v7
    _ = v9
    return 517
}
func genFn1() uint64 {
    if false {
        if false {
            var v10 string = "\n\t"
            _ = v10
        } else {
            for fi11 := 0; fi11 < 18; fi11++ {
                for fi12 := 0; fi12 < 13; fi12++ {
                    fmt.Println(fi11)
                    _ = fi12
                }
                _ = fi11
            }
            if true {
                var v13 uint64 = 211
                v13 = 1
                _ = v13
            }
        }
        if true {
            if false {
                var v14 uint16 = 1
                var v15 int16 = 164
                for fi16 := 0; fi16 < 8; fi16++ {
                    for fi17 := 0; fi17 < 19; fi17++ {
                        v18 := []int{(fi17 + (fi16 * (fi16 / (((fi16 & -9223372036854775808)) | 1)))), (^(fi17)), fi17, 9223372036854775807}
                        v15 = 1000
                        _ = fi17
                        _ = v18
                    }
                    _ = fi16
                }
                _ = v14
                _ = v15
            } else {
                for fi19 := 0; fi19 < 20; fi19++ {
                    v20 := []int{}
                    _ = len(v20)
                    fmt.Println(fi19)
                    _ = fi19
                    _ = v20
                }
            }
            if (175 == 255) {
                for fi21 := 0; fi21 < 7; fi21++ {
                    var v22 uint64 = 42
                    var v23 int32 = -1
                    _ = fi21
                    _ = v22
                    _ = v23
                }
                var v24 int32 = 0
                _ = v24
            }
            var v25 uint8 = 0
            _ = v25
        } else {
            if false {
                for fi26 := 0; fi26 < 1; fi26++ {
                    fmt.Println(fi26)
                    var v27 string = "a"
                    v28 := []int{fi26, fi26, fi26}
                    _ = fi26
                    _ = v27
                    _ = v28
                }
                var v29 uint32 = 443
                var v30 int8 = 1
                _ = v29
                _ = v30
            } else {
                v31 := []int{32768}
                var v32 int8 = -102
                _ = v31
                _ = v32
            }
            switch -9223372036854775808 {
            case 1:
                var v33 uint16 = 2
                _ = v33
            default:
                var v34 uint32 = 2147483648
                _ = v34
            }
        }
        var v35 uint16 = 255
        _ = v35
    }
    var v36 string = ""
    switch v36 {
    case "hello":
        v37 := []int{4294967295, 9223372036854775807, -2147483648, -214}
        _ = v37
    case "foo bar":
        var v38 int = -414
        _ = v38
    default:
        for fi39 := 0; fi39 < 11; fi39++ {
            v36 = ((v36 + "hello") + "foo bar")
            _ = fi39
        }
    }
    _ = v36
    return 42
}

func main() {
    v1 := []int{543}
    var v2 uint32 = 691
    var v3 uint32 = 1
    var v4 uint32 = (v2 / (((v2 << ((255) & 63))) | 1))
    var v5 float32 = 0.0
    v6 := make(map[uint32]int)
    delete(v6, 761)
    v3 -= v4
    for fi40 := 0; fi40 < 8; fi40++ {
        for fi41 := 0; fi41 < 13; fi41++ {
            var v42 float32 = v5
            _ = fi41
            _ = v42
        }
        for fi43 := 0; fi43 < 15; fi43++ {
            v44 := v6[1]
            v4 = v3
            _ = fi43
            _ = v44
        }
        _ = fi40
    }
    if (!true) {
        fmt.Println(v5, v4)
        if false {
            var v45 int8 = 127
            var v46 int16 = 938
            _ = v45
            _ = v46
        }
        var v47 uint64 = 42
        _ = v47
    } else {
        switch "foo bar" {
        case "hello":
            if true {
                v48 := GenS1{F0: "\n\t", F1: -152, F2: (v2 < (v2 << ((42) & 63)))}
                _ = v48
            } else {
                v3 = v4
                switch "\n\t" {
                case "foo bar":
                    v49 := GenS1{F0: (("foo bar" + ("a" + "\n\t")) + "hello"), F1: 32767, F2: (-1 <= 4294967295)}
                    _ = v49
                default:
                    var v50 int16 = 1000
                    _ = v50
                }
                for fi51 := 0; fi51 < 17; fi51++ {
                    v5 += v5
                    _ = fi51
                }
            }
        case "a":
            fmt.Println(v4, v4)
        default:
            _ = len(v6)
        }
        switch 127 {
        case 611:
            var v52 int32 = -1
            _ = v52
        case -832:
            var v53 bool = (42 > 0)
            _ = v53
        }
    }
    var v54 uint8 = 162
    var v55 int8 = 58
    var v56 string = "a"
    var v57 int64 = 42
    var v58 float32 = v5
    v59 := GenS1{F0: ("a" + v56), F1: -433, F2: (((float32(-2.25) - 1.5) < v58) || true)}
    v57 = v57
    if (false && false) {
        v58 -= v58
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        fmt.Println(v2, v55)
    }
    v4 = 65535
    if ((101 <= 965) || (false || false)) {
        switch "" {
        case "":
            var v60 int8 = int8(v54)
            _ = v60
        case "foo bar":
            var v61 int32 = -1
            _ = v61
        }
        for fi62 := 0; fi62 < 3; fi62++ {
            v63 := GenS2{F0: -2147483648}
            _ = fi62
            _ = v63
        }
        var v64 float32 = (0.1 / v5)
        _ = v64
    }
    if true {
        switch ((v56 + v56) + "a") {
        case "foo bar":
            v65 := GenS3{F0: (v54 + (v54 & v54))}
            _ = v65
        case "a":
            if false {
                var v66 int = -2147483648
                for fi67 := 0; fi67 < 7; fi67++ {
                    if (v5 <= 1e10) {
                        if true {
                            fmt.Println(v58)
                        }
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    } else {
                        v2 -= (v2 + (^(v3)))
                        _ = len(v6)
                        if (false || ((!(genFn0() == -925)) && ((!(!true)) && (true || (1 == (v55 - (v55 & v55))))))) {
                            v68 := GenS4{F0: 42, F1: ((v58 + v58) < ((v5 / v58) - v5)), F2: 65535}
                            _ = v68
                        } else {
                            fmt.Println(v66, fi67)
                        }
                    }
                    var v69 float64 = 1.5
                    for fi70 := 0; fi70 < 10; fi70++ {
                        var v71 bool = false
                        v4 = uint32(v54)
                        _ = fi70
                        _ = v71
                    }
                    _ = fi67
                    _ = v69
                }
                _ = v66
            }
        default:
            var v72 int64 = v57
            _ = v72
        }
    }
    for fi73 := 0; fi73 < 8; fi73++ {
        switch fi73 {
        case 128:
            var v74 uint16 = 2
            _ = v74
        case -736:
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        case 1:
            v54 += (v54 | (^(v54)))
        }
        if (uint8(v3) >= 79) {
            fmt.Println(v55)
        } else {
            v3 -= v2
        }
        _ = fi73
    }
    if (true && true) {
        v4 = (^(v4))
    }
    if (false || (!true)) {
        var v75 bool = (!((v56 > v56) || (false && true)))
        var v76 float32 = v58
        _ = v75
        _ = v76
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v54:", v54)
    fmt.Println("v55:", v55)
    fmt.Println("v56:", v56)
    fmt.Println("v57:", v57)
    fmt.Println("v58:", v58)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v54
    _ = v55
    _ = v56
    _ = v57
    _ = v58
    _ = v59
    _ = v6
}
