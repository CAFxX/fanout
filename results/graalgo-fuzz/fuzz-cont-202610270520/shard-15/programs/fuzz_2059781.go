package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint; F2 uint }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int; F1 uint16; F2 uint16 }
type GenS6 struct { F0 uint16 }
func genFn0() uint16 {
    var v7 uint32 = 0
    v8 := GenS1{F0: ("hello" + (("hello" + "\n\t") + "hello")), F1: 186, F2: 33}
    var v9 float32 = (1e10 + 1.5)
    v10 := []int{-368, 128, 32767}
    v7 -= 658
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return 382
}
func genFn1() int {
    if (true && (!(255 != 42))) {
        var v11 uint64 = 18446744073709551615
        v12 := []int{427}
        for fi13 := 0; fi13 < 3; fi13++ {
            var v14 uint16 = 255
            var v15 uint64 = (v11 >> ((65535) & 63))
            _ = fi13
            _ = v14
            _ = v15
        }
        _ = v11
        _ = v12
    }
    var v16 uint32 = 1
    var v17 uint8 = 104
    _ = v16
    _ = v17
    return -9223372036854775808
}

func main() {
    v1 := make(map[int]int)
    var v2 int16 = 1
    v3 := []int{}
    var v4 float32 = -2.25
    var v5 int32 = -1
    var v6 string = "hello"
    fmt.Println(v2)
    switch 1 {
    case -443:
        var v18 int64 = -1
        _ = v18
    case 42:
        if false {
            if true {
                v19 := GenS1{F0: ((v6 + "hello") + v6), F1: 2, F2: 1}
                if len(v3) > 0 {
                    v3[0] = -1
                }
                fmt.Println(v2, v4)
                _ = v19
            } else {
                var v20 int8 = -41
                _ = v20
            }
            if true {
                fmt.Println(v2, v6)
            } else {
                fmt.Println(v5, v4)
                if ((4294967295 > 118) || false) {
                    fmt.Println(v2, v5, v5)
                    _ = len(v1)
                } else {
                    var v21 uint16 = 37
                    if ((v6 + (v6 + (((v6 + v6) + v6) + ((v6 + ("\n\t" + ((("\n\t" + ("hello" + "a")) + "hello") + "hello"))) + v6)))) < v6) {
                        if (!false) {
                            fmt.Println(v6)
                            v22 := GenS2{F0: (v2 << ((18446744073709551615) & 63))}
                            _ = v22
                        } else {
                            v23 := GenS1{F0: (v6 + v6), F1: 65535, F2: 65535}
                            _ = v23
                        }
                        v2 = (v2 & v2)
                        v1[-681] = 747
                    }
                    for fi24 := 0; fi24 < 11; fi24++ {
                        switch (fi24 & (fi24 | (fi24 % (((^(fi24))) | 1)))) {
                        case -507:
                            v21 = (v21 % ((255) | 1))
                        case -380:
                            var v25 int = (fi24 >> ((469) & 63))
                            _ = v25
                        case 0:
                            var v26 int64 = 852
                            _ = v26
                        }
                        v1[(fi24 >> ((42) & 63))] = (fi24 ^ fi24)
                        _ = fi24
                    }
                    _ = v21
                }
                if false {
                    switch -2147483648 {
                    case 9223372036854775807:
                        v27 := GenS2{F0: (v2 << ((42) & 63))}
                        _ = v27
                    }
                    var v28 uint16 = 236
                    for fi29 := 0; fi29 < 20; fi29++ {
                        fmt.Println(v6, v5)
                        var v30 string = ""
                        fmt.Println(v4, v4, v2)
                        _ = fi29
                        _ = v30
                    }
                    _ = v28
                }
            }
        } else {
            var v31 int = 9223372036854775807
            if true {
                for fi32 := 0; fi32 < 3; fi32++ {
                    switch (^(fi32)) {
                    case 42:
                        if (v5 < v5) {
                            switch genFn1() {
                            case -2147483648:
                                var v33 string = v6
                                _ = v33
                            case -816:
                                if true {
                                    v31 += (v31 % (((v31 / ((v31) | 1))) | 1))
                                    var v34 uint = 65535
                                    var v35 uint8 = 255
                                    _ = v34
                                    _ = v35
                                }
                            case 42:
                                var v36 string = "\n\t"
                                _ = v36
                            }
                        }
                    default:
                        _ = len(v1)
                    }
                    if true {
                        v31 = (v31 >> ((18446744073709551615) & 63))
                        for fi37 := 0; fi37 < 8; fi37++ {
                            if (!(!false)) {
                                switch fi37 {
                                case 1:
                                    v38 := v1[-1]
                                    _ = v38
                                case 609:
                                    if false {
                                        v39 := GenS3{F0: 2147483648}
                                        _ = len(v1)
                                        _ = v39
                                    }
                                case 32768:
                                    v40 := GenS4{F0: 301}
                                    _ = v40
                                }
                            }
                            var v41 string = v6
                            _ = fi37
                            _ = v41
                        }
                    }
                    var v42 bool = ((v5 >= (v5 << ((4294967295) & 63))) && true)
                    _ = fi32
                    _ = v42
                }
                v43 := GenS5{F0: 2147483647, F1: 65535, F2: 0}
                _ = v43
            }
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = v31
        }
    case -2147483648:
        for fi44 := 0; fi44 < 20; fi44++ {
            if (0.0 == (1e10 - v4)) {
                if ((1e10 + (((float64(3.14159) * -2.25) + 1.5) * 3.14159)) <= ((0.0 / (1.5 * 1e10)) + 3.14159)) {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    var v45 uint8 = 61
                    _ = v45
                }
            } else {
                var v46 int64 = 941
                _ = v46
            }
            fmt.Println(v4, v6)
            fmt.Println(v2, v2, v2)
            _ = fi44
        }
    }
    var v47 uint8 = 128
    var v48 uint = 2
    if false {
        v2 += (-(v2))
    }
    if len(v3) > 0 {
        v3[0] = -9223372036854775808
    }
    v6 = v6
    switch 128 {
    case 32767:
        v2 = (v2 & 32767)
    case 9223372036854775807:
        if true {
            for fi49 := 0; fi49 < 19; fi49++ {
                switch -361 {
                case 661:
                    v4 -= 3.14159
                case 42:
                    if (true && false) {
                        v50 := GenS6{F0: genFn0()}
                        _ = v50
                    }
                }
                v47 = 2
                for fi51 := 0; fi51 < 5; fi51++ {
                    var v52 int64 = 364
                    var v53 uint32 = 456
                    _ = fi51
                    _ = v52
                    _ = v53
                }
                _ = fi49
            }
            var v54 int16 = (v2 ^ -32768)
            _ = v54
        }
    default:
        for fi55 := 0; fi55 < 7; fi55++ {
            for fi56 := 0; fi56 < 19; fi56++ {
                v5 -= v5
                v2 += 42
                _ = fi56
            }
            _ = fi55
        }
    }
    v47 = (v47 / (((v47 + v47)) | 1))
    if len(v3) > 0 {
        v3[0] = 128
    }
    v2 -= v2
    fmt.Println(v5, v2)
    fmt.Println(v6, v6, v47)
    fmt.Println(v47)
    var v57 float64 = -2.25
    switch -915 {
    case -2147483648:
        var v58 string = "foo bar"
        _ = v58
    case 32768:
        var v59 int32 = v5
        _ = v59
    case 0:
        for fi60 := 0; fi60 < 13; fi60++ {
            v61 := GenS3{F0: 1}
            _ = len(v1)
            _ = fi60
            _ = v61
        }
    default:
        v62 := GenS3{F0: 370}
        _ = v62
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    fmt.Println("v57:", v57)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v48
    _ = v5
    _ = v57
    _ = v6
}
