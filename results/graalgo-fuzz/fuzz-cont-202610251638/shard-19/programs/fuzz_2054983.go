package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16; F2 int32 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint16; F1 uint16; F2 int16 }
type GenS5 struct { F0 uint32; F1 uint64 }
type GenS6 struct { F0 int32; F1 uint8; F2 int64 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 int }
type GenS9 struct { F0 int64; F1 uint16 }
type GenS10 struct { F0 int64; F1 uint }
type GenS11 struct { F0 uint64 }
type GenS12 struct { F0 uint64 }
func genFn0() uint64 {
    var v7 float32 = (-2.25 * float32((1e10 + 1e10)))
    v7 = 1e10
    _ = v7
    return 18446744073709551615
}
func genFn1() int64 {
    v8 := GenS1{F0: (false || true), F1: 32767, F2: -1}
    var v9 int8 = -32
    _ = v8
    _ = v9
    return 42
}

func main() {
    var v1 int64 = 9223372036854775807
    var v2 int16 = 42
    v3 := []string{"", "a", "\n\t", (("hello" + ("a" + "hello")) + "hello")}
    v4 := make(map[uint32]int)
    v5 := []uint32{}
    var v6 int64 = v1
    if true {
        switch 42 {
        case 42:
            v1 -= v1
        case -2147483648:
            for fi10 := 0; fi10 < 17; fi10++ {
                delete(v4, 118)
                fmt.Println(v2)
                fmt.Println(v6, fi10)
                _ = fi10
            }
        case 1:
            v6 = 42
        }
    } else {
        var v11 int8 = -1
        v11 = v11
        v12 := GenS2{F0: -1}
        _ = v11
        _ = v12
    }
    if (1 == 255) {
        switch -9223372036854775808 {
        case 9223372036854775807:
            v13 := GenS2{F0: v6}
            _ = v13
        default:
            v14 := GenS3{F0: 127}
            _ = v14
        }
    }
    for fi15 := 0; fi15 < 8; fi15++ {
        for fi16 := 0; fi16 < 9; fi16++ {
            var v17 int32 = 1
            if (true && (1 <= 42)) {
                v1 -= v6
            }
            var v18 bool = ((v17 >> ((18446744073709551615) & 63)) < (^(v17)))
            _ = fi16
            _ = v17
            _ = v18
        }
        if true {
            v1 -= genFn1()
            _ = len(v4)
            fmt.Println(v1, v2)
        } else {
            for fi19 := 0; fi19 < 2; fi19++ {
                var v20 float32 = 0.0
                fmt.Println(v20, v20, v2)
                _ = fi19
                _ = v20
            }
            var v21 uint8 = 1
            _ = v21
        }
        if ((-128 == -128) && false) {
            fmt.Println(fi15)
            var v22 uint = 4294967295
            _ = v22
        } else {
            if (!(!(!((true && ((((-(v6)) != 1) || (!(true && (0 == 0)))) || ((0.1 >= -2.25) || false))) || true)))) {
                for fi23 := 0; fi23 < 12; fi23++ {
                    v1 = 993
                    var v24 uint = 65535
                    _ = fi23
                    _ = v24
                }
            } else {
                var v25 int32 = 42
                if true {
                    v26 := GenS4{F0: 454, F1: 1, F2: v2}
                    _ = v26
                } else {
                    fmt.Println(v25, v2)
                }
                _ = v25
            }
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        }
        _ = fi15
    }
    var v27 int = -830
    v1 = v6
    v2 = 32767
    switch 32767 {
    case 1:
        v6 -= int64(v27)
    }
    v6 = 1
    switch v27 {
    case -2147483648:
        var v28 uint = 255
        _ = v28
    case 127:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    case 0:
        var v29 string = "\n\t"
        _ = v29
    default:
        for fi30 := 0; fi30 < 5; fi30++ {
            var v31 bool = true
            _ = fi30
            _ = v31
        }
    }
    if (float64(0.0) < 0.0) {
        if false {
            switch v27 {
            case 32768:
                v27 += (v27 ^ (v27 << ((18446744073709551615) & 63)))
            case 42:
                v5 = append(v5, 1)
            case 0:
                v27 = 127
            }
            v27 -= (v27 >> ((255) & 63))
        }
        v32 := GenS5{F0: 179, F1: 42}
        for fi33 := 0; fi33 < 15; fi33++ {
            v34 := GenS6{F0: -117, F1: 0, F2: 428}
            v6 -= (v6 ^ v1)
            v35 := GenS7{F0: false}
            _ = fi33
            _ = v34
            _ = v35
        }
        _ = v32
    }
    switch v27 {
    case -728:
        if ("\n\t" <= "") {
            var v36 uint = 218
            fmt.Println(v27, v36)
            _ = v36
        } else {
            if true {
                var v37 float32 = 1.5
                for fi38 := 0; fi38 < 1; fi38++ {
                    var v39 uint16 = 927
                    fmt.Println(v37, v2)
                    v6 = (v1 / ((v1) | 1))
                    _ = fi38
                    _ = v39
                }
                _ = v37
            } else {
                var v40 int64 = -1
                _ = v40
            }
            v41 := GenS5{F0: 686, F1: 399}
            v42 := GenS2{F0: (v6 | v1)}
            _ = v41
            _ = v42
        }
    case -2147483648:
        var v43 int64 = (v6 / ((1) | 1))
        _ = v43
    case 32768:
        if (!(1 > 0)) {
            v2 = v2
            var v44 uint64 = 2
            if (!(12 != 127)) {
                if false {
                    v44 += v44
                    if (!(!true)) {
                        if false {
                            v1 += (v6 | (v1 | 1))
                            for fi45 := 0; fi45 < 3; fi45++ {
                                var v46 uint64 = (v44 << ((712) & 63))
                                _ = fi45
                                _ = v46
                            }
                            v44 = (v44 << ((18446744073709551615) & 63))
                        }
                        if false {
                            v1 -= v1
                            v1 = v1
                        } else {
                            if (!false) {
                                v27 -= v27
                            }
                            if true {
                                var v47 uint64 = 4294967295
                                v1 -= (v6 & v1)
                                if true {
                                    var v48 string = ("" + "\n\t")
                                    _ = v48
                                }
                                _ = v47
                            }
                            for fi49 := 0; fi49 < 9; fi49++ {
                                var v50 float64 = (3.14159 - 0.1)
                                fmt.Println(v27)
                                _ = fi49
                                _ = v50
                            }
                        }
                        var v51 int8 = -1
                        _ = v51
                    } else {
                        var v52 int = (v27 / (((v27 << ((585) & 63))) | 1))
                        _ = v52
                    }
                }
                v27 -= v27
                v44 += (v44 | (v44 & v44))
            }
            _ = v44
        }
    default:
        for fi53 := 0; fi53 < 3; fi53++ {
            if (1 > 668) {
                v54 := GenS8{F0: (fi53 | (v27 % ((v27) | 1)))}
                _ = v54
            }
            for fi55 := 0; fi55 < 14; fi55++ {
                if false {
                    fmt.Println(v2, v1, v1)
                }
                v56 := GenS9{F0: (v6 >> ((42) & 63)), F1: 773}
                _ = fi55
                _ = v56
            }
            _ = fi53
        }
    }
    var v57 uint16 = 2
    switch v27 {
    case -18:
        if false {
            fmt.Println(v57)
            if false {
                var v58 int16 = v2
                _ = v58
            }
            delete(v4, 4294967295)
        } else {
            v59 := GenS10{F0: 1, F1: 255}
            _ = v59
        }
    case 127:
        v57 = 42
    default:
        for fi60 := 0; fi60 < 19; fi60++ {
            var v61 int = fi60
            _ = len(v4)
            _ = fi60
            _ = v61
        }
    }
    v62 := GenS11{F0: genFn0()}
    for fi63 := 0; fi63 < 12; fi63++ {
        for fi64 := 0; fi64 < 9; fi64++ {
            var v65 float64 = -2.25
            _ = fi64
            _ = v65
        }
        for fi66 := 0; fi66 < 18; fi66++ {
            if (202 > 407) {
                var v67 uint32 = 614
                _ = len(v3)
                v68 := GenS12{F0: 288}
                _ = v67
                _ = v68
            }
            _ = fi66
        }
        _ = fi63
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v57:", v57)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v57
    _ = v6
    _ = v62
}
