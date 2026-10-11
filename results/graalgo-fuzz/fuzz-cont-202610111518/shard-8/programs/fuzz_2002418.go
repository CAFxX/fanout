package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int16; F1 int8 }
type GenS4 struct { F0 uint32; F1 int32; F2 uint16 }
type GenS5 struct { F0 int16; F1 uint; F2 uint8 }
func genFn0(p0 uint16, p1 uint) int {
    var v7 string = "a"
    fmt.Println(v7, p0, v7)
    v8 := []int{int(p0)}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 671
}
func genFn1() uint8 {
    if false {
        if false {
            for fi9 := 0; fi9 < 3; fi9++ {
                v10 := []int{(fi9 - fi9)}
                for fi11 := 0; fi11 < 17; fi11++ {
                    v12 := GenS1{F0: 0, F1: 1}
                    var v13 int32 = 2147483647
                    var v14 int16 = 97
                    if false {
                        _ = len(v10)
                    }
                    _ = fi11
                    _ = v12
                    _ = v13
                    _ = v14
                }
                if (4294967295 == 42) {
                    var v15 int64 = int64(fi9)
                    var v16 int16 = 1
                    var v17 uint = 255
                    _ = v15
                    _ = v16
                    _ = v17
                }
                _ = fi9
                _ = v10
            }
            var v18 float32 = -2.25
            for fi19 := 0; fi19 < 18; fi19++ {
                var v20 uint8 = 127
                switch ("a" + ("\n\t" + ("" + "a"))) {
                case "foo bar":
                    var v21 uint8 = (v20 >> ((18446744073709551615) & 63))
                    _ = v21
                }
                var v22 uint32 = 2147483648
                _ = fi19
                _ = v20
                _ = v22
            }
            _ = v18
        }
    } else {
        var v23 uint8 = 255
        for fi24 := 0; fi24 < 5; fi24++ {
            var v25 bool = true
            var v26 float64 = 1.5
            _ = fi24
            _ = v25
            _ = v26
        }
        _ = v23
    }
    switch 128 {
    case 1:
        var v27 int16 = 32767
        _ = v27
    default:
        var v28 bool = (!((1 >= 128) || false))
        _ = v28
    }
    v29 := []int{2147483647, -2147483648, -775, 805}
    if (-32768 > -32768) {
        v30 := GenS2{F0: genFn0(42, 255)}
        _ = v30
    }
    v31 := GenS2{F0: 32767}
    _ = v29
    _ = v31
    return 2
}

