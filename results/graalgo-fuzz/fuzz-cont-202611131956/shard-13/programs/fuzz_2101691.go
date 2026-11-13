package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8; F2 uint32 }
type GenS2 struct { F0 bool; F1 bool }
type GenS3 struct { F0 int; F1 string }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int; F1 uint16; F2 int64 }
func genFn0(p0 int) uint32 {
    if ((float32(float64(0.1)) > 1e10) || false) {
        v6 := []int{p0, p0}
        _ = v6
    }
    var v7 uint16 = 32767
    var v8 int32 = 1
    v9 := GenS1{F0: -128, F1: 255, F2: 880}
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return 735
}

func main() {
    var v1 uint64 = 0
    v2 := make(map[string]uint32)
    var v3 uint = 1
    v4 := []int{}
    var v5 uint32 = 42
    var v10 uint8 = 2
    if ((-9223372036854775808 <= -9223372036854775808) || (42 != v1)) {
        switch 7 {
        case 2147483647:
            _ = len(v4)
        default:
            var v11 bool = false
            _ = v11
        }
        for fi12 := 0; fi12 < 16; fi12++ {
            var v13 float64 = float64(-2.25)
            v14 := GenS1{F0: -1, F1: (v10 << (((v3 | v3)) & 63)), F2: (v5 / (((v5 & uint32(v10))) | 1))}
            v15 := GenS1{F0: 100, F1: 255, F2: v5}
            _ = fi12
            _ = v13
            _ = v14
            _ = v15
        }
        for fi16 := 0; fi16 < 11; fi16++ {
            if false {
                if false {
                    v3 = v3
                    var v17 float32 = 3.14159
                    switch "\n\t" {
                    case "foo bar":
                        if false {
                            v1 = 4294967295
                            v3 += 387
                            v17 += v17
                        }
                    case "a":
                        var v18 int32 = 42
                        _ = v18
                    case "":
                        var v19 int16 = -683
                        _ = v19
                    default:
                        v20 := GenS1{F0: 53, F1: (^(v10)), F2: v5}
                        _ = v20
                    }
                    _ = v17
                } else {
                    v21 := GenS1{F0: 42, F1: (^(v10)), F2: (^(v5))}
                    _ = v21
                }
            }
            if len(v4) > 0 {
                v4[0] = fi16
            }
            _ = fi16
        }
    } else {
        fmt.Println(v10, v3, v10)
        v4 = append(v4, 2147483647)
    }
    v1 -= 0
    if ("\n\t" < "a") {
        switch 42 {
        case 42:
            var v22 int = -479
            _ = v22
        case 4294967295:
            v23 := GenS1{F0: -61, F1: 1, F2: 2147483648}
            _ = v23
        default:
            switch "foo bar" {
            case "foo bar":
                for fi24 := 0; fi24 < 11; fi24++ {
                    if len(v4) > 0 {
                        v4[0] = fi24
                    }
                    if false {
                        v5 += 4294967295
                    } else {
                        for fi25 := 0; fi25 < 11; fi25++ {
                            var v26 int32 = 2147483647
                            v3 += (v3 | (v3 * 255))
                            if (!((!true) && ("hello" >= "\n\t"))) {
                                delete(v2, "hello")
                            } else {
                                var v27 uint16 = 255
                                _ = v27
                            }
                            _ = fi25
                            _ = v26
                        }
                        fmt.Println(v3, v3)
                        v28 := GenS2{F0: false, F1: (1 == int32(fi24))}
                        _ = v28
                    }
                    _ = fi24
                }
            case "hello":
                fmt.Println(v5, v10, v1)
            case "":
                v1 -= (v1 << (((v3 % (((v3 | v3)) | 1))) & 63))
            }
        }
        var v29 uint16 = 42
        v10 = (v10 % ((128) | 1))
        _ = v29
    }
    fmt.Println(v3, v1)
    switch -9223372036854775808 {
    case 32767:
        if (!false) {
            v30 := GenS3{F0: int(v10), F1: "\n\t"}
            _ = v30
        }
    }
    if false {
        for fi31 := 0; fi31 < 5; fi31++ {
            v5 += (v5 & v5)
            if true {
                for fi32 := 0; fi32 < 7; fi32++ {
                    var v33 uint32 = 42
                    if (42 >= fi32) {
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    }
                    v34 := GenS4{F0: true}
                    _ = fi32
                    _ = v33
                    _ = v34
                }
                v35 := GenS5{F0: (fi31 | 32768), F1: 142, F2: 500}
                _ = v35
            }
            for fi36 := 0; fi36 < 19; fi36++ {
                if len(v4) > 0 {
                    v4[0] = 0
                }
                v37 := GenS1{F0: 100, F1: 1, F2: (v5 + (^(v5)))}
                var v38 uint64 = (v1 / ((v1) | 1))
                v5 += (^(v5))
                if ((fi36 * fi36) >= fi31) {
                    fmt.Println(v10)
                    var v39 float64 = 0.1
                    var v40 float32 = 0.1
                    _ = v39
                    _ = v40
                } else {
                    v41 := v2[("foo bar" + "a")]
                    fmt.Println(v1)
                    _ = len(v4)
                    _ = v41
                }
                _ = fi36
                _ = v37
                _ = v38
            }
            _ = fi31
        }
        var v42 int8 = -1
        _ = v42
    }
    v3 += 255
    for fi43 := 0; fi43 < 8; fi43++ {
        if (!false) {
            if ("a" <= ("a" + "\n\t")) {
                v10 -= (^(v10))
            }
            switch "\n\t" {
            case "hello":
                if (0 == -238) {
                    v5 -= 242
                    var v44 uint64 = (v1 + (v1 >> (((v3 | 18446744073709551615)) & 63)))
                    v45 := GenS4{F0: (0.1 == 0.1)}
                    _ = v44
                    _ = v45
                } else {
                    var v46 int64 = 746
                    _ = v46
                }
            case "a":
                v3 = uint(v10)
            case "foo bar":
                for fi47 := 0; fi47 < 17; fi47++ {
                    fmt.Println(fi43, v5, v10)
                    v48 := GenS4{F0: (int8(v5) != 84)}
                    var v49 uint16 = uint16(v10)
                    _ = fi47
                    _ = v48
                    _ = v49
                }
            default:
                if true {
                    _ = len(v2)
                    if (0.1 < -2.25) {
                        if (!(true || ((v10 <= v10) && (-346 > 0)))) {
                            var v50 bool = false
                            _ = v50
                        } else {
                            for fi51 := 0; fi51 < 1; fi51++ {
                                var v52 int = fi43
                                _ = fi51
                                _ = v52
                            }
                        }
                        for fi53 := 0; fi53 < 20; fi53++ {
                            var v54 float64 = 1.5
                            var v55 int16 = -1
                            if false {
                                if false {
                                    var v56 uint64 = v1
                                    v1 = v56
                                    v10 -= v10
                                    _ = v56
                                } else {
                                    fmt.Println(v10, v54)
                                }
                                v57 := v2["hello"]
                                var v58 string = ""
                                _ = v57
                                _ = v58
                            }
                            _ = fi53
                            _ = v54
                            _ = v55
                        }
                    } else {
                        var v59 uint8 = (v10 & v10)
                        var v60 int64 = -9223372036854775808
                        _ = len(v2)
                        _ = v59
                        _ = v60
                    }
                }
            }
        } else {
            var v61 string = ""
            for fi62 := 0; fi62 < 5; fi62++ {
                fmt.Println(fi62, fi62, fi43)
                var v63 bool = true
                _ = len(v4)
                _ = fi62
                _ = v63
            }
            if (false && (!(false && false))) {
                var v64 float64 = -2.25
                switch fi43 {
                case -9223372036854775808:
                    v10 += (^(v10))
                case 32768:
                    for fi65 := 0; fi65 < 14; fi65++ {
                        var v66 uint16 = 652
                        if true {
                            v1 = v1
                            var v67 uint16 = uint16(v5)
                            _ = len(v4)
                            _ = v67
                        }
                        if ((!true) || ((v61 != v61) || false)) {
                            switch fi65 {
                            case 101:
                                var v68 int32 = -2147483648
                                _ = v68
                            case 4294967295:
                                var v69 string = v61
                                _ = v69
                            }
                        } else {
                            v70 := GenS3{F0: 32767, F1: ((v61 + v61) + v61)}
                            _ = v70
                        }
                        _ = fi65
                        _ = v66
                    }
                default:
                    var v71 uint64 = 701
                    _ = v71
                }
                _ = v64
            }
            _ = v61
        }
        var v72 uint32 = v5
        fmt.Println(v72)
        _ = fi43
        _ = v72
    }
    _ = len(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
