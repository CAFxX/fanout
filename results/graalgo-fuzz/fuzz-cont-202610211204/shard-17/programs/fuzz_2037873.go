package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }
type GenS2 struct { F0 uint16; F1 bool; F2 uint32 }
type GenS3 struct { F0 string }
func genFn0(p0 string, p1 int64) int {
    var v5 string = (((p0 + "a") + "a") + "hello")
    p1 = 0
    if true {
        p1 -= p1
    } else {
        var v6 uint = 255
        var v7 float32 = (0.1 / 1.0)
        _ = v6
        _ = v7
    }
    _ = p0
    _ = p1
    _ = v5
    return 128
}

func main() {
    var v1 bool = true
    var v2 uint8 = 2
    var v3 float32 = float32(1e10)
    var v4 int32 = -2147483648
    if false {
        var v8 uint64 = 2
        if v1 {
            v8 = v8
            v9 := []int{}
            for fi10 := 0; fi10 < 20; fi10++ {
                var v11 int16 = -851
                if (!(((false || (v1 && (v1 || true))) || v1) && ((v1 && v1) || v1))) {
                    for fi12 := 0; fi12 < 11; fi12++ {
                        var v13 bool = ((!v1) && v1)
                        _ = fi12
                        _ = v13
                    }
                    if false {
                        var v14 int8 = 42
                        v1 = false
                        _ = v14
                    }
                    v1 = (fi10 < fi10)
                } else {
                    v8 = (v8 ^ v8)
                    if len(v9) > 0 {
                        v9[0] = (-(fi10))
                    }
                    var v15 string = ("a" + "\n\t")
                    _ = v15
                }
                var v16 float32 = v3
                _ = fi10
                _ = v11
                _ = v16
            }
            _ = v9
        } else {
            if false {
                for fi17 := 0; fi17 < 8; fi17++ {
                    v8 = v8
                    _ = fi17
                }
                for fi18 := 0; fi18 < 2; fi18++ {
                    var v19 uint64 = (^(v8))
                    _ = fi18
                    _ = v19
                }
            }
            for fi20 := 0; fi20 < 2; fi20++ {
                var v21 int8 = -104
                _ = fi20
                _ = v21
            }
            if (0 != 127) {
                v22 := GenS1{F0: 255, F1: 42}
                v23 := GenS2{F0: 12, F1: false, F2: 2}
                _ = v22
                _ = v23
            }
        }
        _ = v8
    } else {
        v24 := []int{9223372036854775807, 128}
        var v25 int8 = int8(v4)
        switch ("foo bar" + "foo bar") {
        case "\n\t":
            v24 = append(v24, int(v2))
        case "hello":
            switch 32767 {
            case 427:
                for fi26 := 0; fi26 < 6; fi26++ {
                    for fi27 := 0; fi27 < 8; fi27++ {
                        var v28 uint = 255
                        _ = fi27
                        _ = v28
                    }
                    var v29 float64 = 0.0
                    _ = fi26
                    _ = v29
                }
            case 32767:
                var v30 float64 = (1e10 + (1e10 - (0.0 + float64(((float32(((3.14159 * 3.14159) + 0.0)) * (((float32(1.5) / 0.1) + (v3 * v3)) * (1e10 / 1.5))) - ((0.1 + (v3 * v3)) * v3))))))
                _ = v30
            default:
                v31 := GenS1{F0: 255, F1: 426}
                _ = v31
            }
        case "a":
            if v1 {
                var v32 uint64 = 2
                v33 := GenS2{F0: 32767, F1: v1, F2: 2}
                _ = v32
                _ = v33
            } else {
                fmt.Println(v2, v1, v25)
                switch 9223372036854775807 {
                case -9223372036854775808:
                    if (!v1) {
                        var v34 uint64 = uint64(v2)
                        _ = v34
                    }
                case -934:
                    v3 += float32(3.14159)
                default:
                    for fi35 := 0; fi35 < 19; fi35++ {
                        switch ("\n\t" + "\n\t") {
                        case "":
                            var v36 bool = (4294967295 == 2)
                            _ = v36
                        }
                        if ((995 > 65535) || false) {
                            if false {
                                v25 -= -1
                                _ = len(v24)
                            } else {
                                switch (fi35 + (fi35 % ((fi35) | 1))) {
                                case -539:
                                    switch ("foo bar" + "\n\t") {
                                    case "":
                                        v3 -= v3
                                    case "a":
                                        var v37 int32 = int32(v2)
                                        _ = v37
                                    case "\n\t":
                                        if len(v24) > 0 {
                                            v24[0] = 127
                                        }
                                    default:
                                        if len(v24) >= 2 {
                                            v24 = v24[1:len(v24)]
                                        }
                                    }
                                case 9223372036854775807:
                                    v38 := GenS3{F0: "hello"}
                                    _ = v38
                                case 128:
                                    for fi39 := 0; fi39 < 17; fi39++ {
                                        var v40 uint32 = 65535
                                        switch fi39 {
                                        case 0:
                                            v40 += 2147483647
                                        case 9223372036854775807:
                                            if len(v24) >= 2 {
                                                v24 = v24[1:len(v24)]
                                            }
                                        case -857:
                                            var v41 uint = 1
                                            _ = v41
                                        default:
                                            switch (fi35 & 128) {
                                            case 128:
                                                var v42 uint64 = 470
                                                if (!v1) {
                                                    if v1 {
                                                        _ = len(v24)
                                                    } else {
                                                        switch "\n\t" {
                                                        case "\n\t":
                                                            switch "" {
                                                            case "a":
                                                                v1 = (0 <= v40)
                                                            }
                                                        case "":
                                                            v4 = 1
                                                        case "a":
                                                            var v43 float64 = 0.1
                                                            _ = v43
                                                        }
                                                        switch "\n\t" {
                                                        case "hello":
                                                            v44 := GenS3{F0: (("hello" + "") + "a")}
                                                            _ = v44
                                                        }
                                                        switch (fi35 / (((^(fi35))) | 1)) {
                                                        case 2147483647:
                                                            v45 := GenS1{F0: 0, F1: (v40 / ((v40) | 1))}
                                                            _ = v45
                                                        default:
                                                            fmt.Println(fi35)
                                                        }
                                                    }
                                                }
                                                v2 = v2
                                                _ = v42
                                            case -650:
                                                var v46 int8 = (v25 * 42)
                                                _ = v46
                                            default:
                                                var v47 bool = true
                                                _ = v47
                                            }
                                        }
                                        _ = fi39
                                        _ = v40
                                    }
                                }
                                for fi48 := 0; fi48 < 3; fi48++ {
                                    _ = len(v24)
                                    _ = fi48
                                }
                                var v49 uint32 = 138
                                _ = v49
                            }
                            _ = len(v24)
                            v50 := GenS2{F0: 690, F1: v1, F2: 1}
                            _ = v50
                        }
                        _ = fi35
                    }
                }
            }
        default:
            v51 := GenS3{F0: "a"}
            _ = v51
        }
        _ = v24
        _ = v25
    }
    v3 -= ((((0.0 - v3) + (v3 + 0.0)) + (float32(3.14159) * 1.5)) + v3)
    v1 = ((!v1) || v1)
    var v52 bool = false
    if v52 {
        for fi53 := 0; fi53 < 15; fi53++ {
            var v54 int64 = 801
            v55 := GenS2{F0: 42, F1: ((!(0.0 > v3)) || v52), F2: 65535}
            _ = fi53
            _ = v54
            _ = v55
        }
        fmt.Println(v52, v1)
    } else {
        var v56 uint64 = 2
        var v57 int64 = 1
        _ = v56
        _ = v57
    }
    if false {
        v58 := GenS1{F0: uint16(v2), F1: 2147483647}
        switch 1 {
        case -2147483648:
            v59 := []int{156, -961, -1}
            _ = v59
        }
        _ = v58
    } else {
        var v60 int8 = 100
        v60 += v60
        _ = v60
    }
    var v61 string = "\n\t"
    var v62 int64 = -1
    if v52 {
        for fi63 := 0; fi63 < 11; fi63++ {
            var v64 int = (-(fi63))
            v62 += -9223372036854775808
            _ = fi63
            _ = v64
        }
        if v52 {
            v65 := GenS2{F0: 65535, F1: v1, F2: 65535}
            v4 -= (v4 * -6)
            _ = v65
        }
    } else {
        v66 := []int{32767}
        _ = v66
    }
    for fi67 := 0; fi67 < 19; fi67++ {
        if (42 == 65535) {
            if (!v52) {
                if v1 {
                    v68 := []int{fi67}
                    _ = v68
                }
                for fi69 := 0; fi69 < 15; fi69++ {
                    if v1 {
                        v4 += int32(fi67)
                        v70 := GenS2{F0: 455, F1: false, F2: 635}
                        v71 := GenS2{F0: 32767, F1: v1, F2: 0}
                        _ = v70
                        _ = v71
                    }
                    fmt.Println(v4, v3, fi67)
                    _ = fi69
                }
                v72 := []int{127}
                _ = v72
            }
            for fi73 := 0; fi73 < 18; fi73++ {
                var v74 uint64 = 4294967295
                var v75 uint = 55
                v76 := GenS1{F0: 86, F1: 404}
                _ = fi73
                _ = v74
                _ = v75
                _ = v76
            }
            v1 = v1
        }
        v1 = false
        _ = fi67
    }
    v77 := GenS3{F0: ((v61 + ((((v61 + (("\n\t" + (("foo bar" + v61) + (("a" + (v61 + v61)) + "\n\t"))) + "")) + (v61 + v61)) + ("\n\t" + v61)) + (v61 + v61))) + (v61 + "a"))}
    var v78 int8 = -1
    for fi79 := 0; fi79 < 16; fi79++ {
        var v80 uint32 = 615
        _ = fi79
        _ = v80
    }
    v78 = v78
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v52:", v52)
    fmt.Println("v61:", v61)
    fmt.Println("v62:", v62)
    fmt.Println("v78:", v78)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v52
    _ = v61
    _ = v62
    _ = v77
    _ = v78
}
