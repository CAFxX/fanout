package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint; F1 uint32; F2 uint8 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 int }

func main() {
    v1 := []uint32{924, 368, 848, 0}
    var v2 int = 32768
    var v3 int16 = 533
    v4 := make(map[uint32]uint32)
    v5 := make(map[uint32]uint32)
    var v6 int16 = 42
    v6 += 738
    switch "\n\t" {
    case "hello":
        fmt.Println(v3, v6, v6)
    case "":
        v7 := GenS1{F0: (false && true)}
        _ = v7
    }
    v8 := GenS1{F0: ((40 < 9223372036854775807) && true)}
    fmt.Println(v6, v3, v6)
    var v9 int8 = -128
    var v10 float64 = (float64((-2.25 - 0.1)) / 0.1)
    if (424 == 1) {
        var v11 uint32 = 65535
        for fi12 := 0; fi12 < 7; fi12++ {
            if len(v1) > 0 {
                v1[0] = (v11 >> ((42) & 63))
            }
            _ = fi12
        }
        if (0 < v2) {
            v3 -= (v3 * v6)
        }
        _ = v11
    } else {
        switch (^(v2)) {
        case 1:
            if ((true || (false && ("foo bar" == ("hello" + "a")))) && false) {
                v13 := GenS1{F0: (!true)}
                for fi14 := 0; fi14 < 20; fi14++ {
                    v15 := GenS1{F0: (true || (!true))}
                    switch 9223372036854775807 {
                    case 128:
                        if ((fi14 & 735) < (v2 | fi14)) {
                            fmt.Println(v6)
                        } else {
                            if true {
                                v16 := GenS1{F0: (950 == 42)}
                                _ = v16
                            } else {
                                fmt.Println(v2, v9)
                                v17 := GenS2{F0: (0 >= 0)}
                                if true {
                                    v2 = (-(v2))
                                    v18 := GenS3{F0: 2, F1: 65535, F2: 128}
                                    v6 -= (v3 % ((v3) | 1))
                                    _ = v18
                                } else {
                                    var v19 uint8 = 42
                                    v20 := GenS4{F0: 2147483648}
                                    if ((430 <= 42) || true) {
                                        if (0 > 42) {
                                            switch (^(fi14)) {
                                            case 42:
                                                var v21 int8 = v9
                                                _ = v21
                                            case -2147483648:
                                                v2 = v2
                                            default:
                                                v2 -= (v2 % ((fi14) | 1))
                                            }
                                            for fi22 := 0; fi22 < 4; fi22++ {
                                                var v23 float32 = -2.25
                                                _ = fi22
                                                _ = v23
                                            }
                                        } else {
                                            for fi24 := 0; fi24 < 11; fi24++ {
                                                fmt.Println(fi14, v10, fi14)
                                                _ = fi24
                                            }
                                            v3 = v3
                                        }
                                        var v25 uint16 = 255
                                        _ = v25
                                    }
                                    _ = v19
                                    _ = v20
                                }
                                _ = v17
                            }
                        }
                    }
                    for fi26 := 0; fi26 < 19; fi26++ {
                        v4[0] = 0
                        var v27 uint64 = 2
                        v28 := GenS2{F0: (1 > 0)}
                        _ = fi26
                        _ = v27
                        _ = v28
                    }
                    _ = fi14
                    _ = v15
                }
                for fi29 := 0; fi29 < 18; fi29++ {
                    for fi30 := 0; fi30 < 13; fi30++ {
                        if false {
                            var v31 bool = false
                            if v31 {
                                _ = len(v1)
                            } else {
                                var v32 uint8 = 0
                                var v33 uint16 = 513
                                fmt.Println(v6, v10, fi29)
                                _ = v32
                                _ = v33
                            }
                            fmt.Println(v2, v31, v3)
                            _ = v31
                        }
                        fmt.Println(fi29, fi30)
                        _ = fi30
                    }
                    if (false && true) {
                        switch fi29 {
                        case 9223372036854775807:
                            v9 = (v9 >> ((0) & 63))
                        }
                        switch "\n\t" {
                        case "foo bar":
                            var v34 int16 = (v6 & v3)
                            _ = v34
                        case "hello":
                            v35 := GenS5{F0: 899}
                            _ = v35
                        default:
                            for fi36 := 0; fi36 < 6; fi36++ {
                                if true {
                                    v2 += v2
                                    var v37 int = -2147483648
                                    _ = v37
                                } else {
                                    var v38 uint64 = 1
                                    if (-2.25 != (3.14159 * ((((float32(float64(3.14159)) - float32((v10 / v10))) / 3.14159) - 1e10) - 1.5))) {
                                        v9 -= (v9 ^ (v9 << ((839) & 63)))
                                        v2 += (fi29 % ((v2) | 1))
                                    }
                                    if false {
                                        _ = len(v5)
                                        var v39 float32 = 1e10
                                        v40 := GenS1{F0: ("foo bar" != ("foo bar" + "\n\t"))}
                                        _ = v39
                                        _ = v40
                                    }
                                    if (!(("hello" + "\n\t") >= ("" + "hello"))) {
                                        var v41 uint = 2
                                        v42 := GenS6{F0: 9223372036854775807}
                                        v3 += v6
                                        _ = v41
                                        _ = v42
                                    } else {
                                        v43 := v5[65535]
                                        var v44 uint8 = 249
                                        _ = v43
                                        _ = v44
                                    }
                                    _ = v38
                                }
                                _ = fi36
                            }
                        }
                        var v45 bool = ((374 == 2147483647) && false)
                        _ = v45
                    }
                    _ = fi29
                }
                _ = v13
            } else {
                var v46 uint64 = 2
                _ = v46
            }
        default:
            for fi47 := 0; fi47 < 16; fi47++ {
                delete(v4, 762)
                _ = fi47
            }
        }
        delete(v5, 128)
    }
    fmt.Println(v2)
    if true {
        switch v2 {
        case 32768:
            delete(v4, 121)
        case 32767:
            for fi48 := 0; fi48 < 18; fi48++ {
                v2 -= fi48
                var v49 uint8 = 0
                fmt.Println(fi48, v9, v2)
                _ = fi48
                _ = v49
            }
        case 935:
            v10 = (float64(3.14159) / (v10 * v10))
        }
        v4[1] = 548
    }
    var v50 uint8 = 127
    if len(v1) > 0 {
        v1[0] = 2147483647
    }
    v2 += v2
    var v51 uint = 18446744073709551615
    v52 := GenS6{F0: v2}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v51
    _ = v52
    _ = v6
    _ = v8
    _ = v9
}
