package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool; F2 int64 }
type GenS2 struct { F0 uint8; F1 int32 }
type GenS3 struct { F0 int16; F1 uint64; F2 uint64 }
type GenS4 struct { F0 uint8; F1 uint32; F2 bool }
type GenS5 struct { F0 uint32; F1 uint8 }
type GenS6 struct { F0 uint16; F1 int; F2 int8 }
type GenS7 struct { F0 uint32; F1 bool; F2 uint8 }
type GenS8 struct { F0 uint32; F1 int8; F2 int }
func genFn0() uint16 {
    var v6 int = 127
    v7 := GenS1{F0: 0, F1: (18446744073709551615 == 42), F2: 0}
    fmt.Println(v6)
    _ = v6
    _ = v7
    return 42
}
func genFn1() uint8 {
    var v8 int = 1
    v8 += v8
    v9 := GenS2{F0: 255, F1: 619}
    fmt.Println(v8)
    _ = v8
    _ = v9
    return 12
}

func main() {
    var v1 uint32 = 826
    var v2 uint32 = v1
    v3 := make(map[uint32]string)
    v4 := []uint32{(v2 >> ((65535) & 63))}
    var v5 int16 = 978
    v10 := GenS1{F0: v1, F1: (!true), F2: 1}
    if (!(1 >= 42)) {
        var v11 bool = false
        var v12 float32 = ((float32(0.1) * 0.0) * 3.14159)
        _ = v11
        _ = v12
    } else {
        v5 -= (v5 | (v5 & -1))
        v13 := GenS1{F0: 513, F1: (-384 != 34), F2: -1}
        _ = v13
    }
    fmt.Println(v1, v1, v2)
    var v14 int = -80
    var v15 float64 = 1.5
    _ = len(v4)
    if true {
        _ = len(v4)
        if false {
            var v16 int32 = -1
            v15 = v15
            if (-173 > v16) {
                v17 := GenS1{F0: 1, F1: false, F2: 42}
                if false {
                    var v18 float64 = (((((-2.25 / v15) * -2.25) / (v15 * v15)) - 1e10) + -2.25)
                    v14 = v14
                    if (!(false && (-113 == 10))) {
                        var v19 uint32 = (^(v2))
                        v18 = (((v18 * 0.1) / 1.0) / (1.5 / v15))
                        v20 := GenS3{F0: (v5 << ((545) & 63)), F1: 940, F2: 0}
                        _ = v19
                        _ = v20
                    } else {
                        switch (v14 << ((466) & 63)) {
                        case 2147483647:
                            if false {
                                v21 := GenS3{F0: v5, F1: 2, F2: 2}
                                _ = v21
                            } else {
                                if true {
                                    for fi22 := 0; fi22 < 18; fi22++ {
                                        v23 := GenS1{F0: (v2 & 452), F1: false, F2: 410}
                                        v5 = (v5 >> ((61) & 63))
                                        _ = fi22
                                        _ = v23
                                    }
                                    v15 += 3.14159
                                    var v24 bool = false
                                    _ = v24
                                }
                                v14 = (v14 % (((v14 - (v14 - (v14 | (-(v14)))))) | 1))
                                v5 -= v5
                            }
                        case -914:
                            v25 := GenS4{F0: 127, F1: (v1 << ((65535) & 63)), F2: true}
                            _ = v25
                        default:
                            switch ("hello" + "hello") {
                            case "a":
                                if (false && (!((v14 + (v14 % (((v14 - (v14 & 178))) | 1))) >= v14))) {
                                    for fi26 := 0; fi26 < 3; fi26++ {
                                        for fi27 := 0; fi27 < 7; fi27++ {
                                            fmt.Println(v2)
                                            _ = fi27
                                        }
                                        v16 += v16
                                        v28 := GenS4{F0: 0, F1: (^(v2)), F2: (!false)}
                                        _ = fi26
                                        _ = v28
                                    }
                                } else {
                                    v14 -= v14
                                    if true {
                                        switch "foo bar" {
                                        case "hello":
                                            var v29 int32 = (v16 ^ (v16 / (((v16 / ((v16) | 1))) | 1)))
                                            _ = v29
                                        case "foo bar":
                                            if (!true) {
                                                v2 += 2
                                                if true {
                                                    if len(v4) > 0 {
                                                        v4[0] = v2
                                                    }
                                                    for fi30 := 0; fi30 < 19; fi30++ {
                                                        var v31 int64 = -606
                                                        switch (fi30 >> ((255) & 63)) {
                                                        case 9223372036854775807:
                                                            v32 := GenS5{F0: 0, F1: 255}
                                                            _ = v32
                                                        }
                                                        _ = fi30
                                                        _ = v31
                                                    }
                                                    var v33 int8 = -1
                                                    _ = v33
                                                } else {
                                                    if true {
                                                        for fi34 := 0; fi34 < 7; fi34++ {
                                                            _ = len(v3)
                                                            switch (^(fi34)) {
                                                            case -1:
                                                                v4 = append(v4, (v1 >> ((255) & 63)))
                                                            }
                                                            v35 := GenS6{F0: 1, F1: fi34, F2: 100}
                                                            _ = fi34
                                                            _ = v35
                                                        }
                                                        v3[v2] = "\n\t"
                                                        for fi36 := 0; fi36 < 3; fi36++ {
                                                            var v37 uint = 1
                                                            if (true && true) {
                                                                for fi38 := 0; fi38 < 15; fi38++ {
                                                                    var v39 int16 = v5
                                                                    v5 += -32768
                                                                    var v40 int64 = -1
                                                                    if len(v4) >= 2 {
                                                                        v4 = v4[1:len(v4)]
                                                                    }
                                                                    v5 = (-(v39))
                                                                    _ = fi38
                                                                    _ = v39
                                                                    _ = v40
                                                                }
                                                            } else {
                                                                delete(v3, (v2 * v2))
                                                                if (255 == 70) {
                                                                    if (((^(v1)) > 2147483647) && true) {
                                                                        v1 = (v1 / ((v2) | 1))
                                                                    }
                                                                } else {
                                                                    if len(v4) >= 2 {
                                                                        v4 = v4[1:len(v4)]
                                                                    }
                                                                    if (true && ("a" >= "a")) {
                                                                        for fi41 := 0; fi41 < 8; fi41++ {
                                                                            if len(v4) > 0 {
                                                                                v4[0] = (v2 >> ((791) & 63))
                                                                            }
                                                                            _ = fi41
                                                                        }
                                                                    }
                                                                }
                                                                if true {
                                                                    _ = len(v3)
                                                                    v37 = (v37 % (((v37 >> ((2) & 63))) | 1))
                                                                } else {
                                                                    v14 += (fi36 | fi36)
                                                                    var v42 uint32 = v2
                                                                    _ = v42
                                                                }
                                                            }
                                                            _ = fi36
                                                            _ = v37
                                                        }
                                                    } else {
                                                        var v43 uint16 = 42
                                                        _ = v43
                                                    }
                                                }
                                            }
                                        case "\n\t":
                                            var v44 float64 = v18
                                            _ = v44
                                        default:
                                            v2 = 785
                                        }
                                        v18 -= v15
                                    } else {
                                        v45 := GenS3{F0: v5, F1: 1, F2: 42}
                                        var v46 uint64 = 4294967295
                                        _ = v45
                                        _ = v46
                                    }
                                    v47 := GenS7{F0: (v1 + 42), F1: false, F2: 1}
                                    _ = v47
                                }
                            case "":
                                if false {
                                    if len(v4) > 0 {
                                        v4[0] = (v1 / (((v1 / (((^(v2))) | 1))) | 1))
                                    }
                                    var v48 float32 = 3.14159
                                    var v49 int64 = 1
                                    _ = v48
                                    _ = v49
                                } else {
                                    if len(v4) > 0 {
                                        v4[0] = v2
                                    }
                                    v50 := GenS8{F0: 1, F1: 124, F2: (v14 + int(v5))}
                                    _ = v50
                                }
                            case "\n\t":
                                fmt.Println(v5, v18, v15)
                            }
                        }
                    }
                    _ = v18
                }
                _ = v17
            } else {
                v51 := GenS4{F0: 2, F1: (v1 - 927), F2: (!true)}
                v15 -= 1.5
                v15 = v15
                _ = v51
            }
            _ = v16
        }
    }
    v52 := GenS8{F0: v1, F1: 89, F2: 9223372036854775807}
    v53 := GenS2{F0: 1, F1: 2147483647}
    delete(v3, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v52
    _ = v53
}