func main() {
    var v1 uint16 = 65535
    var v2 uint8 = 255
    var v3 int8 = 1
    var v4 int32 = 42
    var v5 uint = 42
    v6 := make(map[uint32]uint32)
    v32 := GenS3{F0: 1000, F1: -1}
    v33 := GenS4{F0: 793, F1: (-(v4)), F2: (^(v1))}
    switch 127 {
    case -2147483648:
        var v34 uint16 = (v1 * v1)
        _ = v34
    }
    switch -1 {
    case 128:
        fmt.Println(v4)
    default:
        delete(v6, 586)
    }
    if false {
        if (!(true || false)) {
            if (v3 <= (v3 + v3)) {
                v1 = v1
                for fi35 := 0; fi35 < 6; fi35++ {
                    if (!false) {
                        var v36 string = ("\n\t" + "foo bar")
                        _ = v36
                    } else {
                        v5 -= 0
                        delete(v6, 2147483647)
                        v37 := []int{43, (fi35 + fi35), (fi35 << (((^(v5))) & 63)), (fi35 + (fi35 & (fi35 >> ((v5) & 63))))}
                        _ = v37
                    }
                    v38 := GenS2{F0: int(v2)}
                    _ = fi35
                    _ = v38
                }
                fmt.Println(v5)
            }
        }
    }
    if (!((v2 - 1) < genFn1())) {
        switch -897 {
        case -2147483648:
            if true {
                v3 -= (-(v3))
                v39 := GenS1{F0: 2, F1: v5}
                v5 = (^(v5))
                _ = v39
            } else {
                v4 -= (v4 << (((v5 % (((^(v5))) | 1))) & 63))
                v1 += (^(v1))
            }
        case 341:
            var v40 float64 = 1e10
            _ = v40
        default:
            v3 = (v3 << (((v5 % ((v5) | 1))) & 63))
        }
        var v41 int32 = v4
        _ = v41
    } else {
        if false {
            v42 := GenS3{F0: 1, F1: v3}
            var v43 int8 = v3
            switch "" {
            case "foo bar":
                v44 := GenS5{F0: 1, F1: (^(v5)), F2: v2}
                _ = v44
            case "":
                for fi45 := 0; fi45 < 15; fi45++ {
                    switch "" {
                    case "a":
                        var v46 int = (fi45 - genFn0(v1, v5))
                        _ = v46
                    }
                    var v47 uint16 = (v1 + (v1 | (v1 >> (((v5 / (((v5 | v5)) | 1))) & 63))))
                    fmt.Println(v4)
                    _ = fi45
                    _ = v47
                }
            }
            _ = v42
            _ = v43
        } else {
            v48 := GenS1{F0: 1, F1: (v5 - (v5 >> (((v5 + (v5 | (v5 << (((^(v5))) & 63))))) & 63)))}
            switch -9223372036854775808 {
            case 213:
                if true {
                    v49 := []int{int(v2), 567, 719, 0}
                    var v50 uint = (v5 * (v5 + v5))
                    for fi51 := 0; fi51 < 5; fi51++ {
                        v5 += 255
                        v3 += 0
                        _ = fi51
                    }
                    _ = v49
                    _ = v50
                } else {
                    var v52 int8 = v3
                    var v53 float32 = 0.0
                    var v54 int64 = 0
                    _ = v52
                    _ = v53
                    _ = v54
                }
            case 128:
                _ = len(v6)
            default:
                if false {
                    if true {
                        fmt.Println(v1, v4)
                        v55 := []int{1, 9223372036854775807}
                        _ = v55
                    }
                    v5 += 2
                }
            }
            fmt.Println(v1, v2)
            _ = v48
        }
    }
    for fi56 := 0; fi56 < 11; fi56++ {
        v57 := GenS3{F0: 1000, F1: (^(v3))}
        var v58 uint64 = 4294967295
        _ = fi56
        _ = v57
        _ = v58
    }
    if true {
        v4 = (v4 - v4)
    }
    if false {
        v3 -= (^(v3))
        if ((v3 | v3) != (v3 & 100)) {
            switch ("foo bar" + "foo bar") {
            case "hello":
                var v59 uint16 = v1
                _ = v59
            case "a":
                switch 2147483647 {
                case 42:
                    if true {
                        v60 := GenS2{F0: genFn0((v1 & v1), 255)}
                        for fi61 := 0; fi61 < 5; fi61++ {
                            v3 = v3
                            v1 = (v1 | v1)
                            for fi62 := 0; fi62 < 12; fi62++ {
                                fmt.Println(v1, v3, v3)
                                v2 -= (v2 % (((v2 | genFn1())) | 1))
                                _ = fi62
                            }
                            _ = fi61
                        }
                        _ = v60
                    } else {
                        v1 -= 2
                    }
                case -1:
                    var v63 int8 = v3
                    _ = v63
                }
            default:
                v1 -= (v1 - (v1 ^ (^(v1))))
            }
            if (!true) {
                if true {
                    for fi64 := 0; fi64 < 12; fi64++ {
                        fmt.Println(v2, v1)
                        v65 := GenS4{F0: 388, F1: v4, F2: (v1 / ((v1) | 1))}
                        var v66 int = -1
                        _ = fi64
                        _ = v65
                        _ = v66
                    }
                } else {
                    fmt.Println(v4)
                }
                v6[2147483647] = 912
                v2 = genFn1()
            }
            v1 -= v1
        } else {
            var v67 int16 = -918
            for fi68 := 0; fi68 < 2; fi68++ {
                var v69 int64 = 1
                v67 -= 901
                var v70 bool = (false && false)
                _ = fi68
                _ = v69
                _ = v70
            }
            if (-828 == 1) {
                v1 -= v1
            } else {
                switch "hello" {
                case "":
                    var v71 uint64 = 1
                    _ = v71
                case "a":
                    _ = len(v6)
                }
            }
            _ = v67
        }
    } else {
        var v72 uint8 = (v2 * 152)
        _ = v72
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
